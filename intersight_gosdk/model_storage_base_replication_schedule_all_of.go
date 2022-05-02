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

// StorageBaseReplicationScheduleAllOf Definition of the list of properties defined in 'storage.BaseReplicationSchedule', excluding properties defined in parent classes.
type StorageBaseReplicationScheduleAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Replication frequency. It is an interval at which replication is set to trigger. Examples:     PT2H, Snapshot is generated every 2 hours.     P30D, Snapshot is scheduled for every 30 days.     PT2H34M56.123S is 2 hours, 34 minutes, 56 seconds and 123 milliseconds.
	Frequency *string `json:"Frequency,omitempty"`
	// Replication schedule name.
	Name *string `json:"Name,omitempty"`
	// Preferred time of day at which to replicate the snapshots on target array. It is applicable only if the replication frequency is set for a day or more. Format: hh:mm:ss Example: 15:00:00, Replication is set for 3:00 PM.
	ReplicationTime *string `json:"ReplicationTime,omitempty"`
	// Duration to keep the replicated snapshots on the targets. Replicated snapshots are deleted from target array once the retention period has elapsed. Examples: P30D, Snapshots are available for 30 days. PT2H34M56.123S, 2 hours, 34 minutes, 56 seconds and 123 milliseconds.
	RetentionTime        *string `json:"RetentionTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageBaseReplicationScheduleAllOf StorageBaseReplicationScheduleAllOf

// NewStorageBaseReplicationScheduleAllOf instantiates a new StorageBaseReplicationScheduleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageBaseReplicationScheduleAllOf(classId string, objectType string) *StorageBaseReplicationScheduleAllOf {
	this := StorageBaseReplicationScheduleAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageBaseReplicationScheduleAllOfWithDefaults instantiates a new StorageBaseReplicationScheduleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageBaseReplicationScheduleAllOfWithDefaults() *StorageBaseReplicationScheduleAllOf {
	this := StorageBaseReplicationScheduleAllOf{}
	var classId string = "storage.PureReplicationSchedule"
	this.ClassId = classId
	var objectType string = "storage.PureReplicationSchedule"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageBaseReplicationScheduleAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageBaseReplicationScheduleAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageBaseReplicationScheduleAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageBaseReplicationScheduleAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageBaseReplicationScheduleAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageBaseReplicationScheduleAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetFrequency returns the Frequency field value if set, zero value otherwise.
func (o *StorageBaseReplicationScheduleAllOf) GetFrequency() string {
	if o == nil || o.Frequency == nil {
		var ret string
		return ret
	}
	return *o.Frequency
}

// GetFrequencyOk returns a tuple with the Frequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseReplicationScheduleAllOf) GetFrequencyOk() (*string, bool) {
	if o == nil || o.Frequency == nil {
		return nil, false
	}
	return o.Frequency, true
}

// HasFrequency returns a boolean if a field has been set.
func (o *StorageBaseReplicationScheduleAllOf) HasFrequency() bool {
	if o != nil && o.Frequency != nil {
		return true
	}

	return false
}

// SetFrequency gets a reference to the given string and assigns it to the Frequency field.
func (o *StorageBaseReplicationScheduleAllOf) SetFrequency(v string) {
	o.Frequency = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StorageBaseReplicationScheduleAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseReplicationScheduleAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StorageBaseReplicationScheduleAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StorageBaseReplicationScheduleAllOf) SetName(v string) {
	o.Name = &v
}

// GetReplicationTime returns the ReplicationTime field value if set, zero value otherwise.
func (o *StorageBaseReplicationScheduleAllOf) GetReplicationTime() string {
	if o == nil || o.ReplicationTime == nil {
		var ret string
		return ret
	}
	return *o.ReplicationTime
}

// GetReplicationTimeOk returns a tuple with the ReplicationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseReplicationScheduleAllOf) GetReplicationTimeOk() (*string, bool) {
	if o == nil || o.ReplicationTime == nil {
		return nil, false
	}
	return o.ReplicationTime, true
}

// HasReplicationTime returns a boolean if a field has been set.
func (o *StorageBaseReplicationScheduleAllOf) HasReplicationTime() bool {
	if o != nil && o.ReplicationTime != nil {
		return true
	}

	return false
}

// SetReplicationTime gets a reference to the given string and assigns it to the ReplicationTime field.
func (o *StorageBaseReplicationScheduleAllOf) SetReplicationTime(v string) {
	o.ReplicationTime = &v
}

// GetRetentionTime returns the RetentionTime field value if set, zero value otherwise.
func (o *StorageBaseReplicationScheduleAllOf) GetRetentionTime() string {
	if o == nil || o.RetentionTime == nil {
		var ret string
		return ret
	}
	return *o.RetentionTime
}

// GetRetentionTimeOk returns a tuple with the RetentionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseReplicationScheduleAllOf) GetRetentionTimeOk() (*string, bool) {
	if o == nil || o.RetentionTime == nil {
		return nil, false
	}
	return o.RetentionTime, true
}

// HasRetentionTime returns a boolean if a field has been set.
func (o *StorageBaseReplicationScheduleAllOf) HasRetentionTime() bool {
	if o != nil && o.RetentionTime != nil {
		return true
	}

	return false
}

// SetRetentionTime gets a reference to the given string and assigns it to the RetentionTime field.
func (o *StorageBaseReplicationScheduleAllOf) SetRetentionTime(v string) {
	o.RetentionTime = &v
}

func (o StorageBaseReplicationScheduleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Frequency != nil {
		toSerialize["Frequency"] = o.Frequency
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.ReplicationTime != nil {
		toSerialize["ReplicationTime"] = o.ReplicationTime
	}
	if o.RetentionTime != nil {
		toSerialize["RetentionTime"] = o.RetentionTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageBaseReplicationScheduleAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageBaseReplicationScheduleAllOf := _StorageBaseReplicationScheduleAllOf{}

	if err = json.Unmarshal(bytes, &varStorageBaseReplicationScheduleAllOf); err == nil {
		*o = StorageBaseReplicationScheduleAllOf(varStorageBaseReplicationScheduleAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Frequency")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "ReplicationTime")
		delete(additionalProperties, "RetentionTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageBaseReplicationScheduleAllOf struct {
	value *StorageBaseReplicationScheduleAllOf
	isSet bool
}

func (v NullableStorageBaseReplicationScheduleAllOf) Get() *StorageBaseReplicationScheduleAllOf {
	return v.value
}

func (v *NullableStorageBaseReplicationScheduleAllOf) Set(val *StorageBaseReplicationScheduleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageBaseReplicationScheduleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageBaseReplicationScheduleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageBaseReplicationScheduleAllOf(val *StorageBaseReplicationScheduleAllOf) *NullableStorageBaseReplicationScheduleAllOf {
	return &NullableStorageBaseReplicationScheduleAllOf{value: val, isSet: true}
}

func (v NullableStorageBaseReplicationScheduleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageBaseReplicationScheduleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
