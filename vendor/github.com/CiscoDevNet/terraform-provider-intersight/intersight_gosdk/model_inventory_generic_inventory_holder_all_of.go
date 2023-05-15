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

// InventoryGenericInventoryHolderAllOf Definition of the list of properties defined in 'inventory.GenericInventoryHolder', excluding properties defined in parent classes.
type InventoryGenericInventoryHolderAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The endpoint represented by this holder.
	Endpoint        *string                      `json:"Endpoint,omitempty"`
	ComputeBlade    *ComputeBladeRelationship    `json:"ComputeBlade,omitempty"`
	ComputeRackUnit *ComputeRackUnitRelationship `json:"ComputeRackUnit,omitempty"`
	// An array of relationships to inventoryGenericInventory resources.
	GenericInventory     []InventoryGenericInventoryRelationship `json:"GenericInventory,omitempty"`
	InventoryDeviceInfo  *InventoryDeviceInfoRelationship        `json:"InventoryDeviceInfo,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship    `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InventoryGenericInventoryHolderAllOf InventoryGenericInventoryHolderAllOf

// NewInventoryGenericInventoryHolderAllOf instantiates a new InventoryGenericInventoryHolderAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryGenericInventoryHolderAllOf(classId string, objectType string) *InventoryGenericInventoryHolderAllOf {
	this := InventoryGenericInventoryHolderAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewInventoryGenericInventoryHolderAllOfWithDefaults instantiates a new InventoryGenericInventoryHolderAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryGenericInventoryHolderAllOfWithDefaults() *InventoryGenericInventoryHolderAllOf {
	this := InventoryGenericInventoryHolderAllOf{}
	var classId string = "inventory.GenericInventoryHolder"
	this.ClassId = classId
	var objectType string = "inventory.GenericInventoryHolder"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *InventoryGenericInventoryHolderAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *InventoryGenericInventoryHolderAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *InventoryGenericInventoryHolderAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *InventoryGenericInventoryHolderAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *InventoryGenericInventoryHolderAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *InventoryGenericInventoryHolderAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise.
func (o *InventoryGenericInventoryHolderAllOf) GetEndpoint() string {
	if o == nil || o.Endpoint == nil {
		var ret string
		return ret
	}
	return *o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryGenericInventoryHolderAllOf) GetEndpointOk() (*string, bool) {
	if o == nil || o.Endpoint == nil {
		return nil, false
	}
	return o.Endpoint, true
}

// HasEndpoint returns a boolean if a field has been set.
func (o *InventoryGenericInventoryHolderAllOf) HasEndpoint() bool {
	if o != nil && o.Endpoint != nil {
		return true
	}

	return false
}

// SetEndpoint gets a reference to the given string and assigns it to the Endpoint field.
func (o *InventoryGenericInventoryHolderAllOf) SetEndpoint(v string) {
	o.Endpoint = &v
}

// GetComputeBlade returns the ComputeBlade field value if set, zero value otherwise.
func (o *InventoryGenericInventoryHolderAllOf) GetComputeBlade() ComputeBladeRelationship {
	if o == nil || o.ComputeBlade == nil {
		var ret ComputeBladeRelationship
		return ret
	}
	return *o.ComputeBlade
}

// GetComputeBladeOk returns a tuple with the ComputeBlade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryGenericInventoryHolderAllOf) GetComputeBladeOk() (*ComputeBladeRelationship, bool) {
	if o == nil || o.ComputeBlade == nil {
		return nil, false
	}
	return o.ComputeBlade, true
}

// HasComputeBlade returns a boolean if a field has been set.
func (o *InventoryGenericInventoryHolderAllOf) HasComputeBlade() bool {
	if o != nil && o.ComputeBlade != nil {
		return true
	}

	return false
}

// SetComputeBlade gets a reference to the given ComputeBladeRelationship and assigns it to the ComputeBlade field.
func (o *InventoryGenericInventoryHolderAllOf) SetComputeBlade(v ComputeBladeRelationship) {
	o.ComputeBlade = &v
}

// GetComputeRackUnit returns the ComputeRackUnit field value if set, zero value otherwise.
func (o *InventoryGenericInventoryHolderAllOf) GetComputeRackUnit() ComputeRackUnitRelationship {
	if o == nil || o.ComputeRackUnit == nil {
		var ret ComputeRackUnitRelationship
		return ret
	}
	return *o.ComputeRackUnit
}

