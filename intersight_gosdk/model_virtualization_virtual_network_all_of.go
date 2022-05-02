/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-6484
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// VirtualizationVirtualNetworkAllOf Definition of the list of properties defined in 'virtualization.VirtualNetwork', excluding properties defined in parent classes.
type VirtualizationVirtualNetworkAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Human readable description about this network.
	Description *string `json:"Description,omitempty"`
	// Flag to indicate whether the configuration is created from inventory object.
	Discovered *bool `json:"Discovered,omitempty"`
	// A flag to distinguish if a network belongs to an infrastructure network or a user defined network that guest workloads can use.
	InfraNetwork *bool `json:"InfraNetwork,omitempty"`
	// Maximum transmissible unit of data in bytes that can be sent across the network.
	Mtu *int64 `json:"Mtu,omitempty"`
	// Name of the virtual network. Name must be unique.
	Name *string `json:"Name,omitempty"`
	// Type of network layer, either L2 or L3. * `unknown` - This network is of an unknown network type. * `L2` - A Layer 2 switching network type.
	NetworkType *string  `json:"NetworkType,omitempty"`
	Trunk       []string `json:"Trunk,omitempty"`
	// A VLAN id set on the network attachment point.
	Vlan *int64 `json:"Vlan,omitempty"`
	// Name of the virtual switch.
	Vswitch              *string                                       `json:"Vswitch,omitempty"`
	Cluster              *VirtualizationBaseClusterRelationship        `json:"Cluster,omitempty"`
	Inventory            *VirtualizationBaseVirtualNetworkRelationship `json:"Inventory,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship          `json:"RegisteredDevice,omitempty"`
	WorkflowInfo         *WorkflowWorkflowInfoRelationship             `json:"WorkflowInfo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationVirtualNetworkAllOf VirtualizationVirtualNetworkAllOf

// NewVirtualizationVirtualNetworkAllOf instantiates a new VirtualizationVirtualNetworkAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationVirtualNetworkAllOf(classId string, objectType string) *VirtualizationVirtualNetworkAllOf {
	this := VirtualizationVirtualNetworkAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var networkType string = "unknown"
	this.NetworkType = &networkType
	return &this
}

// NewVirtualizationVirtualNetworkAllOfWithDefaults instantiates a new VirtualizationVirtualNetworkAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationVirtualNetworkAllOfWithDefaults() *VirtualizationVirtualNetworkAllOf {
	this := VirtualizationVirtualNetworkAllOf{}
	var classId string = "virtualization.VirtualNetwork"
	this.ClassId = classId
	var objectType string = "virtualization.VirtualNetwork"
	this.ObjectType = objectType
	var networkType string = "unknown"
	this.NetworkType = &networkType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationVirtualNetworkAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualNetworkAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationVirtualNetworkAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationVirtualNetworkAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualNetworkAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationVirtualNetworkAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *VirtualizationVirtualNetworkAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualNetworkAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *VirtualizationVirtualNetworkAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *VirtualizationVirtualNetworkAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetDiscovered returns the Discovered field value if set, zero value otherwise.
func (o *VirtualizationVirtualNetworkAllOf) GetDiscovered() bool {
	if o == nil || o.Discovered == nil {
		var ret bool
		return ret
	}
	return *o.Discovered
}

// GetDiscoveredOk returns a tuple with the Discovered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualNetworkAllOf) GetDiscoveredOk() (*bool, bool) {
	if o == nil || o.Discovered == nil {
		return nil, false
	}
	return o.Discovered, true
}

// HasDiscovered returns a boolean if a field has been set.
func (o *VirtualizationVirtualNetworkAllOf) HasDiscovered() bool {
	if o != nil && o.Discovered != nil {
		return true
	}

	return false
}

// SetDiscovered gets a reference to the given bool and assigns it to the Discovered field.
func (o *VirtualizationVirtualNetworkAllOf) SetDiscovered(v bool) {
	o.Discovered = &v
}

