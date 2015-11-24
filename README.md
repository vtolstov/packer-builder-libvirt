# packer-builder-libvirt

# Installation

* install packer as suggested on packer.io site
* install go as suggested on golang.org site
* install libvirt development packages
* run on console
```
GOBIN="some path where you install packer binary" go get -u github.com/vtolstov/packer-builder-libvirt
```
* create .packerconfig file with
```
{
  "builders": {
    "libvirt": "packer-builder-libvirt"
  }
}
```
* export PACKER_CONFIG="directory with .packerconfig file"
* run packer build
