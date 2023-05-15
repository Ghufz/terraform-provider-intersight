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

// EquipmentTransceiverAllOf Definition of the list of properties defined in 'equipment.Transceiver', excluding properties defined in parent classes.
type EquipmentTransceiverAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Breakout port member in the Fabric Interconnect.
	AggregatePortId *int64 `json:"AggregatePortId,omitempty"`
	// The cisco extended Id number state of the pluggable SFP.
	CiscoExtendedIdNumber *string `json:"CiscoExtendedIdNumber,omitempty"`
	// Interface type of transceiver copper or fiber.
	InterfaceType *string `json:"InterfaceType,omitempty"`
	// The manufacturer part number of the pluggable SFP.
	ManufacturerPartNumber *string `json:"ManufacturerPartNumber,omitempty"`
	// Fabric extender identifier.
	ModuleId *int64 `json:"ModuleId,omitempty"`
	// The name of the pluggable transceiver.
	Name *string `json:"Name,omitempty"`
	// Operational speed of the transceiver.
	OperSpeed *string `json:"OperSpeed,omitempty"`
	// Operational state of the transceiver.
	OperState *string `json:"OperState,omitempty"`
	// Reason for this transceiver's operational state.
	OperStateQual *string `json:"OperStateQual,omitempty"`
	// Switch physical port identifier.
	PortId *int64 `json:"PortId,omitempty"`
	// Switch expansion slot module identifier.
	SlotId *int64 `json:"SlotId,omitempty"`
	// Status of the pluggable SFP.
	Status *string `json:"Status,omitempty"`
	// Switch Identifier that is local to a cluster.
	SwitchId *string `json:"SwitchId,omitempty"`
	// The type of the transceiver.
	Type                 *string                              `json:"Type,omitempty"`
	EtherHostPort        *EtherHostPortRelationship           `json:"EtherHostPort,omitempty"`
	EtherPhysicalPort    *EtherPhysicalPortRelationship       `json:"EtherPhysicalPort,omitempty"`
	FcPhysicalPort       *FcPhysicalPortRelationship          `json:"FcPhysicalPort,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EquipmentTransceiverAllOf EquipmentTransceiverAllOf

// NewEquipmentTransceiverAllOf instantiates a new EquipmentTransceiverAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquipmentTransceiverAllOf(classId string, objectType string) *EquipmentTransceiverAllOf {
	this := EquipmentTransceiverAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewEquipmentTransceiverAllOfWithDefaults instantiates a new EquipmentTransceiverAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquipmentTransceiverAllOfWithDefaults() *EquipmentTransceiverAllOf {
	this := EquipmentTransceiverAllOf{}
	var classId string = "equipment.Transceiver"
	this.ClassId = classId
	var objectType string = "equipment.Transceiver"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *EquipmentTransceiverAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *EquipmentTransceiverAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *EquipmentTransceiverAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *EquipmentTransceiverAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *EquipmentTransceiverAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *EquipmentTransceiverAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAggregatePortId returns the AggregatePortId field value if set, zero value otherwise.
func (o *EquipmentTransceiverAllOf) GetAggregatePortId() int64 {
	if o == nil || o.AggregatePortId == nil {
		var ret int64
		return ret
	}
	return *o.AggregatePortId
}

// GetAggregatePortIdOk returns a tuple with the AggregatePortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentTransceiverAllOf) GetAggregatePortIdOk() (*int64, bool) {
	if o == nil || o.AggregatePortId == nil {
		return nil, false
	}
	return o.AggregatePortId, true
}

// HasAggregatePortId returns a boolean if a field has been set.
func (o *EquipmentTransceiverAllOf) HasAggregatePortId() bool {
	if o != nil && o.AggregatePortId != nil {
		return true
	}

	return false
}

// SetAggregatePortId gets a reference to the given int64 and assigns it to the AggregatePortId field.
func (o *EquipmentTransceiverAllOf) SetAggregatePortId(v int64) {
	o.AggregatePortId = &v
}

// GetCiscoExtendedIdNumber returns the CiscoExtendedIdNumber field value if set, zero value otherwise.
func (o *EquipmentTransceiverAllOf) GetCiscoExtendedIdNumber() string {
	if o == nil || o.CiscoExtendedIdNumber == nil {
		var ret string
		return ret
	}
	return *o.CiscoExtendedIdNumber
}

// GetCiscoExtendedIdNumberOk returns a tuple with the CiscoExtendedIdNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentTransceiverAllOf) GetCiscoExtendedIdNumberOk() (*string, bool) {
	if o == nil || o.CiscoExtendedIdNumber == nil {
		return nil, false
	}
	return o.CiscoExtendedIdNumber, true
}

// HasCiscoExtendedIdNumber returns a boolean if a field has been set.
func (o *EquipmentTransceiverAllOf) HasCiscoExtendedIdNumber() bool {
	if o != nil && o.CiscoExtendedIdNumber != nil {
		return true
	}

	return false
}

// SetCiscoExtendedIdNumber gets a reference to the given string and assigns it to the CiscoExtendedIdNumber field.
func (o *EquipmentTransceiverAllOf) SetCiscoExtendedIdNumber(v string) {
	o.CiscoExtendedIdNumber = &v
}

// GetInterfaceType returns the InterfaceType field value if set, zero value otherwise.
func (o *EquipmentTransceiverAllOf) GetInterfaceType() string {
	if o == nil || o.InterfaceType == nil {
		var ret string
		return ret
	}
	return *o.InterfaceType
}

// GetInterfaceTypeOk returns a tuple with the InterfaceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentTransceiverAllOf) GetInterfaceTypeOk() (*string, bool) {
	if o == nil || o.InterfaceType == nil {
		return nil, false
	}
	return o.InterfaceType, true
}

// HasInterfaceType returns a boolean if a field has been set.
func (o *EquipmentTransceiverAllOf) HasInterfaceType() bool {
	if o != nil && o.InterfaceType != nil {
		return true
	}

	return false
}

// SetInterfaceType gets a reference to the given string and assigns it to the InterfaceType field.
func (o *EquipmentTransceiverAllOf) SetInterfaceType(v string) {
	o.InterfaceType = &v
}

// GetManufacturerPartNumber returns the ManufacturerPartNumber field value if set, zero value otherwise.
func (o *EquipmentTransceiverAllOf) GetManufacturerPartNumber() string {
	if o == nil || o.ManufacturerPartNumber == nil {
		var ret string
		return ret
	}
	return *o.ManufacturerPartNumber
}

// GetManufacturerPartNumberOk returns a tuple with the ManufacturerPartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentTransceiverAllOf) GetManufacturerPartNumberOk() (*string, bool) {
	if o == nil || o.ManufacturerPartNumber == nil {
		return nil, false
	}
	return o.ManufacturerPartNumber, true
}

// HasManufacturerPartNumber returns a boolean if a field has been set.
func (o *EquipmentTransceiverAllOf) HasManufacturerPartNumber() bool {
	if o != nil && o.ManufacturerPartNumber != nil {
		return true
	}

	return false
}

// SetManufacturerPartNumber gets a reference to the given string and assigns it to the ManufacturerPartNumber field.
func (o *EquipmentTransceiverAllOf) SetManufacturerPartNumber(v string) {
	o.ManufacturerPartNumber = &v
}

// GetModuleId returns the ModuleId field value if set, zero value otherwise.
func (o *EquipmentTransceiverAllOf) GetModuleId() int64 {
	if o == nil || o.ModuleId == nil {
		var ret int64
		return ret
	}
	return *o.ModuleId
}

// GetModuleIdOk returns a tuple with the ModuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentTransceiverAllOf) GetModuleIdOk() (*int64, bool) {
	if o == nil || o.ModuleId == nil {
		return nil, false
	}
	return o.ModuleId, true
}

// HasModuleId returns a boolean if a field has been set.
func (o *EquipmentTransceiverAllOf) HasModuleId() bool {
	if o != nil && o.ModuleId != nil {
		return true
	}

	return false
}

// SetModuleId gets a reference to the given int64 and assigns it to the ModuleId field.
func (o *EquipmentTransceiverAllOf) SetModuleId(v int64) {
	o.ModuleId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EquipmentTransceiverAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentTransceiverAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EquipmentTransceiverAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EquipmentTransceiverAllOf) SetName(v string) {
	o.Name = &v
}

// GetOperSpeed returns the OperSpeed field value if set, zero value otherwise.
func (o *EquipmentTransceiverAllOf) GetOperSpeed() string {
	if o == nil || o.OperSpeed == nil {
		var ret string
		return ret
	}
	return *o.OperSpeed
}

// GetOperSpeedOk returns a tuple with the OperSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentTransceiverAllOf) GetOperSpeedOk() (*string, bool) {
	if o == nil || o.OperSpeed == nil {
		return nil, false
	}
	return o.OperSpeed, true
}

// HasOperSpeed returns a boolean if a field has been set.
func (o *EquipmentTransceiverAllOf) HasOperSpeed() bool {
	if o != nil && o.OperSpeed != nil {
		return true
	}

	return false
}

// SetOperSpeed gets a reference to the given string and assigns it to the OperSpeed field.
func (o *EquipmentTransceiverAllOf) SetOperSpeed(v string) {
	o.OperSpeed = &v
}

// GetOperState returns the OperState field value if set, zero value otherwise.
func (o *EquipmentTransceiverAllOf) GetOperState() string {
	if o == nil || o.OperState == nil {
		var ret string
		return ret
	}
	return *o.OperState
}

// GetOperStateOk returns a tuple with the OperState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentTransceiverAllOf) GetOperStateOk() (*string, bool) {
	if o == nil || o.OperState == nil {
		return nil, false
	}
	return o.OperState, true
}

// HasOperState returns a boolean if a field has been set.
func (o *EquipmentTransceiverAllOf) HasOperState() bool {
	if o != nil && o.OperState != nil {
		return true
	}

	return false
}

// SetOperState gets a reference to the given string and assigns it to the OperState field.
func (o *EquipmentTransceiverAllOf) SetOperState(v string) {
	o.OperState = &v
}

// GetOperStateQual returns the OperStateQual field value if set, zero value otherwise.
func (o *EquipmentTransceiverAllOf) GetOperStateQual() string {
	if o == nil || o.OperStateQual == nil {
		var ret string
		return ret
	}
	return *o.OperStateQual
}

// GetOperStateQualOk returns a tuple with the OperStateQual field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentTransceiverAllOf) GetOperStateQualOk() (*string, bool) {
	if o == nil || o.OperStateQual == nil {
		return nil, false
	}
	return o.OperStateQual, true
}

// HasOperStateQual returns a boolean if a field has been set.
func (o *EquipmentTransceiverAllOf) HasOperStateQual() bool {
	if o != nil && o.OperStateQual != nil {
		return true
	}

	return false
}

// SetOperStateQual gets a reference to the given string and assigns it to the OperStateQual field.
func (o *EquipmentTransceiverAllOf) SetOperStateQual(v string) {
	o.OperStateQual = &v
}

// GetPortId returns the PortId field value if set, zero value otherwise.
func (o *EquipmentTransceiverAllOf) GetPortId() int64 {
	if o == nil || o.PortId == nil {
		var ret int64
		return ret
	}
	return *o.PortId
}

// GetPortIdOk returns a tuple with the PortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentTransceiverAllOf) GetPortIdOk() (*int64, bool) {
	if o == nil || o.PortId == nil {
		return nil, false
	}
	return o.PortId, true
}

// HasPortId returns a boolean if a field has been set.
func (o *EquipmentTransceiverAllOf) HasPortId() bool {
	if o != nil && o.PortId != nil {
		return true
	}

	return false
}

// SetPortId gets a reference to the given int64 and assigns it to the PortId field.
func (o *EquipmentTransceiverAllOf) SetPortId(v int64) {
	o.PortId = &v
}

// GetSlotId returns the SlotId field value if set, zero value otherwise.
func (o *EquipmentTransceiverAllOf) GetSlotId() int64 {
	if o == nil || o.SlotId == nil {
		var ret int64
		return ret
	}
	return *o.SlotId
}

// GetSlotIdOk returns a tuple with the SlotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentTransceiverAllOf) GetSlotIdOk() (*int64, bool) {
	if o == nil || o.SlotId == nil {
		return nil, false
	}
	return o.SlotId, true
}

// HasSlotId returns a boolean if a field has been set.
func (o *EquipmentTransceiverAllOf) HasSlotId() bool {
	if o != nil && o.SlotId != nil {
		return true
	}

	return false
}

// SetSlotId gets a reference to the given int64 and assigns it to the SlotId field.
func (o *EquipmentTransceiverAllOf) SetSlotId(v int64) {
	o.SlotId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *EquipmentTransceiverAllOf) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentTransceiverAllOf) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *EquipmentTransceiverAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *EquipmentTransceiverAllOf) SetStatus(v string) {
	o.Status = &v
}

// GetSwitchId returns the SwitchId field value if set, zero value otherwise.
func (o *EquipmentTransceiverAllOf) GetSwitchId() string {
	if o == nil || o.SwitchId == nil {
		var ret string
		return ret
	}
	return *o.SwitchId
}

// GetSwitchIdOk returns a tuple with the SwitchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentTransceiverAllOf) GetSwitchIdOk() (*string, bool) {
	if o == nil || o.SwitchId == nil {
		return nil, false
	}
	return o.SwitchId, true
}

// HasSwitchId returns a boolean if a field has been set.
func (o *EquipmentTransceiverAllOf) HasSwitchId() bool {
	if o != nil && o.SwitchId != nil {
		return true
	}

	return false
}

// SetSwitchId gets a reference to the given string and assigns it to the SwitchId field.
func (o *EquipmentTransceiverAllOf) SetSwitchId(v string) {
	o.SwitchId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *EquipmentTransceiverAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentTransceiverAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *EquipmentTransceiverAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *EquipmentTransceiverAllOf) SetType(v string) {
	o.Type = &v
}

// GetEtherHostPort returns the EtherHostPort field value if set, zero value otherwise.
func (o *EquipmentTransceiverAllOf) GetEtherHostPort() EtherHostPortRelationship {
	if o == nil || o.EtherHostPort == nil {
		var ret EtherHostPortRelationship
		return ret
	}
	return *o.EtherHostPort
}

// GetEtherHostPortOk returns a tuple with the EtherHostPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentTransceiverAllOf) GetEtherHostPortOk() (*EtherHostPortRelationship, bool) {
	if o == nil || o.EtherHostPort == nil {
		return nil, false
	}
	return o.EtherHostPort, true
}

// HasEtherHostPort returns a boolean if a field has been set.
func (o *EquipmentTransceiverAllOf) HasEtherHostPort() bool {
	if o != nil && o.EtherHostPort != nil {
		return true
	}

	return false
}

// SetEtherHostPort gets a reference to the given EtherHostPortRelationship and assigns it to the EtherHostPort field.
func (o *EquipmentTransceiverAllOf) SetEtherHostPort(v EtherHostPortRelationship) {
	o.EtherHostPort = &v
}

// GetEtherPhysicalPort returns the EtherPhysicalPort field value if set, zero value otherwise.
func (o *EquipmentTransceiverAllOf) GetEtherPhysicalPort() EtherPhysicalPortRelationship {
	if o == nil || o.EtherPhysicalPort == nil {
		var ret EtherPhysicalPortRelationship
		return ret
	}
	return *o.EtherPhysicalPort
}

// GetEtherPhysicalPortOk returns a tuple with the EtherPhysicalPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentTransceiverAllOf) GetEtherPhysicalPortOk() (*EtherPhysicalPortRelationship, bool) {
	if o == nil || o.EtherPhysicalPort == nil {
		return nil, false
	}
	return o.EtherPhysicalPort, true
}

// HasEtherPhysicalPort returns a boolean if a field has been set.
func (o *EquipmentTransceiverAllOf) HasEtherPhysicalPort() bool {
	if o != nil && o.EtherPhysicalPort != nil {
		return true
	}

	return false
}

// SetEtherPhysicalPort gets a reference to the given EtherPhysicalPortRelationship and assigns it to the EtherPhysicalPort field.
func (o *EquipmentTransceiverAllOf) SetEtherPhysicalPort(v EtherPhysicalPortRelationship) {
	o.EtherPhysicalPort = &v
}

// GetFcPhysicalPort returns the FcPhysicalPort field value if set, zero value otherwise.
func (o *EquipmentTransceiverAllOf) GetFcPhysicalPort() FcPhysicalPortRelationship {
	if o == nil || o.FcPhysicalPort == nil {
		var ret FcPhysicalPortRelationship
		return ret
	}
	return *o.FcPhysicalPort
}

// GetFcPhysicalPortOk returns a tuple with the FcPhysicalPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentTransceiverAllOf) GetFcPhysicalPortOk() (*FcPhysicalPortRelationship, bool) {
	if o == nil || o.FcPhysicalPort == nil {
		return nil, false
	}
	return o.FcPhysicalPort, true
}

// HasFcPhysicalPort returns a boolean if a field has been set.
func (o *EquipmentTransceiverAllOf) HasFcPhysicalPort() bool {
	if o != nil && o.FcPhysicalPort != nil {
		return true
	}

	return false
}

// SetFcPhysicalPort gets a reference to the given FcPhysicalPortRelationship and assigns it to the FcPhysicalPort field.
func (o *EquipmentTransceiverAllOf) SetFcPhysicalPort(v FcPhysicalPortRelationship) {
	o.FcPhysicalPort = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *EquipmentTransceiverAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentTransceiverAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *EquipmentTransceiverAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *EquipmentTransceiverAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o EquipmentTransceiverAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AggregatePortId != nil {
		toSerialize["AggregatePortId"] = o.AggregatePortId
	}
	if o.CiscoExtendedIdNumber != nil {
		toSerialize["CiscoExtendedIdNumber"] = o.CiscoExtendedIdNumber
	}
	if o.InterfaceType != nil {
		toSerialize["InterfaceType"] = o.InterfaceType
	}
	if o.ManufacturerPartNumber != nil {
		toSerialize["ManufacturerPartNumber"] = o.ManufacturerPartNumber
	}
	if o.ModuleId != nil {
		toSerialize["ModuleId"] = o.ModuleId
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.OperSpeed != nil {
		toSerialize["OperSpeed"] = o.OperSpeed
	}
	if o.OperState != nil {
		toSerialize["OperState"] = o.OperState
	}
	if o.OperStateQual != nil {
		toSerialize["OperStateQual"] = o.OperStateQual
	}
	if o.PortId != nil {
		toSerialize["PortId"] = o.PortId
	}
	if o.SlotId != nil {
		toSerialize["SlotId"] = o.SlotId
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.SwitchId != nil {
		toSerialize["SwitchId"] = o.SwitchId
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	if o.EtherHostPort != nil {
		toSerialize["EtherHostPort"] = o.EtherHostPort
	}
	if o.EtherPhysicalPort != nil {
		toSerialize["EtherPhysicalPort"] = o.EtherPhysicalPort
	}
	if o.FcPhysicalPort != nil {
		toSerialize["FcPhysicalPort"] = o.FcPhysicalPort
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EquipmentTransceiverAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varEquipmentTransceiverAllOf := _EquipmentTransceiverAllOf{}

	if err = json.Unmarshal(bytes, &varEquipmentTransceiverAllOf); err == nil {
		*o = EquipmentTransceiverAllOf(varEquipmentTransceiverAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AggregatePortId")
		delete(additionalProperties, "CiscoExtendedIdNumber")
		delete(additionalProperties, "InterfaceType")
		delete(additionalProperties, "ManufacturerPartNumber")
		delete(additionalProperties, "ModuleId")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "OperSpeed")
		delete(additionalProperties, "OperState")
		delete(additionalProperties, "OperStateQual")
		delete(additionalProperties, "PortId")
		delete(additionalProperties, "SlotId")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "SwitchId")
		delete(additionalProperties, "Type")
		delete(additionalProperties, "EtherHostPort")
		delete(additionalProperties, "EtherPhysicalPort")
		delete(additionalProperties, "FcPhysicalPort")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEquipmentTransceiverAllOf struct {
	value *EquipmentTransceiverAllOf
	isSet bool
}

func (v NullableEquipmentTransceiverAllOf) Get() *EquipmentTransceiverAllOf {
	return v.value
}

func (v *NullableEquipmentTransceiverAllOf) Set(val *EquipmentTransceiverAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEquipmentTransceiverAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEquipmentTransceiverAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquipmentTransceiverAllOf(val *EquipmentTransceiverAllOf) *NullableEquipmentTransceiverAllOf {
	return &NullableEquipmentTransceiverAllOf{value: val, isSet: true}
}

func (v NullableEquipmentTransceiverAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquipmentTransceiverAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
