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

// StorageNetAppBaseSnapshotPolicy NetApp Snapshot policy controls the behavior and schedule of snapshots when applied to a volume.
type StorageNetAppBaseSnapshotPolicy struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string                                `json:"ObjectType"`
	Copies     []StorageNetAppSnapshotPolicySchedule `json:"Copies,omitempty"`
	// Snapshot policy is enabled or not.
	Enabled *string `json:"Enabled,omitempty"`
	// Name of the NetApp Snapshot Policy.
	Name *string `json:"Name,omitempty"`
	// Identifies whether the Snapshot Policy is owned by the Storage Virtual Machine or the cluster.
	Scope *string `json:"Scope,omitempty"`
	// Uuid of the NetApp Snapshot Policy.
	Uuid                 *string `json:"Uuid,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppBaseSnapshotPolicy StorageNetAppBaseSnapshotPolicy

// NewStorageNetAppBaseSnapshotPolicy instantiates a new StorageNetAppBaseSnapshotPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppBaseSnapshotPolicy(classId string, objectType string) *StorageNetAppBaseSnapshotPolicy {
	this := StorageNetAppBaseSnapshotPolicy{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppBaseSnapshotPolicyWithDefaults instantiates a new StorageNetAppBaseSnapshotPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppBaseSnapshotPolicyWithDefaults() *StorageNetAppBaseSnapshotPolicy {
	this := StorageNetAppBaseSnapshotPolicy{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppBaseSnapshotPolicy) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppBaseSnapshotPolicy) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppBaseSnapshotPolicy) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppBaseSnapshotPolicy) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppBaseSnapshotPolicy) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppBaseSnapshotPolicy) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCopies returns the Copies field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppBaseSnapshotPolicy) GetCopies() []StorageNetAppSnapshotPolicySchedule {
	if o == nil {
		var ret []StorageNetAppSnapshotPolicySchedule
		return ret
	}
	return o.Copies
}

// GetCopiesOk returns a tuple with the Copies field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppBaseSnapshotPolicy) GetCopiesOk() ([]StorageNetAppSnapshotPolicySchedule, bool) {
	if o == nil || o.Copies == nil {
		return nil, false
	}
	return o.Copies, true
}

// HasCopies returns a boolean if a field has been set.
func (o *StorageNetAppBaseSnapshotPolicy) HasCopies() bool {
	if o != nil && o.Copies != nil {
		return true
	}

	return false
}

// SetCopies gets a reference to the given []StorageNetAppSnapshotPolicySchedule and assigns it to the Copies field.
func (o *StorageNetAppBaseSnapshotPolicy) SetCopies(v []StorageNetAppSnapshotPolicySchedule) {
	o.Copies = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *StorageNetAppBaseSnapshotPolicy) GetEnabled() string {
	if o == nil || o.Enabled == nil {
		var ret string
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppBaseSnapshotPolicy) GetEnabledOk() (*string, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *StorageNetAppBaseSnapshotPolicy) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given string and assigns it to the Enabled field.
func (o *StorageNetAppBaseSnapshotPolicy) SetEnabled(v string) {
	o.Enabled = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StorageNetAppBaseSnapshotPolicy) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppBaseSnapshotPolicy) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StorageNetAppBaseSnapshotPolicy) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StorageNetAppBaseSnapshotPolicy) SetName(v string) {
	o.Name = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *StorageNetAppBaseSnapshotPolicy) GetScope() string {
	if o == nil || o.Scope == nil {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppBaseSnapshotPolicy) GetScopeOk() (*string, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *StorageNetAppBaseSnapshotPolicy) HasScope() bool {
	if o != nil && o.Scope != nil {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *StorageNetAppBaseSnapshotPolicy) SetScope(v string) {
	o.Scope = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *StorageNetAppBaseSnapshotPolicy) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppBaseSnapshotPolicy) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *StorageNetAppBaseSnapshotPolicy) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *StorageNetAppBaseSnapshotPolicy) SetUuid(v string) {
	o.Uuid = &v
}

func (o StorageNetAppBaseSnapshotPolicy) MarshalJSON() ([]byte, error) {
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
	if o.Copies != nil {
		toSerialize["Copies"] = o.Copies
	}
	if o.Enabled != nil {
		toSerialize["Enabled"] = o.Enabled
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Scope != nil {
		toSerialize["Scope"] = o.Scope
	}
	if o.Uuid != nil {
		toSerialize["Uuid"] = o.Uuid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageNetAppBaseSnapshotPolicy) UnmarshalJSON(bytes []byte) (err error) {
	type StorageNetAppBaseSnapshotPolicyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string                                `json:"ObjectType"`
		Copies     []StorageNetAppSnapshotPolicySchedule `json:"Copies,omitempty"`
		// Snapshot policy is enabled or not.
		Enabled *string `json:"Enabled,omitempty"`
		// Name of the NetApp Snapshot Policy.
		Name *string `json:"Name,omitempty"`
		// Identifies whether the Snapshot Policy is owned by the Storage Virtual Machine or the cluster.
		Scope *string `json:"Scope,omitempty"`
		// Uuid of the NetApp Snapshot Policy.
		Uuid *string `json:"Uuid,omitempty"`
	}

	varStorageNetAppBaseSnapshotPolicyWithoutEmbeddedStruct := StorageNetAppBaseSnapshotPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageNetAppBaseSnapshotPolicyWithoutEmbeddedStruct)
	if err == nil {
		varStorageNetAppBaseSnapshotPolicy := _StorageNetAppBaseSnapshotPolicy{}
		varStorageNetAppBaseSnapshotPolicy.ClassId = varStorageNetAppBaseSnapshotPolicyWithoutEmbeddedStruct.ClassId
		varStorageNetAppBaseSnapshotPolicy.ObjectType = varStorageNetAppBaseSnapshotPolicyWithoutEmbeddedStruct.ObjectType
		varStorageNetAppBaseSnapshotPolicy.Copies = varStorageNetAppBaseSnapshotPolicyWithoutEmbeddedStruct.Copies
		varStorageNetAppBaseSnapshotPolicy.Enabled = varStorageNetAppBaseSnapshotPolicyWithoutEmbeddedStruct.Enabled
		varStorageNetAppBaseSnapshotPolicy.Name = varStorageNetAppBaseSnapshotPolicyWithoutEmbeddedStruct.Name
		varStorageNetAppBaseSnapshotPolicy.Scope = varStorageNetAppBaseSnapshotPolicyWithoutEmbeddedStruct.Scope
		varStorageNetAppBaseSnapshotPolicy.Uuid = varStorageNetAppBaseSnapshotPolicyWithoutEmbeddedStruct.Uuid
		*o = StorageNetAppBaseSnapshotPolicy(varStorageNetAppBaseSnapshotPolicy)
	} else {
		return err
	}

	varStorageNetAppBaseSnapshotPolicy := _StorageNetAppBaseSnapshotPolicy{}

	err = json.Unmarshal(bytes, &varStorageNetAppBaseSnapshotPolicy)
	if err == nil {
		o.MoBaseMo = varStorageNetAppBaseSnapshotPolicy.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Copies")
		delete(additionalProperties, "Enabled")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Scope")
		delete(additionalProperties, "Uuid")

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

type NullableStorageNetAppBaseSnapshotPolicy struct {
	value *StorageNetAppBaseSnapshotPolicy
	isSet bool
}

func (v NullableStorageNetAppBaseSnapshotPolicy) Get() *StorageNetAppBaseSnapshotPolicy {
	return v.value
}

func (v *NullableStorageNetAppBaseSnapshotPolicy) Set(val *StorageNetAppBaseSnapshotPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppBaseSnapshotPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppBaseSnapshotPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppBaseSnapshotPolicy(val *StorageNetAppBaseSnapshotPolicy) *NullableStorageNetAppBaseSnapshotPolicy {
	return &NullableStorageNetAppBaseSnapshotPolicy{value: val, isSet: true}
}

func (v NullableStorageNetAppBaseSnapshotPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppBaseSnapshotPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
