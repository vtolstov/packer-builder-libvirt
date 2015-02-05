package main

import (
	"github.com/mitchellh/packer/packer/plugin"
	"github.com/vtolstov/packer-builder-libvirt/builder/libvirt"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterBuilder(new(libvirt.Builder))
	server.Serve()
}
