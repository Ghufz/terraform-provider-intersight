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

// NiatelemetryApicNtpAuth Object to capture the authentication details of NTP in APIC.
type NiatelemetryApicNtpAuth struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Authentication state of NTP policy in APIC.
	AuthSt *string `json:"AuthSt,omitempty"`
	// Dn of the NTP Auth in APIC.
	Dn *string `json:"Dn,omitempty"`
	// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
	RecordType *string `json:"RecordType,omitempty"`
	// Version of record being pushed. This determines what was the API version for data available from the device.
	RecordVersion *string `json:"RecordVersion,omitempty"`
	// Name of the APIC site from which this data is being collected.
	SiteName             *string                              `json:"SiteName,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryApicNtpAuth NiatelemetryApicNtpAuth

// NewNiatelemetryApicNtpAuth instantiates a new NiatelemetryApicNtpAuth object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryApicNtpAuth(classId string, objectType string) *NiatelemetryApicNtpAuth {
	this := NiatelemetryApicNtpAuth{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryApicNtpAuthWithDefaults instantiates a new NiatelemetryApicNtpAuth object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryApicNtpAuthWithDefaults() *NiatelemetryApicNtpAuth {
	this := NiatelemetryApicNtpAuth{}
	var classId string = "niatelemetry.ApicNtpAuth"
	this.ClassId = classId
	var objectType string = "niatelemetry.ApicNtpAuth"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryApicNtpAuth) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicNtpAuth) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryApicNtpAuth) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryApicNtpAuth) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicNtpAuth) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryApicNtpAuth) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAuthSt returns the AuthSt field value if set, zero value otherwise.
func (o *NiatelemetryApicNtpAuth) GetAuthSt() string {
	if o == nil || o.AuthSt == nil {
		var ret string
		return ret
	}
	return *o.AuthSt
}

// GetAuthStOk returns a tuple with the AuthSt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicNtpAuth) GetAuthStOk() (*string, bool) {
	if o == nil || o.AuthSt == nil {
		return nil, false
	}
	return o.AuthSt, true
}

// HasAuthSt returns a boolean if a field has been set.
func (o *NiatelemetryApicNtpAuth) HasAuthSt() bool {
	if o != nil && o.AuthSt != nil {
		return true
	}

	return false
}

// SetAuthSt gets a reference to the given string and assigns it to the AuthSt field.
func (o *NiatelemetryApicNtpAuth) SetAuthSt(v string) {
	o.AuthSt = &v
}

// GetDn returns the Dn field value if set, zero value otherwise.
func (o *NiatelemetryApicNtpAuth) GetDn() string {
	if o == nil || o.Dn == nil {
		var ret string
		return ret
	}
	return *o.Dn
}

// GetDnOk returns a tuple with the Dn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicNtpAuth) GetDnOk() (*string, bool) {
	if o == nil || o.Dn == nil {
		return nil, false
	}
	return o.Dn, true
}

// HasDn returns a boolean if a field has been set.
func (o *NiatelemetryApicNtpAuth) HasDn() bool {
	if o != nil && o.Dn != nil {
		return true
	}

	return false
}

// SetDn gets a reference to the given string and assigns it to the Dn field.
func (o *NiatelemetryApicNtpAuth) SetDn(v string) {
	o.Dn = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *NiatelemetryApicNtpAuth) GetRecordType() string {
	if o == nil || o.RecordType == nil {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicNtpAuth) GetRecordTypeOk() (*string, bool) {
	if o == nil || o.RecordType == nil {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *NiatelemetryApicNtpAuth) HasRecordType() bool {
	if o != nil && o.RecordType != nil {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *NiatelemetryApicNtpAuth) SetRecordType(v string) {
	o.RecordType = &v
}

// GetRecordVersion returns the RecordVersion field value if set, zero value otherwise.
func (o *NiatelemetryApicNtpAuth) GetRecordVersion() string {
	if o == nil || o.RecordVersion == nil {
		var ret string
		return ret
	}
	return *o.RecordVersion
}

// GetRecordVersionOk returns a tuple with the RecordVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicNtpAuth) GetRecordVersionOk() (*string, bool) {
	if o == nil || o.RecordVersion == nil {
		return nil, false
	}
	return o.RecordVersion, true
}

// HasRecordVersion returns a boolean if a field has been set.
func (o *NiatelemetryApicNtpAuth) HasRecordVersion() bool {
	if o != nil && o.RecordVersion != nil {
		return true
	}

	return false
}

// SetRecordVersion gets a reference to the given string and assigns it to the RecordVersion field.
func (o *NiatelemetryApicNtpAuth) SetRecordVersion(v string) {
	o.RecordVersion = &v
}

// GetSiteName returns the SiteName field value if set, zero value otherwise.
func (o *NiatelemetryApicNtpAuth) GetSiteName() string {
	if o == nil || o.SiteName == nil {
		var ret string
		return ret
	}
	return *o.SiteName
}

// GetSiteNameOk returns a tuple with the SiteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicNtpAuth) GetSiteNameOk() (*string, bool) {
	if o == nil || o.SiteName == nil {
		return nil, false
	}
	return o.SiteName, true
}

// HasSiteName returns a boolean if a field has been set.
func (o *NiatelemetryApicNtpAuth) HasSiteName() bool {
	if o != nil && o.SiteName != nil {
		return true
	}

	return false
}

// SetSiteName gets a reference to the given string and assigns it to the SiteName field.
func (o *NiatelemetryApicNtpAuth) SetSiteName(v string) {
	o.SiteName = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NiatelemetryApicNtpAuth) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicNtpAuth) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetryApicNtpAuth) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetryApicNtpAuth) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NiatelemetryApicNtpAuth) MarshalJSON() ([]byte, error) {
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
	if o.AuthSt != nil {
		toSerialize["AuthSt"] = o.AuthSt
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
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryApicNtpAuth) UnmarshalJSON(bytes []byte) (err error) {
	type NiatelemetryApicNtpAuthWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Authentication state of NTP policy in APIC.
		AuthSt *string `json:"AuthSt,omitempty"`
		// Dn of the NTP Auth in APIC.
		Dn *string `json:"Dn,omitempty"`
		// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
		RecordType *string `json:"RecordType,omitempty"`
		// Version of record being pushed. This determines what was the API version for data available from the device.
		RecordVersion *string `json:"RecordVersion,omitempty"`
		// Name of the APIC site from which this data is being collected.
		SiteName         *string                              `json:"SiteName,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varNiatelemetryApicNtpAuthWithoutEmbeddedStruct := NiatelemetryApicNtpAuthWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNiatelemetryApicNtpAuthWithoutEmbeddedStruct)
	if err == nil {
		varNiatelemetryApicNtpAuth := _NiatelemetryApicNtpAuth{}
		varNiatelemetryApicNtpAuth.ClassId = varNiatelemetryApicNtpAuthWithoutEmbeddedStruct.ClassId
		varNiatelemetryApicNtpAuth.ObjectType = varNiatelemetryApicNtpAuthWithoutEmbeddedStruct.ObjectType
		varNiatelemetryApicNtpAuth.AuthSt = varNiatelemetryApicNtpAuthWithoutEmbeddedStruct.AuthSt
		varNiatelemetryApicNtpAuth.Dn = varNiatelemetryApicNtpAuthWithoutEmbeddedStruct.Dn
		varNiatelemetryApicNtpAuth.RecordType = varNiatelemetryApicNtpAuthWithoutEmbeddedStruct.RecordType
		varNiatelemetryApicNtpAuth.RecordVersion = varNiatelemetryApicNtpAuthWithoutEmbeddedStruct.RecordVersion
		varNiatelemetryApicNtpAuth.SiteName = varNiatelemetryApicNtpAuthWithoutEmbeddedStruct.SiteName
		varNiatelemetryApicNtpAuth.RegisteredDevice = varNiatelemetryApicNtpAuthWithoutEmbeddedStruct.RegisteredDevice
		*o = NiatelemetryApicNtpAuth(varNiatelemetryApicNtpAuth)
	} else {
		return err
	}

	varNiatelemetryApicNtpAuth := _NiatelemetryApicNtpAuth{}

	err = json.Unmarshal(bytes, &varNiatelemetryApicNtpAuth)
	if err == nil {
		o.MoBaseMo = varNiatelemetryApicNtpAuth.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AuthSt")
		delete(additionalProperties, "Dn")
		delete(additionalProperties, "RecordType")
		delete(additionalProperties, "RecordVersion")
		delete(additionalProperties, "SiteName")
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

type NullableNiatelemetryApicNtpAuth struct {
	value *NiatelemetryApicNtpAuth
	isSet bool
}

func (v NullableNiatelemetryApicNtpAuth) Get() *NiatelemetryApicNtpAuth {
	return v.value
}

func (v *NullableNiatelemetryApicNtpAuth) Set(val *NiatelemetryApicNtpAuth) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryApicNtpAuth) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryApicNtpAuth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryApicNtpAuth(val *NiatelemetryApicNtpAuth) *NullableNiatelemetryApicNtpAuth {
	return &NullableNiatelemetryApicNtpAuth{value: val, isSet: true}
}

func (v NullableNiatelemetryApicNtpAuth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryApicNtpAuth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
