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

// LicenseSubstituteLicense License tiers can be substituted with higher tiers. Will contain information on the type of license being used for substitution, the number of substutions used, and the substitution type.
type LicenseSubstituteLicense struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The substitute license that is used.
	SubstitutedLicense *string `json:"SubstitutedLicense,omitempty"`
	// The number of substitute licenses that are used.
	SubstitutedQuantity *int64 `json:"SubstitutedQuantity,omitempty"`
	// The substitution from lower or from higher tier.
	SubstitutionType     *string `json:"SubstitutionType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LicenseSubstituteLicense LicenseSubstituteLicense

// NewLicenseSubstituteLicense instantiates a new LicenseSubstituteLicense object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicenseSubstituteLicense(classId string, objectType string) *LicenseSubstituteLicense {
	this := LicenseSubstituteLicense{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewLicenseSubstituteLicenseWithDefaults instantiates a new LicenseSubstituteLicense object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicenseSubstituteLicenseWithDefaults() *LicenseSubstituteLicense {
	this := LicenseSubstituteLicense{}
	var classId string = "license.SubstituteLicense"
	this.ClassId = classId
	var objectType string = "license.SubstituteLicense"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *LicenseSubstituteLicense) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *LicenseSubstituteLicense) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *LicenseSubstituteLicense) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *LicenseSubstituteLicense) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *LicenseSubstituteLicense) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *LicenseSubstituteLicense) SetObjectType(v string) {
	o.ObjectType = v
}

// GetSubstitutedLicense returns the SubstitutedLicense field value if set, zero value otherwise.
func (o *LicenseSubstituteLicense) GetSubstitutedLicense() string {
	if o == nil || o.SubstitutedLicense == nil {
		var ret string
		return ret
	}
	return *o.SubstitutedLicense
}

// GetSubstitutedLicenseOk returns a tuple with the SubstitutedLicense field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseSubstituteLicense) GetSubstitutedLicenseOk() (*string, bool) {
	if o == nil || o.SubstitutedLicense == nil {
		return nil, false
	}
	return o.SubstitutedLicense, true
}

// HasSubstitutedLicense returns a boolean if a field has been set.
func (o *LicenseSubstituteLicense) HasSubstitutedLicense() bool {
	if o != nil && o.SubstitutedLicense != nil {
		return true
	}

	return false
}

// SetSubstitutedLicense gets a reference to the given string and assigns it to the SubstitutedLicense field.
func (o *LicenseSubstituteLicense) SetSubstitutedLicense(v string) {
	o.SubstitutedLicense = &v
}

// GetSubstitutedQuantity returns the SubstitutedQuantity field value if set, zero value otherwise.
func (o *LicenseSubstituteLicense) GetSubstitutedQuantity() int64 {
	if o == nil || o.SubstitutedQuantity == nil {
		var ret int64
		return ret
	}
	return *o.SubstitutedQuantity
}

// GetSubstitutedQuantityOk returns a tuple with the SubstitutedQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseSubstituteLicense) GetSubstitutedQuantityOk() (*int64, bool) {
	if o == nil || o.SubstitutedQuantity == nil {
		return nil, false
	}
	return o.SubstitutedQuantity, true
}

// HasSubstitutedQuantity returns a boolean if a field has been set.
func (o *LicenseSubstituteLicense) HasSubstitutedQuantity() bool {
	if o != nil && o.SubstitutedQuantity != nil {
		return true
	}

	return false
}

// SetSubstitutedQuantity gets a reference to the given int64 and assigns it to the SubstitutedQuantity field.
func (o *LicenseSubstituteLicense) SetSubstitutedQuantity(v int64) {
	o.SubstitutedQuantity = &v
}

// GetSubstitutionType returns the SubstitutionType field value if set, zero value otherwise.
func (o *LicenseSubstituteLicense) GetSubstitutionType() string {
	if o == nil || o.SubstitutionType == nil {
		var ret string
		return ret
	}
	return *o.SubstitutionType
}

// GetSubstitutionTypeOk returns a tuple with the SubstitutionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseSubstituteLicense) GetSubstitutionTypeOk() (*string, bool) {
	if o == nil || o.SubstitutionType == nil {
		return nil, false
	}
	return o.SubstitutionType, true
}

// HasSubstitutionType returns a boolean if a field has been set.
func (o *LicenseSubstituteLicense) HasSubstitutionType() bool {
	if o != nil && o.SubstitutionType != nil {
		return true
	}

	return false
}

// SetSubstitutionType gets a reference to the given string and assigns it to the SubstitutionType field.
func (o *LicenseSubstituteLicense) SetSubstitutionType(v string) {
	o.SubstitutionType = &v
}

func (o LicenseSubstituteLicense) MarshalJSON() ([]byte, error) {
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
	if o.SubstitutedLicense != nil {
		toSerialize["SubstitutedLicense"] = o.SubstitutedLicense
	}
	if o.SubstitutedQuantity != nil {
		toSerialize["SubstitutedQuantity"] = o.SubstitutedQuantity
	}
	if o.SubstitutionType != nil {
		toSerialize["SubstitutionType"] = o.SubstitutionType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LicenseSubstituteLicense) UnmarshalJSON(bytes []byte) (err error) {
	type LicenseSubstituteLicenseWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The substitute license that is used.
		SubstitutedLicense *string `json:"SubstitutedLicense,omitempty"`
		// The number of substitute licenses that are used.
		SubstitutedQuantity *int64 `json:"SubstitutedQuantity,omitempty"`
		// The substitution from lower or from higher tier.
		SubstitutionType *string `json:"SubstitutionType,omitempty"`
	}

	varLicenseSubstituteLicenseWithoutEmbeddedStruct := LicenseSubstituteLicenseWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varLicenseSubstituteLicenseWithoutEmbeddedStruct)
	if err == nil {
		varLicenseSubstituteLicense := _LicenseSubstituteLicense{}
		varLicenseSubstituteLicense.ClassId = varLicenseSubstituteLicenseWithoutEmbeddedStruct.ClassId
		varLicenseSubstituteLicense.ObjectType = varLicenseSubstituteLicenseWithoutEmbeddedStruct.ObjectType
		varLicenseSubstituteLicense.SubstitutedLicense = varLicenseSubstituteLicenseWithoutEmbeddedStruct.SubstitutedLicense
		varLicenseSubstituteLicense.SubstitutedQuantity = varLicenseSubstituteLicenseWithoutEmbeddedStruct.SubstitutedQuantity
		varLicenseSubstituteLicense.SubstitutionType = varLicenseSubstituteLicenseWithoutEmbeddedStruct.SubstitutionType
		*o = LicenseSubstituteLicense(varLicenseSubstituteLicense)
	} else {
		return err
	}

	varLicenseSubstituteLicense := _LicenseSubstituteLicense{}

	err = json.Unmarshal(bytes, &varLicenseSubstituteLicense)
	if err == nil {
		o.MoBaseComplexType = varLicenseSubstituteLicense.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "SubstitutedLicense")
		delete(additionalProperties, "SubstitutedQuantity")
		delete(additionalProperties, "SubstitutionType")

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

type NullableLicenseSubstituteLicense struct {
	value *LicenseSubstituteLicense
	isSet bool
}

func (v NullableLicenseSubstituteLicense) Get() *LicenseSubstituteLicense {
	return v.value
}

func (v *NullableLicenseSubstituteLicense) Set(val *LicenseSubstituteLicense) {
	v.value = val
	v.isSet = true
}

func (v NullableLicenseSubstituteLicense) IsSet() bool {
	return v.isSet
}

func (v *NullableLicenseSubstituteLicense) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicenseSubstituteLicense(val *LicenseSubstituteLicense) *NullableLicenseSubstituteLicense {
	return &NullableLicenseSubstituteLicense{value: val, isSet: true}
}

func (v NullableLicenseSubstituteLicense) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicenseSubstituteLicense) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
