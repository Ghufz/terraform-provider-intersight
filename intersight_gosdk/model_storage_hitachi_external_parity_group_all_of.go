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

// StorageHitachiExternalParityGroupAllOf Definition of the list of properties defined in 'storage.HitachiExternalParityGroup', excluding properties defined in parent classes.
type StorageHitachiExternalParityGroupAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// From among the open volumes in the external parity group, the total capacity of volumes to which paths can be allocated (KB).
	AllocatableOpenVolumeCapacity *int64 `json:"AllocatableOpenVolumeCapacity,omitempty"`
	// From among the open volumes in the external parity group, the total capacity of volumes to which paths are allocated (KB).
	AllocatedOpenVolumeCapacity *int64 `json:"AllocatedOpenVolumeCapacity,omitempty"`
	// Available capacity of the external parity group, represented in bytes.
	AvailableVolumeCapacity *int64 `json:"AvailableVolumeCapacity,omitempty"`
	// Number of CLPR to which the external parity group belongs.
	ClprId *int64 `json:"ClprId,omitempty"`
	// Emulation type of the external parity group.
	EmulationType *string `json:"EmulationType,omitempty"`
	// Storage system that is connected using the external storage connection functionality of Universal Volume Manager.
	ExternalProductId *string `json:"ExternalProductId,omitempty"`
	// Maximum capacity of the non-volume areas in the external parity group (KB).
	LargestAvailableCapacity *int64 `json:"LargestAvailableCapacity,omitempty"`
	// External parity group number.
	Name *string `json:"Name,omitempty"`
	// From among the open volumes in the external parity group, the total capacity of volumes which are reserved (KB).
	ReservedOpenVolumeCapacity *int64                  `json:"ReservedOpenVolumeCapacity,omitempty"`
	Spaces                     []StorageSpace          `json:"Spaces,omitempty"`
	StorageUtilization         *StorageHitachiCapacity `json:"StorageUtilization,omitempty"`
	// Total volume capacity of the open volumes in the external parity group (KB).
	TotalOpenVolumeCapacity *int64 `json:"TotalOpenVolumeCapacity,omitempty"`
	// From among the open volumes in the external parity group, the total capacity of volumes to which paths are not allocated (KB).
	UnallocatedOpenVolumeCapacity *int64 `json:"UnallocatedOpenVolumeCapacity,omitempty"`
	// Usage rate of the external parity group.
	UsedCapacityRate     *int64                               `json:"UsedCapacityRate,omitempty"`
	Array                *StorageHitachiArrayRelationship     `json:"Array,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageHitachiExternalParityGroupAllOf StorageHitachiExternalParityGroupAllOf

// NewStorageHitachiExternalParityGroupAllOf instantiates a new StorageHitachiExternalParityGroupAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageHitachiExternalParityGroupAllOf(classId string, objectType string) *StorageHitachiExternalParityGroupAllOf {
	this := StorageHitachiExternalParityGroupAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageHitachiExternalParityGroupAllOfWithDefaults instantiates a new StorageHitachiExternalParityGroupAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageHitachiExternalParityGroupAllOfWithDefaults() *StorageHitachiExternalParityGroupAllOf {
	this := StorageHitachiExternalParityGroupAllOf{}
	var classId string = "storage.HitachiExternalParityGroup"
	this.ClassId = classId
	var objectType string = "storage.HitachiExternalParityGroup"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageHitachiExternalParityGroupAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageHitachiExternalParityGroupAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageHitachiExternalParityGroupAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageHitachiExternalParityGroupAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageHitachiExternalParityGroupAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageHitachiExternalParityGroupAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAllocatableOpenVolumeCapacity returns the AllocatableOpenVolumeCapacity field value if set, zero value otherwise.
func (o *StorageHitachiExternalParityGroupAllOf) GetAllocatableOpenVolumeCapacity() int64 {
	if o == nil || o.AllocatableOpenVolumeCapacity == nil {
		var ret int64
		return ret
	}
	return *o.AllocatableOpenVolumeCapacity
}

// GetAllocatableOpenVolumeCapacityOk returns a tuple with the AllocatableOpenVolumeCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiExternalParityGroupAllOf) GetAllocatableOpenVolumeCapacityOk() (*int64, bool) {
	if o == nil || o.AllocatableOpenVolumeCapacity == nil {
		return nil, false
	}
	return o.AllocatableOpenVolumeCapacity, true
}

// HasAllocatableOpenVolumeCapacity returns a boolean if a field has been set.
func (o *StorageHitachiExternalParityGroupAllOf) HasAllocatableOpenVolumeCapacity() bool {
	if o != nil && o.AllocatableOpenVolumeCapacity != nil {
		return true
	}

	return false
}

// SetAllocatableOpenVolumeCapacity gets a reference to the given int64 and assigns it to the AllocatableOpenVolumeCapacity field.
func (o *StorageHitachiExternalParityGroupAllOf) SetAllocatableOpenVolumeCapacity(v int64) {
	o.AllocatableOpenVolumeCapacity = &v
}

// GetAllocatedOpenVolumeCapacity returns the AllocatedOpenVolumeCapacity field value if set, zero value otherwise.
func (o *StorageHitachiExternalParityGroupAllOf) GetAllocatedOpenVolumeCapacity() int64 {
	if o == nil || o.AllocatedOpenVolumeCapacity == nil {
		var ret int64
		return ret
	}
	return *o.AllocatedOpenVolumeCapacity
}

// GetAllocatedOpenVolumeCapacityOk returns a tuple with the AllocatedOpenVolumeCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiExternalParityGroupAllOf) GetAllocatedOpenVolumeCapacityOk() (*int64, bool) {
	if o == nil || o.AllocatedOpenVolumeCapacity == nil {
		return nil, false
	}
	return o.AllocatedOpenVolumeCapacity, true
}

// HasAllocatedOpenVolumeCapacity returns a boolean if a field has been set.
func (o *StorageHitachiExternalParityGroupAllOf) HasAllocatedOpenVolumeCapacity() bool {
	if o != nil && o.AllocatedOpenVolumeCapacity != nil {
		return true
	}

	return false
}

// SetAllocatedOpenVolumeCapacity gets a reference to the given int64 and assigns it to the AllocatedOpenVolumeCapacity field.
func (o *StorageHitachiExternalParityGroupAllOf) SetAllocatedOpenVolumeCapacity(v int64) {
	o.AllocatedOpenVolumeCapacity = &v
}

// GetAvailableVolumeCapacity returns the AvailableVolumeCapacity field value if set, zero value otherwise.
func (o *StorageHitachiExternalParityGroupAllOf) GetAvailableVolumeCapacity() int64 {
	if o == nil || o.AvailableVolumeCapacity == nil {
		var ret int64
		return ret
	}
	return *o.AvailableVolumeCapacity
}

// GetAvailableVolumeCapacityOk returns a tuple with the AvailableVolumeCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiExternalParityGroupAllOf) GetAvailableVolumeCapacityOk() (*int64, bool) {
	if o == nil || o.AvailableVolumeCapacity == nil {
		return nil, false
	}
	return o.AvailableVolumeCapacity, true
}

// HasAvailableVolumeCapacity returns a boolean if a field has been set.
func (o *StorageHitachiExternalParityGroupAllOf) HasAvailableVolumeCapacity() bool {
	if o != nil && o.AvailableVolumeCapacity != nil {
		return true
	}

	return false
}

// SetAvailableVolumeCapacity gets a reference to the given int64 and assigns it to the AvailableVolumeCapacity field.
func (o *StorageHitachiExternalParityGroupAllOf) SetAvailableVolumeCapacity(v int64) {
	o.AvailableVolumeCapacity = &v
}

// GetClprId returns the ClprId field value if set, zero value otherwise.
func (o *StorageHitachiExternalParityGroupAllOf) GetClprId() int64 {
	if o == nil || o.ClprId == nil {
		var ret int64
		return ret
	}
	return *o.ClprId
}

// GetClprIdOk returns a tuple with the ClprId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiExternalParityGroupAllOf) GetClprIdOk() (*int64, bool) {
	if o == nil || o.ClprId == nil {
		return nil, false
	}
	return o.ClprId, true
}

// HasClprId returns a boolean if a field has been set.
func (o *StorageHitachiExternalParityGroupAllOf) HasClprId() bool {
	if o != nil && o.ClprId != nil {
		return true
	}

	return false
}

// SetClprId gets a reference to the given int64 and assigns it to the ClprId field.
func (o *StorageHitachiExternalParityGroupAllOf) SetClprId(v int64) {
	o.ClprId = &v
}

// GetEmulationType returns the EmulationType field value if set, zero value otherwise.
func (o *StorageHitachiExternalParityGroupAllOf) GetEmulationType() string {
	if o == nil || o.EmulationType == nil {
		var ret string
		return ret
	}
	return *o.EmulationType
}

// GetEmulationTypeOk returns a tuple with the EmulationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiExternalParityGroupAllOf) GetEmulationTypeOk() (*string, bool) {
	if o == nil || o.EmulationType == nil {
		return nil, false
	}
	return o.EmulationType, true
}

// HasEmulationType returns a boolean if a field has been set.
func (o *StorageHitachiExternalParityGroupAllOf) HasEmulationType() bool {
	if o != nil && o.EmulationType != nil {
		return true
	}

	return false
}

// SetEmulationType gets a reference to the given string and assigns it to the EmulationType field.
func (o *StorageHitachiExternalParityGroupAllOf) SetEmulationType(v string) {
	o.EmulationType = &v
}

// GetExternalProductId returns the ExternalProductId field value if set, zero value otherwise.
func (o *StorageHitachiExternalParityGroupAllOf) GetExternalProductId() string {
	if o == nil || o.ExternalProductId == nil {
		var ret string
		return ret
	}
	return *o.ExternalProductId
}

// GetExternalProductIdOk returns a tuple with the ExternalProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiExternalParityGroupAllOf) GetExternalProductIdOk() (*string, bool) {
	if o == nil || o.ExternalProductId == nil {
		return nil, false
	}
	return o.ExternalProductId, true
}

// HasExternalProductId returns a boolean if a field has been set.
func (o *StorageHitachiExternalParityGroupAllOf) HasExternalProductId() bool {
	if o != nil && o.ExternalProductId != nil {
		return true
	}

	return false
}

// SetExternalProductId gets a reference to the given string and assigns it to the ExternalProductId field.
func (o *StorageHitachiExternalParityGroupAllOf) SetExternalProductId(v string) {
	o.ExternalProductId = &v
}

// GetLargestAvailableCapacity returns the LargestAvailableCapacity field value if set, zero value otherwise.
func (o *StorageHitachiExternalParityGroupAllOf) GetLargestAvailableCapacity() int64 {
	if o == nil || o.LargestAvailableCapacity == nil {
		var ret int64
		return ret
	}
	return *o.LargestAvailableCapacity
}

// GetLargestAvailableCapacityOk returns a tuple with the LargestAvailableCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiExternalParityGroupAllOf) GetLargestAvailableCapacityOk() (*int64, bool) {
	if o == nil || o.LargestAvailableCapacity == nil {
		return nil, false
	}
	return o.LargestAvailableCapacity, true
}

// HasLargestAvailableCapacity returns a boolean if a field has been set.
func (o *StorageHitachiExternalParityGroupAllOf) HasLargestAvailableCapacity() bool {
	if o != nil && o.LargestAvailableCapacity != nil {
		return true
	}

	return false
}

// SetLargestAvailableCapacity gets a reference to the given int64 and assigns it to the LargestAvailableCapacity field.
func (o *StorageHitachiExternalParityGroupAllOf) SetLargestAvailableCapacity(v int64) {
	o.LargestAvailableCapacity = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StorageHitachiExternalParityGroupAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiExternalParityGroupAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StorageHitachiExternalParityGroupAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StorageHitachiExternalParityGroupAllOf) SetName(v string) {
	o.Name = &v
}

// GetReservedOpenVolumeCapacity returns the ReservedOpenVolumeCapacity field value if set, zero value otherwise.
func (o *StorageHitachiExternalParityGroupAllOf) GetReservedOpenVolumeCapacity() int64 {
	if o == nil || o.ReservedOpenVolumeCapacity == nil {
		var ret int64
		return ret
	}
	return *o.ReservedOpenVolumeCapacity
}

// GetReservedOpenVolumeCapacityOk returns a tuple with the ReservedOpenVolumeCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiExternalParityGroupAllOf) GetReservedOpenVolumeCapacityOk() (*int64, bool) {
	if o == nil || o.ReservedOpenVolumeCapacity == nil {
		return nil, false
	}
	return o.ReservedOpenVolumeCapacity, true
}

// HasReservedOpenVolumeCapacity returns a boolean if a field has been set.
func (o *StorageHitachiExternalParityGroupAllOf) HasReservedOpenVolumeCapacity() bool {
	if o != nil && o.ReservedOpenVolumeCapacity != nil {
		return true
	}

	return false
}

// SetReservedOpenVolumeCapacity gets a reference to the given int64 and assigns it to the ReservedOpenVolumeCapacity field.
func (o *StorageHitachiExternalParityGroupAllOf) SetReservedOpenVolumeCapacity(v int64) {
	o.ReservedOpenVolumeCapacity = &v
}

// GetSpaces returns the Spaces field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageHitachiExternalParityGroupAllOf) GetSpaces() []StorageSpace {
	if o == nil {
		var ret []StorageSpace
		return ret
	}
	return o.Spaces
}

// GetSpacesOk returns a tuple with the Spaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageHitachiExternalParityGroupAllOf) GetSpacesOk() ([]StorageSpace, bool) {
	if o == nil || o.Spaces == nil {
		return nil, false
	}
	return o.Spaces, true
}

// HasSpaces returns a boolean if a field has been set.
func (o *StorageHitachiExternalParityGroupAllOf) HasSpaces() bool {
	if o != nil && o.Spaces != nil {
		return true
	}

	return false
}

// SetSpaces gets a reference to the given []StorageSpace and assigns it to the Spaces field.
func (o *StorageHitachiExternalParityGroupAllOf) SetSpaces(v []StorageSpace) {
	o.Spaces = v
}

// GetStorageUtilization returns the StorageUtilization field value if set, zero value otherwise.
func (o *StorageHitachiExternalParityGroupAllOf) GetStorageUtilization() StorageHitachiCapacity {
	if o == nil || o.StorageUtilization == nil {
		var ret StorageHitachiCapacity
		return ret
	}
	return *o.StorageUtilization
}

// GetStorageUtilizationOk returns a tuple with the StorageUtilization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiExternalParityGroupAllOf) GetStorageUtilizationOk() (*StorageHitachiCapacity, bool) {
	if o == nil || o.StorageUtilization == nil {
		return nil, false
	}
	return o.StorageUtilization, true
}

// HasStorageUtilization returns a boolean if a field has been set.
func (o *StorageHitachiExternalParityGroupAllOf) HasStorageUtilization() bool {
	if o != nil && o.StorageUtilization != nil {
		return true
	}

	return false
}

// SetStorageUtilization gets a reference to the given StorageHitachiCapacity and assigns it to the StorageUtilization field.
func (o *StorageHitachiExternalParityGroupAllOf) SetStorageUtilization(v StorageHitachiCapacity) {
	o.StorageUtilization = &v
}

// GetTotalOpenVolumeCapacity returns the TotalOpenVolumeCapacity field value if set, zero value otherwise.
func (o *StorageHitachiExternalParityGroupAllOf) GetTotalOpenVolumeCapacity() int64 {
	if o == nil || o.TotalOpenVolumeCapacity == nil {
		var ret int64
		return ret
	}
	return *o.TotalOpenVolumeCapacity
}

// GetTotalOpenVolumeCapacityOk returns a tuple with the TotalOpenVolumeCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiExternalParityGroupAllOf) GetTotalOpenVolumeCapacityOk() (*int64, bool) {
	if o == nil || o.TotalOpenVolumeCapacity == nil {
		return nil, false
	}
	return o.TotalOpenVolumeCapacity, true
}

// HasTotalOpenVolumeCapacity returns a boolean if a field has been set.
func (o *StorageHitachiExternalParityGroupAllOf) HasTotalOpenVolumeCapacity() bool {
	if o != nil && o.TotalOpenVolumeCapacity != nil {
		return true
	}

	return false
}

// SetTotalOpenVolumeCapacity gets a reference to the given int64 and assigns it to the TotalOpenVolumeCapacity field.
func (o *StorageHitachiExternalParityGroupAllOf) SetTotalOpenVolumeCapacity(v int64) {
	o.TotalOpenVolumeCapacity = &v
}

// GetUnallocatedOpenVolumeCapacity returns the UnallocatedOpenVolumeCapacity field value if set, zero value otherwise.
func (o *StorageHitachiExternalParityGroupAllOf) GetUnallocatedOpenVolumeCapacity() int64 {
	if o == nil || o.UnallocatedOpenVolumeCapacity == nil {
		var ret int64
		return ret
	}
	return *o.UnallocatedOpenVolumeCapacity
}

// GetUnallocatedOpenVolumeCapacityOk returns a tuple with the UnallocatedOpenVolumeCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiExternalParityGroupAllOf) GetUnallocatedOpenVolumeCapacityOk() (*int64, bool) {
	if o == nil || o.UnallocatedOpenVolumeCapacity == nil {
		return nil, false
	}
	return o.UnallocatedOpenVolumeCapacity, true
}

// HasUnallocatedOpenVolumeCapacity returns a boolean if a field has been set.
func (o *StorageHitachiExternalParityGroupAllOf) HasUnallocatedOpenVolumeCapacity() bool {
	if o != nil && o.UnallocatedOpenVolumeCapacity != nil {
		return true
	}

	return false
}

// SetUnallocatedOpenVolumeCapacity gets a reference to the given int64 and assigns it to the UnallocatedOpenVolumeCapacity field.
func (o *StorageHitachiExternalParityGroupAllOf) SetUnallocatedOpenVolumeCapacity(v int64) {
	o.UnallocatedOpenVolumeCapacity = &v
}

// GetUsedCapacityRate returns the UsedCapacityRate field value if set, zero value otherwise.
func (o *StorageHitachiExternalParityGroupAllOf) GetUsedCapacityRate() int64 {
	if o == nil || o.UsedCapacityRate == nil {
		var ret int64
		return ret
	}
	return *o.UsedCapacityRate
}

// GetUsedCapacityRateOk returns a tuple with the UsedCapacityRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiExternalParityGroupAllOf) GetUsedCapacityRateOk() (*int64, bool) {
	if o == nil || o.UsedCapacityRate == nil {
		return nil, false
	}
	return o.UsedCapacityRate, true
}

// HasUsedCapacityRate returns a boolean if a field has been set.
func (o *StorageHitachiExternalParityGroupAllOf) HasUsedCapacityRate() bool {
	if o != nil && o.UsedCapacityRate != nil {
		return true
	}

	return false
}

// SetUsedCapacityRate gets a reference to the given int64 and assigns it to the UsedCapacityRate field.
func (o *StorageHitachiExternalParityGroupAllOf) SetUsedCapacityRate(v int64) {
	o.UsedCapacityRate = &v
}

// GetArray returns the Array field value if set, zero value otherwise.
func (o *StorageHitachiExternalParityGroupAllOf) GetArray() StorageHitachiArrayRelationship {
	if o == nil || o.Array == nil {
		var ret StorageHitachiArrayRelationship
		return ret
	}
	return *o.Array
}

// GetArrayOk returns a tuple with the Array field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiExternalParityGroupAllOf) GetArrayOk() (*StorageHitachiArrayRelationship, bool) {
	if o == nil || o.Array == nil {
		return nil, false
	}
	return o.Array, true
}

// HasArray returns a boolean if a field has been set.
func (o *StorageHitachiExternalParityGroupAllOf) HasArray() bool {
	if o != nil && o.Array != nil {
		return true
	}

	return false
}

// SetArray gets a reference to the given StorageHitachiArrayRelationship and assigns it to the Array field.
func (o *StorageHitachiExternalParityGroupAllOf) SetArray(v StorageHitachiArrayRelationship) {
	o.Array = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *StorageHitachiExternalParityGroupAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiExternalParityGroupAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StorageHitachiExternalParityGroupAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StorageHitachiExternalParityGroupAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o StorageHitachiExternalParityGroupAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AllocatableOpenVolumeCapacity != nil {
		toSerialize["AllocatableOpenVolumeCapacity"] = o.AllocatableOpenVolumeCapacity
	}
	if o.AllocatedOpenVolumeCapacity != nil {
		toSerialize["AllocatedOpenVolumeCapacity"] = o.AllocatedOpenVolumeCapacity
	}
	if o.AvailableVolumeCapacity != nil {
		toSerialize["AvailableVolumeCapacity"] = o.AvailableVolumeCapacity
	}
	if o.ClprId != nil {
		toSerialize["ClprId"] = o.ClprId
	}
	if o.EmulationType != nil {
		toSerialize["EmulationType"] = o.EmulationType
	}
	if o.ExternalProductId != nil {
		toSerialize["ExternalProductId"] = o.ExternalProductId
	}
	if o.LargestAvailableCapacity != nil {
		toSerialize["LargestAvailableCapacity"] = o.LargestAvailableCapacity
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.ReservedOpenVolumeCapacity != nil {
		toSerialize["ReservedOpenVolumeCapacity"] = o.ReservedOpenVolumeCapacity
	}
	if o.Spaces != nil {
		toSerialize["Spaces"] = o.Spaces
	}
	if o.StorageUtilization != nil {
		toSerialize["StorageUtilization"] = o.StorageUtilization
	}
	if o.TotalOpenVolumeCapacity != nil {
		toSerialize["TotalOpenVolumeCapacity"] = o.TotalOpenVolumeCapacity
	}
	if o.UnallocatedOpenVolumeCapacity != nil {
		toSerialize["UnallocatedOpenVolumeCapacity"] = o.UnallocatedOpenVolumeCapacity
	}
	if o.UsedCapacityRate != nil {
		toSerialize["UsedCapacityRate"] = o.UsedCapacityRate
	}
	if o.Array != nil {
		toSerialize["Array"] = o.Array
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageHitachiExternalParityGroupAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageHitachiExternalParityGroupAllOf := _StorageHitachiExternalParityGroupAllOf{}

	if err = json.Unmarshal(bytes, &varStorageHitachiExternalParityGroupAllOf); err == nil {
		*o = StorageHitachiExternalParityGroupAllOf(varStorageHitachiExternalParityGroupAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AllocatableOpenVolumeCapacity")
		delete(additionalProperties, "AllocatedOpenVolumeCapacity")
		delete(additionalProperties, "AvailableVolumeCapacity")
		delete(additionalProperties, "ClprId")
		delete(additionalProperties, "EmulationType")
		delete(additionalProperties, "ExternalProductId")
		delete(additionalProperties, "LargestAvailableCapacity")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "ReservedOpenVolumeCapacity")
		delete(additionalProperties, "Spaces")
		delete(additionalProperties, "StorageUtilization")
		delete(additionalProperties, "TotalOpenVolumeCapacity")
		delete(additionalProperties, "UnallocatedOpenVolumeCapacity")
		delete(additionalProperties, "UsedCapacityRate")
		delete(additionalProperties, "Array")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageHitachiExternalParityGroupAllOf struct {
	value *StorageHitachiExternalParityGroupAllOf
	isSet bool
}

func (v NullableStorageHitachiExternalParityGroupAllOf) Get() *StorageHitachiExternalParityGroupAllOf {
	return v.value
}

func (v *NullableStorageHitachiExternalParityGroupAllOf) Set(val *StorageHitachiExternalParityGroupAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageHitachiExternalParityGroupAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageHitachiExternalParityGroupAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageHitachiExternalParityGroupAllOf(val *StorageHitachiExternalParityGroupAllOf) *NullableStorageHitachiExternalParityGroupAllOf {
	return &NullableStorageHitachiExternalParityGroupAllOf{value: val, isSet: true}
}

func (v NullableStorageHitachiExternalParityGroupAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageHitachiExternalParityGroupAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
