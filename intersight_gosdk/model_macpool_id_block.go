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

// MacpoolIdBlock A block of contiguous MAC addresses that are part of a pool.
type MacpoolIdBlock struct {
	PoolAbstractBlock
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                   `json:"ObjectType"`
	MacBlock             *MacpoolBlock            `json:"MacBlock,omitempty"`
	Pool                 *MacpoolPoolRelationship `json:"Pool,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MacpoolIdBlock MacpoolIdBlock

// NewMacpoolIdBlock instantiates a new MacpoolIdBlock object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMacpoolIdBlock(classId string, objectType string) *MacpoolIdBlock {
	this := MacpoolIdBlock{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewMacpoolIdBlockWithDefaults instantiates a new MacpoolIdBlock object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMacpoolIdBlockWithDefaults() *MacpoolIdBlock {
	this := MacpoolIdBlock{}
	var classId string = "macpool.IdBlock"
	this.ClassId = classId
	var objectType string = "macpool.IdBlock"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *MacpoolIdBlock) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *MacpoolIdBlock) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *MacpoolIdBlock) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *MacpoolIdBlock) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *MacpoolIdBlock) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *MacpoolIdBlock) SetObjectType(v string) {
	o.ObjectType = v
}

// GetMacBlock returns the MacBlock field value if set, zero value otherwise.
func (o *MacpoolIdBlock) GetMacBlock() MacpoolBlock {
	if o == nil || o.MacBlock == nil {
		var ret MacpoolBlock
		return ret
	}
	return *o.MacBlock
}

// GetMacBlockOk returns a tuple with the MacBlock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MacpoolIdBlock) GetMacBlockOk() (*MacpoolBlock, bool) {
	if o == nil || o.MacBlock == nil {
		return nil, false
	}
	return o.MacBlock, true
}

// HasMacBlock returns a boolean if a field has been set.
func (o *MacpoolIdBlock) HasMacBlock() bool {
	if o != nil && o.MacBlock != nil {
		return true
	}

	return false
}

// SetMacBlock gets a reference to the given MacpoolBlock and assigns it to the MacBlock field.
func (o *MacpoolIdBlock) SetMacBlock(v MacpoolBlock) {
	o.MacBlock = &v
}

// GetPool returns the Pool field value if set, zero value otherwise.
func (o *MacpoolIdBlock) GetPool() MacpoolPoolRelationship {
	if o == nil || o.Pool == nil {
		var ret MacpoolPoolRelationship
		return ret
	}
	return *o.Pool
}

// GetPoolOk returns a tuple with the Pool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MacpoolIdBlock) GetPoolOk() (*MacpoolPoolRelationship, bool) {
	if o == nil || o.Pool == nil {
		return nil, false
	}
	return o.Pool, true
}

// HasPool returns a boolean if a field has been set.
func (o *MacpoolIdBlock) HasPool() bool {
	if o != nil && o.Pool != nil {
		return true
	}

	return false
}

// SetPool gets a reference to the given MacpoolPoolRelationship and assigns it to the Pool field.
func (o *MacpoolIdBlock) SetPool(v MacpoolPoolRelationship) {
	o.Pool = &v
}

func (o MacpoolIdBlock) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPoolAbstractBlock, errPoolAbstractBlock := json.Marshal(o.PoolAbstractBlock)
	if errPoolAbstractBlock != nil {
		return []byte{}, errPoolAbstractBlock
	}
	errPoolAbstractBlock = json.Unmarshal([]byte(serializedPoolAbstractBlock), &toSerialize)
	if errPoolAbstractBlock != nil {
		return []byte{}, errPoolAbstractBlock
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.MacBlock != nil {
		toSerialize["MacBlock"] = o.MacBlock
	}
	if o.Pool != nil {
		toSerialize["Pool"] = o.Pool
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MacpoolIdBlock) UnmarshalJSON(bytes []byte) (err error) {
	type MacpoolIdBlockWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string                   `json:"ObjectType"`
		MacBlock   *MacpoolBlock            `json:"MacBlock,omitempty"`
		Pool       *MacpoolPoolRelationship `json:"Pool,omitempty"`
	}

	varMacpoolIdBlockWithoutEmbeddedStruct := MacpoolIdBlockWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varMacpoolIdBlockWithoutEmbeddedStruct)
	if err == nil {
		varMacpoolIdBlock := _MacpoolIdBlock{}
		varMacpoolIdBlock.ClassId = varMacpoolIdBlockWithoutEmbeddedStruct.ClassId
		varMacpoolIdBlock.ObjectType = varMacpoolIdBlockWithoutEmbeddedStruct.ObjectType
		varMacpoolIdBlock.MacBlock = varMacpoolIdBlockWithoutEmbeddedStruct.MacBlock
		varMacpoolIdBlock.Pool = varMacpoolIdBlockWithoutEmbeddedStruct.Pool
		*o = MacpoolIdBlock(varMacpoolIdBlock)
	} else {
		return err
	}

	varMacpoolIdBlock := _MacpoolIdBlock{}

	err = json.Unmarshal(bytes, &varMacpoolIdBlock)
	if err == nil {
		o.PoolAbstractBlock = varMacpoolIdBlock.PoolAbstractBlock
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "MacBlock")
		delete(additionalProperties, "Pool")

		// remove fields from embedded structs
		reflectPoolAbstractBlock := reflect.ValueOf(o.PoolAbstractBlock)
		for i := 0; i < reflectPoolAbstractBlock.Type().NumField(); i++ {
			t := reflectPoolAbstractBlock.Type().Field(i)

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

type NullableMacpoolIdBlock struct {
	value *MacpoolIdBlock
	isSet bool
}

func (v NullableMacpoolIdBlock) Get() *MacpoolIdBlock {
	return v.value
}

func (v *NullableMacpoolIdBlock) Set(val *MacpoolIdBlock) {
	v.value = val
	v.isSet = true
}

func (v NullableMacpoolIdBlock) IsSet() bool {
	return v.isSet
}

func (v *NullableMacpoolIdBlock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMacpoolIdBlock(val *MacpoolIdBlock) *NullableMacpoolIdBlock {
	return &NullableMacpoolIdBlock{value: val, isSet: true}
}

func (v NullableMacpoolIdBlock) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMacpoolIdBlock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
