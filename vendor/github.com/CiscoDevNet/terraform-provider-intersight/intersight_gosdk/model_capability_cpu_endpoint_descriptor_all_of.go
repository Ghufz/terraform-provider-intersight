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

// CapabilityCpuEndpointDescriptorAllOf Definition of the list of properties defined in 'capability.CpuEndpointDescriptor', excluding properties defined in parent classes.
type CapabilityCpuEndpointDescriptorAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// This field is to provide description of the cpu.
	Description *string `json:"Description,omitempty"`
	// This field is to provide model of the cpu.
	Model *string `json:"Model,omitempty"`
	// This field is to provide partNumber of the cpu.
	PartNumber *string `json:"PartNumber,omitempty"`
	// This field is to provide pid of the cpu.
	Pid                    *string  `json:"Pid,omitempty"`
	SupportedPlatformsPids []string `json:"SupportedPlatformsPids,omitempty"`
	// This field is to provide vendor of the cpu.
	Vendor               *string `json:"Vendor,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CapabilityCpuEndpointDescriptorAllOf CapabilityCpuEndpointDescriptorAllOf

// NewCapabilityCpuEndpointDescriptorAllOf instantiates a new CapabilityCpuEndpointDescriptorAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilityCpuEndpointDescriptorAllOf(classId string, objectType string) *CapabilityCpuEndpointDescriptorAllOf {
	this := CapabilityCpuEndpointDescriptorAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCapabilityCpuEndpointDescriptorAllOfWithDefaults instantiates a new CapabilityCpuEndpointDescriptorAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilityCpuEndpointDescriptorAllOfWithDefaults() *CapabilityCpuEndpointDescriptorAllOf {
	this := CapabilityCpuEndpointDescriptorAllOf{}
	var classId string = "capability.CpuEndpointDescriptor"
	this.ClassId = classId
	var objectType string = "capability.CpuEndpointDescriptor"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CapabilityCpuEndpointDescriptorAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CapabilityCpuEndpointDescriptorAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CapabilityCpuEndpointDescriptorAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CapabilityCpuEndpointDescriptorAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CapabilityCpuEndpointDescriptorAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CapabilityCpuEndpointDescriptorAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CapabilityCpuEndpointDescriptorAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityCpuEndpointDescriptorAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CapabilityCpuEndpointDescriptorAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CapabilityCpuEndpointDescriptorAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *CapabilityCpuEndpointDescriptorAllOf) GetModel() string {
	if o == nil || o.Model == nil {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityCpuEndpointDescriptorAllOf) GetModelOk() (*string, bool) {
	if o == nil || o.Model == nil {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *CapabilityCpuEndpointDescriptorAllOf) HasModel() bool {
	if o != nil && o.Model != nil {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *CapabilityCpuEndpointDescriptorAllOf) SetModel(v string) {
	o.Model = &v
}

// GetPartNumber returns the PartNumber field value if set, zero value otherwise.
func (o *CapabilityCpuEndpointDescriptorAllOf) GetPartNumber() string {
	if o == nil || o.PartNumber == nil {
		var ret string
		return ret
	}
	return *o.PartNumber
}

// GetPartNumberOk returns a tuple with the PartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityCpuEndpointDescriptorAllOf) GetPartNumberOk() (*string, bool) {
	if o == nil || o.PartNumber == nil {
		return nil, false
	}
	return o.PartNumber, true
}

// HasPartNumber returns a boolean if a field has been set.
func (o *CapabilityCpuEndpointDescriptorAllOf) HasPartNumber() bool {
	if o != nil && o.PartNumber != nil {
		return true
	}

	return false
}

// SetPartNumber gets a reference to the given string and assigns it to the PartNumber field.
func (o *CapabilityCpuEndpointDescriptorAllOf) SetPartNumber(v string) {
	o.PartNumber = &v
}

// GetPid returns the Pid field value if set, zero value otherwise.
func (o *CapabilityCpuEndpointDescriptorAllOf) GetPid() string {
	if o == nil || o.Pid == nil {
		var ret string
		return ret
	}
	return *o.Pid
}

// GetPidOk returns a tuple with the Pid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityCpuEndpointDescriptorAllOf) GetPidOk() (*string, bool) {
	if o == nil || o.Pid == nil {
		return nil, false
	}
	return o.Pid, true
}

// HasPid returns a boolean if a field has been set.
func (o *CapabilityCpuEndpointDescriptorAllOf) HasPid() bool {
	if o != nil && o.Pid != nil {
		return true
	}

	return false
}

// SetPid gets a reference to the given string and assigns it to the Pid field.
func (o *CapabilityCpuEndpointDescriptorAllOf) SetPid(v string) {
	o.Pid = &v
}

// GetSupportedPlatformsPids returns the SupportedPlatformsPids field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CapabilityCpuEndpointDescriptorAllOf) GetSupportedPlatformsPids() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SupportedPlatformsPids
}

// GetSupportedPlatformsPidsOk returns a tuple with the SupportedPlatformsPids field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CapabilityCpuEndpointDescriptorAllOf) GetSupportedPlatformsPidsOk() ([]string, bool) {
	if o == nil || o.SupportedPlatformsPids == nil {
		return nil, false
	}
	return o.SupportedPlatformsPids, true
}

// HasSupportedPlatformsPids returns a boolean if a field has been set.
func (o *CapabilityCpuEndpointDescriptorAllOf) HasSupportedPlatformsPids() bool {
	if o != nil && o.SupportedPlatformsPids != nil {
		return true
	}

	return false
}

// SetSupportedPlatformsPids gets a reference to the given []string and assigns it to the SupportedPlatformsPids field.
func (o *CapabilityCpuEndpointDescriptorAllOf) SetSupportedPlatformsPids(v []string) {
	o.SupportedPlatformsPids = v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *CapabilityCpuEndpointDescriptorAllOf) GetVendor() string {
	if o == nil || o.Vendor == nil {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityCpuEndpointDescriptorAllOf) GetVendorOk() (*string, bool) {
	if o == nil || o.Vendor == nil {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *CapabilityCpuEndpointDescriptorAllOf) HasVendor() bool {
	if o != nil && o.Vendor != nil {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *CapabilityCpuEndpointDescriptorAllOf) SetVendor(v string) {
	o.Vendor = &v
}

func (o CapabilityCpuEndpointDescriptorAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.Model != nil {
		toSerialize["Model"] = o.Model
	}
	if o.PartNumber != nil {
		toSerialize["PartNumber"] = o.PartNumber
	}
	if o.Pid != nil {
		toSerialize["Pid"] = o.Pid
	}
	if o.SupportedPlatformsPids != nil {
		toSerialize["SupportedPlatformsPids"] = o.SupportedPlatformsPids
	}
	if o.Vendor != nil {
		toSerialize["Vendor"] = o.Vendor
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CapabilityCpuEndpointDescriptorAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varCapabilityCpuEndpointDescriptorAllOf := _CapabilityCpuEndpointDescriptorAllOf{}

	if err = json.Unmarshal(bytes, &varCapabilityCpuEndpointDescriptorAllOf); err == nil {
		*o = CapabilityCpuEndpointDescriptorAllOf(varCapabilityCpuEndpointDescriptorAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "Model")
		delete(additionalProperties, "PartNumber")
		delete(additionalProperties, "Pid")
		delete(additionalProperties, "SupportedPlatformsPids")
		delete(additionalProperties, "Vendor")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCapabilityCpuEndpointDescriptorAllOf struct {
	value *CapabilityCpuEndpointDescriptorAllOf
	isSet bool
}

func (v NullableCapabilityCpuEndpointDescriptorAllOf) Get() *CapabilityCpuEndpointDescriptorAllOf {
	return v.value
}

func (v *NullableCapabilityCpuEndpointDescriptorAllOf) Set(val *CapabilityCpuEndpointDescriptorAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilityCpuEndpointDescriptorAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilityCpuEndpointDescriptorAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilityCpuEndpointDescriptorAllOf(val *CapabilityCpuEndpointDescriptorAllOf) *NullableCapabilityCpuEndpointDescriptorAllOf {
	return &NullableCapabilityCpuEndpointDescriptorAllOf{value: val, isSet: true}
}

func (v NullableCapabilityCpuEndpointDescriptorAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilityCpuEndpointDescriptorAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
