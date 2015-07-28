package libvirt

import (
	"fmt"
	"log"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"

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
	var err error
	var key uint

	shiftedChars := "~!@#$%^&*()_+{}|:\"<>?"

	for len(original) > 0 {
		time.Sleep(50 * time.Millisecond)
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
			d.SendKey(libvirt.VIR_KEYCODE_SET_RFB, 50, []uint{ecodes["<esc>"]}, 0)
			original = original[len("<esc>"):]
			continue
		}

		if strings.HasPrefix(original, "<enter>") {
			d.SendKey(libvirt.VIR_KEYCODE_SET_RFB, 50, []uint{ecodes["<enter>"]}, 0)
			original = original[len("<enter>"):]
			continue
		}

		log.Printf("command %s", original)
		r, size := utf8.DecodeRuneInString(original)
		original = original[size:]
		var keys []uint
		if unicode.IsUpper(r) || strings.ContainsRune(shiftedChars, r) {
			keys = append(keys, ecodes["<lshift>"])
		}
		keys = append(keys, ecodes[string(unicode.ToLower(r))])

		log.Printf("find code for char %s %v", string(r), keys)
		//VIR_KEYCODE_SET_LINUX, VIR_KEYCODE_SET_USB, VIR_KEYCODE_SET_RFB, VIR_KEYCODE_SET_WIN32, VIR_KEYCODE_SET_XT_KBD
		if err = d.SendKey(libvirt.VIR_KEYCODE_SET_LINUX, 50, keys, 0); err != nil {
			log.Printf("Sending code %d failed: %s", key, err.Error())
		}
	}

}
