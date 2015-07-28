package libvirt

import (
	"bytes"
	"fmt"
	"html/template"
	"net"
	"net/url"
	"strings"

	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
	"gopkg.in/alexzorin/libvirt-go.v2"
)

type stepCreateDomain struct{}

func (s *stepCreateDomain) Run(state multistep.StateBag) multistep.StepAction {
	config := state.Get("config").(*Config)
	ui := state.Get("ui").(packer.Ui)

	lv, err := libvirt.NewVirConnection(config.LibvirtUrl)
	if err != nil {
		err := fmt.Errorf("Error connecting to libvirt: %s", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}
	defer lv.CloseConnection()
	if lvd, err := lv.LookupDomainByName(config.VMName); err == nil {
		err = lvd.Destroy()
		if err != nil {
			err := fmt.Errorf("Error domain already running")
			state.Put("error", err)
			ui.Error(err.Error())
			return multistep.ActionHalt
		}
	}

	domainXml := bytes.NewBuffer(nil)
	tmpl, err := template.New("domain").Parse(config.DomainXml)
	if err != nil {
		err := fmt.Errorf("Error creating domain template: %s", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	u, err := url.Parse(config.ISOUrl)
	if err != nil {
		err := fmt.Errorf("Error parse iso_url: %s", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}
	var h string
	var p string

	if strings.Index(u.Host, ":") == -1 {
		h = u.Host
	} else {
		h, p, _ = net.SplitHostPort(u.Host)
	}
	if p == "" {
		switch u.Scheme {
		case "https":
			p = "443"
		case "http":
			p = "80"
		}
	}

	data := struct {
		VMName      string
		DiskName    string
		DiskType    string
		PoolName    string
		MemorySize  uint
		ISOUrlProto string
		ISOUrlPath  string
		ISOUrlHost  string
		ISOUrlPort  string
		SSHPort     string
	}{
		config.VMName,
		config.DiskName,
		"raw",
		config.PoolName,
		config.MemorySize,
		u.Scheme,
		u.Path,
		h,
		p,
		"2022",
	}
	err = tmpl.Execute(domainXml, data)
	if err != nil {
		err := fmt.Errorf("Error running domain template: %s", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	//	ui.Say(fmt.Sprintf("domain config: %s", string(domainXml.Bytes())))

	lvd, err := lv.DomainDefineXML(string(domainXml.Bytes()))
	if err != nil {
		err := fmt.Errorf("Error defining domain: %s", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	defer lvd.Free()

	return multistep.ActionContinue
}

func (s *stepCreateDomain) Cleanup(state multistep.StateBag) {

}
