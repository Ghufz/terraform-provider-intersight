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

// WorkflowMoReferenceAutoProperty Captures all the display attributes and rules for selecting the resources.
type WorkflowMoReferenceAutoProperty struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType        string   `json:"ObjectType"`
	DisplayAttributes []string `json:"DisplayAttributes,omitempty"`
	Filters           []string `json:"Filters,omitempty"`
	// Determines  properties that are used to sort the collection of resources.
	OrderBy              *string                                  `json:"OrderBy,omitempty"`
	Rule                 NullableWorkflowAbstractResourceSelector `json:"Rule,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowMoReferenceAutoProperty WorkflowMoReferenceAutoProperty

// NewWorkflowMoReferenceAutoProperty instantiates a new WorkflowMoReferenceAutoProperty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowMoReferenceAutoProperty(classId string, objectType string) *WorkflowMoReferenceAutoProperty {
	this := WorkflowMoReferenceAutoProperty{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowMoReferenceAutoPropertyWithDefaults instantiates a new WorkflowMoReferenceAutoProperty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowMoReferenceAutoPropertyWithDefaults() *WorkflowMoReferenceAutoProperty {
	this := WorkflowMoReferenceAutoProperty{}
	var classId string = "workflow.MoReferenceAutoProperty"
	this.ClassId = classId
	var objectType string = "workflow.MoReferenceAutoProperty"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowMoReferenceAutoProperty) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowMoReferenceAutoProperty) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowMoReferenceAutoProperty) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowMoReferenceAutoProperty) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowMoReferenceAutoProperty) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowMoReferenceAutoProperty) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDisplayAttributes returns the DisplayAttributes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowMoReferenceAutoProperty) GetDisplayAttributes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.DisplayAttributes
}

// GetDisplayAttributesOk returns a tuple with the DisplayAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowMoReferenceAutoProperty) GetDisplayAttributesOk() ([]string, bool) {
	if o == nil || o.DisplayAttributes == nil {
		return nil, false
	}
	return o.DisplayAttributes, true
}

// HasDisplayAttributes returns a boolean if a field has been set.
func (o *WorkflowMoReferenceAutoProperty) HasDisplayAttributes() bool {
	if o != nil && o.DisplayAttributes != nil {
		return true
	}

	return false
}

// SetDisplayAttributes gets a reference to the given []string and assigns it to the DisplayAttributes field.
func (o *WorkflowMoReferenceAutoProperty) SetDisplayAttributes(v []string) {
	o.DisplayAttributes = v
}

// GetFilters returns the Filters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowMoReferenceAutoProperty) GetFilters() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowMoReferenceAutoProperty) GetFiltersOk() ([]string, bool) {
	if o == nil || o.Filters == nil {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *WorkflowMoReferenceAutoProperty) HasFilters() bool {
	if o != nil && o.Filters != nil {
		return true
	}

	return false
}

// SetFilters gets a reference to the given []string and assigns it to the Filters field.
func (o *WorkflowMoReferenceAutoProperty) SetFilters(v []string) {
	o.Filters = v
}

// GetOrderBy returns the OrderBy field value if set, zero value otherwise.
func (o *WorkflowMoReferenceAutoProperty) GetOrderBy() string {
	if o == nil || o.OrderBy == nil {
		var ret string
		return ret
	}
	return *o.OrderBy
}

// GetOrderByOk returns a tuple with the OrderBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowMoReferenceAutoProperty) GetOrderByOk() (*string, bool) {
	if o == nil || o.OrderBy == nil {
		return nil, false
	}
	return o.OrderBy, true
}

// HasOrderBy returns a boolean if a field has been set.
func (o *WorkflowMoReferenceAutoProperty) HasOrderBy() bool {
	if o != nil && o.OrderBy != nil {
		return true
	}

	return false
}

// SetOrderBy gets a reference to the given string and assigns it to the OrderBy field.
func (o *WorkflowMoReferenceAutoProperty) SetOrderBy(v string) {
	o.OrderBy = &v
}

// GetRule returns the Rule field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowMoReferenceAutoProperty) GetRule() WorkflowAbstractResourceSelector {
	if o == nil || o.Rule.Get() == nil {
		var ret WorkflowAbstractResourceSelector
		return ret
	}
	return *o.Rule.Get()
}

// GetRuleOk returns a tuple with the Rule field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowMoReferenceAutoProperty) GetRuleOk() (*WorkflowAbstractResourceSelector, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rule.Get(), o.Rule.IsSet()
}

// HasRule returns a boolean if a field has been set.
func (o *WorkflowMoReferenceAutoProperty) HasRule() bool {
	if o != nil && o.Rule.IsSet() {
		return true
	}

	return false
}

// SetRule gets a reference to the given NullableWorkflowAbstractResourceSelector and assigns it to the Rule field.
func (o *WorkflowMoReferenceAutoProperty) SetRule(v WorkflowAbstractResourceSelector) {
	o.Rule.Set(&v)
}

// SetRuleNil sets the value for Rule to be an explicit nil
func (o *WorkflowMoReferenceAutoProperty) SetRuleNil() {
	o.Rule.Set(nil)
}

// UnsetRule ensures that no value is present for Rule, not even an explicit nil
func (o *WorkflowMoReferenceAutoProperty) UnsetRule() {
	o.Rule.Unset()
}

func (o WorkflowMoReferenceAutoProperty) MarshalJSON() ([]byte, error) {
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
	if o.DisplayAttributes != nil {
		toSerialize["DisplayAttributes"] = o.DisplayAttributes
	}
	if o.Filters != nil {
		toSerialize["Filters"] = o.Filters
	}
	if o.OrderBy != nil {
		toSerialize["OrderBy"] = o.OrderBy
	}
	if o.Rule.IsSet() {
		toSerialize["Rule"] = o.Rule.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowMoReferenceAutoProperty) UnmarshalJSON(bytes []byte) (err error) {
	type WorkflowMoReferenceAutoPropertyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType        string   `json:"ObjectType"`
		DisplayAttributes []string `json:"DisplayAttributes,omitempty"`
		Filters           []string `json:"Filters,omitempty"`
		// Determines  properties that are used to sort the collection of resources.
		OrderBy *string                                  `json:"OrderBy,omitempty"`
		Rule    NullableWorkflowAbstractResourceSelector `json:"Rule,omitempty"`
	}

	varWorkflowMoReferenceAutoPropertyWithoutEmbeddedStruct := WorkflowMoReferenceAutoPropertyWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varWorkflowMoReferenceAutoPropertyWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowMoReferenceAutoProperty := _WorkflowMoReferenceAutoProperty{}
		varWorkflowMoReferenceAutoProperty.ClassId = varWorkflowMoReferenceAutoPropertyWithoutEmbeddedStruct.ClassId
		varWorkflowMoReferenceAutoProperty.ObjectType = varWorkflowMoReferenceAutoPropertyWithoutEmbeddedStruct.ObjectType
		varWorkflowMoReferenceAutoProperty.DisplayAttributes = varWorkflowMoReferenceAutoPropertyWithoutEmbeddedStruct.DisplayAttributes
		varWorkflowMoReferenceAutoProperty.Filters = varWorkflowMoReferenceAutoPropertyWithoutEmbeddedStruct.Filters
		varWorkflowMoReferenceAutoProperty.OrderBy = varWorkflowMoReferenceAutoPropertyWithoutEmbeddedStruct.OrderBy
		varWorkflowMoReferenceAutoProperty.Rule = varWorkflowMoReferenceAutoPropertyWithoutEmbeddedStruct.Rule
		*o = WorkflowMoReferenceAutoProperty(varWorkflowMoReferenceAutoProperty)
	} else {
		return err
	}

	varWorkflowMoReferenceAutoProperty := _WorkflowMoReferenceAutoProperty{}

	err = json.Unmarshal(bytes, &varWorkflowMoReferenceAutoProperty)
	if err == nil {
		o.MoBaseComplexType = varWorkflowMoReferenceAutoProperty.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DisplayAttributes")
		delete(additionalProperties, "Filters")
		delete(additionalProperties, "OrderBy")
		delete(additionalProperties, "Rule")

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

type NullableWorkflowMoReferenceAutoProperty struct {
	value *WorkflowMoReferenceAutoProperty
	isSet bool
}

func (v NullableWorkflowMoReferenceAutoProperty) Get() *WorkflowMoReferenceAutoProperty {
	return v.value
}

func (v *NullableWorkflowMoReferenceAutoProperty) Set(val *WorkflowMoReferenceAutoProperty) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowMoReferenceAutoProperty) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowMoReferenceAutoProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowMoReferenceAutoProperty(val *WorkflowMoReferenceAutoProperty) *NullableWorkflowMoReferenceAutoProperty {
	return &NullableWorkflowMoReferenceAutoProperty{value: val, isSet: true}
}

func (v NullableWorkflowMoReferenceAutoProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowMoReferenceAutoProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
