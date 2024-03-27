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

// ApplianceRemoteFileImport The properties of appliance.RemoteFileImport are used to create an scp or sftp request to import a firmware image.
type ApplianceRemoteFileImport struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The name of the file to be imported.
	Filename *string `json:"Filename,omitempty"`
	// The hostname of the machine where the file is located.
	Hostname *string `json:"Hostname,omitempty"`
	// Indicates whether the value of the 'password' property has been set.
	IsPasswordSet *bool `json:"IsPasswordSet,omitempty"`
	// Password for remote requiest.
	Password *string `json:"Password,omitempty"`
	// The port that should be used for the remote request.
	Path *string `json:"Path,omitempty"`
	// The port that should be used for the remote request.
	Port *int64 `json:"Port,omitempty"`
	// Specifies if this is an scp or sftp request. * `scp` - Secure Copy Protocol (SCP) to access the file server. * `sftp` - SSH File Transfer Protocol (SFTP) to access file server. * `cifs` - Common Internet File System (CIFS) Protocol to access file server.
	Protocol *string `json:"Protocol,omitempty"`
	// The username for the remote request.
	Username             *string                 `json:"Username,omitempty"`
	Account              *IamAccountRelationship `json:"Account,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplianceRemoteFileImport ApplianceRemoteFileImport

// NewApplianceRemoteFileImport instantiates a new ApplianceRemoteFileImport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceRemoteFileImport(classId string, objectType string) *ApplianceRemoteFileImport {
	this := ApplianceRemoteFileImport{}
	this.ClassId = classId
	this.ObjectType = objectType
	var protocol string = "scp"
	this.Protocol = &protocol
	return &this
}

// NewApplianceRemoteFileImportWithDefaults instantiates a new ApplianceRemoteFileImport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceRemoteFileImportWithDefaults() *ApplianceRemoteFileImport {
	this := ApplianceRemoteFileImport{}
	var classId string = "appliance.RemoteFileImport"
	this.ClassId = classId
	var objectType string = "appliance.RemoteFileImport"
	this.ObjectType = objectType
	var protocol string = "scp"
	this.Protocol = &protocol
	return &this
}

// GetClassId returns the ClassId field value
func (o *ApplianceRemoteFileImport) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ApplianceRemoteFileImport) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ApplianceRemoteFileImport) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ApplianceRemoteFileImport) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ApplianceRemoteFileImport) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ApplianceRemoteFileImport) SetObjectType(v string) {
	o.ObjectType = v
}

// GetFilename returns the Filename field value if set, zero value otherwise.
func (o *ApplianceRemoteFileImport) GetFilename() string {
	if o == nil || o.Filename == nil {
		var ret string
		return ret
	}
	return *o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceRemoteFileImport) GetFilenameOk() (*string, bool) {
	if o == nil || o.Filename == nil {
		return nil, false
	}
	return o.Filename, true
}

// HasFilename returns a boolean if a field has been set.
func (o *ApplianceRemoteFileImport) HasFilename() bool {
	if o != nil && o.Filename != nil {
		return true
	}

	return false
}

// SetFilename gets a reference to the given string and assigns it to the Filename field.
func (o *ApplianceRemoteFileImport) SetFilename(v string) {
	o.Filename = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *ApplianceRemoteFileImport) GetHostname() string {
	if o == nil || o.Hostname == nil {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceRemoteFileImport) GetHostnameOk() (*string, bool) {
	if o == nil || o.Hostname == nil {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *ApplianceRemoteFileImport) HasHostname() bool {
	if o != nil && o.Hostname != nil {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *ApplianceRemoteFileImport) SetHostname(v string) {
	o.Hostname = &v
}

// GetIsPasswordSet returns the IsPasswordSet field value if set, zero value otherwise.
func (o *ApplianceRemoteFileImport) GetIsPasswordSet() bool {
	if o == nil || o.IsPasswordSet == nil {
		var ret bool
		return ret
	}
	return *o.IsPasswordSet
}

// GetIsPasswordSetOk returns a tuple with the IsPasswordSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceRemoteFileImport) GetIsPasswordSetOk() (*bool, bool) {
	if o == nil || o.IsPasswordSet == nil {
		return nil, false
	}
	return o.IsPasswordSet, true
}

// HasIsPasswordSet returns a boolean if a field has been set.
func (o *ApplianceRemoteFileImport) HasIsPasswordSet() bool {
	if o != nil && o.IsPasswordSet != nil {
		return true
	}

	return false
}

// SetIsPasswordSet gets a reference to the given bool and assigns it to the IsPasswordSet field.
func (o *ApplianceRemoteFileImport) SetIsPasswordSet(v bool) {
	o.IsPasswordSet = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *ApplianceRemoteFileImport) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceRemoteFileImport) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *ApplianceRemoteFileImport) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *ApplianceRemoteFileImport) SetPassword(v string) {
	o.Password = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *ApplianceRemoteFileImport) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceRemoteFileImport) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *ApplianceRemoteFileImport) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *ApplianceRemoteFileImport) SetPath(v string) {
	o.Path = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *ApplianceRemoteFileImport) GetPort() int64 {
	if o == nil || o.Port == nil {
		var ret int64
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceRemoteFileImport) GetPortOk() (*int64, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *ApplianceRemoteFileImport) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int64 and assigns it to the Port field.
func (o *ApplianceRemoteFileImport) SetPort(v int64) {
	o.Port = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *ApplianceRemoteFileImport) GetProtocol() string {
	if o == nil || o.Protocol == nil {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceRemoteFileImport) GetProtocolOk() (*string, bool) {
	if o == nil || o.Protocol == nil {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *ApplianceRemoteFileImport) HasProtocol() bool {
	if o != nil && o.Protocol != nil {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *ApplianceRemoteFileImport) SetProtocol(v string) {
	o.Protocol = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *ApplianceRemoteFileImport) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceRemoteFileImport) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *ApplianceRemoteFileImport) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *ApplianceRemoteFileImport) SetUsername(v string) {
	o.Username = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *ApplianceRemoteFileImport) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceRemoteFileImport) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *ApplianceRemoteFileImport) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *ApplianceRemoteFileImport) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

func (o ApplianceRemoteFileImport) MarshalJSON() ([]byte, error) {
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
	if o.Filename != nil {
		toSerialize["Filename"] = o.Filename
	}
	if o.Hostname != nil {
		toSerialize["Hostname"] = o.Hostname
	}
	if o.IsPasswordSet != nil {
		toSerialize["IsPasswordSet"] = o.IsPasswordSet
	}
	if o.Password != nil {
		toSerialize["Password"] = o.Password
	}
	if o.Path != nil {
		toSerialize["Path"] = o.Path
	}
	if o.Port != nil {
		toSerialize["Port"] = o.Port
	}
	if o.Protocol != nil {
		toSerialize["Protocol"] = o.Protocol
	}
	if o.Username != nil {
		toSerialize["Username"] = o.Username
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ApplianceRemoteFileImport) UnmarshalJSON(bytes []byte) (err error) {
	type ApplianceRemoteFileImportWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The name of the file to be imported.
		Filename *string `json:"Filename,omitempty"`
		// The hostname of the machine where the file is located.
		Hostname *string `json:"Hostname,omitempty"`
		// Indicates whether the value of the 'password' property has been set.
		IsPasswordSet *bool `json:"IsPasswordSet,omitempty"`
		// Password for remote requiest.
		Password *string `json:"Password,omitempty"`
		// The port that should be used for the remote request.
		Path *string `json:"Path,omitempty"`
		// The port that should be used for the remote request.
		Port *int64 `json:"Port,omitempty"`
		// Specifies if this is an scp or sftp request. * `scp` - Secure Copy Protocol (SCP) to access the file server. * `sftp` - SSH File Transfer Protocol (SFTP) to access file server. * `cifs` - Common Internet File System (CIFS) Protocol to access file server.
		Protocol *string `json:"Protocol,omitempty"`
		// The username for the remote request.
		Username *string                 `json:"Username,omitempty"`
		Account  *IamAccountRelationship `json:"Account,omitempty"`
	}

	varApplianceRemoteFileImportWithoutEmbeddedStruct := ApplianceRemoteFileImportWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varApplianceRemoteFileImportWithoutEmbeddedStruct)
	if err == nil {
		varApplianceRemoteFileImport := _ApplianceRemoteFileImport{}
		varApplianceRemoteFileImport.ClassId = varApplianceRemoteFileImportWithoutEmbeddedStruct.ClassId
		varApplianceRemoteFileImport.ObjectType = varApplianceRemoteFileImportWithoutEmbeddedStruct.ObjectType
		varApplianceRemoteFileImport.Filename = varApplianceRemoteFileImportWithoutEmbeddedStruct.Filename
		varApplianceRemoteFileImport.Hostname = varApplianceRemoteFileImportWithoutEmbeddedStruct.Hostname
		varApplianceRemoteFileImport.IsPasswordSet = varApplianceRemoteFileImportWithoutEmbeddedStruct.IsPasswordSet
		varApplianceRemoteFileImport.Password = varApplianceRemoteFileImportWithoutEmbeddedStruct.Password
		varApplianceRemoteFileImport.Path = varApplianceRemoteFileImportWithoutEmbeddedStruct.Path
		varApplianceRemoteFileImport.Port = varApplianceRemoteFileImportWithoutEmbeddedStruct.Port
		varApplianceRemoteFileImport.Protocol = varApplianceRemoteFileImportWithoutEmbeddedStruct.Protocol
		varApplianceRemoteFileImport.Username = varApplianceRemoteFileImportWithoutEmbeddedStruct.Username
		varApplianceRemoteFileImport.Account = varApplianceRemoteFileImportWithoutEmbeddedStruct.Account
		*o = ApplianceRemoteFileImport(varApplianceRemoteFileImport)
	} else {
		return err
	}

	varApplianceRemoteFileImport := _ApplianceRemoteFileImport{}

	err = json.Unmarshal(bytes, &varApplianceRemoteFileImport)
	if err == nil {
		o.MoBaseMo = varApplianceRemoteFileImport.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Filename")
		delete(additionalProperties, "Hostname")
		delete(additionalProperties, "IsPasswordSet")
		delete(additionalProperties, "Password")
		delete(additionalProperties, "Path")
		delete(additionalProperties, "Port")
		delete(additionalProperties, "Protocol")
		delete(additionalProperties, "Username")
		delete(additionalProperties, "Account")

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

type NullableApplianceRemoteFileImport struct {
	value *ApplianceRemoteFileImport
	isSet bool
}

func (v NullableApplianceRemoteFileImport) Get() *ApplianceRemoteFileImport {
	return v.value
}

func (v *NullableApplianceRemoteFileImport) Set(val *ApplianceRemoteFileImport) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceRemoteFileImport) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceRemoteFileImport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceRemoteFileImport(val *ApplianceRemoteFileImport) *NullableApplianceRemoteFileImport {
	return &NullableApplianceRemoteFileImport{value: val, isSet: true}
}

func (v NullableApplianceRemoteFileImport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceRemoteFileImport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
