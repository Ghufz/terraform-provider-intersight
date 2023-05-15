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

// SyslogLocalClientBase Lists properties that are common to all local logging clients. It serves as base class for all local logging clients.
type SyslogLocalClientBase struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Lowest level of messages to be included in the local log. * `warning` - Use logging level warning for logs classified as warning. * `emergency` - Use logging level emergency for logs classified as emergency. * `alert` - Use logging level alert for logs classified as alert. * `critical` - Use logging level critical for logs classified as critical. * `error` - Use logging level error for logs classified as error. * `notice` - Use logging level notice for logs classified as notice. * `informational` - Use logging level informational for logs classified as informational. * `debug` - Use logging level debug for logs classified as debug.
	MinSeverity          *string `json:"MinSeverity,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SyslogLocalClientBase SyslogLocalClientBase

// NewSyslogLocalClientBase instantiates a new SyslogLocalClientBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyslogLocalClientBase(classId string, objectType string) *SyslogLocalClientBase {
	this := SyslogLocalClientBase{}
	this.ClassId = classId
	this.ObjectType = objectType
	var minSeverity string = "warning"
	this.MinSeverity = &minSeverity
	return &this
}

// NewSyslogLocalClientBaseWithDefaults instantiates a new SyslogLocalClientBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyslogLocalClientBaseWithDefaults() *SyslogLocalClientBase {
	this := SyslogLocalClientBase{}
	var classId string = "syslog.LocalFileLoggingClient"
	this.ClassId = classId
	var objectType string = "syslog.LocalFileLoggingClient"
	this.ObjectType = objectType
	var minSeverity string = "warning"
	this.MinSeverity = &minSeverity
	return &this
}

// GetClassId returns the ClassId field value
func (o *SyslogLocalClientBase) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SyslogLocalClientBase) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SyslogLocalClientBase) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *SyslogLocalClientBase) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SyslogLocalClientBase) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SyslogLocalClientBase) SetObjectType(v string) {
	o.ObjectType = v
}

// GetMinSeverity returns the MinSeverity field value if set, zero value otherwise.
func (o *SyslogLocalClientBase) GetMinSeverity() string {
	if o == nil || o.MinSeverity == nil {
		var ret string
		return ret
	}
	return *o.MinSeverity
}

// GetMinSeverityOk returns a tuple with the MinSeverity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyslogLocalClientBase) GetMinSeverityOk() (*string, bool) {
	if o == nil || o.MinSeverity == nil {
		return nil, false
	}
	return o.MinSeverity, true
}

// HasMinSeverity returns a boolean if a field has been set.
func (o *SyslogLocalClientBase) HasMinSeverity() bool {
	if o != nil && o.MinSeverity != nil {
		return true
	}

	return false
}

// SetMinSeverity gets a reference to the given string and assigns it to the MinSeverity field.
func (o *SyslogLocalClientBase) SetMinSeverity(v string) {
	o.MinSeverity = &v
}

func (o SyslogLocalClientBase) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.MinSeverity != nil {
		toSerialize["MinSeverity"] = o.MinSeverity
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SyslogLocalClientBase) UnmarshalJSON(bytes []byte) (err error) {
	type SyslogLocalClientBaseWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// Lowest level of messages to be included in the local log. * `warning` - Use logging level warning for logs classified as warning. * `emergency` - Use logging level emergency for logs classified as emergency. * `alert` - Use logging level alert for logs classified as alert. * `critical` - Use logging level critical for logs classified as critical. * `error` - Use logging level error for logs classified as error. * `notice` - Use logging level notice for logs classified as notice. * `informational` - Use logging level informational for logs classified as informational. * `debug` - Use logging level debug for logs classified as debug.
		MinSeverity *string `json:"MinSeverity,omitempty"`
	}

	varSyslogLocalClientBaseWithoutEmbeddedStruct := SyslogLocalClientBaseWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varSyslogLocalClientBaseWithoutEmbeddedStruct)
	if err == nil {
		varSyslogLocalClientBase := _SyslogLocalClientBase{}
		varSyslogLocalClientBase.ClassId = varSyslogLocalClientBaseWithoutEmbeddedStruct.ClassId
		varSyslogLocalClientBase.ObjectType = varSyslogLocalClientBaseWithoutEmbeddedStruct.ObjectType
		varSyslogLocalClientBase.MinSeverity = varSyslogLocalClientBaseWithoutEmbeddedStruct.MinSeverity
		*o = SyslogLocalClientBase(varSyslogLocalClientBase)
	} else {
		return err
	}

	varSyslogLocalClientBase := _SyslogLocalClientBase{}

	err = json.Unmarshal(bytes, &varSyslogLocalClientBase)
	if err == nil {
		o.MoBaseComplexType = varSyslogLocalClientBase.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "MinSeverity")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

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

type NullableSyslogLocalClientBase struct {
	value *SyslogLocalClientBase
	isSet bool
}

func (v NullableSyslogLocalClientBase) Get() *SyslogLocalClientBase {
	return v.value
}

func (v *NullableSyslogLocalClientBase) Set(val *SyslogLocalClientBase) {
	v.value = val
	v.isSet = true
}

func (v NullableSyslogLocalClientBase) IsSet() bool {
	return v.isSet
}

func (v *NullableSyslogLocalClientBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyslogLocalClientBase(val *SyslogLocalClientBase) *NullableSyslogLocalClientBase {
	return &NullableSyslogLocalClientBase{value: val, isSet: true}
}

func (v NullableSyslogLocalClientBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyslogLocalClientBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
