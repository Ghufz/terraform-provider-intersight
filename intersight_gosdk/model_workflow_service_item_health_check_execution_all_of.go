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
	"time"
)

// WorkflowServiceItemHealthCheckExecutionAllOf Definition of the list of properties defined in 'workflow.ServiceItemHealthCheckExecution', excluding properties defined in parent classes.
type WorkflowServiceItemHealthCheckExecutionAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType    string                               `json:"ObjectType"`
	ErrorElements []ServiceitemHealthCheckErrorElement `json:"ErrorElements,omitempty"`
	// Error summary of a health check execution failure.
	ErrorSummary *string `json:"ErrorSummary,omitempty"`
	// Timestamp of the last passed health check execution.
	LastPassedTimestamp *time.Time `json:"LastPassedTimestamp,omitempty"`
	// Health check execution result. * `Unknown` - Indicates that the health check results could not be determined. * `Pass` - Indicates that the health check has passed. * `Fail` - Indicates that the health check has failed. * `Warning` - Indicates that the health check completed with a warning. * `NotApplicable` - Indicates that the health check is either unsupported, or not applicable for the service item.
	Result *string `json:"Result,omitempty"`
	// A brief summary of health check execution result.
	Summary *string `json:"Summary,omitempty"`
	// Status of the workflow that is executed as a part of health check execution.
	WorkflowStatus        *string                                               `json:"WorkflowStatus,omitempty"`
	HealthCheckDefinition *WorkflowServiceItemHealthCheckDefinitionRelationship `json:"HealthCheckDefinition,omitempty"`
	ServiceItemInstance   *WorkflowServiceItemInstanceRelationship              `json:"ServiceItemInstance,omitempty"`
	WorkflowInfo          *WorkflowWorkflowInfoRelationship                     `json:"WorkflowInfo,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _WorkflowServiceItemHealthCheckExecutionAllOf WorkflowServiceItemHealthCheckExecutionAllOf

// NewWorkflowServiceItemHealthCheckExecutionAllOf instantiates a new WorkflowServiceItemHealthCheckExecutionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowServiceItemHealthCheckExecutionAllOf(classId string, objectType string) *WorkflowServiceItemHealthCheckExecutionAllOf {
	this := WorkflowServiceItemHealthCheckExecutionAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowServiceItemHealthCheckExecutionAllOfWithDefaults instantiates a new WorkflowServiceItemHealthCheckExecutionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowServiceItemHealthCheckExecutionAllOfWithDefaults() *WorkflowServiceItemHealthCheckExecutionAllOf {
	this := WorkflowServiceItemHealthCheckExecutionAllOf{}
	var classId string = "workflow.ServiceItemHealthCheckExecution"
	this.ClassId = classId
	var objectType string = "workflow.ServiceItemHealthCheckExecution"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowServiceItemHealthCheckExecutionAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemHealthCheckExecutionAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowServiceItemHealthCheckExecutionAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowServiceItemHealthCheckExecutionAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemHealthCheckExecutionAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowServiceItemHealthCheckExecutionAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetErrorElements returns the ErrorElements field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowServiceItemHealthCheckExecutionAllOf) GetErrorElements() []ServiceitemHealthCheckErrorElement {
	if o == nil {
		var ret []ServiceitemHealthCheckErrorElement
		return ret
	}
	return o.ErrorElements
}

// GetErrorElementsOk returns a tuple with the ErrorElements field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowServiceItemHealthCheckExecutionAllOf) GetErrorElementsOk() ([]ServiceitemHealthCheckErrorElement, bool) {
	if o == nil || o.ErrorElements == nil {
		return nil, false
	}
	return o.ErrorElements, true
}

// HasErrorElements returns a boolean if a field has been set.
func (o *WorkflowServiceItemHealthCheckExecutionAllOf) HasErrorElements() bool {
	if o != nil && o.ErrorElements != nil {
		return true
	}

	return false
}

// SetErrorElements gets a reference to the given []ServiceitemHealthCheckErrorElement and assigns it to the ErrorElements field.
func (o *WorkflowServiceItemHealthCheckExecutionAllOf) SetErrorElements(v []ServiceitemHealthCheckErrorElement) {
	o.ErrorElements = v
}

// GetErrorSummary returns the ErrorSummary field value if set, zero value otherwise.
func (o *WorkflowServiceItemHealthCheckExecutionAllOf) GetErrorSummary() string {
	if o == nil || o.ErrorSummary == nil {
		var ret string
		return ret
	}
	return *o.ErrorSummary
}

// GetErrorSummaryOk returns a tuple with the ErrorSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemHealthCheckExecutionAllOf) GetErrorSummaryOk() (*string, bool) {
	if o == nil || o.ErrorSummary == nil {
		return nil, false
	}
	return o.ErrorSummary, true
}

// HasErrorSummary returns a boolean if a field has been set.
func (o *WorkflowServiceItemHealthCheckExecutionAllOf) HasErrorSummary() bool {
	if o != nil && o.ErrorSummary != nil {
		return true
	}

	return false
}

// SetErrorSummary gets a reference to the given string and assigns it to the ErrorSummary field.
func (o *WorkflowServiceItemHealthCheckExecutionAllOf) SetErrorSummary(v string) {
	o.ErrorSummary = &v
}

// GetLastPassedTimestamp returns the LastPassedTimestamp field value if set, zero value otherwise.
func (o *WorkflowServiceItemHealthCheckExecutionAllOf) GetLastPassedTimestamp() time.Time {
	if o == nil || o.LastPassedTimestamp == nil {
		var ret time.Time
		return ret
	}
	return *o.LastPassedTimestamp
}

// GetLastPassedTimestampOk returns a tuple with the LastPassedTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemHealthCheckExecutionAllOf) GetLastPassedTimestampOk() (*time.Time, bool) {
	if o == nil || o.LastPassedTimestamp == nil {
		return nil, false
	}
	return o.LastPassedTimestamp, true
}

// HasLastPassedTimestamp returns a boolean if a field has been set.
func (o *WorkflowServiceItemHealthCheckExecutionAllOf) HasLastPassedTimestamp() bool {
	if o != nil && o.LastPassedTimestamp != nil {
		return true
	}

	return false
}

// SetLastPassedTimestamp gets a reference to the given time.Time and assigns it to the LastPassedTimestamp field.
func (o *WorkflowServiceItemHealthCheckExecutionAllOf) SetLastPassedTimestamp(v time.Time) {
	o.LastPassedTimestamp = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *WorkflowServiceItemHealthCheckExecutionAllOf) GetResult() string {
	if o == nil || o.Result == nil {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemHealthCheckExecutionAllOf) GetResultOk() (*string, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *WorkflowServiceItemHealthCheckExecutionAllOf) HasResult() bool {
	if o != nil && o.Result != nil {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *WorkflowServiceItemHealthCheckExecutionAllOf) SetResult(v string) {
	o.Result = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *WorkflowServiceItemHealthCheckExecutionAllOf) GetSummary() string {
	if o == nil || o.Summary == nil {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemHealthCheckExecutionAllOf) GetSummaryOk() (*string, bool) {
	if o == nil || o.Summary == nil {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *WorkflowServiceItemHealthCheckExecutionAllOf) HasSummary() bool {
	if o != nil && o.Summary != nil {
		return true
	}

	return false
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *WorkflowServiceItemHealthCheckExecutionAllOf) SetSummary(v string) {
	o.Summary = &v
}

// GetWorkflowStatus returns the WorkflowStatus field value if set, zero value otherwise.
func (o *WorkflowServiceItemHealthCheckExecutionAllOf) GetWorkflowStatus() string {
	if o == nil || o.WorkflowStatus == nil {
		var ret string
		return ret
	}
	return *o.WorkflowStatus
}

// GetWorkflowStatusOk returns a tuple with the WorkflowStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemHealthCheckExecutionAllOf) GetWorkflowStatusOk() (*string, bool) {
	if o == nil || o.WorkflowStatus == nil {
		return nil, false
	}
	return o.WorkflowStatus, true
}

// HasWorkflowStatus returns a boolean if a field has been set.
func (o *WorkflowServiceItemHealthCheckExecutionAllOf) HasWorkflowStatus() bool {
	if o != nil && o.WorkflowStatus != nil {
		return true
	}

	return false
}

// SetWorkflowStatus gets a reference to the given string and assigns it to the WorkflowStatus field.
func (o *WorkflowServiceItemHealthCheckExecutionAllOf) SetWorkflowStatus(v string) {
	o.WorkflowStatus = &v
}

// GetHealthCheckDefinition returns the HealthCheckDefinition field value if set, zero value otherwise.
func (o *WorkflowServiceItemHealthCheckExecutionAllOf) GetHealthCheckDefinition() WorkflowServiceItemHealthCheckDefinitionRelationship {
	if o == nil || o.HealthCheckDefinition == nil {
		var ret WorkflowServiceItemHealthCheckDefinitionRelationship
		return ret
	}
	return *o.HealthCheckDefinition
}

// GetHealthCheckDefinitionOk returns a tuple with the HealthCheckDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemHealthCheckExecutionAllOf) GetHealthCheckDefinitionOk() (*WorkflowServiceItemHealthCheckDefinitionRelationship, bool) {
	if o == nil || o.HealthCheckDefinition == nil {
		return nil, false
	}
	return o.HealthCheckDefinition, true
}

// HasHealthCheckDefinition returns a boolean if a field has been set.
func (o *WorkflowServiceItemHealthCheckExecutionAllOf) HasHealthCheckDefinition() bool {
	if o != nil && o.HealthCheckDefinition != nil {
		return true
	}

	return false
}

// SetHealthCheckDefinition gets a reference to the given WorkflowServiceItemHealthCheckDefinitionRelationship and assigns it to the HealthCheckDefinition field.
func (o *WorkflowServiceItemHealthCheckExecutionAllOf) SetHealthCheckDefinition(v WorkflowServiceItemHealthCheckDefinitionRelationship) {
	o.HealthCheckDefinition = &v
}

// GetServiceItemInstance returns the ServiceItemInstance field value if set, zero value otherwise.
func (o *WorkflowServiceItemHealthCheckExecutionAllOf) GetServiceItemInstance() WorkflowServiceItemInstanceRelationship {
	if o == nil || o.ServiceItemInstance == nil {
		var ret WorkflowServiceItemInstanceRelationship
		return ret
	}
	return *o.ServiceItemInstance
}

// GetServiceItemInstanceOk returns a tuple with the ServiceItemInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemHealthCheckExecutionAllOf) GetServiceItemInstanceOk() (*WorkflowServiceItemInstanceRelationship, bool) {
	if o == nil || o.ServiceItemInstance == nil {
		return nil, false
	}
	return o.ServiceItemInstance, true
}

// HasServiceItemInstance returns a boolean if a field has been set.
func (o *WorkflowServiceItemHealthCheckExecutionAllOf) HasServiceItemInstance() bool {
	if o != nil && o.ServiceItemInstance != nil {
		return true
	}

	return false
}

// SetServiceItemInstance gets a reference to the given WorkflowServiceItemInstanceRelationship and assigns it to the ServiceItemInstance field.
func (o *WorkflowServiceItemHealthCheckExecutionAllOf) SetServiceItemInstance(v WorkflowServiceItemInstanceRelationship) {
	o.ServiceItemInstance = &v
}

// GetWorkflowInfo returns the WorkflowInfo field value if set, zero value otherwise.
func (o *WorkflowServiceItemHealthCheckExecutionAllOf) GetWorkflowInfo() WorkflowWorkflowInfoRelationship {
	if o == nil || o.WorkflowInfo == nil {
		var ret WorkflowWorkflowInfoRelationship
		return ret
	}
	return *o.WorkflowInfo
}

// GetWorkflowInfoOk returns a tuple with the WorkflowInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemHealthCheckExecutionAllOf) GetWorkflowInfoOk() (*WorkflowWorkflowInfoRelationship, bool) {
	if o == nil || o.WorkflowInfo == nil {
		return nil, false
	}
	return o.WorkflowInfo, true
}

// HasWorkflowInfo returns a boolean if a field has been set.
func (o *WorkflowServiceItemHealthCheckExecutionAllOf) HasWorkflowInfo() bool {
	if o != nil && o.WorkflowInfo != nil {
		return true
	}

	return false
}

// SetWorkflowInfo gets a reference to the given WorkflowWorkflowInfoRelationship and assigns it to the WorkflowInfo field.
func (o *WorkflowServiceItemHealthCheckExecutionAllOf) SetWorkflowInfo(v WorkflowWorkflowInfoRelationship) {
	o.WorkflowInfo = &v
}

func (o WorkflowServiceItemHealthCheckExecutionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ErrorElements != nil {
		toSerialize["ErrorElements"] = o.ErrorElements
	}
	if o.ErrorSummary != nil {
		toSerialize["ErrorSummary"] = o.ErrorSummary
	}
	if o.LastPassedTimestamp != nil {
		toSerialize["LastPassedTimestamp"] = o.LastPassedTimestamp
	}
	if o.Result != nil {
		toSerialize["Result"] = o.Result
	}
	if o.Summary != nil {
		toSerialize["Summary"] = o.Summary
	}
	if o.WorkflowStatus != nil {
		toSerialize["WorkflowStatus"] = o.WorkflowStatus
	}
	if o.HealthCheckDefinition != nil {
		toSerialize["HealthCheckDefinition"] = o.HealthCheckDefinition
	}
	if o.ServiceItemInstance != nil {
		toSerialize["ServiceItemInstance"] = o.ServiceItemInstance
	}
	if o.WorkflowInfo != nil {
		toSerialize["WorkflowInfo"] = o.WorkflowInfo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowServiceItemHealthCheckExecutionAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowServiceItemHealthCheckExecutionAllOf := _WorkflowServiceItemHealthCheckExecutionAllOf{}

	if err = json.Unmarshal(bytes, &varWorkflowServiceItemHealthCheckExecutionAllOf); err == nil {
		*o = WorkflowServiceItemHealthCheckExecutionAllOf(varWorkflowServiceItemHealthCheckExecutionAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ErrorElements")
		delete(additionalProperties, "ErrorSummary")
		delete(additionalProperties, "LastPassedTimestamp")
		delete(additionalProperties, "Result")
		delete(additionalProperties, "Summary")
		delete(additionalProperties, "WorkflowStatus")
		delete(additionalProperties, "HealthCheckDefinition")
		delete(additionalProperties, "ServiceItemInstance")
		delete(additionalProperties, "WorkflowInfo")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowServiceItemHealthCheckExecutionAllOf struct {
	value *WorkflowServiceItemHealthCheckExecutionAllOf
	isSet bool
}

func (v NullableWorkflowServiceItemHealthCheckExecutionAllOf) Get() *WorkflowServiceItemHealthCheckExecutionAllOf {
	return v.value
}

func (v *NullableWorkflowServiceItemHealthCheckExecutionAllOf) Set(val *WorkflowServiceItemHealthCheckExecutionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowServiceItemHealthCheckExecutionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowServiceItemHealthCheckExecutionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowServiceItemHealthCheckExecutionAllOf(val *WorkflowServiceItemHealthCheckExecutionAllOf) *NullableWorkflowServiceItemHealthCheckExecutionAllOf {
	return &NullableWorkflowServiceItemHealthCheckExecutionAllOf{value: val, isSet: true}
}

func (v NullableWorkflowServiceItemHealthCheckExecutionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowServiceItemHealthCheckExecutionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
