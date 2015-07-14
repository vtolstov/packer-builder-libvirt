package libvirt

import (
	"bytes"
	"testing"
	"text/template"

	"github.com/alexzorin/libvirt-go"
	"github.com/mitchellh/packer/packer"
)

func TestBuilder_ImplementsBuilder(t *testing.T) {
	var raw interface{}
	raw = &Builder{}
	if _, ok := raw.(packer.Builder); !ok {
		t.Error("Builder must implement builder.")
	}
}

func TestBuilder_Network(t *testing.T) {
	var lvn libvirt.VirNetwork
	var err error

	lv := getConn(t)
	defer lv.UnrefAndCloseConnection()

	if lvn, err = lv.LookupNetworkByName("packer"); err != nil {
		lvn, err = lv.NetworkDefineXML(PackerNetwork)
		if err != nil {
			t.Fatalf("Error define network: %s", err)
		}
	}
	defer lvn.Free()

	if ok, err := lvn.IsActive(); !ok && err == nil {
		err = lvn.Create()
		if err != nil {
			t.Fatalf("Error create network: %s", err)
		}
	}
	err = lvn.Destroy()
	if err != nil {
		t.Fatalf("Error destroy network: %s", err)
	}
	err = lvn.Undefine()
	if err != nil {
		t.Fatalf("Error undefine network: %s", err)
	}
}

func getConn(t *testing.T) libvirt.VirConnection {
	lv, err := libvirt.NewVirConnection("qemu:///system")
	if err != nil {
		t.Fatalf("Error connect to libvirt: %s", err)
	}
	return lv
}

func getPool(lv libvirt.VirConnection, t *testing.T) libvirt.VirStoragePool {
	var lvp libvirt.VirStoragePool
	var err error
	if lvp, err = lv.LookupStoragePoolByName("packer"); err != nil {
		lvp, err = lv.StoragePoolDefineXML(PackerPool, 0)
		if err != nil {
			t.Fatalf("Error define pool: %s", err)
		}
	}
	if ok, err := lvp.IsActive(); !ok && err == nil {
		err = lvp.Build(0)
		if err != nil {
			t.Fatalf("Error build pool: %s", err)
		}
		err = lvp.Create(0)
		if err != nil {
			t.Fatalf("Error create pool: %s", err)
		}
	}
	return lvp
}

func TestBuilder_Pool(t *testing.T) {
	lv := getConn(t)
	lvp := getPool(lv, t)
	var err error

	defer lv.UnrefAndCloseConnection()
	defer lvp.Free()

	err = lvp.Destroy()
	if err != nil {
		t.Fatalf("Error destroy pool: %s", err)
	}
	err = lvp.Undefine()
	if err != nil {
		t.Fatalf("Error undefine pool: %s", err)
	}
}

func TestBuilder_Volume(t *testing.T) {
	var err error
	var lvv libvirt.VirStorageVol
	lv := getConn(t)
	lvp := getPool(lv, t)

	defer lv.UnrefAndCloseConnection()
	defer lvp.Free()

	var tmp bytes.Buffer
	tmpl := template.New("packer")
	tmpl, err = tmpl.Parse(PackerVolume)
	if err != nil {
		t.Fatalf("Error parsing volume xml: %s", err)
	}
	tmpl.Execute(&tmp, struct {
		DiskSize int
		DiskName string
	}{1, "test"})
	volumeXml := tmp.String()
	if err != nil {
		t.Fatalf("Error execute volume xml: %s", err)
	}
	if lvv, err = lvp.StorageVolCreateXML(volumeXml, 0); err != nil {
		t.Fatalf("Error creating volume: %s", err)
	}
	err = lvv.Delete(0)
	if err != nil {
		t.Fatalf("Error delete volume: %s", err)
	}
	defer lvv.Free()
}
