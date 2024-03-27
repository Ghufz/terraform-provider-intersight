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

// VnicFcVethInventory A Fibre Channel Quality of Service (QoS) policy assigns a system class to the outgoing traffic for a vHBA. This system class determines the quality of service for the outgoing traffic. For certain adapters additional controls can also be specified like burst and rate on the outgoing traffic.
type VnicFcVethInventory struct {
	PolicyAbstractInventory
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The burst traffic, in bytes, allowed on the vNIC.
	Burst *int64 `json:"Burst,omitempty"`
	// Description of the virtual FC interface.
	Description *string `json:"Description,omitempty"`
	// Name of the virtual FC interface.
	Name *string `json:"Name,omitempty"`
	// The value in Mbps to use for limiting the data rate on the virtual interface. Setting this to zero will turn rate limiting off.
	RateLimit            *int64                `json:"RateLimit,omitempty"`
	TargetMo             *MoBaseMoRelationship `json:"TargetMo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VnicFcVethInventory VnicFcVethInventory

// NewVnicFcVethInventory instantiates a new VnicFcVethInventory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicFcVethInventory(classId string, objectType string) *VnicFcVethInventory {
	this := VnicFcVethInventory{}
	this.ClassId = classId
	this.ObjectType = objectType
	var burst int64 = 1024
	this.Burst = &burst
	var rateLimit int64 = 0
	this.RateLimit = &rateLimit
	return &this
}

// NewVnicFcVethInventoryWithDefaults instantiates a new VnicFcVethInventory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicFcVethInventoryWithDefaults() *VnicFcVethInventory {
	this := VnicFcVethInventory{}
	var classId string = "vnic.FcVethInventory"
	this.ClassId = classId
	var objectType string = "vnic.FcVethInventory"
	this.ObjectType = objectType
	var burst int64 = 1024
	this.Burst = &burst
	var rateLimit int64 = 0
	this.RateLimit = &rateLimit
	return &this
}

// GetClassId returns the ClassId field value
func (o *VnicFcVethInventory) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VnicFcVethInventory) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VnicFcVethInventory) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VnicFcVethInventory) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VnicFcVethInventory) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VnicFcVethInventory) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBurst returns the Burst field value if set, zero value otherwise.
func (o *VnicFcVethInventory) GetBurst() int64 {
	if o == nil || o.Burst == nil {
		var ret int64
		return ret
	}
	return *o.Burst
}

// GetBurstOk returns a tuple with the Burst field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicFcVethInventory) GetBurstOk() (*int64, bool) {
	if o == nil || o.Burst == nil {
		return nil, false
	}
	return o.Burst, true
}

// HasBurst returns a boolean if a field has been set.
func (o *VnicFcVethInventory) HasBurst() bool {
	if o != nil && o.Burst != nil {
		return true
	}

	return false
}

// SetBurst gets a reference to the given int64 and assigns it to the Burst field.
func (o *VnicFcVethInventory) SetBurst(v int64) {
	o.Burst = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *VnicFcVethInventory) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicFcVethInventory) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *VnicFcVethInventory) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *VnicFcVethInventory) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VnicFcVethInventory) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicFcVethInventory) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VnicFcVethInventory) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VnicFcVethInventory) SetName(v string) {
	o.Name = &v
}

// GetRateLimit returns the RateLimit field value if set, zero value otherwise.
func (o *VnicFcVethInventory) GetRateLimit() int64 {
	if o == nil || o.RateLimit == nil {
		var ret int64
		return ret
	}
	return *o.RateLimit
}

// GetRateLimitOk returns a tuple with the RateLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicFcVethInventory) GetRateLimitOk() (*int64, bool) {
	if o == nil || o.RateLimit == nil {
		return nil, false
	}
	return o.RateLimit, true
}

// HasRateLimit returns a boolean if a field has been set.
func (o *VnicFcVethInventory) HasRateLimit() bool {
	if o != nil && o.RateLimit != nil {
		return true
	}

	return false
}

// SetRateLimit gets a reference to the given int64 and assigns it to the RateLimit field.
func (o *VnicFcVethInventory) SetRateLimit(v int64) {
	o.RateLimit = &v
}

// GetTargetMo returns the TargetMo field value if set, zero value otherwise.
func (o *VnicFcVethInventory) GetTargetMo() MoBaseMoRelationship {
	if o == nil || o.TargetMo == nil {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.TargetMo
}

// GetTargetMoOk returns a tuple with the TargetMo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicFcVethInventory) GetTargetMoOk() (*MoBaseMoRelationship, bool) {
	if o == nil || o.TargetMo == nil {
		return nil, false
	}
	return o.TargetMo, true
}

// HasTargetMo returns a boolean if a field has been set.
func (o *VnicFcVethInventory) HasTargetMo() bool {
	if o != nil && o.TargetMo != nil {
		return true
	}

	return false
}

// SetTargetMo gets a reference to the given MoBaseMoRelationship and assigns it to the TargetMo field.
func (o *VnicFcVethInventory) SetTargetMo(v MoBaseMoRelationship) {
	o.TargetMo = &v
}

func (o VnicFcVethInventory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractInventory, errPolicyAbstractInventory := json.Marshal(o.PolicyAbstractInventory)
	if errPolicyAbstractInventory != nil {
		return []byte{}, errPolicyAbstractInventory
	}
	errPolicyAbstractInventory = json.Unmarshal([]byte(serializedPolicyAbstractInventory), &toSerialize)
	if errPolicyAbstractInventory != nil {
		return []byte{}, errPolicyAbstractInventory
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Burst != nil {
		toSerialize["Burst"] = o.Burst
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.RateLimit != nil {
		toSerialize["RateLimit"] = o.RateLimit
	}
	if o.TargetMo != nil {
		toSerialize["TargetMo"] = o.TargetMo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VnicFcVethInventory) UnmarshalJSON(bytes []byte) (err error) {
	type VnicFcVethInventoryWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The burst traffic, in bytes, allowed on the vNIC.
		Burst *int64 `json:"Burst,omitempty"`
		// Description of the virtual FC interface.
		Description *string `json:"Description,omitempty"`
		// Name of the virtual FC interface.
		Name *string `json:"Name,omitempty"`
		// The value in Mbps to use for limiting the data rate on the virtual interface. Setting this to zero will turn rate limiting off.
		RateLimit *int64                `json:"RateLimit,omitempty"`
		TargetMo  *MoBaseMoRelationship `json:"TargetMo,omitempty"`
	}

	varVnicFcVethInventoryWithoutEmbeddedStruct := VnicFcVethInventoryWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVnicFcVethInventoryWithoutEmbeddedStruct)
	if err == nil {
		varVnicFcVethInventory := _VnicFcVethInventory{}
		varVnicFcVethInventory.ClassId = varVnicFcVethInventoryWithoutEmbeddedStruct.ClassId
		varVnicFcVethInventory.ObjectType = varVnicFcVethInventoryWithoutEmbeddedStruct.ObjectType
		varVnicFcVethInventory.Burst = varVnicFcVethInventoryWithoutEmbeddedStruct.Burst
		varVnicFcVethInventory.Description = varVnicFcVethInventoryWithoutEmbeddedStruct.Description
		varVnicFcVethInventory.Name = varVnicFcVethInventoryWithoutEmbeddedStruct.Name
		varVnicFcVethInventory.RateLimit = varVnicFcVethInventoryWithoutEmbeddedStruct.RateLimit
		varVnicFcVethInventory.TargetMo = varVnicFcVethInventoryWithoutEmbeddedStruct.TargetMo
		*o = VnicFcVethInventory(varVnicFcVethInventory)
	} else {
		return err
	}

	varVnicFcVethInventory := _VnicFcVethInventory{}

	err = json.Unmarshal(bytes, &varVnicFcVethInventory)
	if err == nil {
		o.PolicyAbstractInventory = varVnicFcVethInventory.PolicyAbstractInventory
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Burst")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "RateLimit")
		delete(additionalProperties, "TargetMo")

		// remove fields from embedded structs
		reflectPolicyAbstractInventory := reflect.ValueOf(o.PolicyAbstractInventory)
		for i := 0; i < reflectPolicyAbstractInventory.Type().NumField(); i++ {
			t := reflectPolicyAbstractInventory.Type().Field(i)

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

type NullableVnicFcVethInventory struct {
	value *VnicFcVethInventory
	isSet bool
}

func (v NullableVnicFcVethInventory) Get() *VnicFcVethInventory {
	return v.value
}

func (v *NullableVnicFcVethInventory) Set(val *VnicFcVethInventory) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicFcVethInventory) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicFcVethInventory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicFcVethInventory(val *VnicFcVethInventory) *NullableVnicFcVethInventory {
	return &NullableVnicFcVethInventory{value: val, isSet: true}
}

func (v NullableVnicFcVethInventory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicFcVethInventory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
