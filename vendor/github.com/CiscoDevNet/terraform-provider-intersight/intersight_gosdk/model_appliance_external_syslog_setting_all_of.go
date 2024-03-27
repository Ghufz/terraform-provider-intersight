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

// ApplianceExternalSyslogSettingAllOf Definition of the list of properties defined in 'appliance.ExternalSyslogSetting', excluding properties defined in parent classes.
type ApplianceExternalSyslogSettingAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Enable or disable External Syslog Server.
	Enabled *bool `json:"Enabled,omitempty"`
	// If the flag is set, the alarms reported in Intersight Appliances are exported to external syslog server based on the alarm severity selection.
	ExportAlarms *bool `json:"ExportAlarms,omitempty"`
	// Enable or disable exporting of Audit logs.
	ExportAudit *bool `json:"ExportAudit,omitempty"`
	// Enable or disable exporting of Web Server access logs.
	ExportNginx *bool `json:"ExportNginx,omitempty"`
	// External Syslog Server Port for connection establishment.
	Port *int64 `json:"Port,omitempty"`
	// Protocol used to connect to external syslog server. * `TCP` - External Syslog messages sent over TCP. * `UDP` - External Syslog messages sent over UDP. * `TLS` - Secure External Syslog messages sent over TLS.
	Protocol *string `json:"Protocol,omitempty"`
	// External Syslog Server Address, can be IP address or hostname.
	Server *string `json:"Server,omitempty"`
	// Minimum severity level to report. * `None` - The Enum value None represents that there is no severity. * `Info` - The Enum value Info represents the Informational level of severity. * `Critical` - The Enum value Critical represents the Critical level of severity. * `Warning` - The Enum value Warning represents the Warning level of severity. * `Cleared` - The Enum value Cleared represents that the alarm severity has been cleared.
	Severity             *string                 `json:"Severity,omitempty"`
	Account              *IamAccountRelationship `json:"Account,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplianceExternalSyslogSettingAllOf ApplianceExternalSyslogSettingAllOf

// NewApplianceExternalSyslogSettingAllOf instantiates a new ApplianceExternalSyslogSettingAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceExternalSyslogSettingAllOf(classId string, objectType string) *ApplianceExternalSyslogSettingAllOf {
	this := ApplianceExternalSyslogSettingAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var enabled bool = false
	this.Enabled = &enabled
	var exportAlarms bool = false
	this.ExportAlarms = &exportAlarms
	var exportAudit bool = false
	this.ExportAudit = &exportAudit
	var exportNginx bool = false
	this.ExportNginx = &exportNginx
	var port int64 = 10514
	this.Port = &port
	var protocol string = "TCP"
	this.Protocol = &protocol
	var severity string = "None"
	this.Severity = &severity
	return &this
}

// NewApplianceExternalSyslogSettingAllOfWithDefaults instantiates a new ApplianceExternalSyslogSettingAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceExternalSyslogSettingAllOfWithDefaults() *ApplianceExternalSyslogSettingAllOf {
	this := ApplianceExternalSyslogSettingAllOf{}
	var classId string = "appliance.ExternalSyslogSetting"
	this.ClassId = classId
	var objectType string = "appliance.ExternalSyslogSetting"
	this.ObjectType = objectType
	var enabled bool = false
	this.Enabled = &enabled
	var exportAlarms bool = false
	this.ExportAlarms = &exportAlarms
	var exportAudit bool = false
	this.ExportAudit = &exportAudit
	var exportNginx bool = false
	this.ExportNginx = &exportNginx
	var port int64 = 10514
	this.Port = &port
	var protocol string = "TCP"
	this.Protocol = &protocol
	var severity string = "None"
	this.Severity = &severity
	return &this
}

// GetClassId returns the ClassId field value
func (o *ApplianceExternalSyslogSettingAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ApplianceExternalSyslogSettingAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ApplianceExternalSyslogSettingAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ApplianceExternalSyslogSettingAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ApplianceExternalSyslogSettingAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ApplianceExternalSyslogSettingAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ApplianceExternalSyslogSettingAllOf) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceExternalSyslogSettingAllOf) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ApplianceExternalSyslogSettingAllOf) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ApplianceExternalSyslogSettingAllOf) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetExportAlarms returns the ExportAlarms field value if set, zero value otherwise.
func (o *ApplianceExternalSyslogSettingAllOf) GetExportAlarms() bool {
	if o == nil || o.ExportAlarms == nil {
		var ret bool
		return ret
	}
	return *o.ExportAlarms
}

// GetExportAlarmsOk returns a tuple with the ExportAlarms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceExternalSyslogSettingAllOf) GetExportAlarmsOk() (*bool, bool) {
	if o == nil || o.ExportAlarms == nil {
		return nil, false
	}
	return o.ExportAlarms, true
}

// HasExportAlarms returns a boolean if a field has been set.
func (o *ApplianceExternalSyslogSettingAllOf) HasExportAlarms() bool {
	if o != nil && o.ExportAlarms != nil {
		return true
	}

	return false
}

// SetExportAlarms gets a reference to the given bool and assigns it to the ExportAlarms field.
func (o *ApplianceExternalSyslogSettingAllOf) SetExportAlarms(v bool) {
	o.ExportAlarms = &v
}

// GetExportAudit returns the ExportAudit field value if set, zero value otherwise.
func (o *ApplianceExternalSyslogSettingAllOf) GetExportAudit() bool {
	if o == nil || o.ExportAudit == nil {
		var ret bool
		return ret
	}
	return *o.ExportAudit
}

// GetExportAuditOk returns a tuple with the ExportAudit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceExternalSyslogSettingAllOf) GetExportAuditOk() (*bool, bool) {
	if o == nil || o.ExportAudit == nil {
		return nil, false
	}
	return o.ExportAudit, true
}

// HasExportAudit returns a boolean if a field has been set.
func (o *ApplianceExternalSyslogSettingAllOf) HasExportAudit() bool {
	if o != nil && o.ExportAudit != nil {
		return true
	}

	return false
}

// SetExportAudit gets a reference to the given bool and assigns it to the ExportAudit field.
func (o *ApplianceExternalSyslogSettingAllOf) SetExportAudit(v bool) {
	o.ExportAudit = &v
}

// GetExportNginx returns the ExportNginx field value if set, zero value otherwise.
func (o *ApplianceExternalSyslogSettingAllOf) GetExportNginx() bool {
	if o == nil || o.ExportNginx == nil {
		var ret bool
		return ret
	}
	return *o.ExportNginx
}

// GetExportNginxOk returns a tuple with the ExportNginx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceExternalSyslogSettingAllOf) GetExportNginxOk() (*bool, bool) {
	if o == nil || o.ExportNginx == nil {
		return nil, false
	}
	return o.ExportNginx, true
}

// HasExportNginx returns a boolean if a field has been set.
func (o *ApplianceExternalSyslogSettingAllOf) HasExportNginx() bool {
	if o != nil && o.ExportNginx != nil {
		return true
	}

	return false
}

// SetExportNginx gets a reference to the given bool and assigns it to the ExportNginx field.
func (o *ApplianceExternalSyslogSettingAllOf) SetExportNginx(v bool) {
	o.ExportNginx = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *ApplianceExternalSyslogSettingAllOf) GetPort() int64 {
	if o == nil || o.Port == nil {
		var ret int64
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceExternalSyslogSettingAllOf) GetPortOk() (*int64, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *ApplianceExternalSyslogSettingAllOf) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int64 and assigns it to the Port field.
func (o *ApplianceExternalSyslogSettingAllOf) SetPort(v int64) {
	o.Port = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *ApplianceExternalSyslogSettingAllOf) GetProtocol() string {
	if o == nil || o.Protocol == nil {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceExternalSyslogSettingAllOf) GetProtocolOk() (*string, bool) {
	if o == nil || o.Protocol == nil {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *ApplianceExternalSyslogSettingAllOf) HasProtocol() bool {
	if o != nil && o.Protocol != nil {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *ApplianceExternalSyslogSettingAllOf) SetProtocol(v string) {
	o.Protocol = &v
}

// GetServer returns the Server field value if set, zero value otherwise.
func (o *ApplianceExternalSyslogSettingAllOf) GetServer() string {
	if o == nil || o.Server == nil {
		var ret string
		return ret
	}
	return *o.Server
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceExternalSyslogSettingAllOf) GetServerOk() (*string, bool) {
	if o == nil || o.Server == nil {
		return nil, false
	}
	return o.Server, true
}

// HasServer returns a boolean if a field has been set.
func (o *ApplianceExternalSyslogSettingAllOf) HasServer() bool {
	if o != nil && o.Server != nil {
		return true
	}

	return false
}

// SetServer gets a reference to the given string and assigns it to the Server field.
func (o *ApplianceExternalSyslogSettingAllOf) SetServer(v string) {
	o.Server = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *ApplianceExternalSyslogSettingAllOf) GetSeverity() string {
	if o == nil || o.Severity == nil {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceExternalSyslogSettingAllOf) GetSeverityOk() (*string, bool) {
	if o == nil || o.Severity == nil {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *ApplianceExternalSyslogSettingAllOf) HasSeverity() bool {
	if o != nil && o.Severity != nil {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *ApplianceExternalSyslogSettingAllOf) SetSeverity(v string) {
	o.Severity = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *ApplianceExternalSyslogSettingAllOf) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceExternalSyslogSettingAllOf) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *ApplianceExternalSyslogSettingAllOf) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *ApplianceExternalSyslogSettingAllOf) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

func (o ApplianceExternalSyslogSettingAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Enabled != nil {
		toSerialize["Enabled"] = o.Enabled
	}
	if o.ExportAlarms != nil {
		toSerialize["ExportAlarms"] = o.ExportAlarms
	}
	if o.ExportAudit != nil {
		toSerialize["ExportAudit"] = o.ExportAudit
	}
	if o.ExportNginx != nil {
		toSerialize["ExportNginx"] = o.ExportNginx
	}
	if o.Port != nil {
		toSerialize["Port"] = o.Port
	}
	if o.Protocol != nil {
		toSerialize["Protocol"] = o.Protocol
	}
	if o.Server != nil {
		toSerialize["Server"] = o.Server
	}
	if o.Severity != nil {
		toSerialize["Severity"] = o.Severity
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ApplianceExternalSyslogSettingAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varApplianceExternalSyslogSettingAllOf := _ApplianceExternalSyslogSettingAllOf{}

	if err = json.Unmarshal(bytes, &varApplianceExternalSyslogSettingAllOf); err == nil {
		*o = ApplianceExternalSyslogSettingAllOf(varApplianceExternalSyslogSettingAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Enabled")
		delete(additionalProperties, "ExportAlarms")
		delete(additionalProperties, "ExportAudit")
		delete(additionalProperties, "ExportNginx")
		delete(additionalProperties, "Port")
		delete(additionalProperties, "Protocol")
		delete(additionalProperties, "Server")
		delete(additionalProperties, "Severity")
		delete(additionalProperties, "Account")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApplianceExternalSyslogSettingAllOf struct {
	value *ApplianceExternalSyslogSettingAllOf
	isSet bool
}

func (v NullableApplianceExternalSyslogSettingAllOf) Get() *ApplianceExternalSyslogSettingAllOf {
	return v.value
}

func (v *NullableApplianceExternalSyslogSettingAllOf) Set(val *ApplianceExternalSyslogSettingAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceExternalSyslogSettingAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceExternalSyslogSettingAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceExternalSyslogSettingAllOf(val *ApplianceExternalSyslogSettingAllOf) *NullableApplianceExternalSyslogSettingAllOf {
	return &NullableApplianceExternalSyslogSettingAllOf{value: val, isSet: true}
}

func (v NullableApplianceExternalSyslogSettingAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceExternalSyslogSettingAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
