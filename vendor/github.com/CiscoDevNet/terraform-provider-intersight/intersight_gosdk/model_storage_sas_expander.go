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
	"reflect"
	"strings"
)

// StorageSasExpander SAS Expander present in a server.
type StorageSasExpander struct {
	EquipmentBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Unique Identifier of the storage expander.
	ExpanderId *int64 `json:"ExpanderId,omitempty"`
	// The name  of the installed storage expander.
	Name *string `json:"Name,omitempty"`
	// The operational state of the storage expander.
	OperState *string `json:"OperState,omitempty"`
	// The operability status of the storage expander.
	Operability *string `json:"Operability,omitempty"`
	// The SAS address of the SAS expander.
	SasAddress           *string                              `json:"SasAddress,omitempty"`
	ComputeRackUnit      *ComputeRackUnitRelationship         `json:"ComputeRackUnit,omitempty"`
	Controller           *ManagementControllerRelationship    `json:"Controller,omitempty"`
	EquipmentChassis     *EquipmentChassisRelationship        `json:"EquipmentChassis,omitempty"`
	InventoryDeviceInfo  *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageSasExpander StorageSasExpander

// NewStorageSasExpander instantiates a new StorageSasExpander object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageSasExpander(classId string, objectType string) *StorageSasExpander {
	this := StorageSasExpander{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageSasExpanderWithDefaults instantiates a new StorageSasExpander object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageSasExpanderWithDefaults() *StorageSasExpander {
	this := StorageSasExpander{}
	var classId string = "storage.SasExpander"
	this.ClassId = classId
	var objectType string = "storage.SasExpander"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageSasExpander) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageSasExpander) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageSasExpander) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageSasExpander) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageSasExpander) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageSasExpander) SetObjectType(v string) {
	o.ObjectType = v
}

// GetExpanderId returns the ExpanderId field value if set, zero value otherwise.
func (o *StorageSasExpander) GetExpanderId() int64 {
	if o == nil || o.ExpanderId == nil {
		var ret int64
		return ret
	}
	return *o.ExpanderId
}

// GetExpanderIdOk returns a tuple with the ExpanderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageSasExpander) GetExpanderIdOk() (*int64, bool) {
	if o == nil || o.ExpanderId == nil {
		return nil, false
	}
	return o.ExpanderId, true
}

// HasExpanderId returns a boolean if a field has been set.
func (o *StorageSasExpander) HasExpanderId() bool {
	if o != nil && o.ExpanderId != nil {
		return true
	}

	return false
}

// SetExpanderId gets a reference to the given int64 and assigns it to the ExpanderId field.
func (o *StorageSasExpander) SetExpanderId(v int64) {
	o.ExpanderId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StorageSasExpander) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageSasExpander) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StorageSasExpander) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StorageSasExpander) SetName(v string) {
	o.Name = &v
}

// GetOperState returns the OperState field value if set, zero value otherwise.
func (o *StorageSasExpander) GetOperState() string {
	if o == nil || o.OperState == nil {
		var ret string
		return ret
	}
	return *o.OperState
}

// GetOperStateOk returns a tuple with the OperState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageSasExpander) GetOperStateOk() (*string, bool) {
	if o == nil || o.OperState == nil {
		return nil, false
	}
	return o.OperState, true
}

// HasOperState returns a boolean if a field has been set.
func (o *StorageSasExpander) HasOperState() bool {
	if o != nil && o.OperState != nil {
		return true
	}

	return false
}

// SetOperState gets a reference to the given string and assigns it to the OperState field.
func (o *StorageSasExpander) SetOperState(v string) {
	o.OperState = &v
}

// GetOperability returns the Operability field value if set, zero value otherwise.
func (o *StorageSasExpander) GetOperability() string {
	if o == nil || o.Operability == nil {
		var ret string
		return ret
	}
	return *o.Operability
}

// GetOperabilityOk returns a tuple with the Operability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageSasExpander) GetOperabilityOk() (*string, bool) {
	if o == nil || o.Operability == nil {
		return nil, false
	}
	return o.Operability, true
}

// HasOperability returns a boolean if a field has been set.
func (o *StorageSasExpander) HasOperability() bool {
	if o != nil && o.Operability != nil {
		return true
	}

	return false
}

// SetOperability gets a reference to the given string and assigns it to the Operability field.
func (o *StorageSasExpander) SetOperability(v string) {
	o.Operability = &v
}

// GetSasAddress returns the SasAddress field value if set, zero value otherwise.
func (o *StorageSasExpander) GetSasAddress() string {
	if o == nil || o.SasAddress == nil {
		var ret string
		return ret
	}
	return *o.SasAddress
}

