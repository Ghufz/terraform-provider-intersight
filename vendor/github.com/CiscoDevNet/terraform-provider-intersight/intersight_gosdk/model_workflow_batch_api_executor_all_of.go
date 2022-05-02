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
)

// WorkflowBatchApiExecutorAllOf Definition of the list of properties defined in 'workflow.BatchApiExecutor', excluding properties defined in parent classes.
type WorkflowBatchApiExecutorAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                                    `json:"ObjectType"`
	ErrorResponseHandler *WorkflowErrorResponseHandlerRelationship `json:"ErrorResponseHandler,omitempty"`
	TaskDefinition       *WorkflowTaskDefinitionRelationship       `json:"TaskDefinition,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowBatchApiExecutorAllOf WorkflowBatchApiExecutorAllOf

// NewWorkflowBatchApiExecutorAllOf instantiates a new WorkflowBatchApiExecutorAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowBatchApiExecutorAllOf(classId string, objectType string) *WorkflowBatchApiExecutorAllOf {
	this := WorkflowBatchApiExecutorAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowBatchApiExecutorAllOfWithDefaults instantiates a new WorkflowBatchApiExecutorAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowBatchApiExecutorAllOfWithDefaults() *WorkflowBatchApiExecutorAllOf {
	this := WorkflowBatchApiExecutorAllOf{}
	var classId string = "workflow.BatchApiExecutor"
	this.ClassId = classId
	var objectType string = "workflow.BatchApiExecutor"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowBatchApiExecutorAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowBatchApiExecutorAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowBatchApiExecutorAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowBatchApiExecutorAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowBatchApiExecutorAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowBatchApiExecutorAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetErrorResponseHandler returns the ErrorResponseHandler field value if set, zero value otherwise.
func (o *WorkflowBatchApiExecutorAllOf) GetErrorResponseHandler() WorkflowErrorResponseHandlerRelationship {
	if o == nil || o.ErrorResponseHandler == nil {
		var ret WorkflowErrorResponseHandlerRelationship
		return ret
	}
	return *o.ErrorResponseHandler
}

// GetErrorResponseHandlerOk returns a tuple with the ErrorResponseHandler field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowBatchApiExecutorAllOf) GetErrorResponseHandlerOk() (*WorkflowErrorResponseHandlerRelationship, bool) {
	if o == nil || o.ErrorResponseHandler == nil {
		return nil, false
	}
	return o.ErrorResponseHandler, true
}

// HasErrorResponseHandler returns a boolean if a field has been set.
func (o *WorkflowBatchApiExecutorAllOf) HasErrorResponseHandler() bool {
	if o != nil && o.ErrorResponseHandler != nil {
		return true
	}

	return false
}

// SetErrorResponseHandler gets a reference to the given WorkflowErrorResponseHandlerRelationship and assigns it to the ErrorResponseHandler field.
func (o *WorkflowBatchApiExecutorAllOf) SetErrorResponseHandler(v WorkflowErrorResponseHandlerRelationship) {
	o.ErrorResponseHandler = &v
}

// GetTaskDefinition returns the TaskDefinition field value if set, zero value otherwise.
func (o *WorkflowBatchApiExecutorAllOf) GetTaskDefinition() WorkflowTaskDefinitionRelationship {
	if o == nil || o.TaskDefinition == nil {
		var ret WorkflowTaskDefinitionRelationship
		return ret
	}
	return *o.TaskDefinition
}

// GetTaskDefinitionOk returns a tuple with the TaskDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowBatchApiExecutorAllOf) GetTaskDefinitionOk() (*WorkflowTaskDefinitionRelationship, bool) {
	if o == nil || o.TaskDefinition == nil {
		return nil, false
	}
	return o.TaskDefinition, true
}

// HasTaskDefinition returns a boolean if a field has been set.
func (o *WorkflowBatchApiExecutorAllOf) HasTaskDefinition() bool {
	if o != nil && o.TaskDefinition != nil {
		return true
	}

	return false
}

// SetTaskDefinition gets a reference to the given WorkflowTaskDefinitionRelationship and assigns it to the TaskDefinition field.
func (o *WorkflowBatchApiExecutorAllOf) SetTaskDefinition(v WorkflowTaskDefinitionRelationship) {
	o.TaskDefinition = &v
}

func (o WorkflowBatchApiExecutorAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ErrorResponseHandler != nil {
		toSerialize["ErrorResponseHandler"] = o.ErrorResponseHandler
	}
	if o.TaskDefinition != nil {
		toSerialize["TaskDefinition"] = o.TaskDefinition
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowBatchApiExecutorAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowBatchApiExecutorAllOf := _WorkflowBatchApiExecutorAllOf{}

	if err = json.Unmarshal(bytes, &varWorkflowBatchApiExecutorAllOf); err == nil {
		*o = WorkflowBatchApiExecutorAllOf(varWorkflowBatchApiExecutorAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ErrorResponseHandler")
		delete(additionalProperties, "TaskDefinition")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowBatchApiExecutorAllOf struct {
	value *WorkflowBatchApiExecutorAllOf
	isSet bool
}

func (v NullableWorkflowBatchApiExecutorAllOf) Get() *WorkflowBatchApiExecutorAllOf {
	return v.value
}

func (v *NullableWorkflowBatchApiExecutorAllOf) Set(val *WorkflowBatchApiExecutorAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowBatchApiExecutorAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowBatchApiExecutorAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowBatchApiExecutorAllOf(val *WorkflowBatchApiExecutorAllOf) *NullableWorkflowBatchApiExecutorAllOf {
	return &NullableWorkflowBatchApiExecutorAllOf{value: val, isSet: true}
}

func (v NullableWorkflowBatchApiExecutorAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowBatchApiExecutorAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
