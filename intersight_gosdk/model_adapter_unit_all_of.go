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

// AdapterUnitAllOf Definition of the list of properties defined in 'adapter.Unit', excluding properties defined in parent classes.
type AdapterUnitAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Unique Identifier of an adapter Unit within a Rack Interface.
	AdapterId *string `json:"AdapterId,omitempty"`
	// Original Base Mac address of an adapter unit.
	BaseMacAddress *string `json:"BaseMacAddress,omitempty"`
	// Connectivity Status of adapter - A or B or AB.
	ConnectionStatus *string `json:"ConnectionStatus,omitempty"`
	// Cisco Integrated adapter or other type.
	Integrated *string  `json:"Integrated,omitempty"`
	OperReason []string `json:"OperReason,omitempty"`
	// Operational state of an adapter unit.
	OperState *string `json:"OperState,omitempty"`
	// Operability state of an adapter unit.
	Operability *string `json:"Operability,omitempty"`
	// Part number of an adapter unit.
	PartNumber *string `json:"PartNumber,omitempty"`
	// PCIe slot of the adapter in the server.
	PciSlot *string `json:"PciSlot,omitempty"`
	// Power state of an adapter unit.
	Power *string `json:"Power,omitempty"`
	// Thermal state of an adapter unit.
	Thermal *string `json:"Thermal,omitempty"`
	// Virtual Id of the adapter in the server.
	Vid                 *string                           `json:"Vid,omitempty"`
	AdapterUnitExpander *AdapterUnitExpanderRelationship  `json:"AdapterUnitExpander,omitempty"`
	ComputeBlade        *ComputeBladeRelationship         `json:"ComputeBlade,omitempty"`
	ComputeRackUnit     *ComputeRackUnitRelationship      `json:"ComputeRackUnit,omitempty"`
	Controller          *ManagementControllerRelationship `json:"Controller,omitempty"`
	// An array of relationships to adapterExtEthInterface resources.
	ExtEthIfs []AdapterExtEthInterfaceRelationship `json:"ExtEthIfs,omitempty"`
	// An array of relationships to adapterHostEthInterface resources.
	HostEthIfs []AdapterHostEthInterfaceRelationship `json:"HostEthIfs,omitempty"`
	// An array of relationships to adapterHostFcInterface resources.
	HostFcIfs []AdapterHostFcInterfaceRelationship `json:"HostFcIfs,omitempty"`
	// An array of relationships to adapterHostIscsiInterface resources.
	HostIscsiIfs         []AdapterHostIscsiInterfaceRelationship `json:"HostIscsiIfs,omitempty"`
	InventoryDeviceInfo  *InventoryDeviceInfoRelationship        `json:"InventoryDeviceInfo,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship    `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AdapterUnitAllOf AdapterUnitAllOf

// NewAdapterUnitAllOf instantiates a new AdapterUnitAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdapterUnitAllOf(classId string, objectType string) *AdapterUnitAllOf {
	this := AdapterUnitAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAdapterUnitAllOfWithDefaults instantiates a new AdapterUnitAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdapterUnitAllOfWithDefaults() *AdapterUnitAllOf {
	this := AdapterUnitAllOf{}
	var classId string = "adapter.Unit"
	this.ClassId = classId
	var objectType string = "adapter.Unit"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AdapterUnitAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AdapterUnitAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AdapterUnitAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AdapterUnitAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AdapterUnitAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AdapterUnitAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAdapterId returns the AdapterId field value if set, zero value otherwise.
func (o *AdapterUnitAllOf) GetAdapterId() string {
	if o == nil || o.AdapterId == nil {
		var ret string
		return ret
	}
	return *o.AdapterId
}

// GetAdapterIdOk returns a tuple with the AdapterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterUnitAllOf) GetAdapterIdOk() (*string, bool) {
	if o == nil || o.AdapterId == nil {
		return nil, false
	}
	return o.AdapterId, true
}

// HasAdapterId returns a boolean if a field has been set.
func (o *AdapterUnitAllOf) HasAdapterId() bool {
	if o != nil && o.AdapterId != nil {
		return true
	}

	return false
}

// SetAdapterId gets a reference to the given string and assigns it to the AdapterId field.
func (o *AdapterUnitAllOf) SetAdapterId(v string) {
	o.AdapterId = &v
}

// GetBaseMacAddress returns the BaseMacAddress field value if set, zero value otherwise.
func (o *AdapterUnitAllOf) GetBaseMacAddress() string {
	if o == nil || o.BaseMacAddress == nil {
		var ret string
		return ret
	}
	return *o.BaseMacAddress
}

