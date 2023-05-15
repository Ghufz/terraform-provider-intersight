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

// TelemetryDruidNotFilterAllOf struct for TelemetryDruidNotFilterAllOf
type TelemetryDruidNotFilterAllOf struct {
	Field                TelemetryDruidFilter `json:"field"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidNotFilterAllOf TelemetryDruidNotFilterAllOf

// NewTelemetryDruidNotFilterAllOf instantiates a new TelemetryDruidNotFilterAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidNotFilterAllOf(field TelemetryDruidFilter) *TelemetryDruidNotFilterAllOf {
	this := TelemetryDruidNotFilterAllOf{}
	this.Field = field
	return &this
}

// NewTelemetryDruidNotFilterAllOfWithDefaults instantiates a new TelemetryDruidNotFilterAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidNotFilterAllOfWithDefaults() *TelemetryDruidNotFilterAllOf {
	this := TelemetryDruidNotFilterAllOf{}
	return &this
}

// GetField returns the Field field value
func (o *TelemetryDruidNotFilterAllOf) GetField() TelemetryDruidFilter {
	if o == nil {
		var ret TelemetryDruidFilter
		return ret
	}

	return o.Field
}

// GetFieldOk returns a tuple with the Field field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidNotFilterAllOf) GetFieldOk() (*TelemetryDruidFilter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Field, true
}

// SetField sets field value
func (o *TelemetryDruidNotFilterAllOf) SetField(v TelemetryDruidFilter) {
	o.Field = v
}

func (o TelemetryDruidNotFilterAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["field"] = o.Field
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TelemetryDruidNotFilterAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varTelemetryDruidNotFilterAllOf := _TelemetryDruidNotFilterAllOf{}

	if err = json.Unmarshal(bytes, &varTelemetryDruidNotFilterAllOf); err == nil {
		*o = TelemetryDruidNotFilterAllOf(varTelemetryDruidNotFilterAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "field")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidNotFilterAllOf struct {
	value *TelemetryDruidNotFilterAllOf
	isSet bool
}

func (v NullableTelemetryDruidNotFilterAllOf) Get() *TelemetryDruidNotFilterAllOf {
	return v.value
}

func (v *NullableTelemetryDruidNotFilterAllOf) Set(val *TelemetryDruidNotFilterAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidNotFilterAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidNotFilterAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidNotFilterAllOf(val *TelemetryDruidNotFilterAllOf) *NullableTelemetryDruidNotFilterAllOf {
	return &NullableTelemetryDruidNotFilterAllOf{value: val, isSet: true}
}

func (v NullableTelemetryDruidNotFilterAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidNotFilterAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
