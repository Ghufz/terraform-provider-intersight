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

// AssetScopedTargetConnectionAllOf Definition of the list of properties defined in 'asset.ScopedTargetConnection', excluding properties defined in parent classes.
type AssetScopedTargetConnectionAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// When this flag is set to true, every IWO entity in the scope targets will be checked and discovery of the scope target will be regarded as a failure when anyone of these entities cannot be connected and validated.
	FullValidation *bool `json:"FullValidation,omitempty"`
	// The port number to be used to connect to the managed target. Values 1-65535 indicate a port number to be used. A value of 0 is not a valid port number and instead indicates that the default management port, as defined by the documentation of the managed target, should be used to establish a connection.
	Port *int64 `json:"Port,omitempty"`
	// The group id of IWO entities upon which the discover of a scoped target is performed. For example, a database target may be scoped to the group of virtual machines upon which the database application is running. Scope value is group id created for all those virtual machines in this scope.
	Scope                *string `json:"Scope,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetScopedTargetConnectionAllOf AssetScopedTargetConnectionAllOf

// NewAssetScopedTargetConnectionAllOf instantiates a new AssetScopedTargetConnectionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetScopedTargetConnectionAllOf(classId string, objectType string) *AssetScopedTargetConnectionAllOf {
	this := AssetScopedTargetConnectionAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAssetScopedTargetConnectionAllOfWithDefaults instantiates a new AssetScopedTargetConnectionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetScopedTargetConnectionAllOfWithDefaults() *AssetScopedTargetConnectionAllOf {
	this := AssetScopedTargetConnectionAllOf{}
	var classId string = "asset.ScopedTargetConnection"
	this.ClassId = classId
	var objectType string = "asset.ScopedTargetConnection"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetScopedTargetConnectionAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetScopedTargetConnectionAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetScopedTargetConnectionAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AssetScopedTargetConnectionAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetScopedTargetConnectionAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetScopedTargetConnectionAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetFullValidation returns the FullValidation field value if set, zero value otherwise.
func (o *AssetScopedTargetConnectionAllOf) GetFullValidation() bool {
	if o == nil || o.FullValidation == nil {
		var ret bool
		return ret
	}
	return *o.FullValidation
}

// GetFullValidationOk returns a tuple with the FullValidation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetScopedTargetConnectionAllOf) GetFullValidationOk() (*bool, bool) {
	if o == nil || o.FullValidation == nil {
		return nil, false
	}
	return o.FullValidation, true
}

// HasFullValidation returns a boolean if a field has been set.
func (o *AssetScopedTargetConnectionAllOf) HasFullValidation() bool {
	if o != nil && o.FullValidation != nil {
		return true
	}

	return false
}

// SetFullValidation gets a reference to the given bool and assigns it to the FullValidation field.
func (o *AssetScopedTargetConnectionAllOf) SetFullValidation(v bool) {
	o.FullValidation = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *AssetScopedTargetConnectionAllOf) GetPort() int64 {
	if o == nil || o.Port == nil {
		var ret int64
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetScopedTargetConnectionAllOf) GetPortOk() (*int64, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *AssetScopedTargetConnectionAllOf) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int64 and assigns it to the Port field.
func (o *AssetScopedTargetConnectionAllOf) SetPort(v int64) {
	o.Port = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *AssetScopedTargetConnectionAllOf) GetScope() string {
	if o == nil || o.Scope == nil {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetScopedTargetConnectionAllOf) GetScopeOk() (*string, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *AssetScopedTargetConnectionAllOf) HasScope() bool {
	if o != nil && o.Scope != nil {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *AssetScopedTargetConnectionAllOf) SetScope(v string) {
	o.Scope = &v
}

func (o AssetScopedTargetConnectionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.FullValidation != nil {
		toSerialize["FullValidation"] = o.FullValidation
	}
	if o.Port != nil {
		toSerialize["Port"] = o.Port
	}
	if o.Scope != nil {
		toSerialize["Scope"] = o.Scope
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetScopedTargetConnectionAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varAssetScopedTargetConnectionAllOf := _AssetScopedTargetConnectionAllOf{}

	if err = json.Unmarshal(bytes, &varAssetScopedTargetConnectionAllOf); err == nil {
		*o = AssetScopedTargetConnectionAllOf(varAssetScopedTargetConnectionAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "FullValidation")
		delete(additionalProperties, "Port")
		delete(additionalProperties, "Scope")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetScopedTargetConnectionAllOf struct {
	value *AssetScopedTargetConnectionAllOf
	isSet bool
}

func (v NullableAssetScopedTargetConnectionAllOf) Get() *AssetScopedTargetConnectionAllOf {
	return v.value
}

func (v *NullableAssetScopedTargetConnectionAllOf) Set(val *AssetScopedTargetConnectionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetScopedTargetConnectionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetScopedTargetConnectionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetScopedTargetConnectionAllOf(val *AssetScopedTargetConnectionAllOf) *NullableAssetScopedTargetConnectionAllOf {
	return &NullableAssetScopedTargetConnectionAllOf{value: val, isSet: true}
}

func (v NullableAssetScopedTargetConnectionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetScopedTargetConnectionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
