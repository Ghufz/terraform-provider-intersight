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
	"reflect"
	"strings"
)

// NetworkSupervisorCard Concrete class for supervisor card.
type NetworkSupervisorCard struct {
	EquipmentBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The description of the supervisor card.
	Description *string `json:"Description,omitempty"`
	// The hardware version of the supervisor card.
	HardwareVersion *string `json:"HardwareVersion,omitempty"`
	// The name of the supervisor card like Supervisor Card-1.
	Name *string `json:"Name,omitempty"`
	// The number of ports on the supervisor card.
	NumberOfPorts *int64 `json:"NumberOfPorts,omitempty"`
	// The health issue of the supervisor card.
	OperReason *string `json:"OperReason,omitempty"`
	// The operational status of the supervisor card.
	OperState *string `json:"OperState,omitempty"`
	// The part number of the supervisor card.
	PartNumber *string `json:"PartNumber,omitempty"`
	// The power state of the supervisor card.
	PowerState *string `json:"PowerState,omitempty"`
	// The status of the supervisor card.
	Status *string `json:"Status,omitempty"`
	// The identifier for the supervisor card.
	SupervisorId *string `json:"SupervisorId,omitempty"`
	// The type of the supervisor card.
	Type *string `json:"Type,omitempty"`
	// An array of relationships to fcPhysicalPort resources.
	FcPorts              []FcPhysicalPortRelationship         `json:"FcPorts,omitempty"`
	NetworkElement       *NetworkElementRelationship          `json:"NetworkElement,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NetworkSupervisorCard NetworkSupervisorCard

// NewNetworkSupervisorCard instantiates a new NetworkSupervisorCard object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkSupervisorCard(classId string, objectType string) *NetworkSupervisorCard {
	this := NetworkSupervisorCard{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNetworkSupervisorCardWithDefaults instantiates a new NetworkSupervisorCard object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkSupervisorCardWithDefaults() *NetworkSupervisorCard {
	this := NetworkSupervisorCard{}
	var classId string = "network.SupervisorCard"
	this.ClassId = classId
	var objectType string = "network.SupervisorCard"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NetworkSupervisorCard) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NetworkSupervisorCard) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NetworkSupervisorCard) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NetworkSupervisorCard) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NetworkSupervisorCard) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NetworkSupervisorCard) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *NetworkSupervisorCard) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkSupervisorCard) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *NetworkSupervisorCard) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *NetworkSupervisorCard) SetDescription(v string) {
	o.Description = &v
}

// GetHardwareVersion returns the HardwareVersion field value if set, zero value otherwise.
func (o *NetworkSupervisorCard) GetHardwareVersion() string {
	if o == nil || o.HardwareVersion == nil {
		var ret string
		return ret
	}
	return *o.HardwareVersion
}

// GetHardwareVersionOk returns a tuple with the HardwareVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkSupervisorCard) GetHardwareVersionOk() (*string, bool) {
	if o == nil || o.HardwareVersion == nil {
		return nil, false
	}
	return o.HardwareVersion, true
}

// HasHardwareVersion returns a boolean if a field has been set.
func (o *NetworkSupervisorCard) HasHardwareVersion() bool {
	if o != nil && o.HardwareVersion != nil {
		return true
	}

	return false
}

// SetHardwareVersion gets a reference to the given string and assigns it to the HardwareVersion field.
func (o *NetworkSupervisorCard) SetHardwareVersion(v string) {
	o.HardwareVersion = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NetworkSupervisorCard) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkSupervisorCard) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NetworkSupervisorCard) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NetworkSupervisorCard) SetName(v string) {
	o.Name = &v
}

// GetNumberOfPorts returns the NumberOfPorts field value if set, zero value otherwise.
func (o *NetworkSupervisorCard) GetNumberOfPorts() int64 {
	if o == nil || o.NumberOfPorts == nil {
		var ret int64
		return ret
	}
	return *o.NumberOfPorts
}

// GetNumberOfPortsOk returns a tuple with the NumberOfPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkSupervisorCard) GetNumberOfPortsOk() (*int64, bool) {
	if o == nil || o.NumberOfPorts == nil {
		return nil, false
	}
	return o.NumberOfPorts, true
}

// HasNumberOfPorts returns a boolean if a field has been set.
func (o *NetworkSupervisorCard) HasNumberOfPorts() bool {
	if o != nil && o.NumberOfPorts != nil {
		return true
	}

	return false
}

// SetNumberOfPorts gets a reference to the given int64 and assigns it to the NumberOfPorts field.
func (o *NetworkSupervisorCard) SetNumberOfPorts(v int64) {
	o.NumberOfPorts = &v
}

// GetOperReason returns the OperReason field value if set, zero value otherwise.
func (o *NetworkSupervisorCard) GetOperReason() string {
	if o == nil || o.OperReason == nil {
		var ret string
		return ret
	}
	return *o.OperReason
}

// GetOperReasonOk returns a tuple with the OperReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkSupervisorCard) GetOperReasonOk() (*string, bool) {
	if o == nil || o.OperReason == nil {
		return nil, false
	}
	return o.OperReason, true
}

// HasOperReason returns a boolean if a field has been set.
func (o *NetworkSupervisorCard) HasOperReason() bool {
	if o != nil && o.OperReason != nil {
		return true
	}

	return false
}

// SetOperReason gets a reference to the given string and assigns it to the OperReason field.
func (o *NetworkSupervisorCard) SetOperReason(v string) {
	o.OperReason = &v
}

// GetOperState returns the OperState field value if set, zero value otherwise.
func (o *NetworkSupervisorCard) GetOperState() string {
	if o == nil || o.OperState == nil {
		var ret string
		return ret
	}
	return *o.OperState
}

// GetOperStateOk returns a tuple with the OperState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkSupervisorCard) GetOperStateOk() (*string, bool) {
	if o == nil || o.OperState == nil {
		return nil, false
	}
	return o.OperState, true
}

// HasOperState returns a boolean if a field has been set.
func (o *NetworkSupervisorCard) HasOperState() bool {
	if o != nil && o.OperState != nil {
		return true
	}

	return false
}

// SetOperState gets a reference to the given string and assigns it to the OperState field.
func (o *NetworkSupervisorCard) SetOperState(v string) {
	o.OperState = &v
}

// GetPartNumber returns the PartNumber field value if set, zero value otherwise.
func (o *NetworkSupervisorCard) GetPartNumber() string {
	if o == nil || o.PartNumber == nil {
		var ret string
		return ret
	}
	return *o.PartNumber
}

// GetPartNumberOk returns a tuple with the PartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkSupervisorCard) GetPartNumberOk() (*string, bool) {
	if o == nil || o.PartNumber == nil {
		return nil, false
	}
	return o.PartNumber, true
}

// HasPartNumber returns a boolean if a field has been set.
func (o *NetworkSupervisorCard) HasPartNumber() bool {
	if o != nil && o.PartNumber != nil {
		return true
	}

	return false
}

// SetPartNumber gets a reference to the given string and assigns it to the PartNumber field.
func (o *NetworkSupervisorCard) SetPartNumber(v string) {
	o.PartNumber = &v
}

// GetPowerState returns the PowerState field value if set, zero value otherwise.
func (o *NetworkSupervisorCard) GetPowerState() string {
	if o == nil || o.PowerState == nil {
		var ret string
		return ret
	}
	return *o.PowerState
}

// GetPowerStateOk returns a tuple with the PowerState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkSupervisorCard) GetPowerStateOk() (*string, bool) {
	if o == nil || o.PowerState == nil {
		return nil, false
	}
	return o.PowerState, true
}

// HasPowerState returns a boolean if a field has been set.
func (o *NetworkSupervisorCard) HasPowerState() bool {
	if o != nil && o.PowerState != nil {
		return true
	}

	return false
}

// SetPowerState gets a reference to the given string and assigns it to the PowerState field.
func (o *NetworkSupervisorCard) SetPowerState(v string) {
	o.PowerState = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *NetworkSupervisorCard) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkSupervisorCard) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *NetworkSupervisorCard) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *NetworkSupervisorCard) SetStatus(v string) {
	o.Status = &v
}

// GetSupervisorId returns the SupervisorId field value if set, zero value otherwise.
func (o *NetworkSupervisorCard) GetSupervisorId() string {
	if o == nil || o.SupervisorId == nil {
		var ret string
		return ret
	}
	return *o.SupervisorId
}

// GetSupervisorIdOk returns a tuple with the SupervisorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkSupervisorCard) GetSupervisorIdOk() (*string, bool) {
	if o == nil || o.SupervisorId == nil {
		return nil, false
	}
	return o.SupervisorId, true
}

// HasSupervisorId returns a boolean if a field has been set.
func (o *NetworkSupervisorCard) HasSupervisorId() bool {
	if o != nil && o.SupervisorId != nil {
		return true
	}

	return false
}

// SetSupervisorId gets a reference to the given string and assigns it to the SupervisorId field.
func (o *NetworkSupervisorCard) SetSupervisorId(v string) {
	o.SupervisorId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *NetworkSupervisorCard) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkSupervisorCard) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *NetworkSupervisorCard) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *NetworkSupervisorCard) SetType(v string) {
	o.Type = &v
}

// GetFcPorts returns the FcPorts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetworkSupervisorCard) GetFcPorts() []FcPhysicalPortRelationship {
	if o == nil {
		var ret []FcPhysicalPortRelationship
		return ret
	}
	return o.FcPorts
}

// GetFcPortsOk returns a tuple with the FcPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkSupervisorCard) GetFcPortsOk() ([]FcPhysicalPortRelationship, bool) {
	if o == nil || o.FcPorts == nil {
		return nil, false
	}
	return o.FcPorts, true
}

// HasFcPorts returns a boolean if a field has been set.
func (o *NetworkSupervisorCard) HasFcPorts() bool {
	if o != nil && o.FcPorts != nil {
		return true
	}

	return false
}

// SetFcPorts gets a reference to the given []FcPhysicalPortRelationship and assigns it to the FcPorts field.
func (o *NetworkSupervisorCard) SetFcPorts(v []FcPhysicalPortRelationship) {
	o.FcPorts = v
}

// GetNetworkElement returns the NetworkElement field value if set, zero value otherwise.
func (o *NetworkSupervisorCard) GetNetworkElement() NetworkElementRelationship {
	if o == nil || o.NetworkElement == nil {
		var ret NetworkElementRelationship
		return ret
	}
	return *o.NetworkElement
}

// GetNetworkElementOk returns a tuple with the NetworkElement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkSupervisorCard) GetNetworkElementOk() (*NetworkElementRelationship, bool) {
	if o == nil || o.NetworkElement == nil {
		return nil, false
	}
	return o.NetworkElement, true
}

// HasNetworkElement returns a boolean if a field has been set.
func (o *NetworkSupervisorCard) HasNetworkElement() bool {
	if o != nil && o.NetworkElement != nil {
		return true
	}

	return false
}

// SetNetworkElement gets a reference to the given NetworkElementRelationship and assigns it to the NetworkElement field.
func (o *NetworkSupervisorCard) SetNetworkElement(v NetworkElementRelationship) {
	o.NetworkElement = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NetworkSupervisorCard) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkSupervisorCard) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NetworkSupervisorCard) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NetworkSupervisorCard) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NetworkSupervisorCard) MarshalJSON() ([]byte, error) {
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
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.HardwareVersion != nil {
		toSerialize["HardwareVersion"] = o.HardwareVersion
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.NumberOfPorts != nil {
		toSerialize["NumberOfPorts"] = o.NumberOfPorts
	}
	if o.OperReason != nil {
		toSerialize["OperReason"] = o.OperReason
	}
	if o.OperState != nil {
		toSerialize["OperState"] = o.OperState
	}
	if o.PartNumber != nil {
		toSerialize["PartNumber"] = o.PartNumber
	}
	if o.PowerState != nil {
		toSerialize["PowerState"] = o.PowerState
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.SupervisorId != nil {
		toSerialize["SupervisorId"] = o.SupervisorId
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	if o.FcPorts != nil {
		toSerialize["FcPorts"] = o.FcPorts
	}
	if o.NetworkElement != nil {
		toSerialize["NetworkElement"] = o.NetworkElement
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NetworkSupervisorCard) UnmarshalJSON(bytes []byte) (err error) {
	type NetworkSupervisorCardWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The description of the supervisor card.
		Description *string `json:"Description,omitempty"`
		// The hardware version of the supervisor card.
		HardwareVersion *string `json:"HardwareVersion,omitempty"`
		// The name of the supervisor card like Supervisor Card-1.
		Name *string `json:"Name,omitempty"`
		// The number of ports on the supervisor card.
		NumberOfPorts *int64 `json:"NumberOfPorts,omitempty"`
		// The health issue of the supervisor card.
		OperReason *string `json:"OperReason,omitempty"`
		// The operational status of the supervisor card.
		OperState *string `json:"OperState,omitempty"`
		// The part number of the supervisor card.
		PartNumber *string `json:"PartNumber,omitempty"`
		// The power state of the supervisor card.
		PowerState *string `json:"PowerState,omitempty"`
		// The status of the supervisor card.
		Status *string `json:"Status,omitempty"`
		// The identifier for the supervisor card.
		SupervisorId *string `json:"SupervisorId,omitempty"`
		// The type of the supervisor card.
		Type *string `json:"Type,omitempty"`
		// An array of relationships to fcPhysicalPort resources.
		FcPorts          []FcPhysicalPortRelationship         `json:"FcPorts,omitempty"`
		NetworkElement   *NetworkElementRelationship          `json:"NetworkElement,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varNetworkSupervisorCardWithoutEmbeddedStruct := NetworkSupervisorCardWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNetworkSupervisorCardWithoutEmbeddedStruct)
	if err == nil {
		varNetworkSupervisorCard := _NetworkSupervisorCard{}
		varNetworkSupervisorCard.ClassId = varNetworkSupervisorCardWithoutEmbeddedStruct.ClassId
		varNetworkSupervisorCard.ObjectType = varNetworkSupervisorCardWithoutEmbeddedStruct.ObjectType
		varNetworkSupervisorCard.Description = varNetworkSupervisorCardWithoutEmbeddedStruct.Description
		varNetworkSupervisorCard.HardwareVersion = varNetworkSupervisorCardWithoutEmbeddedStruct.HardwareVersion
		varNetworkSupervisorCard.Name = varNetworkSupervisorCardWithoutEmbeddedStruct.Name
		varNetworkSupervisorCard.NumberOfPorts = varNetworkSupervisorCardWithoutEmbeddedStruct.NumberOfPorts
		varNetworkSupervisorCard.OperReason = varNetworkSupervisorCardWithoutEmbeddedStruct.OperReason
		varNetworkSupervisorCard.OperState = varNetworkSupervisorCardWithoutEmbeddedStruct.OperState
		varNetworkSupervisorCard.PartNumber = varNetworkSupervisorCardWithoutEmbeddedStruct.PartNumber
		varNetworkSupervisorCard.PowerState = varNetworkSupervisorCardWithoutEmbeddedStruct.PowerState
		varNetworkSupervisorCard.Status = varNetworkSupervisorCardWithoutEmbeddedStruct.Status
		varNetworkSupervisorCard.SupervisorId = varNetworkSupervisorCardWithoutEmbeddedStruct.SupervisorId
		varNetworkSupervisorCard.Type = varNetworkSupervisorCardWithoutEmbeddedStruct.Type
		varNetworkSupervisorCard.FcPorts = varNetworkSupervisorCardWithoutEmbeddedStruct.FcPorts
		varNetworkSupervisorCard.NetworkElement = varNetworkSupervisorCardWithoutEmbeddedStruct.NetworkElement
		varNetworkSupervisorCard.RegisteredDevice = varNetworkSupervisorCardWithoutEmbeddedStruct.RegisteredDevice
		*o = NetworkSupervisorCard(varNetworkSupervisorCard)
	} else {
		return err
	}

	varNetworkSupervisorCard := _NetworkSupervisorCard{}

	err = json.Unmarshal(bytes, &varNetworkSupervisorCard)
	if err == nil {
		o.EquipmentBase = varNetworkSupervisorCard.EquipmentBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "HardwareVersion")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "NumberOfPorts")
		delete(additionalProperties, "OperReason")
		delete(additionalProperties, "OperState")
		delete(additionalProperties, "PartNumber")
		delete(additionalProperties, "PowerState")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "SupervisorId")
		delete(additionalProperties, "Type")
		delete(additionalProperties, "FcPorts")
		delete(additionalProperties, "NetworkElement")
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

type NullableNetworkSupervisorCard struct {
	value *NetworkSupervisorCard
	isSet bool
}

func (v NullableNetworkSupervisorCard) Get() *NetworkSupervisorCard {
	return v.value
}

func (v *NullableNetworkSupervisorCard) Set(val *NetworkSupervisorCard) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkSupervisorCard) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkSupervisorCard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkSupervisorCard(val *NetworkSupervisorCard) *NullableNetworkSupervisorCard {
	return &NullableNetworkSupervisorCard{value: val, isSet: true}
}

func (v NullableNetworkSupervisorCard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkSupervisorCard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