// GetSasAddressOk returns a tuple with the SasAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageSasExpander) GetSasAddressOk() (*string, bool) {
	if o == nil || o.SasAddress == nil {
		return nil, false
	}
	return o.SasAddress, true
}

// HasSasAddress returns a boolean if a field has been set.
func (o *StorageSasExpander) HasSasAddress() bool {
	if o != nil && o.SasAddress != nil {
		return true
	}

	return false
}

// SetSasAddress gets a reference to the given string and assigns it to the SasAddress field.
func (o *StorageSasExpander) SetSasAddress(v string) {
	o.SasAddress = &v
}

// GetComputeRackUnit returns the ComputeRackUnit field value if set, zero value otherwise.
func (o *StorageSasExpander) GetComputeRackUnit() ComputeRackUnitRelationship {
	if o == nil || o.ComputeRackUnit == nil {
		var ret ComputeRackUnitRelationship
		return ret
	}
	return *o.ComputeRackUnit
}

// GetComputeRackUnitOk returns a tuple with the ComputeRackUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageSasExpander) GetComputeRackUnitOk() (*ComputeRackUnitRelationship, bool) {
	if o == nil || o.ComputeRackUnit == nil {
		return nil, false
	}
	return o.ComputeRackUnit, true
}

// HasComputeRackUnit returns a boolean if a field has been set.
func (o *StorageSasExpander) HasComputeRackUnit() bool {
	if o != nil && o.ComputeRackUnit != nil {
		return true
	}

	return false
}

// SetComputeRackUnit gets a reference to the given ComputeRackUnitRelationship and assigns it to the ComputeRackUnit field.
func (o *StorageSasExpander) SetComputeRackUnit(v ComputeRackUnitRelationship) {
	o.ComputeRackUnit = &v
}

// GetController returns the Controller field value if set, zero value otherwise.
func (o *StorageSasExpander) GetController() ManagementControllerRelationship {
	if o == nil || o.Controller == nil {
		var ret ManagementControllerRelationship
		return ret
	}
	return *o.Controller
}

// GetControllerOk returns a tuple with the Controller field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageSasExpander) GetControllerOk() (*ManagementControllerRelationship, bool) {
	if o == nil || o.Controller == nil {
		return nil, false
	}
	return o.Controller, true
}

// HasController returns a boolean if a field has been set.
func (o *StorageSasExpander) HasController() bool {
	if o != nil && o.Controller != nil {
		return true
	}

	return false
}

// SetController gets a reference to the given ManagementControllerRelationship and assigns it to the Controller field.
func (o *StorageSasExpander) SetController(v ManagementControllerRelationship) {
	o.Controller = &v
}

// GetEquipmentChassis returns the EquipmentChassis field value if set, zero value otherwise.
func (o *StorageSasExpander) GetEquipmentChassis() EquipmentChassisRelationship {
	if o == nil || o.EquipmentChassis == nil {
		var ret EquipmentChassisRelationship
		return ret
	}
	return *o.EquipmentChassis
}

// GetEquipmentChassisOk returns a tuple with the EquipmentChassis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageSasExpander) GetEquipmentChassisOk() (*EquipmentChassisRelationship, bool) {
	if o == nil || o.EquipmentChassis == nil {
		return nil, false
	}
	return o.EquipmentChassis, true
}

// HasEquipmentChassis returns a boolean if a field has been set.
func (o *StorageSasExpander) HasEquipmentChassis() bool {
	if o != nil && o.EquipmentChassis != nil {
		return true
	}

	return false
}

