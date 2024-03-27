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

// NiaapiNewReleaseDetail The detail content of of this post.
type NiaapiNewReleaseDetail struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Description of this new verison release post.
	Description *string `json:"Description,omitempty"`
	// Link of downloading the release file.
	Link *string `json:"Link,omitempty"`
	// The link used to provide a gateway for customer to review the release note.
	ReleaseNoteLink *string `json:"ReleaseNoteLink,omitempty"`
	// The link title used to provide a gateway for customer to review the release note.
	ReleaseNoteLinkTitle *string `json:"ReleaseNoteLinkTitle,omitempty"`
	// The link used to provide a gateway for customer to download the release.
	SoftwareDownloadLink *string `json:"SoftwareDownloadLink,omitempty"`
	// The link title used to provide a gateway for customer to download the release.
	SoftwareDownloadLinkTitle *string `json:"SoftwareDownloadLinkTitle,omitempty"`
	// Title of the new verison release post.
	Title *string `json:"Title,omitempty"`
	// Version number Associate with this Post.
	Version              *string `json:"Version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiaapiNewReleaseDetail NiaapiNewReleaseDetail

// NewNiaapiNewReleaseDetail instantiates a new NiaapiNewReleaseDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiaapiNewReleaseDetail(classId string, objectType string) *NiaapiNewReleaseDetail {
	this := NiaapiNewReleaseDetail{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiaapiNewReleaseDetailWithDefaults instantiates a new NiaapiNewReleaseDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiaapiNewReleaseDetailWithDefaults() *NiaapiNewReleaseDetail {
	this := NiaapiNewReleaseDetail{}
	var classId string = "niaapi.NewReleaseDetail"
	this.ClassId = classId
	var objectType string = "niaapi.NewReleaseDetail"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiaapiNewReleaseDetail) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiaapiNewReleaseDetail) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiaapiNewReleaseDetail) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiaapiNewReleaseDetail) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiaapiNewReleaseDetail) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiaapiNewReleaseDetail) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *NiaapiNewReleaseDetail) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiNewReleaseDetail) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *NiaapiNewReleaseDetail) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *NiaapiNewReleaseDetail) SetDescription(v string) {
	o.Description = &v
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *NiaapiNewReleaseDetail) GetLink() string {
	if o == nil || o.Link == nil {
		var ret string
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiNewReleaseDetail) GetLinkOk() (*string, bool) {
	if o == nil || o.Link == nil {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *NiaapiNewReleaseDetail) HasLink() bool {
	if o != nil && o.Link != nil {
		return true
	}

	return false
}

// SetLink gets a reference to the given string and assigns it to the Link field.
func (o *NiaapiNewReleaseDetail) SetLink(v string) {
	o.Link = &v
}

// GetReleaseNoteLink returns the ReleaseNoteLink field value if set, zero value otherwise.
func (o *NiaapiNewReleaseDetail) GetReleaseNoteLink() string {
	if o == nil || o.ReleaseNoteLink == nil {
		var ret string
		return ret
	}
	return *o.ReleaseNoteLink
}

// GetReleaseNoteLinkOk returns a tuple with the ReleaseNoteLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiNewReleaseDetail) GetReleaseNoteLinkOk() (*string, bool) {
	if o == nil || o.ReleaseNoteLink == nil {
		return nil, false
	}
	return o.ReleaseNoteLink, true
}

// HasReleaseNoteLink returns a boolean if a field has been set.
func (o *NiaapiNewReleaseDetail) HasReleaseNoteLink() bool {
	if o != nil && o.ReleaseNoteLink != nil {
		return true
	}

	return false
}

// SetReleaseNoteLink gets a reference to the given string and assigns it to the ReleaseNoteLink field.
func (o *NiaapiNewReleaseDetail) SetReleaseNoteLink(v string) {
	o.ReleaseNoteLink = &v
}

// GetReleaseNoteLinkTitle returns the ReleaseNoteLinkTitle field value if set, zero value otherwise.
func (o *NiaapiNewReleaseDetail) GetReleaseNoteLinkTitle() string {
	if o == nil || o.ReleaseNoteLinkTitle == nil {
		var ret string
		return ret
	}
	return *o.ReleaseNoteLinkTitle
}

// GetReleaseNoteLinkTitleOk returns a tuple with the ReleaseNoteLinkTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiNewReleaseDetail) GetReleaseNoteLinkTitleOk() (*string, bool) {
	if o == nil || o.ReleaseNoteLinkTitle == nil {
		return nil, false
	}
	return o.ReleaseNoteLinkTitle, true
}

// HasReleaseNoteLinkTitle returns a boolean if a field has been set.
func (o *NiaapiNewReleaseDetail) HasReleaseNoteLinkTitle() bool {
	if o != nil && o.ReleaseNoteLinkTitle != nil {
		return true
	}

	return false
}

// SetReleaseNoteLinkTitle gets a reference to the given string and assigns it to the ReleaseNoteLinkTitle field.
func (o *NiaapiNewReleaseDetail) SetReleaseNoteLinkTitle(v string) {
	o.ReleaseNoteLinkTitle = &v
}

// GetSoftwareDownloadLink returns the SoftwareDownloadLink field value if set, zero value otherwise.
func (o *NiaapiNewReleaseDetail) GetSoftwareDownloadLink() string {
	if o == nil || o.SoftwareDownloadLink == nil {
		var ret string
		return ret
	}
	return *o.SoftwareDownloadLink
}

// GetSoftwareDownloadLinkOk returns a tuple with the SoftwareDownloadLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiNewReleaseDetail) GetSoftwareDownloadLinkOk() (*string, bool) {
	if o == nil || o.SoftwareDownloadLink == nil {
		return nil, false
	}
	return o.SoftwareDownloadLink, true
}

// HasSoftwareDownloadLink returns a boolean if a field has been set.
func (o *NiaapiNewReleaseDetail) HasSoftwareDownloadLink() bool {
	if o != nil && o.SoftwareDownloadLink != nil {
		return true
	}

	return false
}

// SetSoftwareDownloadLink gets a reference to the given string and assigns it to the SoftwareDownloadLink field.
func (o *NiaapiNewReleaseDetail) SetSoftwareDownloadLink(v string) {
	o.SoftwareDownloadLink = &v
}

// GetSoftwareDownloadLinkTitle returns the SoftwareDownloadLinkTitle field value if set, zero value otherwise.
func (o *NiaapiNewReleaseDetail) GetSoftwareDownloadLinkTitle() string {
	if o == nil || o.SoftwareDownloadLinkTitle == nil {
		var ret string
		return ret
	}
	return *o.SoftwareDownloadLinkTitle
}

// GetSoftwareDownloadLinkTitleOk returns a tuple with the SoftwareDownloadLinkTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiNewReleaseDetail) GetSoftwareDownloadLinkTitleOk() (*string, bool) {
	if o == nil || o.SoftwareDownloadLinkTitle == nil {
		return nil, false
	}
	return o.SoftwareDownloadLinkTitle, true
}

// HasSoftwareDownloadLinkTitle returns a boolean if a field has been set.
func (o *NiaapiNewReleaseDetail) HasSoftwareDownloadLinkTitle() bool {
	if o != nil && o.SoftwareDownloadLinkTitle != nil {
		return true
	}

	return false
}

// SetSoftwareDownloadLinkTitle gets a reference to the given string and assigns it to the SoftwareDownloadLinkTitle field.
func (o *NiaapiNewReleaseDetail) SetSoftwareDownloadLinkTitle(v string) {
	o.SoftwareDownloadLinkTitle = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *NiaapiNewReleaseDetail) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiNewReleaseDetail) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *NiaapiNewReleaseDetail) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *NiaapiNewReleaseDetail) SetTitle(v string) {
	o.Title = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *NiaapiNewReleaseDetail) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiNewReleaseDetail) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *NiaapiNewReleaseDetail) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *NiaapiNewReleaseDetail) SetVersion(v string) {
	o.Version = &v
}

func (o NiaapiNewReleaseDetail) MarshalJSON() ([]byte, error) {
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
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.Link != nil {
		toSerialize["Link"] = o.Link
	}
	if o.ReleaseNoteLink != nil {
		toSerialize["ReleaseNoteLink"] = o.ReleaseNoteLink
	}
	if o.ReleaseNoteLinkTitle != nil {
		toSerialize["ReleaseNoteLinkTitle"] = o.ReleaseNoteLinkTitle
	}
	if o.SoftwareDownloadLink != nil {
		toSerialize["SoftwareDownloadLink"] = o.SoftwareDownloadLink
	}
	if o.SoftwareDownloadLinkTitle != nil {
		toSerialize["SoftwareDownloadLinkTitle"] = o.SoftwareDownloadLinkTitle
	}
	if o.Title != nil {
		toSerialize["Title"] = o.Title
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiaapiNewReleaseDetail) UnmarshalJSON(bytes []byte) (err error) {
	type NiaapiNewReleaseDetailWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Description of this new verison release post.
		Description *string `json:"Description,omitempty"`
		// Link of downloading the release file.
		Link *string `json:"Link,omitempty"`
		// The link used to provide a gateway for customer to review the release note.
		ReleaseNoteLink *string `json:"ReleaseNoteLink,omitempty"`
		// The link title used to provide a gateway for customer to review the release note.
		ReleaseNoteLinkTitle *string `json:"ReleaseNoteLinkTitle,omitempty"`
		// The link used to provide a gateway for customer to download the release.
		SoftwareDownloadLink *string `json:"SoftwareDownloadLink,omitempty"`
		// The link title used to provide a gateway for customer to download the release.
		SoftwareDownloadLinkTitle *string `json:"SoftwareDownloadLinkTitle,omitempty"`
		// Title of the new verison release post.
		Title *string `json:"Title,omitempty"`
		// Version number Associate with this Post.
		Version *string `json:"Version,omitempty"`
	}

	varNiaapiNewReleaseDetailWithoutEmbeddedStruct := NiaapiNewReleaseDetailWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNiaapiNewReleaseDetailWithoutEmbeddedStruct)
	if err == nil {
		varNiaapiNewReleaseDetail := _NiaapiNewReleaseDetail{}
		varNiaapiNewReleaseDetail.ClassId = varNiaapiNewReleaseDetailWithoutEmbeddedStruct.ClassId
		varNiaapiNewReleaseDetail.ObjectType = varNiaapiNewReleaseDetailWithoutEmbeddedStruct.ObjectType
		varNiaapiNewReleaseDetail.Description = varNiaapiNewReleaseDetailWithoutEmbeddedStruct.Description
		varNiaapiNewReleaseDetail.Link = varNiaapiNewReleaseDetailWithoutEmbeddedStruct.Link
		varNiaapiNewReleaseDetail.ReleaseNoteLink = varNiaapiNewReleaseDetailWithoutEmbeddedStruct.ReleaseNoteLink
		varNiaapiNewReleaseDetail.ReleaseNoteLinkTitle = varNiaapiNewReleaseDetailWithoutEmbeddedStruct.ReleaseNoteLinkTitle
		varNiaapiNewReleaseDetail.SoftwareDownloadLink = varNiaapiNewReleaseDetailWithoutEmbeddedStruct.SoftwareDownloadLink
		varNiaapiNewReleaseDetail.SoftwareDownloadLinkTitle = varNiaapiNewReleaseDetailWithoutEmbeddedStruct.SoftwareDownloadLinkTitle
		varNiaapiNewReleaseDetail.Title = varNiaapiNewReleaseDetailWithoutEmbeddedStruct.Title
		varNiaapiNewReleaseDetail.Version = varNiaapiNewReleaseDetailWithoutEmbeddedStruct.Version
		*o = NiaapiNewReleaseDetail(varNiaapiNewReleaseDetail)
	} else {
		return err
	}

	varNiaapiNewReleaseDetail := _NiaapiNewReleaseDetail{}

	err = json.Unmarshal(bytes, &varNiaapiNewReleaseDetail)
	if err == nil {
		o.MoBaseComplexType = varNiaapiNewReleaseDetail.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "Link")
		delete(additionalProperties, "ReleaseNoteLink")
		delete(additionalProperties, "ReleaseNoteLinkTitle")
		delete(additionalProperties, "SoftwareDownloadLink")
		delete(additionalProperties, "SoftwareDownloadLinkTitle")
		delete(additionalProperties, "Title")
		delete(additionalProperties, "Version")

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

type NullableNiaapiNewReleaseDetail struct {
	value *NiaapiNewReleaseDetail
	isSet bool
}

func (v NullableNiaapiNewReleaseDetail) Get() *NiaapiNewReleaseDetail {
	return v.value
}

func (v *NullableNiaapiNewReleaseDetail) Set(val *NiaapiNewReleaseDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableNiaapiNewReleaseDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableNiaapiNewReleaseDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiaapiNewReleaseDetail(val *NiaapiNewReleaseDetail) *NullableNiaapiNewReleaseDetail {
	return &NullableNiaapiNewReleaseDetail{value: val, isSet: true}
}

func (v NullableNiaapiNewReleaseDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiaapiNewReleaseDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
