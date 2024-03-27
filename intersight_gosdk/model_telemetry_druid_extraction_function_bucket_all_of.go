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

// TelemetryDruidExtractionFunctionBucketAllOf struct for TelemetryDruidExtractionFunctionBucketAllOf
type TelemetryDruidExtractionFunctionBucketAllOf struct {
	Type string `json:"type"`
	// The size of the buckets (optional, default 1).
	Size *int32 `json:"size,omitempty"`
	// The offset for the buckets (optional, default 0).
	Offset               *int32 `json:"offset,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidExtractionFunctionBucketAllOf TelemetryDruidExtractionFunctionBucketAllOf

// NewTelemetryDruidExtractionFunctionBucketAllOf instantiates a new TelemetryDruidExtractionFunctionBucketAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidExtractionFunctionBucketAllOf(type_ string) *TelemetryDruidExtractionFunctionBucketAllOf {
	this := TelemetryDruidExtractionFunctionBucketAllOf{}
	this.Type = type_
	return &this
}

// NewTelemetryDruidExtractionFunctionBucketAllOfWithDefaults instantiates a new TelemetryDruidExtractionFunctionBucketAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidExtractionFunctionBucketAllOfWithDefaults() *TelemetryDruidExtractionFunctionBucketAllOf {
	this := TelemetryDruidExtractionFunctionBucketAllOf{}
	return &this
}

// GetType returns the Type field value
func (o *TelemetryDruidExtractionFunctionBucketAllOf) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidExtractionFunctionBucketAllOf) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TelemetryDruidExtractionFunctionBucketAllOf) SetType(v string) {
	o.Type = v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *TelemetryDruidExtractionFunctionBucketAllOf) GetSize() int32 {
	if o == nil || o.Size == nil {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidExtractionFunctionBucketAllOf) GetSizeOk() (*int32, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *TelemetryDruidExtractionFunctionBucketAllOf) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *TelemetryDruidExtractionFunctionBucketAllOf) SetSize(v int32) {
	o.Size = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *TelemetryDruidExtractionFunctionBucketAllOf) GetOffset() int32 {
	if o == nil || o.Offset == nil {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidExtractionFunctionBucketAllOf) GetOffsetOk() (*int32, bool) {
	if o == nil || o.Offset == nil {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *TelemetryDruidExtractionFunctionBucketAllOf) HasOffset() bool {
	if o != nil && o.Offset != nil {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *TelemetryDruidExtractionFunctionBucketAllOf) SetOffset(v int32) {
	o.Offset = &v
}

func (o TelemetryDruidExtractionFunctionBucketAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	if o.Offset != nil {
		toSerialize["offset"] = o.Offset
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TelemetryDruidExtractionFunctionBucketAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varTelemetryDruidExtractionFunctionBucketAllOf := _TelemetryDruidExtractionFunctionBucketAllOf{}

	if err = json.Unmarshal(bytes, &varTelemetryDruidExtractionFunctionBucketAllOf); err == nil {
		*o = TelemetryDruidExtractionFunctionBucketAllOf(varTelemetryDruidExtractionFunctionBucketAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "size")
		delete(additionalProperties, "offset")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidExtractionFunctionBucketAllOf struct {
	value *TelemetryDruidExtractionFunctionBucketAllOf
	isSet bool
}

func (v NullableTelemetryDruidExtractionFunctionBucketAllOf) Get() *TelemetryDruidExtractionFunctionBucketAllOf {
	return v.value
}

func (v *NullableTelemetryDruidExtractionFunctionBucketAllOf) Set(val *TelemetryDruidExtractionFunctionBucketAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidExtractionFunctionBucketAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidExtractionFunctionBucketAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidExtractionFunctionBucketAllOf(val *TelemetryDruidExtractionFunctionBucketAllOf) *NullableTelemetryDruidExtractionFunctionBucketAllOf {
	return &NullableTelemetryDruidExtractionFunctionBucketAllOf{value: val, isSet: true}
}

func (v NullableTelemetryDruidExtractionFunctionBucketAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidExtractionFunctionBucketAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
