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

// WorkflowPendingDynamicWorkflowInfo Information for a pending Dynamic Workflow Instance before it is run.  Validation needs to be done on the dynamic workflow tasks before starting.  After it begins, it will be tracked with regular WorkflowInstance.
type WorkflowPendingDynamicWorkflowInfo struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The input data provided for workflow execution.
	Input interface{} `json:"Input,omitempty"`
	// A name for the pending dynamic workflow.
	Name *string `json:"Name,omitempty"`
	// The src is workflow owner service.
	Src *string `json:"Src,omitempty"`
	// The current status of the PendingDynamicWorkflowInfo. * `GatheringTasks` - Dynamic workflow is gathering tasks before workflow can start execution. * `Waiting` - Dynamic workflow is in waiting state and not yet started execution. * `RateLimit` - Dynamic workflow is rate limited and hasn't started execution.
	Status *string `json:"Status,omitempty"`
	// When set to true workflow engine will wait for a duplicate to finish before starting a new one.
	WaitOnDuplicate         *bool                                   `json:"WaitOnDuplicate,omitempty"`
	WorkflowActionTaskLists []WorkflowDynamicWorkflowActionTaskList `json:"WorkflowActionTaskLists,omitempty"`
	WorkflowCtx             NullableWorkflowWorkflowCtx             `json:"WorkflowCtx,omitempty"`
	// The metadata of the workflow.
	WorkflowMeta         interface{}                       `json:"WorkflowMeta,omitempty"`
	WorkflowInfo         *WorkflowWorkflowInfoRelationship `json:"WorkflowInfo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowPendingDynamicWorkflowInfo WorkflowPendingDynamicWorkflowInfo

// NewWorkflowPendingDynamicWorkflowInfo instantiates a new WorkflowPendingDynamicWorkflowInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowPendingDynamicWorkflowInfo(classId string, objectType string) *WorkflowPendingDynamicWorkflowInfo {
	this := WorkflowPendingDynamicWorkflowInfo{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowPendingDynamicWorkflowInfoWithDefaults instantiates a new WorkflowPendingDynamicWorkflowInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowPendingDynamicWorkflowInfoWithDefaults() *WorkflowPendingDynamicWorkflowInfo {
	this := WorkflowPendingDynamicWorkflowInfo{}
	var classId string = "workflow.PendingDynamicWorkflowInfo"
	this.ClassId = classId
	var objectType string = "workflow.PendingDynamicWorkflowInfo"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowPendingDynamicWorkflowInfo) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowPendingDynamicWorkflowInfo) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowPendingDynamicWorkflowInfo) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowPendingDynamicWorkflowInfo) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowPendingDynamicWorkflowInfo) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowPendingDynamicWorkflowInfo) SetObjectType(v string) {
	o.ObjectType = v
}

// GetInput returns the Input field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowPendingDynamicWorkflowInfo) GetInput() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Input
}

// GetInputOk returns a tuple with the Input field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowPendingDynamicWorkflowInfo) GetInputOk() (*interface{}, bool) {
	if o == nil || o.Input == nil {
		return nil, false
	}
	return &o.Input, true
}

// HasInput returns a boolean if a field has been set.
func (o *WorkflowPendingDynamicWorkflowInfo) HasInput() bool {
	if o != nil && o.Input != nil {
		return true
	}

	return false
}

// SetInput gets a reference to the given interface{} and assigns it to the Input field.
func (o *WorkflowPendingDynamicWorkflowInfo) SetInput(v interface{}) {
	o.Input = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowPendingDynamicWorkflowInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowPendingDynamicWorkflowInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowPendingDynamicWorkflowInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowPendingDynamicWorkflowInfo) SetName(v string) {
	o.Name = &v
}

// GetSrc returns the Src field value if set, zero value otherwise.
func (o *WorkflowPendingDynamicWorkflowInfo) GetSrc() string {
	if o == nil || o.Src == nil {
		var ret string
		return ret
	}
	return *o.Src
}

// GetSrcOk returns a tuple with the Src field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowPendingDynamicWorkflowInfo) GetSrcOk() (*string, bool) {
	if o == nil || o.Src == nil {
		return nil, false
	}
	return o.Src, true
}

// HasSrc returns a boolean if a field has been set.
func (o *WorkflowPendingDynamicWorkflowInfo) HasSrc() bool {
	if o != nil && o.Src != nil {
		return true
	}

	return false
}

// SetSrc gets a reference to the given string and assigns it to the Src field.
func (o *WorkflowPendingDynamicWorkflowInfo) SetSrc(v string) {
	o.Src = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *WorkflowPendingDynamicWorkflowInfo) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowPendingDynamicWorkflowInfo) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *WorkflowPendingDynamicWorkflowInfo) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *WorkflowPendingDynamicWorkflowInfo) SetStatus(v string) {
	o.Status = &v
}

// GetWaitOnDuplicate returns the WaitOnDuplicate field value if set, zero value otherwise.
func (o *WorkflowPendingDynamicWorkflowInfo) GetWaitOnDuplicate() bool {
	if o == nil || o.WaitOnDuplicate == nil {
		var ret bool
		return ret
	}
	return *o.WaitOnDuplicate
}

// GetWaitOnDuplicateOk returns a tuple with the WaitOnDuplicate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowPendingDynamicWorkflowInfo) GetWaitOnDuplicateOk() (*bool, bool) {
	if o == nil || o.WaitOnDuplicate == nil {
		return nil, false
	}
	return o.WaitOnDuplicate, true
}

// HasWaitOnDuplicate returns a boolean if a field has been set.
func (o *WorkflowPendingDynamicWorkflowInfo) HasWaitOnDuplicate() bool {
	if o != nil && o.WaitOnDuplicate != nil {
		return true
	}

	return false
}

// SetWaitOnDuplicate gets a reference to the given bool and assigns it to the WaitOnDuplicate field.
func (o *WorkflowPendingDynamicWorkflowInfo) SetWaitOnDuplicate(v bool) {
	o.WaitOnDuplicate = &v
}

// GetWorkflowActionTaskLists returns the WorkflowActionTaskLists field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowPendingDynamicWorkflowInfo) GetWorkflowActionTaskLists() []WorkflowDynamicWorkflowActionTaskList {
	if o == nil {
		var ret []WorkflowDynamicWorkflowActionTaskList
		return ret
	}
	return o.WorkflowActionTaskLists
}

// GetWorkflowActionTaskListsOk returns a tuple with the WorkflowActionTaskLists field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowPendingDynamicWorkflowInfo) GetWorkflowActionTaskListsOk() ([]WorkflowDynamicWorkflowActionTaskList, bool) {
	if o == nil || o.WorkflowActionTaskLists == nil {
		return nil, false
	}
	return o.WorkflowActionTaskLists, true
}

// HasWorkflowActionTaskLists returns a boolean if a field has been set.
func (o *WorkflowPendingDynamicWorkflowInfo) HasWorkflowActionTaskLists() bool {
	if o != nil && o.WorkflowActionTaskLists != nil {
		return true
	}

	return false
}

// SetWorkflowActionTaskLists gets a reference to the given []WorkflowDynamicWorkflowActionTaskList and assigns it to the WorkflowActionTaskLists field.
func (o *WorkflowPendingDynamicWorkflowInfo) SetWorkflowActionTaskLists(v []WorkflowDynamicWorkflowActionTaskList) {
	o.WorkflowActionTaskLists = v
}

// GetWorkflowCtx returns the WorkflowCtx field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowPendingDynamicWorkflowInfo) GetWorkflowCtx() WorkflowWorkflowCtx {
	if o == nil || o.WorkflowCtx.Get() == nil {
		var ret WorkflowWorkflowCtx
		return ret
	}
	return *o.WorkflowCtx.Get()
}

// GetWorkflowCtxOk returns a tuple with the WorkflowCtx field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowPendingDynamicWorkflowInfo) GetWorkflowCtxOk() (*WorkflowWorkflowCtx, bool) {
	if o == nil {
		return nil, false
	}
	return o.WorkflowCtx.Get(), o.WorkflowCtx.IsSet()
}

// HasWorkflowCtx returns a boolean if a field has been set.
func (o *WorkflowPendingDynamicWorkflowInfo) HasWorkflowCtx() bool {
	if o != nil && o.WorkflowCtx.IsSet() {
		return true
	}

	return false
}

// SetWorkflowCtx gets a reference to the given NullableWorkflowWorkflowCtx and assigns it to the WorkflowCtx field.
func (o *WorkflowPendingDynamicWorkflowInfo) SetWorkflowCtx(v WorkflowWorkflowCtx) {
	o.WorkflowCtx.Set(&v)
}

// SetWorkflowCtxNil sets the value for WorkflowCtx to be an explicit nil
func (o *WorkflowPendingDynamicWorkflowInfo) SetWorkflowCtxNil() {
	o.WorkflowCtx.Set(nil)
}

// UnsetWorkflowCtx ensures that no value is present for WorkflowCtx, not even an explicit nil
func (o *WorkflowPendingDynamicWorkflowInfo) UnsetWorkflowCtx() {
	o.WorkflowCtx.Unset()
}

// GetWorkflowMeta returns the WorkflowMeta field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowPendingDynamicWorkflowInfo) GetWorkflowMeta() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.WorkflowMeta
}

// GetWorkflowMetaOk returns a tuple with the WorkflowMeta field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowPendingDynamicWorkflowInfo) GetWorkflowMetaOk() (*interface{}, bool) {
	if o == nil || o.WorkflowMeta == nil {
		return nil, false
	}
	return &o.WorkflowMeta, true
}

// HasWorkflowMeta returns a boolean if a field has been set.
func (o *WorkflowPendingDynamicWorkflowInfo) HasWorkflowMeta() bool {
	if o != nil && o.WorkflowMeta != nil {
		return true
	}

	return false
}

// SetWorkflowMeta gets a reference to the given interface{} and assigns it to the WorkflowMeta field.
func (o *WorkflowPendingDynamicWorkflowInfo) SetWorkflowMeta(v interface{}) {
	o.WorkflowMeta = v
}

// GetWorkflowInfo returns the WorkflowInfo field value if set, zero value otherwise.
func (o *WorkflowPendingDynamicWorkflowInfo) GetWorkflowInfo() WorkflowWorkflowInfoRelationship {
	if o == nil || o.WorkflowInfo == nil {
		var ret WorkflowWorkflowInfoRelationship
		return ret
	}
	return *o.WorkflowInfo
}

// GetWorkflowInfoOk returns a tuple with the WorkflowInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowPendingDynamicWorkflowInfo) GetWorkflowInfoOk() (*WorkflowWorkflowInfoRelationship, bool) {
	if o == nil || o.WorkflowInfo == nil {
		return nil, false
	}
	return o.WorkflowInfo, true
}

// HasWorkflowInfo returns a boolean if a field has been set.
func (o *WorkflowPendingDynamicWorkflowInfo) HasWorkflowInfo() bool {
	if o != nil && o.WorkflowInfo != nil {
		return true
	}

	return false
}

// SetWorkflowInfo gets a reference to the given WorkflowWorkflowInfoRelationship and assigns it to the WorkflowInfo field.
func (o *WorkflowPendingDynamicWorkflowInfo) SetWorkflowInfo(v WorkflowWorkflowInfoRelationship) {
	o.WorkflowInfo = &v
}

func (o WorkflowPendingDynamicWorkflowInfo) MarshalJSON() ([]byte, error) {
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
	if o.Input != nil {
		toSerialize["Input"] = o.Input
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Src != nil {
		toSerialize["Src"] = o.Src
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.WaitOnDuplicate != nil {
		toSerialize["WaitOnDuplicate"] = o.WaitOnDuplicate
	}
	if o.WorkflowActionTaskLists != nil {
		toSerialize["WorkflowActionTaskLists"] = o.WorkflowActionTaskLists
	}
	if o.WorkflowCtx.IsSet() {
		toSerialize["WorkflowCtx"] = o.WorkflowCtx.Get()
	}
	if o.WorkflowMeta != nil {
		toSerialize["WorkflowMeta"] = o.WorkflowMeta
	}
	if o.WorkflowInfo != nil {
		toSerialize["WorkflowInfo"] = o.WorkflowInfo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowPendingDynamicWorkflowInfo) UnmarshalJSON(bytes []byte) (err error) {
	type WorkflowPendingDynamicWorkflowInfoWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The input data provided for workflow execution.
		Input interface{} `json:"Input,omitempty"`
		// A name for the pending dynamic workflow.
		Name *string `json:"Name,omitempty"`
		// The src is workflow owner service.
		Src *string `json:"Src,omitempty"`
		// The current status of the PendingDynamicWorkflowInfo. * `GatheringTasks` - Dynamic workflow is gathering tasks before workflow can start execution. * `Waiting` - Dynamic workflow is in waiting state and not yet started execution. * `RateLimit` - Dynamic workflow is rate limited and hasn't started execution.
		Status *string `json:"Status,omitempty"`
		// When set to true workflow engine will wait for a duplicate to finish before starting a new one.
		WaitOnDuplicate         *bool                                   `json:"WaitOnDuplicate,omitempty"`
		WorkflowActionTaskLists []WorkflowDynamicWorkflowActionTaskList `json:"WorkflowActionTaskLists,omitempty"`
		WorkflowCtx             NullableWorkflowWorkflowCtx             `json:"WorkflowCtx,omitempty"`
		// The metadata of the workflow.
		WorkflowMeta interface{}                       `json:"WorkflowMeta,omitempty"`
		WorkflowInfo *WorkflowWorkflowInfoRelationship `json:"WorkflowInfo,omitempty"`
	}

	varWorkflowPendingDynamicWorkflowInfoWithoutEmbeddedStruct := WorkflowPendingDynamicWorkflowInfoWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varWorkflowPendingDynamicWorkflowInfoWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowPendingDynamicWorkflowInfo := _WorkflowPendingDynamicWorkflowInfo{}
		varWorkflowPendingDynamicWorkflowInfo.ClassId = varWorkflowPendingDynamicWorkflowInfoWithoutEmbeddedStruct.ClassId
		varWorkflowPendingDynamicWorkflowInfo.ObjectType = varWorkflowPendingDynamicWorkflowInfoWithoutEmbeddedStruct.ObjectType
		varWorkflowPendingDynamicWorkflowInfo.Input = varWorkflowPendingDynamicWorkflowInfoWithoutEmbeddedStruct.Input
		varWorkflowPendingDynamicWorkflowInfo.Name = varWorkflowPendingDynamicWorkflowInfoWithoutEmbeddedStruct.Name
		varWorkflowPendingDynamicWorkflowInfo.Src = varWorkflowPendingDynamicWorkflowInfoWithoutEmbeddedStruct.Src
		varWorkflowPendingDynamicWorkflowInfo.Status = varWorkflowPendingDynamicWorkflowInfoWithoutEmbeddedStruct.Status
		varWorkflowPendingDynamicWorkflowInfo.WaitOnDuplicate = varWorkflowPendingDynamicWorkflowInfoWithoutEmbeddedStruct.WaitOnDuplicate
		varWorkflowPendingDynamicWorkflowInfo.WorkflowActionTaskLists = varWorkflowPendingDynamicWorkflowInfoWithoutEmbeddedStruct.WorkflowActionTaskLists
		varWorkflowPendingDynamicWorkflowInfo.WorkflowCtx = varWorkflowPendingDynamicWorkflowInfoWithoutEmbeddedStruct.WorkflowCtx
		varWorkflowPendingDynamicWorkflowInfo.WorkflowMeta = varWorkflowPendingDynamicWorkflowInfoWithoutEmbeddedStruct.WorkflowMeta
		varWorkflowPendingDynamicWorkflowInfo.WorkflowInfo = varWorkflowPendingDynamicWorkflowInfoWithoutEmbeddedStruct.WorkflowInfo
		*o = WorkflowPendingDynamicWorkflowInfo(varWorkflowPendingDynamicWorkflowInfo)
	} else {
		return err
	}

	varWorkflowPendingDynamicWorkflowInfo := _WorkflowPendingDynamicWorkflowInfo{}

	err = json.Unmarshal(bytes, &varWorkflowPendingDynamicWorkflowInfo)
	if err == nil {
		o.MoBaseMo = varWorkflowPendingDynamicWorkflowInfo.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Input")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Src")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "WaitOnDuplicate")
		delete(additionalProperties, "WorkflowActionTaskLists")
		delete(additionalProperties, "WorkflowCtx")
		delete(additionalProperties, "WorkflowMeta")
		delete(additionalProperties, "WorkflowInfo")

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

type NullableWorkflowPendingDynamicWorkflowInfo struct {
	value *WorkflowPendingDynamicWorkflowInfo
	isSet bool
}

func (v NullableWorkflowPendingDynamicWorkflowInfo) Get() *WorkflowPendingDynamicWorkflowInfo {
	return v.value
}

func (v *NullableWorkflowPendingDynamicWorkflowInfo) Set(val *WorkflowPendingDynamicWorkflowInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowPendingDynamicWorkflowInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowPendingDynamicWorkflowInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowPendingDynamicWorkflowInfo(val *WorkflowPendingDynamicWorkflowInfo) *NullableWorkflowPendingDynamicWorkflowInfo {
	return &NullableWorkflowPendingDynamicWorkflowInfo{value: val, isSet: true}
}

func (v NullableWorkflowPendingDynamicWorkflowInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowPendingDynamicWorkflowInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
