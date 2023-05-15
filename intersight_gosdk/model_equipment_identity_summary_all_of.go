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

// EquipmentIdentitySummaryAllOf Definition of the list of properties defined in 'equipment.IdentitySummary', excluding properties defined in parent classes.
type EquipmentIdentitySummaryAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Serial Identifier of an adapter participating in SWM.
	AdapterSerial *string `json:"AdapterSerial,omitempty"`
	// Updated by UI/API to trigger specific action type. * `None` - No operation value for maintenance actions on an equipment. * `Decommission` - Decommission the equipment and temporarily remove it from being managed by Intersight. * `Recommission` - Recommission the equipment. * `Reack` - Reacknowledge the equipment and discover it again. * `Remove` - Remove the equipment permanently from Intersight management. * `Replace` - Replace the equipment with the other one.
	AdminAction *string `json:"AdminAction,omitempty"`
	// The state of Maintenance Action performed. This will have three states. Applying - Action is in progress. Applied - Action is completed and applied. Failed - Action has failed. * `None` - Nil value when no action has been triggered by the user. * `Applied` - User configured settings are in applied state. * `Applying` - User settings are being applied on the target server. * `Failed` - User configured settings could not be applied.
	AdminActionState *string `json:"AdminActionState,omitempty"`
	// Chassis Identifier of a blade server.
	ChassisId *int64 `json:"ChassisId,omitempty"`
	// The id of the chassis that the blade is currently located in.
	CurrentChassisId *int64 `json:"CurrentChassisId,omitempty"`
	// The slot number in the chassis that the blade is currently located in.
	CurrentSlotId *int64 `json:"CurrentSlotId,omitempty"`
	// Describes whether the running CIMC version supports Intersight managed mode. * `Unknown` - The running firmware version is unknown. * `Supported` - The running firmware version is known and supports IMM mode. * `NotSupported` - The running firmware version is known and does not support IMM mode.
	FirmwareSupportability *string `json:"FirmwareSupportability,omitempty"`
	// Numeric Identifier assigned by the management system to the equipment. Identifier can only be changed if it has been PATCHED with the AdminAction property set to 'Recommission'.
	Identifier         *int64                    `json:"Identifier,omitempty"`
	IoCardIdentityList []EquipmentIoCardIdentity `json:"IoCardIdentityList,omitempty"`
	// Denotes the type of discovery that was most recently triggered on this server. * `Unknown` - The last discovery type is unknown. * `Deep` - The last discovery triggered is deep. * `Shallow` - The last discovery triggered is shallow.
	LastDiscoveryTriggered *string `json:"LastDiscoveryTriggered,omitempty"`
	// The equipment's lifecycle status. * `None` - Default state of an equipment. This should be an initial state when no state is defined for an equipment. * `Active` - Default Lifecycle State for a physical entity. * `Decommissioned` - Decommission Lifecycle state. * `DecommissionInProgress` - Decommission Inprogress Lifecycle state. * `RecommissionInProgress` - Recommission Inprogress Lifecycle state. * `OperationFailed` - Failed Operation Lifecycle state. * `ReackInProgress` - ReackInProgress Lifecycle state. * `RemoveInProgress` - RemoveInProgress Lifecycle state. * `Discovered` - Discovered Lifecycle state. * `DiscoveryInProgress` - DiscoveryInProgress Lifecycle state. * `DiscoveryFailed` - DiscoveryFailed Lifecycle state. * `FirmwareUpgradeInProgress` - Firmware upgrade is in progress on given physical entity. * `BladeMigrationInProgress` - Server slot migration is in progress on given physical entity. * `Inactive` - Inactive Lifecycle state. * `ReplaceInProgress` - ReplaceInProgress Lifecycle state. * `SlotMismatch` - The blade server is detected in a different chassis/slot than it was previously.
	Lifecycle *string `json:"Lifecycle,omitempty"`
	// Chassis slot number of the manager compute server.
	ManagerSlotId *int64 `json:"ManagerSlotId,omitempty"`
	// The vendor provided model name for the equipment.
	Model *string `json:"Model,omitempty"`
	// The presence state of the blade server. * `Unknown` - The default presence state. * `Equipped` - The server is equipped in the slot. * `EquippedMismatch` - The slot is equipped, but there is another server currently inventoried in the slot. * `Missing` - The server is not present in the given slot.
	Presence *string `json:"Presence,omitempty"`
	// The serial number of the equipment.
	Serial *string `json:"Serial,omitempty"`
	// Chassis slot number of a blade server.
	SlotId *int64 `json:"SlotId,omitempty"`
	// The source object type of this view MO.
	SourceObjectType *string `json:"SourceObjectType,omitempty"`
	// Switch ID to which Fabric Extender is connected, ID can be either 1 or 2, equalent to A or B.
	SwitchId *int64 `json:"SwitchId,omitempty"`
	// The manufacturer of the equipment.
	Vendor               *string                              `json:"Vendor,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EquipmentIdentitySummaryAllOf EquipmentIdentitySummaryAllOf

// NewEquipmentIdentitySummaryAllOf instantiates a new EquipmentIdentitySummaryAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquipmentIdentitySummaryAllOf(classId string, objectType string) *EquipmentIdentitySummaryAllOf {
	this := EquipmentIdentitySummaryAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewEquipmentIdentitySummaryAllOfWithDefaults instantiates a new EquipmentIdentitySummaryAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquipmentIdentitySummaryAllOfWithDefaults() *EquipmentIdentitySummaryAllOf {
	this := EquipmentIdentitySummaryAllOf{}
	var classId string = "equipment.IdentitySummary"
	this.ClassId = classId
	var objectType string = "equipment.IdentitySummary"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *EquipmentIdentitySummaryAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *EquipmentIdentitySummaryAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *EquipmentIdentitySummaryAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *EquipmentIdentitySummaryAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *EquipmentIdentitySummaryAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *EquipmentIdentitySummaryAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAdapterSerial returns the AdapterSerial field value if set, zero value otherwise.
func (o *EquipmentIdentitySummaryAllOf) GetAdapterSerial() string {
	if o == nil || o.AdapterSerial == nil {
		var ret string
		return ret
	}
	return *o.AdapterSerial
}

// GetAdapterSerialOk returns a tuple with the AdapterSerial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIdentitySummaryAllOf) GetAdapterSerialOk() (*string, bool) {
	if o == nil || o.AdapterSerial == nil {
		return nil, false
	}
	return o.AdapterSerial, true
}

// HasAdapterSerial returns a boolean if a field has been set.
func (o *EquipmentIdentitySummaryAllOf) HasAdapterSerial() bool {
	if o != nil && o.AdapterSerial != nil {
		return true
	}

	return false
}

// SetAdapterSerial gets a reference to the given string and assigns it to the AdapterSerial field.
func (o *EquipmentIdentitySummaryAllOf) SetAdapterSerial(v string) {
	o.AdapterSerial = &v
}

// GetAdminAction returns the AdminAction field value if set, zero value otherwise.
func (o *EquipmentIdentitySummaryAllOf) GetAdminAction() string {
	if o == nil || o.AdminAction == nil {
		var ret string
		return ret
	}
	return *o.AdminAction
}

// GetAdminActionOk returns a tuple with the AdminAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIdentitySummaryAllOf) GetAdminActionOk() (*string, bool) {
	if o == nil || o.AdminAction == nil {
		return nil, false
	}
	return o.AdminAction, true
}

// HasAdminAction returns a boolean if a field has been set.
func (o *EquipmentIdentitySummaryAllOf) HasAdminAction() bool {
	if o != nil && o.AdminAction != nil {
		return true
	}

	return false
}

// SetAdminAction gets a reference to the given string and assigns it to the AdminAction field.
func (o *EquipmentIdentitySummaryAllOf) SetAdminAction(v string) {
	o.AdminAction = &v
}

// GetAdminActionState returns the AdminActionState field value if set, zero value otherwise.
func (o *EquipmentIdentitySummaryAllOf) GetAdminActionState() string {
	if o == nil || o.AdminActionState == nil {
		var ret string
		return ret
	}
	return *o.AdminActionState
}

// GetAdminActionStateOk returns a tuple with the AdminActionState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIdentitySummaryAllOf) GetAdminActionStateOk() (*string, bool) {
	if o == nil || o.AdminActionState == nil {
		return nil, false
	}
	return o.AdminActionState, true
}

// HasAdminActionState returns a boolean if a field has been set.
func (o *EquipmentIdentitySummaryAllOf) HasAdminActionState() bool {
	if o != nil && o.AdminActionState != nil {
		return true
	}

	return false
}

// SetAdminActionState gets a reference to the given string and assigns it to the AdminActionState field.
func (o *EquipmentIdentitySummaryAllOf) SetAdminActionState(v string) {
	o.AdminActionState = &v
}

// GetChassisId returns the ChassisId field value if set, zero value otherwise.
func (o *EquipmentIdentitySummaryAllOf) GetChassisId() int64 {
	if o == nil || o.ChassisId == nil {
		var ret int64
		return ret
	}
	return *o.ChassisId
}

// GetChassisIdOk returns a tuple with the ChassisId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIdentitySummaryAllOf) GetChassisIdOk() (*int64, bool) {
	if o == nil || o.ChassisId == nil {
		return nil, false
	}
	return o.ChassisId, true
}

// HasChassisId returns a boolean if a field has been set.
func (o *EquipmentIdentitySummaryAllOf) HasChassisId() bool {
	if o != nil && o.ChassisId != nil {
		return true
	}

	return false
}

// SetChassisId gets a reference to the given int64 and assigns it to the ChassisId field.
func (o *EquipmentIdentitySummaryAllOf) SetChassisId(v int64) {
	o.ChassisId = &v
}

// GetCurrentChassisId returns the CurrentChassisId field value if set, zero value otherwise.
func (o *EquipmentIdentitySummaryAllOf) GetCurrentChassisId() int64 {
	if o == nil || o.CurrentChassisId == nil {
		var ret int64
		return ret
	}
	return *o.CurrentChassisId
}

// GetCurrentChassisIdOk returns a tuple with the CurrentChassisId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIdentitySummaryAllOf) GetCurrentChassisIdOk() (*int64, bool) {
	if o == nil || o.CurrentChassisId == nil {
		return nil, false
	}
	return o.CurrentChassisId, true
}

// HasCurrentChassisId returns a boolean if a field has been set.
func (o *EquipmentIdentitySummaryAllOf) HasCurrentChassisId() bool {
	if o != nil && o.CurrentChassisId != nil {
		return true
	}

	return false
}

// SetCurrentChassisId gets a reference to the given int64 and assigns it to the CurrentChassisId field.
func (o *EquipmentIdentitySummaryAllOf) SetCurrentChassisId(v int64) {
	o.CurrentChassisId = &v
}

// GetCurrentSlotId returns the CurrentSlotId field value if set, zero value otherwise.
func (o *EquipmentIdentitySummaryAllOf) GetCurrentSlotId() int64 {
	if o == nil || o.CurrentSlotId == nil {
		var ret int64
		return ret
	}
	return *o.CurrentSlotId
}

// GetCurrentSlotIdOk returns a tuple with the CurrentSlotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIdentitySummaryAllOf) GetCurrentSlotIdOk() (*int64, bool) {
	if o == nil || o.CurrentSlotId == nil {
		return nil, false
	}
	return o.CurrentSlotId, true
}

// HasCurrentSlotId returns a boolean if a field has been set.
func (o *EquipmentIdentitySummaryAllOf) HasCurrentSlotId() bool {
	if o != nil && o.CurrentSlotId != nil {
		return true
	}

	return false
}

// SetCurrentSlotId gets a reference to the given int64 and assigns it to the CurrentSlotId field.
func (o *EquipmentIdentitySummaryAllOf) SetCurrentSlotId(v int64) {
	o.CurrentSlotId = &v
}

// GetFirmwareSupportability returns the FirmwareSupportability field value if set, zero value otherwise.
func (o *EquipmentIdentitySummaryAllOf) GetFirmwareSupportability() string {
	if o == nil || o.FirmwareSupportability == nil {
		var ret string
		return ret
	}
	return *o.FirmwareSupportability
}

// GetFirmwareSupportabilityOk returns a tuple with the FirmwareSupportability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIdentitySummaryAllOf) GetFirmwareSupportabilityOk() (*string, bool) {
	if o == nil || o.FirmwareSupportability == nil {
		return nil, false
	}
	return o.FirmwareSupportability, true
}

// HasFirmwareSupportability returns a boolean if a field has been set.
func (o *EquipmentIdentitySummaryAllOf) HasFirmwareSupportability() bool {
	if o != nil && o.FirmwareSupportability != nil {
		return true
	}

	return false
}

// SetFirmwareSupportability gets a reference to the given string and assigns it to the FirmwareSupportability field.
func (o *EquipmentIdentitySummaryAllOf) SetFirmwareSupportability(v string) {
	o.FirmwareSupportability = &v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *EquipmentIdentitySummaryAllOf) GetIdentifier() int64 {
	if o == nil || o.Identifier == nil {
		var ret int64
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIdentitySummaryAllOf) GetIdentifierOk() (*int64, bool) {
	if o == nil || o.Identifier == nil {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *EquipmentIdentitySummaryAllOf) HasIdentifier() bool {
	if o != nil && o.Identifier != nil {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given int64 and assigns it to the Identifier field.
func (o *EquipmentIdentitySummaryAllOf) SetIdentifier(v int64) {
	o.Identifier = &v
}

// GetIoCardIdentityList returns the IoCardIdentityList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquipmentIdentitySummaryAllOf) GetIoCardIdentityList() []EquipmentIoCardIdentity {
	if o == nil {
		var ret []EquipmentIoCardIdentity
		return ret
	}
	return o.IoCardIdentityList
}

// GetIoCardIdentityListOk returns a tuple with the IoCardIdentityList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquipmentIdentitySummaryAllOf) GetIoCardIdentityListOk() ([]EquipmentIoCardIdentity, bool) {
	if o == nil || o.IoCardIdentityList == nil {
		return nil, false
	}
	return o.IoCardIdentityList, true
}

// HasIoCardIdentityList returns a boolean if a field has been set.
func (o *EquipmentIdentitySummaryAllOf) HasIoCardIdentityList() bool {
	if o != nil && o.IoCardIdentityList != nil {
		return true
	}

	return false
}

// SetIoCardIdentityList gets a reference to the given []EquipmentIoCardIdentity and assigns it to the IoCardIdentityList field.
func (o *EquipmentIdentitySummaryAllOf) SetIoCardIdentityList(v []EquipmentIoCardIdentity) {
	o.IoCardIdentityList = v
}

// GetLastDiscoveryTriggered returns the LastDiscoveryTriggered field value if set, zero value otherwise.
func (o *EquipmentIdentitySummaryAllOf) GetLastDiscoveryTriggered() string {
	if o == nil || o.LastDiscoveryTriggered == nil {
		var ret string
		return ret
	}
	return *o.LastDiscoveryTriggered
}

// GetLastDiscoveryTriggeredOk returns a tuple with the LastDiscoveryTriggered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIdentitySummaryAllOf) GetLastDiscoveryTriggeredOk() (*string, bool) {
	if o == nil || o.LastDiscoveryTriggered == nil {
		return nil, false
	}
	return o.LastDiscoveryTriggered, true
}

// HasLastDiscoveryTriggered returns a boolean if a field has been set.
func (o *EquipmentIdentitySummaryAllOf) HasLastDiscoveryTriggered() bool {
	if o != nil && o.LastDiscoveryTriggered != nil {
		return true
	}

	return false
}

// SetLastDiscoveryTriggered gets a reference to the given string and assigns it to the LastDiscoveryTriggered field.
func (o *EquipmentIdentitySummaryAllOf) SetLastDiscoveryTriggered(v string) {
	o.LastDiscoveryTriggered = &v
}

// GetLifecycle returns the Lifecycle field value if set, zero value otherwise.
func (o *EquipmentIdentitySummaryAllOf) GetLifecycle() string {
	if o == nil || o.Lifecycle == nil {
		var ret string
		return ret
	}
	return *o.Lifecycle
}

// GetLifecycleOk returns a tuple with the Lifecycle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIdentitySummaryAllOf) GetLifecycleOk() (*string, bool) {
	if o == nil || o.Lifecycle == nil {
		return nil, false
	}
	return o.Lifecycle, true
}

// HasLifecycle returns a boolean if a field has been set.
func (o *EquipmentIdentitySummaryAllOf) HasLifecycle() bool {
	if o != nil && o.Lifecycle != nil {
		return true
	}

	return false
}

// SetLifecycle gets a reference to the given string and assigns it to the Lifecycle field.
func (o *EquipmentIdentitySummaryAllOf) SetLifecycle(v string) {
	o.Lifecycle = &v
}

// GetManagerSlotId returns the ManagerSlotId field value if set, zero value otherwise.
func (o *EquipmentIdentitySummaryAllOf) GetManagerSlotId() int64 {
	if o == nil || o.ManagerSlotId == nil {
		var ret int64
		return ret
	}
	return *o.ManagerSlotId
}

// GetManagerSlotIdOk returns a tuple with the ManagerSlotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIdentitySummaryAllOf) GetManagerSlotIdOk() (*int64, bool) {
	if o == nil || o.ManagerSlotId == nil {
		return nil, false
	}
	return o.ManagerSlotId, true
}

// HasManagerSlotId returns a boolean if a field has been set.
func (o *EquipmentIdentitySummaryAllOf) HasManagerSlotId() bool {
	if o != nil && o.ManagerSlotId != nil {
		return true
	}

	return false
}

// SetManagerSlotId gets a reference to the given int64 and assigns it to the ManagerSlotId field.
func (o *EquipmentIdentitySummaryAllOf) SetManagerSlotId(v int64) {
	o.ManagerSlotId = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *EquipmentIdentitySummaryAllOf) GetModel() string {
	if o == nil || o.Model == nil {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIdentitySummaryAllOf) GetModelOk() (*string, bool) {
	if o == nil || o.Model == nil {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *EquipmentIdentitySummaryAllOf) HasModel() bool {
	if o != nil && o.Model != nil {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *EquipmentIdentitySummaryAllOf) SetModel(v string) {
	o.Model = &v
}

// GetPresence returns the Presence field value if set, zero value otherwise.
func (o *EquipmentIdentitySummaryAllOf) GetPresence() string {
	if o == nil || o.Presence == nil {
		var ret string
		return ret
	}
	return *o.Presence
}

// GetPresenceOk returns a tuple with the Presence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIdentitySummaryAllOf) GetPresenceOk() (*string, bool) {
	if o == nil || o.Presence == nil {
		return nil, false
	}
	return o.Presence, true
}

// HasPresence returns a boolean if a field has been set.
func (o *EquipmentIdentitySummaryAllOf) HasPresence() bool {
	if o != nil && o.Presence != nil {
		return true
	}

	return false
}

// SetPresence gets a reference to the given string and assigns it to the Presence field.
func (o *EquipmentIdentitySummaryAllOf) SetPresence(v string) {
	o.Presence = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *EquipmentIdentitySummaryAllOf) GetSerial() string {
	if o == nil || o.Serial == nil {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIdentitySummaryAllOf) GetSerialOk() (*string, bool) {
	if o == nil || o.Serial == nil {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *EquipmentIdentitySummaryAllOf) HasSerial() bool {
	if o != nil && o.Serial != nil {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *EquipmentIdentitySummaryAllOf) SetSerial(v string) {
	o.Serial = &v
}

// GetSlotId returns the SlotId field value if set, zero value otherwise.
func (o *EquipmentIdentitySummaryAllOf) GetSlotId() int64 {
	if o == nil || o.SlotId == nil {
		var ret int64
		return ret
	}
	return *o.SlotId
}

// GetSlotIdOk returns a tuple with the SlotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIdentitySummaryAllOf) GetSlotIdOk() (*int64, bool) {
	if o == nil || o.SlotId == nil {
		return nil, false
	}
	return o.SlotId, true
}

// HasSlotId returns a boolean if a field has been set.
func (o *EquipmentIdentitySummaryAllOf) HasSlotId() bool {
	if o != nil && o.SlotId != nil {
		return true
	}

	return false
}

// SetSlotId gets a reference to the given int64 and assigns it to the SlotId field.
func (o *EquipmentIdentitySummaryAllOf) SetSlotId(v int64) {
	o.SlotId = &v
}

// GetSourceObjectType returns the SourceObjectType field value if set, zero value otherwise.
func (o *EquipmentIdentitySummaryAllOf) GetSourceObjectType() string {
	if o == nil || o.SourceObjectType == nil {
		var ret string
		return ret
	}
	return *o.SourceObjectType
}

// GetSourceObjectTypeOk returns a tuple with the SourceObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIdentitySummaryAllOf) GetSourceObjectTypeOk() (*string, bool) {
	if o == nil || o.SourceObjectType == nil {
		return nil, false
	}
	return o.SourceObjectType, true
}

// HasSourceObjectType returns a boolean if a field has been set.
func (o *EquipmentIdentitySummaryAllOf) HasSourceObjectType() bool {
	if o != nil && o.SourceObjectType != nil {
		return true
	}

	return false
}

// SetSourceObjectType gets a reference to the given string and assigns it to the SourceObjectType field.
func (o *EquipmentIdentitySummaryAllOf) SetSourceObjectType(v string) {
	o.SourceObjectType = &v
}

// GetSwitchId returns the SwitchId field value if set, zero value otherwise.
func (o *EquipmentIdentitySummaryAllOf) GetSwitchId() int64 {
	if o == nil || o.SwitchId == nil {
		var ret int64
		return ret
	}
	return *o.SwitchId
}

// GetSwitchIdOk returns a tuple with the SwitchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIdentitySummaryAllOf) GetSwitchIdOk() (*int64, bool) {
	if o == nil || o.SwitchId == nil {
		return nil, false
	}
	return o.SwitchId, true
}

// HasSwitchId returns a boolean if a field has been set.
func (o *EquipmentIdentitySummaryAllOf) HasSwitchId() bool {
	if o != nil && o.SwitchId != nil {
		return true
	}

	return false
}

// SetSwitchId gets a reference to the given int64 and assigns it to the SwitchId field.
func (o *EquipmentIdentitySummaryAllOf) SetSwitchId(v int64) {
	o.SwitchId = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *EquipmentIdentitySummaryAllOf) GetVendor() string {
	if o == nil || o.Vendor == nil {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIdentitySummaryAllOf) GetVendorOk() (*string, bool) {
	if o == nil || o.Vendor == nil {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *EquipmentIdentitySummaryAllOf) HasVendor() bool {
	if o != nil && o.Vendor != nil {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *EquipmentIdentitySummaryAllOf) SetVendor(v string) {
	o.Vendor = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *EquipmentIdentitySummaryAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIdentitySummaryAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *EquipmentIdentitySummaryAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *EquipmentIdentitySummaryAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o EquipmentIdentitySummaryAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AdapterSerial != nil {
		toSerialize["AdapterSerial"] = o.AdapterSerial
	}
	if o.AdminAction != nil {
		toSerialize["AdminAction"] = o.AdminAction
	}
	if o.AdminActionState != nil {
		toSerialize["AdminActionState"] = o.AdminActionState
	}
	if o.ChassisId != nil {
		toSerialize["ChassisId"] = o.ChassisId
	}
	if o.CurrentChassisId != nil {
		toSerialize["CurrentChassisId"] = o.CurrentChassisId
	}
	if o.CurrentSlotId != nil {
		toSerialize["CurrentSlotId"] = o.CurrentSlotId
	}
	if o.FirmwareSupportability != nil {
		toSerialize["FirmwareSupportability"] = o.FirmwareSupportability
	}
	if o.Identifier != nil {
		toSerialize["Identifier"] = o.Identifier
	}
	if o.IoCardIdentityList != nil {
		toSerialize["IoCardIdentityList"] = o.IoCardIdentityList
	}
	if o.LastDiscoveryTriggered != nil {
		toSerialize["LastDiscoveryTriggered"] = o.LastDiscoveryTriggered
	}
	if o.Lifecycle != nil {
		toSerialize["Lifecycle"] = o.Lifecycle
	}
	if o.ManagerSlotId != nil {
		toSerialize["ManagerSlotId"] = o.ManagerSlotId
	}
	if o.Model != nil {
		toSerialize["Model"] = o.Model
	}
	if o.Presence != nil {
		toSerialize["Presence"] = o.Presence
	}
	if o.Serial != nil {
		toSerialize["Serial"] = o.Serial
	}
	if o.SlotId != nil {
		toSerialize["SlotId"] = o.SlotId
	}
	if o.SourceObjectType != nil {
		toSerialize["SourceObjectType"] = o.SourceObjectType
	}
	if o.SwitchId != nil {
		toSerialize["SwitchId"] = o.SwitchId
	}
	if o.Vendor != nil {
		toSerialize["Vendor"] = o.Vendor
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EquipmentIdentitySummaryAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varEquipmentIdentitySummaryAllOf := _EquipmentIdentitySummaryAllOf{}

	if err = json.Unmarshal(bytes, &varEquipmentIdentitySummaryAllOf); err == nil {
		*o = EquipmentIdentitySummaryAllOf(varEquipmentIdentitySummaryAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AdapterSerial")
		delete(additionalProperties, "AdminAction")
		delete(additionalProperties, "AdminActionState")
		delete(additionalProperties, "ChassisId")
		delete(additionalProperties, "CurrentChassisId")
		delete(additionalProperties, "CurrentSlotId")
		delete(additionalProperties, "FirmwareSupportability")
		delete(additionalProperties, "Identifier")
		delete(additionalProperties, "IoCardIdentityList")
		delete(additionalProperties, "LastDiscoveryTriggered")
		delete(additionalProperties, "Lifecycle")
		delete(additionalProperties, "ManagerSlotId")
		delete(additionalProperties, "Model")
		delete(additionalProperties, "Presence")
		delete(additionalProperties, "Serial")
		delete(additionalProperties, "SlotId")
		delete(additionalProperties, "SourceObjectType")
		delete(additionalProperties, "SwitchId")
		delete(additionalProperties, "Vendor")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEquipmentIdentitySummaryAllOf struct {
	value *EquipmentIdentitySummaryAllOf
	isSet bool
}

func (v NullableEquipmentIdentitySummaryAllOf) Get() *EquipmentIdentitySummaryAllOf {
	return v.value
}

func (v *NullableEquipmentIdentitySummaryAllOf) Set(val *EquipmentIdentitySummaryAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEquipmentIdentitySummaryAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEquipmentIdentitySummaryAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquipmentIdentitySummaryAllOf(val *EquipmentIdentitySummaryAllOf) *NullableEquipmentIdentitySummaryAllOf {
	return &NullableEquipmentIdentitySummaryAllOf{value: val, isSet: true}
}

func (v NullableEquipmentIdentitySummaryAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquipmentIdentitySummaryAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
