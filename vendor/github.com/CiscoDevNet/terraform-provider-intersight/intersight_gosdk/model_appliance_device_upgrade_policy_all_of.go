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
	"time"
)

// ApplianceDeviceUpgradePolicyAllOf Definition of the list of properties defined in 'appliance.DeviceUpgradePolicy', excluding properties defined in parent classes.
type ApplianceDeviceUpgradePolicyAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Indicates if the upgrade service is set to automatically start the software upgrade or not. If autoUpgrade is true, then the value of the schedule field is ignored.
	AutoUpgrade *bool `json:"AutoUpgrade,omitempty"`
	// If enabled, allows the user to define a blackout period during which the appliance will not be upgraded.
	BlackoutDatesEnabled *bool `json:"BlackoutDatesEnabled,omitempty"`
	// End date of the black out period.
	BlackoutEndDate *time.Time `json:"BlackoutEndDate,omitempty"`
	// Start date of the black out period. The appliance will not be upgraded during this period.
	BlackoutStartDate *time.Time `json:"BlackoutStartDate,omitempty"`
	// Indicates if the updated metadata files should be synced immediately or at the next upgrade.
	EnableMetaDataSync *bool                  `json:"EnableMetaDataSync,omitempty"`
	Schedule           NullableOnpremSchedule `json:"Schedule,omitempty"`
	// SerialId of the Intersight Appliance. SerialId is generated when the Intersight Appliance is setup. It is a unique UUID string, and serialId will not change for the life time of the Intersight Appliance.
	SerialId *string `json:"SerialId,omitempty"`
	// UpgradeType is used to indicate the kink of software upload to upgrade. * `unknown` - Indicates user setting of upgrade service to unknown. * `connected` - Indicates if the upgrade service is set to upload software to latest version automatically. * `manual` - Indicates if the upgrade service is set to upload software to user picked verison manually.
	SoftwareDownloadType *string                              `json:"SoftwareDownloadType,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplianceDeviceUpgradePolicyAllOf ApplianceDeviceUpgradePolicyAllOf

// NewApplianceDeviceUpgradePolicyAllOf instantiates a new ApplianceDeviceUpgradePolicyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceDeviceUpgradePolicyAllOf(classId string, objectType string) *ApplianceDeviceUpgradePolicyAllOf {
	this := ApplianceDeviceUpgradePolicyAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var enableMetaDataSync bool = true
	this.EnableMetaDataSync = &enableMetaDataSync
	var softwareDownloadType string = "unknown"
	this.SoftwareDownloadType = &softwareDownloadType
	return &this
}

// NewApplianceDeviceUpgradePolicyAllOfWithDefaults instantiates a new ApplianceDeviceUpgradePolicyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceDeviceUpgradePolicyAllOfWithDefaults() *ApplianceDeviceUpgradePolicyAllOf {
	this := ApplianceDeviceUpgradePolicyAllOf{}
	var classId string = "appliance.DeviceUpgradePolicy"
	this.ClassId = classId
	var objectType string = "appliance.DeviceUpgradePolicy"
	this.ObjectType = objectType
	var enableMetaDataSync bool = true
	this.EnableMetaDataSync = &enableMetaDataSync
	var softwareDownloadType string = "unknown"
	this.SoftwareDownloadType = &softwareDownloadType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ApplianceDeviceUpgradePolicyAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ApplianceDeviceUpgradePolicyAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ApplianceDeviceUpgradePolicyAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ApplianceDeviceUpgradePolicyAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ApplianceDeviceUpgradePolicyAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ApplianceDeviceUpgradePolicyAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAutoUpgrade returns the AutoUpgrade field value if set, zero value otherwise.
func (o *ApplianceDeviceUpgradePolicyAllOf) GetAutoUpgrade() bool {
	if o == nil || o.AutoUpgrade == nil {
		var ret bool
		return ret
	}
	return *o.AutoUpgrade
}

// GetAutoUpgradeOk returns a tuple with the AutoUpgrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceDeviceUpgradePolicyAllOf) GetAutoUpgradeOk() (*bool, bool) {
	if o == nil || o.AutoUpgrade == nil {
		return nil, false
	}
	return o.AutoUpgrade, true
}

// HasAutoUpgrade returns a boolean if a field has been set.
func (o *ApplianceDeviceUpgradePolicyAllOf) HasAutoUpgrade() bool {
	if o != nil && o.AutoUpgrade != nil {
		return true
	}

	return false
}

// SetAutoUpgrade gets a reference to the given bool and assigns it to the AutoUpgrade field.
func (o *ApplianceDeviceUpgradePolicyAllOf) SetAutoUpgrade(v bool) {
	o.AutoUpgrade = &v
}

// GetBlackoutDatesEnabled returns the BlackoutDatesEnabled field value if set, zero value otherwise.
func (o *ApplianceDeviceUpgradePolicyAllOf) GetBlackoutDatesEnabled() bool {
	if o == nil || o.BlackoutDatesEnabled == nil {
		var ret bool
		return ret
	}
	return *o.BlackoutDatesEnabled
}

// GetBlackoutDatesEnabledOk returns a tuple with the BlackoutDatesEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceDeviceUpgradePolicyAllOf) GetBlackoutDatesEnabledOk() (*bool, bool) {
	if o == nil || o.BlackoutDatesEnabled == nil {
		return nil, false
	}
	return o.BlackoutDatesEnabled, true
}

// HasBlackoutDatesEnabled returns a boolean if a field has been set.
func (o *ApplianceDeviceUpgradePolicyAllOf) HasBlackoutDatesEnabled() bool {
	if o != nil && o.BlackoutDatesEnabled != nil {
		return true
	}

	return false
}

// SetBlackoutDatesEnabled gets a reference to the given bool and assigns it to the BlackoutDatesEnabled field.
func (o *ApplianceDeviceUpgradePolicyAllOf) SetBlackoutDatesEnabled(v bool) {
	o.BlackoutDatesEnabled = &v
}

// GetBlackoutEndDate returns the BlackoutEndDate field value if set, zero value otherwise.
func (o *ApplianceDeviceUpgradePolicyAllOf) GetBlackoutEndDate() time.Time {
	if o == nil || o.BlackoutEndDate == nil {
		var ret time.Time
		return ret
	}
	return *o.BlackoutEndDate
}

// GetBlackoutEndDateOk returns a tuple with the BlackoutEndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceDeviceUpgradePolicyAllOf) GetBlackoutEndDateOk() (*time.Time, bool) {
	if o == nil || o.BlackoutEndDate == nil {
		return nil, false
	}
	return o.BlackoutEndDate, true
}

// HasBlackoutEndDate returns a boolean if a field has been set.
func (o *ApplianceDeviceUpgradePolicyAllOf) HasBlackoutEndDate() bool {
	if o != nil && o.BlackoutEndDate != nil {
		return true
	}

	return false
}

// SetBlackoutEndDate gets a reference to the given time.Time and assigns it to the BlackoutEndDate field.
func (o *ApplianceDeviceUpgradePolicyAllOf) SetBlackoutEndDate(v time.Time) {
	o.BlackoutEndDate = &v
}

// GetBlackoutStartDate returns the BlackoutStartDate field value if set, zero value otherwise.
func (o *ApplianceDeviceUpgradePolicyAllOf) GetBlackoutStartDate() time.Time {
	if o == nil || o.BlackoutStartDate == nil {
		var ret time.Time
		return ret
	}
	return *o.BlackoutStartDate
}

// GetBlackoutStartDateOk returns a tuple with the BlackoutStartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceDeviceUpgradePolicyAllOf) GetBlackoutStartDateOk() (*time.Time, bool) {
	if o == nil || o.BlackoutStartDate == nil {
		return nil, false
	}
	return o.BlackoutStartDate, true
}

// HasBlackoutStartDate returns a boolean if a field has been set.
func (o *ApplianceDeviceUpgradePolicyAllOf) HasBlackoutStartDate() bool {
	if o != nil && o.BlackoutStartDate != nil {
		return true
	}

	return false
}

// SetBlackoutStartDate gets a reference to the given time.Time and assigns it to the BlackoutStartDate field.
func (o *ApplianceDeviceUpgradePolicyAllOf) SetBlackoutStartDate(v time.Time) {
	o.BlackoutStartDate = &v
}

// GetEnableMetaDataSync returns the EnableMetaDataSync field value if set, zero value otherwise.
func (o *ApplianceDeviceUpgradePolicyAllOf) GetEnableMetaDataSync() bool {
	if o == nil || o.EnableMetaDataSync == nil {
		var ret bool
		return ret
	}
	return *o.EnableMetaDataSync
}

// GetEnableMetaDataSyncOk returns a tuple with the EnableMetaDataSync field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceDeviceUpgradePolicyAllOf) GetEnableMetaDataSyncOk() (*bool, bool) {
	if o == nil || o.EnableMetaDataSync == nil {
		return nil, false
	}
	return o.EnableMetaDataSync, true
}

// HasEnableMetaDataSync returns a boolean if a field has been set.
func (o *ApplianceDeviceUpgradePolicyAllOf) HasEnableMetaDataSync() bool {
	if o != nil && o.EnableMetaDataSync != nil {
		return true
	}

	return false
}

// SetEnableMetaDataSync gets a reference to the given bool and assigns it to the EnableMetaDataSync field.
func (o *ApplianceDeviceUpgradePolicyAllOf) SetEnableMetaDataSync(v bool) {
	o.EnableMetaDataSync = &v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplianceDeviceUpgradePolicyAllOf) GetSchedule() OnpremSchedule {
	if o == nil || o.Schedule.Get() == nil {
		var ret OnpremSchedule
		return ret
	}
	return *o.Schedule.Get()
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplianceDeviceUpgradePolicyAllOf) GetScheduleOk() (*OnpremSchedule, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schedule.Get(), o.Schedule.IsSet()
}

// HasSchedule returns a boolean if a field has been set.
func (o *ApplianceDeviceUpgradePolicyAllOf) HasSchedule() bool {
	if o != nil && o.Schedule.IsSet() {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given NullableOnpremSchedule and assigns it to the Schedule field.
func (o *ApplianceDeviceUpgradePolicyAllOf) SetSchedule(v OnpremSchedule) {
	o.Schedule.Set(&v)
}

// SetScheduleNil sets the value for Schedule to be an explicit nil
func (o *ApplianceDeviceUpgradePolicyAllOf) SetScheduleNil() {
	o.Schedule.Set(nil)
}

// UnsetSchedule ensures that no value is present for Schedule, not even an explicit nil
func (o *ApplianceDeviceUpgradePolicyAllOf) UnsetSchedule() {
	o.Schedule.Unset()
}

// GetSerialId returns the SerialId field value if set, zero value otherwise.
func (o *ApplianceDeviceUpgradePolicyAllOf) GetSerialId() string {
	if o == nil || o.SerialId == nil {
		var ret string
		return ret
	}
	return *o.SerialId
}

// GetSerialIdOk returns a tuple with the SerialId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceDeviceUpgradePolicyAllOf) GetSerialIdOk() (*string, bool) {
	if o == nil || o.SerialId == nil {
		return nil, false
	}
	return o.SerialId, true
}

// HasSerialId returns a boolean if a field has been set.
func (o *ApplianceDeviceUpgradePolicyAllOf) HasSerialId() bool {
	if o != nil && o.SerialId != nil {
		return true
	}

	return false
}

// SetSerialId gets a reference to the given string and assigns it to the SerialId field.
func (o *ApplianceDeviceUpgradePolicyAllOf) SetSerialId(v string) {
	o.SerialId = &v
}

// GetSoftwareDownloadType returns the SoftwareDownloadType field value if set, zero value otherwise.
func (o *ApplianceDeviceUpgradePolicyAllOf) GetSoftwareDownloadType() string {
	if o == nil || o.SoftwareDownloadType == nil {
		var ret string
		return ret
	}
	return *o.SoftwareDownloadType
}

// GetSoftwareDownloadTypeOk returns a tuple with the SoftwareDownloadType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceDeviceUpgradePolicyAllOf) GetSoftwareDownloadTypeOk() (*string, bool) {
	if o == nil || o.SoftwareDownloadType == nil {
		return nil, false
	}
	return o.SoftwareDownloadType, true
}

// HasSoftwareDownloadType returns a boolean if a field has been set.
func (o *ApplianceDeviceUpgradePolicyAllOf) HasSoftwareDownloadType() bool {
	if o != nil && o.SoftwareDownloadType != nil {
		return true
	}

	return false
}

// SetSoftwareDownloadType gets a reference to the given string and assigns it to the SoftwareDownloadType field.
func (o *ApplianceDeviceUpgradePolicyAllOf) SetSoftwareDownloadType(v string) {
	o.SoftwareDownloadType = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *ApplianceDeviceUpgradePolicyAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceDeviceUpgradePolicyAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *ApplianceDeviceUpgradePolicyAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *ApplianceDeviceUpgradePolicyAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o ApplianceDeviceUpgradePolicyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AutoUpgrade != nil {
		toSerialize["AutoUpgrade"] = o.AutoUpgrade
	}
	if o.BlackoutDatesEnabled != nil {
		toSerialize["BlackoutDatesEnabled"] = o.BlackoutDatesEnabled
	}
	if o.BlackoutEndDate != nil {
		toSerialize["BlackoutEndDate"] = o.BlackoutEndDate
	}
	if o.BlackoutStartDate != nil {
		toSerialize["BlackoutStartDate"] = o.BlackoutStartDate
	}
	if o.EnableMetaDataSync != nil {
		toSerialize["EnableMetaDataSync"] = o.EnableMetaDataSync
	}
	if o.Schedule.IsSet() {
		toSerialize["Schedule"] = o.Schedule.Get()
	}
	if o.SerialId != nil {
		toSerialize["SerialId"] = o.SerialId
	}
	if o.SoftwareDownloadType != nil {
		toSerialize["SoftwareDownloadType"] = o.SoftwareDownloadType
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ApplianceDeviceUpgradePolicyAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varApplianceDeviceUpgradePolicyAllOf := _ApplianceDeviceUpgradePolicyAllOf{}

	if err = json.Unmarshal(bytes, &varApplianceDeviceUpgradePolicyAllOf); err == nil {
		*o = ApplianceDeviceUpgradePolicyAllOf(varApplianceDeviceUpgradePolicyAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AutoUpgrade")
		delete(additionalProperties, "BlackoutDatesEnabled")
		delete(additionalProperties, "BlackoutEndDate")
		delete(additionalProperties, "BlackoutStartDate")
		delete(additionalProperties, "EnableMetaDataSync")
		delete(additionalProperties, "Schedule")
		delete(additionalProperties, "SerialId")
		delete(additionalProperties, "SoftwareDownloadType")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApplianceDeviceUpgradePolicyAllOf struct {
	value *ApplianceDeviceUpgradePolicyAllOf
	isSet bool
}

func (v NullableApplianceDeviceUpgradePolicyAllOf) Get() *ApplianceDeviceUpgradePolicyAllOf {
	return v.value
}

func (v *NullableApplianceDeviceUpgradePolicyAllOf) Set(val *ApplianceDeviceUpgradePolicyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceDeviceUpgradePolicyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceDeviceUpgradePolicyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceDeviceUpgradePolicyAllOf(val *ApplianceDeviceUpgradePolicyAllOf) *NullableApplianceDeviceUpgradePolicyAllOf {
	return &NullableApplianceDeviceUpgradePolicyAllOf{value: val, isSet: true}
}

func (v NullableApplianceDeviceUpgradePolicyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceDeviceUpgradePolicyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
