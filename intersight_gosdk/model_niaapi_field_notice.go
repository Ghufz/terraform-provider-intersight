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

// NiaapiFieldNotice This contains the new version release notice.
type NiaapiFieldNotice struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Bug Id associated with this notice.
	Bugid *string `json:"Bugid,omitempty"`
	// Field notice Description.
	FieldNoticeDesc *string `json:"FieldNoticeDesc,omitempty"`
	// Fieldnotice Id of this notice.
	FieldNoticeId *string `json:"FieldNoticeId,omitempty"`
	// Field notice URL link to the notice webpage.
	FieldNoticeUrl *string `json:"FieldNoticeUrl,omitempty"`
	// The headline of this field notice.
	Headline *string `json:"Headline,omitempty"`
	// Hardware PID for affected models.
	Hwpid        *string              `json:"Hwpid,omitempty"`
	RevisionInfo []NiaapiRevisionInfo `json:"RevisionInfo,omitempty"`
	// Software Release number for affected versions.
	SwRelease *string `json:"SwRelease,omitempty"`
	// URL of workaround of this notice.
	WorkaroundUrl        *string `json:"WorkaroundUrl,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiaapiFieldNotice NiaapiFieldNotice

// NewNiaapiFieldNotice instantiates a new NiaapiFieldNotice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiaapiFieldNotice(classId string, objectType string) *NiaapiFieldNotice {
	this := NiaapiFieldNotice{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiaapiFieldNoticeWithDefaults instantiates a new NiaapiFieldNotice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiaapiFieldNoticeWithDefaults() *NiaapiFieldNotice {
	this := NiaapiFieldNotice{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiaapiFieldNotice) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiaapiFieldNotice) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiaapiFieldNotice) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiaapiFieldNotice) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiaapiFieldNotice) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiaapiFieldNotice) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBugid returns the Bugid field value if set, zero value otherwise.
func (o *NiaapiFieldNotice) GetBugid() string {
	if o == nil || o.Bugid == nil {
		var ret string
		return ret
	}
	return *o.Bugid
}

// GetBugidOk returns a tuple with the Bugid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiFieldNotice) GetBugidOk() (*string, bool) {
	if o == nil || o.Bugid == nil {
		return nil, false
	}
	return o.Bugid, true
}

// HasBugid returns a boolean if a field has been set.
func (o *NiaapiFieldNotice) HasBugid() bool {
	if o != nil && o.Bugid != nil {
		return true
	}

	return false
}

// SetBugid gets a reference to the given string and assigns it to the Bugid field.
func (o *NiaapiFieldNotice) SetBugid(v string) {
	o.Bugid = &v
}

// GetFieldNoticeDesc returns the FieldNoticeDesc field value if set, zero value otherwise.
func (o *NiaapiFieldNotice) GetFieldNoticeDesc() string {
	if o == nil || o.FieldNoticeDesc == nil {
		var ret string
		return ret
	}
	return *o.FieldNoticeDesc
}

// GetFieldNoticeDescOk returns a tuple with the FieldNoticeDesc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiFieldNotice) GetFieldNoticeDescOk() (*string, bool) {
	if o == nil || o.FieldNoticeDesc == nil {
		return nil, false
	}
	return o.FieldNoticeDesc, true
}

// HasFieldNoticeDesc returns a boolean if a field has been set.
func (o *NiaapiFieldNotice) HasFieldNoticeDesc() bool {
	if o != nil && o.FieldNoticeDesc != nil {
		return true
	}

	return false
}

// SetFieldNoticeDesc gets a reference to the given string and assigns it to the FieldNoticeDesc field.
func (o *NiaapiFieldNotice) SetFieldNoticeDesc(v string) {
	o.FieldNoticeDesc = &v
}

// GetFieldNoticeId returns the FieldNoticeId field value if set, zero value otherwise.
func (o *NiaapiFieldNotice) GetFieldNoticeId() string {
	if o == nil || o.FieldNoticeId == nil {
		var ret string
		return ret
	}
	return *o.FieldNoticeId
}

// GetFieldNoticeIdOk returns a tuple with the FieldNoticeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiFieldNotice) GetFieldNoticeIdOk() (*string, bool) {
	if o == nil || o.FieldNoticeId == nil {
		return nil, false
	}
	return o.FieldNoticeId, true
}

// HasFieldNoticeId returns a boolean if a field has been set.
func (o *NiaapiFieldNotice) HasFieldNoticeId() bool {
	if o != nil && o.FieldNoticeId != nil {
		return true
	}

	return false
}

// SetFieldNoticeId gets a reference to the given string and assigns it to the FieldNoticeId field.
func (o *NiaapiFieldNotice) SetFieldNoticeId(v string) {
	o.FieldNoticeId = &v
}

// GetFieldNoticeUrl returns the FieldNoticeUrl field value if set, zero value otherwise.
func (o *NiaapiFieldNotice) GetFieldNoticeUrl() string {
	if o == nil || o.FieldNoticeUrl == nil {
		var ret string
		return ret
	}
	return *o.FieldNoticeUrl
}

// GetFieldNoticeUrlOk returns a tuple with the FieldNoticeUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiFieldNotice) GetFieldNoticeUrlOk() (*string, bool) {
	if o == nil || o.FieldNoticeUrl == nil {
		return nil, false
	}
	return o.FieldNoticeUrl, true
}

// HasFieldNoticeUrl returns a boolean if a field has been set.
func (o *NiaapiFieldNotice) HasFieldNoticeUrl() bool {
	if o != nil && o.FieldNoticeUrl != nil {
		return true
	}

	return false
}

// SetFieldNoticeUrl gets a reference to the given string and assigns it to the FieldNoticeUrl field.
func (o *NiaapiFieldNotice) SetFieldNoticeUrl(v string) {
	o.FieldNoticeUrl = &v
}

// GetHeadline returns the Headline field value if set, zero value otherwise.
func (o *NiaapiFieldNotice) GetHeadline() string {
	if o == nil || o.Headline == nil {
		var ret string
		return ret
	}
	return *o.Headline
}

// GetHeadlineOk returns a tuple with the Headline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiFieldNotice) GetHeadlineOk() (*string, bool) {
	if o == nil || o.Headline == nil {
		return nil, false
	}
	return o.Headline, true
}

// HasHeadline returns a boolean if a field has been set.
func (o *NiaapiFieldNotice) HasHeadline() bool {
	if o != nil && o.Headline != nil {
		return true
	}

	return false
}

// SetHeadline gets a reference to the given string and assigns it to the Headline field.
func (o *NiaapiFieldNotice) SetHeadline(v string) {
	o.Headline = &v
}

// GetHwpid returns the Hwpid field value if set, zero value otherwise.
func (o *NiaapiFieldNotice) GetHwpid() string {
	if o == nil || o.Hwpid == nil {
		var ret string
		return ret
	}
	return *o.Hwpid
}

// GetHwpidOk returns a tuple with the Hwpid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiFieldNotice) GetHwpidOk() (*string, bool) {
	if o == nil || o.Hwpid == nil {
		return nil, false
	}
	return o.Hwpid, true
}

// HasHwpid returns a boolean if a field has been set.
func (o *NiaapiFieldNotice) HasHwpid() bool {
	if o != nil && o.Hwpid != nil {
		return true
	}

	return false
}

// SetHwpid gets a reference to the given string and assigns it to the Hwpid field.
func (o *NiaapiFieldNotice) SetHwpid(v string) {
	o.Hwpid = &v
}

// GetRevisionInfo returns the RevisionInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NiaapiFieldNotice) GetRevisionInfo() []NiaapiRevisionInfo {
	if o == nil {
		var ret []NiaapiRevisionInfo
		return ret
	}
	return o.RevisionInfo
}

// GetRevisionInfoOk returns a tuple with the RevisionInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NiaapiFieldNotice) GetRevisionInfoOk() ([]NiaapiRevisionInfo, bool) {
	if o == nil || o.RevisionInfo == nil {
		return nil, false
	}
	return o.RevisionInfo, true
}

// HasRevisionInfo returns a boolean if a field has been set.
func (o *NiaapiFieldNotice) HasRevisionInfo() bool {
	if o != nil && o.RevisionInfo != nil {
		return true
	}

	return false
}

// SetRevisionInfo gets a reference to the given []NiaapiRevisionInfo and assigns it to the RevisionInfo field.
func (o *NiaapiFieldNotice) SetRevisionInfo(v []NiaapiRevisionInfo) {
	o.RevisionInfo = v
}

// GetSwRelease returns the SwRelease field value if set, zero value otherwise.
func (o *NiaapiFieldNotice) GetSwRelease() string {
	if o == nil || o.SwRelease == nil {
		var ret string
		return ret
	}
	return *o.SwRelease
}

// GetSwReleaseOk returns a tuple with the SwRelease field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiFieldNotice) GetSwReleaseOk() (*string, bool) {
	if o == nil || o.SwRelease == nil {
		return nil, false
	}
	return o.SwRelease, true
}

// HasSwRelease returns a boolean if a field has been set.
func (o *NiaapiFieldNotice) HasSwRelease() bool {
	if o != nil && o.SwRelease != nil {
		return true
	}

	return false
}

// SetSwRelease gets a reference to the given string and assigns it to the SwRelease field.
func (o *NiaapiFieldNotice) SetSwRelease(v string) {
	o.SwRelease = &v
}

// GetWorkaroundUrl returns the WorkaroundUrl field value if set, zero value otherwise.
func (o *NiaapiFieldNotice) GetWorkaroundUrl() string {
	if o == nil || o.WorkaroundUrl == nil {
		var ret string
		return ret
	}
	return *o.WorkaroundUrl
}

// GetWorkaroundUrlOk returns a tuple with the WorkaroundUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiFieldNotice) GetWorkaroundUrlOk() (*string, bool) {
	if o == nil || o.WorkaroundUrl == nil {
		return nil, false
	}
	return o.WorkaroundUrl, true
}

// HasWorkaroundUrl returns a boolean if a field has been set.
func (o *NiaapiFieldNotice) HasWorkaroundUrl() bool {
	if o != nil && o.WorkaroundUrl != nil {
		return true
	}

	return false
}

// SetWorkaroundUrl gets a reference to the given string and assigns it to the WorkaroundUrl field.
func (o *NiaapiFieldNotice) SetWorkaroundUrl(v string) {
	o.WorkaroundUrl = &v
}

func (o NiaapiFieldNotice) MarshalJSON() ([]byte, error) {
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
	if o.Bugid != nil {
		toSerialize["Bugid"] = o.Bugid
	}
	if o.FieldNoticeDesc != nil {
		toSerialize["FieldNoticeDesc"] = o.FieldNoticeDesc
	}
	if o.FieldNoticeId != nil {
		toSerialize["FieldNoticeId"] = o.FieldNoticeId
	}
	if o.FieldNoticeUrl != nil {
		toSerialize["FieldNoticeUrl"] = o.FieldNoticeUrl
	}
	if o.Headline != nil {
		toSerialize["Headline"] = o.Headline
	}
	if o.Hwpid != nil {
		toSerialize["Hwpid"] = o.Hwpid
	}
	if o.RevisionInfo != nil {
		toSerialize["RevisionInfo"] = o.RevisionInfo
	}
	if o.SwRelease != nil {
		toSerialize["SwRelease"] = o.SwRelease
	}
	if o.WorkaroundUrl != nil {
		toSerialize["WorkaroundUrl"] = o.WorkaroundUrl
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiaapiFieldNotice) UnmarshalJSON(bytes []byte) (err error) {
	type NiaapiFieldNoticeWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// Bug Id associated with this notice.
		Bugid *string `json:"Bugid,omitempty"`
		// Field notice Description.
		FieldNoticeDesc *string `json:"FieldNoticeDesc,omitempty"`
		// Fieldnotice Id of this notice.
		FieldNoticeId *string `json:"FieldNoticeId,omitempty"`
		// Field notice URL link to the notice webpage.
		FieldNoticeUrl *string `json:"FieldNoticeUrl,omitempty"`
		// The headline of this field notice.
		Headline *string `json:"Headline,omitempty"`
		// Hardware PID for affected models.
		Hwpid        *string              `json:"Hwpid,omitempty"`
		RevisionInfo []NiaapiRevisionInfo `json:"RevisionInfo,omitempty"`
		// Software Release number for affected versions.
		SwRelease *string `json:"SwRelease,omitempty"`
		// URL of workaround of this notice.
		WorkaroundUrl *string `json:"WorkaroundUrl,omitempty"`
	}

	varNiaapiFieldNoticeWithoutEmbeddedStruct := NiaapiFieldNoticeWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNiaapiFieldNoticeWithoutEmbeddedStruct)
	if err == nil {
		varNiaapiFieldNotice := _NiaapiFieldNotice{}
		varNiaapiFieldNotice.ClassId = varNiaapiFieldNoticeWithoutEmbeddedStruct.ClassId
		varNiaapiFieldNotice.ObjectType = varNiaapiFieldNoticeWithoutEmbeddedStruct.ObjectType
		varNiaapiFieldNotice.Bugid = varNiaapiFieldNoticeWithoutEmbeddedStruct.Bugid
		varNiaapiFieldNotice.FieldNoticeDesc = varNiaapiFieldNoticeWithoutEmbeddedStruct.FieldNoticeDesc
		varNiaapiFieldNotice.FieldNoticeId = varNiaapiFieldNoticeWithoutEmbeddedStruct.FieldNoticeId
		varNiaapiFieldNotice.FieldNoticeUrl = varNiaapiFieldNoticeWithoutEmbeddedStruct.FieldNoticeUrl
		varNiaapiFieldNotice.Headline = varNiaapiFieldNoticeWithoutEmbeddedStruct.Headline
		varNiaapiFieldNotice.Hwpid = varNiaapiFieldNoticeWithoutEmbeddedStruct.Hwpid
		varNiaapiFieldNotice.RevisionInfo = varNiaapiFieldNoticeWithoutEmbeddedStruct.RevisionInfo
		varNiaapiFieldNotice.SwRelease = varNiaapiFieldNoticeWithoutEmbeddedStruct.SwRelease
		varNiaapiFieldNotice.WorkaroundUrl = varNiaapiFieldNoticeWithoutEmbeddedStruct.WorkaroundUrl
		*o = NiaapiFieldNotice(varNiaapiFieldNotice)
	} else {
		return err
	}

	varNiaapiFieldNotice := _NiaapiFieldNotice{}

	err = json.Unmarshal(bytes, &varNiaapiFieldNotice)
	if err == nil {
		o.MoBaseMo = varNiaapiFieldNotice.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Bugid")
		delete(additionalProperties, "FieldNoticeDesc")
		delete(additionalProperties, "FieldNoticeId")
		delete(additionalProperties, "FieldNoticeUrl")
		delete(additionalProperties, "Headline")
		delete(additionalProperties, "Hwpid")
		delete(additionalProperties, "RevisionInfo")
		delete(additionalProperties, "SwRelease")
		delete(additionalProperties, "WorkaroundUrl")

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

type NullableNiaapiFieldNotice struct {
	value *NiaapiFieldNotice
	isSet bool
}

func (v NullableNiaapiFieldNotice) Get() *NiaapiFieldNotice {
	return v.value
}

func (v *NullableNiaapiFieldNotice) Set(val *NiaapiFieldNotice) {
	v.value = val
	v.isSet = true
}

func (v NullableNiaapiFieldNotice) IsSet() bool {
	return v.isSet
}

func (v *NullableNiaapiFieldNotice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiaapiFieldNotice(val *NiaapiFieldNotice) *NullableNiaapiFieldNotice {
	return &NullableNiaapiFieldNotice{value: val, isSet: true}
}

func (v NullableNiaapiFieldNotice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiaapiFieldNotice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
