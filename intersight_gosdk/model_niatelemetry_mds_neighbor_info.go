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

// NiatelemetryMdsNeighborInfo Object that carries wwn of Neighbours needed for Neighbor discovery for Autoclaim.
type NiatelemetryMdsNeighborInfo struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Returns deviceWwn info of neigbbor.
	DeviceWwn            *string `json:"DeviceWwn,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryMdsNeighborInfo NiatelemetryMdsNeighborInfo

// NewNiatelemetryMdsNeighborInfo instantiates a new NiatelemetryMdsNeighborInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryMdsNeighborInfo(classId string, objectType string) *NiatelemetryMdsNeighborInfo {
	this := NiatelemetryMdsNeighborInfo{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryMdsNeighborInfoWithDefaults instantiates a new NiatelemetryMdsNeighborInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryMdsNeighborInfoWithDefaults() *NiatelemetryMdsNeighborInfo {
	this := NiatelemetryMdsNeighborInfo{}
	var classId string = "niatelemetry.MdsNeighborInfo"
	this.ClassId = classId
	var objectType string = "niatelemetry.MdsNeighborInfo"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryMdsNeighborInfo) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryMdsNeighborInfo) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryMdsNeighborInfo) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryMdsNeighborInfo) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryMdsNeighborInfo) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryMdsNeighborInfo) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDeviceWwn returns the DeviceWwn field value if set, zero value otherwise.
func (o *NiatelemetryMdsNeighborInfo) GetDeviceWwn() string {
	if o == nil || o.DeviceWwn == nil {
		var ret string
		return ret
	}
	return *o.DeviceWwn
}

// GetDeviceWwnOk returns a tuple with the DeviceWwn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryMdsNeighborInfo) GetDeviceWwnOk() (*string, bool) {
	if o == nil || o.DeviceWwn == nil {
		return nil, false
	}
	return o.DeviceWwn, true
}

// HasDeviceWwn returns a boolean if a field has been set.
func (o *NiatelemetryMdsNeighborInfo) HasDeviceWwn() bool {
	if o != nil && o.DeviceWwn != nil {
		return true
	}

	return false
}

// SetDeviceWwn gets a reference to the given string and assigns it to the DeviceWwn field.
func (o *NiatelemetryMdsNeighborInfo) SetDeviceWwn(v string) {
	o.DeviceWwn = &v
}

func (o NiatelemetryMdsNeighborInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DeviceWwn != nil {
		toSerialize["DeviceWwn"] = o.DeviceWwn
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryMdsNeighborInfo) UnmarshalJSON(bytes []byte) (err error) {
	type NiatelemetryMdsNeighborInfoWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Returns deviceWwn info of neigbbor.
		DeviceWwn *string `json:"DeviceWwn,omitempty"`
	}

	varNiatelemetryMdsNeighborInfoWithoutEmbeddedStruct := NiatelemetryMdsNeighborInfoWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNiatelemetryMdsNeighborInfoWithoutEmbeddedStruct)
	if err == nil {
		varNiatelemetryMdsNeighborInfo := _NiatelemetryMdsNeighborInfo{}
		varNiatelemetryMdsNeighborInfo.ClassId = varNiatelemetryMdsNeighborInfoWithoutEmbeddedStruct.ClassId
		varNiatelemetryMdsNeighborInfo.ObjectType = varNiatelemetryMdsNeighborInfoWithoutEmbeddedStruct.ObjectType
		varNiatelemetryMdsNeighborInfo.DeviceWwn = varNiatelemetryMdsNeighborInfoWithoutEmbeddedStruct.DeviceWwn
		*o = NiatelemetryMdsNeighborInfo(varNiatelemetryMdsNeighborInfo)
	} else {
		return err
	}

	varNiatelemetryMdsNeighborInfo := _NiatelemetryMdsNeighborInfo{}

	err = json.Unmarshal(bytes, &varNiatelemetryMdsNeighborInfo)
	if err == nil {
		o.MoBaseComplexType = varNiatelemetryMdsNeighborInfo.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DeviceWwn")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

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

type NullableNiatelemetryMdsNeighborInfo struct {
	value *NiatelemetryMdsNeighborInfo
	isSet bool
}

func (v NullableNiatelemetryMdsNeighborInfo) Get() *NiatelemetryMdsNeighborInfo {
	return v.value
}

func (v *NullableNiatelemetryMdsNeighborInfo) Set(val *NiatelemetryMdsNeighborInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryMdsNeighborInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryMdsNeighborInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryMdsNeighborInfo(val *NiatelemetryMdsNeighborInfo) *NullableNiatelemetryMdsNeighborInfo {
	return &NullableNiatelemetryMdsNeighborInfo{value: val, isSet: true}
}

func (v NullableNiatelemetryMdsNeighborInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryMdsNeighborInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
