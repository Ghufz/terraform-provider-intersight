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

// VnicSanConnectivityPolicyInventory SAN connectivity policy determines the network storage resources and the connections between the server and the SAN on the network. This policy enables configuration of vHBAs that the servers use to communicate with the storage network.
type VnicSanConnectivityPolicyInventory struct {
	PolicyAbstractPolicyInventory
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The mode used for placement of vNICs on network adapters. It can either be Auto or Custom. * `custom` - The placement of the vNICs / vHBAs on network adapters is manually chosen by the user. * `auto` - The placement of the vNICs / vHBAs on network adapters is automatically determined by the system.
	PlacementMode *string `json:"PlacementMode,omitempty"`
	// The WWNN address for the server node must be in hexadecimal format xx:xx:xx:xx:xx:xx:xx:xx. Allowed ranges are 20:00:00:00:00:00:00:00 to 20:FF:FF:FF:FF:FF:FF:FF or from 50:00:00:00:00:00:00:00 to 5F:FF:FF:FF:FF:FF:FF:FF. To ensure uniqueness of WWN's in the SAN fabric, you are strongly encouraged to use the WWN prefix - 20:00:00:25:B5:xx:xx:xx.
	StaticWwnnAddress *string `json:"StaticWwnnAddress,omitempty"`
	// The platform for which the server profile is applicable. It can either be a server that is operating in standalone mode or which is attached to a Fabric Interconnect managed by Intersight. * `Standalone` - Servers which are operating in standalone mode i.e. not connected to a Fabric Interconnected. * `FIAttached` - Servers which are connected to a Fabric Interconnect that is managed by Intersight.
	TargetPlatform *string `json:"TargetPlatform,omitempty"`
	// Type of allocation selected to assign a WWNN address for the server node. * `POOL` - The user selects a pool from which the mac/wwn address will be leased for the Virtual Interface. * `STATIC` - The user assigns a static mac/wwn address for the Virtual Interface.
	WwnnAddressType *string `json:"WwnnAddressType,omitempty"`
	// An array of relationships to vnicFcIfInventory resources.
	FcIfs                []VnicFcIfInventoryRelationship `json:"FcIfs,omitempty"`
	TargetMo             *MoBaseMoRelationship           `json:"TargetMo,omitempty"`
	WwnnPool             *FcpoolPoolRelationship         `json:"WwnnPool,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VnicSanConnectivityPolicyInventory VnicSanConnectivityPolicyInventory

// NewVnicSanConnectivityPolicyInventory instantiates a new VnicSanConnectivityPolicyInventory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicSanConnectivityPolicyInventory(classId string, objectType string) *VnicSanConnectivityPolicyInventory {
	this := VnicSanConnectivityPolicyInventory{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVnicSanConnectivityPolicyInventoryWithDefaults instantiates a new VnicSanConnectivityPolicyInventory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicSanConnectivityPolicyInventoryWithDefaults() *VnicSanConnectivityPolicyInventory {
	this := VnicSanConnectivityPolicyInventory{}
	var classId string = "vnic.SanConnectivityPolicyInventory"
	this.ClassId = classId
	var objectType string = "vnic.SanConnectivityPolicyInventory"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VnicSanConnectivityPolicyInventory) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VnicSanConnectivityPolicyInventory) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VnicSanConnectivityPolicyInventory) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VnicSanConnectivityPolicyInventory) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VnicSanConnectivityPolicyInventory) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VnicSanConnectivityPolicyInventory) SetObjectType(v string) {
	o.ObjectType = v
}

// GetPlacementMode returns the PlacementMode field value if set, zero value otherwise.
func (o *VnicSanConnectivityPolicyInventory) GetPlacementMode() string {
	if o == nil || o.PlacementMode == nil {
		var ret string
		return ret
	}
	return *o.PlacementMode
}

// GetPlacementModeOk returns a tuple with the PlacementMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicSanConnectivityPolicyInventory) GetPlacementModeOk() (*string, bool) {
	if o == nil || o.PlacementMode == nil {
		return nil, false
	}
	return o.PlacementMode, true
}

// HasPlacementMode returns a boolean if a field has been set.
func (o *VnicSanConnectivityPolicyInventory) HasPlacementMode() bool {
	if o != nil && o.PlacementMode != nil {
		return true
	}

	return false
}

// SetPlacementMode gets a reference to the given string and assigns it to the PlacementMode field.
func (o *VnicSanConnectivityPolicyInventory) SetPlacementMode(v string) {
	o.PlacementMode = &v
}

// GetStaticWwnnAddress returns the StaticWwnnAddress field value if set, zero value otherwise.
func (o *VnicSanConnectivityPolicyInventory) GetStaticWwnnAddress() string {
	if o == nil || o.StaticWwnnAddress == nil {
		var ret string
		return ret
	}
	return *o.StaticWwnnAddress
}

// GetStaticWwnnAddressOk returns a tuple with the StaticWwnnAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicSanConnectivityPolicyInventory) GetStaticWwnnAddressOk() (*string, bool) {
	if o == nil || o.StaticWwnnAddress == nil {
		return nil, false
	}
	return o.StaticWwnnAddress, true
}

// HasStaticWwnnAddress returns a boolean if a field has been set.
func (o *VnicSanConnectivityPolicyInventory) HasStaticWwnnAddress() bool {
	if o != nil && o.StaticWwnnAddress != nil {
		return true
	}

	return false
}

// SetStaticWwnnAddress gets a reference to the given string and assigns it to the StaticWwnnAddress field.
func (o *VnicSanConnectivityPolicyInventory) SetStaticWwnnAddress(v string) {
	o.StaticWwnnAddress = &v
}

// GetTargetPlatform returns the TargetPlatform field value if set, zero value otherwise.
func (o *VnicSanConnectivityPolicyInventory) GetTargetPlatform() string {
	if o == nil || o.TargetPlatform == nil {
		var ret string
		return ret
	}
	return *o.TargetPlatform
}

// GetTargetPlatformOk returns a tuple with the TargetPlatform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicSanConnectivityPolicyInventory) GetTargetPlatformOk() (*string, bool) {
	if o == nil || o.TargetPlatform == nil {
		return nil, false
	}
	return o.TargetPlatform, true
}

// HasTargetPlatform returns a boolean if a field has been set.
func (o *VnicSanConnectivityPolicyInventory) HasTargetPlatform() bool {
	if o != nil && o.TargetPlatform != nil {
		return true
	}

	return false
}

// SetTargetPlatform gets a reference to the given string and assigns it to the TargetPlatform field.
func (o *VnicSanConnectivityPolicyInventory) SetTargetPlatform(v string) {
	o.TargetPlatform = &v
}

// GetWwnnAddressType returns the WwnnAddressType field value if set, zero value otherwise.
func (o *VnicSanConnectivityPolicyInventory) GetWwnnAddressType() string {
	if o == nil || o.WwnnAddressType == nil {
		var ret string
		return ret
	}
	return *o.WwnnAddressType
}

// GetWwnnAddressTypeOk returns a tuple with the WwnnAddressType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicSanConnectivityPolicyInventory) GetWwnnAddressTypeOk() (*string, bool) {
	if o == nil || o.WwnnAddressType == nil {
		return nil, false
	}
	return o.WwnnAddressType, true
}

// HasWwnnAddressType returns a boolean if a field has been set.
func (o *VnicSanConnectivityPolicyInventory) HasWwnnAddressType() bool {
	if o != nil && o.WwnnAddressType != nil {
		return true
	}

	return false
}

// SetWwnnAddressType gets a reference to the given string and assigns it to the WwnnAddressType field.
func (o *VnicSanConnectivityPolicyInventory) SetWwnnAddressType(v string) {
	o.WwnnAddressType = &v
}

// GetFcIfs returns the FcIfs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicSanConnectivityPolicyInventory) GetFcIfs() []VnicFcIfInventoryRelationship {
	if o == nil {
		var ret []VnicFcIfInventoryRelationship
		return ret
	}
	return o.FcIfs
}

// GetFcIfsOk returns a tuple with the FcIfs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicSanConnectivityPolicyInventory) GetFcIfsOk() ([]VnicFcIfInventoryRelationship, bool) {
	if o == nil || o.FcIfs == nil {
		return nil, false
	}
	return o.FcIfs, true
}

// HasFcIfs returns a boolean if a field has been set.
func (o *VnicSanConnectivityPolicyInventory) HasFcIfs() bool {
	if o != nil && o.FcIfs != nil {
		return true
	}

	return false
}

// SetFcIfs gets a reference to the given []VnicFcIfInventoryRelationship and assigns it to the FcIfs field.
func (o *VnicSanConnectivityPolicyInventory) SetFcIfs(v []VnicFcIfInventoryRelationship) {
	o.FcIfs = v
}

// GetTargetMo returns the TargetMo field value if set, zero value otherwise.
func (o *VnicSanConnectivityPolicyInventory) GetTargetMo() MoBaseMoRelationship {
	if o == nil || o.TargetMo == nil {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.TargetMo
}

// GetTargetMoOk returns a tuple with the TargetMo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicSanConnectivityPolicyInventory) GetTargetMoOk() (*MoBaseMoRelationship, bool) {
	if o == nil || o.TargetMo == nil {
		return nil, false
	}
	return o.TargetMo, true
}

// HasTargetMo returns a boolean if a field has been set.
func (o *VnicSanConnectivityPolicyInventory) HasTargetMo() bool {
	if o != nil && o.TargetMo != nil {
		return true
	}

	return false
}

// SetTargetMo gets a reference to the given MoBaseMoRelationship and assigns it to the TargetMo field.
func (o *VnicSanConnectivityPolicyInventory) SetTargetMo(v MoBaseMoRelationship) {
	o.TargetMo = &v
}

// GetWwnnPool returns the WwnnPool field value if set, zero value otherwise.
func (o *VnicSanConnectivityPolicyInventory) GetWwnnPool() FcpoolPoolRelationship {
	if o == nil || o.WwnnPool == nil {
		var ret FcpoolPoolRelationship
		return ret
	}
	return *o.WwnnPool
}

// GetWwnnPoolOk returns a tuple with the WwnnPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicSanConnectivityPolicyInventory) GetWwnnPoolOk() (*FcpoolPoolRelationship, bool) {
	if o == nil || o.WwnnPool == nil {
		return nil, false
	}
	return o.WwnnPool, true
}

// HasWwnnPool returns a boolean if a field has been set.
func (o *VnicSanConnectivityPolicyInventory) HasWwnnPool() bool {
	if o != nil && o.WwnnPool != nil {
		return true
	}

	return false
}

// SetWwnnPool gets a reference to the given FcpoolPoolRelationship and assigns it to the WwnnPool field.
func (o *VnicSanConnectivityPolicyInventory) SetWwnnPool(v FcpoolPoolRelationship) {
	o.WwnnPool = &v
}

func (o VnicSanConnectivityPolicyInventory) MarshalJSON() ([]byte, error) {
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
	if o.PlacementMode != nil {
		toSerialize["PlacementMode"] = o.PlacementMode
	}
	if o.StaticWwnnAddress != nil {
		toSerialize["StaticWwnnAddress"] = o.StaticWwnnAddress
	}
	if o.TargetPlatform != nil {
		toSerialize["TargetPlatform"] = o.TargetPlatform
	}
	if o.WwnnAddressType != nil {
		toSerialize["WwnnAddressType"] = o.WwnnAddressType
	}
	if o.FcIfs != nil {
		toSerialize["FcIfs"] = o.FcIfs
	}
	if o.TargetMo != nil {
		toSerialize["TargetMo"] = o.TargetMo
	}
	if o.WwnnPool != nil {
		toSerialize["WwnnPool"] = o.WwnnPool
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VnicSanConnectivityPolicyInventory) UnmarshalJSON(bytes []byte) (err error) {
	type VnicSanConnectivityPolicyInventoryWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The mode used for placement of vNICs on network adapters. It can either be Auto or Custom. * `custom` - The placement of the vNICs / vHBAs on network adapters is manually chosen by the user. * `auto` - The placement of the vNICs / vHBAs on network adapters is automatically determined by the system.
		PlacementMode *string `json:"PlacementMode,omitempty"`
		// The WWNN address for the server node must be in hexadecimal format xx:xx:xx:xx:xx:xx:xx:xx. Allowed ranges are 20:00:00:00:00:00:00:00 to 20:FF:FF:FF:FF:FF:FF:FF or from 50:00:00:00:00:00:00:00 to 5F:FF:FF:FF:FF:FF:FF:FF. To ensure uniqueness of WWN's in the SAN fabric, you are strongly encouraged to use the WWN prefix - 20:00:00:25:B5:xx:xx:xx.
		StaticWwnnAddress *string `json:"StaticWwnnAddress,omitempty"`
		// The platform for which the server profile is applicable. It can either be a server that is operating in standalone mode or which is attached to a Fabric Interconnect managed by Intersight. * `Standalone` - Servers which are operating in standalone mode i.e. not connected to a Fabric Interconnected. * `FIAttached` - Servers which are connected to a Fabric Interconnect that is managed by Intersight.
		TargetPlatform *string `json:"TargetPlatform,omitempty"`
		// Type of allocation selected to assign a WWNN address for the server node. * `POOL` - The user selects a pool from which the mac/wwn address will be leased for the Virtual Interface. * `STATIC` - The user assigns a static mac/wwn address for the Virtual Interface.
		WwnnAddressType *string `json:"WwnnAddressType,omitempty"`
		// An array of relationships to vnicFcIfInventory resources.
		FcIfs    []VnicFcIfInventoryRelationship `json:"FcIfs,omitempty"`
		TargetMo *MoBaseMoRelationship           `json:"TargetMo,omitempty"`
		WwnnPool *FcpoolPoolRelationship         `json:"WwnnPool,omitempty"`
	}

	varVnicSanConnectivityPolicyInventoryWithoutEmbeddedStruct := VnicSanConnectivityPolicyInventoryWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVnicSanConnectivityPolicyInventoryWithoutEmbeddedStruct)
	if err == nil {
		varVnicSanConnectivityPolicyInventory := _VnicSanConnectivityPolicyInventory{}
		varVnicSanConnectivityPolicyInventory.ClassId = varVnicSanConnectivityPolicyInventoryWithoutEmbeddedStruct.ClassId
		varVnicSanConnectivityPolicyInventory.ObjectType = varVnicSanConnectivityPolicyInventoryWithoutEmbeddedStruct.ObjectType
		varVnicSanConnectivityPolicyInventory.PlacementMode = varVnicSanConnectivityPolicyInventoryWithoutEmbeddedStruct.PlacementMode
		varVnicSanConnectivityPolicyInventory.StaticWwnnAddress = varVnicSanConnectivityPolicyInventoryWithoutEmbeddedStruct.StaticWwnnAddress
		varVnicSanConnectivityPolicyInventory.TargetPlatform = varVnicSanConnectivityPolicyInventoryWithoutEmbeddedStruct.TargetPlatform
		varVnicSanConnectivityPolicyInventory.WwnnAddressType = varVnicSanConnectivityPolicyInventoryWithoutEmbeddedStruct.WwnnAddressType
		varVnicSanConnectivityPolicyInventory.FcIfs = varVnicSanConnectivityPolicyInventoryWithoutEmbeddedStruct.FcIfs
		varVnicSanConnectivityPolicyInventory.TargetMo = varVnicSanConnectivityPolicyInventoryWithoutEmbeddedStruct.TargetMo
		varVnicSanConnectivityPolicyInventory.WwnnPool = varVnicSanConnectivityPolicyInventoryWithoutEmbeddedStruct.WwnnPool
		*o = VnicSanConnectivityPolicyInventory(varVnicSanConnectivityPolicyInventory)
	} else {
		return err
	}

	varVnicSanConnectivityPolicyInventory := _VnicSanConnectivityPolicyInventory{}

	err = json.Unmarshal(bytes, &varVnicSanConnectivityPolicyInventory)
	if err == nil {
		o.PolicyAbstractPolicyInventory = varVnicSanConnectivityPolicyInventory.PolicyAbstractPolicyInventory
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "PlacementMode")
		delete(additionalProperties, "StaticWwnnAddress")
		delete(additionalProperties, "TargetPlatform")
		delete(additionalProperties, "WwnnAddressType")
		delete(additionalProperties, "FcIfs")
		delete(additionalProperties, "TargetMo")
		delete(additionalProperties, "WwnnPool")

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

type NullableVnicSanConnectivityPolicyInventory struct {
	value *VnicSanConnectivityPolicyInventory
	isSet bool
}

func (v NullableVnicSanConnectivityPolicyInventory) Get() *VnicSanConnectivityPolicyInventory {
	return v.value
}

func (v *NullableVnicSanConnectivityPolicyInventory) Set(val *VnicSanConnectivityPolicyInventory) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicSanConnectivityPolicyInventory) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicSanConnectivityPolicyInventory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicSanConnectivityPolicyInventory(val *VnicSanConnectivityPolicyInventory) *NullableVnicSanConnectivityPolicyInventory {
	return &NullableVnicSanConnectivityPolicyInventory{value: val, isSet: true}
}

func (v NullableVnicSanConnectivityPolicyInventory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicSanConnectivityPolicyInventory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
