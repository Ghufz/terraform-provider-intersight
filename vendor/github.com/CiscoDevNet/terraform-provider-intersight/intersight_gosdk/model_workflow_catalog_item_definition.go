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

// WorkflowCatalogItemDefinition Catalog Item definition is a collection of Service items which are associated with workflow definition that can be used to deploy and manage service items.
type WorkflowCatalogItemDefinition struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The description for the catalog item which provides information on what are the pre-requisites to deploy the resource.
	Description *string `json:"Description,omitempty"`
	// A user friendly short name to identify the catalog item. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ) or an underscore (_).
	Label *string `json:"Label,omitempty"`
	// The name for this catalog item definition. You can have multiple versions of the catalog item with the same name. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:) or an underscore (_).
	Name *string `json:"Name,omitempty"`
	// Publish status of the catalog item. * `NotPublished` - A state of the service item or catalog item which is not yet published. * `Published` - A state denoting that the service item or catalog item is published.
	PublishStatus *string                   `json:"PublishStatus,omitempty"`
	ServiceItems  []WorkflowServiceItemType `json:"ServiceItems,omitempty"`
	// The CatalogItem Support depicts the support status of catalog, the values will be any of Supported or Deprecated state. The user can create a Catalog Service Request if the support status is supported, if its Deprecated then it cannot be instantiated. * `Supported` - The definition is a supported version and there will be no changes to the mandatory inputs or outputs. * `Beta` - The definition is a Beta version and this version can under go changes until the version is marked supported. * `Deprecated` - The version of definition is deprecated and typically there will be a higher version of the same definition that has been added.
	SupportStatus *string `json:"SupportStatus,omitempty"`
	// The user identifier who created or updated the catalog item definition.
	UserIdOrEmail         *string                               `json:"UserIdOrEmail,omitempty"`
	ValidationInformation NullableWorkflowValidationInformation `json:"ValidationInformation,omitempty"`
	// The version of the catalog item to support multiple versions.
	Version              *int64                                `json:"Version,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowCatalogItemDefinition WorkflowCatalogItemDefinition

// NewWorkflowCatalogItemDefinition instantiates a new WorkflowCatalogItemDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowCatalogItemDefinition(classId string, objectType string) *WorkflowCatalogItemDefinition {
	this := WorkflowCatalogItemDefinition{}
	this.ClassId = classId
	this.ObjectType = objectType
	var publishStatus string = "NotPublished"
	this.PublishStatus = &publishStatus
	var supportStatus string = "Supported"
	this.SupportStatus = &supportStatus
	var version int64 = 1
	this.Version = &version
	return &this
}

// NewWorkflowCatalogItemDefinitionWithDefaults instantiates a new WorkflowCatalogItemDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowCatalogItemDefinitionWithDefaults() *WorkflowCatalogItemDefinition {
	this := WorkflowCatalogItemDefinition{}
	var classId string = "workflow.CatalogItemDefinition"
	this.ClassId = classId
	var objectType string = "workflow.CatalogItemDefinition"
	this.ObjectType = objectType
	var publishStatus string = "NotPublished"
	this.PublishStatus = &publishStatus
	var supportStatus string = "Supported"
	this.SupportStatus = &supportStatus
	var version int64 = 1
	this.Version = &version
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowCatalogItemDefinition) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowCatalogItemDefinition) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowCatalogItemDefinition) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowCatalogItemDefinition) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowCatalogItemDefinition) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowCatalogItemDefinition) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WorkflowCatalogItemDefinition) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowCatalogItemDefinition) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WorkflowCatalogItemDefinition) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WorkflowCatalogItemDefinition) SetDescription(v string) {
	o.Description = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *WorkflowCatalogItemDefinition) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowCatalogItemDefinition) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *WorkflowCatalogItemDefinition) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *WorkflowCatalogItemDefinition) SetLabel(v string) {
	o.Label = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowCatalogItemDefinition) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowCatalogItemDefinition) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowCatalogItemDefinition) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowCatalogItemDefinition) SetName(v string) {
	o.Name = &v
}

// GetPublishStatus returns the PublishStatus field value if set, zero value otherwise.
func (o *WorkflowCatalogItemDefinition) GetPublishStatus() string {
	if o == nil || o.PublishStatus == nil {
		var ret string
		return ret
	}
	return *o.PublishStatus
}

// GetPublishStatusOk returns a tuple with the PublishStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowCatalogItemDefinition) GetPublishStatusOk() (*string, bool) {
	if o == nil || o.PublishStatus == nil {
		return nil, false
	}
	return o.PublishStatus, true
}

// HasPublishStatus returns a boolean if a field has been set.
func (o *WorkflowCatalogItemDefinition) HasPublishStatus() bool {
	if o != nil && o.PublishStatus != nil {
		return true
	}

	return false
}

// SetPublishStatus gets a reference to the given string and assigns it to the PublishStatus field.
func (o *WorkflowCatalogItemDefinition) SetPublishStatus(v string) {
	o.PublishStatus = &v
}

// GetServiceItems returns the ServiceItems field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowCatalogItemDefinition) GetServiceItems() []WorkflowServiceItemType {
	if o == nil {
		var ret []WorkflowServiceItemType
		return ret
	}
	return o.ServiceItems
}

// GetServiceItemsOk returns a tuple with the ServiceItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowCatalogItemDefinition) GetServiceItemsOk() ([]WorkflowServiceItemType, bool) {
	if o == nil || o.ServiceItems == nil {
		return nil, false
	}
	return o.ServiceItems, true
}

// HasServiceItems returns a boolean if a field has been set.
func (o *WorkflowCatalogItemDefinition) HasServiceItems() bool {
	if o != nil && o.ServiceItems != nil {
		return true
	}

	return false
}

// SetServiceItems gets a reference to the given []WorkflowServiceItemType and assigns it to the ServiceItems field.
func (o *WorkflowCatalogItemDefinition) SetServiceItems(v []WorkflowServiceItemType) {
	o.ServiceItems = v
}

// GetSupportStatus returns the SupportStatus field value if set, zero value otherwise.
func (o *WorkflowCatalogItemDefinition) GetSupportStatus() string {
	if o == nil || o.SupportStatus == nil {
		var ret string
		return ret
	}
	return *o.SupportStatus
}

// GetSupportStatusOk returns a tuple with the SupportStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowCatalogItemDefinition) GetSupportStatusOk() (*string, bool) {
	if o == nil || o.SupportStatus == nil {
		return nil, false
	}
	return o.SupportStatus, true
}

// HasSupportStatus returns a boolean if a field has been set.
func (o *WorkflowCatalogItemDefinition) HasSupportStatus() bool {
	if o != nil && o.SupportStatus != nil {
		return true
	}

	return false
}

// SetSupportStatus gets a reference to the given string and assigns it to the SupportStatus field.
func (o *WorkflowCatalogItemDefinition) SetSupportStatus(v string) {
	o.SupportStatus = &v
}

// GetUserIdOrEmail returns the UserIdOrEmail field value if set, zero value otherwise.
func (o *WorkflowCatalogItemDefinition) GetUserIdOrEmail() string {
	if o == nil || o.UserIdOrEmail == nil {
		var ret string
		return ret
	}
	return *o.UserIdOrEmail
}

// GetUserIdOrEmailOk returns a tuple with the UserIdOrEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowCatalogItemDefinition) GetUserIdOrEmailOk() (*string, bool) {
	if o == nil || o.UserIdOrEmail == nil {
		return nil, false
	}
	return o.UserIdOrEmail, true
}

// HasUserIdOrEmail returns a boolean if a field has been set.
func (o *WorkflowCatalogItemDefinition) HasUserIdOrEmail() bool {
	if o != nil && o.UserIdOrEmail != nil {
		return true
	}

	return false
}

// SetUserIdOrEmail gets a reference to the given string and assigns it to the UserIdOrEmail field.
func (o *WorkflowCatalogItemDefinition) SetUserIdOrEmail(v string) {
	o.UserIdOrEmail = &v
}

// GetValidationInformation returns the ValidationInformation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowCatalogItemDefinition) GetValidationInformation() WorkflowValidationInformation {
	if o == nil || o.ValidationInformation.Get() == nil {
		var ret WorkflowValidationInformation
		return ret
	}
	return *o.ValidationInformation.Get()
}

// GetValidationInformationOk returns a tuple with the ValidationInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowCatalogItemDefinition) GetValidationInformationOk() (*WorkflowValidationInformation, bool) {
	if o == nil {
		return nil, false
	}
	return o.ValidationInformation.Get(), o.ValidationInformation.IsSet()
}

// HasValidationInformation returns a boolean if a field has been set.
func (o *WorkflowCatalogItemDefinition) HasValidationInformation() bool {
	if o != nil && o.ValidationInformation.IsSet() {
		return true
	}

	return false
}

// SetValidationInformation gets a reference to the given NullableWorkflowValidationInformation and assigns it to the ValidationInformation field.
func (o *WorkflowCatalogItemDefinition) SetValidationInformation(v WorkflowValidationInformation) {
	o.ValidationInformation.Set(&v)
}

// SetValidationInformationNil sets the value for ValidationInformation to be an explicit nil
func (o *WorkflowCatalogItemDefinition) SetValidationInformationNil() {
	o.ValidationInformation.Set(nil)
}

// UnsetValidationInformation ensures that no value is present for ValidationInformation, not even an explicit nil
func (o *WorkflowCatalogItemDefinition) UnsetValidationInformation() {
	o.ValidationInformation.Unset()
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *WorkflowCatalogItemDefinition) GetVersion() int64 {
	if o == nil || o.Version == nil {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowCatalogItemDefinition) GetVersionOk() (*int64, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *WorkflowCatalogItemDefinition) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *WorkflowCatalogItemDefinition) SetVersion(v int64) {
	o.Version = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *WorkflowCatalogItemDefinition) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowCatalogItemDefinition) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *WorkflowCatalogItemDefinition) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *WorkflowCatalogItemDefinition) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o WorkflowCatalogItemDefinition) MarshalJSON() ([]byte, error) {
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
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.Label != nil {
		toSerialize["Label"] = o.Label
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.PublishStatus != nil {
		toSerialize["PublishStatus"] = o.PublishStatus
	}
	if o.ServiceItems != nil {
		toSerialize["ServiceItems"] = o.ServiceItems
	}
	if o.SupportStatus != nil {
		toSerialize["SupportStatus"] = o.SupportStatus
	}
	if o.UserIdOrEmail != nil {
		toSerialize["UserIdOrEmail"] = o.UserIdOrEmail
	}
	if o.ValidationInformation.IsSet() {
		toSerialize["ValidationInformation"] = o.ValidationInformation.Get()
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowCatalogItemDefinition) UnmarshalJSON(bytes []byte) (err error) {
	type WorkflowCatalogItemDefinitionWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The description for the catalog item which provides information on what are the pre-requisites to deploy the resource.
		Description *string `json:"Description,omitempty"`
		// A user friendly short name to identify the catalog item. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ) or an underscore (_).
		Label *string `json:"Label,omitempty"`
		// The name for this catalog item definition. You can have multiple versions of the catalog item with the same name. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:) or an underscore (_).
		Name *string `json:"Name,omitempty"`
		// Publish status of the catalog item. * `NotPublished` - A state of the service item or catalog item which is not yet published. * `Published` - A state denoting that the service item or catalog item is published.
		PublishStatus *string                   `json:"PublishStatus,omitempty"`
		ServiceItems  []WorkflowServiceItemType `json:"ServiceItems,omitempty"`
		// The CatalogItem Support depicts the support status of catalog, the values will be any of Supported or Deprecated state. The user can create a Catalog Service Request if the support status is supported, if its Deprecated then it cannot be instantiated. * `Supported` - The definition is a supported version and there will be no changes to the mandatory inputs or outputs. * `Beta` - The definition is a Beta version and this version can under go changes until the version is marked supported. * `Deprecated` - The version of definition is deprecated and typically there will be a higher version of the same definition that has been added.
		SupportStatus *string `json:"SupportStatus,omitempty"`
		// The user identifier who created or updated the catalog item definition.
		UserIdOrEmail         *string                               `json:"UserIdOrEmail,omitempty"`
		ValidationInformation NullableWorkflowValidationInformation `json:"ValidationInformation,omitempty"`
		// The version of the catalog item to support multiple versions.
		Version      *int64                                `json:"Version,omitempty"`
		Organization *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	}

	varWorkflowCatalogItemDefinitionWithoutEmbeddedStruct := WorkflowCatalogItemDefinitionWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varWorkflowCatalogItemDefinitionWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowCatalogItemDefinition := _WorkflowCatalogItemDefinition{}
		varWorkflowCatalogItemDefinition.ClassId = varWorkflowCatalogItemDefinitionWithoutEmbeddedStruct.ClassId
		varWorkflowCatalogItemDefinition.ObjectType = varWorkflowCatalogItemDefinitionWithoutEmbeddedStruct.ObjectType
		varWorkflowCatalogItemDefinition.Description = varWorkflowCatalogItemDefinitionWithoutEmbeddedStruct.Description
		varWorkflowCatalogItemDefinition.Label = varWorkflowCatalogItemDefinitionWithoutEmbeddedStruct.Label
		varWorkflowCatalogItemDefinition.Name = varWorkflowCatalogItemDefinitionWithoutEmbeddedStruct.Name
		varWorkflowCatalogItemDefinition.PublishStatus = varWorkflowCatalogItemDefinitionWithoutEmbeddedStruct.PublishStatus
		varWorkflowCatalogItemDefinition.ServiceItems = varWorkflowCatalogItemDefinitionWithoutEmbeddedStruct.ServiceItems
		varWorkflowCatalogItemDefinition.SupportStatus = varWorkflowCatalogItemDefinitionWithoutEmbeddedStruct.SupportStatus
		varWorkflowCatalogItemDefinition.UserIdOrEmail = varWorkflowCatalogItemDefinitionWithoutEmbeddedStruct.UserIdOrEmail
		varWorkflowCatalogItemDefinition.ValidationInformation = varWorkflowCatalogItemDefinitionWithoutEmbeddedStruct.ValidationInformation
		varWorkflowCatalogItemDefinition.Version = varWorkflowCatalogItemDefinitionWithoutEmbeddedStruct.Version
		varWorkflowCatalogItemDefinition.Organization = varWorkflowCatalogItemDefinitionWithoutEmbeddedStruct.Organization
		*o = WorkflowCatalogItemDefinition(varWorkflowCatalogItemDefinition)
	} else {
		return err
	}

	varWorkflowCatalogItemDefinition := _WorkflowCatalogItemDefinition{}

	err = json.Unmarshal(bytes, &varWorkflowCatalogItemDefinition)
	if err == nil {
		o.MoBaseMo = varWorkflowCatalogItemDefinition.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "Label")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "PublishStatus")
		delete(additionalProperties, "ServiceItems")
		delete(additionalProperties, "SupportStatus")
		delete(additionalProperties, "UserIdOrEmail")
		delete(additionalProperties, "ValidationInformation")
		delete(additionalProperties, "Version")
		delete(additionalProperties, "Organization")

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

type NullableWorkflowCatalogItemDefinition struct {
	value *WorkflowCatalogItemDefinition
	isSet bool
}

func (v NullableWorkflowCatalogItemDefinition) Get() *WorkflowCatalogItemDefinition {
	return v.value
}

func (v *NullableWorkflowCatalogItemDefinition) Set(val *WorkflowCatalogItemDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowCatalogItemDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowCatalogItemDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowCatalogItemDefinition(val *WorkflowCatalogItemDefinition) *NullableWorkflowCatalogItemDefinition {
	return &NullableWorkflowCatalogItemDefinition{value: val, isSet: true}
}

func (v NullableWorkflowCatalogItemDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowCatalogItemDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
