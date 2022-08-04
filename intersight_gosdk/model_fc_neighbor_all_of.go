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
)

// FcNeighborAllOf Definition of the list of properties defined in 'fc.Neighbor', excluding properties defined in parent classes.
type FcNeighborAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// This field defines if neighbor is a switch or an NPV device. * `Switch` - Switch type neighbors of an interface. * `NPV` - N Port Virtualization neighbors of an interface.
	PeerDeviceCapability *string `json:"PeerDeviceCapability,omitempty"`
	// Interface through which the relationship is established.
	PeerInterface *string `json:"PeerInterface,omitempty"`
	// IP address of the peer switch.
	PeerIpAddress *string `json:"PeerIpAddress,omitempty"`
	// Device Id of the neighbor switch.
	PeerSwitchName *string `json:"PeerSwitchName,omitempty"`
	// World Wide Name of the neighbor switch.
	PeerWwn              *string                     `json:"PeerWwn,omitempty"`
	FcPhysicalPort       *FcPhysicalPortRelationship `json:"FcPhysicalPort,omitempty"`
	FcPortChannel        *FcPortChannelRelationship  `json:"FcPortChannel,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FcNeighborAllOf FcNeighborAllOf

// NewFcNeighborAllOf instantiates a new FcNeighborAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFcNeighborAllOf(classId string, objectType string) *FcNeighborAllOf {
	this := FcNeighborAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewFcNeighborAllOfWithDefaults instantiates a new FcNeighborAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFcNeighborAllOfWithDefaults() *FcNeighborAllOf {
	this := FcNeighborAllOf{}
	var classId string = "fc.Neighbor"
	this.ClassId = classId
	var objectType string = "fc.Neighbor"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *FcNeighborAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FcNeighborAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FcNeighborAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FcNeighborAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FcNeighborAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FcNeighborAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetPeerDeviceCapability returns the PeerDeviceCapability field value if set, zero value otherwise.
func (o *FcNeighborAllOf) GetPeerDeviceCapability() string {
	if o == nil || o.PeerDeviceCapability == nil {
		var ret string
		return ret
	}
	return *o.PeerDeviceCapability
}

// GetPeerDeviceCapabilityOk returns a tuple with the PeerDeviceCapability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcNeighborAllOf) GetPeerDeviceCapabilityOk() (*string, bool) {
	if o == nil || o.PeerDeviceCapability == nil {
		return nil, false
	}
	return o.PeerDeviceCapability, true
}

// HasPeerDeviceCapability returns a boolean if a field has been set.
func (o *FcNeighborAllOf) HasPeerDeviceCapability() bool {
	if o != nil && o.PeerDeviceCapability != nil {
		return true
	}

	return false
}

// SetPeerDeviceCapability gets a reference to the given string and assigns it to the PeerDeviceCapability field.
func (o *FcNeighborAllOf) SetPeerDeviceCapability(v string) {
	o.PeerDeviceCapability = &v
}

// GetPeerInterface returns the PeerInterface field value if set, zero value otherwise.
func (o *FcNeighborAllOf) GetPeerInterface() string {
	if o == nil || o.PeerInterface == nil {
		var ret string
		return ret
	}
	return *o.PeerInterface
}

// GetPeerInterfaceOk returns a tuple with the PeerInterface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcNeighborAllOf) GetPeerInterfaceOk() (*string, bool) {
	if o == nil || o.PeerInterface == nil {
		return nil, false
	}
	return o.PeerInterface, true
}

// HasPeerInterface returns a boolean if a field has been set.
func (o *FcNeighborAllOf) HasPeerInterface() bool {
	if o != nil && o.PeerInterface != nil {
		return true
	}

	return false
}

// SetPeerInterface gets a reference to the given string and assigns it to the PeerInterface field.
func (o *FcNeighborAllOf) SetPeerInterface(v string) {
	o.PeerInterface = &v
}

// GetPeerIpAddress returns the PeerIpAddress field value if set, zero value otherwise.
func (o *FcNeighborAllOf) GetPeerIpAddress() string {
	if o == nil || o.PeerIpAddress == nil {
		var ret string
		return ret
	}
	return *o.PeerIpAddress
}

// GetPeerIpAddressOk returns a tuple with the PeerIpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcNeighborAllOf) GetPeerIpAddressOk() (*string, bool) {
	if o == nil || o.PeerIpAddress == nil {
		return nil, false
	}
	return o.PeerIpAddress, true
}

// HasPeerIpAddress returns a boolean if a field has been set.
func (o *FcNeighborAllOf) HasPeerIpAddress() bool {
	if o != nil && o.PeerIpAddress != nil {
		return true
	}

	return false
}

// SetPeerIpAddress gets a reference to the given string and assigns it to the PeerIpAddress field.
func (o *FcNeighborAllOf) SetPeerIpAddress(v string) {
	o.PeerIpAddress = &v
}

// GetPeerSwitchName returns the PeerSwitchName field value if set, zero value otherwise.
func (o *FcNeighborAllOf) GetPeerSwitchName() string {
	if o == nil || o.PeerSwitchName == nil {
		var ret string
		return ret
	}
	return *o.PeerSwitchName
}

// GetPeerSwitchNameOk returns a tuple with the PeerSwitchName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcNeighborAllOf) GetPeerSwitchNameOk() (*string, bool) {
	if o == nil || o.PeerSwitchName == nil {
		return nil, false
	}
	return o.PeerSwitchName, true
}

// HasPeerSwitchName returns a boolean if a field has been set.
func (o *FcNeighborAllOf) HasPeerSwitchName() bool {
	if o != nil && o.PeerSwitchName != nil {
		return true
	}

	return false
}

// SetPeerSwitchName gets a reference to the given string and assigns it to the PeerSwitchName field.
func (o *FcNeighborAllOf) SetPeerSwitchName(v string) {
	o.PeerSwitchName = &v
}

// GetPeerWwn returns the PeerWwn field value if set, zero value otherwise.
func (o *FcNeighborAllOf) GetPeerWwn() string {
	if o == nil || o.PeerWwn == nil {
		var ret string
		return ret
	}
	return *o.PeerWwn
}

// GetPeerWwnOk returns a tuple with the PeerWwn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcNeighborAllOf) GetPeerWwnOk() (*string, bool) {
	if o == nil || o.PeerWwn == nil {
		return nil, false
	}
	return o.PeerWwn, true
}

// HasPeerWwn returns a boolean if a field has been set.
func (o *FcNeighborAllOf) HasPeerWwn() bool {
	if o != nil && o.PeerWwn != nil {
		return true
	}

	return false
}

// SetPeerWwn gets a reference to the given string and assigns it to the PeerWwn field.
func (o *FcNeighborAllOf) SetPeerWwn(v string) {
	o.PeerWwn = &v
}

// GetFcPhysicalPort returns the FcPhysicalPort field value if set, zero value otherwise.
func (o *FcNeighborAllOf) GetFcPhysicalPort() FcPhysicalPortRelationship {
	if o == nil || o.FcPhysicalPort == nil {
		var ret FcPhysicalPortRelationship
		return ret
	}
	return *o.FcPhysicalPort
}

// GetFcPhysicalPortOk returns a tuple with the FcPhysicalPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcNeighborAllOf) GetFcPhysicalPortOk() (*FcPhysicalPortRelationship, bool) {
	if o == nil || o.FcPhysicalPort == nil {
		return nil, false
	}
	return o.FcPhysicalPort, true
}

// HasFcPhysicalPort returns a boolean if a field has been set.
func (o *FcNeighborAllOf) HasFcPhysicalPort() bool {
	if o != nil && o.FcPhysicalPort != nil {
		return true
	}

	return false
}

// SetFcPhysicalPort gets a reference to the given FcPhysicalPortRelationship and assigns it to the FcPhysicalPort field.
func (o *FcNeighborAllOf) SetFcPhysicalPort(v FcPhysicalPortRelationship) {
	o.FcPhysicalPort = &v
}

// GetFcPortChannel returns the FcPortChannel field value if set, zero value otherwise.
func (o *FcNeighborAllOf) GetFcPortChannel() FcPortChannelRelationship {
	if o == nil || o.FcPortChannel == nil {
		var ret FcPortChannelRelationship
		return ret
	}
	return *o.FcPortChannel
}

// GetFcPortChannelOk returns a tuple with the FcPortChannel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcNeighborAllOf) GetFcPortChannelOk() (*FcPortChannelRelationship, bool) {
	if o == nil || o.FcPortChannel == nil {
		return nil, false
	}
	return o.FcPortChannel, true
}

// HasFcPortChannel returns a boolean if a field has been set.
func (o *FcNeighborAllOf) HasFcPortChannel() bool {
	if o != nil && o.FcPortChannel != nil {
		return true
	}

	return false
}

// SetFcPortChannel gets a reference to the given FcPortChannelRelationship and assigns it to the FcPortChannel field.
func (o *FcNeighborAllOf) SetFcPortChannel(v FcPortChannelRelationship) {
	o.FcPortChannel = &v
}

func (o FcNeighborAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.PeerDeviceCapability != nil {
		toSerialize["PeerDeviceCapability"] = o.PeerDeviceCapability
	}
	if o.PeerInterface != nil {
		toSerialize["PeerInterface"] = o.PeerInterface
	}
	if o.PeerIpAddress != nil {
		toSerialize["PeerIpAddress"] = o.PeerIpAddress
	}
	if o.PeerSwitchName != nil {
		toSerialize["PeerSwitchName"] = o.PeerSwitchName
	}
	if o.PeerWwn != nil {
		toSerialize["PeerWwn"] = o.PeerWwn
	}
	if o.FcPhysicalPort != nil {
		toSerialize["FcPhysicalPort"] = o.FcPhysicalPort
	}
	if o.FcPortChannel != nil {
		toSerialize["FcPortChannel"] = o.FcPortChannel
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FcNeighborAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varFcNeighborAllOf := _FcNeighborAllOf{}

	if err = json.Unmarshal(bytes, &varFcNeighborAllOf); err == nil {
		*o = FcNeighborAllOf(varFcNeighborAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "PeerDeviceCapability")
		delete(additionalProperties, "PeerInterface")
		delete(additionalProperties, "PeerIpAddress")
		delete(additionalProperties, "PeerSwitchName")
		delete(additionalProperties, "PeerWwn")
		delete(additionalProperties, "FcPhysicalPort")
		delete(additionalProperties, "FcPortChannel")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFcNeighborAllOf struct {
	value *FcNeighborAllOf
	isSet bool
}

func (v NullableFcNeighborAllOf) Get() *FcNeighborAllOf {
	return v.value
}

func (v *NullableFcNeighborAllOf) Set(val *FcNeighborAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFcNeighborAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFcNeighborAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFcNeighborAllOf(val *FcNeighborAllOf) *NullableFcNeighborAllOf {
	return &NullableFcNeighborAllOf{value: val, isSet: true}
}

func (v NullableFcNeighborAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFcNeighborAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
