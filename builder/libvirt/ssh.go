package libvirt

import (
	"fmt"

	"github.com/mitchellh/multistep"
	commonssh "github.com/mitchellh/packer/common/ssh"
	"github.com/mitchellh/packer/communicator/ssh"
	"github.com/vtolstov/libvirt-go"
	gossh "golang.org/x/crypto/ssh"
)

func commHost(state multistep.StateBag) (string, error) {
	sshHost := state.Get("sshHost").(string)
	return string(sshHost), nil
}

func commPort(state multistep.StateBag) (int, error) {
	sshHostPort := state.Get("sshHostPort").(uint)
	return int(sshHostPort), nil
}

func sshConfig(state multistep.StateBag) (*gossh.ClientConfig, error) {
	config := state.Get("config").(*Config)

	lv, err := libvirt.NewVirConnection(config.LibvirtUrl)
	if err != nil {
		err := fmt.Errorf("Error connecting to libvirt: %s", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return nil, multistep.ActionHalt
	}
	defer lv.CloseConnection()
	if lvd, err := lv.LookupDomainByName(config.VMName); err != nil {
		err := fmt.Errorf("Error lookup domain")
		state.Put("error", err)
		ui.Error(err.Error())
		return nil, multistep.ActionHalt
	}
	defer lvd.Free()

	if config.NetworkType == "nat" {
		ifaces, err := lvd.InterfaceAddresses(0, 0)
		if err != nil {
			err := fmt.Errorf("Error lookup addrs")
			state.Put("error", err)
			ui.Error(err.Error())
			return nil, multistep.ActionHalt
		}
	}
	for _, iface := range ifaces {
		if iface.Name() == "vnet0" {
			state.Put("sshHost", iface.Addrs()[0])
		}
	}

	auth := []gossh.AuthMethod{
		gossh.Password(config.Comm.SSHPassword),
		gossh.KeyboardInteractive(
			ssh.PasswordKeyboardInteractive(config.Comm.SSHPassword)),
	}

	if config.Comm.SSHPrivateKey != "" {
		signer, err := commonssh.FileSigner(config.Comm.SSHPrivateKey)
		if err != nil {
			return nil, err
		}

		auth = append(auth, gossh.PublicKeys(signer))
	}

	return &gossh.ClientConfig{
		User: config.Comm.SSHUsername,
		Auth: auth,
	}, nil
}
