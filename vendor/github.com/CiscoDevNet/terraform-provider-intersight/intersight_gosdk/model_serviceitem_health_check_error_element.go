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

// ServiceitemHealthCheckErrorElement Details about the service item elements of particular type (eg. Server, Storage, etc.) for which the health check has failed.
type ServiceitemHealthCheckErrorElement struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Type of the element of the service item. Examples are Server, Storage, UCS Fabric Interconnect etc.
	ElementType          *string   `json:"ElementType,omitempty"`
	Elements             []MoMoRef `json:"Elements,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServiceitemHealthCheckErrorElement ServiceitemHealthCheckErrorElement

// NewServiceitemHealthCheckErrorElement instantiates a new ServiceitemHealthCheckErrorElement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceitemHealthCheckErrorElement(classId string, objectType string) *ServiceitemHealthCheckErrorElement {
	this := ServiceitemHealthCheckErrorElement{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewServiceitemHealthCheckErrorElementWithDefaults instantiates a new ServiceitemHealthCheckErrorElement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceitemHealthCheckErrorElementWithDefaults() *ServiceitemHealthCheckErrorElement {
	this := ServiceitemHealthCheckErrorElement{}
	var classId string = "serviceitem.HealthCheckErrorElement"
	this.ClassId = classId
	var objectType string = "serviceitem.HealthCheckErrorElement"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ServiceitemHealthCheckErrorElement) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ServiceitemHealthCheckErrorElement) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ServiceitemHealthCheckErrorElement) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ServiceitemHealthCheckErrorElement) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ServiceitemHealthCheckErrorElement) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ServiceitemHealthCheckErrorElement) SetObjectType(v string) {
	o.ObjectType = v
}

// GetElementType returns the ElementType field value if set, zero value otherwise.
func (o *ServiceitemHealthCheckErrorElement) GetElementType() string {
	if o == nil || o.ElementType == nil {
		var ret string
		return ret
	}
	return *o.ElementType
}

// GetElementTypeOk returns a tuple with the ElementType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceitemHealthCheckErrorElement) GetElementTypeOk() (*string, bool) {
	if o == nil || o.ElementType == nil {
		return nil, false
	}
	return o.ElementType, true
}

// HasElementType returns a boolean if a field has been set.
func (o *ServiceitemHealthCheckErrorElement) HasElementType() bool {
	if o != nil && o.ElementType != nil {
		return true
	}

	return false
}

// SetElementType gets a reference to the given string and assigns it to the ElementType field.
func (o *ServiceitemHealthCheckErrorElement) SetElementType(v string) {
	o.ElementType = &v
}

// GetElements returns the Elements field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceitemHealthCheckErrorElement) GetElements() []MoMoRef {
	if o == nil {
		var ret []MoMoRef
		return ret
	}
	return o.Elements
}

// GetElementsOk returns a tuple with the Elements field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceitemHealthCheckErrorElement) GetElementsOk() ([]MoMoRef, bool) {
	if o == nil || o.Elements == nil {
		return nil, false
	}
	return o.Elements, true
}

// HasElements returns a boolean if a field has been set.
func (o *ServiceitemHealthCheckErrorElement) HasElements() bool {
	if o != nil && o.Elements != nil {
		return true
	}

	return false
}

// SetElements gets a reference to the given []MoMoRef and assigns it to the Elements field.
func (o *ServiceitemHealthCheckErrorElement) SetElements(v []MoMoRef) {
	o.Elements = v
}

func (o ServiceitemHealthCheckErrorElement) MarshalJSON() ([]byte, error) {
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
	if o.ElementType != nil {
		toSerialize["ElementType"] = o.ElementType
	}
	if o.Elements != nil {
		toSerialize["Elements"] = o.Elements
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ServiceitemHealthCheckErrorElement) UnmarshalJSON(bytes []byte) (err error) {
	type ServiceitemHealthCheckErrorElementWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Type of the element of the service item. Examples are Server, Storage, UCS Fabric Interconnect etc.
		ElementType *string   `json:"ElementType,omitempty"`
		Elements    []MoMoRef `json:"Elements,omitempty"`
	}

	varServiceitemHealthCheckErrorElementWithoutEmbeddedStruct := ServiceitemHealthCheckErrorElementWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varServiceitemHealthCheckErrorElementWithoutEmbeddedStruct)
	if err == nil {
		varServiceitemHealthCheckErrorElement := _ServiceitemHealthCheckErrorElement{}
		varServiceitemHealthCheckErrorElement.ClassId = varServiceitemHealthCheckErrorElementWithoutEmbeddedStruct.ClassId
		varServiceitemHealthCheckErrorElement.ObjectType = varServiceitemHealthCheckErrorElementWithoutEmbeddedStruct.ObjectType
		varServiceitemHealthCheckErrorElement.ElementType = varServiceitemHealthCheckErrorElementWithoutEmbeddedStruct.ElementType
		varServiceitemHealthCheckErrorElement.Elements = varServiceitemHealthCheckErrorElementWithoutEmbeddedStruct.Elements
		*o = ServiceitemHealthCheckErrorElement(varServiceitemHealthCheckErrorElement)
	} else {
		return err
	}

	varServiceitemHealthCheckErrorElement := _ServiceitemHealthCheckErrorElement{}

	err = json.Unmarshal(bytes, &varServiceitemHealthCheckErrorElement)
	if err == nil {
		o.MoBaseComplexType = varServiceitemHealthCheckErrorElement.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ElementType")
		delete(additionalProperties, "Elements")

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

type NullableServiceitemHealthCheckErrorElement struct {
	value *ServiceitemHealthCheckErrorElement
	isSet bool
}

func (v NullableServiceitemHealthCheckErrorElement) Get() *ServiceitemHealthCheckErrorElement {
	return v.value
}

func (v *NullableServiceitemHealthCheckErrorElement) Set(val *ServiceitemHealthCheckErrorElement) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceitemHealthCheckErrorElement) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceitemHealthCheckErrorElement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceitemHealthCheckErrorElement(val *ServiceitemHealthCheckErrorElement) *NullableServiceitemHealthCheckErrorElement {
	return &NullableServiceitemHealthCheckErrorElement{value: val, isSet: true}
}

func (v NullableServiceitemHealthCheckErrorElement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceitemHealthCheckErrorElement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