// SetEquipmentChassis gets a reference to the given EquipmentChassisRelationship and assigns it to the EquipmentChassis field.
func (o *StorageSasExpander) SetEquipmentChassis(v EquipmentChassisRelationship) {
	o.EquipmentChassis = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *StorageSasExpander) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageSasExpander) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *StorageSasExpander) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *StorageSasExpander) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *StorageSasExpander) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageSasExpander) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StorageSasExpander) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StorageSasExpander) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o StorageSasExpander) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedEquipmentBase, errEquipmentBase := json.Marshal(o.EquipmentBase)
	if errEquipmentBase != nil {
		return []byte{}, errEquipmentBase
	}
	errEquipmentBase = json.Unmarshal([]byte(serializedEquipmentBase), &toSerialize)
	if errEquipmentBase != nil {
		return []byte{}, errEquipmentBase
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ExpanderId != nil {
		toSerialize["ExpanderId"] = o.ExpanderId
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.OperState != nil {
		toSerialize["OperState"] = o.OperState
	}
	if o.Operability != nil {
		toSerialize["Operability"] = o.Operability
	}
	if o.SasAddress != nil {
		toSerialize["SasAddress"] = o.SasAddress
	}
	if o.ComputeRackUnit != nil {
		toSerialize["ComputeRackUnit"] = o.ComputeRackUnit
	}
	if o.Controller != nil {
		toSerialize["Controller"] = o.Controller
	}
	if o.EquipmentChassis != nil {
		toSerialize["EquipmentChassis"] = o.EquipmentChassis
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageSasExpander) UnmarshalJSON(bytes []byte) (err error) {
	type StorageSasExpanderWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Unique Identifier of the storage expander.
		ExpanderId *int64 `json:"ExpanderId,omitempty"`
		// The name  of the installed storage expander.
		Name *string `json:"Name,omitempty"`
		// The operational state of the storage expander.
		OperState *string `json:"OperState,omitempty"`
		// The operability status of the storage expander.
		Operability *string `json:"Operability,omitempty"`
		// The SAS address of the SAS expander.
		SasAddress          *string                              `json:"SasAddress,omitempty"`
		ComputeRackUnit     *ComputeRackUnitRelationship         `json:"ComputeRackUnit,omitempty"`
		Controller          *ManagementControllerRelationship    `json:"Controller,omitempty"`
		EquipmentChassis    *EquipmentChassisRelationship        `json:"EquipmentChassis,omitempty"`
		InventoryDeviceInfo *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
		RegisteredDevice    *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varStorageSasExpanderWithoutEmbeddedStruct := StorageSasExpanderWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageSasExpanderWithoutEmbeddedStruct)
	if err == nil {
		varStorageSasExpander := _StorageSasExpander{}
		varStorageSasExpander.ClassId = varStorageSasExpanderWithoutEmbeddedStruct.ClassId
		varStorageSasExpander.ObjectType = varStorageSasExpanderWithoutEmbeddedStruct.ObjectType
		varStorageSasExpander.ExpanderId = varStorageSasExpanderWithoutEmbeddedStruct.ExpanderId
		varStorageSasExpander.Name = varStorageSasExpanderWithoutEmbeddedStruct.Name
		varStorageSasExpander.OperState = varStorageSasExpanderWithoutEmbeddedStruct.OperState
		varStorageSasExpander.Operability = varStorageSasExpanderWithoutEmbeddedStruct.Operability
		varStorageSasExpander.SasAddress = varStorageSasExpanderWithoutEmbeddedStruct.SasAddress
		varStorageSasExpander.ComputeRackUnit = varStorageSasExpanderWithoutEmbeddedStruct.ComputeRackUnit
		varStorageSasExpander.Controller = varStorageSasExpanderWithoutEmbeddedStruct.Controller
		varStorageSasExpander.EquipmentChassis = varStorageSasExpanderWithoutEmbeddedStruct.EquipmentChassis
		varStorageSasExpander.InventoryDeviceInfo = varStorageSasExpanderWithoutEmbeddedStruct.InventoryDeviceInfo
		varStorageSasExpander.RegisteredDevice = varStorageSasExpanderWithoutEmbeddedStruct.RegisteredDevice
		*o = StorageSasExpander(varStorageSasExpander)
	} else {
		return err
	}

	varStorageSasExpander := _StorageSasExpander{}

	err = json.Unmarshal(bytes, &varStorageSasExpander)
	if err == nil {
		o.EquipmentBase = varStorageSasExpander.EquipmentBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ExpanderId")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "OperState")
		delete(additionalProperties, "Operability")
		delete(additionalProperties, "SasAddress")
		delete(additionalProperties, "ComputeRackUnit")
		delete(additionalProperties, "Controller")
		delete(additionalProperties, "EquipmentChassis")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "RegisteredDevice")

		// remove fields from embedded structs
		reflectEquipmentBase := reflect.ValueOf(o.EquipmentBase)
		for i := 0; i < reflectEquipmentBase.Type().NumField(); i++ {
			t := reflectEquipmentBase.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageSasExpander struct {
	value *StorageSasExpander
	isSet bool
}

func (v NullableStorageSasExpander) Get() *StorageSasExpander {
	return v.value
}

func (v *NullableStorageSasExpander) Set(val *StorageSasExpander) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageSasExpander) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageSasExpander) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageSasExpander(val *StorageSasExpander) *NullableStorageSasExpander {
	return &NullableStorageSasExpander{value: val, isSet: true}
}

func (v NullableStorageSasExpander) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageSasExpander) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
