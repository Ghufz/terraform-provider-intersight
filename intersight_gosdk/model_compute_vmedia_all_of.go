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

// ComputeVmediaAllOf Definition of the list of properties defined in 'compute.Vmedia', excluding properties defined in parent classes.
type ComputeVmediaAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// State of the Virtual Media service on the server.
	Enabled *bool `json:"Enabled,omitempty"`
	// If enabled, allows encryption of all Virtual Media communications.
	Encryption *bool `json:"Encryption,omitempty"`
	// If enabled, the virtual drives appear on the boot selection menu after mapping the image and rebooting the host.
	LowPowerUsb         *bool                            `json:"LowPowerUsb,omitempty"`
	ComputePhysicalUnit *ComputePhysicalRelationship     `json:"ComputePhysicalUnit,omitempty"`
	InventoryDeviceInfo *InventoryDeviceInfoRelationship `json:"InventoryDeviceInfo,omitempty"`
	// An array of relationships to computeMapping resources.
	Mappings             []ComputeMappingRelationship         `json:"Mappings,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ComputeVmediaAllOf ComputeVmediaAllOf

// NewComputeVmediaAllOf instantiates a new ComputeVmediaAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputeVmediaAllOf(classId string, objectType string) *ComputeVmediaAllOf {
	this := ComputeVmediaAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewComputeVmediaAllOfWithDefaults instantiates a new ComputeVmediaAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputeVmediaAllOfWithDefaults() *ComputeVmediaAllOf {
	this := ComputeVmediaAllOf{}
	var classId string = "compute.Vmedia"
	this.ClassId = classId
	var objectType string = "compute.Vmedia"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ComputeVmediaAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ComputeVmediaAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ComputeVmediaAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ComputeVmediaAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ComputeVmediaAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ComputeVmediaAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ComputeVmediaAllOf) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeVmediaAllOf) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ComputeVmediaAllOf) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ComputeVmediaAllOf) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetEncryption returns the Encryption field value if set, zero value otherwise.
func (o *ComputeVmediaAllOf) GetEncryption() bool {
	if o == nil || o.Encryption == nil {
		var ret bool
		return ret
	}
	return *o.Encryption
}

// GetEncryptionOk returns a tuple with the Encryption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeVmediaAllOf) GetEncryptionOk() (*bool, bool) {
	if o == nil || o.Encryption == nil {
		return nil, false
	}
	return o.Encryption, true
}

// HasEncryption returns a boolean if a field has been set.
func (o *ComputeVmediaAllOf) HasEncryption() bool {
	if o != nil && o.Encryption != nil {
		return true
	}

	return false
}

// SetEncryption gets a reference to the given bool and assigns it to the Encryption field.
func (o *ComputeVmediaAllOf) SetEncryption(v bool) {
	o.Encryption = &v
}

// GetLowPowerUsb returns the LowPowerUsb field value if set, zero value otherwise.
func (o *ComputeVmediaAllOf) GetLowPowerUsb() bool {
	if o == nil || o.LowPowerUsb == nil {
		var ret bool
		return ret
	}
	return *o.LowPowerUsb
}

// GetLowPowerUsbOk returns a tuple with the LowPowerUsb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeVmediaAllOf) GetLowPowerUsbOk() (*bool, bool) {
	if o == nil || o.LowPowerUsb == nil {
		return nil, false
	}
	return o.LowPowerUsb, true
}

// HasLowPowerUsb returns a boolean if a field has been set.
func (o *ComputeVmediaAllOf) HasLowPowerUsb() bool {
	if o != nil && o.LowPowerUsb != nil {
		return true
	}

	return false
}

// SetLowPowerUsb gets a reference to the given bool and assigns it to the LowPowerUsb field.
func (o *ComputeVmediaAllOf) SetLowPowerUsb(v bool) {
	o.LowPowerUsb = &v
}

// GetComputePhysicalUnit returns the ComputePhysicalUnit field value if set, zero value otherwise.
func (o *ComputeVmediaAllOf) GetComputePhysicalUnit() ComputePhysicalRelationship {
	if o == nil || o.ComputePhysicalUnit == nil {
		var ret ComputePhysicalRelationship
		return ret
	}
	return *o.ComputePhysicalUnit
}

// GetComputePhysicalUnitOk returns a tuple with the ComputePhysicalUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeVmediaAllOf) GetComputePhysicalUnitOk() (*ComputePhysicalRelationship, bool) {
	if o == nil || o.ComputePhysicalUnit == nil {
		return nil, false
	}
	return o.ComputePhysicalUnit, true
}

// HasComputePhysicalUnit returns a boolean if a field has been set.
func (o *ComputeVmediaAllOf) HasComputePhysicalUnit() bool {
	if o != nil && o.ComputePhysicalUnit != nil {
		return true
	}

	return false
}

// SetComputePhysicalUnit gets a reference to the given ComputePhysicalRelationship and assigns it to the ComputePhysicalUnit field.
func (o *ComputeVmediaAllOf) SetComputePhysicalUnit(v ComputePhysicalRelationship) {
	o.ComputePhysicalUnit = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *ComputeVmediaAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeVmediaAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *ComputeVmediaAllOf) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *ComputeVmediaAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetMappings returns the Mappings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComputeVmediaAllOf) GetMappings() []ComputeMappingRelationship {
	if o == nil {
		var ret []ComputeMappingRelationship
		return ret
	}
	return o.Mappings
}

// GetMappingsOk returns a tuple with the Mappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComputeVmediaAllOf) GetMappingsOk() ([]ComputeMappingRelationship, bool) {
	if o == nil || o.Mappings == nil {
		return nil, false
	}
	return o.Mappings, true
}

// HasMappings returns a boolean if a field has been set.
func (o *ComputeVmediaAllOf) HasMappings() bool {
	if o != nil && o.Mappings != nil {
		return true
	}

	return false
}

// SetMappings gets a reference to the given []ComputeMappingRelationship and assigns it to the Mappings field.
func (o *ComputeVmediaAllOf) SetMappings(v []ComputeMappingRelationship) {
	o.Mappings = v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *ComputeVmediaAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeVmediaAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *ComputeVmediaAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *ComputeVmediaAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o ComputeVmediaAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Enabled != nil {
		toSerialize["Enabled"] = o.Enabled
	}
	if o.Encryption != nil {
		toSerialize["Encryption"] = o.Encryption
	}
	if o.LowPowerUsb != nil {
		toSerialize["LowPowerUsb"] = o.LowPowerUsb
	}
	if o.ComputePhysicalUnit != nil {
		toSerialize["ComputePhysicalUnit"] = o.ComputePhysicalUnit
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.Mappings != nil {
		toSerialize["Mappings"] = o.Mappings
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ComputeVmediaAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varComputeVmediaAllOf := _ComputeVmediaAllOf{}

	if err = json.Unmarshal(bytes, &varComputeVmediaAllOf); err == nil {
		*o = ComputeVmediaAllOf(varComputeVmediaAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Enabled")
		delete(additionalProperties, "Encryption")
		delete(additionalProperties, "LowPowerUsb")
		delete(additionalProperties, "ComputePhysicalUnit")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "Mappings")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableComputeVmediaAllOf struct {
	value *ComputeVmediaAllOf
	isSet bool
}

func (v NullableComputeVmediaAllOf) Get() *ComputeVmediaAllOf {
	return v.value
}

func (v *NullableComputeVmediaAllOf) Set(val *ComputeVmediaAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableComputeVmediaAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableComputeVmediaAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputeVmediaAllOf(val *ComputeVmediaAllOf) *NullableComputeVmediaAllOf {
	return &NullableComputeVmediaAllOf{value: val, isSet: true}
}

func (v NullableComputeVmediaAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputeVmediaAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
