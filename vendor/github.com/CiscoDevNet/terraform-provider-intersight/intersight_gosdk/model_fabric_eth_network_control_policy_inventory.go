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

// FabricEthNetworkControlPolicyInventory The features that are applied on a vethernet that is connected to the vNIC.
type FabricEthNetworkControlPolicyInventory struct {
	PolicyAbstractPolicyInventory
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Enables the CDP on an interface.
	CdpEnabled *bool `json:"CdpEnabled,omitempty"`
	// Determines if the MAC forging is allowed or denied on an interface. * `allow` - Allows mac forging on an interface. * `deny` - Denies mac forging on an interface.
	ForgeMac     *string                    `json:"ForgeMac,omitempty"`
	LldpSettings NullableFabricLldpSettings `json:"LldpSettings,omitempty"`
	// Determines the MAC addresses that have to be registered with the switch. * `nativeVlanOnly` - Register only the MAC addresses learnt on the native VLAN. * `allVlans` - Register all the MAC addresses learnt on all the allowed VLANs.
	MacRegistrationMode *string `json:"MacRegistrationMode,omitempty"`
	// Determines the state of the virtual interface (vethernet / vfc) on the switch when a suitable uplink is not pinned. * `linkDown` - The vethernet will go down in case a suitable uplink is not pinned. * `warning` - The vethernet will remain up even if a suitable uplink is not pinned.
	UplinkFailAction *string `json:"UplinkFailAction,omitempty"`
	// An array of relationships to vnicEthNetworkPolicyInventory resources.
	// Deprecated
	NetworkPolicy        []VnicEthNetworkPolicyInventoryRelationship `json:"NetworkPolicy,omitempty"`
	TargetMo             *MoBaseMoRelationship                       `json:"TargetMo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricEthNetworkControlPolicyInventory FabricEthNetworkControlPolicyInventory

// NewFabricEthNetworkControlPolicyInventory instantiates a new FabricEthNetworkControlPolicyInventory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricEthNetworkControlPolicyInventory(classId string, objectType string) *FabricEthNetworkControlPolicyInventory {
	this := FabricEthNetworkControlPolicyInventory{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewFabricEthNetworkControlPolicyInventoryWithDefaults instantiates a new FabricEthNetworkControlPolicyInventory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricEthNetworkControlPolicyInventoryWithDefaults() *FabricEthNetworkControlPolicyInventory {
	this := FabricEthNetworkControlPolicyInventory{}
	var classId string = "fabric.EthNetworkControlPolicyInventory"
	this.ClassId = classId
	var objectType string = "fabric.EthNetworkControlPolicyInventory"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *FabricEthNetworkControlPolicyInventory) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FabricEthNetworkControlPolicyInventory) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FabricEthNetworkControlPolicyInventory) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FabricEthNetworkControlPolicyInventory) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FabricEthNetworkControlPolicyInventory) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FabricEthNetworkControlPolicyInventory) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCdpEnabled returns the CdpEnabled field value if set, zero value otherwise.
func (o *FabricEthNetworkControlPolicyInventory) GetCdpEnabled() bool {
	if o == nil || o.CdpEnabled == nil {
		var ret bool
		return ret
	}
	return *o.CdpEnabled
}

// GetCdpEnabledOk returns a tuple with the CdpEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricEthNetworkControlPolicyInventory) GetCdpEnabledOk() (*bool, bool) {
	if o == nil || o.CdpEnabled == nil {
		return nil, false
	}
	return o.CdpEnabled, true
}

// HasCdpEnabled returns a boolean if a field has been set.
func (o *FabricEthNetworkControlPolicyInventory) HasCdpEnabled() bool {
	if o != nil && o.CdpEnabled != nil {
		return true
	}

	return false
}

// SetCdpEnabled gets a reference to the given bool and assigns it to the CdpEnabled field.
func (o *FabricEthNetworkControlPolicyInventory) SetCdpEnabled(v bool) {
	o.CdpEnabled = &v
}

// GetForgeMac returns the ForgeMac field value if set, zero value otherwise.
func (o *FabricEthNetworkControlPolicyInventory) GetForgeMac() string {
	if o == nil || o.ForgeMac == nil {
		var ret string
		return ret
	}
	return *o.ForgeMac
}

// GetForgeMacOk returns a tuple with the ForgeMac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricEthNetworkControlPolicyInventory) GetForgeMacOk() (*string, bool) {
	if o == nil || o.ForgeMac == nil {
		return nil, false
	}
	return o.ForgeMac, true
}

// HasForgeMac returns a boolean if a field has been set.
func (o *FabricEthNetworkControlPolicyInventory) HasForgeMac() bool {
	if o != nil && o.ForgeMac != nil {
		return true
	}

	return false
}

// SetForgeMac gets a reference to the given string and assigns it to the ForgeMac field.
func (o *FabricEthNetworkControlPolicyInventory) SetForgeMac(v string) {
	o.ForgeMac = &v
}

// GetLldpSettings returns the LldpSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FabricEthNetworkControlPolicyInventory) GetLldpSettings() FabricLldpSettings {
	if o == nil || o.LldpSettings.Get() == nil {
		var ret FabricLldpSettings
		return ret
	}
	return *o.LldpSettings.Get()
}

// GetLldpSettingsOk returns a tuple with the LldpSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FabricEthNetworkControlPolicyInventory) GetLldpSettingsOk() (*FabricLldpSettings, bool) {
	if o == nil {
		return nil, false
	}
	return o.LldpSettings.Get(), o.LldpSettings.IsSet()
}

// HasLldpSettings returns a boolean if a field has been set.
func (o *FabricEthNetworkControlPolicyInventory) HasLldpSettings() bool {
	if o != nil && o.LldpSettings.IsSet() {
		return true
	}

	return false
}

// SetLldpSettings gets a reference to the given NullableFabricLldpSettings and assigns it to the LldpSettings field.
func (o *FabricEthNetworkControlPolicyInventory) SetLldpSettings(v FabricLldpSettings) {
	o.LldpSettings.Set(&v)
}

// SetLldpSettingsNil sets the value for LldpSettings to be an explicit nil
func (o *FabricEthNetworkControlPolicyInventory) SetLldpSettingsNil() {
	o.LldpSettings.Set(nil)
}

// UnsetLldpSettings ensures that no value is present for LldpSettings, not even an explicit nil
func (o *FabricEthNetworkControlPolicyInventory) UnsetLldpSettings() {
	o.LldpSettings.Unset()
}

// GetMacRegistrationMode returns the MacRegistrationMode field value if set, zero value otherwise.
func (o *FabricEthNetworkControlPolicyInventory) GetMacRegistrationMode() string {
	if o == nil || o.MacRegistrationMode == nil {
		var ret string
		return ret
	}
	return *o.MacRegistrationMode
}

// GetMacRegistrationModeOk returns a tuple with the MacRegistrationMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricEthNetworkControlPolicyInventory) GetMacRegistrationModeOk() (*string, bool) {
	if o == nil || o.MacRegistrationMode == nil {
		return nil, false
	}
	return o.MacRegistrationMode, true
}

// HasMacRegistrationMode returns a boolean if a field has been set.
func (o *FabricEthNetworkControlPolicyInventory) HasMacRegistrationMode() bool {
	if o != nil && o.MacRegistrationMode != nil {
		return true
	}

	return false
}

// SetMacRegistrationMode gets a reference to the given string and assigns it to the MacRegistrationMode field.
func (o *FabricEthNetworkControlPolicyInventory) SetMacRegistrationMode(v string) {
	o.MacRegistrationMode = &v
}

// GetUplinkFailAction returns the UplinkFailAction field value if set, zero value otherwise.
func (o *FabricEthNetworkControlPolicyInventory) GetUplinkFailAction() string {
	if o == nil || o.UplinkFailAction == nil {
		var ret string
		return ret
	}
	return *o.UplinkFailAction
}

// GetUplinkFailActionOk returns a tuple with the UplinkFailAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricEthNetworkControlPolicyInventory) GetUplinkFailActionOk() (*string, bool) {
	if o == nil || o.UplinkFailAction == nil {
		return nil, false
	}
	return o.UplinkFailAction, true
}

// HasUplinkFailAction returns a boolean if a field has been set.
func (o *FabricEthNetworkControlPolicyInventory) HasUplinkFailAction() bool {
	if o != nil && o.UplinkFailAction != nil {
		return true
	}

	return false
}

// SetUplinkFailAction gets a reference to the given string and assigns it to the UplinkFailAction field.
func (o *FabricEthNetworkControlPolicyInventory) SetUplinkFailAction(v string) {
	o.UplinkFailAction = &v
}

// GetNetworkPolicy returns the NetworkPolicy field value if set, zero value otherwise (both if not set or set to explicit null).
// Deprecated
func (o *FabricEthNetworkControlPolicyInventory) GetNetworkPolicy() []VnicEthNetworkPolicyInventoryRelationship {
	if o == nil {
		var ret []VnicEthNetworkPolicyInventoryRelationship
		return ret
	}
	return o.NetworkPolicy
}

// GetNetworkPolicyOk returns a tuple with the NetworkPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
// Deprecated
func (o *FabricEthNetworkControlPolicyInventory) GetNetworkPolicyOk() ([]VnicEthNetworkPolicyInventoryRelationship, bool) {
	if o == nil || o.NetworkPolicy == nil {
		return nil, false
	}
	return o.NetworkPolicy, true
}

// HasNetworkPolicy returns a boolean if a field has been set.
func (o *FabricEthNetworkControlPolicyInventory) HasNetworkPolicy() bool {
	if o != nil && o.NetworkPolicy != nil {
		return true
	}

	return false
}

// SetNetworkPolicy gets a reference to the given []VnicEthNetworkPolicyInventoryRelationship and assigns it to the NetworkPolicy field.
// Deprecated
func (o *FabricEthNetworkControlPolicyInventory) SetNetworkPolicy(v []VnicEthNetworkPolicyInventoryRelationship) {
	o.NetworkPolicy = v
}

// GetTargetMo returns the TargetMo field value if set, zero value otherwise.
func (o *FabricEthNetworkControlPolicyInventory) GetTargetMo() MoBaseMoRelationship {
	if o == nil || o.TargetMo == nil {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.TargetMo
}

// GetTargetMoOk returns a tuple with the TargetMo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricEthNetworkControlPolicyInventory) GetTargetMoOk() (*MoBaseMoRelationship, bool) {
	if o == nil || o.TargetMo == nil {
		return nil, false
	}
	return o.TargetMo, true
}

// HasTargetMo returns a boolean if a field has been set.
func (o *FabricEthNetworkControlPolicyInventory) HasTargetMo() bool {
	if o != nil && o.TargetMo != nil {
		return true
	}

	return false
}

// SetTargetMo gets a reference to the given MoBaseMoRelationship and assigns it to the TargetMo field.
func (o *FabricEthNetworkControlPolicyInventory) SetTargetMo(v MoBaseMoRelationship) {
	o.TargetMo = &v
}

func (o FabricEthNetworkControlPolicyInventory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractPolicyInventory, errPolicyAbstractPolicyInventory := json.Marshal(o.PolicyAbstractPolicyInventory)
	if errPolicyAbstractPolicyInventory != nil {
		return []byte{}, errPolicyAbstractPolicyInventory
	}
	errPolicyAbstractPolicyInventory = json.Unmarshal([]byte(serializedPolicyAbstractPolicyInventory), &toSerialize)
	if errPolicyAbstractPolicyInventory != nil {
		return []byte{}, errPolicyAbstractPolicyInventory
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.CdpEnabled != nil {
		toSerialize["CdpEnabled"] = o.CdpEnabled
	}
	if o.ForgeMac != nil {
		toSerialize["ForgeMac"] = o.ForgeMac
	}
	if o.LldpSettings.IsSet() {
		toSerialize["LldpSettings"] = o.LldpSettings.Get()
	}
	if o.MacRegistrationMode != nil {
		toSerialize["MacRegistrationMode"] = o.MacRegistrationMode
	}
	if o.UplinkFailAction != nil {
		toSerialize["UplinkFailAction"] = o.UplinkFailAction
	}
	if o.NetworkPolicy != nil {
		toSerialize["NetworkPolicy"] = o.NetworkPolicy
	}
	if o.TargetMo != nil {
		toSerialize["TargetMo"] = o.TargetMo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FabricEthNetworkControlPolicyInventory) UnmarshalJSON(bytes []byte) (err error) {
	type FabricEthNetworkControlPolicyInventoryWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Enables the CDP on an interface.
		CdpEnabled *bool `json:"CdpEnabled,omitempty"`
		// Determines if the MAC forging is allowed or denied on an interface. * `allow` - Allows mac forging on an interface. * `deny` - Denies mac forging on an interface.
		ForgeMac     *string                    `json:"ForgeMac,omitempty"`
		LldpSettings NullableFabricLldpSettings `json:"LldpSettings,omitempty"`
		// Determines the MAC addresses that have to be registered with the switch. * `nativeVlanOnly` - Register only the MAC addresses learnt on the native VLAN. * `allVlans` - Register all the MAC addresses learnt on all the allowed VLANs.
		MacRegistrationMode *string `json:"MacRegistrationMode,omitempty"`
		// Determines the state of the virtual interface (vethernet / vfc) on the switch when a suitable uplink is not pinned. * `linkDown` - The vethernet will go down in case a suitable uplink is not pinned. * `warning` - The vethernet will remain up even if a suitable uplink is not pinned.
		UplinkFailAction *string `json:"UplinkFailAction,omitempty"`
		// An array of relationships to vnicEthNetworkPolicyInventory resources.
		// Deprecated
		NetworkPolicy []VnicEthNetworkPolicyInventoryRelationship `json:"NetworkPolicy,omitempty"`
		TargetMo      *MoBaseMoRelationship                       `json:"TargetMo,omitempty"`
	}

	varFabricEthNetworkControlPolicyInventoryWithoutEmbeddedStruct := FabricEthNetworkControlPolicyInventoryWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varFabricEthNetworkControlPolicyInventoryWithoutEmbeddedStruct)
	if err == nil {
		varFabricEthNetworkControlPolicyInventory := _FabricEthNetworkControlPolicyInventory{}
		varFabricEthNetworkControlPolicyInventory.ClassId = varFabricEthNetworkControlPolicyInventoryWithoutEmbeddedStruct.ClassId
		varFabricEthNetworkControlPolicyInventory.ObjectType = varFabricEthNetworkControlPolicyInventoryWithoutEmbeddedStruct.ObjectType
		varFabricEthNetworkControlPolicyInventory.CdpEnabled = varFabricEthNetworkControlPolicyInventoryWithoutEmbeddedStruct.CdpEnabled
		varFabricEthNetworkControlPolicyInventory.ForgeMac = varFabricEthNetworkControlPolicyInventoryWithoutEmbeddedStruct.ForgeMac
		varFabricEthNetworkControlPolicyInventory.LldpSettings = varFabricEthNetworkControlPolicyInventoryWithoutEmbeddedStruct.LldpSettings
		varFabricEthNetworkControlPolicyInventory.MacRegistrationMode = varFabricEthNetworkControlPolicyInventoryWithoutEmbeddedStruct.MacRegistrationMode
		varFabricEthNetworkControlPolicyInventory.UplinkFailAction = varFabricEthNetworkControlPolicyInventoryWithoutEmbeddedStruct.UplinkFailAction
		varFabricEthNetworkControlPolicyInventory.NetworkPolicy = varFabricEthNetworkControlPolicyInventoryWithoutEmbeddedStruct.NetworkPolicy
		varFabricEthNetworkControlPolicyInventory.TargetMo = varFabricEthNetworkControlPolicyInventoryWithoutEmbeddedStruct.TargetMo
		*o = FabricEthNetworkControlPolicyInventory(varFabricEthNetworkControlPolicyInventory)
	} else {
		return err
	}

	varFabricEthNetworkControlPolicyInventory := _FabricEthNetworkControlPolicyInventory{}

	err = json.Unmarshal(bytes, &varFabricEthNetworkControlPolicyInventory)
	if err == nil {
		o.PolicyAbstractPolicyInventory = varFabricEthNetworkControlPolicyInventory.PolicyAbstractPolicyInventory
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CdpEnabled")
		delete(additionalProperties, "ForgeMac")
		delete(additionalProperties, "LldpSettings")
		delete(additionalProperties, "MacRegistrationMode")
		delete(additionalProperties, "UplinkFailAction")
		delete(additionalProperties, "NetworkPolicy")
		delete(additionalProperties, "TargetMo")

		// remove fields from embedded structs
		reflectPolicyAbstractPolicyInventory := reflect.ValueOf(o.PolicyAbstractPolicyInventory)
		for i := 0; i < reflectPolicyAbstractPolicyInventory.Type().NumField(); i++ {
			t := reflectPolicyAbstractPolicyInventory.Type().Field(i)

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

type NullableFabricEthNetworkControlPolicyInventory struct {
	value *FabricEthNetworkControlPolicyInventory
	isSet bool
}

func (v NullableFabricEthNetworkControlPolicyInventory) Get() *FabricEthNetworkControlPolicyInventory {
	return v.value
}

func (v *NullableFabricEthNetworkControlPolicyInventory) Set(val *FabricEthNetworkControlPolicyInventory) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricEthNetworkControlPolicyInventory) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricEthNetworkControlPolicyInventory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricEthNetworkControlPolicyInventory(val *FabricEthNetworkControlPolicyInventory) *NullableFabricEthNetworkControlPolicyInventory {
	return &NullableFabricEthNetworkControlPolicyInventory{value: val, isSet: true}
}

func (v NullableFabricEthNetworkControlPolicyInventory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricEthNetworkControlPolicyInventory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
