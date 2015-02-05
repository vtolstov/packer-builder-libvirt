package libvirt

import (
	"fmt"
	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
	"io/ioutil"
	"os"
	"path/filepath"
)

type xmlTemplateData struct {
	DomainType string
	Name       string
	NetName    string
	MemSize    uint
	DiskType   string
	DiskPath   string
	ISOPath    string
}

// This step creates the XML file for the VM.
//
// Uses:
//   config *config
//   iso_path string
//   ui     packer.Ui
//
// Produces:
//   xml_path string - The path to the XML file.
type stepCreateXML struct{}

func (stepCreateXML) Run(state multistep.StateBag) multistep.StepAction {
	config := state.Get("config").(*config)
	isoPath := state.Get("iso_path").(string)
	ui := state.Get("ui").(packer.Ui)

	ui.Say("Building and writing XML file")

	diskPath, err := filepath.Abs(filepath.Join(config.OutputDir,
		config.DiskName+".img"))
	if err != nil {
		err := fmt.Errorf("Error creating XML file: %s", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	tplData := &xmlTemplateData{
		DomainType: config.DomainType,
		Name:       config.VMName,
		NetName:    config.VMName,
		MemSize:    config.MemSize,
		DiskType:   config.DiskType,
		DiskPath:   diskPath,
		ISOPath:    isoPath,
	}

	xmlTemplate := DefaultXMLTemplate
	if config.XMLTemplatePath != "" {
		f, err := os.Open(config.XMLTemplatePath)
		if err != nil {
			err := fmt.Errorf("Error reading XML template: %s", err)
			state.Put("error", err)
			ui.Error(err.Error())
			return multistep.ActionHalt
		}
		defer f.Close()

		rawBytes, err := ioutil.ReadAll(f)
		if err != nil {
			err := fmt.Errorf("Error reading XML template: %s", err)
			state.Put("error", err)
			ui.Error(err.Error())
			return multistep.ActionHalt
		}

		xmlTemplate = string(rawBytes)
	}

	xmlContents, err := config.tpl.Process(xmlTemplate, tplData)
	if err != nil {
		err := fmt.Errorf("Error procesing XML template: %s", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	xmlData := ParseXML(xmlContents)
	/*
	TODO: figure out how to add this stuff back in!

	if config.XMLData != nil {
		log.Println("Setting custom XML data...")
		for k, v := range config.XMLData {
			log.Printf("Setting XML: '%s' = '%s'", k, v)
			xmlData[k] = v
		}
	}

	if floppyPathRaw, ok := state.GetOk("floppy_path"); ok {
		log.Println("Floppy path present, setting in XML")
		xmlData["floppy0.present"] = "TRUE"
		xmlData["floppy0.fileType"] = "file"
		xmlData["floppy0.fileName"] = floppyPathRaw.(string)
	}
	*/

	xmlPath := filepath.Join(config.OutputDir, config.VMName+".xml")
	if err := WriteXML(xmlPath, xmlData); err != nil {
		err := fmt.Errorf("Error creating XML file: %s", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	state.Put("xml_path", xmlPath)

	return multistep.ActionContinue
}

func (stepCreateXML) Cleanup(multistep.StateBag) {
}

// This is the default XML template used if no other template is given.
// This is hardcoded here. If you wish to use a custom template please
// do so by specifying in the builder configuration.
const DefaultXMLTemplate = `
<domain type="{{ .DomainType }}">
  <name>{{ .Name }}</name>
  <memory unit="MiB">{{ .MemSize }}</memory>
  <vcpu placement="static">1</vcpu>
  <os>
    <type arch="x86_64" machine="pc-i440fx-1.4">hvm</type>
    <boot dev="hd"/>
    <boot dev="cdrom"/>
  </os>
  <features>
    <acpi/>
    <apic/>
    <pae/>
  </features>
  <clock offset="utc"/>
  <on_poweroff>destroy</on_poweroff>
  <on_reboot>restart</on_reboot>
  <on_crash>restart</on_crash>
  <devices>
    <disk type="file" device="disk">
      <driver name="qemu" type="{{ .DiskType }}"/>
      <source file="{{ .DiskPath }}"/>
      <target dev="vda" bus="virtio"/>
    </disk>
    <disk type="file" device="cdrom">
      <driver name="qemu" type="raw"/>
      <source file="{{ .ISOPath }}"/>
      <target dev="hdc" bus="ide"/>
      <readonly/>
    </disk>
    <graphics type="vnc" port="-1" autoport="yes"/>
    <video>
      <model type="cirrus" vram="9216" heads="1"/>
    </video>
    <interface type="network">
      <source network="{{ .NetName }}"/>
      <model type="virtio"/>
    </interface>
  </devices>
</domain>
`
