/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-7658
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// HyperflexHypervisorVirtualMachine A virtual machine belonging to the HyperFlex cluster spawned via the hypervisor.
type HyperflexHypervisorVirtualMachine struct {
	VirtualizationBaseVirtualMachine
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The connectivity state of the HyperFlex virtual machine.
	ConnectionState *string `json:"ConnectionState,omitempty"`
	// Guest operating system state of the HyperFlex virtual machine.
	GuestOsState *string `json:"GuestOsState,omitempty"`
	// Host UUID of the HyperFlex virtual machine.
	HostUuid *string                                `json:"HostUuid,omitempty"`
	Ip       NullableNetworkHyperFlexNetworkAddress `json:"Ip,omitempty"`
	// Directory path where virtual machine is stored.
	Path *string `json:"Path,omitempty"`
	// The instance id of platform which a virtual machine is running on.
	PlatformInstanceId *string `json:"PlatformInstanceId,omitempty"`
	// Total provisioned storage to the HyperFlex virtual machine in bytes.
	StorageProvisionedInBytes *int64 `json:"StorageProvisionedInBytes,omitempty"`
	// Storage used by HyperFlex virtual machine in bytes.
	StorageUsedInBytes *int64 `json:"StorageUsedInBytes,omitempty"`
	// Flag indicating whether or not this virtual machine is a template. Apply to the ESXi platform only.
	Template *bool `json:"Template,omitempty"`
	// The instance UUID of a virtual machine.
	VmInstanceUuid       *string                              `json:"VmInstanceUuid,omitempty"`
	Cluster              *HyperflexClusterRelationship        `json:"Cluster,omitempty"`
	Host                 *HyperflexHypervisorHostRelationship `json:"Host,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexHypervisorVirtualMachine HyperflexHypervisorVirtualMachine

// NewHyperflexHypervisorVirtualMachine instantiates a new HyperflexHypervisorVirtualMachine object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexHypervisorVirtualMachine(classId string, objectType string) *HyperflexHypervisorVirtualMachine {
	this := HyperflexHypervisorVirtualMachine{}
	this.ClassId = classId
	this.ObjectType = objectType
	var hypervisorType string = "ESXi"
	this.HypervisorType = &hypervisorType
	var powerState string = "Unknown"
	this.PowerState = &powerState
	var provider string = "Unknown"
	this.Provider = &provider
	var state string = "None"
	this.State = &state
	return &this
}

// NewHyperflexHypervisorVirtualMachineWithDefaults instantiates a new HyperflexHypervisorVirtualMachine object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexHypervisorVirtualMachineWithDefaults() *HyperflexHypervisorVirtualMachine {
	this := HyperflexHypervisorVirtualMachine{}
	var classId string = "hyperflex.HypervisorVirtualMachine"
	this.ClassId = classId
	var objectType string = "hyperflex.HypervisorVirtualMachine"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexHypervisorVirtualMachine) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexHypervisorVirtualMachine) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexHypervisorVirtualMachine) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexHypervisorVirtualMachine) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexHypervisorVirtualMachine) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexHypervisorVirtualMachine) SetObjectType(v string) {
	o.ObjectType = v
}

// GetConnectionState returns the ConnectionState field value if set, zero value otherwise.
func (o *HyperflexHypervisorVirtualMachine) GetConnectionState() string {
	if o == nil || o.ConnectionState == nil {
		var ret string
		return ret
	}
	return *o.ConnectionState
}

// GetConnectionStateOk returns a tuple with the ConnectionState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHypervisorVirtualMachine) GetConnectionStateOk() (*string, bool) {
	if o == nil || o.ConnectionState == nil {
		return nil, false
	}
	return o.ConnectionState, true
}

// HasConnectionState returns a boolean if a field has been set.
func (o *HyperflexHypervisorVirtualMachine) HasConnectionState() bool {
	if o != nil && o.ConnectionState != nil {
		return true
	}

	return false
}

// SetConnectionState gets a reference to the given string and assigns it to the ConnectionState field.
func (o *HyperflexHypervisorVirtualMachine) SetConnectionState(v string) {
	o.ConnectionState = &v
}

// GetGuestOsState returns the GuestOsState field value if set, zero value otherwise.
func (o *HyperflexHypervisorVirtualMachine) GetGuestOsState() string {
	if o == nil || o.GuestOsState == nil {
		var ret string
		return ret
	}
	return *o.GuestOsState
}

// GetGuestOsStateOk returns a tuple with the GuestOsState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHypervisorVirtualMachine) GetGuestOsStateOk() (*string, bool) {
	if o == nil || o.GuestOsState == nil {
		return nil, false
	}
	return o.GuestOsState, true
}

// HasGuestOsState returns a boolean if a field has been set.
func (o *HyperflexHypervisorVirtualMachine) HasGuestOsState() bool {
	if o != nil && o.GuestOsState != nil {
		return true
	}

	return false
}

// SetGuestOsState gets a reference to the given string and assigns it to the GuestOsState field.
func (o *HyperflexHypervisorVirtualMachine) SetGuestOsState(v string) {
	o.GuestOsState = &v
}

// GetHostUuid returns the HostUuid field value if set, zero value otherwise.
func (o *HyperflexHypervisorVirtualMachine) GetHostUuid() string {
	if o == nil || o.HostUuid == nil {
		var ret string
		return ret
	}
	return *o.HostUuid
}

// GetHostUuidOk returns a tuple with the HostUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHypervisorVirtualMachine) GetHostUuidOk() (*string, bool) {
	if o == nil || o.HostUuid == nil {
		return nil, false
	}
	return o.HostUuid, true
}

// HasHostUuid returns a boolean if a field has been set.
func (o *HyperflexHypervisorVirtualMachine) HasHostUuid() bool {
	if o != nil && o.HostUuid != nil {
		return true
	}

	return false
}

// SetHostUuid gets a reference to the given string and assigns it to the HostUuid field.
func (o *HyperflexHypervisorVirtualMachine) SetHostUuid(v string) {
	o.HostUuid = &v
}

// GetIp returns the Ip field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexHypervisorVirtualMachine) GetIp() NetworkHyperFlexNetworkAddress {
	if o == nil || o.Ip.Get() == nil {
		var ret NetworkHyperFlexNetworkAddress
		return ret
	}
	return *o.Ip.Get()
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexHypervisorVirtualMachine) GetIpOk() (*NetworkHyperFlexNetworkAddress, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ip.Get(), o.Ip.IsSet()
}

// HasIp returns a boolean if a field has been set.
func (o *HyperflexHypervisorVirtualMachine) HasIp() bool {
	if o != nil && o.Ip.IsSet() {
		return true
	}

	return false
}

// SetIp gets a reference to the given NullableNetworkHyperFlexNetworkAddress and assigns it to the Ip field.
func (o *HyperflexHypervisorVirtualMachine) SetIp(v NetworkHyperFlexNetworkAddress) {
	o.Ip.Set(&v)
}

// SetIpNil sets the value for Ip to be an explicit nil
func (o *HyperflexHypervisorVirtualMachine) SetIpNil() {
	o.Ip.Set(nil)
}

// UnsetIp ensures that no value is present for Ip, not even an explicit nil
func (o *HyperflexHypervisorVirtualMachine) UnsetIp() {
	o.Ip.Unset()
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *HyperflexHypervisorVirtualMachine) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHypervisorVirtualMachine) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *HyperflexHypervisorVirtualMachine) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *HyperflexHypervisorVirtualMachine) SetPath(v string) {
	o.Path = &v
}

// GetPlatformInstanceId returns the PlatformInstanceId field value if set, zero value otherwise.
func (o *HyperflexHypervisorVirtualMachine) GetPlatformInstanceId() string {
	if o == nil || o.PlatformInstanceId == nil {
		var ret string
		return ret
	}
	return *o.PlatformInstanceId
}

// GetPlatformInstanceIdOk returns a tuple with the PlatformInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHypervisorVirtualMachine) GetPlatformInstanceIdOk() (*string, bool) {
	if o == nil || o.PlatformInstanceId == nil {
		return nil, false
	}
	return o.PlatformInstanceId, true
}

// HasPlatformInstanceId returns a boolean if a field has been set.
func (o *HyperflexHypervisorVirtualMachine) HasPlatformInstanceId() bool {
	if o != nil && o.PlatformInstanceId != nil {
		return true
	}

	return false
}

// SetPlatformInstanceId gets a reference to the given string and assigns it to the PlatformInstanceId field.
func (o *HyperflexHypervisorVirtualMachine) SetPlatformInstanceId(v string) {
	o.PlatformInstanceId = &v
}

// GetStorageProvisionedInBytes returns the StorageProvisionedInBytes field value if set, zero value otherwise.
func (o *HyperflexHypervisorVirtualMachine) GetStorageProvisionedInBytes() int64 {
	if o == nil || o.StorageProvisionedInBytes == nil {
		var ret int64
		return ret
	}
	return *o.StorageProvisionedInBytes
}

// GetStorageProvisionedInBytesOk returns a tuple with the StorageProvisionedInBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHypervisorVirtualMachine) GetStorageProvisionedInBytesOk() (*int64, bool) {
	if o == nil || o.StorageProvisionedInBytes == nil {
		return nil, false
	}
	return o.StorageProvisionedInBytes, true
}

// HasStorageProvisionedInBytes returns a boolean if a field has been set.
func (o *HyperflexHypervisorVirtualMachine) HasStorageProvisionedInBytes() bool {
	if o != nil && o.StorageProvisionedInBytes != nil {
		return true
	}

	return false
}

// SetStorageProvisionedInBytes gets a reference to the given int64 and assigns it to the StorageProvisionedInBytes field.
func (o *HyperflexHypervisorVirtualMachine) SetStorageProvisionedInBytes(v int64) {
	o.StorageProvisionedInBytes = &v
}

// GetStorageUsedInBytes returns the StorageUsedInBytes field value if set, zero value otherwise.
func (o *HyperflexHypervisorVirtualMachine) GetStorageUsedInBytes() int64 {
	if o == nil || o.StorageUsedInBytes == nil {
		var ret int64
		return ret
	}
	return *o.StorageUsedInBytes
}

// GetStorageUsedInBytesOk returns a tuple with the StorageUsedInBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHypervisorVirtualMachine) GetStorageUsedInBytesOk() (*int64, bool) {
	if o == nil || o.StorageUsedInBytes == nil {
		return nil, false
	}
	return o.StorageUsedInBytes, true
}

// HasStorageUsedInBytes returns a boolean if a field has been set.
func (o *HyperflexHypervisorVirtualMachine) HasStorageUsedInBytes() bool {
	if o != nil && o.StorageUsedInBytes != nil {
		return true
	}

	return false
}

// SetStorageUsedInBytes gets a reference to the given int64 and assigns it to the StorageUsedInBytes field.
func (o *HyperflexHypervisorVirtualMachine) SetStorageUsedInBytes(v int64) {
	o.StorageUsedInBytes = &v
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *HyperflexHypervisorVirtualMachine) GetTemplate() bool {
	if o == nil || o.Template == nil {
		var ret bool
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHypervisorVirtualMachine) GetTemplateOk() (*bool, bool) {
	if o == nil || o.Template == nil {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *HyperflexHypervisorVirtualMachine) HasTemplate() bool {
	if o != nil && o.Template != nil {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given bool and assigns it to the Template field.
func (o *HyperflexHypervisorVirtualMachine) SetTemplate(v bool) {
	o.Template = &v
}

// GetVmInstanceUuid returns the VmInstanceUuid field value if set, zero value otherwise.
func (o *HyperflexHypervisorVirtualMachine) GetVmInstanceUuid() string {
	if o == nil || o.VmInstanceUuid == nil {
		var ret string
		return ret
	}
	return *o.VmInstanceUuid
}

// GetVmInstanceUuidOk returns a tuple with the VmInstanceUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHypervisorVirtualMachine) GetVmInstanceUuidOk() (*string, bool) {
	if o == nil || o.VmInstanceUuid == nil {
		return nil, false
	}
	return o.VmInstanceUuid, true
}

// HasVmInstanceUuid returns a boolean if a field has been set.
func (o *HyperflexHypervisorVirtualMachine) HasVmInstanceUuid() bool {
	if o != nil && o.VmInstanceUuid != nil {
		return true
	}

	return false
}

// SetVmInstanceUuid gets a reference to the given string and assigns it to the VmInstanceUuid field.
func (o *HyperflexHypervisorVirtualMachine) SetVmInstanceUuid(v string) {
	o.VmInstanceUuid = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *HyperflexHypervisorVirtualMachine) GetCluster() HyperflexClusterRelationship {
	if o == nil || o.Cluster == nil {
		var ret HyperflexClusterRelationship
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHypervisorVirtualMachine) GetClusterOk() (*HyperflexClusterRelationship, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *HyperflexHypervisorVirtualMachine) HasCluster() bool {
	if o != nil && o.Cluster != nil {
		return true
	}

	return false
}

// SetCluster gets a reference to the given HyperflexClusterRelationship and assigns it to the Cluster field.
func (o *HyperflexHypervisorVirtualMachine) SetCluster(v HyperflexClusterRelationship) {
	o.Cluster = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *HyperflexHypervisorVirtualMachine) GetHost() HyperflexHypervisorHostRelationship {
	if o == nil || o.Host == nil {
		var ret HyperflexHypervisorHostRelationship
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHypervisorVirtualMachine) GetHostOk() (*HyperflexHypervisorHostRelationship, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *HyperflexHypervisorVirtualMachine) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given HyperflexHypervisorHostRelationship and assigns it to the Host field.
func (o *HyperflexHypervisorVirtualMachine) SetHost(v HyperflexHypervisorHostRelationship) {
	o.Host = &v
}

func (o HyperflexHypervisorVirtualMachine) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedVirtualizationBaseVirtualMachine, errVirtualizationBaseVirtualMachine := json.Marshal(o.VirtualizationBaseVirtualMachine)
	if errVirtualizationBaseVirtualMachine != nil {
		return []byte{}, errVirtualizationBaseVirtualMachine
	}
	errVirtualizationBaseVirtualMachine = json.Unmarshal([]byte(serializedVirtualizationBaseVirtualMachine), &toSerialize)
	if errVirtualizationBaseVirtualMachine != nil {
		return []byte{}, errVirtualizationBaseVirtualMachine
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ConnectionState != nil {
		toSerialize["ConnectionState"] = o.ConnectionState
	}
	if o.GuestOsState != nil {
		toSerialize["GuestOsState"] = o.GuestOsState
	}
	if o.HostUuid != nil {
		toSerialize["HostUuid"] = o.HostUuid
	}
	if o.Ip.IsSet() {
		toSerialize["Ip"] = o.Ip.Get()
	}
	if o.Path != nil {
		toSerialize["Path"] = o.Path
	}
	if o.PlatformInstanceId != nil {
		toSerialize["PlatformInstanceId"] = o.PlatformInstanceId
	}
	if o.StorageProvisionedInBytes != nil {
		toSerialize["StorageProvisionedInBytes"] = o.StorageProvisionedInBytes
	}
	if o.StorageUsedInBytes != nil {
		toSerialize["StorageUsedInBytes"] = o.StorageUsedInBytes
	}
	if o.Template != nil {
		toSerialize["Template"] = o.Template
	}
	if o.VmInstanceUuid != nil {
		toSerialize["VmInstanceUuid"] = o.VmInstanceUuid
	}
	if o.Cluster != nil {
		toSerialize["Cluster"] = o.Cluster
	}
	if o.Host != nil {
		toSerialize["Host"] = o.Host
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexHypervisorVirtualMachine) UnmarshalJSON(bytes []byte) (err error) {
	type HyperflexHypervisorVirtualMachineWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The connectivity state of the HyperFlex virtual machine.
		ConnectionState *string `json:"ConnectionState,omitempty"`
		// Guest operating system state of the HyperFlex virtual machine.
		GuestOsState *string `json:"GuestOsState,omitempty"`
		// Host UUID of the HyperFlex virtual machine.
		HostUuid *string                                `json:"HostUuid,omitempty"`
		Ip       NullableNetworkHyperFlexNetworkAddress `json:"Ip,omitempty"`
		// Directory path where virtual machine is stored.
		Path *string `json:"Path,omitempty"`
		// The instance id of platform which a virtual machine is running on.
		PlatformInstanceId *string `json:"PlatformInstanceId,omitempty"`
		// Total provisioned storage to the HyperFlex virtual machine in bytes.
		StorageProvisionedInBytes *int64 `json:"StorageProvisionedInBytes,omitempty"`
		// Storage used by HyperFlex virtual machine in bytes.
		StorageUsedInBytes *int64 `json:"StorageUsedInBytes,omitempty"`
		// Flag indicating whether or not this virtual machine is a template. Apply to the ESXi platform only.
		Template *bool `json:"Template,omitempty"`
		// The instance UUID of a virtual machine.
		VmInstanceUuid *string                              `json:"VmInstanceUuid,omitempty"`
		Cluster        *HyperflexClusterRelationship        `json:"Cluster,omitempty"`
		Host           *HyperflexHypervisorHostRelationship `json:"Host,omitempty"`
	}

	varHyperflexHypervisorVirtualMachineWithoutEmbeddedStruct := HyperflexHypervisorVirtualMachineWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHyperflexHypervisorVirtualMachineWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexHypervisorVirtualMachine := _HyperflexHypervisorVirtualMachine{}
		varHyperflexHypervisorVirtualMachine.ClassId = varHyperflexHypervisorVirtualMachineWithoutEmbeddedStruct.ClassId
		varHyperflexHypervisorVirtualMachine.ObjectType = varHyperflexHypervisorVirtualMachineWithoutEmbeddedStruct.ObjectType
		varHyperflexHypervisorVirtualMachine.ConnectionState = varHyperflexHypervisorVirtualMachineWithoutEmbeddedStruct.ConnectionState
		varHyperflexHypervisorVirtualMachine.GuestOsState = varHyperflexHypervisorVirtualMachineWithoutEmbeddedStruct.GuestOsState
		varHyperflexHypervisorVirtualMachine.HostUuid = varHyperflexHypervisorVirtualMachineWithoutEmbeddedStruct.HostUuid
		varHyperflexHypervisorVirtualMachine.Ip = varHyperflexHypervisorVirtualMachineWithoutEmbeddedStruct.Ip
		varHyperflexHypervisorVirtualMachine.Path = varHyperflexHypervisorVirtualMachineWithoutEmbeddedStruct.Path
		varHyperflexHypervisorVirtualMachine.PlatformInstanceId = varHyperflexHypervisorVirtualMachineWithoutEmbeddedStruct.PlatformInstanceId
		varHyperflexHypervisorVirtualMachine.StorageProvisionedInBytes = varHyperflexHypervisorVirtualMachineWithoutEmbeddedStruct.StorageProvisionedInBytes
		varHyperflexHypervisorVirtualMachine.StorageUsedInBytes = varHyperflexHypervisorVirtualMachineWithoutEmbeddedStruct.StorageUsedInBytes
		varHyperflexHypervisorVirtualMachine.Template = varHyperflexHypervisorVirtualMachineWithoutEmbeddedStruct.Template
		varHyperflexHypervisorVirtualMachine.VmInstanceUuid = varHyperflexHypervisorVirtualMachineWithoutEmbeddedStruct.VmInstanceUuid
		varHyperflexHypervisorVirtualMachine.Cluster = varHyperflexHypervisorVirtualMachineWithoutEmbeddedStruct.Cluster
		varHyperflexHypervisorVirtualMachine.Host = varHyperflexHypervisorVirtualMachineWithoutEmbeddedStruct.Host
		*o = HyperflexHypervisorVirtualMachine(varHyperflexHypervisorVirtualMachine)
	} else {
		return err
	}

	varHyperflexHypervisorVirtualMachine := _HyperflexHypervisorVirtualMachine{}

	err = json.Unmarshal(bytes, &varHyperflexHypervisorVirtualMachine)
	if err == nil {
		o.VirtualizationBaseVirtualMachine = varHyperflexHypervisorVirtualMachine.VirtualizationBaseVirtualMachine
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ConnectionState")
		delete(additionalProperties, "GuestOsState")
		delete(additionalProperties, "HostUuid")
		delete(additionalProperties, "Ip")
		delete(additionalProperties, "Path")
		delete(additionalProperties, "PlatformInstanceId")
		delete(additionalProperties, "StorageProvisionedInBytes")
		delete(additionalProperties, "StorageUsedInBytes")
		delete(additionalProperties, "Template")
		delete(additionalProperties, "VmInstanceUuid")
		delete(additionalProperties, "Cluster")
		delete(additionalProperties, "Host")

		// remove fields from embedded structs
		reflectVirtualizationBaseVirtualMachine := reflect.ValueOf(o.VirtualizationBaseVirtualMachine)
		for i := 0; i < reflectVirtualizationBaseVirtualMachine.Type().NumField(); i++ {
			t := reflectVirtualizationBaseVirtualMachine.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexHypervisorVirtualMachine struct {
	value *HyperflexHypervisorVirtualMachine
	isSet bool
}

func (v NullableHyperflexHypervisorVirtualMachine) Get() *HyperflexHypervisorVirtualMachine {
	return v.value
}

func (v *NullableHyperflexHypervisorVirtualMachine) Set(val *HyperflexHypervisorVirtualMachine) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexHypervisorVirtualMachine) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexHypervisorVirtualMachine) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexHypervisorVirtualMachine(val *HyperflexHypervisorVirtualMachine) *NullableHyperflexHypervisorVirtualMachine {
	return &NullableHyperflexHypervisorVirtualMachine{value: val, isSet: true}
}

func (v NullableHyperflexHypervisorVirtualMachine) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexHypervisorVirtualMachine) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
