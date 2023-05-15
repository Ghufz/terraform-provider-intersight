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

// AssetWorkloadOptimizerAmazonWebServicesBillingOptions Captures configuration specific to the Amazon web service (AWS) target for Workload Optimizer service. Following explanation is for S3 parameters for configuring billing and costs reports. S3 bucket parameters help in programmatically gaining access to the S3 bucket for billing and cost management. The user must first have Amazon web service billing data being generated for their Amazon web service account. Since this is not enabled in Amazon web service by default, follow the instructions from Amazon web service to generate the bill hourly. Use \"Daily\" time unit frequency on which report data are measured and displayed. [link](https://aws.amazon.com/blogs/aws/new-programmatic-access-to-aws-billing-data/) [link](https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/billing-getting-started.html).
type AssetWorkloadOptimizerAmazonWebServicesBillingOptions struct {
	AssetServiceOptions
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Name of S3 bucket that contains the Amazon web service Cost and Usage report to get cloud service spend.
	CostAndUsageReportBucket *string `json:"CostAndUsageReportBucket,omitempty"`
	// Report path prefix for the Amazon web service cost and usage report to get cloud service spend.
	CostAndUsageReportPath *string `json:"CostAndUsageReportPath,omitempty"`
	// Region for S3 bucket that contains the Amazon web service Cost and Usage report to get cloud service spend.
	CostAndUsageReportRegion *string `json:"CostAndUsageReportRegion,omitempty"`
	AdditionalProperties     map[string]interface{}
}

type _AssetWorkloadOptimizerAmazonWebServicesBillingOptions AssetWorkloadOptimizerAmazonWebServicesBillingOptions

// NewAssetWorkloadOptimizerAmazonWebServicesBillingOptions instantiates a new AssetWorkloadOptimizerAmazonWebServicesBillingOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetWorkloadOptimizerAmazonWebServicesBillingOptions(classId string, objectType string) *AssetWorkloadOptimizerAmazonWebServicesBillingOptions {
	this := AssetWorkloadOptimizerAmazonWebServicesBillingOptions{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAssetWorkloadOptimizerAmazonWebServicesBillingOptionsWithDefaults instantiates a new AssetWorkloadOptimizerAmazonWebServicesBillingOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetWorkloadOptimizerAmazonWebServicesBillingOptionsWithDefaults() *AssetWorkloadOptimizerAmazonWebServicesBillingOptions {
	this := AssetWorkloadOptimizerAmazonWebServicesBillingOptions{}
	var classId string = "asset.WorkloadOptimizerAmazonWebServicesBillingOptions"
	this.ClassId = classId
	var objectType string = "asset.WorkloadOptimizerAmazonWebServicesBillingOptions"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetWorkloadOptimizerAmazonWebServicesBillingOptions) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerAmazonWebServicesBillingOptions) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetWorkloadOptimizerAmazonWebServicesBillingOptions) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AssetWorkloadOptimizerAmazonWebServicesBillingOptions) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerAmazonWebServicesBillingOptions) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetWorkloadOptimizerAmazonWebServicesBillingOptions) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCostAndUsageReportBucket returns the CostAndUsageReportBucket field value if set, zero value otherwise.
func (o *AssetWorkloadOptimizerAmazonWebServicesBillingOptions) GetCostAndUsageReportBucket() string {
	if o == nil || o.CostAndUsageReportBucket == nil {
		var ret string
		return ret
	}
	return *o.CostAndUsageReportBucket
}

// GetCostAndUsageReportBucketOk returns a tuple with the CostAndUsageReportBucket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerAmazonWebServicesBillingOptions) GetCostAndUsageReportBucketOk() (*string, bool) {
	if o == nil || o.CostAndUsageReportBucket == nil {
		return nil, false
	}
	return o.CostAndUsageReportBucket, true
}

// HasCostAndUsageReportBucket returns a boolean if a field has been set.
func (o *AssetWorkloadOptimizerAmazonWebServicesBillingOptions) HasCostAndUsageReportBucket() bool {
	if o != nil && o.CostAndUsageReportBucket != nil {
		return true
	}

	return false
}

