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

// AdapterPortChannelSettings Port Channel setting for this adapter.
type AdapterPortChannelSettings struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// When Port Channel is enabled, two vNICs and two vHBAs are available for use on the adapter card. When disabled, four vNICs and four vHBAs are available for use on the adapter card. Disabling port channel reboots the server. Port Channel is supported only for Cisco VIC 1455/1457 adapters.
	Enabled              *bool `json:"Enabled,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AdapterPortChannelSettings AdapterPortChannelSettings

// NewAdapterPortChannelSettings instantiates a new AdapterPortChannelSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdapterPortChannelSettings(classId string, objectType string) *AdapterPortChannelSettings {
	this := AdapterPortChannelSettings{}
	this.ClassId = classId
	this.ObjectType = objectType
	var enabled bool = true
	this.Enabled = &enabled
	return &this
}

// NewAdapterPortChannelSettingsWithDefaults instantiates a new AdapterPortChannelSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdapterPortChannelSettingsWithDefaults() *AdapterPortChannelSettings {
	this := AdapterPortChannelSettings{}
	var classId string = "adapter.PortChannelSettings"
	this.ClassId = classId
	var objectType string = "adapter.PortChannelSettings"
	this.ObjectType = objectType
	var enabled bool = true
	this.Enabled = &enabled
	return &this
}

// GetClassId returns the ClassId field value
func (o *AdapterPortChannelSettings) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AdapterPortChannelSettings) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AdapterPortChannelSettings) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AdapterPortChannelSettings) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AdapterPortChannelSettings) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AdapterPortChannelSettings) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *AdapterPortChannelSettings) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterPortChannelSettings) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *AdapterPortChannelSettings) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *AdapterPortChannelSettings) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o AdapterPortChannelSettings) MarshalJSON() ([]byte, error) {
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
	if o.Enabled != nil {
		toSerialize["Enabled"] = o.Enabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AdapterPortChannelSettings) UnmarshalJSON(bytes []byte) (err error) {
	type AdapterPortChannelSettingsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// When Port Channel is enabled, two vNICs and two vHBAs are available for use on the adapter card. When disabled, four vNICs and four vHBAs are available for use on the adapter card. Disabling port channel reboots the server. Port Channel is supported only for Cisco VIC 1455/1457 adapters.
		Enabled *bool `json:"Enabled,omitempty"`
	}

	varAdapterPortChannelSettingsWithoutEmbeddedStruct := AdapterPortChannelSettingsWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varAdapterPortChannelSettingsWithoutEmbeddedStruct)
	if err == nil {
		varAdapterPortChannelSettings := _AdapterPortChannelSettings{}
		varAdapterPortChannelSettings.ClassId = varAdapterPortChannelSettingsWithoutEmbeddedStruct.ClassId
		varAdapterPortChannelSettings.ObjectType = varAdapterPortChannelSettingsWithoutEmbeddedStruct.ObjectType
		varAdapterPortChannelSettings.Enabled = varAdapterPortChannelSettingsWithoutEmbeddedStruct.Enabled
		*o = AdapterPortChannelSettings(varAdapterPortChannelSettings)
	} else {
		return err
	}

	varAdapterPortChannelSettings := _AdapterPortChannelSettings{}

	err = json.Unmarshal(bytes, &varAdapterPortChannelSettings)
	if err == nil {
		o.MoBaseComplexType = varAdapterPortChannelSettings.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Enabled")

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

type NullableAdapterPortChannelSettings struct {
	value *AdapterPortChannelSettings
	isSet bool
}

func (v NullableAdapterPortChannelSettings) Get() *AdapterPortChannelSettings {
	return v.value
}

func (v *NullableAdapterPortChannelSettings) Set(val *AdapterPortChannelSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableAdapterPortChannelSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableAdapterPortChannelSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdapterPortChannelSettings(val *AdapterPortChannelSettings) *NullableAdapterPortChannelSettings {
	return &NullableAdapterPortChannelSettings{value: val, isSet: true}
}

func (v NullableAdapterPortChannelSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdapterPortChannelSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
