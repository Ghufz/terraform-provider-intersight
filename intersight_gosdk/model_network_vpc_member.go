/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-15711
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// NetworkVpcMember Concrete class for VPC configured on a network device.
type NetworkVpcMember struct {
	InventoryBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Operational state of the virtual port channel.
	OperationalState *string `json:"OperationalState,omitempty"`
	// Name of the virtual port channel.
	PortChannel *string `json:"PortChannel,omitempty"`
	// Port channel identity of the virtual port channel.
	PortChannelId *int64 `json:"PortChannelId,omitempty"`
	// Identity of the virtual port channel.
	VpcDomainId *int64 `json:"VpcDomainId,omitempty"`
	// Identity of the virtual port channel.
	VpcMemberId          *int64                               `json:"VpcMemberId,omitempty"`
	EtherPortChannel     *EtherPortChannelRelationship        `json:"EtherPortChannel,omitempty"`
	NetworkElement       *NetworkElementRelationship          `json:"NetworkElement,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NetworkVpcMember NetworkVpcMember

// NewNetworkVpcMember instantiates a new NetworkVpcMember object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkVpcMember(classId string, objectType string) *NetworkVpcMember {
	this := NetworkVpcMember{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNetworkVpcMemberWithDefaults instantiates a new NetworkVpcMember object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkVpcMemberWithDefaults() *NetworkVpcMember {
	this := NetworkVpcMember{}
	var classId string = "network.VpcMember"
	this.ClassId = classId
	var objectType string = "network.VpcMember"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NetworkVpcMember) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NetworkVpcMember) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NetworkVpcMember) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NetworkVpcMember) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NetworkVpcMember) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NetworkVpcMember) SetObjectType(v string) {
	o.ObjectType = v
}

// GetOperationalState returns the OperationalState field value if set, zero value otherwise.
func (o *NetworkVpcMember) GetOperationalState() string {
	if o == nil || o.OperationalState == nil {
		var ret string
		return ret
	}
	return *o.OperationalState
}

// GetOperationalStateOk returns a tuple with the OperationalState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkVpcMember) GetOperationalStateOk() (*string, bool) {
	if o == nil || o.OperationalState == nil {
		return nil, false
	}
	return o.OperationalState, true
}

// HasOperationalState returns a boolean if a field has been set.
func (o *NetworkVpcMember) HasOperationalState() bool {
	if o != nil && o.OperationalState != nil {
		return true
	}

	return false
}

// SetOperationalState gets a reference to the given string and assigns it to the OperationalState field.
func (o *NetworkVpcMember) SetOperationalState(v string) {
	o.OperationalState = &v
}

// GetPortChannel returns the PortChannel field value if set, zero value otherwise.
func (o *NetworkVpcMember) GetPortChannel() string {
	if o == nil || o.PortChannel == nil {
		var ret string
		return ret
	}
	return *o.PortChannel
}

// GetPortChannelOk returns a tuple with the PortChannel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkVpcMember) GetPortChannelOk() (*string, bool) {
	if o == nil || o.PortChannel == nil {
		return nil, false
	}
	return o.PortChannel, true
}

// HasPortChannel returns a boolean if a field has been set.
func (o *NetworkVpcMember) HasPortChannel() bool {
	if o != nil && o.PortChannel != nil {
		return true
	}

	return false
}

// SetPortChannel gets a reference to the given string and assigns it to the PortChannel field.
func (o *NetworkVpcMember) SetPortChannel(v string) {
	o.PortChannel = &v
}

// GetPortChannelId returns the PortChannelId field value if set, zero value otherwise.
func (o *NetworkVpcMember) GetPortChannelId() int64 {
	if o == nil || o.PortChannelId == nil {
		var ret int64
		return ret
	}
	return *o.PortChannelId
}

// GetPortChannelIdOk returns a tuple with the PortChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkVpcMember) GetPortChannelIdOk() (*int64, bool) {
	if o == nil || o.PortChannelId == nil {
		return nil, false
	}
	return o.PortChannelId, true
}

// HasPortChannelId returns a boolean if a field has been set.
func (o *NetworkVpcMember) HasPortChannelId() bool {
	if o != nil && o.PortChannelId != nil {
		return true
	}

	return false
}

// SetPortChannelId gets a reference to the given int64 and assigns it to the PortChannelId field.
func (o *NetworkVpcMember) SetPortChannelId(v int64) {
	o.PortChannelId = &v
}

// GetVpcDomainId returns the VpcDomainId field value if set, zero value otherwise.
func (o *NetworkVpcMember) GetVpcDomainId() int64 {
	if o == nil || o.VpcDomainId == nil {
		var ret int64
		return ret
	}
	return *o.VpcDomainId
}

// GetVpcDomainIdOk returns a tuple with the VpcDomainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkVpcMember) GetVpcDomainIdOk() (*int64, bool) {
	if o == nil || o.VpcDomainId == nil {
		return nil, false
	}
	return o.VpcDomainId, true
}

// HasVpcDomainId returns a boolean if a field has been set.
func (o *NetworkVpcMember) HasVpcDomainId() bool {
	if o != nil && o.VpcDomainId != nil {
		return true
	}

	return false
}

// SetVpcDomainId gets a reference to the given int64 and assigns it to the VpcDomainId field.
func (o *NetworkVpcMember) SetVpcDomainId(v int64) {
	o.VpcDomainId = &v
}

// GetVpcMemberId returns the VpcMemberId field value if set, zero value otherwise.
func (o *NetworkVpcMember) GetVpcMemberId() int64 {
	if o == nil || o.VpcMemberId == nil {
		var ret int64
		return ret
	}
	return *o.VpcMemberId
}

// GetVpcMemberIdOk returns a tuple with the VpcMemberId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkVpcMember) GetVpcMemberIdOk() (*int64, bool) {
	if o == nil || o.VpcMemberId == nil {
		return nil, false
	}
	return o.VpcMemberId, true
}

// HasVpcMemberId returns a boolean if a field has been set.
func (o *NetworkVpcMember) HasVpcMemberId() bool {
	if o != nil && o.VpcMemberId != nil {
		return true
	}

	return false
}

// SetVpcMemberId gets a reference to the given int64 and assigns it to the VpcMemberId field.
func (o *NetworkVpcMember) SetVpcMemberId(v int64) {
	o.VpcMemberId = &v
}

// GetEtherPortChannel returns the EtherPortChannel field value if set, zero value otherwise.
func (o *NetworkVpcMember) GetEtherPortChannel() EtherPortChannelRelationship {
	if o == nil || o.EtherPortChannel == nil {
		var ret EtherPortChannelRelationship
		return ret
	}
	return *o.EtherPortChannel
}

// GetEtherPortChannelOk returns a tuple with the EtherPortChannel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkVpcMember) GetEtherPortChannelOk() (*EtherPortChannelRelationship, bool) {
	if o == nil || o.EtherPortChannel == nil {
		return nil, false
	}
	return o.EtherPortChannel, true
}

// HasEtherPortChannel returns a boolean if a field has been set.
func (o *NetworkVpcMember) HasEtherPortChannel() bool {
	if o != nil && o.EtherPortChannel != nil {
		return true
	}

	return false
}

// SetEtherPortChannel gets a reference to the given EtherPortChannelRelationship and assigns it to the EtherPortChannel field.
func (o *NetworkVpcMember) SetEtherPortChannel(v EtherPortChannelRelationship) {
	o.EtherPortChannel = &v
}

// GetNetworkElement returns the NetworkElement field value if set, zero value otherwise.
func (o *NetworkVpcMember) GetNetworkElement() NetworkElementRelationship {
	if o == nil || o.NetworkElement == nil {
		var ret NetworkElementRelationship
		return ret
	}
	return *o.NetworkElement
}

// GetNetworkElementOk returns a tuple with the NetworkElement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkVpcMember) GetNetworkElementOk() (*NetworkElementRelationship, bool) {
	if o == nil || o.NetworkElement == nil {
		return nil, false
	}
	return o.NetworkElement, true
}

// HasNetworkElement returns a boolean if a field has been set.
func (o *NetworkVpcMember) HasNetworkElement() bool {
	if o != nil && o.NetworkElement != nil {
		return true
	}

	return false
}

// SetNetworkElement gets a reference to the given NetworkElementRelationship and assigns it to the NetworkElement field.
func (o *NetworkVpcMember) SetNetworkElement(v NetworkElementRelationship) {
	o.NetworkElement = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NetworkVpcMember) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkVpcMember) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NetworkVpcMember) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NetworkVpcMember) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NetworkVpcMember) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedInventoryBase, errInventoryBase := json.Marshal(o.InventoryBase)
	if errInventoryBase != nil {
		return []byte{}, errInventoryBase
	}
	errInventoryBase = json.Unmarshal([]byte(serializedInventoryBase), &toSerialize)
	if errInventoryBase != nil {
		return []byte{}, errInventoryBase
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.OperationalState != nil {
		toSerialize["OperationalState"] = o.OperationalState
	}
	if o.PortChannel != nil {
		toSerialize["PortChannel"] = o.PortChannel
	}
	if o.PortChannelId != nil {
		toSerialize["PortChannelId"] = o.PortChannelId
	}
	if o.VpcDomainId != nil {
		toSerialize["VpcDomainId"] = o.VpcDomainId
	}
	if o.VpcMemberId != nil {
		toSerialize["VpcMemberId"] = o.VpcMemberId
	}
	if o.EtherPortChannel != nil {
		toSerialize["EtherPortChannel"] = o.EtherPortChannel
	}
	if o.NetworkElement != nil {
		toSerialize["NetworkElement"] = o.NetworkElement
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NetworkVpcMember) UnmarshalJSON(bytes []byte) (err error) {
	type NetworkVpcMemberWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Operational state of the virtual port channel.
		OperationalState *string `json:"OperationalState,omitempty"`
		// Name of the virtual port channel.
		PortChannel *string `json:"PortChannel,omitempty"`
		// Port channel identity of the virtual port channel.
		PortChannelId *int64 `json:"PortChannelId,omitempty"`
		// Identity of the virtual port channel.
		VpcDomainId *int64 `json:"VpcDomainId,omitempty"`
		// Identity of the virtual port channel.
		VpcMemberId      *int64                               `json:"VpcMemberId,omitempty"`
		EtherPortChannel *EtherPortChannelRelationship        `json:"EtherPortChannel,omitempty"`
		NetworkElement   *NetworkElementRelationship          `json:"NetworkElement,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varNetworkVpcMemberWithoutEmbeddedStruct := NetworkVpcMemberWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNetworkVpcMemberWithoutEmbeddedStruct)
	if err == nil {
		varNetworkVpcMember := _NetworkVpcMember{}
		varNetworkVpcMember.ClassId = varNetworkVpcMemberWithoutEmbeddedStruct.ClassId
		varNetworkVpcMember.ObjectType = varNetworkVpcMemberWithoutEmbeddedStruct.ObjectType
		varNetworkVpcMember.OperationalState = varNetworkVpcMemberWithoutEmbeddedStruct.OperationalState
		varNetworkVpcMember.PortChannel = varNetworkVpcMemberWithoutEmbeddedStruct.PortChannel
		varNetworkVpcMember.PortChannelId = varNetworkVpcMemberWithoutEmbeddedStruct.PortChannelId
		varNetworkVpcMember.VpcDomainId = varNetworkVpcMemberWithoutEmbeddedStruct.VpcDomainId
		varNetworkVpcMember.VpcMemberId = varNetworkVpcMemberWithoutEmbeddedStruct.VpcMemberId
		varNetworkVpcMember.EtherPortChannel = varNetworkVpcMemberWithoutEmbeddedStruct.EtherPortChannel
		varNetworkVpcMember.NetworkElement = varNetworkVpcMemberWithoutEmbeddedStruct.NetworkElement
		varNetworkVpcMember.RegisteredDevice = varNetworkVpcMemberWithoutEmbeddedStruct.RegisteredDevice
		*o = NetworkVpcMember(varNetworkVpcMember)
	} else {
		return err
	}

	varNetworkVpcMember := _NetworkVpcMember{}

	err = json.Unmarshal(bytes, &varNetworkVpcMember)
	if err == nil {
		o.InventoryBase = varNetworkVpcMember.InventoryBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "OperationalState")
		delete(additionalProperties, "PortChannel")
		delete(additionalProperties, "PortChannelId")
		delete(additionalProperties, "VpcDomainId")
		delete(additionalProperties, "VpcMemberId")
		delete(additionalProperties, "EtherPortChannel")
		delete(additionalProperties, "NetworkElement")
		delete(additionalProperties, "RegisteredDevice")

		// remove fields from embedded structs
		reflectInventoryBase := reflect.ValueOf(o.InventoryBase)
		for i := 0; i < reflectInventoryBase.Type().NumField(); i++ {
			t := reflectInventoryBase.Type().Field(i)

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

type NullableNetworkVpcMember struct {
	value *NetworkVpcMember
	isSet bool
}

func (v NullableNetworkVpcMember) Get() *NetworkVpcMember {
	return v.value
}

func (v *NullableNetworkVpcMember) Set(val *NetworkVpcMember) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkVpcMember) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkVpcMember) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkVpcMember(val *NetworkVpcMember) *NullableNetworkVpcMember {
	return &NullableNetworkVpcMember{value: val, isSet: true}
}

func (v NullableNetworkVpcMember) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkVpcMember) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