// GetComputeRackUnitOk returns a tuple with the ComputeRackUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryGenericInventoryHolderAllOf) GetComputeRackUnitOk() (*ComputeRackUnitRelationship, bool) {
	if o == nil || o.ComputeRackUnit == nil {
		return nil, false
	}
	return o.ComputeRackUnit, true
}

// HasComputeRackUnit returns a boolean if a field has been set.
func (o *InventoryGenericInventoryHolderAllOf) HasComputeRackUnit() bool {
	if o != nil && o.ComputeRackUnit != nil {
		return true
	}

	return false
}

// SetComputeRackUnit gets a reference to the given ComputeRackUnitRelationship and assigns it to the ComputeRackUnit field.
func (o *InventoryGenericInventoryHolderAllOf) SetComputeRackUnit(v ComputeRackUnitRelationship) {
	o.ComputeRackUnit = &v
}

// GetGenericInventory returns the GenericInventory field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InventoryGenericInventoryHolderAllOf) GetGenericInventory() []InventoryGenericInventoryRelationship {
	if o == nil {
		var ret []InventoryGenericInventoryRelationship
		return ret
	}
	return o.GenericInventory
}

// GetGenericInventoryOk returns a tuple with the GenericInventory field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InventoryGenericInventoryHolderAllOf) GetGenericInventoryOk() ([]InventoryGenericInventoryRelationship, bool) {
	if o == nil || o.GenericInventory == nil {
		return nil, false
	}
	return o.GenericInventory, true
}

// HasGenericInventory returns a boolean if a field has been set.
func (o *InventoryGenericInventoryHolderAllOf) HasGenericInventory() bool {
	if o != nil && o.GenericInventory != nil {
		return true
	}

	return false
}

// SetGenericInventory gets a reference to the given []InventoryGenericInventoryRelationship and assigns it to the GenericInventory field.
func (o *InventoryGenericInventoryHolderAllOf) SetGenericInventory(v []InventoryGenericInventoryRelationship) {
	o.GenericInventory = v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *InventoryGenericInventoryHolderAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryGenericInventoryHolderAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *InventoryGenericInventoryHolderAllOf) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *InventoryGenericInventoryHolderAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *InventoryGenericInventoryHolderAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryGenericInventoryHolderAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *InventoryGenericInventoryHolderAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *InventoryGenericInventoryHolderAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o InventoryGenericInventoryHolderAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Endpoint != nil {
		toSerialize["Endpoint"] = o.Endpoint
	}
	if o.ComputeBlade != nil {
		toSerialize["ComputeBlade"] = o.ComputeBlade
	}
	if o.ComputeRackUnit != nil {
		toSerialize["ComputeRackUnit"] = o.ComputeRackUnit
	}
	if o.GenericInventory != nil {
		toSerialize["GenericInventory"] = o.GenericInventory
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

func (o *InventoryGenericInventoryHolderAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varInventoryGenericInventoryHolderAllOf := _InventoryGenericInventoryHolderAllOf{}

	if err = json.Unmarshal(bytes, &varInventoryGenericInventoryHolderAllOf); err == nil {
		*o = InventoryGenericInventoryHolderAllOf(varInventoryGenericInventoryHolderAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Endpoint")
		delete(additionalProperties, "ComputeBlade")
		delete(additionalProperties, "ComputeRackUnit")
		delete(additionalProperties, "GenericInventory")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInventoryGenericInventoryHolderAllOf struct {
	value *InventoryGenericInventoryHolderAllOf
	isSet bool
}

func (v NullableInventoryGenericInventoryHolderAllOf) Get() *InventoryGenericInventoryHolderAllOf {
	return v.value
}

func (v *NullableInventoryGenericInventoryHolderAllOf) Set(val *InventoryGenericInventoryHolderAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryGenericInventoryHolderAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryGenericInventoryHolderAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryGenericInventoryHolderAllOf(val *InventoryGenericInventoryHolderAllOf) *NullableInventoryGenericInventoryHolderAllOf {
	return &NullableInventoryGenericInventoryHolderAllOf{value: val, isSet: true}
}

func (v NullableInventoryGenericInventoryHolderAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryGenericInventoryHolderAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
