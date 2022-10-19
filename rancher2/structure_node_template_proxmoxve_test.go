package rancher2

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	managementClient "github.com/rancher/rancher/pkg/client/generated/management/v3"
)

var (
	testNodeTemplateProxmoxveNodeTaintsConf      []managementClient.Taint
	testNodeTemplateProxmoxveConf                proxmoxveConfig
	testNodeTemplateProxmoxveInterface           map[string]interface{}
	testNodeTemplateRootConf                     *NodeTemplate
	testNodeTemplateSquashProxmoxveConfInterface map[string]interface{}
	testNodeTemplateExpandProxmoxveConfInterface map[string]interface{}
	testNodeTemplateRootNodeTaintsInterface      interface{}
)

func init() {
	testNodeTemplateProxmoxveNodeTaintsConf = []managementClient.Taint{
		{
			Key:       "key",
			Value:     "value",
			Effect:    "recipient",
			TimeAdded: "time_added",
		},
	}
	testNodeTemplateRootNodeTaintsInterface = []interface{}{
		map[string]interface{}{
			"key":        "key",
			"value":      "value",
			"effect":     "recipient",
			"time_added": "time_added",
		},
	}
	testNodeTemplateProxmoxveConf = proxmoxveConfig{
		ProvisionStrategy: "ProvisionStrategy",
		Host:              "Host",
		Node:              "Node",
		User:              "User",
		Password:          "Password",
		Realm:             "Realm",
		ImageFile:         "ImageFile",
		Pool:              "Pool",
		StoragePath:       "Storage",
		StorageType:       "StorageType",
		StorageSize:       "StorageSize",
		Memory:            "Memory",
		StartOnBoot:       "OnBoot",
		Protection:        "Protection",
		Citype:            "Citype",
		NUMA:              "1",
		CiEnabled:         "CiEnabled",
		NetModel:          "NetModel",
		NetFirewall:       "NetFirewall",
		NetMtu:            "NetMtu",
		NetBridge:         "NetBridge",
		NetVlanTag:        "NetVlanTag",
		ScsiController:    "ScsiController",
		ScsiAttributes:    "ScsiAttributes",
		VMIDRange:         "VMIDRange",
		CloneVMID:         "CloneVMID",
		CloneFull:         "CloneFull",
		GuestUsername:     "GuestUsername",
		GuestPassword:     "GuestPassword",
		GuestSSHPort:      "GuestSSHPort",
		CPU:               "CPU",
		CPUSockets:        "CPUSockets",
		CPUCores:          "CPUCores",
		DriverDebug:       false,
		RestyDebug:        false,
	}
	testNodeTemplateProxmoxveInterface = map[string]interface{}{
		"provision_strategy": "ProvisionStrategy",
		"host":               "Host",
		"node":               "Node",
		"user":               "User",
		"password":           "Password",
		"realm":              "Realm",
		"image_file":         "ImageFile",
		"pool":               "Pool",
		"storage_path":       "Storage",
		"storage_type":       "StorageType",
		"storage_size":       "StorageSize",
		"memory":             "Memory",
		"start_on_boot":      "OnBoot",
		"protection":         "Protection",
		"citype":             "Citype",
		"numa":               "1",
		"ci_enabled":         "CiEnabled",
		"net_model":          "NetModel",
		"net_firewall":       "NetFirewall",
		"net_mtu":            "NetMtu",
		"net_bridge":         "NetBridge",
		"net_vlan_tag":       "NetVlanTag",
		"scsi_controller":    "ScsiController",
		"scsi_attributes":    "ScsiAttributes",
		"vm_id_range":        "VMIDRange",
		"clone_vm_id":        "CloneVMID",
		"clone_full":         "CloneFull",
		"guest_username":     "GuestUsername",
		"guest_password":     "GuestPassword",
		"guest_ssh_port":     "GuestSSHPort",
		"cpu":                "CPU",
		"cpu_sockets":        "CPUSockets",
		"cpu_cores":          "CPUCores",
		"debug_driver":       false,
		"debug_resty":        false,
	}
	testNodeTemplateAnnotationsConf := map[string]string{
		"key": "value",
	}
	testNodeTemplateAnnotationsInterface := map[string]interface{}{
		"key": "value",
	}
	useInternalIP := false
	testNodeTemplateRootConf = &NodeTemplate{
		NodeTemplate: managementClient.NodeTemplate{
			Driver:               "proxmoxve",
			Annotations:          testNodeTemplateAnnotationsConf,
			CloudCredentialID:    "abc-test-123",
			NodeTaints:           testNodeTemplateProxmoxveNodeTaintsConf,
			EngineInstallURL:     "http://fake.url",
			Name:                 "test-node-template",
			UseInternalIPAddress: &useInternalIP,
		},
		ProxmoxveConfig: &testNodeTemplateProxmoxveConf,
	}
	testNodeTemplateSquashProxmoxveConfInterface = map[string]interface{}{
		"annotations":             testNodeTemplateAnnotationsInterface,
		"driver":                  "proxmoxve",
		"cloud_credential_id":     "abc-test-123",
		"use_internal_ip_address": useInternalIP,
		"engine_install_url":      "http://fake.url",
		"name":                    "test-node-template",
	}

	testNodeTemplateExpandProxmoxveConfInterface = map[string]interface{}{
		"annotations":             testNodeTemplateAnnotationsInterface,
		"node_taints":             testNodeTemplateRootNodeTaintsInterface,
		"driver":                  "Proxmoxve",
		"cloud_credential_id":     "abc-test-123",
		"use_internal_ip_address": useInternalIP,
		"engine_install_url":      "http://fake.url",
		"name":                    "test-node-template",
		"proxmoxve_config":        []interface{}{testNodeTemplateProxmoxveInterface},
	}

}

func TestFlattenProxmoxveNodeTemplate(t *testing.T) {
	cases := []struct {
		Input          *NodeTemplate
		ExpectedOutput map[string]interface{}
	}{
		{
			testNodeTemplateRootConf,
			testNodeTemplateSquashProxmoxveConfInterface,
		},
	}

	for _, tc := range cases {
		output := schema.TestResourceDataRaw(t, nodeTemplateFields(), map[string]interface{}{})
		err := flattenNodeTemplate(output, tc.Input)
		if err != nil {
			t.Fatalf("[ERROR] on flattener: %#v", err)
		}
		expectedOutput := map[string]interface{}{}
		for k := range tc.ExpectedOutput {
			expectedOutput[k] = output.Get(k)
		}
		if !reflect.DeepEqual(expectedOutput, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from flattener. \nExpected: %#v\nGiven: %#v", expectedOutput, tc.ExpectedOutput)
		}
	}
}

func TestExpandProxmoxveNodeTemplate(t *testing.T) {
	cases := []struct {
		Input          map[string]interface{}
		ExpectedOutput *NodeTemplate
	}{
		{
			Input:          testNodeTemplateExpandProxmoxveConfInterface,
			ExpectedOutput: testNodeTemplateRootConf,
		},
	}

	for _, tc := range cases {
		inputData := schema.TestResourceDataRaw(t, nodeTemplateFields(), tc.Input)
		output := expandNodeTemplate(inputData)
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			a, _ := json.MarshalIndent(tc.ExpectedOutput, "", " ")
			b, _ := json.MarshalIndent(output, "", " ")
			t.Fatalf("Unexpected output from expander.\nExpected: %s\nGiven: %s", a, b)
		}
	}
}
