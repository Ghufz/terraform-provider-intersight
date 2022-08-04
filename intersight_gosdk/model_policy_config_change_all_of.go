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

// PolicyConfigChangeAllOf Definition of the list of properties defined in 'policy.ConfigChange', excluding properties defined in parent classes.
type PolicyConfigChangeAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string   `json:"ObjectType"`
	Changes              []string `json:"Changes,omitempty"`
	Disruptions          []string `json:"Disruptions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PolicyConfigChangeAllOf PolicyConfigChangeAllOf

// NewPolicyConfigChangeAllOf instantiates a new PolicyConfigChangeAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyConfigChangeAllOf(classId string, objectType string) *PolicyConfigChangeAllOf {
	this := PolicyConfigChangeAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewPolicyConfigChangeAllOfWithDefaults instantiates a new PolicyConfigChangeAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyConfigChangeAllOfWithDefaults() *PolicyConfigChangeAllOf {
	this := PolicyConfigChangeAllOf{}
	var classId string = "policy.ConfigChange"
	this.ClassId = classId
	var objectType string = "policy.ConfigChange"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *PolicyConfigChangeAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *PolicyConfigChangeAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *PolicyConfigChangeAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *PolicyConfigChangeAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *PolicyConfigChangeAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *PolicyConfigChangeAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetChanges returns the Changes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PolicyConfigChangeAllOf) GetChanges() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Changes
}

// GetChangesOk returns a tuple with the Changes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PolicyConfigChangeAllOf) GetChangesOk() ([]string, bool) {
	if o == nil || o.Changes == nil {
		return nil, false
	}
	return o.Changes, true
}

// HasChanges returns a boolean if a field has been set.
func (o *PolicyConfigChangeAllOf) HasChanges() bool {
	if o != nil && o.Changes != nil {
		return true
	}

	return false
}

// SetChanges gets a reference to the given []string and assigns it to the Changes field.
func (o *PolicyConfigChangeAllOf) SetChanges(v []string) {
	o.Changes = v
}

// GetDisruptions returns the Disruptions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PolicyConfigChangeAllOf) GetDisruptions() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Disruptions
}

// GetDisruptionsOk returns a tuple with the Disruptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PolicyConfigChangeAllOf) GetDisruptionsOk() ([]string, bool) {
	if o == nil || o.Disruptions == nil {
		return nil, false
	}
	return o.Disruptions, true
}

// HasDisruptions returns a boolean if a field has been set.
func (o *PolicyConfigChangeAllOf) HasDisruptions() bool {
	if o != nil && o.Disruptions != nil {
		return true
	}

	return false
}

// SetDisruptions gets a reference to the given []string and assigns it to the Disruptions field.
func (o *PolicyConfigChangeAllOf) SetDisruptions(v []string) {
	o.Disruptions = v
}

func (o PolicyConfigChangeAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Changes != nil {
		toSerialize["Changes"] = o.Changes
	}
	if o.Disruptions != nil {
		toSerialize["Disruptions"] = o.Disruptions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PolicyConfigChangeAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varPolicyConfigChangeAllOf := _PolicyConfigChangeAllOf{}

	if err = json.Unmarshal(bytes, &varPolicyConfigChangeAllOf); err == nil {
		*o = PolicyConfigChangeAllOf(varPolicyConfigChangeAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Changes")
		delete(additionalProperties, "Disruptions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePolicyConfigChangeAllOf struct {
	value *PolicyConfigChangeAllOf
	isSet bool
}

func (v NullablePolicyConfigChangeAllOf) Get() *PolicyConfigChangeAllOf {
	return v.value
}

func (v *NullablePolicyConfigChangeAllOf) Set(val *PolicyConfigChangeAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyConfigChangeAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyConfigChangeAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyConfigChangeAllOf(val *PolicyConfigChangeAllOf) *NullablePolicyConfigChangeAllOf {
	return &NullablePolicyConfigChangeAllOf{value: val, isSet: true}
}

func (v NullablePolicyConfigChangeAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyConfigChangeAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
