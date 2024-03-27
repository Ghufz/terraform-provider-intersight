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
)

// AaaAbstractAuditRecordAllOf Definition of the list of properties defined in 'aaa.AbstractAuditRecord', excluding properties defined in parent classes.
type AaaAbstractAuditRecordAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// The operation that was performed on this Managed Object. The event is a compound string that includes the CRUD operation such as Create, Modify, Delete, and a string representing the Managed Object type.
	Event *string `json:"Event,omitempty"`
	// The user-friendly names of the changed MO.
	MoDisplayNames interface{} `json:"MoDisplayNames,omitempty"`
	// The object type of the REST resource that was created, modified or deleted.
	MoType *string `json:"MoType,omitempty"`
	// The Moid of the REST resource that was created, modified or deleted.
	ObjectMoid *string `json:"ObjectMoid,omitempty"`
	// The body of the REST request that was received from a client to create or modify a REST resource, represented as a JSON document.
	Request interface{} `json:"Request,omitempty"`
	// The trace id of the request that was used to create, modify or delete a REST resource. A trace id is a unique identifier for one particular REST request. It may be used for troubleshooting purpose by the Intersight technical support team.
	TraceId   *string              `json:"TraceId,omitempty"`
	UserAgent NullableAaaUserAgent `json:"UserAgent,omitempty"`
	// The raw, string representation of the user agent of the request from the user-agent http request header.
	UserAgentString      *string `json:"UserAgentString,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AaaAbstractAuditRecordAllOf AaaAbstractAuditRecordAllOf

// NewAaaAbstractAuditRecordAllOf instantiates a new AaaAbstractAuditRecordAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAaaAbstractAuditRecordAllOf(classId string, objectType string) *AaaAbstractAuditRecordAllOf {
	this := AaaAbstractAuditRecordAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAaaAbstractAuditRecordAllOfWithDefaults instantiates a new AaaAbstractAuditRecordAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAaaAbstractAuditRecordAllOfWithDefaults() *AaaAbstractAuditRecordAllOf {
	this := AaaAbstractAuditRecordAllOf{}
	var classId string = "aaa.AuditRecord"
	this.ClassId = classId
	var objectType string = "aaa.AuditRecord"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AaaAbstractAuditRecordAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AaaAbstractAuditRecordAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AaaAbstractAuditRecordAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AaaAbstractAuditRecordAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AaaAbstractAuditRecordAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AaaAbstractAuditRecordAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEvent returns the Event field value if set, zero value otherwise.
func (o *AaaAbstractAuditRecordAllOf) GetEvent() string {
	if o == nil || o.Event == nil {
		var ret string
		return ret
	}
	return *o.Event
}

// GetEventOk returns a tuple with the Event field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AaaAbstractAuditRecordAllOf) GetEventOk() (*string, bool) {
	if o == nil || o.Event == nil {
		return nil, false
	}
	return o.Event, true
}

// HasEvent returns a boolean if a field has been set.
func (o *AaaAbstractAuditRecordAllOf) HasEvent() bool {
	if o != nil && o.Event != nil {
		return true
	}

	return false
}

// SetEvent gets a reference to the given string and assigns it to the Event field.
func (o *AaaAbstractAuditRecordAllOf) SetEvent(v string) {
	o.Event = &v
}

// GetMoDisplayNames returns the MoDisplayNames field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AaaAbstractAuditRecordAllOf) GetMoDisplayNames() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.MoDisplayNames
}

// GetMoDisplayNamesOk returns a tuple with the MoDisplayNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AaaAbstractAuditRecordAllOf) GetMoDisplayNamesOk() (*interface{}, bool) {
	if o == nil || o.MoDisplayNames == nil {
		return nil, false
	}
	return &o.MoDisplayNames, true
}

// HasMoDisplayNames returns a boolean if a field has been set.
func (o *AaaAbstractAuditRecordAllOf) HasMoDisplayNames() bool {
	if o != nil && o.MoDisplayNames != nil {
		return true
	}

	return false
}

// SetMoDisplayNames gets a reference to the given interface{} and assigns it to the MoDisplayNames field.
func (o *AaaAbstractAuditRecordAllOf) SetMoDisplayNames(v interface{}) {
	o.MoDisplayNames = v
}

// GetMoType returns the MoType field value if set, zero value otherwise.
func (o *AaaAbstractAuditRecordAllOf) GetMoType() string {
	if o == nil || o.MoType == nil {
		var ret string
		return ret
	}
	return *o.MoType
}

// GetMoTypeOk returns a tuple with the MoType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AaaAbstractAuditRecordAllOf) GetMoTypeOk() (*string, bool) {
	if o == nil || o.MoType == nil {
		return nil, false
	}
	return o.MoType, true
}

// HasMoType returns a boolean if a field has been set.
func (o *AaaAbstractAuditRecordAllOf) HasMoType() bool {
	if o != nil && o.MoType != nil {
		return true
	}

	return false
}

// SetMoType gets a reference to the given string and assigns it to the MoType field.
func (o *AaaAbstractAuditRecordAllOf) SetMoType(v string) {
	o.MoType = &v
}

// GetObjectMoid returns the ObjectMoid field value if set, zero value otherwise.
func (o *AaaAbstractAuditRecordAllOf) GetObjectMoid() string {
	if o == nil || o.ObjectMoid == nil {
		var ret string
		return ret
	}
	return *o.ObjectMoid
}

// GetObjectMoidOk returns a tuple with the ObjectMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AaaAbstractAuditRecordAllOf) GetObjectMoidOk() (*string, bool) {
	if o == nil || o.ObjectMoid == nil {
		return nil, false
	}
	return o.ObjectMoid, true
}

// HasObjectMoid returns a boolean if a field has been set.
func (o *AaaAbstractAuditRecordAllOf) HasObjectMoid() bool {
	if o != nil && o.ObjectMoid != nil {
		return true
	}

	return false
}

// SetObjectMoid gets a reference to the given string and assigns it to the ObjectMoid field.
func (o *AaaAbstractAuditRecordAllOf) SetObjectMoid(v string) {
	o.ObjectMoid = &v
}

// GetRequest returns the Request field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AaaAbstractAuditRecordAllOf) GetRequest() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AaaAbstractAuditRecordAllOf) GetRequestOk() (*interface{}, bool) {
	if o == nil || o.Request == nil {
		return nil, false
	}
	return &o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *AaaAbstractAuditRecordAllOf) HasRequest() bool {
	if o != nil && o.Request != nil {
		return true
	}

	return false
}

// SetRequest gets a reference to the given interface{} and assigns it to the Request field.
func (o *AaaAbstractAuditRecordAllOf) SetRequest(v interface{}) {
	o.Request = v
}

// GetTraceId returns the TraceId field value if set, zero value otherwise.
func (o *AaaAbstractAuditRecordAllOf) GetTraceId() string {
	if o == nil || o.TraceId == nil {
		var ret string
		return ret
	}
	return *o.TraceId
}

// GetTraceIdOk returns a tuple with the TraceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AaaAbstractAuditRecordAllOf) GetTraceIdOk() (*string, bool) {
	if o == nil || o.TraceId == nil {
		return nil, false
	}
	return o.TraceId, true
}

// HasTraceId returns a boolean if a field has been set.
func (o *AaaAbstractAuditRecordAllOf) HasTraceId() bool {
	if o != nil && o.TraceId != nil {
		return true
	}

	return false
}

// SetTraceId gets a reference to the given string and assigns it to the TraceId field.
func (o *AaaAbstractAuditRecordAllOf) SetTraceId(v string) {
	o.TraceId = &v
}

// GetUserAgent returns the UserAgent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AaaAbstractAuditRecordAllOf) GetUserAgent() AaaUserAgent {
	if o == nil || o.UserAgent.Get() == nil {
		var ret AaaUserAgent
		return ret
	}
	return *o.UserAgent.Get()
}

// GetUserAgentOk returns a tuple with the UserAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AaaAbstractAuditRecordAllOf) GetUserAgentOk() (*AaaUserAgent, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserAgent.Get(), o.UserAgent.IsSet()
}

// HasUserAgent returns a boolean if a field has been set.
func (o *AaaAbstractAuditRecordAllOf) HasUserAgent() bool {
	if o != nil && o.UserAgent.IsSet() {
		return true
	}

	return false
}

// SetUserAgent gets a reference to the given NullableAaaUserAgent and assigns it to the UserAgent field.
func (o *AaaAbstractAuditRecordAllOf) SetUserAgent(v AaaUserAgent) {
	o.UserAgent.Set(&v)
}

// SetUserAgentNil sets the value for UserAgent to be an explicit nil
func (o *AaaAbstractAuditRecordAllOf) SetUserAgentNil() {
	o.UserAgent.Set(nil)
}

// UnsetUserAgent ensures that no value is present for UserAgent, not even an explicit nil
func (o *AaaAbstractAuditRecordAllOf) UnsetUserAgent() {
	o.UserAgent.Unset()
}

// GetUserAgentString returns the UserAgentString field value if set, zero value otherwise.
func (o *AaaAbstractAuditRecordAllOf) GetUserAgentString() string {
	if o == nil || o.UserAgentString == nil {
		var ret string
		return ret
	}
	return *o.UserAgentString
}

// GetUserAgentStringOk returns a tuple with the UserAgentString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AaaAbstractAuditRecordAllOf) GetUserAgentStringOk() (*string, bool) {
	if o == nil || o.UserAgentString == nil {
		return nil, false
	}
	return o.UserAgentString, true
}

// HasUserAgentString returns a boolean if a field has been set.
func (o *AaaAbstractAuditRecordAllOf) HasUserAgentString() bool {
	if o != nil && o.UserAgentString != nil {
		return true
	}

	return false
}

// SetUserAgentString gets a reference to the given string and assigns it to the UserAgentString field.
func (o *AaaAbstractAuditRecordAllOf) SetUserAgentString(v string) {
	o.UserAgentString = &v
}

func (o AaaAbstractAuditRecordAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Event != nil {
		toSerialize["Event"] = o.Event
	}
	if o.MoDisplayNames != nil {
		toSerialize["MoDisplayNames"] = o.MoDisplayNames
	}
	if o.MoType != nil {
		toSerialize["MoType"] = o.MoType
	}
	if o.ObjectMoid != nil {
		toSerialize["ObjectMoid"] = o.ObjectMoid
	}
	if o.Request != nil {
		toSerialize["Request"] = o.Request
	}
	if o.TraceId != nil {
		toSerialize["TraceId"] = o.TraceId
	}
	if o.UserAgent.IsSet() {
		toSerialize["UserAgent"] = o.UserAgent.Get()
	}
	if o.UserAgentString != nil {
		toSerialize["UserAgentString"] = o.UserAgentString
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AaaAbstractAuditRecordAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varAaaAbstractAuditRecordAllOf := _AaaAbstractAuditRecordAllOf{}

	if err = json.Unmarshal(bytes, &varAaaAbstractAuditRecordAllOf); err == nil {
		*o = AaaAbstractAuditRecordAllOf(varAaaAbstractAuditRecordAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Event")
		delete(additionalProperties, "MoDisplayNames")
		delete(additionalProperties, "MoType")
		delete(additionalProperties, "ObjectMoid")
		delete(additionalProperties, "Request")
		delete(additionalProperties, "TraceId")
		delete(additionalProperties, "UserAgent")
		delete(additionalProperties, "UserAgentString")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAaaAbstractAuditRecordAllOf struct {
	value *AaaAbstractAuditRecordAllOf
	isSet bool
}

func (v NullableAaaAbstractAuditRecordAllOf) Get() *AaaAbstractAuditRecordAllOf {
	return v.value
}

func (v *NullableAaaAbstractAuditRecordAllOf) Set(val *AaaAbstractAuditRecordAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAaaAbstractAuditRecordAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAaaAbstractAuditRecordAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAaaAbstractAuditRecordAllOf(val *AaaAbstractAuditRecordAllOf) *NullableAaaAbstractAuditRecordAllOf {
	return &NullableAaaAbstractAuditRecordAllOf{value: val, isSet: true}
}

func (v NullableAaaAbstractAuditRecordAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAaaAbstractAuditRecordAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
