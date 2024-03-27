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

// StorageNetAppIscsiService An iSCSI service defines the properties of the iSCSI target for an SVM. There can be at most one iSCSI service for an SVM. An SVM's iSCSI service must be created before iSCSI initiators can log in to the SVM.
type StorageNetAppIscsiService struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Unique identifier for the NetApp Storage Virtual Machine.
	SvmUuid *string `json:"SvmUuid,omitempty"`
	// The iSCSI target alias of the iSCSI service.
	TargetAlias *string `json:"TargetAlias,omitempty"`
	// The iSCSI target name of the iSCSI service.
	TargetName           *string                             `json:"TargetName,omitempty"`
	Tenant               *StorageNetAppStorageVmRelationship `json:"Tenant,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppIscsiService StorageNetAppIscsiService

// NewStorageNetAppIscsiService instantiates a new StorageNetAppIscsiService object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppIscsiService(classId string, objectType string) *StorageNetAppIscsiService {
	this := StorageNetAppIscsiService{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppIscsiServiceWithDefaults instantiates a new StorageNetAppIscsiService object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppIscsiServiceWithDefaults() *StorageNetAppIscsiService {
	this := StorageNetAppIscsiService{}
	var classId string = "storage.NetAppIscsiService"
	this.ClassId = classId
	var objectType string = "storage.NetAppIscsiService"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppIscsiService) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppIscsiService) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppIscsiService) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppIscsiService) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppIscsiService) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppIscsiService) SetObjectType(v string) {
	o.ObjectType = v
}

// GetSvmUuid returns the SvmUuid field value if set, zero value otherwise.
func (o *StorageNetAppIscsiService) GetSvmUuid() string {
	if o == nil || o.SvmUuid == nil {
		var ret string
		return ret
	}
	return *o.SvmUuid
}

// GetSvmUuidOk returns a tuple with the SvmUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppIscsiService) GetSvmUuidOk() (*string, bool) {
	if o == nil || o.SvmUuid == nil {
		return nil, false
	}
	return o.SvmUuid, true
}

// HasSvmUuid returns a boolean if a field has been set.
func (o *StorageNetAppIscsiService) HasSvmUuid() bool {
	if o != nil && o.SvmUuid != nil {
		return true
	}

	return false
}

// SetSvmUuid gets a reference to the given string and assigns it to the SvmUuid field.
func (o *StorageNetAppIscsiService) SetSvmUuid(v string) {
	o.SvmUuid = &v
}

// GetTargetAlias returns the TargetAlias field value if set, zero value otherwise.
func (o *StorageNetAppIscsiService) GetTargetAlias() string {
	if o == nil || o.TargetAlias == nil {
		var ret string
		return ret
	}
	return *o.TargetAlias
}

// GetTargetAliasOk returns a tuple with the TargetAlias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppIscsiService) GetTargetAliasOk() (*string, bool) {
	if o == nil || o.TargetAlias == nil {
		return nil, false
	}
	return o.TargetAlias, true
}

// HasTargetAlias returns a boolean if a field has been set.
func (o *StorageNetAppIscsiService) HasTargetAlias() bool {
	if o != nil && o.TargetAlias != nil {
		return true
	}

	return false
}

// SetTargetAlias gets a reference to the given string and assigns it to the TargetAlias field.
func (o *StorageNetAppIscsiService) SetTargetAlias(v string) {
	o.TargetAlias = &v
}

// GetTargetName returns the TargetName field value if set, zero value otherwise.
func (o *StorageNetAppIscsiService) GetTargetName() string {
	if o == nil || o.TargetName == nil {
		var ret string
		return ret
	}
	return *o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppIscsiService) GetTargetNameOk() (*string, bool) {
	if o == nil || o.TargetName == nil {
		return nil, false
	}
	return o.TargetName, true
}

// HasTargetName returns a boolean if a field has been set.
func (o *StorageNetAppIscsiService) HasTargetName() bool {
	if o != nil && o.TargetName != nil {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given string and assigns it to the TargetName field.
func (o *StorageNetAppIscsiService) SetTargetName(v string) {
	o.TargetName = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *StorageNetAppIscsiService) GetTenant() StorageNetAppStorageVmRelationship {
	if o == nil || o.Tenant == nil {
		var ret StorageNetAppStorageVmRelationship
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppIscsiService) GetTenantOk() (*StorageNetAppStorageVmRelationship, bool) {
	if o == nil || o.Tenant == nil {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *StorageNetAppIscsiService) HasTenant() bool {
	if o != nil && o.Tenant != nil {
		return true
	}

	return false
}

// SetTenant gets a reference to the given StorageNetAppStorageVmRelationship and assigns it to the Tenant field.
func (o *StorageNetAppIscsiService) SetTenant(v StorageNetAppStorageVmRelationship) {
	o.Tenant = &v
}

func (o StorageNetAppIscsiService) MarshalJSON() ([]byte, error) {
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
	if o.SvmUuid != nil {
		toSerialize["SvmUuid"] = o.SvmUuid
	}
	if o.TargetAlias != nil {
		toSerialize["TargetAlias"] = o.TargetAlias
	}
	if o.TargetName != nil {
		toSerialize["TargetName"] = o.TargetName
	}
	if o.Tenant != nil {
		toSerialize["Tenant"] = o.Tenant
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageNetAppIscsiService) UnmarshalJSON(bytes []byte) (err error) {
	type StorageNetAppIscsiServiceWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Unique identifier for the NetApp Storage Virtual Machine.
		SvmUuid *string `json:"SvmUuid,omitempty"`
		// The iSCSI target alias of the iSCSI service.
		TargetAlias *string `json:"TargetAlias,omitempty"`
		// The iSCSI target name of the iSCSI service.
		TargetName *string                             `json:"TargetName,omitempty"`
		Tenant     *StorageNetAppStorageVmRelationship `json:"Tenant,omitempty"`
	}

	varStorageNetAppIscsiServiceWithoutEmbeddedStruct := StorageNetAppIscsiServiceWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageNetAppIscsiServiceWithoutEmbeddedStruct)
	if err == nil {
		varStorageNetAppIscsiService := _StorageNetAppIscsiService{}
		varStorageNetAppIscsiService.ClassId = varStorageNetAppIscsiServiceWithoutEmbeddedStruct.ClassId
		varStorageNetAppIscsiService.ObjectType = varStorageNetAppIscsiServiceWithoutEmbeddedStruct.ObjectType
		varStorageNetAppIscsiService.SvmUuid = varStorageNetAppIscsiServiceWithoutEmbeddedStruct.SvmUuid
		varStorageNetAppIscsiService.TargetAlias = varStorageNetAppIscsiServiceWithoutEmbeddedStruct.TargetAlias
		varStorageNetAppIscsiService.TargetName = varStorageNetAppIscsiServiceWithoutEmbeddedStruct.TargetName
		varStorageNetAppIscsiService.Tenant = varStorageNetAppIscsiServiceWithoutEmbeddedStruct.Tenant
		*o = StorageNetAppIscsiService(varStorageNetAppIscsiService)
	} else {
		return err
	}

	varStorageNetAppIscsiService := _StorageNetAppIscsiService{}

	err = json.Unmarshal(bytes, &varStorageNetAppIscsiService)
	if err == nil {
		o.MoBaseMo = varStorageNetAppIscsiService.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "SvmUuid")
		delete(additionalProperties, "TargetAlias")
		delete(additionalProperties, "TargetName")
		delete(additionalProperties, "Tenant")

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

type NullableStorageNetAppIscsiService struct {
	value *StorageNetAppIscsiService
	isSet bool
}

func (v NullableStorageNetAppIscsiService) Get() *StorageNetAppIscsiService {
	return v.value
}

func (v *NullableStorageNetAppIscsiService) Set(val *StorageNetAppIscsiService) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppIscsiService) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppIscsiService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppIscsiService(val *StorageNetAppIscsiService) *NullableStorageNetAppIscsiService {
	return &NullableStorageNetAppIscsiService{value: val, isSet: true}
}

func (v NullableStorageNetAppIscsiService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppIscsiService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
