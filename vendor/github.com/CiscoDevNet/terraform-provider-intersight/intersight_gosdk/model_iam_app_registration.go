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
	"time"
)

// IamAppRegistration AppRegistration encapsulates the meta-data values of a registered OAuth2 client application, as described in https://tools.ietf.org/html/rfc7591#section-2. Registered client applications have a set of metadata values associated with their client identifier at the Intersight authorization server, including the list of valid redirection URIs or a display name. The meta-data is used to specify how a client application can retrieve a OAuth2 Access Token and subsequently invoke Intersight API on behalf of this AppRegistration. To register an OAuth2 application, the following information must be provided. 1) Application name 2) An icon for the application 3) URL to the application's home page 4) A short description of the application 5) A list of redirect URLs When an AppRegistration is created, a unique OAuth2 clientId is generated and returned in the HTTP response.
type IamAppRegistration struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// A unique identifier for the OAuth2 client application. The client ID is auto-generated when the AppRegistration object is created.
	ClientId *string `json:"ClientId,omitempty"`
	// App Registration name specified by user.
	ClientName *string `json:"ClientName,omitempty"`
	// The OAuth2 client secret. The value of this property is generated when grantType includes 'client-credentials'. Otherwise, no client-secret is generated.
	ClientSecret *string `json:"ClientSecret,omitempty"`
	// The type of the OAuth2 client (public or confidential), as specified in https://tools.ietf.org/html/rfc6749#section-2.1. * `public` - Clients incapable of maintaining the confidentiality of their credentials.This includes clients executing on the device used by the resource owner,such as mobile applications, installed native application or a webbrowser-based application. * `confidential` - Clients capable of maintaining the confidentiality of their credentials.For example, this could be a client implemented on a secure server withrestricted access to the client credentials.To maintain the confidentiality of the OAuth2 credentials, two use cases areconsidered.1) The application is running as a service within Intersight. The application automatically   obtains the OAuth2 credentials when the application starts and the credentials are not   exposed to the end-user.   Because end-users (even account administrators) do not have access the OAuth2 credentials,   they cannot take the credentials with them when they leave their organization.2) The application is under the control of a \"trusted\" end-user. For example,   the end-user may create a native application running outside Intersight. The application   uses OAuth2 credentials to interact with the Intersight API. In that case, the Intersight   account administrator may generate OAuth2 credentials with a registered application   using \"client_credentials\" grant type.   In that case, the end-user is responsible for maintaining the confidentiality of the   OAuth2 credentials. If the end-user leaves the organization, you should revoke the   credentials and issue new Oauth2 credentials.Here is a possible workflow for handling OAuth2 tokens.1) User Alice (Intersight Account Administrator) logins to Intersight and deploys an Intersight   application that requires an OAuth2 token.2) Intersight automatically deploys the application. The application is assigned a OAuth2 token,   possibly linked to Alice. The application must NOT expose the OAuth2 secret to Alice, otherwise   Alice would be able to use the token after she leaves the company.3) The application can make API calls to Intersight using its assigned OAuth2 token. For example,   the application could make weekly scheduled API calls to Intersight.4) Separately, Alice may also get OAuth2 tokens that she can use to make API calls from the   Intersight SDK through the northbound API. In that case, Alice will get the associated OAuth2   secrets, but not the one assigned in step #2.5) Alice leaves the organization. The OAuth2 tokens assigned in step #2 must retain their validity   even after Alice has left the organization. Because the OAuth2 secrets were never shared with   Alice, there is no risk Alice can reuse the OAuth2 secrets.   On the other hand, the OAuth2 tokens assigned in step #4 must be invalidated because Alice had   the OAuth2 tokens in her possession.
	ClientType *string `json:"ClientType,omitempty"`
	// Description of the application.
	Description  *string  `json:"Description,omitempty"`
	GrantTypes   []string `json:"GrantTypes,omitempty"`
	RedirectUris []string `json:"RedirectUris,omitempty"`
	// Set value to true to renew the client-secret. Applicable to client_credentials grant type.
	RenewClientSecret *bool    `json:"RenewClientSecret,omitempty"`
	ResponseTypes     []string `json:"ResponseTypes,omitempty"`
	// Used to perform revocation for tokens of AppRegistration. Updated only internally is case Revoke property come from UI with value true. On each request with OAuth2 access token the CreationTime of the OAuth2 token will be compared to RevokationTimestamp of the corresponding App Registration.
	RevocationTimestamp *time.Time `json:"RevocationTimestamp,omitempty"`
	// Used to trigger update the revocationTimestamp value. If UI sent updating request with the Revoke value is true, then update RevocationTimestamp.
	Revoke  *bool                   `json:"Revoke,omitempty"`
	Account *IamAccountRelationship `json:"Account,omitempty"`
	// An array of relationships to iamOAuthToken resources.
	OauthTokens []IamOAuthTokenRelationship `json:"OauthTokens,omitempty"`
	Permission  *IamPermissionRelationship  `json:"Permission,omitempty"`
	// An array of relationships to iamRole resources.
	// Deprecated
	Roles                []IamRoleRelationship `json:"Roles,omitempty"`
	User                 *IamUserRelationship  `json:"User,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamAppRegistration IamAppRegistration

// NewIamAppRegistration instantiates a new IamAppRegistration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamAppRegistration(classId string, objectType string) *IamAppRegistration {
	this := IamAppRegistration{}
	this.ClassId = classId
	this.ObjectType = objectType
	var clientType string = "public"
	this.ClientType = &clientType
	var renewClientSecret bool = false
	this.RenewClientSecret = &renewClientSecret
	var revoke bool = false
	this.Revoke = &revoke
	return &this
}

// NewIamAppRegistrationWithDefaults instantiates a new IamAppRegistration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamAppRegistrationWithDefaults() *IamAppRegistration {
	this := IamAppRegistration{}
	var classId string = "iam.AppRegistration"
	this.ClassId = classId
	var objectType string = "iam.AppRegistration"
	this.ObjectType = objectType
	var clientType string = "public"
	this.ClientType = &clientType
	var renewClientSecret bool = false
	this.RenewClientSecret = &renewClientSecret
	var revoke bool = false
	this.Revoke = &revoke
	return &this
}

// GetClassId returns the ClassId field value
func (o *IamAppRegistration) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IamAppRegistration) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IamAppRegistration) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IamAppRegistration) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IamAppRegistration) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IamAppRegistration) SetObjectType(v string) {
	o.ObjectType = v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *IamAppRegistration) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamAppRegistration) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *IamAppRegistration) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *IamAppRegistration) SetClientId(v string) {
	o.ClientId = &v
}

// GetClientName returns the ClientName field value if set, zero value otherwise.
func (o *IamAppRegistration) GetClientName() string {
	if o == nil || o.ClientName == nil {
		var ret string
		return ret
	}
	return *o.ClientName
}

// GetClientNameOk returns a tuple with the ClientName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamAppRegistration) GetClientNameOk() (*string, bool) {
	if o == nil || o.ClientName == nil {
		return nil, false
	}
	return o.ClientName, true
}

// HasClientName returns a boolean if a field has been set.
func (o *IamAppRegistration) HasClientName() bool {
	if o != nil && o.ClientName != nil {
		return true
	}

	return false
}

// SetClientName gets a reference to the given string and assigns it to the ClientName field.
func (o *IamAppRegistration) SetClientName(v string) {
	o.ClientName = &v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *IamAppRegistration) GetClientSecret() string {
	if o == nil || o.ClientSecret == nil {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamAppRegistration) GetClientSecretOk() (*string, bool) {
	if o == nil || o.ClientSecret == nil {
		return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *IamAppRegistration) HasClientSecret() bool {
	if o != nil && o.ClientSecret != nil {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *IamAppRegistration) SetClientSecret(v string) {
	o.ClientSecret = &v
}

// GetClientType returns the ClientType field value if set, zero value otherwise.
func (o *IamAppRegistration) GetClientType() string {
	if o == nil || o.ClientType == nil {
		var ret string
		return ret
	}
	return *o.ClientType
}

// GetClientTypeOk returns a tuple with the ClientType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamAppRegistration) GetClientTypeOk() (*string, bool) {
	if o == nil || o.ClientType == nil {
		return nil, false
	}
	return o.ClientType, true
}

// HasClientType returns a boolean if a field has been set.
func (o *IamAppRegistration) HasClientType() bool {
	if o != nil && o.ClientType != nil {
		return true
	}

	return false
}

// SetClientType gets a reference to the given string and assigns it to the ClientType field.
func (o *IamAppRegistration) SetClientType(v string) {
	o.ClientType = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IamAppRegistration) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamAppRegistration) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IamAppRegistration) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IamAppRegistration) SetDescription(v string) {
	o.Description = &v
}

// GetGrantTypes returns the GrantTypes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamAppRegistration) GetGrantTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.GrantTypes
}

// GetGrantTypesOk returns a tuple with the GrantTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamAppRegistration) GetGrantTypesOk() ([]string, bool) {
	if o == nil || o.GrantTypes == nil {
		return nil, false
	}
	return o.GrantTypes, true
}

// HasGrantTypes returns a boolean if a field has been set.
func (o *IamAppRegistration) HasGrantTypes() bool {
	if o != nil && o.GrantTypes != nil {
		return true
	}

	return false
}

// SetGrantTypes gets a reference to the given []string and assigns it to the GrantTypes field.
func (o *IamAppRegistration) SetGrantTypes(v []string) {
	o.GrantTypes = v
}

// GetRedirectUris returns the RedirectUris field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamAppRegistration) GetRedirectUris() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.RedirectUris
}

// GetRedirectUrisOk returns a tuple with the RedirectUris field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamAppRegistration) GetRedirectUrisOk() ([]string, bool) {
	if o == nil || o.RedirectUris == nil {
		return nil, false
	}
	return o.RedirectUris, true
}

// HasRedirectUris returns a boolean if a field has been set.
func (o *IamAppRegistration) HasRedirectUris() bool {
	if o != nil && o.RedirectUris != nil {
		return true
	}

	return false
}

// SetRedirectUris gets a reference to the given []string and assigns it to the RedirectUris field.
func (o *IamAppRegistration) SetRedirectUris(v []string) {
	o.RedirectUris = v
}

// GetRenewClientSecret returns the RenewClientSecret field value if set, zero value otherwise.
func (o *IamAppRegistration) GetRenewClientSecret() bool {
	if o == nil || o.RenewClientSecret == nil {
		var ret bool
		return ret
	}
	return *o.RenewClientSecret
}

// GetRenewClientSecretOk returns a tuple with the RenewClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamAppRegistration) GetRenewClientSecretOk() (*bool, bool) {
	if o == nil || o.RenewClientSecret == nil {
		return nil, false
	}
	return o.RenewClientSecret, true
}

// HasRenewClientSecret returns a boolean if a field has been set.
func (o *IamAppRegistration) HasRenewClientSecret() bool {
	if o != nil && o.RenewClientSecret != nil {
		return true
	}

	return false
}

// SetRenewClientSecret gets a reference to the given bool and assigns it to the RenewClientSecret field.
func (o *IamAppRegistration) SetRenewClientSecret(v bool) {
	o.RenewClientSecret = &v
}

// GetResponseTypes returns the ResponseTypes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamAppRegistration) GetResponseTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ResponseTypes
}

// GetResponseTypesOk returns a tuple with the ResponseTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamAppRegistration) GetResponseTypesOk() ([]string, bool) {
	if o == nil || o.ResponseTypes == nil {
		return nil, false
	}
	return o.ResponseTypes, true
}

// HasResponseTypes returns a boolean if a field has been set.
func (o *IamAppRegistration) HasResponseTypes() bool {
	if o != nil && o.ResponseTypes != nil {
		return true
	}

	return false
}

// SetResponseTypes gets a reference to the given []string and assigns it to the ResponseTypes field.
func (o *IamAppRegistration) SetResponseTypes(v []string) {
	o.ResponseTypes = v
}

// GetRevocationTimestamp returns the RevocationTimestamp field value if set, zero value otherwise.
func (o *IamAppRegistration) GetRevocationTimestamp() time.Time {
	if o == nil || o.RevocationTimestamp == nil {
		var ret time.Time
		return ret
	}
	return *o.RevocationTimestamp
}

// GetRevocationTimestampOk returns a tuple with the RevocationTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamAppRegistration) GetRevocationTimestampOk() (*time.Time, bool) {
	if o == nil || o.RevocationTimestamp == nil {
		return nil, false
	}
	return o.RevocationTimestamp, true
}

// HasRevocationTimestamp returns a boolean if a field has been set.
func (o *IamAppRegistration) HasRevocationTimestamp() bool {
	if o != nil && o.RevocationTimestamp != nil {
		return true
	}

	return false
}

// SetRevocationTimestamp gets a reference to the given time.Time and assigns it to the RevocationTimestamp field.
func (o *IamAppRegistration) SetRevocationTimestamp(v time.Time) {
	o.RevocationTimestamp = &v
}

// GetRevoke returns the Revoke field value if set, zero value otherwise.
func (o *IamAppRegistration) GetRevoke() bool {
	if o == nil || o.Revoke == nil {
		var ret bool
		return ret
	}
	return *o.Revoke
}

// GetRevokeOk returns a tuple with the Revoke field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamAppRegistration) GetRevokeOk() (*bool, bool) {
	if o == nil || o.Revoke == nil {
		return nil, false
	}
	return o.Revoke, true
}

// HasRevoke returns a boolean if a field has been set.
func (o *IamAppRegistration) HasRevoke() bool {
	if o != nil && o.Revoke != nil {
		return true
	}

	return false
}

// SetRevoke gets a reference to the given bool and assigns it to the Revoke field.
func (o *IamAppRegistration) SetRevoke(v bool) {
	o.Revoke = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *IamAppRegistration) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamAppRegistration) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *IamAppRegistration) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *IamAppRegistration) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

// GetOauthTokens returns the OauthTokens field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamAppRegistration) GetOauthTokens() []IamOAuthTokenRelationship {
	if o == nil {
		var ret []IamOAuthTokenRelationship
		return ret
	}
	return o.OauthTokens
}

// GetOauthTokensOk returns a tuple with the OauthTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamAppRegistration) GetOauthTokensOk() ([]IamOAuthTokenRelationship, bool) {
	if o == nil || o.OauthTokens == nil {
		return nil, false
	}
	return o.OauthTokens, true
}

// HasOauthTokens returns a boolean if a field has been set.
func (o *IamAppRegistration) HasOauthTokens() bool {
	if o != nil && o.OauthTokens != nil {
		return true
	}

	return false
}

// SetOauthTokens gets a reference to the given []IamOAuthTokenRelationship and assigns it to the OauthTokens field.
func (o *IamAppRegistration) SetOauthTokens(v []IamOAuthTokenRelationship) {
	o.OauthTokens = v
}

// GetPermission returns the Permission field value if set, zero value otherwise.
func (o *IamAppRegistration) GetPermission() IamPermissionRelationship {
	if o == nil || o.Permission == nil {
		var ret IamPermissionRelationship
		return ret
	}
	return *o.Permission
}

// GetPermissionOk returns a tuple with the Permission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamAppRegistration) GetPermissionOk() (*IamPermissionRelationship, bool) {
	if o == nil || o.Permission == nil {
		return nil, false
	}
	return o.Permission, true
}

// HasPermission returns a boolean if a field has been set.
func (o *IamAppRegistration) HasPermission() bool {
	if o != nil && o.Permission != nil {
		return true
	}

	return false
}

// SetPermission gets a reference to the given IamPermissionRelationship and assigns it to the Permission field.
func (o *IamAppRegistration) SetPermission(v IamPermissionRelationship) {
	o.Permission = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise (both if not set or set to explicit null).
// Deprecated
func (o *IamAppRegistration) GetRoles() []IamRoleRelationship {
	if o == nil {
		var ret []IamRoleRelationship
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
// Deprecated
func (o *IamAppRegistration) GetRolesOk() ([]IamRoleRelationship, bool) {
	if o == nil || o.Roles == nil {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *IamAppRegistration) HasRoles() bool {
	if o != nil && o.Roles != nil {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []IamRoleRelationship and assigns it to the Roles field.
// Deprecated
func (o *IamAppRegistration) SetRoles(v []IamRoleRelationship) {
	o.Roles = v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *IamAppRegistration) GetUser() IamUserRelationship {
	if o == nil || o.User == nil {
		var ret IamUserRelationship
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamAppRegistration) GetUserOk() (*IamUserRelationship, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *IamAppRegistration) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given IamUserRelationship and assigns it to the User field.
func (o *IamAppRegistration) SetUser(v IamUserRelationship) {
	o.User = &v
}

func (o IamAppRegistration) MarshalJSON() ([]byte, error) {
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
	if o.ClientId != nil {
		toSerialize["ClientId"] = o.ClientId
	}
	if o.ClientName != nil {
		toSerialize["ClientName"] = o.ClientName
	}
	if o.ClientSecret != nil {
		toSerialize["ClientSecret"] = o.ClientSecret
	}
	if o.ClientType != nil {
		toSerialize["ClientType"] = o.ClientType
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.GrantTypes != nil {
		toSerialize["GrantTypes"] = o.GrantTypes
	}
	if o.RedirectUris != nil {
		toSerialize["RedirectUris"] = o.RedirectUris
	}
	if o.RenewClientSecret != nil {
		toSerialize["RenewClientSecret"] = o.RenewClientSecret
	}
	if o.ResponseTypes != nil {
		toSerialize["ResponseTypes"] = o.ResponseTypes
	}
	if o.RevocationTimestamp != nil {
		toSerialize["RevocationTimestamp"] = o.RevocationTimestamp
	}
	if o.Revoke != nil {
		toSerialize["Revoke"] = o.Revoke
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}
	if o.OauthTokens != nil {
		toSerialize["OauthTokens"] = o.OauthTokens
	}
	if o.Permission != nil {
		toSerialize["Permission"] = o.Permission
	}
	if o.Roles != nil {
		toSerialize["Roles"] = o.Roles
	}
	if o.User != nil {
		toSerialize["User"] = o.User
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IamAppRegistration) UnmarshalJSON(bytes []byte) (err error) {
	type IamAppRegistrationWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// A unique identifier for the OAuth2 client application. The client ID is auto-generated when the AppRegistration object is created.
		ClientId *string `json:"ClientId,omitempty"`
		// App Registration name specified by user.
		ClientName *string `json:"ClientName,omitempty"`
		// The OAuth2 client secret. The value of this property is generated when grantType includes 'client-credentials'. Otherwise, no client-secret is generated.
		ClientSecret *string `json:"ClientSecret,omitempty"`
		// The type of the OAuth2 client (public or confidential), as specified in https://tools.ietf.org/html/rfc6749#section-2.1. * `public` - Clients incapable of maintaining the confidentiality of their credentials.This includes clients executing on the device used by the resource owner,such as mobile applications, installed native application or a webbrowser-based application. * `confidential` - Clients capable of maintaining the confidentiality of their credentials.For example, this could be a client implemented on a secure server withrestricted access to the client credentials.To maintain the confidentiality of the OAuth2 credentials, two use cases areconsidered.1) The application is running as a service within Intersight. The application automatically   obtains the OAuth2 credentials when the application starts and the credentials are not   exposed to the end-user.   Because end-users (even account administrators) do not have access the OAuth2 credentials,   they cannot take the credentials with them when they leave their organization.2) The application is under the control of a \"trusted\" end-user. For example,   the end-user may create a native application running outside Intersight. The application   uses OAuth2 credentials to interact with the Intersight API. In that case, the Intersight   account administrator may generate OAuth2 credentials with a registered application   using \"client_credentials\" grant type.   In that case, the end-user is responsible for maintaining the confidentiality of the   OAuth2 credentials. If the end-user leaves the organization, you should revoke the   credentials and issue new Oauth2 credentials.Here is a possible workflow for handling OAuth2 tokens.1) User Alice (Intersight Account Administrator) logins to Intersight and deploys an Intersight   application that requires an OAuth2 token.2) Intersight automatically deploys the application. The application is assigned a OAuth2 token,   possibly linked to Alice. The application must NOT expose the OAuth2 secret to Alice, otherwise   Alice would be able to use the token after she leaves the company.3) The application can make API calls to Intersight using its assigned OAuth2 token. For example,   the application could make weekly scheduled API calls to Intersight.4) Separately, Alice may also get OAuth2 tokens that she can use to make API calls from the   Intersight SDK through the northbound API. In that case, Alice will get the associated OAuth2   secrets, but not the one assigned in step #2.5) Alice leaves the organization. The OAuth2 tokens assigned in step #2 must retain their validity   even after Alice has left the organization. Because the OAuth2 secrets were never shared with   Alice, there is no risk Alice can reuse the OAuth2 secrets.   On the other hand, the OAuth2 tokens assigned in step #4 must be invalidated because Alice had   the OAuth2 tokens in her possession.
		ClientType *string `json:"ClientType,omitempty"`
		// Description of the application.
		Description  *string  `json:"Description,omitempty"`
		GrantTypes   []string `json:"GrantTypes,omitempty"`
		RedirectUris []string `json:"RedirectUris,omitempty"`
		// Set value to true to renew the client-secret. Applicable to client_credentials grant type.
		RenewClientSecret *bool    `json:"RenewClientSecret,omitempty"`
		ResponseTypes     []string `json:"ResponseTypes,omitempty"`
		// Used to perform revocation for tokens of AppRegistration. Updated only internally is case Revoke property come from UI with value true. On each request with OAuth2 access token the CreationTime of the OAuth2 token will be compared to RevokationTimestamp of the corresponding App Registration.
		RevocationTimestamp *time.Time `json:"RevocationTimestamp,omitempty"`
		// Used to trigger update the revocationTimestamp value. If UI sent updating request with the Revoke value is true, then update RevocationTimestamp.
		Revoke  *bool                   `json:"Revoke,omitempty"`
		Account *IamAccountRelationship `json:"Account,omitempty"`
		// An array of relationships to iamOAuthToken resources.
		OauthTokens []IamOAuthTokenRelationship `json:"OauthTokens,omitempty"`
		Permission  *IamPermissionRelationship  `json:"Permission,omitempty"`
		// An array of relationships to iamRole resources.
		// Deprecated
		Roles []IamRoleRelationship `json:"Roles,omitempty"`
		User  *IamUserRelationship  `json:"User,omitempty"`
	}

	varIamAppRegistrationWithoutEmbeddedStruct := IamAppRegistrationWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varIamAppRegistrationWithoutEmbeddedStruct)
	if err == nil {
		varIamAppRegistration := _IamAppRegistration{}
		varIamAppRegistration.ClassId = varIamAppRegistrationWithoutEmbeddedStruct.ClassId
		varIamAppRegistration.ObjectType = varIamAppRegistrationWithoutEmbeddedStruct.ObjectType
		varIamAppRegistration.ClientId = varIamAppRegistrationWithoutEmbeddedStruct.ClientId
		varIamAppRegistration.ClientName = varIamAppRegistrationWithoutEmbeddedStruct.ClientName
		varIamAppRegistration.ClientSecret = varIamAppRegistrationWithoutEmbeddedStruct.ClientSecret
		varIamAppRegistration.ClientType = varIamAppRegistrationWithoutEmbeddedStruct.ClientType
		varIamAppRegistration.Description = varIamAppRegistrationWithoutEmbeddedStruct.Description
		varIamAppRegistration.GrantTypes = varIamAppRegistrationWithoutEmbeddedStruct.GrantTypes
		varIamAppRegistration.RedirectUris = varIamAppRegistrationWithoutEmbeddedStruct.RedirectUris
		varIamAppRegistration.RenewClientSecret = varIamAppRegistrationWithoutEmbeddedStruct.RenewClientSecret
		varIamAppRegistration.ResponseTypes = varIamAppRegistrationWithoutEmbeddedStruct.ResponseTypes
		varIamAppRegistration.RevocationTimestamp = varIamAppRegistrationWithoutEmbeddedStruct.RevocationTimestamp
		varIamAppRegistration.Revoke = varIamAppRegistrationWithoutEmbeddedStruct.Revoke
		varIamAppRegistration.Account = varIamAppRegistrationWithoutEmbeddedStruct.Account
		varIamAppRegistration.OauthTokens = varIamAppRegistrationWithoutEmbeddedStruct.OauthTokens
		varIamAppRegistration.Permission = varIamAppRegistrationWithoutEmbeddedStruct.Permission
		varIamAppRegistration.Roles = varIamAppRegistrationWithoutEmbeddedStruct.Roles
		varIamAppRegistration.User = varIamAppRegistrationWithoutEmbeddedStruct.User
		*o = IamAppRegistration(varIamAppRegistration)
	} else {
		return err
	}

	varIamAppRegistration := _IamAppRegistration{}

	err = json.Unmarshal(bytes, &varIamAppRegistration)
	if err == nil {
		o.MoBaseMo = varIamAppRegistration.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ClientId")
		delete(additionalProperties, "ClientName")
		delete(additionalProperties, "ClientSecret")
		delete(additionalProperties, "ClientType")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "GrantTypes")
		delete(additionalProperties, "RedirectUris")
		delete(additionalProperties, "RenewClientSecret")
		delete(additionalProperties, "ResponseTypes")
		delete(additionalProperties, "RevocationTimestamp")
		delete(additionalProperties, "Revoke")
		delete(additionalProperties, "Account")
		delete(additionalProperties, "OauthTokens")
		delete(additionalProperties, "Permission")
		delete(additionalProperties, "Roles")
		delete(additionalProperties, "User")

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

type NullableIamAppRegistration struct {
	value *IamAppRegistration
	isSet bool
}

func (v NullableIamAppRegistration) Get() *IamAppRegistration {
	return v.value
}

func (v *NullableIamAppRegistration) Set(val *IamAppRegistration) {
	v.value = val
	v.isSet = true
}

func (v NullableIamAppRegistration) IsSet() bool {
	return v.isSet
}

func (v *NullableIamAppRegistration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamAppRegistration(val *IamAppRegistration) *NullableIamAppRegistration {
	return &NullableIamAppRegistration{value: val, isSet: true}
}

func (v NullableIamAppRegistration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamAppRegistration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
