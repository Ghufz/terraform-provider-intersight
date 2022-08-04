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

// KubernetesNetworkPolicy A policy specifying the CIDR for internal networks in a Kubernetes cluster like Pod network, and Service network.
type KubernetesNetworkPolicy struct {
	PolicyAbstractPolicy
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string                      `json:"ObjectType"`
	CniConfig  NullableKubernetesCniConfig `json:"CniConfig,omitempty"`
	// Supported CNI type. Currently we only support Calico. * `Calico` - Calico CNI plugin as described in https://github.com/projectcalico/cni-plugin. * `Aci` - Cisco ACI Container Network Interface plugin.
	CniType *string `json:"CniType,omitempty"`
	// CIDR block to allocate Pod network IP addresses from.
	PodNetworkCidr *string `json:"PodNetworkCidr,omitempty"`
	// CIDR block to allocate cluster service IP addresses from.
	ServiceCidr *string `json:"ServiceCidr,omitempty"`
	// An array of relationships to kubernetesClusterProfile resources.
	ClusterProfiles      []KubernetesClusterProfileRelationship `json:"ClusterProfiles,omitempty"`
	Organization         *OrganizationOrganizationRelationship  `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesNetworkPolicy KubernetesNetworkPolicy

// NewKubernetesNetworkPolicy instantiates a new KubernetesNetworkPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesNetworkPolicy(classId string, objectType string) *KubernetesNetworkPolicy {
	this := KubernetesNetworkPolicy{}
	this.ClassId = classId
	this.ObjectType = objectType
	var cniType string = "Calico"
	this.CniType = &cniType
	return &this
}

// NewKubernetesNetworkPolicyWithDefaults instantiates a new KubernetesNetworkPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesNetworkPolicyWithDefaults() *KubernetesNetworkPolicy {
	this := KubernetesNetworkPolicy{}
	var classId string = "kubernetes.NetworkPolicy"
	this.ClassId = classId
	var objectType string = "kubernetes.NetworkPolicy"
	this.ObjectType = objectType
	var cniType string = "Calico"
	this.CniType = &cniType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesNetworkPolicy) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesNetworkPolicy) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesNetworkPolicy) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesNetworkPolicy) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesNetworkPolicy) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesNetworkPolicy) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCniConfig returns the CniConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesNetworkPolicy) GetCniConfig() KubernetesCniConfig {
	if o == nil || o.CniConfig.Get() == nil {
		var ret KubernetesCniConfig
		return ret
	}
	return *o.CniConfig.Get()
}

// GetCniConfigOk returns a tuple with the CniConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNetworkPolicy) GetCniConfigOk() (*KubernetesCniConfig, bool) {
	if o == nil {
		return nil, false
	}
	return o.CniConfig.Get(), o.CniConfig.IsSet()
}

// HasCniConfig returns a boolean if a field has been set.
func (o *KubernetesNetworkPolicy) HasCniConfig() bool {
	if o != nil && o.CniConfig.IsSet() {
		return true
	}

	return false
}

// SetCniConfig gets a reference to the given NullableKubernetesCniConfig and assigns it to the CniConfig field.
func (o *KubernetesNetworkPolicy) SetCniConfig(v KubernetesCniConfig) {
	o.CniConfig.Set(&v)
}

// SetCniConfigNil sets the value for CniConfig to be an explicit nil
func (o *KubernetesNetworkPolicy) SetCniConfigNil() {
	o.CniConfig.Set(nil)
}

// UnsetCniConfig ensures that no value is present for CniConfig, not even an explicit nil
func (o *KubernetesNetworkPolicy) UnsetCniConfig() {
	o.CniConfig.Unset()
}

// GetCniType returns the CniType field value if set, zero value otherwise.
func (o *KubernetesNetworkPolicy) GetCniType() string {
	if o == nil || o.CniType == nil {
		var ret string
		return ret
	}
	return *o.CniType
}

// GetCniTypeOk returns a tuple with the CniType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNetworkPolicy) GetCniTypeOk() (*string, bool) {
	if o == nil || o.CniType == nil {
		return nil, false
	}
	return o.CniType, true
}

// HasCniType returns a boolean if a field has been set.
func (o *KubernetesNetworkPolicy) HasCniType() bool {
	if o != nil && o.CniType != nil {
		return true
	}

	return false
}

// SetCniType gets a reference to the given string and assigns it to the CniType field.
func (o *KubernetesNetworkPolicy) SetCniType(v string) {
	o.CniType = &v
}

// GetPodNetworkCidr returns the PodNetworkCidr field value if set, zero value otherwise.
func (o *KubernetesNetworkPolicy) GetPodNetworkCidr() string {
	if o == nil || o.PodNetworkCidr == nil {
		var ret string
		return ret
	}
	return *o.PodNetworkCidr
}

// GetPodNetworkCidrOk returns a tuple with the PodNetworkCidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNetworkPolicy) GetPodNetworkCidrOk() (*string, bool) {
	if o == nil || o.PodNetworkCidr == nil {
		return nil, false
	}
	return o.PodNetworkCidr, true
}

// HasPodNetworkCidr returns a boolean if a field has been set.
func (o *KubernetesNetworkPolicy) HasPodNetworkCidr() bool {
	if o != nil && o.PodNetworkCidr != nil {
		return true
	}

	return false
}

// SetPodNetworkCidr gets a reference to the given string and assigns it to the PodNetworkCidr field.
func (o *KubernetesNetworkPolicy) SetPodNetworkCidr(v string) {
	o.PodNetworkCidr = &v
}

// GetServiceCidr returns the ServiceCidr field value if set, zero value otherwise.
func (o *KubernetesNetworkPolicy) GetServiceCidr() string {
	if o == nil || o.ServiceCidr == nil {
		var ret string
		return ret
	}
	return *o.ServiceCidr
}

// GetServiceCidrOk returns a tuple with the ServiceCidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNetworkPolicy) GetServiceCidrOk() (*string, bool) {
	if o == nil || o.ServiceCidr == nil {
		return nil, false
	}
	return o.ServiceCidr, true
}

// HasServiceCidr returns a boolean if a field has been set.
func (o *KubernetesNetworkPolicy) HasServiceCidr() bool {
	if o != nil && o.ServiceCidr != nil {
		return true
	}

	return false
}

// SetServiceCidr gets a reference to the given string and assigns it to the ServiceCidr field.
func (o *KubernetesNetworkPolicy) SetServiceCidr(v string) {
	o.ServiceCidr = &v
}

// GetClusterProfiles returns the ClusterProfiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesNetworkPolicy) GetClusterProfiles() []KubernetesClusterProfileRelationship {
	if o == nil {
		var ret []KubernetesClusterProfileRelationship
		return ret
	}
	return o.ClusterProfiles
}

// GetClusterProfilesOk returns a tuple with the ClusterProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNetworkPolicy) GetClusterProfilesOk() ([]KubernetesClusterProfileRelationship, bool) {
	if o == nil || o.ClusterProfiles == nil {
		return nil, false
	}
	return o.ClusterProfiles, true
}

// HasClusterProfiles returns a boolean if a field has been set.
func (o *KubernetesNetworkPolicy) HasClusterProfiles() bool {
	if o != nil && o.ClusterProfiles != nil {
		return true
	}

	return false
}

// SetClusterProfiles gets a reference to the given []KubernetesClusterProfileRelationship and assigns it to the ClusterProfiles field.
func (o *KubernetesNetworkPolicy) SetClusterProfiles(v []KubernetesClusterProfileRelationship) {
	o.ClusterProfiles = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *KubernetesNetworkPolicy) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNetworkPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *KubernetesNetworkPolicy) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *KubernetesNetworkPolicy) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o KubernetesNetworkPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractPolicy, errPolicyAbstractPolicy := json.Marshal(o.PolicyAbstractPolicy)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	errPolicyAbstractPolicy = json.Unmarshal([]byte(serializedPolicyAbstractPolicy), &toSerialize)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.CniConfig.IsSet() {
		toSerialize["CniConfig"] = o.CniConfig.Get()
	}
	if o.CniType != nil {
		toSerialize["CniType"] = o.CniType
	}
	if o.PodNetworkCidr != nil {
		toSerialize["PodNetworkCidr"] = o.PodNetworkCidr
	}
	if o.ServiceCidr != nil {
		toSerialize["ServiceCidr"] = o.ServiceCidr
	}
	if o.ClusterProfiles != nil {
		toSerialize["ClusterProfiles"] = o.ClusterProfiles
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesNetworkPolicy) UnmarshalJSON(bytes []byte) (err error) {
	type KubernetesNetworkPolicyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string                      `json:"ObjectType"`
		CniConfig  NullableKubernetesCniConfig `json:"CniConfig,omitempty"`
		// Supported CNI type. Currently we only support Calico. * `Calico` - Calico CNI plugin as described in https://github.com/projectcalico/cni-plugin. * `Aci` - Cisco ACI Container Network Interface plugin.
		CniType *string `json:"CniType,omitempty"`
		// CIDR block to allocate Pod network IP addresses from.
		PodNetworkCidr *string `json:"PodNetworkCidr,omitempty"`
		// CIDR block to allocate cluster service IP addresses from.
		ServiceCidr *string `json:"ServiceCidr,omitempty"`
		// An array of relationships to kubernetesClusterProfile resources.
		ClusterProfiles []KubernetesClusterProfileRelationship `json:"ClusterProfiles,omitempty"`
		Organization    *OrganizationOrganizationRelationship  `json:"Organization,omitempty"`
	}

	varKubernetesNetworkPolicyWithoutEmbeddedStruct := KubernetesNetworkPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varKubernetesNetworkPolicyWithoutEmbeddedStruct)
	if err == nil {
		varKubernetesNetworkPolicy := _KubernetesNetworkPolicy{}
		varKubernetesNetworkPolicy.ClassId = varKubernetesNetworkPolicyWithoutEmbeddedStruct.ClassId
		varKubernetesNetworkPolicy.ObjectType = varKubernetesNetworkPolicyWithoutEmbeddedStruct.ObjectType
		varKubernetesNetworkPolicy.CniConfig = varKubernetesNetworkPolicyWithoutEmbeddedStruct.CniConfig
		varKubernetesNetworkPolicy.CniType = varKubernetesNetworkPolicyWithoutEmbeddedStruct.CniType
		varKubernetesNetworkPolicy.PodNetworkCidr = varKubernetesNetworkPolicyWithoutEmbeddedStruct.PodNetworkCidr
		varKubernetesNetworkPolicy.ServiceCidr = varKubernetesNetworkPolicyWithoutEmbeddedStruct.ServiceCidr
		varKubernetesNetworkPolicy.ClusterProfiles = varKubernetesNetworkPolicyWithoutEmbeddedStruct.ClusterProfiles
		varKubernetesNetworkPolicy.Organization = varKubernetesNetworkPolicyWithoutEmbeddedStruct.Organization
		*o = KubernetesNetworkPolicy(varKubernetesNetworkPolicy)
	} else {
		return err
	}

	varKubernetesNetworkPolicy := _KubernetesNetworkPolicy{}

	err = json.Unmarshal(bytes, &varKubernetesNetworkPolicy)
	if err == nil {
		o.PolicyAbstractPolicy = varKubernetesNetworkPolicy.PolicyAbstractPolicy
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CniConfig")
		delete(additionalProperties, "CniType")
		delete(additionalProperties, "PodNetworkCidr")
		delete(additionalProperties, "ServiceCidr")
		delete(additionalProperties, "ClusterProfiles")
		delete(additionalProperties, "Organization")

		// remove fields from embedded structs
		reflectPolicyAbstractPolicy := reflect.ValueOf(o.PolicyAbstractPolicy)
		for i := 0; i < reflectPolicyAbstractPolicy.Type().NumField(); i++ {
			t := reflectPolicyAbstractPolicy.Type().Field(i)

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

type NullableKubernetesNetworkPolicy struct {
	value *KubernetesNetworkPolicy
	isSet bool
}

func (v NullableKubernetesNetworkPolicy) Get() *KubernetesNetworkPolicy {
	return v.value
}

func (v *NullableKubernetesNetworkPolicy) Set(val *KubernetesNetworkPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesNetworkPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesNetworkPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesNetworkPolicy(val *KubernetesNetworkPolicy) *NullableKubernetesNetworkPolicy {
	return &NullableKubernetesNetworkPolicy{value: val, isSet: true}
}

func (v NullableKubernetesNetworkPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesNetworkPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
