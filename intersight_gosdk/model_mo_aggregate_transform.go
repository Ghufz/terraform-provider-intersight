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

// MoAggregateTransform The output of a request that includes the $apply query parameter. The schema of an aggregation query is dynamically determined based on the request query parameters.  See https://intersight.com/apidocs/introduction/query/#apply-query-option for more details.
type MoAggregateTransform struct {
	MoBaseResponse
	// The results of the aggregation query.
	Results              []map[string]interface{} `json:"Results,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MoAggregateTransform MoAggregateTransform

// NewMoAggregateTransform instantiates a new MoAggregateTransform object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMoAggregateTransform(objectType string) *MoAggregateTransform {
	this := MoAggregateTransform{}
	this.ObjectType = objectType
	return &this
}

// NewMoAggregateTransformWithDefaults instantiates a new MoAggregateTransform object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMoAggregateTransformWithDefaults() *MoAggregateTransform {
	this := MoAggregateTransform{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MoAggregateTransform) GetResults() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MoAggregateTransform) GetResultsOk() ([]map[string]interface{}, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *MoAggregateTransform) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []map[string]interface{} and assigns it to the Results field.
func (o *MoAggregateTransform) SetResults(v []map[string]interface{}) {
	o.Results = v
}

func (o MoAggregateTransform) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseResponse, errMoBaseResponse := json.Marshal(o.MoBaseResponse)
	if errMoBaseResponse != nil {
		return []byte{}, errMoBaseResponse
	}
	errMoBaseResponse = json.Unmarshal([]byte(serializedMoBaseResponse), &toSerialize)
	if errMoBaseResponse != nil {
		return []byte{}, errMoBaseResponse
	}
	if o.Results != nil {
		toSerialize["Results"] = o.Results
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MoAggregateTransform) UnmarshalJSON(bytes []byte) (err error) {
	type MoAggregateTransformWithoutEmbeddedStruct struct {
		// The results of the aggregation query.
		Results []map[string]interface{} `json:"Results,omitempty"`
	}

	varMoAggregateTransformWithoutEmbeddedStruct := MoAggregateTransformWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varMoAggregateTransformWithoutEmbeddedStruct)
	if err == nil {
		varMoAggregateTransform := _MoAggregateTransform{}
		varMoAggregateTransform.Results = varMoAggregateTransformWithoutEmbeddedStruct.Results
		*o = MoAggregateTransform(varMoAggregateTransform)
	} else {
		return err
	}

	varMoAggregateTransform := _MoAggregateTransform{}

	err = json.Unmarshal(bytes, &varMoAggregateTransform)
	if err == nil {
		o.MoBaseResponse = varMoAggregateTransform.MoBaseResponse
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Results")

		// remove fields from embedded structs
		reflectMoBaseResponse := reflect.ValueOf(o.MoBaseResponse)
		for i := 0; i < reflectMoBaseResponse.Type().NumField(); i++ {
			t := reflectMoBaseResponse.Type().Field(i)

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

type NullableMoAggregateTransform struct {
	value *MoAggregateTransform
	isSet bool
}

func (v NullableMoAggregateTransform) Get() *MoAggregateTransform {
	return v.value
}

func (v *NullableMoAggregateTransform) Set(val *MoAggregateTransform) {
	v.value = val
	v.isSet = true
}

func (v NullableMoAggregateTransform) IsSet() bool {
	return v.isSet
}

func (v *NullableMoAggregateTransform) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMoAggregateTransform(val *MoAggregateTransform) *NullableMoAggregateTransform {
	return &NullableMoAggregateTransform{value: val, isSet: true}
}

func (v NullableMoAggregateTransform) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMoAggregateTransform) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
