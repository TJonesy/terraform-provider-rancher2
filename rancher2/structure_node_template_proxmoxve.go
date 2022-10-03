package rancher2

import "log"
import "encoding/json"

// Flatteners

func flattenProxmoxveConfig(in *proxmoxveConfig) []interface{} {

	in_pretty, _ := json.MarshalIndent(in, "", " ")
	log.Printf("[TRACE] flattenProxmoxveConfig IN: %s", in_pretty)
	obj := make(map[string]interface{})
	if in == nil {
		return []interface{}{}
	}
	obj["host"] = in.Host
	obj["user"] = in.User
	obj["password"] = in.Password
	obj["realm"] = in.Realm

	if len(in.ProvisionStrategy) > 0 {
		obj["provision_strategy"] = in.ProvisionStrategy
	}
	if len(in.Node) > 0 {
		obj["node"] = in.Node
	}

	if len(in.ImageFile) > 0 {
		obj["image_file"] = in.ImageFile
	}
	if len(in.Pool) > 0 {
		obj["pool"] = in.Pool
	}
	if len(in.StoragePath) > 0 {
		obj["storage_path"] = in.StoragePath
	}
	if len(in.StorageType) > 0 {
		obj["storage_type"] = in.StorageType
	}
	if len(in.StorageSize) > 0 {
		obj["storage_size"] = in.StorageSize
	}
	if len(in.Memory) > 0 {
		obj["memory"] = in.Memory
	}
	if len(in.Onboot) > 0 {
		obj["onboot"] = in.Onboot
	}
	if len(in.Protection) > 0 {
		obj["protection"] = in.Protection
	}
	if len(in.Citype) > 0 {
		obj["citype"] = in.Citype
	}
	if len(in.NUMA) > 0 {
		obj["numa"] = in.NUMA
	}
	if len(in.CiEnabled) > 0 {
		obj["ci_enabled"] = in.CiEnabled
	}
	if len(in.NetModel) > 0 {
		obj["net_model"] = in.NetModel
	}
	if len(in.NetFirewall) > 0 {
		obj["net_firewall"] = in.NetFirewall
	}
	if len(in.NetMtu) > 0 {
		obj["net_mtu"] = in.NetMtu
	}
	if len(in.NetBridge) > 0 {
		obj["net_bridge"] = in.NetBridge
	}
	if len(in.NetVlanTag) > 0 {
		obj["net_vlan_tag"] = in.NetVlanTag
	}
	if len(in.ScsiController) > 0 {
		obj["scsi_controller"] = in.ScsiController
	}
	if len(in.ScsiAttributes) > 0 {
		obj["scsi_attributes"] = in.ScsiAttributes
	}
	if len(in.VMIDRange) > 0 {
		obj["vm_id_range"] = in.VMIDRange
	}
	if len(in.CloneVMID) > 0 {
		obj["clone_vm_id"] = in.CloneVMID
	}
	if len(in.CloneFull) > 0 {
		obj["clone_full"] = in.CloneFull
	}
	if len(in.GuestUsername) > 0 {
		obj["guest_username"] = in.GuestUsername
	}
	if len(in.GuestPassword) > 0 {
		obj["guest_password"] = in.GuestPassword
	}
	if len(in.GuestSSHPort) > 0 {
		obj["guest_ssh_port"] = in.GuestSSHPort
	}
	if len(in.CPU) > 0 {
		obj["cpu"] = in.CPU
	}
	if len(in.CPUSockets) > 0 {
		obj["cpu_sockets"] = in.CPUSockets
	}
	if len(in.CPUCores) > 0 {
		obj["cpu_cores"] = in.CPUCores
	}
	obj["driver_debug"] = in.DriverDebug
	obj["resty_debug"] = in.RestyDebug

	obj_pretty, _ := json.MarshalIndent(obj, "", " ")
	log.Printf("[TRACE] flattenProxmoxveConfig OBJ: %s", obj_pretty)

	return []interface{}{obj}
}

// Expanders

