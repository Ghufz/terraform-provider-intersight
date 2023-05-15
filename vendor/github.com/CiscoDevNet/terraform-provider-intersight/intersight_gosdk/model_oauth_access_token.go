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
	"time"
)

// OauthAccessToken Api access token for a given account.
type OauthAccessToken struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Type of OAuth Api. For example, Smart-licensing-API. * `Unknown` - Unknown is the default API type. * `SmartLicensing-API` - Smart licensing API type. * `CommerceEstimate-API` - Commerce Estimate API type.
	ApiType *string `json:"ApiType,omitempty"`
	// The date and time when the access token expires.
	Expiry *time.Time `json:"Expiry,omitempty"`
	// Issuer of OAuth access token.
	Issuer *string `json:"Issuer,omitempty"`
	// The date and time when the refresh token expires.
	RefreshExpiry        *time.Time              `json:"RefreshExpiry,omitempty"`
	Account              *IamAccountRelationship `json:"Account,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OauthAccessToken OauthAccessToken

// NewOauthAccessToken instantiates a new OauthAccessToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOauthAccessToken(classId string, objectType string) *OauthAccessToken {
	this := OauthAccessToken{}
	this.ClassId = classId
	this.ObjectType = objectType
	var apiType string = "Unknown"
	this.ApiType = &apiType
	return &this
}

// NewOauthAccessTokenWithDefaults instantiates a new OauthAccessToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOauthAccessTokenWithDefaults() *OauthAccessToken {
	this := OauthAccessToken{}
	var classId string = "oauth.AccessToken"
	this.ClassId = classId
	var objectType string = "oauth.AccessToken"
	this.ObjectType = objectType
	var apiType string = "Unknown"
	this.ApiType = &apiType
	return &this
}

// GetClassId returns the ClassId field value
func (o *OauthAccessToken) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *OauthAccessToken) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *OauthAccessToken) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *OauthAccessToken) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *OauthAccessToken) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *OauthAccessToken) SetObjectType(v string) {
	o.ObjectType = v
}

// GetApiType returns the ApiType field value if set, zero value otherwise.
func (o *OauthAccessToken) GetApiType() string {
	if o == nil || o.ApiType == nil {
		var ret string
		return ret
	}
	return *o.ApiType
}

// GetApiTypeOk returns a tuple with the ApiType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OauthAccessToken) GetApiTypeOk() (*string, bool) {
	if o == nil || o.ApiType == nil {
		return nil, false
	}
	return o.ApiType, true
}

// HasApiType returns a boolean if a field has been set.
func (o *OauthAccessToken) HasApiType() bool {
	if o != nil && o.ApiType != nil {
		return true
	}

	return false
}

// SetApiType gets a reference to the given string and assigns it to the ApiType field.
func (o *OauthAccessToken) SetApiType(v string) {
	o.ApiType = &v
}

// GetExpiry returns the Expiry field value if set, zero value otherwise.
func (o *OauthAccessToken) GetExpiry() time.Time {
	if o == nil || o.Expiry == nil {
		var ret time.Time
		return ret
	}
	return *o.Expiry
}

// GetExpiryOk returns a tuple with the Expiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OauthAccessToken) GetExpiryOk() (*time.Time, bool) {
	if o == nil || o.Expiry == nil {
		return nil, false
	}
	return o.Expiry, true
}

// HasExpiry returns a boolean if a field has been set.
func (o *OauthAccessToken) HasExpiry() bool {
	if o != nil && o.Expiry != nil {
		return true
	}

	return false
}

// SetExpiry gets a reference to the given time.Time and assigns it to the Expiry field.
func (o *OauthAccessToken) SetExpiry(v time.Time) {
	o.Expiry = &v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *OauthAccessToken) GetIssuer() string {
	if o == nil || o.Issuer == nil {
		var ret string
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OauthAccessToken) GetIssuerOk() (*string, bool) {
	if o == nil || o.Issuer == nil {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *OauthAccessToken) HasIssuer() bool {
	if o != nil && o.Issuer != nil {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given string and assigns it to the Issuer field.
func (o *OauthAccessToken) SetIssuer(v string) {
	o.Issuer = &v
}

// GetRefreshExpiry returns the RefreshExpiry field value if set, zero value otherwise.
func (o *OauthAccessToken) GetRefreshExpiry() time.Time {
	if o == nil || o.RefreshExpiry == nil {
		var ret time.Time
		return ret
	}
	return *o.RefreshExpiry
}

// GetRefreshExpiryOk returns a tuple with the RefreshExpiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OauthAccessToken) GetRefreshExpiryOk() (*time.Time, bool) {
	if o == nil || o.RefreshExpiry == nil {
		return nil, false
	}
	return o.RefreshExpiry, true
}

// HasRefreshExpiry returns a boolean if a field has been set.
func (o *OauthAccessToken) HasRefreshExpiry() bool {
	if o != nil && o.RefreshExpiry != nil {
		return true
	}

	return false
}

// SetRefreshExpiry gets a reference to the given time.Time and assigns it to the RefreshExpiry field.
func (o *OauthAccessToken) SetRefreshExpiry(v time.Time) {
	o.RefreshExpiry = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *OauthAccessToken) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OauthAccessToken) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *OauthAccessToken) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *OauthAccessToken) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

func (o OauthAccessToken) MarshalJSON() ([]byte, error) {
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
	if o.ApiType != nil {
		toSerialize["ApiType"] = o.ApiType
	}
	if o.Expiry != nil {
		toSerialize["Expiry"] = o.Expiry
	}
	if o.Issuer != nil {
		toSerialize["Issuer"] = o.Issuer
	}
	if o.RefreshExpiry != nil {
		toSerialize["RefreshExpiry"] = o.RefreshExpiry
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OauthAccessToken) UnmarshalJSON(bytes []byte) (err error) {
	type OauthAccessTokenWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Type of OAuth Api. For example, Smart-licensing-API. * `Unknown` - Unknown is the default API type. * `SmartLicensing-API` - Smart licensing API type. * `CommerceEstimate-API` - Commerce Estimate API type.
		ApiType *string `json:"ApiType,omitempty"`
		// The date and time when the access token expires.
		Expiry *time.Time `json:"Expiry,omitempty"`
		// Issuer of OAuth access token.
		Issuer *string `json:"Issuer,omitempty"`
		// The date and time when the refresh token expires.
		RefreshExpiry *time.Time              `json:"RefreshExpiry,omitempty"`
		Account       *IamAccountRelationship `json:"Account,omitempty"`
	}

	varOauthAccessTokenWithoutEmbeddedStruct := OauthAccessTokenWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varOauthAccessTokenWithoutEmbeddedStruct)
	if err == nil {
		varOauthAccessToken := _OauthAccessToken{}
		varOauthAccessToken.ClassId = varOauthAccessTokenWithoutEmbeddedStruct.ClassId
		varOauthAccessToken.ObjectType = varOauthAccessTokenWithoutEmbeddedStruct.ObjectType
		varOauthAccessToken.ApiType = varOauthAccessTokenWithoutEmbeddedStruct.ApiType
		varOauthAccessToken.Expiry = varOauthAccessTokenWithoutEmbeddedStruct.Expiry
		varOauthAccessToken.Issuer = varOauthAccessTokenWithoutEmbeddedStruct.Issuer
		varOauthAccessToken.RefreshExpiry = varOauthAccessTokenWithoutEmbeddedStruct.RefreshExpiry
		varOauthAccessToken.Account = varOauthAccessTokenWithoutEmbeddedStruct.Account
		*o = OauthAccessToken(varOauthAccessToken)
	} else {
		return err
	}

	varOauthAccessToken := _OauthAccessToken{}

	err = json.Unmarshal(bytes, &varOauthAccessToken)
	if err == nil {
		o.MoBaseMo = varOauthAccessToken.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ApiType")
		delete(additionalProperties, "Expiry")
		delete(additionalProperties, "Issuer")
		delete(additionalProperties, "RefreshExpiry")
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

type NullableOauthAccessToken struct {
	value *OauthAccessToken
	isSet bool
}

func (v NullableOauthAccessToken) Get() *OauthAccessToken {
	return v.value
}

func (v *NullableOauthAccessToken) Set(val *OauthAccessToken) {
	v.value = val
	v.isSet = true
}

func (v NullableOauthAccessToken) IsSet() bool {
	return v.isSet
}

func (v *NullableOauthAccessToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOauthAccessToken(val *OauthAccessToken) *NullableOauthAccessToken {
	return &NullableOauthAccessToken{value: val, isSet: true}
}

func (v NullableOauthAccessToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOauthAccessToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
