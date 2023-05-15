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

// StorageFlexFlashControllerPropsAllOf Definition of the list of properties defined in 'storage.FlexFlashControllerProps', excluding properties defined in parent classes.
type StorageFlexFlashControllerPropsAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Manageable card on the flex flash controller.
	CardsManageable *string `json:"CardsManageable,omitempty"`
	// Mode configured on the flex flash controller.
	ConfiguredMode *string `json:"ConfiguredMode,omitempty"`
	// The current name of the flex flash controller.
	ControllerName *string `json:"ControllerName,omitempty"`
	// The current status of the flex flash controller.
	ControllerStatus *string `json:"ControllerStatus,omitempty"`
	// Firmware version of the flex flash controller.
	FwVersion *string `json:"FwVersion,omitempty"`
	// Internal state of the flex flash controller.
	InternalState *string `json:"InternalState,omitempty"`
	// Operating mode of flex flash controller.
	OperatingMode *string `json:"OperatingMode,omitempty"`
	// Number of connected physical drives to a specific Flex flash controller.
	PhysicalDriveCount *string `json:"PhysicalDriveCount,omitempty"`
	// Product name of the flex flash controller.
	ProductName *string `json:"ProductName,omitempty"`
	// Startup firmware version of the Flex flash controller.
	StartupFwVersion *string `json:"StartupFwVersion,omitempty"`
	// Number of virtual drives for a specific Flex flash controller.
	VirtualDriveCount          *string                                 `json:"VirtualDriveCount,omitempty"`
	InventoryDeviceInfo        *InventoryDeviceInfoRelationship        `json:"InventoryDeviceInfo,omitempty"`
	RegisteredDevice           *AssetDeviceRegistrationRelationship    `json:"RegisteredDevice,omitempty"`
	StorageFlexFlashController *StorageFlexFlashControllerRelationship `json:"StorageFlexFlashController,omitempty"`
	AdditionalProperties       map[string]interface{}
}

type _StorageFlexFlashControllerPropsAllOf StorageFlexFlashControllerPropsAllOf

// NewStorageFlexFlashControllerPropsAllOf instantiates a new StorageFlexFlashControllerPropsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageFlexFlashControllerPropsAllOf(classId string, objectType string) *StorageFlexFlashControllerPropsAllOf {
	this := StorageFlexFlashControllerPropsAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageFlexFlashControllerPropsAllOfWithDefaults instantiates a new StorageFlexFlashControllerPropsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageFlexFlashControllerPropsAllOfWithDefaults() *StorageFlexFlashControllerPropsAllOf {
	this := StorageFlexFlashControllerPropsAllOf{}
	var classId string = "storage.FlexFlashControllerProps"
	this.ClassId = classId
	var objectType string = "storage.FlexFlashControllerProps"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageFlexFlashControllerPropsAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashControllerPropsAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageFlexFlashControllerPropsAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageFlexFlashControllerPropsAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashControllerPropsAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageFlexFlashControllerPropsAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCardsManageable returns the CardsManageable field value if set, zero value otherwise.
func (o *StorageFlexFlashControllerPropsAllOf) GetCardsManageable() string {
	if o == nil || o.CardsManageable == nil {
		var ret string
		return ret
	}
	return *o.CardsManageable
}

// GetCardsManageableOk returns a tuple with the CardsManageable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashControllerPropsAllOf) GetCardsManageableOk() (*string, bool) {
	if o == nil || o.CardsManageable == nil {
		return nil, false
	}
	return o.CardsManageable, true
}

// HasCardsManageable returns a boolean if a field has been set.
func (o *StorageFlexFlashControllerPropsAllOf) HasCardsManageable() bool {
	if o != nil && o.CardsManageable != nil {
		return true
	}

	return false
}

// SetCardsManageable gets a reference to the given string and assigns it to the CardsManageable field.
func (o *StorageFlexFlashControllerPropsAllOf) SetCardsManageable(v string) {
	o.CardsManageable = &v
}

// GetConfiguredMode returns the ConfiguredMode field value if set, zero value otherwise.
func (o *StorageFlexFlashControllerPropsAllOf) GetConfiguredMode() string {
	if o == nil || o.ConfiguredMode == nil {
		var ret string
		return ret
	}
	return *o.ConfiguredMode
}

// GetConfiguredModeOk returns a tuple with the ConfiguredMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashControllerPropsAllOf) GetConfiguredModeOk() (*string, bool) {
	if o == nil || o.ConfiguredMode == nil {
		return nil, false
	}
	return o.ConfiguredMode, true
}

// HasConfiguredMode returns a boolean if a field has been set.
func (o *StorageFlexFlashControllerPropsAllOf) HasConfiguredMode() bool {
	if o != nil && o.ConfiguredMode != nil {
		return true
	}

	return false
}

