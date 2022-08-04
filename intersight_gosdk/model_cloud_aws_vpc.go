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
	"reflect"
	"strings"
)

// CloudAwsVpc VPC (Virtual Private Cloud) object in AWS inventory.It is a service that lets you launch AWS resources in a logically isolated virtual network.
type CloudAwsVpc struct {
	CloudBasePlacement
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// If true, instances in the vpc get public DNS hostnames.
	DnsHostName *bool `json:"DnsHostName,omitempty"`
	// Indicates whether the Dns resolution is supported.
	DnsResolution *bool    `json:"DnsResolution,omitempty"`
	Ipv4Cidr      []string `json:"Ipv4Cidr,omitempty"`
	Ipv6Cidr      []string `json:"Ipv6Cidr,omitempty"`
	// If true, indicates that this is default VPC.
	IsDefault *bool `json:"IsDefault,omitempty"`
	// The state of the VPC (pending | available).
	State *string `json:"State,omitempty"`
	// The allowed tenancy of instances launched into the VPC.
	Tenancy              *string                          `json:"Tenancy,omitempty"`
	VpcTags              []CloudCloudTag                  `json:"VpcTags,omitempty"`
	AwsBillingUnit       *CloudAwsBillingUnitRelationship `json:"AwsBillingUnit,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CloudAwsVpc CloudAwsVpc

// NewCloudAwsVpc instantiates a new CloudAwsVpc object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudAwsVpc(classId string, objectType string) *CloudAwsVpc {
	this := CloudAwsVpc{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCloudAwsVpcWithDefaults instantiates a new CloudAwsVpc object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudAwsVpcWithDefaults() *CloudAwsVpc {
	this := CloudAwsVpc{}
	var classId string = "cloud.AwsVpc"
	this.ClassId = classId
	var objectType string = "cloud.AwsVpc"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CloudAwsVpc) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CloudAwsVpc) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CloudAwsVpc) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CloudAwsVpc) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CloudAwsVpc) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CloudAwsVpc) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDnsHostName returns the DnsHostName field value if set, zero value otherwise.
func (o *CloudAwsVpc) GetDnsHostName() bool {
	if o == nil || o.DnsHostName == nil {
		var ret bool
		return ret
	}
	return *o.DnsHostName
}

// GetDnsHostNameOk returns a tuple with the DnsHostName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudAwsVpc) GetDnsHostNameOk() (*bool, bool) {
	if o == nil || o.DnsHostName == nil {
		return nil, false
	}
	return o.DnsHostName, true
}

// HasDnsHostName returns a boolean if a field has been set.
func (o *CloudAwsVpc) HasDnsHostName() bool {
	if o != nil && o.DnsHostName != nil {
		return true
	}

	return false
}

// SetDnsHostName gets a reference to the given bool and assigns it to the DnsHostName field.
func (o *CloudAwsVpc) SetDnsHostName(v bool) {
	o.DnsHostName = &v
}

// GetDnsResolution returns the DnsResolution field value if set, zero value otherwise.
func (o *CloudAwsVpc) GetDnsResolution() bool {
	if o == nil || o.DnsResolution == nil {
		var ret bool
		return ret
	}
	return *o.DnsResolution
}

// GetDnsResolutionOk returns a tuple with the DnsResolution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudAwsVpc) GetDnsResolutionOk() (*bool, bool) {
	if o == nil || o.DnsResolution == nil {
		return nil, false
	}
	return o.DnsResolution, true
}

// HasDnsResolution returns a boolean if a field has been set.
func (o *CloudAwsVpc) HasDnsResolution() bool {
	if o != nil && o.DnsResolution != nil {
		return true
	}

	return false
}

// SetDnsResolution gets a reference to the given bool and assigns it to the DnsResolution field.
func (o *CloudAwsVpc) SetDnsResolution(v bool) {
	o.DnsResolution = &v
}

// GetIpv4Cidr returns the Ipv4Cidr field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudAwsVpc) GetIpv4Cidr() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Ipv4Cidr
}

// GetIpv4CidrOk returns a tuple with the Ipv4Cidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudAwsVpc) GetIpv4CidrOk() ([]string, bool) {
	if o == nil || o.Ipv4Cidr == nil {
		return nil, false
	}
	return o.Ipv4Cidr, true
}

// HasIpv4Cidr returns a boolean if a field has been set.
func (o *CloudAwsVpc) HasIpv4Cidr() bool {
	if o != nil && o.Ipv4Cidr != nil {
		return true
	}

	return false
}

// SetIpv4Cidr gets a reference to the given []string and assigns it to the Ipv4Cidr field.
func (o *CloudAwsVpc) SetIpv4Cidr(v []string) {
	o.Ipv4Cidr = v
}

// GetIpv6Cidr returns the Ipv6Cidr field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudAwsVpc) GetIpv6Cidr() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Ipv6Cidr
}

// GetIpv6CidrOk returns a tuple with the Ipv6Cidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudAwsVpc) GetIpv6CidrOk() ([]string, bool) {
	if o == nil || o.Ipv6Cidr == nil {
		return nil, false
	}
	return o.Ipv6Cidr, true
}

// HasIpv6Cidr returns a boolean if a field has been set.
func (o *CloudAwsVpc) HasIpv6Cidr() bool {
	if o != nil && o.Ipv6Cidr != nil {
		return true
	}

	return false
}

// SetIpv6Cidr gets a reference to the given []string and assigns it to the Ipv6Cidr field.
func (o *CloudAwsVpc) SetIpv6Cidr(v []string) {
	o.Ipv6Cidr = v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *CloudAwsVpc) GetIsDefault() bool {
	if o == nil || o.IsDefault == nil {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudAwsVpc) GetIsDefaultOk() (*bool, bool) {
	if o == nil || o.IsDefault == nil {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *CloudAwsVpc) HasIsDefault() bool {
	if o != nil && o.IsDefault != nil {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *CloudAwsVpc) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *CloudAwsVpc) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudAwsVpc) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *CloudAwsVpc) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *CloudAwsVpc) SetState(v string) {
	o.State = &v
}

// GetTenancy returns the Tenancy field value if set, zero value otherwise.
func (o *CloudAwsVpc) GetTenancy() string {
	if o == nil || o.Tenancy == nil {
		var ret string
		return ret
	}
	return *o.Tenancy
}

// GetTenancyOk returns a tuple with the Tenancy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudAwsVpc) GetTenancyOk() (*string, bool) {
	if o == nil || o.Tenancy == nil {
		return nil, false
	}
	return o.Tenancy, true
}

// HasTenancy returns a boolean if a field has been set.
func (o *CloudAwsVpc) HasTenancy() bool {
	if o != nil && o.Tenancy != nil {
		return true
	}

	return false
}

// SetTenancy gets a reference to the given string and assigns it to the Tenancy field.
func (o *CloudAwsVpc) SetTenancy(v string) {
	o.Tenancy = &v
}

// GetVpcTags returns the VpcTags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudAwsVpc) GetVpcTags() []CloudCloudTag {
	if o == nil {
		var ret []CloudCloudTag
		return ret
	}
	return o.VpcTags
}

// GetVpcTagsOk returns a tuple with the VpcTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudAwsVpc) GetVpcTagsOk() ([]CloudCloudTag, bool) {
	if o == nil || o.VpcTags == nil {
		return nil, false
	}
	return o.VpcTags, true
}

// HasVpcTags returns a boolean if a field has been set.
func (o *CloudAwsVpc) HasVpcTags() bool {
	if o != nil && o.VpcTags != nil {
		return true
	}

	return false
}

// SetVpcTags gets a reference to the given []CloudCloudTag and assigns it to the VpcTags field.
func (o *CloudAwsVpc) SetVpcTags(v []CloudCloudTag) {
	o.VpcTags = v
}

// GetAwsBillingUnit returns the AwsBillingUnit field value if set, zero value otherwise.
func (o *CloudAwsVpc) GetAwsBillingUnit() CloudAwsBillingUnitRelationship {
	if o == nil || o.AwsBillingUnit == nil {
		var ret CloudAwsBillingUnitRelationship
		return ret
	}
	return *o.AwsBillingUnit
}

// GetAwsBillingUnitOk returns a tuple with the AwsBillingUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudAwsVpc) GetAwsBillingUnitOk() (*CloudAwsBillingUnitRelationship, bool) {
	if o == nil || o.AwsBillingUnit == nil {
		return nil, false
	}
	return o.AwsBillingUnit, true
}

// HasAwsBillingUnit returns a boolean if a field has been set.
func (o *CloudAwsVpc) HasAwsBillingUnit() bool {
	if o != nil && o.AwsBillingUnit != nil {
		return true
	}

	return false
}

// SetAwsBillingUnit gets a reference to the given CloudAwsBillingUnitRelationship and assigns it to the AwsBillingUnit field.
func (o *CloudAwsVpc) SetAwsBillingUnit(v CloudAwsBillingUnitRelationship) {
	o.AwsBillingUnit = &v
}

func (o CloudAwsVpc) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedCloudBasePlacement, errCloudBasePlacement := json.Marshal(o.CloudBasePlacement)
	if errCloudBasePlacement != nil {
		return []byte{}, errCloudBasePlacement
	}
	errCloudBasePlacement = json.Unmarshal([]byte(serializedCloudBasePlacement), &toSerialize)
	if errCloudBasePlacement != nil {
		return []byte{}, errCloudBasePlacement
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DnsHostName != nil {
		toSerialize["DnsHostName"] = o.DnsHostName
	}
	if o.DnsResolution != nil {
		toSerialize["DnsResolution"] = o.DnsResolution
	}
	if o.Ipv4Cidr != nil {
		toSerialize["Ipv4Cidr"] = o.Ipv4Cidr
	}
	if o.Ipv6Cidr != nil {
		toSerialize["Ipv6Cidr"] = o.Ipv6Cidr
	}
	if o.IsDefault != nil {
		toSerialize["IsDefault"] = o.IsDefault
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}
	if o.Tenancy != nil {
		toSerialize["Tenancy"] = o.Tenancy
	}
	if o.VpcTags != nil {
		toSerialize["VpcTags"] = o.VpcTags
	}
	if o.AwsBillingUnit != nil {
		toSerialize["AwsBillingUnit"] = o.AwsBillingUnit
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CloudAwsVpc) UnmarshalJSON(bytes []byte) (err error) {
	type CloudAwsVpcWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// If true, instances in the vpc get public DNS hostnames.
		DnsHostName *bool `json:"DnsHostName,omitempty"`
		// Indicates whether the Dns resolution is supported.
		DnsResolution *bool    `json:"DnsResolution,omitempty"`
		Ipv4Cidr      []string `json:"Ipv4Cidr,omitempty"`
		Ipv6Cidr      []string `json:"Ipv6Cidr,omitempty"`
		// If true, indicates that this is default VPC.
		IsDefault *bool `json:"IsDefault,omitempty"`
		// The state of the VPC (pending | available).
		State *string `json:"State,omitempty"`
		// The allowed tenancy of instances launched into the VPC.
		Tenancy        *string                          `json:"Tenancy,omitempty"`
		VpcTags        []CloudCloudTag                  `json:"VpcTags,omitempty"`
		AwsBillingUnit *CloudAwsBillingUnitRelationship `json:"AwsBillingUnit,omitempty"`
	}

	varCloudAwsVpcWithoutEmbeddedStruct := CloudAwsVpcWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varCloudAwsVpcWithoutEmbeddedStruct)
	if err == nil {
		varCloudAwsVpc := _CloudAwsVpc{}
		varCloudAwsVpc.ClassId = varCloudAwsVpcWithoutEmbeddedStruct.ClassId
		varCloudAwsVpc.ObjectType = varCloudAwsVpcWithoutEmbeddedStruct.ObjectType
		varCloudAwsVpc.DnsHostName = varCloudAwsVpcWithoutEmbeddedStruct.DnsHostName
		varCloudAwsVpc.DnsResolution = varCloudAwsVpcWithoutEmbeddedStruct.DnsResolution
		varCloudAwsVpc.Ipv4Cidr = varCloudAwsVpcWithoutEmbeddedStruct.Ipv4Cidr
		varCloudAwsVpc.Ipv6Cidr = varCloudAwsVpcWithoutEmbeddedStruct.Ipv6Cidr
		varCloudAwsVpc.IsDefault = varCloudAwsVpcWithoutEmbeddedStruct.IsDefault
		varCloudAwsVpc.State = varCloudAwsVpcWithoutEmbeddedStruct.State
		varCloudAwsVpc.Tenancy = varCloudAwsVpcWithoutEmbeddedStruct.Tenancy
		varCloudAwsVpc.VpcTags = varCloudAwsVpcWithoutEmbeddedStruct.VpcTags
		varCloudAwsVpc.AwsBillingUnit = varCloudAwsVpcWithoutEmbeddedStruct.AwsBillingUnit
		*o = CloudAwsVpc(varCloudAwsVpc)
	} else {
		return err
	}

	varCloudAwsVpc := _CloudAwsVpc{}

	err = json.Unmarshal(bytes, &varCloudAwsVpc)
	if err == nil {
		o.CloudBasePlacement = varCloudAwsVpc.CloudBasePlacement
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DnsHostName")
		delete(additionalProperties, "DnsResolution")
		delete(additionalProperties, "Ipv4Cidr")
		delete(additionalProperties, "Ipv6Cidr")
		delete(additionalProperties, "IsDefault")
		delete(additionalProperties, "State")
		delete(additionalProperties, "Tenancy")
		delete(additionalProperties, "VpcTags")
		delete(additionalProperties, "AwsBillingUnit")

		// remove fields from embedded structs
		reflectCloudBasePlacement := reflect.ValueOf(o.CloudBasePlacement)
		for i := 0; i < reflectCloudBasePlacement.Type().NumField(); i++ {
			t := reflectCloudBasePlacement.Type().Field(i)

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

type NullableCloudAwsVpc struct {
	value *CloudAwsVpc
	isSet bool
}

func (v NullableCloudAwsVpc) Get() *CloudAwsVpc {
	return v.value
}

func (v *NullableCloudAwsVpc) Set(val *CloudAwsVpc) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudAwsVpc) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudAwsVpc) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudAwsVpc(val *CloudAwsVpc) *NullableCloudAwsVpc {
	return &NullableCloudAwsVpc{value: val, isSet: true}
}

func (v NullableCloudAwsVpc) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudAwsVpc) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
