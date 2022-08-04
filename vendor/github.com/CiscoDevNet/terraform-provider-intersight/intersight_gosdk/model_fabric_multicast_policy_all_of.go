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
)

// FabricMulticastPolicyAllOf Definition of the list of properties defined in 'fabric.MulticastPolicy', excluding properties defined in parent classes.
type FabricMulticastPolicyAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Used to define the IGMP Querier IP address.
	QuerierIpAddress *string `json:"QuerierIpAddress,omitempty"`
	// Used to define the IGMP Querier IP address of the peer switch.
	QuerierIpAddressPeer *string `json:"QuerierIpAddressPeer,omitempty"`
	// Administrative state of the IGMP Querier for this VLAN. * `Disabled` - Admin configured Disabled State. * `Enabled` - Admin configured Enabled State.
	QuerierState *string `json:"QuerierState,omitempty"`
	// Administrative state of the IGMP Snooping for this VLAN. * `Enabled` - Admin configured Enabled State. * `Disabled` - Admin configured Disabled State.
	SnoopingState        *string                               `json:"SnoopingState,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricMulticastPolicyAllOf FabricMulticastPolicyAllOf

// NewFabricMulticastPolicyAllOf instantiates a new FabricMulticastPolicyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricMulticastPolicyAllOf(classId string, objectType string) *FabricMulticastPolicyAllOf {
	this := FabricMulticastPolicyAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var querierState string = "Disabled"
	this.QuerierState = &querierState
	var snoopingState string = "Enabled"
	this.SnoopingState = &snoopingState
	return &this
}

// NewFabricMulticastPolicyAllOfWithDefaults instantiates a new FabricMulticastPolicyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricMulticastPolicyAllOfWithDefaults() *FabricMulticastPolicyAllOf {
	this := FabricMulticastPolicyAllOf{}
	var classId string = "fabric.MulticastPolicy"
	this.ClassId = classId
	var objectType string = "fabric.MulticastPolicy"
	this.ObjectType = objectType
	var querierState string = "Disabled"
	this.QuerierState = &querierState
	var snoopingState string = "Enabled"
	this.SnoopingState = &snoopingState
	return &this
}

// GetClassId returns the ClassId field value
func (o *FabricMulticastPolicyAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FabricMulticastPolicyAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FabricMulticastPolicyAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FabricMulticastPolicyAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FabricMulticastPolicyAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FabricMulticastPolicyAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetQuerierIpAddress returns the QuerierIpAddress field value if set, zero value otherwise.
func (o *FabricMulticastPolicyAllOf) GetQuerierIpAddress() string {
	if o == nil || o.QuerierIpAddress == nil {
		var ret string
		return ret
	}
	return *o.QuerierIpAddress
}

// GetQuerierIpAddressOk returns a tuple with the QuerierIpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricMulticastPolicyAllOf) GetQuerierIpAddressOk() (*string, bool) {
	if o == nil || o.QuerierIpAddress == nil {
		return nil, false
	}
	return o.QuerierIpAddress, true
}

// HasQuerierIpAddress returns a boolean if a field has been set.
func (o *FabricMulticastPolicyAllOf) HasQuerierIpAddress() bool {
	if o != nil && o.QuerierIpAddress != nil {
		return true
	}

	return false
}

// SetQuerierIpAddress gets a reference to the given string and assigns it to the QuerierIpAddress field.
func (o *FabricMulticastPolicyAllOf) SetQuerierIpAddress(v string) {
	o.QuerierIpAddress = &v
}

// GetQuerierIpAddressPeer returns the QuerierIpAddressPeer field value if set, zero value otherwise.
func (o *FabricMulticastPolicyAllOf) GetQuerierIpAddressPeer() string {
	if o == nil || o.QuerierIpAddressPeer == nil {
		var ret string
		return ret
	}
	return *o.QuerierIpAddressPeer
}

// GetQuerierIpAddressPeerOk returns a tuple with the QuerierIpAddressPeer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricMulticastPolicyAllOf) GetQuerierIpAddressPeerOk() (*string, bool) {
	if o == nil || o.QuerierIpAddressPeer == nil {
		return nil, false
	}
	return o.QuerierIpAddressPeer, true
}

// HasQuerierIpAddressPeer returns a boolean if a field has been set.
func (o *FabricMulticastPolicyAllOf) HasQuerierIpAddressPeer() bool {
	if o != nil && o.QuerierIpAddressPeer != nil {
		return true
	}

	return false
}

// SetQuerierIpAddressPeer gets a reference to the given string and assigns it to the QuerierIpAddressPeer field.
func (o *FabricMulticastPolicyAllOf) SetQuerierIpAddressPeer(v string) {
	o.QuerierIpAddressPeer = &v
}

// GetQuerierState returns the QuerierState field value if set, zero value otherwise.
func (o *FabricMulticastPolicyAllOf) GetQuerierState() string {
	if o == nil || o.QuerierState == nil {
		var ret string
		return ret
	}
	return *o.QuerierState
}

// GetQuerierStateOk returns a tuple with the QuerierState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricMulticastPolicyAllOf) GetQuerierStateOk() (*string, bool) {
	if o == nil || o.QuerierState == nil {
		return nil, false
	}
	return o.QuerierState, true
}

// HasQuerierState returns a boolean if a field has been set.
func (o *FabricMulticastPolicyAllOf) HasQuerierState() bool {
	if o != nil && o.QuerierState != nil {
		return true
	}

	return false
}

// SetQuerierState gets a reference to the given string and assigns it to the QuerierState field.
func (o *FabricMulticastPolicyAllOf) SetQuerierState(v string) {
	o.QuerierState = &v
}

// GetSnoopingState returns the SnoopingState field value if set, zero value otherwise.
func (o *FabricMulticastPolicyAllOf) GetSnoopingState() string {
	if o == nil || o.SnoopingState == nil {
		var ret string
		return ret
	}
	return *o.SnoopingState
}

// GetSnoopingStateOk returns a tuple with the SnoopingState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricMulticastPolicyAllOf) GetSnoopingStateOk() (*string, bool) {
	if o == nil || o.SnoopingState == nil {
		return nil, false
	}
	return o.SnoopingState, true
}

// HasSnoopingState returns a boolean if a field has been set.
func (o *FabricMulticastPolicyAllOf) HasSnoopingState() bool {
	if o != nil && o.SnoopingState != nil {
		return true
	}

	return false
}

// SetSnoopingState gets a reference to the given string and assigns it to the SnoopingState field.
func (o *FabricMulticastPolicyAllOf) SetSnoopingState(v string) {
	o.SnoopingState = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *FabricMulticastPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricMulticastPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *FabricMulticastPolicyAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *FabricMulticastPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o FabricMulticastPolicyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.QuerierIpAddress != nil {
		toSerialize["QuerierIpAddress"] = o.QuerierIpAddress
	}
	if o.QuerierIpAddressPeer != nil {
		toSerialize["QuerierIpAddressPeer"] = o.QuerierIpAddressPeer
	}
	if o.QuerierState != nil {
		toSerialize["QuerierState"] = o.QuerierState
	}
	if o.SnoopingState != nil {
		toSerialize["SnoopingState"] = o.SnoopingState
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FabricMulticastPolicyAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varFabricMulticastPolicyAllOf := _FabricMulticastPolicyAllOf{}

	if err = json.Unmarshal(bytes, &varFabricMulticastPolicyAllOf); err == nil {
		*o = FabricMulticastPolicyAllOf(varFabricMulticastPolicyAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "QuerierIpAddress")
		delete(additionalProperties, "QuerierIpAddressPeer")
		delete(additionalProperties, "QuerierState")
		delete(additionalProperties, "SnoopingState")
		delete(additionalProperties, "Organization")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFabricMulticastPolicyAllOf struct {
	value *FabricMulticastPolicyAllOf
	isSet bool
}

func (v NullableFabricMulticastPolicyAllOf) Get() *FabricMulticastPolicyAllOf {
	return v.value
}

func (v *NullableFabricMulticastPolicyAllOf) Set(val *FabricMulticastPolicyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricMulticastPolicyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricMulticastPolicyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricMulticastPolicyAllOf(val *FabricMulticastPolicyAllOf) *NullableFabricMulticastPolicyAllOf {
	return &NullableFabricMulticastPolicyAllOf{value: val, isSet: true}
}

func (v NullableFabricMulticastPolicyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricMulticastPolicyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
