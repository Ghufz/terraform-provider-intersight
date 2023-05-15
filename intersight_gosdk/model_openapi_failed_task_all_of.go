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

// OpenapiFailedTaskAllOf Definition of the list of properties defined in 'openapi.FailedTask', excluding properties defined in parent classes.
type OpenapiFailedTaskAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Name of the task.
	Name *string `json:"Name,omitempty"`
	// REST API path of the task.
	Path *string `json:"Path,omitempty"`
	// Indicates the reason for task creation failure.
	Reason               *string `json:"Reason,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OpenapiFailedTaskAllOf OpenapiFailedTaskAllOf

// NewOpenapiFailedTaskAllOf instantiates a new OpenapiFailedTaskAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenapiFailedTaskAllOf(classId string, objectType string) *OpenapiFailedTaskAllOf {
	this := OpenapiFailedTaskAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewOpenapiFailedTaskAllOfWithDefaults instantiates a new OpenapiFailedTaskAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenapiFailedTaskAllOfWithDefaults() *OpenapiFailedTaskAllOf {
	this := OpenapiFailedTaskAllOf{}
	var classId string = "openapi.FailedTask"
	this.ClassId = classId
	var objectType string = "openapi.FailedTask"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *OpenapiFailedTaskAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *OpenapiFailedTaskAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *OpenapiFailedTaskAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *OpenapiFailedTaskAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *OpenapiFailedTaskAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *OpenapiFailedTaskAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OpenapiFailedTaskAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenapiFailedTaskAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OpenapiFailedTaskAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OpenapiFailedTaskAllOf) SetName(v string) {
	o.Name = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *OpenapiFailedTaskAllOf) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenapiFailedTaskAllOf) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *OpenapiFailedTaskAllOf) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *OpenapiFailedTaskAllOf) SetPath(v string) {
	o.Path = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *OpenapiFailedTaskAllOf) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenapiFailedTaskAllOf) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *OpenapiFailedTaskAllOf) HasReason() bool {
	if o != nil && o.Reason != nil {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *OpenapiFailedTaskAllOf) SetReason(v string) {
	o.Reason = &v
}

func (o OpenapiFailedTaskAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Path != nil {
		toSerialize["Path"] = o.Path
	}
	if o.Reason != nil {
		toSerialize["Reason"] = o.Reason
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OpenapiFailedTaskAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varOpenapiFailedTaskAllOf := _OpenapiFailedTaskAllOf{}

	if err = json.Unmarshal(bytes, &varOpenapiFailedTaskAllOf); err == nil {
		*o = OpenapiFailedTaskAllOf(varOpenapiFailedTaskAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Path")
		delete(additionalProperties, "Reason")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOpenapiFailedTaskAllOf struct {
	value *OpenapiFailedTaskAllOf
	isSet bool
}

func (v NullableOpenapiFailedTaskAllOf) Get() *OpenapiFailedTaskAllOf {
	return v.value
}

func (v *NullableOpenapiFailedTaskAllOf) Set(val *OpenapiFailedTaskAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenapiFailedTaskAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenapiFailedTaskAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenapiFailedTaskAllOf(val *OpenapiFailedTaskAllOf) *NullableOpenapiFailedTaskAllOf {
	return &NullableOpenapiFailedTaskAllOf{value: val, isSet: true}
}

func (v NullableOpenapiFailedTaskAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenapiFailedTaskAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
