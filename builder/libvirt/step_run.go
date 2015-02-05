package libvirt

import (
	"fmt"
	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
	"time"
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
	config := state.Get("config").(*config)
	ui := state.Get("ui").(packer.Ui)
	xmlPath := state.Get("xml_path").(string)

	ui.Say("Starting the virtual machine...")
	/*
	guiArgument := "gui"
	if config.Headless == true {
		ui.Message("WARNING: The VM will be started in headless mode, as configured.\n" +
			"In headless mode, errors during the boot sequence or OS setup\n" +
			"won't be easily visible. Use at your own discretion.")
		guiArgument = "headless"
	}
	*/
	if _, _, err := virsh("create", xmlPath); err != nil {
		err := fmt.Errorf("Error starting VM: %s", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	s.vmName = config.VMName

	if int64(config.bootWait) > 0 {
		ui.Say(fmt.Sprintf("Waiting %s for boot...", config.bootWait))
		time.Sleep(config.bootWait)
	}

	return multistep.ActionContinue
}

func (s *stepRun) Cleanup(state multistep.StateBag) {
	if s.vmName == "" {
		return
	}

	ui := state.Get("ui").(packer.Ui)

	if running, _ := isRunning(s.vmName); running {
		if _, _, err := virsh("destroy", s.vmName); err != nil {
			ui.Error(fmt.Sprintf("Error shutting down VM: %s", err))
		}
	}
}
