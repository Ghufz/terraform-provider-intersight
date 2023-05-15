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

// AssetWorkloadOptimizerMicrosoftAzureBillingOptionsAllOf Definition of the list of properties defined in 'asset.WorkloadOptimizerMicrosoftAzureBillingOptions', excluding properties defined in parent classes.
type AssetWorkloadOptimizerMicrosoftAzureBillingOptionsAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The Microsoft Azure Billing Account ID.
	BillingAccountId *string `json:"BillingAccountId,omitempty"`
	// Name of the Cost Export Data that exports cost management data.
	CostExportName *string `json:"CostExportName,omitempty"`
	// Id of the tenant used while authenticating the managed target.
	TenantId             *string `json:"TenantId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetWorkloadOptimizerMicrosoftAzureBillingOptionsAllOf AssetWorkloadOptimizerMicrosoftAzureBillingOptionsAllOf

// NewAssetWorkloadOptimizerMicrosoftAzureBillingOptionsAllOf instantiates a new AssetWorkloadOptimizerMicrosoftAzureBillingOptionsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetWorkloadOptimizerMicrosoftAzureBillingOptionsAllOf(classId string, objectType string) *AssetWorkloadOptimizerMicrosoftAzureBillingOptionsAllOf {
	this := AssetWorkloadOptimizerMicrosoftAzureBillingOptionsAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAssetWorkloadOptimizerMicrosoftAzureBillingOptionsAllOfWithDefaults instantiates a new AssetWorkloadOptimizerMicrosoftAzureBillingOptionsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetWorkloadOptimizerMicrosoftAzureBillingOptionsAllOfWithDefaults() *AssetWorkloadOptimizerMicrosoftAzureBillingOptionsAllOf {
	this := AssetWorkloadOptimizerMicrosoftAzureBillingOptionsAllOf{}
	var classId string = "asset.WorkloadOptimizerMicrosoftAzureBillingOptions"
	this.ClassId = classId
	var objectType string = "asset.WorkloadOptimizerMicrosoftAzureBillingOptions"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetWorkloadOptimizerMicrosoftAzureBillingOptionsAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerMicrosoftAzureBillingOptionsAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetWorkloadOptimizerMicrosoftAzureBillingOptionsAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AssetWorkloadOptimizerMicrosoftAzureBillingOptionsAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerMicrosoftAzureBillingOptionsAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetWorkloadOptimizerMicrosoftAzureBillingOptionsAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBillingAccountId returns the BillingAccountId field value if set, zero value otherwise.
func (o *AssetWorkloadOptimizerMicrosoftAzureBillingOptionsAllOf) GetBillingAccountId() string {
	if o == nil || o.BillingAccountId == nil {
		var ret string
		return ret
	}
	return *o.BillingAccountId
}

// GetBillingAccountIdOk returns a tuple with the BillingAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerMicrosoftAzureBillingOptionsAllOf) GetBillingAccountIdOk() (*string, bool) {
	if o == nil || o.BillingAccountId == nil {
		return nil, false
	}
	return o.BillingAccountId, true
}

// HasBillingAccountId returns a boolean if a field has been set.
func (o *AssetWorkloadOptimizerMicrosoftAzureBillingOptionsAllOf) HasBillingAccountId() bool {
	if o != nil && o.BillingAccountId != nil {
		return true
	}

	return false
}

// SetBillingAccountId gets a reference to the given string and assigns it to the BillingAccountId field.
func (o *AssetWorkloadOptimizerMicrosoftAzureBillingOptionsAllOf) SetBillingAccountId(v string) {
	o.BillingAccountId = &v
}

// GetCostExportName returns the CostExportName field value if set, zero value otherwise.
func (o *AssetWorkloadOptimizerMicrosoftAzureBillingOptionsAllOf) GetCostExportName() string {
	if o == nil || o.CostExportName == nil {
		var ret string
		return ret
	}
	return *o.CostExportName
}

// GetCostExportNameOk returns a tuple with the CostExportName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerMicrosoftAzureBillingOptionsAllOf) GetCostExportNameOk() (*string, bool) {
	if o == nil || o.CostExportName == nil {
		return nil, false
	}
	return o.CostExportName, true
}

// HasCostExportName returns a boolean if a field has been set.
func (o *AssetWorkloadOptimizerMicrosoftAzureBillingOptionsAllOf) HasCostExportName() bool {
	if o != nil && o.CostExportName != nil {
		return true
	}

	return false
}

// SetCostExportName gets a reference to the given string and assigns it to the CostExportName field.
func (o *AssetWorkloadOptimizerMicrosoftAzureBillingOptionsAllOf) SetCostExportName(v string) {
	o.CostExportName = &v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *AssetWorkloadOptimizerMicrosoftAzureBillingOptionsAllOf) GetTenantId() string {
	if o == nil || o.TenantId == nil {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerMicrosoftAzureBillingOptionsAllOf) GetTenantIdOk() (*string, bool) {
	if o == nil || o.TenantId == nil {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *AssetWorkloadOptimizerMicrosoftAzureBillingOptionsAllOf) HasTenantId() bool {
	if o != nil && o.TenantId != nil {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *AssetWorkloadOptimizerMicrosoftAzureBillingOptionsAllOf) SetTenantId(v string) {
	o.TenantId = &v
}

func (o AssetWorkloadOptimizerMicrosoftAzureBillingOptionsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.BillingAccountId != nil {
		toSerialize["BillingAccountId"] = o.BillingAccountId
	}
	if o.CostExportName != nil {
		toSerialize["CostExportName"] = o.CostExportName
	}
	if o.TenantId != nil {
		toSerialize["TenantId"] = o.TenantId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetWorkloadOptimizerMicrosoftAzureBillingOptionsAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varAssetWorkloadOptimizerMicrosoftAzureBillingOptionsAllOf := _AssetWorkloadOptimizerMicrosoftAzureBillingOptionsAllOf{}

	if err = json.Unmarshal(bytes, &varAssetWorkloadOptimizerMicrosoftAzureBillingOptionsAllOf); err == nil {
		*o = AssetWorkloadOptimizerMicrosoftAzureBillingOptionsAllOf(varAssetWorkloadOptimizerMicrosoftAzureBillingOptionsAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BillingAccountId")
		delete(additionalProperties, "CostExportName")
		delete(additionalProperties, "TenantId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetWorkloadOptimizerMicrosoftAzureBillingOptionsAllOf struct {
	value *AssetWorkloadOptimizerMicrosoftAzureBillingOptionsAllOf
	isSet bool
}

func (v NullableAssetWorkloadOptimizerMicrosoftAzureBillingOptionsAllOf) Get() *AssetWorkloadOptimizerMicrosoftAzureBillingOptionsAllOf {
	return v.value
}

func (v *NullableAssetWorkloadOptimizerMicrosoftAzureBillingOptionsAllOf) Set(val *AssetWorkloadOptimizerMicrosoftAzureBillingOptionsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetWorkloadOptimizerMicrosoftAzureBillingOptionsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetWorkloadOptimizerMicrosoftAzureBillingOptionsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetWorkloadOptimizerMicrosoftAzureBillingOptionsAllOf(val *AssetWorkloadOptimizerMicrosoftAzureBillingOptionsAllOf) *NullableAssetWorkloadOptimizerMicrosoftAzureBillingOptionsAllOf {
	return &NullableAssetWorkloadOptimizerMicrosoftAzureBillingOptionsAllOf{value: val, isSet: true}
}

func (v NullableAssetWorkloadOptimizerMicrosoftAzureBillingOptionsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetWorkloadOptimizerMicrosoftAzureBillingOptionsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
