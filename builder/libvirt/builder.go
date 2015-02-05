package libvirt

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/common"
	commonssh "github.com/mitchellh/packer/common/ssh"
	"github.com/mitchellh/packer/packer"
)

const BuilderId = "qur.libvirt"

type Builder struct {
	config config
	runner multistep.Runner
}

type config struct {
	common.PackerConfig `mapstructure:",squash"`

	BootCommand     []string `mapstructure:"boot_command"`
	MemSize         uint     `mapstructure:"mem_size"`
	DomainType      string   `mapstructure:"domain_type"`
	DiskName        string   `mapstructure:"disk_name"`
	DiskType        string   `mapstructure:"disk_type"`
	DiskSize        uint     `mapstructure:"disk_size"`
	FloppyFiles     []string `mapstructure:"floppy_files"`
	Headless        bool     `mapstructure:"headless"`
	HTTPDir         string   `mapstructure:"http_directory"`
	HTTPPortMin     uint     `mapstructure:"http_port_min"`
	HTTPPortMax     uint     `mapstructure:"http_port_max"`
	ISOChecksum     string   `mapstructure:"iso_checksum"`
	ISOChecksumType string   `mapstructure:"iso_checksum_type"`
	ISOUrls         []string `mapstructure:"iso_urls"`
	OutputDir       string   `mapstructure:"output_directory"`
	ShutdownCommand string   `mapstructure:"shutdown_command"`
	SSHHostPortMin  uint     `mapstructure:"ssh_host_port_min"`
	SSHHostPortMax  uint     `mapstructure:"ssh_host_port_max"`
	SSHKeyPath      string   `mapstructure:"ssh_key_path"`
	SSHPassword     string   `mapstructure:"ssh_password"`
	SSHPort         uint     `mapstructure:"ssh_port"`
	SSHUser         string   `mapstructure:"ssh_username"`
	VMName          string   `mapstructure:"vm_name"`
	XMLTemplatePath string   `mapstructure:"xml_template_path"`

	URI string `mapstructure:"uri"`

	RawBootWait        string `mapstructure:"boot_wait"`
	RawSingleISOUrl    string `mapstructure:"iso_url"`
	RawShutdownTimeout string `mapstructure:"shutdown_timeout"`
	RawSSHWaitTimeout  string `mapstructure:"ssh_wait_timeout"`

	bootWait        time.Duration ``
	shutdownTimeout time.Duration ``
	sshWaitTimeout  time.Duration ``
	tpl             *packer.ConfigTemplate
}

