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

// FabricSanPinGroupAllOf Definition of the list of properties defined in 'fabric.SanPinGroup', excluding properties defined in parent classes.
type FabricSanPinGroupAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType             string                                   `json:"ObjectType"`
	PinTargetInterfaceRole *FabricAbstractInterfaceRoleRelationship `json:"PinTargetInterfaceRole,omitempty"`
	AdditionalProperties   map[string]interface{}
}

type _FabricSanPinGroupAllOf FabricSanPinGroupAllOf

// NewFabricSanPinGroupAllOf instantiates a new FabricSanPinGroupAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricSanPinGroupAllOf(classId string, objectType string) *FabricSanPinGroupAllOf {
	this := FabricSanPinGroupAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewFabricSanPinGroupAllOfWithDefaults instantiates a new FabricSanPinGroupAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricSanPinGroupAllOfWithDefaults() *FabricSanPinGroupAllOf {
	this := FabricSanPinGroupAllOf{}
	var classId string = "fabric.SanPinGroup"
	this.ClassId = classId
	var objectType string = "fabric.SanPinGroup"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *FabricSanPinGroupAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FabricSanPinGroupAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FabricSanPinGroupAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FabricSanPinGroupAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FabricSanPinGroupAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FabricSanPinGroupAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetPinTargetInterfaceRole returns the PinTargetInterfaceRole field value if set, zero value otherwise.
func (o *FabricSanPinGroupAllOf) GetPinTargetInterfaceRole() FabricAbstractInterfaceRoleRelationship {
	if o == nil || o.PinTargetInterfaceRole == nil {
		var ret FabricAbstractInterfaceRoleRelationship
		return ret
	}
	return *o.PinTargetInterfaceRole
}

// GetPinTargetInterfaceRoleOk returns a tuple with the PinTargetInterfaceRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricSanPinGroupAllOf) GetPinTargetInterfaceRoleOk() (*FabricAbstractInterfaceRoleRelationship, bool) {
	if o == nil || o.PinTargetInterfaceRole == nil {
		return nil, false
	}
	return o.PinTargetInterfaceRole, true
}

// HasPinTargetInterfaceRole returns a boolean if a field has been set.
func (o *FabricSanPinGroupAllOf) HasPinTargetInterfaceRole() bool {
	if o != nil && o.PinTargetInterfaceRole != nil {
		return true
	}

	return false
}

// SetPinTargetInterfaceRole gets a reference to the given FabricAbstractInterfaceRoleRelationship and assigns it to the PinTargetInterfaceRole field.
func (o *FabricSanPinGroupAllOf) SetPinTargetInterfaceRole(v FabricAbstractInterfaceRoleRelationship) {
	o.PinTargetInterfaceRole = &v
}

func (o FabricSanPinGroupAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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

func (o *FabricSanPinGroupAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varFabricSanPinGroupAllOf := _FabricSanPinGroupAllOf{}

	if err = json.Unmarshal(bytes, &varFabricSanPinGroupAllOf); err == nil {
		*o = FabricSanPinGroupAllOf(varFabricSanPinGroupAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "PinTargetInterfaceRole")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFabricSanPinGroupAllOf struct {
	value *FabricSanPinGroupAllOf
	isSet bool
}

func (v NullableFabricSanPinGroupAllOf) Get() *FabricSanPinGroupAllOf {
	return v.value
}

func (v *NullableFabricSanPinGroupAllOf) Set(val *FabricSanPinGroupAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricSanPinGroupAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricSanPinGroupAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricSanPinGroupAllOf(val *FabricSanPinGroupAllOf) *NullableFabricSanPinGroupAllOf {
	return &NullableFabricSanPinGroupAllOf{value: val, isSet: true}
}

func (v NullableFabricSanPinGroupAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricSanPinGroupAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
