package libvirt

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/common"
	commonssh "github.com/mitchellh/packer/common/ssh"
	"github.com/mitchellh/packer/helper/communicator"
	"github.com/mitchellh/packer/helper/config"
	"github.com/mitchellh/packer/packer"
	"github.com/mitchellh/packer/template/interpolate"
)

const BuilderId = "vtolstov.libvirt"

type Builder struct {
	config Config
	runner multistep.Runner
}

type Config struct {
	common.PackerConfig `mapstructure:",squash"`
	Comm                communicator.Config `mapstructure:",squash"`
	DomainType          string              `mapstructure:"domain_type"`
	BootCommand         []string            `mapstructure:"boot_command"`
	MemorySize          uint                `mapstructure:"memory_size"`
	Arch                string              `mapstructure:"arch"`
	DiskName            string              `mapstructure:"disk_name"`
	DiskType            string              `mapstructure:"disk_type"`
	DiskSize            uint                `mapstructure:"disk_size"`
	FloppyFiles         []string            `mapstructure:"floppy_files"`
	HTTPDir             string              `mapstructure:"http_directory"`
	HTTPPortMin         uint                `mapstructure:"http_port_min"`
	HTTPPortMax         uint                `mapstructure:"http_port_max"`
	ISOUrl              string              `mapstructure:"iso_url"`
	OutputDir           string              `mapstructure:"output_directory"`
	ShutdownCommand     string              `mapstructure:"shutdown_command"`
	SSHHostPortMin      uint                `mapstructure:"ssh_host_port_min"`
	SSHHostPortMax      uint                `mapstructure:"ssh_host_port_max"`
	SSHKeyPath          string              `mapstructure:"ssh_key_path"`
	SSHPassword         string              `mapstructure:"ssh_password"`
	SSHPort             uint                `mapstructure:"ssh_port"`
	SSHUser             string              `mapstructure:"ssh_username"`
	VMName              string              `mapstructure:"vm_name"`

	DomainXml   string `mapstructure:"domain_xml"`
	VolumeXml   string `mapstructure:"volume_xml"`
	PoolName    string `mapstructure:"pool_name"`
	PoolXml     string `mapstructure:"pool_xml"`
	NetworkName string `mapstructure:"network_name"`
	NetworkXml  string `mapstructure:"network_xml"`
	NetworkType string `mapstructure:"network_type"`
	LibvirtUrl  string `mapstructure:"libvirt_url"`

	RawBootWait        string `mapstructure:"boot_wait"`
	RawShutdownTimeout string `mapstructure:"shutdown_timeout"`
	RawSSHWaitTimeout  string `mapstructure:"ssh_wait_timeout"`

	bootWait        time.Duration ``
	shutdownTimeout time.Duration ``
	sshWaitTimeout  time.Duration ``
	ctx             interpolate.Context
}

func (b *Builder) Prepare(raws ...interface{}) ([]string, error) {
	var errs *packer.MultiError
	warnings := make([]string, 0)

	err := config.Decode(&b.config, &config.DecodeOpts{
		Interpolate:        true,
		InterpolateContext: &b.config.ctx,
		InterpolateFilter: &interpolate.RenderFilter{
			Exclude: []string{
				"boot_command",
				"qemuargs",
			},
		},
	}, raws...)
	if err != nil {
		return nil, err
	}

	if b.config.MemorySize < 1 {
		b.config.MemorySize = 512
	}

	if b.config.DomainType == "" {
		b.config.DomainType = "kvm"
	}

	if b.config.DiskType == "" {
		b.config.DiskType = "raw"
	}

	if b.config.DiskSize == 0 {
		b.config.DiskSize = 5000
	}

	if b.config.PoolXml == "" {
		b.config.PoolXml = PackerPool
	}

	if b.config.PoolName == "" {
		b.config.PoolName = "packer"
	}

	if b.config.NetworkName == "" {
		b.config.NetworkName = "packer"
	}

	if b.config.VolumeXml == "" {
		b.config.VolumeXml = PackerVolume
	}

	if b.config.NetworkType == "" {
		b.config.NetworkType = "user"
	}

	if b.config.NetworkType != "user" {
		if b.config.NetworkXml == "" {
			b.config.NetworkXml = PackerNetwork
		}
	}

	if b.config.DomainXml == "" {
		b.config.DomainXml = PackerQemuXML
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

	if b.config.DiskName == "" {
		b.config.DiskName = b.config.VMName
	}

	/*
		// Errors
		templates := map[string]*string{
			"disk_name":        &b.config.DiskName,
			"http_directory":   &b.config.HTTPDir,
			"output_directory": &b.config.OutputDir,
			"shutdown_command": &b.config.ShutdownCommand,
			"ssh_password":     &b.config.SSHPassword,
			"ssh_username":     &b.config.SSHUser,
			"vm_name":          &b.config.VMName,
			"boot_wait":        &b.config.RawBootWait,
			"shutdown_timeout": &b.config.RawShutdownTimeout,
			"ssh_wait_timeout": &b.config.RawSSHWaitTimeout,
			"domain_xml":       &b.config.DomainXml,
			"volume_xml":       &b.config.VolumeXml,
			"pool_name":        &b.config.PoolName,
			"pool_xml":         &b.config.PoolXml,
			"network_name":     &b.config.NetworkName,
			"network_xml":      &b.config.NetworkXml,
		}
	*/
	b.config.ISOUrl, err = common.DownloadableURL(b.config.ISOUrl)
	if err != nil {
		errs = packer.MultiErrorAppend(
			errs, fmt.Errorf("Failed to parse iso_url: %s", err))
	}

	if b.config.HTTPPortMin > b.config.HTTPPortMax {
		errs = packer.MultiErrorAppend(
			errs, errors.New("http_port_min must be less than http_port_max"))
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

	if b.config.SSHUser == "" && !strings.HasPrefix(b.config.LibvirtUrl, "lxc") {
		errs = packer.MultiErrorAppend(
			errs, errors.New("An ssh_username must be specified."))
	}

	b.config.sshWaitTimeout, err = time.ParseDuration(b.config.RawSSHWaitTimeout)
	if err != nil {
		errs = packer.MultiErrorAppend(
			errs, fmt.Errorf("Failed parsing ssh_wait_timeout: %s", err))
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

	steps := []multistep.Step{
		new(stepPrepareOutputDir),
		&common.StepCreateFloppy{
			Files: b.config.FloppyFiles,
		},
		new(stepHTTPServer),
		new(stepCreatePool),
		new(stepCreateNetwork),
		new(stepCreateVolume),
		new(stepCreateDomain),
		new(stepRun),
		new(stepTypeBootCommand),
		&communicator.StepConnect{
			Config:    &b.config.Comm,
			Host:      commHost,
			SSHConfig: sshConfig,
			SSHPort:   commPort,
		},
		// can we upload any guest helpers?
		new(common.StepProvision),
		new(stepShutdown),
		// compact disk?
	}
	/*
	   ▶       ▶       new(stepDownloadVolume),
	   ▶       ▶       new(stepDeleteVolume),

	*/
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
