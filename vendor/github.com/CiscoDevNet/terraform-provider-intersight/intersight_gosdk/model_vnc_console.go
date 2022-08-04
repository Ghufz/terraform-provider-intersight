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

// VncConsole API to launch the Intersight Workload Engine virtual machine console.
type VncConsole struct {
	TunnelingTunnel
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                                       `json:"ObjectType"`
	VirtualMachine       *VirtualizationIweVirtualMachineRelationship `json:"VirtualMachine,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VncConsole VncConsole

// NewVncConsole instantiates a new VncConsole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVncConsole(classId string, objectType string) *VncConsole {
	this := VncConsole{}
	this.ClassId = classId
	this.ObjectType = objectType
	var status string = "Active"
	this.Status = &status
	return &this
}

// NewVncConsoleWithDefaults instantiates a new VncConsole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVncConsoleWithDefaults() *VncConsole {
	this := VncConsole{}
	var classId string = "vnc.Console"
	this.ClassId = classId
	var objectType string = "vnc.Console"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VncConsole) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VncConsole) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VncConsole) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VncConsole) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VncConsole) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VncConsole) SetObjectType(v string) {
	o.ObjectType = v
}

// GetVirtualMachine returns the VirtualMachine field value if set, zero value otherwise.
func (o *VncConsole) GetVirtualMachine() VirtualizationIweVirtualMachineRelationship {
	if o == nil || o.VirtualMachine == nil {
		var ret VirtualizationIweVirtualMachineRelationship
		return ret
	}
	return *o.VirtualMachine
}

// GetVirtualMachineOk returns a tuple with the VirtualMachine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VncConsole) GetVirtualMachineOk() (*VirtualizationIweVirtualMachineRelationship, bool) {
	if o == nil || o.VirtualMachine == nil {
		return nil, false
	}
	return o.VirtualMachine, true
}

// HasVirtualMachine returns a boolean if a field has been set.
func (o *VncConsole) HasVirtualMachine() bool {
	if o != nil && o.VirtualMachine != nil {
		return true
	}

	return false
}

// SetVirtualMachine gets a reference to the given VirtualizationIweVirtualMachineRelationship and assigns it to the VirtualMachine field.
func (o *VncConsole) SetVirtualMachine(v VirtualizationIweVirtualMachineRelationship) {
	o.VirtualMachine = &v
}

func (o VncConsole) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedTunnelingTunnel, errTunnelingTunnel := json.Marshal(o.TunnelingTunnel)
	if errTunnelingTunnel != nil {
		return []byte{}, errTunnelingTunnel
	}
	errTunnelingTunnel = json.Unmarshal([]byte(serializedTunnelingTunnel), &toSerialize)
	if errTunnelingTunnel != nil {
		return []byte{}, errTunnelingTunnel
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.VirtualMachine != nil {
		toSerialize["VirtualMachine"] = o.VirtualMachine
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VncConsole) UnmarshalJSON(bytes []byte) (err error) {
	type VncConsoleWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType     string                                       `json:"ObjectType"`
		VirtualMachine *VirtualizationIweVirtualMachineRelationship `json:"VirtualMachine,omitempty"`
	}

	varVncConsoleWithoutEmbeddedStruct := VncConsoleWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVncConsoleWithoutEmbeddedStruct)
	if err == nil {
		varVncConsole := _VncConsole{}
		varVncConsole.ClassId = varVncConsoleWithoutEmbeddedStruct.ClassId
		varVncConsole.ObjectType = varVncConsoleWithoutEmbeddedStruct.ObjectType
		varVncConsole.VirtualMachine = varVncConsoleWithoutEmbeddedStruct.VirtualMachine
		*o = VncConsole(varVncConsole)
	} else {
		return err
	}

	varVncConsole := _VncConsole{}

	err = json.Unmarshal(bytes, &varVncConsole)
	if err == nil {
		o.TunnelingTunnel = varVncConsole.TunnelingTunnel
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "VirtualMachine")

		// remove fields from embedded structs
		reflectTunnelingTunnel := reflect.ValueOf(o.TunnelingTunnel)
		for i := 0; i < reflectTunnelingTunnel.Type().NumField(); i++ {
			t := reflectTunnelingTunnel.Type().Field(i)

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

type NullableVncConsole struct {
	value *VncConsole
	isSet bool
}

func (v NullableVncConsole) Get() *VncConsole {
	return v.value
}

func (v *NullableVncConsole) Set(val *VncConsole) {
	v.value = val
	v.isSet = true
}

func (v NullableVncConsole) IsSet() bool {
	return v.isSet
}

func (v *NullableVncConsole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVncConsole(val *VncConsole) *NullableVncConsole {
	return &NullableVncConsole{value: val, isSet: true}
}

func (v NullableVncConsole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVncConsole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
