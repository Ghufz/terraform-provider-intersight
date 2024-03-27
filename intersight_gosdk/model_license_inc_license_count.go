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
	"reflect"
	"strings"
)

// LicenseIncLicenseCount Customer operation object to request reservation code.
type LicenseIncLicenseCount struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The total number of devices claimed in the premier 100G fixed tier Intersight Nexus Cloud.
	Premier100GfxCount *int64 `json:"Premier100GfxCount,omitempty"`
	// The total number of devices claimed in the premier 10G fixed tier Intersight Nexus Cloud.
	Premier10GfxCount *int64 `json:"Premier10GfxCount,omitempty"`
	// The total number of devices claimed in the premier 1G fixed tier Intersight Nexus Cloud.
	Premier1GfxCount *int64 `json:"Premier1GfxCount,omitempty"`
	// The total number of devices claimed in the CentralizedMod8Slot premier tier Intersight Nexus Cloud.
	PremierCentralizedMod8SlotCount *int64 `json:"PremierCentralizedMod8SlotCount,omitempty"`
	// The total number of devices claimed in the D2Ops Fixed premier tier Intersight Nexus Cloud.
	PremierD2OpsFixedCount *int64 `json:"PremierD2OpsFixedCount,omitempty"`
	// The total number of devices claimed in the D2Ops modular premier tier Intersight Nexus Cloud.
	PremierD2OpsModCount *int64 `json:"PremierD2OpsModCount,omitempty"`
	// The total number of devices claimed in the DistributedMod8Slot premier tier Intersight Nexus Cloud.
	PremierDistributedMod8SlotCount *int64 `json:"PremierDistributedMod8SlotCount,omitempty"`
	// The total number of devices claimed in the modular 4 slot premier tier Intersight Nexus Cloud.
	PremierMod4SlotCount *int64 `json:"PremierMod4SlotCount,omitempty"`
	// The total number of devices claimed in the modular 8 slot premier tier Intersight Nexus Cloud.
	PremierMod8SlotCount *int64                                 `json:"PremierMod8SlotCount,omitempty"`
	AccountLicenseData   *LicenseAccountLicenseDataRelationship `json:"AccountLicenseData,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LicenseIncLicenseCount LicenseIncLicenseCount

// NewLicenseIncLicenseCount instantiates a new LicenseIncLicenseCount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicenseIncLicenseCount(classId string, objectType string) *LicenseIncLicenseCount {
	this := LicenseIncLicenseCount{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewLicenseIncLicenseCountWithDefaults instantiates a new LicenseIncLicenseCount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicenseIncLicenseCountWithDefaults() *LicenseIncLicenseCount {
	this := LicenseIncLicenseCount{}
	var classId string = "license.IncLicenseCount"
	this.ClassId = classId
	var objectType string = "license.IncLicenseCount"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *LicenseIncLicenseCount) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *LicenseIncLicenseCount) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *LicenseIncLicenseCount) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *LicenseIncLicenseCount) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *LicenseIncLicenseCount) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *LicenseIncLicenseCount) SetObjectType(v string) {
	o.ObjectType = v
}

// GetPremier100GfxCount returns the Premier100GfxCount field value if set, zero value otherwise.
func (o *LicenseIncLicenseCount) GetPremier100GfxCount() int64 {
	if o == nil || o.Premier100GfxCount == nil {
		var ret int64
		return ret
	}
	return *o.Premier100GfxCount
}

// GetPremier100GfxCountOk returns a tuple with the Premier100GfxCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseIncLicenseCount) GetPremier100GfxCountOk() (*int64, bool) {
	if o == nil || o.Premier100GfxCount == nil {
		return nil, false
	}
	return o.Premier100GfxCount, true
}

// HasPremier100GfxCount returns a boolean if a field has been set.
func (o *LicenseIncLicenseCount) HasPremier100GfxCount() bool {
	if o != nil && o.Premier100GfxCount != nil {
		return true
	}

	return false
}

// SetPremier100GfxCount gets a reference to the given int64 and assigns it to the Premier100GfxCount field.
func (o *LicenseIncLicenseCount) SetPremier100GfxCount(v int64) {
	o.Premier100GfxCount = &v
}

// GetPremier10GfxCount returns the Premier10GfxCount field value if set, zero value otherwise.
func (o *LicenseIncLicenseCount) GetPremier10GfxCount() int64 {
	if o == nil || o.Premier10GfxCount == nil {
		var ret int64
		return ret
	}
	return *o.Premier10GfxCount
}

// GetPremier10GfxCountOk returns a tuple with the Premier10GfxCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseIncLicenseCount) GetPremier10GfxCountOk() (*int64, bool) {
	if o == nil || o.Premier10GfxCount == nil {
		return nil, false
	}
	return o.Premier10GfxCount, true
}

// HasPremier10GfxCount returns a boolean if a field has been set.
func (o *LicenseIncLicenseCount) HasPremier10GfxCount() bool {
	if o != nil && o.Premier10GfxCount != nil {
		return true
	}

	return false
}

// SetPremier10GfxCount gets a reference to the given int64 and assigns it to the Premier10GfxCount field.
func (o *LicenseIncLicenseCount) SetPremier10GfxCount(v int64) {
	o.Premier10GfxCount = &v
}

// GetPremier1GfxCount returns the Premier1GfxCount field value if set, zero value otherwise.
func (o *LicenseIncLicenseCount) GetPremier1GfxCount() int64 {
	if o == nil || o.Premier1GfxCount == nil {
		var ret int64
		return ret
	}
	return *o.Premier1GfxCount
}

// GetPremier1GfxCountOk returns a tuple with the Premier1GfxCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseIncLicenseCount) GetPremier1GfxCountOk() (*int64, bool) {
	if o == nil || o.Premier1GfxCount == nil {
		return nil, false
	}
	return o.Premier1GfxCount, true
}

// HasPremier1GfxCount returns a boolean if a field has been set.
func (o *LicenseIncLicenseCount) HasPremier1GfxCount() bool {
	if o != nil && o.Premier1GfxCount != nil {
		return true
	}

	return false
}

// SetPremier1GfxCount gets a reference to the given int64 and assigns it to the Premier1GfxCount field.
func (o *LicenseIncLicenseCount) SetPremier1GfxCount(v int64) {
	o.Premier1GfxCount = &v
}

// GetPremierCentralizedMod8SlotCount returns the PremierCentralizedMod8SlotCount field value if set, zero value otherwise.
func (o *LicenseIncLicenseCount) GetPremierCentralizedMod8SlotCount() int64 {
	if o == nil || o.PremierCentralizedMod8SlotCount == nil {
		var ret int64
		return ret
	}
	return *o.PremierCentralizedMod8SlotCount
}

// GetPremierCentralizedMod8SlotCountOk returns a tuple with the PremierCentralizedMod8SlotCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseIncLicenseCount) GetPremierCentralizedMod8SlotCountOk() (*int64, bool) {
	if o == nil || o.PremierCentralizedMod8SlotCount == nil {
		return nil, false
	}
	return o.PremierCentralizedMod8SlotCount, true
}

// HasPremierCentralizedMod8SlotCount returns a boolean if a field has been set.
func (o *LicenseIncLicenseCount) HasPremierCentralizedMod8SlotCount() bool {
	if o != nil && o.PremierCentralizedMod8SlotCount != nil {
		return true
	}

	return false
}

// SetPremierCentralizedMod8SlotCount gets a reference to the given int64 and assigns it to the PremierCentralizedMod8SlotCount field.
func (o *LicenseIncLicenseCount) SetPremierCentralizedMod8SlotCount(v int64) {
	o.PremierCentralizedMod8SlotCount = &v
}

// GetPremierD2OpsFixedCount returns the PremierD2OpsFixedCount field value if set, zero value otherwise.
func (o *LicenseIncLicenseCount) GetPremierD2OpsFixedCount() int64 {
	if o == nil || o.PremierD2OpsFixedCount == nil {
		var ret int64
		return ret
	}
	return *o.PremierD2OpsFixedCount
}

// GetPremierD2OpsFixedCountOk returns a tuple with the PremierD2OpsFixedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseIncLicenseCount) GetPremierD2OpsFixedCountOk() (*int64, bool) {
	if o == nil || o.PremierD2OpsFixedCount == nil {
		return nil, false
	}
	return o.PremierD2OpsFixedCount, true
}

// HasPremierD2OpsFixedCount returns a boolean if a field has been set.
func (o *LicenseIncLicenseCount) HasPremierD2OpsFixedCount() bool {
	if o != nil && o.PremierD2OpsFixedCount != nil {
		return true
	}

	return false
}

// SetPremierD2OpsFixedCount gets a reference to the given int64 and assigns it to the PremierD2OpsFixedCount field.
func (o *LicenseIncLicenseCount) SetPremierD2OpsFixedCount(v int64) {
	o.PremierD2OpsFixedCount = &v
}

// GetPremierD2OpsModCount returns the PremierD2OpsModCount field value if set, zero value otherwise.
func (o *LicenseIncLicenseCount) GetPremierD2OpsModCount() int64 {
	if o == nil || o.PremierD2OpsModCount == nil {
		var ret int64
		return ret
	}
	return *o.PremierD2OpsModCount
}

// GetPremierD2OpsModCountOk returns a tuple with the PremierD2OpsModCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseIncLicenseCount) GetPremierD2OpsModCountOk() (*int64, bool) {
	if o == nil || o.PremierD2OpsModCount == nil {
		return nil, false
	}
	return o.PremierD2OpsModCount, true
}

// HasPremierD2OpsModCount returns a boolean if a field has been set.
func (o *LicenseIncLicenseCount) HasPremierD2OpsModCount() bool {
	if o != nil && o.PremierD2OpsModCount != nil {
		return true
	}

	return false
}

// SetPremierD2OpsModCount gets a reference to the given int64 and assigns it to the PremierD2OpsModCount field.
func (o *LicenseIncLicenseCount) SetPremierD2OpsModCount(v int64) {
	o.PremierD2OpsModCount = &v
}

// GetPremierDistributedMod8SlotCount returns the PremierDistributedMod8SlotCount field value if set, zero value otherwise.
func (o *LicenseIncLicenseCount) GetPremierDistributedMod8SlotCount() int64 {
	if o == nil || o.PremierDistributedMod8SlotCount == nil {
		var ret int64
		return ret
	}
	return *o.PremierDistributedMod8SlotCount
}

// GetPremierDistributedMod8SlotCountOk returns a tuple with the PremierDistributedMod8SlotCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseIncLicenseCount) GetPremierDistributedMod8SlotCountOk() (*int64, bool) {
	if o == nil || o.PremierDistributedMod8SlotCount == nil {
		return nil, false
	}
	return o.PremierDistributedMod8SlotCount, true
}

// HasPremierDistributedMod8SlotCount returns a boolean if a field has been set.
func (o *LicenseIncLicenseCount) HasPremierDistributedMod8SlotCount() bool {
	if o != nil && o.PremierDistributedMod8SlotCount != nil {
		return true
	}

	return false
}

// SetPremierDistributedMod8SlotCount gets a reference to the given int64 and assigns it to the PremierDistributedMod8SlotCount field.
func (o *LicenseIncLicenseCount) SetPremierDistributedMod8SlotCount(v int64) {
	o.PremierDistributedMod8SlotCount = &v
}

// GetPremierMod4SlotCount returns the PremierMod4SlotCount field value if set, zero value otherwise.
func (o *LicenseIncLicenseCount) GetPremierMod4SlotCount() int64 {
	if o == nil || o.PremierMod4SlotCount == nil {
		var ret int64
		return ret
	}
	return *o.PremierMod4SlotCount
}

// GetPremierMod4SlotCountOk returns a tuple with the PremierMod4SlotCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseIncLicenseCount) GetPremierMod4SlotCountOk() (*int64, bool) {
	if o == nil || o.PremierMod4SlotCount == nil {
		return nil, false
	}
	return o.PremierMod4SlotCount, true
}

// HasPremierMod4SlotCount returns a boolean if a field has been set.
func (o *LicenseIncLicenseCount) HasPremierMod4SlotCount() bool {
	if o != nil && o.PremierMod4SlotCount != nil {
		return true
	}

	return false
}

// SetPremierMod4SlotCount gets a reference to the given int64 and assigns it to the PremierMod4SlotCount field.
func (o *LicenseIncLicenseCount) SetPremierMod4SlotCount(v int64) {
	o.PremierMod4SlotCount = &v
}

// GetPremierMod8SlotCount returns the PremierMod8SlotCount field value if set, zero value otherwise.
func (o *LicenseIncLicenseCount) GetPremierMod8SlotCount() int64 {
	if o == nil || o.PremierMod8SlotCount == nil {
		var ret int64
		return ret
	}
	return *o.PremierMod8SlotCount
}

// GetPremierMod8SlotCountOk returns a tuple with the PremierMod8SlotCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseIncLicenseCount) GetPremierMod8SlotCountOk() (*int64, bool) {
	if o == nil || o.PremierMod8SlotCount == nil {
		return nil, false
	}
	return o.PremierMod8SlotCount, true
}

// HasPremierMod8SlotCount returns a boolean if a field has been set.
func (o *LicenseIncLicenseCount) HasPremierMod8SlotCount() bool {
	if o != nil && o.PremierMod8SlotCount != nil {
		return true
	}

	return false
}

// SetPremierMod8SlotCount gets a reference to the given int64 and assigns it to the PremierMod8SlotCount field.
func (o *LicenseIncLicenseCount) SetPremierMod8SlotCount(v int64) {
	o.PremierMod8SlotCount = &v
}

// GetAccountLicenseData returns the AccountLicenseData field value if set, zero value otherwise.
func (o *LicenseIncLicenseCount) GetAccountLicenseData() LicenseAccountLicenseDataRelationship {
	if o == nil || o.AccountLicenseData == nil {
		var ret LicenseAccountLicenseDataRelationship
		return ret
	}
	return *o.AccountLicenseData
}

// GetAccountLicenseDataOk returns a tuple with the AccountLicenseData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseIncLicenseCount) GetAccountLicenseDataOk() (*LicenseAccountLicenseDataRelationship, bool) {
	if o == nil || o.AccountLicenseData == nil {
		return nil, false
	}
	return o.AccountLicenseData, true
}

// HasAccountLicenseData returns a boolean if a field has been set.
func (o *LicenseIncLicenseCount) HasAccountLicenseData() bool {
	if o != nil && o.AccountLicenseData != nil {
		return true
	}

	return false
}

// SetAccountLicenseData gets a reference to the given LicenseAccountLicenseDataRelationship and assigns it to the AccountLicenseData field.
func (o *LicenseIncLicenseCount) SetAccountLicenseData(v LicenseAccountLicenseDataRelationship) {
	o.AccountLicenseData = &v
}

func (o LicenseIncLicenseCount) MarshalJSON() ([]byte, error) {
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
	if o.Premier100GfxCount != nil {
		toSerialize["Premier100GfxCount"] = o.Premier100GfxCount
	}
	if o.Premier10GfxCount != nil {
		toSerialize["Premier10GfxCount"] = o.Premier10GfxCount
	}
	if o.Premier1GfxCount != nil {
		toSerialize["Premier1GfxCount"] = o.Premier1GfxCount
	}
	if o.PremierCentralizedMod8SlotCount != nil {
		toSerialize["PremierCentralizedMod8SlotCount"] = o.PremierCentralizedMod8SlotCount
	}
	if o.PremierD2OpsFixedCount != nil {
		toSerialize["PremierD2OpsFixedCount"] = o.PremierD2OpsFixedCount
	}
	if o.PremierD2OpsModCount != nil {
		toSerialize["PremierD2OpsModCount"] = o.PremierD2OpsModCount
	}
	if o.PremierDistributedMod8SlotCount != nil {
		toSerialize["PremierDistributedMod8SlotCount"] = o.PremierDistributedMod8SlotCount
	}
	if o.PremierMod4SlotCount != nil {
		toSerialize["PremierMod4SlotCount"] = o.PremierMod4SlotCount
	}
	if o.PremierMod8SlotCount != nil {
		toSerialize["PremierMod8SlotCount"] = o.PremierMod8SlotCount
	}
	if o.AccountLicenseData != nil {
		toSerialize["AccountLicenseData"] = o.AccountLicenseData
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LicenseIncLicenseCount) UnmarshalJSON(bytes []byte) (err error) {
	type LicenseIncLicenseCountWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The total number of devices claimed in the premier 100G fixed tier Intersight Nexus Cloud.
		Premier100GfxCount *int64 `json:"Premier100GfxCount,omitempty"`
		// The total number of devices claimed in the premier 10G fixed tier Intersight Nexus Cloud.
		Premier10GfxCount *int64 `json:"Premier10GfxCount,omitempty"`
		// The total number of devices claimed in the premier 1G fixed tier Intersight Nexus Cloud.
		Premier1GfxCount *int64 `json:"Premier1GfxCount,omitempty"`
		// The total number of devices claimed in the CentralizedMod8Slot premier tier Intersight Nexus Cloud.
		PremierCentralizedMod8SlotCount *int64 `json:"PremierCentralizedMod8SlotCount,omitempty"`
		// The total number of devices claimed in the D2Ops Fixed premier tier Intersight Nexus Cloud.
		PremierD2OpsFixedCount *int64 `json:"PremierD2OpsFixedCount,omitempty"`
		// The total number of devices claimed in the D2Ops modular premier tier Intersight Nexus Cloud.
		PremierD2OpsModCount *int64 `json:"PremierD2OpsModCount,omitempty"`
		// The total number of devices claimed in the DistributedMod8Slot premier tier Intersight Nexus Cloud.
		PremierDistributedMod8SlotCount *int64 `json:"PremierDistributedMod8SlotCount,omitempty"`
		// The total number of devices claimed in the modular 4 slot premier tier Intersight Nexus Cloud.
		PremierMod4SlotCount *int64 `json:"PremierMod4SlotCount,omitempty"`
		// The total number of devices claimed in the modular 8 slot premier tier Intersight Nexus Cloud.
		PremierMod8SlotCount *int64                                 `json:"PremierMod8SlotCount,omitempty"`
		AccountLicenseData   *LicenseAccountLicenseDataRelationship `json:"AccountLicenseData,omitempty"`
	}

	varLicenseIncLicenseCountWithoutEmbeddedStruct := LicenseIncLicenseCountWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varLicenseIncLicenseCountWithoutEmbeddedStruct)
	if err == nil {
		varLicenseIncLicenseCount := _LicenseIncLicenseCount{}
		varLicenseIncLicenseCount.ClassId = varLicenseIncLicenseCountWithoutEmbeddedStruct.ClassId
		varLicenseIncLicenseCount.ObjectType = varLicenseIncLicenseCountWithoutEmbeddedStruct.ObjectType
		varLicenseIncLicenseCount.Premier100GfxCount = varLicenseIncLicenseCountWithoutEmbeddedStruct.Premier100GfxCount
		varLicenseIncLicenseCount.Premier10GfxCount = varLicenseIncLicenseCountWithoutEmbeddedStruct.Premier10GfxCount
		varLicenseIncLicenseCount.Premier1GfxCount = varLicenseIncLicenseCountWithoutEmbeddedStruct.Premier1GfxCount
		varLicenseIncLicenseCount.PremierCentralizedMod8SlotCount = varLicenseIncLicenseCountWithoutEmbeddedStruct.PremierCentralizedMod8SlotCount
		varLicenseIncLicenseCount.PremierD2OpsFixedCount = varLicenseIncLicenseCountWithoutEmbeddedStruct.PremierD2OpsFixedCount
		varLicenseIncLicenseCount.PremierD2OpsModCount = varLicenseIncLicenseCountWithoutEmbeddedStruct.PremierD2OpsModCount
		varLicenseIncLicenseCount.PremierDistributedMod8SlotCount = varLicenseIncLicenseCountWithoutEmbeddedStruct.PremierDistributedMod8SlotCount
		varLicenseIncLicenseCount.PremierMod4SlotCount = varLicenseIncLicenseCountWithoutEmbeddedStruct.PremierMod4SlotCount
		varLicenseIncLicenseCount.PremierMod8SlotCount = varLicenseIncLicenseCountWithoutEmbeddedStruct.PremierMod8SlotCount
		varLicenseIncLicenseCount.AccountLicenseData = varLicenseIncLicenseCountWithoutEmbeddedStruct.AccountLicenseData
		*o = LicenseIncLicenseCount(varLicenseIncLicenseCount)
	} else {
		return err
	}

	varLicenseIncLicenseCount := _LicenseIncLicenseCount{}

	err = json.Unmarshal(bytes, &varLicenseIncLicenseCount)
	if err == nil {
		o.MoBaseMo = varLicenseIncLicenseCount.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Premier100GfxCount")
		delete(additionalProperties, "Premier10GfxCount")
		delete(additionalProperties, "Premier1GfxCount")
		delete(additionalProperties, "PremierCentralizedMod8SlotCount")
		delete(additionalProperties, "PremierD2OpsFixedCount")
		delete(additionalProperties, "PremierD2OpsModCount")
		delete(additionalProperties, "PremierDistributedMod8SlotCount")
		delete(additionalProperties, "PremierMod4SlotCount")
		delete(additionalProperties, "PremierMod8SlotCount")
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

type NullableLicenseIncLicenseCount struct {
	value *LicenseIncLicenseCount
	isSet bool
}

func (v NullableLicenseIncLicenseCount) Get() *LicenseIncLicenseCount {
	return v.value
}

func (v *NullableLicenseIncLicenseCount) Set(val *LicenseIncLicenseCount) {
	v.value = val
	v.isSet = true
}

func (v NullableLicenseIncLicenseCount) IsSet() bool {
	return v.isSet
}

func (v *NullableLicenseIncLicenseCount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicenseIncLicenseCount(val *LicenseIncLicenseCount) *NullableLicenseIncLicenseCount {
	return &NullableLicenseIncLicenseCount{value: val, isSet: true}
}

func (v NullableLicenseIncLicenseCount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicenseIncLicenseCount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
