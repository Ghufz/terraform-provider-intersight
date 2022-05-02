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
)

// StorageBaseCapacityAllOf Definition of the list of properties defined in 'storage.BaseCapacity', excluding properties defined in parent classes.
type StorageBaseCapacityAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Total consumable storage capacity represented in bytes. System may reserve some space for internal purposes which is excluded from total capacity.
	Available *int64 `json:"Available,omitempty"`
	// Percentage of used capacity.
	CapacityUtilization *float32 `json:"CapacityUtilization,omitempty"`
	// Unused space available for applications to consume, represented in bytes.
	Free *int64 `json:"Free,omitempty"`
	// Total storage capacity, represented in bytes. It is set by the component manufacturer.
	Total *int64 `json:"Total,omitempty"`
	// Used or consumed storage capacity, represented in bytes.
	Used                 *int64 `json:"Used,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageBaseCapacityAllOf StorageBaseCapacityAllOf

// NewStorageBaseCapacityAllOf instantiates a new StorageBaseCapacityAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageBaseCapacityAllOf(classId string, objectType string) *StorageBaseCapacityAllOf {
	this := StorageBaseCapacityAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageBaseCapacityAllOfWithDefaults instantiates a new StorageBaseCapacityAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageBaseCapacityAllOfWithDefaults() *StorageBaseCapacityAllOf {
	this := StorageBaseCapacityAllOf{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageBaseCapacityAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageBaseCapacityAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageBaseCapacityAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageBaseCapacityAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageBaseCapacityAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageBaseCapacityAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAvailable returns the Available field value if set, zero value otherwise.
func (o *StorageBaseCapacityAllOf) GetAvailable() int64 {
	if o == nil || o.Available == nil {
		var ret int64
		return ret
	}
	return *o.Available
}

// GetAvailableOk returns a tuple with the Available field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseCapacityAllOf) GetAvailableOk() (*int64, bool) {
	if o == nil || o.Available == nil {
		return nil, false
	}
	return o.Available, true
}

// HasAvailable returns a boolean if a field has been set.
func (o *StorageBaseCapacityAllOf) HasAvailable() bool {
	if o != nil && o.Available != nil {
		return true
	}

	return false
}

// SetAvailable gets a reference to the given int64 and assigns it to the Available field.
func (o *StorageBaseCapacityAllOf) SetAvailable(v int64) {
	o.Available = &v
}

// GetCapacityUtilization returns the CapacityUtilization field value if set, zero value otherwise.
func (o *StorageBaseCapacityAllOf) GetCapacityUtilization() float32 {
	if o == nil || o.CapacityUtilization == nil {
		var ret float32
		return ret
	}
	return *o.CapacityUtilization
}

// GetCapacityUtilizationOk returns a tuple with the CapacityUtilization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseCapacityAllOf) GetCapacityUtilizationOk() (*float32, bool) {
	if o == nil || o.CapacityUtilization == nil {
		return nil, false
	}
	return o.CapacityUtilization, true
}

// HasCapacityUtilization returns a boolean if a field has been set.
func (o *StorageBaseCapacityAllOf) HasCapacityUtilization() bool {
	if o != nil && o.CapacityUtilization != nil {
		return true
	}

	return false
}

// SetCapacityUtilization gets a reference to the given float32 and assigns it to the CapacityUtilization field.
func (o *StorageBaseCapacityAllOf) SetCapacityUtilization(v float32) {
	o.CapacityUtilization = &v
}

// GetFree returns the Free field value if set, zero value otherwise.
func (o *StorageBaseCapacityAllOf) GetFree() int64 {
	if o == nil || o.Free == nil {
		var ret int64
		return ret
	}
	return *o.Free
}

// GetFreeOk returns a tuple with the Free field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseCapacityAllOf) GetFreeOk() (*int64, bool) {
	if o == nil || o.Free == nil {
		return nil, false
	}
	return o.Free, true
}

// HasFree returns a boolean if a field has been set.
func (o *StorageBaseCapacityAllOf) HasFree() bool {
	if o != nil && o.Free != nil {
		return true
	}

	return false
}

// SetFree gets a reference to the given int64 and assigns it to the Free field.
func (o *StorageBaseCapacityAllOf) SetFree(v int64) {
	o.Free = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *StorageBaseCapacityAllOf) GetTotal() int64 {
	if o == nil || o.Total == nil {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseCapacityAllOf) GetTotalOk() (*int64, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *StorageBaseCapacityAllOf) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *StorageBaseCapacityAllOf) SetTotal(v int64) {
	o.Total = &v
}

// GetUsed returns the Used field value if set, zero value otherwise.
func (o *StorageBaseCapacityAllOf) GetUsed() int64 {
	if o == nil || o.Used == nil {
		var ret int64
		return ret
	}
	return *o.Used
}

// GetUsedOk returns a tuple with the Used field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseCapacityAllOf) GetUsedOk() (*int64, bool) {
	if o == nil || o.Used == nil {
		return nil, false
	}
	return o.Used, true
}

// HasUsed returns a boolean if a field has been set.
func (o *StorageBaseCapacityAllOf) HasUsed() bool {
	if o != nil && o.Used != nil {
		return true
	}

	return false
}

// SetUsed gets a reference to the given int64 and assigns it to the Used field.
func (o *StorageBaseCapacityAllOf) SetUsed(v int64) {
	o.Used = &v
}

func (o StorageBaseCapacityAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Available != nil {
		toSerialize["Available"] = o.Available
	}
	if o.CapacityUtilization != nil {
		toSerialize["CapacityUtilization"] = o.CapacityUtilization
	}
	if o.Free != nil {
		toSerialize["Free"] = o.Free
	}
	if o.Total != nil {
		toSerialize["Total"] = o.Total
	}
	if o.Used != nil {
		toSerialize["Used"] = o.Used
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageBaseCapacityAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageBaseCapacityAllOf := _StorageBaseCapacityAllOf{}

	if err = json.Unmarshal(bytes, &varStorageBaseCapacityAllOf); err == nil {
		*o = StorageBaseCapacityAllOf(varStorageBaseCapacityAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Available")
		delete(additionalProperties, "CapacityUtilization")
		delete(additionalProperties, "Free")
		delete(additionalProperties, "Total")
		delete(additionalProperties, "Used")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageBaseCapacityAllOf struct {
	value *StorageBaseCapacityAllOf
	isSet bool
}

func (v NullableStorageBaseCapacityAllOf) Get() *StorageBaseCapacityAllOf {
	return v.value
}

func (v *NullableStorageBaseCapacityAllOf) Set(val *StorageBaseCapacityAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageBaseCapacityAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageBaseCapacityAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageBaseCapacityAllOf(val *StorageBaseCapacityAllOf) *NullableStorageBaseCapacityAllOf {
	return &NullableStorageBaseCapacityAllOf{value: val, isSet: true}
}

func (v NullableStorageBaseCapacityAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageBaseCapacityAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