func (b *Builder) Prepare(raws ...interface{}) ([]string, error) {
	md, err := common.DecodeConfig(&b.config, raws...)
	if err != nil {
		return nil, err
	}

	b.config.tpl, err = packer.NewConfigTemplate()
	if err != nil {
		return nil, err
	}
	b.config.tpl.UserVars = b.config.PackerUserVars

	// Accumulate any errors
	errs := common.CheckUnusedConfig(md)
	warnings := make([]string, 0)

	if b.config.DomainType == "" {
		b.config.DomainType = "kvm"
	}

	if b.config.DiskType == "" {
		b.config.DiskType = "qcow2"
	}

	if b.config.DiskName == "" {
		b.config.DiskName = "disk"
	}

	if b.config.DiskSize == 0 {
		b.config.DiskSize = 40000
	}

	if b.config.FloppyFiles == nil {
		b.config.FloppyFiles = make([]string, 0)
	}

	if b.config.HTTPPortMin == 0 {
		b.config.HTTPPortMin = 8000
	}

	if b.config.HTTPPortMax == 0 {
		b.config.HTTPPortMax = 9000
	}

	if b.config.OutputDir == "" {
		b.config.OutputDir = fmt.Sprintf("output-%s", b.config.PackerBuildName)
	}

	if b.config.RawBootWait == "" {
		b.config.RawBootWait = "10s"
	}

	if b.config.SSHHostPortMin == 0 {
		b.config.SSHHostPortMin = 2222
	}

	if b.config.SSHHostPortMax == 0 {
		b.config.SSHHostPortMax = 4444
	}

	if b.config.SSHPort == 0 {
		b.config.SSHPort = 22
	}

	if b.config.VMName == "" {
		b.config.VMName = fmt.Sprintf("packer-%s", b.config.PackerBuildName)
	}

	// Errors
	templates := map[string]*string{
		"disk_name":         &b.config.DiskName,
		"http_directory":    &b.config.HTTPDir,
		"iso_checksum":      &b.config.ISOChecksum,
		"iso_checksum_type": &b.config.ISOChecksumType,
		"iso_url":           &b.config.RawSingleISOUrl,
		"output_directory":  &b.config.OutputDir,
		"shutdown_command":  &b.config.ShutdownCommand,
		"ssh_password":      &b.config.SSHPassword,
		"ssh_username":      &b.config.SSHUser,
		"vm_name":           &b.config.VMName,
		"boot_wait":         &b.config.RawBootWait,
		"shutdown_timeout":  &b.config.RawShutdownTimeout,
		"ssh_wait_timeout":  &b.config.RawSSHWaitTimeout,
		"xml_template_path": &b.config.XMLTemplatePath,
	}

	for n, ptr := range templates {
		var err error
		*ptr, err = b.config.tpl.Process(*ptr, nil)
		if err != nil {
			errs = packer.MultiErrorAppend(
				errs, fmt.Errorf("Error processing %s: %s", n, err))
		}
	}

	for i, url := range b.config.ISOUrls {
		var err error
		b.config.ISOUrls[i], err = b.config.tpl.Process(url, nil)
		if err != nil {
			errs = packer.MultiErrorAppend(
				errs, fmt.Errorf("Error processing iso_urls[%d]: %s", i, err))
		}
	}

	validates := map[string]*string{}

	for n, ptr := range validates {
		if err := b.config.tpl.Validate(*ptr); err != nil {
			errs = packer.MultiErrorAppend(
				errs, fmt.Errorf("Error parsing %s: %s", n, err))
		}
	}

	for i, command := range b.config.BootCommand {
		if err := b.config.tpl.Validate(command); err != nil {
			errs = packer.MultiErrorAppend(errs,
				fmt.Errorf("Error processing boot_command[%d]: %s", i, err))
		}
	}

	for i, file := range b.config.FloppyFiles {
		var err error
		b.config.FloppyFiles[i], err = b.config.tpl.Process(file, nil)
		if err != nil {
			errs = packer.MultiErrorAppend(errs,
				fmt.Errorf("Error processing floppy_files[%d]: %s",
					i, err))
		}
	}

	if b.config.HTTPPortMin > b.config.HTTPPortMax {
		errs = packer.MultiErrorAppend(
			errs, errors.New("http_port_min must be less than http_port_max"))
	}

	if b.config.ISOChecksum == "" {
		errs = packer.MultiErrorAppend(
			errs, errors.New("Due to large file sizes, an iso_checksum is required"))
	} else {
		b.config.ISOChecksum = strings.ToLower(b.config.ISOChecksum)
	}

	if b.config.ISOChecksumType == "" {
		errs = packer.MultiErrorAppend(
			errs, errors.New("The iso_checksum_type must be specified."))
	} else {
		b.config.ISOChecksumType = strings.ToLower(b.config.ISOChecksumType)
		if h := common.HashForType(b.config.ISOChecksumType); h == nil {
			errs = packer.MultiErrorAppend(
				errs,
				fmt.Errorf("Unsupported checksum type: %s", b.config.ISOChecksumType))
		}
	}

	if b.config.RawSingleISOUrl == "" && len(b.config.ISOUrls) == 0 {
		errs = packer.MultiErrorAppend(
			errs, errors.New("One of iso_url or iso_urls must be specified."))
	} else if b.config.RawSingleISOUrl != "" && len(b.config.ISOUrls) > 0 {
		errs = packer.MultiErrorAppend(
			errs, errors.New("Only one of iso_url or iso_urls may be specified."))
	} else if b.config.RawSingleISOUrl != "" {
		b.config.ISOUrls = []string{b.config.RawSingleISOUrl}
	}

	for i, url := range b.config.ISOUrls {
		b.config.ISOUrls[i], err = common.DownloadableURL(url)
		if err != nil {
			errs = packer.MultiErrorAppend(
				errs, fmt.Errorf("Failed to parse iso_url %d: %s", i+1, err))
		}
	}

	if !b.config.PackerForce {
		if _, err := os.Stat(b.config.OutputDir); err == nil {
			errs = packer.MultiErrorAppend(
				errs,
				fmt.Errorf("Output directory '%s' already exists. It must not exist.", b.config.OutputDir))
		}
	}

	b.config.bootWait, err = time.ParseDuration(b.config.RawBootWait)
	if err != nil {
		errs = packer.MultiErrorAppend(
			errs, fmt.Errorf("Failed parsing boot_wait: %s", err))
	}

	if b.config.RawShutdownTimeout == "" {
		b.config.RawShutdownTimeout = "5m"
	}

	if b.config.RawSSHWaitTimeout == "" {
		b.config.RawSSHWaitTimeout = "20m"
	}

	b.config.shutdownTimeout, err = time.ParseDuration(b.config.RawShutdownTimeout)
	if err != nil {
		errs = packer.MultiErrorAppend(
			errs, fmt.Errorf("Failed parsing shutdown_timeout: %s", err))
	}

	if b.config.SSHKeyPath != "" {
		if _, err := os.Stat(b.config.SSHKeyPath); err != nil {
			errs = packer.MultiErrorAppend(
				errs, fmt.Errorf("ssh_key_path is invalid: %s", err))
		} else if _, err := commonssh.FileSigner(b.config.SSHKeyPath); err != nil {
			errs = packer.MultiErrorAppend(
				errs, fmt.Errorf("ssh_key_path is invalid: %s", err))
		}
	}

	if b.config.SSHHostPortMin > b.config.SSHHostPortMax {
		errs = packer.MultiErrorAppend(
			errs, errors.New("ssh_host_port_min must be less than ssh_host_port_max"))
	}

	if b.config.SSHUser == "" {
		errs = packer.MultiErrorAppend(
			errs, errors.New("An ssh_username must be specified."))
	}

	b.config.sshWaitTimeout, err = time.ParseDuration(b.config.RawSSHWaitTimeout)
	if err != nil {
		errs = packer.MultiErrorAppend(
			errs, fmt.Errorf("Failed parsing ssh_wait_timeout: %s", err))
	}

	if b.config.XMLTemplatePath != "" {
		if err := b.validateXMLTemplatePath(); err != nil {
			errs = packer.MultiErrorAppend(
				errs, fmt.Errorf("xml_template_path is invalid: %s", err))
		}

	}

	// Warnings
	if b.config.ShutdownCommand == "" {
		warnings = append(warnings,
			"A shutdown_command was not specified. Without a shutdown command, Packer\n"+
				"will forcibly halt the virtual machine, which may result in data loss.")
	}

	if errs != nil && len(errs.Errors) > 0 {
		return warnings, errs
	}

	return warnings, nil
}

