package libvirt

import (
	"os"
)

func ParseXML(raw string) string {
	return raw
}

func WriteXML(path string, data string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.Write([]byte(data)); err != nil {
		return err
	}

	return nil
}
