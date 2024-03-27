/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-15711
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// VnicEthNetworkPolicyAllOf Definition of the list of properties defined in 'vnic.EthNetworkPolicy', excluding properties defined in parent classes.
type VnicEthNetworkPolicyAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The platform for which the server profile is applicable. It can either be a server that is operating in standalone mode or which is attached to a Fabric Interconnect managed by Intersight. * `Standalone` - Servers which are operating in standalone mode i.e. not connected to a Fabric Interconnected. * `FIAttached` - Servers which are connected to a Fabric Interconnect that is managed by Intersight.
	// Deprecated
	TargetPlatform       *string                               `json:"TargetPlatform,omitempty"`
	VlanSettings         NullableVnicVlanSettings              `json:"VlanSettings,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VnicEthNetworkPolicyAllOf VnicEthNetworkPolicyAllOf

// NewVnicEthNetworkPolicyAllOf instantiates a new VnicEthNetworkPolicyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicEthNetworkPolicyAllOf(classId string, objectType string) *VnicEthNetworkPolicyAllOf {
	this := VnicEthNetworkPolicyAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var targetPlatform string = "Standalone"
	this.TargetPlatform = &targetPlatform
	return &this
}

// NewVnicEthNetworkPolicyAllOfWithDefaults instantiates a new VnicEthNetworkPolicyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicEthNetworkPolicyAllOfWithDefaults() *VnicEthNetworkPolicyAllOf {
	this := VnicEthNetworkPolicyAllOf{}
	var classId string = "vnic.EthNetworkPolicy"
	this.ClassId = classId
	var objectType string = "vnic.EthNetworkPolicy"
	this.ObjectType = objectType
	var targetPlatform string = "Standalone"
	this.TargetPlatform = &targetPlatform
	return &this
}

// GetClassId returns the ClassId field value
func (o *VnicEthNetworkPolicyAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VnicEthNetworkPolicyAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VnicEthNetworkPolicyAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VnicEthNetworkPolicyAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VnicEthNetworkPolicyAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VnicEthNetworkPolicyAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetTargetPlatform returns the TargetPlatform field value if set, zero value otherwise.
// Deprecated
func (o *VnicEthNetworkPolicyAllOf) GetTargetPlatform() string {
	if o == nil || o.TargetPlatform == nil {
		var ret string
		return ret
	}
	return *o.TargetPlatform
}

// GetTargetPlatformOk returns a tuple with the TargetPlatform field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *VnicEthNetworkPolicyAllOf) GetTargetPlatformOk() (*string, bool) {
	if o == nil || o.TargetPlatform == nil {
		return nil, false
	}
	return o.TargetPlatform, true
}

// HasTargetPlatform returns a boolean if a field has been set.
func (o *VnicEthNetworkPolicyAllOf) HasTargetPlatform() bool {
	if o != nil && o.TargetPlatform != nil {
		return true
	}

	return false
}

// SetTargetPlatform gets a reference to the given string and assigns it to the TargetPlatform field.
// Deprecated
func (o *VnicEthNetworkPolicyAllOf) SetTargetPlatform(v string) {
	o.TargetPlatform = &v
}

// GetVlanSettings returns the VlanSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicEthNetworkPolicyAllOf) GetVlanSettings() VnicVlanSettings {
	if o == nil || o.VlanSettings.Get() == nil {
		var ret VnicVlanSettings
		return ret
	}
	return *o.VlanSettings.Get()
}

// GetVlanSettingsOk returns a tuple with the VlanSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicEthNetworkPolicyAllOf) GetVlanSettingsOk() (*VnicVlanSettings, bool) {
	if o == nil {
		return nil, false
	}
	return o.VlanSettings.Get(), o.VlanSettings.IsSet()
}

// HasVlanSettings returns a boolean if a field has been set.
func (o *VnicEthNetworkPolicyAllOf) HasVlanSettings() bool {
	if o != nil && o.VlanSettings.IsSet() {
		return true
	}

	return false
}

// SetVlanSettings gets a reference to the given NullableVnicVlanSettings and assigns it to the VlanSettings field.
func (o *VnicEthNetworkPolicyAllOf) SetVlanSettings(v VnicVlanSettings) {
	o.VlanSettings.Set(&v)
}

// SetVlanSettingsNil sets the value for VlanSettings to be an explicit nil
func (o *VnicEthNetworkPolicyAllOf) SetVlanSettingsNil() {
	o.VlanSettings.Set(nil)
}

// UnsetVlanSettings ensures that no value is present for VlanSettings, not even an explicit nil
func (o *VnicEthNetworkPolicyAllOf) UnsetVlanSettings() {
	o.VlanSettings.Unset()
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *VnicEthNetworkPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthNetworkPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *VnicEthNetworkPolicyAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *VnicEthNetworkPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o VnicEthNetworkPolicyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.TargetPlatform != nil {
		toSerialize["TargetPlatform"] = o.TargetPlatform
	}
	if o.VlanSettings.IsSet() {
		toSerialize["VlanSettings"] = o.VlanSettings.Get()
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VnicEthNetworkPolicyAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varVnicEthNetworkPolicyAllOf := _VnicEthNetworkPolicyAllOf{}

	if err = json.Unmarshal(bytes, &varVnicEthNetworkPolicyAllOf); err == nil {
		*o = VnicEthNetworkPolicyAllOf(varVnicEthNetworkPolicyAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "TargetPlatform")
		delete(additionalProperties, "VlanSettings")
		delete(additionalProperties, "Organization")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVnicEthNetworkPolicyAllOf struct {
	value *VnicEthNetworkPolicyAllOf
	isSet bool
}

func (v NullableVnicEthNetworkPolicyAllOf) Get() *VnicEthNetworkPolicyAllOf {
	return v.value
}

func (v *NullableVnicEthNetworkPolicyAllOf) Set(val *VnicEthNetworkPolicyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicEthNetworkPolicyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicEthNetworkPolicyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicEthNetworkPolicyAllOf(val *VnicEthNetworkPolicyAllOf) *NullableVnicEthNetworkPolicyAllOf {
	return &NullableVnicEthNetworkPolicyAllOf{value: val, isSet: true}
}

func (v NullableVnicEthNetworkPolicyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicEthNetworkPolicyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
