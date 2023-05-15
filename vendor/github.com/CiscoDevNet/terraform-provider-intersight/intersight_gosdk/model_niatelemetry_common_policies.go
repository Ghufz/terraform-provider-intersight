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

// NiatelemetryCommonPolicies Object to capture Common Policy details.
type NiatelemetryCommonPolicies struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Dn of the Common Policy in APIC.
	Dn *string `json:"Dn,omitempty"`
	// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
	RecordType *string `json:"RecordType,omitempty"`
	// Version of record being pushed. This determines what was the API version for data available from the device.
	RecordVersion *string `json:"RecordVersion,omitempty"`
	// Name of the APIC site from which this data is being collected.
	SiteName *string `json:"SiteName,omitempty"`
	// List of Dn of SNMP Src for the above common pol.
	SnmpSrc *string `json:"SnmpSrc,omitempty"`
	// List of Dn of Syslog Src for the above common pol.
	SyslogSrc *string `json:"SyslogSrc,omitempty"`
	// List of Dn of Syslog Sys Msg the above common pol.
	SyslogSysMsg         *string                              `json:"SyslogSysMsg,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryCommonPolicies NiatelemetryCommonPolicies

// NewNiatelemetryCommonPolicies instantiates a new NiatelemetryCommonPolicies object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryCommonPolicies(classId string, objectType string) *NiatelemetryCommonPolicies {
	this := NiatelemetryCommonPolicies{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryCommonPoliciesWithDefaults instantiates a new NiatelemetryCommonPolicies object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryCommonPoliciesWithDefaults() *NiatelemetryCommonPolicies {
	this := NiatelemetryCommonPolicies{}
	var classId string = "niatelemetry.CommonPolicies"
	this.ClassId = classId
	var objectType string = "niatelemetry.CommonPolicies"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryCommonPolicies) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryCommonPolicies) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryCommonPolicies) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryCommonPolicies) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryCommonPolicies) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryCommonPolicies) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDn returns the Dn field value if set, zero value otherwise.
func (o *NiatelemetryCommonPolicies) GetDn() string {
	if o == nil || o.Dn == nil {
		var ret string
		return ret
	}
	return *o.Dn
}

// GetDnOk returns a tuple with the Dn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryCommonPolicies) GetDnOk() (*string, bool) {
	if o == nil || o.Dn == nil {
		return nil, false
	}
	return o.Dn, true
}

// HasDn returns a boolean if a field has been set.
func (o *NiatelemetryCommonPolicies) HasDn() bool {
	if o != nil && o.Dn != nil {
		return true
	}

	return false
}

// SetDn gets a reference to the given string and assigns it to the Dn field.
func (o *NiatelemetryCommonPolicies) SetDn(v string) {
	o.Dn = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *NiatelemetryCommonPolicies) GetRecordType() string {
	if o == nil || o.RecordType == nil {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryCommonPolicies) GetRecordTypeOk() (*string, bool) {
	if o == nil || o.RecordType == nil {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *NiatelemetryCommonPolicies) HasRecordType() bool {
	if o != nil && o.RecordType != nil {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *NiatelemetryCommonPolicies) SetRecordType(v string) {
	o.RecordType = &v
}

// GetRecordVersion returns the RecordVersion field value if set, zero value otherwise.
func (o *NiatelemetryCommonPolicies) GetRecordVersion() string {
	if o == nil || o.RecordVersion == nil {
		var ret string
		return ret
	}
	return *o.RecordVersion
}

// GetRecordVersionOk returns a tuple with the RecordVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryCommonPolicies) GetRecordVersionOk() (*string, bool) {
	if o == nil || o.RecordVersion == nil {
		return nil, false
	}
	return o.RecordVersion, true
}

// HasRecordVersion returns a boolean if a field has been set.
func (o *NiatelemetryCommonPolicies) HasRecordVersion() bool {
	if o != nil && o.RecordVersion != nil {
		return true
	}

	return false
}

// SetRecordVersion gets a reference to the given string and assigns it to the RecordVersion field.
func (o *NiatelemetryCommonPolicies) SetRecordVersion(v string) {
	o.RecordVersion = &v
}

// GetSiteName returns the SiteName field value if set, zero value otherwise.
func (o *NiatelemetryCommonPolicies) GetSiteName() string {
	if o == nil || o.SiteName == nil {
		var ret string
		return ret
	}
	return *o.SiteName
}

// GetSiteNameOk returns a tuple with the SiteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryCommonPolicies) GetSiteNameOk() (*string, bool) {
	if o == nil || o.SiteName == nil {
		return nil, false
	}
	return o.SiteName, true
}

// HasSiteName returns a boolean if a field has been set.
func (o *NiatelemetryCommonPolicies) HasSiteName() bool {
	if o != nil && o.SiteName != nil {
		return true
	}

	return false
}

// SetSiteName gets a reference to the given string and assigns it to the SiteName field.
func (o *NiatelemetryCommonPolicies) SetSiteName(v string) {
	o.SiteName = &v
}

// GetSnmpSrc returns the SnmpSrc field value if set, zero value otherwise.
func (o *NiatelemetryCommonPolicies) GetSnmpSrc() string {
	if o == nil || o.SnmpSrc == nil {
		var ret string
		return ret
	}
	return *o.SnmpSrc
}

// GetSnmpSrcOk returns a tuple with the SnmpSrc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryCommonPolicies) GetSnmpSrcOk() (*string, bool) {
	if o == nil || o.SnmpSrc == nil {
		return nil, false
	}
	return o.SnmpSrc, true
}

// HasSnmpSrc returns a boolean if a field has been set.
func (o *NiatelemetryCommonPolicies) HasSnmpSrc() bool {
	if o != nil && o.SnmpSrc != nil {
		return true
	}

	return false
}

// SetSnmpSrc gets a reference to the given string and assigns it to the SnmpSrc field.
func (o *NiatelemetryCommonPolicies) SetSnmpSrc(v string) {
	o.SnmpSrc = &v
}

// GetSyslogSrc returns the SyslogSrc field value if set, zero value otherwise.
func (o *NiatelemetryCommonPolicies) GetSyslogSrc() string {
	if o == nil || o.SyslogSrc == nil {
		var ret string
		return ret
	}
	return *o.SyslogSrc
}

// GetSyslogSrcOk returns a tuple with the SyslogSrc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryCommonPolicies) GetSyslogSrcOk() (*string, bool) {
	if o == nil || o.SyslogSrc == nil {
		return nil, false
	}
	return o.SyslogSrc, true
}

// HasSyslogSrc returns a boolean if a field has been set.
func (o *NiatelemetryCommonPolicies) HasSyslogSrc() bool {
	if o != nil && o.SyslogSrc != nil {
		return true
	}

	return false
}

// SetSyslogSrc gets a reference to the given string and assigns it to the SyslogSrc field.
func (o *NiatelemetryCommonPolicies) SetSyslogSrc(v string) {
	o.SyslogSrc = &v
}

// GetSyslogSysMsg returns the SyslogSysMsg field value if set, zero value otherwise.
func (o *NiatelemetryCommonPolicies) GetSyslogSysMsg() string {
	if o == nil || o.SyslogSysMsg == nil {
		var ret string
		return ret
	}
	return *o.SyslogSysMsg
}

// GetSyslogSysMsgOk returns a tuple with the SyslogSysMsg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryCommonPolicies) GetSyslogSysMsgOk() (*string, bool) {
	if o == nil || o.SyslogSysMsg == nil {
		return nil, false
	}
	return o.SyslogSysMsg, true
}

// HasSyslogSysMsg returns a boolean if a field has been set.
func (o *NiatelemetryCommonPolicies) HasSyslogSysMsg() bool {
	if o != nil && o.SyslogSysMsg != nil {
		return true
	}

	return false
}

// SetSyslogSysMsg gets a reference to the given string and assigns it to the SyslogSysMsg field.
func (o *NiatelemetryCommonPolicies) SetSyslogSysMsg(v string) {
	o.SyslogSysMsg = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NiatelemetryCommonPolicies) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryCommonPolicies) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetryCommonPolicies) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetryCommonPolicies) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NiatelemetryCommonPolicies) MarshalJSON() ([]byte, error) {
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
	if o.Dn != nil {
		toSerialize["Dn"] = o.Dn
	}
	if o.RecordType != nil {
		toSerialize["RecordType"] = o.RecordType
	}
	if o.RecordVersion != nil {
		toSerialize["RecordVersion"] = o.RecordVersion
	}
	if o.SiteName != nil {
		toSerialize["SiteName"] = o.SiteName
	}
	if o.SnmpSrc != nil {
		toSerialize["SnmpSrc"] = o.SnmpSrc
	}
	if o.SyslogSrc != nil {
		toSerialize["SyslogSrc"] = o.SyslogSrc
	}
	if o.SyslogSysMsg != nil {
		toSerialize["SyslogSysMsg"] = o.SyslogSysMsg
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryCommonPolicies) UnmarshalJSON(bytes []byte) (err error) {
	type NiatelemetryCommonPoliciesWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Dn of the Common Policy in APIC.
		Dn *string `json:"Dn,omitempty"`
		// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
		RecordType *string `json:"RecordType,omitempty"`
		// Version of record being pushed. This determines what was the API version for data available from the device.
		RecordVersion *string `json:"RecordVersion,omitempty"`
		// Name of the APIC site from which this data is being collected.
		SiteName *string `json:"SiteName,omitempty"`
		// List of Dn of SNMP Src for the above common pol.
		SnmpSrc *string `json:"SnmpSrc,omitempty"`
		// List of Dn of Syslog Src for the above common pol.
		SyslogSrc *string `json:"SyslogSrc,omitempty"`
		// List of Dn of Syslog Sys Msg the above common pol.
		SyslogSysMsg     *string                              `json:"SyslogSysMsg,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varNiatelemetryCommonPoliciesWithoutEmbeddedStruct := NiatelemetryCommonPoliciesWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNiatelemetryCommonPoliciesWithoutEmbeddedStruct)
	if err == nil {
		varNiatelemetryCommonPolicies := _NiatelemetryCommonPolicies{}
		varNiatelemetryCommonPolicies.ClassId = varNiatelemetryCommonPoliciesWithoutEmbeddedStruct.ClassId
		varNiatelemetryCommonPolicies.ObjectType = varNiatelemetryCommonPoliciesWithoutEmbeddedStruct.ObjectType
		varNiatelemetryCommonPolicies.Dn = varNiatelemetryCommonPoliciesWithoutEmbeddedStruct.Dn
		varNiatelemetryCommonPolicies.RecordType = varNiatelemetryCommonPoliciesWithoutEmbeddedStruct.RecordType
		varNiatelemetryCommonPolicies.RecordVersion = varNiatelemetryCommonPoliciesWithoutEmbeddedStruct.RecordVersion
		varNiatelemetryCommonPolicies.SiteName = varNiatelemetryCommonPoliciesWithoutEmbeddedStruct.SiteName
		varNiatelemetryCommonPolicies.SnmpSrc = varNiatelemetryCommonPoliciesWithoutEmbeddedStruct.SnmpSrc
		varNiatelemetryCommonPolicies.SyslogSrc = varNiatelemetryCommonPoliciesWithoutEmbeddedStruct.SyslogSrc
		varNiatelemetryCommonPolicies.SyslogSysMsg = varNiatelemetryCommonPoliciesWithoutEmbeddedStruct.SyslogSysMsg
		varNiatelemetryCommonPolicies.RegisteredDevice = varNiatelemetryCommonPoliciesWithoutEmbeddedStruct.RegisteredDevice
		*o = NiatelemetryCommonPolicies(varNiatelemetryCommonPolicies)
	} else {
		return err
	}

	varNiatelemetryCommonPolicies := _NiatelemetryCommonPolicies{}

	err = json.Unmarshal(bytes, &varNiatelemetryCommonPolicies)
	if err == nil {
		o.MoBaseMo = varNiatelemetryCommonPolicies.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Dn")
		delete(additionalProperties, "RecordType")
		delete(additionalProperties, "RecordVersion")
		delete(additionalProperties, "SiteName")
		delete(additionalProperties, "SnmpSrc")
		delete(additionalProperties, "SyslogSrc")
		delete(additionalProperties, "SyslogSysMsg")
		delete(additionalProperties, "RegisteredDevice")

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

type NullableNiatelemetryCommonPolicies struct {
	value *NiatelemetryCommonPolicies
	isSet bool
}

func (v NullableNiatelemetryCommonPolicies) Get() *NiatelemetryCommonPolicies {
	return v.value
}

func (v *NullableNiatelemetryCommonPolicies) Set(val *NiatelemetryCommonPolicies) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryCommonPolicies) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryCommonPolicies) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryCommonPolicies(val *NiatelemetryCommonPolicies) *NullableNiatelemetryCommonPolicies {
	return &NullableNiatelemetryCommonPolicies{value: val, isSet: true}
}

func (v NullableNiatelemetryCommonPolicies) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryCommonPolicies) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
