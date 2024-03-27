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

// IssueDefinition The definition of an issue which encompases the criteria for identifying when the issue exists, documentation of the detected issue, and actions to be taken by Intersight, e.g. raising of an alarm or triggering of a workflow to perform remediation.
type IssueDefinition struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string                 `json:"ObjectType"`
	Condition  NullableIssueCondition `json:"Condition,omitempty"`
	// A description of the issue which is common to all instances of the issue.
	Description *string `json:"Description,omitempty"`
	// An informational display name.
	Name *string `json:"Name,omitempty"`
	// An explanation of the likely causes of the detected issue.
	ProbableCause *string `json:"ProbableCause,omitempty"`
	// An explanation of the steps to perform to remediate the detected issue.
	Remediation           *string  `json:"Remediation,omitempty"`
	SystemClassifications []string `json:"SystemClassifications,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _IssueDefinition IssueDefinition

// NewIssueDefinition instantiates a new IssueDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIssueDefinition(classId string, objectType string) *IssueDefinition {
	this := IssueDefinition{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIssueDefinitionWithDefaults instantiates a new IssueDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIssueDefinitionWithDefaults() *IssueDefinition {
	this := IssueDefinition{}
	var classId string = "cond.AlarmDefinition"
	this.ClassId = classId
	var objectType string = "cond.AlarmDefinition"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IssueDefinition) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IssueDefinition) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IssueDefinition) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IssueDefinition) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IssueDefinition) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IssueDefinition) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCondition returns the Condition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssueDefinition) GetCondition() IssueCondition {
	if o == nil || o.Condition.Get() == nil {
		var ret IssueCondition
		return ret
	}
	return *o.Condition.Get()
}

// GetConditionOk returns a tuple with the Condition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssueDefinition) GetConditionOk() (*IssueCondition, bool) {
	if o == nil {
		return nil, false
	}
	return o.Condition.Get(), o.Condition.IsSet()
}

// HasCondition returns a boolean if a field has been set.
func (o *IssueDefinition) HasCondition() bool {
	if o != nil && o.Condition.IsSet() {
		return true
	}

	return false
}

// SetCondition gets a reference to the given NullableIssueCondition and assigns it to the Condition field.
func (o *IssueDefinition) SetCondition(v IssueCondition) {
	o.Condition.Set(&v)
}

// SetConditionNil sets the value for Condition to be an explicit nil
func (o *IssueDefinition) SetConditionNil() {
	o.Condition.Set(nil)
}

// UnsetCondition ensures that no value is present for Condition, not even an explicit nil
func (o *IssueDefinition) UnsetCondition() {
	o.Condition.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IssueDefinition) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueDefinition) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IssueDefinition) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IssueDefinition) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IssueDefinition) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueDefinition) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IssueDefinition) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IssueDefinition) SetName(v string) {
	o.Name = &v
}

// GetProbableCause returns the ProbableCause field value if set, zero value otherwise.
func (o *IssueDefinition) GetProbableCause() string {
	if o == nil || o.ProbableCause == nil {
		var ret string
		return ret
	}
	return *o.ProbableCause
}

// GetProbableCauseOk returns a tuple with the ProbableCause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueDefinition) GetProbableCauseOk() (*string, bool) {
	if o == nil || o.ProbableCause == nil {
		return nil, false
	}
	return o.ProbableCause, true
}

// HasProbableCause returns a boolean if a field has been set.
func (o *IssueDefinition) HasProbableCause() bool {
	if o != nil && o.ProbableCause != nil {
		return true
	}

	return false
}

// SetProbableCause gets a reference to the given string and assigns it to the ProbableCause field.
func (o *IssueDefinition) SetProbableCause(v string) {
	o.ProbableCause = &v
}

// GetRemediation returns the Remediation field value if set, zero value otherwise.
func (o *IssueDefinition) GetRemediation() string {
	if o == nil || o.Remediation == nil {
		var ret string
		return ret
	}
	return *o.Remediation
}

// GetRemediationOk returns a tuple with the Remediation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueDefinition) GetRemediationOk() (*string, bool) {
	if o == nil || o.Remediation == nil {
		return nil, false
	}
	return o.Remediation, true
}

// HasRemediation returns a boolean if a field has been set.
func (o *IssueDefinition) HasRemediation() bool {
	if o != nil && o.Remediation != nil {
		return true
	}

	return false
}

// SetRemediation gets a reference to the given string and assigns it to the Remediation field.
func (o *IssueDefinition) SetRemediation(v string) {
	o.Remediation = &v
}

// GetSystemClassifications returns the SystemClassifications field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssueDefinition) GetSystemClassifications() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SystemClassifications
}

// GetSystemClassificationsOk returns a tuple with the SystemClassifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssueDefinition) GetSystemClassificationsOk() ([]string, bool) {
	if o == nil || o.SystemClassifications == nil {
		return nil, false
	}
	return o.SystemClassifications, true
}

// HasSystemClassifications returns a boolean if a field has been set.
func (o *IssueDefinition) HasSystemClassifications() bool {
	if o != nil && o.SystemClassifications != nil {
		return true
	}

	return false
}

// SetSystemClassifications gets a reference to the given []string and assigns it to the SystemClassifications field.
func (o *IssueDefinition) SetSystemClassifications(v []string) {
	o.SystemClassifications = v
}

func (o IssueDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Condition.IsSet() {
		toSerialize["Condition"] = o.Condition.Get()
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.ProbableCause != nil {
		toSerialize["ProbableCause"] = o.ProbableCause
	}
	if o.Remediation != nil {
		toSerialize["Remediation"] = o.Remediation
	}
	if o.SystemClassifications != nil {
		toSerialize["SystemClassifications"] = o.SystemClassifications
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IssueDefinition) UnmarshalJSON(bytes []byte) (err error) {
	type IssueDefinitionWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string                 `json:"ObjectType"`
		Condition  NullableIssueCondition `json:"Condition,omitempty"`
		// A description of the issue which is common to all instances of the issue.
		Description *string `json:"Description,omitempty"`
		// An informational display name.
		Name *string `json:"Name,omitempty"`
		// An explanation of the likely causes of the detected issue.
		ProbableCause *string `json:"ProbableCause,omitempty"`
		// An explanation of the steps to perform to remediate the detected issue.
		Remediation           *string  `json:"Remediation,omitempty"`
		SystemClassifications []string `json:"SystemClassifications,omitempty"`
	}

	varIssueDefinitionWithoutEmbeddedStruct := IssueDefinitionWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varIssueDefinitionWithoutEmbeddedStruct)
	if err == nil {
		varIssueDefinition := _IssueDefinition{}
		varIssueDefinition.ClassId = varIssueDefinitionWithoutEmbeddedStruct.ClassId
		varIssueDefinition.ObjectType = varIssueDefinitionWithoutEmbeddedStruct.ObjectType
		varIssueDefinition.Condition = varIssueDefinitionWithoutEmbeddedStruct.Condition
		varIssueDefinition.Description = varIssueDefinitionWithoutEmbeddedStruct.Description
		varIssueDefinition.Name = varIssueDefinitionWithoutEmbeddedStruct.Name
		varIssueDefinition.ProbableCause = varIssueDefinitionWithoutEmbeddedStruct.ProbableCause
		varIssueDefinition.Remediation = varIssueDefinitionWithoutEmbeddedStruct.Remediation
		varIssueDefinition.SystemClassifications = varIssueDefinitionWithoutEmbeddedStruct.SystemClassifications
		*o = IssueDefinition(varIssueDefinition)
	} else {
		return err
	}

	varIssueDefinition := _IssueDefinition{}

	err = json.Unmarshal(bytes, &varIssueDefinition)
	if err == nil {
		o.MoBaseMo = varIssueDefinition.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Condition")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "ProbableCause")
		delete(additionalProperties, "Remediation")
		delete(additionalProperties, "SystemClassifications")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

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

type NullableIssueDefinition struct {
	value *IssueDefinition
	isSet bool
}

func (v NullableIssueDefinition) Get() *IssueDefinition {
	return v.value
}

func (v *NullableIssueDefinition) Set(val *IssueDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableIssueDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableIssueDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssueDefinition(val *IssueDefinition) *NullableIssueDefinition {
	return &NullableIssueDefinition{value: val, isSet: true}
}

func (v NullableIssueDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIssueDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
