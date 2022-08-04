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

// VirtualizationCpuInfoAllOf Definition of the list of properties defined in 'virtualization.CpuInfo', excluding properties defined in parent classes.
type VirtualizationCpuInfoAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Number of cores per CPU, as reported by the manufacturer.
	Cores *int64 `json:"Cores,omitempty"`
	// The vendor provided description of the CPU. For example, Intel Xeon E5-2640 v3 @ 2.60GHz.
	Description *string `json:"Description,omitempty"`
	// Number of CPU sockets available.
	Sockets *int64 `json:"Sockets,omitempty"`
	// Speed of the CPUs in Hertz. For example, 2593749663.
	Speed *int64 `json:"Speed,omitempty"`
	// Manufacturer of the CPU . For example, Intel.
	Vendor               *string `json:"Vendor,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationCpuInfoAllOf VirtualizationCpuInfoAllOf

// NewVirtualizationCpuInfoAllOf instantiates a new VirtualizationCpuInfoAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationCpuInfoAllOf(classId string, objectType string) *VirtualizationCpuInfoAllOf {
	this := VirtualizationCpuInfoAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVirtualizationCpuInfoAllOfWithDefaults instantiates a new VirtualizationCpuInfoAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationCpuInfoAllOfWithDefaults() *VirtualizationCpuInfoAllOf {
	this := VirtualizationCpuInfoAllOf{}
	var classId string = "virtualization.CpuInfo"
	this.ClassId = classId
	var objectType string = "virtualization.CpuInfo"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationCpuInfoAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationCpuInfoAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationCpuInfoAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationCpuInfoAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationCpuInfoAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationCpuInfoAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCores returns the Cores field value if set, zero value otherwise.
func (o *VirtualizationCpuInfoAllOf) GetCores() int64 {
	if o == nil || o.Cores == nil {
		var ret int64
		return ret
	}
	return *o.Cores
}

// GetCoresOk returns a tuple with the Cores field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationCpuInfoAllOf) GetCoresOk() (*int64, bool) {
	if o == nil || o.Cores == nil {
		return nil, false
	}
	return o.Cores, true
}

// HasCores returns a boolean if a field has been set.
func (o *VirtualizationCpuInfoAllOf) HasCores() bool {
	if o != nil && o.Cores != nil {
		return true
	}

	return false
}

// SetCores gets a reference to the given int64 and assigns it to the Cores field.
func (o *VirtualizationCpuInfoAllOf) SetCores(v int64) {
	o.Cores = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *VirtualizationCpuInfoAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationCpuInfoAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *VirtualizationCpuInfoAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *VirtualizationCpuInfoAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetSockets returns the Sockets field value if set, zero value otherwise.
func (o *VirtualizationCpuInfoAllOf) GetSockets() int64 {
	if o == nil || o.Sockets == nil {
		var ret int64
		return ret
	}
	return *o.Sockets
}

// GetSocketsOk returns a tuple with the Sockets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationCpuInfoAllOf) GetSocketsOk() (*int64, bool) {
	if o == nil || o.Sockets == nil {
		return nil, false
	}
	return o.Sockets, true
}

// HasSockets returns a boolean if a field has been set.
func (o *VirtualizationCpuInfoAllOf) HasSockets() bool {
	if o != nil && o.Sockets != nil {
		return true
	}

	return false
}

// SetSockets gets a reference to the given int64 and assigns it to the Sockets field.
func (o *VirtualizationCpuInfoAllOf) SetSockets(v int64) {
	o.Sockets = &v
}

// GetSpeed returns the Speed field value if set, zero value otherwise.
func (o *VirtualizationCpuInfoAllOf) GetSpeed() int64 {
	if o == nil || o.Speed == nil {
		var ret int64
		return ret
	}
	return *o.Speed
}

// GetSpeedOk returns a tuple with the Speed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationCpuInfoAllOf) GetSpeedOk() (*int64, bool) {
	if o == nil || o.Speed == nil {
		return nil, false
	}
	return o.Speed, true
}

// HasSpeed returns a boolean if a field has been set.
func (o *VirtualizationCpuInfoAllOf) HasSpeed() bool {
	if o != nil && o.Speed != nil {
		return true
	}

	return false
}

// SetSpeed gets a reference to the given int64 and assigns it to the Speed field.
func (o *VirtualizationCpuInfoAllOf) SetSpeed(v int64) {
	o.Speed = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *VirtualizationCpuInfoAllOf) GetVendor() string {
	if o == nil || o.Vendor == nil {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationCpuInfoAllOf) GetVendorOk() (*string, bool) {
	if o == nil || o.Vendor == nil {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *VirtualizationCpuInfoAllOf) HasVendor() bool {
	if o != nil && o.Vendor != nil {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *VirtualizationCpuInfoAllOf) SetVendor(v string) {
	o.Vendor = &v
}

func (o VirtualizationCpuInfoAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Cores != nil {
		toSerialize["Cores"] = o.Cores
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.Sockets != nil {
		toSerialize["Sockets"] = o.Sockets
	}
	if o.Speed != nil {
		toSerialize["Speed"] = o.Speed
	}
	if o.Vendor != nil {
		toSerialize["Vendor"] = o.Vendor
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationCpuInfoAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varVirtualizationCpuInfoAllOf := _VirtualizationCpuInfoAllOf{}

	if err = json.Unmarshal(bytes, &varVirtualizationCpuInfoAllOf); err == nil {
		*o = VirtualizationCpuInfoAllOf(varVirtualizationCpuInfoAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Cores")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "Sockets")
		delete(additionalProperties, "Speed")
		delete(additionalProperties, "Vendor")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVirtualizationCpuInfoAllOf struct {
	value *VirtualizationCpuInfoAllOf
	isSet bool
}

func (v NullableVirtualizationCpuInfoAllOf) Get() *VirtualizationCpuInfoAllOf {
	return v.value
}

func (v *NullableVirtualizationCpuInfoAllOf) Set(val *VirtualizationCpuInfoAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationCpuInfoAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationCpuInfoAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationCpuInfoAllOf(val *VirtualizationCpuInfoAllOf) *NullableVirtualizationCpuInfoAllOf {
	return &NullableVirtualizationCpuInfoAllOf{value: val, isSet: true}
}

func (v NullableVirtualizationCpuInfoAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationCpuInfoAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
