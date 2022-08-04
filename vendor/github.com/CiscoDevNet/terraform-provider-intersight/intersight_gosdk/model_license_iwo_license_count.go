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
	"reflect"
	"strings"
)

// LicenseIwoLicenseCount Customer operation object to request reservation code.
type LicenseIwoLicenseCount struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The total number of devices claimed in the Intersight account.
	VmLicenseCount       *int64                                 `json:"VmLicenseCount,omitempty"`
	AccountLicenseData   *LicenseAccountLicenseDataRelationship `json:"AccountLicenseData,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LicenseIwoLicenseCount LicenseIwoLicenseCount

// NewLicenseIwoLicenseCount instantiates a new LicenseIwoLicenseCount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicenseIwoLicenseCount(classId string, objectType string) *LicenseIwoLicenseCount {
	this := LicenseIwoLicenseCount{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewLicenseIwoLicenseCountWithDefaults instantiates a new LicenseIwoLicenseCount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicenseIwoLicenseCountWithDefaults() *LicenseIwoLicenseCount {
	this := LicenseIwoLicenseCount{}
	var classId string = "license.IwoLicenseCount"
	this.ClassId = classId
	var objectType string = "license.IwoLicenseCount"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *LicenseIwoLicenseCount) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *LicenseIwoLicenseCount) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *LicenseIwoLicenseCount) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *LicenseIwoLicenseCount) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *LicenseIwoLicenseCount) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *LicenseIwoLicenseCount) SetObjectType(v string) {
	o.ObjectType = v
}

// GetVmLicenseCount returns the VmLicenseCount field value if set, zero value otherwise.
func (o *LicenseIwoLicenseCount) GetVmLicenseCount() int64 {
	if o == nil || o.VmLicenseCount == nil {
		var ret int64
		return ret
	}
	return *o.VmLicenseCount
}

// GetVmLicenseCountOk returns a tuple with the VmLicenseCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseIwoLicenseCount) GetVmLicenseCountOk() (*int64, bool) {
	if o == nil || o.VmLicenseCount == nil {
		return nil, false
	}
	return o.VmLicenseCount, true
}

// HasVmLicenseCount returns a boolean if a field has been set.
func (o *LicenseIwoLicenseCount) HasVmLicenseCount() bool {
	if o != nil && o.VmLicenseCount != nil {
		return true
	}

	return false
}

// SetVmLicenseCount gets a reference to the given int64 and assigns it to the VmLicenseCount field.
func (o *LicenseIwoLicenseCount) SetVmLicenseCount(v int64) {
	o.VmLicenseCount = &v
}

// GetAccountLicenseData returns the AccountLicenseData field value if set, zero value otherwise.
func (o *LicenseIwoLicenseCount) GetAccountLicenseData() LicenseAccountLicenseDataRelationship {
	if o == nil || o.AccountLicenseData == nil {
		var ret LicenseAccountLicenseDataRelationship
		return ret
	}
	return *o.AccountLicenseData
}

// GetAccountLicenseDataOk returns a tuple with the AccountLicenseData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseIwoLicenseCount) GetAccountLicenseDataOk() (*LicenseAccountLicenseDataRelationship, bool) {
	if o == nil || o.AccountLicenseData == nil {
		return nil, false
	}
	return o.AccountLicenseData, true
}

// HasAccountLicenseData returns a boolean if a field has been set.
func (o *LicenseIwoLicenseCount) HasAccountLicenseData() bool {
	if o != nil && o.AccountLicenseData != nil {
		return true
	}

	return false
}

// SetAccountLicenseData gets a reference to the given LicenseAccountLicenseDataRelationship and assigns it to the AccountLicenseData field.
func (o *LicenseIwoLicenseCount) SetAccountLicenseData(v LicenseAccountLicenseDataRelationship) {
	o.AccountLicenseData = &v
}

func (o LicenseIwoLicenseCount) MarshalJSON() ([]byte, error) {
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
	if o.VmLicenseCount != nil {
		toSerialize["VmLicenseCount"] = o.VmLicenseCount
	}
	if o.AccountLicenseData != nil {
		toSerialize["AccountLicenseData"] = o.AccountLicenseData
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LicenseIwoLicenseCount) UnmarshalJSON(bytes []byte) (err error) {
	type LicenseIwoLicenseCountWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The total number of devices claimed in the Intersight account.
		VmLicenseCount     *int64                                 `json:"VmLicenseCount,omitempty"`
		AccountLicenseData *LicenseAccountLicenseDataRelationship `json:"AccountLicenseData,omitempty"`
	}

	varLicenseIwoLicenseCountWithoutEmbeddedStruct := LicenseIwoLicenseCountWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varLicenseIwoLicenseCountWithoutEmbeddedStruct)
	if err == nil {
		varLicenseIwoLicenseCount := _LicenseIwoLicenseCount{}
		varLicenseIwoLicenseCount.ClassId = varLicenseIwoLicenseCountWithoutEmbeddedStruct.ClassId
		varLicenseIwoLicenseCount.ObjectType = varLicenseIwoLicenseCountWithoutEmbeddedStruct.ObjectType
		varLicenseIwoLicenseCount.VmLicenseCount = varLicenseIwoLicenseCountWithoutEmbeddedStruct.VmLicenseCount
		varLicenseIwoLicenseCount.AccountLicenseData = varLicenseIwoLicenseCountWithoutEmbeddedStruct.AccountLicenseData
		*o = LicenseIwoLicenseCount(varLicenseIwoLicenseCount)
	} else {
		return err
	}

	varLicenseIwoLicenseCount := _LicenseIwoLicenseCount{}

	err = json.Unmarshal(bytes, &varLicenseIwoLicenseCount)
	if err == nil {
		o.MoBaseMo = varLicenseIwoLicenseCount.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "VmLicenseCount")
		delete(additionalProperties, "AccountLicenseData")

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

type NullableLicenseIwoLicenseCount struct {
	value *LicenseIwoLicenseCount
	isSet bool
}

func (v NullableLicenseIwoLicenseCount) Get() *LicenseIwoLicenseCount {
	return v.value
}

func (v *NullableLicenseIwoLicenseCount) Set(val *LicenseIwoLicenseCount) {
	v.value = val
	v.isSet = true
}

func (v NullableLicenseIwoLicenseCount) IsSet() bool {
	return v.isSet
}

func (v *NullableLicenseIwoLicenseCount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicenseIwoLicenseCount(val *LicenseIwoLicenseCount) *NullableLicenseIwoLicenseCount {
	return &NullableLicenseIwoLicenseCount{value: val, isSet: true}
}

func (v NullableLicenseIwoLicenseCount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicenseIwoLicenseCount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
