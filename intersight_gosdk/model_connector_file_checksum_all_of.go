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

// ConnectorFileChecksumAllOf Definition of the list of properties defined in 'connector.FileChecksum', excluding properties defined in parent classes.
type ConnectorFileChecksumAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The calculated hash of the contents using the algorithm.
	Hash *string `json:"Hash,omitempty"`
	// The hash algorithm used to calculate the checksum. * `crc` - A CRC hash as definded by RFC 3385. Generated with the IEEE polynomial. * `sha256` - An SHA256 hash as defined by RFC 4634.
	HashAlgorithm        *string `json:"HashAlgorithm,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConnectorFileChecksumAllOf ConnectorFileChecksumAllOf

// NewConnectorFileChecksumAllOf instantiates a new ConnectorFileChecksumAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorFileChecksumAllOf(classId string, objectType string) *ConnectorFileChecksumAllOf {
	this := ConnectorFileChecksumAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var hashAlgorithm string = "crc"
	this.HashAlgorithm = &hashAlgorithm
	return &this
}

// NewConnectorFileChecksumAllOfWithDefaults instantiates a new ConnectorFileChecksumAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorFileChecksumAllOfWithDefaults() *ConnectorFileChecksumAllOf {
	this := ConnectorFileChecksumAllOf{}
	var classId string = "connector.FileChecksum"
	this.ClassId = classId
	var objectType string = "connector.FileChecksum"
	this.ObjectType = objectType
	var hashAlgorithm string = "crc"
	this.HashAlgorithm = &hashAlgorithm
	return &this
}

// GetClassId returns the ClassId field value
func (o *ConnectorFileChecksumAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ConnectorFileChecksumAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ConnectorFileChecksumAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ConnectorFileChecksumAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ConnectorFileChecksumAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ConnectorFileChecksumAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetHash returns the Hash field value if set, zero value otherwise.
func (o *ConnectorFileChecksumAllOf) GetHash() string {
	if o == nil || o.Hash == nil {
		var ret string
		return ret
	}
	return *o.Hash
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorFileChecksumAllOf) GetHashOk() (*string, bool) {
	if o == nil || o.Hash == nil {
		return nil, false
	}
	return o.Hash, true
}

// HasHash returns a boolean if a field has been set.
func (o *ConnectorFileChecksumAllOf) HasHash() bool {
	if o != nil && o.Hash != nil {
		return true
	}

	return false
}

// SetHash gets a reference to the given string and assigns it to the Hash field.
func (o *ConnectorFileChecksumAllOf) SetHash(v string) {
	o.Hash = &v
}

// GetHashAlgorithm returns the HashAlgorithm field value if set, zero value otherwise.
func (o *ConnectorFileChecksumAllOf) GetHashAlgorithm() string {
	if o == nil || o.HashAlgorithm == nil {
		var ret string
		return ret
	}
	return *o.HashAlgorithm
}

// GetHashAlgorithmOk returns a tuple with the HashAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorFileChecksumAllOf) GetHashAlgorithmOk() (*string, bool) {
	if o == nil || o.HashAlgorithm == nil {
		return nil, false
	}
	return o.HashAlgorithm, true
}

// HasHashAlgorithm returns a boolean if a field has been set.
func (o *ConnectorFileChecksumAllOf) HasHashAlgorithm() bool {
	if o != nil && o.HashAlgorithm != nil {
		return true
	}

	return false
}

// SetHashAlgorithm gets a reference to the given string and assigns it to the HashAlgorithm field.
func (o *ConnectorFileChecksumAllOf) SetHashAlgorithm(v string) {
	o.HashAlgorithm = &v
}

func (o ConnectorFileChecksumAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Hash != nil {
		toSerialize["Hash"] = o.Hash
	}
	if o.HashAlgorithm != nil {
		toSerialize["HashAlgorithm"] = o.HashAlgorithm
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ConnectorFileChecksumAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varConnectorFileChecksumAllOf := _ConnectorFileChecksumAllOf{}

	if err = json.Unmarshal(bytes, &varConnectorFileChecksumAllOf); err == nil {
		*o = ConnectorFileChecksumAllOf(varConnectorFileChecksumAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Hash")
		delete(additionalProperties, "HashAlgorithm")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConnectorFileChecksumAllOf struct {
	value *ConnectorFileChecksumAllOf
	isSet bool
}

func (v NullableConnectorFileChecksumAllOf) Get() *ConnectorFileChecksumAllOf {
	return v.value
}

func (v *NullableConnectorFileChecksumAllOf) Set(val *ConnectorFileChecksumAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorFileChecksumAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorFileChecksumAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorFileChecksumAllOf(val *ConnectorFileChecksumAllOf) *NullableConnectorFileChecksumAllOf {
	return &NullableConnectorFileChecksumAllOf{value: val, isSet: true}
}

func (v NullableConnectorFileChecksumAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorFileChecksumAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
