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
	"reflect"
	"strings"
)

// EquipmentSwitchOperation Models the operational states of a Switch in Intersight.
type EquipmentSwitchOperation struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Action performed on the locator LED of the switch. * `None` - No operation action for the Locator Led of an equipment. * `TurnOn` - Turn on the Locator Led of an equipment. * `TurnOff` - Turn off the Locator Led of an equipment.
	AdminLocatorLedAction *string `json:"AdminLocatorLedAction,omitempty"`
	// Defines status of action performed on AdminLocatorLedState. * `None` - Nil value when no action has been triggered by the user. * `Applied` - User configured settings are in applied state. * `Applying` - User settings are being applied on the target server. * `Failed` - User configured settings could not be applied.
	AdminLocatorLedActionState *string                              `json:"AdminLocatorLedActionState,omitempty"`
	DeviceRegistration         *AssetDeviceRegistrationRelationship `json:"DeviceRegistration,omitempty"`
	NetworkElement             *NetworkElementRelationship          `json:"NetworkElement,omitempty"`
	AdditionalProperties       map[string]interface{}
}

type _EquipmentSwitchOperation EquipmentSwitchOperation

// NewEquipmentSwitchOperation instantiates a new EquipmentSwitchOperation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquipmentSwitchOperation(classId string, objectType string) *EquipmentSwitchOperation {
	this := EquipmentSwitchOperation{}
	this.ClassId = classId
	this.ObjectType = objectType
	var adminLocatorLedAction string = "None"
	this.AdminLocatorLedAction = &adminLocatorLedAction
	var adminLocatorLedActionState string = "None"
	this.AdminLocatorLedActionState = &adminLocatorLedActionState
	return &this
}

// NewEquipmentSwitchOperationWithDefaults instantiates a new EquipmentSwitchOperation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquipmentSwitchOperationWithDefaults() *EquipmentSwitchOperation {
	this := EquipmentSwitchOperation{}
	var classId string = "equipment.SwitchOperation"
	this.ClassId = classId
	var objectType string = "equipment.SwitchOperation"
	this.ObjectType = objectType
	var adminLocatorLedAction string = "None"
	this.AdminLocatorLedAction = &adminLocatorLedAction
	var adminLocatorLedActionState string = "None"
	this.AdminLocatorLedActionState = &adminLocatorLedActionState
	return &this
}

// GetClassId returns the ClassId field value
func (o *EquipmentSwitchOperation) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *EquipmentSwitchOperation) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *EquipmentSwitchOperation) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *EquipmentSwitchOperation) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *EquipmentSwitchOperation) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *EquipmentSwitchOperation) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAdminLocatorLedAction returns the AdminLocatorLedAction field value if set, zero value otherwise.
func (o *EquipmentSwitchOperation) GetAdminLocatorLedAction() string {
	if o == nil || o.AdminLocatorLedAction == nil {
		var ret string
		return ret
	}
	return *o.AdminLocatorLedAction
}

// GetAdminLocatorLedActionOk returns a tuple with the AdminLocatorLedAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSwitchOperation) GetAdminLocatorLedActionOk() (*string, bool) {
	if o == nil || o.AdminLocatorLedAction == nil {
		return nil, false
	}
	return o.AdminLocatorLedAction, true
}

// HasAdminLocatorLedAction returns a boolean if a field has been set.
func (o *EquipmentSwitchOperation) HasAdminLocatorLedAction() bool {
	if o != nil && o.AdminLocatorLedAction != nil {
		return true
	}

	return false
}

// SetAdminLocatorLedAction gets a reference to the given string and assigns it to the AdminLocatorLedAction field.
func (o *EquipmentSwitchOperation) SetAdminLocatorLedAction(v string) {
	o.AdminLocatorLedAction = &v
}

// GetAdminLocatorLedActionState returns the AdminLocatorLedActionState field value if set, zero value otherwise.
func (o *EquipmentSwitchOperation) GetAdminLocatorLedActionState() string {
	if o == nil || o.AdminLocatorLedActionState == nil {
		var ret string
		return ret
	}
	return *o.AdminLocatorLedActionState
}

// GetAdminLocatorLedActionStateOk returns a tuple with the AdminLocatorLedActionState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSwitchOperation) GetAdminLocatorLedActionStateOk() (*string, bool) {
	if o == nil || o.AdminLocatorLedActionState == nil {
		return nil, false
	}
	return o.AdminLocatorLedActionState, true
}

// HasAdminLocatorLedActionState returns a boolean if a field has been set.
func (o *EquipmentSwitchOperation) HasAdminLocatorLedActionState() bool {
	if o != nil && o.AdminLocatorLedActionState != nil {
		return true
	}

	return false
}

// SetAdminLocatorLedActionState gets a reference to the given string and assigns it to the AdminLocatorLedActionState field.
func (o *EquipmentSwitchOperation) SetAdminLocatorLedActionState(v string) {
	o.AdminLocatorLedActionState = &v
}

// GetDeviceRegistration returns the DeviceRegistration field value if set, zero value otherwise.
func (o *EquipmentSwitchOperation) GetDeviceRegistration() AssetDeviceRegistrationRelationship {
	if o == nil || o.DeviceRegistration == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.DeviceRegistration
}

// GetDeviceRegistrationOk returns a tuple with the DeviceRegistration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSwitchOperation) GetDeviceRegistrationOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.DeviceRegistration == nil {
		return nil, false
	}
	return o.DeviceRegistration, true
}

