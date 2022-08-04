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

// VirtualizationVmwareDistributedNetwork The VMware Distributed Virtual PortGroup object is represented here.
type VirtualizationVmwareDistributedNetwork struct {
	VirtualizationBaseDistributedNetwork
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// If forgedTransmits property value is set to reject, outbound frames with a source MAC address different from the one set on the adapter are dropped. If property value is set to accept, no filtering is performed and all outbound frames are passed. * `Reject` - Indicates that the security policy is rejected. * `Accept` - Indicates that the security policy is accepted.
	ForgedTransmits *string `json:"ForgedTransmits,omitempty"`
	// If macAddressChanges property value is set to reject and the MAC address of the adapter is changed to a value other than the one specified in .vmx configuration file, all inbound frames are dropped. If property value is set to accept and the MAC address is changed, inbound frames to the new MAC address are received. * `Reject` - Indicates that the security policy is rejected. * `Accept` - Indicates that the security policy is accepted.
	MacAddressChanges     *string                                        `json:"MacAddressChanges,omitempty"`
	NicTeamingAndFailover NullableVirtualizationVmwareTeamingAndFailover `json:"NicTeamingAndFailover,omitempty"`
	// The total number of hosts connected to this distributed virtual network.
	NumHosts *int64 `json:"NumHosts,omitempty"`
	// The total number of ports in the distributed virtual network.
	NumPorts *int64 `json:"NumPorts,omitempty"`
	// If promiscuousMode property value is set to reject, incoming traffic only targeted to that network will be visible. If property value is set to accept, objects defined within the network can see all incoming traffic on the virtual switch based on the VLAN policy. * `Reject` - Indicates that the security policy is rejected. * `Accept` - Indicates that the security policy is accepted.
	PromiscuousMode *string `json:"PromiscuousMode,omitempty"`
	// Indicates if the distributed virtual network is a uplink.
	UpLink    *bool                           `json:"UpLink,omitempty"`
	VlanRange []VirtualizationVmwareVlanRange `json:"VlanRange,omitempty"`
	// VLAN type of the distributed virtual network. It can be None, VLAN, VLAN Trunking or Private VLAN. * `None` - Do not tag traffic with any VLAN Id. * `VLAN` - Tag traffic with the Id from the VLAN Id field. * `VLAN trunking` - Pass VLAN traffic with Id within the VLAN trunk range to guest operating system. * `Private VLAN` - Associate the traffic with a private VLAN created on the distributed switch.
	VlanType          *string                                            `json:"VlanType,omitempty"`
	DistributedSwitch *VirtualizationVmwareDistributedSwitchRelationship `json:"DistributedSwitch,omitempty"`
	// An array of relationships to virtualizationVmwareHost resources.
	Hosts                []VirtualizationVmwareHostRelationship `json:"Hosts,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationVmwareDistributedNetwork VirtualizationVmwareDistributedNetwork

// NewVirtualizationVmwareDistributedNetwork instantiates a new VirtualizationVmwareDistributedNetwork object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationVmwareDistributedNetwork(classId string, objectType string) *VirtualizationVmwareDistributedNetwork {
	this := VirtualizationVmwareDistributedNetwork{}
	this.ClassId = classId
	this.ObjectType = objectType
	var forgedTransmits string = "Reject"
	this.ForgedTransmits = &forgedTransmits
	var macAddressChanges string = "Reject"
	this.MacAddressChanges = &macAddressChanges
	var promiscuousMode string = "Reject"
	this.PromiscuousMode = &promiscuousMode
	var vlanType string = "None"
	this.VlanType = &vlanType
	return &this
}

// NewVirtualizationVmwareDistributedNetworkWithDefaults instantiates a new VirtualizationVmwareDistributedNetwork object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationVmwareDistributedNetworkWithDefaults() *VirtualizationVmwareDistributedNetwork {
	this := VirtualizationVmwareDistributedNetwork{}
	var classId string = "virtualization.VmwareDistributedNetwork"
	this.ClassId = classId
	var objectType string = "virtualization.VmwareDistributedNetwork"
	this.ObjectType = objectType
	var forgedTransmits string = "Reject"
	this.ForgedTransmits = &forgedTransmits
	var macAddressChanges string = "Reject"
	this.MacAddressChanges = &macAddressChanges
	var promiscuousMode string = "Reject"
	this.PromiscuousMode = &promiscuousMode
	var vlanType string = "None"
	this.VlanType = &vlanType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationVmwareDistributedNetwork) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDistributedNetwork) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationVmwareDistributedNetwork) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationVmwareDistributedNetwork) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDistributedNetwork) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationVmwareDistributedNetwork) SetObjectType(v string) {
	o.ObjectType = v
}

// GetForgedTransmits returns the ForgedTransmits field value if set, zero value otherwise.
func (o *VirtualizationVmwareDistributedNetwork) GetForgedTransmits() string {
	if o == nil || o.ForgedTransmits == nil {
		var ret string
		return ret
	}
	return *o.ForgedTransmits
}

// GetForgedTransmitsOk returns a tuple with the ForgedTransmits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDistributedNetwork) GetForgedTransmitsOk() (*string, bool) {
	if o == nil || o.ForgedTransmits == nil {
		return nil, false
	}
	return o.ForgedTransmits, true
}

// HasForgedTransmits returns a boolean if a field has been set.
func (o *VirtualizationVmwareDistributedNetwork) HasForgedTransmits() bool {
	if o != nil && o.ForgedTransmits != nil {
		return true
	}

	return false
}

// SetForgedTransmits gets a reference to the given string and assigns it to the ForgedTransmits field.
func (o *VirtualizationVmwareDistributedNetwork) SetForgedTransmits(v string) {
	o.ForgedTransmits = &v
}

// GetMacAddressChanges returns the MacAddressChanges field value if set, zero value otherwise.
func (o *VirtualizationVmwareDistributedNetwork) GetMacAddressChanges() string {
	if o == nil || o.MacAddressChanges == nil {
		var ret string
		return ret
	}
	return *o.MacAddressChanges
}

// GetMacAddressChangesOk returns a tuple with the MacAddressChanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDistributedNetwork) GetMacAddressChangesOk() (*string, bool) {
	if o == nil || o.MacAddressChanges == nil {
		return nil, false
	}
	return o.MacAddressChanges, true
}

// HasMacAddressChanges returns a boolean if a field has been set.
func (o *VirtualizationVmwareDistributedNetwork) HasMacAddressChanges() bool {
	if o != nil && o.MacAddressChanges != nil {
		return true
	}

	return false
}

// SetMacAddressChanges gets a reference to the given string and assigns it to the MacAddressChanges field.
func (o *VirtualizationVmwareDistributedNetwork) SetMacAddressChanges(v string) {
	o.MacAddressChanges = &v
}

// GetNicTeamingAndFailover returns the NicTeamingAndFailover field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationVmwareDistributedNetwork) GetNicTeamingAndFailover() VirtualizationVmwareTeamingAndFailover {
	if o == nil || o.NicTeamingAndFailover.Get() == nil {
		var ret VirtualizationVmwareTeamingAndFailover
		return ret
	}
	return *o.NicTeamingAndFailover.Get()
}

// GetNicTeamingAndFailoverOk returns a tuple with the NicTeamingAndFailover field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationVmwareDistributedNetwork) GetNicTeamingAndFailoverOk() (*VirtualizationVmwareTeamingAndFailover, bool) {
	if o == nil {
		return nil, false
	}
	return o.NicTeamingAndFailover.Get(), o.NicTeamingAndFailover.IsSet()
}

// HasNicTeamingAndFailover returns a boolean if a field has been set.
func (o *VirtualizationVmwareDistributedNetwork) HasNicTeamingAndFailover() bool {
	if o != nil && o.NicTeamingAndFailover.IsSet() {
		return true
	}

	return false
}

// SetNicTeamingAndFailover gets a reference to the given NullableVirtualizationVmwareTeamingAndFailover and assigns it to the NicTeamingAndFailover field.
func (o *VirtualizationVmwareDistributedNetwork) SetNicTeamingAndFailover(v VirtualizationVmwareTeamingAndFailover) {
	o.NicTeamingAndFailover.Set(&v)
}

// SetNicTeamingAndFailoverNil sets the value for NicTeamingAndFailover to be an explicit nil
func (o *VirtualizationVmwareDistributedNetwork) SetNicTeamingAndFailoverNil() {
	o.NicTeamingAndFailover.Set(nil)
}

// UnsetNicTeamingAndFailover ensures that no value is present for NicTeamingAndFailover, not even an explicit nil
func (o *VirtualizationVmwareDistributedNetwork) UnsetNicTeamingAndFailover() {
	o.NicTeamingAndFailover.Unset()
}

// GetNumHosts returns the NumHosts field value if set, zero value otherwise.
func (o *VirtualizationVmwareDistributedNetwork) GetNumHosts() int64 {
	if o == nil || o.NumHosts == nil {
		var ret int64
		return ret
	}
	return *o.NumHosts
}

// GetNumHostsOk returns a tuple with the NumHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDistributedNetwork) GetNumHostsOk() (*int64, bool) {
	if o == nil || o.NumHosts == nil {
		return nil, false
	}
	return o.NumHosts, true
}

// HasNumHosts returns a boolean if a field has been set.
func (o *VirtualizationVmwareDistributedNetwork) HasNumHosts() bool {
	if o != nil && o.NumHosts != nil {
		return true
	}

	return false
}

// SetNumHosts gets a reference to the given int64 and assigns it to the NumHosts field.
func (o *VirtualizationVmwareDistributedNetwork) SetNumHosts(v int64) {
	o.NumHosts = &v
}

// GetNumPorts returns the NumPorts field value if set, zero value otherwise.
func (o *VirtualizationVmwareDistributedNetwork) GetNumPorts() int64 {
	if o == nil || o.NumPorts == nil {
		var ret int64
		return ret
	}
	return *o.NumPorts
}

// GetNumPortsOk returns a tuple with the NumPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDistributedNetwork) GetNumPortsOk() (*int64, bool) {
	if o == nil || o.NumPorts == nil {
		return nil, false
	}
	return o.NumPorts, true
}

// HasNumPorts returns a boolean if a field has been set.
func (o *VirtualizationVmwareDistributedNetwork) HasNumPorts() bool {
	if o != nil && o.NumPorts != nil {
		return true
	}

	return false
}

// SetNumPorts gets a reference to the given int64 and assigns it to the NumPorts field.
func (o *VirtualizationVmwareDistributedNetwork) SetNumPorts(v int64) {
	o.NumPorts = &v
}

// GetPromiscuousMode returns the PromiscuousMode field value if set, zero value otherwise.
func (o *VirtualizationVmwareDistributedNetwork) GetPromiscuousMode() string {
	if o == nil || o.PromiscuousMode == nil {
		var ret string
		return ret
	}
	return *o.PromiscuousMode
}

// GetPromiscuousModeOk returns a tuple with the PromiscuousMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDistributedNetwork) GetPromiscuousModeOk() (*string, bool) {
	if o == nil || o.PromiscuousMode == nil {
		return nil, false
	}
	return o.PromiscuousMode, true
}

// HasPromiscuousMode returns a boolean if a field has been set.
func (o *VirtualizationVmwareDistributedNetwork) HasPromiscuousMode() bool {
	if o != nil && o.PromiscuousMode != nil {
		return true
	}

	return false
}

// SetPromiscuousMode gets a reference to the given string and assigns it to the PromiscuousMode field.
func (o *VirtualizationVmwareDistributedNetwork) SetPromiscuousMode(v string) {
	o.PromiscuousMode = &v
}

// GetUpLink returns the UpLink field value if set, zero value otherwise.
func (o *VirtualizationVmwareDistributedNetwork) GetUpLink() bool {
	if o == nil || o.UpLink == nil {
		var ret bool
		return ret
	}
	return *o.UpLink
}

// GetUpLinkOk returns a tuple with the UpLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDistributedNetwork) GetUpLinkOk() (*bool, bool) {
	if o == nil || o.UpLink == nil {
		return nil, false
	}
	return o.UpLink, true
}

// HasUpLink returns a boolean if a field has been set.
func (o *VirtualizationVmwareDistributedNetwork) HasUpLink() bool {
	if o != nil && o.UpLink != nil {
		return true
	}

	return false
}

// SetUpLink gets a reference to the given bool and assigns it to the UpLink field.
func (o *VirtualizationVmwareDistributedNetwork) SetUpLink(v bool) {
	o.UpLink = &v
}

// GetVlanRange returns the VlanRange field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationVmwareDistributedNetwork) GetVlanRange() []VirtualizationVmwareVlanRange {
	if o == nil {
		var ret []VirtualizationVmwareVlanRange
		return ret
	}
	return o.VlanRange
}

// GetVlanRangeOk returns a tuple with the VlanRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationVmwareDistributedNetwork) GetVlanRangeOk() ([]VirtualizationVmwareVlanRange, bool) {
	if o == nil || o.VlanRange == nil {
		return nil, false
	}
	return o.VlanRange, true
}

// HasVlanRange returns a boolean if a field has been set.
func (o *VirtualizationVmwareDistributedNetwork) HasVlanRange() bool {
	if o != nil && o.VlanRange != nil {
		return true
	}

	return false
}

// SetVlanRange gets a reference to the given []VirtualizationVmwareVlanRange and assigns it to the VlanRange field.
func (o *VirtualizationVmwareDistributedNetwork) SetVlanRange(v []VirtualizationVmwareVlanRange) {
	o.VlanRange = v
}

// GetVlanType returns the VlanType field value if set, zero value otherwise.
func (o *VirtualizationVmwareDistributedNetwork) GetVlanType() string {
	if o == nil || o.VlanType == nil {
		var ret string
		return ret
	}
	return *o.VlanType
}

// GetVlanTypeOk returns a tuple with the VlanType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDistributedNetwork) GetVlanTypeOk() (*string, bool) {
	if o == nil || o.VlanType == nil {
		return nil, false
	}
	return o.VlanType, true
}

// HasVlanType returns a boolean if a field has been set.
func (o *VirtualizationVmwareDistributedNetwork) HasVlanType() bool {
	if o != nil && o.VlanType != nil {
		return true
	}

	return false
}

// SetVlanType gets a reference to the given string and assigns it to the VlanType field.
func (o *VirtualizationVmwareDistributedNetwork) SetVlanType(v string) {
	o.VlanType = &v
}

// GetDistributedSwitch returns the DistributedSwitch field value if set, zero value otherwise.
func (o *VirtualizationVmwareDistributedNetwork) GetDistributedSwitch() VirtualizationVmwareDistributedSwitchRelationship {
	if o == nil || o.DistributedSwitch == nil {
		var ret VirtualizationVmwareDistributedSwitchRelationship
		return ret
	}
	return *o.DistributedSwitch
}

// GetDistributedSwitchOk returns a tuple with the DistributedSwitch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDistributedNetwork) GetDistributedSwitchOk() (*VirtualizationVmwareDistributedSwitchRelationship, bool) {
	if o == nil || o.DistributedSwitch == nil {
		return nil, false
	}
	return o.DistributedSwitch, true
}

// HasDistributedSwitch returns a boolean if a field has been set.
func (o *VirtualizationVmwareDistributedNetwork) HasDistributedSwitch() bool {
	if o != nil && o.DistributedSwitch != nil {
		return true
	}

	return false
}

// SetDistributedSwitch gets a reference to the given VirtualizationVmwareDistributedSwitchRelationship and assigns it to the DistributedSwitch field.
func (o *VirtualizationVmwareDistributedNetwork) SetDistributedSwitch(v VirtualizationVmwareDistributedSwitchRelationship) {
	o.DistributedSwitch = &v
}

// GetHosts returns the Hosts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationVmwareDistributedNetwork) GetHosts() []VirtualizationVmwareHostRelationship {
	if o == nil {
		var ret []VirtualizationVmwareHostRelationship
		return ret
	}
	return o.Hosts
}

// GetHostsOk returns a tuple with the Hosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationVmwareDistributedNetwork) GetHostsOk() ([]VirtualizationVmwareHostRelationship, bool) {
	if o == nil || o.Hosts == nil {
		return nil, false
	}
	return o.Hosts, true
}

// HasHosts returns a boolean if a field has been set.
func (o *VirtualizationVmwareDistributedNetwork) HasHosts() bool {
	if o != nil && o.Hosts != nil {
		return true
	}

	return false
}

// SetHosts gets a reference to the given []VirtualizationVmwareHostRelationship and assigns it to the Hosts field.
func (o *VirtualizationVmwareDistributedNetwork) SetHosts(v []VirtualizationVmwareHostRelationship) {
	o.Hosts = v
}

func (o VirtualizationVmwareDistributedNetwork) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedVirtualizationBaseDistributedNetwork, errVirtualizationBaseDistributedNetwork := json.Marshal(o.VirtualizationBaseDistributedNetwork)
	if errVirtualizationBaseDistributedNetwork != nil {
		return []byte{}, errVirtualizationBaseDistributedNetwork
	}
	errVirtualizationBaseDistributedNetwork = json.Unmarshal([]byte(serializedVirtualizationBaseDistributedNetwork), &toSerialize)
	if errVirtualizationBaseDistributedNetwork != nil {
		return []byte{}, errVirtualizationBaseDistributedNetwork
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ForgedTransmits != nil {
		toSerialize["ForgedTransmits"] = o.ForgedTransmits
	}
	if o.MacAddressChanges != nil {
		toSerialize["MacAddressChanges"] = o.MacAddressChanges
	}
	if o.NicTeamingAndFailover.IsSet() {
		toSerialize["NicTeamingAndFailover"] = o.NicTeamingAndFailover.Get()
	}
	if o.NumHosts != nil {
		toSerialize["NumHosts"] = o.NumHosts
	}
	if o.NumPorts != nil {
		toSerialize["NumPorts"] = o.NumPorts
	}
	if o.PromiscuousMode != nil {
		toSerialize["PromiscuousMode"] = o.PromiscuousMode
	}
	if o.UpLink != nil {
		toSerialize["UpLink"] = o.UpLink
	}
	if o.VlanRange != nil {
		toSerialize["VlanRange"] = o.VlanRange
	}
	if o.VlanType != nil {
		toSerialize["VlanType"] = o.VlanType
	}
	if o.DistributedSwitch != nil {
		toSerialize["DistributedSwitch"] = o.DistributedSwitch
	}
	if o.Hosts != nil {
		toSerialize["Hosts"] = o.Hosts
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationVmwareDistributedNetwork) UnmarshalJSON(bytes []byte) (err error) {
	type VirtualizationVmwareDistributedNetworkWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// If forgedTransmits property value is set to reject, outbound frames with a source MAC address different from the one set on the adapter are dropped. If property value is set to accept, no filtering is performed and all outbound frames are passed. * `Reject` - Indicates that the security policy is rejected. * `Accept` - Indicates that the security policy is accepted.
		ForgedTransmits *string `json:"ForgedTransmits,omitempty"`
		// If macAddressChanges property value is set to reject and the MAC address of the adapter is changed to a value other than the one specified in .vmx configuration file, all inbound frames are dropped. If property value is set to accept and the MAC address is changed, inbound frames to the new MAC address are received. * `Reject` - Indicates that the security policy is rejected. * `Accept` - Indicates that the security policy is accepted.
		MacAddressChanges     *string                                        `json:"MacAddressChanges,omitempty"`
		NicTeamingAndFailover NullableVirtualizationVmwareTeamingAndFailover `json:"NicTeamingAndFailover,omitempty"`
		// The total number of hosts connected to this distributed virtual network.
		NumHosts *int64 `json:"NumHosts,omitempty"`
		// The total number of ports in the distributed virtual network.
		NumPorts *int64 `json:"NumPorts,omitempty"`
		// If promiscuousMode property value is set to reject, incoming traffic only targeted to that network will be visible. If property value is set to accept, objects defined within the network can see all incoming traffic on the virtual switch based on the VLAN policy. * `Reject` - Indicates that the security policy is rejected. * `Accept` - Indicates that the security policy is accepted.
		PromiscuousMode *string `json:"PromiscuousMode,omitempty"`
		// Indicates if the distributed virtual network is a uplink.
		UpLink    *bool                           `json:"UpLink,omitempty"`
		VlanRange []VirtualizationVmwareVlanRange `json:"VlanRange,omitempty"`
		// VLAN type of the distributed virtual network. It can be None, VLAN, VLAN Trunking or Private VLAN. * `None` - Do not tag traffic with any VLAN Id. * `VLAN` - Tag traffic with the Id from the VLAN Id field. * `VLAN trunking` - Pass VLAN traffic with Id within the VLAN trunk range to guest operating system. * `Private VLAN` - Associate the traffic with a private VLAN created on the distributed switch.
		VlanType          *string                                            `json:"VlanType,omitempty"`
		DistributedSwitch *VirtualizationVmwareDistributedSwitchRelationship `json:"DistributedSwitch,omitempty"`
		// An array of relationships to virtualizationVmwareHost resources.
		Hosts []VirtualizationVmwareHostRelationship `json:"Hosts,omitempty"`
	}

	varVirtualizationVmwareDistributedNetworkWithoutEmbeddedStruct := VirtualizationVmwareDistributedNetworkWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVirtualizationVmwareDistributedNetworkWithoutEmbeddedStruct)
	if err == nil {
		varVirtualizationVmwareDistributedNetwork := _VirtualizationVmwareDistributedNetwork{}
		varVirtualizationVmwareDistributedNetwork.ClassId = varVirtualizationVmwareDistributedNetworkWithoutEmbeddedStruct.ClassId
		varVirtualizationVmwareDistributedNetwork.ObjectType = varVirtualizationVmwareDistributedNetworkWithoutEmbeddedStruct.ObjectType
		varVirtualizationVmwareDistributedNetwork.ForgedTransmits = varVirtualizationVmwareDistributedNetworkWithoutEmbeddedStruct.ForgedTransmits
		varVirtualizationVmwareDistributedNetwork.MacAddressChanges = varVirtualizationVmwareDistributedNetworkWithoutEmbeddedStruct.MacAddressChanges
		varVirtualizationVmwareDistributedNetwork.NicTeamingAndFailover = varVirtualizationVmwareDistributedNetworkWithoutEmbeddedStruct.NicTeamingAndFailover
		varVirtualizationVmwareDistributedNetwork.NumHosts = varVirtualizationVmwareDistributedNetworkWithoutEmbeddedStruct.NumHosts
		varVirtualizationVmwareDistributedNetwork.NumPorts = varVirtualizationVmwareDistributedNetworkWithoutEmbeddedStruct.NumPorts
		varVirtualizationVmwareDistributedNetwork.PromiscuousMode = varVirtualizationVmwareDistributedNetworkWithoutEmbeddedStruct.PromiscuousMode
		varVirtualizationVmwareDistributedNetwork.UpLink = varVirtualizationVmwareDistributedNetworkWithoutEmbeddedStruct.UpLink
		varVirtualizationVmwareDistributedNetwork.VlanRange = varVirtualizationVmwareDistributedNetworkWithoutEmbeddedStruct.VlanRange
		varVirtualizationVmwareDistributedNetwork.VlanType = varVirtualizationVmwareDistributedNetworkWithoutEmbeddedStruct.VlanType
		varVirtualizationVmwareDistributedNetwork.DistributedSwitch = varVirtualizationVmwareDistributedNetworkWithoutEmbeddedStruct.DistributedSwitch
		varVirtualizationVmwareDistributedNetwork.Hosts = varVirtualizationVmwareDistributedNetworkWithoutEmbeddedStruct.Hosts
		*o = VirtualizationVmwareDistributedNetwork(varVirtualizationVmwareDistributedNetwork)
	} else {
		return err
	}

	varVirtualizationVmwareDistributedNetwork := _VirtualizationVmwareDistributedNetwork{}

	err = json.Unmarshal(bytes, &varVirtualizationVmwareDistributedNetwork)
	if err == nil {
		o.VirtualizationBaseDistributedNetwork = varVirtualizationVmwareDistributedNetwork.VirtualizationBaseDistributedNetwork
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ForgedTransmits")
		delete(additionalProperties, "MacAddressChanges")
		delete(additionalProperties, "NicTeamingAndFailover")
		delete(additionalProperties, "NumHosts")
		delete(additionalProperties, "NumPorts")
		delete(additionalProperties, "PromiscuousMode")
		delete(additionalProperties, "UpLink")
		delete(additionalProperties, "VlanRange")
		delete(additionalProperties, "VlanType")
		delete(additionalProperties, "DistributedSwitch")
		delete(additionalProperties, "Hosts")

		// remove fields from embedded structs
		reflectVirtualizationBaseDistributedNetwork := reflect.ValueOf(o.VirtualizationBaseDistributedNetwork)
		for i := 0; i < reflectVirtualizationBaseDistributedNetwork.Type().NumField(); i++ {
			t := reflectVirtualizationBaseDistributedNetwork.Type().Field(i)

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

type NullableVirtualizationVmwareDistributedNetwork struct {
	value *VirtualizationVmwareDistributedNetwork
	isSet bool
}

func (v NullableVirtualizationVmwareDistributedNetwork) Get() *VirtualizationVmwareDistributedNetwork {
	return v.value
}

func (v *NullableVirtualizationVmwareDistributedNetwork) Set(val *VirtualizationVmwareDistributedNetwork) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationVmwareDistributedNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationVmwareDistributedNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationVmwareDistributedNetwork(val *VirtualizationVmwareDistributedNetwork) *NullableVirtualizationVmwareDistributedNetwork {
	return &NullableVirtualizationVmwareDistributedNetwork{value: val, isSet: true}
}

func (v NullableVirtualizationVmwareDistributedNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationVmwareDistributedNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
