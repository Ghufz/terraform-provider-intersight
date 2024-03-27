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

// StorageNetAppSvmEventAllOf Definition of the list of properties defined in 'storage.NetAppSvmEvent', excluding properties defined in parent classes.
type StorageNetAppSvmEventAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                              `json:"ObjectType"`
	Tenant               *StorageNetAppStorageVmRelationship `json:"Tenant,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppSvmEventAllOf StorageNetAppSvmEventAllOf

// NewStorageNetAppSvmEventAllOf instantiates a new StorageNetAppSvmEventAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppSvmEventAllOf(classId string, objectType string) *StorageNetAppSvmEventAllOf {
	this := StorageNetAppSvmEventAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppSvmEventAllOfWithDefaults instantiates a new StorageNetAppSvmEventAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppSvmEventAllOfWithDefaults() *StorageNetAppSvmEventAllOf {
	this := StorageNetAppSvmEventAllOf{}
	var classId string = "storage.NetAppSvmEvent"
	this.ClassId = classId
	var objectType string = "storage.NetAppSvmEvent"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppSvmEventAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppSvmEventAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppSvmEventAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppSvmEventAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppSvmEventAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppSvmEventAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *StorageNetAppSvmEventAllOf) GetTenant() StorageNetAppStorageVmRelationship {
	if o == nil || o.Tenant == nil {
		var ret StorageNetAppStorageVmRelationship
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppSvmEventAllOf) GetTenantOk() (*StorageNetAppStorageVmRelationship, bool) {
	if o == nil || o.Tenant == nil {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *StorageNetAppSvmEventAllOf) HasTenant() bool {
	if o != nil && o.Tenant != nil {
		return true
	}

	return false
}

// SetTenant gets a reference to the given StorageNetAppStorageVmRelationship and assigns it to the Tenant field.
func (o *StorageNetAppSvmEventAllOf) SetTenant(v StorageNetAppStorageVmRelationship) {
	o.Tenant = &v
}

func (o StorageNetAppSvmEventAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Tenant != nil {
		toSerialize["Tenant"] = o.Tenant
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageNetAppSvmEventAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageNetAppSvmEventAllOf := _StorageNetAppSvmEventAllOf{}

	if err = json.Unmarshal(bytes, &varStorageNetAppSvmEventAllOf); err == nil {
		*o = StorageNetAppSvmEventAllOf(varStorageNetAppSvmEventAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Tenant")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageNetAppSvmEventAllOf struct {
	value *StorageNetAppSvmEventAllOf
	isSet bool
}

func (v NullableStorageNetAppSvmEventAllOf) Get() *StorageNetAppSvmEventAllOf {
	return v.value
}

func (v *NullableStorageNetAppSvmEventAllOf) Set(val *StorageNetAppSvmEventAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppSvmEventAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppSvmEventAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppSvmEventAllOf(val *StorageNetAppSvmEventAllOf) *NullableStorageNetAppSvmEventAllOf {
	return &NullableStorageNetAppSvmEventAllOf{value: val, isSet: true}
}

func (v NullableStorageNetAppSvmEventAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppSvmEventAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
