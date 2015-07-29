package libvirt

import (
	"fmt"

	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
	"github.com/vtolstov/libvirt-go"
)

type stepCreatePool struct{}

func (stepCreatePool) Run(state multistep.StateBag) multistep.StepAction {
	config := state.Get("config").(*Config)
	ui := state.Get("ui").(packer.Ui)
	var lvp libvirt.VirStoragePool
	lv, err := libvirt.NewVirConnection(config.LibvirtUrl)
	if err != nil {
		err := fmt.Errorf("Error connecting to libvirt: %s", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}
	defer lv.CloseConnection()
	if lvp, err = lv.LookupStoragePoolByName(config.PoolName); err != nil {
		lvp, err = lv.StoragePoolDefineXML(config.PoolXml, 0)
		if err != nil {
			err := fmt.Errorf("Error defining pool: %s", err)
			state.Put("error", err)
			ui.Error(err.Error())
			return multistep.ActionHalt
		}
	}
	defer lvp.Free()
	if ok, err := lvp.IsActive(); !ok && err == nil {
		err = lvp.Build(0)
		if err != nil {
			err := fmt.Errorf("Error building pool: %s", err)
			state.Put("error", err)
			ui.Error(err.Error())
			return multistep.ActionHalt
		}
		err = lvp.Create(0)
		if err != nil {
			err := fmt.Errorf("Error creating network: %s", err)
			state.Put("error", err)
			ui.Error(err.Error())
			return multistep.ActionHalt
		}
	}
	return multistep.ActionContinue
}

func (stepCreatePool) Cleanup(state multistep.StateBag) {
	config := state.Get("config").(*Config)
	ui := state.Get("ui").(packer.Ui)
	var lvp libvirt.VirStoragePool
	lv, err := libvirt.NewVirConnection(config.LibvirtUrl)
	if err != nil {
		err := fmt.Errorf("Error connecting to libvirt: %s", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return
	}

	if lvp, err = lv.LookupStoragePoolByName(config.PoolName); err == nil {
		defer lvp.Free()
		if ok, err := lvp.IsActive(); !ok && err == nil {
			err = lvp.Destroy()
			if err != nil {
				ui.Error(fmt.Sprintf("Error destroying network: %s", err))
			}
		}
	}
}
