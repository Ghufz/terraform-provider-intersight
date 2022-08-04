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

// NiatelemetryApicSysLogSrcAllOf Definition of the list of properties defined in 'niatelemetry.ApicSysLogSrc', excluding properties defined in parent classes.
type NiatelemetryApicSysLogSrcAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Dn of the SysLogSrc in APIC.
	Dn *string `json:"Dn,omitempty"`
	// Minimum sevirity on SysLogSrc MO in APIC.
	MinSev *string `json:"MinSev,omitempty"`
	// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
	RecordType *string `json:"RecordType,omitempty"`
	// Version of record being pushed. This determines what was the API version for data available from the device.
	RecordVersion *string `json:"RecordVersion,omitempty"`
	// Name of the APIC site from which this data is being collected.
	SiteName *string `json:"SiteName,omitempty"`
	// List of Syslog remote destination for SyslogSrc in APIC.
	SyslogRemoteDest *string `json:"SyslogRemoteDest,omitempty"`
	// Syslog destination grp for SysLogSrc in APIC.
	SyslogRsDestGrp      *string                              `json:"SyslogRsDestGrp,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryApicSysLogSrcAllOf NiatelemetryApicSysLogSrcAllOf

// NewNiatelemetryApicSysLogSrcAllOf instantiates a new NiatelemetryApicSysLogSrcAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryApicSysLogSrcAllOf(classId string, objectType string) *NiatelemetryApicSysLogSrcAllOf {
	this := NiatelemetryApicSysLogSrcAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryApicSysLogSrcAllOfWithDefaults instantiates a new NiatelemetryApicSysLogSrcAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryApicSysLogSrcAllOfWithDefaults() *NiatelemetryApicSysLogSrcAllOf {
	this := NiatelemetryApicSysLogSrcAllOf{}
	var classId string = "niatelemetry.ApicSysLogSrc"
	this.ClassId = classId
	var objectType string = "niatelemetry.ApicSysLogSrc"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryApicSysLogSrcAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicSysLogSrcAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryApicSysLogSrcAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryApicSysLogSrcAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicSysLogSrcAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryApicSysLogSrcAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDn returns the Dn field value if set, zero value otherwise.
func (o *NiatelemetryApicSysLogSrcAllOf) GetDn() string {
	if o == nil || o.Dn == nil {
		var ret string
		return ret
	}
	return *o.Dn
}

// GetDnOk returns a tuple with the Dn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicSysLogSrcAllOf) GetDnOk() (*string, bool) {
	if o == nil || o.Dn == nil {
		return nil, false
	}
	return o.Dn, true
}

// HasDn returns a boolean if a field has been set.
func (o *NiatelemetryApicSysLogSrcAllOf) HasDn() bool {
	if o != nil && o.Dn != nil {
		return true
	}

	return false
}

// SetDn gets a reference to the given string and assigns it to the Dn field.
func (o *NiatelemetryApicSysLogSrcAllOf) SetDn(v string) {
	o.Dn = &v
}

// GetMinSev returns the MinSev field value if set, zero value otherwise.
func (o *NiatelemetryApicSysLogSrcAllOf) GetMinSev() string {
	if o == nil || o.MinSev == nil {
		var ret string
		return ret
	}
	return *o.MinSev
}

// GetMinSevOk returns a tuple with the MinSev field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicSysLogSrcAllOf) GetMinSevOk() (*string, bool) {
	if o == nil || o.MinSev == nil {
		return nil, false
	}
	return o.MinSev, true
}

// HasMinSev returns a boolean if a field has been set.
func (o *NiatelemetryApicSysLogSrcAllOf) HasMinSev() bool {
	if o != nil && o.MinSev != nil {
		return true
	}

	return false
}

// SetMinSev gets a reference to the given string and assigns it to the MinSev field.
func (o *NiatelemetryApicSysLogSrcAllOf) SetMinSev(v string) {
	o.MinSev = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *NiatelemetryApicSysLogSrcAllOf) GetRecordType() string {
	if o == nil || o.RecordType == nil {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicSysLogSrcAllOf) GetRecordTypeOk() (*string, bool) {
	if o == nil || o.RecordType == nil {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *NiatelemetryApicSysLogSrcAllOf) HasRecordType() bool {
	if o != nil && o.RecordType != nil {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *NiatelemetryApicSysLogSrcAllOf) SetRecordType(v string) {
	o.RecordType = &v
}

// GetRecordVersion returns the RecordVersion field value if set, zero value otherwise.
func (o *NiatelemetryApicSysLogSrcAllOf) GetRecordVersion() string {
	if o == nil || o.RecordVersion == nil {
		var ret string
		return ret
	}
	return *o.RecordVersion
}

// GetRecordVersionOk returns a tuple with the RecordVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicSysLogSrcAllOf) GetRecordVersionOk() (*string, bool) {
	if o == nil || o.RecordVersion == nil {
		return nil, false
	}
	return o.RecordVersion, true
}

// HasRecordVersion returns a boolean if a field has been set.
func (o *NiatelemetryApicSysLogSrcAllOf) HasRecordVersion() bool {
	if o != nil && o.RecordVersion != nil {
		return true
	}

	return false
}

// SetRecordVersion gets a reference to the given string and assigns it to the RecordVersion field.
func (o *NiatelemetryApicSysLogSrcAllOf) SetRecordVersion(v string) {
	o.RecordVersion = &v
}

// GetSiteName returns the SiteName field value if set, zero value otherwise.
func (o *NiatelemetryApicSysLogSrcAllOf) GetSiteName() string {
	if o == nil || o.SiteName == nil {
		var ret string
		return ret
	}
	return *o.SiteName
}

// GetSiteNameOk returns a tuple with the SiteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicSysLogSrcAllOf) GetSiteNameOk() (*string, bool) {
	if o == nil || o.SiteName == nil {
		return nil, false
	}
	return o.SiteName, true
}

// HasSiteName returns a boolean if a field has been set.
func (o *NiatelemetryApicSysLogSrcAllOf) HasSiteName() bool {
	if o != nil && o.SiteName != nil {
		return true
	}

	return false
}

// SetSiteName gets a reference to the given string and assigns it to the SiteName field.
func (o *NiatelemetryApicSysLogSrcAllOf) SetSiteName(v string) {
	o.SiteName = &v
}

// GetSyslogRemoteDest returns the SyslogRemoteDest field value if set, zero value otherwise.
func (o *NiatelemetryApicSysLogSrcAllOf) GetSyslogRemoteDest() string {
	if o == nil || o.SyslogRemoteDest == nil {
		var ret string
		return ret
	}
	return *o.SyslogRemoteDest
}

// GetSyslogRemoteDestOk returns a tuple with the SyslogRemoteDest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicSysLogSrcAllOf) GetSyslogRemoteDestOk() (*string, bool) {
	if o == nil || o.SyslogRemoteDest == nil {
		return nil, false
	}
	return o.SyslogRemoteDest, true
}

// HasSyslogRemoteDest returns a boolean if a field has been set.
func (o *NiatelemetryApicSysLogSrcAllOf) HasSyslogRemoteDest() bool {
	if o != nil && o.SyslogRemoteDest != nil {
		return true
	}

	return false
}

// SetSyslogRemoteDest gets a reference to the given string and assigns it to the SyslogRemoteDest field.
func (o *NiatelemetryApicSysLogSrcAllOf) SetSyslogRemoteDest(v string) {
	o.SyslogRemoteDest = &v
}

// GetSyslogRsDestGrp returns the SyslogRsDestGrp field value if set, zero value otherwise.
func (o *NiatelemetryApicSysLogSrcAllOf) GetSyslogRsDestGrp() string {
	if o == nil || o.SyslogRsDestGrp == nil {
		var ret string
		return ret
	}
	return *o.SyslogRsDestGrp
}

// GetSyslogRsDestGrpOk returns a tuple with the SyslogRsDestGrp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicSysLogSrcAllOf) GetSyslogRsDestGrpOk() (*string, bool) {
	if o == nil || o.SyslogRsDestGrp == nil {
		return nil, false
	}
	return o.SyslogRsDestGrp, true
}

// HasSyslogRsDestGrp returns a boolean if a field has been set.
func (o *NiatelemetryApicSysLogSrcAllOf) HasSyslogRsDestGrp() bool {
	if o != nil && o.SyslogRsDestGrp != nil {
		return true
	}

	return false
}

// SetSyslogRsDestGrp gets a reference to the given string and assigns it to the SyslogRsDestGrp field.
func (o *NiatelemetryApicSysLogSrcAllOf) SetSyslogRsDestGrp(v string) {
	o.SyslogRsDestGrp = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NiatelemetryApicSysLogSrcAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicSysLogSrcAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetryApicSysLogSrcAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetryApicSysLogSrcAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NiatelemetryApicSysLogSrcAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Dn != nil {
		toSerialize["Dn"] = o.Dn
	}
	if o.MinSev != nil {
		toSerialize["MinSev"] = o.MinSev
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
	if o.SyslogRemoteDest != nil {
		toSerialize["SyslogRemoteDest"] = o.SyslogRemoteDest
	}
	if o.SyslogRsDestGrp != nil {
		toSerialize["SyslogRsDestGrp"] = o.SyslogRsDestGrp
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryApicSysLogSrcAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varNiatelemetryApicSysLogSrcAllOf := _NiatelemetryApicSysLogSrcAllOf{}

	if err = json.Unmarshal(bytes, &varNiatelemetryApicSysLogSrcAllOf); err == nil {
		*o = NiatelemetryApicSysLogSrcAllOf(varNiatelemetryApicSysLogSrcAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Dn")
		delete(additionalProperties, "MinSev")
		delete(additionalProperties, "RecordType")
		delete(additionalProperties, "RecordVersion")
		delete(additionalProperties, "SiteName")
		delete(additionalProperties, "SyslogRemoteDest")
		delete(additionalProperties, "SyslogRsDestGrp")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNiatelemetryApicSysLogSrcAllOf struct {
	value *NiatelemetryApicSysLogSrcAllOf
	isSet bool
}

func (v NullableNiatelemetryApicSysLogSrcAllOf) Get() *NiatelemetryApicSysLogSrcAllOf {
	return v.value
}

func (v *NullableNiatelemetryApicSysLogSrcAllOf) Set(val *NiatelemetryApicSysLogSrcAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryApicSysLogSrcAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryApicSysLogSrcAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryApicSysLogSrcAllOf(val *NiatelemetryApicSysLogSrcAllOf) *NullableNiatelemetryApicSysLogSrcAllOf {
	return &NullableNiatelemetryApicSysLogSrcAllOf{value: val, isSet: true}
}

func (v NullableNiatelemetryApicSysLogSrcAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryApicSysLogSrcAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