// SetCostAndUsageReportBucket gets a reference to the given string and assigns it to the CostAndUsageReportBucket field.
func (o *AssetWorkloadOptimizerAmazonWebServicesBillingOptions) SetCostAndUsageReportBucket(v string) {
	o.CostAndUsageReportBucket = &v
}

// GetCostAndUsageReportPath returns the CostAndUsageReportPath field value if set, zero value otherwise.
func (o *AssetWorkloadOptimizerAmazonWebServicesBillingOptions) GetCostAndUsageReportPath() string {
	if o == nil || o.CostAndUsageReportPath == nil {
		var ret string
		return ret
	}
	return *o.CostAndUsageReportPath
}

// GetCostAndUsageReportPathOk returns a tuple with the CostAndUsageReportPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerAmazonWebServicesBillingOptions) GetCostAndUsageReportPathOk() (*string, bool) {
	if o == nil || o.CostAndUsageReportPath == nil {
		return nil, false
	}
	return o.CostAndUsageReportPath, true
}

// HasCostAndUsageReportPath returns a boolean if a field has been set.
func (o *AssetWorkloadOptimizerAmazonWebServicesBillingOptions) HasCostAndUsageReportPath() bool {
	if o != nil && o.CostAndUsageReportPath != nil {
		return true
	}

	return false
}

// SetCostAndUsageReportPath gets a reference to the given string and assigns it to the CostAndUsageReportPath field.
func (o *AssetWorkloadOptimizerAmazonWebServicesBillingOptions) SetCostAndUsageReportPath(v string) {
	o.CostAndUsageReportPath = &v
}

// GetCostAndUsageReportRegion returns the CostAndUsageReportRegion field value if set, zero value otherwise.
func (o *AssetWorkloadOptimizerAmazonWebServicesBillingOptions) GetCostAndUsageReportRegion() string {
	if o == nil || o.CostAndUsageReportRegion == nil {
		var ret string
		return ret
	}
	return *o.CostAndUsageReportRegion
}

// GetCostAndUsageReportRegionOk returns a tuple with the CostAndUsageReportRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerAmazonWebServicesBillingOptions) GetCostAndUsageReportRegionOk() (*string, bool) {
	if o == nil || o.CostAndUsageReportRegion == nil {
		return nil, false
	}
	return o.CostAndUsageReportRegion, true
}

// HasCostAndUsageReportRegion returns a boolean if a field has been set.
func (o *AssetWorkloadOptimizerAmazonWebServicesBillingOptions) HasCostAndUsageReportRegion() bool {
	if o != nil && o.CostAndUsageReportRegion != nil {
		return true
	}

	return false
}

// SetCostAndUsageReportRegion gets a reference to the given string and assigns it to the CostAndUsageReportRegion field.
func (o *AssetWorkloadOptimizerAmazonWebServicesBillingOptions) SetCostAndUsageReportRegion(v string) {
	o.CostAndUsageReportRegion = &v
}

