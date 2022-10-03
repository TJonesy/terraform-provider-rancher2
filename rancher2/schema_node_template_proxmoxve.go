package rancher2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

const (
	proxmoxveConfigDriver = "proxmoxve"
)

//Types

type proxmoxveConfig struct {
	ProvisionStrategy string `json:"provisionStrategy,omitempty" yaml:"provisionStrategy,omitempty"`
	Host              string `json:"proxmoxHost,omitempty" yaml:"proxmoxHost,omitempty"`
	Node              string `json:"proxmoxNode,omitempty" yaml:"proxmoxNode,omitempty"`
	User              string `json:"proxmoxUserName,omitempty" yaml:"proxmoxUserName,omitempty"`
	Password          string `json:"proxmoxUserPassword,omitempty" yaml:"proxmoxUserPassword,omitempty"`
	Realm             string `json:"proxmoxRealm,omitempty" yaml:"proxmoxRealm,omitempty"`
	ImageFile         string `json:"vmImageFile,omitempty" yaml:"vmImageFile,omitempty"`
	Pool              string `json:"proxmoxPool,omitempty" yaml:"proxmoxPool,omitempty"`
	StoragePath       string `json:"vmStoragePath,omitempty" yaml:"vmStoragePath,omitempty"`
	StorageType       string `json:"vmStorageType,omitempty" yaml:"vmStorageType,omitempty"`
	StorageSize       string `json:"vmStorageSize,omitempty" yaml:"vmStorageSize,omitempty"`
	Memory            string `json:"vmMemory,omitempty" yaml:"vmMemory,omitempty"`
	Onboot            string `json:"vmStartOnBoot,omitempty" yaml:"vmStartOnBoot,omitempty"`
	Protection        string `json:"vmProtection,omitempty" yaml:"vmProtection,omitempty"`
	Citype            string `json:"vmCitype,omitempty" yaml:"vmCitype,omitempty"`
	NUMA              string `json:"vmNuma,omitempty" yaml:"vmNuma,omitempty"`
	CiEnabled         string `json:"vmCienabled,omitempty" yaml:"vmCienabled,omitempty"`
	NetModel          string `json:"vmNetModel,omitempty" yaml:"vmNetModel,omitempty"`
	NetFirewall       string `json:"vmNetFirewall,omitempty" yaml:"netFirewall,omitempty"`
	NetMtu            string `json:"vmNetMtu,omitempty" yaml:"vmNetMtu,omitempty"`
	NetBridge         string `json:"vmNetBridge,omitempty" yaml:"netBridge,omitempty"`
	NetVlanTag        string `json:"vmNetTag,omitempty" yaml:"vmNetTag,omitempty"`
	ScsiController    string `json:"vmScsiController,omitempty" yaml:"vmScsiController,omitempty"`
	ScsiAttributes    string `json:"vmScsiAttributes,omitempty" yaml:"vmScsiAttributes,omitempty"`
	VMIDRange         string `json:"vmVmidRange,omitempty" yaml:"vmVmidRange,omitempty"`
	CloneVMID         string `json:"vmCloneVmid,omitempty" yaml:"vmCloneVmid,omitempty"`
	CloneFull         string `json:"vmCloneFull,omitempty" yaml:"vmCloneFull,omitempty"`
	GuestUsername     string `json:"sshUsername,omitempty" yaml:"sshUsername,omitempty"`
	GuestPassword     string `json:"sshPassword,omitempty" yaml:"sshPassword,omitempty"`
	GuestSSHPort      string `json:"sshPort,omitempty" yaml:"sshPort,omitempty"`
	CPU               string `json:"vmCpu,omitempty" yaml:"vmCpu,omitempty"`
	CPUSockets        string `json:"vmCpuSockets,omitempty" yaml:"vmCpuSockets,omitempty"`
	CPUCores          string `json:"vmCpuCores,omitempty" yaml:"vmCpuCores,omitempty"`
	DriverDebug       bool   `json:"debugDriver" yaml:"debugDriver"`
	RestyDebug        bool   `json:"debugResty" yaml:"debugResty"`
}

//Schemas

func proxmoxveConfigFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{

		"provision_strategy": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "clone",
			Description: "Strategy to use for provisioning (clone|cdrom)",
		},
		"host": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Proxmox API Host to connect to",
		},
		"node": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "optional, node to create VM on, host used if omitted but must match internal node name",
		},
		"user": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Proxmox API username",
		},
		"password": {
			Type:        schema.TypeString,
			Required:    true,
			Sensitive:   true,
			Description: "Promxmox API password",
		},
		"realm": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "ProxmoxAPI Realm, e.g. pam, pve, etc.",
		},
		"image_file": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "CDROM boot Image File in the format <storagename>:iso/<filename>.iso",
		},
		"pool": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Pool to add the VM to (necessary for users with only pool permission)",
		},
		"storage_path": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Internal PVE storage name, leave blank for clone default behavior",
		},
		"storage_type": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Type of the storage (currently QCOW2 and RAW)",
		},
		"storage_size": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Override disk size in GB",
		},
		"memory": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Memory GB",
		},
		"onboot": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Specifies whether a VM will be started during system bootup.",
		},
		"protection": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Sets the protection flag of the VM. This will disable the remove VM and remove disk operations.",
		},
		"citype": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "cloud-init type (nocloud|configdrive2), leave blank for clone default behavior",
		},
		"numa": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Enable/disable NUMA",
		},

		"ci_enabled": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Whether to use cloudinit or not (0|1), Leave blank for clone default behavior",
		},
		"net_model": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Net Interface Model, [e1000, virtio, realtek, etc...]",
		},
		"net_firewall": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Enable/disable firewall",
		},
		"net_mtu": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "set nic MTU",
		},
		"net_bridge": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "bridge applied to network interface",
		},
		"net_vlan_tag": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "vlan tag",
		},
		"scsi_controller": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "SCSI Controller",
		},
		"scsi_attributes": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "SCSI Tag",
		},
		"vm_id_range": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Acceptable range of VMIDs (<low>[:<high>])",
		},
		"clone_vm_id": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "VM ID to clone",
		},
		"clone_full": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Make a full (detached) clone from parent (0=LinkedClone, 1=FullClone, 2=ProxmoxDefaultBehavior)",
		},
		"guest_username": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "User to log into the guest OS to copy the public key",
		},
		"guest_password": {
			Type:        schema.TypeString,
			Optional:    true,
			Sensitive:   true,
			Description: "Password to log into the guest OS to copy the public key",
		},
		"guest_ssh_port": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     22,
			Description: "SSH port to log into the guest OS to copy the public key",
		},
		"cpu": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Emulated CPU type.",
		},
		"cpu_sockets": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The number of cpu sockets.",
		},
		"cpu_cores": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The number of cores per socket.",
		},
		"driver_debug": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "Enables debugging in the driver",
		},
		"resty_debug": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "Enables resty debugging",
		},
	}

	return s
}
