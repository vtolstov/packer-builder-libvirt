package libvirt

import (
	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
	"os"
	"fmt"
	"path/filepath"
)

type stepCreateNetwork struct{}

func (stepCreateNetwork) Run(state multistep.StateBag) multistep.StepAction {
	config := state.Get("config").(*config)
	ui := state.Get("ui").(packer.Ui)

	// TODO: Find IP address & network ...
	// TODO: config.HostIp

	ni := &netInfo{
		VMName:     config.VMName,
		IP:         "172.13.92.1",
		Netmask:    "255.255.255.0",
		RangeStart: "172.13.92.2",
		RangeEnd:   "172.13.92.250",
	}

	state.Put("host_ip", ni.IP)

	netContents, err := config.tpl.Process(netTemplate, ni)
	if err != nil {
		err := fmt.Errorf("Error procesing network template: %s", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	netPath := filepath.Join(config.OutputDir, "net.xml")
	if err := write(netPath, netContents); err != nil {
		err := fmt.Errorf("Error creating network XML file: %s", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	_, _, err = virsh("net-create", netPath)
	if err != nil {
		err := fmt.Errorf("Error creating network: %s", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	return multistep.ActionContinue
}

func (stepCreateNetwork) Cleanup(state multistep.StateBag) {
	config := state.Get("config").(*config)
	ui := state.Get("ui").(packer.Ui)

	_, _, err := virsh("net-destroy", config.VMName)
	if err != nil {
		ui.Error(fmt.Sprintf("Error destroying network: %s", err))
	}
}

func write(path string, data string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.WriteString(data); err != nil {
		return err
	}

	return nil
}

type netInfo struct {
	VMName     string
	IP         string
	Netmask    string
	RangeStart string
	RangeEnd   string
}

const netTemplate = `
<network>
  <name>{{ .VMName }}</name>
  <forward mode="nat"/>
  <ip address="{{ .IP }}" netmask="{{ .Netmask }}">
    <dhcp>
	  <range start="{{ .RangeStart }}" end="{{ .RangeEnd }}" />
	</dhcp>
  </ip>
</network>
`
