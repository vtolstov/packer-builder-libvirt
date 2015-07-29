package libvirt

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
	"github.com/vtolstov/libvirt-go"
)

type stepCreateVolume struct{}

func (s *stepCreateVolume) Run(state multistep.StateBag) multistep.StepAction {
	config := state.Get("config").(*Config)
	ui := state.Get("ui").(packer.Ui)
	var lvp libvirt.VirStoragePool
	var lvv libvirt.VirStorageVol
	lv, err := libvirt.NewVirConnection(config.LibvirtUrl)
	if err != nil {
		err := fmt.Errorf("Error connecting to libvirt: %s", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}
	defer lv.CloseConnection()
	if lvp, err = lv.LookupStoragePoolByName(config.PoolName); err != nil {
		err := fmt.Errorf("Error getting pool %s: %s", config.PoolName, err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	volumeXml := bytes.NewBuffer(nil)
	tmpl, err := template.New("volume").Parse(config.VolumeXml)
	if err != nil {
		err := fmt.Errorf("Error creating volume: %s", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}
	data := struct {
		DiskName string
		DiskSize uint
	}{
		config.DiskName,
		config.DiskSize,
	}
	err = tmpl.Execute(volumeXml, data)
	if err != nil {
		err := fmt.Errorf("Error creating volume: %s", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	if config.PackerForce {
		if lvv, err = lvp.LookupStorageVolByName(config.DiskName); err == nil {
			if err = lvv.Delete(0); err != nil {
				err := fmt.Errorf("Error creating volume: %s", err)
				ui.Error(err.Error())
				return multistep.ActionHalt
			}
		}
	}

	if _, err := lvp.StorageVolCreateXML(string(volumeXml.Bytes()), 0); err != nil {
		err := fmt.Errorf("Error creating volume: %s", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	return multistep.ActionContinue
}

func (s *stepCreateVolume) Cleanup(state multistep.StateBag) {
	config := state.Get("config").(*Config)
	ui := state.Get("ui").(packer.Ui)
	var lvp libvirt.VirStoragePool
	var lvv libvirt.VirStorageVol

	lv, err := libvirt.NewVirConnection(config.LibvirtUrl)
	if err != nil {
		err := fmt.Errorf("Error connecting to libvirt: %s", err)
		ui.Error(err.Error())
		return
	}
	defer lv.CloseConnection()
	if lvp, err = lv.LookupStoragePoolByName(config.PoolName); err != nil {
		err := fmt.Errorf("Error getting pool %s: %s", config.PoolName, err)
		ui.Error(err.Error())
		return
	}

	if lvv, err = lvp.LookupStorageVolByName(config.DiskName); err != nil {
		err := fmt.Errorf("Error creating volume: %s", err)
		ui.Error(err.Error())
		return
	}

	defer lvv.Delete(0)
}
