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

// HyperflexNamedVsan A VSAN with a name and ID. VSANs are used when defining Fibre Channel external storage policies for the cluster.
type HyperflexNamedVsan struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The name of the VSAN. The name can be from 1 to 32 characters long and can contain a combination of alphanumeric characters, underscores, and hyphens.
	Name *string `json:"Name,omitempty"`
	// The ID of the named VSAN. The ID can be any number between 1 and 4093, inclusive.
	VsanId               *int64 `json:"VsanId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexNamedVsan HyperflexNamedVsan

// NewHyperflexNamedVsan instantiates a new HyperflexNamedVsan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexNamedVsan(classId string, objectType string) *HyperflexNamedVsan {
	this := HyperflexNamedVsan{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexNamedVsanWithDefaults instantiates a new HyperflexNamedVsan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexNamedVsanWithDefaults() *HyperflexNamedVsan {
	this := HyperflexNamedVsan{}
	var classId string = "hyperflex.NamedVsan"
	this.ClassId = classId
	var objectType string = "hyperflex.NamedVsan"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexNamedVsan) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexNamedVsan) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexNamedVsan) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexNamedVsan) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexNamedVsan) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexNamedVsan) SetObjectType(v string) {
	o.ObjectType = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *HyperflexNamedVsan) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexNamedVsan) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *HyperflexNamedVsan) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *HyperflexNamedVsan) SetName(v string) {
	o.Name = &v
}

// GetVsanId returns the VsanId field value if set, zero value otherwise.
func (o *HyperflexNamedVsan) GetVsanId() int64 {
	if o == nil || o.VsanId == nil {
		var ret int64
		return ret
	}
	return *o.VsanId
}

// GetVsanIdOk returns a tuple with the VsanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexNamedVsan) GetVsanIdOk() (*int64, bool) {
	if o == nil || o.VsanId == nil {
		return nil, false
	}
	return o.VsanId, true
}

// HasVsanId returns a boolean if a field has been set.
func (o *HyperflexNamedVsan) HasVsanId() bool {
	if o != nil && o.VsanId != nil {
		return true
	}

	return false
}

// SetVsanId gets a reference to the given int64 and assigns it to the VsanId field.
func (o *HyperflexNamedVsan) SetVsanId(v int64) {
	o.VsanId = &v
}

func (o HyperflexNamedVsan) MarshalJSON() ([]byte, error) {
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
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.VsanId != nil {
		toSerialize["VsanId"] = o.VsanId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexNamedVsan) UnmarshalJSON(bytes []byte) (err error) {
	type HyperflexNamedVsanWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The name of the VSAN. The name can be from 1 to 32 characters long and can contain a combination of alphanumeric characters, underscores, and hyphens.
		Name *string `json:"Name,omitempty"`
		// The ID of the named VSAN. The ID can be any number between 1 and 4093, inclusive.
		VsanId *int64 `json:"VsanId,omitempty"`
	}

	varHyperflexNamedVsanWithoutEmbeddedStruct := HyperflexNamedVsanWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHyperflexNamedVsanWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexNamedVsan := _HyperflexNamedVsan{}
		varHyperflexNamedVsan.ClassId = varHyperflexNamedVsanWithoutEmbeddedStruct.ClassId
		varHyperflexNamedVsan.ObjectType = varHyperflexNamedVsanWithoutEmbeddedStruct.ObjectType
		varHyperflexNamedVsan.Name = varHyperflexNamedVsanWithoutEmbeddedStruct.Name
		varHyperflexNamedVsan.VsanId = varHyperflexNamedVsanWithoutEmbeddedStruct.VsanId
		*o = HyperflexNamedVsan(varHyperflexNamedVsan)
	} else {
		return err
	}

	varHyperflexNamedVsan := _HyperflexNamedVsan{}

	err = json.Unmarshal(bytes, &varHyperflexNamedVsan)
	if err == nil {
		o.MoBaseComplexType = varHyperflexNamedVsan.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "VsanId")

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

type NullableHyperflexNamedVsan struct {
	value *HyperflexNamedVsan
	isSet bool
}

func (v NullableHyperflexNamedVsan) Get() *HyperflexNamedVsan {
	return v.value
}

func (v *NullableHyperflexNamedVsan) Set(val *HyperflexNamedVsan) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexNamedVsan) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexNamedVsan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexNamedVsan(val *HyperflexNamedVsan) *NullableHyperflexNamedVsan {
	return &NullableHyperflexNamedVsan{value: val, isSet: true}
}

func (v NullableHyperflexNamedVsan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexNamedVsan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
