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

// WorkflowTemplateFunctionMeta The function template object.
type WorkflowTemplateFunctionMeta struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string                   `json:"ObjectType"`
	Comments   NullableWorkflowComments `json:"Comments,omitempty"`
	Inputs     []WorkflowBaseDataType   `json:"Inputs,omitempty"`
	// The flag indicates whether a guided mode template is supported for it or not.
	IsGuidedModeSupported *bool `json:"IsGuidedModeSupported,omitempty"`
	// The template function name.
	Name                 *string                `json:"Name,omitempty"`
	Outputs              []WorkflowBaseDataType `json:"Outputs,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowTemplateFunctionMeta WorkflowTemplateFunctionMeta

// NewWorkflowTemplateFunctionMeta instantiates a new WorkflowTemplateFunctionMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowTemplateFunctionMeta(classId string, objectType string) *WorkflowTemplateFunctionMeta {
	this := WorkflowTemplateFunctionMeta{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowTemplateFunctionMetaWithDefaults instantiates a new WorkflowTemplateFunctionMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowTemplateFunctionMetaWithDefaults() *WorkflowTemplateFunctionMeta {
	this := WorkflowTemplateFunctionMeta{}
	var classId string = "workflow.TemplateFunctionMeta"
	this.ClassId = classId
	var objectType string = "workflow.TemplateFunctionMeta"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowTemplateFunctionMeta) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowTemplateFunctionMeta) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowTemplateFunctionMeta) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowTemplateFunctionMeta) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowTemplateFunctionMeta) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowTemplateFunctionMeta) SetObjectType(v string) {
	o.ObjectType = v
}

// GetComments returns the Comments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowTemplateFunctionMeta) GetComments() WorkflowComments {
	if o == nil || o.Comments.Get() == nil {
		var ret WorkflowComments
		return ret
	}
	return *o.Comments.Get()
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowTemplateFunctionMeta) GetCommentsOk() (*WorkflowComments, bool) {
	if o == nil {
		return nil, false
	}
	return o.Comments.Get(), o.Comments.IsSet()
}

// HasComments returns a boolean if a field has been set.
func (o *WorkflowTemplateFunctionMeta) HasComments() bool {
	if o != nil && o.Comments.IsSet() {
		return true
	}

	return false
}

// SetComments gets a reference to the given NullableWorkflowComments and assigns it to the Comments field.
func (o *WorkflowTemplateFunctionMeta) SetComments(v WorkflowComments) {
	o.Comments.Set(&v)
}

// SetCommentsNil sets the value for Comments to be an explicit nil
func (o *WorkflowTemplateFunctionMeta) SetCommentsNil() {
	o.Comments.Set(nil)
}

// UnsetComments ensures that no value is present for Comments, not even an explicit nil
func (o *WorkflowTemplateFunctionMeta) UnsetComments() {
	o.Comments.Unset()
}

// GetInputs returns the Inputs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowTemplateFunctionMeta) GetInputs() []WorkflowBaseDataType {
	if o == nil {
		var ret []WorkflowBaseDataType
		return ret
	}
	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowTemplateFunctionMeta) GetInputsOk() ([]WorkflowBaseDataType, bool) {
	if o == nil || o.Inputs == nil {
		return nil, false
	}
	return o.Inputs, true
}

// HasInputs returns a boolean if a field has been set.
func (o *WorkflowTemplateFunctionMeta) HasInputs() bool {
	if o != nil && o.Inputs != nil {
		return true
	}

	return false
}

// SetInputs gets a reference to the given []WorkflowBaseDataType and assigns it to the Inputs field.
func (o *WorkflowTemplateFunctionMeta) SetInputs(v []WorkflowBaseDataType) {
	o.Inputs = v
}

// GetIsGuidedModeSupported returns the IsGuidedModeSupported field value if set, zero value otherwise.
func (o *WorkflowTemplateFunctionMeta) GetIsGuidedModeSupported() bool {
	if o == nil || o.IsGuidedModeSupported == nil {
		var ret bool
		return ret
	}
	return *o.IsGuidedModeSupported
}

// GetIsGuidedModeSupportedOk returns a tuple with the IsGuidedModeSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTemplateFunctionMeta) GetIsGuidedModeSupportedOk() (*bool, bool) {
	if o == nil || o.IsGuidedModeSupported == nil {
		return nil, false
	}
	return o.IsGuidedModeSupported, true
}

// HasIsGuidedModeSupported returns a boolean if a field has been set.
func (o *WorkflowTemplateFunctionMeta) HasIsGuidedModeSupported() bool {
	if o != nil && o.IsGuidedModeSupported != nil {
		return true
	}

	return false
}

// SetIsGuidedModeSupported gets a reference to the given bool and assigns it to the IsGuidedModeSupported field.
func (o *WorkflowTemplateFunctionMeta) SetIsGuidedModeSupported(v bool) {
	o.IsGuidedModeSupported = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowTemplateFunctionMeta) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTemplateFunctionMeta) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowTemplateFunctionMeta) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowTemplateFunctionMeta) SetName(v string) {
	o.Name = &v
}

// GetOutputs returns the Outputs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowTemplateFunctionMeta) GetOutputs() []WorkflowBaseDataType {
	if o == nil {
		var ret []WorkflowBaseDataType
		return ret
	}
	return o.Outputs
}

// GetOutputsOk returns a tuple with the Outputs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowTemplateFunctionMeta) GetOutputsOk() ([]WorkflowBaseDataType, bool) {
	if o == nil || o.Outputs == nil {
		return nil, false
	}
	return o.Outputs, true
}

// HasOutputs returns a boolean if a field has been set.
func (o *WorkflowTemplateFunctionMeta) HasOutputs() bool {
	if o != nil && o.Outputs != nil {
		return true
	}

	return false
}

// SetOutputs gets a reference to the given []WorkflowBaseDataType and assigns it to the Outputs field.
func (o *WorkflowTemplateFunctionMeta) SetOutputs(v []WorkflowBaseDataType) {
	o.Outputs = v
}

func (o WorkflowTemplateFunctionMeta) MarshalJSON() ([]byte, error) {
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
	if o.Comments.IsSet() {
		toSerialize["Comments"] = o.Comments.Get()
	}
	if o.Inputs != nil {
		toSerialize["Inputs"] = o.Inputs
	}
	if o.IsGuidedModeSupported != nil {
		toSerialize["IsGuidedModeSupported"] = o.IsGuidedModeSupported
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Outputs != nil {
		toSerialize["Outputs"] = o.Outputs
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowTemplateFunctionMeta) UnmarshalJSON(bytes []byte) (err error) {
	type WorkflowTemplateFunctionMetaWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string                   `json:"ObjectType"`
		Comments   NullableWorkflowComments `json:"Comments,omitempty"`
		Inputs     []WorkflowBaseDataType   `json:"Inputs,omitempty"`
		// The flag indicates whether a guided mode template is supported for it or not.
		IsGuidedModeSupported *bool `json:"IsGuidedModeSupported,omitempty"`
		// The template function name.
		Name    *string                `json:"Name,omitempty"`
		Outputs []WorkflowBaseDataType `json:"Outputs,omitempty"`
	}

	varWorkflowTemplateFunctionMetaWithoutEmbeddedStruct := WorkflowTemplateFunctionMetaWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varWorkflowTemplateFunctionMetaWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowTemplateFunctionMeta := _WorkflowTemplateFunctionMeta{}
		varWorkflowTemplateFunctionMeta.ClassId = varWorkflowTemplateFunctionMetaWithoutEmbeddedStruct.ClassId
		varWorkflowTemplateFunctionMeta.ObjectType = varWorkflowTemplateFunctionMetaWithoutEmbeddedStruct.ObjectType
		varWorkflowTemplateFunctionMeta.Comments = varWorkflowTemplateFunctionMetaWithoutEmbeddedStruct.Comments
		varWorkflowTemplateFunctionMeta.Inputs = varWorkflowTemplateFunctionMetaWithoutEmbeddedStruct.Inputs
		varWorkflowTemplateFunctionMeta.IsGuidedModeSupported = varWorkflowTemplateFunctionMetaWithoutEmbeddedStruct.IsGuidedModeSupported
		varWorkflowTemplateFunctionMeta.Name = varWorkflowTemplateFunctionMetaWithoutEmbeddedStruct.Name
		varWorkflowTemplateFunctionMeta.Outputs = varWorkflowTemplateFunctionMetaWithoutEmbeddedStruct.Outputs
		*o = WorkflowTemplateFunctionMeta(varWorkflowTemplateFunctionMeta)
	} else {
		return err
	}

	varWorkflowTemplateFunctionMeta := _WorkflowTemplateFunctionMeta{}

	err = json.Unmarshal(bytes, &varWorkflowTemplateFunctionMeta)
	if err == nil {
		o.MoBaseMo = varWorkflowTemplateFunctionMeta.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Comments")
		delete(additionalProperties, "Inputs")
		delete(additionalProperties, "IsGuidedModeSupported")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Outputs")

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

type NullableWorkflowTemplateFunctionMeta struct {
	value *WorkflowTemplateFunctionMeta
	isSet bool
}

func (v NullableWorkflowTemplateFunctionMeta) Get() *WorkflowTemplateFunctionMeta {
	return v.value
}

func (v *NullableWorkflowTemplateFunctionMeta) Set(val *WorkflowTemplateFunctionMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowTemplateFunctionMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowTemplateFunctionMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowTemplateFunctionMeta(val *WorkflowTemplateFunctionMeta) *NullableWorkflowTemplateFunctionMeta {
	return &NullableWorkflowTemplateFunctionMeta{value: val, isSet: true}
}

func (v NullableWorkflowTemplateFunctionMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowTemplateFunctionMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
