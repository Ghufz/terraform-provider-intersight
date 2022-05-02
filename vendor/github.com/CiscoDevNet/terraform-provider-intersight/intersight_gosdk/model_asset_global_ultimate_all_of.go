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

// AssetGlobalUltimateAllOf Definition of the list of properties defined in 'asset.GlobalUltimate', excluding properties defined in parent classes.
type AssetGlobalUltimateAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// ID of the user in BillToGlobal.
	Id *string `json:"Id,omitempty"`
	// Name of the user in BillToGlobal.
	Name                 *string `json:"Name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetGlobalUltimateAllOf AssetGlobalUltimateAllOf

// NewAssetGlobalUltimateAllOf instantiates a new AssetGlobalUltimateAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetGlobalUltimateAllOf(classId string, objectType string) *AssetGlobalUltimateAllOf {
	this := AssetGlobalUltimateAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAssetGlobalUltimateAllOfWithDefaults instantiates a new AssetGlobalUltimateAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetGlobalUltimateAllOfWithDefaults() *AssetGlobalUltimateAllOf {
	this := AssetGlobalUltimateAllOf{}
	var classId string = "asset.GlobalUltimate"
	this.ClassId = classId
	var objectType string = "asset.GlobalUltimate"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetGlobalUltimateAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetGlobalUltimateAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetGlobalUltimateAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AssetGlobalUltimateAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetGlobalUltimateAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetGlobalUltimateAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AssetGlobalUltimateAllOf) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetGlobalUltimateAllOf) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AssetGlobalUltimateAllOf) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AssetGlobalUltimateAllOf) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AssetGlobalUltimateAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetGlobalUltimateAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AssetGlobalUltimateAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AssetGlobalUltimateAllOf) SetName(v string) {
	o.Name = &v
}

func (o AssetGlobalUltimateAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Id != nil {
		toSerialize["Id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetGlobalUltimateAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varAssetGlobalUltimateAllOf := _AssetGlobalUltimateAllOf{}

	if err = json.Unmarshal(bytes, &varAssetGlobalUltimateAllOf); err == nil {
		*o = AssetGlobalUltimateAllOf(varAssetGlobalUltimateAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Id")
		delete(additionalProperties, "Name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetGlobalUltimateAllOf struct {
	value *AssetGlobalUltimateAllOf
	isSet bool
}

func (v NullableAssetGlobalUltimateAllOf) Get() *AssetGlobalUltimateAllOf {
	return v.value
}

func (v *NullableAssetGlobalUltimateAllOf) Set(val *AssetGlobalUltimateAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetGlobalUltimateAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetGlobalUltimateAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetGlobalUltimateAllOf(val *AssetGlobalUltimateAllOf) *NullableAssetGlobalUltimateAllOf {
	return &NullableAssetGlobalUltimateAllOf{value: val, isSet: true}
}

func (v NullableAssetGlobalUltimateAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetGlobalUltimateAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