func expandProxmoxveConfig(p []interface{}) *proxmoxveConfig {
	obj := &proxmoxveConfig{}
	if len(p) == 0 || p[0] == nil {
		return obj
	}
	in := p[0].(map[string]interface{})

	in_pretty, _ := json.MarshalIndent(in, "", " ")
	log.Printf("[TRACE] expandProxmoxveConfig IN: %s", in_pretty)

	if v, ok := in["provision_strategy"].(string); ok {
		obj.ProvisionStrategy = v
	}
	if v, ok := in["host"].(string); ok {
		obj.Host = v
	}
	if v, ok := in["user"].(string); ok {
		obj.User = v
	}
	if v, ok := in["password"].(string); ok {
		obj.Password = v
	}
	if v, ok := in["realm"].(string); ok {
		obj.Realm = v
	}

	if v, ok := in["node"].(string); ok && len(v) > 0 {
		obj.Node = v
	}
	if v, ok := in["image_file"].(string); ok && len(v) > 0 {
		obj.ImageFile = v
	}
	if v, ok := in["pool"].(string); ok && len(v) > 0 {
		obj.Pool = v
	}
	if v, ok := in["storage_path"].(string); ok && len(v) > 0 {
		obj.StoragePath = v
	}
	if v, ok := in["storage_type"].(string); ok && len(v) > 0 {
		obj.StorageType = v
	}
	if v, ok := in["citype"].(string); ok && len(v) > 0 {
		obj.Citype = v
	}
	if v, ok := in["ci_enabled"].(string); ok && len(v) > 0 {
		obj.CiEnabled = v
	}
	if v, ok := in["net_model"].(string); ok && len(v) > 0 {
		obj.NetModel = v
	}
	if v, ok := in["net_mtu"].(string); ok && len(v) > 0 {
		obj.NetMtu = v
	}
	if v, ok := in["net_bridge"].(string); ok && len(v) > 0 {
		obj.NetBridge = v
	}
	if v, ok := in["scsi_controller"].(string); ok && len(v) > 0 {
		obj.ScsiController = v
	}
	if v, ok := in["scsi_attributes"].(string); ok && len(v) > 0 {
		obj.ScsiAttributes = v
	}
	if v, ok := in["vm_id_range"].(string); ok && len(v) > 0 {
		obj.VMIDRange = v
	}
	if v, ok := in["guest_username"].(string); ok && len(v) > 0 {
		obj.GuestUsername = v
	}
	if v, ok := in["guest_password"].(string); ok && len(v) > 0 {
		obj.GuestPassword = v
	}

	if v, ok := in["onboot"].(string); ok {
		obj.Onboot = v
	}

	if v, ok := in["protection"].(string); ok {
		obj.Protection = v
	}

	if v, ok := in["numa"].(string); ok {
		obj.NUMA = v
	}

	if v, ok := in["net_firewall"].(string); ok {
		obj.NetFirewall = v
	}

	if v, ok := in["driver_debug"].(bool); ok {
		obj.DriverDebug = v
	}
	if v, ok := in["resty_debug"].(bool); ok {
		obj.RestyDebug = v
	}

	if v, ok := in["storage_size"].(string); ok && len(v) > 0 {
		obj.StorageSize = v
	}

	if v, ok := in["memory"].(string); ok && len(v) > 0 {
		obj.Memory = v
	}

	if v, ok := in["net_vlan_tag"].(string); ok && len(v) > 0 {
		obj.NetVlanTag = v
	}

	if v, ok := in["clone_vm_id"].(string); ok && len(v) > 0 {
		obj.CloneVMID = v
	}

	if v, ok := in["clone_full"].(string); ok && len(v) > 0 {
		obj.CloneFull = v
	}

	if v, ok := in["guest_ssh_port"].(string); ok && len(v) > 0 {
		obj.GuestSSHPort = v
	}

	if v, ok := in["cpu"].(string); ok && len(v) > 0 {
		obj.CPU = v
	}

	if v, ok := in["cpu_sockets"].(string); ok && len(v) > 0 {
		obj.CPUSockets = v
	}

	if v, ok := in["cpu_cores"].(string); ok && len(v) > 0 {
		obj.CPUCores = v
	}

	obj_pretty, _ := json.MarshalIndent(obj, "", " ")
	log.Printf("[TRACE] expandProxmoxveConfig OBJ: %s", obj_pretty)

	return obj
}
