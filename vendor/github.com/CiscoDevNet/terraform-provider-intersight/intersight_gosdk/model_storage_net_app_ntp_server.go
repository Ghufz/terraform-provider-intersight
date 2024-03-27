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

// StorageNetAppNtpServer External NTP time servers ONTAP uses for time adjustment and correction.
type StorageNetAppNtpServer struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Indicates whether or not NTP symmetric authentication is enabled.
	// Deprecated
	AuthenticationEnabled *bool `json:"AuthenticationEnabled,omitempty"`
	// NTP symmetric authentication key identifier or index number (ID).
	AuthenticationKeyId *string `json:"AuthenticationKeyId,omitempty"`
	// Unique identity of the device.
	ClusterUuid *string `json:"ClusterUuid,omitempty"`
	// Indicates whether or not NTP symmetric authentication is enabled.
	IsAuthenticationEnabled *string `json:"IsAuthenticationEnabled,omitempty"`
	// NTP server host name, IPv4, or IPv6 address.
	Server *string `json:"Server,omitempty"`
	// NTP protocol version for server. Valid versions are 3, 4, or auto. * `none` - Default unknown NTP protocol version. * `3` - NTP protocol version is 3. * `4` - NTP protocol version is 4. * `auto` - NTP protocol version is auto.
	Version              *string                           `json:"Version,omitempty"`
	Array                *StorageNetAppClusterRelationship `json:"Array,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppNtpServer StorageNetAppNtpServer

// NewStorageNetAppNtpServer instantiates a new StorageNetAppNtpServer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppNtpServer(classId string, objectType string) *StorageNetAppNtpServer {
	this := StorageNetAppNtpServer{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppNtpServerWithDefaults instantiates a new StorageNetAppNtpServer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppNtpServerWithDefaults() *StorageNetAppNtpServer {
	this := StorageNetAppNtpServer{}
	var classId string = "storage.NetAppNtpServer"
	this.ClassId = classId
	var objectType string = "storage.NetAppNtpServer"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppNtpServer) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppNtpServer) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppNtpServer) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppNtpServer) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppNtpServer) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppNtpServer) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAuthenticationEnabled returns the AuthenticationEnabled field value if set, zero value otherwise.
// Deprecated
func (o *StorageNetAppNtpServer) GetAuthenticationEnabled() bool {
	if o == nil || o.AuthenticationEnabled == nil {
		var ret bool
		return ret
	}
	return *o.AuthenticationEnabled
}

// GetAuthenticationEnabledOk returns a tuple with the AuthenticationEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *StorageNetAppNtpServer) GetAuthenticationEnabledOk() (*bool, bool) {
	if o == nil || o.AuthenticationEnabled == nil {
		return nil, false
	}
	return o.AuthenticationEnabled, true
}

// HasAuthenticationEnabled returns a boolean if a field has been set.
func (o *StorageNetAppNtpServer) HasAuthenticationEnabled() bool {
	if o != nil && o.AuthenticationEnabled != nil {
		return true
	}

	return false
}

// SetAuthenticationEnabled gets a reference to the given bool and assigns it to the AuthenticationEnabled field.
// Deprecated
func (o *StorageNetAppNtpServer) SetAuthenticationEnabled(v bool) {
	o.AuthenticationEnabled = &v
}

// GetAuthenticationKeyId returns the AuthenticationKeyId field value if set, zero value otherwise.
func (o *StorageNetAppNtpServer) GetAuthenticationKeyId() string {
	if o == nil || o.AuthenticationKeyId == nil {
		var ret string
		return ret
	}
	return *o.AuthenticationKeyId
}

// GetAuthenticationKeyIdOk returns a tuple with the AuthenticationKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppNtpServer) GetAuthenticationKeyIdOk() (*string, bool) {
	if o == nil || o.AuthenticationKeyId == nil {
		return nil, false
	}
	return o.AuthenticationKeyId, true
}

// HasAuthenticationKeyId returns a boolean if a field has been set.
func (o *StorageNetAppNtpServer) HasAuthenticationKeyId() bool {
	if o != nil && o.AuthenticationKeyId != nil {
		return true
	}

	return false
}

// SetAuthenticationKeyId gets a reference to the given string and assigns it to the AuthenticationKeyId field.
func (o *StorageNetAppNtpServer) SetAuthenticationKeyId(v string) {
	o.AuthenticationKeyId = &v
}

// GetClusterUuid returns the ClusterUuid field value if set, zero value otherwise.
func (o *StorageNetAppNtpServer) GetClusterUuid() string {
	if o == nil || o.ClusterUuid == nil {
		var ret string
		return ret
	}
	return *o.ClusterUuid
}

// GetClusterUuidOk returns a tuple with the ClusterUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppNtpServer) GetClusterUuidOk() (*string, bool) {
	if o == nil || o.ClusterUuid == nil {
		return nil, false
	}
	return o.ClusterUuid, true
}

// HasClusterUuid returns a boolean if a field has been set.
func (o *StorageNetAppNtpServer) HasClusterUuid() bool {
	if o != nil && o.ClusterUuid != nil {
		return true
	}

	return false
}

// SetClusterUuid gets a reference to the given string and assigns it to the ClusterUuid field.
func (o *StorageNetAppNtpServer) SetClusterUuid(v string) {
	o.ClusterUuid = &v
}

// GetIsAuthenticationEnabled returns the IsAuthenticationEnabled field value if set, zero value otherwise.
func (o *StorageNetAppNtpServer) GetIsAuthenticationEnabled() string {
	if o == nil || o.IsAuthenticationEnabled == nil {
		var ret string
		return ret
	}
	return *o.IsAuthenticationEnabled
}

// GetIsAuthenticationEnabledOk returns a tuple with the IsAuthenticationEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppNtpServer) GetIsAuthenticationEnabledOk() (*string, bool) {
	if o == nil || o.IsAuthenticationEnabled == nil {
		return nil, false
	}
	return o.IsAuthenticationEnabled, true
}

// HasIsAuthenticationEnabled returns a boolean if a field has been set.
func (o *StorageNetAppNtpServer) HasIsAuthenticationEnabled() bool {
	if o != nil && o.IsAuthenticationEnabled != nil {
		return true
	}

	return false
}

// SetIsAuthenticationEnabled gets a reference to the given string and assigns it to the IsAuthenticationEnabled field.
func (o *StorageNetAppNtpServer) SetIsAuthenticationEnabled(v string) {
	o.IsAuthenticationEnabled = &v
}

// GetServer returns the Server field value if set, zero value otherwise.
func (o *StorageNetAppNtpServer) GetServer() string {
	if o == nil || o.Server == nil {
		var ret string
		return ret
	}
	return *o.Server
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppNtpServer) GetServerOk() (*string, bool) {
	if o == nil || o.Server == nil {
		return nil, false
	}
	return o.Server, true
}

// HasServer returns a boolean if a field has been set.
func (o *StorageNetAppNtpServer) HasServer() bool {
	if o != nil && o.Server != nil {
		return true
	}

	return false
}

// SetServer gets a reference to the given string and assigns it to the Server field.
func (o *StorageNetAppNtpServer) SetServer(v string) {
	o.Server = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *StorageNetAppNtpServer) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppNtpServer) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *StorageNetAppNtpServer) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *StorageNetAppNtpServer) SetVersion(v string) {
	o.Version = &v
}

// GetArray returns the Array field value if set, zero value otherwise.
func (o *StorageNetAppNtpServer) GetArray() StorageNetAppClusterRelationship {
	if o == nil || o.Array == nil {
		var ret StorageNetAppClusterRelationship
		return ret
	}
	return *o.Array
}

// GetArrayOk returns a tuple with the Array field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppNtpServer) GetArrayOk() (*StorageNetAppClusterRelationship, bool) {
	if o == nil || o.Array == nil {
		return nil, false
	}
	return o.Array, true
}

// HasArray returns a boolean if a field has been set.
func (o *StorageNetAppNtpServer) HasArray() bool {
	if o != nil && o.Array != nil {
		return true
	}

	return false
}

// SetArray gets a reference to the given StorageNetAppClusterRelationship and assigns it to the Array field.
func (o *StorageNetAppNtpServer) SetArray(v StorageNetAppClusterRelationship) {
	o.Array = &v
}

func (o StorageNetAppNtpServer) MarshalJSON() ([]byte, error) {
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
	if o.AuthenticationEnabled != nil {
		toSerialize["AuthenticationEnabled"] = o.AuthenticationEnabled
	}
	if o.AuthenticationKeyId != nil {
		toSerialize["AuthenticationKeyId"] = o.AuthenticationKeyId
	}
	if o.ClusterUuid != nil {
		toSerialize["ClusterUuid"] = o.ClusterUuid
	}
	if o.IsAuthenticationEnabled != nil {
		toSerialize["IsAuthenticationEnabled"] = o.IsAuthenticationEnabled
	}
	if o.Server != nil {
		toSerialize["Server"] = o.Server
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	if o.Array != nil {
		toSerialize["Array"] = o.Array
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageNetAppNtpServer) UnmarshalJSON(bytes []byte) (err error) {
	type StorageNetAppNtpServerWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Indicates whether or not NTP symmetric authentication is enabled.
		// Deprecated
		AuthenticationEnabled *bool `json:"AuthenticationEnabled,omitempty"`
		// NTP symmetric authentication key identifier or index number (ID).
		AuthenticationKeyId *string `json:"AuthenticationKeyId,omitempty"`
		// Unique identity of the device.
		ClusterUuid *string `json:"ClusterUuid,omitempty"`
		// Indicates whether or not NTP symmetric authentication is enabled.
		IsAuthenticationEnabled *string `json:"IsAuthenticationEnabled,omitempty"`
		// NTP server host name, IPv4, or IPv6 address.
		Server *string `json:"Server,omitempty"`
		// NTP protocol version for server. Valid versions are 3, 4, or auto. * `none` - Default unknown NTP protocol version. * `3` - NTP protocol version is 3. * `4` - NTP protocol version is 4. * `auto` - NTP protocol version is auto.
		Version *string                           `json:"Version,omitempty"`
		Array   *StorageNetAppClusterRelationship `json:"Array,omitempty"`
	}

	varStorageNetAppNtpServerWithoutEmbeddedStruct := StorageNetAppNtpServerWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageNetAppNtpServerWithoutEmbeddedStruct)
	if err == nil {
		varStorageNetAppNtpServer := _StorageNetAppNtpServer{}
		varStorageNetAppNtpServer.ClassId = varStorageNetAppNtpServerWithoutEmbeddedStruct.ClassId
		varStorageNetAppNtpServer.ObjectType = varStorageNetAppNtpServerWithoutEmbeddedStruct.ObjectType
		varStorageNetAppNtpServer.AuthenticationEnabled = varStorageNetAppNtpServerWithoutEmbeddedStruct.AuthenticationEnabled
		varStorageNetAppNtpServer.AuthenticationKeyId = varStorageNetAppNtpServerWithoutEmbeddedStruct.AuthenticationKeyId
		varStorageNetAppNtpServer.ClusterUuid = varStorageNetAppNtpServerWithoutEmbeddedStruct.ClusterUuid
		varStorageNetAppNtpServer.IsAuthenticationEnabled = varStorageNetAppNtpServerWithoutEmbeddedStruct.IsAuthenticationEnabled
		varStorageNetAppNtpServer.Server = varStorageNetAppNtpServerWithoutEmbeddedStruct.Server
		varStorageNetAppNtpServer.Version = varStorageNetAppNtpServerWithoutEmbeddedStruct.Version
		varStorageNetAppNtpServer.Array = varStorageNetAppNtpServerWithoutEmbeddedStruct.Array
		*o = StorageNetAppNtpServer(varStorageNetAppNtpServer)
	} else {
		return err
	}

	varStorageNetAppNtpServer := _StorageNetAppNtpServer{}

	err = json.Unmarshal(bytes, &varStorageNetAppNtpServer)
	if err == nil {
		o.MoBaseMo = varStorageNetAppNtpServer.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AuthenticationEnabled")
		delete(additionalProperties, "AuthenticationKeyId")
		delete(additionalProperties, "ClusterUuid")
		delete(additionalProperties, "IsAuthenticationEnabled")
		delete(additionalProperties, "Server")
		delete(additionalProperties, "Version")
		delete(additionalProperties, "Array")

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

type NullableStorageNetAppNtpServer struct {
	value *StorageNetAppNtpServer
	isSet bool
}

func (v NullableStorageNetAppNtpServer) Get() *StorageNetAppNtpServer {
	return v.value
}

func (v *NullableStorageNetAppNtpServer) Set(val *StorageNetAppNtpServer) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppNtpServer) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppNtpServer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppNtpServer(val *StorageNetAppNtpServer) *NullableStorageNetAppNtpServer {
	return &NullableStorageNetAppNtpServer{value: val, isSet: true}
}

func (v NullableStorageNetAppNtpServer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppNtpServer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
