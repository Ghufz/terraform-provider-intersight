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
)

// FirmwareModelBundleVersionAllOf Definition of the list of properties defined in 'firmware.ModelBundleVersion', excluding properties defined in parent classes.
type FirmwareModelBundleVersionAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The bundle version to which the server will be upgraded.
	BundleVersion *string `json:"BundleVersion,omitempty"`
	// The server family that will be impacted by this upgrade. * `UCSC-C220-M5` - The upgrade on all C220-M5 servers claimed in setup. * `UCSC-C220-M4` - The upgrade on all C220-M4 servers claimed in setup. * `UCSC-C240-M4` - The upgrade on all C240-M4 servers claimed in setup. * `UCSC-C460-M4` - The upgrade on all C460-M4 servers claimed in setup. * `UCSC-C240-M5` - The upgrade on all C240-M5 servers claimed in setup. * `UCSC-C480-M5` - The upgrade on all C480-M5 servers claimed in setup. * `UCSB-B200-M5` - The upgrade on all B200-M5 servers claimed in setup. * `UCSB-B480-M5` - The upgrade on all B480-M5 servers claimed in setup. * `UCSC-C220-M6` - The upgrade on all C220-M6 servers claimed in setup. * `UCSC-C240-M6` - The upgrade on all C240-M6 servers claimed in setup. * `UCSC-C225-M6` - The upgrade on all C225-M6 servers claimed in setup. * `UCSC-C245-M6` - The upgrade on all C245-M6 servers claimed in setup. * `UCSB-B200-M6` - The upgrade on all B200-M6 servers claimed in setup. * `UCSX-210C-M6` - The upgrade on all 210C-M6 servers claimed in setup. * `UCSX-210C-M7` - The upgrade on all 210C-M7 servers claimed in setup. * `UCSC-C220-M7` - The upgrade on all C220-M7 servers claimed in setup. * `UCSC-C240-M7` - The upgrade on all C240-M7 servers claimed in setup. * `UCSC-C125` - The upgrade on all C125 servers claimed in setup. * `UCSX-410C-M7` - The upgrade on all 410C-M7 servers claimed in setup.
	ModelFamily          *string `json:"ModelFamily,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FirmwareModelBundleVersionAllOf FirmwareModelBundleVersionAllOf

// NewFirmwareModelBundleVersionAllOf instantiates a new FirmwareModelBundleVersionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirmwareModelBundleVersionAllOf(classId string, objectType string) *FirmwareModelBundleVersionAllOf {
	this := FirmwareModelBundleVersionAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var modelFamily string = "UCSC-C220-M5"
	this.ModelFamily = &modelFamily
	return &this
}

// NewFirmwareModelBundleVersionAllOfWithDefaults instantiates a new FirmwareModelBundleVersionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirmwareModelBundleVersionAllOfWithDefaults() *FirmwareModelBundleVersionAllOf {
	this := FirmwareModelBundleVersionAllOf{}
	var classId string = "firmware.ModelBundleVersion"
	this.ClassId = classId
	var objectType string = "firmware.ModelBundleVersion"
	this.ObjectType = objectType
	var modelFamily string = "UCSC-C220-M5"
	this.ModelFamily = &modelFamily
	return &this
}

// GetClassId returns the ClassId field value
func (o *FirmwareModelBundleVersionAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FirmwareModelBundleVersionAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FirmwareModelBundleVersionAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FirmwareModelBundleVersionAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FirmwareModelBundleVersionAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FirmwareModelBundleVersionAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBundleVersion returns the BundleVersion field value if set, zero value otherwise.
func (o *FirmwareModelBundleVersionAllOf) GetBundleVersion() string {
	if o == nil || o.BundleVersion == nil {
		var ret string
		return ret
	}
	return *o.BundleVersion
}

// GetBundleVersionOk returns a tuple with the BundleVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareModelBundleVersionAllOf) GetBundleVersionOk() (*string, bool) {
	if o == nil || o.BundleVersion == nil {
		return nil, false
	}
	return o.BundleVersion, true
}

// HasBundleVersion returns a boolean if a field has been set.
func (o *FirmwareModelBundleVersionAllOf) HasBundleVersion() bool {
	if o != nil && o.BundleVersion != nil {
		return true
	}

	return false
}

// SetBundleVersion gets a reference to the given string and assigns it to the BundleVersion field.
func (o *FirmwareModelBundleVersionAllOf) SetBundleVersion(v string) {
	o.BundleVersion = &v
}

// GetModelFamily returns the ModelFamily field value if set, zero value otherwise.
func (o *FirmwareModelBundleVersionAllOf) GetModelFamily() string {
	if o == nil || o.ModelFamily == nil {
		var ret string
		return ret
	}
	return *o.ModelFamily
}

// GetModelFamilyOk returns a tuple with the ModelFamily field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareModelBundleVersionAllOf) GetModelFamilyOk() (*string, bool) {
	if o == nil || o.ModelFamily == nil {
		return nil, false
	}
	return o.ModelFamily, true
}

// HasModelFamily returns a boolean if a field has been set.
func (o *FirmwareModelBundleVersionAllOf) HasModelFamily() bool {
	if o != nil && o.ModelFamily != nil {
		return true
	}

	return false
}

// SetModelFamily gets a reference to the given string and assigns it to the ModelFamily field.
func (o *FirmwareModelBundleVersionAllOf) SetModelFamily(v string) {
	o.ModelFamily = &v
}

func (o FirmwareModelBundleVersionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.BundleVersion != nil {
		toSerialize["BundleVersion"] = o.BundleVersion
	}
	if o.ModelFamily != nil {
		toSerialize["ModelFamily"] = o.ModelFamily
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FirmwareModelBundleVersionAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varFirmwareModelBundleVersionAllOf := _FirmwareModelBundleVersionAllOf{}

	if err = json.Unmarshal(bytes, &varFirmwareModelBundleVersionAllOf); err == nil {
		*o = FirmwareModelBundleVersionAllOf(varFirmwareModelBundleVersionAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BundleVersion")
		delete(additionalProperties, "ModelFamily")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFirmwareModelBundleVersionAllOf struct {
	value *FirmwareModelBundleVersionAllOf
	isSet bool
}

func (v NullableFirmwareModelBundleVersionAllOf) Get() *FirmwareModelBundleVersionAllOf {
	return v.value
}

func (v *NullableFirmwareModelBundleVersionAllOf) Set(val *FirmwareModelBundleVersionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFirmwareModelBundleVersionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFirmwareModelBundleVersionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirmwareModelBundleVersionAllOf(val *FirmwareModelBundleVersionAllOf) *NullableFirmwareModelBundleVersionAllOf {
	return &NullableFirmwareModelBundleVersionAllOf{value: val, isSet: true}
}

func (v NullableFirmwareModelBundleVersionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirmwareModelBundleVersionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
