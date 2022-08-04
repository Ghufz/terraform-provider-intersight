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

// VirtualizationBaseVirtualMachinePciDevice Common attributes of a PCI device on a VM.
type VirtualizationBaseVirtualMachinePciDevice struct {
	VirtualizationBaseSourceDevice
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// The backing physical host PCI device Id for this device.
	BackingPciId *string `json:"BackingPciId,omitempty"`
	// Name of this virtual machine PCI device.
	Name *string `json:"Name,omitempty"`
	// Indicates if this virtual machine PCI device is enabled via passthrough from the host.
	Passthrough          *bool                                         `json:"Passthrough,omitempty"`
	BackingPciDevice     *VirtualizationBaseHostPciDeviceRelationship  `json:"BackingPciDevice,omitempty"`
	VirtualMachine       *VirtualizationBaseVirtualMachineRelationship `json:"VirtualMachine,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationBaseVirtualMachinePciDevice VirtualizationBaseVirtualMachinePciDevice

// NewVirtualizationBaseVirtualMachinePciDevice instantiates a new VirtualizationBaseVirtualMachinePciDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationBaseVirtualMachinePciDevice(classId string, objectType string) *VirtualizationBaseVirtualMachinePciDevice {
	this := VirtualizationBaseVirtualMachinePciDevice{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVirtualizationBaseVirtualMachinePciDeviceWithDefaults instantiates a new VirtualizationBaseVirtualMachinePciDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationBaseVirtualMachinePciDeviceWithDefaults() *VirtualizationBaseVirtualMachinePciDevice {
	this := VirtualizationBaseVirtualMachinePciDevice{}
	var classId string = "virtualization.VmwareVirtualMachineGpu"
	this.ClassId = classId
	var objectType string = "virtualization.VmwareVirtualMachineGpu"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationBaseVirtualMachinePciDevice) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseVirtualMachinePciDevice) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationBaseVirtualMachinePciDevice) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationBaseVirtualMachinePciDevice) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseVirtualMachinePciDevice) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationBaseVirtualMachinePciDevice) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBackingPciId returns the BackingPciId field value if set, zero value otherwise.
func (o *VirtualizationBaseVirtualMachinePciDevice) GetBackingPciId() string {
	if o == nil || o.BackingPciId == nil {
		var ret string
		return ret
	}
	return *o.BackingPciId
}

// GetBackingPciIdOk returns a tuple with the BackingPciId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseVirtualMachinePciDevice) GetBackingPciIdOk() (*string, bool) {
	if o == nil || o.BackingPciId == nil {
		return nil, false
	}
	return o.BackingPciId, true
}

// HasBackingPciId returns a boolean if a field has been set.
func (o *VirtualizationBaseVirtualMachinePciDevice) HasBackingPciId() bool {
	if o != nil && o.BackingPciId != nil {
		return true
	}

	return false
}

// SetBackingPciId gets a reference to the given string and assigns it to the BackingPciId field.
func (o *VirtualizationBaseVirtualMachinePciDevice) SetBackingPciId(v string) {
	o.BackingPciId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VirtualizationBaseVirtualMachinePciDevice) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseVirtualMachinePciDevice) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VirtualizationBaseVirtualMachinePciDevice) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VirtualizationBaseVirtualMachinePciDevice) SetName(v string) {
	o.Name = &v
}

// GetPassthrough returns the Passthrough field value if set, zero value otherwise.
func (o *VirtualizationBaseVirtualMachinePciDevice) GetPassthrough() bool {
	if o == nil || o.Passthrough == nil {
		var ret bool
		return ret
	}
	return *o.Passthrough
}

// GetPassthroughOk returns a tuple with the Passthrough field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseVirtualMachinePciDevice) GetPassthroughOk() (*bool, bool) {
	if o == nil || o.Passthrough == nil {
		return nil, false
	}
	return o.Passthrough, true
}

// HasPassthrough returns a boolean if a field has been set.
func (o *VirtualizationBaseVirtualMachinePciDevice) HasPassthrough() bool {
	if o != nil && o.Passthrough != nil {
		return true
	}

	return false
}

// SetPassthrough gets a reference to the given bool and assigns it to the Passthrough field.
func (o *VirtualizationBaseVirtualMachinePciDevice) SetPassthrough(v bool) {
	o.Passthrough = &v
}

// GetBackingPciDevice returns the BackingPciDevice field value if set, zero value otherwise.
func (o *VirtualizationBaseVirtualMachinePciDevice) GetBackingPciDevice() VirtualizationBaseHostPciDeviceRelationship {
	if o == nil || o.BackingPciDevice == nil {
		var ret VirtualizationBaseHostPciDeviceRelationship
		return ret
	}
	return *o.BackingPciDevice
}

// GetBackingPciDeviceOk returns a tuple with the BackingPciDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseVirtualMachinePciDevice) GetBackingPciDeviceOk() (*VirtualizationBaseHostPciDeviceRelationship, bool) {
	if o == nil || o.BackingPciDevice == nil {
		return nil, false
	}
	return o.BackingPciDevice, true
}

// HasBackingPciDevice returns a boolean if a field has been set.
func (o *VirtualizationBaseVirtualMachinePciDevice) HasBackingPciDevice() bool {
	if o != nil && o.BackingPciDevice != nil {
		return true
	}

	return false
}

// SetBackingPciDevice gets a reference to the given VirtualizationBaseHostPciDeviceRelationship and assigns it to the BackingPciDevice field.
func (o *VirtualizationBaseVirtualMachinePciDevice) SetBackingPciDevice(v VirtualizationBaseHostPciDeviceRelationship) {
	o.BackingPciDevice = &v
}

// GetVirtualMachine returns the VirtualMachine field value if set, zero value otherwise.
func (o *VirtualizationBaseVirtualMachinePciDevice) GetVirtualMachine() VirtualizationBaseVirtualMachineRelationship {
	if o == nil || o.VirtualMachine == nil {
		var ret VirtualizationBaseVirtualMachineRelationship
		return ret
	}
	return *o.VirtualMachine
}

// GetVirtualMachineOk returns a tuple with the VirtualMachine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseVirtualMachinePciDevice) GetVirtualMachineOk() (*VirtualizationBaseVirtualMachineRelationship, bool) {
	if o == nil || o.VirtualMachine == nil {
		return nil, false
	}
	return o.VirtualMachine, true
}

// HasVirtualMachine returns a boolean if a field has been set.
func (o *VirtualizationBaseVirtualMachinePciDevice) HasVirtualMachine() bool {
	if o != nil && o.VirtualMachine != nil {
		return true
	}

	return false
}

// SetVirtualMachine gets a reference to the given VirtualizationBaseVirtualMachineRelationship and assigns it to the VirtualMachine field.
func (o *VirtualizationBaseVirtualMachinePciDevice) SetVirtualMachine(v VirtualizationBaseVirtualMachineRelationship) {
	o.VirtualMachine = &v
}

func (o VirtualizationBaseVirtualMachinePciDevice) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedVirtualizationBaseSourceDevice, errVirtualizationBaseSourceDevice := json.Marshal(o.VirtualizationBaseSourceDevice)
	if errVirtualizationBaseSourceDevice != nil {
		return []byte{}, errVirtualizationBaseSourceDevice
	}
	errVirtualizationBaseSourceDevice = json.Unmarshal([]byte(serializedVirtualizationBaseSourceDevice), &toSerialize)
	if errVirtualizationBaseSourceDevice != nil {
		return []byte{}, errVirtualizationBaseSourceDevice
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.BackingPciId != nil {
		toSerialize["BackingPciId"] = o.BackingPciId
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Passthrough != nil {
		toSerialize["Passthrough"] = o.Passthrough
	}
	if o.BackingPciDevice != nil {
		toSerialize["BackingPciDevice"] = o.BackingPciDevice
	}
	if o.VirtualMachine != nil {
		toSerialize["VirtualMachine"] = o.VirtualMachine
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationBaseVirtualMachinePciDevice) UnmarshalJSON(bytes []byte) (err error) {
	type VirtualizationBaseVirtualMachinePciDeviceWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// The backing physical host PCI device Id for this device.
		BackingPciId *string `json:"BackingPciId,omitempty"`
		// Name of this virtual machine PCI device.
		Name *string `json:"Name,omitempty"`
		// Indicates if this virtual machine PCI device is enabled via passthrough from the host.
		Passthrough      *bool                                         `json:"Passthrough,omitempty"`
		BackingPciDevice *VirtualizationBaseHostPciDeviceRelationship  `json:"BackingPciDevice,omitempty"`
		VirtualMachine   *VirtualizationBaseVirtualMachineRelationship `json:"VirtualMachine,omitempty"`
	}

	varVirtualizationBaseVirtualMachinePciDeviceWithoutEmbeddedStruct := VirtualizationBaseVirtualMachinePciDeviceWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVirtualizationBaseVirtualMachinePciDeviceWithoutEmbeddedStruct)
	if err == nil {
		varVirtualizationBaseVirtualMachinePciDevice := _VirtualizationBaseVirtualMachinePciDevice{}
		varVirtualizationBaseVirtualMachinePciDevice.ClassId = varVirtualizationBaseVirtualMachinePciDeviceWithoutEmbeddedStruct.ClassId
		varVirtualizationBaseVirtualMachinePciDevice.ObjectType = varVirtualizationBaseVirtualMachinePciDeviceWithoutEmbeddedStruct.ObjectType
		varVirtualizationBaseVirtualMachinePciDevice.BackingPciId = varVirtualizationBaseVirtualMachinePciDeviceWithoutEmbeddedStruct.BackingPciId
		varVirtualizationBaseVirtualMachinePciDevice.Name = varVirtualizationBaseVirtualMachinePciDeviceWithoutEmbeddedStruct.Name
		varVirtualizationBaseVirtualMachinePciDevice.Passthrough = varVirtualizationBaseVirtualMachinePciDeviceWithoutEmbeddedStruct.Passthrough
		varVirtualizationBaseVirtualMachinePciDevice.BackingPciDevice = varVirtualizationBaseVirtualMachinePciDeviceWithoutEmbeddedStruct.BackingPciDevice
		varVirtualizationBaseVirtualMachinePciDevice.VirtualMachine = varVirtualizationBaseVirtualMachinePciDeviceWithoutEmbeddedStruct.VirtualMachine
		*o = VirtualizationBaseVirtualMachinePciDevice(varVirtualizationBaseVirtualMachinePciDevice)
	} else {
		return err
	}

	varVirtualizationBaseVirtualMachinePciDevice := _VirtualizationBaseVirtualMachinePciDevice{}

	err = json.Unmarshal(bytes, &varVirtualizationBaseVirtualMachinePciDevice)
	if err == nil {
		o.VirtualizationBaseSourceDevice = varVirtualizationBaseVirtualMachinePciDevice.VirtualizationBaseSourceDevice
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BackingPciId")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Passthrough")
		delete(additionalProperties, "BackingPciDevice")
		delete(additionalProperties, "VirtualMachine")

		// remove fields from embedded structs
		reflectVirtualizationBaseSourceDevice := reflect.ValueOf(o.VirtualizationBaseSourceDevice)
		for i := 0; i < reflectVirtualizationBaseSourceDevice.Type().NumField(); i++ {
			t := reflectVirtualizationBaseSourceDevice.Type().Field(i)

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

type NullableVirtualizationBaseVirtualMachinePciDevice struct {
	value *VirtualizationBaseVirtualMachinePciDevice
	isSet bool
}

func (v NullableVirtualizationBaseVirtualMachinePciDevice) Get() *VirtualizationBaseVirtualMachinePciDevice {
	return v.value
}

func (v *NullableVirtualizationBaseVirtualMachinePciDevice) Set(val *VirtualizationBaseVirtualMachinePciDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationBaseVirtualMachinePciDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationBaseVirtualMachinePciDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationBaseVirtualMachinePciDevice(val *VirtualizationBaseVirtualMachinePciDevice) *NullableVirtualizationBaseVirtualMachinePciDevice {
	return &NullableVirtualizationBaseVirtualMachinePciDevice{value: val, isSet: true}
}

func (v NullableVirtualizationBaseVirtualMachinePciDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationBaseVirtualMachinePciDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
