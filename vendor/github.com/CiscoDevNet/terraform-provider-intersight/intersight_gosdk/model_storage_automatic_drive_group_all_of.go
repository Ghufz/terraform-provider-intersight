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

// StorageAutomaticDriveGroupAllOf Definition of the list of properties defined in 'storage.AutomaticDriveGroup', excluding properties defined in parent classes.
type StorageAutomaticDriveGroupAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Type of drive that should be used for this RAID group. * `Any` - Any type of drive can be used for virtual drive creation. * `HDD` - Hard disk drives should be used for virtual drive creation. * `SSD` - Solid state drives should be used for virtual drive creation.
	DriveType *string `json:"DriveType,omitempty"`
	// Number of drives within this span group. The minimum number of disks needed in a span group varies based on RAID level. RAID0 requires at least one disk. RAID1 and RAID10 requires at least 2 and in multiples of . RAID5 and RAID50 require at least 3 disks in a span group. RAID6 and RAID60 require atleast 4 disks in a span.
	DrivesPerSpan *int64 `json:"DrivesPerSpan,omitempty"`
	// Minimum size of the drive to be used for creating this RAID group.
	MinimumDriveSize *int64 `json:"MinimumDriveSize,omitempty"`
	// Number of dedicated hot spare disks for this RAID group. Allowed value is a comma or hyphen separated number range.
	NumDedicatedHotSpares *string `json:"NumDedicatedHotSpares,omitempty"`
	// Number of span groups to be created for this RAID group. Non-nested RAID levels have a single span.
	NumberOfSpans *int64 `json:"NumberOfSpans,omitempty"`
	// This flag enables the drive group to use all the remaining drives on the server.
	UseRemainingDrives   *bool `json:"UseRemainingDrives,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageAutomaticDriveGroupAllOf StorageAutomaticDriveGroupAllOf

// NewStorageAutomaticDriveGroupAllOf instantiates a new StorageAutomaticDriveGroupAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageAutomaticDriveGroupAllOf(classId string, objectType string) *StorageAutomaticDriveGroupAllOf {
	this := StorageAutomaticDriveGroupAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var driveType string = "Any"
	this.DriveType = &driveType
	var drivesPerSpan int64 = 1
	this.DrivesPerSpan = &drivesPerSpan
	var numberOfSpans int64 = 0
	this.NumberOfSpans = &numberOfSpans
	return &this
}

// NewStorageAutomaticDriveGroupAllOfWithDefaults instantiates a new StorageAutomaticDriveGroupAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageAutomaticDriveGroupAllOfWithDefaults() *StorageAutomaticDriveGroupAllOf {
	this := StorageAutomaticDriveGroupAllOf{}
	var classId string = "storage.AutomaticDriveGroup"
	this.ClassId = classId
	var objectType string = "storage.AutomaticDriveGroup"
	this.ObjectType = objectType
	var driveType string = "Any"
	this.DriveType = &driveType
	var drivesPerSpan int64 = 1
	this.DrivesPerSpan = &drivesPerSpan
	var numberOfSpans int64 = 0
	this.NumberOfSpans = &numberOfSpans
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageAutomaticDriveGroupAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageAutomaticDriveGroupAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageAutomaticDriveGroupAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageAutomaticDriveGroupAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageAutomaticDriveGroupAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageAutomaticDriveGroupAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDriveType returns the DriveType field value if set, zero value otherwise.
func (o *StorageAutomaticDriveGroupAllOf) GetDriveType() string {
	if o == nil || o.DriveType == nil {
		var ret string
		return ret
	}
	return *o.DriveType
}

// GetDriveTypeOk returns a tuple with the DriveType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageAutomaticDriveGroupAllOf) GetDriveTypeOk() (*string, bool) {
	if o == nil || o.DriveType == nil {
		return nil, false
	}
	return o.DriveType, true
}

// HasDriveType returns a boolean if a field has been set.
func (o *StorageAutomaticDriveGroupAllOf) HasDriveType() bool {
	if o != nil && o.DriveType != nil {
		return true
	}

	return false
}

// SetDriveType gets a reference to the given string and assigns it to the DriveType field.
func (o *StorageAutomaticDriveGroupAllOf) SetDriveType(v string) {
	o.DriveType = &v
}

// GetDrivesPerSpan returns the DrivesPerSpan field value if set, zero value otherwise.
func (o *StorageAutomaticDriveGroupAllOf) GetDrivesPerSpan() int64 {
	if o == nil || o.DrivesPerSpan == nil {
		var ret int64
		return ret
	}
	return *o.DrivesPerSpan
}

// GetDrivesPerSpanOk returns a tuple with the DrivesPerSpan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageAutomaticDriveGroupAllOf) GetDrivesPerSpanOk() (*int64, bool) {
	if o == nil || o.DrivesPerSpan == nil {
		return nil, false
	}
	return o.DrivesPerSpan, true
}

// HasDrivesPerSpan returns a boolean if a field has been set.
func (o *StorageAutomaticDriveGroupAllOf) HasDrivesPerSpan() bool {
	if o != nil && o.DrivesPerSpan != nil {
		return true
	}

	return false
}

// SetDrivesPerSpan gets a reference to the given int64 and assigns it to the DrivesPerSpan field.
func (o *StorageAutomaticDriveGroupAllOf) SetDrivesPerSpan(v int64) {
	o.DrivesPerSpan = &v
}

// GetMinimumDriveSize returns the MinimumDriveSize field value if set, zero value otherwise.
func (o *StorageAutomaticDriveGroupAllOf) GetMinimumDriveSize() int64 {
	if o == nil || o.MinimumDriveSize == nil {
		var ret int64
		return ret
	}
	return *o.MinimumDriveSize
}

// GetMinimumDriveSizeOk returns a tuple with the MinimumDriveSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageAutomaticDriveGroupAllOf) GetMinimumDriveSizeOk() (*int64, bool) {
	if o == nil || o.MinimumDriveSize == nil {
		return nil, false
	}
	return o.MinimumDriveSize, true
}

// HasMinimumDriveSize returns a boolean if a field has been set.
func (o *StorageAutomaticDriveGroupAllOf) HasMinimumDriveSize() bool {
	if o != nil && o.MinimumDriveSize != nil {
		return true
	}

	return false
}

// SetMinimumDriveSize gets a reference to the given int64 and assigns it to the MinimumDriveSize field.
func (o *StorageAutomaticDriveGroupAllOf) SetMinimumDriveSize(v int64) {
	o.MinimumDriveSize = &v
}

// GetNumDedicatedHotSpares returns the NumDedicatedHotSpares field value if set, zero value otherwise.
func (o *StorageAutomaticDriveGroupAllOf) GetNumDedicatedHotSpares() string {
	if o == nil || o.NumDedicatedHotSpares == nil {
		var ret string
		return ret
	}
	return *o.NumDedicatedHotSpares
}

// GetNumDedicatedHotSparesOk returns a tuple with the NumDedicatedHotSpares field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageAutomaticDriveGroupAllOf) GetNumDedicatedHotSparesOk() (*string, bool) {
	if o == nil || o.NumDedicatedHotSpares == nil {
		return nil, false
	}
	return o.NumDedicatedHotSpares, true
}

// HasNumDedicatedHotSpares returns a boolean if a field has been set.
func (o *StorageAutomaticDriveGroupAllOf) HasNumDedicatedHotSpares() bool {
	if o != nil && o.NumDedicatedHotSpares != nil {
		return true
	}

	return false
}

// SetNumDedicatedHotSpares gets a reference to the given string and assigns it to the NumDedicatedHotSpares field.
func (o *StorageAutomaticDriveGroupAllOf) SetNumDedicatedHotSpares(v string) {
	o.NumDedicatedHotSpares = &v
}

// GetNumberOfSpans returns the NumberOfSpans field value if set, zero value otherwise.
func (o *StorageAutomaticDriveGroupAllOf) GetNumberOfSpans() int64 {
	if o == nil || o.NumberOfSpans == nil {
		var ret int64
		return ret
	}
	return *o.NumberOfSpans
}

// GetNumberOfSpansOk returns a tuple with the NumberOfSpans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageAutomaticDriveGroupAllOf) GetNumberOfSpansOk() (*int64, bool) {
	if o == nil || o.NumberOfSpans == nil {
		return nil, false
	}
	return o.NumberOfSpans, true
}

// HasNumberOfSpans returns a boolean if a field has been set.
func (o *StorageAutomaticDriveGroupAllOf) HasNumberOfSpans() bool {
	if o != nil && o.NumberOfSpans != nil {
		return true
	}

	return false
}

// SetNumberOfSpans gets a reference to the given int64 and assigns it to the NumberOfSpans field.
func (o *StorageAutomaticDriveGroupAllOf) SetNumberOfSpans(v int64) {
	o.NumberOfSpans = &v
}

// GetUseRemainingDrives returns the UseRemainingDrives field value if set, zero value otherwise.
func (o *StorageAutomaticDriveGroupAllOf) GetUseRemainingDrives() bool {
	if o == nil || o.UseRemainingDrives == nil {
		var ret bool
		return ret
	}
	return *o.UseRemainingDrives
}

// GetUseRemainingDrivesOk returns a tuple with the UseRemainingDrives field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageAutomaticDriveGroupAllOf) GetUseRemainingDrivesOk() (*bool, bool) {
	if o == nil || o.UseRemainingDrives == nil {
		return nil, false
	}
	return o.UseRemainingDrives, true
}

// HasUseRemainingDrives returns a boolean if a field has been set.
func (o *StorageAutomaticDriveGroupAllOf) HasUseRemainingDrives() bool {
	if o != nil && o.UseRemainingDrives != nil {
		return true
	}

	return false
}

// SetUseRemainingDrives gets a reference to the given bool and assigns it to the UseRemainingDrives field.
func (o *StorageAutomaticDriveGroupAllOf) SetUseRemainingDrives(v bool) {
	o.UseRemainingDrives = &v
}

func (o StorageAutomaticDriveGroupAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DriveType != nil {
		toSerialize["DriveType"] = o.DriveType
	}
	if o.DrivesPerSpan != nil {
		toSerialize["DrivesPerSpan"] = o.DrivesPerSpan
	}
	if o.MinimumDriveSize != nil {
		toSerialize["MinimumDriveSize"] = o.MinimumDriveSize
	}
	if o.NumDedicatedHotSpares != nil {
		toSerialize["NumDedicatedHotSpares"] = o.NumDedicatedHotSpares
	}
	if o.NumberOfSpans != nil {
		toSerialize["NumberOfSpans"] = o.NumberOfSpans
	}
	if o.UseRemainingDrives != nil {
		toSerialize["UseRemainingDrives"] = o.UseRemainingDrives
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageAutomaticDriveGroupAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageAutomaticDriveGroupAllOf := _StorageAutomaticDriveGroupAllOf{}

	if err = json.Unmarshal(bytes, &varStorageAutomaticDriveGroupAllOf); err == nil {
		*o = StorageAutomaticDriveGroupAllOf(varStorageAutomaticDriveGroupAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DriveType")
		delete(additionalProperties, "DrivesPerSpan")
		delete(additionalProperties, "MinimumDriveSize")
		delete(additionalProperties, "NumDedicatedHotSpares")
		delete(additionalProperties, "NumberOfSpans")
		delete(additionalProperties, "UseRemainingDrives")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageAutomaticDriveGroupAllOf struct {
	value *StorageAutomaticDriveGroupAllOf
	isSet bool
}

func (v NullableStorageAutomaticDriveGroupAllOf) Get() *StorageAutomaticDriveGroupAllOf {
	return v.value
}

func (v *NullableStorageAutomaticDriveGroupAllOf) Set(val *StorageAutomaticDriveGroupAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageAutomaticDriveGroupAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageAutomaticDriveGroupAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageAutomaticDriveGroupAllOf(val *StorageAutomaticDriveGroupAllOf) *NullableStorageAutomaticDriveGroupAllOf {
	return &NullableStorageAutomaticDriveGroupAllOf{value: val, isSet: true}
}

func (v NullableStorageAutomaticDriveGroupAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageAutomaticDriveGroupAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
