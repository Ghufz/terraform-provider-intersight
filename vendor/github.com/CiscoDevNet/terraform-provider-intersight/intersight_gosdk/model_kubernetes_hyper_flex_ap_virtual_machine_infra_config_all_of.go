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

// KubernetesHyperFlexApVirtualMachineInfraConfigAllOf Definition of the list of properties defined in 'kubernetes.HyperFlexApVirtualMachineInfraConfig', excluding properties defined in parent classes.
type KubernetesHyperFlexApVirtualMachineInfraConfigAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Disk mode to use for volumes. * `Block` - It is a Block virtual disk. * `Filesystem` - It is a File system virtual disk. * `` - Disk mode is either unknown or not supported.
	DiskMode             *string `json:"DiskMode,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesHyperFlexApVirtualMachineInfraConfigAllOf KubernetesHyperFlexApVirtualMachineInfraConfigAllOf

// NewKubernetesHyperFlexApVirtualMachineInfraConfigAllOf instantiates a new KubernetesHyperFlexApVirtualMachineInfraConfigAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesHyperFlexApVirtualMachineInfraConfigAllOf(classId string, objectType string) *KubernetesHyperFlexApVirtualMachineInfraConfigAllOf {
	this := KubernetesHyperFlexApVirtualMachineInfraConfigAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var diskMode string = "Block"
	this.DiskMode = &diskMode
	return &this
}

// NewKubernetesHyperFlexApVirtualMachineInfraConfigAllOfWithDefaults instantiates a new KubernetesHyperFlexApVirtualMachineInfraConfigAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesHyperFlexApVirtualMachineInfraConfigAllOfWithDefaults() *KubernetesHyperFlexApVirtualMachineInfraConfigAllOf {
	this := KubernetesHyperFlexApVirtualMachineInfraConfigAllOf{}
	var classId string = "kubernetes.HyperFlexApVirtualMachineInfraConfig"
	this.ClassId = classId
	var objectType string = "kubernetes.HyperFlexApVirtualMachineInfraConfig"
	this.ObjectType = objectType
	var diskMode string = "Block"
	this.DiskMode = &diskMode
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesHyperFlexApVirtualMachineInfraConfigAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesHyperFlexApVirtualMachineInfraConfigAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesHyperFlexApVirtualMachineInfraConfigAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesHyperFlexApVirtualMachineInfraConfigAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesHyperFlexApVirtualMachineInfraConfigAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesHyperFlexApVirtualMachineInfraConfigAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDiskMode returns the DiskMode field value if set, zero value otherwise.
func (o *KubernetesHyperFlexApVirtualMachineInfraConfigAllOf) GetDiskMode() string {
	if o == nil || o.DiskMode == nil {
		var ret string
		return ret
	}
	return *o.DiskMode
}

// GetDiskModeOk returns a tuple with the DiskMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesHyperFlexApVirtualMachineInfraConfigAllOf) GetDiskModeOk() (*string, bool) {
	if o == nil || o.DiskMode == nil {
		return nil, false
	}
	return o.DiskMode, true
}

// HasDiskMode returns a boolean if a field has been set.
func (o *KubernetesHyperFlexApVirtualMachineInfraConfigAllOf) HasDiskMode() bool {
	if o != nil && o.DiskMode != nil {
		return true
	}

	return false
}

// SetDiskMode gets a reference to the given string and assigns it to the DiskMode field.
func (o *KubernetesHyperFlexApVirtualMachineInfraConfigAllOf) SetDiskMode(v string) {
	o.DiskMode = &v
}

func (o KubernetesHyperFlexApVirtualMachineInfraConfigAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DiskMode != nil {
		toSerialize["DiskMode"] = o.DiskMode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesHyperFlexApVirtualMachineInfraConfigAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varKubernetesHyperFlexApVirtualMachineInfraConfigAllOf := _KubernetesHyperFlexApVirtualMachineInfraConfigAllOf{}

	if err = json.Unmarshal(bytes, &varKubernetesHyperFlexApVirtualMachineInfraConfigAllOf); err == nil {
		*o = KubernetesHyperFlexApVirtualMachineInfraConfigAllOf(varKubernetesHyperFlexApVirtualMachineInfraConfigAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DiskMode")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKubernetesHyperFlexApVirtualMachineInfraConfigAllOf struct {
	value *KubernetesHyperFlexApVirtualMachineInfraConfigAllOf
	isSet bool
}

func (v NullableKubernetesHyperFlexApVirtualMachineInfraConfigAllOf) Get() *KubernetesHyperFlexApVirtualMachineInfraConfigAllOf {
	return v.value
}

func (v *NullableKubernetesHyperFlexApVirtualMachineInfraConfigAllOf) Set(val *KubernetesHyperFlexApVirtualMachineInfraConfigAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesHyperFlexApVirtualMachineInfraConfigAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesHyperFlexApVirtualMachineInfraConfigAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesHyperFlexApVirtualMachineInfraConfigAllOf(val *KubernetesHyperFlexApVirtualMachineInfraConfigAllOf) *NullableKubernetesHyperFlexApVirtualMachineInfraConfigAllOf {
	return &NullableKubernetesHyperFlexApVirtualMachineInfraConfigAllOf{value: val, isSet: true}
}

func (v NullableKubernetesHyperFlexApVirtualMachineInfraConfigAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesHyperFlexApVirtualMachineInfraConfigAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
