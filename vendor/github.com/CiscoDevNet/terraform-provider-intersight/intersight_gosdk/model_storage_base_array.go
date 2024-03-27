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

// StorageBaseArray Abstract object which contains common attributes of a storage array. Every storage array should inherit from this object.
type StorageBaseArray struct {
	EquipmentAbstractDevice
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType           string                      `json:"ObjectType"`
	StorageUtilization   NullableStorageBaseCapacity `json:"StorageUtilization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageBaseArray StorageBaseArray

// NewStorageBaseArray instantiates a new StorageBaseArray object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageBaseArray(classId string, objectType string) *StorageBaseArray {
	this := StorageBaseArray{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageBaseArrayWithDefaults instantiates a new StorageBaseArray object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageBaseArrayWithDefaults() *StorageBaseArray {
	this := StorageBaseArray{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageBaseArray) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageBaseArray) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageBaseArray) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageBaseArray) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageBaseArray) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageBaseArray) SetObjectType(v string) {
	o.ObjectType = v
}

// GetStorageUtilization returns the StorageUtilization field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageBaseArray) GetStorageUtilization() StorageBaseCapacity {
	if o == nil || o.StorageUtilization.Get() == nil {
		var ret StorageBaseCapacity
		return ret
	}
	return *o.StorageUtilization.Get()
}

// GetStorageUtilizationOk returns a tuple with the StorageUtilization field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageBaseArray) GetStorageUtilizationOk() (*StorageBaseCapacity, bool) {
	if o == nil {
		return nil, false
	}
	return o.StorageUtilization.Get(), o.StorageUtilization.IsSet()
}

// HasStorageUtilization returns a boolean if a field has been set.
func (o *StorageBaseArray) HasStorageUtilization() bool {
	if o != nil && o.StorageUtilization.IsSet() {
		return true
	}

	return false
}

// SetStorageUtilization gets a reference to the given NullableStorageBaseCapacity and assigns it to the StorageUtilization field.
func (o *StorageBaseArray) SetStorageUtilization(v StorageBaseCapacity) {
	o.StorageUtilization.Set(&v)
}

// SetStorageUtilizationNil sets the value for StorageUtilization to be an explicit nil
func (o *StorageBaseArray) SetStorageUtilizationNil() {
	o.StorageUtilization.Set(nil)
}

// UnsetStorageUtilization ensures that no value is present for StorageUtilization, not even an explicit nil
func (o *StorageBaseArray) UnsetStorageUtilization() {
	o.StorageUtilization.Unset()
}

func (o StorageBaseArray) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedEquipmentAbstractDevice, errEquipmentAbstractDevice := json.Marshal(o.EquipmentAbstractDevice)
	if errEquipmentAbstractDevice != nil {
		return []byte{}, errEquipmentAbstractDevice
	}
	errEquipmentAbstractDevice = json.Unmarshal([]byte(serializedEquipmentAbstractDevice), &toSerialize)
	if errEquipmentAbstractDevice != nil {
		return []byte{}, errEquipmentAbstractDevice
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.StorageUtilization.IsSet() {
		toSerialize["StorageUtilization"] = o.StorageUtilization.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageBaseArray) UnmarshalJSON(bytes []byte) (err error) {
	type StorageBaseArrayWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType         string                      `json:"ObjectType"`
		StorageUtilization NullableStorageBaseCapacity `json:"StorageUtilization,omitempty"`
	}

	varStorageBaseArrayWithoutEmbeddedStruct := StorageBaseArrayWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageBaseArrayWithoutEmbeddedStruct)
	if err == nil {
		varStorageBaseArray := _StorageBaseArray{}
		varStorageBaseArray.ClassId = varStorageBaseArrayWithoutEmbeddedStruct.ClassId
		varStorageBaseArray.ObjectType = varStorageBaseArrayWithoutEmbeddedStruct.ObjectType
		varStorageBaseArray.StorageUtilization = varStorageBaseArrayWithoutEmbeddedStruct.StorageUtilization
		*o = StorageBaseArray(varStorageBaseArray)
	} else {
		return err
	}

	varStorageBaseArray := _StorageBaseArray{}

	err = json.Unmarshal(bytes, &varStorageBaseArray)
	if err == nil {
		o.EquipmentAbstractDevice = varStorageBaseArray.EquipmentAbstractDevice
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "StorageUtilization")

		// remove fields from embedded structs
		reflectEquipmentAbstractDevice := reflect.ValueOf(o.EquipmentAbstractDevice)
		for i := 0; i < reflectEquipmentAbstractDevice.Type().NumField(); i++ {
			t := reflectEquipmentAbstractDevice.Type().Field(i)

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

type NullableStorageBaseArray struct {
	value *StorageBaseArray
	isSet bool
}

func (v NullableStorageBaseArray) Get() *StorageBaseArray {
	return v.value
}

func (v *NullableStorageBaseArray) Set(val *StorageBaseArray) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageBaseArray) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageBaseArray) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageBaseArray(val *StorageBaseArray) *NullableStorageBaseArray {
	return &NullableStorageBaseArray{value: val, isSet: true}
}

func (v NullableStorageBaseArray) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageBaseArray) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
