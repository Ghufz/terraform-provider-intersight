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

// ComputeServerConfigAllOf Definition of the list of properties defined in 'compute.ServerConfig', excluding properties defined in parent classes.
type ComputeServerConfigAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// User defined asset tag of the server.
	AssetTag *string `json:"AssetTag,omitempty"`
	// User defined description of the server.
	UserLabel            *string `json:"UserLabel,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ComputeServerConfigAllOf ComputeServerConfigAllOf

// NewComputeServerConfigAllOf instantiates a new ComputeServerConfigAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputeServerConfigAllOf(classId string, objectType string) *ComputeServerConfigAllOf {
	this := ComputeServerConfigAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewComputeServerConfigAllOfWithDefaults instantiates a new ComputeServerConfigAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputeServerConfigAllOfWithDefaults() *ComputeServerConfigAllOf {
	this := ComputeServerConfigAllOf{}
	var classId string = "compute.ServerConfig"
	this.ClassId = classId
	var objectType string = "compute.ServerConfig"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ComputeServerConfigAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ComputeServerConfigAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ComputeServerConfigAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ComputeServerConfigAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ComputeServerConfigAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ComputeServerConfigAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAssetTag returns the AssetTag field value if set, zero value otherwise.
func (o *ComputeServerConfigAllOf) GetAssetTag() string {
	if o == nil || o.AssetTag == nil {
		var ret string
		return ret
	}
	return *o.AssetTag
}

// GetAssetTagOk returns a tuple with the AssetTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeServerConfigAllOf) GetAssetTagOk() (*string, bool) {
	if o == nil || o.AssetTag == nil {
		return nil, false
	}
	return o.AssetTag, true
}

// HasAssetTag returns a boolean if a field has been set.
func (o *ComputeServerConfigAllOf) HasAssetTag() bool {
	if o != nil && o.AssetTag != nil {
		return true
	}

	return false
}

// SetAssetTag gets a reference to the given string and assigns it to the AssetTag field.
func (o *ComputeServerConfigAllOf) SetAssetTag(v string) {
	o.AssetTag = &v
}

// GetUserLabel returns the UserLabel field value if set, zero value otherwise.
func (o *ComputeServerConfigAllOf) GetUserLabel() string {
	if o == nil || o.UserLabel == nil {
		var ret string
		return ret
	}
	return *o.UserLabel
}

// GetUserLabelOk returns a tuple with the UserLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeServerConfigAllOf) GetUserLabelOk() (*string, bool) {
	if o == nil || o.UserLabel == nil {
		return nil, false
	}
	return o.UserLabel, true
}

// HasUserLabel returns a boolean if a field has been set.
func (o *ComputeServerConfigAllOf) HasUserLabel() bool {
	if o != nil && o.UserLabel != nil {
		return true
	}

	return false
}

// SetUserLabel gets a reference to the given string and assigns it to the UserLabel field.
func (o *ComputeServerConfigAllOf) SetUserLabel(v string) {
	o.UserLabel = &v
}

func (o ComputeServerConfigAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AssetTag != nil {
		toSerialize["AssetTag"] = o.AssetTag
	}
	if o.UserLabel != nil {
		toSerialize["UserLabel"] = o.UserLabel
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ComputeServerConfigAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varComputeServerConfigAllOf := _ComputeServerConfigAllOf{}

	if err = json.Unmarshal(bytes, &varComputeServerConfigAllOf); err == nil {
		*o = ComputeServerConfigAllOf(varComputeServerConfigAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AssetTag")
		delete(additionalProperties, "UserLabel")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableComputeServerConfigAllOf struct {
	value *ComputeServerConfigAllOf
	isSet bool
}

func (v NullableComputeServerConfigAllOf) Get() *ComputeServerConfigAllOf {
	return v.value
}

func (v *NullableComputeServerConfigAllOf) Set(val *ComputeServerConfigAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableComputeServerConfigAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableComputeServerConfigAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputeServerConfigAllOf(val *ComputeServerConfigAllOf) *NullableComputeServerConfigAllOf {
	return &NullableComputeServerConfigAllOf{value: val, isSet: true}
}

func (v NullableComputeServerConfigAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputeServerConfigAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
