/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-6484
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// KvmTunnelAllOf Definition of the list of properties defined in 'kvm.Tunnel', excluding properties defined in parent classes.
type KvmTunnelAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                               `json:"ObjectType"`
	Device               *AssetDeviceRegistrationRelationship `json:"Device,omitempty"`
	KvmSession           *KvmSessionRelationship              `json:"KvmSession,omitempty"`
	Server               *ComputePhysicalRelationship         `json:"Server,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KvmTunnelAllOf KvmTunnelAllOf

// NewKvmTunnelAllOf instantiates a new KvmTunnelAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKvmTunnelAllOf(classId string, objectType string) *KvmTunnelAllOf {
	this := KvmTunnelAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKvmTunnelAllOfWithDefaults instantiates a new KvmTunnelAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKvmTunnelAllOfWithDefaults() *KvmTunnelAllOf {
	this := KvmTunnelAllOf{}
	var classId string = "kvm.Tunnel"
	this.ClassId = classId
	var objectType string = "kvm.Tunnel"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KvmTunnelAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KvmTunnelAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KvmTunnelAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KvmTunnelAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KvmTunnelAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KvmTunnelAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *KvmTunnelAllOf) GetDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.Device == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvmTunnelAllOf) GetDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.Device == nil {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *KvmTunnelAllOf) HasDevice() bool {
	if o != nil && o.Device != nil {
		return true
	}

	return false
}

// SetDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the Device field.
func (o *KvmTunnelAllOf) SetDevice(v AssetDeviceRegistrationRelationship) {
	o.Device = &v
}

// GetKvmSession returns the KvmSession field value if set, zero value otherwise.
func (o *KvmTunnelAllOf) GetKvmSession() KvmSessionRelationship {
	if o == nil || o.KvmSession == nil {
		var ret KvmSessionRelationship
		return ret
	}
	return *o.KvmSession
}

// GetKvmSessionOk returns a tuple with the KvmSession field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvmTunnelAllOf) GetKvmSessionOk() (*KvmSessionRelationship, bool) {
	if o == nil || o.KvmSession == nil {
		return nil, false
	}
	return o.KvmSession, true
}

// HasKvmSession returns a boolean if a field has been set.
func (o *KvmTunnelAllOf) HasKvmSession() bool {
	if o != nil && o.KvmSession != nil {
		return true
	}

	return false
}

// SetKvmSession gets a reference to the given KvmSessionRelationship and assigns it to the KvmSession field.
func (o *KvmTunnelAllOf) SetKvmSession(v KvmSessionRelationship) {
	o.KvmSession = &v
}

// GetServer returns the Server field value if set, zero value otherwise.
func (o *KvmTunnelAllOf) GetServer() ComputePhysicalRelationship {
	if o == nil || o.Server == nil {
		var ret ComputePhysicalRelationship
		return ret
	}
	return *o.Server
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvmTunnelAllOf) GetServerOk() (*ComputePhysicalRelationship, bool) {
	if o == nil || o.Server == nil {
		return nil, false
	}
	return o.Server, true
}

// HasServer returns a boolean if a field has been set.
func (o *KvmTunnelAllOf) HasServer() bool {
	if o != nil && o.Server != nil {
		return true
	}

	return false
}

// SetServer gets a reference to the given ComputePhysicalRelationship and assigns it to the Server field.
func (o *KvmTunnelAllOf) SetServer(v ComputePhysicalRelationship) {
	o.Server = &v
}

func (o KvmTunnelAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Device != nil {
		toSerialize["Device"] = o.Device
	}
	if o.KvmSession != nil {
		toSerialize["KvmSession"] = o.KvmSession
	}
	if o.Server != nil {
		toSerialize["Server"] = o.Server
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KvmTunnelAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varKvmTunnelAllOf := _KvmTunnelAllOf{}

	if err = json.Unmarshal(bytes, &varKvmTunnelAllOf); err == nil {
		*o = KvmTunnelAllOf(varKvmTunnelAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Device")
		delete(additionalProperties, "KvmSession")
		delete(additionalProperties, "Server")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKvmTunnelAllOf struct {
	value *KvmTunnelAllOf
	isSet bool
}

func (v NullableKvmTunnelAllOf) Get() *KvmTunnelAllOf {
	return v.value
}

func (v *NullableKvmTunnelAllOf) Set(val *KvmTunnelAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableKvmTunnelAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableKvmTunnelAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKvmTunnelAllOf(val *KvmTunnelAllOf) *NullableKvmTunnelAllOf {
	return &NullableKvmTunnelAllOf{value: val, isSet: true}
}

func (v NullableKvmTunnelAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKvmTunnelAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
