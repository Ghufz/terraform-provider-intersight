/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-16342
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"time"
)

// BulkResultAllOf Definition of the list of properties defined in 'bulk.Result', excluding properties defined in parent classes.
type BulkResultAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The action that will be performed when an error occurs during processing of the request. * `Stop` - Stop the processing of the request after the first error. * `Proceed` - Proceed with the processing of the request even when an error occurs.
	ActionOnError *string `json:"ActionOnError,omitempty"`
	// The timestamp in UTC when the request processing is completed.
	CompletionTime *time.Time `json:"CompletionTime,omitempty"`
	// The number of subrequests received in this request.
	NumSubRequests *int64 `json:"NumSubRequests,omitempty"`
	// The individual request to be executed asynchronously.
	Request interface{} `json:"Request,omitempty"`
	// The timestamp in UTC when the request was received.
	RequestReceivedTime *time.Time `json:"RequestReceivedTime,omitempty"`
	// The processing status of the request. * `NotStarted` - Indicates that the request processing has not begun yet. * `ObjPresenceCheckInProgress` - Indicates that the object presence check is in progress for this request. * `ObjPresenceCheckComplete` - Indicates that the object presence check is complete. * `ExecutionInProgress` - Indicates that the request processing is in progress. * `Completed` - Indicates that the request processing has been completed successfully. * `CompletedWithErrors` - Indicates that the request processing has one or more failed subrequests. * `Failed` - Indicates that the processing of this request failed. * `TimedOut` - Indicates that the request processing timed out.
	Status *string `json:"Status,omitempty"`
	// The status message shows the error details in human readable format when the request goes to failed state. No additional information is shown for success case.
	StatusMessage *string `json:"StatusMessage,omitempty"`
	// The URI on which this async operation is being performed.
	Uri          *string                               `json:"Uri,omitempty"`
	MoCloner     *BulkMoClonerRelationship             `json:"MoCloner,omitempty"`
	MoDeepCloner *BulkMoDeepClonerRelationship         `json:"MoDeepCloner,omitempty"`
	MoMerger     *BulkMoMergerRelationship             `json:"MoMerger,omitempty"`
	Organization *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	// An array of relationships to bulkSubRequestObj resources.
	Results              []BulkSubRequestObjRelationship   `json:"Results,omitempty"`
	WorkflowInfo         *WorkflowWorkflowInfoRelationship `json:"WorkflowInfo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BulkResultAllOf BulkResultAllOf

// NewBulkResultAllOf instantiates a new BulkResultAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkResultAllOf(classId string, objectType string) *BulkResultAllOf {
	this := BulkResultAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewBulkResultAllOfWithDefaults instantiates a new BulkResultAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkResultAllOfWithDefaults() *BulkResultAllOf {
	this := BulkResultAllOf{}
	var classId string = "bulk.Result"
	this.ClassId = classId
	var objectType string = "bulk.Result"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *BulkResultAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *BulkResultAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *BulkResultAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *BulkResultAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *BulkResultAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *BulkResultAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetActionOnError returns the ActionOnError field value if set, zero value otherwise.
func (o *BulkResultAllOf) GetActionOnError() string {
	if o == nil || o.ActionOnError == nil {
		var ret string
		return ret
	}
	return *o.ActionOnError
}

// GetActionOnErrorOk returns a tuple with the ActionOnError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkResultAllOf) GetActionOnErrorOk() (*string, bool) {
	if o == nil || o.ActionOnError == nil {
		return nil, false
	}
	return o.ActionOnError, true
}

// HasActionOnError returns a boolean if a field has been set.
func (o *BulkResultAllOf) HasActionOnError() bool {
	if o != nil && o.ActionOnError != nil {
		return true
	}

	return false
}

// SetActionOnError gets a reference to the given string and assigns it to the ActionOnError field.
func (o *BulkResultAllOf) SetActionOnError(v string) {
	o.ActionOnError = &v
}

// GetCompletionTime returns the CompletionTime field value if set, zero value otherwise.
func (o *BulkResultAllOf) GetCompletionTime() time.Time {
	if o == nil || o.CompletionTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CompletionTime
}

// GetCompletionTimeOk returns a tuple with the CompletionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkResultAllOf) GetCompletionTimeOk() (*time.Time, bool) {
	if o == nil || o.CompletionTime == nil {
		return nil, false
	}
	return o.CompletionTime, true
}

// HasCompletionTime returns a boolean if a field has been set.
func (o *BulkResultAllOf) HasCompletionTime() bool {
	if o != nil && o.CompletionTime != nil {
		return true
	}

	return false
}

// SetCompletionTime gets a reference to the given time.Time and assigns it to the CompletionTime field.
func (o *BulkResultAllOf) SetCompletionTime(v time.Time) {
	o.CompletionTime = &v
}

// GetNumSubRequests returns the NumSubRequests field value if set, zero value otherwise.
func (o *BulkResultAllOf) GetNumSubRequests() int64 {
	if o == nil || o.NumSubRequests == nil {
		var ret int64
		return ret
	}
	return *o.NumSubRequests
}

// GetNumSubRequestsOk returns a tuple with the NumSubRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkResultAllOf) GetNumSubRequestsOk() (*int64, bool) {
	if o == nil || o.NumSubRequests == nil {
		return nil, false
	}
	return o.NumSubRequests, true
}

// HasNumSubRequests returns a boolean if a field has been set.
func (o *BulkResultAllOf) HasNumSubRequests() bool {
	if o != nil && o.NumSubRequests != nil {
		return true
	}

	return false
}

// SetNumSubRequests gets a reference to the given int64 and assigns it to the NumSubRequests field.
func (o *BulkResultAllOf) SetNumSubRequests(v int64) {
	o.NumSubRequests = &v
}

// GetRequest returns the Request field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkResultAllOf) GetRequest() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkResultAllOf) GetRequestOk() (*interface{}, bool) {
	if o == nil || o.Request == nil {
		return nil, false
	}
	return &o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *BulkResultAllOf) HasRequest() bool {
	if o != nil && o.Request != nil {
		return true
	}

	return false
}

// SetRequest gets a reference to the given interface{} and assigns it to the Request field.
func (o *BulkResultAllOf) SetRequest(v interface{}) {
	o.Request = v
}

// GetRequestReceivedTime returns the RequestReceivedTime field value if set, zero value otherwise.
func (o *BulkResultAllOf) GetRequestReceivedTime() time.Time {
	if o == nil || o.RequestReceivedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.RequestReceivedTime
}

// GetRequestReceivedTimeOk returns a tuple with the RequestReceivedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkResultAllOf) GetRequestReceivedTimeOk() (*time.Time, bool) {
	if o == nil || o.RequestReceivedTime == nil {
		return nil, false
	}
	return o.RequestReceivedTime, true
}

// HasRequestReceivedTime returns a boolean if a field has been set.
func (o *BulkResultAllOf) HasRequestReceivedTime() bool {
	if o != nil && o.RequestReceivedTime != nil {
		return true
	}

	return false
}

// SetRequestReceivedTime gets a reference to the given time.Time and assigns it to the RequestReceivedTime field.
func (o *BulkResultAllOf) SetRequestReceivedTime(v time.Time) {
	o.RequestReceivedTime = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BulkResultAllOf) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkResultAllOf) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BulkResultAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *BulkResultAllOf) SetStatus(v string) {
	o.Status = &v
}

// GetStatusMessage returns the StatusMessage field value if set, zero value otherwise.
func (o *BulkResultAllOf) GetStatusMessage() string {
	if o == nil || o.StatusMessage == nil {
		var ret string
		return ret
	}
	return *o.StatusMessage
}

// GetStatusMessageOk returns a tuple with the StatusMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkResultAllOf) GetStatusMessageOk() (*string, bool) {
	if o == nil || o.StatusMessage == nil {
		return nil, false
	}
	return o.StatusMessage, true
}

// HasStatusMessage returns a boolean if a field has been set.
func (o *BulkResultAllOf) HasStatusMessage() bool {
	if o != nil && o.StatusMessage != nil {
		return true
	}

	return false
}

// SetStatusMessage gets a reference to the given string and assigns it to the StatusMessage field.
func (o *BulkResultAllOf) SetStatusMessage(v string) {
	o.StatusMessage = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *BulkResultAllOf) GetUri() string {
	if o == nil || o.Uri == nil {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkResultAllOf) GetUriOk() (*string, bool) {
	if o == nil || o.Uri == nil {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *BulkResultAllOf) HasUri() bool {
	if o != nil && o.Uri != nil {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *BulkResultAllOf) SetUri(v string) {
	o.Uri = &v
}

// GetMoCloner returns the MoCloner field value if set, zero value otherwise.
func (o *BulkResultAllOf) GetMoCloner() BulkMoClonerRelationship {
	if o == nil || o.MoCloner == nil {
		var ret BulkMoClonerRelationship
		return ret
	}
	return *o.MoCloner
}

// GetMoClonerOk returns a tuple with the MoCloner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkResultAllOf) GetMoClonerOk() (*BulkMoClonerRelationship, bool) {
	if o == nil || o.MoCloner == nil {
		return nil, false
	}
	return o.MoCloner, true
}

// HasMoCloner returns a boolean if a field has been set.
func (o *BulkResultAllOf) HasMoCloner() bool {
	if o != nil && o.MoCloner != nil {
		return true
	}

	return false
}

// SetMoCloner gets a reference to the given BulkMoClonerRelationship and assigns it to the MoCloner field.
func (o *BulkResultAllOf) SetMoCloner(v BulkMoClonerRelationship) {
	o.MoCloner = &v
}

// GetMoDeepCloner returns the MoDeepCloner field value if set, zero value otherwise.
func (o *BulkResultAllOf) GetMoDeepCloner() BulkMoDeepClonerRelationship {
	if o == nil || o.MoDeepCloner == nil {
		var ret BulkMoDeepClonerRelationship
		return ret
	}
	return *o.MoDeepCloner
}

// GetMoDeepClonerOk returns a tuple with the MoDeepCloner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkResultAllOf) GetMoDeepClonerOk() (*BulkMoDeepClonerRelationship, bool) {
	if o == nil || o.MoDeepCloner == nil {
		return nil, false
	}
	return o.MoDeepCloner, true
}

// HasMoDeepCloner returns a boolean if a field has been set.
func (o *BulkResultAllOf) HasMoDeepCloner() bool {
	if o != nil && o.MoDeepCloner != nil {
		return true
	}

	return false
}

// SetMoDeepCloner gets a reference to the given BulkMoDeepClonerRelationship and assigns it to the MoDeepCloner field.
func (o *BulkResultAllOf) SetMoDeepCloner(v BulkMoDeepClonerRelationship) {
	o.MoDeepCloner = &v
}

// GetMoMerger returns the MoMerger field value if set, zero value otherwise.
func (o *BulkResultAllOf) GetMoMerger() BulkMoMergerRelationship {
	if o == nil || o.MoMerger == nil {
		var ret BulkMoMergerRelationship
		return ret
	}
	return *o.MoMerger
}

// GetMoMergerOk returns a tuple with the MoMerger field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkResultAllOf) GetMoMergerOk() (*BulkMoMergerRelationship, bool) {
	if o == nil || o.MoMerger == nil {
		return nil, false
	}
	return o.MoMerger, true
}

// HasMoMerger returns a boolean if a field has been set.
func (o *BulkResultAllOf) HasMoMerger() bool {
	if o != nil && o.MoMerger != nil {
		return true
	}

	return false
}

// SetMoMerger gets a reference to the given BulkMoMergerRelationship and assigns it to the MoMerger field.
func (o *BulkResultAllOf) SetMoMerger(v BulkMoMergerRelationship) {
	o.MoMerger = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *BulkResultAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkResultAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *BulkResultAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *BulkResultAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetResults returns the Results field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkResultAllOf) GetResults() []BulkSubRequestObjRelationship {
	if o == nil {
		var ret []BulkSubRequestObjRelationship
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkResultAllOf) GetResultsOk() ([]BulkSubRequestObjRelationship, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *BulkResultAllOf) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []BulkSubRequestObjRelationship and assigns it to the Results field.
func (o *BulkResultAllOf) SetResults(v []BulkSubRequestObjRelationship) {
	o.Results = v
}

// GetWorkflowInfo returns the WorkflowInfo field value if set, zero value otherwise.
func (o *BulkResultAllOf) GetWorkflowInfo() WorkflowWorkflowInfoRelationship {
	if o == nil || o.WorkflowInfo == nil {
		var ret WorkflowWorkflowInfoRelationship
		return ret
	}
	return *o.WorkflowInfo
}

// GetWorkflowInfoOk returns a tuple with the WorkflowInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkResultAllOf) GetWorkflowInfoOk() (*WorkflowWorkflowInfoRelationship, bool) {
	if o == nil || o.WorkflowInfo == nil {
		return nil, false
	}
	return o.WorkflowInfo, true
}

// HasWorkflowInfo returns a boolean if a field has been set.
func (o *BulkResultAllOf) HasWorkflowInfo() bool {
	if o != nil && o.WorkflowInfo != nil {
		return true
	}

	return false
}

// SetWorkflowInfo gets a reference to the given WorkflowWorkflowInfoRelationship and assigns it to the WorkflowInfo field.
func (o *BulkResultAllOf) SetWorkflowInfo(v WorkflowWorkflowInfoRelationship) {
	o.WorkflowInfo = &v
}

func (o BulkResultAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ActionOnError != nil {
		toSerialize["ActionOnError"] = o.ActionOnError
	}
	if o.CompletionTime != nil {
		toSerialize["CompletionTime"] = o.CompletionTime
	}
	if o.NumSubRequests != nil {
		toSerialize["NumSubRequests"] = o.NumSubRequests
	}
	if o.Request != nil {
		toSerialize["Request"] = o.Request
	}
	if o.RequestReceivedTime != nil {
		toSerialize["RequestReceivedTime"] = o.RequestReceivedTime
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.StatusMessage != nil {
		toSerialize["StatusMessage"] = o.StatusMessage
	}
	if o.Uri != nil {
		toSerialize["Uri"] = o.Uri
	}
	if o.MoCloner != nil {
		toSerialize["MoCloner"] = o.MoCloner
	}
	if o.MoDeepCloner != nil {
		toSerialize["MoDeepCloner"] = o.MoDeepCloner
	}
	if o.MoMerger != nil {
		toSerialize["MoMerger"] = o.MoMerger
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.Results != nil {
		toSerialize["Results"] = o.Results
	}
	if o.WorkflowInfo != nil {
		toSerialize["WorkflowInfo"] = o.WorkflowInfo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BulkResultAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varBulkResultAllOf := _BulkResultAllOf{}

	if err = json.Unmarshal(bytes, &varBulkResultAllOf); err == nil {
		*o = BulkResultAllOf(varBulkResultAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ActionOnError")
		delete(additionalProperties, "CompletionTime")
		delete(additionalProperties, "NumSubRequests")
		delete(additionalProperties, "Request")
		delete(additionalProperties, "RequestReceivedTime")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "StatusMessage")
		delete(additionalProperties, "Uri")
		delete(additionalProperties, "MoCloner")
		delete(additionalProperties, "MoDeepCloner")
		delete(additionalProperties, "MoMerger")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "Results")
		delete(additionalProperties, "WorkflowInfo")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBulkResultAllOf struct {
	value *BulkResultAllOf
	isSet bool
}

func (v NullableBulkResultAllOf) Get() *BulkResultAllOf {
	return v.value
}

func (v *NullableBulkResultAllOf) Set(val *BulkResultAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkResultAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkResultAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkResultAllOf(val *BulkResultAllOf) *NullableBulkResultAllOf {
	return &NullableBulkResultAllOf{value: val, isSet: true}
}

func (v NullableBulkResultAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkResultAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
