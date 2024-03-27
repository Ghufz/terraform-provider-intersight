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
)

// NetworkHyperFlexInternetProtocolAddressRangeAllOf Definition of the list of properties defined in 'network.HyperFlexInternetProtocolAddressRange', excluding properties defined in parent classes.
type NetworkHyperFlexInternetProtocolAddressRangeAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Begining IP address. IPv4 only.
	BeginAddress *string `json:"BeginAddress,omitempty"`
	// Ending IP address. IPv4 only.
	EndAddress           *string `json:"EndAddress,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NetworkHyperFlexInternetProtocolAddressRangeAllOf NetworkHyperFlexInternetProtocolAddressRangeAllOf

// NewNetworkHyperFlexInternetProtocolAddressRangeAllOf instantiates a new NetworkHyperFlexInternetProtocolAddressRangeAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkHyperFlexInternetProtocolAddressRangeAllOf(classId string, objectType string) *NetworkHyperFlexInternetProtocolAddressRangeAllOf {
	this := NetworkHyperFlexInternetProtocolAddressRangeAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNetworkHyperFlexInternetProtocolAddressRangeAllOfWithDefaults instantiates a new NetworkHyperFlexInternetProtocolAddressRangeAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkHyperFlexInternetProtocolAddressRangeAllOfWithDefaults() *NetworkHyperFlexInternetProtocolAddressRangeAllOf {
	this := NetworkHyperFlexInternetProtocolAddressRangeAllOf{}
	var classId string = "network.HyperFlexInternetProtocolAddressRange"
	this.ClassId = classId
	var objectType string = "network.HyperFlexInternetProtocolAddressRange"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NetworkHyperFlexInternetProtocolAddressRangeAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NetworkHyperFlexInternetProtocolAddressRangeAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NetworkHyperFlexInternetProtocolAddressRangeAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NetworkHyperFlexInternetProtocolAddressRangeAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NetworkHyperFlexInternetProtocolAddressRangeAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NetworkHyperFlexInternetProtocolAddressRangeAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBeginAddress returns the BeginAddress field value if set, zero value otherwise.
func (o *NetworkHyperFlexInternetProtocolAddressRangeAllOf) GetBeginAddress() string {
	if o == nil || o.BeginAddress == nil {
		var ret string
		return ret
	}
	return *o.BeginAddress
}

// GetBeginAddressOk returns a tuple with the BeginAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkHyperFlexInternetProtocolAddressRangeAllOf) GetBeginAddressOk() (*string, bool) {
	if o == nil || o.BeginAddress == nil {
		return nil, false
	}
	return o.BeginAddress, true
}

// HasBeginAddress returns a boolean if a field has been set.
func (o *NetworkHyperFlexInternetProtocolAddressRangeAllOf) HasBeginAddress() bool {
	if o != nil && o.BeginAddress != nil {
		return true
	}

	return false
}

// SetBeginAddress gets a reference to the given string and assigns it to the BeginAddress field.
func (o *NetworkHyperFlexInternetProtocolAddressRangeAllOf) SetBeginAddress(v string) {
	o.BeginAddress = &v
}

// GetEndAddress returns the EndAddress field value if set, zero value otherwise.
func (o *NetworkHyperFlexInternetProtocolAddressRangeAllOf) GetEndAddress() string {
	if o == nil || o.EndAddress == nil {
		var ret string
		return ret
	}
	return *o.EndAddress
}

// GetEndAddressOk returns a tuple with the EndAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkHyperFlexInternetProtocolAddressRangeAllOf) GetEndAddressOk() (*string, bool) {
	if o == nil || o.EndAddress == nil {
		return nil, false
	}
	return o.EndAddress, true
}

// HasEndAddress returns a boolean if a field has been set.
func (o *NetworkHyperFlexInternetProtocolAddressRangeAllOf) HasEndAddress() bool {
	if o != nil && o.EndAddress != nil {
		return true
	}

	return false
}

// SetEndAddress gets a reference to the given string and assigns it to the EndAddress field.
func (o *NetworkHyperFlexInternetProtocolAddressRangeAllOf) SetEndAddress(v string) {
	o.EndAddress = &v
}

func (o NetworkHyperFlexInternetProtocolAddressRangeAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.BeginAddress != nil {
		toSerialize["BeginAddress"] = o.BeginAddress
	}
	if o.EndAddress != nil {
		toSerialize["EndAddress"] = o.EndAddress
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NetworkHyperFlexInternetProtocolAddressRangeAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varNetworkHyperFlexInternetProtocolAddressRangeAllOf := _NetworkHyperFlexInternetProtocolAddressRangeAllOf{}

	if err = json.Unmarshal(bytes, &varNetworkHyperFlexInternetProtocolAddressRangeAllOf); err == nil {
		*o = NetworkHyperFlexInternetProtocolAddressRangeAllOf(varNetworkHyperFlexInternetProtocolAddressRangeAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BeginAddress")
		delete(additionalProperties, "EndAddress")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNetworkHyperFlexInternetProtocolAddressRangeAllOf struct {
	value *NetworkHyperFlexInternetProtocolAddressRangeAllOf
	isSet bool
}

func (v NullableNetworkHyperFlexInternetProtocolAddressRangeAllOf) Get() *NetworkHyperFlexInternetProtocolAddressRangeAllOf {
	return v.value
}

func (v *NullableNetworkHyperFlexInternetProtocolAddressRangeAllOf) Set(val *NetworkHyperFlexInternetProtocolAddressRangeAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkHyperFlexInternetProtocolAddressRangeAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkHyperFlexInternetProtocolAddressRangeAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkHyperFlexInternetProtocolAddressRangeAllOf(val *NetworkHyperFlexInternetProtocolAddressRangeAllOf) *NullableNetworkHyperFlexInternetProtocolAddressRangeAllOf {
	return &NullableNetworkHyperFlexInternetProtocolAddressRangeAllOf{value: val, isSet: true}
}

func (v NullableNetworkHyperFlexInternetProtocolAddressRangeAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkHyperFlexInternetProtocolAddressRangeAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
