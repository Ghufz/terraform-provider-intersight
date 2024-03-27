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

// NiatelemetryNexusCloudAccountAllOf Definition of the list of properties defined in 'niatelemetry.NexusCloudAccount', excluding properties defined in parent classes.
type NiatelemetryNexusCloudAccountAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Count of ACI-type site devices.
	AciCount *int64 `json:"AciCount,omitempty"`
	// Count of NXOS-type site devices.
	NxosCount            *int64                  `json:"NxosCount,omitempty"`
	Account              *IamAccountRelationship `json:"Account,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryNexusCloudAccountAllOf NiatelemetryNexusCloudAccountAllOf

// NewNiatelemetryNexusCloudAccountAllOf instantiates a new NiatelemetryNexusCloudAccountAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryNexusCloudAccountAllOf(classId string, objectType string) *NiatelemetryNexusCloudAccountAllOf {
	this := NiatelemetryNexusCloudAccountAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryNexusCloudAccountAllOfWithDefaults instantiates a new NiatelemetryNexusCloudAccountAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryNexusCloudAccountAllOfWithDefaults() *NiatelemetryNexusCloudAccountAllOf {
	this := NiatelemetryNexusCloudAccountAllOf{}
	var classId string = "niatelemetry.NexusCloudAccount"
	this.ClassId = classId
	var objectType string = "niatelemetry.NexusCloudAccount"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryNexusCloudAccountAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusCloudAccountAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryNexusCloudAccountAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryNexusCloudAccountAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusCloudAccountAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryNexusCloudAccountAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAciCount returns the AciCount field value if set, zero value otherwise.
func (o *NiatelemetryNexusCloudAccountAllOf) GetAciCount() int64 {
	if o == nil || o.AciCount == nil {
		var ret int64
		return ret
	}
	return *o.AciCount
}

// GetAciCountOk returns a tuple with the AciCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusCloudAccountAllOf) GetAciCountOk() (*int64, bool) {
	if o == nil || o.AciCount == nil {
		return nil, false
	}
	return o.AciCount, true
}

// HasAciCount returns a boolean if a field has been set.
func (o *NiatelemetryNexusCloudAccountAllOf) HasAciCount() bool {
	if o != nil && o.AciCount != nil {
		return true
	}

	return false
}

// SetAciCount gets a reference to the given int64 and assigns it to the AciCount field.
func (o *NiatelemetryNexusCloudAccountAllOf) SetAciCount(v int64) {
	o.AciCount = &v
}

// GetNxosCount returns the NxosCount field value if set, zero value otherwise.
func (o *NiatelemetryNexusCloudAccountAllOf) GetNxosCount() int64 {
	if o == nil || o.NxosCount == nil {
		var ret int64
		return ret
	}
	return *o.NxosCount
}

// GetNxosCountOk returns a tuple with the NxosCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusCloudAccountAllOf) GetNxosCountOk() (*int64, bool) {
	if o == nil || o.NxosCount == nil {
		return nil, false
	}
	return o.NxosCount, true
}

// HasNxosCount returns a boolean if a field has been set.
func (o *NiatelemetryNexusCloudAccountAllOf) HasNxosCount() bool {
	if o != nil && o.NxosCount != nil {
		return true
	}

	return false
}

// SetNxosCount gets a reference to the given int64 and assigns it to the NxosCount field.
func (o *NiatelemetryNexusCloudAccountAllOf) SetNxosCount(v int64) {
	o.NxosCount = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *NiatelemetryNexusCloudAccountAllOf) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNexusCloudAccountAllOf) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *NiatelemetryNexusCloudAccountAllOf) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *NiatelemetryNexusCloudAccountAllOf) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

func (o NiatelemetryNexusCloudAccountAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AciCount != nil {
		toSerialize["AciCount"] = o.AciCount
	}
	if o.NxosCount != nil {
		toSerialize["NxosCount"] = o.NxosCount
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryNexusCloudAccountAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varNiatelemetryNexusCloudAccountAllOf := _NiatelemetryNexusCloudAccountAllOf{}

	if err = json.Unmarshal(bytes, &varNiatelemetryNexusCloudAccountAllOf); err == nil {
		*o = NiatelemetryNexusCloudAccountAllOf(varNiatelemetryNexusCloudAccountAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AciCount")
		delete(additionalProperties, "NxosCount")
		delete(additionalProperties, "Account")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNiatelemetryNexusCloudAccountAllOf struct {
	value *NiatelemetryNexusCloudAccountAllOf
	isSet bool
}

func (v NullableNiatelemetryNexusCloudAccountAllOf) Get() *NiatelemetryNexusCloudAccountAllOf {
	return v.value
}

func (v *NullableNiatelemetryNexusCloudAccountAllOf) Set(val *NiatelemetryNexusCloudAccountAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryNexusCloudAccountAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryNexusCloudAccountAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryNexusCloudAccountAllOf(val *NiatelemetryNexusCloudAccountAllOf) *NullableNiatelemetryNexusCloudAccountAllOf {
	return &NullableNiatelemetryNexusCloudAccountAllOf{value: val, isSet: true}
}

func (v NullableNiatelemetryNexusCloudAccountAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryNexusCloudAccountAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
