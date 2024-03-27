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

// AssetAlarmSummaryAllOf Definition of the list of properties defined in 'asset.AlarmSummary', excluding properties defined in parent classes.
type AssetAlarmSummaryAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The count of active alarms that have severity type Critical.
	Critical *int64 `json:"Critical,omitempty"`
	// Health of the managed end point. The highest severity computed from alarmSummary property is set as the health. * `Healthy` - The Enum value represents that the entity is healthy. * `Warning` - The Enum value Warning represents that the entity has one or more active warnings on it. * `Critical` - The Enum value Critical represents that the entity is in a critical state.
	Health *string `json:"Health,omitempty"`
	// The count of active alarms that have severity type Info.
	Info *int64 `json:"Info,omitempty"`
	// The count of active suppressed alarms that have severity type Critical.
	SuppressedCritical *int64 `json:"SuppressedCritical,omitempty"`
	// The count of active suppressed alarms that have severity type Info.
	SuppressedInfo *int64 `json:"SuppressedInfo,omitempty"`
	// The count of active suppressed alarms that have severity type Warning.
	SuppressedWarning *int64 `json:"SuppressedWarning,omitempty"`
	// The count of active alarms that have severity type Warning.
	Warning              *int64 `json:"Warning,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetAlarmSummaryAllOf AssetAlarmSummaryAllOf

// NewAssetAlarmSummaryAllOf instantiates a new AssetAlarmSummaryAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetAlarmSummaryAllOf(classId string, objectType string) *AssetAlarmSummaryAllOf {
	this := AssetAlarmSummaryAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAssetAlarmSummaryAllOfWithDefaults instantiates a new AssetAlarmSummaryAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetAlarmSummaryAllOfWithDefaults() *AssetAlarmSummaryAllOf {
	this := AssetAlarmSummaryAllOf{}
	var classId string = "asset.AlarmSummary"
	this.ClassId = classId
	var objectType string = "asset.AlarmSummary"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetAlarmSummaryAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetAlarmSummaryAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetAlarmSummaryAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AssetAlarmSummaryAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetAlarmSummaryAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetAlarmSummaryAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCritical returns the Critical field value if set, zero value otherwise.
func (o *AssetAlarmSummaryAllOf) GetCritical() int64 {
	if o == nil || o.Critical == nil {
		var ret int64
		return ret
	}
	return *o.Critical
}

// GetCriticalOk returns a tuple with the Critical field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetAlarmSummaryAllOf) GetCriticalOk() (*int64, bool) {
	if o == nil || o.Critical == nil {
		return nil, false
	}
	return o.Critical, true
}

// HasCritical returns a boolean if a field has been set.
func (o *AssetAlarmSummaryAllOf) HasCritical() bool {
	if o != nil && o.Critical != nil {
		return true
	}

	return false
}

// SetCritical gets a reference to the given int64 and assigns it to the Critical field.
func (o *AssetAlarmSummaryAllOf) SetCritical(v int64) {
	o.Critical = &v
}

// GetHealth returns the Health field value if set, zero value otherwise.
func (o *AssetAlarmSummaryAllOf) GetHealth() string {
	if o == nil || o.Health == nil {
		var ret string
		return ret
	}
	return *o.Health
}

// GetHealthOk returns a tuple with the Health field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetAlarmSummaryAllOf) GetHealthOk() (*string, bool) {
	if o == nil || o.Health == nil {
		return nil, false
	}
	return o.Health, true
}

// HasHealth returns a boolean if a field has been set.
func (o *AssetAlarmSummaryAllOf) HasHealth() bool {
	if o != nil && o.Health != nil {
		return true
	}

	return false
}

// SetHealth gets a reference to the given string and assigns it to the Health field.
func (o *AssetAlarmSummaryAllOf) SetHealth(v string) {
	o.Health = &v
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *AssetAlarmSummaryAllOf) GetInfo() int64 {
	if o == nil || o.Info == nil {
		var ret int64
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetAlarmSummaryAllOf) GetInfoOk() (*int64, bool) {
	if o == nil || o.Info == nil {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *AssetAlarmSummaryAllOf) HasInfo() bool {
	if o != nil && o.Info != nil {
		return true
	}

	return false
}

// SetInfo gets a reference to the given int64 and assigns it to the Info field.
func (o *AssetAlarmSummaryAllOf) SetInfo(v int64) {
	o.Info = &v
}

// GetSuppressedCritical returns the SuppressedCritical field value if set, zero value otherwise.
func (o *AssetAlarmSummaryAllOf) GetSuppressedCritical() int64 {
	if o == nil || o.SuppressedCritical == nil {
		var ret int64
		return ret
	}
	return *o.SuppressedCritical
}

// GetSuppressedCriticalOk returns a tuple with the SuppressedCritical field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetAlarmSummaryAllOf) GetSuppressedCriticalOk() (*int64, bool) {
	if o == nil || o.SuppressedCritical == nil {
		return nil, false
	}
	return o.SuppressedCritical, true
}

// HasSuppressedCritical returns a boolean if a field has been set.
func (o *AssetAlarmSummaryAllOf) HasSuppressedCritical() bool {
	if o != nil && o.SuppressedCritical != nil {
		return true
	}

	return false
}

// SetSuppressedCritical gets a reference to the given int64 and assigns it to the SuppressedCritical field.
func (o *AssetAlarmSummaryAllOf) SetSuppressedCritical(v int64) {
	o.SuppressedCritical = &v
}

// GetSuppressedInfo returns the SuppressedInfo field value if set, zero value otherwise.
func (o *AssetAlarmSummaryAllOf) GetSuppressedInfo() int64 {
	if o == nil || o.SuppressedInfo == nil {
		var ret int64
		return ret
	}
	return *o.SuppressedInfo
}

// GetSuppressedInfoOk returns a tuple with the SuppressedInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetAlarmSummaryAllOf) GetSuppressedInfoOk() (*int64, bool) {
	if o == nil || o.SuppressedInfo == nil {
		return nil, false
	}
	return o.SuppressedInfo, true
}

// HasSuppressedInfo returns a boolean if a field has been set.
func (o *AssetAlarmSummaryAllOf) HasSuppressedInfo() bool {
	if o != nil && o.SuppressedInfo != nil {
		return true
	}

	return false
}

// SetSuppressedInfo gets a reference to the given int64 and assigns it to the SuppressedInfo field.
func (o *AssetAlarmSummaryAllOf) SetSuppressedInfo(v int64) {
	o.SuppressedInfo = &v
}

// GetSuppressedWarning returns the SuppressedWarning field value if set, zero value otherwise.
func (o *AssetAlarmSummaryAllOf) GetSuppressedWarning() int64 {
	if o == nil || o.SuppressedWarning == nil {
		var ret int64
		return ret
	}
	return *o.SuppressedWarning
}

// GetSuppressedWarningOk returns a tuple with the SuppressedWarning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetAlarmSummaryAllOf) GetSuppressedWarningOk() (*int64, bool) {
	if o == nil || o.SuppressedWarning == nil {
		return nil, false
	}
	return o.SuppressedWarning, true
}

// HasSuppressedWarning returns a boolean if a field has been set.
func (o *AssetAlarmSummaryAllOf) HasSuppressedWarning() bool {
	if o != nil && o.SuppressedWarning != nil {
		return true
	}

	return false
}

// SetSuppressedWarning gets a reference to the given int64 and assigns it to the SuppressedWarning field.
func (o *AssetAlarmSummaryAllOf) SetSuppressedWarning(v int64) {
	o.SuppressedWarning = &v
}

// GetWarning returns the Warning field value if set, zero value otherwise.
func (o *AssetAlarmSummaryAllOf) GetWarning() int64 {
	if o == nil || o.Warning == nil {
		var ret int64
		return ret
	}
	return *o.Warning
}

// GetWarningOk returns a tuple with the Warning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetAlarmSummaryAllOf) GetWarningOk() (*int64, bool) {
	if o == nil || o.Warning == nil {
		return nil, false
	}
	return o.Warning, true
}

// HasWarning returns a boolean if a field has been set.
func (o *AssetAlarmSummaryAllOf) HasWarning() bool {
	if o != nil && o.Warning != nil {
		return true
	}

	return false
}

// SetWarning gets a reference to the given int64 and assigns it to the Warning field.
func (o *AssetAlarmSummaryAllOf) SetWarning(v int64) {
	o.Warning = &v
}

func (o AssetAlarmSummaryAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Critical != nil {
		toSerialize["Critical"] = o.Critical
	}
	if o.Health != nil {
		toSerialize["Health"] = o.Health
	}
	if o.Info != nil {
		toSerialize["Info"] = o.Info
	}
	if o.SuppressedCritical != nil {
		toSerialize["SuppressedCritical"] = o.SuppressedCritical
	}
	if o.SuppressedInfo != nil {
		toSerialize["SuppressedInfo"] = o.SuppressedInfo
	}
	if o.SuppressedWarning != nil {
		toSerialize["SuppressedWarning"] = o.SuppressedWarning
	}
	if o.Warning != nil {
		toSerialize["Warning"] = o.Warning
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetAlarmSummaryAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varAssetAlarmSummaryAllOf := _AssetAlarmSummaryAllOf{}

	if err = json.Unmarshal(bytes, &varAssetAlarmSummaryAllOf); err == nil {
		*o = AssetAlarmSummaryAllOf(varAssetAlarmSummaryAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Critical")
		delete(additionalProperties, "Health")
		delete(additionalProperties, "Info")
		delete(additionalProperties, "SuppressedCritical")
		delete(additionalProperties, "SuppressedInfo")
		delete(additionalProperties, "SuppressedWarning")
		delete(additionalProperties, "Warning")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetAlarmSummaryAllOf struct {
	value *AssetAlarmSummaryAllOf
	isSet bool
}

func (v NullableAssetAlarmSummaryAllOf) Get() *AssetAlarmSummaryAllOf {
	return v.value
}

func (v *NullableAssetAlarmSummaryAllOf) Set(val *AssetAlarmSummaryAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetAlarmSummaryAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetAlarmSummaryAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetAlarmSummaryAllOf(val *AssetAlarmSummaryAllOf) *NullableAssetAlarmSummaryAllOf {
	return &NullableAssetAlarmSummaryAllOf{value: val, isSet: true}
}

func (v NullableAssetAlarmSummaryAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetAlarmSummaryAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
