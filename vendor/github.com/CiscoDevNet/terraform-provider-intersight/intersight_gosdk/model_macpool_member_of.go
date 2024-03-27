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

// MacpoolMemberOf The pool and block to which a pool member belongs to.
type MacpoolMemberOf struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The moid of the block of a pool to which this ID belongs.
	IdBlockMoid *string `json:"IdBlockMoid,omitempty"`
	// The moid of the pool to which this ID belongs.
	PoolMoid             *string `json:"PoolMoid,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MacpoolMemberOf MacpoolMemberOf

// NewMacpoolMemberOf instantiates a new MacpoolMemberOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMacpoolMemberOf(classId string, objectType string) *MacpoolMemberOf {
	this := MacpoolMemberOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewMacpoolMemberOfWithDefaults instantiates a new MacpoolMemberOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMacpoolMemberOfWithDefaults() *MacpoolMemberOf {
	this := MacpoolMemberOf{}
	var classId string = "macpool.MemberOf"
	this.ClassId = classId
	var objectType string = "macpool.MemberOf"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *MacpoolMemberOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *MacpoolMemberOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *MacpoolMemberOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *MacpoolMemberOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *MacpoolMemberOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *MacpoolMemberOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIdBlockMoid returns the IdBlockMoid field value if set, zero value otherwise.
func (o *MacpoolMemberOf) GetIdBlockMoid() string {
	if o == nil || o.IdBlockMoid == nil {
		var ret string
		return ret
	}
	return *o.IdBlockMoid
}

// GetIdBlockMoidOk returns a tuple with the IdBlockMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MacpoolMemberOf) GetIdBlockMoidOk() (*string, bool) {
	if o == nil || o.IdBlockMoid == nil {
		return nil, false
	}
	return o.IdBlockMoid, true
}

// HasIdBlockMoid returns a boolean if a field has been set.
func (o *MacpoolMemberOf) HasIdBlockMoid() bool {
	if o != nil && o.IdBlockMoid != nil {
		return true
	}

	return false
}

// SetIdBlockMoid gets a reference to the given string and assigns it to the IdBlockMoid field.
func (o *MacpoolMemberOf) SetIdBlockMoid(v string) {
	o.IdBlockMoid = &v
}

// GetPoolMoid returns the PoolMoid field value if set, zero value otherwise.
func (o *MacpoolMemberOf) GetPoolMoid() string {
	if o == nil || o.PoolMoid == nil {
		var ret string
		return ret
	}
	return *o.PoolMoid
}

// GetPoolMoidOk returns a tuple with the PoolMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MacpoolMemberOf) GetPoolMoidOk() (*string, bool) {
	if o == nil || o.PoolMoid == nil {
		return nil, false
	}
	return o.PoolMoid, true
}

// HasPoolMoid returns a boolean if a field has been set.
func (o *MacpoolMemberOf) HasPoolMoid() bool {
	if o != nil && o.PoolMoid != nil {
		return true
	}

	return false
}

// SetPoolMoid gets a reference to the given string and assigns it to the PoolMoid field.
func (o *MacpoolMemberOf) SetPoolMoid(v string) {
	o.PoolMoid = &v
}

func (o MacpoolMemberOf) MarshalJSON() ([]byte, error) {
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
	if o.IdBlockMoid != nil {
		toSerialize["IdBlockMoid"] = o.IdBlockMoid
	}
	if o.PoolMoid != nil {
		toSerialize["PoolMoid"] = o.PoolMoid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MacpoolMemberOf) UnmarshalJSON(bytes []byte) (err error) {
	type MacpoolMemberOfWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The moid of the block of a pool to which this ID belongs.
		IdBlockMoid *string `json:"IdBlockMoid,omitempty"`
		// The moid of the pool to which this ID belongs.
		PoolMoid *string `json:"PoolMoid,omitempty"`
	}

	varMacpoolMemberOfWithoutEmbeddedStruct := MacpoolMemberOfWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varMacpoolMemberOfWithoutEmbeddedStruct)
	if err == nil {
		varMacpoolMemberOf := _MacpoolMemberOf{}
		varMacpoolMemberOf.ClassId = varMacpoolMemberOfWithoutEmbeddedStruct.ClassId
		varMacpoolMemberOf.ObjectType = varMacpoolMemberOfWithoutEmbeddedStruct.ObjectType
		varMacpoolMemberOf.IdBlockMoid = varMacpoolMemberOfWithoutEmbeddedStruct.IdBlockMoid
		varMacpoolMemberOf.PoolMoid = varMacpoolMemberOfWithoutEmbeddedStruct.PoolMoid
		*o = MacpoolMemberOf(varMacpoolMemberOf)
	} else {
		return err
	}

	varMacpoolMemberOf := _MacpoolMemberOf{}

	err = json.Unmarshal(bytes, &varMacpoolMemberOf)
	if err == nil {
		o.MoBaseComplexType = varMacpoolMemberOf.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "IdBlockMoid")
		delete(additionalProperties, "PoolMoid")

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

type NullableMacpoolMemberOf struct {
	value *MacpoolMemberOf
	isSet bool
}

func (v NullableMacpoolMemberOf) Get() *MacpoolMemberOf {
	return v.value
}

func (v *NullableMacpoolMemberOf) Set(val *MacpoolMemberOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMacpoolMemberOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMacpoolMemberOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMacpoolMemberOf(val *MacpoolMemberOf) *NullableMacpoolMemberOf {
	return &NullableMacpoolMemberOf{value: val, isSet: true}
}

func (v NullableMacpoolMemberOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMacpoolMemberOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
