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

// WorkflowJoinTask A JoinTask is a control task that must follow a fork task and specify all the fork tasks that must complete and join before the worfklow can proceed to the task specified in the OnSuccess transition.
type WorkflowJoinTask struct {
	WorkflowControlTask
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType  string   `json:"ObjectType"`
	JoinOnTasks []string `json:"JoinOnTasks,omitempty"`
	// Name of the next task to run if all fork path specified in the JoinOnTasks list succeeds which is the unique name given to the task instance within the workflow. In a graph model, denotes an edge to another Task Node.
	OnSuccess            *string `json:"OnSuccess,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowJoinTask WorkflowJoinTask

// NewWorkflowJoinTask instantiates a new WorkflowJoinTask object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowJoinTask(classId string, objectType string) *WorkflowJoinTask {
	this := WorkflowJoinTask{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowJoinTaskWithDefaults instantiates a new WorkflowJoinTask object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowJoinTaskWithDefaults() *WorkflowJoinTask {
	this := WorkflowJoinTask{}
	var classId string = "workflow.JoinTask"
	this.ClassId = classId
	var objectType string = "workflow.JoinTask"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowJoinTask) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowJoinTask) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowJoinTask) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowJoinTask) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowJoinTask) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowJoinTask) SetObjectType(v string) {
	o.ObjectType = v
}

// GetJoinOnTasks returns the JoinOnTasks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowJoinTask) GetJoinOnTasks() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.JoinOnTasks
}

// GetJoinOnTasksOk returns a tuple with the JoinOnTasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowJoinTask) GetJoinOnTasksOk() ([]string, bool) {
	if o == nil || o.JoinOnTasks == nil {
		return nil, false
	}
	return o.JoinOnTasks, true
}

// HasJoinOnTasks returns a boolean if a field has been set.
func (o *WorkflowJoinTask) HasJoinOnTasks() bool {
	if o != nil && o.JoinOnTasks != nil {
		return true
	}

	return false
}

// SetJoinOnTasks gets a reference to the given []string and assigns it to the JoinOnTasks field.
func (o *WorkflowJoinTask) SetJoinOnTasks(v []string) {
	o.JoinOnTasks = v
}

// GetOnSuccess returns the OnSuccess field value if set, zero value otherwise.
func (o *WorkflowJoinTask) GetOnSuccess() string {
	if o == nil || o.OnSuccess == nil {
		var ret string
		return ret
	}
	return *o.OnSuccess
}

// GetOnSuccessOk returns a tuple with the OnSuccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowJoinTask) GetOnSuccessOk() (*string, bool) {
	if o == nil || o.OnSuccess == nil {
		return nil, false
	}
	return o.OnSuccess, true
}

// HasOnSuccess returns a boolean if a field has been set.
func (o *WorkflowJoinTask) HasOnSuccess() bool {
	if o != nil && o.OnSuccess != nil {
		return true
	}

	return false
}

// SetOnSuccess gets a reference to the given string and assigns it to the OnSuccess field.
func (o *WorkflowJoinTask) SetOnSuccess(v string) {
	o.OnSuccess = &v
}

func (o WorkflowJoinTask) MarshalJSON() ([]byte, error) {
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
	if o.JoinOnTasks != nil {
		toSerialize["JoinOnTasks"] = o.JoinOnTasks
	}
	if o.OnSuccess != nil {
		toSerialize["OnSuccess"] = o.OnSuccess
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowJoinTask) UnmarshalJSON(bytes []byte) (err error) {
	type WorkflowJoinTaskWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType  string   `json:"ObjectType"`
		JoinOnTasks []string `json:"JoinOnTasks,omitempty"`
		// Name of the next task to run if all fork path specified in the JoinOnTasks list succeeds which is the unique name given to the task instance within the workflow. In a graph model, denotes an edge to another Task Node.
		OnSuccess *string `json:"OnSuccess,omitempty"`
	}

	varWorkflowJoinTaskWithoutEmbeddedStruct := WorkflowJoinTaskWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varWorkflowJoinTaskWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowJoinTask := _WorkflowJoinTask{}
		varWorkflowJoinTask.ClassId = varWorkflowJoinTaskWithoutEmbeddedStruct.ClassId
		varWorkflowJoinTask.ObjectType = varWorkflowJoinTaskWithoutEmbeddedStruct.ObjectType
		varWorkflowJoinTask.JoinOnTasks = varWorkflowJoinTaskWithoutEmbeddedStruct.JoinOnTasks
		varWorkflowJoinTask.OnSuccess = varWorkflowJoinTaskWithoutEmbeddedStruct.OnSuccess
		*o = WorkflowJoinTask(varWorkflowJoinTask)
	} else {
		return err
	}

	varWorkflowJoinTask := _WorkflowJoinTask{}

	err = json.Unmarshal(bytes, &varWorkflowJoinTask)
	if err == nil {
		o.WorkflowControlTask = varWorkflowJoinTask.WorkflowControlTask
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "JoinOnTasks")
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

type NullableWorkflowJoinTask struct {
	value *WorkflowJoinTask
	isSet bool
}

func (v NullableWorkflowJoinTask) Get() *WorkflowJoinTask {
	return v.value
}

func (v *NullableWorkflowJoinTask) Set(val *WorkflowJoinTask) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowJoinTask) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowJoinTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowJoinTask(val *WorkflowJoinTask) *NullableWorkflowJoinTask {
	return &NullableWorkflowJoinTask{value: val, isSet: true}
}

func (v NullableWorkflowJoinTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowJoinTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
