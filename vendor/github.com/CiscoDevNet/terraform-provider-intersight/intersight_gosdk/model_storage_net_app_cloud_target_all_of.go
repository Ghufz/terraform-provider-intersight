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

// StorageNetAppCloudTargetAllOf Definition of the list of properties defined in 'storage.NetAppCloudTarget', excluding properties defined in parent classes.
type StorageNetAppCloudTargetAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Access key ID for AWS_S3 and other S3 compatible provider types.
	AccessKey *string `json:"AccessKey,omitempty"`
	// Authentication used to access the target.
	AuthenticationType *string `json:"AuthenticationType,omitempty"`
	// Azure cloud account name.
	AzureAccount *string `json:"AzureAccount,omitempty"`
	// Specifies a full URL of the request to a CAP server for retrieving temporary credentials (access-key, secret-password, and session token) for accessing the object store.
	CapUrl *string `json:"CapUrl,omitempty"`
	// Is SSL/TLS certificate validation enabled? The default value is true. This can only be modified for SGWS, IBM_COS, and ONTAP_S3 provider types.
	CertificateValidationEnabled *bool    `json:"CertificateValidationEnabled,omitempty"`
	CloudStorage                 []string `json:"CloudStorage,omitempty"`
	// Data bucket/container name.
	Container *string `json:"Container,omitempty"`
	// IPspace to use in order to reach the cloud target.
	Ipspace *string `json:"Ipspace,omitempty"`
	// Name of the cloud target.
	Name *string `json:"Name,omitempty"`
	// Owner of the target. Allowed values are FabricPool or SnapMirror. * `FabricPool` - NetApp FabricPool Cloud Target owner. * `SnapMirror` - NetApp SnapMirror Cloud Target owner.
	Owner *string `json:"Owner,omitempty"`
	// Port number of the object store that ONTAP uses when establishing a connection.
	Port *int64 `json:"Port,omitempty"`
	// Type of cloud provider, e.g., AWS_S3 or ONTAP_S3. * `ONTAP_S3` - Cloud target provider type ONTAP_S3. * `AliCloud` - Cloud target provider type AliCloud. * `AWS_S3` - Cloud target provider type AWS S3. * `Azure_Cloud` - Cloud target provider type Azure_Cloud. * `GoogleCloud` - Cloud target provider type GoogleCloud. * `IBM_COS` - Cloud target provider type IBM_COS. * `SGWS` - Cloud target provider type SGWS.
	ProviderType *string `json:"ProviderType,omitempty"`
	// Fully qualified domain name of the object store server.
	Server *string `json:"Server,omitempty"`
	// Use of the cloud target by SnapMirror. * `data` - Data is stored in the SnapMirror target. * `metadata` - Metadata is stored in the SnapMirror target.
	SnapMirrorUse *string `json:"SnapMirrorUse,omitempty"`
	// SSL/HTTPS enabled or not.
	SslEnabled *bool `json:"SslEnabled,omitempty"`
	// The amount of cloud space used by all the aggregates attached to the target, in bytes.
	Used *int64 `json:"Used,omitempty"`
	// Uuid of the cloud target.
	Uuid                 *string                           `json:"Uuid,omitempty"`
	Array                *StorageNetAppClusterRelationship `json:"Array,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppCloudTargetAllOf StorageNetAppCloudTargetAllOf

// NewStorageNetAppCloudTargetAllOf instantiates a new StorageNetAppCloudTargetAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppCloudTargetAllOf(classId string, objectType string) *StorageNetAppCloudTargetAllOf {
	this := StorageNetAppCloudTargetAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppCloudTargetAllOfWithDefaults instantiates a new StorageNetAppCloudTargetAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppCloudTargetAllOfWithDefaults() *StorageNetAppCloudTargetAllOf {
	this := StorageNetAppCloudTargetAllOf{}
	var classId string = "storage.NetAppCloudTarget"
	this.ClassId = classId
	var objectType string = "storage.NetAppCloudTarget"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppCloudTargetAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppCloudTargetAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppCloudTargetAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppCloudTargetAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppCloudTargetAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppCloudTargetAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAccessKey returns the AccessKey field value if set, zero value otherwise.
func (o *StorageNetAppCloudTargetAllOf) GetAccessKey() string {
	if o == nil || o.AccessKey == nil {
		var ret string
		return ret
	}
	return *o.AccessKey
}

// GetAccessKeyOk returns a tuple with the AccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppCloudTargetAllOf) GetAccessKeyOk() (*string, bool) {
	if o == nil || o.AccessKey == nil {
		return nil, false
	}
	return o.AccessKey, true
}

// HasAccessKey returns a boolean if a field has been set.
func (o *StorageNetAppCloudTargetAllOf) HasAccessKey() bool {
	if o != nil && o.AccessKey != nil {
		return true
	}

	return false
}

// SetAccessKey gets a reference to the given string and assigns it to the AccessKey field.
func (o *StorageNetAppCloudTargetAllOf) SetAccessKey(v string) {
	o.AccessKey = &v
}

// GetAuthenticationType returns the AuthenticationType field value if set, zero value otherwise.
func (o *StorageNetAppCloudTargetAllOf) GetAuthenticationType() string {
	if o == nil || o.AuthenticationType == nil {
		var ret string
		return ret
	}
	return *o.AuthenticationType
}

// GetAuthenticationTypeOk returns a tuple with the AuthenticationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppCloudTargetAllOf) GetAuthenticationTypeOk() (*string, bool) {
	if o == nil || o.AuthenticationType == nil {
		return nil, false
	}
	return o.AuthenticationType, true
}

// HasAuthenticationType returns a boolean if a field has been set.
func (o *StorageNetAppCloudTargetAllOf) HasAuthenticationType() bool {
	if o != nil && o.AuthenticationType != nil {
		return true
	}

	return false
}

// SetAuthenticationType gets a reference to the given string and assigns it to the AuthenticationType field.
func (o *StorageNetAppCloudTargetAllOf) SetAuthenticationType(v string) {
	o.AuthenticationType = &v
}

// GetAzureAccount returns the AzureAccount field value if set, zero value otherwise.
func (o *StorageNetAppCloudTargetAllOf) GetAzureAccount() string {
	if o == nil || o.AzureAccount == nil {
		var ret string
		return ret
	}
	return *o.AzureAccount
}

// GetAzureAccountOk returns a tuple with the AzureAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppCloudTargetAllOf) GetAzureAccountOk() (*string, bool) {
	if o == nil || o.AzureAccount == nil {
		return nil, false
	}
	return o.AzureAccount, true
}

// HasAzureAccount returns a boolean if a field has been set.
func (o *StorageNetAppCloudTargetAllOf) HasAzureAccount() bool {
	if o != nil && o.AzureAccount != nil {
		return true
	}

	return false
}

// SetAzureAccount gets a reference to the given string and assigns it to the AzureAccount field.
func (o *StorageNetAppCloudTargetAllOf) SetAzureAccount(v string) {
	o.AzureAccount = &v
}

// GetCapUrl returns the CapUrl field value if set, zero value otherwise.
func (o *StorageNetAppCloudTargetAllOf) GetCapUrl() string {
	if o == nil || o.CapUrl == nil {
		var ret string
		return ret
	}
	return *o.CapUrl
}

// GetCapUrlOk returns a tuple with the CapUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppCloudTargetAllOf) GetCapUrlOk() (*string, bool) {
	if o == nil || o.CapUrl == nil {
		return nil, false
	}
	return o.CapUrl, true
}

// HasCapUrl returns a boolean if a field has been set.
func (o *StorageNetAppCloudTargetAllOf) HasCapUrl() bool {
	if o != nil && o.CapUrl != nil {
		return true
	}

	return false
}

// SetCapUrl gets a reference to the given string and assigns it to the CapUrl field.
func (o *StorageNetAppCloudTargetAllOf) SetCapUrl(v string) {
	o.CapUrl = &v
}

// GetCertificateValidationEnabled returns the CertificateValidationEnabled field value if set, zero value otherwise.
func (o *StorageNetAppCloudTargetAllOf) GetCertificateValidationEnabled() bool {
	if o == nil || o.CertificateValidationEnabled == nil {
		var ret bool
		return ret
	}
	return *o.CertificateValidationEnabled
}

// GetCertificateValidationEnabledOk returns a tuple with the CertificateValidationEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppCloudTargetAllOf) GetCertificateValidationEnabledOk() (*bool, bool) {
	if o == nil || o.CertificateValidationEnabled == nil {
		return nil, false
	}
	return o.CertificateValidationEnabled, true
}

// HasCertificateValidationEnabled returns a boolean if a field has been set.
func (o *StorageNetAppCloudTargetAllOf) HasCertificateValidationEnabled() bool {
	if o != nil && o.CertificateValidationEnabled != nil {
		return true
	}

	return false
}

// SetCertificateValidationEnabled gets a reference to the given bool and assigns it to the CertificateValidationEnabled field.
func (o *StorageNetAppCloudTargetAllOf) SetCertificateValidationEnabled(v bool) {
	o.CertificateValidationEnabled = &v
}

// GetCloudStorage returns the CloudStorage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppCloudTargetAllOf) GetCloudStorage() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.CloudStorage
}

// GetCloudStorageOk returns a tuple with the CloudStorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppCloudTargetAllOf) GetCloudStorageOk() ([]string, bool) {
	if o == nil || o.CloudStorage == nil {
		return nil, false
	}
	return o.CloudStorage, true
}

// HasCloudStorage returns a boolean if a field has been set.
func (o *StorageNetAppCloudTargetAllOf) HasCloudStorage() bool {
	if o != nil && o.CloudStorage != nil {
		return true
	}

	return false
}

// SetCloudStorage gets a reference to the given []string and assigns it to the CloudStorage field.
func (o *StorageNetAppCloudTargetAllOf) SetCloudStorage(v []string) {
	o.CloudStorage = v
}

// GetContainer returns the Container field value if set, zero value otherwise.
func (o *StorageNetAppCloudTargetAllOf) GetContainer() string {
	if o == nil || o.Container == nil {
		var ret string
		return ret
	}
	return *o.Container
}

// GetContainerOk returns a tuple with the Container field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppCloudTargetAllOf) GetContainerOk() (*string, bool) {
	if o == nil || o.Container == nil {
		return nil, false
	}
	return o.Container, true
}

// HasContainer returns a boolean if a field has been set.
func (o *StorageNetAppCloudTargetAllOf) HasContainer() bool {
	if o != nil && o.Container != nil {
		return true
	}

	return false
}

// SetContainer gets a reference to the given string and assigns it to the Container field.
func (o *StorageNetAppCloudTargetAllOf) SetContainer(v string) {
	o.Container = &v
}

// GetIpspace returns the Ipspace field value if set, zero value otherwise.
func (o *StorageNetAppCloudTargetAllOf) GetIpspace() string {
	if o == nil || o.Ipspace == nil {
		var ret string
		return ret
	}
	return *o.Ipspace
}

// GetIpspaceOk returns a tuple with the Ipspace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppCloudTargetAllOf) GetIpspaceOk() (*string, bool) {
	if o == nil || o.Ipspace == nil {
		return nil, false
	}
	return o.Ipspace, true
}

// HasIpspace returns a boolean if a field has been set.
func (o *StorageNetAppCloudTargetAllOf) HasIpspace() bool {
	if o != nil && o.Ipspace != nil {
		return true
	}

	return false
}

// SetIpspace gets a reference to the given string and assigns it to the Ipspace field.
func (o *StorageNetAppCloudTargetAllOf) SetIpspace(v string) {
	o.Ipspace = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StorageNetAppCloudTargetAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppCloudTargetAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StorageNetAppCloudTargetAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StorageNetAppCloudTargetAllOf) SetName(v string) {
	o.Name = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *StorageNetAppCloudTargetAllOf) GetOwner() string {
	if o == nil || o.Owner == nil {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppCloudTargetAllOf) GetOwnerOk() (*string, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *StorageNetAppCloudTargetAllOf) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *StorageNetAppCloudTargetAllOf) SetOwner(v string) {
	o.Owner = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *StorageNetAppCloudTargetAllOf) GetPort() int64 {
	if o == nil || o.Port == nil {
		var ret int64
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppCloudTargetAllOf) GetPortOk() (*int64, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *StorageNetAppCloudTargetAllOf) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int64 and assigns it to the Port field.
func (o *StorageNetAppCloudTargetAllOf) SetPort(v int64) {
	o.Port = &v
}

// GetProviderType returns the ProviderType field value if set, zero value otherwise.
func (o *StorageNetAppCloudTargetAllOf) GetProviderType() string {
	if o == nil || o.ProviderType == nil {
		var ret string
		return ret
	}
	return *o.ProviderType
}

// GetProviderTypeOk returns a tuple with the ProviderType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppCloudTargetAllOf) GetProviderTypeOk() (*string, bool) {
	if o == nil || o.ProviderType == nil {
		return nil, false
	}
	return o.ProviderType, true
}

// HasProviderType returns a boolean if a field has been set.
func (o *StorageNetAppCloudTargetAllOf) HasProviderType() bool {
	if o != nil && o.ProviderType != nil {
		return true
	}

	return false
}

// SetProviderType gets a reference to the given string and assigns it to the ProviderType field.
func (o *StorageNetAppCloudTargetAllOf) SetProviderType(v string) {
	o.ProviderType = &v
}

// GetServer returns the Server field value if set, zero value otherwise.
func (o *StorageNetAppCloudTargetAllOf) GetServer() string {
	if o == nil || o.Server == nil {
		var ret string
		return ret
	}
	return *o.Server
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppCloudTargetAllOf) GetServerOk() (*string, bool) {
	if o == nil || o.Server == nil {
		return nil, false
	}
	return o.Server, true
}

// HasServer returns a boolean if a field has been set.
func (o *StorageNetAppCloudTargetAllOf) HasServer() bool {
	if o != nil && o.Server != nil {
		return true
	}

	return false
}

// SetServer gets a reference to the given string and assigns it to the Server field.
func (o *StorageNetAppCloudTargetAllOf) SetServer(v string) {
	o.Server = &v
}

// GetSnapMirrorUse returns the SnapMirrorUse field value if set, zero value otherwise.
func (o *StorageNetAppCloudTargetAllOf) GetSnapMirrorUse() string {
	if o == nil || o.SnapMirrorUse == nil {
		var ret string
		return ret
	}
	return *o.SnapMirrorUse
}

// GetSnapMirrorUseOk returns a tuple with the SnapMirrorUse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppCloudTargetAllOf) GetSnapMirrorUseOk() (*string, bool) {
	if o == nil || o.SnapMirrorUse == nil {
		return nil, false
	}
	return o.SnapMirrorUse, true
}

// HasSnapMirrorUse returns a boolean if a field has been set.
func (o *StorageNetAppCloudTargetAllOf) HasSnapMirrorUse() bool {
	if o != nil && o.SnapMirrorUse != nil {
		return true
	}

	return false
}

// SetSnapMirrorUse gets a reference to the given string and assigns it to the SnapMirrorUse field.
func (o *StorageNetAppCloudTargetAllOf) SetSnapMirrorUse(v string) {
	o.SnapMirrorUse = &v
}

// GetSslEnabled returns the SslEnabled field value if set, zero value otherwise.
func (o *StorageNetAppCloudTargetAllOf) GetSslEnabled() bool {
	if o == nil || o.SslEnabled == nil {
		var ret bool
		return ret
	}
	return *o.SslEnabled
}

// GetSslEnabledOk returns a tuple with the SslEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppCloudTargetAllOf) GetSslEnabledOk() (*bool, bool) {
	if o == nil || o.SslEnabled == nil {
		return nil, false
	}
	return o.SslEnabled, true
}

// HasSslEnabled returns a boolean if a field has been set.
func (o *StorageNetAppCloudTargetAllOf) HasSslEnabled() bool {
	if o != nil && o.SslEnabled != nil {
		return true
	}

	return false
}

// SetSslEnabled gets a reference to the given bool and assigns it to the SslEnabled field.
func (o *StorageNetAppCloudTargetAllOf) SetSslEnabled(v bool) {
	o.SslEnabled = &v
}

// GetUsed returns the Used field value if set, zero value otherwise.
func (o *StorageNetAppCloudTargetAllOf) GetUsed() int64 {
	if o == nil || o.Used == nil {
		var ret int64
		return ret
	}
	return *o.Used
}

// GetUsedOk returns a tuple with the Used field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppCloudTargetAllOf) GetUsedOk() (*int64, bool) {
	if o == nil || o.Used == nil {
		return nil, false
	}
	return o.Used, true
}

// HasUsed returns a boolean if a field has been set.
func (o *StorageNetAppCloudTargetAllOf) HasUsed() bool {
	if o != nil && o.Used != nil {
		return true
	}

	return false
}

// SetUsed gets a reference to the given int64 and assigns it to the Used field.
func (o *StorageNetAppCloudTargetAllOf) SetUsed(v int64) {
	o.Used = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *StorageNetAppCloudTargetAllOf) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppCloudTargetAllOf) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *StorageNetAppCloudTargetAllOf) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *StorageNetAppCloudTargetAllOf) SetUuid(v string) {
	o.Uuid = &v
}

// GetArray returns the Array field value if set, zero value otherwise.
func (o *StorageNetAppCloudTargetAllOf) GetArray() StorageNetAppClusterRelationship {
	if o == nil || o.Array == nil {
		var ret StorageNetAppClusterRelationship
		return ret
	}
	return *o.Array
}

// GetArrayOk returns a tuple with the Array field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppCloudTargetAllOf) GetArrayOk() (*StorageNetAppClusterRelationship, bool) {
	if o == nil || o.Array == nil {
		return nil, false
	}
	return o.Array, true
}

// HasArray returns a boolean if a field has been set.
func (o *StorageNetAppCloudTargetAllOf) HasArray() bool {
	if o != nil && o.Array != nil {
		return true
	}

	return false
}

// SetArray gets a reference to the given StorageNetAppClusterRelationship and assigns it to the Array field.
func (o *StorageNetAppCloudTargetAllOf) SetArray(v StorageNetAppClusterRelationship) {
	o.Array = &v
}

func (o StorageNetAppCloudTargetAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AccessKey != nil {
		toSerialize["AccessKey"] = o.AccessKey
	}
	if o.AuthenticationType != nil {
		toSerialize["AuthenticationType"] = o.AuthenticationType
	}
	if o.AzureAccount != nil {
		toSerialize["AzureAccount"] = o.AzureAccount
	}
	if o.CapUrl != nil {
		toSerialize["CapUrl"] = o.CapUrl
	}
	if o.CertificateValidationEnabled != nil {
		toSerialize["CertificateValidationEnabled"] = o.CertificateValidationEnabled
	}
	if o.CloudStorage != nil {
		toSerialize["CloudStorage"] = o.CloudStorage
	}
	if o.Container != nil {
		toSerialize["Container"] = o.Container
	}
	if o.Ipspace != nil {
		toSerialize["Ipspace"] = o.Ipspace
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Owner != nil {
		toSerialize["Owner"] = o.Owner
	}
	if o.Port != nil {
		toSerialize["Port"] = o.Port
	}
	if o.ProviderType != nil {
		toSerialize["ProviderType"] = o.ProviderType
	}
	if o.Server != nil {
		toSerialize["Server"] = o.Server
	}
	if o.SnapMirrorUse != nil {
		toSerialize["SnapMirrorUse"] = o.SnapMirrorUse
	}
	if o.SslEnabled != nil {
		toSerialize["SslEnabled"] = o.SslEnabled
	}
	if o.Used != nil {
		toSerialize["Used"] = o.Used
	}
	if o.Uuid != nil {
		toSerialize["Uuid"] = o.Uuid
	}
	if o.Array != nil {
		toSerialize["Array"] = o.Array
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageNetAppCloudTargetAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageNetAppCloudTargetAllOf := _StorageNetAppCloudTargetAllOf{}

	if err = json.Unmarshal(bytes, &varStorageNetAppCloudTargetAllOf); err == nil {
		*o = StorageNetAppCloudTargetAllOf(varStorageNetAppCloudTargetAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AccessKey")
		delete(additionalProperties, "AuthenticationType")
		delete(additionalProperties, "AzureAccount")
		delete(additionalProperties, "CapUrl")
		delete(additionalProperties, "CertificateValidationEnabled")
		delete(additionalProperties, "CloudStorage")
		delete(additionalProperties, "Container")
		delete(additionalProperties, "Ipspace")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Owner")
		delete(additionalProperties, "Port")
		delete(additionalProperties, "ProviderType")
		delete(additionalProperties, "Server")
		delete(additionalProperties, "SnapMirrorUse")
		delete(additionalProperties, "SslEnabled")
		delete(additionalProperties, "Used")
		delete(additionalProperties, "Uuid")
		delete(additionalProperties, "Array")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageNetAppCloudTargetAllOf struct {
	value *StorageNetAppCloudTargetAllOf
	isSet bool
}

func (v NullableStorageNetAppCloudTargetAllOf) Get() *StorageNetAppCloudTargetAllOf {
	return v.value
}

func (v *NullableStorageNetAppCloudTargetAllOf) Set(val *StorageNetAppCloudTargetAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppCloudTargetAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppCloudTargetAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppCloudTargetAllOf(val *StorageNetAppCloudTargetAllOf) *NullableStorageNetAppCloudTargetAllOf {
	return &NullableStorageNetAppCloudTargetAllOf{value: val, isSet: true}
}

func (v NullableStorageNetAppCloudTargetAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppCloudTargetAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
