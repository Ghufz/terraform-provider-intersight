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

// FabricPortOperation PortOperation objects allows the user to alter the state of the port.
type FabricPortOperation struct {
	FabricPortBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Admin configured state to disable the port. * `Enabled` - Admin configured Enabled State. * `Disabled` - Admin configured Disabled State.
	AdminState *string `json:"AdminState,omitempty"`
	// The configured state of these settings in the target chassis. The value is any one of Applied, Applying, Failed. Applied - This state denotes that the admin state changes are applied successfully in the target FI domain. Applying - This state denotes that the admin state changes are being applied in the target FI domain. Failed - This state denotes that the admin state changes could not be applied in the target FI domain. * `None` - Nil value when no action has been triggered by the user. * `Applied` - User configured settings are in applied state. * `Applying` - User settings are being applied on the target server. * `Failed` - User configured settings could not be applied.
	ConfigState          *string                     `json:"ConfigState,omitempty"`
	NetworkElement       *NetworkElementRelationship `json:"NetworkElement,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricPortOperation FabricPortOperation

// NewFabricPortOperation instantiates a new FabricPortOperation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricPortOperation(classId string, objectType string) *FabricPortOperation {
	this := FabricPortOperation{}
	this.ClassId = classId
	this.ObjectType = objectType
	var adminState string = "Enabled"
	this.AdminState = &adminState
	var configState string = "None"
	this.ConfigState = &configState
	return &this
}

// NewFabricPortOperationWithDefaults instantiates a new FabricPortOperation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricPortOperationWithDefaults() *FabricPortOperation {
	this := FabricPortOperation{}
	var classId string = "fabric.PortOperation"
	this.ClassId = classId
	var objectType string = "fabric.PortOperation"
	this.ObjectType = objectType
	var adminState string = "Enabled"
	this.AdminState = &adminState
	var configState string = "None"
	this.ConfigState = &configState
	return &this
}

// GetClassId returns the ClassId field value
func (o *FabricPortOperation) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FabricPortOperation) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FabricPortOperation) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FabricPortOperation) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FabricPortOperation) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FabricPortOperation) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAdminState returns the AdminState field value if set, zero value otherwise.
func (o *FabricPortOperation) GetAdminState() string {
	if o == nil || o.AdminState == nil {
		var ret string
		return ret
	}
	return *o.AdminState
}

// GetAdminStateOk returns a tuple with the AdminState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricPortOperation) GetAdminStateOk() (*string, bool) {
	if o == nil || o.AdminState == nil {
		return nil, false
	}
	return o.AdminState, true
}

// HasAdminState returns a boolean if a field has been set.
func (o *FabricPortOperation) HasAdminState() bool {
	if o != nil && o.AdminState != nil {
		return true
	}

	return false
}

// SetAdminState gets a reference to the given string and assigns it to the AdminState field.
func (o *FabricPortOperation) SetAdminState(v string) {
	o.AdminState = &v
}

// GetConfigState returns the ConfigState field value if set, zero value otherwise.
func (o *FabricPortOperation) GetConfigState() string {
	if o == nil || o.ConfigState == nil {
		var ret string
		return ret
	}
	return *o.ConfigState
}

// GetConfigStateOk returns a tuple with the ConfigState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricPortOperation) GetConfigStateOk() (*string, bool) {
	if o == nil || o.ConfigState == nil {
		return nil, false
	}
	return o.ConfigState, true
}

// HasConfigState returns a boolean if a field has been set.
func (o *FabricPortOperation) HasConfigState() bool {
	if o != nil && o.ConfigState != nil {
		return true
	}

	return false
}

// SetConfigState gets a reference to the given string and assigns it to the ConfigState field.
func (o *FabricPortOperation) SetConfigState(v string) {
	o.ConfigState = &v
}

// GetNetworkElement returns the NetworkElement field value if set, zero value otherwise.
func (o *FabricPortOperation) GetNetworkElement() NetworkElementRelationship {
	if o == nil || o.NetworkElement == nil {
		var ret NetworkElementRelationship
		return ret
	}
	return *o.NetworkElement
}

// GetNetworkElementOk returns a tuple with the NetworkElement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricPortOperation) GetNetworkElementOk() (*NetworkElementRelationship, bool) {
	if o == nil || o.NetworkElement == nil {
		return nil, false
	}
	return o.NetworkElement, true
}

// HasNetworkElement returns a boolean if a field has been set.
func (o *FabricPortOperation) HasNetworkElement() bool {
	if o != nil && o.NetworkElement != nil {
		return true
	}

	return false
}

// SetNetworkElement gets a reference to the given NetworkElementRelationship and assigns it to the NetworkElement field.
func (o *FabricPortOperation) SetNetworkElement(v NetworkElementRelationship) {
	o.NetworkElement = &v
}

func (o FabricPortOperation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedFabricPortBase, errFabricPortBase := json.Marshal(o.FabricPortBase)
	if errFabricPortBase != nil {
		return []byte{}, errFabricPortBase
	}
	errFabricPortBase = json.Unmarshal([]byte(serializedFabricPortBase), &toSerialize)
	if errFabricPortBase != nil {
		return []byte{}, errFabricPortBase
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AdminState != nil {
		toSerialize["AdminState"] = o.AdminState
	}
	if o.ConfigState != nil {
		toSerialize["ConfigState"] = o.ConfigState
	}
	if o.NetworkElement != nil {
		toSerialize["NetworkElement"] = o.NetworkElement
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FabricPortOperation) UnmarshalJSON(bytes []byte) (err error) {
	type FabricPortOperationWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Admin configured state to disable the port. * `Enabled` - Admin configured Enabled State. * `Disabled` - Admin configured Disabled State.
		AdminState *string `json:"AdminState,omitempty"`
		// The configured state of these settings in the target chassis. The value is any one of Applied, Applying, Failed. Applied - This state denotes that the admin state changes are applied successfully in the target FI domain. Applying - This state denotes that the admin state changes are being applied in the target FI domain. Failed - This state denotes that the admin state changes could not be applied in the target FI domain. * `None` - Nil value when no action has been triggered by the user. * `Applied` - User configured settings are in applied state. * `Applying` - User settings are being applied on the target server. * `Failed` - User configured settings could not be applied.
		ConfigState    *string                     `json:"ConfigState,omitempty"`
		NetworkElement *NetworkElementRelationship `json:"NetworkElement,omitempty"`
	}

	varFabricPortOperationWithoutEmbeddedStruct := FabricPortOperationWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varFabricPortOperationWithoutEmbeddedStruct)
	if err == nil {
		varFabricPortOperation := _FabricPortOperation{}
		varFabricPortOperation.ClassId = varFabricPortOperationWithoutEmbeddedStruct.ClassId
		varFabricPortOperation.ObjectType = varFabricPortOperationWithoutEmbeddedStruct.ObjectType
		varFabricPortOperation.AdminState = varFabricPortOperationWithoutEmbeddedStruct.AdminState
		varFabricPortOperation.ConfigState = varFabricPortOperationWithoutEmbeddedStruct.ConfigState
		varFabricPortOperation.NetworkElement = varFabricPortOperationWithoutEmbeddedStruct.NetworkElement
		*o = FabricPortOperation(varFabricPortOperation)
	} else {
		return err
	}

	varFabricPortOperation := _FabricPortOperation{}

	err = json.Unmarshal(bytes, &varFabricPortOperation)
	if err == nil {
		o.FabricPortBase = varFabricPortOperation.FabricPortBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AdminState")
		delete(additionalProperties, "ConfigState")
		delete(additionalProperties, "NetworkElement")

		// remove fields from embedded structs
		reflectFabricPortBase := reflect.ValueOf(o.FabricPortBase)
		for i := 0; i < reflectFabricPortBase.Type().NumField(); i++ {
			t := reflectFabricPortBase.Type().Field(i)

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

type NullableFabricPortOperation struct {
	value *FabricPortOperation
	isSet bool
}

func (v NullableFabricPortOperation) Get() *FabricPortOperation {
	return v.value
}

func (v *NullableFabricPortOperation) Set(val *FabricPortOperation) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricPortOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricPortOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricPortOperation(val *FabricPortOperation) *NullableFabricPortOperation {
	return &NullableFabricPortOperation{value: val, isSet: true}
}

func (v NullableFabricPortOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricPortOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