// GetBaseMacAddressOk returns a tuple with the BaseMacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterUnitAllOf) GetBaseMacAddressOk() (*string, bool) {
	if o == nil || o.BaseMacAddress == nil {
		return nil, false
	}
	return o.BaseMacAddress, true
}

// HasBaseMacAddress returns a boolean if a field has been set.
func (o *AdapterUnitAllOf) HasBaseMacAddress() bool {
	if o != nil && o.BaseMacAddress != nil {
		return true
	}

	return false
}

// SetBaseMacAddress gets a reference to the given string and assigns it to the BaseMacAddress field.
func (o *AdapterUnitAllOf) SetBaseMacAddress(v string) {
	o.BaseMacAddress = &v
}

// GetConnectionStatus returns the ConnectionStatus field value if set, zero value otherwise.
func (o *AdapterUnitAllOf) GetConnectionStatus() string {
	if o == nil || o.ConnectionStatus == nil {
		var ret string
		return ret
	}
	return *o.ConnectionStatus
}

// GetConnectionStatusOk returns a tuple with the ConnectionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterUnitAllOf) GetConnectionStatusOk() (*string, bool) {
	if o == nil || o.ConnectionStatus == nil {
		return nil, false
	}
	return o.ConnectionStatus, true
}

// HasConnectionStatus returns a boolean if a field has been set.
func (o *AdapterUnitAllOf) HasConnectionStatus() bool {
	if o != nil && o.ConnectionStatus != nil {
		return true
	}

	return false
}

// SetConnectionStatus gets a reference to the given string and assigns it to the ConnectionStatus field.
func (o *AdapterUnitAllOf) SetConnectionStatus(v string) {
	o.ConnectionStatus = &v
}

// GetIntegrated returns the Integrated field value if set, zero value otherwise.
func (o *AdapterUnitAllOf) GetIntegrated() string {
	if o == nil || o.Integrated == nil {
		var ret string
		return ret
	}
	return *o.Integrated
}

// GetIntegratedOk returns a tuple with the Integrated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterUnitAllOf) GetIntegratedOk() (*string, bool) {
	if o == nil || o.Integrated == nil {
		return nil, false
	}
	return o.Integrated, true
}

// HasIntegrated returns a boolean if a field has been set.
func (o *AdapterUnitAllOf) HasIntegrated() bool {
	if o != nil && o.Integrated != nil {
		return true
	}

	return false
}

// SetIntegrated gets a reference to the given string and assigns it to the Integrated field.
func (o *AdapterUnitAllOf) SetIntegrated(v string) {
	o.Integrated = &v
}

// GetOperReason returns the OperReason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AdapterUnitAllOf) GetOperReason() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.OperReason
}

// GetOperReasonOk returns a tuple with the OperReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdapterUnitAllOf) GetOperReasonOk() ([]string, bool) {
	if o == nil || o.OperReason == nil {
		return nil, false
	}
	return o.OperReason, true
}

// HasOperReason returns a boolean if a field has been set.
func (o *AdapterUnitAllOf) HasOperReason() bool {
	if o != nil && o.OperReason != nil {
		return true
	}

	return false
}

// SetOperReason gets a reference to the given []string and assigns it to the OperReason field.
func (o *AdapterUnitAllOf) SetOperReason(v []string) {
	o.OperReason = v
}

// GetOperState returns the OperState field value if set, zero value otherwise.
func (o *AdapterUnitAllOf) GetOperState() string {
	if o == nil || o.OperState == nil {
		var ret string
		return ret
	}
	return *o.OperState
}

// GetOperStateOk returns a tuple with the OperState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterUnitAllOf) GetOperStateOk() (*string, bool) {
	if o == nil || o.OperState == nil {
		return nil, false
	}
	return o.OperState, true
}

// HasOperState returns a boolean if a field has been set.
func (o *AdapterUnitAllOf) HasOperState() bool {
	if o != nil && o.OperState != nil {
		return true
	}

	return false
}

// SetOperState gets a reference to the given string and assigns it to the OperState field.
func (o *AdapterUnitAllOf) SetOperState(v string) {
	o.OperState = &v
}

// GetOperability returns the Operability field value if set, zero value otherwise.
func (o *AdapterUnitAllOf) GetOperability() string {
	if o == nil || o.Operability == nil {
		var ret string
		return ret
	}
	return *o.Operability
}