// SetConfiguredMode gets a reference to the given string and assigns it to the ConfiguredMode field.
func (o *StorageFlexFlashControllerPropsAllOf) SetConfiguredMode(v string) {
	o.ConfiguredMode = &v
}

// GetControllerName returns the ControllerName field value if set, zero value otherwise.
func (o *StorageFlexFlashControllerPropsAllOf) GetControllerName() string {
	if o == nil || o.ControllerName == nil {
		var ret string
		return ret
	}
	return *o.ControllerName
}

// GetControllerNameOk returns a tuple with the ControllerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashControllerPropsAllOf) GetControllerNameOk() (*string, bool) {
	if o == nil || o.ControllerName == nil {
		return nil, false
	}
	return o.ControllerName, true
}

// HasControllerName returns a boolean if a field has been set.
func (o *StorageFlexFlashControllerPropsAllOf) HasControllerName() bool {
	if o != nil && o.ControllerName != nil {
		return true
	}

	return false
}

// SetControllerName gets a reference to the given string and assigns it to the ControllerName field.
func (o *StorageFlexFlashControllerPropsAllOf) SetControllerName(v string) {
	o.ControllerName = &v
}

// GetControllerStatus returns the ControllerStatus field value if set, zero value otherwise.
func (o *StorageFlexFlashControllerPropsAllOf) GetControllerStatus() string {
	if o == nil || o.ControllerStatus == nil {
		var ret string
		return ret
	}
	return *o.ControllerStatus
}

// GetControllerStatusOk returns a tuple with the ControllerStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashControllerPropsAllOf) GetControllerStatusOk() (*string, bool) {
	if o == nil || o.ControllerStatus == nil {
		return nil, false
	}
	return o.ControllerStatus, true
}

// HasControllerStatus returns a boolean if a field has been set.
func (o *StorageFlexFlashControllerPropsAllOf) HasControllerStatus() bool {
	if o != nil && o.ControllerStatus != nil {
		return true
	}

	return false
}

// SetControllerStatus gets a reference to the given string and assigns it to the ControllerStatus field.
func (o *StorageFlexFlashControllerPropsAllOf) SetControllerStatus(v string) {
	o.ControllerStatus = &v
}

// GetFwVersion returns the FwVersion field value if set, zero value otherwise.
func (o *StorageFlexFlashControllerPropsAllOf) GetFwVersion() string {
	if o == nil || o.FwVersion == nil {
		var ret string
		return ret
	}
	return *o.FwVersion
}

// GetFwVersionOk returns a tuple with the FwVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashControllerPropsAllOf) GetFwVersionOk() (*string, bool) {
	if o == nil || o.FwVersion == nil {
		return nil, false
	}
	return o.FwVersion, true
}

// HasFwVersion returns a boolean if a field has been set.
func (o *StorageFlexFlashControllerPropsAllOf) HasFwVersion() bool {
	if o != nil && o.FwVersion != nil {
		return true
	}

	return false
}

// SetFwVersion gets a reference to the given string and assigns it to the FwVersion field.
func (o *StorageFlexFlashControllerPropsAllOf) SetFwVersion(v string) {
	o.FwVersion = &v
}

// GetInternalState returns the InternalState field value if set, zero value otherwise.
func (o *StorageFlexFlashControllerPropsAllOf) GetInternalState() string {
	if o == nil || o.InternalState == nil {
		var ret string
		return ret
	}
	return *o.InternalState
}

// GetInternalStateOk returns a tuple with the InternalState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashControllerPropsAllOf) GetInternalStateOk() (*string, bool) {
	if o == nil || o.InternalState == nil {
		return nil, false
	}
	return o.InternalState, true
}

// HasInternalState returns a boolean if a field has been set.
func (o *StorageFlexFlashControllerPropsAllOf) HasInternalState() bool {
	if o != nil && o.InternalState != nil {
		return true
	}

	return false
}

// SetInternalState gets a reference to the given string and assigns it to the InternalState field.
func (o *StorageFlexFlashControllerPropsAllOf) SetInternalState(v string) {
	o.InternalState = &v
}

// GetOperatingMode returns the OperatingMode field value if set, zero value otherwise.
func (o *StorageFlexFlashControllerPropsAllOf) GetOperatingMode() string {
	if o == nil || o.OperatingMode == nil {
		var ret string
		return ret
	}
	return *o.OperatingMode
}

// GetOperatingModeOk returns a tuple with the OperatingMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashControllerPropsAllOf) GetOperatingModeOk() (*string, bool) {
	if o == nil || o.OperatingMode == nil {
		return nil, false
	}
	return o.OperatingMode, true
}

