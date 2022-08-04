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
)

// WorkflowSelectorPropertyAllOf Definition of the list of properties defined in 'workflow.SelectorProperty', excluding properties defined in parent classes.
type WorkflowSelectorPropertyAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Content of the request body to send for POST request.
	Body interface{} `json:"Body,omitempty"`
	// The HTTP method to be used. * `GET` - The HTTP GET method requests a representation of the specified resource. * `POST` - The HTTP POST method sends data to the server.
	Method               *string `json:"Method,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowSelectorPropertyAllOf WorkflowSelectorPropertyAllOf

// NewWorkflowSelectorPropertyAllOf instantiates a new WorkflowSelectorPropertyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowSelectorPropertyAllOf(classId string, objectType string) *WorkflowSelectorPropertyAllOf {
	this := WorkflowSelectorPropertyAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var method string = "GET"
	this.Method = &method
	return &this
}

// NewWorkflowSelectorPropertyAllOfWithDefaults instantiates a new WorkflowSelectorPropertyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowSelectorPropertyAllOfWithDefaults() *WorkflowSelectorPropertyAllOf {
	this := WorkflowSelectorPropertyAllOf{}
	var classId string = "workflow.SelectorProperty"
	this.ClassId = classId
	var objectType string = "workflow.SelectorProperty"
	this.ObjectType = objectType
	var method string = "GET"
	this.Method = &method
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowSelectorPropertyAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowSelectorPropertyAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowSelectorPropertyAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowSelectorPropertyAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowSelectorPropertyAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowSelectorPropertyAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBody returns the Body field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowSelectorPropertyAllOf) GetBody() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowSelectorPropertyAllOf) GetBodyOk() (*interface{}, bool) {
	if o == nil || o.Body == nil {
		return nil, false
	}
	return &o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *WorkflowSelectorPropertyAllOf) HasBody() bool {
	if o != nil && o.Body != nil {
		return true
	}

	return false
}

// SetBody gets a reference to the given interface{} and assigns it to the Body field.
func (o *WorkflowSelectorPropertyAllOf) SetBody(v interface{}) {
	o.Body = v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *WorkflowSelectorPropertyAllOf) GetMethod() string {
	if o == nil || o.Method == nil {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSelectorPropertyAllOf) GetMethodOk() (*string, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *WorkflowSelectorPropertyAllOf) HasMethod() bool {
	if o != nil && o.Method != nil {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *WorkflowSelectorPropertyAllOf) SetMethod(v string) {
	o.Method = &v
}

func (o WorkflowSelectorPropertyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Body != nil {
		toSerialize["Body"] = o.Body
	}
	if o.Method != nil {
		toSerialize["Method"] = o.Method
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowSelectorPropertyAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowSelectorPropertyAllOf := _WorkflowSelectorPropertyAllOf{}

	if err = json.Unmarshal(bytes, &varWorkflowSelectorPropertyAllOf); err == nil {
		*o = WorkflowSelectorPropertyAllOf(varWorkflowSelectorPropertyAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Body")
		delete(additionalProperties, "Method")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowSelectorPropertyAllOf struct {
	value *WorkflowSelectorPropertyAllOf
	isSet bool
}

func (v NullableWorkflowSelectorPropertyAllOf) Get() *WorkflowSelectorPropertyAllOf {
	return v.value
}

func (v *NullableWorkflowSelectorPropertyAllOf) Set(val *WorkflowSelectorPropertyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowSelectorPropertyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowSelectorPropertyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowSelectorPropertyAllOf(val *WorkflowSelectorPropertyAllOf) *NullableWorkflowSelectorPropertyAllOf {
	return &NullableWorkflowSelectorPropertyAllOf{value: val, isSet: true}
}

func (v NullableWorkflowSelectorPropertyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowSelectorPropertyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
