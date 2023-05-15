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

// StorageNetAppClusterSnapMirrorPolicy NetApp SnapMirror policy owned by the cluster. NetApp SnapMirror policy when applied to a SnapMirror relationship, controls the behavior of the relationship and specifies the configuration attributes for that relationship.
type StorageNetAppClusterSnapMirrorPolicy struct {
	StorageNetAppBaseSnapMirrorPolicy
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                            `json:"ObjectType"`
	Array                *StorageNetAppClusterRelationship `json:"Array,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppClusterSnapMirrorPolicy StorageNetAppClusterSnapMirrorPolicy

// NewStorageNetAppClusterSnapMirrorPolicy instantiates a new StorageNetAppClusterSnapMirrorPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppClusterSnapMirrorPolicy(classId string, objectType string) *StorageNetAppClusterSnapMirrorPolicy {
	this := StorageNetAppClusterSnapMirrorPolicy{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppClusterSnapMirrorPolicyWithDefaults instantiates a new StorageNetAppClusterSnapMirrorPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppClusterSnapMirrorPolicyWithDefaults() *StorageNetAppClusterSnapMirrorPolicy {
	this := StorageNetAppClusterSnapMirrorPolicy{}
	var classId string = "storage.NetAppClusterSnapMirrorPolicy"
	this.ClassId = classId
	var objectType string = "storage.NetAppClusterSnapMirrorPolicy"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppClusterSnapMirrorPolicy) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppClusterSnapMirrorPolicy) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppClusterSnapMirrorPolicy) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppClusterSnapMirrorPolicy) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppClusterSnapMirrorPolicy) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppClusterSnapMirrorPolicy) SetObjectType(v string) {
	o.ObjectType = v
}

// GetArray returns the Array field value if set, zero value otherwise.
func (o *StorageNetAppClusterSnapMirrorPolicy) GetArray() StorageNetAppClusterRelationship {
	if o == nil || o.Array == nil {
		var ret StorageNetAppClusterRelationship
		return ret
	}
	return *o.Array
}

// GetArrayOk returns a tuple with the Array field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppClusterSnapMirrorPolicy) GetArrayOk() (*StorageNetAppClusterRelationship, bool) {
	if o == nil || o.Array == nil {
		return nil, false
	}
	return o.Array, true
}

// HasArray returns a boolean if a field has been set.
func (o *StorageNetAppClusterSnapMirrorPolicy) HasArray() bool {
	if o != nil && o.Array != nil {
		return true
	}

	return false
}

// SetArray gets a reference to the given StorageNetAppClusterRelationship and assigns it to the Array field.
func (o *StorageNetAppClusterSnapMirrorPolicy) SetArray(v StorageNetAppClusterRelationship) {
	o.Array = &v
}

func (o StorageNetAppClusterSnapMirrorPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedStorageNetAppBaseSnapMirrorPolicy, errStorageNetAppBaseSnapMirrorPolicy := json.Marshal(o.StorageNetAppBaseSnapMirrorPolicy)
	if errStorageNetAppBaseSnapMirrorPolicy != nil {
		return []byte{}, errStorageNetAppBaseSnapMirrorPolicy
	}
	errStorageNetAppBaseSnapMirrorPolicy = json.Unmarshal([]byte(serializedStorageNetAppBaseSnapMirrorPolicy), &toSerialize)
	if errStorageNetAppBaseSnapMirrorPolicy != nil {
		return []byte{}, errStorageNetAppBaseSnapMirrorPolicy
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Array != nil {
		toSerialize["Array"] = o.Array
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageNetAppClusterSnapMirrorPolicy) UnmarshalJSON(bytes []byte) (err error) {
	type StorageNetAppClusterSnapMirrorPolicyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string                            `json:"ObjectType"`
		Array      *StorageNetAppClusterRelationship `json:"Array,omitempty"`
	}

	varStorageNetAppClusterSnapMirrorPolicyWithoutEmbeddedStruct := StorageNetAppClusterSnapMirrorPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageNetAppClusterSnapMirrorPolicyWithoutEmbeddedStruct)
	if err == nil {
		varStorageNetAppClusterSnapMirrorPolicy := _StorageNetAppClusterSnapMirrorPolicy{}
		varStorageNetAppClusterSnapMirrorPolicy.ClassId = varStorageNetAppClusterSnapMirrorPolicyWithoutEmbeddedStruct.ClassId
		varStorageNetAppClusterSnapMirrorPolicy.ObjectType = varStorageNetAppClusterSnapMirrorPolicyWithoutEmbeddedStruct.ObjectType
		varStorageNetAppClusterSnapMirrorPolicy.Array = varStorageNetAppClusterSnapMirrorPolicyWithoutEmbeddedStruct.Array
		*o = StorageNetAppClusterSnapMirrorPolicy(varStorageNetAppClusterSnapMirrorPolicy)
	} else {
		return err
	}

	varStorageNetAppClusterSnapMirrorPolicy := _StorageNetAppClusterSnapMirrorPolicy{}

	err = json.Unmarshal(bytes, &varStorageNetAppClusterSnapMirrorPolicy)
	if err == nil {
		o.StorageNetAppBaseSnapMirrorPolicy = varStorageNetAppClusterSnapMirrorPolicy.StorageNetAppBaseSnapMirrorPolicy
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Array")

		// remove fields from embedded structs
		reflectStorageNetAppBaseSnapMirrorPolicy := reflect.ValueOf(o.StorageNetAppBaseSnapMirrorPolicy)
		for i := 0; i < reflectStorageNetAppBaseSnapMirrorPolicy.Type().NumField(); i++ {
			t := reflectStorageNetAppBaseSnapMirrorPolicy.Type().Field(i)

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

type NullableStorageNetAppClusterSnapMirrorPolicy struct {
	value *StorageNetAppClusterSnapMirrorPolicy
	isSet bool
}

func (v NullableStorageNetAppClusterSnapMirrorPolicy) Get() *StorageNetAppClusterSnapMirrorPolicy {
	return v.value
}

func (v *NullableStorageNetAppClusterSnapMirrorPolicy) Set(val *StorageNetAppClusterSnapMirrorPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppClusterSnapMirrorPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppClusterSnapMirrorPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppClusterSnapMirrorPolicy(val *StorageNetAppClusterSnapMirrorPolicy) *NullableStorageNetAppClusterSnapMirrorPolicy {
	return &NullableStorageNetAppClusterSnapMirrorPolicy{value: val, isSet: true}
}

func (v NullableStorageNetAppClusterSnapMirrorPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppClusterSnapMirrorPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
