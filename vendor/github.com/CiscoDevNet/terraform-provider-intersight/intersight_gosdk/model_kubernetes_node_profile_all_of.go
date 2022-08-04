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

// KubernetesNodeProfileAllOf Definition of the list of properties defined in 'kubernetes.NodeProfile', excluding properties defined in parent classes.
type KubernetesNodeProfileAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Cloud provider for this node profile. * `noProvider` - Enables the use of no cloud provider. * `external` - Out of tree cloud provider, e.g. CPI for vsphere.
	CloudProvider        *string                                 `json:"CloudProvider,omitempty"`
	ConfigResult         *KubernetesConfigResultRelationship     `json:"ConfigResult,omitempty"`
	NodeGroup            *KubernetesNodeGroupProfileRelationship `json:"NodeGroup,omitempty"`
	Target               *AssetDeviceRegistrationRelationship    `json:"Target,omitempty"`
	Version              *KubernetesVersionRelationship          `json:"Version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesNodeProfileAllOf KubernetesNodeProfileAllOf

// NewKubernetesNodeProfileAllOf instantiates a new KubernetesNodeProfileAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesNodeProfileAllOf(classId string, objectType string) *KubernetesNodeProfileAllOf {
	this := KubernetesNodeProfileAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var cloudProvider string = "noProvider"
	this.CloudProvider = &cloudProvider
	return &this
}

// NewKubernetesNodeProfileAllOfWithDefaults instantiates a new KubernetesNodeProfileAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesNodeProfileAllOfWithDefaults() *KubernetesNodeProfileAllOf {
	this := KubernetesNodeProfileAllOf{}
	var cloudProvider string = "noProvider"
	this.CloudProvider = &cloudProvider
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesNodeProfileAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesNodeProfileAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesNodeProfileAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesNodeProfileAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesNodeProfileAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesNodeProfileAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCloudProvider returns the CloudProvider field value if set, zero value otherwise.
func (o *KubernetesNodeProfileAllOf) GetCloudProvider() string {
	if o == nil || o.CloudProvider == nil {
		var ret string
		return ret
	}
	return *o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeProfileAllOf) GetCloudProviderOk() (*string, bool) {
	if o == nil || o.CloudProvider == nil {
		return nil, false
	}
	return o.CloudProvider, true
}

// HasCloudProvider returns a boolean if a field has been set.
func (o *KubernetesNodeProfileAllOf) HasCloudProvider() bool {
	if o != nil && o.CloudProvider != nil {
		return true
	}

	return false
}

// SetCloudProvider gets a reference to the given string and assigns it to the CloudProvider field.
func (o *KubernetesNodeProfileAllOf) SetCloudProvider(v string) {
	o.CloudProvider = &v
}

// GetConfigResult returns the ConfigResult field value if set, zero value otherwise.
func (o *KubernetesNodeProfileAllOf) GetConfigResult() KubernetesConfigResultRelationship {
	if o == nil || o.ConfigResult == nil {
		var ret KubernetesConfigResultRelationship
		return ret
	}
	return *o.ConfigResult
}

// GetConfigResultOk returns a tuple with the ConfigResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeProfileAllOf) GetConfigResultOk() (*KubernetesConfigResultRelationship, bool) {
	if o == nil || o.ConfigResult == nil {
		return nil, false
	}
	return o.ConfigResult, true
}

// HasConfigResult returns a boolean if a field has been set.
func (o *KubernetesNodeProfileAllOf) HasConfigResult() bool {
	if o != nil && o.ConfigResult != nil {
		return true
	}

	return false
}

// SetConfigResult gets a reference to the given KubernetesConfigResultRelationship and assigns it to the ConfigResult field.
func (o *KubernetesNodeProfileAllOf) SetConfigResult(v KubernetesConfigResultRelationship) {
	o.ConfigResult = &v
}

// GetNodeGroup returns the NodeGroup field value if set, zero value otherwise.
func (o *KubernetesNodeProfileAllOf) GetNodeGroup() KubernetesNodeGroupProfileRelationship {
	if o == nil || o.NodeGroup == nil {
		var ret KubernetesNodeGroupProfileRelationship
		return ret
	}
	return *o.NodeGroup
}

// GetNodeGroupOk returns a tuple with the NodeGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeProfileAllOf) GetNodeGroupOk() (*KubernetesNodeGroupProfileRelationship, bool) {
	if o == nil || o.NodeGroup == nil {
		return nil, false
	}
	return o.NodeGroup, true
}

// HasNodeGroup returns a boolean if a field has been set.
func (o *KubernetesNodeProfileAllOf) HasNodeGroup() bool {
	if o != nil && o.NodeGroup != nil {
		return true
	}

	return false
}

// SetNodeGroup gets a reference to the given KubernetesNodeGroupProfileRelationship and assigns it to the NodeGroup field.
func (o *KubernetesNodeProfileAllOf) SetNodeGroup(v KubernetesNodeGroupProfileRelationship) {
	o.NodeGroup = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *KubernetesNodeProfileAllOf) GetTarget() AssetDeviceRegistrationRelationship {
	if o == nil || o.Target == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeProfileAllOf) GetTargetOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *KubernetesNodeProfileAllOf) HasTarget() bool {
	if o != nil && o.Target != nil {
		return true
	}

	return false
}

// SetTarget gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the Target field.
func (o *KubernetesNodeProfileAllOf) SetTarget(v AssetDeviceRegistrationRelationship) {
	o.Target = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *KubernetesNodeProfileAllOf) GetVersion() KubernetesVersionRelationship {
	if o == nil || o.Version == nil {
		var ret KubernetesVersionRelationship
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeProfileAllOf) GetVersionOk() (*KubernetesVersionRelationship, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *KubernetesNodeProfileAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given KubernetesVersionRelationship and assigns it to the Version field.
func (o *KubernetesNodeProfileAllOf) SetVersion(v KubernetesVersionRelationship) {
	o.Version = &v
}

func (o KubernetesNodeProfileAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.CloudProvider != nil {
		toSerialize["CloudProvider"] = o.CloudProvider
	}
	if o.ConfigResult != nil {
		toSerialize["ConfigResult"] = o.ConfigResult
	}
	if o.NodeGroup != nil {
		toSerialize["NodeGroup"] = o.NodeGroup
	}
	if o.Target != nil {
		toSerialize["Target"] = o.Target
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesNodeProfileAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varKubernetesNodeProfileAllOf := _KubernetesNodeProfileAllOf{}

	if err = json.Unmarshal(bytes, &varKubernetesNodeProfileAllOf); err == nil {
		*o = KubernetesNodeProfileAllOf(varKubernetesNodeProfileAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CloudProvider")
		delete(additionalProperties, "ConfigResult")
		delete(additionalProperties, "NodeGroup")
		delete(additionalProperties, "Target")
		delete(additionalProperties, "Version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKubernetesNodeProfileAllOf struct {
	value *KubernetesNodeProfileAllOf
	isSet bool
}

func (v NullableKubernetesNodeProfileAllOf) Get() *KubernetesNodeProfileAllOf {
	return v.value
}

func (v *NullableKubernetesNodeProfileAllOf) Set(val *KubernetesNodeProfileAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesNodeProfileAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesNodeProfileAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesNodeProfileAllOf(val *KubernetesNodeProfileAllOf) *NullableKubernetesNodeProfileAllOf {
	return &NullableKubernetesNodeProfileAllOf{value: val, isSet: true}
}

func (v NullableKubernetesNodeProfileAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesNodeProfileAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
