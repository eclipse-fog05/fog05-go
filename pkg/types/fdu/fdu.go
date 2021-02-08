/*
* Copyright (c) 2014,2021 Contributors to the Eclipse Foundation
* See the NOTICE file(s) distributed with this work for additional
* information regarding copyright ownership.
* This program and the accompanying materials are made available under the
* terms of the Eclipse Public License 2.0 which is available at
* http://www.eclipse.org/legal/epl-2.0, or the Apache License, Version 2.0
* which is available at https://www.apache.org/licenses/LICENSE-2.0.
* SPDX-License-Identifier: EPL-2.0 OR Apache-2.0
* Contributors: Gabriele Baldoni, ADLINK Technology Inc.
* golang APIs
 */

package fdu

import (
	"net"

	"encoding/json"
	// Leaving it here as the plan it to use it.
	_ "github.com/Masterminds/semver"

	"github.com/google/uuid"
)

const (
	// LIVE is Live Migration kind
	LIVE string = "LIVE"

	// COLD is cold Migration kind
	COLD string = "COLD"

	//SCRIPT is script configuration kind
	SCRIPT string = "SCRIPT"

	//CLOUDINIT is cloud init configuration kind
	CLOUDINIT string = "CLOUD_INIT"

	//INTERNAL is internal interface kind
	INTERNAL string = "INTERNAL"

	//EXTERNAL is external interface kind
	EXTERNAL string = "EXTERNAL"

	//WLAN is WLAN interface kind
	WLAN string = "WLAN"

	//BLUETOOTH is Bluetooth interface kind
	BLUETOOTH string = "BLUETOOTH"

	//PARAVIRT is paravirtualised interface kind
	PARAVIRT string = "PARAVIRT"

	//FOSMGMT is fog05 management interface kind
	FOSMGMT string = "FOS_MGMT"

	//PCIPASSTHROUGH is PCI passthrough interface kind
	PCIPASSTHROUGH string = "PCI_PASSTHROUGH"

	//SRIOV is SR-IOV interface kind
	SRIOV string = "SR_IOV"

	//E1000 is e1000 interface kind
	E1000 string = "E1000"

	//RTL8139 is rtl8139 interface kind
	RTL8139 string = "RTL8139"

	//PHYSICAL  is physical interface kind
	PHYSICAL string = "PHYSICAL"

	//BRIDGED is bridged interface kind
	BRIDGED string = "BRIDGED"

	//GPIO is GPIO port kind
	GPIO string = "GPIO"

	//I2C is I2C port kind
	I2C string = "I2C"

	//BUS is BUS port kind
	BUS string = "BUS"

	//COM is COM port kind
	COM string = "COM"

	//CAN is CAN port kind
	CAN string = "CAN"

	BLOCK  string = "BLOCK"
	FILE   string = "FILE"
	OBJECT string = "OBJECT"

	DEFINED    string = "DEFINED"
	CONFIGURED string = "CONFIGURED"
	RUNNING    string = "RUNNING"
)

// ScalingPolicy represent the scaling policy for an FDU Instance
type ScalingPolicy struct {
	Metric               string      `json:"metric"`
	ScaleUpThreshold     json.Number `json:"scale_up_threshold"`
	ScaleDownThreshold   json.Number `json:"scale_down_threshold"`
	ThresholdSensibility uint8       `json:"threshold_sensibility"`
	ProbeInterveal       json.Number `json:"probe_interveal"`
	MinReplicas          uint8       `json:"min_replicas"`
	MaxReplicas          uint8       `json:"max_replicas"`
}

// Position represents the FDU Position
type Position struct {
	Latitude  string      `json:"lat"`
	Longitude string      `json:"lon"`
	Radius    json.Number `json:"radius"`
}

// Proximity represents the FDU Proximity
type Proximity struct {
	Neighbor string      `json:"neighbor"`
	Radius   json.Number `json:"radius"`
}

// Configuration represents the FDU Configuration
type Configuration struct {
	ConfType string   `json:"conf_type"`
	Script   string   `json:"script"`
	SSHKeys  []string `json:"ssh_keys,omitempty"`
}

// Image represents an FDU image
type Image struct {
	UUID     *uuid.UUID `json:"uuid,omitempty"`
	Name     *string    `json:"name,omitempty"`
	URI      string     `json:"uri"`
	Checksum string     `json:"checksum"` //SHA256SUM
	Format   string     `json:"format"`
}

// ComputationalRequirements represents the FDU Computational Requirements aka Flavor
type ComputationalRequirements struct {
	CPUArch         string  `json:"cpu_arch"`
	CPUMinFrequency uint64  `json:"cpu_min_freq"`
	CPUMinCount     uint8   `json:"cpu_min_count"`
	GPUMinCount     uint8   `json:"gpu_min_count"`
	FPGAMinCount    uint8   `json:"fpga_min_count"`
	OperatingSystem *string `json:"operating_system,omitempty"`
	RAMSizeMB       uint32  `json:"ram_size_mb"`
	StorageSizeMB   uint32  `json:"storage_size_gb"`
}

