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

// StorageNetAppAggregateAllOf Definition of the list of properties defined in 'storage.NetAppAggregate', excluding properties defined in parent classes.
type StorageNetAppAggregateAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Storage disk type for NetApp aggregate. * `HDD` - Hard Disk Drive disk type. * `Hybrid` - Solid State Hard Disk Drive. * `Hybrid (Flash Pool)` - SSHD in a flash pool disk type. * `SSD` - Solid State Disk disk type. * `SSD (FabricPool)` - SSD in a flash pool disk type. * `VMDisk (SDS)` - Storage disk with Hard disk drive. * `VMDisk (FabricPool)` - Storage disk with Non-volatile random-access memory drives. * `LUN (FlexArray)` - LUN (FlexArray) disk type. * `Not Mapped` - Storage disk is not mapped.
	AggregateType         *string                                 `json:"AggregateType,omitempty"`
	AvgPerformanceMetrics *StorageNetAppPerformanceMetricsAverage `json:"AvgPerformanceMetrics,omitempty"`
	CloudStorage          []string                                `json:"CloudStorage,omitempty"`
	// Data reduction ratio (logical_used / used).
	EfficiencyRatio *float32 `json:"EfficiencyRatio,omitempty"`
	// Unique identifier of NetApp Aggregate across data center.
	Key *string `json:"Key,omitempty"`
	// The node name for the aggregate.
	NodeName *string `json:"NodeName,omitempty"`
	// Size of the RAID group represented by number of disks in the group.
	RaidSize *int64 `json:"RaidSize,omitempty"`
	// The RAID configuration type. * `Unknown` - Default unknown RAID type. * `RAID0` - RAID0 splits (\"stripes\") data evenly across two or more disks, without parity information. * `RAID1` - RAID1 requires a minimum of two disks to work, and provides data redundancy and failover. * `RAID4` - RAID4 stripes block level data and dedicates a disk to parity. * `RAID5` - RAID5  distributes striping and parity at a block level. * `RAID6` - RAID6 level operates like RAID5 with distributed parity and striping. * `RAID10` - RAID10 requires a minimum of four disks in the array. It stripes across disks for higher performance, and mirrors for redundancy. * `RAIDDP` - RAIDDP uses up to two spare disks to replace and reconstruct the data from up to two simultaneously failed disks within the RAID group. * `RAIDTEC` - With RAIDTEC protection, use up to three spare disks to replace and reconstruct the data from up to three simultaneously failed disks within the RAID group.
	RaidType *string `json:"RaidType,omitempty"`
	// Current state of the NetApp aggregate. * `Unknown` - Specifies that the aggregate is discovered, but the aggregate information is not yet retrieved by the Unified Manager server. * `Online` - Aggregate is ready and available. * `Onlining` - The state is transitioning online. * `Offline` - Aggregate is unavailable. * `Offlining` - The state is transitioning offline. * `Relocating` - The aggregate is being relocated. * `Restricted` - Limited operations (e.g., parity reconstruction) are allowed, but data access is not allowed. * `Failed` - The aggregate cannot be brought online. * `Inconsistent` - The aggregate has been marked corrupted; contact technical support. * `Unmounted` - The aggregate is not mounted.
	State *string `json:"State,omitempty"`
	// Uuid of NetApp Aggregate.
	Uuid            *string                        `json:"Uuid,omitempty"`
	ArrayController *StorageNetAppNodeRelationship `json:"ArrayController,omitempty"`
	// An array of relationships to storageNetAppAggregateEvent resources.
	Events               []StorageNetAppAggregateEventRelationship `json:"Events,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppAggregateAllOf StorageNetAppAggregateAllOf

// NewStorageNetAppAggregateAllOf instantiates a new StorageNetAppAggregateAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppAggregateAllOf(classId string, objectType string) *StorageNetAppAggregateAllOf {
	this := StorageNetAppAggregateAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppAggregateAllOfWithDefaults instantiates a new StorageNetAppAggregateAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppAggregateAllOfWithDefaults() *StorageNetAppAggregateAllOf {
	this := StorageNetAppAggregateAllOf{}
	var classId string = "storage.NetAppAggregate"
	this.ClassId = classId
	var objectType string = "storage.NetAppAggregate"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppAggregateAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppAggregateAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppAggregateAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppAggregateAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppAggregateAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppAggregateAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAggregateType returns the AggregateType field value if set, zero value otherwise.
func (o *StorageNetAppAggregateAllOf) GetAggregateType() string {
	if o == nil || o.AggregateType == nil {
		var ret string
		return ret
	}
	return *o.AggregateType
}

// GetAggregateTypeOk returns a tuple with the AggregateType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppAggregateAllOf) GetAggregateTypeOk() (*string, bool) {
	if o == nil || o.AggregateType == nil {
		return nil, false
	}
	return o.AggregateType, true
}

// HasAggregateType returns a boolean if a field has been set.
func (o *StorageNetAppAggregateAllOf) HasAggregateType() bool {
	if o != nil && o.AggregateType != nil {
		return true
	}

	return false
}

// SetAggregateType gets a reference to the given string and assigns it to the AggregateType field.
func (o *StorageNetAppAggregateAllOf) SetAggregateType(v string) {
	o.AggregateType = &v
}

// GetAvgPerformanceMetrics returns the AvgPerformanceMetrics field value if set, zero value otherwise.
func (o *StorageNetAppAggregateAllOf) GetAvgPerformanceMetrics() StorageNetAppPerformanceMetricsAverage {
	if o == nil || o.AvgPerformanceMetrics == nil {
		var ret StorageNetAppPerformanceMetricsAverage
		return ret
	}
	return *o.AvgPerformanceMetrics
}

// GetAvgPerformanceMetricsOk returns a tuple with the AvgPerformanceMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppAggregateAllOf) GetAvgPerformanceMetricsOk() (*StorageNetAppPerformanceMetricsAverage, bool) {
	if o == nil || o.AvgPerformanceMetrics == nil {
		return nil, false
	}
	return o.AvgPerformanceMetrics, true
}

// HasAvgPerformanceMetrics returns a boolean if a field has been set.
func (o *StorageNetAppAggregateAllOf) HasAvgPerformanceMetrics() bool {
	if o != nil && o.AvgPerformanceMetrics != nil {
		return true
	}

	return false
}

// SetAvgPerformanceMetrics gets a reference to the given StorageNetAppPerformanceMetricsAverage and assigns it to the AvgPerformanceMetrics field.
func (o *StorageNetAppAggregateAllOf) SetAvgPerformanceMetrics(v StorageNetAppPerformanceMetricsAverage) {
	o.AvgPerformanceMetrics = &v
}

// GetCloudStorage returns the CloudStorage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppAggregateAllOf) GetCloudStorage() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.CloudStorage
}

// GetCloudStorageOk returns a tuple with the CloudStorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppAggregateAllOf) GetCloudStorageOk() ([]string, bool) {
	if o == nil || o.CloudStorage == nil {
		return nil, false
	}
	return o.CloudStorage, true
}

// HasCloudStorage returns a boolean if a field has been set.
func (o *StorageNetAppAggregateAllOf) HasCloudStorage() bool {
	if o != nil && o.CloudStorage != nil {
		return true
	}

	return false
}

// SetCloudStorage gets a reference to the given []string and assigns it to the CloudStorage field.
func (o *StorageNetAppAggregateAllOf) SetCloudStorage(v []string) {
	o.CloudStorage = v
}

// GetEfficiencyRatio returns the EfficiencyRatio field value if set, zero value otherwise.
func (o *StorageNetAppAggregateAllOf) GetEfficiencyRatio() float32 {
	if o == nil || o.EfficiencyRatio == nil {
		var ret float32
		return ret
	}
	return *o.EfficiencyRatio
}

// GetEfficiencyRatioOk returns a tuple with the EfficiencyRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppAggregateAllOf) GetEfficiencyRatioOk() (*float32, bool) {
	if o == nil || o.EfficiencyRatio == nil {
		return nil, false
	}
	return o.EfficiencyRatio, true
}

// HasEfficiencyRatio returns a boolean if a field has been set.
func (o *StorageNetAppAggregateAllOf) HasEfficiencyRatio() bool {
	if o != nil && o.EfficiencyRatio != nil {
		return true
	}

	return false
}

// SetEfficiencyRatio gets a reference to the given float32 and assigns it to the EfficiencyRatio field.
func (o *StorageNetAppAggregateAllOf) SetEfficiencyRatio(v float32) {
	o.EfficiencyRatio = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *StorageNetAppAggregateAllOf) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppAggregateAllOf) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *StorageNetAppAggregateAllOf) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *StorageNetAppAggregateAllOf) SetKey(v string) {
	o.Key = &v
}

// GetNodeName returns the NodeName field value if set, zero value otherwise.
func (o *StorageNetAppAggregateAllOf) GetNodeName() string {
	if o == nil || o.NodeName == nil {
		var ret string
		return ret
	}
	return *o.NodeName
}

// GetNodeNameOk returns a tuple with the NodeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppAggregateAllOf) GetNodeNameOk() (*string, bool) {
	if o == nil || o.NodeName == nil {
		return nil, false
	}
	return o.NodeName, true
}

// HasNodeName returns a boolean if a field has been set.
func (o *StorageNetAppAggregateAllOf) HasNodeName() bool {
	if o != nil && o.NodeName != nil {
		return true
	}

	return false
}

// SetNodeName gets a reference to the given string and assigns it to the NodeName field.
func (o *StorageNetAppAggregateAllOf) SetNodeName(v string) {
	o.NodeName = &v
}

// GetRaidSize returns the RaidSize field value if set, zero value otherwise.
func (o *StorageNetAppAggregateAllOf) GetRaidSize() int64 {
	if o == nil || o.RaidSize == nil {
		var ret int64
		return ret
	}
	return *o.RaidSize
}

// GetRaidSizeOk returns a tuple with the RaidSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppAggregateAllOf) GetRaidSizeOk() (*int64, bool) {
	if o == nil || o.RaidSize == nil {
		return nil, false
	}
	return o.RaidSize, true
}

// HasRaidSize returns a boolean if a field has been set.
func (o *StorageNetAppAggregateAllOf) HasRaidSize() bool {
	if o != nil && o.RaidSize != nil {
		return true
	}

	return false
}

// SetRaidSize gets a reference to the given int64 and assigns it to the RaidSize field.
func (o *StorageNetAppAggregateAllOf) SetRaidSize(v int64) {
	o.RaidSize = &v
}

// GetRaidType returns the RaidType field value if set, zero value otherwise.
func (o *StorageNetAppAggregateAllOf) GetRaidType() string {
	if o == nil || o.RaidType == nil {
		var ret string
		return ret
	}
	return *o.RaidType
}

// GetRaidTypeOk returns a tuple with the RaidType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppAggregateAllOf) GetRaidTypeOk() (*string, bool) {
	if o == nil || o.RaidType == nil {
		return nil, false
	}
	return o.RaidType, true
}

// HasRaidType returns a boolean if a field has been set.
func (o *StorageNetAppAggregateAllOf) HasRaidType() bool {
	if o != nil && o.RaidType != nil {
		return true
	}

	return false
}

// SetRaidType gets a reference to the given string and assigns it to the RaidType field.
func (o *StorageNetAppAggregateAllOf) SetRaidType(v string) {
	o.RaidType = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *StorageNetAppAggregateAllOf) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppAggregateAllOf) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *StorageNetAppAggregateAllOf) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *StorageNetAppAggregateAllOf) SetState(v string) {
	o.State = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *StorageNetAppAggregateAllOf) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppAggregateAllOf) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *StorageNetAppAggregateAllOf) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *StorageNetAppAggregateAllOf) SetUuid(v string) {
	o.Uuid = &v
}

// GetArrayController returns the ArrayController field value if set, zero value otherwise.
func (o *StorageNetAppAggregateAllOf) GetArrayController() StorageNetAppNodeRelationship {
	if o == nil || o.ArrayController == nil {
		var ret StorageNetAppNodeRelationship
		return ret
	}
	return *o.ArrayController
}

// GetArrayControllerOk returns a tuple with the ArrayController field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppAggregateAllOf) GetArrayControllerOk() (*StorageNetAppNodeRelationship, bool) {
	if o == nil || o.ArrayController == nil {
		return nil, false
	}
	return o.ArrayController, true
}

// HasArrayController returns a boolean if a field has been set.
func (o *StorageNetAppAggregateAllOf) HasArrayController() bool {
	if o != nil && o.ArrayController != nil {
		return true
	}

	return false
}

// SetArrayController gets a reference to the given StorageNetAppNodeRelationship and assigns it to the ArrayController field.
func (o *StorageNetAppAggregateAllOf) SetArrayController(v StorageNetAppNodeRelationship) {
	o.ArrayController = &v
}

// GetEvents returns the Events field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppAggregateAllOf) GetEvents() []StorageNetAppAggregateEventRelationship {
	if o == nil {
		var ret []StorageNetAppAggregateEventRelationship
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppAggregateAllOf) GetEventsOk() ([]StorageNetAppAggregateEventRelationship, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *StorageNetAppAggregateAllOf) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []StorageNetAppAggregateEventRelationship and assigns it to the Events field.
func (o *StorageNetAppAggregateAllOf) SetEvents(v []StorageNetAppAggregateEventRelationship) {
	o.Events = v
}

func (o StorageNetAppAggregateAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AggregateType != nil {
		toSerialize["AggregateType"] = o.AggregateType
	}
	if o.AvgPerformanceMetrics != nil {
		toSerialize["AvgPerformanceMetrics"] = o.AvgPerformanceMetrics
	}
	if o.CloudStorage != nil {
		toSerialize["CloudStorage"] = o.CloudStorage
	}
	if o.EfficiencyRatio != nil {
		toSerialize["EfficiencyRatio"] = o.EfficiencyRatio
	}
	if o.Key != nil {
		toSerialize["Key"] = o.Key
	}
	if o.NodeName != nil {
		toSerialize["NodeName"] = o.NodeName
	}
	if o.RaidSize != nil {
		toSerialize["RaidSize"] = o.RaidSize
	}
	if o.RaidType != nil {
		toSerialize["RaidType"] = o.RaidType
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}
	if o.Uuid != nil {
		toSerialize["Uuid"] = o.Uuid
	}
	if o.ArrayController != nil {
		toSerialize["ArrayController"] = o.ArrayController
	}
	if o.Events != nil {
		toSerialize["Events"] = o.Events
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageNetAppAggregateAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageNetAppAggregateAllOf := _StorageNetAppAggregateAllOf{}

	if err = json.Unmarshal(bytes, &varStorageNetAppAggregateAllOf); err == nil {
		*o = StorageNetAppAggregateAllOf(varStorageNetAppAggregateAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AggregateType")
		delete(additionalProperties, "AvgPerformanceMetrics")
		delete(additionalProperties, "CloudStorage")
		delete(additionalProperties, "EfficiencyRatio")
		delete(additionalProperties, "Key")
		delete(additionalProperties, "NodeName")
		delete(additionalProperties, "RaidSize")
		delete(additionalProperties, "RaidType")
		delete(additionalProperties, "State")
		delete(additionalProperties, "Uuid")
		delete(additionalProperties, "ArrayController")
		delete(additionalProperties, "Events")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageNetAppAggregateAllOf struct {
	value *StorageNetAppAggregateAllOf
	isSet bool
}

func (v NullableStorageNetAppAggregateAllOf) Get() *StorageNetAppAggregateAllOf {
	return v.value
}

func (v *NullableStorageNetAppAggregateAllOf) Set(val *StorageNetAppAggregateAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppAggregateAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppAggregateAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppAggregateAllOf(val *StorageNetAppAggregateAllOf) *NullableStorageNetAppAggregateAllOf {
	return &NullableStorageNetAppAggregateAllOf{value: val, isSet: true}
}

func (v NullableStorageNetAppAggregateAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppAggregateAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