// GetInfraNetwork returns the InfraNetwork field value if set, zero value otherwise.
func (o *VirtualizationVirtualNetworkAllOf) GetInfraNetwork() bool {
	if o == nil || o.InfraNetwork == nil {
		var ret bool
		return ret
	}
	return *o.InfraNetwork
}

// GetInfraNetworkOk returns a tuple with the InfraNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualNetworkAllOf) GetInfraNetworkOk() (*bool, bool) {
	if o == nil || o.InfraNetwork == nil {
		return nil, false
	}
	return o.InfraNetwork, true
}

// HasInfraNetwork returns a boolean if a field has been set.
func (o *VirtualizationVirtualNetworkAllOf) HasInfraNetwork() bool {
	if o != nil && o.InfraNetwork != nil {
		return true
	}

	return false
}

// SetInfraNetwork gets a reference to the given bool and assigns it to the InfraNetwork field.
func (o *VirtualizationVirtualNetworkAllOf) SetInfraNetwork(v bool) {
	o.InfraNetwork = &v
}

// GetMtu returns the Mtu field value if set, zero value otherwise.
func (o *VirtualizationVirtualNetworkAllOf) GetMtu() int64 {
	if o == nil || o.Mtu == nil {
		var ret int64
		return ret
	}
	return *o.Mtu
}

// GetMtuOk returns a tuple with the Mtu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualNetworkAllOf) GetMtuOk() (*int64, bool) {
	if o == nil || o.Mtu == nil {
		return nil, false
	}
	return o.Mtu, true
}

// HasMtu returns a boolean if a field has been set.
func (o *VirtualizationVirtualNetworkAllOf) HasMtu() bool {
	if o != nil && o.Mtu != nil {
		return true
	}

	return false
}

// SetMtu gets a reference to the given int64 and assigns it to the Mtu field.
func (o *VirtualizationVirtualNetworkAllOf) SetMtu(v int64) {
	o.Mtu = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VirtualizationVirtualNetworkAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualNetworkAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VirtualizationVirtualNetworkAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VirtualizationVirtualNetworkAllOf) SetName(v string) {
	o.Name = &v
}

// GetNetworkType returns the NetworkType field value if set, zero value otherwise.
func (o *VirtualizationVirtualNetworkAllOf) GetNetworkType() string {
	if o == nil || o.NetworkType == nil {
		var ret string
		return ret
	}
	return *o.NetworkType
}

// GetNetworkTypeOk returns a tuple with the NetworkType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualNetworkAllOf) GetNetworkTypeOk() (*string, bool) {
	if o == nil || o.NetworkType == nil {
		return nil, false
	}
	return o.NetworkType, true
}

// HasNetworkType returns a boolean if a field has been set.
func (o *VirtualizationVirtualNetworkAllOf) HasNetworkType() bool {
	if o != nil && o.NetworkType != nil {
		return true
	}

	return false
}

// SetNetworkType gets a reference to the given string and assigns it to the NetworkType field.
func (o *VirtualizationVirtualNetworkAllOf) SetNetworkType(v string) {
	o.NetworkType = &v
}

// GetTrunk returns the Trunk field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationVirtualNetworkAllOf) GetTrunk() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Trunk
}

// GetTrunkOk returns a tuple with the Trunk field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationVirtualNetworkAllOf) GetTrunkOk() ([]string, bool) {
	if o == nil || o.Trunk == nil {
		return nil, false
	}
	return o.Trunk, true
}

// HasTrunk returns a boolean if a field has been set.
func (o *VirtualizationVirtualNetworkAllOf) HasTrunk() bool {
	if o != nil && o.Trunk != nil {
		return true
	}

	return false
}

// SetTrunk gets a reference to the given []string and assigns it to the Trunk field.
func (o *VirtualizationVirtualNetworkAllOf) SetTrunk(v []string) {
	o.Trunk = v
}

