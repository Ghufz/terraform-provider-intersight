/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-7658
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// ManagementInterface Interface that provides access to the management controller.
type ManagementInterface struct {
	InventoryBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Default gateway for the interface.
	Gateway *string `json:"Gateway,omitempty"`
	// Hostname configured for the interface.
	HostName *string `json:"HostName,omitempty"`
	// IP address of the interface.
	IpAddress *string `json:"IpAddress,omitempty"`
	// IPv4 address of the interface.
	Ipv4Address *string `json:"Ipv4Address,omitempty"`
	// IPv4 default gateway for the interface.
	Ipv4Gateway *string `json:"Ipv4Gateway,omitempty"`
	// IPv4 Netmask for the interface.
	Ipv4Mask *string `json:"Ipv4Mask,omitempty"`
	// IPv6 address of the interface.
	Ipv6Address *string `json:"Ipv6Address,omitempty"`
	// IPv6 default gateway for the interface.
	Ipv6Gateway *string `json:"Ipv6Gateway,omitempty"`
	// IPv6 prefix for the interface.
	Ipv6Prefix *int64 `json:"Ipv6Prefix,omitempty"`
	// MAC address configured for the interface.
	MacAddress *string `json:"MacAddress,omitempty"`
	// Netmask for the interface.
	Mask *string `json:"Mask,omitempty"`
	// Switch Id connected to the interface.
	SwitchId *string `json:"SwitchId,omitempty"`
	// The event channel connection status for the interface.
	UemConnStatus *string `json:"UemConnStatus,omitempty"`
	// Virtual hostname configured for the interface in case of clustered environment.
	VirtualHostName *string `json:"VirtualHostName,omitempty"`
	// VlanId configured for the interface.
	VlanId               *int64                               `json:"VlanId,omitempty"`
	EquipmentChassis     *EquipmentChassisRelationship        `json:"EquipmentChassis,omitempty"`
	InventoryDeviceInfo  *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
	ManagementController *ManagementControllerRelationship    `json:"ManagementController,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ManagementInterface ManagementInterface

// NewManagementInterface instantiates a new ManagementInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagementInterface(classId string, objectType string) *ManagementInterface {
	this := ManagementInterface{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewManagementInterfaceWithDefaults instantiates a new ManagementInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagementInterfaceWithDefaults() *ManagementInterface {
	this := ManagementInterface{}
	var classId string = "management.Interface"
	this.ClassId = classId
	var objectType string = "management.Interface"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ManagementInterface) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ManagementInterface) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ManagementInterface) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ManagementInterface) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ManagementInterface) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ManagementInterface) SetObjectType(v string) {
	o.ObjectType = v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *ManagementInterface) GetGateway() string {
	if o == nil || o.Gateway == nil {
		var ret string
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementInterface) GetGatewayOk() (*string, bool) {
	if o == nil || o.Gateway == nil {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *ManagementInterface) HasGateway() bool {
	if o != nil && o.Gateway != nil {
		return true
	}

	return false
}

// SetGateway gets a reference to the given string and assigns it to the Gateway field.
func (o *ManagementInterface) SetGateway(v string) {
	o.Gateway = &v
}

// GetHostName returns the HostName field value if set, zero value otherwise.
func (o *ManagementInterface) GetHostName() string {
	if o == nil || o.HostName == nil {
		var ret string
		return ret
	}
	return *o.HostName
}

// GetHostNameOk returns a tuple with the HostName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementInterface) GetHostNameOk() (*string, bool) {
	if o == nil || o.HostName == nil {
		return nil, false
	}
	return o.HostName, true
}

// HasHostName returns a boolean if a field has been set.
func (o *ManagementInterface) HasHostName() bool {
	if o != nil && o.HostName != nil {
		return true
	}

	return false
}

// SetHostName gets a reference to the given string and assigns it to the HostName field.
func (o *ManagementInterface) SetHostName(v string) {
	o.HostName = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *ManagementInterface) GetIpAddress() string {
	if o == nil || o.IpAddress == nil {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementInterface) GetIpAddressOk() (*string, bool) {
	if o == nil || o.IpAddress == nil {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *ManagementInterface) HasIpAddress() bool {
	if o != nil && o.IpAddress != nil {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *ManagementInterface) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetIpv4Address returns the Ipv4Address field value if set, zero value otherwise.
func (o *ManagementInterface) GetIpv4Address() string {
	if o == nil || o.Ipv4Address == nil {
		var ret string
		return ret
	}
	return *o.Ipv4Address
}

// GetIpv4AddressOk returns a tuple with the Ipv4Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementInterface) GetIpv4AddressOk() (*string, bool) {
	if o == nil || o.Ipv4Address == nil {
		return nil, false
	}
	return o.Ipv4Address, true
}

// HasIpv4Address returns a boolean if a field has been set.
func (o *ManagementInterface) HasIpv4Address() bool {
	if o != nil && o.Ipv4Address != nil {
		return true
	}

	return false
}

// SetIpv4Address gets a reference to the given string and assigns it to the Ipv4Address field.
func (o *ManagementInterface) SetIpv4Address(v string) {
	o.Ipv4Address = &v
}

// GetIpv4Gateway returns the Ipv4Gateway field value if set, zero value otherwise.
func (o *ManagementInterface) GetIpv4Gateway() string {
	if o == nil || o.Ipv4Gateway == nil {
		var ret string
		return ret
	}
	return *o.Ipv4Gateway
}

// GetIpv4GatewayOk returns a tuple with the Ipv4Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementInterface) GetIpv4GatewayOk() (*string, bool) {
	if o == nil || o.Ipv4Gateway == nil {
		return nil, false
	}
	return o.Ipv4Gateway, true
}

// HasIpv4Gateway returns a boolean if a field has been set.
func (o *ManagementInterface) HasIpv4Gateway() bool {
	if o != nil && o.Ipv4Gateway != nil {
		return true
	}

	return false
}

// SetIpv4Gateway gets a reference to the given string and assigns it to the Ipv4Gateway field.
func (o *ManagementInterface) SetIpv4Gateway(v string) {
	o.Ipv4Gateway = &v
}

// GetIpv4Mask returns the Ipv4Mask field value if set, zero value otherwise.
func (o *ManagementInterface) GetIpv4Mask() string {
	if o == nil || o.Ipv4Mask == nil {
		var ret string
		return ret
	}
	return *o.Ipv4Mask
}

// GetIpv4MaskOk returns a tuple with the Ipv4Mask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementInterface) GetIpv4MaskOk() (*string, bool) {
	if o == nil || o.Ipv4Mask == nil {
		return nil, false
	}
	return o.Ipv4Mask, true
}

// HasIpv4Mask returns a boolean if a field has been set.
func (o *ManagementInterface) HasIpv4Mask() bool {
	if o != nil && o.Ipv4Mask != nil {
		return true
	}

	return false
}

// SetIpv4Mask gets a reference to the given string and assigns it to the Ipv4Mask field.
func (o *ManagementInterface) SetIpv4Mask(v string) {
	o.Ipv4Mask = &v
}

// GetIpv6Address returns the Ipv6Address field value if set, zero value otherwise.
func (o *ManagementInterface) GetIpv6Address() string {
	if o == nil || o.Ipv6Address == nil {
		var ret string
		return ret
	}
	return *o.Ipv6Address
}

// GetIpv6AddressOk returns a tuple with the Ipv6Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementInterface) GetIpv6AddressOk() (*string, bool) {
	if o == nil || o.Ipv6Address == nil {
		return nil, false
	}
	return o.Ipv6Address, true
}

// HasIpv6Address returns a boolean if a field has been set.
func (o *ManagementInterface) HasIpv6Address() bool {
	if o != nil && o.Ipv6Address != nil {
		return true
	}

	return false
}

// SetIpv6Address gets a reference to the given string and assigns it to the Ipv6Address field.
func (o *ManagementInterface) SetIpv6Address(v string) {
	o.Ipv6Address = &v
}

// GetIpv6Gateway returns the Ipv6Gateway field value if set, zero value otherwise.
func (o *ManagementInterface) GetIpv6Gateway() string {
	if o == nil || o.Ipv6Gateway == nil {
		var ret string
		return ret
	}
	return *o.Ipv6Gateway
}

// GetIpv6GatewayOk returns a tuple with the Ipv6Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementInterface) GetIpv6GatewayOk() (*string, bool) {
	if o == nil || o.Ipv6Gateway == nil {
		return nil, false
	}
	return o.Ipv6Gateway, true
}

// HasIpv6Gateway returns a boolean if a field has been set.
func (o *ManagementInterface) HasIpv6Gateway() bool {
	if o != nil && o.Ipv6Gateway != nil {
		return true
	}

	return false
}

// SetIpv6Gateway gets a reference to the given string and assigns it to the Ipv6Gateway field.
func (o *ManagementInterface) SetIpv6Gateway(v string) {
	o.Ipv6Gateway = &v
}

// GetIpv6Prefix returns the Ipv6Prefix field value if set, zero value otherwise.
func (o *ManagementInterface) GetIpv6Prefix() int64 {
	if o == nil || o.Ipv6Prefix == nil {
		var ret int64
		return ret
	}
	return *o.Ipv6Prefix
}

// GetIpv6PrefixOk returns a tuple with the Ipv6Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementInterface) GetIpv6PrefixOk() (*int64, bool) {
	if o == nil || o.Ipv6Prefix == nil {
		return nil, false
	}
	return o.Ipv6Prefix, true
}

// HasIpv6Prefix returns a boolean if a field has been set.
func (o *ManagementInterface) HasIpv6Prefix() bool {
	if o != nil && o.Ipv6Prefix != nil {
		return true
	}

	return false
}

// SetIpv6Prefix gets a reference to the given int64 and assigns it to the Ipv6Prefix field.
func (o *ManagementInterface) SetIpv6Prefix(v int64) {
	o.Ipv6Prefix = &v
}

// GetMacAddress returns the MacAddress field value if set, zero value otherwise.
func (o *ManagementInterface) GetMacAddress() string {
	if o == nil || o.MacAddress == nil {
		var ret string
		return ret
	}
	return *o.MacAddress
}

// GetMacAddressOk returns a tuple with the MacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementInterface) GetMacAddressOk() (*string, bool) {
	if o == nil || o.MacAddress == nil {
		return nil, false
	}
	return o.MacAddress, true
}

// HasMacAddress returns a boolean if a field has been set.
func (o *ManagementInterface) HasMacAddress() bool {
	if o != nil && o.MacAddress != nil {
		return true
	}

	return false
}

// SetMacAddress gets a reference to the given string and assigns it to the MacAddress field.
func (o *ManagementInterface) SetMacAddress(v string) {
	o.MacAddress = &v
}

// GetMask returns the Mask field value if set, zero value otherwise.
func (o *ManagementInterface) GetMask() string {
	if o == nil || o.Mask == nil {
		var ret string
		return ret
	}
	return *o.Mask
}

// GetMaskOk returns a tuple with the Mask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementInterface) GetMaskOk() (*string, bool) {
	if o == nil || o.Mask == nil {
		return nil, false
	}
	return o.Mask, true
}

// HasMask returns a boolean if a field has been set.
func (o *ManagementInterface) HasMask() bool {
	if o != nil && o.Mask != nil {
		return true
	}

	return false
}

// SetMask gets a reference to the given string and assigns it to the Mask field.
func (o *ManagementInterface) SetMask(v string) {
	o.Mask = &v
}

// GetSwitchId returns the SwitchId field value if set, zero value otherwise.
func (o *ManagementInterface) GetSwitchId() string {
	if o == nil || o.SwitchId == nil {
		var ret string
		return ret
	}
	return *o.SwitchId
}

// GetSwitchIdOk returns a tuple with the SwitchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementInterface) GetSwitchIdOk() (*string, bool) {
	if o == nil || o.SwitchId == nil {
		return nil, false
	}
	return o.SwitchId, true
}

// HasSwitchId returns a boolean if a field has been set.
func (o *ManagementInterface) HasSwitchId() bool {
	if o != nil && o.SwitchId != nil {
		return true
	}

	return false
}

// SetSwitchId gets a reference to the given string and assigns it to the SwitchId field.
func (o *ManagementInterface) SetSwitchId(v string) {
	o.SwitchId = &v
}

// GetUemConnStatus returns the UemConnStatus field value if set, zero value otherwise.
func (o *ManagementInterface) GetUemConnStatus() string {
	if o == nil || o.UemConnStatus == nil {
		var ret string
		return ret
	}
	return *o.UemConnStatus
}

// GetUemConnStatusOk returns a tuple with the UemConnStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementInterface) GetUemConnStatusOk() (*string, bool) {
	if o == nil || o.UemConnStatus == nil {
		return nil, false
	}
	return o.UemConnStatus, true
}

// HasUemConnStatus returns a boolean if a field has been set.
func (o *ManagementInterface) HasUemConnStatus() bool {
	if o != nil && o.UemConnStatus != nil {
		return true
	}

	return false
}

// SetUemConnStatus gets a reference to the given string and assigns it to the UemConnStatus field.
func (o *ManagementInterface) SetUemConnStatus(v string) {
	o.UemConnStatus = &v
}

// GetVirtualHostName returns the VirtualHostName field value if set, zero value otherwise.
func (o *ManagementInterface) GetVirtualHostName() string {
	if o == nil || o.VirtualHostName == nil {
		var ret string
		return ret
	}
	return *o.VirtualHostName
}

// GetVirtualHostNameOk returns a tuple with the VirtualHostName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementInterface) GetVirtualHostNameOk() (*string, bool) {
	if o == nil || o.VirtualHostName == nil {
		return nil, false
	}
	return o.VirtualHostName, true
}

// HasVirtualHostName returns a boolean if a field has been set.
func (o *ManagementInterface) HasVirtualHostName() bool {
	if o != nil && o.VirtualHostName != nil {
		return true
	}

	return false
}

// SetVirtualHostName gets a reference to the given string and assigns it to the VirtualHostName field.
func (o *ManagementInterface) SetVirtualHostName(v string) {
	o.VirtualHostName = &v
}

// GetVlanId returns the VlanId field value if set, zero value otherwise.
func (o *ManagementInterface) GetVlanId() int64 {
	if o == nil || o.VlanId == nil {
		var ret int64
		return ret
	}
	return *o.VlanId
}

// GetVlanIdOk returns a tuple with the VlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementInterface) GetVlanIdOk() (*int64, bool) {
	if o == nil || o.VlanId == nil {
		return nil, false
	}
	return o.VlanId, true
}

// HasVlanId returns a boolean if a field has been set.
func (o *ManagementInterface) HasVlanId() bool {
	if o != nil && o.VlanId != nil {
		return true
	}

	return false
}

// SetVlanId gets a reference to the given int64 and assigns it to the VlanId field.
func (o *ManagementInterface) SetVlanId(v int64) {
	o.VlanId = &v
}

// GetEquipmentChassis returns the EquipmentChassis field value if set, zero value otherwise.
func (o *ManagementInterface) GetEquipmentChassis() EquipmentChassisRelationship {
	if o == nil || o.EquipmentChassis == nil {
		var ret EquipmentChassisRelationship
		return ret
	}
	return *o.EquipmentChassis
}

// GetEquipmentChassisOk returns a tuple with the EquipmentChassis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementInterface) GetEquipmentChassisOk() (*EquipmentChassisRelationship, bool) {
	if o == nil || o.EquipmentChassis == nil {
		return nil, false
	}
	return o.EquipmentChassis, true
}

// HasEquipmentChassis returns a boolean if a field has been set.
func (o *ManagementInterface) HasEquipmentChassis() bool {
	if o != nil && o.EquipmentChassis != nil {
		return true
	}

	return false
}

// SetEquipmentChassis gets a reference to the given EquipmentChassisRelationship and assigns it to the EquipmentChassis field.
func (o *ManagementInterface) SetEquipmentChassis(v EquipmentChassisRelationship) {
	o.EquipmentChassis = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *ManagementInterface) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementInterface) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *ManagementInterface) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *ManagementInterface) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetManagementController returns the ManagementController field value if set, zero value otherwise.
func (o *ManagementInterface) GetManagementController() ManagementControllerRelationship {
	if o == nil || o.ManagementController == nil {
		var ret ManagementControllerRelationship
		return ret
	}
	return *o.ManagementController
}

// GetManagementControllerOk returns a tuple with the ManagementController field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementInterface) GetManagementControllerOk() (*ManagementControllerRelationship, bool) {
	if o == nil || o.ManagementController == nil {
		return nil, false
	}
	return o.ManagementController, true
}

// HasManagementController returns a boolean if a field has been set.
func (o *ManagementInterface) HasManagementController() bool {
	if o != nil && o.ManagementController != nil {
		return true
	}

	return false
}

// SetManagementController gets a reference to the given ManagementControllerRelationship and assigns it to the ManagementController field.
func (o *ManagementInterface) SetManagementController(v ManagementControllerRelationship) {
	o.ManagementController = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *ManagementInterface) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementInterface) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *ManagementInterface) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *ManagementInterface) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o ManagementInterface) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedInventoryBase, errInventoryBase := json.Marshal(o.InventoryBase)
	if errInventoryBase != nil {
		return []byte{}, errInventoryBase
	}
	errInventoryBase = json.Unmarshal([]byte(serializedInventoryBase), &toSerialize)
	if errInventoryBase != nil {
		return []byte{}, errInventoryBase
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Gateway != nil {
		toSerialize["Gateway"] = o.Gateway
	}
	if o.HostName != nil {
		toSerialize["HostName"] = o.HostName
	}
	if o.IpAddress != nil {
		toSerialize["IpAddress"] = o.IpAddress
	}
	if o.Ipv4Address != nil {
		toSerialize["Ipv4Address"] = o.Ipv4Address
	}
	if o.Ipv4Gateway != nil {
		toSerialize["Ipv4Gateway"] = o.Ipv4Gateway
	}
	if o.Ipv4Mask != nil {
		toSerialize["Ipv4Mask"] = o.Ipv4Mask
	}
	if o.Ipv6Address != nil {
		toSerialize["Ipv6Address"] = o.Ipv6Address
	}
	if o.Ipv6Gateway != nil {
		toSerialize["Ipv6Gateway"] = o.Ipv6Gateway
	}
	if o.Ipv6Prefix != nil {
		toSerialize["Ipv6Prefix"] = o.Ipv6Prefix
	}
	if o.MacAddress != nil {
		toSerialize["MacAddress"] = o.MacAddress
	}
	if o.Mask != nil {
		toSerialize["Mask"] = o.Mask
	}
	if o.SwitchId != nil {
		toSerialize["SwitchId"] = o.SwitchId
	}
	if o.UemConnStatus != nil {
		toSerialize["UemConnStatus"] = o.UemConnStatus
	}
	if o.VirtualHostName != nil {
		toSerialize["VirtualHostName"] = o.VirtualHostName
	}
	if o.VlanId != nil {
		toSerialize["VlanId"] = o.VlanId
	}
	if o.EquipmentChassis != nil {
		toSerialize["EquipmentChassis"] = o.EquipmentChassis
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.ManagementController != nil {
		toSerialize["ManagementController"] = o.ManagementController
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ManagementInterface) UnmarshalJSON(bytes []byte) (err error) {
	type ManagementInterfaceWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Default gateway for the interface.
		Gateway *string `json:"Gateway,omitempty"`
		// Hostname configured for the interface.
		HostName *string `json:"HostName,omitempty"`
		// IP address of the interface.
		IpAddress *string `json:"IpAddress,omitempty"`
		// IPv4 address of the interface.
		Ipv4Address *string `json:"Ipv4Address,omitempty"`
		// IPv4 default gateway for the interface.
		Ipv4Gateway *string `json:"Ipv4Gateway,omitempty"`
		// IPv4 Netmask for the interface.
		Ipv4Mask *string `json:"Ipv4Mask,omitempty"`
		// IPv6 address of the interface.
		Ipv6Address *string `json:"Ipv6Address,omitempty"`
		// IPv6 default gateway for the interface.
		Ipv6Gateway *string `json:"Ipv6Gateway,omitempty"`
		// IPv6 prefix for the interface.
		Ipv6Prefix *int64 `json:"Ipv6Prefix,omitempty"`
		// MAC address configured for the interface.
		MacAddress *string `json:"MacAddress,omitempty"`
		// Netmask for the interface.
		Mask *string `json:"Mask,omitempty"`
		// Switch Id connected to the interface.
		SwitchId *string `json:"SwitchId,omitempty"`
		// The event channel connection status for the interface.
		UemConnStatus *string `json:"UemConnStatus,omitempty"`
		// Virtual hostname configured for the interface in case of clustered environment.
		VirtualHostName *string `json:"VirtualHostName,omitempty"`
		// VlanId configured for the interface.
		VlanId               *int64                               `json:"VlanId,omitempty"`
		EquipmentChassis     *EquipmentChassisRelationship        `json:"EquipmentChassis,omitempty"`
		InventoryDeviceInfo  *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
		ManagementController *ManagementControllerRelationship    `json:"ManagementController,omitempty"`
		RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varManagementInterfaceWithoutEmbeddedStruct := ManagementInterfaceWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varManagementInterfaceWithoutEmbeddedStruct)
	if err == nil {
		varManagementInterface := _ManagementInterface{}
		varManagementInterface.ClassId = varManagementInterfaceWithoutEmbeddedStruct.ClassId
		varManagementInterface.ObjectType = varManagementInterfaceWithoutEmbeddedStruct.ObjectType
		varManagementInterface.Gateway = varManagementInterfaceWithoutEmbeddedStruct.Gateway
		varManagementInterface.HostName = varManagementInterfaceWithoutEmbeddedStruct.HostName
		varManagementInterface.IpAddress = varManagementInterfaceWithoutEmbeddedStruct.IpAddress
		varManagementInterface.Ipv4Address = varManagementInterfaceWithoutEmbeddedStruct.Ipv4Address
		varManagementInterface.Ipv4Gateway = varManagementInterfaceWithoutEmbeddedStruct.Ipv4Gateway
		varManagementInterface.Ipv4Mask = varManagementInterfaceWithoutEmbeddedStruct.Ipv4Mask
		varManagementInterface.Ipv6Address = varManagementInterfaceWithoutEmbeddedStruct.Ipv6Address
		varManagementInterface.Ipv6Gateway = varManagementInterfaceWithoutEmbeddedStruct.Ipv6Gateway
		varManagementInterface.Ipv6Prefix = varManagementInterfaceWithoutEmbeddedStruct.Ipv6Prefix
		varManagementInterface.MacAddress = varManagementInterfaceWithoutEmbeddedStruct.MacAddress
		varManagementInterface.Mask = varManagementInterfaceWithoutEmbeddedStruct.Mask
		varManagementInterface.SwitchId = varManagementInterfaceWithoutEmbeddedStruct.SwitchId
		varManagementInterface.UemConnStatus = varManagementInterfaceWithoutEmbeddedStruct.UemConnStatus
		varManagementInterface.VirtualHostName = varManagementInterfaceWithoutEmbeddedStruct.VirtualHostName
		varManagementInterface.VlanId = varManagementInterfaceWithoutEmbeddedStruct.VlanId
		varManagementInterface.EquipmentChassis = varManagementInterfaceWithoutEmbeddedStruct.EquipmentChassis
		varManagementInterface.InventoryDeviceInfo = varManagementInterfaceWithoutEmbeddedStruct.InventoryDeviceInfo
		varManagementInterface.ManagementController = varManagementInterfaceWithoutEmbeddedStruct.ManagementController
		varManagementInterface.RegisteredDevice = varManagementInterfaceWithoutEmbeddedStruct.RegisteredDevice
		*o = ManagementInterface(varManagementInterface)
	} else {
		return err
	}

	varManagementInterface := _ManagementInterface{}

	err = json.Unmarshal(bytes, &varManagementInterface)
	if err == nil {
		o.InventoryBase = varManagementInterface.InventoryBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Gateway")
		delete(additionalProperties, "HostName")
		delete(additionalProperties, "IpAddress")
		delete(additionalProperties, "Ipv4Address")
		delete(additionalProperties, "Ipv4Gateway")
		delete(additionalProperties, "Ipv4Mask")
		delete(additionalProperties, "Ipv6Address")
		delete(additionalProperties, "Ipv6Gateway")
		delete(additionalProperties, "Ipv6Prefix")
		delete(additionalProperties, "MacAddress")
		delete(additionalProperties, "Mask")
		delete(additionalProperties, "SwitchId")
		delete(additionalProperties, "UemConnStatus")
		delete(additionalProperties, "VirtualHostName")
		delete(additionalProperties, "VlanId")
		delete(additionalProperties, "EquipmentChassis")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "ManagementController")
		delete(additionalProperties, "RegisteredDevice")

		// remove fields from embedded structs
		reflectInventoryBase := reflect.ValueOf(o.InventoryBase)
		for i := 0; i < reflectInventoryBase.Type().NumField(); i++ {
			t := reflectInventoryBase.Type().Field(i)

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

type NullableManagementInterface struct {
	value *ManagementInterface
	isSet bool
}

func (v NullableManagementInterface) Get() *ManagementInterface {
	return v.value
}

func (v *NullableManagementInterface) Set(val *ManagementInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableManagementInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableManagementInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagementInterface(val *ManagementInterface) *NullableManagementInterface {
	return &NullableManagementInterface{value: val, isSet: true}
}

func (v NullableManagementInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagementInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
