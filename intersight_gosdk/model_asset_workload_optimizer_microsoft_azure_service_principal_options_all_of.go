/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-7658
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsAllOf Definition of the list of properties defined in 'asset.WorkloadOptimizerMicrosoftAzureServicePrincipalOptions', excluding properties defined in parent classes.
type AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Azure has different endpoints for managing Germany subscriptions. Azure cloud type helps in differentiating between regular subscriptions and Germany subscriptions to manage the Azure services for workload optimization. Documentation for germany cloud [link](https://docs.microsoft.com/en-us/azure/germany/germany-manage-subscriptions). * `Global` - Global cloud type for Azure subscription. * `Germany` - Germany cloud type for Azure subscription.
	AzureCloudType *string `json:"AzureCloudType,omitempty"`
	// ID of the tenant used while authenticating the managed target.
	TenantId             *string `json:"TenantId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsAllOf AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsAllOf

// NewAssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsAllOf instantiates a new AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsAllOf(classId string, objectType string) *AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsAllOf {
	this := AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var azureCloudType string = "Global"
	this.AzureCloudType = &azureCloudType
	return &this
}

// NewAssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsAllOfWithDefaults instantiates a new AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsAllOfWithDefaults() *AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsAllOf {
	this := AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsAllOf{}
	var classId string = "asset.WorkloadOptimizerMicrosoftAzureServicePrincipalOptions"
	this.ClassId = classId
	var objectType string = "asset.WorkloadOptimizerMicrosoftAzureServicePrincipalOptions"
	this.ObjectType = objectType
	var azureCloudType string = "Global"
	this.AzureCloudType = &azureCloudType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAzureCloudType returns the AzureCloudType field value if set, zero value otherwise.
func (o *AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsAllOf) GetAzureCloudType() string {
	if o == nil || o.AzureCloudType == nil {
		var ret string
		return ret
	}
	return *o.AzureCloudType
}

// GetAzureCloudTypeOk returns a tuple with the AzureCloudType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsAllOf) GetAzureCloudTypeOk() (*string, bool) {
	if o == nil || o.AzureCloudType == nil {
		return nil, false
	}
	return o.AzureCloudType, true
}

// HasAzureCloudType returns a boolean if a field has been set.
func (o *AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsAllOf) HasAzureCloudType() bool {
	if o != nil && o.AzureCloudType != nil {
		return true
	}

	return false
}

// SetAzureCloudType gets a reference to the given string and assigns it to the AzureCloudType field.
func (o *AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsAllOf) SetAzureCloudType(v string) {
	o.AzureCloudType = &v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsAllOf) GetTenantId() string {
	if o == nil || o.TenantId == nil {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsAllOf) GetTenantIdOk() (*string, bool) {
	if o == nil || o.TenantId == nil {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsAllOf) HasTenantId() bool {
	if o != nil && o.TenantId != nil {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsAllOf) SetTenantId(v string) {
	o.TenantId = &v
}

func (o AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AzureCloudType != nil {
		toSerialize["AzureCloudType"] = o.AzureCloudType
	}
	if o.TenantId != nil {
		toSerialize["TenantId"] = o.TenantId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varAssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsAllOf := _AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsAllOf{}

	if err = json.Unmarshal(bytes, &varAssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsAllOf); err == nil {
		*o = AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsAllOf(varAssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AzureCloudType")
		delete(additionalProperties, "TenantId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsAllOf struct {
	value *AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsAllOf
	isSet bool
}

func (v NullableAssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsAllOf) Get() *AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsAllOf {
	return v.value
}

func (v *NullableAssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsAllOf) Set(val *AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsAllOf(val *AssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsAllOf) *NullableAssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsAllOf {
	return &NullableAssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsAllOf{value: val, isSet: true}
}

func (v NullableAssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetWorkloadOptimizerMicrosoftAzureServicePrincipalOptionsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
