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

// BiosUnitAllOf Definition of the list of properties defined in 'bios.Unit', excluding properties defined in parent classes.
type BiosUnitAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The initSeq of the equipment.
	InitSeq *string `json:"InitSeq,omitempty"`
	// The initTs of the equipment.
	InitTs              *string                              `json:"InitTs,omitempty"`
	ComputeBlade        *ComputeBladeRelationship            `json:"ComputeBlade,omitempty"`
	ComputeRackUnit     *ComputeRackUnitRelationship         `json:"ComputeRackUnit,omitempty"`
	InventoryDeviceInfo *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
	RegisteredDevice    *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	// An array of relationships to firmwareRunningFirmware resources.
	RunningFirmware      []FirmwareRunningFirmwareRelationship `json:"RunningFirmware,omitempty"`
	SystemBootOrder      *BiosSystemBootOrderRelationship      `json:"SystemBootOrder,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BiosUnitAllOf BiosUnitAllOf

// NewBiosUnitAllOf instantiates a new BiosUnitAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBiosUnitAllOf(classId string, objectType string) *BiosUnitAllOf {
	this := BiosUnitAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewBiosUnitAllOfWithDefaults instantiates a new BiosUnitAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBiosUnitAllOfWithDefaults() *BiosUnitAllOf {
	this := BiosUnitAllOf{}
	var classId string = "bios.Unit"
	this.ClassId = classId
	var objectType string = "bios.Unit"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *BiosUnitAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *BiosUnitAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *BiosUnitAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *BiosUnitAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *BiosUnitAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *BiosUnitAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetInitSeq returns the InitSeq field value if set, zero value otherwise.
func (o *BiosUnitAllOf) GetInitSeq() string {
	if o == nil || o.InitSeq == nil {
		var ret string
		return ret
	}
	return *o.InitSeq
}

// GetInitSeqOk returns a tuple with the InitSeq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BiosUnitAllOf) GetInitSeqOk() (*string, bool) {
	if o == nil || o.InitSeq == nil {
		return nil, false
	}
	return o.InitSeq, true
}

// HasInitSeq returns a boolean if a field has been set.
func (o *BiosUnitAllOf) HasInitSeq() bool {
	if o != nil && o.InitSeq != nil {
		return true
	}

	return false
}

// SetInitSeq gets a reference to the given string and assigns it to the InitSeq field.
func (o *BiosUnitAllOf) SetInitSeq(v string) {
	o.InitSeq = &v
}

// GetInitTs returns the InitTs field value if set, zero value otherwise.
func (o *BiosUnitAllOf) GetInitTs() string {
	if o == nil || o.InitTs == nil {
		var ret string
		return ret
	}
	return *o.InitTs
}

// GetInitTsOk returns a tuple with the InitTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BiosUnitAllOf) GetInitTsOk() (*string, bool) {
	if o == nil || o.InitTs == nil {
		return nil, false
	}
	return o.InitTs, true
}

// HasInitTs returns a boolean if a field has been set.
func (o *BiosUnitAllOf) HasInitTs() bool {
	if o != nil && o.InitTs != nil {
		return true
	}

	return false
}

// SetInitTs gets a reference to the given string and assigns it to the InitTs field.
func (o *BiosUnitAllOf) SetInitTs(v string) {
	o.InitTs = &v
}

// GetComputeBlade returns the ComputeBlade field value if set, zero value otherwise.
func (o *BiosUnitAllOf) GetComputeBlade() ComputeBladeRelationship {
	if o == nil || o.ComputeBlade == nil {
		var ret ComputeBladeRelationship
		return ret
	}
	return *o.ComputeBlade
}

// GetComputeBladeOk returns a tuple with the ComputeBlade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BiosUnitAllOf) GetComputeBladeOk() (*ComputeBladeRelationship, bool) {
	if o == nil || o.ComputeBlade == nil {
		return nil, false
	}
	return o.ComputeBlade, true
}

// HasComputeBlade returns a boolean if a field has been set.
func (o *BiosUnitAllOf) HasComputeBlade() bool {
	if o != nil && o.ComputeBlade != nil {
		return true
	}

	return false
}

// SetComputeBlade gets a reference to the given ComputeBladeRelationship and assigns it to the ComputeBlade field.
func (o *BiosUnitAllOf) SetComputeBlade(v ComputeBladeRelationship) {
	o.ComputeBlade = &v
}

// GetComputeRackUnit returns the ComputeRackUnit field value if set, zero value otherwise.
func (o *BiosUnitAllOf) GetComputeRackUnit() ComputeRackUnitRelationship {
	if o == nil || o.ComputeRackUnit == nil {
		var ret ComputeRackUnitRelationship
		return ret
	}
	return *o.ComputeRackUnit
}

// GetComputeRackUnitOk returns a tuple with the ComputeRackUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BiosUnitAllOf) GetComputeRackUnitOk() (*ComputeRackUnitRelationship, bool) {
	if o == nil || o.ComputeRackUnit == nil {
		return nil, false
	}
	return o.ComputeRackUnit, true
}

// HasComputeRackUnit returns a boolean if a field has been set.
func (o *BiosUnitAllOf) HasComputeRackUnit() bool {
	if o != nil && o.ComputeRackUnit != nil {
		return true
	}

	return false
}

// SetComputeRackUnit gets a reference to the given ComputeRackUnitRelationship and assigns it to the ComputeRackUnit field.
func (o *BiosUnitAllOf) SetComputeRackUnit(v ComputeRackUnitRelationship) {
	o.ComputeRackUnit = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *BiosUnitAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BiosUnitAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *BiosUnitAllOf) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *BiosUnitAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *BiosUnitAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BiosUnitAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *BiosUnitAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *BiosUnitAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

// GetRunningFirmware returns the RunningFirmware field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BiosUnitAllOf) GetRunningFirmware() []FirmwareRunningFirmwareRelationship {
	if o == nil {
		var ret []FirmwareRunningFirmwareRelationship
		return ret
	}
	return o.RunningFirmware
}

// GetRunningFirmwareOk returns a tuple with the RunningFirmware field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BiosUnitAllOf) GetRunningFirmwareOk() ([]FirmwareRunningFirmwareRelationship, bool) {
	if o == nil || o.RunningFirmware == nil {
		return nil, false
	}
	return o.RunningFirmware, true
}

// HasRunningFirmware returns a boolean if a field has been set.
func (o *BiosUnitAllOf) HasRunningFirmware() bool {
	if o != nil && o.RunningFirmware != nil {
		return true
	}

	return false
}

// SetRunningFirmware gets a reference to the given []FirmwareRunningFirmwareRelationship and assigns it to the RunningFirmware field.
func (o *BiosUnitAllOf) SetRunningFirmware(v []FirmwareRunningFirmwareRelationship) {
	o.RunningFirmware = v
}

// GetSystemBootOrder returns the SystemBootOrder field value if set, zero value otherwise.
func (o *BiosUnitAllOf) GetSystemBootOrder() BiosSystemBootOrderRelationship {
	if o == nil || o.SystemBootOrder == nil {
		var ret BiosSystemBootOrderRelationship
		return ret
	}
	return *o.SystemBootOrder
}

// GetSystemBootOrderOk returns a tuple with the SystemBootOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BiosUnitAllOf) GetSystemBootOrderOk() (*BiosSystemBootOrderRelationship, bool) {
	if o == nil || o.SystemBootOrder == nil {
		return nil, false
	}
	return o.SystemBootOrder, true
}

// HasSystemBootOrder returns a boolean if a field has been set.
func (o *BiosUnitAllOf) HasSystemBootOrder() bool {
	if o != nil && o.SystemBootOrder != nil {
		return true
	}

	return false
}

// SetSystemBootOrder gets a reference to the given BiosSystemBootOrderRelationship and assigns it to the SystemBootOrder field.
func (o *BiosUnitAllOf) SetSystemBootOrder(v BiosSystemBootOrderRelationship) {
	o.SystemBootOrder = &v
}

func (o BiosUnitAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.InitSeq != nil {
		toSerialize["InitSeq"] = o.InitSeq
	}
	if o.InitTs != nil {
		toSerialize["InitTs"] = o.InitTs
	}
	if o.ComputeBlade != nil {
		toSerialize["ComputeBlade"] = o.ComputeBlade
	}
	if o.ComputeRackUnit != nil {
		toSerialize["ComputeRackUnit"] = o.ComputeRackUnit
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	if o.RunningFirmware != nil {
		toSerialize["RunningFirmware"] = o.RunningFirmware
	}
	if o.SystemBootOrder != nil {
		toSerialize["SystemBootOrder"] = o.SystemBootOrder
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BiosUnitAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varBiosUnitAllOf := _BiosUnitAllOf{}

	if err = json.Unmarshal(bytes, &varBiosUnitAllOf); err == nil {
		*o = BiosUnitAllOf(varBiosUnitAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "InitSeq")
		delete(additionalProperties, "InitTs")
		delete(additionalProperties, "ComputeBlade")
		delete(additionalProperties, "ComputeRackUnit")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "RegisteredDevice")
		delete(additionalProperties, "RunningFirmware")
		delete(additionalProperties, "SystemBootOrder")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBiosUnitAllOf struct {
	value *BiosUnitAllOf
	isSet bool
}

func (v NullableBiosUnitAllOf) Get() *BiosUnitAllOf {
	return v.value
}

func (v *NullableBiosUnitAllOf) Set(val *BiosUnitAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBiosUnitAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBiosUnitAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBiosUnitAllOf(val *BiosUnitAllOf) *NullableBiosUnitAllOf {
	return &NullableBiosUnitAllOf{value: val, isSet: true}
}

func (v NullableBiosUnitAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBiosUnitAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