// GetOperabilityOk returns a tuple with the Operability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterUnitAllOf) GetOperabilityOk() (*string, bool) {
	if o == nil || o.Operability == nil {
		return nil, false
	}
	return o.Operability, true
}

// HasOperability returns a boolean if a field has been set.
func (o *AdapterUnitAllOf) HasOperability() bool {
	if o != nil && o.Operability != nil {
		return true
	}

	return false
}

// SetOperability gets a reference to the given string and assigns it to the Operability field.
func (o *AdapterUnitAllOf) SetOperability(v string) {
	o.Operability = &v
}

// GetPartNumber returns the PartNumber field value if set, zero value otherwise.
func (o *AdapterUnitAllOf) GetPartNumber() string {
	if o == nil || o.PartNumber == nil {
		var ret string
		return ret
	}
	return *o.PartNumber
}

// GetPartNumberOk returns a tuple with the PartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterUnitAllOf) GetPartNumberOk() (*string, bool) {
	if o == nil || o.PartNumber == nil {
		return nil, false
	}
	return o.PartNumber, true
}

// HasPartNumber returns a boolean if a field has been set.
func (o *AdapterUnitAllOf) HasPartNumber() bool {
	if o != nil && o.PartNumber != nil {
		return true
	}

	return false
}

// SetPartNumber gets a reference to the given string and assigns it to the PartNumber field.
func (o *AdapterUnitAllOf) SetPartNumber(v string) {
	o.PartNumber = &v
}

// GetPciSlot returns the PciSlot field value if set, zero value otherwise.
func (o *AdapterUnitAllOf) GetPciSlot() string {
	if o == nil || o.PciSlot == nil {
		var ret string
		return ret
	}
	return *o.PciSlot
}

// GetPciSlotOk returns a tuple with the PciSlot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterUnitAllOf) GetPciSlotOk() (*string, bool) {
	if o == nil || o.PciSlot == nil {
		return nil, false
	}
	return o.PciSlot, true
}

// HasPciSlot returns a boolean if a field has been set.
func (o *AdapterUnitAllOf) HasPciSlot() bool {
	if o != nil && o.PciSlot != nil {
		return true
	}

	return false
}

// SetPciSlot gets a reference to the given string and assigns it to the PciSlot field.
func (o *AdapterUnitAllOf) SetPciSlot(v string) {
	o.PciSlot = &v
}

// GetPower returns the Power field value if set, zero value otherwise.
func (o *AdapterUnitAllOf) GetPower() string {
	if o == nil || o.Power == nil {
		var ret string
		return ret
	}
	return *o.Power
}

// GetPowerOk returns a tuple with the Power field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterUnitAllOf) GetPowerOk() (*string, bool) {
	if o == nil || o.Power == nil {
		return nil, false
	}
	return o.Power, true
}

// HasPower returns a boolean if a field has been set.
func (o *AdapterUnitAllOf) HasPower() bool {
	if o != nil && o.Power != nil {
		return true
	}

	return false
}

// SetPower gets a reference to the given string and assigns it to the Power field.
func (o *AdapterUnitAllOf) SetPower(v string) {
	o.Power = &v
}

// GetThermal returns the Thermal field value if set, zero value otherwise.
func (o *AdapterUnitAllOf) GetThermal() string {
	if o == nil || o.Thermal == nil {
		var ret string
		return ret
	}
	return *o.Thermal
}

// GetThermalOk returns a tuple with the Thermal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterUnitAllOf) GetThermalOk() (*string, bool) {
	if o == nil || o.Thermal == nil {
		return nil, false
	}
	return o.Thermal, true
}

// HasThermal returns a boolean if a field has been set.
func (o *AdapterUnitAllOf) HasThermal() bool {
	if o != nil && o.Thermal != nil {
		return true
	}

	return false
}

// SetThermal gets a reference to the given string and assigns it to the Thermal field.
func (o *AdapterUnitAllOf) SetThermal(v string) {
	o.Thermal = &v
}

// GetVid returns the Vid field value if set, zero value otherwise.
func (o *AdapterUnitAllOf) GetVid() string {
	if o == nil || o.Vid == nil {
		var ret string
		return ret
	}
	return *o.Vid
}

// GetVidOk returns a tuple with the Vid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterUnitAllOf) GetVidOk() (*string, bool) {
	if o == nil || o.Vid == nil {
		return nil, false
	}
	return o.Vid, true
}

