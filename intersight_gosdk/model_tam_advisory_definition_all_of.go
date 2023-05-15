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
	"time"
)

// TamAdvisoryDefinitionAllOf Definition of the list of properties defined in 'tam.AdvisoryDefinition', excluding properties defined in parent classes.
type TamAdvisoryDefinitionAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType      string                         `json:"ObjectType"`
	Actions         []TamAction                    `json:"Actions,omitempty"`
	AdvisoryDetails NullableTamBaseAdvisoryDetails `json:"AdvisoryDetails,omitempty"`
	// Cisco generated identifier for the published security/field-notice/end-of-life advisory.
	AdvisoryId     *string            `json:"AdvisoryId,omitempty"`
	ApiDataSources []TamApiDataSource `json:"ApiDataSources,omitempty"`
	// Date when the security/field-notice/end-of-life advisory was first published by Cisco.
	DatePublished *time.Time `json:"DatePublished,omitempty"`
	// Date when the security/field-notice/end-of-life advisory was last updated by Cisco.
	DateUpdated *time.Time `json:"DateUpdated,omitempty"`
	// Orion pod on which this advisory should process. * `tier1` - Advisory processing will be taken care in first advisory driver of multinode cluster. * `tier2` - Advisory processing will be taken care in second advisory driver of multinode cluster.
	ExecuteOnPod *string `json:"ExecuteOnPod,omitempty"`
	// A link to an external URL describing security Advisory in more details.
	ExternalUrl  *string  `json:"ExternalUrl,omitempty"`
	OtherRefUrls []string `json:"OtherRefUrls,omitempty"`
	// Recommended action to resolve the security advisory.
	Recommendation *string           `json:"Recommendation,omitempty"`
	S3DataSources  []TamS3DataSource `json:"S3DataSources,omitempty"`
	// The type (field notice, security advisory, end-of-life milestone advisory etc.) of Intersight advisory. * `securityAdvisory` - Respresents the psirt alert type (https://tools.cisco.com/security/center/publicationListing.x). * `fieldNotice` - Respresents the field notice alert type (https://www.cisco.com/c/en/us/support/web/tsd-products-field-notice-summary.html). * `eolAdvisory` - Represents product End of Life (EOL) type (https://www.cisco.com/c/en/us/products/eos-eol-policy.html).
	Type *string `json:"Type,omitempty"`
	// Cisco assigned advisory/field-notice/end-of-life version after latest revision.
	Version *string `json:"Version,omitempty"`
	// Workarounds available for the advisory.
	Workaround           *string                               `json:"Workaround,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TamAdvisoryDefinitionAllOf TamAdvisoryDefinitionAllOf

// NewTamAdvisoryDefinitionAllOf instantiates a new TamAdvisoryDefinitionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTamAdvisoryDefinitionAllOf(classId string, objectType string) *TamAdvisoryDefinitionAllOf {
	this := TamAdvisoryDefinitionAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var executeOnPod string = "tier1"
	this.ExecuteOnPod = &executeOnPod
	var type_ string = "securityAdvisory"
	this.Type = &type_
	return &this
}

// NewTamAdvisoryDefinitionAllOfWithDefaults instantiates a new TamAdvisoryDefinitionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTamAdvisoryDefinitionAllOfWithDefaults() *TamAdvisoryDefinitionAllOf {
	this := TamAdvisoryDefinitionAllOf{}
	var classId string = "tam.AdvisoryDefinition"
	this.ClassId = classId
	var objectType string = "tam.AdvisoryDefinition"
	this.ObjectType = objectType
	var executeOnPod string = "tier1"
	this.ExecuteOnPod = &executeOnPod
	var type_ string = "securityAdvisory"
	this.Type = &type_
	return &this
}

// GetClassId returns the ClassId field value
func (o *TamAdvisoryDefinitionAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *TamAdvisoryDefinitionAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *TamAdvisoryDefinitionAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *TamAdvisoryDefinitionAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *TamAdvisoryDefinitionAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *TamAdvisoryDefinitionAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetActions returns the Actions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TamAdvisoryDefinitionAllOf) GetActions() []TamAction {
	if o == nil {
		var ret []TamAction
		return ret
	}
	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TamAdvisoryDefinitionAllOf) GetActionsOk() ([]TamAction, bool) {
	if o == nil || o.Actions == nil {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *TamAdvisoryDefinitionAllOf) HasActions() bool {
	if o != nil && o.Actions != nil {
		return true
	}

	return false
}

// SetActions gets a reference to the given []TamAction and assigns it to the Actions field.
func (o *TamAdvisoryDefinitionAllOf) SetActions(v []TamAction) {
	o.Actions = v
}

// GetAdvisoryDetails returns the AdvisoryDetails field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TamAdvisoryDefinitionAllOf) GetAdvisoryDetails() TamBaseAdvisoryDetails {
	if o == nil || o.AdvisoryDetails.Get() == nil {
		var ret TamBaseAdvisoryDetails
		return ret
	}
	return *o.AdvisoryDetails.Get()
}

// GetAdvisoryDetailsOk returns a tuple with the AdvisoryDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TamAdvisoryDefinitionAllOf) GetAdvisoryDetailsOk() (*TamBaseAdvisoryDetails, bool) {
	if o == nil {
		return nil, false
	}
	return o.AdvisoryDetails.Get(), o.AdvisoryDetails.IsSet()
}

// HasAdvisoryDetails returns a boolean if a field has been set.
func (o *TamAdvisoryDefinitionAllOf) HasAdvisoryDetails() bool {
	if o != nil && o.AdvisoryDetails.IsSet() {
		return true
	}

	return false
}

// SetAdvisoryDetails gets a reference to the given NullableTamBaseAdvisoryDetails and assigns it to the AdvisoryDetails field.
func (o *TamAdvisoryDefinitionAllOf) SetAdvisoryDetails(v TamBaseAdvisoryDetails) {
	o.AdvisoryDetails.Set(&v)
}

// SetAdvisoryDetailsNil sets the value for AdvisoryDetails to be an explicit nil
func (o *TamAdvisoryDefinitionAllOf) SetAdvisoryDetailsNil() {
	o.AdvisoryDetails.Set(nil)
}

// UnsetAdvisoryDetails ensures that no value is present for AdvisoryDetails, not even an explicit nil
func (o *TamAdvisoryDefinitionAllOf) UnsetAdvisoryDetails() {
	o.AdvisoryDetails.Unset()
}

// GetAdvisoryId returns the AdvisoryId field value if set, zero value otherwise.
func (o *TamAdvisoryDefinitionAllOf) GetAdvisoryId() string {
	if o == nil || o.AdvisoryId == nil {
		var ret string
		return ret
	}
	return *o.AdvisoryId
}

// GetAdvisoryIdOk returns a tuple with the AdvisoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamAdvisoryDefinitionAllOf) GetAdvisoryIdOk() (*string, bool) {
	if o == nil || o.AdvisoryId == nil {
		return nil, false
	}
	return o.AdvisoryId, true
}

// HasAdvisoryId returns a boolean if a field has been set.
func (o *TamAdvisoryDefinitionAllOf) HasAdvisoryId() bool {
	if o != nil && o.AdvisoryId != nil {
		return true
	}

	return false
}

// SetAdvisoryId gets a reference to the given string and assigns it to the AdvisoryId field.
func (o *TamAdvisoryDefinitionAllOf) SetAdvisoryId(v string) {
	o.AdvisoryId = &v
}

// GetApiDataSources returns the ApiDataSources field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TamAdvisoryDefinitionAllOf) GetApiDataSources() []TamApiDataSource {
	if o == nil {
		var ret []TamApiDataSource
		return ret
	}
	return o.ApiDataSources
}

// GetApiDataSourcesOk returns a tuple with the ApiDataSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TamAdvisoryDefinitionAllOf) GetApiDataSourcesOk() ([]TamApiDataSource, bool) {
	if o == nil || o.ApiDataSources == nil {
		return nil, false
	}
	return o.ApiDataSources, true
}

// HasApiDataSources returns a boolean if a field has been set.
func (o *TamAdvisoryDefinitionAllOf) HasApiDataSources() bool {
	if o != nil && o.ApiDataSources != nil {
		return true
	}

	return false
}

// SetApiDataSources gets a reference to the given []TamApiDataSource and assigns it to the ApiDataSources field.
func (o *TamAdvisoryDefinitionAllOf) SetApiDataSources(v []TamApiDataSource) {
	o.ApiDataSources = v
}

// GetDatePublished returns the DatePublished field value if set, zero value otherwise.
func (o *TamAdvisoryDefinitionAllOf) GetDatePublished() time.Time {
	if o == nil || o.DatePublished == nil {
		var ret time.Time
		return ret
	}
	return *o.DatePublished
}

// GetDatePublishedOk returns a tuple with the DatePublished field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamAdvisoryDefinitionAllOf) GetDatePublishedOk() (*time.Time, bool) {
	if o == nil || o.DatePublished == nil {
		return nil, false
	}
	return o.DatePublished, true
}

// HasDatePublished returns a boolean if a field has been set.
func (o *TamAdvisoryDefinitionAllOf) HasDatePublished() bool {
	if o != nil && o.DatePublished != nil {
		return true
	}

	return false
}

// SetDatePublished gets a reference to the given time.Time and assigns it to the DatePublished field.
func (o *TamAdvisoryDefinitionAllOf) SetDatePublished(v time.Time) {
	o.DatePublished = &v
}

// GetDateUpdated returns the DateUpdated field value if set, zero value otherwise.
func (o *TamAdvisoryDefinitionAllOf) GetDateUpdated() time.Time {
	if o == nil || o.DateUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.DateUpdated
}

// GetDateUpdatedOk returns a tuple with the DateUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamAdvisoryDefinitionAllOf) GetDateUpdatedOk() (*time.Time, bool) {
	if o == nil || o.DateUpdated == nil {
		return nil, false
	}
	return o.DateUpdated, true
}

// HasDateUpdated returns a boolean if a field has been set.
func (o *TamAdvisoryDefinitionAllOf) HasDateUpdated() bool {
	if o != nil && o.DateUpdated != nil {
		return true
	}

	return false
}

// SetDateUpdated gets a reference to the given time.Time and assigns it to the DateUpdated field.
func (o *TamAdvisoryDefinitionAllOf) SetDateUpdated(v time.Time) {
	o.DateUpdated = &v
}

// GetExecuteOnPod returns the ExecuteOnPod field value if set, zero value otherwise.
func (o *TamAdvisoryDefinitionAllOf) GetExecuteOnPod() string {
	if o == nil || o.ExecuteOnPod == nil {
		var ret string
		return ret
	}
	return *o.ExecuteOnPod
}

// GetExecuteOnPodOk returns a tuple with the ExecuteOnPod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamAdvisoryDefinitionAllOf) GetExecuteOnPodOk() (*string, bool) {
	if o == nil || o.ExecuteOnPod == nil {
		return nil, false
	}
	return o.ExecuteOnPod, true
}

// HasExecuteOnPod returns a boolean if a field has been set.
func (o *TamAdvisoryDefinitionAllOf) HasExecuteOnPod() bool {
	if o != nil && o.ExecuteOnPod != nil {
		return true
	}

	return false
}

// SetExecuteOnPod gets a reference to the given string and assigns it to the ExecuteOnPod field.
func (o *TamAdvisoryDefinitionAllOf) SetExecuteOnPod(v string) {
	o.ExecuteOnPod = &v
}

// GetExternalUrl returns the ExternalUrl field value if set, zero value otherwise.
func (o *TamAdvisoryDefinitionAllOf) GetExternalUrl() string {
	if o == nil || o.ExternalUrl == nil {
		var ret string
		return ret
	}
	return *o.ExternalUrl
}

// GetExternalUrlOk returns a tuple with the ExternalUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamAdvisoryDefinitionAllOf) GetExternalUrlOk() (*string, bool) {
	if o == nil || o.ExternalUrl == nil {
		return nil, false
	}
	return o.ExternalUrl, true
}

// HasExternalUrl returns a boolean if a field has been set.
func (o *TamAdvisoryDefinitionAllOf) HasExternalUrl() bool {
	if o != nil && o.ExternalUrl != nil {
		return true
	}

	return false
}

// SetExternalUrl gets a reference to the given string and assigns it to the ExternalUrl field.
func (o *TamAdvisoryDefinitionAllOf) SetExternalUrl(v string) {
	o.ExternalUrl = &v
}

// GetOtherRefUrls returns the OtherRefUrls field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TamAdvisoryDefinitionAllOf) GetOtherRefUrls() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.OtherRefUrls
}

// GetOtherRefUrlsOk returns a tuple with the OtherRefUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TamAdvisoryDefinitionAllOf) GetOtherRefUrlsOk() ([]string, bool) {
	if o == nil || o.OtherRefUrls == nil {
		return nil, false
	}
	return o.OtherRefUrls, true
}

// HasOtherRefUrls returns a boolean if a field has been set.
func (o *TamAdvisoryDefinitionAllOf) HasOtherRefUrls() bool {
	if o != nil && o.OtherRefUrls != nil {
		return true
	}

	return false
}

// SetOtherRefUrls gets a reference to the given []string and assigns it to the OtherRefUrls field.
func (o *TamAdvisoryDefinitionAllOf) SetOtherRefUrls(v []string) {
	o.OtherRefUrls = v
}

// GetRecommendation returns the Recommendation field value if set, zero value otherwise.
func (o *TamAdvisoryDefinitionAllOf) GetRecommendation() string {
	if o == nil || o.Recommendation == nil {
		var ret string
		return ret
	}
	return *o.Recommendation
}

// GetRecommendationOk returns a tuple with the Recommendation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamAdvisoryDefinitionAllOf) GetRecommendationOk() (*string, bool) {
	if o == nil || o.Recommendation == nil {
		return nil, false
	}
	return o.Recommendation, true
}

// HasRecommendation returns a boolean if a field has been set.
func (o *TamAdvisoryDefinitionAllOf) HasRecommendation() bool {
	if o != nil && o.Recommendation != nil {
		return true
	}

	return false
}

// SetRecommendation gets a reference to the given string and assigns it to the Recommendation field.
func (o *TamAdvisoryDefinitionAllOf) SetRecommendation(v string) {
	o.Recommendation = &v
}

// GetS3DataSources returns the S3DataSources field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TamAdvisoryDefinitionAllOf) GetS3DataSources() []TamS3DataSource {
	if o == nil {
		var ret []TamS3DataSource
		return ret
	}
	return o.S3DataSources
}

// GetS3DataSourcesOk returns a tuple with the S3DataSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TamAdvisoryDefinitionAllOf) GetS3DataSourcesOk() ([]TamS3DataSource, bool) {
	if o == nil || o.S3DataSources == nil {
		return nil, false
	}
	return o.S3DataSources, true
}

// HasS3DataSources returns a boolean if a field has been set.
func (o *TamAdvisoryDefinitionAllOf) HasS3DataSources() bool {
	if o != nil && o.S3DataSources != nil {
		return true
	}

	return false
}

// SetS3DataSources gets a reference to the given []TamS3DataSource and assigns it to the S3DataSources field.
func (o *TamAdvisoryDefinitionAllOf) SetS3DataSources(v []TamS3DataSource) {
	o.S3DataSources = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TamAdvisoryDefinitionAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamAdvisoryDefinitionAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TamAdvisoryDefinitionAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TamAdvisoryDefinitionAllOf) SetType(v string) {
	o.Type = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *TamAdvisoryDefinitionAllOf) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamAdvisoryDefinitionAllOf) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *TamAdvisoryDefinitionAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *TamAdvisoryDefinitionAllOf) SetVersion(v string) {
	o.Version = &v
}

// GetWorkaround returns the Workaround field value if set, zero value otherwise.
func (o *TamAdvisoryDefinitionAllOf) GetWorkaround() string {
	if o == nil || o.Workaround == nil {
		var ret string
		return ret
	}
	return *o.Workaround
}

// GetWorkaroundOk returns a tuple with the Workaround field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamAdvisoryDefinitionAllOf) GetWorkaroundOk() (*string, bool) {
	if o == nil || o.Workaround == nil {
		return nil, false
	}
	return o.Workaround, true
}

// HasWorkaround returns a boolean if a field has been set.
func (o *TamAdvisoryDefinitionAllOf) HasWorkaround() bool {
	if o != nil && o.Workaround != nil {
		return true
	}

	return false
}

// SetWorkaround gets a reference to the given string and assigns it to the Workaround field.
func (o *TamAdvisoryDefinitionAllOf) SetWorkaround(v string) {
	o.Workaround = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *TamAdvisoryDefinitionAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamAdvisoryDefinitionAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *TamAdvisoryDefinitionAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *TamAdvisoryDefinitionAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o TamAdvisoryDefinitionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Actions != nil {
		toSerialize["Actions"] = o.Actions
	}
	if o.AdvisoryDetails.IsSet() {
		toSerialize["AdvisoryDetails"] = o.AdvisoryDetails.Get()
	}
	if o.AdvisoryId != nil {
		toSerialize["AdvisoryId"] = o.AdvisoryId
	}
	if o.ApiDataSources != nil {
		toSerialize["ApiDataSources"] = o.ApiDataSources
	}
	if o.DatePublished != nil {
		toSerialize["DatePublished"] = o.DatePublished
	}
	if o.DateUpdated != nil {
		toSerialize["DateUpdated"] = o.DateUpdated
	}
	if o.ExecuteOnPod != nil {
		toSerialize["ExecuteOnPod"] = o.ExecuteOnPod
	}
	if o.ExternalUrl != nil {
		toSerialize["ExternalUrl"] = o.ExternalUrl
	}
	if o.OtherRefUrls != nil {
		toSerialize["OtherRefUrls"] = o.OtherRefUrls
	}
	if o.Recommendation != nil {
		toSerialize["Recommendation"] = o.Recommendation
	}
	if o.S3DataSources != nil {
		toSerialize["S3DataSources"] = o.S3DataSources
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	if o.Workaround != nil {
		toSerialize["Workaround"] = o.Workaround
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TamAdvisoryDefinitionAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varTamAdvisoryDefinitionAllOf := _TamAdvisoryDefinitionAllOf{}

	if err = json.Unmarshal(bytes, &varTamAdvisoryDefinitionAllOf); err == nil {
		*o = TamAdvisoryDefinitionAllOf(varTamAdvisoryDefinitionAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Actions")
		delete(additionalProperties, "AdvisoryDetails")
		delete(additionalProperties, "AdvisoryId")
		delete(additionalProperties, "ApiDataSources")
		delete(additionalProperties, "DatePublished")
		delete(additionalProperties, "DateUpdated")
		delete(additionalProperties, "ExecuteOnPod")
		delete(additionalProperties, "ExternalUrl")
		delete(additionalProperties, "OtherRefUrls")
		delete(additionalProperties, "Recommendation")
		delete(additionalProperties, "S3DataSources")
		delete(additionalProperties, "Type")
		delete(additionalProperties, "Version")
		delete(additionalProperties, "Workaround")
		delete(additionalProperties, "Organization")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTamAdvisoryDefinitionAllOf struct {
	value *TamAdvisoryDefinitionAllOf
	isSet bool
}

func (v NullableTamAdvisoryDefinitionAllOf) Get() *TamAdvisoryDefinitionAllOf {
	return v.value
}

func (v *NullableTamAdvisoryDefinitionAllOf) Set(val *TamAdvisoryDefinitionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTamAdvisoryDefinitionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTamAdvisoryDefinitionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTamAdvisoryDefinitionAllOf(val *TamAdvisoryDefinitionAllOf) *NullableTamAdvisoryDefinitionAllOf {
	return &NullableTamAdvisoryDefinitionAllOf{value: val, isSet: true}
}

func (v NullableTamAdvisoryDefinitionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTamAdvisoryDefinitionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
