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

// ServerServerAssignTypeSlotAllOf Definition of the list of properties defined in 'server.ServerAssignTypeSlot', excluding properties defined in parent classes.
type ServerServerAssignTypeSlotAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Chassis-id of the slot that would be assigned to this pre-assigned server profile.
	ChassisId *int64 `json:"ChassisId,omitempty"`
	// Domain name of the Fabric Interconnect to which the chassis is or to be connected. It can be any string that adheres to the following constraints: It should start and end with an alphanumeric character. It can have underscores and hyphens. It cannot be more than 30 characters.
	DomainName *string `json:"DomainName,omitempty"`
	// Slot-id of the server that would be assigned to this pre-assigned server profile.
	SlotId               *int64 `json:"SlotId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServerServerAssignTypeSlotAllOf ServerServerAssignTypeSlotAllOf

// NewServerServerAssignTypeSlotAllOf instantiates a new ServerServerAssignTypeSlotAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerServerAssignTypeSlotAllOf(classId string, objectType string) *ServerServerAssignTypeSlotAllOf {
	this := ServerServerAssignTypeSlotAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var chassisId int64 = 0
	this.ChassisId = &chassisId
	var slotId int64 = 0
	this.SlotId = &slotId
	return &this
}

// NewServerServerAssignTypeSlotAllOfWithDefaults instantiates a new ServerServerAssignTypeSlotAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerServerAssignTypeSlotAllOfWithDefaults() *ServerServerAssignTypeSlotAllOf {
	this := ServerServerAssignTypeSlotAllOf{}
	var classId string = "server.ServerAssignTypeSlot"
	this.ClassId = classId
	var objectType string = "server.ServerAssignTypeSlot"
	this.ObjectType = objectType
	var chassisId int64 = 0
	this.ChassisId = &chassisId
	var slotId int64 = 0
	this.SlotId = &slotId
	return &this
}

// GetClassId returns the ClassId field value
func (o *ServerServerAssignTypeSlotAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ServerServerAssignTypeSlotAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ServerServerAssignTypeSlotAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ServerServerAssignTypeSlotAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ServerServerAssignTypeSlotAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ServerServerAssignTypeSlotAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetChassisId returns the ChassisId field value if set, zero value otherwise.
func (o *ServerServerAssignTypeSlotAllOf) GetChassisId() int64 {
	if o == nil || o.ChassisId == nil {
		var ret int64
		return ret
	}
	return *o.ChassisId
}

// GetChassisIdOk returns a tuple with the ChassisId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerServerAssignTypeSlotAllOf) GetChassisIdOk() (*int64, bool) {
	if o == nil || o.ChassisId == nil {
		return nil, false
	}
	return o.ChassisId, true
}

// HasChassisId returns a boolean if a field has been set.
func (o *ServerServerAssignTypeSlotAllOf) HasChassisId() bool {
	if o != nil && o.ChassisId != nil {
		return true
	}

	return false
}

// SetChassisId gets a reference to the given int64 and assigns it to the ChassisId field.
func (o *ServerServerAssignTypeSlotAllOf) SetChassisId(v int64) {
	o.ChassisId = &v
}

// GetDomainName returns the DomainName field value if set, zero value otherwise.
func (o *ServerServerAssignTypeSlotAllOf) GetDomainName() string {
	if o == nil || o.DomainName == nil {
		var ret string
		return ret
	}
	return *o.DomainName
}

// GetDomainNameOk returns a tuple with the DomainName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerServerAssignTypeSlotAllOf) GetDomainNameOk() (*string, bool) {
	if o == nil || o.DomainName == nil {
		return nil, false
	}
	return o.DomainName, true
}

// HasDomainName returns a boolean if a field has been set.
func (o *ServerServerAssignTypeSlotAllOf) HasDomainName() bool {
	if o != nil && o.DomainName != nil {
		return true
	}

	return false
}

// SetDomainName gets a reference to the given string and assigns it to the DomainName field.
func (o *ServerServerAssignTypeSlotAllOf) SetDomainName(v string) {
	o.DomainName = &v
}

// GetSlotId returns the SlotId field value if set, zero value otherwise.
func (o *ServerServerAssignTypeSlotAllOf) GetSlotId() int64 {
	if o == nil || o.SlotId == nil {
		var ret int64
		return ret
	}
	return *o.SlotId
}

// GetSlotIdOk returns a tuple with the SlotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerServerAssignTypeSlotAllOf) GetSlotIdOk() (*int64, bool) {
	if o == nil || o.SlotId == nil {
		return nil, false
	}
	return o.SlotId, true
}

// HasSlotId returns a boolean if a field has been set.
func (o *ServerServerAssignTypeSlotAllOf) HasSlotId() bool {
	if o != nil && o.SlotId != nil {
		return true
	}

	return false
}

// SetSlotId gets a reference to the given int64 and assigns it to the SlotId field.
func (o *ServerServerAssignTypeSlotAllOf) SetSlotId(v int64) {
	o.SlotId = &v
}

func (o ServerServerAssignTypeSlotAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ChassisId != nil {
		toSerialize["ChassisId"] = o.ChassisId
	}
	if o.DomainName != nil {
		toSerialize["DomainName"] = o.DomainName
	}
	if o.SlotId != nil {
		toSerialize["SlotId"] = o.SlotId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ServerServerAssignTypeSlotAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varServerServerAssignTypeSlotAllOf := _ServerServerAssignTypeSlotAllOf{}

	if err = json.Unmarshal(bytes, &varServerServerAssignTypeSlotAllOf); err == nil {
		*o = ServerServerAssignTypeSlotAllOf(varServerServerAssignTypeSlotAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ChassisId")
		delete(additionalProperties, "DomainName")
		delete(additionalProperties, "SlotId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServerServerAssignTypeSlotAllOf struct {
	value *ServerServerAssignTypeSlotAllOf
	isSet bool
}

func (v NullableServerServerAssignTypeSlotAllOf) Get() *ServerServerAssignTypeSlotAllOf {
	return v.value
}

func (v *NullableServerServerAssignTypeSlotAllOf) Set(val *ServerServerAssignTypeSlotAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableServerServerAssignTypeSlotAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableServerServerAssignTypeSlotAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerServerAssignTypeSlotAllOf(val *ServerServerAssignTypeSlotAllOf) *NullableServerServerAssignTypeSlotAllOf {
	return &NullableServerServerAssignTypeSlotAllOf{value: val, isSet: true}
}

func (v NullableServerServerAssignTypeSlotAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerServerAssignTypeSlotAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
