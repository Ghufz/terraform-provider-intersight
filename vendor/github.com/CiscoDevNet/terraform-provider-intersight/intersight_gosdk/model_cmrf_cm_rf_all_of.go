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

// CmrfCmRfAllOf Definition of the list of properties defined in 'cmrf.CmRf', excluding properties defined in parent classes.
type CmrfCmRfAllOf struct {
	// The Moid of the referenced REST resource.
	Moid                 *string `json:"Moid,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CmrfCmRfAllOf CmrfCmRfAllOf

// NewCmrfCmRfAllOf instantiates a new CmrfCmRfAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCmrfCmRfAllOf() *CmrfCmRfAllOf {
	this := CmrfCmRfAllOf{}
	return &this
}

// NewCmrfCmRfAllOfWithDefaults instantiates a new CmrfCmRfAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCmrfCmRfAllOfWithDefaults() *CmrfCmRfAllOf {
	this := CmrfCmRfAllOf{}
	return &this
}

// GetMoid returns the Moid field value if set, zero value otherwise.
func (o *CmrfCmRfAllOf) GetMoid() string {
	if o == nil || o.Moid == nil {
		var ret string
		return ret
	}
	return *o.Moid
}

// GetMoidOk returns a tuple with the Moid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmrfCmRfAllOf) GetMoidOk() (*string, bool) {
	if o == nil || o.Moid == nil {
		return nil, false
	}
	return o.Moid, true
}

// HasMoid returns a boolean if a field has been set.
func (o *CmrfCmRfAllOf) HasMoid() bool {
	if o != nil && o.Moid != nil {
		return true
	}

	return false
}

// SetMoid gets a reference to the given string and assigns it to the Moid field.
func (o *CmrfCmRfAllOf) SetMoid(v string) {
	o.Moid = &v
}

func (o CmrfCmRfAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Moid != nil {
		toSerialize["Moid"] = o.Moid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CmrfCmRfAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varCmrfCmRfAllOf := _CmrfCmRfAllOf{}

	if err = json.Unmarshal(bytes, &varCmrfCmRfAllOf); err == nil {
		*o = CmrfCmRfAllOf(varCmrfCmRfAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Moid")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCmrfCmRfAllOf struct {
	value *CmrfCmRfAllOf
	isSet bool
}

func (v NullableCmrfCmRfAllOf) Get() *CmrfCmRfAllOf {
	return v.value
}

func (v *NullableCmrfCmRfAllOf) Set(val *CmrfCmRfAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCmrfCmRfAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCmrfCmRfAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCmrfCmRfAllOf(val *CmrfCmRfAllOf) *NullableCmrfCmRfAllOf {
	return &NullableCmrfCmRfAllOf{value: val, isSet: true}
}

func (v NullableCmrfCmRfAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCmrfCmRfAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
