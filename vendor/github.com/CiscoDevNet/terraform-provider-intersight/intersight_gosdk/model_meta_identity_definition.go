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

// MetaIdentityDefinition Definitions for the identity constraints in meta.
type MetaIdentityDefinition struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string   `json:"ObjectType"`
	Fields               []string `json:"Fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MetaIdentityDefinition MetaIdentityDefinition

// NewMetaIdentityDefinition instantiates a new MetaIdentityDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetaIdentityDefinition(classId string, objectType string) *MetaIdentityDefinition {
	this := MetaIdentityDefinition{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewMetaIdentityDefinitionWithDefaults instantiates a new MetaIdentityDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetaIdentityDefinitionWithDefaults() *MetaIdentityDefinition {
	this := MetaIdentityDefinition{}
	var classId string = "meta.IdentityDefinition"
	this.ClassId = classId
	var objectType string = "meta.IdentityDefinition"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *MetaIdentityDefinition) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *MetaIdentityDefinition) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *MetaIdentityDefinition) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *MetaIdentityDefinition) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *MetaIdentityDefinition) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *MetaIdentityDefinition) SetObjectType(v string) {
	o.ObjectType = v
}

// GetFields returns the Fields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetaIdentityDefinition) GetFields() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetaIdentityDefinition) GetFieldsOk() ([]string, bool) {
	if o == nil || o.Fields == nil {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *MetaIdentityDefinition) HasFields() bool {
	if o != nil && o.Fields != nil {
		return true
	}

	return false
}

// SetFields gets a reference to the given []string and assigns it to the Fields field.
func (o *MetaIdentityDefinition) SetFields(v []string) {
	o.Fields = v
}

func (o MetaIdentityDefinition) MarshalJSON() ([]byte, error) {
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
	if o.Fields != nil {
		toSerialize["Fields"] = o.Fields
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MetaIdentityDefinition) UnmarshalJSON(bytes []byte) (err error) {
	type MetaIdentityDefinitionWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string   `json:"ObjectType"`
		Fields     []string `json:"Fields,omitempty"`
	}

	varMetaIdentityDefinitionWithoutEmbeddedStruct := MetaIdentityDefinitionWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varMetaIdentityDefinitionWithoutEmbeddedStruct)
	if err == nil {
		varMetaIdentityDefinition := _MetaIdentityDefinition{}
		varMetaIdentityDefinition.ClassId = varMetaIdentityDefinitionWithoutEmbeddedStruct.ClassId
		varMetaIdentityDefinition.ObjectType = varMetaIdentityDefinitionWithoutEmbeddedStruct.ObjectType
		varMetaIdentityDefinition.Fields = varMetaIdentityDefinitionWithoutEmbeddedStruct.Fields
		*o = MetaIdentityDefinition(varMetaIdentityDefinition)
	} else {
		return err
	}

	varMetaIdentityDefinition := _MetaIdentityDefinition{}

	err = json.Unmarshal(bytes, &varMetaIdentityDefinition)
	if err == nil {
		o.MoBaseComplexType = varMetaIdentityDefinition.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Fields")

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

type NullableMetaIdentityDefinition struct {
	value *MetaIdentityDefinition
	isSet bool
}

func (v NullableMetaIdentityDefinition) Get() *MetaIdentityDefinition {
	return v.value
}

func (v *NullableMetaIdentityDefinition) Set(val *MetaIdentityDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableMetaIdentityDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableMetaIdentityDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetaIdentityDefinition(val *MetaIdentityDefinition) *NullableMetaIdentityDefinition {
	return &NullableMetaIdentityDefinition{value: val, isSet: true}
}

func (v NullableMetaIdentityDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetaIdentityDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
