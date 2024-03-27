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
	"reflect"
	"strings"
)

// FabricLinkAggregationPolicy A policy to configure the link settings for all the port channels (including LACP).
type FabricLinkAggregationPolicy struct {
	PolicyAbstractPolicy
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Flag used to indicate whether LACP PDUs are to be sent 'fast', i.e., every 1 second. * `normal` - The expanded 4th generation UCS Fabric Interconnect with 108 ports. * `fast` - The standard 4th generation UCS Fabric Interconnect with 54 ports.
	LacpRate *string `json:"LacpRate,omitempty"`
	// Flag tells the switch whether to suspend the port if it didn’t receive LACP PDU.
	SuspendIndividual    *bool                                 `json:"SuspendIndividual,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricLinkAggregationPolicy FabricLinkAggregationPolicy

// NewFabricLinkAggregationPolicy instantiates a new FabricLinkAggregationPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricLinkAggregationPolicy(classId string, objectType string) *FabricLinkAggregationPolicy {
	this := FabricLinkAggregationPolicy{}
	this.ClassId = classId
	this.ObjectType = objectType
	var lacpRate string = "normal"
	this.LacpRate = &lacpRate
	return &this
}

// NewFabricLinkAggregationPolicyWithDefaults instantiates a new FabricLinkAggregationPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricLinkAggregationPolicyWithDefaults() *FabricLinkAggregationPolicy {
	this := FabricLinkAggregationPolicy{}
	var classId string = "fabric.LinkAggregationPolicy"
	this.ClassId = classId
	var objectType string = "fabric.LinkAggregationPolicy"
	this.ObjectType = objectType
	var lacpRate string = "normal"
	this.LacpRate = &lacpRate
	return &this
}

// GetClassId returns the ClassId field value
func (o *FabricLinkAggregationPolicy) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FabricLinkAggregationPolicy) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FabricLinkAggregationPolicy) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FabricLinkAggregationPolicy) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FabricLinkAggregationPolicy) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FabricLinkAggregationPolicy) SetObjectType(v string) {
	o.ObjectType = v
}

// GetLacpRate returns the LacpRate field value if set, zero value otherwise.
func (o *FabricLinkAggregationPolicy) GetLacpRate() string {
	if o == nil || o.LacpRate == nil {
		var ret string
		return ret
	}
	return *o.LacpRate
}

// GetLacpRateOk returns a tuple with the LacpRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricLinkAggregationPolicy) GetLacpRateOk() (*string, bool) {
	if o == nil || o.LacpRate == nil {
		return nil, false
	}
	return o.LacpRate, true
}

// HasLacpRate returns a boolean if a field has been set.
func (o *FabricLinkAggregationPolicy) HasLacpRate() bool {
	if o != nil && o.LacpRate != nil {
		return true
	}

	return false
}

// SetLacpRate gets a reference to the given string and assigns it to the LacpRate field.
func (o *FabricLinkAggregationPolicy) SetLacpRate(v string) {
	o.LacpRate = &v
}

// GetSuspendIndividual returns the SuspendIndividual field value if set, zero value otherwise.
func (o *FabricLinkAggregationPolicy) GetSuspendIndividual() bool {
	if o == nil || o.SuspendIndividual == nil {
		var ret bool
		return ret
	}
	return *o.SuspendIndividual
}

// GetSuspendIndividualOk returns a tuple with the SuspendIndividual field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricLinkAggregationPolicy) GetSuspendIndividualOk() (*bool, bool) {
	if o == nil || o.SuspendIndividual == nil {
		return nil, false
	}
	return o.SuspendIndividual, true
}

// HasSuspendIndividual returns a boolean if a field has been set.
func (o *FabricLinkAggregationPolicy) HasSuspendIndividual() bool {
	if o != nil && o.SuspendIndividual != nil {
		return true
	}

	return false
}

// SetSuspendIndividual gets a reference to the given bool and assigns it to the SuspendIndividual field.
func (o *FabricLinkAggregationPolicy) SetSuspendIndividual(v bool) {
	o.SuspendIndividual = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *FabricLinkAggregationPolicy) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricLinkAggregationPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *FabricLinkAggregationPolicy) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *FabricLinkAggregationPolicy) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o FabricLinkAggregationPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractPolicy, errPolicyAbstractPolicy := json.Marshal(o.PolicyAbstractPolicy)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	errPolicyAbstractPolicy = json.Unmarshal([]byte(serializedPolicyAbstractPolicy), &toSerialize)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.LacpRate != nil {
		toSerialize["LacpRate"] = o.LacpRate
	}
	if o.SuspendIndividual != nil {
		toSerialize["SuspendIndividual"] = o.SuspendIndividual
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FabricLinkAggregationPolicy) UnmarshalJSON(bytes []byte) (err error) {
	type FabricLinkAggregationPolicyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Flag used to indicate whether LACP PDUs are to be sent 'fast', i.e., every 1 second. * `normal` - The expanded 4th generation UCS Fabric Interconnect with 108 ports. * `fast` - The standard 4th generation UCS Fabric Interconnect with 54 ports.
		LacpRate *string `json:"LacpRate,omitempty"`
		// Flag tells the switch whether to suspend the port if it didn’t receive LACP PDU.
		SuspendIndividual *bool                                 `json:"SuspendIndividual,omitempty"`
		Organization      *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	}

	varFabricLinkAggregationPolicyWithoutEmbeddedStruct := FabricLinkAggregationPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varFabricLinkAggregationPolicyWithoutEmbeddedStruct)
	if err == nil {
		varFabricLinkAggregationPolicy := _FabricLinkAggregationPolicy{}
		varFabricLinkAggregationPolicy.ClassId = varFabricLinkAggregationPolicyWithoutEmbeddedStruct.ClassId
		varFabricLinkAggregationPolicy.ObjectType = varFabricLinkAggregationPolicyWithoutEmbeddedStruct.ObjectType
		varFabricLinkAggregationPolicy.LacpRate = varFabricLinkAggregationPolicyWithoutEmbeddedStruct.LacpRate
		varFabricLinkAggregationPolicy.SuspendIndividual = varFabricLinkAggregationPolicyWithoutEmbeddedStruct.SuspendIndividual
		varFabricLinkAggregationPolicy.Organization = varFabricLinkAggregationPolicyWithoutEmbeddedStruct.Organization
		*o = FabricLinkAggregationPolicy(varFabricLinkAggregationPolicy)
	} else {
		return err
	}

	varFabricLinkAggregationPolicy := _FabricLinkAggregationPolicy{}

	err = json.Unmarshal(bytes, &varFabricLinkAggregationPolicy)
	if err == nil {
		o.PolicyAbstractPolicy = varFabricLinkAggregationPolicy.PolicyAbstractPolicy
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "LacpRate")
		delete(additionalProperties, "SuspendIndividual")
		delete(additionalProperties, "Organization")

		// remove fields from embedded structs
		reflectPolicyAbstractPolicy := reflect.ValueOf(o.PolicyAbstractPolicy)
		for i := 0; i < reflectPolicyAbstractPolicy.Type().NumField(); i++ {
			t := reflectPolicyAbstractPolicy.Type().Field(i)

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

type NullableFabricLinkAggregationPolicy struct {
	value *FabricLinkAggregationPolicy
	isSet bool
}

func (v NullableFabricLinkAggregationPolicy) Get() *FabricLinkAggregationPolicy {
	return v.value
}

func (v *NullableFabricLinkAggregationPolicy) Set(val *FabricLinkAggregationPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricLinkAggregationPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricLinkAggregationPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricLinkAggregationPolicy(val *FabricLinkAggregationPolicy) *NullableFabricLinkAggregationPolicy {
	return &NullableFabricLinkAggregationPolicy{value: val, isSet: true}
}

func (v NullableFabricLinkAggregationPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricLinkAggregationPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
