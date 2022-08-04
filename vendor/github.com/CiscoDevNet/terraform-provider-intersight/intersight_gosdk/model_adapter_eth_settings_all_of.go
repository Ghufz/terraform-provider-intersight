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

// AdapterEthSettingsAllOf Definition of the list of properties defined in 'adapter.EthSettings', excluding properties defined in parent classes.
type AdapterEthSettingsAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Status of LLDP protocol on the adapter interfaces.
	LldpEnabled          *bool `json:"LldpEnabled,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AdapterEthSettingsAllOf AdapterEthSettingsAllOf

// NewAdapterEthSettingsAllOf instantiates a new AdapterEthSettingsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdapterEthSettingsAllOf(classId string, objectType string) *AdapterEthSettingsAllOf {
	this := AdapterEthSettingsAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var lldpEnabled bool = true
	this.LldpEnabled = &lldpEnabled
	return &this
}

// NewAdapterEthSettingsAllOfWithDefaults instantiates a new AdapterEthSettingsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdapterEthSettingsAllOfWithDefaults() *AdapterEthSettingsAllOf {
	this := AdapterEthSettingsAllOf{}
	var classId string = "adapter.EthSettings"
	this.ClassId = classId
	var objectType string = "adapter.EthSettings"
	this.ObjectType = objectType
	var lldpEnabled bool = true
	this.LldpEnabled = &lldpEnabled
	return &this
}

// GetClassId returns the ClassId field value
func (o *AdapterEthSettingsAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AdapterEthSettingsAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AdapterEthSettingsAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AdapterEthSettingsAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AdapterEthSettingsAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AdapterEthSettingsAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetLldpEnabled returns the LldpEnabled field value if set, zero value otherwise.
func (o *AdapterEthSettingsAllOf) GetLldpEnabled() bool {
	if o == nil || o.LldpEnabled == nil {
		var ret bool
		return ret
	}
	return *o.LldpEnabled
}

// GetLldpEnabledOk returns a tuple with the LldpEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterEthSettingsAllOf) GetLldpEnabledOk() (*bool, bool) {
	if o == nil || o.LldpEnabled == nil {
		return nil, false
	}
	return o.LldpEnabled, true
}

// HasLldpEnabled returns a boolean if a field has been set.
func (o *AdapterEthSettingsAllOf) HasLldpEnabled() bool {
	if o != nil && o.LldpEnabled != nil {
		return true
	}

	return false
}

// SetLldpEnabled gets a reference to the given bool and assigns it to the LldpEnabled field.
func (o *AdapterEthSettingsAllOf) SetLldpEnabled(v bool) {
	o.LldpEnabled = &v
}

func (o AdapterEthSettingsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.LldpEnabled != nil {
		toSerialize["LldpEnabled"] = o.LldpEnabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AdapterEthSettingsAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varAdapterEthSettingsAllOf := _AdapterEthSettingsAllOf{}

	if err = json.Unmarshal(bytes, &varAdapterEthSettingsAllOf); err == nil {
		*o = AdapterEthSettingsAllOf(varAdapterEthSettingsAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "LldpEnabled")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAdapterEthSettingsAllOf struct {
	value *AdapterEthSettingsAllOf
	isSet bool
}

func (v NullableAdapterEthSettingsAllOf) Get() *AdapterEthSettingsAllOf {
	return v.value
}

func (v *NullableAdapterEthSettingsAllOf) Set(val *AdapterEthSettingsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAdapterEthSettingsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAdapterEthSettingsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdapterEthSettingsAllOf(val *AdapterEthSettingsAllOf) *NullableAdapterEthSettingsAllOf {
	return &NullableAdapterEthSettingsAllOf{value: val, isSet: true}
}

func (v NullableAdapterEthSettingsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdapterEthSettingsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
