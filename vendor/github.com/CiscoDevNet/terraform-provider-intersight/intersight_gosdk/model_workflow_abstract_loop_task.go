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
	"reflect"
	"strings"
)

// WorkflowAbstractLoopTask An abstract loop task which is used to model different types of loops like serial and parallel.
type WorkflowAbstractLoopTask struct {
	WorkflowControlTask
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Count value for the loop, this can be a static value defined as a constant at design time or can be a dynamic value defined as an expression that will evaluate to an integer value at execution time. Dynamic values for count must be specified as a template. For example, if a loop must run for a count which matches the length of a workflow input called StringArray, then the count must be specified using a template '{{ len .global.workflow.input.StringArray }}'. The count must be less than or equal to 100. If count is given as a dynamic value, and during execution time if count evaluates to be a value greater than 100, then the loop task will fail.
	Count *string `json:"Count,omitempty"`
	// Start task where the list of tasks will be executed multiple times based on the count or condition value.
	LoopStartTask *string `json:"LoopStartTask,omitempty"`
	// This specifies the name of the next task to run if all iterations of the loop task succeeds. The unique name given to the task instance within the workflow must be provided here. In a graph model, denotes an edge to another Task Node.
	OnSuccess            *string `json:"OnSuccess,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowAbstractLoopTask WorkflowAbstractLoopTask

// NewWorkflowAbstractLoopTask instantiates a new WorkflowAbstractLoopTask object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowAbstractLoopTask(classId string, objectType string) *WorkflowAbstractLoopTask {
	this := WorkflowAbstractLoopTask{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowAbstractLoopTaskWithDefaults instantiates a new WorkflowAbstractLoopTask object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowAbstractLoopTaskWithDefaults() *WorkflowAbstractLoopTask {
	this := WorkflowAbstractLoopTask{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowAbstractLoopTask) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowAbstractLoopTask) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowAbstractLoopTask) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowAbstractLoopTask) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowAbstractLoopTask) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowAbstractLoopTask) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *WorkflowAbstractLoopTask) GetCount() string {
	if o == nil || o.Count == nil {
		var ret string
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowAbstractLoopTask) GetCountOk() (*string, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *WorkflowAbstractLoopTask) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given string and assigns it to the Count field.
func (o *WorkflowAbstractLoopTask) SetCount(v string) {
	o.Count = &v
}

// GetLoopStartTask returns the LoopStartTask field value if set, zero value otherwise.
func (o *WorkflowAbstractLoopTask) GetLoopStartTask() string {
	if o == nil || o.LoopStartTask == nil {
		var ret string
		return ret
	}
	return *o.LoopStartTask
}

// GetLoopStartTaskOk returns a tuple with the LoopStartTask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowAbstractLoopTask) GetLoopStartTaskOk() (*string, bool) {
	if o == nil || o.LoopStartTask == nil {
		return nil, false
	}
	return o.LoopStartTask, true
}

// HasLoopStartTask returns a boolean if a field has been set.
func (o *WorkflowAbstractLoopTask) HasLoopStartTask() bool {
	if o != nil && o.LoopStartTask != nil {
		return true
	}

	return false
}

// SetLoopStartTask gets a reference to the given string and assigns it to the LoopStartTask field.
func (o *WorkflowAbstractLoopTask) SetLoopStartTask(v string) {
	o.LoopStartTask = &v
}

// GetOnSuccess returns the OnSuccess field value if set, zero value otherwise.
func (o *WorkflowAbstractLoopTask) GetOnSuccess() string {
	if o == nil || o.OnSuccess == nil {
		var ret string
		return ret
	}
	return *o.OnSuccess
}

// GetOnSuccessOk returns a tuple with the OnSuccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowAbstractLoopTask) GetOnSuccessOk() (*string, bool) {
	if o == nil || o.OnSuccess == nil {
		return nil, false
	}
	return o.OnSuccess, true
}

// HasOnSuccess returns a boolean if a field has been set.
func (o *WorkflowAbstractLoopTask) HasOnSuccess() bool {
	if o != nil && o.OnSuccess != nil {
		return true
	}

	return false
}

// SetOnSuccess gets a reference to the given string and assigns it to the OnSuccess field.
func (o *WorkflowAbstractLoopTask) SetOnSuccess(v string) {
	o.OnSuccess = &v
}

func (o WorkflowAbstractLoopTask) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedWorkflowControlTask, errWorkflowControlTask := json.Marshal(o.WorkflowControlTask)
	if errWorkflowControlTask != nil {
		return []byte{}, errWorkflowControlTask
	}
	errWorkflowControlTask = json.Unmarshal([]byte(serializedWorkflowControlTask), &toSerialize)
	if errWorkflowControlTask != nil {
		return []byte{}, errWorkflowControlTask
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Count != nil {
		toSerialize["Count"] = o.Count
	}
	if o.LoopStartTask != nil {
		toSerialize["LoopStartTask"] = o.LoopStartTask
	}
	if o.OnSuccess != nil {
		toSerialize["OnSuccess"] = o.OnSuccess
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowAbstractLoopTask) UnmarshalJSON(bytes []byte) (err error) {
	type WorkflowAbstractLoopTaskWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// Count value for the loop, this can be a static value defined as a constant at design time or can be a dynamic value defined as an expression that will evaluate to an integer value at execution time. Dynamic values for count must be specified as a template. For example, if a loop must run for a count which matches the length of a workflow input called StringArray, then the count must be specified using a template '{{ len .global.workflow.input.StringArray }}'. The count must be less than or equal to 100. If count is given as a dynamic value, and during execution time if count evaluates to be a value greater than 100, then the loop task will fail.
		Count *string `json:"Count,omitempty"`
		// Start task where the list of tasks will be executed multiple times based on the count or condition value.
		LoopStartTask *string `json:"LoopStartTask,omitempty"`
		// This specifies the name of the next task to run if all iterations of the loop task succeeds. The unique name given to the task instance within the workflow must be provided here. In a graph model, denotes an edge to another Task Node.
		OnSuccess *string `json:"OnSuccess,omitempty"`
	}

	varWorkflowAbstractLoopTaskWithoutEmbeddedStruct := WorkflowAbstractLoopTaskWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varWorkflowAbstractLoopTaskWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowAbstractLoopTask := _WorkflowAbstractLoopTask{}
		varWorkflowAbstractLoopTask.ClassId = varWorkflowAbstractLoopTaskWithoutEmbeddedStruct.ClassId
		varWorkflowAbstractLoopTask.ObjectType = varWorkflowAbstractLoopTaskWithoutEmbeddedStruct.ObjectType
		varWorkflowAbstractLoopTask.Count = varWorkflowAbstractLoopTaskWithoutEmbeddedStruct.Count
		varWorkflowAbstractLoopTask.LoopStartTask = varWorkflowAbstractLoopTaskWithoutEmbeddedStruct.LoopStartTask
		varWorkflowAbstractLoopTask.OnSuccess = varWorkflowAbstractLoopTaskWithoutEmbeddedStruct.OnSuccess
		*o = WorkflowAbstractLoopTask(varWorkflowAbstractLoopTask)
	} else {
		return err
	}

	varWorkflowAbstractLoopTask := _WorkflowAbstractLoopTask{}

	err = json.Unmarshal(bytes, &varWorkflowAbstractLoopTask)
	if err == nil {
		o.WorkflowControlTask = varWorkflowAbstractLoopTask.WorkflowControlTask
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Count")
		delete(additionalProperties, "LoopStartTask")
		delete(additionalProperties, "OnSuccess")

		// remove fields from embedded structs
		reflectWorkflowControlTask := reflect.ValueOf(o.WorkflowControlTask)
		for i := 0; i < reflectWorkflowControlTask.Type().NumField(); i++ {
			t := reflectWorkflowControlTask.Type().Field(i)

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

type NullableWorkflowAbstractLoopTask struct {
	value *WorkflowAbstractLoopTask
	isSet bool
}

func (v NullableWorkflowAbstractLoopTask) Get() *WorkflowAbstractLoopTask {
	return v.value
}

func (v *NullableWorkflowAbstractLoopTask) Set(val *WorkflowAbstractLoopTask) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowAbstractLoopTask) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowAbstractLoopTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowAbstractLoopTask(val *WorkflowAbstractLoopTask) *NullableWorkflowAbstractLoopTask {
	return &NullableWorkflowAbstractLoopTask{value: val, isSet: true}
}

func (v NullableWorkflowAbstractLoopTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowAbstractLoopTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
