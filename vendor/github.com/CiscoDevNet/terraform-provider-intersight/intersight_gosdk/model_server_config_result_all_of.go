/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-6484
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// ServerConfigResultAllOf Definition of the list of properties defined in 'server.ConfigResult', excluding properties defined in parent classes.
type ServerConfigResultAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType      string                         `json:"ObjectType"`
	AppliedPolicies []PolicyPolicyStatus           `json:"AppliedPolicies,omitempty"`
	Profile         *ServerBaseProfileRelationship `json:"Profile,omitempty"`
	// An array of relationships to serverConfigResultEntry resources.
	ResultEntries        []ServerConfigResultEntryRelationship `json:"ResultEntries,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServerConfigResultAllOf ServerConfigResultAllOf

// NewServerConfigResultAllOf instantiates a new ServerConfigResultAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerConfigResultAllOf(classId string, objectType string) *ServerConfigResultAllOf {
	this := ServerConfigResultAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewServerConfigResultAllOfWithDefaults instantiates a new ServerConfigResultAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerConfigResultAllOfWithDefaults() *ServerConfigResultAllOf {
	this := ServerConfigResultAllOf{}
	var classId string = "server.ConfigResult"
	this.ClassId = classId
	var objectType string = "server.ConfigResult"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ServerConfigResultAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ServerConfigResultAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ServerConfigResultAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ServerConfigResultAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ServerConfigResultAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ServerConfigResultAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAppliedPolicies returns the AppliedPolicies field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServerConfigResultAllOf) GetAppliedPolicies() []PolicyPolicyStatus {
	if o == nil {
		var ret []PolicyPolicyStatus
		return ret
	}
	return o.AppliedPolicies
}

// GetAppliedPoliciesOk returns a tuple with the AppliedPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServerConfigResultAllOf) GetAppliedPoliciesOk() ([]PolicyPolicyStatus, bool) {
	if o == nil || o.AppliedPolicies == nil {
		return nil, false
	}
	return o.AppliedPolicies, true
}

// HasAppliedPolicies returns a boolean if a field has been set.
func (o *ServerConfigResultAllOf) HasAppliedPolicies() bool {
	if o != nil && o.AppliedPolicies != nil {
		return true
	}

	return false
}

// SetAppliedPolicies gets a reference to the given []PolicyPolicyStatus and assigns it to the AppliedPolicies field.
func (o *ServerConfigResultAllOf) SetAppliedPolicies(v []PolicyPolicyStatus) {
	o.AppliedPolicies = v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *ServerConfigResultAllOf) GetProfile() ServerBaseProfileRelationship {
	if o == nil || o.Profile == nil {
		var ret ServerBaseProfileRelationship
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigResultAllOf) GetProfileOk() (*ServerBaseProfileRelationship, bool) {
	if o == nil || o.Profile == nil {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *ServerConfigResultAllOf) HasProfile() bool {
	if o != nil && o.Profile != nil {
		return true
	}

	return false
}

// SetProfile gets a reference to the given ServerBaseProfileRelationship and assigns it to the Profile field.
func (o *ServerConfigResultAllOf) SetProfile(v ServerBaseProfileRelationship) {
	o.Profile = &v
}

// GetResultEntries returns the ResultEntries field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServerConfigResultAllOf) GetResultEntries() []ServerConfigResultEntryRelationship {
	if o == nil {
		var ret []ServerConfigResultEntryRelationship
		return ret
	}
	return o.ResultEntries
}

// GetResultEntriesOk returns a tuple with the ResultEntries field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServerConfigResultAllOf) GetResultEntriesOk() ([]ServerConfigResultEntryRelationship, bool) {
	if o == nil || o.ResultEntries == nil {
		return nil, false
	}
	return o.ResultEntries, true
}

// HasResultEntries returns a boolean if a field has been set.
func (o *ServerConfigResultAllOf) HasResultEntries() bool {
	if o != nil && o.ResultEntries != nil {
		return true
	}

	return false
}

// SetResultEntries gets a reference to the given []ServerConfigResultEntryRelationship and assigns it to the ResultEntries field.
func (o *ServerConfigResultAllOf) SetResultEntries(v []ServerConfigResultEntryRelationship) {
	o.ResultEntries = v
}

func (o ServerConfigResultAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AppliedPolicies != nil {
		toSerialize["AppliedPolicies"] = o.AppliedPolicies
	}
	if o.Profile != nil {
		toSerialize["Profile"] = o.Profile
	}
	if o.ResultEntries != nil {
		toSerialize["ResultEntries"] = o.ResultEntries
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ServerConfigResultAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varServerConfigResultAllOf := _ServerConfigResultAllOf{}

	if err = json.Unmarshal(bytes, &varServerConfigResultAllOf); err == nil {
		*o = ServerConfigResultAllOf(varServerConfigResultAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AppliedPolicies")
		delete(additionalProperties, "Profile")
		delete(additionalProperties, "ResultEntries")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServerConfigResultAllOf struct {
	value *ServerConfigResultAllOf
	isSet bool
}

func (v NullableServerConfigResultAllOf) Get() *ServerConfigResultAllOf {
	return v.value
}

func (v *NullableServerConfigResultAllOf) Set(val *ServerConfigResultAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableServerConfigResultAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableServerConfigResultAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerConfigResultAllOf(val *ServerConfigResultAllOf) *NullableServerConfigResultAllOf {
	return &NullableServerConfigResultAllOf{value: val, isSet: true}
}

func (v NullableServerConfigResultAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerConfigResultAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