func (b *Builder) Run(ui packer.Ui, hook packer.Hook, cache packer.Cache) (packer.Artifact, error) {
	// Check we have the required tools
	err := findTools()
	if err != nil {
		return nil, fmt.Errorf("Failed to find required tools: %s", err)
	}

	steps := []multistep.Step{
		&common.StepDownload{
			Checksum:     b.config.ISOChecksum,
			ChecksumType: b.config.ISOChecksumType,
			Description:  "ISO",
			ResultKey:    "iso_path",
			Url:          b.config.ISOUrls,
		},
		new(stepPrepareOutputDir),
		&common.StepCreateFloppy{
			Files: b.config.FloppyFiles,
		},
		new(stepHTTPServer),
		new(stepCreateDisk),
		new(stepCreateXML),
		new(stepCreateNetwork),
		new(stepRun),
		new(stepTypeBootCommand),
		&common.StepConnectSSH{
			SSHAddress:     sshAddress,
			SSHConfig:      sshConfig,
			SSHWaitTimeout: b.config.sshWaitTimeout,
		},
		// can we upload any guest helpers?
		new(common.StepProvision),
		new(stepShutdown),
		// compact disk?
	}

	// Setup the state bag
	state := &multistep.BasicStateBag{}
	state.Put("cache", cache)
	state.Put("config", &b.config)
	state.Put("hook", hook)
	state.Put("ui", ui)

	// Run
	if b.config.PackerDebug {
		b.runner = &multistep.DebugRunner{
			Steps:   steps,
			PauseFn: common.MultistepDebugFn(ui),
		}
	} else {
		b.runner = &multistep.BasicRunner{Steps: steps}
	}

	b.runner.Run(state)

	// If there was an error, return that
	if rawErr, ok := state.GetOk("error"); ok {
		return nil, rawErr.(error)
	}

	// If we were interrupted or cancelled, then just exit.
	if _, ok := state.GetOk(multistep.StateCancelled); ok {
		return nil, errors.New("Build was cancelled.")
	}

	if _, ok := state.GetOk(multistep.StateHalted); ok {
		return nil, errors.New("Build was halted.")
	}

	// Compile the artifact list
	files := make([]string, 0, 5)
	visit := func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}

		return err
	}

	if err := filepath.Walk(b.config.OutputDir, visit); err != nil {
		return nil, err
	}

	artifact := &Artifact{
		name:  b.config.VMName,
		dir:   b.config.OutputDir,
		f:     files,
		state: make(map[string]interface{}),
	}

	artifact.state["diskType"] = b.config.DiskType
	artifact.state["diskSize"] = b.config.DiskSize
	artifact.state["domainType"] = "kvm"

	return artifact, nil
}

func (b *Builder) Cancel() {
	if b.runner != nil {
		log.Println("Cancelling the step runner...")
		b.runner.Cancel()
	}
}

func (b *Builder) validateXMLTemplatePath() error {
	f, err := os.Open(b.config.XMLTemplatePath)
	if err != nil {
		return err
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	return b.config.tpl.Validate(string(data))
}
