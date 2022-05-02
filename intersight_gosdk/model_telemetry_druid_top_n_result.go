/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-6484
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"time"
)

// TelemetryDruidTopNResult Sorted time series results for a particular interval based on some criteria
type TelemetryDruidTopNResult struct {
	// The ISO 8601 timestamp.
	Timestamp *time.Time `json:"timestamp,omitempty"`
	// A sorted list of dimension values
	Result               []map[string]interface{} `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidTopNResult TelemetryDruidTopNResult

// NewTelemetryDruidTopNResult instantiates a new TelemetryDruidTopNResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidTopNResult() *TelemetryDruidTopNResult {
	this := TelemetryDruidTopNResult{}
	return &this
}

// NewTelemetryDruidTopNResultWithDefaults instantiates a new TelemetryDruidTopNResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidTopNResultWithDefaults() *TelemetryDruidTopNResult {
	this := TelemetryDruidTopNResult{}
	return &this
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *TelemetryDruidTopNResult) GetTimestamp() time.Time {
	if o == nil || o.Timestamp == nil {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidTopNResult) GetTimestampOk() (*time.Time, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *TelemetryDruidTopNResult) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *TelemetryDruidTopNResult) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *TelemetryDruidTopNResult) GetResult() []map[string]interface{} {
	if o == nil || o.Result == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidTopNResult) GetResultOk() ([]map[string]interface{}, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *TelemetryDruidTopNResult) HasResult() bool {
	if o != nil && o.Result != nil {
		return true
	}

	return false
}

// SetResult gets a reference to the given []map[string]interface{} and assigns it to the Result field.
func (o *TelemetryDruidTopNResult) SetResult(v []map[string]interface{}) {
	o.Result = v
}

func (o TelemetryDruidTopNResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TelemetryDruidTopNResult) UnmarshalJSON(bytes []byte) (err error) {
	varTelemetryDruidTopNResult := _TelemetryDruidTopNResult{}

	if err = json.Unmarshal(bytes, &varTelemetryDruidTopNResult); err == nil {
		*o = TelemetryDruidTopNResult(varTelemetryDruidTopNResult)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "timestamp")
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidTopNResult struct {
	value *TelemetryDruidTopNResult
	isSet bool
}

func (v NullableTelemetryDruidTopNResult) Get() *TelemetryDruidTopNResult {
	return v.value
}

func (v *NullableTelemetryDruidTopNResult) Set(val *TelemetryDruidTopNResult) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidTopNResult) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidTopNResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidTopNResult(val *TelemetryDruidTopNResult) *NullableTelemetryDruidTopNResult {
	return &NullableTelemetryDruidTopNResult{value: val, isSet: true}
}

func (v NullableTelemetryDruidTopNResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidTopNResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
