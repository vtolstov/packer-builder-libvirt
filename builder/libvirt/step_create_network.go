package libvirt

import (
	"fmt"

	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
	"github.com/vtolstov/libvirt-go"
)

type stepCreateNetwork struct{}

func (stepCreateNetwork) Run(state multistep.StateBag) multistep.StepAction {
	config := state.Get("config").(*Config)
	ui := state.Get("ui").(packer.Ui)

	if config.NetworkType == "user" {
		return multistep.ActionContinue
	}

	var lvn libvirt.VirNetwork
	lv, err := libvirt.NewVirConnection(config.LibvirtUrl)
	if err != nil {
		err := fmt.Errorf("Error connecting to libvirt: %s", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}
	defer lv.CloseConnection()
	if lvn, err = lv.LookupNetworkByName(config.NetworkName); err != nil {
		lvn, err = lv.NetworkDefineXML(config.NetworkXml)
		if err != nil {
			err := fmt.Errorf("Error defining network: %s", err)
			state.Put("error", err)
			ui.Error(err.Error())
			return multistep.ActionHalt
		}
	}
	defer lvn.Free()
	if ok, err := lvn.IsActive(); !ok && err == nil {
		err = lvn.Create()
		if err != nil {
			err := fmt.Errorf("Error creating network: %s", err)
			state.Put("error", err)
			ui.Error(err.Error())
			return multistep.ActionHalt
		}
	}
	return multistep.ActionContinue
}

func (stepCreateNetwork) Cleanup(state multistep.StateBag) {
	config := state.Get("config").(*Config)
	ui := state.Get("ui").(packer.Ui)

	if config.NetworkType == "user" {
		return
	}

	lv, err := libvirt.NewVirConnection(config.LibvirtUrl)
	if err != nil {
		err := fmt.Errorf("Error connecting to libvirt: %s", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return
	}

	if lvn, err := lv.LookupNetworkByName(config.NetworkName); err == nil {
		defer lvn.Free()
		if ok, err := lvn.IsActive(); !ok && err == nil {
			err = lvn.Destroy()
			if err != nil {
				ui.Error(fmt.Sprintf("Error destroying network: %s", err))
			}
		}
	}
}
