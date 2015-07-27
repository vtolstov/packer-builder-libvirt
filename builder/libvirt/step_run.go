package libvirt

import (
	"fmt"
	"time"

	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
	"gopkg.in/alexzorin/libvirt-go.v2"
)

// This step starts the virtual machine.
//
// Uses:
//
// Produces:
type stepRun struct {
	vmName string
}

func (s *stepRun) Run(state multistep.StateBag) multistep.StepAction {
	config := state.Get("config").(*Config)
	ui := state.Get("ui").(packer.Ui)

	ui.Say("Starting the virtual machine...")
	var lvd libvirt.VirDomain
	lv, err := libvirt.NewVirConnection(config.LibvirtUrl)
	if err != nil {
		err := fmt.Errorf("Error connecting to libvirt: %s", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}
	defer lv.CloseConnection()
	if lvd, err = lv.LookupDomainByName(config.VMName); err != nil {
		err := fmt.Errorf("Error lookup domain: %s", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}
	if err = lvd.Create(); err != nil {
		err := fmt.Errorf("Error creating domain: %s", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	defer lvd.Free()

	s.vmName = config.VMName

	if int64(config.bootWait) > 0 {
		ui.Say(fmt.Sprintf("Waiting %s for boot...", config.bootWait))
		time.Sleep(config.bootWait)
	}

	return multistep.ActionContinue
}

func (s *stepRun) Cleanup(state multistep.StateBag) {
	config := state.Get("config").(*Config)
	if s.vmName == "" {
		return
	}
	var lvd libvirt.VirDomain
	ui := state.Get("ui").(packer.Ui)

	lv, err := libvirt.NewVirConnection(config.LibvirtUrl)
	if err != nil {
		ui.Error(fmt.Sprintf("Error connecting to libvirt: %s", err))
		return
	}
	defer lv.CloseConnection()
	if lvd, err = lv.LookupDomainByName(s.vmName); err != nil {
		ui.Error(fmt.Sprintf("Error creating domain: %s", err))
		return
	}
	defer lvd.Free()
	if ok, err := lvd.IsActive(); ok && err == nil {
		if err = lvd.Destroy(); err != nil {
			ui.Error(fmt.Sprintf("Error shutting down domain: %s", err))
		}
	}

	if err = lvd.Undefine(); err != nil {
		ui.Error(fmt.Sprintf("Error undefine domain: %s", err))
	}
	return
}
