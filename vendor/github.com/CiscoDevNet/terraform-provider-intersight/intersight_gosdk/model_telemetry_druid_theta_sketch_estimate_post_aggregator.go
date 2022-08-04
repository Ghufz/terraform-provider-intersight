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

// TelemetryDruidThetaSketchEstimatePostAggregator Post aggregator of type fieldAccess that refers to a thetaSketch aggregator or that of type thetaSketchSetOp.
type TelemetryDruidThetaSketchEstimatePostAggregator struct {
	// The post-aggregator type.
	Type string `json:"type"`
	// Fields processed by post aggregator
	Fields []TelemetryDruidPostAggregator `json:"fields,omitempty"`
	// Output name for the post-aggregator.
	Name *string `json:"name,omitempty"`
	// Post aggregator of type fieldAccess that refers to a thetaSketch aggregator or that of type thetaSketchSetOp.
	Field                *string `json:"field,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidThetaSketchEstimatePostAggregator TelemetryDruidThetaSketchEstimatePostAggregator

// NewTelemetryDruidThetaSketchEstimatePostAggregator instantiates a new TelemetryDruidThetaSketchEstimatePostAggregator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidThetaSketchEstimatePostAggregator(type_ string) *TelemetryDruidThetaSketchEstimatePostAggregator {
	this := TelemetryDruidThetaSketchEstimatePostAggregator{}
	this.Type = type_
	return &this
}

// NewTelemetryDruidThetaSketchEstimatePostAggregatorWithDefaults instantiates a new TelemetryDruidThetaSketchEstimatePostAggregator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidThetaSketchEstimatePostAggregatorWithDefaults() *TelemetryDruidThetaSketchEstimatePostAggregator {
	this := TelemetryDruidThetaSketchEstimatePostAggregator{}
	return &this
}

// GetType returns the Type field value
func (o *TelemetryDruidThetaSketchEstimatePostAggregator) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidThetaSketchEstimatePostAggregator) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TelemetryDruidThetaSketchEstimatePostAggregator) SetType(v string) {
	o.Type = v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *TelemetryDruidThetaSketchEstimatePostAggregator) GetFields() []TelemetryDruidPostAggregator {
	if o == nil || o.Fields == nil {
		var ret []TelemetryDruidPostAggregator
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidThetaSketchEstimatePostAggregator) GetFieldsOk() ([]TelemetryDruidPostAggregator, bool) {
	if o == nil || o.Fields == nil {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *TelemetryDruidThetaSketchEstimatePostAggregator) HasFields() bool {
	if o != nil && o.Fields != nil {
		return true
	}

	return false
}

// SetFields gets a reference to the given []TelemetryDruidPostAggregator and assigns it to the Fields field.
func (o *TelemetryDruidThetaSketchEstimatePostAggregator) SetFields(v []TelemetryDruidPostAggregator) {
	o.Fields = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TelemetryDruidThetaSketchEstimatePostAggregator) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidThetaSketchEstimatePostAggregator) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TelemetryDruidThetaSketchEstimatePostAggregator) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TelemetryDruidThetaSketchEstimatePostAggregator) SetName(v string) {
	o.Name = &v
}

// GetField returns the Field field value if set, zero value otherwise.
func (o *TelemetryDruidThetaSketchEstimatePostAggregator) GetField() string {
	if o == nil || o.Field == nil {
		var ret string
		return ret
	}
	return *o.Field
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidThetaSketchEstimatePostAggregator) GetFieldOk() (*string, bool) {
	if o == nil || o.Field == nil {
		return nil, false
	}
	return o.Field, true
}

// HasField returns a boolean if a field has been set.
func (o *TelemetryDruidThetaSketchEstimatePostAggregator) HasField() bool {
	if o != nil && o.Field != nil {
		return true
	}

	return false
}

// SetField gets a reference to the given string and assigns it to the Field field.
func (o *TelemetryDruidThetaSketchEstimatePostAggregator) SetField(v string) {
	o.Field = &v
}

func (o TelemetryDruidThetaSketchEstimatePostAggregator) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.Fields != nil {
		toSerialize["fields"] = o.Fields
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Field != nil {
		toSerialize["field"] = o.Field
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TelemetryDruidThetaSketchEstimatePostAggregator) UnmarshalJSON(bytes []byte) (err error) {
	varTelemetryDruidThetaSketchEstimatePostAggregator := _TelemetryDruidThetaSketchEstimatePostAggregator{}

	if err = json.Unmarshal(bytes, &varTelemetryDruidThetaSketchEstimatePostAggregator); err == nil {
		*o = TelemetryDruidThetaSketchEstimatePostAggregator(varTelemetryDruidThetaSketchEstimatePostAggregator)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "fields")
		delete(additionalProperties, "name")
		delete(additionalProperties, "field")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidThetaSketchEstimatePostAggregator struct {
	value *TelemetryDruidThetaSketchEstimatePostAggregator
	isSet bool
}

func (v NullableTelemetryDruidThetaSketchEstimatePostAggregator) Get() *TelemetryDruidThetaSketchEstimatePostAggregator {
	return v.value
}

func (v *NullableTelemetryDruidThetaSketchEstimatePostAggregator) Set(val *TelemetryDruidThetaSketchEstimatePostAggregator) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidThetaSketchEstimatePostAggregator) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidThetaSketchEstimatePostAggregator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidThetaSketchEstimatePostAggregator(val *TelemetryDruidThetaSketchEstimatePostAggregator) *NullableTelemetryDruidThetaSketchEstimatePostAggregator {
	return &NullableTelemetryDruidThetaSketchEstimatePostAggregator{value: val, isSet: true}
}

func (v NullableTelemetryDruidThetaSketchEstimatePostAggregator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidThetaSketchEstimatePostAggregator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
