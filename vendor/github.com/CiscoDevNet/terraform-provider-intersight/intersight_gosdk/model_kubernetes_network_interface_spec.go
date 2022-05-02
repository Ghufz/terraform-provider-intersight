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
	"reflect"
	"strings"
)

// KubernetesNetworkInterfaceSpec InfraConfigPolicy's Specification for a Node Profile's Network Interface.
type KubernetesNetworkInterfaceSpec struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The MTU for this Network Interface.  If left blank a default value will apply by the Operating System.
	Mtu *int64 `json:"Mtu,omitempty"`
	// NetworkInterfaceSpec is the specification for network interfaces - including configuration of IP Pools and VRF to determine IP configuration, the operating system device settings, and virtual adapter network settings. It can be left empty when used with VirtualMachineInfraConfigPolicy - it will be filled out based on the hypervisor platform type and will match the naming and order of interfaces provided by the hypervisor.
	Name  *string   `json:"Name,omitempty"`
	Pools []MoMoRef `json:"Pools,omitempty"`
	// In other words, to which named network from the Instructure Provider will the port be connected.
	ProviderName         *string  `json:"ProviderName,omitempty"`
	Vrf                  *MoMoRef `json:"Vrf,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesNetworkInterfaceSpec KubernetesNetworkInterfaceSpec

// NewKubernetesNetworkInterfaceSpec instantiates a new KubernetesNetworkInterfaceSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesNetworkInterfaceSpec(classId string, objectType string) *KubernetesNetworkInterfaceSpec {
	this := KubernetesNetworkInterfaceSpec{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKubernetesNetworkInterfaceSpecWithDefaults instantiates a new KubernetesNetworkInterfaceSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesNetworkInterfaceSpecWithDefaults() *KubernetesNetworkInterfaceSpec {
	this := KubernetesNetworkInterfaceSpec{}
	var classId string = "kubernetes.NetworkInterfaceSpec"
	this.ClassId = classId
	var objectType string = "kubernetes.NetworkInterfaceSpec"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesNetworkInterfaceSpec) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesNetworkInterfaceSpec) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesNetworkInterfaceSpec) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesNetworkInterfaceSpec) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesNetworkInterfaceSpec) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesNetworkInterfaceSpec) SetObjectType(v string) {
	o.ObjectType = v
}

// GetMtu returns the Mtu field value if set, zero value otherwise.
func (o *KubernetesNetworkInterfaceSpec) GetMtu() int64 {
	if o == nil || o.Mtu == nil {
		var ret int64
		return ret
	}
	return *o.Mtu
}

// GetMtuOk returns a tuple with the Mtu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNetworkInterfaceSpec) GetMtuOk() (*int64, bool) {
	if o == nil || o.Mtu == nil {
		return nil, false
	}
	return o.Mtu, true
}

// HasMtu returns a boolean if a field has been set.
func (o *KubernetesNetworkInterfaceSpec) HasMtu() bool {
	if o != nil && o.Mtu != nil {
		return true
	}

	return false
}

// SetMtu gets a reference to the given int64 and assigns it to the Mtu field.
func (o *KubernetesNetworkInterfaceSpec) SetMtu(v int64) {
	o.Mtu = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *KubernetesNetworkInterfaceSpec) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNetworkInterfaceSpec) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *KubernetesNetworkInterfaceSpec) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *KubernetesNetworkInterfaceSpec) SetName(v string) {
	o.Name = &v
}

// GetPools returns the Pools field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesNetworkInterfaceSpec) GetPools() []MoMoRef {
	if o == nil {
		var ret []MoMoRef
		return ret
	}
	return o.Pools
}

// GetPoolsOk returns a tuple with the Pools field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNetworkInterfaceSpec) GetPoolsOk() ([]MoMoRef, bool) {
	if o == nil || o.Pools == nil {
		return nil, false
	}
	return o.Pools, true
}

// HasPools returns a boolean if a field has been set.
func (o *KubernetesNetworkInterfaceSpec) HasPools() bool {
	if o != nil && o.Pools != nil {
		return true
	}

	return false
}

// SetPools gets a reference to the given []MoMoRef and assigns it to the Pools field.
func (o *KubernetesNetworkInterfaceSpec) SetPools(v []MoMoRef) {
	o.Pools = v
}

// GetProviderName returns the ProviderName field value if set, zero value otherwise.
func (o *KubernetesNetworkInterfaceSpec) GetProviderName() string {
	if o == nil || o.ProviderName == nil {
		var ret string
		return ret
	}
	return *o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNetworkInterfaceSpec) GetProviderNameOk() (*string, bool) {
	if o == nil || o.ProviderName == nil {
		return nil, false
	}
	return o.ProviderName, true
}

// HasProviderName returns a boolean if a field has been set.
func (o *KubernetesNetworkInterfaceSpec) HasProviderName() bool {
	if o != nil && o.ProviderName != nil {
		return true
	}

	return false
}

// SetProviderName gets a reference to the given string and assigns it to the ProviderName field.
func (o *KubernetesNetworkInterfaceSpec) SetProviderName(v string) {
	o.ProviderName = &v
}

// GetVrf returns the Vrf field value if set, zero value otherwise.
func (o *KubernetesNetworkInterfaceSpec) GetVrf() MoMoRef {
	if o == nil || o.Vrf == nil {
		var ret MoMoRef
		return ret
	}
	return *o.Vrf
}

// GetVrfOk returns a tuple with the Vrf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNetworkInterfaceSpec) GetVrfOk() (*MoMoRef, bool) {
	if o == nil || o.Vrf == nil {
		return nil, false
	}
	return o.Vrf, true
}

// HasVrf returns a boolean if a field has been set.
func (o *KubernetesNetworkInterfaceSpec) HasVrf() bool {
	if o != nil && o.Vrf != nil {
		return true
	}

	return false
}

// SetVrf gets a reference to the given MoMoRef and assigns it to the Vrf field.
func (o *KubernetesNetworkInterfaceSpec) SetVrf(v MoMoRef) {
	o.Vrf = &v
}

func (o KubernetesNetworkInterfaceSpec) MarshalJSON() ([]byte, error) {
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
	if o.Mtu != nil {
		toSerialize["Mtu"] = o.Mtu
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Pools != nil {
		toSerialize["Pools"] = o.Pools
	}
	if o.ProviderName != nil {
		toSerialize["ProviderName"] = o.ProviderName
	}
	if o.Vrf != nil {
		toSerialize["Vrf"] = o.Vrf
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesNetworkInterfaceSpec) UnmarshalJSON(bytes []byte) (err error) {
	type KubernetesNetworkInterfaceSpecWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The MTU for this Network Interface.  If left blank a default value will apply by the Operating System.
		Mtu *int64 `json:"Mtu,omitempty"`
		// NetworkInterfaceSpec is the specification for network interfaces - including configuration of IP Pools and VRF to determine IP configuration, the operating system device settings, and virtual adapter network settings. It can be left empty when used with VirtualMachineInfraConfigPolicy - it will be filled out based on the hypervisor platform type and will match the naming and order of interfaces provided by the hypervisor.
		Name  *string   `json:"Name,omitempty"`
		Pools []MoMoRef `json:"Pools,omitempty"`
		// In other words, to which named network from the Instructure Provider will the port be connected.
		ProviderName *string  `json:"ProviderName,omitempty"`
		Vrf          *MoMoRef `json:"Vrf,omitempty"`
	}

	varKubernetesNetworkInterfaceSpecWithoutEmbeddedStruct := KubernetesNetworkInterfaceSpecWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varKubernetesNetworkInterfaceSpecWithoutEmbeddedStruct)
	if err == nil {
		varKubernetesNetworkInterfaceSpec := _KubernetesNetworkInterfaceSpec{}
		varKubernetesNetworkInterfaceSpec.ClassId = varKubernetesNetworkInterfaceSpecWithoutEmbeddedStruct.ClassId
		varKubernetesNetworkInterfaceSpec.ObjectType = varKubernetesNetworkInterfaceSpecWithoutEmbeddedStruct.ObjectType
		varKubernetesNetworkInterfaceSpec.Mtu = varKubernetesNetworkInterfaceSpecWithoutEmbeddedStruct.Mtu
		varKubernetesNetworkInterfaceSpec.Name = varKubernetesNetworkInterfaceSpecWithoutEmbeddedStruct.Name
		varKubernetesNetworkInterfaceSpec.Pools = varKubernetesNetworkInterfaceSpecWithoutEmbeddedStruct.Pools
		varKubernetesNetworkInterfaceSpec.ProviderName = varKubernetesNetworkInterfaceSpecWithoutEmbeddedStruct.ProviderName
		varKubernetesNetworkInterfaceSpec.Vrf = varKubernetesNetworkInterfaceSpecWithoutEmbeddedStruct.Vrf
		*o = KubernetesNetworkInterfaceSpec(varKubernetesNetworkInterfaceSpec)
	} else {
		return err
	}

	varKubernetesNetworkInterfaceSpec := _KubernetesNetworkInterfaceSpec{}

	err = json.Unmarshal(bytes, &varKubernetesNetworkInterfaceSpec)
	if err == nil {
		o.MoBaseComplexType = varKubernetesNetworkInterfaceSpec.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Mtu")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Pools")
		delete(additionalProperties, "ProviderName")
		delete(additionalProperties, "Vrf")

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

type NullableKubernetesNetworkInterfaceSpec struct {
	value *KubernetesNetworkInterfaceSpec
	isSet bool
}

func (v NullableKubernetesNetworkInterfaceSpec) Get() *KubernetesNetworkInterfaceSpec {
	return v.value
}

func (v *NullableKubernetesNetworkInterfaceSpec) Set(val *KubernetesNetworkInterfaceSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesNetworkInterfaceSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesNetworkInterfaceSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesNetworkInterfaceSpec(val *KubernetesNetworkInterfaceSpec) *NullableKubernetesNetworkInterfaceSpec {
	return &NullableKubernetesNetworkInterfaceSpec{value: val, isSet: true}
}

func (v NullableKubernetesNetworkInterfaceSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesNetworkInterfaceSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
