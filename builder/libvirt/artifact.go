package libvirt

import (
	"fmt"
	"os"
)

// Artifact is the result of running the libvirt builder, namely a set
// of files associated with the resulting machine.
type Artifact struct {
	name  string
	dir   string
	f     []string
	state map[string]interface{}
}

func (*Artifact) BuilderId() string {
	return BuilderId
}

func (a *Artifact) Files() []string {
	return a.f
}

func (a *Artifact) Id() string {
	return a.name
}

func (a *Artifact) String() string {
	return fmt.Sprintf("VM files in directory: %s", a.dir)
}

func (a *Artifact) State(name string) interface{} {
	value, _ := a.state[name]
	return value
}

func (a *Artifact) Destroy() error {
	return os.RemoveAll(a.dir)
}
