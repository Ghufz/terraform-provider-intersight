/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-11765
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// ConsoleConsoleConfigAllOf Definition of the list of properties defined in 'console.ConsoleConfig', excluding properties defined in parent classes.
type ConsoleConsoleConfigAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The databits numbers in console.
	DataBits *int64 `json:"DataBits,omitempty"`
	// The parity bit of the console. * `none` - If Parity is none, parity checking is not performed and the parity bit is not transmitted. * `odd` - If Parity is odd, the number of mark bits (1s) in the data is counted, and the parity bit is asserted or unasserted to obtain an odd number of mark bits. * `even` - If Parity is even, the number of mark bits in the data is counted, and the parity bit is asserted or unasserted to obtain an even number of mark bits. * `mark` - If Parity is mark, the parity bit is asserted. * `space` - If Parity is space, the parity bit is unasserted.
	Parity *string `json:"Parity,omitempty"`
	// The speed of the console.
	Speed *int64 `json:"Speed,omitempty"`
	// The stopbits of async line in console.
	StopBits             *int64                               `json:"StopBits,omitempty"`
	NetworkElement       *NetworkElementRelationship          `json:"NetworkElement,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConsoleConsoleConfigAllOf ConsoleConsoleConfigAllOf

// NewConsoleConsoleConfigAllOf instantiates a new ConsoleConsoleConfigAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsoleConsoleConfigAllOf(classId string, objectType string) *ConsoleConsoleConfigAllOf {
	this := ConsoleConsoleConfigAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var parity string = "none"
	this.Parity = &parity
	return &this
}

// NewConsoleConsoleConfigAllOfWithDefaults instantiates a new ConsoleConsoleConfigAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsoleConsoleConfigAllOfWithDefaults() *ConsoleConsoleConfigAllOf {
	this := ConsoleConsoleConfigAllOf{}
	var classId string = "console.ConsoleConfig"
	this.ClassId = classId
	var objectType string = "console.ConsoleConfig"
	this.ObjectType = objectType
	var parity string = "none"
	this.Parity = &parity
	return &this
}

// GetClassId returns the ClassId field value
func (o *ConsoleConsoleConfigAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ConsoleConsoleConfigAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ConsoleConsoleConfigAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ConsoleConsoleConfigAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ConsoleConsoleConfigAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ConsoleConsoleConfigAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDataBits returns the DataBits field value if set, zero value otherwise.
func (o *ConsoleConsoleConfigAllOf) GetDataBits() int64 {
	if o == nil || o.DataBits == nil {
		var ret int64
		return ret
	}
	return *o.DataBits
}

// GetDataBitsOk returns a tuple with the DataBits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsoleConsoleConfigAllOf) GetDataBitsOk() (*int64, bool) {
	if o == nil || o.DataBits == nil {
		return nil, false
	}
	return o.DataBits, true
}

// HasDataBits returns a boolean if a field has been set.
func (o *ConsoleConsoleConfigAllOf) HasDataBits() bool {
	if o != nil && o.DataBits != nil {
		return true
	}

	return false
}

// SetDataBits gets a reference to the given int64 and assigns it to the DataBits field.
func (o *ConsoleConsoleConfigAllOf) SetDataBits(v int64) {
	o.DataBits = &v
}

// GetParity returns the Parity field value if set, zero value otherwise.
func (o *ConsoleConsoleConfigAllOf) GetParity() string {
	if o == nil || o.Parity == nil {
		var ret string
		return ret
	}
	return *o.Parity
}

// GetParityOk returns a tuple with the Parity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsoleConsoleConfigAllOf) GetParityOk() (*string, bool) {
	if o == nil || o.Parity == nil {
		return nil, false
	}
	return o.Parity, true
}

// HasParity returns a boolean if a field has been set.
func (o *ConsoleConsoleConfigAllOf) HasParity() bool {
	if o != nil && o.Parity != nil {
		return true
	}

	return false
}

// SetParity gets a reference to the given string and assigns it to the Parity field.
func (o *ConsoleConsoleConfigAllOf) SetParity(v string) {
	o.Parity = &v
}

// GetSpeed returns the Speed field value if set, zero value otherwise.
func (o *ConsoleConsoleConfigAllOf) GetSpeed() int64 {
	if o == nil || o.Speed == nil {
		var ret int64
		return ret
	}
	return *o.Speed
}

// GetSpeedOk returns a tuple with the Speed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsoleConsoleConfigAllOf) GetSpeedOk() (*int64, bool) {
	if o == nil || o.Speed == nil {
		return nil, false
	}
	return o.Speed, true
}

// HasSpeed returns a boolean if a field has been set.
func (o *ConsoleConsoleConfigAllOf) HasSpeed() bool {
	if o != nil && o.Speed != nil {
		return true
	}

	return false
}

// SetSpeed gets a reference to the given int64 and assigns it to the Speed field.
func (o *ConsoleConsoleConfigAllOf) SetSpeed(v int64) {
	o.Speed = &v
}

// GetStopBits returns the StopBits field value if set, zero value otherwise.
func (o *ConsoleConsoleConfigAllOf) GetStopBits() int64 {
	if o == nil || o.StopBits == nil {
		var ret int64
		return ret
	}
	return *o.StopBits
}

// GetStopBitsOk returns a tuple with the StopBits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsoleConsoleConfigAllOf) GetStopBitsOk() (*int64, bool) {
	if o == nil || o.StopBits == nil {
		return nil, false
	}
	return o.StopBits, true
}

// HasStopBits returns a boolean if a field has been set.
func (o *ConsoleConsoleConfigAllOf) HasStopBits() bool {
	if o != nil && o.StopBits != nil {
		return true
	}

	return false
}

// SetStopBits gets a reference to the given int64 and assigns it to the StopBits field.
func (o *ConsoleConsoleConfigAllOf) SetStopBits(v int64) {
	o.StopBits = &v
}

// GetNetworkElement returns the NetworkElement field value if set, zero value otherwise.
func (o *ConsoleConsoleConfigAllOf) GetNetworkElement() NetworkElementRelationship {
	if o == nil || o.NetworkElement == nil {
		var ret NetworkElementRelationship
		return ret
	}
	return *o.NetworkElement
}

// GetNetworkElementOk returns a tuple with the NetworkElement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsoleConsoleConfigAllOf) GetNetworkElementOk() (*NetworkElementRelationship, bool) {
	if o == nil || o.NetworkElement == nil {
		return nil, false
	}
	return o.NetworkElement, true
}

// HasNetworkElement returns a boolean if a field has been set.
func (o *ConsoleConsoleConfigAllOf) HasNetworkElement() bool {
	if o != nil && o.NetworkElement != nil {
		return true
	}

	return false
}

// SetNetworkElement gets a reference to the given NetworkElementRelationship and assigns it to the NetworkElement field.
func (o *ConsoleConsoleConfigAllOf) SetNetworkElement(v NetworkElementRelationship) {
	o.NetworkElement = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *ConsoleConsoleConfigAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsoleConsoleConfigAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *ConsoleConsoleConfigAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *ConsoleConsoleConfigAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o ConsoleConsoleConfigAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DataBits != nil {
		toSerialize["DataBits"] = o.DataBits
	}
	if o.Parity != nil {
		toSerialize["Parity"] = o.Parity
	}
	if o.Speed != nil {
		toSerialize["Speed"] = o.Speed
	}
	if o.StopBits != nil {
		toSerialize["StopBits"] = o.StopBits
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

func (o *ConsoleConsoleConfigAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varConsoleConsoleConfigAllOf := _ConsoleConsoleConfigAllOf{}

	if err = json.Unmarshal(bytes, &varConsoleConsoleConfigAllOf); err == nil {
		*o = ConsoleConsoleConfigAllOf(varConsoleConsoleConfigAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DataBits")
		delete(additionalProperties, "Parity")
		delete(additionalProperties, "Speed")
		delete(additionalProperties, "StopBits")
		delete(additionalProperties, "NetworkElement")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConsoleConsoleConfigAllOf struct {
	value *ConsoleConsoleConfigAllOf
	isSet bool
}

func (v NullableConsoleConsoleConfigAllOf) Get() *ConsoleConsoleConfigAllOf {
	return v.value
}

func (v *NullableConsoleConsoleConfigAllOf) Set(val *ConsoleConsoleConfigAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableConsoleConsoleConfigAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableConsoleConsoleConfigAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsoleConsoleConfigAllOf(val *ConsoleConsoleConfigAllOf) *NullableConsoleConsoleConfigAllOf {
	return &NullableConsoleConsoleConfigAllOf{value: val, isSet: true}
}

func (v NullableConsoleConsoleConfigAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsoleConsoleConfigAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
