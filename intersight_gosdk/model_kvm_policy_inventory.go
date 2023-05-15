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

// KvmPolicyInventory Policy to configure KVM Launch settings.
type KvmPolicyInventory struct {
	PolicyAbstractPolicyInventory
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

type _KvmPolicyInventory KvmPolicyInventory

// NewKvmPolicyInventory instantiates a new KvmPolicyInventory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKvmPolicyInventory(classId string, objectType string) *KvmPolicyInventory {
	this := KvmPolicyInventory{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKvmPolicyInventoryWithDefaults instantiates a new KvmPolicyInventory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKvmPolicyInventoryWithDefaults() *KvmPolicyInventory {
	this := KvmPolicyInventory{}
	var classId string = "kvm.PolicyInventory"
	this.ClassId = classId
	var objectType string = "kvm.PolicyInventory"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KvmPolicyInventory) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KvmPolicyInventory) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KvmPolicyInventory) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KvmPolicyInventory) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KvmPolicyInventory) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KvmPolicyInventory) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEnableLocalServerVideo returns the EnableLocalServerVideo field value if set, zero value otherwise.
func (o *KvmPolicyInventory) GetEnableLocalServerVideo() bool {
	if o == nil || o.EnableLocalServerVideo == nil {
		var ret bool
		return ret
	}
	return *o.EnableLocalServerVideo
}

// GetEnableLocalServerVideoOk returns a tuple with the EnableLocalServerVideo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvmPolicyInventory) GetEnableLocalServerVideoOk() (*bool, bool) {
	if o == nil || o.EnableLocalServerVideo == nil {
		return nil, false
	}
	return o.EnableLocalServerVideo, true
}

// HasEnableLocalServerVideo returns a boolean if a field has been set.
func (o *KvmPolicyInventory) HasEnableLocalServerVideo() bool {
	if o != nil && o.EnableLocalServerVideo != nil {
		return true
	}

	return false
}

// SetEnableLocalServerVideo gets a reference to the given bool and assigns it to the EnableLocalServerVideo field.
func (o *KvmPolicyInventory) SetEnableLocalServerVideo(v bool) {
	o.EnableLocalServerVideo = &v
}

// GetEnableVideoEncryption returns the EnableVideoEncryption field value if set, zero value otherwise.
func (o *KvmPolicyInventory) GetEnableVideoEncryption() bool {
	if o == nil || o.EnableVideoEncryption == nil {
		var ret bool
		return ret
	}
	return *o.EnableVideoEncryption
}

// GetEnableVideoEncryptionOk returns a tuple with the EnableVideoEncryption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvmPolicyInventory) GetEnableVideoEncryptionOk() (*bool, bool) {
	if o == nil || o.EnableVideoEncryption == nil {
		return nil, false
	}
	return o.EnableVideoEncryption, true
}

// HasEnableVideoEncryption returns a boolean if a field has been set.
func (o *KvmPolicyInventory) HasEnableVideoEncryption() bool {
	if o != nil && o.EnableVideoEncryption != nil {
		return true
	}

	return false
}

// SetEnableVideoEncryption gets a reference to the given bool and assigns it to the EnableVideoEncryption field.
func (o *KvmPolicyInventory) SetEnableVideoEncryption(v bool) {
	o.EnableVideoEncryption = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *KvmPolicyInventory) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvmPolicyInventory) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *KvmPolicyInventory) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *KvmPolicyInventory) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetMaximumSessions returns the MaximumSessions field value if set, zero value otherwise.
func (o *KvmPolicyInventory) GetMaximumSessions() int64 {
	if o == nil || o.MaximumSessions == nil {
		var ret int64
		return ret
	}
	return *o.MaximumSessions
}

// GetMaximumSessionsOk returns a tuple with the MaximumSessions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvmPolicyInventory) GetMaximumSessionsOk() (*int64, bool) {
	if o == nil || o.MaximumSessions == nil {
		return nil, false
	}
	return o.MaximumSessions, true
}

// HasMaximumSessions returns a boolean if a field has been set.
func (o *KvmPolicyInventory) HasMaximumSessions() bool {
	if o != nil && o.MaximumSessions != nil {
		return true
	}

	return false
}

// SetMaximumSessions gets a reference to the given int64 and assigns it to the MaximumSessions field.
func (o *KvmPolicyInventory) SetMaximumSessions(v int64) {
	o.MaximumSessions = &v
}

// GetRemotePort returns the RemotePort field value if set, zero value otherwise.
func (o *KvmPolicyInventory) GetRemotePort() int64 {
	if o == nil || o.RemotePort == nil {
		var ret int64
		return ret
	}
	return *o.RemotePort
}

// GetRemotePortOk returns a tuple with the RemotePort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvmPolicyInventory) GetRemotePortOk() (*int64, bool) {
	if o == nil || o.RemotePort == nil {
		return nil, false
	}
	return o.RemotePort, true
}

// HasRemotePort returns a boolean if a field has been set.
func (o *KvmPolicyInventory) HasRemotePort() bool {
	if o != nil && o.RemotePort != nil {
		return true
	}

	return false
}

// SetRemotePort gets a reference to the given int64 and assigns it to the RemotePort field.
func (o *KvmPolicyInventory) SetRemotePort(v int64) {
	o.RemotePort = &v
}

// GetTunneledKvmEnabled returns the TunneledKvmEnabled field value if set, zero value otherwise.
func (o *KvmPolicyInventory) GetTunneledKvmEnabled() bool {
	if o == nil || o.TunneledKvmEnabled == nil {
		var ret bool
		return ret
	}
	return *o.TunneledKvmEnabled
}

