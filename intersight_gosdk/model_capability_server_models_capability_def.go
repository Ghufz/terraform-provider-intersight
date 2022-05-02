/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-6484
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// CapabilityServerModelsCapabilityDef Used to categorize server models.
type CapabilityServerModelsCapabilityDef struct {
	CapabilityCapability
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string   `json:"ObjectType"`
	Models     []string `json:"Models,omitempty"`
	// Type of the server. Example, BladeM6, RackM5.
	ServerType           *string `json:"ServerType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CapabilityServerModelsCapabilityDef CapabilityServerModelsCapabilityDef

// NewCapabilityServerModelsCapabilityDef instantiates a new CapabilityServerModelsCapabilityDef object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilityServerModelsCapabilityDef(classId string, objectType string) *CapabilityServerModelsCapabilityDef {
	this := CapabilityServerModelsCapabilityDef{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCapabilityServerModelsCapabilityDefWithDefaults instantiates a new CapabilityServerModelsCapabilityDef object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilityServerModelsCapabilityDefWithDefaults() *CapabilityServerModelsCapabilityDef {
	this := CapabilityServerModelsCapabilityDef{}
	var classId string = "capability.ServerModelsCapabilityDef"
	this.ClassId = classId
	var objectType string = "capability.ServerModelsCapabilityDef"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CapabilityServerModelsCapabilityDef) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CapabilityServerModelsCapabilityDef) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CapabilityServerModelsCapabilityDef) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CapabilityServerModelsCapabilityDef) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CapabilityServerModelsCapabilityDef) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CapabilityServerModelsCapabilityDef) SetObjectType(v string) {
	o.ObjectType = v
}

// GetModels returns the Models field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CapabilityServerModelsCapabilityDef) GetModels() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Models
}

// GetModelsOk returns a tuple with the Models field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CapabilityServerModelsCapabilityDef) GetModelsOk() ([]string, bool) {
	if o == nil || o.Models == nil {
		return nil, false
	}
	return o.Models, true
}

// HasModels returns a boolean if a field has been set.
func (o *CapabilityServerModelsCapabilityDef) HasModels() bool {
	if o != nil && o.Models != nil {
		return true
	}

	return false
}

// SetModels gets a reference to the given []string and assigns it to the Models field.
func (o *CapabilityServerModelsCapabilityDef) SetModels(v []string) {
	o.Models = v
}

// GetServerType returns the ServerType field value if set, zero value otherwise.
func (o *CapabilityServerModelsCapabilityDef) GetServerType() string {
	if o == nil || o.ServerType == nil {
		var ret string
		return ret
	}
	return *o.ServerType
}

// GetServerTypeOk returns a tuple with the ServerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityServerModelsCapabilityDef) GetServerTypeOk() (*string, bool) {
	if o == nil || o.ServerType == nil {
		return nil, false
	}
	return o.ServerType, true
}

// HasServerType returns a boolean if a field has been set.
func (o *CapabilityServerModelsCapabilityDef) HasServerType() bool {
	if o != nil && o.ServerType != nil {
		return true
	}

	return false
}

// SetServerType gets a reference to the given string and assigns it to the ServerType field.
func (o *CapabilityServerModelsCapabilityDef) SetServerType(v string) {
	o.ServerType = &v
}

func (o CapabilityServerModelsCapabilityDef) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedCapabilityCapability, errCapabilityCapability := json.Marshal(o.CapabilityCapability)
	if errCapabilityCapability != nil {
		return []byte{}, errCapabilityCapability
	}
	errCapabilityCapability = json.Unmarshal([]byte(serializedCapabilityCapability), &toSerialize)
	if errCapabilityCapability != nil {
		return []byte{}, errCapabilityCapability
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Models != nil {
		toSerialize["Models"] = o.Models
	}
	if o.ServerType != nil {
		toSerialize["ServerType"] = o.ServerType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CapabilityServerModelsCapabilityDef) UnmarshalJSON(bytes []byte) (err error) {
	type CapabilityServerModelsCapabilityDefWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string   `json:"ObjectType"`
		Models     []string `json:"Models,omitempty"`
		// Type of the server. Example, BladeM6, RackM5.
		ServerType *string `json:"ServerType,omitempty"`
	}

	varCapabilityServerModelsCapabilityDefWithoutEmbeddedStruct := CapabilityServerModelsCapabilityDefWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varCapabilityServerModelsCapabilityDefWithoutEmbeddedStruct)
	if err == nil {
		varCapabilityServerModelsCapabilityDef := _CapabilityServerModelsCapabilityDef{}
		varCapabilityServerModelsCapabilityDef.ClassId = varCapabilityServerModelsCapabilityDefWithoutEmbeddedStruct.ClassId
		varCapabilityServerModelsCapabilityDef.ObjectType = varCapabilityServerModelsCapabilityDefWithoutEmbeddedStruct.ObjectType
		varCapabilityServerModelsCapabilityDef.Models = varCapabilityServerModelsCapabilityDefWithoutEmbeddedStruct.Models
		varCapabilityServerModelsCapabilityDef.ServerType = varCapabilityServerModelsCapabilityDefWithoutEmbeddedStruct.ServerType
		*o = CapabilityServerModelsCapabilityDef(varCapabilityServerModelsCapabilityDef)
	} else {
		return err
	}

	varCapabilityServerModelsCapabilityDef := _CapabilityServerModelsCapabilityDef{}

	err = json.Unmarshal(bytes, &varCapabilityServerModelsCapabilityDef)
	if err == nil {
		o.CapabilityCapability = varCapabilityServerModelsCapabilityDef.CapabilityCapability
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Models")
		delete(additionalProperties, "ServerType")

		// remove fields from embedded structs
		reflectCapabilityCapability := reflect.ValueOf(o.CapabilityCapability)
		for i := 0; i < reflectCapabilityCapability.Type().NumField(); i++ {
			t := reflectCapabilityCapability.Type().Field(i)

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

type NullableCapabilityServerModelsCapabilityDef struct {
	value *CapabilityServerModelsCapabilityDef
	isSet bool
}

func (v NullableCapabilityServerModelsCapabilityDef) Get() *CapabilityServerModelsCapabilityDef {
	return v.value
}

func (v *NullableCapabilityServerModelsCapabilityDef) Set(val *CapabilityServerModelsCapabilityDef) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilityServerModelsCapabilityDef) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilityServerModelsCapabilityDef) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilityServerModelsCapabilityDef(val *CapabilityServerModelsCapabilityDef) *NullableCapabilityServerModelsCapabilityDef {
	return &NullableCapabilityServerModelsCapabilityDef{value: val, isSet: true}
}

func (v NullableCapabilityServerModelsCapabilityDef) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilityServerModelsCapabilityDef) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
