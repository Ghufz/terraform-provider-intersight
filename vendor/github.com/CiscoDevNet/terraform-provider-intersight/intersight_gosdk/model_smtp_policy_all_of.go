/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-7658
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// SmtpPolicyAllOf Definition of the list of properties defined in 'smtp.Policy', excluding properties defined in parent classes.
type SmtpPolicyAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Authorization password for the process.
	AuthPassword *string `json:"AuthPassword,omitempty"`
	// If enabled, lets user input username and password.
	EnableAuth *bool `json:"EnableAuth,omitempty"`
	// If enabled, lets user input valid CA certificates for authorization.
	EnableTls *bool `json:"EnableTls,omitempty"`
	// If enabled, controls the state of the SMTP client service on the managed device.
	Enabled *bool `json:"Enabled,omitempty"`
	// Indicates whether the value of the 'authPassword' property has been set.
	IsAuthPasswordSet *bool `json:"IsAuthPasswordSet,omitempty"`
	// Minimum fault severity level to receive email notifications. Email notifications are sent for all faults whose severity is equal to or greater than the chosen level. * `critical` - Minimum severity to report is critical. * `condition` - Minimum severity to report is informational. * `warning` - Minimum severity to report is warning. * `minor` - Minimum severity to report is minor. * `major` - Minimum severity to report is major.
	MinSeverity *string `json:"MinSeverity,omitempty"`
	// The email address entered here will be displayed as the from address (mail received from address) of all the SMTP mail alerts that are received. If not configured, the hostname of the server is used in the from address field.
	SenderEmail *string `json:"SenderEmail,omitempty"`
	// Port number used by the SMTP server for outgoing SMTP communication.
	SmtpPort       *int64   `json:"SmtpPort,omitempty"`
	SmtpRecipients []string `json:"SmtpRecipients,omitempty"`
	// IP address or hostname of the SMTP server. The SMTP server is used by the managed device to send email notifications.
	SmtpServer *string `json:"SmtpServer,omitempty"`
	// SMTP username from which email notification is sent.
	UserName         *string                               `json:"UserName,omitempty"`
	ApplianceAccount *IamAccountRelationship               `json:"ApplianceAccount,omitempty"`
	Certificate      *IamTrustPointRelationship            `json:"Certificate,omitempty"`
	Organization     *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	// An array of relationships to policyAbstractConfigProfile resources.
	Profiles             []PolicyAbstractConfigProfileRelationship `json:"Profiles,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SmtpPolicyAllOf SmtpPolicyAllOf

// NewSmtpPolicyAllOf instantiates a new SmtpPolicyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmtpPolicyAllOf(classId string, objectType string) *SmtpPolicyAllOf {
	this := SmtpPolicyAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var enableAuth bool = false
	this.EnableAuth = &enableAuth
	var enableTls bool = false
	this.EnableTls = &enableTls
	var enabled bool = true
	this.Enabled = &enabled
	var minSeverity string = "critical"
	this.MinSeverity = &minSeverity
	var smtpPort int64 = 25
	this.SmtpPort = &smtpPort
	return &this
}

// NewSmtpPolicyAllOfWithDefaults instantiates a new SmtpPolicyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmtpPolicyAllOfWithDefaults() *SmtpPolicyAllOf {
	this := SmtpPolicyAllOf{}
	var classId string = "smtp.Policy"
	this.ClassId = classId
	var objectType string = "smtp.Policy"
	this.ObjectType = objectType
	var enableAuth bool = false
	this.EnableAuth = &enableAuth
	var enableTls bool = false
	this.EnableTls = &enableTls
	var enabled bool = true
	this.Enabled = &enabled
	var minSeverity string = "critical"
	this.MinSeverity = &minSeverity
	var smtpPort int64 = 25
	this.SmtpPort = &smtpPort
	return &this
}

// GetClassId returns the ClassId field value
func (o *SmtpPolicyAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SmtpPolicyAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SmtpPolicyAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *SmtpPolicyAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SmtpPolicyAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SmtpPolicyAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAuthPassword returns the AuthPassword field value if set, zero value otherwise.
func (o *SmtpPolicyAllOf) GetAuthPassword() string {
	if o == nil || o.AuthPassword == nil {
		var ret string
		return ret
	}
	return *o.AuthPassword
}

// GetAuthPasswordOk returns a tuple with the AuthPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmtpPolicyAllOf) GetAuthPasswordOk() (*string, bool) {
	if o == nil || o.AuthPassword == nil {
		return nil, false
	}
	return o.AuthPassword, true
}

// HasAuthPassword returns a boolean if a field has been set.
func (o *SmtpPolicyAllOf) HasAuthPassword() bool {
	if o != nil && o.AuthPassword != nil {
		return true
	}

	return false
}

// SetAuthPassword gets a reference to the given string and assigns it to the AuthPassword field.
func (o *SmtpPolicyAllOf) SetAuthPassword(v string) {
	o.AuthPassword = &v
}

// GetEnableAuth returns the EnableAuth field value if set, zero value otherwise.
func (o *SmtpPolicyAllOf) GetEnableAuth() bool {
	if o == nil || o.EnableAuth == nil {
		var ret bool
		return ret
	}
	return *o.EnableAuth
}

// GetEnableAuthOk returns a tuple with the EnableAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmtpPolicyAllOf) GetEnableAuthOk() (*bool, bool) {
	if o == nil || o.EnableAuth == nil {
		return nil, false
	}
	return o.EnableAuth, true
}

// HasEnableAuth returns a boolean if a field has been set.
func (o *SmtpPolicyAllOf) HasEnableAuth() bool {
	if o != nil && o.EnableAuth != nil {
		return true
	}

	return false
}

// SetEnableAuth gets a reference to the given bool and assigns it to the EnableAuth field.
func (o *SmtpPolicyAllOf) SetEnableAuth(v bool) {
	o.EnableAuth = &v
}

// GetEnableTls returns the EnableTls field value if set, zero value otherwise.
func (o *SmtpPolicyAllOf) GetEnableTls() bool {
	if o == nil || o.EnableTls == nil {
		var ret bool
		return ret
	}
	return *o.EnableTls
}

// GetEnableTlsOk returns a tuple with the EnableTls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmtpPolicyAllOf) GetEnableTlsOk() (*bool, bool) {
	if o == nil || o.EnableTls == nil {
		return nil, false
	}
	return o.EnableTls, true
}

// HasEnableTls returns a boolean if a field has been set.
func (o *SmtpPolicyAllOf) HasEnableTls() bool {
	if o != nil && o.EnableTls != nil {
		return true
	}

	return false
}

// SetEnableTls gets a reference to the given bool and assigns it to the EnableTls field.
func (o *SmtpPolicyAllOf) SetEnableTls(v bool) {
	o.EnableTls = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SmtpPolicyAllOf) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmtpPolicyAllOf) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SmtpPolicyAllOf) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SmtpPolicyAllOf) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetIsAuthPasswordSet returns the IsAuthPasswordSet field value if set, zero value otherwise.
func (o *SmtpPolicyAllOf) GetIsAuthPasswordSet() bool {
	if o == nil || o.IsAuthPasswordSet == nil {
		var ret bool
		return ret
	}
	return *o.IsAuthPasswordSet
}

// GetIsAuthPasswordSetOk returns a tuple with the IsAuthPasswordSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmtpPolicyAllOf) GetIsAuthPasswordSetOk() (*bool, bool) {
	if o == nil || o.IsAuthPasswordSet == nil {
		return nil, false
	}
	return o.IsAuthPasswordSet, true
}

// HasIsAuthPasswordSet returns a boolean if a field has been set.
func (o *SmtpPolicyAllOf) HasIsAuthPasswordSet() bool {
	if o != nil && o.IsAuthPasswordSet != nil {
		return true
	}

	return false
}

// SetIsAuthPasswordSet gets a reference to the given bool and assigns it to the IsAuthPasswordSet field.
func (o *SmtpPolicyAllOf) SetIsAuthPasswordSet(v bool) {
	o.IsAuthPasswordSet = &v
}

// GetMinSeverity returns the MinSeverity field value if set, zero value otherwise.
func (o *SmtpPolicyAllOf) GetMinSeverity() string {
	if o == nil || o.MinSeverity == nil {
		var ret string
		return ret
	}
	return *o.MinSeverity
}

// GetMinSeverityOk returns a tuple with the MinSeverity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmtpPolicyAllOf) GetMinSeverityOk() (*string, bool) {
	if o == nil || o.MinSeverity == nil {
		return nil, false
	}
	return o.MinSeverity, true
}

// HasMinSeverity returns a boolean if a field has been set.
func (o *SmtpPolicyAllOf) HasMinSeverity() bool {
	if o != nil && o.MinSeverity != nil {
		return true
	}

	return false
}

// SetMinSeverity gets a reference to the given string and assigns it to the MinSeverity field.
func (o *SmtpPolicyAllOf) SetMinSeverity(v string) {
	o.MinSeverity = &v
}

// GetSenderEmail returns the SenderEmail field value if set, zero value otherwise.
func (o *SmtpPolicyAllOf) GetSenderEmail() string {
	if o == nil || o.SenderEmail == nil {
		var ret string
		return ret
	}
	return *o.SenderEmail
}

// GetSenderEmailOk returns a tuple with the SenderEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmtpPolicyAllOf) GetSenderEmailOk() (*string, bool) {
	if o == nil || o.SenderEmail == nil {
		return nil, false
	}
	return o.SenderEmail, true
}

// HasSenderEmail returns a boolean if a field has been set.
func (o *SmtpPolicyAllOf) HasSenderEmail() bool {
	if o != nil && o.SenderEmail != nil {
		return true
	}

	return false
}

// SetSenderEmail gets a reference to the given string and assigns it to the SenderEmail field.
func (o *SmtpPolicyAllOf) SetSenderEmail(v string) {
	o.SenderEmail = &v
}

// GetSmtpPort returns the SmtpPort field value if set, zero value otherwise.
func (o *SmtpPolicyAllOf) GetSmtpPort() int64 {
	if o == nil || o.SmtpPort == nil {
		var ret int64
		return ret
	}
	return *o.SmtpPort
}

// GetSmtpPortOk returns a tuple with the SmtpPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmtpPolicyAllOf) GetSmtpPortOk() (*int64, bool) {
	if o == nil || o.SmtpPort == nil {
		return nil, false
	}
	return o.SmtpPort, true
}

// HasSmtpPort returns a boolean if a field has been set.
func (o *SmtpPolicyAllOf) HasSmtpPort() bool {
	if o != nil && o.SmtpPort != nil {
		return true
	}

	return false
}

// SetSmtpPort gets a reference to the given int64 and assigns it to the SmtpPort field.
func (o *SmtpPolicyAllOf) SetSmtpPort(v int64) {
	o.SmtpPort = &v
}

// GetSmtpRecipients returns the SmtpRecipients field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SmtpPolicyAllOf) GetSmtpRecipients() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SmtpRecipients
}

// GetSmtpRecipientsOk returns a tuple with the SmtpRecipients field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SmtpPolicyAllOf) GetSmtpRecipientsOk() ([]string, bool) {
	if o == nil || o.SmtpRecipients == nil {
		return nil, false
	}
	return o.SmtpRecipients, true
}

// HasSmtpRecipients returns a boolean if a field has been set.
func (o *SmtpPolicyAllOf) HasSmtpRecipients() bool {
	if o != nil && o.SmtpRecipients != nil {
		return true
	}

	return false
}

// SetSmtpRecipients gets a reference to the given []string and assigns it to the SmtpRecipients field.
func (o *SmtpPolicyAllOf) SetSmtpRecipients(v []string) {
	o.SmtpRecipients = v
}

// GetSmtpServer returns the SmtpServer field value if set, zero value otherwise.
func (o *SmtpPolicyAllOf) GetSmtpServer() string {
	if o == nil || o.SmtpServer == nil {
		var ret string
		return ret
	}
	return *o.SmtpServer
}

// GetSmtpServerOk returns a tuple with the SmtpServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmtpPolicyAllOf) GetSmtpServerOk() (*string, bool) {
	if o == nil || o.SmtpServer == nil {
		return nil, false
	}
	return o.SmtpServer, true
}

// HasSmtpServer returns a boolean if a field has been set.
func (o *SmtpPolicyAllOf) HasSmtpServer() bool {
	if o != nil && o.SmtpServer != nil {
		return true
	}

	return false
}

// SetSmtpServer gets a reference to the given string and assigns it to the SmtpServer field.
func (o *SmtpPolicyAllOf) SetSmtpServer(v string) {
	o.SmtpServer = &v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *SmtpPolicyAllOf) GetUserName() string {
	if o == nil || o.UserName == nil {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmtpPolicyAllOf) GetUserNameOk() (*string, bool) {
	if o == nil || o.UserName == nil {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *SmtpPolicyAllOf) HasUserName() bool {
	if o != nil && o.UserName != nil {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *SmtpPolicyAllOf) SetUserName(v string) {
	o.UserName = &v
}

// GetApplianceAccount returns the ApplianceAccount field value if set, zero value otherwise.
func (o *SmtpPolicyAllOf) GetApplianceAccount() IamAccountRelationship {
	if o == nil || o.ApplianceAccount == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.ApplianceAccount
}

// GetApplianceAccountOk returns a tuple with the ApplianceAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmtpPolicyAllOf) GetApplianceAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.ApplianceAccount == nil {
		return nil, false
	}
	return o.ApplianceAccount, true
}

// HasApplianceAccount returns a boolean if a field has been set.
func (o *SmtpPolicyAllOf) HasApplianceAccount() bool {
	if o != nil && o.ApplianceAccount != nil {
		return true
	}

	return false
}

// SetApplianceAccount gets a reference to the given IamAccountRelationship and assigns it to the ApplianceAccount field.
func (o *SmtpPolicyAllOf) SetApplianceAccount(v IamAccountRelationship) {
	o.ApplianceAccount = &v
}

// GetCertificate returns the Certificate field value if set, zero value otherwise.
func (o *SmtpPolicyAllOf) GetCertificate() IamTrustPointRelationship {
	if o == nil || o.Certificate == nil {
		var ret IamTrustPointRelationship
		return ret
	}
	return *o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmtpPolicyAllOf) GetCertificateOk() (*IamTrustPointRelationship, bool) {
	if o == nil || o.Certificate == nil {
		return nil, false
	}
	return o.Certificate, true
}

// HasCertificate returns a boolean if a field has been set.
func (o *SmtpPolicyAllOf) HasCertificate() bool {
	if o != nil && o.Certificate != nil {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given IamTrustPointRelationship and assigns it to the Certificate field.
func (o *SmtpPolicyAllOf) SetCertificate(v IamTrustPointRelationship) {
	o.Certificate = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *SmtpPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmtpPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *SmtpPolicyAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *SmtpPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetProfiles returns the Profiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SmtpPolicyAllOf) GetProfiles() []PolicyAbstractConfigProfileRelationship {
	if o == nil {
		var ret []PolicyAbstractConfigProfileRelationship
		return ret
	}
	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SmtpPolicyAllOf) GetProfilesOk() ([]PolicyAbstractConfigProfileRelationship, bool) {
	if o == nil || o.Profiles == nil {
		return nil, false
	}
	return o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *SmtpPolicyAllOf) HasProfiles() bool {
	if o != nil && o.Profiles != nil {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []PolicyAbstractConfigProfileRelationship and assigns it to the Profiles field.
func (o *SmtpPolicyAllOf) SetProfiles(v []PolicyAbstractConfigProfileRelationship) {
	o.Profiles = v
}

func (o SmtpPolicyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AuthPassword != nil {
		toSerialize["AuthPassword"] = o.AuthPassword
	}
	if o.EnableAuth != nil {
		toSerialize["EnableAuth"] = o.EnableAuth
	}
	if o.EnableTls != nil {
		toSerialize["EnableTls"] = o.EnableTls
	}
	if o.Enabled != nil {
		toSerialize["Enabled"] = o.Enabled
	}
	if o.IsAuthPasswordSet != nil {
		toSerialize["IsAuthPasswordSet"] = o.IsAuthPasswordSet
	}
	if o.MinSeverity != nil {
		toSerialize["MinSeverity"] = o.MinSeverity
	}
	if o.SenderEmail != nil {
		toSerialize["SenderEmail"] = o.SenderEmail
	}
	if o.SmtpPort != nil {
		toSerialize["SmtpPort"] = o.SmtpPort
	}
	if o.SmtpRecipients != nil {
		toSerialize["SmtpRecipients"] = o.SmtpRecipients
	}
	if o.SmtpServer != nil {
		toSerialize["SmtpServer"] = o.SmtpServer
	}
	if o.UserName != nil {
		toSerialize["UserName"] = o.UserName
	}
	if o.ApplianceAccount != nil {
		toSerialize["ApplianceAccount"] = o.ApplianceAccount
	}
	if o.Certificate != nil {
		toSerialize["Certificate"] = o.Certificate
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.Profiles != nil {
		toSerialize["Profiles"] = o.Profiles
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SmtpPolicyAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varSmtpPolicyAllOf := _SmtpPolicyAllOf{}

	if err = json.Unmarshal(bytes, &varSmtpPolicyAllOf); err == nil {
		*o = SmtpPolicyAllOf(varSmtpPolicyAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AuthPassword")
		delete(additionalProperties, "EnableAuth")
		delete(additionalProperties, "EnableTls")
		delete(additionalProperties, "Enabled")
		delete(additionalProperties, "IsAuthPasswordSet")
		delete(additionalProperties, "MinSeverity")
		delete(additionalProperties, "SenderEmail")
		delete(additionalProperties, "SmtpPort")
		delete(additionalProperties, "SmtpRecipients")
		delete(additionalProperties, "SmtpServer")
		delete(additionalProperties, "UserName")
		delete(additionalProperties, "ApplianceAccount")
		delete(additionalProperties, "Certificate")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "Profiles")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSmtpPolicyAllOf struct {
	value *SmtpPolicyAllOf
	isSet bool
}

func (v NullableSmtpPolicyAllOf) Get() *SmtpPolicyAllOf {
	return v.value
}

func (v *NullableSmtpPolicyAllOf) Set(val *SmtpPolicyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSmtpPolicyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSmtpPolicyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmtpPolicyAllOf(val *SmtpPolicyAllOf) *NullableSmtpPolicyAllOf {
	return &NullableSmtpPolicyAllOf{value: val, isSet: true}
}

func (v NullableSmtpPolicyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmtpPolicyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
