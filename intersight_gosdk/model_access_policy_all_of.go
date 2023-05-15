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

// AccessPolicyAllOf Definition of the list of properties defined in 'access.Policy', excluding properties defined in parent classes.
type AccessPolicyAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType        string                          `json:"ObjectType"`
	AddressType       NullableAccessAddressType       `json:"AddressType,omitempty"`
	ConfigurationType NullableAccessConfigurationType `json:"ConfigurationType,omitempty"`
	// VLAN to be used for server access over Inband network.
	InbandVlan      *int64                                `json:"InbandVlan,omitempty"`
	InbandIpPool    *IppoolPoolRelationship               `json:"InbandIpPool,omitempty"`
	InbandVrf       *VrfVrfRelationship                   `json:"InbandVrf,omitempty"`
	Organization    *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	OutOfBandIpPool *IppoolPoolRelationship               `json:"OutOfBandIpPool,omitempty"`
	OutOfBandVrf    *VrfVrfRelationship                   `json:"OutOfBandVrf,omitempty"`
	// An array of relationships to policyAbstractConfigProfile resources.
	Profiles             []PolicyAbstractConfigProfileRelationship `json:"Profiles,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccessPolicyAllOf AccessPolicyAllOf

// NewAccessPolicyAllOf instantiates a new AccessPolicyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessPolicyAllOf(classId string, objectType string) *AccessPolicyAllOf {
	this := AccessPolicyAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAccessPolicyAllOfWithDefaults instantiates a new AccessPolicyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessPolicyAllOfWithDefaults() *AccessPolicyAllOf {
	this := AccessPolicyAllOf{}
	var classId string = "access.Policy"
	this.ClassId = classId
	var objectType string = "access.Policy"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AccessPolicyAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AccessPolicyAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AccessPolicyAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AccessPolicyAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AccessPolicyAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AccessPolicyAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAddressType returns the AddressType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessPolicyAllOf) GetAddressType() AccessAddressType {
	if o == nil || o.AddressType.Get() == nil {
		var ret AccessAddressType
		return ret
	}
	return *o.AddressType.Get()
}

// GetAddressTypeOk returns a tuple with the AddressType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessPolicyAllOf) GetAddressTypeOk() (*AccessAddressType, bool) {
	if o == nil {
		return nil, false
	}
	return o.AddressType.Get(), o.AddressType.IsSet()
}

// HasAddressType returns a boolean if a field has been set.
func (o *AccessPolicyAllOf) HasAddressType() bool {
	if o != nil && o.AddressType.IsSet() {
		return true
	}

	return false
}

// SetAddressType gets a reference to the given NullableAccessAddressType and assigns it to the AddressType field.
func (o *AccessPolicyAllOf) SetAddressType(v AccessAddressType) {
	o.AddressType.Set(&v)
}

// SetAddressTypeNil sets the value for AddressType to be an explicit nil
func (o *AccessPolicyAllOf) SetAddressTypeNil() {
	o.AddressType.Set(nil)
}

// UnsetAddressType ensures that no value is present for AddressType, not even an explicit nil
func (o *AccessPolicyAllOf) UnsetAddressType() {
	o.AddressType.Unset()
}

// GetConfigurationType returns the ConfigurationType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessPolicyAllOf) GetConfigurationType() AccessConfigurationType {
	if o == nil || o.ConfigurationType.Get() == nil {
		var ret AccessConfigurationType
		return ret
	}
	return *o.ConfigurationType.Get()
}

// GetConfigurationTypeOk returns a tuple with the ConfigurationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessPolicyAllOf) GetConfigurationTypeOk() (*AccessConfigurationType, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigurationType.Get(), o.ConfigurationType.IsSet()
}

// HasConfigurationType returns a boolean if a field has been set.
func (o *AccessPolicyAllOf) HasConfigurationType() bool {
	if o != nil && o.ConfigurationType.IsSet() {
		return true
	}

	return false
}

// SetConfigurationType gets a reference to the given NullableAccessConfigurationType and assigns it to the ConfigurationType field.
func (o *AccessPolicyAllOf) SetConfigurationType(v AccessConfigurationType) {
	o.ConfigurationType.Set(&v)
}

// SetConfigurationTypeNil sets the value for ConfigurationType to be an explicit nil
func (o *AccessPolicyAllOf) SetConfigurationTypeNil() {
	o.ConfigurationType.Set(nil)
}

// UnsetConfigurationType ensures that no value is present for ConfigurationType, not even an explicit nil
func (o *AccessPolicyAllOf) UnsetConfigurationType() {
	o.ConfigurationType.Unset()
}

// GetInbandVlan returns the InbandVlan field value if set, zero value otherwise.
func (o *AccessPolicyAllOf) GetInbandVlan() int64 {
	if o == nil || o.InbandVlan == nil {
		var ret int64
		return ret
	}
	return *o.InbandVlan
}

// GetInbandVlanOk returns a tuple with the InbandVlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicyAllOf) GetInbandVlanOk() (*int64, bool) {
	if o == nil || o.InbandVlan == nil {
		return nil, false
	}
	return o.InbandVlan, true
}

// HasInbandVlan returns a boolean if a field has been set.
func (o *AccessPolicyAllOf) HasInbandVlan() bool {
	if o != nil && o.InbandVlan != nil {
		return true
	}

	return false
}

// SetInbandVlan gets a reference to the given int64 and assigns it to the InbandVlan field.
func (o *AccessPolicyAllOf) SetInbandVlan(v int64) {
	o.InbandVlan = &v
}

// GetInbandIpPool returns the InbandIpPool field value if set, zero value otherwise.
func (o *AccessPolicyAllOf) GetInbandIpPool() IppoolPoolRelationship {
	if o == nil || o.InbandIpPool == nil {
		var ret IppoolPoolRelationship
		return ret
	}
	return *o.InbandIpPool
}

// GetInbandIpPoolOk returns a tuple with the InbandIpPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicyAllOf) GetInbandIpPoolOk() (*IppoolPoolRelationship, bool) {
	if o == nil || o.InbandIpPool == nil {
		return nil, false
	}
	return o.InbandIpPool, true
}

// HasInbandIpPool returns a boolean if a field has been set.
func (o *AccessPolicyAllOf) HasInbandIpPool() bool {
	if o != nil && o.InbandIpPool != nil {
		return true
	}

	return false
}

// SetInbandIpPool gets a reference to the given IppoolPoolRelationship and assigns it to the InbandIpPool field.
func (o *AccessPolicyAllOf) SetInbandIpPool(v IppoolPoolRelationship) {
	o.InbandIpPool = &v
}

// GetInbandVrf returns the InbandVrf field value if set, zero value otherwise.
func (o *AccessPolicyAllOf) GetInbandVrf() VrfVrfRelationship {
	if o == nil || o.InbandVrf == nil {
		var ret VrfVrfRelationship
		return ret
	}
	return *o.InbandVrf
}

// GetInbandVrfOk returns a tuple with the InbandVrf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicyAllOf) GetInbandVrfOk() (*VrfVrfRelationship, bool) {
	if o == nil || o.InbandVrf == nil {
		return nil, false
	}
	return o.InbandVrf, true
}

// HasInbandVrf returns a boolean if a field has been set.
func (o *AccessPolicyAllOf) HasInbandVrf() bool {
	if o != nil && o.InbandVrf != nil {
		return true
	}

	return false
}

// SetInbandVrf gets a reference to the given VrfVrfRelationship and assigns it to the InbandVrf field.
func (o *AccessPolicyAllOf) SetInbandVrf(v VrfVrfRelationship) {
	o.InbandVrf = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *AccessPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *AccessPolicyAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *AccessPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetOutOfBandIpPool returns the OutOfBandIpPool field value if set, zero value otherwise.
func (o *AccessPolicyAllOf) GetOutOfBandIpPool() IppoolPoolRelationship {
	if o == nil || o.OutOfBandIpPool == nil {
		var ret IppoolPoolRelationship
		return ret
	}
	return *o.OutOfBandIpPool
}

// GetOutOfBandIpPoolOk returns a tuple with the OutOfBandIpPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicyAllOf) GetOutOfBandIpPoolOk() (*IppoolPoolRelationship, bool) {
	if o == nil || o.OutOfBandIpPool == nil {
		return nil, false
	}
	return o.OutOfBandIpPool, true
}

// HasOutOfBandIpPool returns a boolean if a field has been set.
func (o *AccessPolicyAllOf) HasOutOfBandIpPool() bool {
	if o != nil && o.OutOfBandIpPool != nil {
		return true
	}

	return false
}

// SetOutOfBandIpPool gets a reference to the given IppoolPoolRelationship and assigns it to the OutOfBandIpPool field.
func (o *AccessPolicyAllOf) SetOutOfBandIpPool(v IppoolPoolRelationship) {
	o.OutOfBandIpPool = &v
}

// GetOutOfBandVrf returns the OutOfBandVrf field value if set, zero value otherwise.
func (o *AccessPolicyAllOf) GetOutOfBandVrf() VrfVrfRelationship {
	if o == nil || o.OutOfBandVrf == nil {
		var ret VrfVrfRelationship
		return ret
	}
	return *o.OutOfBandVrf
}

// GetOutOfBandVrfOk returns a tuple with the OutOfBandVrf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicyAllOf) GetOutOfBandVrfOk() (*VrfVrfRelationship, bool) {
	if o == nil || o.OutOfBandVrf == nil {
		return nil, false
	}
	return o.OutOfBandVrf, true
}

// HasOutOfBandVrf returns a boolean if a field has been set.
func (o *AccessPolicyAllOf) HasOutOfBandVrf() bool {
	if o != nil && o.OutOfBandVrf != nil {
		return true
	}

	return false
}

// SetOutOfBandVrf gets a reference to the given VrfVrfRelationship and assigns it to the OutOfBandVrf field.
func (o *AccessPolicyAllOf) SetOutOfBandVrf(v VrfVrfRelationship) {
	o.OutOfBandVrf = &v
}

// GetProfiles returns the Profiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessPolicyAllOf) GetProfiles() []PolicyAbstractConfigProfileRelationship {
	if o == nil {
		var ret []PolicyAbstractConfigProfileRelationship
		return ret
	}
	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessPolicyAllOf) GetProfilesOk() ([]PolicyAbstractConfigProfileRelationship, bool) {
	if o == nil || o.Profiles == nil {
		return nil, false
	}
	return o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *AccessPolicyAllOf) HasProfiles() bool {
	if o != nil && o.Profiles != nil {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []PolicyAbstractConfigProfileRelationship and assigns it to the Profiles field.
func (o *AccessPolicyAllOf) SetProfiles(v []PolicyAbstractConfigProfileRelationship) {
	o.Profiles = v
}

func (o AccessPolicyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AddressType.IsSet() {
		toSerialize["AddressType"] = o.AddressType.Get()
	}
	if o.ConfigurationType.IsSet() {
		toSerialize["ConfigurationType"] = o.ConfigurationType.Get()
	}
	if o.InbandVlan != nil {
		toSerialize["InbandVlan"] = o.InbandVlan
	}
	if o.InbandIpPool != nil {
		toSerialize["InbandIpPool"] = o.InbandIpPool
	}
	if o.InbandVrf != nil {
		toSerialize["InbandVrf"] = o.InbandVrf
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.OutOfBandIpPool != nil {
		toSerialize["OutOfBandIpPool"] = o.OutOfBandIpPool
	}
	if o.OutOfBandVrf != nil {
		toSerialize["OutOfBandVrf"] = o.OutOfBandVrf
	}
	if o.Profiles != nil {
		toSerialize["Profiles"] = o.Profiles
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AccessPolicyAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varAccessPolicyAllOf := _AccessPolicyAllOf{}

	if err = json.Unmarshal(bytes, &varAccessPolicyAllOf); err == nil {
		*o = AccessPolicyAllOf(varAccessPolicyAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AddressType")
		delete(additionalProperties, "ConfigurationType")
		delete(additionalProperties, "InbandVlan")
		delete(additionalProperties, "InbandIpPool")
		delete(additionalProperties, "InbandVrf")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "OutOfBandIpPool")
		delete(additionalProperties, "OutOfBandVrf")
		delete(additionalProperties, "Profiles")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccessPolicyAllOf struct {
	value *AccessPolicyAllOf
	isSet bool
}

func (v NullableAccessPolicyAllOf) Get() *AccessPolicyAllOf {
	return v.value
}

func (v *NullableAccessPolicyAllOf) Set(val *AccessPolicyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessPolicyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessPolicyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessPolicyAllOf(val *AccessPolicyAllOf) *NullableAccessPolicyAllOf {
	return &NullableAccessPolicyAllOf{value: val, isSet: true}
}

func (v NullableAccessPolicyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessPolicyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
