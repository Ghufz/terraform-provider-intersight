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

// CondAlarmSpecAllOf Definition of the list of properties defined in 'cond.AlarmSpec', excluding properties defined in parent classes.
type CondAlarmSpecAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The severity of the alarm. * `None` - The Enum value None represents that there is no severity. * `Info` - The Enum value Info represents the Informational level of severity. * `Critical` - The Enum value Critical represents the Critical level of severity. * `Warning` - The Enum value Warning represents the Warning level of severity. * `Cleared` - The Enum value Cleared represents that the alarm severity has been cleared.
	Severity             *string `json:"Severity,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CondAlarmSpecAllOf CondAlarmSpecAllOf

// NewCondAlarmSpecAllOf instantiates a new CondAlarmSpecAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCondAlarmSpecAllOf(classId string, objectType string) *CondAlarmSpecAllOf {
	this := CondAlarmSpecAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var severity string = "None"
	this.Severity = &severity
	return &this
}

// NewCondAlarmSpecAllOfWithDefaults instantiates a new CondAlarmSpecAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCondAlarmSpecAllOfWithDefaults() *CondAlarmSpecAllOf {
	this := CondAlarmSpecAllOf{}
	var classId string = "cond.AlarmSpec"
	this.ClassId = classId
	var objectType string = "cond.AlarmSpec"
	this.ObjectType = objectType
	var severity string = "None"
	this.Severity = &severity
	return &this
}

// GetClassId returns the ClassId field value
func (o *CondAlarmSpecAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CondAlarmSpecAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CondAlarmSpecAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CondAlarmSpecAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CondAlarmSpecAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CondAlarmSpecAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *CondAlarmSpecAllOf) GetSeverity() string {
	if o == nil || o.Severity == nil {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondAlarmSpecAllOf) GetSeverityOk() (*string, bool) {
	if o == nil || o.Severity == nil {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *CondAlarmSpecAllOf) HasSeverity() bool {
	if o != nil && o.Severity != nil {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *CondAlarmSpecAllOf) SetSeverity(v string) {
	o.Severity = &v
}

func (o CondAlarmSpecAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Severity != nil {
		toSerialize["Severity"] = o.Severity
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CondAlarmSpecAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varCondAlarmSpecAllOf := _CondAlarmSpecAllOf{}

	if err = json.Unmarshal(bytes, &varCondAlarmSpecAllOf); err == nil {
		*o = CondAlarmSpecAllOf(varCondAlarmSpecAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Severity")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCondAlarmSpecAllOf struct {
	value *CondAlarmSpecAllOf
	isSet bool
}

func (v NullableCondAlarmSpecAllOf) Get() *CondAlarmSpecAllOf {
	return v.value
}

func (v *NullableCondAlarmSpecAllOf) Set(val *CondAlarmSpecAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCondAlarmSpecAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCondAlarmSpecAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCondAlarmSpecAllOf(val *CondAlarmSpecAllOf) *NullableCondAlarmSpecAllOf {
	return &NullableCondAlarmSpecAllOf{value: val, isSet: true}
}

func (v NullableCondAlarmSpecAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCondAlarmSpecAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
