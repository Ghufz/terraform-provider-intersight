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
	"reflect"
	"strings"
)

// RproxyReverseProxy REST API endpoint for reverse proxy implementation. It receives an REST API request with URL and target type. Based on the target type, credentials is fetched from the target prior to being forwarded to the right upstream endpoint.
type RproxyReverseProxy struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                  `json:"ObjectType"`
	Account              *IamAccountRelationship `json:"Account,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RproxyReverseProxy RproxyReverseProxy

// NewRproxyReverseProxy instantiates a new RproxyReverseProxy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRproxyReverseProxy(classId string, objectType string) *RproxyReverseProxy {
	this := RproxyReverseProxy{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewRproxyReverseProxyWithDefaults instantiates a new RproxyReverseProxy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRproxyReverseProxyWithDefaults() *RproxyReverseProxy {
	this := RproxyReverseProxy{}
	var classId string = "rproxy.ReverseProxy"
	this.ClassId = classId
	var objectType string = "rproxy.ReverseProxy"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *RproxyReverseProxy) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *RproxyReverseProxy) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *RproxyReverseProxy) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *RproxyReverseProxy) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *RproxyReverseProxy) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *RproxyReverseProxy) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *RproxyReverseProxy) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RproxyReverseProxy) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *RproxyReverseProxy) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *RproxyReverseProxy) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

func (o RproxyReverseProxy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RproxyReverseProxy) UnmarshalJSON(bytes []byte) (err error) {
	type RproxyReverseProxyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string                  `json:"ObjectType"`
		Account    *IamAccountRelationship `json:"Account,omitempty"`
	}

	varRproxyReverseProxyWithoutEmbeddedStruct := RproxyReverseProxyWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varRproxyReverseProxyWithoutEmbeddedStruct)
	if err == nil {
		varRproxyReverseProxy := _RproxyReverseProxy{}
		varRproxyReverseProxy.ClassId = varRproxyReverseProxyWithoutEmbeddedStruct.ClassId
		varRproxyReverseProxy.ObjectType = varRproxyReverseProxyWithoutEmbeddedStruct.ObjectType
		varRproxyReverseProxy.Account = varRproxyReverseProxyWithoutEmbeddedStruct.Account
		*o = RproxyReverseProxy(varRproxyReverseProxy)
	} else {
		return err
	}

	varRproxyReverseProxy := _RproxyReverseProxy{}

	err = json.Unmarshal(bytes, &varRproxyReverseProxy)
	if err == nil {
		o.MoBaseMo = varRproxyReverseProxy.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Account")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

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

type NullableRproxyReverseProxy struct {
	value *RproxyReverseProxy
	isSet bool
}

func (v NullableRproxyReverseProxy) Get() *RproxyReverseProxy {
	return v.value
}

func (v *NullableRproxyReverseProxy) Set(val *RproxyReverseProxy) {
	v.value = val
	v.isSet = true
}

func (v NullableRproxyReverseProxy) IsSet() bool {
	return v.isSet
}

func (v *NullableRproxyReverseProxy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRproxyReverseProxy(val *RproxyReverseProxy) *NullableRproxyReverseProxy {
	return &NullableRproxyReverseProxy{value: val, isSet: true}
}

func (v NullableRproxyReverseProxy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRproxyReverseProxy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
