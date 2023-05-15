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

// KvmPolicyInventoryAllOf Definition of the list of properties defined in 'kvm.PolicyInventory', excluding properties defined in parent classes.
type KvmPolicyInventoryAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// If enabled, displays KVM session on any monitor attached to the server.
	EnableLocalServerVideo *bool `json:"EnableLocalServerVideo,omitempty"`
	// If enabled, encrypts all video information sent through KVM. Please note that this can no longer be disabled for servers running versions 4.2 and above.
	EnableVideoEncryption *bool `json:"EnableVideoEncryption,omitempty"`
	// State of the vKVM service on the endpoint.
	Enabled *bool `json:"Enabled,omitempty"`
	// The maximum number of concurrent KVM sessions allowed.
	MaximumSessions *int64 `json:"MaximumSessions,omitempty"`
	// The port used for KVM communication.
	RemotePort *int64 `json:"RemotePort,omitempty"`
	// Enables Tunneled vKVM on the endpoint. Applicable only for Device Connectors that support Tunneled vKVM.
	TunneledKvmEnabled   *bool                 `json:"TunneledKvmEnabled,omitempty"`
	TargetMo             *MoBaseMoRelationship `json:"TargetMo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KvmPolicyInventoryAllOf KvmPolicyInventoryAllOf

// NewKvmPolicyInventoryAllOf instantiates a new KvmPolicyInventoryAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKvmPolicyInventoryAllOf(classId string, objectType string) *KvmPolicyInventoryAllOf {
	this := KvmPolicyInventoryAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKvmPolicyInventoryAllOfWithDefaults instantiates a new KvmPolicyInventoryAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKvmPolicyInventoryAllOfWithDefaults() *KvmPolicyInventoryAllOf {
	this := KvmPolicyInventoryAllOf{}
	var classId string = "kvm.PolicyInventory"
	this.ClassId = classId
	var objectType string = "kvm.PolicyInventory"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KvmPolicyInventoryAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KvmPolicyInventoryAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KvmPolicyInventoryAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KvmPolicyInventoryAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KvmPolicyInventoryAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KvmPolicyInventoryAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEnableLocalServerVideo returns the EnableLocalServerVideo field value if set, zero value otherwise.
func (o *KvmPolicyInventoryAllOf) GetEnableLocalServerVideo() bool {
	if o == nil || o.EnableLocalServerVideo == nil {
		var ret bool
		return ret
	}
	return *o.EnableLocalServerVideo
}

// GetEnableLocalServerVideoOk returns a tuple with the EnableLocalServerVideo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvmPolicyInventoryAllOf) GetEnableLocalServerVideoOk() (*bool, bool) {
	if o == nil || o.EnableLocalServerVideo == nil {
		return nil, false
	}
	return o.EnableLocalServerVideo, true
}

// HasEnableLocalServerVideo returns a boolean if a field has been set.
func (o *KvmPolicyInventoryAllOf) HasEnableLocalServerVideo() bool {
	if o != nil && o.EnableLocalServerVideo != nil {
		return true
	}

	return false
}

// SetEnableLocalServerVideo gets a reference to the given bool and assigns it to the EnableLocalServerVideo field.
func (o *KvmPolicyInventoryAllOf) SetEnableLocalServerVideo(v bool) {
	o.EnableLocalServerVideo = &v
}

// GetEnableVideoEncryption returns the EnableVideoEncryption field value if set, zero value otherwise.
func (o *KvmPolicyInventoryAllOf) GetEnableVideoEncryption() bool {
	if o == nil || o.EnableVideoEncryption == nil {
		var ret bool
		return ret
	}
	return *o.EnableVideoEncryption
}

// GetEnableVideoEncryptionOk returns a tuple with the EnableVideoEncryption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvmPolicyInventoryAllOf) GetEnableVideoEncryptionOk() (*bool, bool) {
	if o == nil || o.EnableVideoEncryption == nil {
		return nil, false
	}
	return o.EnableVideoEncryption, true
}

// HasEnableVideoEncryption returns a boolean if a field has been set.
func (o *KvmPolicyInventoryAllOf) HasEnableVideoEncryption() bool {
	if o != nil && o.EnableVideoEncryption != nil {
		return true
	}

	return false
}

// SetEnableVideoEncryption gets a reference to the given bool and assigns it to the EnableVideoEncryption field.
func (o *KvmPolicyInventoryAllOf) SetEnableVideoEncryption(v bool) {
	o.EnableVideoEncryption = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *KvmPolicyInventoryAllOf) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvmPolicyInventoryAllOf) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *KvmPolicyInventoryAllOf) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *KvmPolicyInventoryAllOf) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetMaximumSessions returns the MaximumSessions field value if set, zero value otherwise.
func (o *KvmPolicyInventoryAllOf) GetMaximumSessions() int64 {
	if o == nil || o.MaximumSessions == nil {
		var ret int64
		return ret
	}
	return *o.MaximumSessions
}

// GetMaximumSessionsOk returns a tuple with the MaximumSessions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvmPolicyInventoryAllOf) GetMaximumSessionsOk() (*int64, bool) {
	if o == nil || o.MaximumSessions == nil {
		return nil, false
	}
	return o.MaximumSessions, true
}

// HasMaximumSessions returns a boolean if a field has been set.
func (o *KvmPolicyInventoryAllOf) HasMaximumSessions() bool {
	if o != nil && o.MaximumSessions != nil {
		return true
	}

	return false
}

// SetMaximumSessions gets a reference to the given int64 and assigns it to the MaximumSessions field.
func (o *KvmPolicyInventoryAllOf) SetMaximumSessions(v int64) {
	o.MaximumSessions = &v
}

// GetRemotePort returns the RemotePort field value if set, zero value otherwise.
func (o *KvmPolicyInventoryAllOf) GetRemotePort() int64 {
	if o == nil || o.RemotePort == nil {
		var ret int64
		return ret
	}
	return *o.RemotePort
}

// GetRemotePortOk returns a tuple with the RemotePort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvmPolicyInventoryAllOf) GetRemotePortOk() (*int64, bool) {
	if o == nil || o.RemotePort == nil {
		return nil, false
	}
	return o.RemotePort, true
}

// HasRemotePort returns a boolean if a field has been set.
func (o *KvmPolicyInventoryAllOf) HasRemotePort() bool {
	if o != nil && o.RemotePort != nil {
		return true
	}

	return false
}

// SetRemotePort gets a reference to the given int64 and assigns it to the RemotePort field.
func (o *KvmPolicyInventoryAllOf) SetRemotePort(v int64) {
	o.RemotePort = &v
}

// GetTunneledKvmEnabled returns the TunneledKvmEnabled field value if set, zero value otherwise.
func (o *KvmPolicyInventoryAllOf) GetTunneledKvmEnabled() bool {
	if o == nil || o.TunneledKvmEnabled == nil {
		var ret bool
		return ret
	}
	return *o.TunneledKvmEnabled
}

// GetTunneledKvmEnabledOk returns a tuple with the TunneledKvmEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvmPolicyInventoryAllOf) GetTunneledKvmEnabledOk() (*bool, bool) {
	if o == nil || o.TunneledKvmEnabled == nil {
		return nil, false
	}
	return o.TunneledKvmEnabled, true
}

// HasTunneledKvmEnabled returns a boolean if a field has been set.
func (o *KvmPolicyInventoryAllOf) HasTunneledKvmEnabled() bool {
	if o != nil && o.TunneledKvmEnabled != nil {
		return true
	}

	return false
}

// SetTunneledKvmEnabled gets a reference to the given bool and assigns it to the TunneledKvmEnabled field.
func (o *KvmPolicyInventoryAllOf) SetTunneledKvmEnabled(v bool) {
	o.TunneledKvmEnabled = &v
}

// GetTargetMo returns the TargetMo field value if set, zero value otherwise.
func (o *KvmPolicyInventoryAllOf) GetTargetMo() MoBaseMoRelationship {
	if o == nil || o.TargetMo == nil {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.TargetMo
}

// GetTargetMoOk returns a tuple with the TargetMo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvmPolicyInventoryAllOf) GetTargetMoOk() (*MoBaseMoRelationship, bool) {
	if o == nil || o.TargetMo == nil {
		return nil, false
	}
	return o.TargetMo, true
}

// HasTargetMo returns a boolean if a field has been set.
func (o *KvmPolicyInventoryAllOf) HasTargetMo() bool {
	if o != nil && o.TargetMo != nil {
		return true
	}

	return false
}

// SetTargetMo gets a reference to the given MoBaseMoRelationship and assigns it to the TargetMo field.
func (o *KvmPolicyInventoryAllOf) SetTargetMo(v MoBaseMoRelationship) {
	o.TargetMo = &v
}

func (o KvmPolicyInventoryAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.EnableLocalServerVideo != nil {
		toSerialize["EnableLocalServerVideo"] = o.EnableLocalServerVideo
	}
	if o.EnableVideoEncryption != nil {
		toSerialize["EnableVideoEncryption"] = o.EnableVideoEncryption
	}
	if o.Enabled != nil {
		toSerialize["Enabled"] = o.Enabled
	}
	if o.MaximumSessions != nil {
		toSerialize["MaximumSessions"] = o.MaximumSessions
	}
	if o.RemotePort != nil {
		toSerialize["RemotePort"] = o.RemotePort
	}
	if o.TunneledKvmEnabled != nil {
		toSerialize["TunneledKvmEnabled"] = o.TunneledKvmEnabled
	}
	if o.TargetMo != nil {
		toSerialize["TargetMo"] = o.TargetMo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KvmPolicyInventoryAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varKvmPolicyInventoryAllOf := _KvmPolicyInventoryAllOf{}

	if err = json.Unmarshal(bytes, &varKvmPolicyInventoryAllOf); err == nil {
		*o = KvmPolicyInventoryAllOf(varKvmPolicyInventoryAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "EnableLocalServerVideo")
		delete(additionalProperties, "EnableVideoEncryption")
		delete(additionalProperties, "Enabled")
		delete(additionalProperties, "MaximumSessions")
		delete(additionalProperties, "RemotePort")
		delete(additionalProperties, "TunneledKvmEnabled")
		delete(additionalProperties, "TargetMo")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKvmPolicyInventoryAllOf struct {
	value *KvmPolicyInventoryAllOf
	isSet bool
}

func (v NullableKvmPolicyInventoryAllOf) Get() *KvmPolicyInventoryAllOf {
	return v.value
}

func (v *NullableKvmPolicyInventoryAllOf) Set(val *KvmPolicyInventoryAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableKvmPolicyInventoryAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableKvmPolicyInventoryAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKvmPolicyInventoryAllOf(val *KvmPolicyInventoryAllOf) *NullableKvmPolicyInventoryAllOf {
	return &NullableKvmPolicyInventoryAllOf{value: val, isSet: true}
}

func (v NullableKvmPolicyInventoryAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKvmPolicyInventoryAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
