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

// CloudNetworkAccessConfig Public IP and DNS (Domain Name System) information of Network interface.
type CloudNetworkAccessConfig struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType  string                `json:"ObjectType"`
	ExternalIps []CloudNetworkAddress `json:"ExternalIps,omitempty"`
	// Private DNS name assigned to the network interface.
	PrivateDns *string `json:"PrivateDns,omitempty"`
	// Public DNS name assigned to the network interface.
	PublicDns            *string `json:"PublicDns,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CloudNetworkAccessConfig CloudNetworkAccessConfig

// NewCloudNetworkAccessConfig instantiates a new CloudNetworkAccessConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudNetworkAccessConfig(classId string, objectType string) *CloudNetworkAccessConfig {
	this := CloudNetworkAccessConfig{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCloudNetworkAccessConfigWithDefaults instantiates a new CloudNetworkAccessConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudNetworkAccessConfigWithDefaults() *CloudNetworkAccessConfig {
	this := CloudNetworkAccessConfig{}
	var classId string = "cloud.NetworkAccessConfig"
	this.ClassId = classId
	var objectType string = "cloud.NetworkAccessConfig"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CloudNetworkAccessConfig) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CloudNetworkAccessConfig) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CloudNetworkAccessConfig) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CloudNetworkAccessConfig) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CloudNetworkAccessConfig) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CloudNetworkAccessConfig) SetObjectType(v string) {
	o.ObjectType = v
}

// GetExternalIps returns the ExternalIps field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudNetworkAccessConfig) GetExternalIps() []CloudNetworkAddress {
	if o == nil {
		var ret []CloudNetworkAddress
		return ret
	}
	return o.ExternalIps
}

// GetExternalIpsOk returns a tuple with the ExternalIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudNetworkAccessConfig) GetExternalIpsOk() ([]CloudNetworkAddress, bool) {
	if o == nil || o.ExternalIps == nil {
		return nil, false
	}
	return o.ExternalIps, true
}

// HasExternalIps returns a boolean if a field has been set.
func (o *CloudNetworkAccessConfig) HasExternalIps() bool {
	if o != nil && o.ExternalIps != nil {
		return true
	}

	return false
}

// SetExternalIps gets a reference to the given []CloudNetworkAddress and assigns it to the ExternalIps field.
func (o *CloudNetworkAccessConfig) SetExternalIps(v []CloudNetworkAddress) {
	o.ExternalIps = v
}

// GetPrivateDns returns the PrivateDns field value if set, zero value otherwise.
func (o *CloudNetworkAccessConfig) GetPrivateDns() string {
	if o == nil || o.PrivateDns == nil {
		var ret string
		return ret
	}
	return *o.PrivateDns
}

// GetPrivateDnsOk returns a tuple with the PrivateDns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudNetworkAccessConfig) GetPrivateDnsOk() (*string, bool) {
	if o == nil || o.PrivateDns == nil {
		return nil, false
	}
	return o.PrivateDns, true
}

// HasPrivateDns returns a boolean if a field has been set.
func (o *CloudNetworkAccessConfig) HasPrivateDns() bool {
	if o != nil && o.PrivateDns != nil {
		return true
	}

	return false
}

// SetPrivateDns gets a reference to the given string and assigns it to the PrivateDns field.
func (o *CloudNetworkAccessConfig) SetPrivateDns(v string) {
	o.PrivateDns = &v
}

// GetPublicDns returns the PublicDns field value if set, zero value otherwise.
func (o *CloudNetworkAccessConfig) GetPublicDns() string {
	if o == nil || o.PublicDns == nil {
		var ret string
		return ret
	}
	return *o.PublicDns
}

// GetPublicDnsOk returns a tuple with the PublicDns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudNetworkAccessConfig) GetPublicDnsOk() (*string, bool) {
	if o == nil || o.PublicDns == nil {
		return nil, false
	}
	return o.PublicDns, true
}

// HasPublicDns returns a boolean if a field has been set.
func (o *CloudNetworkAccessConfig) HasPublicDns() bool {
	if o != nil && o.PublicDns != nil {
		return true
	}

	return false
}

// SetPublicDns gets a reference to the given string and assigns it to the PublicDns field.
func (o *CloudNetworkAccessConfig) SetPublicDns(v string) {
	o.PublicDns = &v
}

func (o CloudNetworkAccessConfig) MarshalJSON() ([]byte, error) {
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
	if o.ExternalIps != nil {
		toSerialize["ExternalIps"] = o.ExternalIps
	}
	if o.PrivateDns != nil {
		toSerialize["PrivateDns"] = o.PrivateDns
	}
	if o.PublicDns != nil {
		toSerialize["PublicDns"] = o.PublicDns
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CloudNetworkAccessConfig) UnmarshalJSON(bytes []byte) (err error) {
	type CloudNetworkAccessConfigWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType  string                `json:"ObjectType"`
		ExternalIps []CloudNetworkAddress `json:"ExternalIps,omitempty"`
		// Private DNS name assigned to the network interface.
		PrivateDns *string `json:"PrivateDns,omitempty"`
		// Public DNS name assigned to the network interface.
		PublicDns *string `json:"PublicDns,omitempty"`
	}

	varCloudNetworkAccessConfigWithoutEmbeddedStruct := CloudNetworkAccessConfigWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varCloudNetworkAccessConfigWithoutEmbeddedStruct)
	if err == nil {
		varCloudNetworkAccessConfig := _CloudNetworkAccessConfig{}
		varCloudNetworkAccessConfig.ClassId = varCloudNetworkAccessConfigWithoutEmbeddedStruct.ClassId
		varCloudNetworkAccessConfig.ObjectType = varCloudNetworkAccessConfigWithoutEmbeddedStruct.ObjectType
		varCloudNetworkAccessConfig.ExternalIps = varCloudNetworkAccessConfigWithoutEmbeddedStruct.ExternalIps
		varCloudNetworkAccessConfig.PrivateDns = varCloudNetworkAccessConfigWithoutEmbeddedStruct.PrivateDns
		varCloudNetworkAccessConfig.PublicDns = varCloudNetworkAccessConfigWithoutEmbeddedStruct.PublicDns
		*o = CloudNetworkAccessConfig(varCloudNetworkAccessConfig)
	} else {
		return err
	}

	varCloudNetworkAccessConfig := _CloudNetworkAccessConfig{}

	err = json.Unmarshal(bytes, &varCloudNetworkAccessConfig)
	if err == nil {
		o.MoBaseComplexType = varCloudNetworkAccessConfig.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ExternalIps")
		delete(additionalProperties, "PrivateDns")
		delete(additionalProperties, "PublicDns")

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

type NullableCloudNetworkAccessConfig struct {
	value *CloudNetworkAccessConfig
	isSet bool
}

func (v NullableCloudNetworkAccessConfig) Get() *CloudNetworkAccessConfig {
	return v.value
}

func (v *NullableCloudNetworkAccessConfig) Set(val *CloudNetworkAccessConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudNetworkAccessConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudNetworkAccessConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudNetworkAccessConfig(val *CloudNetworkAccessConfig) *NullableCloudNetworkAccessConfig {
	return &NullableCloudNetworkAccessConfig{value: val, isSet: true}
}

func (v NullableCloudNetworkAccessConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudNetworkAccessConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