// HasVid returns a boolean if a field has been set.
func (o *AdapterUnitAllOf) HasVid() bool {
	if o != nil && o.Vid != nil {
		return true
	}

	return false
}

// SetVid gets a reference to the given string and assigns it to the Vid field.
func (o *AdapterUnitAllOf) SetVid(v string) {
	o.Vid = &v
}

// GetAdapterUnitExpander returns the AdapterUnitExpander field value if set, zero value otherwise.
func (o *AdapterUnitAllOf) GetAdapterUnitExpander() AdapterUnitExpanderRelationship {
	if o == nil || o.AdapterUnitExpander == nil {
		var ret AdapterUnitExpanderRelationship
		return ret
	}
	return *o.AdapterUnitExpander
}

// GetAdapterUnitExpanderOk returns a tuple with the AdapterUnitExpander field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterUnitAllOf) GetAdapterUnitExpanderOk() (*AdapterUnitExpanderRelationship, bool) {
	if o == nil || o.AdapterUnitExpander == nil {
		return nil, false
	}
	return o.AdapterUnitExpander, true
}

// HasAdapterUnitExpander returns a boolean if a field has been set.
func (o *AdapterUnitAllOf) HasAdapterUnitExpander() bool {
	if o != nil && o.AdapterUnitExpander != nil {
		return true
	}

	return false
}

// SetAdapterUnitExpander gets a reference to the given AdapterUnitExpanderRelationship and assigns it to the AdapterUnitExpander field.
func (o *AdapterUnitAllOf) SetAdapterUnitExpander(v AdapterUnitExpanderRelationship) {
	o.AdapterUnitExpander = &v
}

// GetComputeBlade returns the ComputeBlade field value if set, zero value otherwise.
func (o *AdapterUnitAllOf) GetComputeBlade() ComputeBladeRelationship {
	if o == nil || o.ComputeBlade == nil {
		var ret ComputeBladeRelationship
		return ret
	}
	return *o.ComputeBlade
}

// GetComputeBladeOk returns a tuple with the ComputeBlade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterUnitAllOf) GetComputeBladeOk() (*ComputeBladeRelationship, bool) {
	if o == nil || o.ComputeBlade == nil {
		return nil, false
	}
	return o.ComputeBlade, true
}

// HasComputeBlade returns a boolean if a field has been set.
func (o *AdapterUnitAllOf) HasComputeBlade() bool {
	if o != nil && o.ComputeBlade != nil {
		return true
	}

	return false
}

// SetComputeBlade gets a reference to the given ComputeBladeRelationship and assigns it to the ComputeBlade field.
func (o *AdapterUnitAllOf) SetComputeBlade(v ComputeBladeRelationship) {
	o.ComputeBlade = &v
}

// GetComputeRackUnit returns the ComputeRackUnit field value if set, zero value otherwise.
func (o *AdapterUnitAllOf) GetComputeRackUnit() ComputeRackUnitRelationship {
	if o == nil || o.ComputeRackUnit == nil {
		var ret ComputeRackUnitRelationship
		return ret
	}
	return *o.ComputeRackUnit
}

// GetComputeRackUnitOk returns a tuple with the ComputeRackUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterUnitAllOf) GetComputeRackUnitOk() (*ComputeRackUnitRelationship, bool) {
	if o == nil || o.ComputeRackUnit == nil {
		return nil, false
	}
	return o.ComputeRackUnit, true
}

// HasComputeRackUnit returns a boolean if a field has been set.
func (o *AdapterUnitAllOf) HasComputeRackUnit() bool {
	if o != nil && o.ComputeRackUnit != nil {
		return true
	}

	return false
}

// SetComputeRackUnit gets a reference to the given ComputeRackUnitRelationship and assigns it to the ComputeRackUnit field.
func (o *AdapterUnitAllOf) SetComputeRackUnit(v ComputeRackUnitRelationship) {
	o.ComputeRackUnit = &v
}

// GetController returns the Controller field value if set, zero value otherwise.
func (o *AdapterUnitAllOf) GetController() ManagementControllerRelationship {
	if o == nil || o.Controller == nil {
		var ret ManagementControllerRelationship
		return ret
	}
	return *o.Controller
}

// GetControllerOk returns a tuple with the Controller field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterUnitAllOf) GetControllerOk() (*ManagementControllerRelationship, bool) {
	if o == nil || o.Controller == nil {
		return nil, false
	}
	return o.Controller, true
}

// HasController returns a boolean if a field has been set.
func (o *AdapterUnitAllOf) HasController() bool {
	if o != nil && o.Controller != nil {
		return true
	}

	return false
}

