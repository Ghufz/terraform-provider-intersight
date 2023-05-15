/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-11765
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// VirtualizationVolumeInfoAllOf Definition of the list of properties defined in 'virtualization.VolumeInfo', excluding properties defined in parent classes.
type VirtualizationVolumeInfoAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Set to true, if the volume should be a root disk.
	Bootable *bool `json:"Bootable,omitempty"`
	// Set to true, to delete the volume on termination of the VM the volume is attached to.
	DeleteOnTermination *bool `json:"DeleteOnTermination,omitempty"`
	// Set to true, if the volume should be encrypted.
	Encryption *bool `json:"Encryption,omitempty"`
	// IOPS for the volume for applicable volume types.
	Iops *int64 `json:"Iops,omitempty"`
	// Order of the disk attachment to the VM.
	Order *int64 `json:"Order,omitempty"`
	// Throughput for the volume for applicable volume types.
	Throughput *int64 `json:"Throughput,omitempty"`
	// Unique volume id assigned by the cloud provider.
	VolumeId *string `json:"VolumeId,omitempty"`
	// Name assigned to the volume created.
	VolumeName *string `json:"VolumeName,omitempty"`
	// Size of the volume created in GB.
	VolumeSize *int64 `json:"VolumeSize,omitempty"`
	// Id of the volume or storage type of this volume.
	VolumeType           *string `json:"VolumeType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationVolumeInfoAllOf VirtualizationVolumeInfoAllOf

// NewVirtualizationVolumeInfoAllOf instantiates a new VirtualizationVolumeInfoAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationVolumeInfoAllOf(classId string, objectType string) *VirtualizationVolumeInfoAllOf {
	this := VirtualizationVolumeInfoAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVirtualizationVolumeInfoAllOfWithDefaults instantiates a new VirtualizationVolumeInfoAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationVolumeInfoAllOfWithDefaults() *VirtualizationVolumeInfoAllOf {
	this := VirtualizationVolumeInfoAllOf{}
	var classId string = "virtualization.VolumeInfo"
	this.ClassId = classId
	var objectType string = "virtualization.VolumeInfo"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationVolumeInfoAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVolumeInfoAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationVolumeInfoAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationVolumeInfoAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVolumeInfoAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationVolumeInfoAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBootable returns the Bootable field value if set, zero value otherwise.
func (o *VirtualizationVolumeInfoAllOf) GetBootable() bool {
	if o == nil || o.Bootable == nil {
		var ret bool
		return ret
	}
	return *o.Bootable
}

// GetBootableOk returns a tuple with the Bootable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVolumeInfoAllOf) GetBootableOk() (*bool, bool) {
	if o == nil || o.Bootable == nil {
		return nil, false
	}
	return o.Bootable, true
}

// HasBootable returns a boolean if a field has been set.
func (o *VirtualizationVolumeInfoAllOf) HasBootable() bool {
	if o != nil && o.Bootable != nil {
		return true
	}

	return false
}

// SetBootable gets a reference to the given bool and assigns it to the Bootable field.
func (o *VirtualizationVolumeInfoAllOf) SetBootable(v bool) {
	o.Bootable = &v
}

// GetDeleteOnTermination returns the DeleteOnTermination field value if set, zero value otherwise.
func (o *VirtualizationVolumeInfoAllOf) GetDeleteOnTermination() bool {
	if o == nil || o.DeleteOnTermination == nil {
		var ret bool
		return ret
	}
	return *o.DeleteOnTermination
}

// GetDeleteOnTerminationOk returns a tuple with the DeleteOnTermination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVolumeInfoAllOf) GetDeleteOnTerminationOk() (*bool, bool) {
	if o == nil || o.DeleteOnTermination == nil {
		return nil, false
	}
	return o.DeleteOnTermination, true
}

// HasDeleteOnTermination returns a boolean if a field has been set.
func (o *VirtualizationVolumeInfoAllOf) HasDeleteOnTermination() bool {
	if o != nil && o.DeleteOnTermination != nil {
		return true
	}

	return false
}

// SetDeleteOnTermination gets a reference to the given bool and assigns it to the DeleteOnTermination field.
func (o *VirtualizationVolumeInfoAllOf) SetDeleteOnTermination(v bool) {
	o.DeleteOnTermination = &v
}

// GetEncryption returns the Encryption field value if set, zero value otherwise.
func (o *VirtualizationVolumeInfoAllOf) GetEncryption() bool {
	if o == nil || o.Encryption == nil {
		var ret bool
		return ret
	}
	return *o.Encryption
}

// GetEncryptionOk returns a tuple with the Encryption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVolumeInfoAllOf) GetEncryptionOk() (*bool, bool) {
	if o == nil || o.Encryption == nil {
		return nil, false
	}
	return o.Encryption, true
}

// HasEncryption returns a boolean if a field has been set.
func (o *VirtualizationVolumeInfoAllOf) HasEncryption() bool {
	if o != nil && o.Encryption != nil {
		return true
	}

	return false
}

// SetEncryption gets a reference to the given bool and assigns it to the Encryption field.
func (o *VirtualizationVolumeInfoAllOf) SetEncryption(v bool) {
	o.Encryption = &v
}

// GetIops returns the Iops field value if set, zero value otherwise.
func (o *VirtualizationVolumeInfoAllOf) GetIops() int64 {
	if o == nil || o.Iops == nil {
		var ret int64
		return ret
	}
	return *o.Iops
}

// GetIopsOk returns a tuple with the Iops field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVolumeInfoAllOf) GetIopsOk() (*int64, bool) {
	if o == nil || o.Iops == nil {
		return nil, false
	}
	return o.Iops, true
}

// HasIops returns a boolean if a field has been set.
func (o *VirtualizationVolumeInfoAllOf) HasIops() bool {
	if o != nil && o.Iops != nil {
		return true
	}

	return false
}

// SetIops gets a reference to the given int64 and assigns it to the Iops field.
func (o *VirtualizationVolumeInfoAllOf) SetIops(v int64) {
	o.Iops = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *VirtualizationVolumeInfoAllOf) GetOrder() int64 {
	if o == nil || o.Order == nil {
		var ret int64
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVolumeInfoAllOf) GetOrderOk() (*int64, bool) {
	if o == nil || o.Order == nil {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *VirtualizationVolumeInfoAllOf) HasOrder() bool {
	if o != nil && o.Order != nil {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int64 and assigns it to the Order field.
func (o *VirtualizationVolumeInfoAllOf) SetOrder(v int64) {
	o.Order = &v
}

// GetThroughput returns the Throughput field value if set, zero value otherwise.
func (o *VirtualizationVolumeInfoAllOf) GetThroughput() int64 {
	if o == nil || o.Throughput == nil {
		var ret int64
		return ret
	}
	return *o.Throughput
}

// GetThroughputOk returns a tuple with the Throughput field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVolumeInfoAllOf) GetThroughputOk() (*int64, bool) {
	if o == nil || o.Throughput == nil {
		return nil, false
	}
	return o.Throughput, true
}

// HasThroughput returns a boolean if a field has been set.
func (o *VirtualizationVolumeInfoAllOf) HasThroughput() bool {
	if o != nil && o.Throughput != nil {
		return true
	}

	return false
}

// SetThroughput gets a reference to the given int64 and assigns it to the Throughput field.
func (o *VirtualizationVolumeInfoAllOf) SetThroughput(v int64) {
	o.Throughput = &v
}

// GetVolumeId returns the VolumeId field value if set, zero value otherwise.
func (o *VirtualizationVolumeInfoAllOf) GetVolumeId() string {
	if o == nil || o.VolumeId == nil {
		var ret string
		return ret
	}
	return *o.VolumeId
}

// GetVolumeIdOk returns a tuple with the VolumeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVolumeInfoAllOf) GetVolumeIdOk() (*string, bool) {
	if o == nil || o.VolumeId == nil {
		return nil, false
	}
	return o.VolumeId, true
}

// HasVolumeId returns a boolean if a field has been set.
func (o *VirtualizationVolumeInfoAllOf) HasVolumeId() bool {
	if o != nil && o.VolumeId != nil {
		return true
	}

	return false
}

// SetVolumeId gets a reference to the given string and assigns it to the VolumeId field.
func (o *VirtualizationVolumeInfoAllOf) SetVolumeId(v string) {
	o.VolumeId = &v
}

// GetVolumeName returns the VolumeName field value if set, zero value otherwise.
func (o *VirtualizationVolumeInfoAllOf) GetVolumeName() string {
	if o == nil || o.VolumeName == nil {
		var ret string
		return ret
	}
	return *o.VolumeName
}

// GetVolumeNameOk returns a tuple with the VolumeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVolumeInfoAllOf) GetVolumeNameOk() (*string, bool) {
	if o == nil || o.VolumeName == nil {
		return nil, false
	}
	return o.VolumeName, true
}

// HasVolumeName returns a boolean if a field has been set.
func (o *VirtualizationVolumeInfoAllOf) HasVolumeName() bool {
	if o != nil && o.VolumeName != nil {
		return true
	}

	return false
}

// SetVolumeName gets a reference to the given string and assigns it to the VolumeName field.
func (o *VirtualizationVolumeInfoAllOf) SetVolumeName(v string) {
	o.VolumeName = &v
}

// GetVolumeSize returns the VolumeSize field value if set, zero value otherwise.
func (o *VirtualizationVolumeInfoAllOf) GetVolumeSize() int64 {
	if o == nil || o.VolumeSize == nil {
		var ret int64
		return ret
	}
	return *o.VolumeSize
}

// GetVolumeSizeOk returns a tuple with the VolumeSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVolumeInfoAllOf) GetVolumeSizeOk() (*int64, bool) {
	if o == nil || o.VolumeSize == nil {
		return nil, false
	}
	return o.VolumeSize, true
}

// HasVolumeSize returns a boolean if a field has been set.
func (o *VirtualizationVolumeInfoAllOf) HasVolumeSize() bool {
	if o != nil && o.VolumeSize != nil {
		return true
	}

	return false
}

// SetVolumeSize gets a reference to the given int64 and assigns it to the VolumeSize field.
func (o *VirtualizationVolumeInfoAllOf) SetVolumeSize(v int64) {
	o.VolumeSize = &v
}

// GetVolumeType returns the VolumeType field value if set, zero value otherwise.
func (o *VirtualizationVolumeInfoAllOf) GetVolumeType() string {
	if o == nil || o.VolumeType == nil {
		var ret string
		return ret
	}
	return *o.VolumeType
}

// GetVolumeTypeOk returns a tuple with the VolumeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVolumeInfoAllOf) GetVolumeTypeOk() (*string, bool) {
	if o == nil || o.VolumeType == nil {
		return nil, false
	}
	return o.VolumeType, true
}

// HasVolumeType returns a boolean if a field has been set.
func (o *VirtualizationVolumeInfoAllOf) HasVolumeType() bool {
	if o != nil && o.VolumeType != nil {
		return true
	}

	return false
}

// SetVolumeType gets a reference to the given string and assigns it to the VolumeType field.
func (o *VirtualizationVolumeInfoAllOf) SetVolumeType(v string) {
	o.VolumeType = &v
}

func (o VirtualizationVolumeInfoAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Bootable != nil {
		toSerialize["Bootable"] = o.Bootable
	}
	if o.DeleteOnTermination != nil {
		toSerialize["DeleteOnTermination"] = o.DeleteOnTermination
	}
	if o.Encryption != nil {
		toSerialize["Encryption"] = o.Encryption
	}
	if o.Iops != nil {
		toSerialize["Iops"] = o.Iops
	}
	if o.Order != nil {
		toSerialize["Order"] = o.Order
	}
	if o.Throughput != nil {
		toSerialize["Throughput"] = o.Throughput
	}
	if o.VolumeId != nil {
		toSerialize["VolumeId"] = o.VolumeId
	}
	if o.VolumeName != nil {
		toSerialize["VolumeName"] = o.VolumeName
	}
	if o.VolumeSize != nil {
		toSerialize["VolumeSize"] = o.VolumeSize
	}
	if o.VolumeType != nil {
		toSerialize["VolumeType"] = o.VolumeType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationVolumeInfoAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varVirtualizationVolumeInfoAllOf := _VirtualizationVolumeInfoAllOf{}

	if err = json.Unmarshal(bytes, &varVirtualizationVolumeInfoAllOf); err == nil {
		*o = VirtualizationVolumeInfoAllOf(varVirtualizationVolumeInfoAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Bootable")
		delete(additionalProperties, "DeleteOnTermination")
		delete(additionalProperties, "Encryption")
		delete(additionalProperties, "Iops")
		delete(additionalProperties, "Order")
		delete(additionalProperties, "Throughput")
		delete(additionalProperties, "VolumeId")
		delete(additionalProperties, "VolumeName")
		delete(additionalProperties, "VolumeSize")
		delete(additionalProperties, "VolumeType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVirtualizationVolumeInfoAllOf struct {
	value *VirtualizationVolumeInfoAllOf
	isSet bool
}

func (v NullableVirtualizationVolumeInfoAllOf) Get() *VirtualizationVolumeInfoAllOf {
	return v.value
}

func (v *NullableVirtualizationVolumeInfoAllOf) Set(val *VirtualizationVolumeInfoAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationVolumeInfoAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationVolumeInfoAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationVolumeInfoAllOf(val *VirtualizationVolumeInfoAllOf) *NullableVirtualizationVolumeInfoAllOf {
	return &NullableVirtualizationVolumeInfoAllOf{value: val, isSet: true}
}

func (v NullableVirtualizationVolumeInfoAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationVolumeInfoAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
