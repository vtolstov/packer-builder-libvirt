package libvirt

const PackerPool = `
<pool type="dir">
  <name>{{.PoolName}}</name>
  <target>
    <path>/var/lib/libvirt/images</path>
  </target>
</pool>
`

const PackerVolume = `
<volume>
  <name>{{.DiskName}}</name>
  <allocation>0</allocation>
  <capacity unit="M">{{.DiskSize}}</capacity>
  <target>
    <path>{{.DiskName}}</path>
  </target>
</volume>
`
const PackerNetwork = `
<network>
  <name>packer</name>
  <forward mode='nat'/>
  <bridge name='packer0' stp='on' delay='0'/>
  <mac address='52:54:00:2e:85:ab'/>
  <ip address='10.0.2.1' netmask='255.255.255.0'>
    <dhcp>
      <range start='10.0.2.2' end='10.0.2.254'/>
    </dhcp>
  </ip>
</network>
`

const PackerQemuXML = `
<domain type='kvm'>
  <name>{{.VMName}}</name>

  <memory unit='M'>{{.MemorySize}}</memory>
  <type arch='x86_64' machine='pc-i440fx-1.5'>hvm</type>

  <features>
    <acpi/>
    <apic/>
    <pae/>
    <viridian/>
  </features>

  <clock offset='utc'/>
  <on_poweroff>destroy</on_poweroff>
  <on_reboot>restart</on_reboot>
  <on_crash>destroy</on_crash>

  <devices>
  <emulator>/usr/bin/qemu-system-x86_64</emulator>
    <disk type='volume' device='disk'>
      <driver name='qemu' type='{{.DiskType}}' cache='none' io='native' discard='unmap'/>
      <source pool='{{.PoolName}}' volume='{{.DiskName}}'/>
      <alias name='scsi-disk0'/>
      <target dev='sda' bus='scsi'/>
      <boot order='1'/>
      <address type='drive' controller='0' bus='0' unit='0'/>
    </disk>

    <disk type='network' device='cdrom'>
      <driver name='qemu' type='raw'/>
      <target dev='sdb' bus='scsi'/>
      <source protocol="{{.ISOUrlProto}}" name="{{.ISOUrlPath}}">
        <host name="{{.ISOUrlHost}}" port="{{.ISOUrlPort}}"/>
      </source>
      <boot order='2'/>
      <address type='drive' controller='0' target='1' bus='0' unit='0'/>
      <readonly/>
    </disk>

    <controller type='scsi' model='virtio-scsi' index='0'>
      <alias name='scsi0'/>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x04' function='0x0'/>
    </controller>

    <controller type='usb' index='0'>
      <alias name='usb0'/>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x01' function='0x2'/>
    </controller>

    <interface type='user'>
      <alias name='net0'/>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x03' function='0x0'/>
      <rom bar='off'/>
      <model type='virtio'/>
    </interface>

    <serial type='pty'>
      <source path='/dev/pts/4'/>
      <target port='0'/>
      <alias name='serial0'/>
    </serial>

    <console type='pty' tty='/dev/pts/4'>
      <source path='/dev/pts/4'/>
      <target type='serial' port='0'/>
      <alias name='serial0'/>
     </console>

     <input type='mouse' bus='usb'/>

     <graphics type='vnc' port='-1' autoport='yes'>
       <listen type='address' address='::'/>
     </graphics>

     <video>
       <model type='vga' vram='9216' heads='1'/>
       <alias name='video0'/>
       <address type='pci' domain='0x0000' bus='0x00' slot='0x02' function='0x0'/>
     </video>

     <memballoon model='virtio'>
       <alias name='balloon0'/>
       <address type='pci' domain='0x0000' bus='0x00' slot='0x05' function='0x0'/>
       <stats period='30'/>
     </memballoon>

     <rng model='virtio'>
       <rate period="1000" bytes="1024"/>
       <backend model='random'>/dev/random</backend>
     </rng>
  </devices>
  <qemu:commandline xmlns:qemu='http://libvirt.org/schemas/domain/qemu/1.0'>
    <qemu:arg value='-redir'/>
    <qemu:arg value='tcp:{{.SSHPort}}::22'/>
  </qemu:commandline>
</domain>
`