// SetController gets a reference to the given ManagementControllerRelationship and assigns it to the Controller field.
func (o *AdapterUnitAllOf) SetController(v ManagementControllerRelationship) {
	o.Controller = &v
}

// GetExtEthIfs returns the ExtEthIfs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AdapterUnitAllOf) GetExtEthIfs() []AdapterExtEthInterfaceRelationship {
	if o == nil {
		var ret []AdapterExtEthInterfaceRelationship
		return ret
	}
	return o.ExtEthIfs
}

// GetExtEthIfsOk returns a tuple with the ExtEthIfs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdapterUnitAllOf) GetExtEthIfsOk() ([]AdapterExtEthInterfaceRelationship, bool) {
	if o == nil || o.ExtEthIfs == nil {
		return nil, false
	}
	return o.ExtEthIfs, true
}

// HasExtEthIfs returns a boolean if a field has been set.
func (o *AdapterUnitAllOf) HasExtEthIfs() bool {
	if o != nil && o.ExtEthIfs != nil {
		return true
	}

	return false
}

// SetExtEthIfs gets a reference to the given []AdapterExtEthInterfaceRelationship and assigns it to the ExtEthIfs field.
func (o *AdapterUnitAllOf) SetExtEthIfs(v []AdapterExtEthInterfaceRelationship) {
	o.ExtEthIfs = v
}

// GetHostEthIfs returns the HostEthIfs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AdapterUnitAllOf) GetHostEthIfs() []AdapterHostEthInterfaceRelationship {
	if o == nil {
		var ret []AdapterHostEthInterfaceRelationship
		return ret
	}
	return o.HostEthIfs
}

// GetHostEthIfsOk returns a tuple with the HostEthIfs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdapterUnitAllOf) GetHostEthIfsOk() ([]AdapterHostEthInterfaceRelationship, bool) {
	if o == nil || o.HostEthIfs == nil {
		return nil, false
	}
	return o.HostEthIfs, true
}

// HasHostEthIfs returns a boolean if a field has been set.
func (o *AdapterUnitAllOf) HasHostEthIfs() bool {
	if o != nil && o.HostEthIfs != nil {
		return true
	}

	return false
}

// SetHostEthIfs gets a reference to the given []AdapterHostEthInterfaceRelationship and assigns it to the HostEthIfs field.
func (o *AdapterUnitAllOf) SetHostEthIfs(v []AdapterHostEthInterfaceRelationship) {
	o.HostEthIfs = v
}

// GetHostFcIfs returns the HostFcIfs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AdapterUnitAllOf) GetHostFcIfs() []AdapterHostFcInterfaceRelationship {
	if o == nil {
		var ret []AdapterHostFcInterfaceRelationship
		return ret
	}
	return o.HostFcIfs
}

// GetHostFcIfsOk returns a tuple with the HostFcIfs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdapterUnitAllOf) GetHostFcIfsOk() ([]AdapterHostFcInterfaceRelationship, bool) {
	if o == nil || o.HostFcIfs == nil {
		return nil, false
	}
	return o.HostFcIfs, true
}

// HasHostFcIfs returns a boolean if a field has been set.
func (o *AdapterUnitAllOf) HasHostFcIfs() bool {
	if o != nil && o.HostFcIfs != nil {
		return true
	}

	return false
}

// SetHostFcIfs gets a reference to the given []AdapterHostFcInterfaceRelationship and assigns it to the HostFcIfs field.
func (o *AdapterUnitAllOf) SetHostFcIfs(v []AdapterHostFcInterfaceRelationship) {
	o.HostFcIfs = v
}

// GetHostIscsiIfs returns the HostIscsiIfs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AdapterUnitAllOf) GetHostIscsiIfs() []AdapterHostIscsiInterfaceRelationship {
	if o == nil {
		var ret []AdapterHostIscsiInterfaceRelationship
		return ret
	}
	return o.HostIscsiIfs
}

// GetHostIscsiIfsOk returns a tuple with the HostIscsiIfs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdapterUnitAllOf) GetHostIscsiIfsOk() ([]AdapterHostIscsiInterfaceRelationship, bool) {
	if o == nil || o.HostIscsiIfs == nil {
		return nil, false
	}
	return o.HostIscsiIfs, true
}

// HasHostIscsiIfs returns a boolean if a field has been set.
func (o *AdapterUnitAllOf) HasHostIscsiIfs() bool {
	if o != nil && o.HostIscsiIfs != nil {
		return true
	}

	return false
}

