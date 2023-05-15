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
	"reflect"
	"strings"
)

// HclHwCatalogInfo Hardware catalog information for HyperFlex Data Platform, required for BOM validation.
type HclHwCatalogInfo struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Server model information for HyperFlex servers.
	ServerModel *string `json:"ServerModel,omitempty"`
	// Server type of the server hardware. For example, server type AF is for an all-flash server.
	ServerType *string `json:"ServerType,omitempty"`
	// An array of relationships to hclServerHwCatalogInfo resources.
	ServerHwCatalogInfo  []HclServerHwCatalogInfoRelationship `json:"ServerHwCatalogInfo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HclHwCatalogInfo HclHwCatalogInfo

// NewHclHwCatalogInfo instantiates a new HclHwCatalogInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHclHwCatalogInfo(classId string, objectType string) *HclHwCatalogInfo {
	this := HclHwCatalogInfo{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHclHwCatalogInfoWithDefaults instantiates a new HclHwCatalogInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHclHwCatalogInfoWithDefaults() *HclHwCatalogInfo {
	this := HclHwCatalogInfo{}
	var classId string = "hcl.HwCatalogInfo"
	this.ClassId = classId
	var objectType string = "hcl.HwCatalogInfo"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HclHwCatalogInfo) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HclHwCatalogInfo) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HclHwCatalogInfo) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HclHwCatalogInfo) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HclHwCatalogInfo) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HclHwCatalogInfo) SetObjectType(v string) {
	o.ObjectType = v
}

// GetServerModel returns the ServerModel field value if set, zero value otherwise.
func (o *HclHwCatalogInfo) GetServerModel() string {
	if o == nil || o.ServerModel == nil {
		var ret string
		return ret
	}
	return *o.ServerModel
}

// GetServerModelOk returns a tuple with the ServerModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclHwCatalogInfo) GetServerModelOk() (*string, bool) {
	if o == nil || o.ServerModel == nil {
		return nil, false
	}
	return o.ServerModel, true
}

// HasServerModel returns a boolean if a field has been set.
func (o *HclHwCatalogInfo) HasServerModel() bool {
	if o != nil && o.ServerModel != nil {
		return true
	}

	return false
}

// SetServerModel gets a reference to the given string and assigns it to the ServerModel field.
func (o *HclHwCatalogInfo) SetServerModel(v string) {
	o.ServerModel = &v
}

// GetServerType returns the ServerType field value if set, zero value otherwise.
func (o *HclHwCatalogInfo) GetServerType() string {
	if o == nil || o.ServerType == nil {
		var ret string
		return ret
	}
	return *o.ServerType
}

// GetServerTypeOk returns a tuple with the ServerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclHwCatalogInfo) GetServerTypeOk() (*string, bool) {
	if o == nil || o.ServerType == nil {
		return nil, false
	}
	return o.ServerType, true
}

// HasServerType returns a boolean if a field has been set.
func (o *HclHwCatalogInfo) HasServerType() bool {
	if o != nil && o.ServerType != nil {
		return true
	}

	return false
}

// SetServerType gets a reference to the given string and assigns it to the ServerType field.
func (o *HclHwCatalogInfo) SetServerType(v string) {
	o.ServerType = &v
}

// GetServerHwCatalogInfo returns the ServerHwCatalogInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HclHwCatalogInfo) GetServerHwCatalogInfo() []HclServerHwCatalogInfoRelationship {
	if o == nil {
		var ret []HclServerHwCatalogInfoRelationship
		return ret
	}
	return o.ServerHwCatalogInfo
}

// GetServerHwCatalogInfoOk returns a tuple with the ServerHwCatalogInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HclHwCatalogInfo) GetServerHwCatalogInfoOk() ([]HclServerHwCatalogInfoRelationship, bool) {
	if o == nil || o.ServerHwCatalogInfo == nil {
		return nil, false
	}
	return o.ServerHwCatalogInfo, true
}

// HasServerHwCatalogInfo returns a boolean if a field has been set.
func (o *HclHwCatalogInfo) HasServerHwCatalogInfo() bool {
	if o != nil && o.ServerHwCatalogInfo != nil {
		return true
	}

	return false
}

// SetServerHwCatalogInfo gets a reference to the given []HclServerHwCatalogInfoRelationship and assigns it to the ServerHwCatalogInfo field.
func (o *HclHwCatalogInfo) SetServerHwCatalogInfo(v []HclServerHwCatalogInfoRelationship) {
	o.ServerHwCatalogInfo = v
}

func (o HclHwCatalogInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ServerModel != nil {
		toSerialize["ServerModel"] = o.ServerModel
	}
	if o.ServerType != nil {
		toSerialize["ServerType"] = o.ServerType
	}
	if o.ServerHwCatalogInfo != nil {
		toSerialize["ServerHwCatalogInfo"] = o.ServerHwCatalogInfo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HclHwCatalogInfo) UnmarshalJSON(bytes []byte) (err error) {
	type HclHwCatalogInfoWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Server model information for HyperFlex servers.
		ServerModel *string `json:"ServerModel,omitempty"`
		// Server type of the server hardware. For example, server type AF is for an all-flash server.
		ServerType *string `json:"ServerType,omitempty"`
		// An array of relationships to hclServerHwCatalogInfo resources.
		ServerHwCatalogInfo []HclServerHwCatalogInfoRelationship `json:"ServerHwCatalogInfo,omitempty"`
	}

	varHclHwCatalogInfoWithoutEmbeddedStruct := HclHwCatalogInfoWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHclHwCatalogInfoWithoutEmbeddedStruct)
	if err == nil {
		varHclHwCatalogInfo := _HclHwCatalogInfo{}
		varHclHwCatalogInfo.ClassId = varHclHwCatalogInfoWithoutEmbeddedStruct.ClassId
		varHclHwCatalogInfo.ObjectType = varHclHwCatalogInfoWithoutEmbeddedStruct.ObjectType
		varHclHwCatalogInfo.ServerModel = varHclHwCatalogInfoWithoutEmbeddedStruct.ServerModel
		varHclHwCatalogInfo.ServerType = varHclHwCatalogInfoWithoutEmbeddedStruct.ServerType
		varHclHwCatalogInfo.ServerHwCatalogInfo = varHclHwCatalogInfoWithoutEmbeddedStruct.ServerHwCatalogInfo
		*o = HclHwCatalogInfo(varHclHwCatalogInfo)
	} else {
		return err
	}

	varHclHwCatalogInfo := _HclHwCatalogInfo{}

	err = json.Unmarshal(bytes, &varHclHwCatalogInfo)
	if err == nil {
		o.MoBaseMo = varHclHwCatalogInfo.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ServerModel")
		delete(additionalProperties, "ServerType")
		delete(additionalProperties, "ServerHwCatalogInfo")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

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

type NullableHclHwCatalogInfo struct {
	value *HclHwCatalogInfo
	isSet bool
}

func (v NullableHclHwCatalogInfo) Get() *HclHwCatalogInfo {
	return v.value
}

func (v *NullableHclHwCatalogInfo) Set(val *HclHwCatalogInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableHclHwCatalogInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableHclHwCatalogInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHclHwCatalogInfo(val *HclHwCatalogInfo) *NullableHclHwCatalogInfo {
	return &NullableHclHwCatalogInfo{value: val, isSet: true}
}

func (v NullableHclHwCatalogInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHclHwCatalogInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
