package libvirt

import (
	"fmt"

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
	if lvd, err := lv.LookupDomainByName(config.VMName); err != nil {
		lvd, err = lv.DomainCreateXML(config.DomainXml, 0)
		if err != nil {
			err := fmt.Errorf("Error creating domain: %s", err)
			state.Put("error", err)
			ui.Error(err.Error())
			return multistep.ActionHalt
		}
	} else {
		defer lvd.Free()
	}
	return multistep.ActionContinue
}

func (s *stepCreateDomain) Cleanup(state multistep.StateBag) {

}
