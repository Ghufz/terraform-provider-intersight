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

// OpenapiOpenApiSpecification An OpenAPI specification file uploaded to generate workflow tasks.
type OpenapiOpenApiSpecification struct {
	SoftwarerepositoryFile
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The path of the file in S3/minio bucket.
	FilePath *string `json:"FilePath,omitempty"`
	// A unique tracking ID for the file being uploaded.
	FileUploadId         *string                               `json:"FileUploadId,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OpenapiOpenApiSpecification OpenapiOpenApiSpecification

// NewOpenapiOpenApiSpecification instantiates a new OpenapiOpenApiSpecification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenapiOpenApiSpecification(classId string, objectType string) *OpenapiOpenApiSpecification {
	this := OpenapiOpenApiSpecification{}
	this.ClassId = classId
	this.ObjectType = objectType
	var importAction string = "None"
	this.ImportAction = &importAction
	return &this
}

// NewOpenapiOpenApiSpecificationWithDefaults instantiates a new OpenapiOpenApiSpecification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenapiOpenApiSpecificationWithDefaults() *OpenapiOpenApiSpecification {
	this := OpenapiOpenApiSpecification{}
	var classId string = "openapi.OpenApiSpecification"
	this.ClassId = classId
	var objectType string = "openapi.OpenApiSpecification"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *OpenapiOpenApiSpecification) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *OpenapiOpenApiSpecification) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *OpenapiOpenApiSpecification) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *OpenapiOpenApiSpecification) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *OpenapiOpenApiSpecification) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *OpenapiOpenApiSpecification) SetObjectType(v string) {
	o.ObjectType = v
}

// GetFilePath returns the FilePath field value if set, zero value otherwise.
func (o *OpenapiOpenApiSpecification) GetFilePath() string {
	if o == nil || o.FilePath == nil {
		var ret string
		return ret
	}
	return *o.FilePath
}

// GetFilePathOk returns a tuple with the FilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenapiOpenApiSpecification) GetFilePathOk() (*string, bool) {
	if o == nil || o.FilePath == nil {
		return nil, false
	}
	return o.FilePath, true
}

// HasFilePath returns a boolean if a field has been set.
func (o *OpenapiOpenApiSpecification) HasFilePath() bool {
	if o != nil && o.FilePath != nil {
		return true
	}

	return false
}

// SetFilePath gets a reference to the given string and assigns it to the FilePath field.
func (o *OpenapiOpenApiSpecification) SetFilePath(v string) {
	o.FilePath = &v
}

// GetFileUploadId returns the FileUploadId field value if set, zero value otherwise.
func (o *OpenapiOpenApiSpecification) GetFileUploadId() string {
	if o == nil || o.FileUploadId == nil {
		var ret string
		return ret
	}
	return *o.FileUploadId
}

// GetFileUploadIdOk returns a tuple with the FileUploadId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenapiOpenApiSpecification) GetFileUploadIdOk() (*string, bool) {
	if o == nil || o.FileUploadId == nil {
		return nil, false
	}
	return o.FileUploadId, true
}

// HasFileUploadId returns a boolean if a field has been set.
func (o *OpenapiOpenApiSpecification) HasFileUploadId() bool {
	if o != nil && o.FileUploadId != nil {
		return true
	}

	return false
}

// SetFileUploadId gets a reference to the given string and assigns it to the FileUploadId field.
func (o *OpenapiOpenApiSpecification) SetFileUploadId(v string) {
	o.FileUploadId = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *OpenapiOpenApiSpecification) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenapiOpenApiSpecification) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *OpenapiOpenApiSpecification) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *OpenapiOpenApiSpecification) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o OpenapiOpenApiSpecification) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedSoftwarerepositoryFile, errSoftwarerepositoryFile := json.Marshal(o.SoftwarerepositoryFile)
	if errSoftwarerepositoryFile != nil {
		return []byte{}, errSoftwarerepositoryFile
	}
	errSoftwarerepositoryFile = json.Unmarshal([]byte(serializedSoftwarerepositoryFile), &toSerialize)
	if errSoftwarerepositoryFile != nil {
		return []byte{}, errSoftwarerepositoryFile
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.FilePath != nil {
		toSerialize["FilePath"] = o.FilePath
	}
	if o.FileUploadId != nil {
		toSerialize["FileUploadId"] = o.FileUploadId
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OpenapiOpenApiSpecification) UnmarshalJSON(bytes []byte) (err error) {
	type OpenapiOpenApiSpecificationWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The path of the file in S3/minio bucket.
		FilePath *string `json:"FilePath,omitempty"`
		// A unique tracking ID for the file being uploaded.
		FileUploadId *string                               `json:"FileUploadId,omitempty"`
		Organization *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	}

	varOpenapiOpenApiSpecificationWithoutEmbeddedStruct := OpenapiOpenApiSpecificationWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varOpenapiOpenApiSpecificationWithoutEmbeddedStruct)
	if err == nil {
		varOpenapiOpenApiSpecification := _OpenapiOpenApiSpecification{}
		varOpenapiOpenApiSpecification.ClassId = varOpenapiOpenApiSpecificationWithoutEmbeddedStruct.ClassId
		varOpenapiOpenApiSpecification.ObjectType = varOpenapiOpenApiSpecificationWithoutEmbeddedStruct.ObjectType
		varOpenapiOpenApiSpecification.FilePath = varOpenapiOpenApiSpecificationWithoutEmbeddedStruct.FilePath
		varOpenapiOpenApiSpecification.FileUploadId = varOpenapiOpenApiSpecificationWithoutEmbeddedStruct.FileUploadId
		varOpenapiOpenApiSpecification.Organization = varOpenapiOpenApiSpecificationWithoutEmbeddedStruct.Organization
		*o = OpenapiOpenApiSpecification(varOpenapiOpenApiSpecification)
	} else {
		return err
	}

	varOpenapiOpenApiSpecification := _OpenapiOpenApiSpecification{}

	err = json.Unmarshal(bytes, &varOpenapiOpenApiSpecification)
	if err == nil {
		o.SoftwarerepositoryFile = varOpenapiOpenApiSpecification.SoftwarerepositoryFile
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "FilePath")
		delete(additionalProperties, "FileUploadId")
		delete(additionalProperties, "Organization")

		// remove fields from embedded structs
		reflectSoftwarerepositoryFile := reflect.ValueOf(o.SoftwarerepositoryFile)
		for i := 0; i < reflectSoftwarerepositoryFile.Type().NumField(); i++ {
			t := reflectSoftwarerepositoryFile.Type().Field(i)

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

type NullableOpenapiOpenApiSpecification struct {
	value *OpenapiOpenApiSpecification
	isSet bool
}

func (v NullableOpenapiOpenApiSpecification) Get() *OpenapiOpenApiSpecification {
	return v.value
}

func (v *NullableOpenapiOpenApiSpecification) Set(val *OpenapiOpenApiSpecification) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenapiOpenApiSpecification) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenapiOpenApiSpecification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenapiOpenApiSpecification(val *OpenapiOpenApiSpecification) *NullableOpenapiOpenApiSpecification {
	return &NullableOpenapiOpenApiSpecification{value: val, isSet: true}
}

func (v NullableOpenapiOpenApiSpecification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenapiOpenApiSpecification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