// SetHostIscsiIfs gets a reference to the given []AdapterHostIscsiInterfaceRelationship and assigns it to the HostIscsiIfs field.
func (o *AdapterUnitAllOf) SetHostIscsiIfs(v []AdapterHostIscsiInterfaceRelationship) {
	o.HostIscsiIfs = v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *AdapterUnitAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterUnitAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *AdapterUnitAllOf) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *AdapterUnitAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *AdapterUnitAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterUnitAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *AdapterUnitAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *AdapterUnitAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o AdapterUnitAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AdapterId != nil {
		toSerialize["AdapterId"] = o.AdapterId
	}
	if o.BaseMacAddress != nil {
		toSerialize["BaseMacAddress"] = o.BaseMacAddress
	}
	if o.ConnectionStatus != nil {
		toSerialize["ConnectionStatus"] = o.ConnectionStatus
	}
	if o.Integrated != nil {
		toSerialize["Integrated"] = o.Integrated
	}
	if o.OperReason != nil {
		toSerialize["OperReason"] = o.OperReason
	}
	if o.OperState != nil {
		toSerialize["OperState"] = o.OperState
	}
	if o.Operability != nil {
		toSerialize["Operability"] = o.Operability
	}
	if o.PartNumber != nil {
		toSerialize["PartNumber"] = o.PartNumber
	}
	if o.PciSlot != nil {
		toSerialize["PciSlot"] = o.PciSlot
	}
	if o.Power != nil {
		toSerialize["Power"] = o.Power
	}
	if o.Thermal != nil {
		toSerialize["Thermal"] = o.Thermal
	}
	if o.Vid != nil {
		toSerialize["Vid"] = o.Vid
	}
	if o.AdapterUnitExpander != nil {
		toSerialize["AdapterUnitExpander"] = o.AdapterUnitExpander
	}
	if o.ComputeBlade != nil {
		toSerialize["ComputeBlade"] = o.ComputeBlade
	}
	if o.ComputeRackUnit != nil {
		toSerialize["ComputeRackUnit"] = o.ComputeRackUnit
	}
	if o.Controller != nil {
		toSerialize["Controller"] = o.Controller
	}
	if o.ExtEthIfs != nil {
		toSerialize["ExtEthIfs"] = o.ExtEthIfs
	}
	if o.HostEthIfs != nil {
		toSerialize["HostEthIfs"] = o.HostEthIfs
	}
	if o.HostFcIfs != nil {
		toSerialize["HostFcIfs"] = o.HostFcIfs
	}
	if o.HostIscsiIfs != nil {
		toSerialize["HostIscsiIfs"] = o.HostIscsiIfs
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

func (o *AdapterUnitAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varAdapterUnitAllOf := _AdapterUnitAllOf{}

	if err = json.Unmarshal(bytes, &varAdapterUnitAllOf); err == nil {
		*o = AdapterUnitAllOf(varAdapterUnitAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AdapterId")
		delete(additionalProperties, "BaseMacAddress")
		delete(additionalProperties, "ConnectionStatus")
		delete(additionalProperties, "Integrated")
		delete(additionalProperties, "OperReason")
		delete(additionalProperties, "OperState")
		delete(additionalProperties, "Operability")
		delete(additionalProperties, "PartNumber")
		delete(additionalProperties, "PciSlot")
		delete(additionalProperties, "Power")
		delete(additionalProperties, "Thermal")
		delete(additionalProperties, "Vid")
		delete(additionalProperties, "AdapterUnitExpander")
		delete(additionalProperties, "ComputeBlade")
		delete(additionalProperties, "ComputeRackUnit")
		delete(additionalProperties, "Controller")
		delete(additionalProperties, "ExtEthIfs")
		delete(additionalProperties, "HostEthIfs")
		delete(additionalProperties, "HostFcIfs")
		delete(additionalProperties, "HostIscsiIfs")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAdapterUnitAllOf struct {
	value *AdapterUnitAllOf
	isSet bool
}

func (v NullableAdapterUnitAllOf) Get() *AdapterUnitAllOf {
	return v.value
}

func (v *NullableAdapterUnitAllOf) Set(val *AdapterUnitAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAdapterUnitAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAdapterUnitAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdapterUnitAllOf(val *AdapterUnitAllOf) *NullableAdapterUnitAllOf {
	return &NullableAdapterUnitAllOf{value: val, isSet: true}
}

func (v NullableAdapterUnitAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdapterUnitAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