// HasDeviceRegistration returns a boolean if a field has been set.
func (o *EquipmentSwitchOperation) HasDeviceRegistration() bool {
	if o != nil && o.DeviceRegistration != nil {
		return true
	}

	return false
}

// SetDeviceRegistration gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the DeviceRegistration field.
func (o *EquipmentSwitchOperation) SetDeviceRegistration(v AssetDeviceRegistrationRelationship) {
	o.DeviceRegistration = &v
}

// GetNetworkElement returns the NetworkElement field value if set, zero value otherwise.
func (o *EquipmentSwitchOperation) GetNetworkElement() NetworkElementRelationship {
	if o == nil || o.NetworkElement == nil {
		var ret NetworkElementRelationship
		return ret
	}
	return *o.NetworkElement
}

// GetNetworkElementOk returns a tuple with the NetworkElement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSwitchOperation) GetNetworkElementOk() (*NetworkElementRelationship, bool) {
	if o == nil || o.NetworkElement == nil {
		return nil, false
	}
	return o.NetworkElement, true
}

// HasNetworkElement returns a boolean if a field has been set.
func (o *EquipmentSwitchOperation) HasNetworkElement() bool {
	if o != nil && o.NetworkElement != nil {
		return true
	}

	return false
}

// SetNetworkElement gets a reference to the given NetworkElementRelationship and assigns it to the NetworkElement field.
func (o *EquipmentSwitchOperation) SetNetworkElement(v NetworkElementRelationship) {
	o.NetworkElement = &v
}

func (o EquipmentSwitchOperation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AdminLocatorLedAction != nil {
		toSerialize["AdminLocatorLedAction"] = o.AdminLocatorLedAction
	}
	if o.AdminLocatorLedActionState != nil {
		toSerialize["AdminLocatorLedActionState"] = o.AdminLocatorLedActionState
	}
	if o.DeviceRegistration != nil {
		toSerialize["DeviceRegistration"] = o.DeviceRegistration
	}
	if o.NetworkElement != nil {
		toSerialize["NetworkElement"] = o.NetworkElement
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EquipmentSwitchOperation) UnmarshalJSON(bytes []byte) (err error) {
	type EquipmentSwitchOperationWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Action performed on the locator LED of the switch. * `None` - No operation action for the Locator Led of an equipment. * `TurnOn` - Turn on the Locator Led of an equipment. * `TurnOff` - Turn off the Locator Led of an equipment.
		AdminLocatorLedAction *string `json:"AdminLocatorLedAction,omitempty"`
		// Defines status of action performed on AdminLocatorLedState. * `None` - Nil value when no action has been triggered by the user. * `Applied` - User configured settings are in applied state. * `Applying` - User settings are being applied on the target server. * `Failed` - User configured settings could not be applied.
		AdminLocatorLedActionState *string                              `json:"AdminLocatorLedActionState,omitempty"`
		DeviceRegistration         *AssetDeviceRegistrationRelationship `json:"DeviceRegistration,omitempty"`
		NetworkElement             *NetworkElementRelationship          `json:"NetworkElement,omitempty"`
	}

	varEquipmentSwitchOperationWithoutEmbeddedStruct := EquipmentSwitchOperationWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varEquipmentSwitchOperationWithoutEmbeddedStruct)
	if err == nil {
		varEquipmentSwitchOperation := _EquipmentSwitchOperation{}
		varEquipmentSwitchOperation.ClassId = varEquipmentSwitchOperationWithoutEmbeddedStruct.ClassId
		varEquipmentSwitchOperation.ObjectType = varEquipmentSwitchOperationWithoutEmbeddedStruct.ObjectType
		varEquipmentSwitchOperation.AdminLocatorLedAction = varEquipmentSwitchOperationWithoutEmbeddedStruct.AdminLocatorLedAction
		varEquipmentSwitchOperation.AdminLocatorLedActionState = varEquipmentSwitchOperationWithoutEmbeddedStruct.AdminLocatorLedActionState
		varEquipmentSwitchOperation.DeviceRegistration = varEquipmentSwitchOperationWithoutEmbeddedStruct.DeviceRegistration
		varEquipmentSwitchOperation.NetworkElement = varEquipmentSwitchOperationWithoutEmbeddedStruct.NetworkElement
		*o = EquipmentSwitchOperation(varEquipmentSwitchOperation)
	} else {
		return err
	}

	varEquipmentSwitchOperation := _EquipmentSwitchOperation{}

	err = json.Unmarshal(bytes, &varEquipmentSwitchOperation)
	if err == nil {
		o.MoBaseMo = varEquipmentSwitchOperation.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AdminLocatorLedAction")
		delete(additionalProperties, "AdminLocatorLedActionState")
		delete(additionalProperties, "DeviceRegistration")
		delete(additionalProperties, "NetworkElement")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

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

type NullableEquipmentSwitchOperation struct {
	value *EquipmentSwitchOperation
	isSet bool
}

func (v NullableEquipmentSwitchOperation) Get() *EquipmentSwitchOperation {
	return v.value
}

func (v *NullableEquipmentSwitchOperation) Set(val *EquipmentSwitchOperation) {
	v.value = val
	v.isSet = true
}

func (v NullableEquipmentSwitchOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableEquipmentSwitchOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquipmentSwitchOperation(val *EquipmentSwitchOperation) *NullableEquipmentSwitchOperation {
	return &NullableEquipmentSwitchOperation{value: val, isSet: true}
}

func (v NullableEquipmentSwitchOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquipmentSwitchOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
