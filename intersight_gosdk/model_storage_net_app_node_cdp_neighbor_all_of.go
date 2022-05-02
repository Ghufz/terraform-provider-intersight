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

// StorageNetAppNodeCdpNeighborAllOf Definition of the list of properties defined in 'storage.NetAppNodeCdpNeighbor', excluding properties defined in parent classes.
type StorageNetAppNodeCdpNeighborAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType   string   `json:"ObjectType"`
	Capabilities []string `json:"Capabilities,omitempty"`
	// The IP address of the CDP neighbor.
	DeviceIp *string `json:"DeviceIp,omitempty"`
	// The name of the CDP neighbor.
	DiscoveredDevice *string `json:"DiscoveredDevice,omitempty"`
	// The period of time for which CDP advertisements are cached.
	HoldTimeRemaining *int64 `json:"HoldTimeRemaining,omitempty"`
	// The interface of the CDP neighbor.
	Interface *string `json:"Interface,omitempty"`
	// The node that owns the port for this CDP neighbor.
	NodeName *string `json:"NodeName,omitempty"`
	// The platform of the CDP neighbor.
	Platform *string `json:"Platform,omitempty"`
	// The port for this CDP neighbor.
	Port *string `json:"Port,omitempty"`
	// The protocol used for CDP.
	Protocol *string `json:"Protocol,omitempty"`
	// The version of the operating system running on the CDP neighbor.
	Version              *string                        `json:"Version,omitempty"`
	ArrayController      *StorageNetAppNodeRelationship `json:"ArrayController,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppNodeCdpNeighborAllOf StorageNetAppNodeCdpNeighborAllOf

// NewStorageNetAppNodeCdpNeighborAllOf instantiates a new StorageNetAppNodeCdpNeighborAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppNodeCdpNeighborAllOf(classId string, objectType string) *StorageNetAppNodeCdpNeighborAllOf {
	this := StorageNetAppNodeCdpNeighborAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppNodeCdpNeighborAllOfWithDefaults instantiates a new StorageNetAppNodeCdpNeighborAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppNodeCdpNeighborAllOfWithDefaults() *StorageNetAppNodeCdpNeighborAllOf {
	this := StorageNetAppNodeCdpNeighborAllOf{}
	var classId string = "storage.NetAppNodeCdpNeighbor"
	this.ClassId = classId
	var objectType string = "storage.NetAppNodeCdpNeighbor"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppNodeCdpNeighborAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppNodeCdpNeighborAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppNodeCdpNeighborAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppNodeCdpNeighborAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppNodeCdpNeighborAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppNodeCdpNeighborAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCapabilities returns the Capabilities field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppNodeCdpNeighborAllOf) GetCapabilities() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppNodeCdpNeighborAllOf) GetCapabilitiesOk() ([]string, bool) {
	if o == nil || o.Capabilities == nil {
		return nil, false
	}
	return o.Capabilities, true
}

// HasCapabilities returns a boolean if a field has been set.
func (o *StorageNetAppNodeCdpNeighborAllOf) HasCapabilities() bool {
	if o != nil && o.Capabilities != nil {
		return true
	}

	return false
}

// SetCapabilities gets a reference to the given []string and assigns it to the Capabilities field.
func (o *StorageNetAppNodeCdpNeighborAllOf) SetCapabilities(v []string) {
	o.Capabilities = v
}

// GetDeviceIp returns the DeviceIp field value if set, zero value otherwise.
func (o *StorageNetAppNodeCdpNeighborAllOf) GetDeviceIp() string {
	if o == nil || o.DeviceIp == nil {
		var ret string
		return ret
	}
	return *o.DeviceIp
}

// GetDeviceIpOk returns a tuple with the DeviceIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppNodeCdpNeighborAllOf) GetDeviceIpOk() (*string, bool) {
	if o == nil || o.DeviceIp == nil {
		return nil, false
	}
	return o.DeviceIp, true
}

// HasDeviceIp returns a boolean if a field has been set.
func (o *StorageNetAppNodeCdpNeighborAllOf) HasDeviceIp() bool {
	if o != nil && o.DeviceIp != nil {
		return true
	}

	return false
}

// SetDeviceIp gets a reference to the given string and assigns it to the DeviceIp field.
func (o *StorageNetAppNodeCdpNeighborAllOf) SetDeviceIp(v string) {
	o.DeviceIp = &v
}

// GetDiscoveredDevice returns the DiscoveredDevice field value if set, zero value otherwise.
func (o *StorageNetAppNodeCdpNeighborAllOf) GetDiscoveredDevice() string {
	if o == nil || o.DiscoveredDevice == nil {
		var ret string
		return ret
	}
	return *o.DiscoveredDevice
}

// GetDiscoveredDeviceOk returns a tuple with the DiscoveredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppNodeCdpNeighborAllOf) GetDiscoveredDeviceOk() (*string, bool) {
	if o == nil || o.DiscoveredDevice == nil {
		return nil, false
	}
	return o.DiscoveredDevice, true
}

// HasDiscoveredDevice returns a boolean if a field has been set.
func (o *StorageNetAppNodeCdpNeighborAllOf) HasDiscoveredDevice() bool {
	if o != nil && o.DiscoveredDevice != nil {
		return true
	}

	return false
}

// SetDiscoveredDevice gets a reference to the given string and assigns it to the DiscoveredDevice field.
func (o *StorageNetAppNodeCdpNeighborAllOf) SetDiscoveredDevice(v string) {
	o.DiscoveredDevice = &v
}

// GetHoldTimeRemaining returns the HoldTimeRemaining field value if set, zero value otherwise.
func (o *StorageNetAppNodeCdpNeighborAllOf) GetHoldTimeRemaining() int64 {
	if o == nil || o.HoldTimeRemaining == nil {
		var ret int64
		return ret
	}
	return *o.HoldTimeRemaining
}

// GetHoldTimeRemainingOk returns a tuple with the HoldTimeRemaining field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppNodeCdpNeighborAllOf) GetHoldTimeRemainingOk() (*int64, bool) {
	if o == nil || o.HoldTimeRemaining == nil {
		return nil, false
	}
	return o.HoldTimeRemaining, true
}

// HasHoldTimeRemaining returns a boolean if a field has been set.
func (o *StorageNetAppNodeCdpNeighborAllOf) HasHoldTimeRemaining() bool {
	if o != nil && o.HoldTimeRemaining != nil {
		return true
	}

	return false
}

// SetHoldTimeRemaining gets a reference to the given int64 and assigns it to the HoldTimeRemaining field.
func (o *StorageNetAppNodeCdpNeighborAllOf) SetHoldTimeRemaining(v int64) {
	o.HoldTimeRemaining = &v
}

// GetInterface returns the Interface field value if set, zero value otherwise.
func (o *StorageNetAppNodeCdpNeighborAllOf) GetInterface() string {
	if o == nil || o.Interface == nil {
		var ret string
		return ret
	}
	return *o.Interface
}

// GetInterfaceOk returns a tuple with the Interface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppNodeCdpNeighborAllOf) GetInterfaceOk() (*string, bool) {
	if o == nil || o.Interface == nil {
		return nil, false
	}
	return o.Interface, true
}

// HasInterface returns a boolean if a field has been set.
func (o *StorageNetAppNodeCdpNeighborAllOf) HasInterface() bool {
	if o != nil && o.Interface != nil {
		return true
	}

	return false
}

// SetInterface gets a reference to the given string and assigns it to the Interface field.
func (o *StorageNetAppNodeCdpNeighborAllOf) SetInterface(v string) {
	o.Interface = &v
}

// GetNodeName returns the NodeName field value if set, zero value otherwise.
func (o *StorageNetAppNodeCdpNeighborAllOf) GetNodeName() string {
	if o == nil || o.NodeName == nil {
		var ret string
		return ret
	}
	return *o.NodeName
}

// GetNodeNameOk returns a tuple with the NodeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppNodeCdpNeighborAllOf) GetNodeNameOk() (*string, bool) {
	if o == nil || o.NodeName == nil {
		return nil, false
	}
	return o.NodeName, true
}

// HasNodeName returns a boolean if a field has been set.
func (o *StorageNetAppNodeCdpNeighborAllOf) HasNodeName() bool {
	if o != nil && o.NodeName != nil {
		return true
	}

	return false
}

// SetNodeName gets a reference to the given string and assigns it to the NodeName field.
func (o *StorageNetAppNodeCdpNeighborAllOf) SetNodeName(v string) {
	o.NodeName = &v
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *StorageNetAppNodeCdpNeighborAllOf) GetPlatform() string {
	if o == nil || o.Platform == nil {
		var ret string
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppNodeCdpNeighborAllOf) GetPlatformOk() (*string, bool) {
	if o == nil || o.Platform == nil {
		return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *StorageNetAppNodeCdpNeighborAllOf) HasPlatform() bool {
	if o != nil && o.Platform != nil {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given string and assigns it to the Platform field.
func (o *StorageNetAppNodeCdpNeighborAllOf) SetPlatform(v string) {
	o.Platform = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *StorageNetAppNodeCdpNeighborAllOf) GetPort() string {
	if o == nil || o.Port == nil {
		var ret string
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppNodeCdpNeighborAllOf) GetPortOk() (*string, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *StorageNetAppNodeCdpNeighborAllOf) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given string and assigns it to the Port field.
func (o *StorageNetAppNodeCdpNeighborAllOf) SetPort(v string) {
	o.Port = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *StorageNetAppNodeCdpNeighborAllOf) GetProtocol() string {
	if o == nil || o.Protocol == nil {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppNodeCdpNeighborAllOf) GetProtocolOk() (*string, bool) {
	if o == nil || o.Protocol == nil {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *StorageNetAppNodeCdpNeighborAllOf) HasProtocol() bool {
	if o != nil && o.Protocol != nil {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *StorageNetAppNodeCdpNeighborAllOf) SetProtocol(v string) {
	o.Protocol = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *StorageNetAppNodeCdpNeighborAllOf) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppNodeCdpNeighborAllOf) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *StorageNetAppNodeCdpNeighborAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *StorageNetAppNodeCdpNeighborAllOf) SetVersion(v string) {
	o.Version = &v
}

// GetArrayController returns the ArrayController field value if set, zero value otherwise.
func (o *StorageNetAppNodeCdpNeighborAllOf) GetArrayController() StorageNetAppNodeRelationship {
	if o == nil || o.ArrayController == nil {
		var ret StorageNetAppNodeRelationship
		return ret
	}
	return *o.ArrayController
}

// GetArrayControllerOk returns a tuple with the ArrayController field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppNodeCdpNeighborAllOf) GetArrayControllerOk() (*StorageNetAppNodeRelationship, bool) {
	if o == nil || o.ArrayController == nil {
		return nil, false
	}
	return o.ArrayController, true
}

// HasArrayController returns a boolean if a field has been set.
func (o *StorageNetAppNodeCdpNeighborAllOf) HasArrayController() bool {
	if o != nil && o.ArrayController != nil {
		return true
	}

	return false
}

// SetArrayController gets a reference to the given StorageNetAppNodeRelationship and assigns it to the ArrayController field.
func (o *StorageNetAppNodeCdpNeighborAllOf) SetArrayController(v StorageNetAppNodeRelationship) {
	o.ArrayController = &v
}

func (o StorageNetAppNodeCdpNeighborAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Capabilities != nil {
		toSerialize["Capabilities"] = o.Capabilities
	}
	if o.DeviceIp != nil {
		toSerialize["DeviceIp"] = o.DeviceIp
	}
	if o.DiscoveredDevice != nil {
		toSerialize["DiscoveredDevice"] = o.DiscoveredDevice
	}
	if o.HoldTimeRemaining != nil {
		toSerialize["HoldTimeRemaining"] = o.HoldTimeRemaining
	}
	if o.Interface != nil {
		toSerialize["Interface"] = o.Interface
	}
	if o.NodeName != nil {
		toSerialize["NodeName"] = o.NodeName
	}
	if o.Platform != nil {
		toSerialize["Platform"] = o.Platform
	}
	if o.Port != nil {
		toSerialize["Port"] = o.Port
	}
	if o.Protocol != nil {
		toSerialize["Protocol"] = o.Protocol
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	if o.ArrayController != nil {
		toSerialize["ArrayController"] = o.ArrayController
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageNetAppNodeCdpNeighborAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageNetAppNodeCdpNeighborAllOf := _StorageNetAppNodeCdpNeighborAllOf{}

	if err = json.Unmarshal(bytes, &varStorageNetAppNodeCdpNeighborAllOf); err == nil {
		*o = StorageNetAppNodeCdpNeighborAllOf(varStorageNetAppNodeCdpNeighborAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Capabilities")
		delete(additionalProperties, "DeviceIp")
		delete(additionalProperties, "DiscoveredDevice")
		delete(additionalProperties, "HoldTimeRemaining")
		delete(additionalProperties, "Interface")
		delete(additionalProperties, "NodeName")
		delete(additionalProperties, "Platform")
		delete(additionalProperties, "Port")
		delete(additionalProperties, "Protocol")
		delete(additionalProperties, "Version")
		delete(additionalProperties, "ArrayController")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageNetAppNodeCdpNeighborAllOf struct {
	value *StorageNetAppNodeCdpNeighborAllOf
	isSet bool
}

func (v NullableStorageNetAppNodeCdpNeighborAllOf) Get() *StorageNetAppNodeCdpNeighborAllOf {
	return v.value
}

func (v *NullableStorageNetAppNodeCdpNeighborAllOf) Set(val *StorageNetAppNodeCdpNeighborAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppNodeCdpNeighborAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppNodeCdpNeighborAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppNodeCdpNeighborAllOf(val *StorageNetAppNodeCdpNeighborAllOf) *NullableStorageNetAppNodeCdpNeighborAllOf {
	return &NullableStorageNetAppNodeCdpNeighborAllOf{value: val, isSet: true}
}

func (v NullableStorageNetAppNodeCdpNeighborAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppNodeCdpNeighborAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
