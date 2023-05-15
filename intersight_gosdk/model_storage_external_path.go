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

// StorageExternalPath The following properties are restored for each external path.
type StorageExternalPath struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The time (in seconds) until the external parity group is blocked after all paths to the external parity group are disconnected.
	BlockedPathMonitoring *int64 `json:"BlockedPathMonitoring,omitempty"`
	// WWN of the external storage system.
	ExternalWwn *string `json:"ExternalWwn,omitempty"`
	// The value (in seconds) set for the I/O time over for the external parity group.
	IoTimeOut *int64 `json:"IoTimeOut,omitempty"`
	// Port number of external path on the local storage system.
	PortId *string `json:"PortId,omitempty"`
	// Number of Read-Write commands that can be queued to the external parity group.
	QueueDepth           *int64 `json:"QueueDepth,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageExternalPath StorageExternalPath

// NewStorageExternalPath instantiates a new StorageExternalPath object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageExternalPath(classId string, objectType string) *StorageExternalPath {
	this := StorageExternalPath{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageExternalPathWithDefaults instantiates a new StorageExternalPath object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageExternalPathWithDefaults() *StorageExternalPath {
	this := StorageExternalPath{}
	var classId string = "storage.ExternalPath"
	this.ClassId = classId
	var objectType string = "storage.ExternalPath"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageExternalPath) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageExternalPath) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageExternalPath) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageExternalPath) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageExternalPath) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageExternalPath) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBlockedPathMonitoring returns the BlockedPathMonitoring field value if set, zero value otherwise.
func (o *StorageExternalPath) GetBlockedPathMonitoring() int64 {
	if o == nil || o.BlockedPathMonitoring == nil {
		var ret int64
		return ret
	}
	return *o.BlockedPathMonitoring
}

// GetBlockedPathMonitoringOk returns a tuple with the BlockedPathMonitoring field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageExternalPath) GetBlockedPathMonitoringOk() (*int64, bool) {
	if o == nil || o.BlockedPathMonitoring == nil {
		return nil, false
	}
	return o.BlockedPathMonitoring, true
}

// HasBlockedPathMonitoring returns a boolean if a field has been set.
func (o *StorageExternalPath) HasBlockedPathMonitoring() bool {
	if o != nil && o.BlockedPathMonitoring != nil {
		return true
	}

	return false
}

// SetBlockedPathMonitoring gets a reference to the given int64 and assigns it to the BlockedPathMonitoring field.
func (o *StorageExternalPath) SetBlockedPathMonitoring(v int64) {
	o.BlockedPathMonitoring = &v
}

// GetExternalWwn returns the ExternalWwn field value if set, zero value otherwise.
func (o *StorageExternalPath) GetExternalWwn() string {
	if o == nil || o.ExternalWwn == nil {
		var ret string
		return ret
	}
	return *o.ExternalWwn
}

// GetExternalWwnOk returns a tuple with the ExternalWwn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageExternalPath) GetExternalWwnOk() (*string, bool) {
	if o == nil || o.ExternalWwn == nil {
		return nil, false
	}
	return o.ExternalWwn, true
}

// HasExternalWwn returns a boolean if a field has been set.
func (o *StorageExternalPath) HasExternalWwn() bool {
	if o != nil && o.ExternalWwn != nil {
		return true
	}

	return false
}

// SetExternalWwn gets a reference to the given string and assigns it to the ExternalWwn field.
func (o *StorageExternalPath) SetExternalWwn(v string) {
	o.ExternalWwn = &v
}

// GetIoTimeOut returns the IoTimeOut field value if set, zero value otherwise.
func (o *StorageExternalPath) GetIoTimeOut() int64 {
	if o == nil || o.IoTimeOut == nil {
		var ret int64
		return ret
	}
	return *o.IoTimeOut
}

// GetIoTimeOutOk returns a tuple with the IoTimeOut field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageExternalPath) GetIoTimeOutOk() (*int64, bool) {
	if o == nil || o.IoTimeOut == nil {
		return nil, false
	}
	return o.IoTimeOut, true
}

// HasIoTimeOut returns a boolean if a field has been set.
func (o *StorageExternalPath) HasIoTimeOut() bool {
	if o != nil && o.IoTimeOut != nil {
		return true
	}

	return false
}

// SetIoTimeOut gets a reference to the given int64 and assigns it to the IoTimeOut field.
func (o *StorageExternalPath) SetIoTimeOut(v int64) {
	o.IoTimeOut = &v
}

// GetPortId returns the PortId field value if set, zero value otherwise.
func (o *StorageExternalPath) GetPortId() string {
	if o == nil || o.PortId == nil {
		var ret string
		return ret
	}
	return *o.PortId
}

// GetPortIdOk returns a tuple with the PortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageExternalPath) GetPortIdOk() (*string, bool) {
	if o == nil || o.PortId == nil {
		return nil, false
	}
	return o.PortId, true
}

// HasPortId returns a boolean if a field has been set.
func (o *StorageExternalPath) HasPortId() bool {
	if o != nil && o.PortId != nil {
		return true
	}

	return false
}

// SetPortId gets a reference to the given string and assigns it to the PortId field.
func (o *StorageExternalPath) SetPortId(v string) {
	o.PortId = &v
}

// GetQueueDepth returns the QueueDepth field value if set, zero value otherwise.
func (o *StorageExternalPath) GetQueueDepth() int64 {
	if o == nil || o.QueueDepth == nil {
		var ret int64
		return ret
	}
	return *o.QueueDepth
}

// GetQueueDepthOk returns a tuple with the QueueDepth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageExternalPath) GetQueueDepthOk() (*int64, bool) {
	if o == nil || o.QueueDepth == nil {
		return nil, false
	}
	return o.QueueDepth, true
}

// HasQueueDepth returns a boolean if a field has been set.
func (o *StorageExternalPath) HasQueueDepth() bool {
	if o != nil && o.QueueDepth != nil {
		return true
	}

	return false
}

// SetQueueDepth gets a reference to the given int64 and assigns it to the QueueDepth field.
func (o *StorageExternalPath) SetQueueDepth(v int64) {
	o.QueueDepth = &v
}

func (o StorageExternalPath) MarshalJSON() ([]byte, error) {
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
	if o.BlockedPathMonitoring != nil {
		toSerialize["BlockedPathMonitoring"] = o.BlockedPathMonitoring
	}
	if o.ExternalWwn != nil {
		toSerialize["ExternalWwn"] = o.ExternalWwn
	}
	if o.IoTimeOut != nil {
		toSerialize["IoTimeOut"] = o.IoTimeOut
	}
	if o.PortId != nil {
		toSerialize["PortId"] = o.PortId
	}
	if o.QueueDepth != nil {
		toSerialize["QueueDepth"] = o.QueueDepth
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageExternalPath) UnmarshalJSON(bytes []byte) (err error) {
	type StorageExternalPathWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The time (in seconds) until the external parity group is blocked after all paths to the external parity group are disconnected.
		BlockedPathMonitoring *int64 `json:"BlockedPathMonitoring,omitempty"`
		// WWN of the external storage system.
		ExternalWwn *string `json:"ExternalWwn,omitempty"`
		// The value (in seconds) set for the I/O time over for the external parity group.
		IoTimeOut *int64 `json:"IoTimeOut,omitempty"`
		// Port number of external path on the local storage system.
		PortId *string `json:"PortId,omitempty"`
		// Number of Read-Write commands that can be queued to the external parity group.
		QueueDepth *int64 `json:"QueueDepth,omitempty"`
	}

	varStorageExternalPathWithoutEmbeddedStruct := StorageExternalPathWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageExternalPathWithoutEmbeddedStruct)
	if err == nil {
		varStorageExternalPath := _StorageExternalPath{}
		varStorageExternalPath.ClassId = varStorageExternalPathWithoutEmbeddedStruct.ClassId
		varStorageExternalPath.ObjectType = varStorageExternalPathWithoutEmbeddedStruct.ObjectType
		varStorageExternalPath.BlockedPathMonitoring = varStorageExternalPathWithoutEmbeddedStruct.BlockedPathMonitoring
		varStorageExternalPath.ExternalWwn = varStorageExternalPathWithoutEmbeddedStruct.ExternalWwn
		varStorageExternalPath.IoTimeOut = varStorageExternalPathWithoutEmbeddedStruct.IoTimeOut
		varStorageExternalPath.PortId = varStorageExternalPathWithoutEmbeddedStruct.PortId
		varStorageExternalPath.QueueDepth = varStorageExternalPathWithoutEmbeddedStruct.QueueDepth
		*o = StorageExternalPath(varStorageExternalPath)
	} else {
		return err
	}

	varStorageExternalPath := _StorageExternalPath{}

	err = json.Unmarshal(bytes, &varStorageExternalPath)
	if err == nil {
		o.MoBaseComplexType = varStorageExternalPath.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BlockedPathMonitoring")
		delete(additionalProperties, "ExternalWwn")
		delete(additionalProperties, "IoTimeOut")
		delete(additionalProperties, "PortId")
		delete(additionalProperties, "QueueDepth")

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

type NullableStorageExternalPath struct {
	value *StorageExternalPath
	isSet bool
}

func (v NullableStorageExternalPath) Get() *StorageExternalPath {
	return v.value
}

func (v *NullableStorageExternalPath) Set(val *StorageExternalPath) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageExternalPath) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageExternalPath) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageExternalPath(val *StorageExternalPath) *NullableStorageExternalPath {
	return &NullableStorageExternalPath{value: val, isSet: true}
}

func (v NullableStorageExternalPath) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageExternalPath) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
