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

// CapabilityServerUpgradeSupportMetaAllOf Definition of the list of properties defined in 'capability.ServerUpgradeSupportMeta', excluding properties defined in parent classes.
type CapabilityServerUpgradeSupportMetaAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Verbose description regarding this group of server.
	Description *string `json:"Description,omitempty"`
	// The target platform for which the mapping is defined.
	Platform *string `json:"Platform,omitempty"`
	// Classification of a set of server models.
	ServerFamily         *string  `json:"ServerFamily,omitempty"`
	SupportedModels      []string `json:"SupportedModels,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CapabilityServerUpgradeSupportMetaAllOf CapabilityServerUpgradeSupportMetaAllOf

// NewCapabilityServerUpgradeSupportMetaAllOf instantiates a new CapabilityServerUpgradeSupportMetaAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilityServerUpgradeSupportMetaAllOf(classId string, objectType string) *CapabilityServerUpgradeSupportMetaAllOf {
	this := CapabilityServerUpgradeSupportMetaAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCapabilityServerUpgradeSupportMetaAllOfWithDefaults instantiates a new CapabilityServerUpgradeSupportMetaAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilityServerUpgradeSupportMetaAllOfWithDefaults() *CapabilityServerUpgradeSupportMetaAllOf {
	this := CapabilityServerUpgradeSupportMetaAllOf{}
	var classId string = "capability.ServerUpgradeSupportMeta"
	this.ClassId = classId
	var objectType string = "capability.ServerUpgradeSupportMeta"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CapabilityServerUpgradeSupportMetaAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CapabilityServerUpgradeSupportMetaAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CapabilityServerUpgradeSupportMetaAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CapabilityServerUpgradeSupportMetaAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CapabilityServerUpgradeSupportMetaAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CapabilityServerUpgradeSupportMetaAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CapabilityServerUpgradeSupportMetaAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityServerUpgradeSupportMetaAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CapabilityServerUpgradeSupportMetaAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CapabilityServerUpgradeSupportMetaAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *CapabilityServerUpgradeSupportMetaAllOf) GetPlatform() string {
	if o == nil || o.Platform == nil {
		var ret string
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityServerUpgradeSupportMetaAllOf) GetPlatformOk() (*string, bool) {
	if o == nil || o.Platform == nil {
		return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *CapabilityServerUpgradeSupportMetaAllOf) HasPlatform() bool {
	if o != nil && o.Platform != nil {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given string and assigns it to the Platform field.
func (o *CapabilityServerUpgradeSupportMetaAllOf) SetPlatform(v string) {
	o.Platform = &v
}

// GetServerFamily returns the ServerFamily field value if set, zero value otherwise.
func (o *CapabilityServerUpgradeSupportMetaAllOf) GetServerFamily() string {
	if o == nil || o.ServerFamily == nil {
		var ret string
		return ret
	}
	return *o.ServerFamily
}

// GetServerFamilyOk returns a tuple with the ServerFamily field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityServerUpgradeSupportMetaAllOf) GetServerFamilyOk() (*string, bool) {
	if o == nil || o.ServerFamily == nil {
		return nil, false
	}
	return o.ServerFamily, true
}

// HasServerFamily returns a boolean if a field has been set.
func (o *CapabilityServerUpgradeSupportMetaAllOf) HasServerFamily() bool {
	if o != nil && o.ServerFamily != nil {
		return true
	}

	return false
}

// SetServerFamily gets a reference to the given string and assigns it to the ServerFamily field.
func (o *CapabilityServerUpgradeSupportMetaAllOf) SetServerFamily(v string) {
	o.ServerFamily = &v
}

// GetSupportedModels returns the SupportedModels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CapabilityServerUpgradeSupportMetaAllOf) GetSupportedModels() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SupportedModels
}

// GetSupportedModelsOk returns a tuple with the SupportedModels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CapabilityServerUpgradeSupportMetaAllOf) GetSupportedModelsOk() ([]string, bool) {
	if o == nil || o.SupportedModels == nil {
		return nil, false
	}
	return o.SupportedModels, true
}

// HasSupportedModels returns a boolean if a field has been set.
func (o *CapabilityServerUpgradeSupportMetaAllOf) HasSupportedModels() bool {
	if o != nil && o.SupportedModels != nil {
		return true
	}

	return false
}

// SetSupportedModels gets a reference to the given []string and assigns it to the SupportedModels field.
func (o *CapabilityServerUpgradeSupportMetaAllOf) SetSupportedModels(v []string) {
	o.SupportedModels = v
}

func (o CapabilityServerUpgradeSupportMetaAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.Platform != nil {
		toSerialize["Platform"] = o.Platform
	}
	if o.ServerFamily != nil {
		toSerialize["ServerFamily"] = o.ServerFamily
	}
	if o.SupportedModels != nil {
		toSerialize["SupportedModels"] = o.SupportedModels
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CapabilityServerUpgradeSupportMetaAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varCapabilityServerUpgradeSupportMetaAllOf := _CapabilityServerUpgradeSupportMetaAllOf{}

	if err = json.Unmarshal(bytes, &varCapabilityServerUpgradeSupportMetaAllOf); err == nil {
		*o = CapabilityServerUpgradeSupportMetaAllOf(varCapabilityServerUpgradeSupportMetaAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "Platform")
		delete(additionalProperties, "ServerFamily")
		delete(additionalProperties, "SupportedModels")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCapabilityServerUpgradeSupportMetaAllOf struct {
	value *CapabilityServerUpgradeSupportMetaAllOf
	isSet bool
}

func (v NullableCapabilityServerUpgradeSupportMetaAllOf) Get() *CapabilityServerUpgradeSupportMetaAllOf {
	return v.value
}

func (v *NullableCapabilityServerUpgradeSupportMetaAllOf) Set(val *CapabilityServerUpgradeSupportMetaAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilityServerUpgradeSupportMetaAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilityServerUpgradeSupportMetaAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilityServerUpgradeSupportMetaAllOf(val *CapabilityServerUpgradeSupportMetaAllOf) *NullableCapabilityServerUpgradeSupportMetaAllOf {
	return &NullableCapabilityServerUpgradeSupportMetaAllOf{value: val, isSet: true}
}

func (v NullableCapabilityServerUpgradeSupportMetaAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilityServerUpgradeSupportMetaAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
