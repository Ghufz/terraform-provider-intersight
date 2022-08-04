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

// IppoolPoolMember PoolMember represents a single IPv4 and or IPv6 address that is part of a pool.
type IppoolPoolMember struct {
	PoolAbstractPoolMember
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Type of the IP address requested. * `IPv4` - IP V4 address type requested. * `IPv6` - IP V6 address type requested.
	IpType *string `json:"IpType,omitempty"`
	// IPv4 Address of this pool member.
	IpV4Address *string `json:"IpV4Address,omitempty"`
	// IPv6 Address of this pool member.
	IpV6Address          *string                        `json:"IpV6Address,omitempty"`
	AssignedToEntity     *MoBaseMoRelationship          `json:"AssignedToEntity,omitempty"`
	IpBlock              *IppoolShadowBlockRelationship `json:"IpBlock,omitempty"`
	Peer                 *IppoolIpLeaseRelationship     `json:"Peer,omitempty"`
	Pool                 *IppoolShadowPoolRelationship  `json:"Pool,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IppoolPoolMember IppoolPoolMember

// NewIppoolPoolMember instantiates a new IppoolPoolMember object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIppoolPoolMember(classId string, objectType string) *IppoolPoolMember {
	this := IppoolPoolMember{}
	this.ClassId = classId
	this.ObjectType = objectType
	var assigned bool = false
	this.Assigned = &assigned
	var ipType string = "IPv4"
	this.IpType = &ipType
	return &this
}

// NewIppoolPoolMemberWithDefaults instantiates a new IppoolPoolMember object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIppoolPoolMemberWithDefaults() *IppoolPoolMember {
	this := IppoolPoolMember{}
	var classId string = "ippool.PoolMember"
	this.ClassId = classId
	var objectType string = "ippool.PoolMember"
	this.ObjectType = objectType
	var ipType string = "IPv4"
	this.IpType = &ipType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IppoolPoolMember) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IppoolPoolMember) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IppoolPoolMember) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IppoolPoolMember) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IppoolPoolMember) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IppoolPoolMember) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIpType returns the IpType field value if set, zero value otherwise.
func (o *IppoolPoolMember) GetIpType() string {
	if o == nil || o.IpType == nil {
		var ret string
		return ret
	}
	return *o.IpType
}

// GetIpTypeOk returns a tuple with the IpType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolPoolMember) GetIpTypeOk() (*string, bool) {
	if o == nil || o.IpType == nil {
		return nil, false
	}
	return o.IpType, true
}

// HasIpType returns a boolean if a field has been set.
func (o *IppoolPoolMember) HasIpType() bool {
	if o != nil && o.IpType != nil {
		return true
	}

	return false
}

// SetIpType gets a reference to the given string and assigns it to the IpType field.
func (o *IppoolPoolMember) SetIpType(v string) {
	o.IpType = &v
}

// GetIpV4Address returns the IpV4Address field value if set, zero value otherwise.
func (o *IppoolPoolMember) GetIpV4Address() string {
	if o == nil || o.IpV4Address == nil {
		var ret string
		return ret
	}
	return *o.IpV4Address
}

// GetIpV4AddressOk returns a tuple with the IpV4Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolPoolMember) GetIpV4AddressOk() (*string, bool) {
	if o == nil || o.IpV4Address == nil {
		return nil, false
	}
	return o.IpV4Address, true
}

// HasIpV4Address returns a boolean if a field has been set.
func (o *IppoolPoolMember) HasIpV4Address() bool {
	if o != nil && o.IpV4Address != nil {
		return true
	}

	return false
}

// SetIpV4Address gets a reference to the given string and assigns it to the IpV4Address field.
func (o *IppoolPoolMember) SetIpV4Address(v string) {
	o.IpV4Address = &v
}

// GetIpV6Address returns the IpV6Address field value if set, zero value otherwise.
func (o *IppoolPoolMember) GetIpV6Address() string {
	if o == nil || o.IpV6Address == nil {
		var ret string
		return ret
	}
	return *o.IpV6Address
}

// GetIpV6AddressOk returns a tuple with the IpV6Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolPoolMember) GetIpV6AddressOk() (*string, bool) {
	if o == nil || o.IpV6Address == nil {
		return nil, false
	}
	return o.IpV6Address, true
}

// HasIpV6Address returns a boolean if a field has been set.
func (o *IppoolPoolMember) HasIpV6Address() bool {
	if o != nil && o.IpV6Address != nil {
		return true
	}

	return false
}

// SetIpV6Address gets a reference to the given string and assigns it to the IpV6Address field.
func (o *IppoolPoolMember) SetIpV6Address(v string) {
	o.IpV6Address = &v
}

// GetAssignedToEntity returns the AssignedToEntity field value if set, zero value otherwise.
func (o *IppoolPoolMember) GetAssignedToEntity() MoBaseMoRelationship {
	if o == nil || o.AssignedToEntity == nil {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.AssignedToEntity
}

// GetAssignedToEntityOk returns a tuple with the AssignedToEntity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolPoolMember) GetAssignedToEntityOk() (*MoBaseMoRelationship, bool) {
	if o == nil || o.AssignedToEntity == nil {
		return nil, false
	}
	return o.AssignedToEntity, true
}

// HasAssignedToEntity returns a boolean if a field has been set.
func (o *IppoolPoolMember) HasAssignedToEntity() bool {
	if o != nil && o.AssignedToEntity != nil {
		return true
	}

	return false
}

// SetAssignedToEntity gets a reference to the given MoBaseMoRelationship and assigns it to the AssignedToEntity field.
func (o *IppoolPoolMember) SetAssignedToEntity(v MoBaseMoRelationship) {
	o.AssignedToEntity = &v
}

// GetIpBlock returns the IpBlock field value if set, zero value otherwise.
func (o *IppoolPoolMember) GetIpBlock() IppoolShadowBlockRelationship {
	if o == nil || o.IpBlock == nil {
		var ret IppoolShadowBlockRelationship
		return ret
	}
	return *o.IpBlock
}

// GetIpBlockOk returns a tuple with the IpBlock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolPoolMember) GetIpBlockOk() (*IppoolShadowBlockRelationship, bool) {
	if o == nil || o.IpBlock == nil {
		return nil, false
	}
	return o.IpBlock, true
}

// HasIpBlock returns a boolean if a field has been set.
func (o *IppoolPoolMember) HasIpBlock() bool {
	if o != nil && o.IpBlock != nil {
		return true
	}

	return false
}

// SetIpBlock gets a reference to the given IppoolShadowBlockRelationship and assigns it to the IpBlock field.
func (o *IppoolPoolMember) SetIpBlock(v IppoolShadowBlockRelationship) {
	o.IpBlock = &v
}

// GetPeer returns the Peer field value if set, zero value otherwise.
func (o *IppoolPoolMember) GetPeer() IppoolIpLeaseRelationship {
	if o == nil || o.Peer == nil {
		var ret IppoolIpLeaseRelationship
		return ret
	}
	return *o.Peer
}

// GetPeerOk returns a tuple with the Peer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolPoolMember) GetPeerOk() (*IppoolIpLeaseRelationship, bool) {
	if o == nil || o.Peer == nil {
		return nil, false
	}
	return o.Peer, true
}

// HasPeer returns a boolean if a field has been set.
func (o *IppoolPoolMember) HasPeer() bool {
	if o != nil && o.Peer != nil {
		return true
	}

	return false
}

// SetPeer gets a reference to the given IppoolIpLeaseRelationship and assigns it to the Peer field.
func (o *IppoolPoolMember) SetPeer(v IppoolIpLeaseRelationship) {
	o.Peer = &v
}

// GetPool returns the Pool field value if set, zero value otherwise.
func (o *IppoolPoolMember) GetPool() IppoolShadowPoolRelationship {
	if o == nil || o.Pool == nil {
		var ret IppoolShadowPoolRelationship
		return ret
	}
	return *o.Pool
}

// GetPoolOk returns a tuple with the Pool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolPoolMember) GetPoolOk() (*IppoolShadowPoolRelationship, bool) {
	if o == nil || o.Pool == nil {
		return nil, false
	}
	return o.Pool, true
}

// HasPool returns a boolean if a field has been set.
func (o *IppoolPoolMember) HasPool() bool {
	if o != nil && o.Pool != nil {
		return true
	}

	return false
}

// SetPool gets a reference to the given IppoolShadowPoolRelationship and assigns it to the Pool field.
func (o *IppoolPoolMember) SetPool(v IppoolShadowPoolRelationship) {
	o.Pool = &v
}

func (o IppoolPoolMember) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPoolAbstractPoolMember, errPoolAbstractPoolMember := json.Marshal(o.PoolAbstractPoolMember)
	if errPoolAbstractPoolMember != nil {
		return []byte{}, errPoolAbstractPoolMember
	}
	errPoolAbstractPoolMember = json.Unmarshal([]byte(serializedPoolAbstractPoolMember), &toSerialize)
	if errPoolAbstractPoolMember != nil {
		return []byte{}, errPoolAbstractPoolMember
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.IpType != nil {
		toSerialize["IpType"] = o.IpType
	}
	if o.IpV4Address != nil {
		toSerialize["IpV4Address"] = o.IpV4Address
	}
	if o.IpV6Address != nil {
		toSerialize["IpV6Address"] = o.IpV6Address
	}
	if o.AssignedToEntity != nil {
		toSerialize["AssignedToEntity"] = o.AssignedToEntity
	}
	if o.IpBlock != nil {
		toSerialize["IpBlock"] = o.IpBlock
	}
	if o.Peer != nil {
		toSerialize["Peer"] = o.Peer
	}
	if o.Pool != nil {
		toSerialize["Pool"] = o.Pool
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IppoolPoolMember) UnmarshalJSON(bytes []byte) (err error) {
	type IppoolPoolMemberWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Type of the IP address requested. * `IPv4` - IP V4 address type requested. * `IPv6` - IP V6 address type requested.
		IpType *string `json:"IpType,omitempty"`
		// IPv4 Address of this pool member.
		IpV4Address *string `json:"IpV4Address,omitempty"`
		// IPv6 Address of this pool member.
		IpV6Address      *string                        `json:"IpV6Address,omitempty"`
		AssignedToEntity *MoBaseMoRelationship          `json:"AssignedToEntity,omitempty"`
		IpBlock          *IppoolShadowBlockRelationship `json:"IpBlock,omitempty"`
		Peer             *IppoolIpLeaseRelationship     `json:"Peer,omitempty"`
		Pool             *IppoolShadowPoolRelationship  `json:"Pool,omitempty"`
	}

	varIppoolPoolMemberWithoutEmbeddedStruct := IppoolPoolMemberWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varIppoolPoolMemberWithoutEmbeddedStruct)
	if err == nil {
		varIppoolPoolMember := _IppoolPoolMember{}
		varIppoolPoolMember.ClassId = varIppoolPoolMemberWithoutEmbeddedStruct.ClassId
		varIppoolPoolMember.ObjectType = varIppoolPoolMemberWithoutEmbeddedStruct.ObjectType
		varIppoolPoolMember.IpType = varIppoolPoolMemberWithoutEmbeddedStruct.IpType
		varIppoolPoolMember.IpV4Address = varIppoolPoolMemberWithoutEmbeddedStruct.IpV4Address
		varIppoolPoolMember.IpV6Address = varIppoolPoolMemberWithoutEmbeddedStruct.IpV6Address
		varIppoolPoolMember.AssignedToEntity = varIppoolPoolMemberWithoutEmbeddedStruct.AssignedToEntity
		varIppoolPoolMember.IpBlock = varIppoolPoolMemberWithoutEmbeddedStruct.IpBlock
		varIppoolPoolMember.Peer = varIppoolPoolMemberWithoutEmbeddedStruct.Peer
		varIppoolPoolMember.Pool = varIppoolPoolMemberWithoutEmbeddedStruct.Pool
		*o = IppoolPoolMember(varIppoolPoolMember)
	} else {
		return err
	}

	varIppoolPoolMember := _IppoolPoolMember{}

	err = json.Unmarshal(bytes, &varIppoolPoolMember)
	if err == nil {
		o.PoolAbstractPoolMember = varIppoolPoolMember.PoolAbstractPoolMember
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "IpType")
		delete(additionalProperties, "IpV4Address")
		delete(additionalProperties, "IpV6Address")
		delete(additionalProperties, "AssignedToEntity")
		delete(additionalProperties, "IpBlock")
		delete(additionalProperties, "Peer")
		delete(additionalProperties, "Pool")

		// remove fields from embedded structs
		reflectPoolAbstractPoolMember := reflect.ValueOf(o.PoolAbstractPoolMember)
		for i := 0; i < reflectPoolAbstractPoolMember.Type().NumField(); i++ {
			t := reflectPoolAbstractPoolMember.Type().Field(i)

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

type NullableIppoolPoolMember struct {
	value *IppoolPoolMember
	isSet bool
}

func (v NullableIppoolPoolMember) Get() *IppoolPoolMember {
	return v.value
}

func (v *NullableIppoolPoolMember) Set(val *IppoolPoolMember) {
	v.value = val
	v.isSet = true
}

func (v NullableIppoolPoolMember) IsSet() bool {
	return v.isSet
}

func (v *NullableIppoolPoolMember) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIppoolPoolMember(val *IppoolPoolMember) *NullableIppoolPoolMember {
	return &NullableIppoolPoolMember{value: val, isSet: true}
}

func (v NullableIppoolPoolMember) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIppoolPoolMember) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
