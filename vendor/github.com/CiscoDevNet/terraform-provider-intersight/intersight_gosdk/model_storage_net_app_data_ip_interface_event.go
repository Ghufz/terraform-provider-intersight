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
	"reflect"
	"strings"
)

// StorageNetAppDataIpInterfaceEvent An event where the impacted resource type is an ip interface.
type StorageNetAppDataIpInterfaceEvent struct {
	StorageNetAppBaseEvent
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                                    `json:"ObjectType"`
	IpInterface          *StorageNetAppDataIpInterfaceRelationship `json:"IpInterface,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppDataIpInterfaceEvent StorageNetAppDataIpInterfaceEvent

// NewStorageNetAppDataIpInterfaceEvent instantiates a new StorageNetAppDataIpInterfaceEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppDataIpInterfaceEvent(classId string, objectType string) *StorageNetAppDataIpInterfaceEvent {
	this := StorageNetAppDataIpInterfaceEvent{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppDataIpInterfaceEventWithDefaults instantiates a new StorageNetAppDataIpInterfaceEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppDataIpInterfaceEventWithDefaults() *StorageNetAppDataIpInterfaceEvent {
	this := StorageNetAppDataIpInterfaceEvent{}
	var classId string = "storage.NetAppDataIpInterfaceEvent"
	this.ClassId = classId
	var objectType string = "storage.NetAppDataIpInterfaceEvent"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppDataIpInterfaceEvent) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppDataIpInterfaceEvent) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppDataIpInterfaceEvent) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppDataIpInterfaceEvent) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppDataIpInterfaceEvent) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppDataIpInterfaceEvent) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIpInterface returns the IpInterface field value if set, zero value otherwise.
func (o *StorageNetAppDataIpInterfaceEvent) GetIpInterface() StorageNetAppDataIpInterfaceRelationship {
	if o == nil || o.IpInterface == nil {
		var ret StorageNetAppDataIpInterfaceRelationship
		return ret
	}
	return *o.IpInterface
}

// GetIpInterfaceOk returns a tuple with the IpInterface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppDataIpInterfaceEvent) GetIpInterfaceOk() (*StorageNetAppDataIpInterfaceRelationship, bool) {
	if o == nil || o.IpInterface == nil {
		return nil, false
	}
	return o.IpInterface, true
}

// HasIpInterface returns a boolean if a field has been set.
func (o *StorageNetAppDataIpInterfaceEvent) HasIpInterface() bool {
	if o != nil && o.IpInterface != nil {
		return true
	}

	return false
}

// SetIpInterface gets a reference to the given StorageNetAppDataIpInterfaceRelationship and assigns it to the IpInterface field.
func (o *StorageNetAppDataIpInterfaceEvent) SetIpInterface(v StorageNetAppDataIpInterfaceRelationship) {
	o.IpInterface = &v
}

func (o StorageNetAppDataIpInterfaceEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedStorageNetAppBaseEvent, errStorageNetAppBaseEvent := json.Marshal(o.StorageNetAppBaseEvent)
	if errStorageNetAppBaseEvent != nil {
		return []byte{}, errStorageNetAppBaseEvent
	}
	errStorageNetAppBaseEvent = json.Unmarshal([]byte(serializedStorageNetAppBaseEvent), &toSerialize)
	if errStorageNetAppBaseEvent != nil {
		return []byte{}, errStorageNetAppBaseEvent
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.IpInterface != nil {
		toSerialize["IpInterface"] = o.IpInterface
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageNetAppDataIpInterfaceEvent) UnmarshalJSON(bytes []byte) (err error) {
	type StorageNetAppDataIpInterfaceEventWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType  string                                    `json:"ObjectType"`
		IpInterface *StorageNetAppDataIpInterfaceRelationship `json:"IpInterface,omitempty"`
	}

	varStorageNetAppDataIpInterfaceEventWithoutEmbeddedStruct := StorageNetAppDataIpInterfaceEventWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageNetAppDataIpInterfaceEventWithoutEmbeddedStruct)
	if err == nil {
		varStorageNetAppDataIpInterfaceEvent := _StorageNetAppDataIpInterfaceEvent{}
		varStorageNetAppDataIpInterfaceEvent.ClassId = varStorageNetAppDataIpInterfaceEventWithoutEmbeddedStruct.ClassId
		varStorageNetAppDataIpInterfaceEvent.ObjectType = varStorageNetAppDataIpInterfaceEventWithoutEmbeddedStruct.ObjectType
		varStorageNetAppDataIpInterfaceEvent.IpInterface = varStorageNetAppDataIpInterfaceEventWithoutEmbeddedStruct.IpInterface
		*o = StorageNetAppDataIpInterfaceEvent(varStorageNetAppDataIpInterfaceEvent)
	} else {
		return err
	}

	varStorageNetAppDataIpInterfaceEvent := _StorageNetAppDataIpInterfaceEvent{}

	err = json.Unmarshal(bytes, &varStorageNetAppDataIpInterfaceEvent)
	if err == nil {
		o.StorageNetAppBaseEvent = varStorageNetAppDataIpInterfaceEvent.StorageNetAppBaseEvent
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "IpInterface")

		// remove fields from embedded structs
		reflectStorageNetAppBaseEvent := reflect.ValueOf(o.StorageNetAppBaseEvent)
		for i := 0; i < reflectStorageNetAppBaseEvent.Type().NumField(); i++ {
			t := reflectStorageNetAppBaseEvent.Type().Field(i)

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

type NullableStorageNetAppDataIpInterfaceEvent struct {
	value *StorageNetAppDataIpInterfaceEvent
	isSet bool
}

func (v NullableStorageNetAppDataIpInterfaceEvent) Get() *StorageNetAppDataIpInterfaceEvent {
	return v.value
}

func (v *NullableStorageNetAppDataIpInterfaceEvent) Set(val *StorageNetAppDataIpInterfaceEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppDataIpInterfaceEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppDataIpInterfaceEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppDataIpInterfaceEvent(val *StorageNetAppDataIpInterfaceEvent) *NullableStorageNetAppDataIpInterfaceEvent {
	return &NullableStorageNetAppDataIpInterfaceEvent{value: val, isSet: true}
}

func (v NullableStorageNetAppDataIpInterfaceEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppDataIpInterfaceEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
