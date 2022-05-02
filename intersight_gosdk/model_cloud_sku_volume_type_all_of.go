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
)

// CloudSkuVolumeTypeAllOf Definition of the list of properties defined in 'cloud.SkuVolumeType', excluding properties defined in parent classes.
type CloudSkuVolumeTypeAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The units to measure IOPS.
	IopsUnit *string `json:"IopsUnit,omitempty"`
	// Indicates if this volume can be used as a boot volume.
	IsBootable *bool `json:"IsBootable,omitempty"`
	// Flag to indicate if this is a default volume. Default volumes will be used when an instance type is launched unless another volume type is specified.
	IsDefault *bool `json:"IsDefault,omitempty"`
	// Indicates if this volume type supports provisioned IOPS.
	IsProvisionedIops *bool `json:"IsProvisionedIops,omitempty"`
	// The max I/O operations per second that this volume type supports. Read or write numbers does not go beyond this value.
	MaxIops *float64 `json:"MaxIops,omitempty"`
	// The maximum read IOPS that this volume type supports.
	MaxReadIops *float64 `json:"MaxReadIops,omitempty"`
	// The maximum read throughput limit for this volume type.
	MaxReadThroughput *float64 `json:"MaxReadThroughput,omitempty"`
	// The maximum throughput limit for this volume type.
	MaxThroughput *float64 `json:"MaxThroughput,omitempty"`
	// The maximum storage size that this volume type supports.
	MaxVolumeSize *float64 `json:"MaxVolumeSize,omitempty"`
	// The maximum write IOPS that this volume type supports.
	MaxWriteIops *float64 `json:"MaxWriteIops,omitempty"`
	// The maximum write throughput limit for this volume type.
	MaxWriteThroughput *float64 `json:"MaxWriteThroughput,omitempty"`
	// The minimum storage size that this volume type supports.
	MinVolumeSize *float64 `json:"MinVolumeSize,omitempty"`
	// The units for measuring throughput.
	ThroughputUnit *string `json:"ThroughputUnit,omitempty"`
	// The volume type like gp2 or st1.
	Type *string `json:"Type,omitempty"`
	// The units for measuring volume size.
	VolumeSizeUnit       *string `json:"VolumeSizeUnit,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CloudSkuVolumeTypeAllOf CloudSkuVolumeTypeAllOf

// NewCloudSkuVolumeTypeAllOf instantiates a new CloudSkuVolumeTypeAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudSkuVolumeTypeAllOf(classId string, objectType string) *CloudSkuVolumeTypeAllOf {
	this := CloudSkuVolumeTypeAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCloudSkuVolumeTypeAllOfWithDefaults instantiates a new CloudSkuVolumeTypeAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudSkuVolumeTypeAllOfWithDefaults() *CloudSkuVolumeTypeAllOf {
	this := CloudSkuVolumeTypeAllOf{}
	var classId string = "cloud.SkuVolumeType"
	this.ClassId = classId
	var objectType string = "cloud.SkuVolumeType"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CloudSkuVolumeTypeAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CloudSkuVolumeTypeAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CloudSkuVolumeTypeAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CloudSkuVolumeTypeAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CloudSkuVolumeTypeAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CloudSkuVolumeTypeAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIopsUnit returns the IopsUnit field value if set, zero value otherwise.
func (o *CloudSkuVolumeTypeAllOf) GetIopsUnit() string {
	if o == nil || o.IopsUnit == nil {
		var ret string
		return ret
	}
	return *o.IopsUnit
}

// GetIopsUnitOk returns a tuple with the IopsUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudSkuVolumeTypeAllOf) GetIopsUnitOk() (*string, bool) {
	if o == nil || o.IopsUnit == nil {
		return nil, false
	}
	return o.IopsUnit, true
}

// HasIopsUnit returns a boolean if a field has been set.
func (o *CloudSkuVolumeTypeAllOf) HasIopsUnit() bool {
	if o != nil && o.IopsUnit != nil {
		return true
	}

	return false
}

// SetIopsUnit gets a reference to the given string and assigns it to the IopsUnit field.
func (o *CloudSkuVolumeTypeAllOf) SetIopsUnit(v string) {
	o.IopsUnit = &v
}

// GetIsBootable returns the IsBootable field value if set, zero value otherwise.
func (o *CloudSkuVolumeTypeAllOf) GetIsBootable() bool {
	if o == nil || o.IsBootable == nil {
		var ret bool
		return ret
	}
	return *o.IsBootable
}

// GetIsBootableOk returns a tuple with the IsBootable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudSkuVolumeTypeAllOf) GetIsBootableOk() (*bool, bool) {
	if o == nil || o.IsBootable == nil {
		return nil, false
	}
	return o.IsBootable, true
}

// HasIsBootable returns a boolean if a field has been set.
func (o *CloudSkuVolumeTypeAllOf) HasIsBootable() bool {
	if o != nil && o.IsBootable != nil {
		return true
	}

	return false
}

// SetIsBootable gets a reference to the given bool and assigns it to the IsBootable field.
func (o *CloudSkuVolumeTypeAllOf) SetIsBootable(v bool) {
	o.IsBootable = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *CloudSkuVolumeTypeAllOf) GetIsDefault() bool {
	if o == nil || o.IsDefault == nil {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudSkuVolumeTypeAllOf) GetIsDefaultOk() (*bool, bool) {
	if o == nil || o.IsDefault == nil {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *CloudSkuVolumeTypeAllOf) HasIsDefault() bool {
	if o != nil && o.IsDefault != nil {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *CloudSkuVolumeTypeAllOf) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// GetIsProvisionedIops returns the IsProvisionedIops field value if set, zero value otherwise.
func (o *CloudSkuVolumeTypeAllOf) GetIsProvisionedIops() bool {
	if o == nil || o.IsProvisionedIops == nil {
		var ret bool
		return ret
	}
	return *o.IsProvisionedIops
}

// GetIsProvisionedIopsOk returns a tuple with the IsProvisionedIops field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudSkuVolumeTypeAllOf) GetIsProvisionedIopsOk() (*bool, bool) {
	if o == nil || o.IsProvisionedIops == nil {
		return nil, false
	}
	return o.IsProvisionedIops, true
}

// HasIsProvisionedIops returns a boolean if a field has been set.
func (o *CloudSkuVolumeTypeAllOf) HasIsProvisionedIops() bool {
	if o != nil && o.IsProvisionedIops != nil {
		return true
	}

	return false
}

// SetIsProvisionedIops gets a reference to the given bool and assigns it to the IsProvisionedIops field.
func (o *CloudSkuVolumeTypeAllOf) SetIsProvisionedIops(v bool) {
	o.IsProvisionedIops = &v
}

// GetMaxIops returns the MaxIops field value if set, zero value otherwise.
func (o *CloudSkuVolumeTypeAllOf) GetMaxIops() float64 {
	if o == nil || o.MaxIops == nil {
		var ret float64
		return ret
	}
	return *o.MaxIops
}

// GetMaxIopsOk returns a tuple with the MaxIops field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudSkuVolumeTypeAllOf) GetMaxIopsOk() (*float64, bool) {
	if o == nil || o.MaxIops == nil {
		return nil, false
	}
	return o.MaxIops, true
}

// HasMaxIops returns a boolean if a field has been set.
func (o *CloudSkuVolumeTypeAllOf) HasMaxIops() bool {
	if o != nil && o.MaxIops != nil {
		return true
	}

	return false
}

// SetMaxIops gets a reference to the given float64 and assigns it to the MaxIops field.
func (o *CloudSkuVolumeTypeAllOf) SetMaxIops(v float64) {
	o.MaxIops = &v
}

// GetMaxReadIops returns the MaxReadIops field value if set, zero value otherwise.
func (o *CloudSkuVolumeTypeAllOf) GetMaxReadIops() float64 {
	if o == nil || o.MaxReadIops == nil {
		var ret float64
		return ret
	}
	return *o.MaxReadIops
}

// GetMaxReadIopsOk returns a tuple with the MaxReadIops field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudSkuVolumeTypeAllOf) GetMaxReadIopsOk() (*float64, bool) {
	if o == nil || o.MaxReadIops == nil {
		return nil, false
	}
	return o.MaxReadIops, true
}

// HasMaxReadIops returns a boolean if a field has been set.
func (o *CloudSkuVolumeTypeAllOf) HasMaxReadIops() bool {
	if o != nil && o.MaxReadIops != nil {
		return true
	}

	return false
}

// SetMaxReadIops gets a reference to the given float64 and assigns it to the MaxReadIops field.
func (o *CloudSkuVolumeTypeAllOf) SetMaxReadIops(v float64) {
	o.MaxReadIops = &v
}

// GetMaxReadThroughput returns the MaxReadThroughput field value if set, zero value otherwise.
func (o *CloudSkuVolumeTypeAllOf) GetMaxReadThroughput() float64 {
	if o == nil || o.MaxReadThroughput == nil {
		var ret float64
		return ret
	}
	return *o.MaxReadThroughput
}

// GetMaxReadThroughputOk returns a tuple with the MaxReadThroughput field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudSkuVolumeTypeAllOf) GetMaxReadThroughputOk() (*float64, bool) {
	if o == nil || o.MaxReadThroughput == nil {
		return nil, false
	}
	return o.MaxReadThroughput, true
}

// HasMaxReadThroughput returns a boolean if a field has been set.
func (o *CloudSkuVolumeTypeAllOf) HasMaxReadThroughput() bool {
	if o != nil && o.MaxReadThroughput != nil {
		return true
	}

	return false
}

// SetMaxReadThroughput gets a reference to the given float64 and assigns it to the MaxReadThroughput field.
func (o *CloudSkuVolumeTypeAllOf) SetMaxReadThroughput(v float64) {
	o.MaxReadThroughput = &v
}

// GetMaxThroughput returns the MaxThroughput field value if set, zero value otherwise.
func (o *CloudSkuVolumeTypeAllOf) GetMaxThroughput() float64 {
	if o == nil || o.MaxThroughput == nil {
		var ret float64
		return ret
	}
	return *o.MaxThroughput
}

// GetMaxThroughputOk returns a tuple with the MaxThroughput field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudSkuVolumeTypeAllOf) GetMaxThroughputOk() (*float64, bool) {
	if o == nil || o.MaxThroughput == nil {
		return nil, false
	}
	return o.MaxThroughput, true
}

// HasMaxThroughput returns a boolean if a field has been set.
func (o *CloudSkuVolumeTypeAllOf) HasMaxThroughput() bool {
	if o != nil && o.MaxThroughput != nil {
		return true
	}

	return false
}

// SetMaxThroughput gets a reference to the given float64 and assigns it to the MaxThroughput field.
func (o *CloudSkuVolumeTypeAllOf) SetMaxThroughput(v float64) {
	o.MaxThroughput = &v
}

// GetMaxVolumeSize returns the MaxVolumeSize field value if set, zero value otherwise.
func (o *CloudSkuVolumeTypeAllOf) GetMaxVolumeSize() float64 {
	if o == nil || o.MaxVolumeSize == nil {
		var ret float64
		return ret
	}
	return *o.MaxVolumeSize
}

// GetMaxVolumeSizeOk returns a tuple with the MaxVolumeSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudSkuVolumeTypeAllOf) GetMaxVolumeSizeOk() (*float64, bool) {
	if o == nil || o.MaxVolumeSize == nil {
		return nil, false
	}
	return o.MaxVolumeSize, true
}

// HasMaxVolumeSize returns a boolean if a field has been set.
func (o *CloudSkuVolumeTypeAllOf) HasMaxVolumeSize() bool {
	if o != nil && o.MaxVolumeSize != nil {
		return true
	}

	return false
}

// SetMaxVolumeSize gets a reference to the given float64 and assigns it to the MaxVolumeSize field.
func (o *CloudSkuVolumeTypeAllOf) SetMaxVolumeSize(v float64) {
	o.MaxVolumeSize = &v
}

// GetMaxWriteIops returns the MaxWriteIops field value if set, zero value otherwise.
func (o *CloudSkuVolumeTypeAllOf) GetMaxWriteIops() float64 {
	if o == nil || o.MaxWriteIops == nil {
		var ret float64
		return ret
	}
	return *o.MaxWriteIops
}

// GetMaxWriteIopsOk returns a tuple with the MaxWriteIops field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudSkuVolumeTypeAllOf) GetMaxWriteIopsOk() (*float64, bool) {
	if o == nil || o.MaxWriteIops == nil {
		return nil, false
	}
	return o.MaxWriteIops, true
}

// HasMaxWriteIops returns a boolean if a field has been set.
func (o *CloudSkuVolumeTypeAllOf) HasMaxWriteIops() bool {
	if o != nil && o.MaxWriteIops != nil {
		return true
	}

	return false
}

// SetMaxWriteIops gets a reference to the given float64 and assigns it to the MaxWriteIops field.
func (o *CloudSkuVolumeTypeAllOf) SetMaxWriteIops(v float64) {
	o.MaxWriteIops = &v
}

// GetMaxWriteThroughput returns the MaxWriteThroughput field value if set, zero value otherwise.
func (o *CloudSkuVolumeTypeAllOf) GetMaxWriteThroughput() float64 {
	if o == nil || o.MaxWriteThroughput == nil {
		var ret float64
		return ret
	}
	return *o.MaxWriteThroughput
}

// GetMaxWriteThroughputOk returns a tuple with the MaxWriteThroughput field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudSkuVolumeTypeAllOf) GetMaxWriteThroughputOk() (*float64, bool) {
	if o == nil || o.MaxWriteThroughput == nil {
		return nil, false
	}
	return o.MaxWriteThroughput, true
}

// HasMaxWriteThroughput returns a boolean if a field has been set.
func (o *CloudSkuVolumeTypeAllOf) HasMaxWriteThroughput() bool {
	if o != nil && o.MaxWriteThroughput != nil {
		return true
	}

	return false
}

// SetMaxWriteThroughput gets a reference to the given float64 and assigns it to the MaxWriteThroughput field.
func (o *CloudSkuVolumeTypeAllOf) SetMaxWriteThroughput(v float64) {
	o.MaxWriteThroughput = &v
}

// GetMinVolumeSize returns the MinVolumeSize field value if set, zero value otherwise.
func (o *CloudSkuVolumeTypeAllOf) GetMinVolumeSize() float64 {
	if o == nil || o.MinVolumeSize == nil {
		var ret float64
		return ret
	}
	return *o.MinVolumeSize
}

// GetMinVolumeSizeOk returns a tuple with the MinVolumeSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudSkuVolumeTypeAllOf) GetMinVolumeSizeOk() (*float64, bool) {
	if o == nil || o.MinVolumeSize == nil {
		return nil, false
	}
	return o.MinVolumeSize, true
}

// HasMinVolumeSize returns a boolean if a field has been set.
func (o *CloudSkuVolumeTypeAllOf) HasMinVolumeSize() bool {
	if o != nil && o.MinVolumeSize != nil {
		return true
	}

	return false
}

// SetMinVolumeSize gets a reference to the given float64 and assigns it to the MinVolumeSize field.
func (o *CloudSkuVolumeTypeAllOf) SetMinVolumeSize(v float64) {
	o.MinVolumeSize = &v
}

// GetThroughputUnit returns the ThroughputUnit field value if set, zero value otherwise.
func (o *CloudSkuVolumeTypeAllOf) GetThroughputUnit() string {
	if o == nil || o.ThroughputUnit == nil {
		var ret string
		return ret
	}
	return *o.ThroughputUnit
}

// GetThroughputUnitOk returns a tuple with the ThroughputUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudSkuVolumeTypeAllOf) GetThroughputUnitOk() (*string, bool) {
	if o == nil || o.ThroughputUnit == nil {
		return nil, false
	}
	return o.ThroughputUnit, true
}

// HasThroughputUnit returns a boolean if a field has been set.
func (o *CloudSkuVolumeTypeAllOf) HasThroughputUnit() bool {
	if o != nil && o.ThroughputUnit != nil {
		return true
	}

	return false
}

// SetThroughputUnit gets a reference to the given string and assigns it to the ThroughputUnit field.
func (o *CloudSkuVolumeTypeAllOf) SetThroughputUnit(v string) {
	o.ThroughputUnit = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CloudSkuVolumeTypeAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudSkuVolumeTypeAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CloudSkuVolumeTypeAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CloudSkuVolumeTypeAllOf) SetType(v string) {
	o.Type = &v
}

// GetVolumeSizeUnit returns the VolumeSizeUnit field value if set, zero value otherwise.
func (o *CloudSkuVolumeTypeAllOf) GetVolumeSizeUnit() string {
	if o == nil || o.VolumeSizeUnit == nil {
		var ret string
		return ret
	}
	return *o.VolumeSizeUnit
}

// GetVolumeSizeUnitOk returns a tuple with the VolumeSizeUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudSkuVolumeTypeAllOf) GetVolumeSizeUnitOk() (*string, bool) {
	if o == nil || o.VolumeSizeUnit == nil {
		return nil, false
	}
	return o.VolumeSizeUnit, true
}

// HasVolumeSizeUnit returns a boolean if a field has been set.
func (o *CloudSkuVolumeTypeAllOf) HasVolumeSizeUnit() bool {
	if o != nil && o.VolumeSizeUnit != nil {
		return true
	}

	return false
}

// SetVolumeSizeUnit gets a reference to the given string and assigns it to the VolumeSizeUnit field.
func (o *CloudSkuVolumeTypeAllOf) SetVolumeSizeUnit(v string) {
	o.VolumeSizeUnit = &v
}

func (o CloudSkuVolumeTypeAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.IopsUnit != nil {
		toSerialize["IopsUnit"] = o.IopsUnit
	}
	if o.IsBootable != nil {
		toSerialize["IsBootable"] = o.IsBootable
	}
	if o.IsDefault != nil {
		toSerialize["IsDefault"] = o.IsDefault
	}
	if o.IsProvisionedIops != nil {
		toSerialize["IsProvisionedIops"] = o.IsProvisionedIops
	}
	if o.MaxIops != nil {
		toSerialize["MaxIops"] = o.MaxIops
	}
	if o.MaxReadIops != nil {
		toSerialize["MaxReadIops"] = o.MaxReadIops
	}
	if o.MaxReadThroughput != nil {
		toSerialize["MaxReadThroughput"] = o.MaxReadThroughput
	}
	if o.MaxThroughput != nil {
		toSerialize["MaxThroughput"] = o.MaxThroughput
	}
	if o.MaxVolumeSize != nil {
		toSerialize["MaxVolumeSize"] = o.MaxVolumeSize
	}
	if o.MaxWriteIops != nil {
		toSerialize["MaxWriteIops"] = o.MaxWriteIops
	}
	if o.MaxWriteThroughput != nil {
		toSerialize["MaxWriteThroughput"] = o.MaxWriteThroughput
	}
	if o.MinVolumeSize != nil {
		toSerialize["MinVolumeSize"] = o.MinVolumeSize
	}
	if o.ThroughputUnit != nil {
		toSerialize["ThroughputUnit"] = o.ThroughputUnit
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	if o.VolumeSizeUnit != nil {
		toSerialize["VolumeSizeUnit"] = o.VolumeSizeUnit
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CloudSkuVolumeTypeAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varCloudSkuVolumeTypeAllOf := _CloudSkuVolumeTypeAllOf{}

	if err = json.Unmarshal(bytes, &varCloudSkuVolumeTypeAllOf); err == nil {
		*o = CloudSkuVolumeTypeAllOf(varCloudSkuVolumeTypeAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "IopsUnit")
		delete(additionalProperties, "IsBootable")
		delete(additionalProperties, "IsDefault")
		delete(additionalProperties, "IsProvisionedIops")
		delete(additionalProperties, "MaxIops")
		delete(additionalProperties, "MaxReadIops")
		delete(additionalProperties, "MaxReadThroughput")
		delete(additionalProperties, "MaxThroughput")
		delete(additionalProperties, "MaxVolumeSize")
		delete(additionalProperties, "MaxWriteIops")
		delete(additionalProperties, "MaxWriteThroughput")
		delete(additionalProperties, "MinVolumeSize")
		delete(additionalProperties, "ThroughputUnit")
		delete(additionalProperties, "Type")
		delete(additionalProperties, "VolumeSizeUnit")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCloudSkuVolumeTypeAllOf struct {
	value *CloudSkuVolumeTypeAllOf
	isSet bool
}

func (v NullableCloudSkuVolumeTypeAllOf) Get() *CloudSkuVolumeTypeAllOf {
	return v.value
}

func (v *NullableCloudSkuVolumeTypeAllOf) Set(val *CloudSkuVolumeTypeAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudSkuVolumeTypeAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudSkuVolumeTypeAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudSkuVolumeTypeAllOf(val *CloudSkuVolumeTypeAllOf) *NullableCloudSkuVolumeTypeAllOf {
	return &NullableCloudSkuVolumeTypeAllOf{value: val, isSet: true}
}

func (v NullableCloudSkuVolumeTypeAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudSkuVolumeTypeAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