// HasOperatingMode returns a boolean if a field has been set.
func (o *StorageFlexFlashControllerPropsAllOf) HasOperatingMode() bool {
	if o != nil && o.OperatingMode != nil {
		return true
	}

	return false
}

// SetOperatingMode gets a reference to the given string and assigns it to the OperatingMode field.
func (o *StorageFlexFlashControllerPropsAllOf) SetOperatingMode(v string) {
	o.OperatingMode = &v
}

// GetPhysicalDriveCount returns the PhysicalDriveCount field value if set, zero value otherwise.
func (o *StorageFlexFlashControllerPropsAllOf) GetPhysicalDriveCount() string {
	if o == nil || o.PhysicalDriveCount == nil {
		var ret string
		return ret
	}
	return *o.PhysicalDriveCount
}

// GetPhysicalDriveCountOk returns a tuple with the PhysicalDriveCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashControllerPropsAllOf) GetPhysicalDriveCountOk() (*string, bool) {
	if o == nil || o.PhysicalDriveCount == nil {
		return nil, false
	}
	return o.PhysicalDriveCount, true
}

// HasPhysicalDriveCount returns a boolean if a field has been set.
func (o *StorageFlexFlashControllerPropsAllOf) HasPhysicalDriveCount() bool {
	if o != nil && o.PhysicalDriveCount != nil {
		return true
	}

	return false
}

// SetPhysicalDriveCount gets a reference to the given string and assigns it to the PhysicalDriveCount field.
func (o *StorageFlexFlashControllerPropsAllOf) SetPhysicalDriveCount(v string) {
	o.PhysicalDriveCount = &v
}

// GetProductName returns the ProductName field value if set, zero value otherwise.
func (o *StorageFlexFlashControllerPropsAllOf) GetProductName() string {
	if o == nil || o.ProductName == nil {
		var ret string
		return ret
	}
	return *o.ProductName
}

// GetProductNameOk returns a tuple with the ProductName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashControllerPropsAllOf) GetProductNameOk() (*string, bool) {
	if o == nil || o.ProductName == nil {
		return nil, false
	}
	return o.ProductName, true
}

// HasProductName returns a boolean if a field has been set.
func (o *StorageFlexFlashControllerPropsAllOf) HasProductName() bool {
	if o != nil && o.ProductName != nil {
		return true
	}

	return false
}

// SetProductName gets a reference to the given string and assigns it to the ProductName field.
func (o *StorageFlexFlashControllerPropsAllOf) SetProductName(v string) {
	o.ProductName = &v
}

// GetStartupFwVersion returns the StartupFwVersion field value if set, zero value otherwise.
func (o *StorageFlexFlashControllerPropsAllOf) GetStartupFwVersion() string {
	if o == nil || o.StartupFwVersion == nil {
		var ret string
		return ret
	}
	return *o.StartupFwVersion
}

// GetStartupFwVersionOk returns a tuple with the StartupFwVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashControllerPropsAllOf) GetStartupFwVersionOk() (*string, bool) {
	if o == nil || o.StartupFwVersion == nil {
		return nil, false
	}
	return o.StartupFwVersion, true
}

// HasStartupFwVersion returns a boolean if a field has been set.
func (o *StorageFlexFlashControllerPropsAllOf) HasStartupFwVersion() bool {
	if o != nil && o.StartupFwVersion != nil {
		return true
	}

	return false
}

// SetStartupFwVersion gets a reference to the given string and assigns it to the StartupFwVersion field.
func (o *StorageFlexFlashControllerPropsAllOf) SetStartupFwVersion(v string) {
	o.StartupFwVersion = &v
}

// GetVirtualDriveCount returns the VirtualDriveCount field value if set, zero value otherwise.
func (o *StorageFlexFlashControllerPropsAllOf) GetVirtualDriveCount() string {
	if o == nil || o.VirtualDriveCount == nil {
		var ret string
		return ret
	}
	return *o.VirtualDriveCount
}

// GetVirtualDriveCountOk returns a tuple with the VirtualDriveCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashControllerPropsAllOf) GetVirtualDriveCountOk() (*string, bool) {
	if o == nil || o.VirtualDriveCount == nil {
		return nil, false
	}
	return o.VirtualDriveCount, true
}

// HasVirtualDriveCount returns a boolean if a field has been set.
func (o *StorageFlexFlashControllerPropsAllOf) HasVirtualDriveCount() bool {
	if o != nil && o.VirtualDriveCount != nil {
		return true
	}

	return false
}

// SetVirtualDriveCount gets a reference to the given string and assigns it to the VirtualDriveCount field.
func (o *StorageFlexFlashControllerPropsAllOf) SetVirtualDriveCount(v string) {
	o.VirtualDriveCount = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *StorageFlexFlashControllerPropsAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashControllerPropsAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *StorageFlexFlashControllerPropsAllOf) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *StorageFlexFlashControllerPropsAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *StorageFlexFlashControllerPropsAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashControllerPropsAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StorageFlexFlashControllerPropsAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StorageFlexFlashControllerPropsAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