// GeographicalRequirements represents the FDU Geographical Requirements
type GeographicalRequirements struct {
	Position  *Position    `json:"position,omitempty"`
	Proximity *[]Proximity `json:"proximity,omitempty"`
}

// VirtualInterface represents the FDU Virtual Interface
type VirtualInterface struct {
	VIfKind   string  `json:"vif_kind"`
	Parent    *string `json:"parent,omitempty"`
	Bandwidth *uint8  `json:"bandwidth,omitempty"`
}

// ConnectionPointDescriptor ...
type ConnectionPointDescriptor struct {
	UUID   *uuid.UUID `json:"uuid,omitempty"`
	Name   string     `json:"name"`
	ID     string     `json:"id"`
	VldRef *string    `json:"vld_ref,omitempty"`
}

// Interface represent and FDU Network Interface descriptor
type Interface struct {
	Name             string            `json:"name"`
	Kind             string            `json:"kind"`
	MacAddress       *net.HardwareAddr `json:"mac_address"`
	VirtualInterface VirtualInterface  `json:"virtual_interface"`
	CpID             *string           `json:"cp_id,omitemtpy"`
}

// StorageDescriptor represents an FDU Storage Descriptor
type StorageDescriptor struct {
	ID          string `json:"id"`
	StorageKind string `json:"storage_kind"`
	Size        uint32 `json:"size"`
}

// FDUDescriptor represent and FDU descriptor
type FDUDescriptor struct {
	UUID                     *uuid.UUID                  `json:"uuid,omitempty"`
	ID                       string                      `json:"id"`
	Name                     string                      `json:"name"`
	Version                  string                      `json:"version"`
	FDUVersion               string                      `json:"fdu_version"`
	Description              *string                     `json:"description,omitempty"`
	Hypervisor               string                      `json:"hypervisor"`
	Image                    *Image                      `json:"image,omitempty"`
	HypervisorSpecific       *string                     `json:"hypervisor_specific,omitempty"`
	ComputationRequirements  ComputationalRequirements   `json:"computation_requirements"`
	GeographicalRequirements *GeographicalRequirements   `json:"geographical_requirements,omitempty"`
	Interfaces               []Interface                 `json:"interfaces"`
	Storage                  []StorageDescriptor         `json:"storage"`
	ConnectionPoints         []ConnectionPointDescriptor `json:"connection_points"`
	Configuration            *Configuration              `json:"configuration,omitempty"`
	MigrationKind            string                      `json:"migration_kind"`
	Replicas                 *uint8                      `json:"replicas,omitempty"`
	DependsOn                []string                    `json:"depends_on"`
}

// FDURecord represent an FDU instance record
type FDURecord struct {
	UUID               uuid.UUID               `json:"uuid"`
	FDUUUID            uuid.UUID               `json:"fdu_uuid"`
	Node               uuid.UUID               `json:"node"`
	Interfaces         *[]InterfaceRecord      `json:"interfaces,omitempty"`
	ConnectionPoints   []ConnectionPointRecord `json:"connection_points"`
	Status             string                  `json:"status"`
	Error              *string                 `json:"error,omitempty"`
	HypervisorSpecific *[]byte                 `json:"hypervisor_specific,omitempty"`
	Restarts           uint32                  `json:"restarts"`
}

// StorageRecord represent an FDU Storage Record
type StorageRecord struct {
	UUID               string  `json:"uuid"`
	StorageID          string  `json:"storage_id"`
	StorageType        string  `json:"storage_type"`
	Size               uint32  `json:"size"`
	FileSystemProtocol *string `json:"file_system_protocol,omitempty"`
	CPID               *string `json:"cp_id,omitempty"`
}

// InterfaceRecord represent an FDU Interface Record
type InterfaceRecord struct {
	Name             string                 `json:"name"`
	Kind             string                 `json:"kind"`
	MacAddress       *net.HardwareAddr      `json:"mac_address"`
	VirtualInterface VirtualInterfaceRecord `json:"virtual_interface"`
	CpUUID           *uuid.UUID             `json:"cp_id,omitemtpy"`
	IntfUUID         uuid.UUID              `json:"intf_uuid"`
}

// VirtualInterfaceRecord represents the FDU Virtual Interface
type VirtualInterfaceRecord struct {
	VIfKind   string `json:"vif_kind"`
	Bandwidth *uint8 `json:"bandwidth,omitempty"`
}

// ConnectionPointRecord ...
type ConnectionPointRecord struct {
	UUID *uuid.UUID `json:"uuid,omitempty"`
	ID   string     `json:"id"`
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FDUDescriptor) DeepCopyInto(out *FDUDescriptor) {
	*out = *in
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FDURecord) DeepCopyInto(out *FDURecord) {
	*out = *in
	return
}
