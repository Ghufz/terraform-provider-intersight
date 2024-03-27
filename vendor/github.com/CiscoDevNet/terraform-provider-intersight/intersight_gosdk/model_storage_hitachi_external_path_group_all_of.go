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

// StorageHitachiExternalPathGroupAllOf Definition of the list of properties defined in 'storage.HitachiExternalPathGroup', excluding properties defined in parent classes.
type StorageHitachiExternalPathGroupAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                       `json:"ObjectType"`
	ExternalParityGroups []StorageExternalParityGroup `json:"ExternalParityGroups,omitempty"`
	ExternalPaths        []StorageExternalPath        `json:"ExternalPaths,omitempty"`
	// Product ID of the external storage system.
	ExternalProductId *string `json:"ExternalProductId,omitempty"`
	// Serial number of the external storage system.
	ExternalSerialNumber *string `json:"ExternalSerialNumber,omitempty"`
	// External path group number.
	Name                 *string                              `json:"Name,omitempty"`
	Array                *StorageHitachiArrayRelationship     `json:"Array,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageHitachiExternalPathGroupAllOf StorageHitachiExternalPathGroupAllOf

// NewStorageHitachiExternalPathGroupAllOf instantiates a new StorageHitachiExternalPathGroupAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageHitachiExternalPathGroupAllOf(classId string, objectType string) *StorageHitachiExternalPathGroupAllOf {
	this := StorageHitachiExternalPathGroupAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageHitachiExternalPathGroupAllOfWithDefaults instantiates a new StorageHitachiExternalPathGroupAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageHitachiExternalPathGroupAllOfWithDefaults() *StorageHitachiExternalPathGroupAllOf {
	this := StorageHitachiExternalPathGroupAllOf{}
	var classId string = "storage.HitachiExternalPathGroup"
	this.ClassId = classId
	var objectType string = "storage.HitachiExternalPathGroup"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageHitachiExternalPathGroupAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageHitachiExternalPathGroupAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageHitachiExternalPathGroupAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageHitachiExternalPathGroupAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageHitachiExternalPathGroupAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageHitachiExternalPathGroupAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetExternalParityGroups returns the ExternalParityGroups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageHitachiExternalPathGroupAllOf) GetExternalParityGroups() []StorageExternalParityGroup {
	if o == nil {
		var ret []StorageExternalParityGroup
		return ret
	}
	return o.ExternalParityGroups
}

// GetExternalParityGroupsOk returns a tuple with the ExternalParityGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageHitachiExternalPathGroupAllOf) GetExternalParityGroupsOk() ([]StorageExternalParityGroup, bool) {
	if o == nil || o.ExternalParityGroups == nil {
		return nil, false
	}
	return o.ExternalParityGroups, true
}

// HasExternalParityGroups returns a boolean if a field has been set.
func (o *StorageHitachiExternalPathGroupAllOf) HasExternalParityGroups() bool {
	if o != nil && o.ExternalParityGroups != nil {
		return true
	}

	return false
}

// SetExternalParityGroups gets a reference to the given []StorageExternalParityGroup and assigns it to the ExternalParityGroups field.
func (o *StorageHitachiExternalPathGroupAllOf) SetExternalParityGroups(v []StorageExternalParityGroup) {
	o.ExternalParityGroups = v
}

// GetExternalPaths returns the ExternalPaths field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageHitachiExternalPathGroupAllOf) GetExternalPaths() []StorageExternalPath {
	if o == nil {
		var ret []StorageExternalPath
		return ret
	}
	return o.ExternalPaths
}

// GetExternalPathsOk returns a tuple with the ExternalPaths field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageHitachiExternalPathGroupAllOf) GetExternalPathsOk() ([]StorageExternalPath, bool) {
	if o == nil || o.ExternalPaths == nil {
		return nil, false
	}
	return o.ExternalPaths, true
}

// HasExternalPaths returns a boolean if a field has been set.
func (o *StorageHitachiExternalPathGroupAllOf) HasExternalPaths() bool {
	if o != nil && o.ExternalPaths != nil {
		return true
	}

	return false
}

// SetExternalPaths gets a reference to the given []StorageExternalPath and assigns it to the ExternalPaths field.
func (o *StorageHitachiExternalPathGroupAllOf) SetExternalPaths(v []StorageExternalPath) {
	o.ExternalPaths = v
}

// GetExternalProductId returns the ExternalProductId field value if set, zero value otherwise.
func (o *StorageHitachiExternalPathGroupAllOf) GetExternalProductId() string {
	if o == nil || o.ExternalProductId == nil {
		var ret string
		return ret
	}
	return *o.ExternalProductId
}

// GetExternalProductIdOk returns a tuple with the ExternalProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiExternalPathGroupAllOf) GetExternalProductIdOk() (*string, bool) {
	if o == nil || o.ExternalProductId == nil {
		return nil, false
	}
	return o.ExternalProductId, true
}

// HasExternalProductId returns a boolean if a field has been set.
func (o *StorageHitachiExternalPathGroupAllOf) HasExternalProductId() bool {
	if o != nil && o.ExternalProductId != nil {
		return true
	}

	return false
}

// SetExternalProductId gets a reference to the given string and assigns it to the ExternalProductId field.
func (o *StorageHitachiExternalPathGroupAllOf) SetExternalProductId(v string) {
	o.ExternalProductId = &v
}

// GetExternalSerialNumber returns the ExternalSerialNumber field value if set, zero value otherwise.
func (o *StorageHitachiExternalPathGroupAllOf) GetExternalSerialNumber() string {
	if o == nil || o.ExternalSerialNumber == nil {
		var ret string
		return ret
	}
	return *o.ExternalSerialNumber
}

// GetExternalSerialNumberOk returns a tuple with the ExternalSerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiExternalPathGroupAllOf) GetExternalSerialNumberOk() (*string, bool) {
	if o == nil || o.ExternalSerialNumber == nil {
		return nil, false
	}
	return o.ExternalSerialNumber, true
}

// HasExternalSerialNumber returns a boolean if a field has been set.
func (o *StorageHitachiExternalPathGroupAllOf) HasExternalSerialNumber() bool {
	if o != nil && o.ExternalSerialNumber != nil {
		return true
	}

	return false
}

// SetExternalSerialNumber gets a reference to the given string and assigns it to the ExternalSerialNumber field.
func (o *StorageHitachiExternalPathGroupAllOf) SetExternalSerialNumber(v string) {
	o.ExternalSerialNumber = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StorageHitachiExternalPathGroupAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiExternalPathGroupAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StorageHitachiExternalPathGroupAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StorageHitachiExternalPathGroupAllOf) SetName(v string) {
	o.Name = &v
}

// GetArray returns the Array field value if set, zero value otherwise.
func (o *StorageHitachiExternalPathGroupAllOf) GetArray() StorageHitachiArrayRelationship {
	if o == nil || o.Array == nil {
		var ret StorageHitachiArrayRelationship
		return ret
	}
	return *o.Array
}

// GetArrayOk returns a tuple with the Array field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiExternalPathGroupAllOf) GetArrayOk() (*StorageHitachiArrayRelationship, bool) {
	if o == nil || o.Array == nil {
		return nil, false
	}
	return o.Array, true
}

// HasArray returns a boolean if a field has been set.
func (o *StorageHitachiExternalPathGroupAllOf) HasArray() bool {
	if o != nil && o.Array != nil {
		return true
	}

	return false
}

// SetArray gets a reference to the given StorageHitachiArrayRelationship and assigns it to the Array field.
func (o *StorageHitachiExternalPathGroupAllOf) SetArray(v StorageHitachiArrayRelationship) {
	o.Array = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *StorageHitachiExternalPathGroupAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiExternalPathGroupAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StorageHitachiExternalPathGroupAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StorageHitachiExternalPathGroupAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o StorageHitachiExternalPathGroupAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ExternalParityGroups != nil {
		toSerialize["ExternalParityGroups"] = o.ExternalParityGroups
	}
	if o.ExternalPaths != nil {
		toSerialize["ExternalPaths"] = o.ExternalPaths
	}
	if o.ExternalProductId != nil {
		toSerialize["ExternalProductId"] = o.ExternalProductId
	}
	if o.ExternalSerialNumber != nil {
		toSerialize["ExternalSerialNumber"] = o.ExternalSerialNumber
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Array != nil {
		toSerialize["Array"] = o.Array
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageHitachiExternalPathGroupAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageHitachiExternalPathGroupAllOf := _StorageHitachiExternalPathGroupAllOf{}

	if err = json.Unmarshal(bytes, &varStorageHitachiExternalPathGroupAllOf); err == nil {
		*o = StorageHitachiExternalPathGroupAllOf(varStorageHitachiExternalPathGroupAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ExternalParityGroups")
		delete(additionalProperties, "ExternalPaths")
		delete(additionalProperties, "ExternalProductId")
		delete(additionalProperties, "ExternalSerialNumber")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Array")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageHitachiExternalPathGroupAllOf struct {
	value *StorageHitachiExternalPathGroupAllOf
	isSet bool
}

func (v NullableStorageHitachiExternalPathGroupAllOf) Get() *StorageHitachiExternalPathGroupAllOf {
	return v.value
}

func (v *NullableStorageHitachiExternalPathGroupAllOf) Set(val *StorageHitachiExternalPathGroupAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageHitachiExternalPathGroupAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageHitachiExternalPathGroupAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageHitachiExternalPathGroupAllOf(val *StorageHitachiExternalPathGroupAllOf) *NullableStorageHitachiExternalPathGroupAllOf {
	return &NullableStorageHitachiExternalPathGroupAllOf{value: val, isSet: true}
}

func (v NullableStorageHitachiExternalPathGroupAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageHitachiExternalPathGroupAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
