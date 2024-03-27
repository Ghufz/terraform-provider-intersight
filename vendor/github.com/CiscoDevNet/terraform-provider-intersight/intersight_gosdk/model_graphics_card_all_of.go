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

// GraphicsCardAllOf Definition of the list of properties defined in 'graphics.Card', excluding properties defined in parent classes.
type GraphicsCardAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The id of the graphics card.
	CardId *int64 `json:"CardId,omitempty"`
	// This field displays the description of the Graphics Processing Unit.
	Description *string `json:"Description,omitempty"`
	// The device id of the graphics card.
	DeviceId *int64 `json:"DeviceId,omitempty"`
	// The expander slot information of the card.
	ExpanderSlot *string `json:"ExpanderSlot,omitempty"`
	// The firmware version of the graphics card.
	FirmwareVersion *string `json:"FirmwareVersion,omitempty"`
	// The identifier of the graphics processor unit.
	GpuId *string `json:"GpuId,omitempty"`
	// This field indicates whether the Graphics Processing Unit is supported on the server or not.
	IsPlatformSupported *bool `json:"IsPlatformSupported,omitempty"`
	// The current mode of the graphics card.
	Mode *string `json:"Mode,omitempty"`
	// The number of controllers under each card.
	NumGpus    *string  `json:"NumGpus,omitempty"`
	OperReason []string `json:"OperReason,omitempty"`
	// The current operational state of the graphics card.
	OperState *string `json:"OperState,omitempty"`
	// This field displays the part number of the Graphics Processing Unit.
	PartNumber *string `json:"PartNumber,omitempty"`
	// The PCI address of the graphics card.
	PciAddress *string `json:"PciAddress,omitempty"`
	// This list contains the PCI address of all controllers for corresponding card.
	PciAddressList *string `json:"PciAddressList,omitempty"`
	// The PCI slot name of the graphics card.
	PciSlot *string `json:"PciSlot,omitempty"`
	// This field displays the product ID of the Graphics Processing Unit.
	Pid *string `json:"Pid,omitempty"`
	// The sub device id of the graphics processor unit.
	SubDeviceId *int64 `json:"SubDeviceId,omitempty"`
	// The sub vendor id of the graphics processor unit.
	SubVendorId *int64 `json:"SubVendorId,omitempty"`
	// The vendor id of the graphics processor unit.
	VendorId        *int64                       `json:"VendorId,omitempty"`
	ComputeBlade    *ComputeBladeRelationship    `json:"ComputeBlade,omitempty"`
	ComputeBoard    *ComputeBoardRelationship    `json:"ComputeBoard,omitempty"`
	ComputeRackUnit *ComputeRackUnitRelationship `json:"ComputeRackUnit,omitempty"`
	// An array of relationships to graphicsController resources.
	GraphicsControllers []GraphicsControllerRelationship     `json:"GraphicsControllers,omitempty"`
	InventoryDeviceInfo *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
	PciDevice           *PciDeviceRelationship               `json:"PciDevice,omitempty"`
	PciNode             *PciNodeRelationship                 `json:"PciNode,omitempty"`
	RegisteredDevice    *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	// An array of relationships to firmwareRunningFirmware resources.
	RunningFirmware      []FirmwareRunningFirmwareRelationship `json:"RunningFirmware,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GraphicsCardAllOf GraphicsCardAllOf

// NewGraphicsCardAllOf instantiates a new GraphicsCardAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGraphicsCardAllOf(classId string, objectType string) *GraphicsCardAllOf {
	this := GraphicsCardAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewGraphicsCardAllOfWithDefaults instantiates a new GraphicsCardAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGraphicsCardAllOfWithDefaults() *GraphicsCardAllOf {
	this := GraphicsCardAllOf{}
	var classId string = "graphics.Card"
	this.ClassId = classId
	var objectType string = "graphics.Card"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *GraphicsCardAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *GraphicsCardAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *GraphicsCardAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *GraphicsCardAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *GraphicsCardAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *GraphicsCardAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCardId returns the CardId field value if set, zero value otherwise.
func (o *GraphicsCardAllOf) GetCardId() int64 {
	if o == nil || o.CardId == nil {
		var ret int64
		return ret
	}
	return *o.CardId
}

// GetCardIdOk returns a tuple with the CardId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphicsCardAllOf) GetCardIdOk() (*int64, bool) {
	if o == nil || o.CardId == nil {
		return nil, false
	}
	return o.CardId, true
}

// HasCardId returns a boolean if a field has been set.
func (o *GraphicsCardAllOf) HasCardId() bool {
	if o != nil && o.CardId != nil {
		return true
	}

	return false
}

// SetCardId gets a reference to the given int64 and assigns it to the CardId field.
func (o *GraphicsCardAllOf) SetCardId(v int64) {
	o.CardId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GraphicsCardAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphicsCardAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GraphicsCardAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GraphicsCardAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *GraphicsCardAllOf) GetDeviceId() int64 {
	if o == nil || o.DeviceId == nil {
		var ret int64
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphicsCardAllOf) GetDeviceIdOk() (*int64, bool) {
	if o == nil || o.DeviceId == nil {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *GraphicsCardAllOf) HasDeviceId() bool {
	if o != nil && o.DeviceId != nil {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given int64 and assigns it to the DeviceId field.
func (o *GraphicsCardAllOf) SetDeviceId(v int64) {
	o.DeviceId = &v
}

// GetExpanderSlot returns the ExpanderSlot field value if set, zero value otherwise.
func (o *GraphicsCardAllOf) GetExpanderSlot() string {
	if o == nil || o.ExpanderSlot == nil {
		var ret string
		return ret
	}
	return *o.ExpanderSlot
}

// GetExpanderSlotOk returns a tuple with the ExpanderSlot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphicsCardAllOf) GetExpanderSlotOk() (*string, bool) {
	if o == nil || o.ExpanderSlot == nil {
		return nil, false
	}
	return o.ExpanderSlot, true
}

// HasExpanderSlot returns a boolean if a field has been set.
func (o *GraphicsCardAllOf) HasExpanderSlot() bool {
	if o != nil && o.ExpanderSlot != nil {
		return true
	}

	return false
}

// SetExpanderSlot gets a reference to the given string and assigns it to the ExpanderSlot field.
func (o *GraphicsCardAllOf) SetExpanderSlot(v string) {
	o.ExpanderSlot = &v
}

// GetFirmwareVersion returns the FirmwareVersion field value if set, zero value otherwise.
func (o *GraphicsCardAllOf) GetFirmwareVersion() string {
	if o == nil || o.FirmwareVersion == nil {
		var ret string
		return ret
	}
	return *o.FirmwareVersion
}

// GetFirmwareVersionOk returns a tuple with the FirmwareVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphicsCardAllOf) GetFirmwareVersionOk() (*string, bool) {
	if o == nil || o.FirmwareVersion == nil {
		return nil, false
	}
	return o.FirmwareVersion, true
}

// HasFirmwareVersion returns a boolean if a field has been set.
func (o *GraphicsCardAllOf) HasFirmwareVersion() bool {
	if o != nil && o.FirmwareVersion != nil {
		return true
	}

	return false
}

// SetFirmwareVersion gets a reference to the given string and assigns it to the FirmwareVersion field.
func (o *GraphicsCardAllOf) SetFirmwareVersion(v string) {
	o.FirmwareVersion = &v
}

// GetGpuId returns the GpuId field value if set, zero value otherwise.
func (o *GraphicsCardAllOf) GetGpuId() string {
	if o == nil || o.GpuId == nil {
		var ret string
		return ret
	}
	return *o.GpuId
}

// GetGpuIdOk returns a tuple with the GpuId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphicsCardAllOf) GetGpuIdOk() (*string, bool) {
	if o == nil || o.GpuId == nil {
		return nil, false
	}
	return o.GpuId, true
}

// HasGpuId returns a boolean if a field has been set.
func (o *GraphicsCardAllOf) HasGpuId() bool {
	if o != nil && o.GpuId != nil {
		return true
	}

	return false
}

// SetGpuId gets a reference to the given string and assigns it to the GpuId field.
func (o *GraphicsCardAllOf) SetGpuId(v string) {
	o.GpuId = &v
}

// GetIsPlatformSupported returns the IsPlatformSupported field value if set, zero value otherwise.
func (o *GraphicsCardAllOf) GetIsPlatformSupported() bool {
	if o == nil || o.IsPlatformSupported == nil {
		var ret bool
		return ret
	}
	return *o.IsPlatformSupported
}

// GetIsPlatformSupportedOk returns a tuple with the IsPlatformSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphicsCardAllOf) GetIsPlatformSupportedOk() (*bool, bool) {
	if o == nil || o.IsPlatformSupported == nil {
		return nil, false
	}
	return o.IsPlatformSupported, true
}

// HasIsPlatformSupported returns a boolean if a field has been set.
func (o *GraphicsCardAllOf) HasIsPlatformSupported() bool {
	if o != nil && o.IsPlatformSupported != nil {
		return true
	}

	return false
}

// SetIsPlatformSupported gets a reference to the given bool and assigns it to the IsPlatformSupported field.
func (o *GraphicsCardAllOf) SetIsPlatformSupported(v bool) {
	o.IsPlatformSupported = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *GraphicsCardAllOf) GetMode() string {
	if o == nil || o.Mode == nil {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphicsCardAllOf) GetModeOk() (*string, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *GraphicsCardAllOf) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *GraphicsCardAllOf) SetMode(v string) {
	o.Mode = &v
}

// GetNumGpus returns the NumGpus field value if set, zero value otherwise.
func (o *GraphicsCardAllOf) GetNumGpus() string {
	if o == nil || o.NumGpus == nil {
		var ret string
		return ret
	}
	return *o.NumGpus
}

// GetNumGpusOk returns a tuple with the NumGpus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphicsCardAllOf) GetNumGpusOk() (*string, bool) {
	if o == nil || o.NumGpus == nil {
		return nil, false
	}
	return o.NumGpus, true
}

// HasNumGpus returns a boolean if a field has been set.
func (o *GraphicsCardAllOf) HasNumGpus() bool {
	if o != nil && o.NumGpus != nil {
		return true
	}

	return false
}

// SetNumGpus gets a reference to the given string and assigns it to the NumGpus field.
func (o *GraphicsCardAllOf) SetNumGpus(v string) {
	o.NumGpus = &v
}

// GetOperReason returns the OperReason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GraphicsCardAllOf) GetOperReason() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.OperReason
}

// GetOperReasonOk returns a tuple with the OperReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GraphicsCardAllOf) GetOperReasonOk() ([]string, bool) {
	if o == nil || o.OperReason == nil {
		return nil, false
	}
	return o.OperReason, true
}

// HasOperReason returns a boolean if a field has been set.
func (o *GraphicsCardAllOf) HasOperReason() bool {
	if o != nil && o.OperReason != nil {
		return true
	}

	return false
}

// SetOperReason gets a reference to the given []string and assigns it to the OperReason field.
func (o *GraphicsCardAllOf) SetOperReason(v []string) {
	o.OperReason = v
}

// GetOperState returns the OperState field value if set, zero value otherwise.
func (o *GraphicsCardAllOf) GetOperState() string {
	if o == nil || o.OperState == nil {
		var ret string
		return ret
	}
	return *o.OperState
}

// GetOperStateOk returns a tuple with the OperState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphicsCardAllOf) GetOperStateOk() (*string, bool) {
	if o == nil || o.OperState == nil {
		return nil, false
	}
	return o.OperState, true
}

// HasOperState returns a boolean if a field has been set.
func (o *GraphicsCardAllOf) HasOperState() bool {
	if o != nil && o.OperState != nil {
		return true
	}

	return false
}

// SetOperState gets a reference to the given string and assigns it to the OperState field.
func (o *GraphicsCardAllOf) SetOperState(v string) {
	o.OperState = &v
}

// GetPartNumber returns the PartNumber field value if set, zero value otherwise.
func (o *GraphicsCardAllOf) GetPartNumber() string {
	if o == nil || o.PartNumber == nil {
		var ret string
		return ret
	}
	return *o.PartNumber
}

// GetPartNumberOk returns a tuple with the PartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphicsCardAllOf) GetPartNumberOk() (*string, bool) {
	if o == nil || o.PartNumber == nil {
		return nil, false
	}
	return o.PartNumber, true
}

// HasPartNumber returns a boolean if a field has been set.
func (o *GraphicsCardAllOf) HasPartNumber() bool {
	if o != nil && o.PartNumber != nil {
		return true
	}

	return false
}

// SetPartNumber gets a reference to the given string and assigns it to the PartNumber field.
func (o *GraphicsCardAllOf) SetPartNumber(v string) {
	o.PartNumber = &v
}

// GetPciAddress returns the PciAddress field value if set, zero value otherwise.
func (o *GraphicsCardAllOf) GetPciAddress() string {
	if o == nil || o.PciAddress == nil {
		var ret string
		return ret
	}
	return *o.PciAddress
}

// GetPciAddressOk returns a tuple with the PciAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphicsCardAllOf) GetPciAddressOk() (*string, bool) {
	if o == nil || o.PciAddress == nil {
		return nil, false
	}
	return o.PciAddress, true
}

// HasPciAddress returns a boolean if a field has been set.
func (o *GraphicsCardAllOf) HasPciAddress() bool {
	if o != nil && o.PciAddress != nil {
		return true
	}

	return false
}

// SetPciAddress gets a reference to the given string and assigns it to the PciAddress field.
func (o *GraphicsCardAllOf) SetPciAddress(v string) {
	o.PciAddress = &v
}

// GetPciAddressList returns the PciAddressList field value if set, zero value otherwise.
func (o *GraphicsCardAllOf) GetPciAddressList() string {
	if o == nil || o.PciAddressList == nil {
		var ret string
		return ret
	}
	return *o.PciAddressList
}

// GetPciAddressListOk returns a tuple with the PciAddressList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphicsCardAllOf) GetPciAddressListOk() (*string, bool) {
	if o == nil || o.PciAddressList == nil {
		return nil, false
	}
	return o.PciAddressList, true
}

// HasPciAddressList returns a boolean if a field has been set.
func (o *GraphicsCardAllOf) HasPciAddressList() bool {
	if o != nil && o.PciAddressList != nil {
		return true
	}

	return false
}

// SetPciAddressList gets a reference to the given string and assigns it to the PciAddressList field.
func (o *GraphicsCardAllOf) SetPciAddressList(v string) {
	o.PciAddressList = &v
}

// GetPciSlot returns the PciSlot field value if set, zero value otherwise.
func (o *GraphicsCardAllOf) GetPciSlot() string {
	if o == nil || o.PciSlot == nil {
		var ret string
		return ret
	}
	return *o.PciSlot
}

// GetPciSlotOk returns a tuple with the PciSlot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphicsCardAllOf) GetPciSlotOk() (*string, bool) {
	if o == nil || o.PciSlot == nil {
		return nil, false
	}
	return o.PciSlot, true
}

// HasPciSlot returns a boolean if a field has been set.
func (o *GraphicsCardAllOf) HasPciSlot() bool {
	if o != nil && o.PciSlot != nil {
		return true
	}

	return false
}

// SetPciSlot gets a reference to the given string and assigns it to the PciSlot field.
func (o *GraphicsCardAllOf) SetPciSlot(v string) {
	o.PciSlot = &v
}

// GetPid returns the Pid field value if set, zero value otherwise.
func (o *GraphicsCardAllOf) GetPid() string {
	if o == nil || o.Pid == nil {
		var ret string
		return ret
	}
	return *o.Pid
}

// GetPidOk returns a tuple with the Pid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphicsCardAllOf) GetPidOk() (*string, bool) {
	if o == nil || o.Pid == nil {
		return nil, false
	}
	return o.Pid, true
}

// HasPid returns a boolean if a field has been set.
func (o *GraphicsCardAllOf) HasPid() bool {
	if o != nil && o.Pid != nil {
		return true
	}

	return false
}

// SetPid gets a reference to the given string and assigns it to the Pid field.
func (o *GraphicsCardAllOf) SetPid(v string) {
	o.Pid = &v
}

// GetSubDeviceId returns the SubDeviceId field value if set, zero value otherwise.
func (o *GraphicsCardAllOf) GetSubDeviceId() int64 {
	if o == nil || o.SubDeviceId == nil {
		var ret int64
		return ret
	}
	return *o.SubDeviceId
}

// GetSubDeviceIdOk returns a tuple with the SubDeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphicsCardAllOf) GetSubDeviceIdOk() (*int64, bool) {
	if o == nil || o.SubDeviceId == nil {
		return nil, false
	}
	return o.SubDeviceId, true
}

// HasSubDeviceId returns a boolean if a field has been set.
func (o *GraphicsCardAllOf) HasSubDeviceId() bool {
	if o != nil && o.SubDeviceId != nil {
		return true
	}

	return false
}

// SetSubDeviceId gets a reference to the given int64 and assigns it to the SubDeviceId field.
func (o *GraphicsCardAllOf) SetSubDeviceId(v int64) {
	o.SubDeviceId = &v
}

// GetSubVendorId returns the SubVendorId field value if set, zero value otherwise.
func (o *GraphicsCardAllOf) GetSubVendorId() int64 {
	if o == nil || o.SubVendorId == nil {
		var ret int64
		return ret
	}
	return *o.SubVendorId
}

// GetSubVendorIdOk returns a tuple with the SubVendorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphicsCardAllOf) GetSubVendorIdOk() (*int64, bool) {
	if o == nil || o.SubVendorId == nil {
		return nil, false
	}
	return o.SubVendorId, true
}

// HasSubVendorId returns a boolean if a field has been set.
func (o *GraphicsCardAllOf) HasSubVendorId() bool {
	if o != nil && o.SubVendorId != nil {
		return true
	}

	return false
}

// SetSubVendorId gets a reference to the given int64 and assigns it to the SubVendorId field.
func (o *GraphicsCardAllOf) SetSubVendorId(v int64) {
	o.SubVendorId = &v
}

// GetVendorId returns the VendorId field value if set, zero value otherwise.
func (o *GraphicsCardAllOf) GetVendorId() int64 {
	if o == nil || o.VendorId == nil {
		var ret int64
		return ret
	}
	return *o.VendorId
}

// GetVendorIdOk returns a tuple with the VendorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphicsCardAllOf) GetVendorIdOk() (*int64, bool) {
	if o == nil || o.VendorId == nil {
		return nil, false
	}
	return o.VendorId, true
}

// HasVendorId returns a boolean if a field has been set.
func (o *GraphicsCardAllOf) HasVendorId() bool {
	if o != nil && o.VendorId != nil {
		return true
	}

	return false
}

// SetVendorId gets a reference to the given int64 and assigns it to the VendorId field.
func (o *GraphicsCardAllOf) SetVendorId(v int64) {
	o.VendorId = &v
}

// GetComputeBlade returns the ComputeBlade field value if set, zero value otherwise.
func (o *GraphicsCardAllOf) GetComputeBlade() ComputeBladeRelationship {
	if o == nil || o.ComputeBlade == nil {
		var ret ComputeBladeRelationship
		return ret
	}
	return *o.ComputeBlade
}

// GetComputeBladeOk returns a tuple with the ComputeBlade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphicsCardAllOf) GetComputeBladeOk() (*ComputeBladeRelationship, bool) {
	if o == nil || o.ComputeBlade == nil {
		return nil, false
	}
	return o.ComputeBlade, true
}

// HasComputeBlade returns a boolean if a field has been set.
func (o *GraphicsCardAllOf) HasComputeBlade() bool {
	if o != nil && o.ComputeBlade != nil {
		return true
	}

	return false
}

// SetComputeBlade gets a reference to the given ComputeBladeRelationship and assigns it to the ComputeBlade field.
func (o *GraphicsCardAllOf) SetComputeBlade(v ComputeBladeRelationship) {
	o.ComputeBlade = &v
}

// GetComputeBoard returns the ComputeBoard field value if set, zero value otherwise.
func (o *GraphicsCardAllOf) GetComputeBoard() ComputeBoardRelationship {
	if o == nil || o.ComputeBoard == nil {
		var ret ComputeBoardRelationship
		return ret
	}
	return *o.ComputeBoard
}

// GetComputeBoardOk returns a tuple with the ComputeBoard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphicsCardAllOf) GetComputeBoardOk() (*ComputeBoardRelationship, bool) {
	if o == nil || o.ComputeBoard == nil {
		return nil, false
	}
	return o.ComputeBoard, true
}

// HasComputeBoard returns a boolean if a field has been set.
func (o *GraphicsCardAllOf) HasComputeBoard() bool {
	if o != nil && o.ComputeBoard != nil {
		return true
	}

	return false
}

// SetComputeBoard gets a reference to the given ComputeBoardRelationship and assigns it to the ComputeBoard field.
func (o *GraphicsCardAllOf) SetComputeBoard(v ComputeBoardRelationship) {
	o.ComputeBoard = &v
}

// GetComputeRackUnit returns the ComputeRackUnit field value if set, zero value otherwise.
func (o *GraphicsCardAllOf) GetComputeRackUnit() ComputeRackUnitRelationship {
	if o == nil || o.ComputeRackUnit == nil {
		var ret ComputeRackUnitRelationship
		return ret
	}
	return *o.ComputeRackUnit
}

// GetComputeRackUnitOk returns a tuple with the ComputeRackUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphicsCardAllOf) GetComputeRackUnitOk() (*ComputeRackUnitRelationship, bool) {
	if o == nil || o.ComputeRackUnit == nil {
		return nil, false
	}
	return o.ComputeRackUnit, true
}

// HasComputeRackUnit returns a boolean if a field has been set.
func (o *GraphicsCardAllOf) HasComputeRackUnit() bool {
	if o != nil && o.ComputeRackUnit != nil {
		return true
	}

	return false
}

// SetComputeRackUnit gets a reference to the given ComputeRackUnitRelationship and assigns it to the ComputeRackUnit field.
func (o *GraphicsCardAllOf) SetComputeRackUnit(v ComputeRackUnitRelationship) {
	o.ComputeRackUnit = &v
}

// GetGraphicsControllers returns the GraphicsControllers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GraphicsCardAllOf) GetGraphicsControllers() []GraphicsControllerRelationship {
	if o == nil {
		var ret []GraphicsControllerRelationship
		return ret
	}
	return o.GraphicsControllers
}

// GetGraphicsControllersOk returns a tuple with the GraphicsControllers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GraphicsCardAllOf) GetGraphicsControllersOk() ([]GraphicsControllerRelationship, bool) {
	if o == nil || o.GraphicsControllers == nil {
		return nil, false
	}
	return o.GraphicsControllers, true
}

// HasGraphicsControllers returns a boolean if a field has been set.
func (o *GraphicsCardAllOf) HasGraphicsControllers() bool {
	if o != nil && o.GraphicsControllers != nil {
		return true
	}

	return false
}

// SetGraphicsControllers gets a reference to the given []GraphicsControllerRelationship and assigns it to the GraphicsControllers field.
func (o *GraphicsCardAllOf) SetGraphicsControllers(v []GraphicsControllerRelationship) {
	o.GraphicsControllers = v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *GraphicsCardAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphicsCardAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *GraphicsCardAllOf) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *GraphicsCardAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetPciDevice returns the PciDevice field value if set, zero value otherwise.
func (o *GraphicsCardAllOf) GetPciDevice() PciDeviceRelationship {
	if o == nil || o.PciDevice == nil {
		var ret PciDeviceRelationship
		return ret
	}
	return *o.PciDevice
}

// GetPciDeviceOk returns a tuple with the PciDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphicsCardAllOf) GetPciDeviceOk() (*PciDeviceRelationship, bool) {
	if o == nil || o.PciDevice == nil {
		return nil, false
	}
	return o.PciDevice, true
}

// HasPciDevice returns a boolean if a field has been set.
func (o *GraphicsCardAllOf) HasPciDevice() bool {
	if o != nil && o.PciDevice != nil {
		return true
	}

	return false
}

// SetPciDevice gets a reference to the given PciDeviceRelationship and assigns it to the PciDevice field.
func (o *GraphicsCardAllOf) SetPciDevice(v PciDeviceRelationship) {
	o.PciDevice = &v
}

// GetPciNode returns the PciNode field value if set, zero value otherwise.
func (o *GraphicsCardAllOf) GetPciNode() PciNodeRelationship {
	if o == nil || o.PciNode == nil {
		var ret PciNodeRelationship
		return ret
	}
	return *o.PciNode
}

// GetPciNodeOk returns a tuple with the PciNode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphicsCardAllOf) GetPciNodeOk() (*PciNodeRelationship, bool) {
	if o == nil || o.PciNode == nil {
		return nil, false
	}
	return o.PciNode, true
}

// HasPciNode returns a boolean if a field has been set.
func (o *GraphicsCardAllOf) HasPciNode() bool {
	if o != nil && o.PciNode != nil {
		return true
	}

	return false
}

// SetPciNode gets a reference to the given PciNodeRelationship and assigns it to the PciNode field.
func (o *GraphicsCardAllOf) SetPciNode(v PciNodeRelationship) {
	o.PciNode = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *GraphicsCardAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphicsCardAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *GraphicsCardAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *GraphicsCardAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

// GetRunningFirmware returns the RunningFirmware field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GraphicsCardAllOf) GetRunningFirmware() []FirmwareRunningFirmwareRelationship {
	if o == nil {
		var ret []FirmwareRunningFirmwareRelationship
		return ret
	}
	return o.RunningFirmware
}

// GetRunningFirmwareOk returns a tuple with the RunningFirmware field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GraphicsCardAllOf) GetRunningFirmwareOk() ([]FirmwareRunningFirmwareRelationship, bool) {
	if o == nil || o.RunningFirmware == nil {
		return nil, false
	}
	return o.RunningFirmware, true
}

// HasRunningFirmware returns a boolean if a field has been set.
func (o *GraphicsCardAllOf) HasRunningFirmware() bool {
	if o != nil && o.RunningFirmware != nil {
		return true
	}

	return false
}

// SetRunningFirmware gets a reference to the given []FirmwareRunningFirmwareRelationship and assigns it to the RunningFirmware field.
func (o *GraphicsCardAllOf) SetRunningFirmware(v []FirmwareRunningFirmwareRelationship) {
	o.RunningFirmware = v
}

func (o GraphicsCardAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.CardId != nil {
		toSerialize["CardId"] = o.CardId
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.DeviceId != nil {
		toSerialize["DeviceId"] = o.DeviceId
	}
	if o.ExpanderSlot != nil {
		toSerialize["ExpanderSlot"] = o.ExpanderSlot
	}
	if o.FirmwareVersion != nil {
		toSerialize["FirmwareVersion"] = o.FirmwareVersion
	}
	if o.GpuId != nil {
		toSerialize["GpuId"] = o.GpuId
	}
	if o.IsPlatformSupported != nil {
		toSerialize["IsPlatformSupported"] = o.IsPlatformSupported
	}
	if o.Mode != nil {
		toSerialize["Mode"] = o.Mode
	}
	if o.NumGpus != nil {
		toSerialize["NumGpus"] = o.NumGpus
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
	if o.PciAddress != nil {
		toSerialize["PciAddress"] = o.PciAddress
	}
	if o.PciAddressList != nil {
		toSerialize["PciAddressList"] = o.PciAddressList
	}
	if o.PciSlot != nil {
		toSerialize["PciSlot"] = o.PciSlot
	}
	if o.Pid != nil {
		toSerialize["Pid"] = o.Pid
	}
	if o.SubDeviceId != nil {
		toSerialize["SubDeviceId"] = o.SubDeviceId
	}
	if o.SubVendorId != nil {
		toSerialize["SubVendorId"] = o.SubVendorId
	}
	if o.VendorId != nil {
		toSerialize["VendorId"] = o.VendorId
	}
	if o.ComputeBlade != nil {
		toSerialize["ComputeBlade"] = o.ComputeBlade
	}
	if o.ComputeBoard != nil {
		toSerialize["ComputeBoard"] = o.ComputeBoard
	}
	if o.ComputeRackUnit != nil {
		toSerialize["ComputeRackUnit"] = o.ComputeRackUnit
	}
	if o.GraphicsControllers != nil {
		toSerialize["GraphicsControllers"] = o.GraphicsControllers
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.PciDevice != nil {
		toSerialize["PciDevice"] = o.PciDevice
	}
	if o.PciNode != nil {
		toSerialize["PciNode"] = o.PciNode
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	if o.RunningFirmware != nil {
		toSerialize["RunningFirmware"] = o.RunningFirmware
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *GraphicsCardAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varGraphicsCardAllOf := _GraphicsCardAllOf{}

	if err = json.Unmarshal(bytes, &varGraphicsCardAllOf); err == nil {
		*o = GraphicsCardAllOf(varGraphicsCardAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CardId")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "DeviceId")
		delete(additionalProperties, "ExpanderSlot")
		delete(additionalProperties, "FirmwareVersion")
		delete(additionalProperties, "GpuId")
		delete(additionalProperties, "IsPlatformSupported")
		delete(additionalProperties, "Mode")
		delete(additionalProperties, "NumGpus")
		delete(additionalProperties, "OperReason")
		delete(additionalProperties, "OperState")
		delete(additionalProperties, "PartNumber")
		delete(additionalProperties, "PciAddress")
		delete(additionalProperties, "PciAddressList")
		delete(additionalProperties, "PciSlot")
		delete(additionalProperties, "Pid")
		delete(additionalProperties, "SubDeviceId")
		delete(additionalProperties, "SubVendorId")
		delete(additionalProperties, "VendorId")
		delete(additionalProperties, "ComputeBlade")
		delete(additionalProperties, "ComputeBoard")
		delete(additionalProperties, "ComputeRackUnit")
		delete(additionalProperties, "GraphicsControllers")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "PciDevice")
		delete(additionalProperties, "PciNode")
		delete(additionalProperties, "RegisteredDevice")
		delete(additionalProperties, "RunningFirmware")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGraphicsCardAllOf struct {
	value *GraphicsCardAllOf
	isSet bool
}

func (v NullableGraphicsCardAllOf) Get() *GraphicsCardAllOf {
	return v.value
}

func (v *NullableGraphicsCardAllOf) Set(val *GraphicsCardAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableGraphicsCardAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableGraphicsCardAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGraphicsCardAllOf(val *GraphicsCardAllOf) *NullableGraphicsCardAllOf {
	return &NullableGraphicsCardAllOf{value: val, isSet: true}
}

func (v NullableGraphicsCardAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGraphicsCardAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
