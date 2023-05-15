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
)

// NiatelemetryAaaTacacsProviderDetailsAllOf Definition of the list of properties defined in 'niatelemetry.AaaTacacsProviderDetails', excluding properties defined in parent classes.
type NiatelemetryAaaTacacsProviderDetailsAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Dn of the AAA tacacs provider in APIC.
	Dn *string `json:"Dn,omitempty"`
	// Owner Key of the AAA tacacs provider in APIC.
	OwnerKey *string `json:"OwnerKey,omitempty"`
	// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
	RecordType *string `json:"RecordType,omitempty"`
	// Version of record being pushed. This determines what was the API version for data available from the device.
	RecordVersion *string `json:"RecordVersion,omitempty"`
	// Name of the APIC site from which this data is being collected.
	SiteName             *string                              `json:"SiteName,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryAaaTacacsProviderDetailsAllOf NiatelemetryAaaTacacsProviderDetailsAllOf

// NewNiatelemetryAaaTacacsProviderDetailsAllOf instantiates a new NiatelemetryAaaTacacsProviderDetailsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryAaaTacacsProviderDetailsAllOf(classId string, objectType string) *NiatelemetryAaaTacacsProviderDetailsAllOf {
	this := NiatelemetryAaaTacacsProviderDetailsAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryAaaTacacsProviderDetailsAllOfWithDefaults instantiates a new NiatelemetryAaaTacacsProviderDetailsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryAaaTacacsProviderDetailsAllOfWithDefaults() *NiatelemetryAaaTacacsProviderDetailsAllOf {
	this := NiatelemetryAaaTacacsProviderDetailsAllOf{}
	var classId string = "niatelemetry.AaaTacacsProviderDetails"
	this.ClassId = classId
	var objectType string = "niatelemetry.AaaTacacsProviderDetails"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryAaaTacacsProviderDetailsAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryAaaTacacsProviderDetailsAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryAaaTacacsProviderDetailsAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryAaaTacacsProviderDetailsAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryAaaTacacsProviderDetailsAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryAaaTacacsProviderDetailsAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDn returns the Dn field value if set, zero value otherwise.
func (o *NiatelemetryAaaTacacsProviderDetailsAllOf) GetDn() string {
	if o == nil || o.Dn == nil {
		var ret string
		return ret
	}
	return *o.Dn
}

// GetDnOk returns a tuple with the Dn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryAaaTacacsProviderDetailsAllOf) GetDnOk() (*string, bool) {
	if o == nil || o.Dn == nil {
		return nil, false
	}
	return o.Dn, true
}

// HasDn returns a boolean if a field has been set.
func (o *NiatelemetryAaaTacacsProviderDetailsAllOf) HasDn() bool {
	if o != nil && o.Dn != nil {
		return true
	}

	return false
}

// SetDn gets a reference to the given string and assigns it to the Dn field.
func (o *NiatelemetryAaaTacacsProviderDetailsAllOf) SetDn(v string) {
	o.Dn = &v
}

// GetOwnerKey returns the OwnerKey field value if set, zero value otherwise.
func (o *NiatelemetryAaaTacacsProviderDetailsAllOf) GetOwnerKey() string {
	if o == nil || o.OwnerKey == nil {
		var ret string
		return ret
	}
	return *o.OwnerKey
}

// GetOwnerKeyOk returns a tuple with the OwnerKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryAaaTacacsProviderDetailsAllOf) GetOwnerKeyOk() (*string, bool) {
	if o == nil || o.OwnerKey == nil {
		return nil, false
	}
	return o.OwnerKey, true
}

// HasOwnerKey returns a boolean if a field has been set.
func (o *NiatelemetryAaaTacacsProviderDetailsAllOf) HasOwnerKey() bool {
	if o != nil && o.OwnerKey != nil {
		return true
	}

	return false
}

// SetOwnerKey gets a reference to the given string and assigns it to the OwnerKey field.
func (o *NiatelemetryAaaTacacsProviderDetailsAllOf) SetOwnerKey(v string) {
	o.OwnerKey = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *NiatelemetryAaaTacacsProviderDetailsAllOf) GetRecordType() string {
	if o == nil || o.RecordType == nil {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryAaaTacacsProviderDetailsAllOf) GetRecordTypeOk() (*string, bool) {
	if o == nil || o.RecordType == nil {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *NiatelemetryAaaTacacsProviderDetailsAllOf) HasRecordType() bool {
	if o != nil && o.RecordType != nil {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *NiatelemetryAaaTacacsProviderDetailsAllOf) SetRecordType(v string) {
	o.RecordType = &v
}

// GetRecordVersion returns the RecordVersion field value if set, zero value otherwise.
func (o *NiatelemetryAaaTacacsProviderDetailsAllOf) GetRecordVersion() string {
	if o == nil || o.RecordVersion == nil {
		var ret string
		return ret
	}
	return *o.RecordVersion
}

// GetRecordVersionOk returns a tuple with the RecordVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryAaaTacacsProviderDetailsAllOf) GetRecordVersionOk() (*string, bool) {
	if o == nil || o.RecordVersion == nil {
		return nil, false
	}
	return o.RecordVersion, true
}

// HasRecordVersion returns a boolean if a field has been set.
func (o *NiatelemetryAaaTacacsProviderDetailsAllOf) HasRecordVersion() bool {
	if o != nil && o.RecordVersion != nil {
		return true
	}

	return false
}

// SetRecordVersion gets a reference to the given string and assigns it to the RecordVersion field.
func (o *NiatelemetryAaaTacacsProviderDetailsAllOf) SetRecordVersion(v string) {
	o.RecordVersion = &v
}

// GetSiteName returns the SiteName field value if set, zero value otherwise.
func (o *NiatelemetryAaaTacacsProviderDetailsAllOf) GetSiteName() string {
	if o == nil || o.SiteName == nil {
		var ret string
		return ret
	}
	return *o.SiteName
}

// GetSiteNameOk returns a tuple with the SiteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryAaaTacacsProviderDetailsAllOf) GetSiteNameOk() (*string, bool) {
	if o == nil || o.SiteName == nil {
		return nil, false
	}
	return o.SiteName, true
}

// HasSiteName returns a boolean if a field has been set.
func (o *NiatelemetryAaaTacacsProviderDetailsAllOf) HasSiteName() bool {
	if o != nil && o.SiteName != nil {
		return true
	}

	return false
}

// SetSiteName gets a reference to the given string and assigns it to the SiteName field.
func (o *NiatelemetryAaaTacacsProviderDetailsAllOf) SetSiteName(v string) {
	o.SiteName = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NiatelemetryAaaTacacsProviderDetailsAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryAaaTacacsProviderDetailsAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetryAaaTacacsProviderDetailsAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetryAaaTacacsProviderDetailsAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NiatelemetryAaaTacacsProviderDetailsAllOf) MarshalJSON() ([]byte, error) {
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
	if o.OwnerKey != nil {
		toSerialize["OwnerKey"] = o.OwnerKey
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

func (o *NiatelemetryAaaTacacsProviderDetailsAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varNiatelemetryAaaTacacsProviderDetailsAllOf := _NiatelemetryAaaTacacsProviderDetailsAllOf{}

	if err = json.Unmarshal(bytes, &varNiatelemetryAaaTacacsProviderDetailsAllOf); err == nil {
		*o = NiatelemetryAaaTacacsProviderDetailsAllOf(varNiatelemetryAaaTacacsProviderDetailsAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Dn")
		delete(additionalProperties, "OwnerKey")
		delete(additionalProperties, "RecordType")
		delete(additionalProperties, "RecordVersion")
		delete(additionalProperties, "SiteName")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNiatelemetryAaaTacacsProviderDetailsAllOf struct {
	value *NiatelemetryAaaTacacsProviderDetailsAllOf
	isSet bool
}

func (v NullableNiatelemetryAaaTacacsProviderDetailsAllOf) Get() *NiatelemetryAaaTacacsProviderDetailsAllOf {
	return v.value
}

func (v *NullableNiatelemetryAaaTacacsProviderDetailsAllOf) Set(val *NiatelemetryAaaTacacsProviderDetailsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryAaaTacacsProviderDetailsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryAaaTacacsProviderDetailsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryAaaTacacsProviderDetailsAllOf(val *NiatelemetryAaaTacacsProviderDetailsAllOf) *NullableNiatelemetryAaaTacacsProviderDetailsAllOf {
	return &NullableNiatelemetryAaaTacacsProviderDetailsAllOf{value: val, isSet: true}
}

func (v NullableNiatelemetryAaaTacacsProviderDetailsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryAaaTacacsProviderDetailsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
