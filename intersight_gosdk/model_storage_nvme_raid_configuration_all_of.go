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

// StorageNvmeRaidConfigurationAllOf Definition of the list of properties defined in 'storage.NvmeRaidConfiguration', excluding properties defined in parent classes.
type StorageNvmeRaidConfigurationAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The storage controller Dn Name for which Nvme RAID is created at endpoint.
	ControllerDn *string `json:"ControllerDn,omitempty"`
	// The storage controller Moid for which Nvme RAID creation is supported.
	ControllerMoid       *string                           `json:"ControllerMoid,omitempty"`
	DiskStates           []StorageNvmePhysicalDiskState    `json:"DiskStates,omitempty"`
	DriveGroups          []StorageNvmeRaidDriveGroup       `json:"DriveGroups,omitempty"`
	ServerProfile        *ServerProfileRelationship        `json:"ServerProfile,omitempty"`
	StoragePolicy        *StorageStoragePolicyRelationship `json:"StoragePolicy,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNvmeRaidConfigurationAllOf StorageNvmeRaidConfigurationAllOf

// NewStorageNvmeRaidConfigurationAllOf instantiates a new StorageNvmeRaidConfigurationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNvmeRaidConfigurationAllOf(classId string, objectType string) *StorageNvmeRaidConfigurationAllOf {
	this := StorageNvmeRaidConfigurationAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNvmeRaidConfigurationAllOfWithDefaults instantiates a new StorageNvmeRaidConfigurationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNvmeRaidConfigurationAllOfWithDefaults() *StorageNvmeRaidConfigurationAllOf {
	this := StorageNvmeRaidConfigurationAllOf{}
	var classId string = "storage.NvmeRaidConfiguration"
	this.ClassId = classId
	var objectType string = "storage.NvmeRaidConfiguration"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNvmeRaidConfigurationAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNvmeRaidConfigurationAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNvmeRaidConfigurationAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageNvmeRaidConfigurationAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNvmeRaidConfigurationAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNvmeRaidConfigurationAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetControllerDn returns the ControllerDn field value if set, zero value otherwise.
func (o *StorageNvmeRaidConfigurationAllOf) GetControllerDn() string {
	if o == nil || o.ControllerDn == nil {
		var ret string
		return ret
	}
	return *o.ControllerDn
}

// GetControllerDnOk returns a tuple with the ControllerDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNvmeRaidConfigurationAllOf) GetControllerDnOk() (*string, bool) {
	if o == nil || o.ControllerDn == nil {
		return nil, false
	}
	return o.ControllerDn, true
}

// HasControllerDn returns a boolean if a field has been set.
func (o *StorageNvmeRaidConfigurationAllOf) HasControllerDn() bool {
	if o != nil && o.ControllerDn != nil {
		return true
	}

	return false
}

// SetControllerDn gets a reference to the given string and assigns it to the ControllerDn field.
func (o *StorageNvmeRaidConfigurationAllOf) SetControllerDn(v string) {
	o.ControllerDn = &v
}

// GetControllerMoid returns the ControllerMoid field value if set, zero value otherwise.
func (o *StorageNvmeRaidConfigurationAllOf) GetControllerMoid() string {
	if o == nil || o.ControllerMoid == nil {
		var ret string
		return ret
	}
	return *o.ControllerMoid
}

// GetControllerMoidOk returns a tuple with the ControllerMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNvmeRaidConfigurationAllOf) GetControllerMoidOk() (*string, bool) {
	if o == nil || o.ControllerMoid == nil {
		return nil, false
	}
	return o.ControllerMoid, true
}

// HasControllerMoid returns a boolean if a field has been set.
func (o *StorageNvmeRaidConfigurationAllOf) HasControllerMoid() bool {
	if o != nil && o.ControllerMoid != nil {
		return true
	}

	return false
}

// SetControllerMoid gets a reference to the given string and assigns it to the ControllerMoid field.
func (o *StorageNvmeRaidConfigurationAllOf) SetControllerMoid(v string) {
	o.ControllerMoid = &v
}

// GetDiskStates returns the DiskStates field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNvmeRaidConfigurationAllOf) GetDiskStates() []StorageNvmePhysicalDiskState {
	if o == nil {
		var ret []StorageNvmePhysicalDiskState
		return ret
	}
	return o.DiskStates
}

// GetDiskStatesOk returns a tuple with the DiskStates field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNvmeRaidConfigurationAllOf) GetDiskStatesOk() ([]StorageNvmePhysicalDiskState, bool) {
	if o == nil || o.DiskStates == nil {
		return nil, false
	}
	return o.DiskStates, true
}

// HasDiskStates returns a boolean if a field has been set.
func (o *StorageNvmeRaidConfigurationAllOf) HasDiskStates() bool {
	if o != nil && o.DiskStates != nil {
		return true
	}

	return false
}

// SetDiskStates gets a reference to the given []StorageNvmePhysicalDiskState and assigns it to the DiskStates field.
func (o *StorageNvmeRaidConfigurationAllOf) SetDiskStates(v []StorageNvmePhysicalDiskState) {
	o.DiskStates = v
}

// GetDriveGroups returns the DriveGroups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNvmeRaidConfigurationAllOf) GetDriveGroups() []StorageNvmeRaidDriveGroup {
	if o == nil {
		var ret []StorageNvmeRaidDriveGroup
		return ret
	}
	return o.DriveGroups
}

// GetDriveGroupsOk returns a tuple with the DriveGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNvmeRaidConfigurationAllOf) GetDriveGroupsOk() ([]StorageNvmeRaidDriveGroup, bool) {
	if o == nil || o.DriveGroups == nil {
		return nil, false
	}
	return o.DriveGroups, true
}

// HasDriveGroups returns a boolean if a field has been set.
func (o *StorageNvmeRaidConfigurationAllOf) HasDriveGroups() bool {
	if o != nil && o.DriveGroups != nil {
		return true
	}

	return false
}

// SetDriveGroups gets a reference to the given []StorageNvmeRaidDriveGroup and assigns it to the DriveGroups field.
func (o *StorageNvmeRaidConfigurationAllOf) SetDriveGroups(v []StorageNvmeRaidDriveGroup) {
	o.DriveGroups = v
}

// GetServerProfile returns the ServerProfile field value if set, zero value otherwise.
func (o *StorageNvmeRaidConfigurationAllOf) GetServerProfile() ServerProfileRelationship {
	if o == nil || o.ServerProfile == nil {
		var ret ServerProfileRelationship
		return ret
	}
	return *o.ServerProfile
}

// GetServerProfileOk returns a tuple with the ServerProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNvmeRaidConfigurationAllOf) GetServerProfileOk() (*ServerProfileRelationship, bool) {
	if o == nil || o.ServerProfile == nil {
		return nil, false
	}
	return o.ServerProfile, true
}

// HasServerProfile returns a boolean if a field has been set.
func (o *StorageNvmeRaidConfigurationAllOf) HasServerProfile() bool {
	if o != nil && o.ServerProfile != nil {
		return true
	}

	return false
}

// SetServerProfile gets a reference to the given ServerProfileRelationship and assigns it to the ServerProfile field.
func (o *StorageNvmeRaidConfigurationAllOf) SetServerProfile(v ServerProfileRelationship) {
	o.ServerProfile = &v
}

// GetStoragePolicy returns the StoragePolicy field value if set, zero value otherwise.
func (o *StorageNvmeRaidConfigurationAllOf) GetStoragePolicy() StorageStoragePolicyRelationship {
	if o == nil || o.StoragePolicy == nil {
		var ret StorageStoragePolicyRelationship
		return ret
	}
	return *o.StoragePolicy
}

// GetStoragePolicyOk returns a tuple with the StoragePolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNvmeRaidConfigurationAllOf) GetStoragePolicyOk() (*StorageStoragePolicyRelationship, bool) {
	if o == nil || o.StoragePolicy == nil {
		return nil, false
	}
	return o.StoragePolicy, true
}

// HasStoragePolicy returns a boolean if a field has been set.
func (o *StorageNvmeRaidConfigurationAllOf) HasStoragePolicy() bool {
	if o != nil && o.StoragePolicy != nil {
		return true
	}

	return false
}

// SetStoragePolicy gets a reference to the given StorageStoragePolicyRelationship and assigns it to the StoragePolicy field.
func (o *StorageNvmeRaidConfigurationAllOf) SetStoragePolicy(v StorageStoragePolicyRelationship) {
	o.StoragePolicy = &v
}

func (o StorageNvmeRaidConfigurationAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ControllerDn != nil {
		toSerialize["ControllerDn"] = o.ControllerDn
	}
	if o.ControllerMoid != nil {
		toSerialize["ControllerMoid"] = o.ControllerMoid
	}
	if o.DiskStates != nil {
		toSerialize["DiskStates"] = o.DiskStates
	}
	if o.DriveGroups != nil {
		toSerialize["DriveGroups"] = o.DriveGroups
	}
	if o.ServerProfile != nil {
		toSerialize["ServerProfile"] = o.ServerProfile
	}
	if o.StoragePolicy != nil {
		toSerialize["StoragePolicy"] = o.StoragePolicy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageNvmeRaidConfigurationAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageNvmeRaidConfigurationAllOf := _StorageNvmeRaidConfigurationAllOf{}

	if err = json.Unmarshal(bytes, &varStorageNvmeRaidConfigurationAllOf); err == nil {
		*o = StorageNvmeRaidConfigurationAllOf(varStorageNvmeRaidConfigurationAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ControllerDn")
		delete(additionalProperties, "ControllerMoid")
		delete(additionalProperties, "DiskStates")
		delete(additionalProperties, "DriveGroups")
		delete(additionalProperties, "ServerProfile")
		delete(additionalProperties, "StoragePolicy")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageNvmeRaidConfigurationAllOf struct {
	value *StorageNvmeRaidConfigurationAllOf
	isSet bool
}

func (v NullableStorageNvmeRaidConfigurationAllOf) Get() *StorageNvmeRaidConfigurationAllOf {
	return v.value
}

func (v *NullableStorageNvmeRaidConfigurationAllOf) Set(val *StorageNvmeRaidConfigurationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNvmeRaidConfigurationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNvmeRaidConfigurationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNvmeRaidConfigurationAllOf(val *StorageNvmeRaidConfigurationAllOf) *NullableStorageNvmeRaidConfigurationAllOf {
	return &NullableStorageNvmeRaidConfigurationAllOf{value: val, isSet: true}
}

func (v NullableStorageNvmeRaidConfigurationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNvmeRaidConfigurationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