// GetStorageFlexFlashController returns the StorageFlexFlashController field value if set, zero value otherwise.
func (o *StorageFlexFlashControllerPropsAllOf) GetStorageFlexFlashController() StorageFlexFlashControllerRelationship {
	if o == nil || o.StorageFlexFlashController == nil {
		var ret StorageFlexFlashControllerRelationship
		return ret
	}
	return *o.StorageFlexFlashController
}

// GetStorageFlexFlashControllerOk returns a tuple with the StorageFlexFlashController field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashControllerPropsAllOf) GetStorageFlexFlashControllerOk() (*StorageFlexFlashControllerRelationship, bool) {
	if o == nil || o.StorageFlexFlashController == nil {
		return nil, false
	}
	return o.StorageFlexFlashController, true
}

// HasStorageFlexFlashController returns a boolean if a field has been set.
func (o *StorageFlexFlashControllerPropsAllOf) HasStorageFlexFlashController() bool {
	if o != nil && o.StorageFlexFlashController != nil {
		return true
	}

	return false
}

// SetStorageFlexFlashController gets a reference to the given StorageFlexFlashControllerRelationship and assigns it to the StorageFlexFlashController field.
func (o *StorageFlexFlashControllerPropsAllOf) SetStorageFlexFlashController(v StorageFlexFlashControllerRelationship) {
	o.StorageFlexFlashController = &v
}

func (o StorageFlexFlashControllerPropsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.CardsManageable != nil {
		toSerialize["CardsManageable"] = o.CardsManageable
	}
	if o.ConfiguredMode != nil {
		toSerialize["ConfiguredMode"] = o.ConfiguredMode
	}
	if o.ControllerName != nil {
		toSerialize["ControllerName"] = o.ControllerName
	}
	if o.ControllerStatus != nil {
		toSerialize["ControllerStatus"] = o.ControllerStatus
	}
	if o.FwVersion != nil {
		toSerialize["FwVersion"] = o.FwVersion
	}
	if o.InternalState != nil {
		toSerialize["InternalState"] = o.InternalState
	}
	if o.OperatingMode != nil {
		toSerialize["OperatingMode"] = o.OperatingMode
	}
	if o.PhysicalDriveCount != nil {
		toSerialize["PhysicalDriveCount"] = o.PhysicalDriveCount
	}
	if o.ProductName != nil {
		toSerialize["ProductName"] = o.ProductName
	}
	if o.StartupFwVersion != nil {
		toSerialize["StartupFwVersion"] = o.StartupFwVersion
	}
	if o.VirtualDriveCount != nil {
		toSerialize["VirtualDriveCount"] = o.VirtualDriveCount
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	if o.StorageFlexFlashController != nil {
		toSerialize["StorageFlexFlashController"] = o.StorageFlexFlashController
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageFlexFlashControllerPropsAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageFlexFlashControllerPropsAllOf := _StorageFlexFlashControllerPropsAllOf{}

	if err = json.Unmarshal(bytes, &varStorageFlexFlashControllerPropsAllOf); err == nil {
		*o = StorageFlexFlashControllerPropsAllOf(varStorageFlexFlashControllerPropsAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CardsManageable")
		delete(additionalProperties, "ConfiguredMode")
		delete(additionalProperties, "ControllerName")
		delete(additionalProperties, "ControllerStatus")
		delete(additionalProperties, "FwVersion")
		delete(additionalProperties, "InternalState")
		delete(additionalProperties, "OperatingMode")
		delete(additionalProperties, "PhysicalDriveCount")
		delete(additionalProperties, "ProductName")
		delete(additionalProperties, "StartupFwVersion")
		delete(additionalProperties, "VirtualDriveCount")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "RegisteredDevice")
		delete(additionalProperties, "StorageFlexFlashController")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageFlexFlashControllerPropsAllOf struct {
	value *StorageFlexFlashControllerPropsAllOf
	isSet bool
}

func (v NullableStorageFlexFlashControllerPropsAllOf) Get() *StorageFlexFlashControllerPropsAllOf {
	return v.value
}

func (v *NullableStorageFlexFlashControllerPropsAllOf) Set(val *StorageFlexFlashControllerPropsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageFlexFlashControllerPropsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageFlexFlashControllerPropsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageFlexFlashControllerPropsAllOf(val *StorageFlexFlashControllerPropsAllOf) *NullableStorageFlexFlashControllerPropsAllOf {
	return &NullableStorageFlexFlashControllerPropsAllOf{value: val, isSet: true}
}

func (v NullableStorageFlexFlashControllerPropsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageFlexFlashControllerPropsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
