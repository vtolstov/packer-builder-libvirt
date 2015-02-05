package main

import (
	"github.com/mitchellh/packer/packer/plugin"
	"github.com/vtolstov/packer-builder-lxc/builder/lxc"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterBuilder(new(lxc.Builder))
	server.Serve()
}
