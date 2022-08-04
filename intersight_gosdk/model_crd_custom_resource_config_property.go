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

// CrdCustomResourceConfigProperty Depicts a configurable property for the custom resource.
type CrdCustomResourceConfigProperty struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Name of the key/value property.
	Key *string `json:"Key,omitempty"`
	// Value of the key/value property.
	Value                *string `json:"Value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CrdCustomResourceConfigProperty CrdCustomResourceConfigProperty

// NewCrdCustomResourceConfigProperty instantiates a new CrdCustomResourceConfigProperty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCrdCustomResourceConfigProperty(classId string, objectType string) *CrdCustomResourceConfigProperty {
	this := CrdCustomResourceConfigProperty{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCrdCustomResourceConfigPropertyWithDefaults instantiates a new CrdCustomResourceConfigProperty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCrdCustomResourceConfigPropertyWithDefaults() *CrdCustomResourceConfigProperty {
	this := CrdCustomResourceConfigProperty{}
	var classId string = "crd.CustomResourceConfigProperty"
	this.ClassId = classId
	var objectType string = "crd.CustomResourceConfigProperty"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CrdCustomResourceConfigProperty) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CrdCustomResourceConfigProperty) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CrdCustomResourceConfigProperty) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CrdCustomResourceConfigProperty) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CrdCustomResourceConfigProperty) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CrdCustomResourceConfigProperty) SetObjectType(v string) {
	o.ObjectType = v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *CrdCustomResourceConfigProperty) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CrdCustomResourceConfigProperty) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *CrdCustomResourceConfigProperty) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *CrdCustomResourceConfigProperty) SetKey(v string) {
	o.Key = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CrdCustomResourceConfigProperty) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CrdCustomResourceConfigProperty) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CrdCustomResourceConfigProperty) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *CrdCustomResourceConfigProperty) SetValue(v string) {
	o.Value = &v
}

func (o CrdCustomResourceConfigProperty) MarshalJSON() ([]byte, error) {
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
	if o.Key != nil {
		toSerialize["Key"] = o.Key
	}
	if o.Value != nil {
		toSerialize["Value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CrdCustomResourceConfigProperty) UnmarshalJSON(bytes []byte) (err error) {
	type CrdCustomResourceConfigPropertyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Name of the key/value property.
		Key *string `json:"Key,omitempty"`
		// Value of the key/value property.
		Value *string `json:"Value,omitempty"`
	}

	varCrdCustomResourceConfigPropertyWithoutEmbeddedStruct := CrdCustomResourceConfigPropertyWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varCrdCustomResourceConfigPropertyWithoutEmbeddedStruct)
	if err == nil {
		varCrdCustomResourceConfigProperty := _CrdCustomResourceConfigProperty{}
		varCrdCustomResourceConfigProperty.ClassId = varCrdCustomResourceConfigPropertyWithoutEmbeddedStruct.ClassId
		varCrdCustomResourceConfigProperty.ObjectType = varCrdCustomResourceConfigPropertyWithoutEmbeddedStruct.ObjectType
		varCrdCustomResourceConfigProperty.Key = varCrdCustomResourceConfigPropertyWithoutEmbeddedStruct.Key
		varCrdCustomResourceConfigProperty.Value = varCrdCustomResourceConfigPropertyWithoutEmbeddedStruct.Value
		*o = CrdCustomResourceConfigProperty(varCrdCustomResourceConfigProperty)
	} else {
		return err
	}

	varCrdCustomResourceConfigProperty := _CrdCustomResourceConfigProperty{}

	err = json.Unmarshal(bytes, &varCrdCustomResourceConfigProperty)
	if err == nil {
		o.MoBaseComplexType = varCrdCustomResourceConfigProperty.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Key")
		delete(additionalProperties, "Value")

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

type NullableCrdCustomResourceConfigProperty struct {
	value *CrdCustomResourceConfigProperty
	isSet bool
}

func (v NullableCrdCustomResourceConfigProperty) Get() *CrdCustomResourceConfigProperty {
	return v.value
}

func (v *NullableCrdCustomResourceConfigProperty) Set(val *CrdCustomResourceConfigProperty) {
	v.value = val
	v.isSet = true
}

func (v NullableCrdCustomResourceConfigProperty) IsSet() bool {
	return v.isSet
}

func (v *NullableCrdCustomResourceConfigProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCrdCustomResourceConfigProperty(val *CrdCustomResourceConfigProperty) *NullableCrdCustomResourceConfigProperty {
	return &NullableCrdCustomResourceConfigProperty{value: val, isSet: true}
}

func (v NullableCrdCustomResourceConfigProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCrdCustomResourceConfigProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
