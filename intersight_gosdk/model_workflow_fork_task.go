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

// WorkflowForkTask A ForkTask is a control task that forks tasks for parallel execution in a workflow.
type WorkflowForkTask struct {
	WorkflowControlTask
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType  string   `json:"ObjectType"`
	ForkedTasks []string `json:"ForkedTasks,omitempty"`
	// Task name for the join control task that must follow a fork control task.
	JoinTask             *string `json:"JoinTask,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowForkTask WorkflowForkTask

// NewWorkflowForkTask instantiates a new WorkflowForkTask object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowForkTask(classId string, objectType string) *WorkflowForkTask {
	this := WorkflowForkTask{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowForkTaskWithDefaults instantiates a new WorkflowForkTask object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowForkTaskWithDefaults() *WorkflowForkTask {
	this := WorkflowForkTask{}
	var classId string = "workflow.ForkTask"
	this.ClassId = classId
	var objectType string = "workflow.ForkTask"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowForkTask) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowForkTask) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowForkTask) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowForkTask) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowForkTask) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowForkTask) SetObjectType(v string) {
	o.ObjectType = v
}

// GetForkedTasks returns the ForkedTasks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowForkTask) GetForkedTasks() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ForkedTasks
}

// GetForkedTasksOk returns a tuple with the ForkedTasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowForkTask) GetForkedTasksOk() ([]string, bool) {
	if o == nil || o.ForkedTasks == nil {
		return nil, false
	}
	return o.ForkedTasks, true
}

// HasForkedTasks returns a boolean if a field has been set.
func (o *WorkflowForkTask) HasForkedTasks() bool {
	if o != nil && o.ForkedTasks != nil {
		return true
	}

	return false
}

// SetForkedTasks gets a reference to the given []string and assigns it to the ForkedTasks field.
func (o *WorkflowForkTask) SetForkedTasks(v []string) {
	o.ForkedTasks = v
}

// GetJoinTask returns the JoinTask field value if set, zero value otherwise.
func (o *WorkflowForkTask) GetJoinTask() string {
	if o == nil || o.JoinTask == nil {
		var ret string
		return ret
	}
	return *o.JoinTask
}

// GetJoinTaskOk returns a tuple with the JoinTask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowForkTask) GetJoinTaskOk() (*string, bool) {
	if o == nil || o.JoinTask == nil {
		return nil, false
	}
	return o.JoinTask, true
}

// HasJoinTask returns a boolean if a field has been set.
func (o *WorkflowForkTask) HasJoinTask() bool {
	if o != nil && o.JoinTask != nil {
		return true
	}

	return false
}

// SetJoinTask gets a reference to the given string and assigns it to the JoinTask field.
func (o *WorkflowForkTask) SetJoinTask(v string) {
	o.JoinTask = &v
}

func (o WorkflowForkTask) MarshalJSON() ([]byte, error) {
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
	if o.ForkedTasks != nil {
		toSerialize["ForkedTasks"] = o.ForkedTasks
	}
	if o.JoinTask != nil {
		toSerialize["JoinTask"] = o.JoinTask
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowForkTask) UnmarshalJSON(bytes []byte) (err error) {
	type WorkflowForkTaskWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType  string   `json:"ObjectType"`
		ForkedTasks []string `json:"ForkedTasks,omitempty"`
		// Task name for the join control task that must follow a fork control task.
		JoinTask *string `json:"JoinTask,omitempty"`
	}

	varWorkflowForkTaskWithoutEmbeddedStruct := WorkflowForkTaskWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varWorkflowForkTaskWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowForkTask := _WorkflowForkTask{}
		varWorkflowForkTask.ClassId = varWorkflowForkTaskWithoutEmbeddedStruct.ClassId
		varWorkflowForkTask.ObjectType = varWorkflowForkTaskWithoutEmbeddedStruct.ObjectType
		varWorkflowForkTask.ForkedTasks = varWorkflowForkTaskWithoutEmbeddedStruct.ForkedTasks
		varWorkflowForkTask.JoinTask = varWorkflowForkTaskWithoutEmbeddedStruct.JoinTask
		*o = WorkflowForkTask(varWorkflowForkTask)
	} else {
		return err
	}

	varWorkflowForkTask := _WorkflowForkTask{}

	err = json.Unmarshal(bytes, &varWorkflowForkTask)
	if err == nil {
		o.WorkflowControlTask = varWorkflowForkTask.WorkflowControlTask
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ForkedTasks")
		delete(additionalProperties, "JoinTask")

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

type NullableWorkflowForkTask struct {
	value *WorkflowForkTask
	isSet bool
}

func (v NullableWorkflowForkTask) Get() *WorkflowForkTask {
	return v.value
}

func (v *NullableWorkflowForkTask) Set(val *WorkflowForkTask) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowForkTask) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowForkTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowForkTask(val *WorkflowForkTask) *NullableWorkflowForkTask {
	return &NullableWorkflowForkTask{value: val, isSet: true}
}

func (v NullableWorkflowForkTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowForkTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
