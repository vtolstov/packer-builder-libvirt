package libvirt

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
	"github.com/vtolstov/libvirt-go"
)

// This step shuts down the machine. It first attempts to do so gracefully,
// but ultimately forcefully shuts it down if that fails.
//
// Uses:
//   communicator packer.Communicator
//   config *config
//   ui     packer.Ui
//
// Produces:
//   <nothing>
type stepShutdown struct{}

func (s *stepShutdown) Run(state multistep.StateBag) multistep.StepAction {
	comm := state.Get("communicator").(packer.Communicator)
	config := state.Get("config").(*Config)
	ui := state.Get("ui").(packer.Ui)
	var lvd libvirt.VirDomain
	if config.ShutdownCommand != "" {
		ui.Say("Gracefully halting virtual machine...")
		log.Printf("Executing shutdown command: %s", config.ShutdownCommand)
		cmd := &packer.RemoteCmd{Command: config.ShutdownCommand}
		if err := cmd.StartWithUi(comm, ui); err != nil {
			err := fmt.Errorf("Failed to send shutdown command: %s", err)
			state.Put("error", err)
			ui.Error(err.Error())
			return multistep.ActionHalt
		}

		// Wait for the machine to actually shut down
		log.Printf("Waiting max %s for shutdown to complete", config.shutdownTimeout)
		shutdownTimer := time.After(config.shutdownTimeout)
		for {
			//			running, _ := isRunning(config.VMName)
			//			if !running {
			//				break
			//			}

			select {
			case <-shutdownTimer:
				err := errors.New("Timeout while waiting for machine to shut down.")
				state.Put("error", err)
				ui.Error(err.Error())
				return multistep.ActionHalt
			default:
				time.Sleep(1 * time.Second)
			}
		}
	} else {
		ui.Say("Halting the virtual machine...")
		lv, err := libvirt.NewVirConnection(config.LibvirtUrl)
		if err != nil {
			err := fmt.Errorf("Error connecting to libvirt: %s", err)
			state.Put("error", err)
			ui.Error(err.Error())
			return multistep.ActionHalt
		}
		defer lv.CloseConnection()
		if lvd, err = lv.LookupDomainByName(config.VMName); err != nil {
			err := fmt.Errorf("Error creating domain: %s", err)
			state.Put("error", err)
			ui.Error(err.Error())
			return multistep.ActionHalt
		}
		defer lvd.Free()
		if ok, err := lvd.IsActive(); ok && err == nil {
			if err = lvd.Destroy(); err != nil {
				err := fmt.Errorf("Error shut down domain: %s", err)
				state.Put("error", err)
				ui.Error(err.Error())
				return multistep.ActionHalt

			}
		}
		if err = lvd.Undefine(); err != nil {
			err := fmt.Errorf("Error undefine domain: %s", err)
			state.Put("error", err)
			ui.Error(err.Error())
			return multistep.ActionHalt

		}
	}

	log.Println("VM shut down.")
	return multistep.ActionContinue
}

func (s *stepShutdown) Cleanup(state multistep.StateBag) {}
