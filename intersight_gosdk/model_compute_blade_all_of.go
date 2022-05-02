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
)

// ComputeBladeAllOf Definition of the list of properties defined in 'compute.Blade', excluding properties defined in parent classes.
type ComputeBladeAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The id of the chassis that the blade is discovered in.
	ChassisId *string `json:"ChassisId,omitempty"`
	// The mode of the server that determines it is scaled.
	ScaledMode *string `json:"ScaledMode,omitempty"`
	// The slot number in the chassis that the blade is discovered in.
	SlotId *int64 `json:"SlotId,omitempty"`
	// An array of relationships to adapterUnit resources.
	Adapters          []AdapterUnitRelationship      `json:"Adapters,omitempty"`
	BiosBootmode      *BiosBootModeRelationship      `json:"BiosBootmode,omitempty"`
	BiosTokenSettings *BiosTokenSettingsRelationship `json:"BiosTokenSettings,omitempty"`
	// An array of relationships to biosUnit resources.
	BiosUnits                          []BiosUnitRelationship                          `json:"BiosUnits,omitempty"`
	BiosVfSelectMemoryRasConfiguration *BiosVfSelectMemoryRasConfigurationRelationship `json:"BiosVfSelectMemoryRasConfiguration,omitempty"`
	Bmc                                *ManagementControllerRelationship               `json:"Bmc,omitempty"`
	Board                              *ComputeBoardRelationship                       `json:"Board,omitempty"`
	BootDeviceBootmode                 *BootDeviceBootModeRelationship                 `json:"BootDeviceBootmode,omitempty"`
	EquipmentChassis                   *EquipmentChassisRelationship                   `json:"EquipmentChassis,omitempty"`
	// An array of relationships to equipmentIoExpander resources.
	EquipmentIoExpanders []EquipmentIoExpanderRelationship `json:"EquipmentIoExpanders,omitempty"`
	// An array of relationships to inventoryGenericInventoryHolder resources.
	GenericInventoryHolders []InventoryGenericInventoryHolderRelationship `json:"GenericInventoryHolders,omitempty"`
	// An array of relationships to graphicsCard resources.
	GraphicsCards       []GraphicsCardRelationship       `json:"GraphicsCards,omitempty"`
	InventoryDeviceInfo *InventoryDeviceInfoRelationship `json:"InventoryDeviceInfo,omitempty"`
	LocatorLed          *EquipmentLocatorLedRelationship `json:"LocatorLed,omitempty"`
	// An array of relationships to memoryArray resources.
	MemoryArrays []MemoryArrayRelationship `json:"MemoryArrays,omitempty"`
	// An array of relationships to pciDevice resources.
	PciDevices []PciDeviceRelationship `json:"PciDevices,omitempty"`
	// An array of relationships to processorUnit resources.
	Processors       []ProcessorUnitRelationship          `json:"Processors,omitempty"`
	RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	// An array of relationships to storageController resources.
	StorageControllers []StorageControllerRelationship `json:"StorageControllers,omitempty"`
	// An array of relationships to storageEnclosure resources.
	StorageEnclosures    []StorageEnclosureRelationship `json:"StorageEnclosures,omitempty"`
	TopSystem            *TopSystemRelationship         `json:"TopSystem,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ComputeBladeAllOf ComputeBladeAllOf

// NewComputeBladeAllOf instantiates a new ComputeBladeAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputeBladeAllOf(classId string, objectType string) *ComputeBladeAllOf {
	this := ComputeBladeAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewComputeBladeAllOfWithDefaults instantiates a new ComputeBladeAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputeBladeAllOfWithDefaults() *ComputeBladeAllOf {
	this := ComputeBladeAllOf{}
	var classId string = "compute.Blade"
	this.ClassId = classId
	var objectType string = "compute.Blade"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ComputeBladeAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ComputeBladeAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ComputeBladeAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ComputeBladeAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ComputeBladeAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ComputeBladeAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetChassisId returns the ChassisId field value if set, zero value otherwise.
func (o *ComputeBladeAllOf) GetChassisId() string {
	if o == nil || o.ChassisId == nil {
		var ret string
		return ret
	}
	return *o.ChassisId
}

// GetChassisIdOk returns a tuple with the ChassisId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeBladeAllOf) GetChassisIdOk() (*string, bool) {
	if o == nil || o.ChassisId == nil {
		return nil, false
	}
	return o.ChassisId, true
}

// HasChassisId returns a boolean if a field has been set.
func (o *ComputeBladeAllOf) HasChassisId() bool {
	if o != nil && o.ChassisId != nil {
		return true
	}

	return false
}

// SetChassisId gets a reference to the given string and assigns it to the ChassisId field.
func (o *ComputeBladeAllOf) SetChassisId(v string) {
	o.ChassisId = &v
}

// GetScaledMode returns the ScaledMode field value if set, zero value otherwise.
func (o *ComputeBladeAllOf) GetScaledMode() string {
	if o == nil || o.ScaledMode == nil {
		var ret string
		return ret
	}
	return *o.ScaledMode
}

// GetScaledModeOk returns a tuple with the ScaledMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeBladeAllOf) GetScaledModeOk() (*string, bool) {
	if o == nil || o.ScaledMode == nil {
		return nil, false
	}
	return o.ScaledMode, true
}

// HasScaledMode returns a boolean if a field has been set.
func (o *ComputeBladeAllOf) HasScaledMode() bool {
	if o != nil && o.ScaledMode != nil {
		return true
	}

	return false
}

// SetScaledMode gets a reference to the given string and assigns it to the ScaledMode field.
func (o *ComputeBladeAllOf) SetScaledMode(v string) {
	o.ScaledMode = &v
}

// GetSlotId returns the SlotId field value if set, zero value otherwise.
func (o *ComputeBladeAllOf) GetSlotId() int64 {
	if o == nil || o.SlotId == nil {
		var ret int64
		return ret
	}
	return *o.SlotId
}

// GetSlotIdOk returns a tuple with the SlotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeBladeAllOf) GetSlotIdOk() (*int64, bool) {
	if o == nil || o.SlotId == nil {
		return nil, false
	}
	return o.SlotId, true
}

// HasSlotId returns a boolean if a field has been set.
func (o *ComputeBladeAllOf) HasSlotId() bool {
	if o != nil && o.SlotId != nil {
		return true
	}

	return false
}

// SetSlotId gets a reference to the given int64 and assigns it to the SlotId field.
func (o *ComputeBladeAllOf) SetSlotId(v int64) {
	o.SlotId = &v
}

// GetAdapters returns the Adapters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComputeBladeAllOf) GetAdapters() []AdapterUnitRelationship {
	if o == nil {
		var ret []AdapterUnitRelationship
		return ret
	}
	return o.Adapters
}

// GetAdaptersOk returns a tuple with the Adapters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComputeBladeAllOf) GetAdaptersOk() ([]AdapterUnitRelationship, bool) {
	if o == nil || o.Adapters == nil {
		return nil, false
	}
	return o.Adapters, true
}

// HasAdapters returns a boolean if a field has been set.
func (o *ComputeBladeAllOf) HasAdapters() bool {
	if o != nil && o.Adapters != nil {
		return true
	}

	return false
}

// SetAdapters gets a reference to the given []AdapterUnitRelationship and assigns it to the Adapters field.
func (o *ComputeBladeAllOf) SetAdapters(v []AdapterUnitRelationship) {
	o.Adapters = v
}

// GetBiosBootmode returns the BiosBootmode field value if set, zero value otherwise.
func (o *ComputeBladeAllOf) GetBiosBootmode() BiosBootModeRelationship {
	if o == nil || o.BiosBootmode == nil {
		var ret BiosBootModeRelationship
		return ret
	}
	return *o.BiosBootmode
}

// GetBiosBootmodeOk returns a tuple with the BiosBootmode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeBladeAllOf) GetBiosBootmodeOk() (*BiosBootModeRelationship, bool) {
	if o == nil || o.BiosBootmode == nil {
		return nil, false
	}
	return o.BiosBootmode, true
}

// HasBiosBootmode returns a boolean if a field has been set.
func (o *ComputeBladeAllOf) HasBiosBootmode() bool {
	if o != nil && o.BiosBootmode != nil {
		return true
	}

	return false
}

// SetBiosBootmode gets a reference to the given BiosBootModeRelationship and assigns it to the BiosBootmode field.
func (o *ComputeBladeAllOf) SetBiosBootmode(v BiosBootModeRelationship) {
	o.BiosBootmode = &v
}

// GetBiosTokenSettings returns the BiosTokenSettings field value if set, zero value otherwise.
func (o *ComputeBladeAllOf) GetBiosTokenSettings() BiosTokenSettingsRelationship {
	if o == nil || o.BiosTokenSettings == nil {
		var ret BiosTokenSettingsRelationship
		return ret
	}
	return *o.BiosTokenSettings
}

// GetBiosTokenSettingsOk returns a tuple with the BiosTokenSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeBladeAllOf) GetBiosTokenSettingsOk() (*BiosTokenSettingsRelationship, bool) {
	if o == nil || o.BiosTokenSettings == nil {
		return nil, false
	}
	return o.BiosTokenSettings, true
}

// HasBiosTokenSettings returns a boolean if a field has been set.
func (o *ComputeBladeAllOf) HasBiosTokenSettings() bool {
	if o != nil && o.BiosTokenSettings != nil {
		return true
	}

	return false
}

// SetBiosTokenSettings gets a reference to the given BiosTokenSettingsRelationship and assigns it to the BiosTokenSettings field.
func (o *ComputeBladeAllOf) SetBiosTokenSettings(v BiosTokenSettingsRelationship) {
	o.BiosTokenSettings = &v
}

// GetBiosUnits returns the BiosUnits field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComputeBladeAllOf) GetBiosUnits() []BiosUnitRelationship {
	if o == nil {
		var ret []BiosUnitRelationship
		return ret
	}
	return o.BiosUnits
}

// GetBiosUnitsOk returns a tuple with the BiosUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComputeBladeAllOf) GetBiosUnitsOk() ([]BiosUnitRelationship, bool) {
	if o == nil || o.BiosUnits == nil {
		return nil, false
	}
	return o.BiosUnits, true
}

// HasBiosUnits returns a boolean if a field has been set.
func (o *ComputeBladeAllOf) HasBiosUnits() bool {
	if o != nil && o.BiosUnits != nil {
		return true
	}

	return false
}

// SetBiosUnits gets a reference to the given []BiosUnitRelationship and assigns it to the BiosUnits field.
func (o *ComputeBladeAllOf) SetBiosUnits(v []BiosUnitRelationship) {
	o.BiosUnits = v
}

// GetBiosVfSelectMemoryRasConfiguration returns the BiosVfSelectMemoryRasConfiguration field value if set, zero value otherwise.
func (o *ComputeBladeAllOf) GetBiosVfSelectMemoryRasConfiguration() BiosVfSelectMemoryRasConfigurationRelationship {
	if o == nil || o.BiosVfSelectMemoryRasConfiguration == nil {
		var ret BiosVfSelectMemoryRasConfigurationRelationship
		return ret
	}
	return *o.BiosVfSelectMemoryRasConfiguration
}

// GetBiosVfSelectMemoryRasConfigurationOk returns a tuple with the BiosVfSelectMemoryRasConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeBladeAllOf) GetBiosVfSelectMemoryRasConfigurationOk() (*BiosVfSelectMemoryRasConfigurationRelationship, bool) {
	if o == nil || o.BiosVfSelectMemoryRasConfiguration == nil {
		return nil, false
	}
	return o.BiosVfSelectMemoryRasConfiguration, true
}

// HasBiosVfSelectMemoryRasConfiguration returns a boolean if a field has been set.
func (o *ComputeBladeAllOf) HasBiosVfSelectMemoryRasConfiguration() bool {
	if o != nil && o.BiosVfSelectMemoryRasConfiguration != nil {
		return true
	}

	return false
}

// SetBiosVfSelectMemoryRasConfiguration gets a reference to the given BiosVfSelectMemoryRasConfigurationRelationship and assigns it to the BiosVfSelectMemoryRasConfiguration field.
func (o *ComputeBladeAllOf) SetBiosVfSelectMemoryRasConfiguration(v BiosVfSelectMemoryRasConfigurationRelationship) {
	o.BiosVfSelectMemoryRasConfiguration = &v
}

// GetBmc returns the Bmc field value if set, zero value otherwise.
func (o *ComputeBladeAllOf) GetBmc() ManagementControllerRelationship {
	if o == nil || o.Bmc == nil {
		var ret ManagementControllerRelationship
		return ret
	}
	return *o.Bmc
}

// GetBmcOk returns a tuple with the Bmc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeBladeAllOf) GetBmcOk() (*ManagementControllerRelationship, bool) {
	if o == nil || o.Bmc == nil {
		return nil, false
	}
	return o.Bmc, true
}

// HasBmc returns a boolean if a field has been set.
func (o *ComputeBladeAllOf) HasBmc() bool {
	if o != nil && o.Bmc != nil {
		return true
	}

	return false
}

// SetBmc gets a reference to the given ManagementControllerRelationship and assigns it to the Bmc field.
func (o *ComputeBladeAllOf) SetBmc(v ManagementControllerRelationship) {
	o.Bmc = &v
}

// GetBoard returns the Board field value if set, zero value otherwise.
func (o *ComputeBladeAllOf) GetBoard() ComputeBoardRelationship {
	if o == nil || o.Board == nil {
		var ret ComputeBoardRelationship
		return ret
	}
	return *o.Board
}

// GetBoardOk returns a tuple with the Board field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeBladeAllOf) GetBoardOk() (*ComputeBoardRelationship, bool) {
	if o == nil || o.Board == nil {
		return nil, false
	}
	return o.Board, true
}

// HasBoard returns a boolean if a field has been set.
func (o *ComputeBladeAllOf) HasBoard() bool {
	if o != nil && o.Board != nil {
		return true
	}

	return false
}

// SetBoard gets a reference to the given ComputeBoardRelationship and assigns it to the Board field.
func (o *ComputeBladeAllOf) SetBoard(v ComputeBoardRelationship) {
	o.Board = &v
}

// GetBootDeviceBootmode returns the BootDeviceBootmode field value if set, zero value otherwise.
func (o *ComputeBladeAllOf) GetBootDeviceBootmode() BootDeviceBootModeRelationship {
	if o == nil || o.BootDeviceBootmode == nil {
		var ret BootDeviceBootModeRelationship
		return ret
	}
	return *o.BootDeviceBootmode
}

// GetBootDeviceBootmodeOk returns a tuple with the BootDeviceBootmode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeBladeAllOf) GetBootDeviceBootmodeOk() (*BootDeviceBootModeRelationship, bool) {
	if o == nil || o.BootDeviceBootmode == nil {
		return nil, false
	}
	return o.BootDeviceBootmode, true
}

// HasBootDeviceBootmode returns a boolean if a field has been set.
func (o *ComputeBladeAllOf) HasBootDeviceBootmode() bool {
	if o != nil && o.BootDeviceBootmode != nil {
		return true
	}

	return false
}

// SetBootDeviceBootmode gets a reference to the given BootDeviceBootModeRelationship and assigns it to the BootDeviceBootmode field.
func (o *ComputeBladeAllOf) SetBootDeviceBootmode(v BootDeviceBootModeRelationship) {
	o.BootDeviceBootmode = &v
}

// GetEquipmentChassis returns the EquipmentChassis field value if set, zero value otherwise.
func (o *ComputeBladeAllOf) GetEquipmentChassis() EquipmentChassisRelationship {
	if o == nil || o.EquipmentChassis == nil {
		var ret EquipmentChassisRelationship
		return ret
	}
	return *o.EquipmentChassis
}

// GetEquipmentChassisOk returns a tuple with the EquipmentChassis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeBladeAllOf) GetEquipmentChassisOk() (*EquipmentChassisRelationship, bool) {
	if o == nil || o.EquipmentChassis == nil {
		return nil, false
	}
	return o.EquipmentChassis, true
}

// HasEquipmentChassis returns a boolean if a field has been set.
func (o *ComputeBladeAllOf) HasEquipmentChassis() bool {
	if o != nil && o.EquipmentChassis != nil {
		return true
	}

	return false
}

// SetEquipmentChassis gets a reference to the given EquipmentChassisRelationship and assigns it to the EquipmentChassis field.
func (o *ComputeBladeAllOf) SetEquipmentChassis(v EquipmentChassisRelationship) {
	o.EquipmentChassis = &v
}

// GetEquipmentIoExpanders returns the EquipmentIoExpanders field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComputeBladeAllOf) GetEquipmentIoExpanders() []EquipmentIoExpanderRelationship {
	if o == nil {
		var ret []EquipmentIoExpanderRelationship
		return ret
	}
	return o.EquipmentIoExpanders
}

// GetEquipmentIoExpandersOk returns a tuple with the EquipmentIoExpanders field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComputeBladeAllOf) GetEquipmentIoExpandersOk() ([]EquipmentIoExpanderRelationship, bool) {
	if o == nil || o.EquipmentIoExpanders == nil {
		return nil, false
	}
	return o.EquipmentIoExpanders, true
}

// HasEquipmentIoExpanders returns a boolean if a field has been set.
func (o *ComputeBladeAllOf) HasEquipmentIoExpanders() bool {
	if o != nil && o.EquipmentIoExpanders != nil {
		return true
	}

	return false
}

// SetEquipmentIoExpanders gets a reference to the given []EquipmentIoExpanderRelationship and assigns it to the EquipmentIoExpanders field.
func (o *ComputeBladeAllOf) SetEquipmentIoExpanders(v []EquipmentIoExpanderRelationship) {
	o.EquipmentIoExpanders = v
}

// GetGenericInventoryHolders returns the GenericInventoryHolders field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComputeBladeAllOf) GetGenericInventoryHolders() []InventoryGenericInventoryHolderRelationship {
	if o == nil {
		var ret []InventoryGenericInventoryHolderRelationship
		return ret
	}
	return o.GenericInventoryHolders
}

// GetGenericInventoryHoldersOk returns a tuple with the GenericInventoryHolders field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComputeBladeAllOf) GetGenericInventoryHoldersOk() ([]InventoryGenericInventoryHolderRelationship, bool) {
	if o == nil || o.GenericInventoryHolders == nil {
		return nil, false
	}
	return o.GenericInventoryHolders, true
}

// HasGenericInventoryHolders returns a boolean if a field has been set.
func (o *ComputeBladeAllOf) HasGenericInventoryHolders() bool {
	if o != nil && o.GenericInventoryHolders != nil {
		return true
	}

	return false
}

// SetGenericInventoryHolders gets a reference to the given []InventoryGenericInventoryHolderRelationship and assigns it to the GenericInventoryHolders field.
func (o *ComputeBladeAllOf) SetGenericInventoryHolders(v []InventoryGenericInventoryHolderRelationship) {
	o.GenericInventoryHolders = v
}

// GetGraphicsCards returns the GraphicsCards field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComputeBladeAllOf) GetGraphicsCards() []GraphicsCardRelationship {
	if o == nil {
		var ret []GraphicsCardRelationship
		return ret
	}
	return o.GraphicsCards
}

// GetGraphicsCardsOk returns a tuple with the GraphicsCards field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComputeBladeAllOf) GetGraphicsCardsOk() ([]GraphicsCardRelationship, bool) {
	if o == nil || o.GraphicsCards == nil {
		return nil, false
	}
	return o.GraphicsCards, true
}

// HasGraphicsCards returns a boolean if a field has been set.
func (o *ComputeBladeAllOf) HasGraphicsCards() bool {
	if o != nil && o.GraphicsCards != nil {
		return true
	}

	return false
}

// SetGraphicsCards gets a reference to the given []GraphicsCardRelationship and assigns it to the GraphicsCards field.
func (o *ComputeBladeAllOf) SetGraphicsCards(v []GraphicsCardRelationship) {
	o.GraphicsCards = v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *ComputeBladeAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeBladeAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *ComputeBladeAllOf) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *ComputeBladeAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetLocatorLed returns the LocatorLed field value if set, zero value otherwise.
func (o *ComputeBladeAllOf) GetLocatorLed() EquipmentLocatorLedRelationship {
	if o == nil || o.LocatorLed == nil {
		var ret EquipmentLocatorLedRelationship
		return ret
	}
	return *o.LocatorLed
}

// GetLocatorLedOk returns a tuple with the LocatorLed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeBladeAllOf) GetLocatorLedOk() (*EquipmentLocatorLedRelationship, bool) {
	if o == nil || o.LocatorLed == nil {
		return nil, false
	}
	return o.LocatorLed, true
}

// HasLocatorLed returns a boolean if a field has been set.
func (o *ComputeBladeAllOf) HasLocatorLed() bool {
	if o != nil && o.LocatorLed != nil {
		return true
	}

	return false
}

// SetLocatorLed gets a reference to the given EquipmentLocatorLedRelationship and assigns it to the LocatorLed field.
func (o *ComputeBladeAllOf) SetLocatorLed(v EquipmentLocatorLedRelationship) {
	o.LocatorLed = &v
}

// GetMemoryArrays returns the MemoryArrays field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComputeBladeAllOf) GetMemoryArrays() []MemoryArrayRelationship {
	if o == nil {
		var ret []MemoryArrayRelationship
		return ret
	}
	return o.MemoryArrays
}

// GetMemoryArraysOk returns a tuple with the MemoryArrays field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComputeBladeAllOf) GetMemoryArraysOk() ([]MemoryArrayRelationship, bool) {
	if o == nil || o.MemoryArrays == nil {
		return nil, false
	}
	return o.MemoryArrays, true
}

// HasMemoryArrays returns a boolean if a field has been set.
func (o *ComputeBladeAllOf) HasMemoryArrays() bool {
	if o != nil && o.MemoryArrays != nil {
		return true
	}

	return false
}

// SetMemoryArrays gets a reference to the given []MemoryArrayRelationship and assigns it to the MemoryArrays field.
func (o *ComputeBladeAllOf) SetMemoryArrays(v []MemoryArrayRelationship) {
	o.MemoryArrays = v
}

// GetPciDevices returns the PciDevices field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComputeBladeAllOf) GetPciDevices() []PciDeviceRelationship {
	if o == nil {
		var ret []PciDeviceRelationship
		return ret
	}
	return o.PciDevices
}

// GetPciDevicesOk returns a tuple with the PciDevices field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComputeBladeAllOf) GetPciDevicesOk() ([]PciDeviceRelationship, bool) {
	if o == nil || o.PciDevices == nil {
		return nil, false
	}
	return o.PciDevices, true
}

// HasPciDevices returns a boolean if a field has been set.
func (o *ComputeBladeAllOf) HasPciDevices() bool {
	if o != nil && o.PciDevices != nil {
		return true
	}

	return false
}

// SetPciDevices gets a reference to the given []PciDeviceRelationship and assigns it to the PciDevices field.
func (o *ComputeBladeAllOf) SetPciDevices(v []PciDeviceRelationship) {
	o.PciDevices = v
}

// GetProcessors returns the Processors field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComputeBladeAllOf) GetProcessors() []ProcessorUnitRelationship {
	if o == nil {
		var ret []ProcessorUnitRelationship
		return ret
	}
	return o.Processors
}

// GetProcessorsOk returns a tuple with the Processors field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComputeBladeAllOf) GetProcessorsOk() ([]ProcessorUnitRelationship, bool) {
	if o == nil || o.Processors == nil {
		return nil, false
	}
	return o.Processors, true
}

// HasProcessors returns a boolean if a field has been set.
func (o *ComputeBladeAllOf) HasProcessors() bool {
	if o != nil && o.Processors != nil {
		return true
	}

	return false
}

// SetProcessors gets a reference to the given []ProcessorUnitRelationship and assigns it to the Processors field.
func (o *ComputeBladeAllOf) SetProcessors(v []ProcessorUnitRelationship) {
	o.Processors = v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *ComputeBladeAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeBladeAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *ComputeBladeAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *ComputeBladeAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

// GetStorageControllers returns the StorageControllers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComputeBladeAllOf) GetStorageControllers() []StorageControllerRelationship {
	if o == nil {
		var ret []StorageControllerRelationship
		return ret
	}
	return o.StorageControllers
}

// GetStorageControllersOk returns a tuple with the StorageControllers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComputeBladeAllOf) GetStorageControllersOk() ([]StorageControllerRelationship, bool) {
	if o == nil || o.StorageControllers == nil {
		return nil, false
	}
	return o.StorageControllers, true
}

// HasStorageControllers returns a boolean if a field has been set.
func (o *ComputeBladeAllOf) HasStorageControllers() bool {
	if o != nil && o.StorageControllers != nil {
		return true
	}

	return false
}

// SetStorageControllers gets a reference to the given []StorageControllerRelationship and assigns it to the StorageControllers field.
func (o *ComputeBladeAllOf) SetStorageControllers(v []StorageControllerRelationship) {
	o.StorageControllers = v
}

// GetStorageEnclosures returns the StorageEnclosures field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComputeBladeAllOf) GetStorageEnclosures() []StorageEnclosureRelationship {
	if o == nil {
		var ret []StorageEnclosureRelationship
		return ret
	}
	return o.StorageEnclosures
}

// GetStorageEnclosuresOk returns a tuple with the StorageEnclosures field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComputeBladeAllOf) GetStorageEnclosuresOk() ([]StorageEnclosureRelationship, bool) {
	if o == nil || o.StorageEnclosures == nil {
		return nil, false
	}
	return o.StorageEnclosures, true
}

// HasStorageEnclosures returns a boolean if a field has been set.
func (o *ComputeBladeAllOf) HasStorageEnclosures() bool {
	if o != nil && o.StorageEnclosures != nil {
		return true
	}

	return false
}

// SetStorageEnclosures gets a reference to the given []StorageEnclosureRelationship and assigns it to the StorageEnclosures field.
func (o *ComputeBladeAllOf) SetStorageEnclosures(v []StorageEnclosureRelationship) {
	o.StorageEnclosures = v
}

// GetTopSystem returns the TopSystem field value if set, zero value otherwise.
func (o *ComputeBladeAllOf) GetTopSystem() TopSystemRelationship {
	if o == nil || o.TopSystem == nil {
		var ret TopSystemRelationship
		return ret
	}
	return *o.TopSystem
}

// GetTopSystemOk returns a tuple with the TopSystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeBladeAllOf) GetTopSystemOk() (*TopSystemRelationship, bool) {
	if o == nil || o.TopSystem == nil {
		return nil, false
	}
	return o.TopSystem, true
}

// HasTopSystem returns a boolean if a field has been set.
func (o *ComputeBladeAllOf) HasTopSystem() bool {
	if o != nil && o.TopSystem != nil {
		return true
	}

	return false
}

// SetTopSystem gets a reference to the given TopSystemRelationship and assigns it to the TopSystem field.
func (o *ComputeBladeAllOf) SetTopSystem(v TopSystemRelationship) {
	o.TopSystem = &v
}

func (o ComputeBladeAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ChassisId != nil {
		toSerialize["ChassisId"] = o.ChassisId
	}
	if o.ScaledMode != nil {
		toSerialize["ScaledMode"] = o.ScaledMode
	}
	if o.SlotId != nil {
		toSerialize["SlotId"] = o.SlotId
	}
	if o.Adapters != nil {
		toSerialize["Adapters"] = o.Adapters
	}
	if o.BiosBootmode != nil {
		toSerialize["BiosBootmode"] = o.BiosBootmode
	}
	if o.BiosTokenSettings != nil {
		toSerialize["BiosTokenSettings"] = o.BiosTokenSettings
	}
	if o.BiosUnits != nil {
		toSerialize["BiosUnits"] = o.BiosUnits
	}
	if o.BiosVfSelectMemoryRasConfiguration != nil {
		toSerialize["BiosVfSelectMemoryRasConfiguration"] = o.BiosVfSelectMemoryRasConfiguration
	}
	if o.Bmc != nil {
		toSerialize["Bmc"] = o.Bmc
	}
	if o.Board != nil {
		toSerialize["Board"] = o.Board
	}
	if o.BootDeviceBootmode != nil {
		toSerialize["BootDeviceBootmode"] = o.BootDeviceBootmode
	}
	if o.EquipmentChassis != nil {
		toSerialize["EquipmentChassis"] = o.EquipmentChassis
	}
	if o.EquipmentIoExpanders != nil {
		toSerialize["EquipmentIoExpanders"] = o.EquipmentIoExpanders
	}
	if o.GenericInventoryHolders != nil {
		toSerialize["GenericInventoryHolders"] = o.GenericInventoryHolders
	}
	if o.GraphicsCards != nil {
		toSerialize["GraphicsCards"] = o.GraphicsCards
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.LocatorLed != nil {
		toSerialize["LocatorLed"] = o.LocatorLed
	}
	if o.MemoryArrays != nil {
		toSerialize["MemoryArrays"] = o.MemoryArrays
	}
	if o.PciDevices != nil {
		toSerialize["PciDevices"] = o.PciDevices
	}
	if o.Processors != nil {
		toSerialize["Processors"] = o.Processors
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	if o.StorageControllers != nil {
		toSerialize["StorageControllers"] = o.StorageControllers
	}
	if o.StorageEnclosures != nil {
		toSerialize["StorageEnclosures"] = o.StorageEnclosures
	}
	if o.TopSystem != nil {
		toSerialize["TopSystem"] = o.TopSystem
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ComputeBladeAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varComputeBladeAllOf := _ComputeBladeAllOf{}

	if err = json.Unmarshal(bytes, &varComputeBladeAllOf); err == nil {
		*o = ComputeBladeAllOf(varComputeBladeAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ChassisId")
		delete(additionalProperties, "ScaledMode")
		delete(additionalProperties, "SlotId")
		delete(additionalProperties, "Adapters")
		delete(additionalProperties, "BiosBootmode")
		delete(additionalProperties, "BiosTokenSettings")
		delete(additionalProperties, "BiosUnits")
		delete(additionalProperties, "BiosVfSelectMemoryRasConfiguration")
		delete(additionalProperties, "Bmc")
		delete(additionalProperties, "Board")
		delete(additionalProperties, "BootDeviceBootmode")
		delete(additionalProperties, "EquipmentChassis")
		delete(additionalProperties, "EquipmentIoExpanders")
		delete(additionalProperties, "GenericInventoryHolders")
		delete(additionalProperties, "GraphicsCards")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "LocatorLed")
		delete(additionalProperties, "MemoryArrays")
		delete(additionalProperties, "PciDevices")
		delete(additionalProperties, "Processors")
		delete(additionalProperties, "RegisteredDevice")
		delete(additionalProperties, "StorageControllers")
		delete(additionalProperties, "StorageEnclosures")
		delete(additionalProperties, "TopSystem")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableComputeBladeAllOf struct {
	value *ComputeBladeAllOf
	isSet bool
}

func (v NullableComputeBladeAllOf) Get() *ComputeBladeAllOf {
	return v.value
}

func (v *NullableComputeBladeAllOf) Set(val *ComputeBladeAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableComputeBladeAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableComputeBladeAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputeBladeAllOf(val *ComputeBladeAllOf) *NullableComputeBladeAllOf {
	return &NullableComputeBladeAllOf{value: val, isSet: true}
}

func (v NullableComputeBladeAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputeBladeAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
