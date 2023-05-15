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

// StorageNetAppNodeEventAllOf Definition of the list of properties defined in 'storage.NetAppNodeEvent', excluding properties defined in parent classes.
type StorageNetAppNodeEventAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                         `json:"ObjectType"`
	ArrayController      *StorageNetAppNodeRelationship `json:"ArrayController,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppNodeEventAllOf StorageNetAppNodeEventAllOf

// NewStorageNetAppNodeEventAllOf instantiates a new StorageNetAppNodeEventAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppNodeEventAllOf(classId string, objectType string) *StorageNetAppNodeEventAllOf {
	this := StorageNetAppNodeEventAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppNodeEventAllOfWithDefaults instantiates a new StorageNetAppNodeEventAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppNodeEventAllOfWithDefaults() *StorageNetAppNodeEventAllOf {
	this := StorageNetAppNodeEventAllOf{}
	var classId string = "storage.NetAppNodeEvent"
	this.ClassId = classId
	var objectType string = "storage.NetAppNodeEvent"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppNodeEventAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppNodeEventAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppNodeEventAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppNodeEventAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppNodeEventAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppNodeEventAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetArrayController returns the ArrayController field value if set, zero value otherwise.
func (o *StorageNetAppNodeEventAllOf) GetArrayController() StorageNetAppNodeRelationship {
	if o == nil || o.ArrayController == nil {
		var ret StorageNetAppNodeRelationship
		return ret
	}
	return *o.ArrayController
}

// GetArrayControllerOk returns a tuple with the ArrayController field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppNodeEventAllOf) GetArrayControllerOk() (*StorageNetAppNodeRelationship, bool) {
	if o == nil || o.ArrayController == nil {
		return nil, false
	}
	return o.ArrayController, true
}

// HasArrayController returns a boolean if a field has been set.
func (o *StorageNetAppNodeEventAllOf) HasArrayController() bool {
	if o != nil && o.ArrayController != nil {
		return true
	}

	return false
}

// SetArrayController gets a reference to the given StorageNetAppNodeRelationship and assigns it to the ArrayController field.
func (o *StorageNetAppNodeEventAllOf) SetArrayController(v StorageNetAppNodeRelationship) {
	o.ArrayController = &v
}

func (o StorageNetAppNodeEventAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ArrayController != nil {
		toSerialize["ArrayController"] = o.ArrayController
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageNetAppNodeEventAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageNetAppNodeEventAllOf := _StorageNetAppNodeEventAllOf{}

	if err = json.Unmarshal(bytes, &varStorageNetAppNodeEventAllOf); err == nil {
		*o = StorageNetAppNodeEventAllOf(varStorageNetAppNodeEventAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ArrayController")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageNetAppNodeEventAllOf struct {
	value *StorageNetAppNodeEventAllOf
	isSet bool
}

func (v NullableStorageNetAppNodeEventAllOf) Get() *StorageNetAppNodeEventAllOf {
	return v.value
}

func (v *NullableStorageNetAppNodeEventAllOf) Set(val *StorageNetAppNodeEventAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppNodeEventAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppNodeEventAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppNodeEventAllOf(val *StorageNetAppNodeEventAllOf) *NullableStorageNetAppNodeEventAllOf {
	return &NullableStorageNetAppNodeEventAllOf{value: val, isSet: true}
}

func (v NullableStorageNetAppNodeEventAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppNodeEventAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