// GetVlan returns the Vlan field value if set, zero value otherwise.
func (o *VirtualizationVirtualNetworkAllOf) GetVlan() int64 {
	if o == nil || o.Vlan == nil {
		var ret int64
		return ret
	}
	return *o.Vlan
}

// GetVlanOk returns a tuple with the Vlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualNetworkAllOf) GetVlanOk() (*int64, bool) {
	if o == nil || o.Vlan == nil {
		return nil, false
	}
	return o.Vlan, true
}

// HasVlan returns a boolean if a field has been set.
func (o *VirtualizationVirtualNetworkAllOf) HasVlan() bool {
	if o != nil && o.Vlan != nil {
		return true
	}

	return false
}

// SetVlan gets a reference to the given int64 and assigns it to the Vlan field.
func (o *VirtualizationVirtualNetworkAllOf) SetVlan(v int64) {
	o.Vlan = &v
}

// GetVswitch returns the Vswitch field value if set, zero value otherwise.
func (o *VirtualizationVirtualNetworkAllOf) GetVswitch() string {
	if o == nil || o.Vswitch == nil {
		var ret string
		return ret
	}
	return *o.Vswitch
}

// GetVswitchOk returns a tuple with the Vswitch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualNetworkAllOf) GetVswitchOk() (*string, bool) {
	if o == nil || o.Vswitch == nil {
		return nil, false
	}
	return o.Vswitch, true
}

// HasVswitch returns a boolean if a field has been set.
func (o *VirtualizationVirtualNetworkAllOf) HasVswitch() bool {
	if o != nil && o.Vswitch != nil {
		return true
	}

	return false
}

// SetVswitch gets a reference to the given string and assigns it to the Vswitch field.
func (o *VirtualizationVirtualNetworkAllOf) SetVswitch(v string) {
	o.Vswitch = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *VirtualizationVirtualNetworkAllOf) GetCluster() VirtualizationBaseClusterRelationship {
	if o == nil || o.Cluster == nil {
		var ret VirtualizationBaseClusterRelationship
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualNetworkAllOf) GetClusterOk() (*VirtualizationBaseClusterRelationship, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *VirtualizationVirtualNetworkAllOf) HasCluster() bool {
	if o != nil && o.Cluster != nil {
		return true
	}

	return false
}

// SetCluster gets a reference to the given VirtualizationBaseClusterRelationship and assigns it to the Cluster field.
func (o *VirtualizationVirtualNetworkAllOf) SetCluster(v VirtualizationBaseClusterRelationship) {
	o.Cluster = &v
}

// GetInventory returns the Inventory field value if set, zero value otherwise.
func (o *VirtualizationVirtualNetworkAllOf) GetInventory() VirtualizationBaseVirtualNetworkRelationship {
	if o == nil || o.Inventory == nil {
		var ret VirtualizationBaseVirtualNetworkRelationship
		return ret
	}
	return *o.Inventory
}

// GetInventoryOk returns a tuple with the Inventory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualNetworkAllOf) GetInventoryOk() (*VirtualizationBaseVirtualNetworkRelationship, bool) {
	if o == nil || o.Inventory == nil {
		return nil, false
	}
	return o.Inventory, true
}

// HasInventory returns a boolean if a field has been set.
func (o *VirtualizationVirtualNetworkAllOf) HasInventory() bool {
	if o != nil && o.Inventory != nil {
		return true
	}

	return false
}

