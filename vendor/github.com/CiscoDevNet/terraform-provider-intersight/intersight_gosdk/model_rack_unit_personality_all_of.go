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
)

// RackUnitPersonalityAllOf Definition of the list of properties defined in 'rack.UnitPersonality', excluding properties defined in parent classes.
type RackUnitPersonalityAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Additional info about the added software personality.
	AdditionalInfo *string `json:"AdditionalInfo,omitempty"`
	// Name of the software personality.
	Name *string `json:"Name,omitempty"`
	// Unique identity of added software personality.
	PersonalityId        *int64                               `json:"PersonalityId,omitempty"`
	ComputeRackUnit      *ComputeRackUnitRelationship         `json:"ComputeRackUnit,omitempty"`
	InventoryDeviceInfo  *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RackUnitPersonalityAllOf RackUnitPersonalityAllOf

// NewRackUnitPersonalityAllOf instantiates a new RackUnitPersonalityAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRackUnitPersonalityAllOf(classId string, objectType string) *RackUnitPersonalityAllOf {
	this := RackUnitPersonalityAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewRackUnitPersonalityAllOfWithDefaults instantiates a new RackUnitPersonalityAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRackUnitPersonalityAllOfWithDefaults() *RackUnitPersonalityAllOf {
	this := RackUnitPersonalityAllOf{}
	var classId string = "rack.UnitPersonality"
	this.ClassId = classId
	var objectType string = "rack.UnitPersonality"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *RackUnitPersonalityAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *RackUnitPersonalityAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *RackUnitPersonalityAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *RackUnitPersonalityAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *RackUnitPersonalityAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *RackUnitPersonalityAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAdditionalInfo returns the AdditionalInfo field value if set, zero value otherwise.
func (o *RackUnitPersonalityAllOf) GetAdditionalInfo() string {
	if o == nil || o.AdditionalInfo == nil {
		var ret string
		return ret
	}
	return *o.AdditionalInfo
}

// GetAdditionalInfoOk returns a tuple with the AdditionalInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RackUnitPersonalityAllOf) GetAdditionalInfoOk() (*string, bool) {
	if o == nil || o.AdditionalInfo == nil {
		return nil, false
	}
	return o.AdditionalInfo, true
}

// HasAdditionalInfo returns a boolean if a field has been set.
func (o *RackUnitPersonalityAllOf) HasAdditionalInfo() bool {
	if o != nil && o.AdditionalInfo != nil {
		return true
	}

	return false
}

// SetAdditionalInfo gets a reference to the given string and assigns it to the AdditionalInfo field.
func (o *RackUnitPersonalityAllOf) SetAdditionalInfo(v string) {
	o.AdditionalInfo = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RackUnitPersonalityAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RackUnitPersonalityAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RackUnitPersonalityAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RackUnitPersonalityAllOf) SetName(v string) {
	o.Name = &v
}

// GetPersonalityId returns the PersonalityId field value if set, zero value otherwise.
func (o *RackUnitPersonalityAllOf) GetPersonalityId() int64 {
	if o == nil || o.PersonalityId == nil {
		var ret int64
		return ret
	}
	return *o.PersonalityId
}

// GetPersonalityIdOk returns a tuple with the PersonalityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RackUnitPersonalityAllOf) GetPersonalityIdOk() (*int64, bool) {
	if o == nil || o.PersonalityId == nil {
		return nil, false
	}
	return o.PersonalityId, true
}

// HasPersonalityId returns a boolean if a field has been set.
func (o *RackUnitPersonalityAllOf) HasPersonalityId() bool {
	if o != nil && o.PersonalityId != nil {
		return true
	}

	return false
}

// SetPersonalityId gets a reference to the given int64 and assigns it to the PersonalityId field.
func (o *RackUnitPersonalityAllOf) SetPersonalityId(v int64) {
	o.PersonalityId = &v
}

// GetComputeRackUnit returns the ComputeRackUnit field value if set, zero value otherwise.
func (o *RackUnitPersonalityAllOf) GetComputeRackUnit() ComputeRackUnitRelationship {
	if o == nil || o.ComputeRackUnit == nil {
		var ret ComputeRackUnitRelationship
		return ret
	}
	return *o.ComputeRackUnit
}

// GetComputeRackUnitOk returns a tuple with the ComputeRackUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RackUnitPersonalityAllOf) GetComputeRackUnitOk() (*ComputeRackUnitRelationship, bool) {
	if o == nil || o.ComputeRackUnit == nil {
		return nil, false
	}
	return o.ComputeRackUnit, true
}

// HasComputeRackUnit returns a boolean if a field has been set.
func (o *RackUnitPersonalityAllOf) HasComputeRackUnit() bool {
	if o != nil && o.ComputeRackUnit != nil {
		return true
	}

	return false
}

// SetComputeRackUnit gets a reference to the given ComputeRackUnitRelationship and assigns it to the ComputeRackUnit field.
func (o *RackUnitPersonalityAllOf) SetComputeRackUnit(v ComputeRackUnitRelationship) {
	o.ComputeRackUnit = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *RackUnitPersonalityAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RackUnitPersonalityAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *RackUnitPersonalityAllOf) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *RackUnitPersonalityAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *RackUnitPersonalityAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RackUnitPersonalityAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *RackUnitPersonalityAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *RackUnitPersonalityAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o RackUnitPersonalityAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AdditionalInfo != nil {
		toSerialize["AdditionalInfo"] = o.AdditionalInfo
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.PersonalityId != nil {
		toSerialize["PersonalityId"] = o.PersonalityId
	}
	if o.ComputeRackUnit != nil {
		toSerialize["ComputeRackUnit"] = o.ComputeRackUnit
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RackUnitPersonalityAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varRackUnitPersonalityAllOf := _RackUnitPersonalityAllOf{}

	if err = json.Unmarshal(bytes, &varRackUnitPersonalityAllOf); err == nil {
		*o = RackUnitPersonalityAllOf(varRackUnitPersonalityAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AdditionalInfo")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "PersonalityId")
		delete(additionalProperties, "ComputeRackUnit")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRackUnitPersonalityAllOf struct {
	value *RackUnitPersonalityAllOf
	isSet bool
}

func (v NullableRackUnitPersonalityAllOf) Get() *RackUnitPersonalityAllOf {
	return v.value
}

func (v *NullableRackUnitPersonalityAllOf) Set(val *RackUnitPersonalityAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableRackUnitPersonalityAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableRackUnitPersonalityAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRackUnitPersonalityAllOf(val *RackUnitPersonalityAllOf) *NullableRackUnitPersonalityAllOf {
	return &NullableRackUnitPersonalityAllOf{value: val, isSet: true}
}

func (v NullableRackUnitPersonalityAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRackUnitPersonalityAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
