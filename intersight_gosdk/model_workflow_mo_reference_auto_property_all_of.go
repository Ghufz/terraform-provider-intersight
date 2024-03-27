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
)

// WorkflowMoReferenceAutoPropertyAllOf Definition of the list of properties defined in 'workflow.MoReferenceAutoProperty', excluding properties defined in parent classes.
type WorkflowMoReferenceAutoPropertyAllOf struct {
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

type _WorkflowMoReferenceAutoPropertyAllOf WorkflowMoReferenceAutoPropertyAllOf

// NewWorkflowMoReferenceAutoPropertyAllOf instantiates a new WorkflowMoReferenceAutoPropertyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowMoReferenceAutoPropertyAllOf(classId string, objectType string) *WorkflowMoReferenceAutoPropertyAllOf {
	this := WorkflowMoReferenceAutoPropertyAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowMoReferenceAutoPropertyAllOfWithDefaults instantiates a new WorkflowMoReferenceAutoPropertyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowMoReferenceAutoPropertyAllOfWithDefaults() *WorkflowMoReferenceAutoPropertyAllOf {
	this := WorkflowMoReferenceAutoPropertyAllOf{}
	var classId string = "workflow.MoReferenceAutoProperty"
	this.ClassId = classId
	var objectType string = "workflow.MoReferenceAutoProperty"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowMoReferenceAutoPropertyAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowMoReferenceAutoPropertyAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowMoReferenceAutoPropertyAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowMoReferenceAutoPropertyAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowMoReferenceAutoPropertyAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowMoReferenceAutoPropertyAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDisplayAttributes returns the DisplayAttributes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowMoReferenceAutoPropertyAllOf) GetDisplayAttributes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.DisplayAttributes
}

// GetDisplayAttributesOk returns a tuple with the DisplayAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowMoReferenceAutoPropertyAllOf) GetDisplayAttributesOk() ([]string, bool) {
	if o == nil || o.DisplayAttributes == nil {
		return nil, false
	}
	return o.DisplayAttributes, true
}

// HasDisplayAttributes returns a boolean if a field has been set.
func (o *WorkflowMoReferenceAutoPropertyAllOf) HasDisplayAttributes() bool {
	if o != nil && o.DisplayAttributes != nil {
		return true
	}

	return false
}

// SetDisplayAttributes gets a reference to the given []string and assigns it to the DisplayAttributes field.
func (o *WorkflowMoReferenceAutoPropertyAllOf) SetDisplayAttributes(v []string) {
	o.DisplayAttributes = v
}

// GetFilters returns the Filters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowMoReferenceAutoPropertyAllOf) GetFilters() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowMoReferenceAutoPropertyAllOf) GetFiltersOk() ([]string, bool) {
	if o == nil || o.Filters == nil {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *WorkflowMoReferenceAutoPropertyAllOf) HasFilters() bool {
	if o != nil && o.Filters != nil {
		return true
	}

	return false
}

// SetFilters gets a reference to the given []string and assigns it to the Filters field.
func (o *WorkflowMoReferenceAutoPropertyAllOf) SetFilters(v []string) {
	o.Filters = v
}

// GetOrderBy returns the OrderBy field value if set, zero value otherwise.
func (o *WorkflowMoReferenceAutoPropertyAllOf) GetOrderBy() string {
	if o == nil || o.OrderBy == nil {
		var ret string
		return ret
	}
	return *o.OrderBy
}

// GetOrderByOk returns a tuple with the OrderBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowMoReferenceAutoPropertyAllOf) GetOrderByOk() (*string, bool) {
	if o == nil || o.OrderBy == nil {
		return nil, false
	}
	return o.OrderBy, true
}

// HasOrderBy returns a boolean if a field has been set.
func (o *WorkflowMoReferenceAutoPropertyAllOf) HasOrderBy() bool {
	if o != nil && o.OrderBy != nil {
		return true
	}

	return false
}

// SetOrderBy gets a reference to the given string and assigns it to the OrderBy field.
func (o *WorkflowMoReferenceAutoPropertyAllOf) SetOrderBy(v string) {
	o.OrderBy = &v
}

// GetRule returns the Rule field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowMoReferenceAutoPropertyAllOf) GetRule() WorkflowAbstractResourceSelector {
	if o == nil || o.Rule.Get() == nil {
		var ret WorkflowAbstractResourceSelector
		return ret
	}
	return *o.Rule.Get()
}

// GetRuleOk returns a tuple with the Rule field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowMoReferenceAutoPropertyAllOf) GetRuleOk() (*WorkflowAbstractResourceSelector, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rule.Get(), o.Rule.IsSet()
}

// HasRule returns a boolean if a field has been set.
func (o *WorkflowMoReferenceAutoPropertyAllOf) HasRule() bool {
	if o != nil && o.Rule.IsSet() {
		return true
	}

	return false
}

// SetRule gets a reference to the given NullableWorkflowAbstractResourceSelector and assigns it to the Rule field.
func (o *WorkflowMoReferenceAutoPropertyAllOf) SetRule(v WorkflowAbstractResourceSelector) {
	o.Rule.Set(&v)
}

// SetRuleNil sets the value for Rule to be an explicit nil
func (o *WorkflowMoReferenceAutoPropertyAllOf) SetRuleNil() {
	o.Rule.Set(nil)
}

// UnsetRule ensures that no value is present for Rule, not even an explicit nil
func (o *WorkflowMoReferenceAutoPropertyAllOf) UnsetRule() {
	o.Rule.Unset()
}

func (o WorkflowMoReferenceAutoPropertyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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

func (o *WorkflowMoReferenceAutoPropertyAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowMoReferenceAutoPropertyAllOf := _WorkflowMoReferenceAutoPropertyAllOf{}

	if err = json.Unmarshal(bytes, &varWorkflowMoReferenceAutoPropertyAllOf); err == nil {
		*o = WorkflowMoReferenceAutoPropertyAllOf(varWorkflowMoReferenceAutoPropertyAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DisplayAttributes")
		delete(additionalProperties, "Filters")
		delete(additionalProperties, "OrderBy")
		delete(additionalProperties, "Rule")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowMoReferenceAutoPropertyAllOf struct {
	value *WorkflowMoReferenceAutoPropertyAllOf
	isSet bool
}

func (v NullableWorkflowMoReferenceAutoPropertyAllOf) Get() *WorkflowMoReferenceAutoPropertyAllOf {
	return v.value
}

func (v *NullableWorkflowMoReferenceAutoPropertyAllOf) Set(val *WorkflowMoReferenceAutoPropertyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowMoReferenceAutoPropertyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowMoReferenceAutoPropertyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowMoReferenceAutoPropertyAllOf(val *WorkflowMoReferenceAutoPropertyAllOf) *NullableWorkflowMoReferenceAutoPropertyAllOf {
	return &NullableWorkflowMoReferenceAutoPropertyAllOf{value: val, isSet: true}
}

func (v NullableWorkflowMoReferenceAutoPropertyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowMoReferenceAutoPropertyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
