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
)

// WorkflowCancelableTypeAllOf Definition of the list of properties defined in 'workflow.CancelableType', excluding properties defined in parent classes.
type WorkflowCancelableTypeAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType       string   `json:"ObjectType"`
	CancelableStates []string `json:"CancelableStates,omitempty"`
	// When true the workflow can be cancelled. The action can be further restricted by the mode and cancelableStates properties.
	Enabled *bool `json:"Enabled,omitempty"`
	// Mode controls how the workflow can be canceled. * `ApiOnly` - The workflow can only be canceled via API call. * `All` - The workflow can be canceled from API or from the user interface.
	Mode                 *string `json:"Mode,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowCancelableTypeAllOf WorkflowCancelableTypeAllOf

// NewWorkflowCancelableTypeAllOf instantiates a new WorkflowCancelableTypeAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowCancelableTypeAllOf(classId string, objectType string) *WorkflowCancelableTypeAllOf {
	this := WorkflowCancelableTypeAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var enabled bool = true
	this.Enabled = &enabled
	var mode string = "ApiOnly"
	this.Mode = &mode
	return &this
}

// NewWorkflowCancelableTypeAllOfWithDefaults instantiates a new WorkflowCancelableTypeAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowCancelableTypeAllOfWithDefaults() *WorkflowCancelableTypeAllOf {
	this := WorkflowCancelableTypeAllOf{}
	var classId string = "workflow.CancelableType"
	this.ClassId = classId
	var objectType string = "workflow.CancelableType"
	this.ObjectType = objectType
	var enabled bool = true
	this.Enabled = &enabled
	var mode string = "ApiOnly"
	this.Mode = &mode
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowCancelableTypeAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowCancelableTypeAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowCancelableTypeAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowCancelableTypeAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowCancelableTypeAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowCancelableTypeAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCancelableStates returns the CancelableStates field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowCancelableTypeAllOf) GetCancelableStates() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.CancelableStates
}

// GetCancelableStatesOk returns a tuple with the CancelableStates field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowCancelableTypeAllOf) GetCancelableStatesOk() ([]string, bool) {
	if o == nil || o.CancelableStates == nil {
		return nil, false
	}
	return o.CancelableStates, true
}

// HasCancelableStates returns a boolean if a field has been set.
func (o *WorkflowCancelableTypeAllOf) HasCancelableStates() bool {
	if o != nil && o.CancelableStates != nil {
		return true
	}

	return false
}

// SetCancelableStates gets a reference to the given []string and assigns it to the CancelableStates field.
func (o *WorkflowCancelableTypeAllOf) SetCancelableStates(v []string) {
	o.CancelableStates = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *WorkflowCancelableTypeAllOf) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowCancelableTypeAllOf) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *WorkflowCancelableTypeAllOf) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *WorkflowCancelableTypeAllOf) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *WorkflowCancelableTypeAllOf) GetMode() string {
	if o == nil || o.Mode == nil {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowCancelableTypeAllOf) GetModeOk() (*string, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *WorkflowCancelableTypeAllOf) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *WorkflowCancelableTypeAllOf) SetMode(v string) {
	o.Mode = &v
}

func (o WorkflowCancelableTypeAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.CancelableStates != nil {
		toSerialize["CancelableStates"] = o.CancelableStates
	}
	if o.Enabled != nil {
		toSerialize["Enabled"] = o.Enabled
	}
	if o.Mode != nil {
		toSerialize["Mode"] = o.Mode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowCancelableTypeAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowCancelableTypeAllOf := _WorkflowCancelableTypeAllOf{}

	if err = json.Unmarshal(bytes, &varWorkflowCancelableTypeAllOf); err == nil {
		*o = WorkflowCancelableTypeAllOf(varWorkflowCancelableTypeAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CancelableStates")
		delete(additionalProperties, "Enabled")
		delete(additionalProperties, "Mode")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowCancelableTypeAllOf struct {
	value *WorkflowCancelableTypeAllOf
	isSet bool
}

func (v NullableWorkflowCancelableTypeAllOf) Get() *WorkflowCancelableTypeAllOf {
	return v.value
}

func (v *NullableWorkflowCancelableTypeAllOf) Set(val *WorkflowCancelableTypeAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowCancelableTypeAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowCancelableTypeAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowCancelableTypeAllOf(val *WorkflowCancelableTypeAllOf) *NullableWorkflowCancelableTypeAllOf {
	return &NullableWorkflowCancelableTypeAllOf{value: val, isSet: true}
}

func (v NullableWorkflowCancelableTypeAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowCancelableTypeAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
