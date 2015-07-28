package libvirt

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
	"gopkg.in/alexzorin/libvirt-go.v2"
)

const KeyLeftShift uint32 = 0xFFE1

type bootCommandTemplateData struct {
	HTTPIP   string
	HTTPPort uint
	Name     string
}

// This step "types" the boot command into the VM over VNC.
//
// Uses:
//   config *config
//   http_port int
//   ui     packer.Ui
//
// Produces:
//   <nothing>
type stepTypeBootCommand struct{}

func (s *stepTypeBootCommand) Run(state multistep.StateBag) multistep.StepAction {
	config := state.Get("config").(*Config)
	//	httpPort := state.Get("http_port").(uint)
	//	hostIp := state.Get("host_ip").(string)
	ui := state.Get("ui").(packer.Ui)

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
	defer lvd.Free()

	//	tplData := &bootCommandTemplateData{
	//		hostIp,
	//		httpPort,
	//		config.VMName,
	//	}

	ui.Say("Typing the boot command...")
	for _, command := range config.BootCommand {
		//		command, err := config.tpl.Process(command, tplData)
		//		if err != nil {
		//			err := fmt.Errorf("Error preparing boot command: %s", err)
		//			state.Put("error", err)
		//			ui.Error(err.Error())
		//			return multistep.ActionHalt
		//		}

		// Check for interrupts between typing things so we can cancel
		// since this isn't the fastest thing.
		if _, ok := state.GetOk(multistep.StateCancelled); ok {
			return multistep.ActionHalt
		}

		sendBootString(lvd, command)
	}

	return multistep.ActionContinue
}

func (*stepTypeBootCommand) Cleanup(multistep.StateBag) {}

func sendBootString(d libvirt.VirDomain, original string) {
	//	shiftedChars := "~!@#$%^&*()_+{}|:\"<>?"
	var keys []uint
	var key uint
	var ok bool
	var err error

	for len(original) > 0 {
		//		var keyCode uint
		//		keyShift := false

		if strings.HasPrefix(original, "<wait>") {
			log.Printf("Special code '<wait>' found, sleeping one second")
			time.Sleep(1 * time.Second)
			original = original[len("<wait>"):]
			continue
		}

		if strings.HasPrefix(original, "<wait5>") {
			log.Printf("Special code '<wait5>' found, sleeping 5 seconds")
			time.Sleep(5 * time.Second)
			original = original[len("<wait5>"):]
			continue
		}

		if strings.HasPrefix(original, "<wait10>") {
			log.Printf("Special code '<wait10>' found, sleeping 10 seconds")
			time.Sleep(10 * time.Second)
			original = original[len("<wait10>"):]
			continue
		}

		if strings.HasPrefix(original, "<esc>") {
			keys = append(keys, ecodes["<esc>"])
			original = original[len("<esc>"):]
		}
		if strings.HasPrefix(original, "<enter>") {
			keys = append(keys, ecodes["<enter>"])
			original = original[len("<enter>"):]
		}

		char := original[0]
		log.Printf("try to find code for char %s", string(char))
		if key, ok = ecodes[string(char)]; ok {
			log.Printf("find code for char %s %d", string(char), key)
			keys = append(keys, key)
			//			keyShift = unicode.IsUpper(r) || strings.ContainsRune(shiftedChars, r)
		}
	}
	//VIR_KEYCODE_SET_LINUX, VIR_KEYCODE_SET_USB, VIR_KEYCODE_SET_RFB, VIR_KEYCODE_SET_WIN32, VIR_KEYCODE_SET_XT_KBD
	for _, key := range keys {
		log.Printf("send code %d", key)
		if err = d.SendKey(libvirt.VIR_KEYCODE_SET_RFB, 1000, []uint{key}, 0); err != nil {
			log.Printf("Sending code %d failed: %s", key, err.Error())
		}
	}

}