// GetTunneledKvmEnabledOk returns a tuple with the TunneledKvmEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvmPolicyInventory) GetTunneledKvmEnabledOk() (*bool, bool) {
	if o == nil || o.TunneledKvmEnabled == nil {
		return nil, false
	}
	return o.TunneledKvmEnabled, true
}

// HasTunneledKvmEnabled returns a boolean if a field has been set.
func (o *KvmPolicyInventory) HasTunneledKvmEnabled() bool {
	if o != nil && o.TunneledKvmEnabled != nil {
		return true
	}

	return false
}

// SetTunneledKvmEnabled gets a reference to the given bool and assigns it to the TunneledKvmEnabled field.
func (o *KvmPolicyInventory) SetTunneledKvmEnabled(v bool) {
	o.TunneledKvmEnabled = &v
}

// GetTargetMo returns the TargetMo field value if set, zero value otherwise.
func (o *KvmPolicyInventory) GetTargetMo() MoBaseMoRelationship {
	if o == nil || o.TargetMo == nil {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.TargetMo
}

// GetTargetMoOk returns a tuple with the TargetMo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvmPolicyInventory) GetTargetMoOk() (*MoBaseMoRelationship, bool) {
	if o == nil || o.TargetMo == nil {
		return nil, false
	}
	return o.TargetMo, true
}

// HasTargetMo returns a boolean if a field has been set.
func (o *KvmPolicyInventory) HasTargetMo() bool {
	if o != nil && o.TargetMo != nil {
		return true
	}

	return false
}

// SetTargetMo gets a reference to the given MoBaseMoRelationship and assigns it to the TargetMo field.
func (o *KvmPolicyInventory) SetTargetMo(v MoBaseMoRelationship) {
	o.TargetMo = &v
}

func (o KvmPolicyInventory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractPolicyInventory, errPolicyAbstractPolicyInventory := json.Marshal(o.PolicyAbstractPolicyInventory)
	if errPolicyAbstractPolicyInventory != nil {
		return []byte{}, errPolicyAbstractPolicyInventory
	}
	errPolicyAbstractPolicyInventory = json.Unmarshal([]byte(serializedPolicyAbstractPolicyInventory), &toSerialize)
	if errPolicyAbstractPolicyInventory != nil {
		return []byte{}, errPolicyAbstractPolicyInventory
	}
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

func (o *KvmPolicyInventory) UnmarshalJSON(bytes []byte) (err error) {
	type KvmPolicyInventoryWithoutEmbeddedStruct struct {
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
		TunneledKvmEnabled *bool                 `json:"TunneledKvmEnabled,omitempty"`
		TargetMo           *MoBaseMoRelationship `json:"TargetMo,omitempty"`
	}

	varKvmPolicyInventoryWithoutEmbeddedStruct := KvmPolicyInventoryWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varKvmPolicyInventoryWithoutEmbeddedStruct)
	if err == nil {
		varKvmPolicyInventory := _KvmPolicyInventory{}
		varKvmPolicyInventory.ClassId = varKvmPolicyInventoryWithoutEmbeddedStruct.ClassId
		varKvmPolicyInventory.ObjectType = varKvmPolicyInventoryWithoutEmbeddedStruct.ObjectType
		varKvmPolicyInventory.EnableLocalServerVideo = varKvmPolicyInventoryWithoutEmbeddedStruct.EnableLocalServerVideo
		varKvmPolicyInventory.EnableVideoEncryption = varKvmPolicyInventoryWithoutEmbeddedStruct.EnableVideoEncryption
		varKvmPolicyInventory.Enabled = varKvmPolicyInventoryWithoutEmbeddedStruct.Enabled
		varKvmPolicyInventory.MaximumSessions = varKvmPolicyInventoryWithoutEmbeddedStruct.MaximumSessions
		varKvmPolicyInventory.RemotePort = varKvmPolicyInventoryWithoutEmbeddedStruct.RemotePort
		varKvmPolicyInventory.TunneledKvmEnabled = varKvmPolicyInventoryWithoutEmbeddedStruct.TunneledKvmEnabled
		varKvmPolicyInventory.TargetMo = varKvmPolicyInventoryWithoutEmbeddedStruct.TargetMo
		*o = KvmPolicyInventory(varKvmPolicyInventory)
	} else {
		return err
	}

	varKvmPolicyInventory := _KvmPolicyInventory{}

	err = json.Unmarshal(bytes, &varKvmPolicyInventory)
	if err == nil {
		o.PolicyAbstractPolicyInventory = varKvmPolicyInventory.PolicyAbstractPolicyInventory
	} else {
		return err
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

		// remove fields from embedded structs
		reflectPolicyAbstractPolicyInventory := reflect.ValueOf(o.PolicyAbstractPolicyInventory)
		for i := 0; i < reflectPolicyAbstractPolicyInventory.Type().NumField(); i++ {
			t := reflectPolicyAbstractPolicyInventory.Type().Field(i)

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

type NullableKvmPolicyInventory struct {
	value *KvmPolicyInventory
	isSet bool
}

func (v NullableKvmPolicyInventory) Get() *KvmPolicyInventory {
	return v.value
}

func (v *NullableKvmPolicyInventory) Set(val *KvmPolicyInventory) {
	v.value = val
	v.isSet = true
}

func (v NullableKvmPolicyInventory) IsSet() bool {
	return v.isSet
}

func (v *NullableKvmPolicyInventory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKvmPolicyInventory(val *KvmPolicyInventory) *NullableKvmPolicyInventory {
	return &NullableKvmPolicyInventory{value: val, isSet: true}
}

func (v NullableKvmPolicyInventory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKvmPolicyInventory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
