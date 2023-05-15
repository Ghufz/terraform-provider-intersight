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

// StorageNetAppCifsShare NetApp CIFS share is a named access point in a volume which is tied to the CIFS server on the SVM.
type StorageNetAppCifsShare struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Description of the CIFS share.
	Comment *string `json:"Comment,omitempty"`
	// Indicates that SMB encryption must be used when accessing the share.
	Encryption *string `json:"Encryption,omitempty"`
	// Indicates that the share is a home directory share, where the share and path names are dynamic.
	HomeDirectory *string `json:"HomeDirectory,omitempty"`
	// Name of the NetApp CIFS share.
	Name          *string                `json:"Name,omitempty"`
	NetAppCifsAcl []StorageNetAppCifsAcl `json:"NetAppCifsAcl,omitempty"`
	// The fully-qualified pathname in the owning SVM namespace that is shared through the share.
	Path *string `json:"Path,omitempty"`
	// The storage virtual machine name for the CIFS share.
	SvmName *string `json:"SvmName,omitempty"`
	// Unique identifier for the NetApp Storage Virtual Machine.
	SvmUuid              *string                             `json:"SvmUuid,omitempty"`
	StorageContainer     *StorageNetAppVolumeRelationship    `json:"StorageContainer,omitempty"`
	Tenant               *StorageNetAppStorageVmRelationship `json:"Tenant,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppCifsShare StorageNetAppCifsShare

// NewStorageNetAppCifsShare instantiates a new StorageNetAppCifsShare object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppCifsShare(classId string, objectType string) *StorageNetAppCifsShare {
	this := StorageNetAppCifsShare{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppCifsShareWithDefaults instantiates a new StorageNetAppCifsShare object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppCifsShareWithDefaults() *StorageNetAppCifsShare {
	this := StorageNetAppCifsShare{}
	var classId string = "storage.NetAppCifsShare"
	this.ClassId = classId
	var objectType string = "storage.NetAppCifsShare"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppCifsShare) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppCifsShare) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppCifsShare) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppCifsShare) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppCifsShare) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppCifsShare) SetObjectType(v string) {
	o.ObjectType = v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *StorageNetAppCifsShare) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppCifsShare) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *StorageNetAppCifsShare) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *StorageNetAppCifsShare) SetComment(v string) {
	o.Comment = &v
}

// GetEncryption returns the Encryption field value if set, zero value otherwise.
func (o *StorageNetAppCifsShare) GetEncryption() string {
	if o == nil || o.Encryption == nil {
		var ret string
		return ret
	}
	return *o.Encryption
}

// GetEncryptionOk returns a tuple with the Encryption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppCifsShare) GetEncryptionOk() (*string, bool) {
	if o == nil || o.Encryption == nil {
		return nil, false
	}
	return o.Encryption, true
}

// HasEncryption returns a boolean if a field has been set.
func (o *StorageNetAppCifsShare) HasEncryption() bool {
	if o != nil && o.Encryption != nil {
		return true
	}

	return false
}

// SetEncryption gets a reference to the given string and assigns it to the Encryption field.
func (o *StorageNetAppCifsShare) SetEncryption(v string) {
	o.Encryption = &v
}

// GetHomeDirectory returns the HomeDirectory field value if set, zero value otherwise.
func (o *StorageNetAppCifsShare) GetHomeDirectory() string {
	if o == nil || o.HomeDirectory == nil {
		var ret string
		return ret
	}
	return *o.HomeDirectory
}

// GetHomeDirectoryOk returns a tuple with the HomeDirectory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppCifsShare) GetHomeDirectoryOk() (*string, bool) {
	if o == nil || o.HomeDirectory == nil {
		return nil, false
	}
	return o.HomeDirectory, true
}

// HasHomeDirectory returns a boolean if a field has been set.
func (o *StorageNetAppCifsShare) HasHomeDirectory() bool {
	if o != nil && o.HomeDirectory != nil {
		return true
	}

	return false
}

// SetHomeDirectory gets a reference to the given string and assigns it to the HomeDirectory field.
func (o *StorageNetAppCifsShare) SetHomeDirectory(v string) {
	o.HomeDirectory = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StorageNetAppCifsShare) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppCifsShare) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StorageNetAppCifsShare) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StorageNetAppCifsShare) SetName(v string) {
	o.Name = &v
}

// GetNetAppCifsAcl returns the NetAppCifsAcl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppCifsShare) GetNetAppCifsAcl() []StorageNetAppCifsAcl {
	if o == nil {
		var ret []StorageNetAppCifsAcl
		return ret
	}
	return o.NetAppCifsAcl
}

// GetNetAppCifsAclOk returns a tuple with the NetAppCifsAcl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppCifsShare) GetNetAppCifsAclOk() ([]StorageNetAppCifsAcl, bool) {
	if o == nil || o.NetAppCifsAcl == nil {
		return nil, false
	}
	return o.NetAppCifsAcl, true
}

// HasNetAppCifsAcl returns a boolean if a field has been set.
func (o *StorageNetAppCifsShare) HasNetAppCifsAcl() bool {
	if o != nil && o.NetAppCifsAcl != nil {
		return true
	}

	return false
}

// SetNetAppCifsAcl gets a reference to the given []StorageNetAppCifsAcl and assigns it to the NetAppCifsAcl field.
func (o *StorageNetAppCifsShare) SetNetAppCifsAcl(v []StorageNetAppCifsAcl) {
	o.NetAppCifsAcl = v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *StorageNetAppCifsShare) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppCifsShare) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *StorageNetAppCifsShare) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *StorageNetAppCifsShare) SetPath(v string) {
	o.Path = &v
}

// GetSvmName returns the SvmName field value if set, zero value otherwise.
func (o *StorageNetAppCifsShare) GetSvmName() string {
	if o == nil || o.SvmName == nil {
		var ret string
		return ret
	}
	return *o.SvmName
}

// GetSvmNameOk returns a tuple with the SvmName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppCifsShare) GetSvmNameOk() (*string, bool) {
	if o == nil || o.SvmName == nil {
		return nil, false
	}
	return o.SvmName, true
}

// HasSvmName returns a boolean if a field has been set.
func (o *StorageNetAppCifsShare) HasSvmName() bool {
	if o != nil && o.SvmName != nil {
		return true
	}

	return false
}

// SetSvmName gets a reference to the given string and assigns it to the SvmName field.
func (o *StorageNetAppCifsShare) SetSvmName(v string) {
	o.SvmName = &v
}

// GetSvmUuid returns the SvmUuid field value if set, zero value otherwise.
func (o *StorageNetAppCifsShare) GetSvmUuid() string {
	if o == nil || o.SvmUuid == nil {
		var ret string
		return ret
	}
	return *o.SvmUuid
}

// GetSvmUuidOk returns a tuple with the SvmUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppCifsShare) GetSvmUuidOk() (*string, bool) {
	if o == nil || o.SvmUuid == nil {
		return nil, false
	}
	return o.SvmUuid, true
}

// HasSvmUuid returns a boolean if a field has been set.
func (o *StorageNetAppCifsShare) HasSvmUuid() bool {
	if o != nil && o.SvmUuid != nil {
		return true
	}

	return false
}

// SetSvmUuid gets a reference to the given string and assigns it to the SvmUuid field.
func (o *StorageNetAppCifsShare) SetSvmUuid(v string) {
	o.SvmUuid = &v
}

// GetStorageContainer returns the StorageContainer field value if set, zero value otherwise.
func (o *StorageNetAppCifsShare) GetStorageContainer() StorageNetAppVolumeRelationship {
	if o == nil || o.StorageContainer == nil {
		var ret StorageNetAppVolumeRelationship
		return ret
	}
	return *o.StorageContainer
}

// GetStorageContainerOk returns a tuple with the StorageContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppCifsShare) GetStorageContainerOk() (*StorageNetAppVolumeRelationship, bool) {
	if o == nil || o.StorageContainer == nil {
		return nil, false
	}
	return o.StorageContainer, true
}

// HasStorageContainer returns a boolean if a field has been set.
func (o *StorageNetAppCifsShare) HasStorageContainer() bool {
	if o != nil && o.StorageContainer != nil {
		return true
	}

	return false
}

// SetStorageContainer gets a reference to the given StorageNetAppVolumeRelationship and assigns it to the StorageContainer field.
func (o *StorageNetAppCifsShare) SetStorageContainer(v StorageNetAppVolumeRelationship) {
	o.StorageContainer = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *StorageNetAppCifsShare) GetTenant() StorageNetAppStorageVmRelationship {
	if o == nil || o.Tenant == nil {
		var ret StorageNetAppStorageVmRelationship
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppCifsShare) GetTenantOk() (*StorageNetAppStorageVmRelationship, bool) {
	if o == nil || o.Tenant == nil {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *StorageNetAppCifsShare) HasTenant() bool {
	if o != nil && o.Tenant != nil {
		return true
	}

	return false
}

// SetTenant gets a reference to the given StorageNetAppStorageVmRelationship and assigns it to the Tenant field.
func (o *StorageNetAppCifsShare) SetTenant(v StorageNetAppStorageVmRelationship) {
	o.Tenant = &v
}

func (o StorageNetAppCifsShare) MarshalJSON() ([]byte, error) {
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
	if o.Comment != nil {
		toSerialize["Comment"] = o.Comment
	}
	if o.Encryption != nil {
		toSerialize["Encryption"] = o.Encryption
	}
	if o.HomeDirectory != nil {
		toSerialize["HomeDirectory"] = o.HomeDirectory
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.NetAppCifsAcl != nil {
		toSerialize["NetAppCifsAcl"] = o.NetAppCifsAcl
	}
	if o.Path != nil {
		toSerialize["Path"] = o.Path
	}
	if o.SvmName != nil {
		toSerialize["SvmName"] = o.SvmName
	}
	if o.SvmUuid != nil {
		toSerialize["SvmUuid"] = o.SvmUuid
	}
	if o.StorageContainer != nil {
		toSerialize["StorageContainer"] = o.StorageContainer
	}
	if o.Tenant != nil {
		toSerialize["Tenant"] = o.Tenant
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageNetAppCifsShare) UnmarshalJSON(bytes []byte) (err error) {
	type StorageNetAppCifsShareWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Description of the CIFS share.
		Comment *string `json:"Comment,omitempty"`
		// Indicates that SMB encryption must be used when accessing the share.
		Encryption *string `json:"Encryption,omitempty"`
		// Indicates that the share is a home directory share, where the share and path names are dynamic.
		HomeDirectory *string `json:"HomeDirectory,omitempty"`
		// Name of the NetApp CIFS share.
		Name          *string                `json:"Name,omitempty"`
		NetAppCifsAcl []StorageNetAppCifsAcl `json:"NetAppCifsAcl,omitempty"`
		// The fully-qualified pathname in the owning SVM namespace that is shared through the share.
		Path *string `json:"Path,omitempty"`
		// The storage virtual machine name for the CIFS share.
		SvmName *string `json:"SvmName,omitempty"`
		// Unique identifier for the NetApp Storage Virtual Machine.
		SvmUuid          *string                             `json:"SvmUuid,omitempty"`
		StorageContainer *StorageNetAppVolumeRelationship    `json:"StorageContainer,omitempty"`
		Tenant           *StorageNetAppStorageVmRelationship `json:"Tenant,omitempty"`
	}

	varStorageNetAppCifsShareWithoutEmbeddedStruct := StorageNetAppCifsShareWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageNetAppCifsShareWithoutEmbeddedStruct)
	if err == nil {
		varStorageNetAppCifsShare := _StorageNetAppCifsShare{}
		varStorageNetAppCifsShare.ClassId = varStorageNetAppCifsShareWithoutEmbeddedStruct.ClassId
		varStorageNetAppCifsShare.ObjectType = varStorageNetAppCifsShareWithoutEmbeddedStruct.ObjectType
		varStorageNetAppCifsShare.Comment = varStorageNetAppCifsShareWithoutEmbeddedStruct.Comment
		varStorageNetAppCifsShare.Encryption = varStorageNetAppCifsShareWithoutEmbeddedStruct.Encryption
		varStorageNetAppCifsShare.HomeDirectory = varStorageNetAppCifsShareWithoutEmbeddedStruct.HomeDirectory
		varStorageNetAppCifsShare.Name = varStorageNetAppCifsShareWithoutEmbeddedStruct.Name
		varStorageNetAppCifsShare.NetAppCifsAcl = varStorageNetAppCifsShareWithoutEmbeddedStruct.NetAppCifsAcl
		varStorageNetAppCifsShare.Path = varStorageNetAppCifsShareWithoutEmbeddedStruct.Path
		varStorageNetAppCifsShare.SvmName = varStorageNetAppCifsShareWithoutEmbeddedStruct.SvmName
		varStorageNetAppCifsShare.SvmUuid = varStorageNetAppCifsShareWithoutEmbeddedStruct.SvmUuid
		varStorageNetAppCifsShare.StorageContainer = varStorageNetAppCifsShareWithoutEmbeddedStruct.StorageContainer
		varStorageNetAppCifsShare.Tenant = varStorageNetAppCifsShareWithoutEmbeddedStruct.Tenant
		*o = StorageNetAppCifsShare(varStorageNetAppCifsShare)
	} else {
		return err
	}

	varStorageNetAppCifsShare := _StorageNetAppCifsShare{}

	err = json.Unmarshal(bytes, &varStorageNetAppCifsShare)
	if err == nil {
		o.MoBaseMo = varStorageNetAppCifsShare.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Comment")
		delete(additionalProperties, "Encryption")
		delete(additionalProperties, "HomeDirectory")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "NetAppCifsAcl")
		delete(additionalProperties, "Path")
		delete(additionalProperties, "SvmName")
		delete(additionalProperties, "SvmUuid")
		delete(additionalProperties, "StorageContainer")
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

type NullableStorageNetAppCifsShare struct {
	value *StorageNetAppCifsShare
	isSet bool
}

func (v NullableStorageNetAppCifsShare) Get() *StorageNetAppCifsShare {
	return v.value
}

func (v *NullableStorageNetAppCifsShare) Set(val *StorageNetAppCifsShare) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppCifsShare) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppCifsShare) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppCifsShare(val *StorageNetAppCifsShare) *NullableStorageNetAppCifsShare {
	return &NullableStorageNetAppCifsShare{value: val, isSet: true}
}

func (v NullableStorageNetAppCifsShare) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppCifsShare) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
