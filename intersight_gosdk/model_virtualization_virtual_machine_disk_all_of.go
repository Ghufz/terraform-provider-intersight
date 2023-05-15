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

// VirtualizationVirtualMachineDiskAllOf Definition of the list of properties defined in 'virtualization.VirtualMachineDisk', excluding properties defined in parent classes.
type VirtualizationVirtualMachineDiskAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Disk bus name given for a virtual machine. * `virtio` - Disk uses the same paths as a bare-metal system. This simplifies physical-to-virtual and virtual-to-virtual migration. * `sata` - Serial ATA (SATA, abbreviated from Serial AT Attachment) is a computer bus interface that connects host bus adapters to mass storage devices such as hard disk drives, optical drives, and solid-state drives. * `scsi` - SCSI (Small Computer System Interface) bus used..
	Bus *string `json:"Bus,omitempty"`
	// Virtual machine network bridge name.
	Name *string `json:"Name,omitempty"`
	// Priority order of the disk.
	Order *int64 `json:"Order,omitempty"`
	// Disk type hdd or cdrom for a virtual machine. * `hdd` - Allows the virtual machine to mount disk from hard disk drive (hdd) image. * `cdrom` - Allows the virtual machine to mount disk from compact disk (cd) image.
	Type        *string                                 `json:"Type,omitempty"`
	VirtualDisk NullableVirtualizationVirtualDiskConfig `json:"VirtualDisk,omitempty"`
	// Name of the existing virtual disk to be attached to the Virtual Machine.
	VirtualDiskReference *string `json:"VirtualDiskReference,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationVirtualMachineDiskAllOf VirtualizationVirtualMachineDiskAllOf

// NewVirtualizationVirtualMachineDiskAllOf instantiates a new VirtualizationVirtualMachineDiskAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationVirtualMachineDiskAllOf(classId string, objectType string) *VirtualizationVirtualMachineDiskAllOf {
	this := VirtualizationVirtualMachineDiskAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var bus string = "virtio"
	this.Bus = &bus
	var type_ string = "hdd"
	this.Type = &type_
	return &this
}

// NewVirtualizationVirtualMachineDiskAllOfWithDefaults instantiates a new VirtualizationVirtualMachineDiskAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationVirtualMachineDiskAllOfWithDefaults() *VirtualizationVirtualMachineDiskAllOf {
	this := VirtualizationVirtualMachineDiskAllOf{}
	var classId string = "virtualization.VirtualMachineDisk"
	this.ClassId = classId
	var objectType string = "virtualization.VirtualMachineDisk"
	this.ObjectType = objectType
	var bus string = "virtio"
	this.Bus = &bus
	var type_ string = "hdd"
	this.Type = &type_
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationVirtualMachineDiskAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualMachineDiskAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationVirtualMachineDiskAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationVirtualMachineDiskAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualMachineDiskAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationVirtualMachineDiskAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBus returns the Bus field value if set, zero value otherwise.
func (o *VirtualizationVirtualMachineDiskAllOf) GetBus() string {
	if o == nil || o.Bus == nil {
		var ret string
		return ret
	}
	return *o.Bus
}

// GetBusOk returns a tuple with the Bus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualMachineDiskAllOf) GetBusOk() (*string, bool) {
	if o == nil || o.Bus == nil {
		return nil, false
	}
	return o.Bus, true
}

// HasBus returns a boolean if a field has been set.
func (o *VirtualizationVirtualMachineDiskAllOf) HasBus() bool {
	if o != nil && o.Bus != nil {
		return true
	}

	return false
}

// SetBus gets a reference to the given string and assigns it to the Bus field.
func (o *VirtualizationVirtualMachineDiskAllOf) SetBus(v string) {
	o.Bus = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VirtualizationVirtualMachineDiskAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualMachineDiskAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VirtualizationVirtualMachineDiskAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VirtualizationVirtualMachineDiskAllOf) SetName(v string) {
	o.Name = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *VirtualizationVirtualMachineDiskAllOf) GetOrder() int64 {
	if o == nil || o.Order == nil {
		var ret int64
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualMachineDiskAllOf) GetOrderOk() (*int64, bool) {
	if o == nil || o.Order == nil {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *VirtualizationVirtualMachineDiskAllOf) HasOrder() bool {
	if o != nil && o.Order != nil {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int64 and assigns it to the Order field.
func (o *VirtualizationVirtualMachineDiskAllOf) SetOrder(v int64) {
	o.Order = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *VirtualizationVirtualMachineDiskAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualMachineDiskAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *VirtualizationVirtualMachineDiskAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *VirtualizationVirtualMachineDiskAllOf) SetType(v string) {
	o.Type = &v
}

// GetVirtualDisk returns the VirtualDisk field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationVirtualMachineDiskAllOf) GetVirtualDisk() VirtualizationVirtualDiskConfig {
	if o == nil || o.VirtualDisk.Get() == nil {
		var ret VirtualizationVirtualDiskConfig
		return ret
	}
	return *o.VirtualDisk.Get()
}

// GetVirtualDiskOk returns a tuple with the VirtualDisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationVirtualMachineDiskAllOf) GetVirtualDiskOk() (*VirtualizationVirtualDiskConfig, bool) {
	if o == nil {
		return nil, false
	}
	return o.VirtualDisk.Get(), o.VirtualDisk.IsSet()
}

// HasVirtualDisk returns a boolean if a field has been set.
func (o *VirtualizationVirtualMachineDiskAllOf) HasVirtualDisk() bool {
	if o != nil && o.VirtualDisk.IsSet() {
		return true
	}

	return false
}

// SetVirtualDisk gets a reference to the given NullableVirtualizationVirtualDiskConfig and assigns it to the VirtualDisk field.
func (o *VirtualizationVirtualMachineDiskAllOf) SetVirtualDisk(v VirtualizationVirtualDiskConfig) {
	o.VirtualDisk.Set(&v)
}

// SetVirtualDiskNil sets the value for VirtualDisk to be an explicit nil
func (o *VirtualizationVirtualMachineDiskAllOf) SetVirtualDiskNil() {
	o.VirtualDisk.Set(nil)
}

// UnsetVirtualDisk ensures that no value is present for VirtualDisk, not even an explicit nil
func (o *VirtualizationVirtualMachineDiskAllOf) UnsetVirtualDisk() {
	o.VirtualDisk.Unset()
}

// GetVirtualDiskReference returns the VirtualDiskReference field value if set, zero value otherwise.
func (o *VirtualizationVirtualMachineDiskAllOf) GetVirtualDiskReference() string {
	if o == nil || o.VirtualDiskReference == nil {
		var ret string
		return ret
	}
	return *o.VirtualDiskReference
}

// GetVirtualDiskReferenceOk returns a tuple with the VirtualDiskReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualMachineDiskAllOf) GetVirtualDiskReferenceOk() (*string, bool) {
	if o == nil || o.VirtualDiskReference == nil {
		return nil, false
	}
	return o.VirtualDiskReference, true
}

// HasVirtualDiskReference returns a boolean if a field has been set.
func (o *VirtualizationVirtualMachineDiskAllOf) HasVirtualDiskReference() bool {
	if o != nil && o.VirtualDiskReference != nil {
		return true
	}

	return false
}

// SetVirtualDiskReference gets a reference to the given string and assigns it to the VirtualDiskReference field.
func (o *VirtualizationVirtualMachineDiskAllOf) SetVirtualDiskReference(v string) {
	o.VirtualDiskReference = &v
}

func (o VirtualizationVirtualMachineDiskAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Bus != nil {
		toSerialize["Bus"] = o.Bus
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Order != nil {
		toSerialize["Order"] = o.Order
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	if o.VirtualDisk.IsSet() {
		toSerialize["VirtualDisk"] = o.VirtualDisk.Get()
	}
	if o.VirtualDiskReference != nil {
		toSerialize["VirtualDiskReference"] = o.VirtualDiskReference
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationVirtualMachineDiskAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varVirtualizationVirtualMachineDiskAllOf := _VirtualizationVirtualMachineDiskAllOf{}

	if err = json.Unmarshal(bytes, &varVirtualizationVirtualMachineDiskAllOf); err == nil {
		*o = VirtualizationVirtualMachineDiskAllOf(varVirtualizationVirtualMachineDiskAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Bus")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Order")
		delete(additionalProperties, "Type")
		delete(additionalProperties, "VirtualDisk")
		delete(additionalProperties, "VirtualDiskReference")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVirtualizationVirtualMachineDiskAllOf struct {
	value *VirtualizationVirtualMachineDiskAllOf
	isSet bool
}

func (v NullableVirtualizationVirtualMachineDiskAllOf) Get() *VirtualizationVirtualMachineDiskAllOf {
	return v.value
}

func (v *NullableVirtualizationVirtualMachineDiskAllOf) Set(val *VirtualizationVirtualMachineDiskAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationVirtualMachineDiskAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationVirtualMachineDiskAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationVirtualMachineDiskAllOf(val *VirtualizationVirtualMachineDiskAllOf) *NullableVirtualizationVirtualMachineDiskAllOf {
	return &NullableVirtualizationVirtualMachineDiskAllOf{value: val, isSet: true}
}

func (v NullableVirtualizationVirtualMachineDiskAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationVirtualMachineDiskAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
