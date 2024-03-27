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

// VirtualizationEsxiVmNetworkConfiguration Specify ESXi virtual machine network configuration data.
type VirtualizationEsxiVmNetworkConfiguration struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                           `json:"ObjectType"`
	Interfaces           []VirtualizationNetworkInterface `json:"Interfaces,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationEsxiVmNetworkConfiguration VirtualizationEsxiVmNetworkConfiguration

// NewVirtualizationEsxiVmNetworkConfiguration instantiates a new VirtualizationEsxiVmNetworkConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationEsxiVmNetworkConfiguration(classId string, objectType string) *VirtualizationEsxiVmNetworkConfiguration {
	this := VirtualizationEsxiVmNetworkConfiguration{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVirtualizationEsxiVmNetworkConfigurationWithDefaults instantiates a new VirtualizationEsxiVmNetworkConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationEsxiVmNetworkConfigurationWithDefaults() *VirtualizationEsxiVmNetworkConfiguration {
	this := VirtualizationEsxiVmNetworkConfiguration{}
	var classId string = "virtualization.EsxiVmNetworkConfiguration"
	this.ClassId = classId
	var objectType string = "virtualization.EsxiVmNetworkConfiguration"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationEsxiVmNetworkConfiguration) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationEsxiVmNetworkConfiguration) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationEsxiVmNetworkConfiguration) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationEsxiVmNetworkConfiguration) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationEsxiVmNetworkConfiguration) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationEsxiVmNetworkConfiguration) SetObjectType(v string) {
	o.ObjectType = v
}

// GetInterfaces returns the Interfaces field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationEsxiVmNetworkConfiguration) GetInterfaces() []VirtualizationNetworkInterface {
	if o == nil {
		var ret []VirtualizationNetworkInterface
		return ret
	}
	return o.Interfaces
}

// GetInterfacesOk returns a tuple with the Interfaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationEsxiVmNetworkConfiguration) GetInterfacesOk() ([]VirtualizationNetworkInterface, bool) {
	if o == nil || o.Interfaces == nil {
		return nil, false
	}
	return o.Interfaces, true
}

// HasInterfaces returns a boolean if a field has been set.
func (o *VirtualizationEsxiVmNetworkConfiguration) HasInterfaces() bool {
	if o != nil && o.Interfaces != nil {
		return true
	}

	return false
}

// SetInterfaces gets a reference to the given []VirtualizationNetworkInterface and assigns it to the Interfaces field.
func (o *VirtualizationEsxiVmNetworkConfiguration) SetInterfaces(v []VirtualizationNetworkInterface) {
	o.Interfaces = v
}

func (o VirtualizationEsxiVmNetworkConfiguration) MarshalJSON() ([]byte, error) {
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
	if o.Interfaces != nil {
		toSerialize["Interfaces"] = o.Interfaces
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationEsxiVmNetworkConfiguration) UnmarshalJSON(bytes []byte) (err error) {
	type VirtualizationEsxiVmNetworkConfigurationWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string                           `json:"ObjectType"`
		Interfaces []VirtualizationNetworkInterface `json:"Interfaces,omitempty"`
	}

	varVirtualizationEsxiVmNetworkConfigurationWithoutEmbeddedStruct := VirtualizationEsxiVmNetworkConfigurationWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVirtualizationEsxiVmNetworkConfigurationWithoutEmbeddedStruct)
	if err == nil {
		varVirtualizationEsxiVmNetworkConfiguration := _VirtualizationEsxiVmNetworkConfiguration{}
		varVirtualizationEsxiVmNetworkConfiguration.ClassId = varVirtualizationEsxiVmNetworkConfigurationWithoutEmbeddedStruct.ClassId
		varVirtualizationEsxiVmNetworkConfiguration.ObjectType = varVirtualizationEsxiVmNetworkConfigurationWithoutEmbeddedStruct.ObjectType
		varVirtualizationEsxiVmNetworkConfiguration.Interfaces = varVirtualizationEsxiVmNetworkConfigurationWithoutEmbeddedStruct.Interfaces
		*o = VirtualizationEsxiVmNetworkConfiguration(varVirtualizationEsxiVmNetworkConfiguration)
	} else {
		return err
	}

	varVirtualizationEsxiVmNetworkConfiguration := _VirtualizationEsxiVmNetworkConfiguration{}

	err = json.Unmarshal(bytes, &varVirtualizationEsxiVmNetworkConfiguration)
	if err == nil {
		o.MoBaseComplexType = varVirtualizationEsxiVmNetworkConfiguration.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Interfaces")

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

type NullableVirtualizationEsxiVmNetworkConfiguration struct {
	value *VirtualizationEsxiVmNetworkConfiguration
	isSet bool
}

func (v NullableVirtualizationEsxiVmNetworkConfiguration) Get() *VirtualizationEsxiVmNetworkConfiguration {
	return v.value
}

func (v *NullableVirtualizationEsxiVmNetworkConfiguration) Set(val *VirtualizationEsxiVmNetworkConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationEsxiVmNetworkConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationEsxiVmNetworkConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationEsxiVmNetworkConfiguration(val *VirtualizationEsxiVmNetworkConfiguration) *NullableVirtualizationEsxiVmNetworkConfiguration {
	return &NullableVirtualizationEsxiVmNetworkConfiguration{value: val, isSet: true}
}

func (v NullableVirtualizationEsxiVmNetworkConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationEsxiVmNetworkConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
