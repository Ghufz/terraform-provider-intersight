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

// OsCatalogAllOf Definition of the list of properties defined in 'os.Catalog', excluding properties defined in parent classes.
type OsCatalogAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The catalog name. There will be one for system and one for each user account.
	Name *string `json:"Name,omitempty"`
	// An array of relationships to osConfigurationFile resources.
	ConfigurationFiles []OsConfigurationFileRelationship `json:"ConfigurationFiles,omitempty"`
	// An array of relationships to osDistribution resources.
	Distributions        []OsDistributionRelationship          `json:"Distributions,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OsCatalogAllOf OsCatalogAllOf

// NewOsCatalogAllOf instantiates a new OsCatalogAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOsCatalogAllOf(classId string, objectType string) *OsCatalogAllOf {
	this := OsCatalogAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewOsCatalogAllOfWithDefaults instantiates a new OsCatalogAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOsCatalogAllOfWithDefaults() *OsCatalogAllOf {
	this := OsCatalogAllOf{}
	var classId string = "os.Catalog"
	this.ClassId = classId
	var objectType string = "os.Catalog"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *OsCatalogAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *OsCatalogAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *OsCatalogAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *OsCatalogAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *OsCatalogAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *OsCatalogAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OsCatalogAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsCatalogAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OsCatalogAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OsCatalogAllOf) SetName(v string) {
	o.Name = &v
}

// GetConfigurationFiles returns the ConfigurationFiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OsCatalogAllOf) GetConfigurationFiles() []OsConfigurationFileRelationship {
	if o == nil {
		var ret []OsConfigurationFileRelationship
		return ret
	}
	return o.ConfigurationFiles
}

// GetConfigurationFilesOk returns a tuple with the ConfigurationFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OsCatalogAllOf) GetConfigurationFilesOk() ([]OsConfigurationFileRelationship, bool) {
	if o == nil || o.ConfigurationFiles == nil {
		return nil, false
	}
	return o.ConfigurationFiles, true
}

// HasConfigurationFiles returns a boolean if a field has been set.
func (o *OsCatalogAllOf) HasConfigurationFiles() bool {
	if o != nil && o.ConfigurationFiles != nil {
		return true
	}

	return false
}

// SetConfigurationFiles gets a reference to the given []OsConfigurationFileRelationship and assigns it to the ConfigurationFiles field.
func (o *OsCatalogAllOf) SetConfigurationFiles(v []OsConfigurationFileRelationship) {
	o.ConfigurationFiles = v
}

// GetDistributions returns the Distributions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OsCatalogAllOf) GetDistributions() []OsDistributionRelationship {
	if o == nil {
		var ret []OsDistributionRelationship
		return ret
	}
	return o.Distributions
}

// GetDistributionsOk returns a tuple with the Distributions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OsCatalogAllOf) GetDistributionsOk() ([]OsDistributionRelationship, bool) {
	if o == nil || o.Distributions == nil {
		return nil, false
	}
	return o.Distributions, true
}

// HasDistributions returns a boolean if a field has been set.
func (o *OsCatalogAllOf) HasDistributions() bool {
	if o != nil && o.Distributions != nil {
		return true
	}

	return false
}

// SetDistributions gets a reference to the given []OsDistributionRelationship and assigns it to the Distributions field.
func (o *OsCatalogAllOf) SetDistributions(v []OsDistributionRelationship) {
	o.Distributions = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *OsCatalogAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsCatalogAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *OsCatalogAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *OsCatalogAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o OsCatalogAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.ConfigurationFiles != nil {
		toSerialize["ConfigurationFiles"] = o.ConfigurationFiles
	}
	if o.Distributions != nil {
		toSerialize["Distributions"] = o.Distributions
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OsCatalogAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varOsCatalogAllOf := _OsCatalogAllOf{}

	if err = json.Unmarshal(bytes, &varOsCatalogAllOf); err == nil {
		*o = OsCatalogAllOf(varOsCatalogAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "ConfigurationFiles")
		delete(additionalProperties, "Distributions")
		delete(additionalProperties, "Organization")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOsCatalogAllOf struct {
	value *OsCatalogAllOf
	isSet bool
}

func (v NullableOsCatalogAllOf) Get() *OsCatalogAllOf {
	return v.value
}

func (v *NullableOsCatalogAllOf) Set(val *OsCatalogAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableOsCatalogAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableOsCatalogAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOsCatalogAllOf(val *OsCatalogAllOf) *NullableOsCatalogAllOf {
	return &NullableOsCatalogAllOf{value: val, isSet: true}
}

func (v NullableOsCatalogAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOsCatalogAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
