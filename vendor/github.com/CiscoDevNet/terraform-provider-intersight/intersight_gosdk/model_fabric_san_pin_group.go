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

// FabricSanPinGroup SAN PinGroup configuration sent by user for static pinning.
type FabricSanPinGroup struct {
	FabricPinGroup
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType             string                                   `json:"ObjectType"`
	PinTargetInterfaceRole *FabricAbstractInterfaceRoleRelationship `json:"PinTargetInterfaceRole,omitempty"`
	AdditionalProperties   map[string]interface{}
}

type _FabricSanPinGroup FabricSanPinGroup

// NewFabricSanPinGroup instantiates a new FabricSanPinGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricSanPinGroup(classId string, objectType string) *FabricSanPinGroup {
	this := FabricSanPinGroup{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewFabricSanPinGroupWithDefaults instantiates a new FabricSanPinGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricSanPinGroupWithDefaults() *FabricSanPinGroup {
	this := FabricSanPinGroup{}
	var classId string = "fabric.SanPinGroup"
	this.ClassId = classId
	var objectType string = "fabric.SanPinGroup"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *FabricSanPinGroup) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FabricSanPinGroup) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FabricSanPinGroup) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FabricSanPinGroup) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FabricSanPinGroup) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FabricSanPinGroup) SetObjectType(v string) {
	o.ObjectType = v
}

// GetPinTargetInterfaceRole returns the PinTargetInterfaceRole field value if set, zero value otherwise.
func (o *FabricSanPinGroup) GetPinTargetInterfaceRole() FabricAbstractInterfaceRoleRelationship {
	if o == nil || o.PinTargetInterfaceRole == nil {
		var ret FabricAbstractInterfaceRoleRelationship
		return ret
	}
	return *o.PinTargetInterfaceRole
}

// GetPinTargetInterfaceRoleOk returns a tuple with the PinTargetInterfaceRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricSanPinGroup) GetPinTargetInterfaceRoleOk() (*FabricAbstractInterfaceRoleRelationship, bool) {
	if o == nil || o.PinTargetInterfaceRole == nil {
		return nil, false
	}
	return o.PinTargetInterfaceRole, true
}

// HasPinTargetInterfaceRole returns a boolean if a field has been set.
func (o *FabricSanPinGroup) HasPinTargetInterfaceRole() bool {
	if o != nil && o.PinTargetInterfaceRole != nil {
		return true
	}

	return false
}

// SetPinTargetInterfaceRole gets a reference to the given FabricAbstractInterfaceRoleRelationship and assigns it to the PinTargetInterfaceRole field.
func (o *FabricSanPinGroup) SetPinTargetInterfaceRole(v FabricAbstractInterfaceRoleRelationship) {
	o.PinTargetInterfaceRole = &v
}

func (o FabricSanPinGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedFabricPinGroup, errFabricPinGroup := json.Marshal(o.FabricPinGroup)
	if errFabricPinGroup != nil {
		return []byte{}, errFabricPinGroup
	}
	errFabricPinGroup = json.Unmarshal([]byte(serializedFabricPinGroup), &toSerialize)
	if errFabricPinGroup != nil {
		return []byte{}, errFabricPinGroup
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.PinTargetInterfaceRole != nil {
		toSerialize["PinTargetInterfaceRole"] = o.PinTargetInterfaceRole
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FabricSanPinGroup) UnmarshalJSON(bytes []byte) (err error) {
	type FabricSanPinGroupWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType             string                                   `json:"ObjectType"`
		PinTargetInterfaceRole *FabricAbstractInterfaceRoleRelationship `json:"PinTargetInterfaceRole,omitempty"`
	}

	varFabricSanPinGroupWithoutEmbeddedStruct := FabricSanPinGroupWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varFabricSanPinGroupWithoutEmbeddedStruct)
	if err == nil {
		varFabricSanPinGroup := _FabricSanPinGroup{}
		varFabricSanPinGroup.ClassId = varFabricSanPinGroupWithoutEmbeddedStruct.ClassId
		varFabricSanPinGroup.ObjectType = varFabricSanPinGroupWithoutEmbeddedStruct.ObjectType
		varFabricSanPinGroup.PinTargetInterfaceRole = varFabricSanPinGroupWithoutEmbeddedStruct.PinTargetInterfaceRole
		*o = FabricSanPinGroup(varFabricSanPinGroup)
	} else {
		return err
	}

	varFabricSanPinGroup := _FabricSanPinGroup{}

	err = json.Unmarshal(bytes, &varFabricSanPinGroup)
	if err == nil {
		o.FabricPinGroup = varFabricSanPinGroup.FabricPinGroup
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "PinTargetInterfaceRole")

		// remove fields from embedded structs
		reflectFabricPinGroup := reflect.ValueOf(o.FabricPinGroup)
		for i := 0; i < reflectFabricPinGroup.Type().NumField(); i++ {
			t := reflectFabricPinGroup.Type().Field(i)

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

type NullableFabricSanPinGroup struct {
	value *FabricSanPinGroup
	isSet bool
}

func (v NullableFabricSanPinGroup) Get() *FabricSanPinGroup {
	return v.value
}

func (v *NullableFabricSanPinGroup) Set(val *FabricSanPinGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricSanPinGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricSanPinGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricSanPinGroup(val *FabricSanPinGroup) *NullableFabricSanPinGroup {
	return &NullableFabricSanPinGroup{value: val, isSet: true}
}

func (v NullableFabricSanPinGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricSanPinGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
