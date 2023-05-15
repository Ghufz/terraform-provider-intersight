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
	"reflect"
	"strings"
)

// ComputeStorageVirtualDrive Storage Virtual Drives on which an operation has to be performed.
type ComputeStorageVirtualDrive struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Virtual Drive ID of the storage on the server.
	Id                   *string `json:"Id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ComputeStorageVirtualDrive ComputeStorageVirtualDrive

// NewComputeStorageVirtualDrive instantiates a new ComputeStorageVirtualDrive object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputeStorageVirtualDrive(classId string, objectType string) *ComputeStorageVirtualDrive {
	this := ComputeStorageVirtualDrive{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewComputeStorageVirtualDriveWithDefaults instantiates a new ComputeStorageVirtualDrive object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputeStorageVirtualDriveWithDefaults() *ComputeStorageVirtualDrive {
	this := ComputeStorageVirtualDrive{}
	var classId string = "compute.StorageVirtualDrive"
	this.ClassId = classId
	var objectType string = "compute.StorageVirtualDrive"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ComputeStorageVirtualDrive) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ComputeStorageVirtualDrive) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ComputeStorageVirtualDrive) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ComputeStorageVirtualDrive) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ComputeStorageVirtualDrive) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ComputeStorageVirtualDrive) SetObjectType(v string) {
	o.ObjectType = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ComputeStorageVirtualDrive) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeStorageVirtualDrive) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ComputeStorageVirtualDrive) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ComputeStorageVirtualDrive) SetId(v string) {
	o.Id = &v
}

func (o ComputeStorageVirtualDrive) MarshalJSON() ([]byte, error) {
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
	if o.Id != nil {
		toSerialize["Id"] = o.Id
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ComputeStorageVirtualDrive) UnmarshalJSON(bytes []byte) (err error) {
	type ComputeStorageVirtualDriveWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Virtual Drive ID of the storage on the server.
		Id *string `json:"Id,omitempty"`
	}

	varComputeStorageVirtualDriveWithoutEmbeddedStruct := ComputeStorageVirtualDriveWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varComputeStorageVirtualDriveWithoutEmbeddedStruct)
	if err == nil {
		varComputeStorageVirtualDrive := _ComputeStorageVirtualDrive{}
		varComputeStorageVirtualDrive.ClassId = varComputeStorageVirtualDriveWithoutEmbeddedStruct.ClassId
		varComputeStorageVirtualDrive.ObjectType = varComputeStorageVirtualDriveWithoutEmbeddedStruct.ObjectType
		varComputeStorageVirtualDrive.Id = varComputeStorageVirtualDriveWithoutEmbeddedStruct.Id
		*o = ComputeStorageVirtualDrive(varComputeStorageVirtualDrive)
	} else {
		return err
	}

	varComputeStorageVirtualDrive := _ComputeStorageVirtualDrive{}

	err = json.Unmarshal(bytes, &varComputeStorageVirtualDrive)
	if err == nil {
		o.MoBaseComplexType = varComputeStorageVirtualDrive.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Id")

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

type NullableComputeStorageVirtualDrive struct {
	value *ComputeStorageVirtualDrive
	isSet bool
}

func (v NullableComputeStorageVirtualDrive) Get() *ComputeStorageVirtualDrive {
	return v.value
}

func (v *NullableComputeStorageVirtualDrive) Set(val *ComputeStorageVirtualDrive) {
	v.value = val
	v.isSet = true
}

func (v NullableComputeStorageVirtualDrive) IsSet() bool {
	return v.isSet
}

func (v *NullableComputeStorageVirtualDrive) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputeStorageVirtualDrive(val *ComputeStorageVirtualDrive) *NullableComputeStorageVirtualDrive {
	return &NullableComputeStorageVirtualDrive{value: val, isSet: true}
}

func (v NullableComputeStorageVirtualDrive) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputeStorageVirtualDrive) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
