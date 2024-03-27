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

// TelemetryDruidExtractionFunctionTimeParsing Parses dimension values as timestamps using the given input format, and returns them formatted using the given output format. Note, if you are working with the __time dimension, you should consider using the time extraction function instead instead, which works on time value directly as opposed to string values. If \"joda\" is true, time formats are described in the Joda DateTimeFormat documentation. If \"joda\" is false (or unspecified) then formats are described in the SimpleDateFormat documentation. In general, we recommend setting \"joda\" to true since Joda format strings are more common in Druid APIs and since Joda handles certain edge cases (like weeks and weekyears near the start and end of calendar years) in a more ISO8601 compliant way. If a value cannot be parsed using the provided timeFormat, it will be returned as-is.
type TelemetryDruidExtractionFunctionTimeParsing struct {
	Type                 string  `json:"type"`
	TimeFormat           *string `json:"timeFormat,omitempty"`
	ResultFormat         *string `json:"resultFormat,omitempty"`
	Joda                 *bool   `json:"joda,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidExtractionFunctionTimeParsing TelemetryDruidExtractionFunctionTimeParsing

// NewTelemetryDruidExtractionFunctionTimeParsing instantiates a new TelemetryDruidExtractionFunctionTimeParsing object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidExtractionFunctionTimeParsing(type_ string) *TelemetryDruidExtractionFunctionTimeParsing {
	this := TelemetryDruidExtractionFunctionTimeParsing{}
	this.Type = type_
	return &this
}

// NewTelemetryDruidExtractionFunctionTimeParsingWithDefaults instantiates a new TelemetryDruidExtractionFunctionTimeParsing object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidExtractionFunctionTimeParsingWithDefaults() *TelemetryDruidExtractionFunctionTimeParsing {
	this := TelemetryDruidExtractionFunctionTimeParsing{}
	return &this
}

// GetType returns the Type field value
func (o *TelemetryDruidExtractionFunctionTimeParsing) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidExtractionFunctionTimeParsing) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TelemetryDruidExtractionFunctionTimeParsing) SetType(v string) {
	o.Type = v
}

// GetTimeFormat returns the TimeFormat field value if set, zero value otherwise.
func (o *TelemetryDruidExtractionFunctionTimeParsing) GetTimeFormat() string {
	if o == nil || o.TimeFormat == nil {
		var ret string
		return ret
	}
	return *o.TimeFormat
}

// GetTimeFormatOk returns a tuple with the TimeFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidExtractionFunctionTimeParsing) GetTimeFormatOk() (*string, bool) {
	if o == nil || o.TimeFormat == nil {
		return nil, false
	}
	return o.TimeFormat, true
}

// HasTimeFormat returns a boolean if a field has been set.
func (o *TelemetryDruidExtractionFunctionTimeParsing) HasTimeFormat() bool {
	if o != nil && o.TimeFormat != nil {
		return true
	}

	return false
}

// SetTimeFormat gets a reference to the given string and assigns it to the TimeFormat field.
func (o *TelemetryDruidExtractionFunctionTimeParsing) SetTimeFormat(v string) {
	o.TimeFormat = &v
}

// GetResultFormat returns the ResultFormat field value if set, zero value otherwise.
func (o *TelemetryDruidExtractionFunctionTimeParsing) GetResultFormat() string {
	if o == nil || o.ResultFormat == nil {
		var ret string
		return ret
	}
	return *o.ResultFormat
}

// GetResultFormatOk returns a tuple with the ResultFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidExtractionFunctionTimeParsing) GetResultFormatOk() (*string, bool) {
	if o == nil || o.ResultFormat == nil {
		return nil, false
	}
	return o.ResultFormat, true
}

// HasResultFormat returns a boolean if a field has been set.
func (o *TelemetryDruidExtractionFunctionTimeParsing) HasResultFormat() bool {
	if o != nil && o.ResultFormat != nil {
		return true
	}

	return false
}

// SetResultFormat gets a reference to the given string and assigns it to the ResultFormat field.
func (o *TelemetryDruidExtractionFunctionTimeParsing) SetResultFormat(v string) {
	o.ResultFormat = &v
}

// GetJoda returns the Joda field value if set, zero value otherwise.
func (o *TelemetryDruidExtractionFunctionTimeParsing) GetJoda() bool {
	if o == nil || o.Joda == nil {
		var ret bool
		return ret
	}
	return *o.Joda
}

// GetJodaOk returns a tuple with the Joda field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidExtractionFunctionTimeParsing) GetJodaOk() (*bool, bool) {
	if o == nil || o.Joda == nil {
		return nil, false
	}
	return o.Joda, true
}

// HasJoda returns a boolean if a field has been set.
func (o *TelemetryDruidExtractionFunctionTimeParsing) HasJoda() bool {
	if o != nil && o.Joda != nil {
		return true
	}

	return false
}

// SetJoda gets a reference to the given bool and assigns it to the Joda field.
func (o *TelemetryDruidExtractionFunctionTimeParsing) SetJoda(v bool) {
	o.Joda = &v
}

func (o TelemetryDruidExtractionFunctionTimeParsing) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.TimeFormat != nil {
		toSerialize["timeFormat"] = o.TimeFormat
	}
	if o.ResultFormat != nil {
		toSerialize["resultFormat"] = o.ResultFormat
	}
	if o.Joda != nil {
		toSerialize["joda"] = o.Joda
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TelemetryDruidExtractionFunctionTimeParsing) UnmarshalJSON(bytes []byte) (err error) {
	varTelemetryDruidExtractionFunctionTimeParsing := _TelemetryDruidExtractionFunctionTimeParsing{}

	if err = json.Unmarshal(bytes, &varTelemetryDruidExtractionFunctionTimeParsing); err == nil {
		*o = TelemetryDruidExtractionFunctionTimeParsing(varTelemetryDruidExtractionFunctionTimeParsing)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "timeFormat")
		delete(additionalProperties, "resultFormat")
		delete(additionalProperties, "joda")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidExtractionFunctionTimeParsing struct {
	value *TelemetryDruidExtractionFunctionTimeParsing
	isSet bool
}

func (v NullableTelemetryDruidExtractionFunctionTimeParsing) Get() *TelemetryDruidExtractionFunctionTimeParsing {
	return v.value
}

func (v *NullableTelemetryDruidExtractionFunctionTimeParsing) Set(val *TelemetryDruidExtractionFunctionTimeParsing) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidExtractionFunctionTimeParsing) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidExtractionFunctionTimeParsing) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidExtractionFunctionTimeParsing(val *TelemetryDruidExtractionFunctionTimeParsing) *NullableTelemetryDruidExtractionFunctionTimeParsing {
	return &NullableTelemetryDruidExtractionFunctionTimeParsing{value: val, isSet: true}
}

func (v NullableTelemetryDruidExtractionFunctionTimeParsing) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidExtractionFunctionTimeParsing) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
