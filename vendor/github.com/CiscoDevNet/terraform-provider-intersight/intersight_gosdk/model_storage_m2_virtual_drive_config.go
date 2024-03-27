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

// StorageM2VirtualDriveConfig This models the options for creating a virtual drive on M.2 RAID controller.
type StorageM2VirtualDriveConfig struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Select the M.2 RAID controller slot on which the virtual drive is to be created. Select 'MSTOR-RAID-1' to create virtual drive on the M.2 RAID controller in the first slot or in the MSTOR-RAID slot, 'MSTOR-RAID-2' for second slot, 'MSTOR-RAID-1, MSTOR-RAID-2' for both slots or either slot. * `MSTOR-RAID-1` - Virtual drive  will be created on the M.2 RAID controller in the first slot. * `MSTOR-RAID-2` - Virtual drive  will be created on the M.2 RAID controller in the second slot, if available. * `MSTOR-RAID-1,MSTOR-RAID-2` - Virtual drive  will be created on the M.2 RAID controller in both the slots, if available.
	ControllerSlot *string `json:"ControllerSlot,omitempty"`
	// If enabled, this will create a virtual drive on the M.2 RAID controller.
	Enable *bool `json:"Enable,omitempty"`
	// The name of the virtual drive. The name can be between 1 and 15 alphanumeric characters. Spaces or any special characters other than - (hyphen), _ (underscore), : (colon), and . (period) are not allowed. This field will be pre-populated with the default or user configured value which can be edited.
	Name                 *string `json:"Name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageM2VirtualDriveConfig StorageM2VirtualDriveConfig

// NewStorageM2VirtualDriveConfig instantiates a new StorageM2VirtualDriveConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageM2VirtualDriveConfig(classId string, objectType string) *StorageM2VirtualDriveConfig {
	this := StorageM2VirtualDriveConfig{}
	this.ClassId = classId
	this.ObjectType = objectType
	var controllerSlot string = "MSTOR-RAID-1"
	this.ControllerSlot = &controllerSlot
	var enable bool = false
	this.Enable = &enable
	var name string = "MStorBootVd"
	this.Name = &name
	return &this
}

// NewStorageM2VirtualDriveConfigWithDefaults instantiates a new StorageM2VirtualDriveConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageM2VirtualDriveConfigWithDefaults() *StorageM2VirtualDriveConfig {
	this := StorageM2VirtualDriveConfig{}
	var classId string = "storage.M2VirtualDriveConfig"
	this.ClassId = classId
	var objectType string = "storage.M2VirtualDriveConfig"
	this.ObjectType = objectType
	var controllerSlot string = "MSTOR-RAID-1"
	this.ControllerSlot = &controllerSlot
	var enable bool = false
	this.Enable = &enable
	var name string = "MStorBootVd"
	this.Name = &name
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageM2VirtualDriveConfig) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageM2VirtualDriveConfig) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageM2VirtualDriveConfig) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageM2VirtualDriveConfig) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageM2VirtualDriveConfig) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageM2VirtualDriveConfig) SetObjectType(v string) {
	o.ObjectType = v
}

// GetControllerSlot returns the ControllerSlot field value if set, zero value otherwise.
func (o *StorageM2VirtualDriveConfig) GetControllerSlot() string {
	if o == nil || o.ControllerSlot == nil {
		var ret string
		return ret
	}
	return *o.ControllerSlot
}

// GetControllerSlotOk returns a tuple with the ControllerSlot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageM2VirtualDriveConfig) GetControllerSlotOk() (*string, bool) {
	if o == nil || o.ControllerSlot == nil {
		return nil, false
	}
	return o.ControllerSlot, true
}

// HasControllerSlot returns a boolean if a field has been set.
func (o *StorageM2VirtualDriveConfig) HasControllerSlot() bool {
	if o != nil && o.ControllerSlot != nil {
		return true
	}

	return false
}

// SetControllerSlot gets a reference to the given string and assigns it to the ControllerSlot field.
func (o *StorageM2VirtualDriveConfig) SetControllerSlot(v string) {
	o.ControllerSlot = &v
}

// GetEnable returns the Enable field value if set, zero value otherwise.
func (o *StorageM2VirtualDriveConfig) GetEnable() bool {
	if o == nil || o.Enable == nil {
		var ret bool
		return ret
	}
	return *o.Enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageM2VirtualDriveConfig) GetEnableOk() (*bool, bool) {
	if o == nil || o.Enable == nil {
		return nil, false
	}
	return o.Enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *StorageM2VirtualDriveConfig) HasEnable() bool {
	if o != nil && o.Enable != nil {
		return true
	}

	return false
}

// SetEnable gets a reference to the given bool and assigns it to the Enable field.
func (o *StorageM2VirtualDriveConfig) SetEnable(v bool) {
	o.Enable = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StorageM2VirtualDriveConfig) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageM2VirtualDriveConfig) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StorageM2VirtualDriveConfig) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StorageM2VirtualDriveConfig) SetName(v string) {
	o.Name = &v
}

func (o StorageM2VirtualDriveConfig) MarshalJSON() ([]byte, error) {
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
	if o.ControllerSlot != nil {
		toSerialize["ControllerSlot"] = o.ControllerSlot
	}
	if o.Enable != nil {
		toSerialize["Enable"] = o.Enable
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageM2VirtualDriveConfig) UnmarshalJSON(bytes []byte) (err error) {
	type StorageM2VirtualDriveConfigWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Select the M.2 RAID controller slot on which the virtual drive is to be created. Select 'MSTOR-RAID-1' to create virtual drive on the M.2 RAID controller in the first slot or in the MSTOR-RAID slot, 'MSTOR-RAID-2' for second slot, 'MSTOR-RAID-1, MSTOR-RAID-2' for both slots or either slot. * `MSTOR-RAID-1` - Virtual drive  will be created on the M.2 RAID controller in the first slot. * `MSTOR-RAID-2` - Virtual drive  will be created on the M.2 RAID controller in the second slot, if available. * `MSTOR-RAID-1,MSTOR-RAID-2` - Virtual drive  will be created on the M.2 RAID controller in both the slots, if available.
		ControllerSlot *string `json:"ControllerSlot,omitempty"`
		// If enabled, this will create a virtual drive on the M.2 RAID controller.
		Enable *bool `json:"Enable,omitempty"`
		// The name of the virtual drive. The name can be between 1 and 15 alphanumeric characters. Spaces or any special characters other than - (hyphen), _ (underscore), : (colon), and . (period) are not allowed. This field will be pre-populated with the default or user configured value which can be edited.
		Name *string `json:"Name,omitempty"`
	}

	varStorageM2VirtualDriveConfigWithoutEmbeddedStruct := StorageM2VirtualDriveConfigWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageM2VirtualDriveConfigWithoutEmbeddedStruct)
	if err == nil {
		varStorageM2VirtualDriveConfig := _StorageM2VirtualDriveConfig{}
		varStorageM2VirtualDriveConfig.ClassId = varStorageM2VirtualDriveConfigWithoutEmbeddedStruct.ClassId
		varStorageM2VirtualDriveConfig.ObjectType = varStorageM2VirtualDriveConfigWithoutEmbeddedStruct.ObjectType
		varStorageM2VirtualDriveConfig.ControllerSlot = varStorageM2VirtualDriveConfigWithoutEmbeddedStruct.ControllerSlot
		varStorageM2VirtualDriveConfig.Enable = varStorageM2VirtualDriveConfigWithoutEmbeddedStruct.Enable
		varStorageM2VirtualDriveConfig.Name = varStorageM2VirtualDriveConfigWithoutEmbeddedStruct.Name
		*o = StorageM2VirtualDriveConfig(varStorageM2VirtualDriveConfig)
	} else {
		return err
	}

	varStorageM2VirtualDriveConfig := _StorageM2VirtualDriveConfig{}

	err = json.Unmarshal(bytes, &varStorageM2VirtualDriveConfig)
	if err == nil {
		o.MoBaseComplexType = varStorageM2VirtualDriveConfig.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ControllerSlot")
		delete(additionalProperties, "Enable")
		delete(additionalProperties, "Name")

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

type NullableStorageM2VirtualDriveConfig struct {
	value *StorageM2VirtualDriveConfig
	isSet bool
}

func (v NullableStorageM2VirtualDriveConfig) Get() *StorageM2VirtualDriveConfig {
	return v.value
}

func (v *NullableStorageM2VirtualDriveConfig) Set(val *StorageM2VirtualDriveConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageM2VirtualDriveConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageM2VirtualDriveConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageM2VirtualDriveConfig(val *StorageM2VirtualDriveConfig) *NullableStorageM2VirtualDriveConfig {
	return &NullableStorageM2VirtualDriveConfig{value: val, isSet: true}
}

func (v NullableStorageM2VirtualDriveConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageM2VirtualDriveConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