func (o AssetWorkloadOptimizerAmazonWebServicesBillingOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedAssetServiceOptions, errAssetServiceOptions := json.Marshal(o.AssetServiceOptions)
	if errAssetServiceOptions != nil {
		return []byte{}, errAssetServiceOptions
	}
	errAssetServiceOptions = json.Unmarshal([]byte(serializedAssetServiceOptions), &toSerialize)
	if errAssetServiceOptions != nil {
		return []byte{}, errAssetServiceOptions
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.CostAndUsageReportBucket != nil {
		toSerialize["CostAndUsageReportBucket"] = o.CostAndUsageReportBucket
	}
	if o.CostAndUsageReportPath != nil {
		toSerialize["CostAndUsageReportPath"] = o.CostAndUsageReportPath
	}
	if o.CostAndUsageReportRegion != nil {
		toSerialize["CostAndUsageReportRegion"] = o.CostAndUsageReportRegion
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetWorkloadOptimizerAmazonWebServicesBillingOptions) UnmarshalJSON(bytes []byte) (err error) {
	type AssetWorkloadOptimizerAmazonWebServicesBillingOptionsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Name of S3 bucket that contains the Amazon web service Cost and Usage report to get cloud service spend.
		CostAndUsageReportBucket *string `json:"CostAndUsageReportBucket,omitempty"`
		// Report path prefix for the Amazon web service cost and usage report to get cloud service spend.
		CostAndUsageReportPath *string `json:"CostAndUsageReportPath,omitempty"`
		// Region for S3 bucket that contains the Amazon web service Cost and Usage report to get cloud service spend.
		CostAndUsageReportRegion *string `json:"CostAndUsageReportRegion,omitempty"`
	}

	varAssetWorkloadOptimizerAmazonWebServicesBillingOptionsWithoutEmbeddedStruct := AssetWorkloadOptimizerAmazonWebServicesBillingOptionsWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varAssetWorkloadOptimizerAmazonWebServicesBillingOptionsWithoutEmbeddedStruct)
	if err == nil {
		varAssetWorkloadOptimizerAmazonWebServicesBillingOptions := _AssetWorkloadOptimizerAmazonWebServicesBillingOptions{}
		varAssetWorkloadOptimizerAmazonWebServicesBillingOptions.ClassId = varAssetWorkloadOptimizerAmazonWebServicesBillingOptionsWithoutEmbeddedStruct.ClassId
		varAssetWorkloadOptimizerAmazonWebServicesBillingOptions.ObjectType = varAssetWorkloadOptimizerAmazonWebServicesBillingOptionsWithoutEmbeddedStruct.ObjectType
		varAssetWorkloadOptimizerAmazonWebServicesBillingOptions.CostAndUsageReportBucket = varAssetWorkloadOptimizerAmazonWebServicesBillingOptionsWithoutEmbeddedStruct.CostAndUsageReportBucket
		varAssetWorkloadOptimizerAmazonWebServicesBillingOptions.CostAndUsageReportPath = varAssetWorkloadOptimizerAmazonWebServicesBillingOptionsWithoutEmbeddedStruct.CostAndUsageReportPath
		varAssetWorkloadOptimizerAmazonWebServicesBillingOptions.CostAndUsageReportRegion = varAssetWorkloadOptimizerAmazonWebServicesBillingOptionsWithoutEmbeddedStruct.CostAndUsageReportRegion
		*o = AssetWorkloadOptimizerAmazonWebServicesBillingOptions(varAssetWorkloadOptimizerAmazonWebServicesBillingOptions)
	} else {
		return err
	}

	varAssetWorkloadOptimizerAmazonWebServicesBillingOptions := _AssetWorkloadOptimizerAmazonWebServicesBillingOptions{}

	err = json.Unmarshal(bytes, &varAssetWorkloadOptimizerAmazonWebServicesBillingOptions)
	if err == nil {
		o.AssetServiceOptions = varAssetWorkloadOptimizerAmazonWebServicesBillingOptions.AssetServiceOptions
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CostAndUsageReportBucket")
		delete(additionalProperties, "CostAndUsageReportPath")
		delete(additionalProperties, "CostAndUsageReportRegion")

		// remove fields from embedded structs
		reflectAssetServiceOptions := reflect.ValueOf(o.AssetServiceOptions)
		for i := 0; i < reflectAssetServiceOptions.Type().NumField(); i++ {
			t := reflectAssetServiceOptions.Type().Field(i)

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

type NullableAssetWorkloadOptimizerAmazonWebServicesBillingOptions struct {
	value *AssetWorkloadOptimizerAmazonWebServicesBillingOptions
	isSet bool
}

func (v NullableAssetWorkloadOptimizerAmazonWebServicesBillingOptions) Get() *AssetWorkloadOptimizerAmazonWebServicesBillingOptions {
	return v.value
}

func (v *NullableAssetWorkloadOptimizerAmazonWebServicesBillingOptions) Set(val *AssetWorkloadOptimizerAmazonWebServicesBillingOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetWorkloadOptimizerAmazonWebServicesBillingOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetWorkloadOptimizerAmazonWebServicesBillingOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetWorkloadOptimizerAmazonWebServicesBillingOptions(val *AssetWorkloadOptimizerAmazonWebServicesBillingOptions) *NullableAssetWorkloadOptimizerAmazonWebServicesBillingOptions {
	return &NullableAssetWorkloadOptimizerAmazonWebServicesBillingOptions{value: val, isSet: true}
}

func (v NullableAssetWorkloadOptimizerAmazonWebServicesBillingOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetWorkloadOptimizerAmazonWebServicesBillingOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
