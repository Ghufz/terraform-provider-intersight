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

// TamEolAdvisoryDetails Details pertaining to the milestone defined by an end-of-life (EOL) milestone advisory.
type TamEolAdvisoryDetails struct {
	TamBaseAdvisoryDetails
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType    string               `json:"ObjectType"`
	AllMilestones []TamMilestone       `json:"AllMilestones,omitempty"`
	Milestone     NullableTamMilestone `json:"Milestone,omitempty"`
	// The name of the impacted release this advisory milestone intends to address, e.g. \"3.5 (2x)\".
	Release              *string `json:"Release,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TamEolAdvisoryDetails TamEolAdvisoryDetails

// NewTamEolAdvisoryDetails instantiates a new TamEolAdvisoryDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTamEolAdvisoryDetails(classId string, objectType string) *TamEolAdvisoryDetails {
	this := TamEolAdvisoryDetails{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewTamEolAdvisoryDetailsWithDefaults instantiates a new TamEolAdvisoryDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTamEolAdvisoryDetailsWithDefaults() *TamEolAdvisoryDetails {
	this := TamEolAdvisoryDetails{}
	var classId string = "tam.EolAdvisoryDetails"
	this.ClassId = classId
	var objectType string = "tam.EolAdvisoryDetails"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *TamEolAdvisoryDetails) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *TamEolAdvisoryDetails) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *TamEolAdvisoryDetails) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *TamEolAdvisoryDetails) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *TamEolAdvisoryDetails) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *TamEolAdvisoryDetails) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAllMilestones returns the AllMilestones field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TamEolAdvisoryDetails) GetAllMilestones() []TamMilestone {
	if o == nil {
		var ret []TamMilestone
		return ret
	}
	return o.AllMilestones
}

// GetAllMilestonesOk returns a tuple with the AllMilestones field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TamEolAdvisoryDetails) GetAllMilestonesOk() ([]TamMilestone, bool) {
	if o == nil || o.AllMilestones == nil {
		return nil, false
	}
	return o.AllMilestones, true
}

// HasAllMilestones returns a boolean if a field has been set.
func (o *TamEolAdvisoryDetails) HasAllMilestones() bool {
	if o != nil && o.AllMilestones != nil {
		return true
	}

	return false
}

// SetAllMilestones gets a reference to the given []TamMilestone and assigns it to the AllMilestones field.
func (o *TamEolAdvisoryDetails) SetAllMilestones(v []TamMilestone) {
	o.AllMilestones = v
}

// GetMilestone returns the Milestone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TamEolAdvisoryDetails) GetMilestone() TamMilestone {
	if o == nil || o.Milestone.Get() == nil {
		var ret TamMilestone
		return ret
	}
	return *o.Milestone.Get()
}

// GetMilestoneOk returns a tuple with the Milestone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TamEolAdvisoryDetails) GetMilestoneOk() (*TamMilestone, bool) {
	if o == nil {
		return nil, false
	}
	return o.Milestone.Get(), o.Milestone.IsSet()
}

// HasMilestone returns a boolean if a field has been set.
func (o *TamEolAdvisoryDetails) HasMilestone() bool {
	if o != nil && o.Milestone.IsSet() {
		return true
	}

	return false
}

// SetMilestone gets a reference to the given NullableTamMilestone and assigns it to the Milestone field.
func (o *TamEolAdvisoryDetails) SetMilestone(v TamMilestone) {
	o.Milestone.Set(&v)
}

// SetMilestoneNil sets the value for Milestone to be an explicit nil
func (o *TamEolAdvisoryDetails) SetMilestoneNil() {
	o.Milestone.Set(nil)
}

// UnsetMilestone ensures that no value is present for Milestone, not even an explicit nil
func (o *TamEolAdvisoryDetails) UnsetMilestone() {
	o.Milestone.Unset()
}

// GetRelease returns the Release field value if set, zero value otherwise.
func (o *TamEolAdvisoryDetails) GetRelease() string {
	if o == nil || o.Release == nil {
		var ret string
		return ret
	}
	return *o.Release
}

// GetReleaseOk returns a tuple with the Release field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamEolAdvisoryDetails) GetReleaseOk() (*string, bool) {
	if o == nil || o.Release == nil {
		return nil, false
	}
	return o.Release, true
}

// HasRelease returns a boolean if a field has been set.
func (o *TamEolAdvisoryDetails) HasRelease() bool {
	if o != nil && o.Release != nil {
		return true
	}

	return false
}

// SetRelease gets a reference to the given string and assigns it to the Release field.
func (o *TamEolAdvisoryDetails) SetRelease(v string) {
	o.Release = &v
}

func (o TamEolAdvisoryDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedTamBaseAdvisoryDetails, errTamBaseAdvisoryDetails := json.Marshal(o.TamBaseAdvisoryDetails)
	if errTamBaseAdvisoryDetails != nil {
		return []byte{}, errTamBaseAdvisoryDetails
	}
	errTamBaseAdvisoryDetails = json.Unmarshal([]byte(serializedTamBaseAdvisoryDetails), &toSerialize)
	if errTamBaseAdvisoryDetails != nil {
		return []byte{}, errTamBaseAdvisoryDetails
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AllMilestones != nil {
		toSerialize["AllMilestones"] = o.AllMilestones
	}
	if o.Milestone.IsSet() {
		toSerialize["Milestone"] = o.Milestone.Get()
	}
	if o.Release != nil {
		toSerialize["Release"] = o.Release
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TamEolAdvisoryDetails) UnmarshalJSON(bytes []byte) (err error) {
	type TamEolAdvisoryDetailsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType    string               `json:"ObjectType"`
		AllMilestones []TamMilestone       `json:"AllMilestones,omitempty"`
		Milestone     NullableTamMilestone `json:"Milestone,omitempty"`
		// The name of the impacted release this advisory milestone intends to address, e.g. \"3.5 (2x)\".
		Release *string `json:"Release,omitempty"`
	}

	varTamEolAdvisoryDetailsWithoutEmbeddedStruct := TamEolAdvisoryDetailsWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varTamEolAdvisoryDetailsWithoutEmbeddedStruct)
	if err == nil {
		varTamEolAdvisoryDetails := _TamEolAdvisoryDetails{}
		varTamEolAdvisoryDetails.ClassId = varTamEolAdvisoryDetailsWithoutEmbeddedStruct.ClassId
		varTamEolAdvisoryDetails.ObjectType = varTamEolAdvisoryDetailsWithoutEmbeddedStruct.ObjectType
		varTamEolAdvisoryDetails.AllMilestones = varTamEolAdvisoryDetailsWithoutEmbeddedStruct.AllMilestones
		varTamEolAdvisoryDetails.Milestone = varTamEolAdvisoryDetailsWithoutEmbeddedStruct.Milestone
		varTamEolAdvisoryDetails.Release = varTamEolAdvisoryDetailsWithoutEmbeddedStruct.Release
		*o = TamEolAdvisoryDetails(varTamEolAdvisoryDetails)
	} else {
		return err
	}

	varTamEolAdvisoryDetails := _TamEolAdvisoryDetails{}

	err = json.Unmarshal(bytes, &varTamEolAdvisoryDetails)
	if err == nil {
		o.TamBaseAdvisoryDetails = varTamEolAdvisoryDetails.TamBaseAdvisoryDetails
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AllMilestones")
		delete(additionalProperties, "Milestone")
		delete(additionalProperties, "Release")

		// remove fields from embedded structs
		reflectTamBaseAdvisoryDetails := reflect.ValueOf(o.TamBaseAdvisoryDetails)
		for i := 0; i < reflectTamBaseAdvisoryDetails.Type().NumField(); i++ {
			t := reflectTamBaseAdvisoryDetails.Type().Field(i)

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

type NullableTamEolAdvisoryDetails struct {
	value *TamEolAdvisoryDetails
	isSet bool
}

func (v NullableTamEolAdvisoryDetails) Get() *TamEolAdvisoryDetails {
	return v.value
}

func (v *NullableTamEolAdvisoryDetails) Set(val *TamEolAdvisoryDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableTamEolAdvisoryDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableTamEolAdvisoryDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTamEolAdvisoryDetails(val *TamEolAdvisoryDetails) *NullableTamEolAdvisoryDetails {
	return &NullableTamEolAdvisoryDetails{value: val, isSet: true}
}

func (v NullableTamEolAdvisoryDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTamEolAdvisoryDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
