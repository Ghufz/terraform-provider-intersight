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

// StorageSpanDrives This models a single span group of disks in a RAID group.
type StorageSpanDrives struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Collection of local disks that are part of this span group. Allowed value is a comma or hyphen separated number range. The minimum number of disks needed in a span group varies based on RAID level. RAID0 requires at least one disk, RAID1 and RAID10 requires at least 2 and in multiples of 2, RAID5 RAID50 RAID6 and RAID60 require at least 3 disks in a span group.
	Slots                *string `json:"Slots,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageSpanDrives StorageSpanDrives

// NewStorageSpanDrives instantiates a new StorageSpanDrives object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageSpanDrives(classId string, objectType string) *StorageSpanDrives {
	this := StorageSpanDrives{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageSpanDrivesWithDefaults instantiates a new StorageSpanDrives object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageSpanDrivesWithDefaults() *StorageSpanDrives {
	this := StorageSpanDrives{}
	var classId string = "storage.SpanDrives"
	this.ClassId = classId
	var objectType string = "storage.SpanDrives"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageSpanDrives) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageSpanDrives) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageSpanDrives) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageSpanDrives) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageSpanDrives) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageSpanDrives) SetObjectType(v string) {
	o.ObjectType = v
}

// GetSlots returns the Slots field value if set, zero value otherwise.
func (o *StorageSpanDrives) GetSlots() string {
	if o == nil || o.Slots == nil {
		var ret string
		return ret
	}
	return *o.Slots
}

// GetSlotsOk returns a tuple with the Slots field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageSpanDrives) GetSlotsOk() (*string, bool) {
	if o == nil || o.Slots == nil {
		return nil, false
	}
	return o.Slots, true
}

// HasSlots returns a boolean if a field has been set.
func (o *StorageSpanDrives) HasSlots() bool {
	if o != nil && o.Slots != nil {
		return true
	}

	return false
}

// SetSlots gets a reference to the given string and assigns it to the Slots field.
func (o *StorageSpanDrives) SetSlots(v string) {
	o.Slots = &v
}

func (o StorageSpanDrives) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Slots != nil {
		toSerialize["Slots"] = o.Slots
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageSpanDrives) UnmarshalJSON(bytes []byte) (err error) {
	type StorageSpanDrivesWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Collection of local disks that are part of this span group. Allowed value is a comma or hyphen separated number range. The minimum number of disks needed in a span group varies based on RAID level. RAID0 requires at least one disk, RAID1 and RAID10 requires at least 2 and in multiples of 2, RAID5 RAID50 RAID6 and RAID60 require at least 3 disks in a span group.
		Slots *string `json:"Slots,omitempty"`
	}

	varStorageSpanDrivesWithoutEmbeddedStruct := StorageSpanDrivesWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageSpanDrivesWithoutEmbeddedStruct)
	if err == nil {
		varStorageSpanDrives := _StorageSpanDrives{}
		varStorageSpanDrives.ClassId = varStorageSpanDrivesWithoutEmbeddedStruct.ClassId
		varStorageSpanDrives.ObjectType = varStorageSpanDrivesWithoutEmbeddedStruct.ObjectType
		varStorageSpanDrives.Slots = varStorageSpanDrivesWithoutEmbeddedStruct.Slots
		*o = StorageSpanDrives(varStorageSpanDrives)
	} else {
		return err
	}

	varStorageSpanDrives := _StorageSpanDrives{}

	err = json.Unmarshal(bytes, &varStorageSpanDrives)
	if err == nil {
		o.MoBaseComplexType = varStorageSpanDrives.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Slots")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

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

type NullableStorageSpanDrives struct {
	value *StorageSpanDrives
	isSet bool
}

func (v NullableStorageSpanDrives) Get() *StorageSpanDrives {
	return v.value
}

func (v *NullableStorageSpanDrives) Set(val *StorageSpanDrives) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageSpanDrives) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageSpanDrives) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageSpanDrives(val *StorageSpanDrives) *NullableStorageSpanDrives {
	return &NullableStorageSpanDrives{value: val, isSet: true}
}

func (v NullableStorageSpanDrives) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageSpanDrives) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