// SetInventory gets a reference to the given VirtualizationBaseVirtualNetworkRelationship and assigns it to the Inventory field.
func (o *VirtualizationVirtualNetworkAllOf) SetInventory(v VirtualizationBaseVirtualNetworkRelationship) {
	o.Inventory = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *VirtualizationVirtualNetworkAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualNetworkAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *VirtualizationVirtualNetworkAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *VirtualizationVirtualNetworkAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

// GetWorkflowInfo returns the WorkflowInfo field value if set, zero value otherwise.
func (o *VirtualizationVirtualNetworkAllOf) GetWorkflowInfo() WorkflowWorkflowInfoRelationship {
	if o == nil || o.WorkflowInfo == nil {
		var ret WorkflowWorkflowInfoRelationship
		return ret
	}
	return *o.WorkflowInfo
}

// GetWorkflowInfoOk returns a tuple with the WorkflowInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualNetworkAllOf) GetWorkflowInfoOk() (*WorkflowWorkflowInfoRelationship, bool) {
	if o == nil || o.WorkflowInfo == nil {
		return nil, false
	}
	return o.WorkflowInfo, true
}

// HasWorkflowInfo returns a boolean if a field has been set.
func (o *VirtualizationVirtualNetworkAllOf) HasWorkflowInfo() bool {
	if o != nil && o.WorkflowInfo != nil {
		return true
	}

	return false
}

// SetWorkflowInfo gets a reference to the given WorkflowWorkflowInfoRelationship and assigns it to the WorkflowInfo field.
func (o *VirtualizationVirtualNetworkAllOf) SetWorkflowInfo(v WorkflowWorkflowInfoRelationship) {
	o.WorkflowInfo = &v
}

func (o VirtualizationVirtualNetworkAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.Discovered != nil {
		toSerialize["Discovered"] = o.Discovered
	}
	if o.InfraNetwork != nil {
		toSerialize["InfraNetwork"] = o.InfraNetwork
	}
	if o.Mtu != nil {
		toSerialize["Mtu"] = o.Mtu
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.NetworkType != nil {
		toSerialize["NetworkType"] = o.NetworkType
	}
	if o.Trunk != nil {
		toSerialize["Trunk"] = o.Trunk
	}
	if o.Vlan != nil {
		toSerialize["Vlan"] = o.Vlan
	}
	if o.Vswitch != nil {
		toSerialize["Vswitch"] = o.Vswitch
	}
	if o.Cluster != nil {
		toSerialize["Cluster"] = o.Cluster
	}
	if o.Inventory != nil {
		toSerialize["Inventory"] = o.Inventory
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	if o.WorkflowInfo != nil {
		toSerialize["WorkflowInfo"] = o.WorkflowInfo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationVirtualNetworkAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varVirtualizationVirtualNetworkAllOf := _VirtualizationVirtualNetworkAllOf{}

	if err = json.Unmarshal(bytes, &varVirtualizationVirtualNetworkAllOf); err == nil {
		*o = VirtualizationVirtualNetworkAllOf(varVirtualizationVirtualNetworkAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "Discovered")
		delete(additionalProperties, "InfraNetwork")
		delete(additionalProperties, "Mtu")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "NetworkType")
		delete(additionalProperties, "Trunk")
		delete(additionalProperties, "Vlan")
		delete(additionalProperties, "Vswitch")
		delete(additionalProperties, "Cluster")
		delete(additionalProperties, "Inventory")
		delete(additionalProperties, "RegisteredDevice")
		delete(additionalProperties, "WorkflowInfo")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVirtualizationVirtualNetworkAllOf struct {
	value *VirtualizationVirtualNetworkAllOf
	isSet bool
}

func (v NullableVirtualizationVirtualNetworkAllOf) Get() *VirtualizationVirtualNetworkAllOf {
	return v.value
}

func (v *NullableVirtualizationVirtualNetworkAllOf) Set(val *VirtualizationVirtualNetworkAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationVirtualNetworkAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationVirtualNetworkAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationVirtualNetworkAllOf(val *VirtualizationVirtualNetworkAllOf) *NullableVirtualizationVirtualNetworkAllOf {
	return &NullableVirtualizationVirtualNetworkAllOf{value: val, isSet: true}
}

func (v NullableVirtualizationVirtualNetworkAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationVirtualNetworkAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
