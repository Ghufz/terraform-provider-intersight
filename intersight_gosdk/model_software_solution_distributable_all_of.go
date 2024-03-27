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

// SoftwareSolutionDistributableAllOf Definition of the list of properties defined in 'software.SolutionDistributable', excluding properties defined in parent classes.
type SoftwareSolutionDistributableAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The path of the file in S3/minio bucket.
	FilePath *string `json:"FilePath,omitempty"`
	// The name of the solution in which the image belongs.
	SolutionName *string `json:"SolutionName,omitempty"`
	// The type of the file like OS image, Script etc. * `osimage` - The solution OS image for deployment. * `script` - The Python script for the solution VM configuration and deployment.
	SubType              *string                                `json:"SubType,omitempty"`
	Catalog              *SoftwarerepositoryCatalogRelationship `json:"Catalog,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SoftwareSolutionDistributableAllOf SoftwareSolutionDistributableAllOf

// NewSoftwareSolutionDistributableAllOf instantiates a new SoftwareSolutionDistributableAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSoftwareSolutionDistributableAllOf(classId string, objectType string) *SoftwareSolutionDistributableAllOf {
	this := SoftwareSolutionDistributableAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var subType string = "osimage"
	this.SubType = &subType
	return &this
}

// NewSoftwareSolutionDistributableAllOfWithDefaults instantiates a new SoftwareSolutionDistributableAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSoftwareSolutionDistributableAllOfWithDefaults() *SoftwareSolutionDistributableAllOf {
	this := SoftwareSolutionDistributableAllOf{}
	var classId string = "software.SolutionDistributable"
	this.ClassId = classId
	var objectType string = "software.SolutionDistributable"
	this.ObjectType = objectType
	var subType string = "osimage"
	this.SubType = &subType
	return &this
}

// GetClassId returns the ClassId field value
func (o *SoftwareSolutionDistributableAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SoftwareSolutionDistributableAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SoftwareSolutionDistributableAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *SoftwareSolutionDistributableAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SoftwareSolutionDistributableAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SoftwareSolutionDistributableAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetFilePath returns the FilePath field value if set, zero value otherwise.
func (o *SoftwareSolutionDistributableAllOf) GetFilePath() string {
	if o == nil || o.FilePath == nil {
		var ret string
		return ret
	}
	return *o.FilePath
}

// GetFilePathOk returns a tuple with the FilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareSolutionDistributableAllOf) GetFilePathOk() (*string, bool) {
	if o == nil || o.FilePath == nil {
		return nil, false
	}
	return o.FilePath, true
}

// HasFilePath returns a boolean if a field has been set.
func (o *SoftwareSolutionDistributableAllOf) HasFilePath() bool {
	if o != nil && o.FilePath != nil {
		return true
	}

	return false
}

// SetFilePath gets a reference to the given string and assigns it to the FilePath field.
func (o *SoftwareSolutionDistributableAllOf) SetFilePath(v string) {
	o.FilePath = &v
}

// GetSolutionName returns the SolutionName field value if set, zero value otherwise.
func (o *SoftwareSolutionDistributableAllOf) GetSolutionName() string {
	if o == nil || o.SolutionName == nil {
		var ret string
		return ret
	}
	return *o.SolutionName
}

// GetSolutionNameOk returns a tuple with the SolutionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareSolutionDistributableAllOf) GetSolutionNameOk() (*string, bool) {
	if o == nil || o.SolutionName == nil {
		return nil, false
	}
	return o.SolutionName, true
}

// HasSolutionName returns a boolean if a field has been set.
func (o *SoftwareSolutionDistributableAllOf) HasSolutionName() bool {
	if o != nil && o.SolutionName != nil {
		return true
	}

	return false
}

// SetSolutionName gets a reference to the given string and assigns it to the SolutionName field.
func (o *SoftwareSolutionDistributableAllOf) SetSolutionName(v string) {
	o.SolutionName = &v
}

// GetSubType returns the SubType field value if set, zero value otherwise.
func (o *SoftwareSolutionDistributableAllOf) GetSubType() string {
	if o == nil || o.SubType == nil {
		var ret string
		return ret
	}
	return *o.SubType
}

// GetSubTypeOk returns a tuple with the SubType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareSolutionDistributableAllOf) GetSubTypeOk() (*string, bool) {
	if o == nil || o.SubType == nil {
		return nil, false
	}
	return o.SubType, true
}

// HasSubType returns a boolean if a field has been set.
func (o *SoftwareSolutionDistributableAllOf) HasSubType() bool {
	if o != nil && o.SubType != nil {
		return true
	}

	return false
}

// SetSubType gets a reference to the given string and assigns it to the SubType field.
func (o *SoftwareSolutionDistributableAllOf) SetSubType(v string) {
	o.SubType = &v
}

// GetCatalog returns the Catalog field value if set, zero value otherwise.
func (o *SoftwareSolutionDistributableAllOf) GetCatalog() SoftwarerepositoryCatalogRelationship {
	if o == nil || o.Catalog == nil {
		var ret SoftwarerepositoryCatalogRelationship
		return ret
	}
	return *o.Catalog
}

// GetCatalogOk returns a tuple with the Catalog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareSolutionDistributableAllOf) GetCatalogOk() (*SoftwarerepositoryCatalogRelationship, bool) {
	if o == nil || o.Catalog == nil {
		return nil, false
	}
	return o.Catalog, true
}

// HasCatalog returns a boolean if a field has been set.
func (o *SoftwareSolutionDistributableAllOf) HasCatalog() bool {
	if o != nil && o.Catalog != nil {
		return true
	}

	return false
}

// SetCatalog gets a reference to the given SoftwarerepositoryCatalogRelationship and assigns it to the Catalog field.
func (o *SoftwareSolutionDistributableAllOf) SetCatalog(v SoftwarerepositoryCatalogRelationship) {
	o.Catalog = &v
}

func (o SoftwareSolutionDistributableAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.FilePath != nil {
		toSerialize["FilePath"] = o.FilePath
	}
	if o.SolutionName != nil {
		toSerialize["SolutionName"] = o.SolutionName
	}
	if o.SubType != nil {
		toSerialize["SubType"] = o.SubType
	}
	if o.Catalog != nil {
		toSerialize["Catalog"] = o.Catalog
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SoftwareSolutionDistributableAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varSoftwareSolutionDistributableAllOf := _SoftwareSolutionDistributableAllOf{}

	if err = json.Unmarshal(bytes, &varSoftwareSolutionDistributableAllOf); err == nil {
		*o = SoftwareSolutionDistributableAllOf(varSoftwareSolutionDistributableAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "FilePath")
		delete(additionalProperties, "SolutionName")
		delete(additionalProperties, "SubType")
		delete(additionalProperties, "Catalog")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSoftwareSolutionDistributableAllOf struct {
	value *SoftwareSolutionDistributableAllOf
	isSet bool
}

func (v NullableSoftwareSolutionDistributableAllOf) Get() *SoftwareSolutionDistributableAllOf {
	return v.value
}

func (v *NullableSoftwareSolutionDistributableAllOf) Set(val *SoftwareSolutionDistributableAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSoftwareSolutionDistributableAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSoftwareSolutionDistributableAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSoftwareSolutionDistributableAllOf(val *SoftwareSolutionDistributableAllOf) *NullableSoftwareSolutionDistributableAllOf {
	return &NullableSoftwareSolutionDistributableAllOf{value: val, isSet: true}
}

func (v NullableSoftwareSolutionDistributableAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSoftwareSolutionDistributableAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
