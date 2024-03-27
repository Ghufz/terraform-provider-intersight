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

// OsBulkInstallInfo This MO models the CSV file content which the user uploaded for OS installation. As part of the handler, necessary filed in the model can be populated along with respective validation.
type OsBulkInstallInfo struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The content of the entire CSV file is stored as value. The content can hold complete OS install parameters in two sections. The first section holds generic information about the OS Install like OS Image, SCU Image etc. The second section holds parameters which are specific to each server level data.
	FileContent  *string                `json:"FileContent,omitempty"`
	GlobalConfig NullableOsGlobalConfig `json:"GlobalConfig,omitempty"`
	// Indicates whether the value of the 'fileContent' property has been set.
	IsFileContentSet *bool `json:"IsFileContentSet,omitempty"`
	// The name of the CSV file, which holds the OS install parameters.
	Name *string `json:"Name,omitempty"`
	// Denotes if the operating is pending, in_progress, completed_ok, completed_error. * `Pending` - The initial value of the OperStatus. * `InProgress` - The OperStatus value will be InProgress during execution. * `CompletedOk` - The API is successful with operation then OperStatus will be marked as CompletedOk. * `CompletedError` - The API is failed with operation then OperStatus will be marked as CompletedError. * `CompletedWarning` - The API is completed with some warning then OperStatus will be CompletedWarning.
	OperState         *string                                                      `json:"OperState,omitempty"`
	ServerConfigs     []OsServerConfig                                             `json:"ServerConfigs,omitempty"`
	ValidationInfos   []OsValidationInformation                                    `json:"ValidationInfos,omitempty"`
	ConfigurationFile *OsConfigurationFileRelationship                             `json:"ConfigurationFile,omitempty"`
	Organization      *OrganizationOrganizationRelationship                        `json:"Organization,omitempty"`
	OsImage           *SoftwarerepositoryOperatingSystemFileRelationship           `json:"OsImage,omitempty"`
	ScuImage          *FirmwareServerConfigurationUtilityDistributableRelationship `json:"ScuImage,omitempty"`
	// An array of relationships to computePhysical resources.
	Servers              []ComputePhysicalRelationship `json:"Servers,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OsBulkInstallInfo OsBulkInstallInfo

// NewOsBulkInstallInfo instantiates a new OsBulkInstallInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOsBulkInstallInfo(classId string, objectType string) *OsBulkInstallInfo {
	this := OsBulkInstallInfo{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewOsBulkInstallInfoWithDefaults instantiates a new OsBulkInstallInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOsBulkInstallInfoWithDefaults() *OsBulkInstallInfo {
	this := OsBulkInstallInfo{}
	var classId string = "os.BulkInstallInfo"
	this.ClassId = classId
	var objectType string = "os.BulkInstallInfo"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *OsBulkInstallInfo) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *OsBulkInstallInfo) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *OsBulkInstallInfo) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *OsBulkInstallInfo) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *OsBulkInstallInfo) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *OsBulkInstallInfo) SetObjectType(v string) {
	o.ObjectType = v
}

// GetFileContent returns the FileContent field value if set, zero value otherwise.
func (o *OsBulkInstallInfo) GetFileContent() string {
	if o == nil || o.FileContent == nil {
		var ret string
		return ret
	}
	return *o.FileContent
}

// GetFileContentOk returns a tuple with the FileContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsBulkInstallInfo) GetFileContentOk() (*string, bool) {
	if o == nil || o.FileContent == nil {
		return nil, false
	}
	return o.FileContent, true
}

// HasFileContent returns a boolean if a field has been set.
func (o *OsBulkInstallInfo) HasFileContent() bool {
	if o != nil && o.FileContent != nil {
		return true
	}

	return false
}

// SetFileContent gets a reference to the given string and assigns it to the FileContent field.
func (o *OsBulkInstallInfo) SetFileContent(v string) {
	o.FileContent = &v
}

// GetGlobalConfig returns the GlobalConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OsBulkInstallInfo) GetGlobalConfig() OsGlobalConfig {
	if o == nil || o.GlobalConfig.Get() == nil {
		var ret OsGlobalConfig
		return ret
	}
	return *o.GlobalConfig.Get()
}

// GetGlobalConfigOk returns a tuple with the GlobalConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OsBulkInstallInfo) GetGlobalConfigOk() (*OsGlobalConfig, bool) {
	if o == nil {
		return nil, false
	}
	return o.GlobalConfig.Get(), o.GlobalConfig.IsSet()
}

// HasGlobalConfig returns a boolean if a field has been set.
func (o *OsBulkInstallInfo) HasGlobalConfig() bool {
	if o != nil && o.GlobalConfig.IsSet() {
		return true
	}

	return false
}

// SetGlobalConfig gets a reference to the given NullableOsGlobalConfig and assigns it to the GlobalConfig field.
func (o *OsBulkInstallInfo) SetGlobalConfig(v OsGlobalConfig) {
	o.GlobalConfig.Set(&v)
}

// SetGlobalConfigNil sets the value for GlobalConfig to be an explicit nil
func (o *OsBulkInstallInfo) SetGlobalConfigNil() {
	o.GlobalConfig.Set(nil)
}

// UnsetGlobalConfig ensures that no value is present for GlobalConfig, not even an explicit nil
func (o *OsBulkInstallInfo) UnsetGlobalConfig() {
	o.GlobalConfig.Unset()
}

// GetIsFileContentSet returns the IsFileContentSet field value if set, zero value otherwise.
func (o *OsBulkInstallInfo) GetIsFileContentSet() bool {
	if o == nil || o.IsFileContentSet == nil {
		var ret bool
		return ret
	}
	return *o.IsFileContentSet
}

// GetIsFileContentSetOk returns a tuple with the IsFileContentSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsBulkInstallInfo) GetIsFileContentSetOk() (*bool, bool) {
	if o == nil || o.IsFileContentSet == nil {
		return nil, false
	}
	return o.IsFileContentSet, true
}

// HasIsFileContentSet returns a boolean if a field has been set.
func (o *OsBulkInstallInfo) HasIsFileContentSet() bool {
	if o != nil && o.IsFileContentSet != nil {
		return true
	}

	return false
}

// SetIsFileContentSet gets a reference to the given bool and assigns it to the IsFileContentSet field.
func (o *OsBulkInstallInfo) SetIsFileContentSet(v bool) {
	o.IsFileContentSet = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OsBulkInstallInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsBulkInstallInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OsBulkInstallInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OsBulkInstallInfo) SetName(v string) {
	o.Name = &v
}

// GetOperState returns the OperState field value if set, zero value otherwise.
func (o *OsBulkInstallInfo) GetOperState() string {
	if o == nil || o.OperState == nil {
		var ret string
		return ret
	}
	return *o.OperState
}

// GetOperStateOk returns a tuple with the OperState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsBulkInstallInfo) GetOperStateOk() (*string, bool) {
	if o == nil || o.OperState == nil {
		return nil, false
	}
	return o.OperState, true
}

// HasOperState returns a boolean if a field has been set.
func (o *OsBulkInstallInfo) HasOperState() bool {
	if o != nil && o.OperState != nil {
		return true
	}

	return false
}

// SetOperState gets a reference to the given string and assigns it to the OperState field.
func (o *OsBulkInstallInfo) SetOperState(v string) {
	o.OperState = &v
}

// GetServerConfigs returns the ServerConfigs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OsBulkInstallInfo) GetServerConfigs() []OsServerConfig {
	if o == nil {
		var ret []OsServerConfig
		return ret
	}
	return o.ServerConfigs
}

// GetServerConfigsOk returns a tuple with the ServerConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OsBulkInstallInfo) GetServerConfigsOk() ([]OsServerConfig, bool) {
	if o == nil || o.ServerConfigs == nil {
		return nil, false
	}
	return o.ServerConfigs, true
}

// HasServerConfigs returns a boolean if a field has been set.
func (o *OsBulkInstallInfo) HasServerConfigs() bool {
	if o != nil && o.ServerConfigs != nil {
		return true
	}

	return false
}

// SetServerConfigs gets a reference to the given []OsServerConfig and assigns it to the ServerConfigs field.
func (o *OsBulkInstallInfo) SetServerConfigs(v []OsServerConfig) {
	o.ServerConfigs = v
}

// GetValidationInfos returns the ValidationInfos field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OsBulkInstallInfo) GetValidationInfos() []OsValidationInformation {
	if o == nil {
		var ret []OsValidationInformation
		return ret
	}
	return o.ValidationInfos
}

// GetValidationInfosOk returns a tuple with the ValidationInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OsBulkInstallInfo) GetValidationInfosOk() ([]OsValidationInformation, bool) {
	if o == nil || o.ValidationInfos == nil {
		return nil, false
	}
	return o.ValidationInfos, true
}

// HasValidationInfos returns a boolean if a field has been set.
func (o *OsBulkInstallInfo) HasValidationInfos() bool {
	if o != nil && o.ValidationInfos != nil {
		return true
	}

	return false
}

// SetValidationInfos gets a reference to the given []OsValidationInformation and assigns it to the ValidationInfos field.
func (o *OsBulkInstallInfo) SetValidationInfos(v []OsValidationInformation) {
	o.ValidationInfos = v
}

// GetConfigurationFile returns the ConfigurationFile field value if set, zero value otherwise.
func (o *OsBulkInstallInfo) GetConfigurationFile() OsConfigurationFileRelationship {
	if o == nil || o.ConfigurationFile == nil {
		var ret OsConfigurationFileRelationship
		return ret
	}
	return *o.ConfigurationFile
}

// GetConfigurationFileOk returns a tuple with the ConfigurationFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsBulkInstallInfo) GetConfigurationFileOk() (*OsConfigurationFileRelationship, bool) {
	if o == nil || o.ConfigurationFile == nil {
		return nil, false
	}
	return o.ConfigurationFile, true
}

// HasConfigurationFile returns a boolean if a field has been set.
func (o *OsBulkInstallInfo) HasConfigurationFile() bool {
	if o != nil && o.ConfigurationFile != nil {
		return true
	}

	return false
}

// SetConfigurationFile gets a reference to the given OsConfigurationFileRelationship and assigns it to the ConfigurationFile field.
func (o *OsBulkInstallInfo) SetConfigurationFile(v OsConfigurationFileRelationship) {
	o.ConfigurationFile = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *OsBulkInstallInfo) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsBulkInstallInfo) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *OsBulkInstallInfo) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *OsBulkInstallInfo) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetOsImage returns the OsImage field value if set, zero value otherwise.
func (o *OsBulkInstallInfo) GetOsImage() SoftwarerepositoryOperatingSystemFileRelationship {
	if o == nil || o.OsImage == nil {
		var ret SoftwarerepositoryOperatingSystemFileRelationship
		return ret
	}
	return *o.OsImage
}

// GetOsImageOk returns a tuple with the OsImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsBulkInstallInfo) GetOsImageOk() (*SoftwarerepositoryOperatingSystemFileRelationship, bool) {
	if o == nil || o.OsImage == nil {
		return nil, false
	}
	return o.OsImage, true
}

// HasOsImage returns a boolean if a field has been set.
func (o *OsBulkInstallInfo) HasOsImage() bool {
	if o != nil && o.OsImage != nil {
		return true
	}

	return false
}

// SetOsImage gets a reference to the given SoftwarerepositoryOperatingSystemFileRelationship and assigns it to the OsImage field.
func (o *OsBulkInstallInfo) SetOsImage(v SoftwarerepositoryOperatingSystemFileRelationship) {
	o.OsImage = &v
}

// GetScuImage returns the ScuImage field value if set, zero value otherwise.
func (o *OsBulkInstallInfo) GetScuImage() FirmwareServerConfigurationUtilityDistributableRelationship {
	if o == nil || o.ScuImage == nil {
		var ret FirmwareServerConfigurationUtilityDistributableRelationship
		return ret
	}
	return *o.ScuImage
}

// GetScuImageOk returns a tuple with the ScuImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsBulkInstallInfo) GetScuImageOk() (*FirmwareServerConfigurationUtilityDistributableRelationship, bool) {
	if o == nil || o.ScuImage == nil {
		return nil, false
	}
	return o.ScuImage, true
}

// HasScuImage returns a boolean if a field has been set.
func (o *OsBulkInstallInfo) HasScuImage() bool {
	if o != nil && o.ScuImage != nil {
		return true
	}

	return false
}

// SetScuImage gets a reference to the given FirmwareServerConfigurationUtilityDistributableRelationship and assigns it to the ScuImage field.
func (o *OsBulkInstallInfo) SetScuImage(v FirmwareServerConfigurationUtilityDistributableRelationship) {
	o.ScuImage = &v
}

// GetServers returns the Servers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OsBulkInstallInfo) GetServers() []ComputePhysicalRelationship {
	if o == nil {
		var ret []ComputePhysicalRelationship
		return ret
	}
	return o.Servers
}

// GetServersOk returns a tuple with the Servers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OsBulkInstallInfo) GetServersOk() ([]ComputePhysicalRelationship, bool) {
	if o == nil || o.Servers == nil {
		return nil, false
	}
	return o.Servers, true
}

// HasServers returns a boolean if a field has been set.
func (o *OsBulkInstallInfo) HasServers() bool {
	if o != nil && o.Servers != nil {
		return true
	}

	return false
}

// SetServers gets a reference to the given []ComputePhysicalRelationship and assigns it to the Servers field.
func (o *OsBulkInstallInfo) SetServers(v []ComputePhysicalRelationship) {
	o.Servers = v
}

func (o OsBulkInstallInfo) MarshalJSON() ([]byte, error) {
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
	if o.FileContent != nil {
		toSerialize["FileContent"] = o.FileContent
	}
	if o.GlobalConfig.IsSet() {
		toSerialize["GlobalConfig"] = o.GlobalConfig.Get()
	}
	if o.IsFileContentSet != nil {
		toSerialize["IsFileContentSet"] = o.IsFileContentSet
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.OperState != nil {
		toSerialize["OperState"] = o.OperState
	}
	if o.ServerConfigs != nil {
		toSerialize["ServerConfigs"] = o.ServerConfigs
	}
	if o.ValidationInfos != nil {
		toSerialize["ValidationInfos"] = o.ValidationInfos
	}
	if o.ConfigurationFile != nil {
		toSerialize["ConfigurationFile"] = o.ConfigurationFile
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.OsImage != nil {
		toSerialize["OsImage"] = o.OsImage
	}
	if o.ScuImage != nil {
		toSerialize["ScuImage"] = o.ScuImage
	}
	if o.Servers != nil {
		toSerialize["Servers"] = o.Servers
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OsBulkInstallInfo) UnmarshalJSON(bytes []byte) (err error) {
	type OsBulkInstallInfoWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The content of the entire CSV file is stored as value. The content can hold complete OS install parameters in two sections. The first section holds generic information about the OS Install like OS Image, SCU Image etc. The second section holds parameters which are specific to each server level data.
		FileContent  *string                `json:"FileContent,omitempty"`
		GlobalConfig NullableOsGlobalConfig `json:"GlobalConfig,omitempty"`
		// Indicates whether the value of the 'fileContent' property has been set.
		IsFileContentSet *bool `json:"IsFileContentSet,omitempty"`
		// The name of the CSV file, which holds the OS install parameters.
		Name *string `json:"Name,omitempty"`
		// Denotes if the operating is pending, in_progress, completed_ok, completed_error. * `Pending` - The initial value of the OperStatus. * `InProgress` - The OperStatus value will be InProgress during execution. * `CompletedOk` - The API is successful with operation then OperStatus will be marked as CompletedOk. * `CompletedError` - The API is failed with operation then OperStatus will be marked as CompletedError. * `CompletedWarning` - The API is completed with some warning then OperStatus will be CompletedWarning.
		OperState         *string                                                      `json:"OperState,omitempty"`
		ServerConfigs     []OsServerConfig                                             `json:"ServerConfigs,omitempty"`
		ValidationInfos   []OsValidationInformation                                    `json:"ValidationInfos,omitempty"`
		ConfigurationFile *OsConfigurationFileRelationship                             `json:"ConfigurationFile,omitempty"`
		Organization      *OrganizationOrganizationRelationship                        `json:"Organization,omitempty"`
		OsImage           *SoftwarerepositoryOperatingSystemFileRelationship           `json:"OsImage,omitempty"`
		ScuImage          *FirmwareServerConfigurationUtilityDistributableRelationship `json:"ScuImage,omitempty"`
		// An array of relationships to computePhysical resources.
		Servers []ComputePhysicalRelationship `json:"Servers,omitempty"`
	}

	varOsBulkInstallInfoWithoutEmbeddedStruct := OsBulkInstallInfoWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varOsBulkInstallInfoWithoutEmbeddedStruct)
	if err == nil {
		varOsBulkInstallInfo := _OsBulkInstallInfo{}
		varOsBulkInstallInfo.ClassId = varOsBulkInstallInfoWithoutEmbeddedStruct.ClassId
		varOsBulkInstallInfo.ObjectType = varOsBulkInstallInfoWithoutEmbeddedStruct.ObjectType
		varOsBulkInstallInfo.FileContent = varOsBulkInstallInfoWithoutEmbeddedStruct.FileContent
		varOsBulkInstallInfo.GlobalConfig = varOsBulkInstallInfoWithoutEmbeddedStruct.GlobalConfig
		varOsBulkInstallInfo.IsFileContentSet = varOsBulkInstallInfoWithoutEmbeddedStruct.IsFileContentSet
		varOsBulkInstallInfo.Name = varOsBulkInstallInfoWithoutEmbeddedStruct.Name
		varOsBulkInstallInfo.OperState = varOsBulkInstallInfoWithoutEmbeddedStruct.OperState
		varOsBulkInstallInfo.ServerConfigs = varOsBulkInstallInfoWithoutEmbeddedStruct.ServerConfigs
		varOsBulkInstallInfo.ValidationInfos = varOsBulkInstallInfoWithoutEmbeddedStruct.ValidationInfos
		varOsBulkInstallInfo.ConfigurationFile = varOsBulkInstallInfoWithoutEmbeddedStruct.ConfigurationFile
		varOsBulkInstallInfo.Organization = varOsBulkInstallInfoWithoutEmbeddedStruct.Organization
		varOsBulkInstallInfo.OsImage = varOsBulkInstallInfoWithoutEmbeddedStruct.OsImage
		varOsBulkInstallInfo.ScuImage = varOsBulkInstallInfoWithoutEmbeddedStruct.ScuImage
		varOsBulkInstallInfo.Servers = varOsBulkInstallInfoWithoutEmbeddedStruct.Servers
		*o = OsBulkInstallInfo(varOsBulkInstallInfo)
	} else {
		return err
	}

	varOsBulkInstallInfo := _OsBulkInstallInfo{}

	err = json.Unmarshal(bytes, &varOsBulkInstallInfo)
	if err == nil {
		o.MoBaseMo = varOsBulkInstallInfo.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "FileContent")
		delete(additionalProperties, "GlobalConfig")
		delete(additionalProperties, "IsFileContentSet")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "OperState")
		delete(additionalProperties, "ServerConfigs")
		delete(additionalProperties, "ValidationInfos")
		delete(additionalProperties, "ConfigurationFile")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "OsImage")
		delete(additionalProperties, "ScuImage")
		delete(additionalProperties, "Servers")

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

type NullableOsBulkInstallInfo struct {
	value *OsBulkInstallInfo
	isSet bool
}

func (v NullableOsBulkInstallInfo) Get() *OsBulkInstallInfo {
	return v.value
}

func (v *NullableOsBulkInstallInfo) Set(val *OsBulkInstallInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableOsBulkInstallInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableOsBulkInstallInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOsBulkInstallInfo(val *OsBulkInstallInfo) *NullableOsBulkInstallInfo {
	return &NullableOsBulkInstallInfo{value: val, isSet: true}
}

func (v NullableOsBulkInstallInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOsBulkInstallInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
