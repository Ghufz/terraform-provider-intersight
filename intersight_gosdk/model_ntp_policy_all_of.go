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

// NtpPolicyAllOf Definition of the list of properties defined in 'ntp.Policy', excluding properties defined in parent classes.
type NtpPolicyAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType              string             `json:"ObjectType"`
	AuthenticatedNtpServers []NtpAuthNtpServer `json:"AuthenticatedNtpServers,omitempty"`
	// State of NTP service on the endpoint.
	Enabled    *bool    `json:"Enabled,omitempty"`
	NtpServers []string `json:"NtpServers,omitempty"`
	// Timezone of services on the endpoint. * `Pacific/Niue` -  * `Pacific/Pago_Pago` -  * `Pacific/Honolulu` -  * `Pacific/Rarotonga` -  * `Pacific/Tahiti` -  * `Pacific/Marquesas` -  * `America/Anchorage` -  * `Pacific/Gambier` -  * `America/Los_Angeles` -  * `America/Tijuana` -  * `America/Vancouver` -  * `America/Whitehorse` -  * `Pacific/Pitcairn` -  * `America/Dawson_Creek` -  * `America/Denver` -  * `America/Edmonton` -  * `America/Hermosillo` -  * `America/Mazatlan` -  * `America/Phoenix` -  * `America/Yellowknife` -  * `America/Belize` -  * `America/Chicago` -  * `America/Costa_Rica` -  * `America/El_Salvador` -  * `America/Guatemala` -  * `America/Managua` -  * `America/Mexico_City` -  * `America/Regina` -  * `America/Tegucigalpa` -  * `America/Winnipeg` -  * `Pacific/Galapagos` -  * `America/Bogota` -  * `America/Cancun` -  * `America/Cayman` -  * `America/Guayaquil` -  * `America/Havana` -  * `America/Iqaluit` -  * `America/Jamaica` -  * `America/Lima` -  * `America/Nassau` -  * `America/New_York` -  * `America/Nuuk` -  * `America/Panama` -  * `America/Port-au-Prince` -  * `America/Rio_Branco` -  * `America/Toronto` -  * `Pacific/Easter` -  * `America/Caracas` -  * `America/Asuncion` -  * `America/Barbados` -  * `America/Boa_Vista` -  * `America/Campo_Grande` -  * `America/Cuiaba` -  * `America/Curacao` -  * `America/Grand_Turk` -  * `America/Guyana` -  * `America/Halifax` -  * `America/La_Paz` -  * `America/Manaus` -  * `America/Martinique` -  * `America/Port_of_Spain` -  * `America/Porto_Velho` -  * `America/Puerto_Rico` -  * `America/Santo_Domingo` -  * `America/Thule` -  * `Atlantic/Bermuda` -  * `America/St_Johns` -  * `America/Araguaina` -  * `America/Argentina/Buenos_Aires` -  * `America/Bahia` -  * `America/Belem` -  * `America/Cayenne` -  * `America/Fortaleza` -  * `America/Godthab` -  * `America/Maceio` -  * `America/Miquelon` -  * `America/Montevideo` -  * `America/Paramaribo` -  * `America/Recife` -  * `America/Santiago` -  * `America/Sao_Paulo` -  * `Antarctica/Palmer` -  * `Antarctica/Rothera` -  * `Atlantic/Stanley` -  * `America/Noronha` -  * `Atlantic/South_Georgia` -  * `America/Scoresbysund` -  * `Atlantic/Azores` -  * `Atlantic/Cape_Verde` -  * `Africa/Abidjan` -  * `Africa/Accra` -  * `Africa/Bissau` -  * `Africa/Casablanca` -  * `Africa/El_Aaiun` -  * `Africa/Monrovia` -  * `America/Danmarkshavn` -  * `Atlantic/Canary` -  * `Atlantic/Faroe` -  * `Atlantic/Reykjavik` -  * `Etc/GMT` -  * `Europe/Dublin` -  * `Europe/Lisbon` -  * `Europe/London` -  * `Africa/Algiers` -  * `Africa/Ceuta` -  * `Africa/Lagos` -  * `Africa/Ndjamena` -  * `Africa/Tunis` -  * `Africa/Windhoek` -  * `Europe/Amsterdam` -  * `Europe/Andorra` -  * `Europe/Belgrade` -  * `Europe/Berlin` -  * `Europe/Brussels` -  * `Europe/Budapest` -  * `Europe/Copenhagen` -  * `Europe/Gibraltar` -  * `Europe/Luxembourg` -  * `Europe/Madrid` -  * `Europe/Malta` -  * `Europe/Monaco` -  * `Europe/Oslo` -  * `Europe/Paris` -  * `Europe/Prague` -  * `Europe/Rome` -  * `Europe/Stockholm` -  * `Europe/Tirane` -  * `Europe/Vienna` -  * `Europe/Warsaw` -  * `Europe/Zurich` -  * `Africa/Cairo` -  * `Africa/Johannesburg` -  * `Africa/Maputo` -  * `Africa/Tripoli` -  * `Asia/Amman` -  * `Asia/Beirut` -  * `Asia/Damascus` -  * `Asia/Gaza` -  * `Asia/Jerusalem` -  * `Asia/Nicosia` -  * `Europe/Athens` -  * `Europe/Bucharest` -  * `Europe/Chisinau` -  * `Europe/Helsinki` -  * `Europe/Istanbul` -  * `Europe/Kaliningrad` -  * `Europe/Kiev` -  * `Europe/Riga` -  * `Europe/Sofia` -  * `Europe/Tallinn` -  * `Europe/Vilnius` -  * `Africa/Khartoum` -  * `Africa/Nairobi` -  * `Antarctica/Syowa` -  * `Asia/Baghdad` -  * `Asia/Qatar` -  * `Asia/Riyadh` -  * `Europe/Minsk` -  * `Europe/Moscow` -  * `Asia/Tehran` -  * `Asia/Baku` -  * `Asia/Dubai` -  * `Asia/Tbilisi` -  * `Asia/Yerevan` -  * `Europe/Samara` -  * `Indian/Mahe` -  * `Indian/Mauritius` -  * `Indian/Reunion` -  * `Asia/Kabul` -  * `Antarctica/Mawson` -  * `Asia/Aqtau` -  * `Asia/Aqtobe` -  * `Asia/Ashgabat` -  * `Asia/Dushanbe` -  * `Asia/Karachi` -  * `Asia/Tashkent` -  * `Asia/Yekaterinburg` -  * `Indian/Kerguelen` -  * `Indian/Maldives` -  * `Asia/Calcutta` -  * `Asia/Kolkata` -  * `Asia/Colombo` -  * `Asia/Kathmandu` -  * `Asia/Katmandu` -  * `Antarctica/Vostok` -  * `Asia/Almaty` -  * `Asia/Bishkek` -  * `Asia/Dhaka` -  * `Asia/Omsk` -  * `Asia/Thimphu` -  * `Indian/Chagos` -  * `Asia/Rangoon` -  * `Indian/Cocos` -  * `Antarctica/Davis` -  * `Asia/Bangkok` -  * `Asia/Ho_Chi_Minh` -  * `Asia/Hovd` -  * `Asia/Jakarta` -  * `Asia/Krasnoyarsk` -  * `Asia/Saigon` -  * `Indian/Christmas` -  * `Antarctica/Casey` -  * `Asia/Brunei` -  * `Asia/Choibalsan` -  * `Asia/Hong_Kong` -  * `Asia/Irkutsk` -  * `Asia/Kuala_Lumpur` -  * `Asia/Macau` -  * `Asia/Makassar` -  * `Asia/Manila` -  * `Asia/Shanghai` -  * `Asia/Singapore` -  * `Asia/Taipei` -  * `Asia/Ulaanbaatar` -  * `Australia/Perth` -  * `Asia/Pyongyang` -  * `Asia/Dili` -  * `Asia/Jayapura` -  * `Asia/Seoul` -  * `Asia/Tokyo` -  * `Asia/Yakutsk` -  * `Asia/Yangon` -  * `Pacific/Palau` -  * `Australia/Adelaide` -  * `Australia/Darwin` -  * `Antarctica/DumontDUrville` -  * `Asia/Magadan` -  * `Asia/Vladivostok` -  * `Australia/Brisbane` -  * `Australia/Hobart` -  * `Australia/Sydney` -  * `Pacific/Chuuk` -  * `Pacific/Guam` -  * `Pacific/Port_Moresby` -  * `Pacific/Efate` -  * `Pacific/Guadalcanal` -  * `Pacific/Kosrae` -  * `Pacific/Norfolk` -  * `Pacific/Noumea` -  * `Pacific/Pohnpei` -  * `Asia/Kamchatka` -  * `Pacific/Auckland` -  * `Pacific/Fiji` -  * `Pacific/Funafuti` -  * `Pacific/Kwajalein` -  * `Pacific/Majuro` -  * `Pacific/Nauru` -  * `Pacific/Tarawa` -  * `Pacific/Wake` -  * `Pacific/Wallis` -  * `Pacific/Apia` -  * `Pacific/Enderbury` -  * `Pacific/Fakaofo` -  * `Pacific/Tongatapu` -  * `Pacific/Kiritimati` -  * `UTC` -
	Timezone         *string                               `json:"Timezone,omitempty"`
	ApplianceAccount *IamAccountRelationship               `json:"ApplianceAccount,omitempty"`
	Organization     *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	// An array of relationships to policyAbstractConfigProfile resources.
	Profiles             []PolicyAbstractConfigProfileRelationship `json:"Profiles,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NtpPolicyAllOf NtpPolicyAllOf

// NewNtpPolicyAllOf instantiates a new NtpPolicyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNtpPolicyAllOf(classId string, objectType string) *NtpPolicyAllOf {
	this := NtpPolicyAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var enabled bool = true
	this.Enabled = &enabled
	var timezone string = "Pacific/Niue"
	this.Timezone = &timezone
	return &this
}

// NewNtpPolicyAllOfWithDefaults instantiates a new NtpPolicyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNtpPolicyAllOfWithDefaults() *NtpPolicyAllOf {
	this := NtpPolicyAllOf{}
	var classId string = "ntp.Policy"
	this.ClassId = classId
	var objectType string = "ntp.Policy"
	this.ObjectType = objectType
	var enabled bool = true
	this.Enabled = &enabled
	var timezone string = "Pacific/Niue"
	this.Timezone = &timezone
	return &this
}

// GetClassId returns the ClassId field value
func (o *NtpPolicyAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NtpPolicyAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NtpPolicyAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NtpPolicyAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NtpPolicyAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NtpPolicyAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAuthenticatedNtpServers returns the AuthenticatedNtpServers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NtpPolicyAllOf) GetAuthenticatedNtpServers() []NtpAuthNtpServer {
	if o == nil {
		var ret []NtpAuthNtpServer
		return ret
	}
	return o.AuthenticatedNtpServers
}

// GetAuthenticatedNtpServersOk returns a tuple with the AuthenticatedNtpServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NtpPolicyAllOf) GetAuthenticatedNtpServersOk() ([]NtpAuthNtpServer, bool) {
	if o == nil || o.AuthenticatedNtpServers == nil {
		return nil, false
	}
	return o.AuthenticatedNtpServers, true
}

// HasAuthenticatedNtpServers returns a boolean if a field has been set.
func (o *NtpPolicyAllOf) HasAuthenticatedNtpServers() bool {
	if o != nil && o.AuthenticatedNtpServers != nil {
		return true
	}

	return false
}

// SetAuthenticatedNtpServers gets a reference to the given []NtpAuthNtpServer and assigns it to the AuthenticatedNtpServers field.
func (o *NtpPolicyAllOf) SetAuthenticatedNtpServers(v []NtpAuthNtpServer) {
	o.AuthenticatedNtpServers = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *NtpPolicyAllOf) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NtpPolicyAllOf) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *NtpPolicyAllOf) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *NtpPolicyAllOf) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetNtpServers returns the NtpServers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NtpPolicyAllOf) GetNtpServers() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.NtpServers
}

// GetNtpServersOk returns a tuple with the NtpServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NtpPolicyAllOf) GetNtpServersOk() ([]string, bool) {
	if o == nil || o.NtpServers == nil {
		return nil, false
	}
	return o.NtpServers, true
}

// HasNtpServers returns a boolean if a field has been set.
func (o *NtpPolicyAllOf) HasNtpServers() bool {
	if o != nil && o.NtpServers != nil {
		return true
	}

	return false
}

// SetNtpServers gets a reference to the given []string and assigns it to the NtpServers field.
func (o *NtpPolicyAllOf) SetNtpServers(v []string) {
	o.NtpServers = v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *NtpPolicyAllOf) GetTimezone() string {
	if o == nil || o.Timezone == nil {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NtpPolicyAllOf) GetTimezoneOk() (*string, bool) {
	if o == nil || o.Timezone == nil {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *NtpPolicyAllOf) HasTimezone() bool {
	if o != nil && o.Timezone != nil {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *NtpPolicyAllOf) SetTimezone(v string) {
	o.Timezone = &v
}

// GetApplianceAccount returns the ApplianceAccount field value if set, zero value otherwise.
func (o *NtpPolicyAllOf) GetApplianceAccount() IamAccountRelationship {
	if o == nil || o.ApplianceAccount == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.ApplianceAccount
}

// GetApplianceAccountOk returns a tuple with the ApplianceAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NtpPolicyAllOf) GetApplianceAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.ApplianceAccount == nil {
		return nil, false
	}
	return o.ApplianceAccount, true
}

// HasApplianceAccount returns a boolean if a field has been set.
func (o *NtpPolicyAllOf) HasApplianceAccount() bool {
	if o != nil && o.ApplianceAccount != nil {
		return true
	}

	return false
}

// SetApplianceAccount gets a reference to the given IamAccountRelationship and assigns it to the ApplianceAccount field.
func (o *NtpPolicyAllOf) SetApplianceAccount(v IamAccountRelationship) {
	o.ApplianceAccount = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *NtpPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NtpPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *NtpPolicyAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *NtpPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetProfiles returns the Profiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NtpPolicyAllOf) GetProfiles() []PolicyAbstractConfigProfileRelationship {
	if o == nil {
		var ret []PolicyAbstractConfigProfileRelationship
		return ret
	}
	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NtpPolicyAllOf) GetProfilesOk() ([]PolicyAbstractConfigProfileRelationship, bool) {
	if o == nil || o.Profiles == nil {
		return nil, false
	}
	return o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *NtpPolicyAllOf) HasProfiles() bool {
	if o != nil && o.Profiles != nil {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []PolicyAbstractConfigProfileRelationship and assigns it to the Profiles field.
func (o *NtpPolicyAllOf) SetProfiles(v []PolicyAbstractConfigProfileRelationship) {
	o.Profiles = v
}

func (o NtpPolicyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AuthenticatedNtpServers != nil {
		toSerialize["AuthenticatedNtpServers"] = o.AuthenticatedNtpServers
	}
	if o.Enabled != nil {
		toSerialize["Enabled"] = o.Enabled
	}
	if o.NtpServers != nil {
		toSerialize["NtpServers"] = o.NtpServers
	}
	if o.Timezone != nil {
		toSerialize["Timezone"] = o.Timezone
	}
	if o.ApplianceAccount != nil {
		toSerialize["ApplianceAccount"] = o.ApplianceAccount
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

func (o *NtpPolicyAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varNtpPolicyAllOf := _NtpPolicyAllOf{}

	if err = json.Unmarshal(bytes, &varNtpPolicyAllOf); err == nil {
		*o = NtpPolicyAllOf(varNtpPolicyAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AuthenticatedNtpServers")
		delete(additionalProperties, "Enabled")
		delete(additionalProperties, "NtpServers")
		delete(additionalProperties, "Timezone")
		delete(additionalProperties, "ApplianceAccount")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "Profiles")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNtpPolicyAllOf struct {
	value *NtpPolicyAllOf
	isSet bool
}

func (v NullableNtpPolicyAllOf) Get() *NtpPolicyAllOf {
	return v.value
}

func (v *NullableNtpPolicyAllOf) Set(val *NtpPolicyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNtpPolicyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNtpPolicyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNtpPolicyAllOf(val *NtpPolicyAllOf) *NullableNtpPolicyAllOf {
	return &NullableNtpPolicyAllOf{value: val, isSet: true}
}

func (v NullableNtpPolicyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNtpPolicyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
