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

// KubernetesClusterManagementConfig Configuration settings for a Kubernetes cluster.
type KubernetesClusterManagementConfig struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Indicates whether the value of the 'tacPasswd' property has been set.
	IsTacPasswdSet *bool `json:"IsTacPasswdSet,omitempty"`
	// Number of IP addresses to reserve for load balancer services.
	LoadBalancerCount *int64   `json:"LoadBalancerCount,omitempty"`
	LoadBalancers     []string `json:"LoadBalancers,omitempty"`
	// VIP for the cluster Kubernetes API server. If this is empty and a cluster IP pool is specified, it will be allocated from the IP pool.
	MasterVip *string  `json:"MasterVip,omitempty"`
	SshKeys   []string `json:"SshKeys,omitempty"`
	// Name of the user to SSH to nodes in a cluster.
	SshUser *string `json:"SshUser,omitempty"`
	// Hashed password of the TAC user.
	TacPasswd            *string `json:"TacPasswd,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesClusterManagementConfig KubernetesClusterManagementConfig

// NewKubernetesClusterManagementConfig instantiates a new KubernetesClusterManagementConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesClusterManagementConfig(classId string, objectType string) *KubernetesClusterManagementConfig {
	this := KubernetesClusterManagementConfig{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKubernetesClusterManagementConfigWithDefaults instantiates a new KubernetesClusterManagementConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesClusterManagementConfigWithDefaults() *KubernetesClusterManagementConfig {
	this := KubernetesClusterManagementConfig{}
	var classId string = "kubernetes.ClusterManagementConfig"
	this.ClassId = classId
	var objectType string = "kubernetes.ClusterManagementConfig"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesClusterManagementConfig) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesClusterManagementConfig) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesClusterManagementConfig) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesClusterManagementConfig) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesClusterManagementConfig) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesClusterManagementConfig) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIsTacPasswdSet returns the IsTacPasswdSet field value if set, zero value otherwise.
func (o *KubernetesClusterManagementConfig) GetIsTacPasswdSet() bool {
	if o == nil || o.IsTacPasswdSet == nil {
		var ret bool
		return ret
	}
	return *o.IsTacPasswdSet
}

// GetIsTacPasswdSetOk returns a tuple with the IsTacPasswdSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterManagementConfig) GetIsTacPasswdSetOk() (*bool, bool) {
	if o == nil || o.IsTacPasswdSet == nil {
		return nil, false
	}
	return o.IsTacPasswdSet, true
}

// HasIsTacPasswdSet returns a boolean if a field has been set.
func (o *KubernetesClusterManagementConfig) HasIsTacPasswdSet() bool {
	if o != nil && o.IsTacPasswdSet != nil {
		return true
	}

	return false
}

// SetIsTacPasswdSet gets a reference to the given bool and assigns it to the IsTacPasswdSet field.
func (o *KubernetesClusterManagementConfig) SetIsTacPasswdSet(v bool) {
	o.IsTacPasswdSet = &v
}

// GetLoadBalancerCount returns the LoadBalancerCount field value if set, zero value otherwise.
func (o *KubernetesClusterManagementConfig) GetLoadBalancerCount() int64 {
	if o == nil || o.LoadBalancerCount == nil {
		var ret int64
		return ret
	}
	return *o.LoadBalancerCount
}

// GetLoadBalancerCountOk returns a tuple with the LoadBalancerCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterManagementConfig) GetLoadBalancerCountOk() (*int64, bool) {
	if o == nil || o.LoadBalancerCount == nil {
		return nil, false
	}
	return o.LoadBalancerCount, true
}

// HasLoadBalancerCount returns a boolean if a field has been set.
func (o *KubernetesClusterManagementConfig) HasLoadBalancerCount() bool {
	if o != nil && o.LoadBalancerCount != nil {
		return true
	}

	return false
}

// SetLoadBalancerCount gets a reference to the given int64 and assigns it to the LoadBalancerCount field.
func (o *KubernetesClusterManagementConfig) SetLoadBalancerCount(v int64) {
	o.LoadBalancerCount = &v
}

// GetLoadBalancers returns the LoadBalancers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesClusterManagementConfig) GetLoadBalancers() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.LoadBalancers
}

// GetLoadBalancersOk returns a tuple with the LoadBalancers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesClusterManagementConfig) GetLoadBalancersOk() ([]string, bool) {
	if o == nil || o.LoadBalancers == nil {
		return nil, false
	}
	return o.LoadBalancers, true
}

// HasLoadBalancers returns a boolean if a field has been set.
func (o *KubernetesClusterManagementConfig) HasLoadBalancers() bool {
	if o != nil && o.LoadBalancers != nil {
		return true
	}

	return false
}

// SetLoadBalancers gets a reference to the given []string and assigns it to the LoadBalancers field.
func (o *KubernetesClusterManagementConfig) SetLoadBalancers(v []string) {
	o.LoadBalancers = v
}

// GetMasterVip returns the MasterVip field value if set, zero value otherwise.
func (o *KubernetesClusterManagementConfig) GetMasterVip() string {
	if o == nil || o.MasterVip == nil {
		var ret string
		return ret
	}
	return *o.MasterVip
}

// GetMasterVipOk returns a tuple with the MasterVip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterManagementConfig) GetMasterVipOk() (*string, bool) {
	if o == nil || o.MasterVip == nil {
		return nil, false
	}
	return o.MasterVip, true
}

// HasMasterVip returns a boolean if a field has been set.
func (o *KubernetesClusterManagementConfig) HasMasterVip() bool {
	if o != nil && o.MasterVip != nil {
		return true
	}

	return false
}

// SetMasterVip gets a reference to the given string and assigns it to the MasterVip field.
func (o *KubernetesClusterManagementConfig) SetMasterVip(v string) {
	o.MasterVip = &v
}

// GetSshKeys returns the SshKeys field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesClusterManagementConfig) GetSshKeys() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SshKeys
}

// GetSshKeysOk returns a tuple with the SshKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesClusterManagementConfig) GetSshKeysOk() ([]string, bool) {
	if o == nil || o.SshKeys == nil {
		return nil, false
	}
	return o.SshKeys, true
}

// HasSshKeys returns a boolean if a field has been set.
func (o *KubernetesClusterManagementConfig) HasSshKeys() bool {
	if o != nil && o.SshKeys != nil {
		return true
	}

	return false
}

// SetSshKeys gets a reference to the given []string and assigns it to the SshKeys field.
func (o *KubernetesClusterManagementConfig) SetSshKeys(v []string) {
	o.SshKeys = v
}

// GetSshUser returns the SshUser field value if set, zero value otherwise.
func (o *KubernetesClusterManagementConfig) GetSshUser() string {
	if o == nil || o.SshUser == nil {
		var ret string
		return ret
	}
	return *o.SshUser
}

// GetSshUserOk returns a tuple with the SshUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterManagementConfig) GetSshUserOk() (*string, bool) {
	if o == nil || o.SshUser == nil {
		return nil, false
	}
	return o.SshUser, true
}

// HasSshUser returns a boolean if a field has been set.
func (o *KubernetesClusterManagementConfig) HasSshUser() bool {
	if o != nil && o.SshUser != nil {
		return true
	}

	return false
}

// SetSshUser gets a reference to the given string and assigns it to the SshUser field.
func (o *KubernetesClusterManagementConfig) SetSshUser(v string) {
	o.SshUser = &v
}

// GetTacPasswd returns the TacPasswd field value if set, zero value otherwise.
func (o *KubernetesClusterManagementConfig) GetTacPasswd() string {
	if o == nil || o.TacPasswd == nil {
		var ret string
		return ret
	}
	return *o.TacPasswd
}

// GetTacPasswdOk returns a tuple with the TacPasswd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterManagementConfig) GetTacPasswdOk() (*string, bool) {
	if o == nil || o.TacPasswd == nil {
		return nil, false
	}
	return o.TacPasswd, true
}

// HasTacPasswd returns a boolean if a field has been set.
func (o *KubernetesClusterManagementConfig) HasTacPasswd() bool {
	if o != nil && o.TacPasswd != nil {
		return true
	}

	return false
}

// SetTacPasswd gets a reference to the given string and assigns it to the TacPasswd field.
func (o *KubernetesClusterManagementConfig) SetTacPasswd(v string) {
	o.TacPasswd = &v
}

func (o KubernetesClusterManagementConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.IsTacPasswdSet != nil {
		toSerialize["IsTacPasswdSet"] = o.IsTacPasswdSet
	}
	if o.LoadBalancerCount != nil {
		toSerialize["LoadBalancerCount"] = o.LoadBalancerCount
	}
	if o.LoadBalancers != nil {
		toSerialize["LoadBalancers"] = o.LoadBalancers
	}
	if o.MasterVip != nil {
		toSerialize["MasterVip"] = o.MasterVip
	}
	if o.SshKeys != nil {
		toSerialize["SshKeys"] = o.SshKeys
	}
	if o.SshUser != nil {
		toSerialize["SshUser"] = o.SshUser
	}
	if o.TacPasswd != nil {
		toSerialize["TacPasswd"] = o.TacPasswd
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesClusterManagementConfig) UnmarshalJSON(bytes []byte) (err error) {
	type KubernetesClusterManagementConfigWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Indicates whether the value of the 'tacPasswd' property has been set.
		IsTacPasswdSet *bool `json:"IsTacPasswdSet,omitempty"`
		// Number of IP addresses to reserve for load balancer services.
		LoadBalancerCount *int64   `json:"LoadBalancerCount,omitempty"`
		LoadBalancers     []string `json:"LoadBalancers,omitempty"`
		// VIP for the cluster Kubernetes API server. If this is empty and a cluster IP pool is specified, it will be allocated from the IP pool.
		MasterVip *string  `json:"MasterVip,omitempty"`
		SshKeys   []string `json:"SshKeys,omitempty"`
		// Name of the user to SSH to nodes in a cluster.
		SshUser *string `json:"SshUser,omitempty"`
		// Hashed password of the TAC user.
		TacPasswd *string `json:"TacPasswd,omitempty"`
	}

	varKubernetesClusterManagementConfigWithoutEmbeddedStruct := KubernetesClusterManagementConfigWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varKubernetesClusterManagementConfigWithoutEmbeddedStruct)
	if err == nil {
		varKubernetesClusterManagementConfig := _KubernetesClusterManagementConfig{}
		varKubernetesClusterManagementConfig.ClassId = varKubernetesClusterManagementConfigWithoutEmbeddedStruct.ClassId
		varKubernetesClusterManagementConfig.ObjectType = varKubernetesClusterManagementConfigWithoutEmbeddedStruct.ObjectType
		varKubernetesClusterManagementConfig.IsTacPasswdSet = varKubernetesClusterManagementConfigWithoutEmbeddedStruct.IsTacPasswdSet
		varKubernetesClusterManagementConfig.LoadBalancerCount = varKubernetesClusterManagementConfigWithoutEmbeddedStruct.LoadBalancerCount
		varKubernetesClusterManagementConfig.LoadBalancers = varKubernetesClusterManagementConfigWithoutEmbeddedStruct.LoadBalancers
		varKubernetesClusterManagementConfig.MasterVip = varKubernetesClusterManagementConfigWithoutEmbeddedStruct.MasterVip
		varKubernetesClusterManagementConfig.SshKeys = varKubernetesClusterManagementConfigWithoutEmbeddedStruct.SshKeys
		varKubernetesClusterManagementConfig.SshUser = varKubernetesClusterManagementConfigWithoutEmbeddedStruct.SshUser
		varKubernetesClusterManagementConfig.TacPasswd = varKubernetesClusterManagementConfigWithoutEmbeddedStruct.TacPasswd
		*o = KubernetesClusterManagementConfig(varKubernetesClusterManagementConfig)
	} else {
		return err
	}

	varKubernetesClusterManagementConfig := _KubernetesClusterManagementConfig{}

	err = json.Unmarshal(bytes, &varKubernetesClusterManagementConfig)
	if err == nil {
		o.MoBaseComplexType = varKubernetesClusterManagementConfig.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "IsTacPasswdSet")
		delete(additionalProperties, "LoadBalancerCount")
		delete(additionalProperties, "LoadBalancers")
		delete(additionalProperties, "MasterVip")
		delete(additionalProperties, "SshKeys")
		delete(additionalProperties, "SshUser")
		delete(additionalProperties, "TacPasswd")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

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

type NullableKubernetesClusterManagementConfig struct {
	value *KubernetesClusterManagementConfig
	isSet bool
}

func (v NullableKubernetesClusterManagementConfig) Get() *KubernetesClusterManagementConfig {
	return v.value
}

func (v *NullableKubernetesClusterManagementConfig) Set(val *KubernetesClusterManagementConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesClusterManagementConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesClusterManagementConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesClusterManagementConfig(val *KubernetesClusterManagementConfig) *NullableKubernetesClusterManagementConfig {
	return &NullableKubernetesClusterManagementConfig{value: val, isSet: true}
}

func (v NullableKubernetesClusterManagementConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesClusterManagementConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
