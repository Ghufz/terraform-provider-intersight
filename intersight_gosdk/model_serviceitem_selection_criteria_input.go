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

// ServiceitemSelectionCriteriaInput The policy based input is available at service item action definition and catalog item which is evaluated at runtime, SelectionCriteriaInput is used to capture the input which is extracted and possible condition is used to extract the value is captured.
type ServiceitemSelectionCriteriaInput struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType       string             `json:"ObjectType"`
	FilterConditions []ResourceSelector `json:"FilterConditions,omitempty"`
	// Name of the Policy Input.
	InputName *string `json:"InputName,omitempty"`
	// The value extracted from the filter conditions.
	InputValue           interface{} `json:"InputValue,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServiceitemSelectionCriteriaInput ServiceitemSelectionCriteriaInput

// NewServiceitemSelectionCriteriaInput instantiates a new ServiceitemSelectionCriteriaInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceitemSelectionCriteriaInput(classId string, objectType string) *ServiceitemSelectionCriteriaInput {
	this := ServiceitemSelectionCriteriaInput{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewServiceitemSelectionCriteriaInputWithDefaults instantiates a new ServiceitemSelectionCriteriaInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceitemSelectionCriteriaInputWithDefaults() *ServiceitemSelectionCriteriaInput {
	this := ServiceitemSelectionCriteriaInput{}
	var classId string = "serviceitem.SelectionCriteriaInput"
	this.ClassId = classId
	var objectType string = "serviceitem.SelectionCriteriaInput"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ServiceitemSelectionCriteriaInput) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ServiceitemSelectionCriteriaInput) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ServiceitemSelectionCriteriaInput) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ServiceitemSelectionCriteriaInput) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ServiceitemSelectionCriteriaInput) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ServiceitemSelectionCriteriaInput) SetObjectType(v string) {
	o.ObjectType = v
}

// GetFilterConditions returns the FilterConditions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceitemSelectionCriteriaInput) GetFilterConditions() []ResourceSelector {
	if o == nil {
		var ret []ResourceSelector
		return ret
	}
	return o.FilterConditions
}

// GetFilterConditionsOk returns a tuple with the FilterConditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceitemSelectionCriteriaInput) GetFilterConditionsOk() ([]ResourceSelector, bool) {
	if o == nil || o.FilterConditions == nil {
		return nil, false
	}
	return o.FilterConditions, true
}

// HasFilterConditions returns a boolean if a field has been set.
func (o *ServiceitemSelectionCriteriaInput) HasFilterConditions() bool {
	if o != nil && o.FilterConditions != nil {
		return true
	}

	return false
}

// SetFilterConditions gets a reference to the given []ResourceSelector and assigns it to the FilterConditions field.
func (o *ServiceitemSelectionCriteriaInput) SetFilterConditions(v []ResourceSelector) {
	o.FilterConditions = v
}

// GetInputName returns the InputName field value if set, zero value otherwise.
func (o *ServiceitemSelectionCriteriaInput) GetInputName() string {
	if o == nil || o.InputName == nil {
		var ret string
		return ret
	}
	return *o.InputName
}

// GetInputNameOk returns a tuple with the InputName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceitemSelectionCriteriaInput) GetInputNameOk() (*string, bool) {
	if o == nil || o.InputName == nil {
		return nil, false
	}
	return o.InputName, true
}

// HasInputName returns a boolean if a field has been set.
func (o *ServiceitemSelectionCriteriaInput) HasInputName() bool {
	if o != nil && o.InputName != nil {
		return true
	}

	return false
}

// SetInputName gets a reference to the given string and assigns it to the InputName field.
func (o *ServiceitemSelectionCriteriaInput) SetInputName(v string) {
	o.InputName = &v
}

// GetInputValue returns the InputValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceitemSelectionCriteriaInput) GetInputValue() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.InputValue
}

// GetInputValueOk returns a tuple with the InputValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceitemSelectionCriteriaInput) GetInputValueOk() (*interface{}, bool) {
	if o == nil || o.InputValue == nil {
		return nil, false
	}
	return &o.InputValue, true
}

// HasInputValue returns a boolean if a field has been set.
func (o *ServiceitemSelectionCriteriaInput) HasInputValue() bool {
	if o != nil && o.InputValue != nil {
		return true
	}

	return false
}

// SetInputValue gets a reference to the given interface{} and assigns it to the InputValue field.
func (o *ServiceitemSelectionCriteriaInput) SetInputValue(v interface{}) {
	o.InputValue = v
}

func (o ServiceitemSelectionCriteriaInput) MarshalJSON() ([]byte, error) {
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
	if o.FilterConditions != nil {
		toSerialize["FilterConditions"] = o.FilterConditions
	}
	if o.InputName != nil {
		toSerialize["InputName"] = o.InputName
	}
	if o.InputValue != nil {
		toSerialize["InputValue"] = o.InputValue
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ServiceitemSelectionCriteriaInput) UnmarshalJSON(bytes []byte) (err error) {
	type ServiceitemSelectionCriteriaInputWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType       string             `json:"ObjectType"`
		FilterConditions []ResourceSelector `json:"FilterConditions,omitempty"`
		// Name of the Policy Input.
		InputName *string `json:"InputName,omitempty"`
		// The value extracted from the filter conditions.
		InputValue interface{} `json:"InputValue,omitempty"`
	}

	varServiceitemSelectionCriteriaInputWithoutEmbeddedStruct := ServiceitemSelectionCriteriaInputWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varServiceitemSelectionCriteriaInputWithoutEmbeddedStruct)
	if err == nil {
		varServiceitemSelectionCriteriaInput := _ServiceitemSelectionCriteriaInput{}
		varServiceitemSelectionCriteriaInput.ClassId = varServiceitemSelectionCriteriaInputWithoutEmbeddedStruct.ClassId
		varServiceitemSelectionCriteriaInput.ObjectType = varServiceitemSelectionCriteriaInputWithoutEmbeddedStruct.ObjectType
		varServiceitemSelectionCriteriaInput.FilterConditions = varServiceitemSelectionCriteriaInputWithoutEmbeddedStruct.FilterConditions
		varServiceitemSelectionCriteriaInput.InputName = varServiceitemSelectionCriteriaInputWithoutEmbeddedStruct.InputName
		varServiceitemSelectionCriteriaInput.InputValue = varServiceitemSelectionCriteriaInputWithoutEmbeddedStruct.InputValue
		*o = ServiceitemSelectionCriteriaInput(varServiceitemSelectionCriteriaInput)
	} else {
		return err
	}

	varServiceitemSelectionCriteriaInput := _ServiceitemSelectionCriteriaInput{}

	err = json.Unmarshal(bytes, &varServiceitemSelectionCriteriaInput)
	if err == nil {
		o.MoBaseComplexType = varServiceitemSelectionCriteriaInput.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "FilterConditions")
		delete(additionalProperties, "InputName")
		delete(additionalProperties, "InputValue")

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

type NullableServiceitemSelectionCriteriaInput struct {
	value *ServiceitemSelectionCriteriaInput
	isSet bool
}

func (v NullableServiceitemSelectionCriteriaInput) Get() *ServiceitemSelectionCriteriaInput {
	return v.value
}

func (v *NullableServiceitemSelectionCriteriaInput) Set(val *ServiceitemSelectionCriteriaInput) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceitemSelectionCriteriaInput) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceitemSelectionCriteriaInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceitemSelectionCriteriaInput(val *ServiceitemSelectionCriteriaInput) *NullableServiceitemSelectionCriteriaInput {
	return &NullableServiceitemSelectionCriteriaInput{value: val, isSet: true}
}

func (v NullableServiceitemSelectionCriteriaInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceitemSelectionCriteriaInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
