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

// KubernetesNvidiaGpuProduct Information of a Nvidia GPU product.
type KubernetesNvidiaGpuProduct struct {
	KubernetesBaseGpuProduct
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// True if this Nvidia GPU supports MIG.
	MigCapable           *bool    `json:"MigCapable,omitempty"`
	MigProfiles          []string `json:"MigProfiles,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesNvidiaGpuProduct KubernetesNvidiaGpuProduct

// NewKubernetesNvidiaGpuProduct instantiates a new KubernetesNvidiaGpuProduct object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesNvidiaGpuProduct(classId string, objectType string) *KubernetesNvidiaGpuProduct {
	this := KubernetesNvidiaGpuProduct{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKubernetesNvidiaGpuProductWithDefaults instantiates a new KubernetesNvidiaGpuProduct object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesNvidiaGpuProductWithDefaults() *KubernetesNvidiaGpuProduct {
	this := KubernetesNvidiaGpuProduct{}
	var classId string = "kubernetes.NvidiaGpuProduct"
	this.ClassId = classId
	var objectType string = "kubernetes.NvidiaGpuProduct"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesNvidiaGpuProduct) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesNvidiaGpuProduct) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesNvidiaGpuProduct) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesNvidiaGpuProduct) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesNvidiaGpuProduct) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesNvidiaGpuProduct) SetObjectType(v string) {
	o.ObjectType = v
}

// GetMigCapable returns the MigCapable field value if set, zero value otherwise.
func (o *KubernetesNvidiaGpuProduct) GetMigCapable() bool {
	if o == nil || o.MigCapable == nil {
		var ret bool
		return ret
	}
	return *o.MigCapable
}

// GetMigCapableOk returns a tuple with the MigCapable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNvidiaGpuProduct) GetMigCapableOk() (*bool, bool) {
	if o == nil || o.MigCapable == nil {
		return nil, false
	}
	return o.MigCapable, true
}

// HasMigCapable returns a boolean if a field has been set.
func (o *KubernetesNvidiaGpuProduct) HasMigCapable() bool {
	if o != nil && o.MigCapable != nil {
		return true
	}

	return false
}

// SetMigCapable gets a reference to the given bool and assigns it to the MigCapable field.
func (o *KubernetesNvidiaGpuProduct) SetMigCapable(v bool) {
	o.MigCapable = &v
}

// GetMigProfiles returns the MigProfiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesNvidiaGpuProduct) GetMigProfiles() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.MigProfiles
}

// GetMigProfilesOk returns a tuple with the MigProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNvidiaGpuProduct) GetMigProfilesOk() ([]string, bool) {
	if o == nil || o.MigProfiles == nil {
		return nil, false
	}
	return o.MigProfiles, true
}

// HasMigProfiles returns a boolean if a field has been set.
func (o *KubernetesNvidiaGpuProduct) HasMigProfiles() bool {
	if o != nil && o.MigProfiles != nil {
		return true
	}

	return false
}

// SetMigProfiles gets a reference to the given []string and assigns it to the MigProfiles field.
func (o *KubernetesNvidiaGpuProduct) SetMigProfiles(v []string) {
	o.MigProfiles = v
}

func (o KubernetesNvidiaGpuProduct) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKubernetesBaseGpuProduct, errKubernetesBaseGpuProduct := json.Marshal(o.KubernetesBaseGpuProduct)
	if errKubernetesBaseGpuProduct != nil {
		return []byte{}, errKubernetesBaseGpuProduct
	}
	errKubernetesBaseGpuProduct = json.Unmarshal([]byte(serializedKubernetesBaseGpuProduct), &toSerialize)
	if errKubernetesBaseGpuProduct != nil {
		return []byte{}, errKubernetesBaseGpuProduct
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.MigCapable != nil {
		toSerialize["MigCapable"] = o.MigCapable
	}
	if o.MigProfiles != nil {
		toSerialize["MigProfiles"] = o.MigProfiles
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesNvidiaGpuProduct) UnmarshalJSON(bytes []byte) (err error) {
	type KubernetesNvidiaGpuProductWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// True if this Nvidia GPU supports MIG.
		MigCapable  *bool    `json:"MigCapable,omitempty"`
		MigProfiles []string `json:"MigProfiles,omitempty"`
	}

	varKubernetesNvidiaGpuProductWithoutEmbeddedStruct := KubernetesNvidiaGpuProductWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varKubernetesNvidiaGpuProductWithoutEmbeddedStruct)
	if err == nil {
		varKubernetesNvidiaGpuProduct := _KubernetesNvidiaGpuProduct{}
		varKubernetesNvidiaGpuProduct.ClassId = varKubernetesNvidiaGpuProductWithoutEmbeddedStruct.ClassId
		varKubernetesNvidiaGpuProduct.ObjectType = varKubernetesNvidiaGpuProductWithoutEmbeddedStruct.ObjectType
		varKubernetesNvidiaGpuProduct.MigCapable = varKubernetesNvidiaGpuProductWithoutEmbeddedStruct.MigCapable
		varKubernetesNvidiaGpuProduct.MigProfiles = varKubernetesNvidiaGpuProductWithoutEmbeddedStruct.MigProfiles
		*o = KubernetesNvidiaGpuProduct(varKubernetesNvidiaGpuProduct)
	} else {
		return err
	}

	varKubernetesNvidiaGpuProduct := _KubernetesNvidiaGpuProduct{}

	err = json.Unmarshal(bytes, &varKubernetesNvidiaGpuProduct)
	if err == nil {
		o.KubernetesBaseGpuProduct = varKubernetesNvidiaGpuProduct.KubernetesBaseGpuProduct
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "MigCapable")
		delete(additionalProperties, "MigProfiles")

		// remove fields from embedded structs
		reflectKubernetesBaseGpuProduct := reflect.ValueOf(o.KubernetesBaseGpuProduct)
		for i := 0; i < reflectKubernetesBaseGpuProduct.Type().NumField(); i++ {
			t := reflectKubernetesBaseGpuProduct.Type().Field(i)

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

type NullableKubernetesNvidiaGpuProduct struct {
	value *KubernetesNvidiaGpuProduct
	isSet bool
}

func (v NullableKubernetesNvidiaGpuProduct) Get() *KubernetesNvidiaGpuProduct {
	return v.value
}

func (v *NullableKubernetesNvidiaGpuProduct) Set(val *KubernetesNvidiaGpuProduct) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesNvidiaGpuProduct) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesNvidiaGpuProduct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesNvidiaGpuProduct(val *KubernetesNvidiaGpuProduct) *NullableKubernetesNvidiaGpuProduct {
	return &NullableKubernetesNvidiaGpuProduct{value: val, isSet: true}
}

func (v NullableKubernetesNvidiaGpuProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesNvidiaGpuProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
