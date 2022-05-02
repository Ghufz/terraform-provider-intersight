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
	"time"
)

// ForecastInstanceAllOf Definition of the list of properties defined in 'forecast.Instance', excluding properties defined in parent classes.
type ForecastInstanceAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Action to be triggered on forecast instance. Default value is None. * `None` - The Enum value None represents that no action is triggered on the forecast Instance managed object. * `Evaluate` - The Enum value Evaluate represents that a re-evaluation of the forecast needs to be triggered.
	Action   *string   `json:"Action,omitempty"`
	AltModel []float32 `json:"AltModel,omitempty"`
	// The time interval (in days) for the data to be used for computing forecast model.
	DataInterval *int64 `json:"DataInterval,omitempty"`
	// The start date from when the data should be used for computing forecast model.
	DataStartDate *time.Time `json:"DataStartDate,omitempty"`
	// The Moid of the Intersight managed device instance for which regression model is derived.
	DeviceId *string `json:"DeviceId,omitempty"`
	// The number of days remaining before the device reaches its full functional capacity.
	FullCapDays *int64 `json:"FullCapDays,omitempty"`
	// The time when the forecast model was last updated.
	LastModelUpdateTime *time.Time `json:"LastModelUpdateTime,omitempty"`
	// The name of the metric for which regression model is generated.
	MetricName *string `json:"MetricName,omitempty"`
	// The minimum number of days the HyperFlex cluster should be up for computing forecast.
	MinDaysForForecast *int64                `json:"MinDaysForForecast,omitempty"`
	Model              NullableForecastModel `json:"Model,omitempty"`
	// The number of days remaining before the device reaches the specified threshold for the metric as defined in definition.
	ThresholdDays        *int64                               `json:"ThresholdDays,omitempty"`
	ForecastDef          *ForecastDefinitionRelationship      `json:"ForecastDef,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ForecastInstanceAllOf ForecastInstanceAllOf

// NewForecastInstanceAllOf instantiates a new ForecastInstanceAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewForecastInstanceAllOf(classId string, objectType string) *ForecastInstanceAllOf {
	this := ForecastInstanceAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var action string = "None"
	this.Action = &action
	var dataInterval int64 = 180
	this.DataInterval = &dataInterval
	return &this
}

// NewForecastInstanceAllOfWithDefaults instantiates a new ForecastInstanceAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewForecastInstanceAllOfWithDefaults() *ForecastInstanceAllOf {
	this := ForecastInstanceAllOf{}
	var classId string = "forecast.Instance"
	this.ClassId = classId
	var objectType string = "forecast.Instance"
	this.ObjectType = objectType
	var action string = "None"
	this.Action = &action
	var dataInterval int64 = 180
	this.DataInterval = &dataInterval
	return &this
}

// GetClassId returns the ClassId field value
func (o *ForecastInstanceAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ForecastInstanceAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ForecastInstanceAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ForecastInstanceAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ForecastInstanceAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ForecastInstanceAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *ForecastInstanceAllOf) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForecastInstanceAllOf) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *ForecastInstanceAllOf) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *ForecastInstanceAllOf) SetAction(v string) {
	o.Action = &v
}

// GetAltModel returns the AltModel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ForecastInstanceAllOf) GetAltModel() []float32 {
	if o == nil {
		var ret []float32
		return ret
	}
	return o.AltModel
}

// GetAltModelOk returns a tuple with the AltModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ForecastInstanceAllOf) GetAltModelOk() ([]float32, bool) {
	if o == nil || o.AltModel == nil {
		return nil, false
	}
	return o.AltModel, true
}

// HasAltModel returns a boolean if a field has been set.
func (o *ForecastInstanceAllOf) HasAltModel() bool {
	if o != nil && o.AltModel != nil {
		return true
	}

	return false
}

// SetAltModel gets a reference to the given []float32 and assigns it to the AltModel field.
func (o *ForecastInstanceAllOf) SetAltModel(v []float32) {
	o.AltModel = v
}

// GetDataInterval returns the DataInterval field value if set, zero value otherwise.
func (o *ForecastInstanceAllOf) GetDataInterval() int64 {
	if o == nil || o.DataInterval == nil {
		var ret int64
		return ret
	}
	return *o.DataInterval
}

// GetDataIntervalOk returns a tuple with the DataInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForecastInstanceAllOf) GetDataIntervalOk() (*int64, bool) {
	if o == nil || o.DataInterval == nil {
		return nil, false
	}
	return o.DataInterval, true
}

// HasDataInterval returns a boolean if a field has been set.
func (o *ForecastInstanceAllOf) HasDataInterval() bool {
	if o != nil && o.DataInterval != nil {
		return true
	}

	return false
}

// SetDataInterval gets a reference to the given int64 and assigns it to the DataInterval field.
func (o *ForecastInstanceAllOf) SetDataInterval(v int64) {
	o.DataInterval = &v
}

// GetDataStartDate returns the DataStartDate field value if set, zero value otherwise.
func (o *ForecastInstanceAllOf) GetDataStartDate() time.Time {
	if o == nil || o.DataStartDate == nil {
		var ret time.Time
		return ret
	}
	return *o.DataStartDate
}

// GetDataStartDateOk returns a tuple with the DataStartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForecastInstanceAllOf) GetDataStartDateOk() (*time.Time, bool) {
	if o == nil || o.DataStartDate == nil {
		return nil, false
	}
	return o.DataStartDate, true
}

// HasDataStartDate returns a boolean if a field has been set.
func (o *ForecastInstanceAllOf) HasDataStartDate() bool {
	if o != nil && o.DataStartDate != nil {
		return true
	}

	return false
}

// SetDataStartDate gets a reference to the given time.Time and assigns it to the DataStartDate field.
func (o *ForecastInstanceAllOf) SetDataStartDate(v time.Time) {
	o.DataStartDate = &v
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *ForecastInstanceAllOf) GetDeviceId() string {
	if o == nil || o.DeviceId == nil {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForecastInstanceAllOf) GetDeviceIdOk() (*string, bool) {
	if o == nil || o.DeviceId == nil {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *ForecastInstanceAllOf) HasDeviceId() bool {
	if o != nil && o.DeviceId != nil {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *ForecastInstanceAllOf) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetFullCapDays returns the FullCapDays field value if set, zero value otherwise.
func (o *ForecastInstanceAllOf) GetFullCapDays() int64 {
	if o == nil || o.FullCapDays == nil {
		var ret int64
		return ret
	}
	return *o.FullCapDays
}

// GetFullCapDaysOk returns a tuple with the FullCapDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForecastInstanceAllOf) GetFullCapDaysOk() (*int64, bool) {
	if o == nil || o.FullCapDays == nil {
		return nil, false
	}
	return o.FullCapDays, true
}

// HasFullCapDays returns a boolean if a field has been set.
func (o *ForecastInstanceAllOf) HasFullCapDays() bool {
	if o != nil && o.FullCapDays != nil {
		return true
	}

	return false
}

// SetFullCapDays gets a reference to the given int64 and assigns it to the FullCapDays field.
func (o *ForecastInstanceAllOf) SetFullCapDays(v int64) {
	o.FullCapDays = &v
}

// GetLastModelUpdateTime returns the LastModelUpdateTime field value if set, zero value otherwise.
func (o *ForecastInstanceAllOf) GetLastModelUpdateTime() time.Time {
	if o == nil || o.LastModelUpdateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModelUpdateTime
}

// GetLastModelUpdateTimeOk returns a tuple with the LastModelUpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForecastInstanceAllOf) GetLastModelUpdateTimeOk() (*time.Time, bool) {
	if o == nil || o.LastModelUpdateTime == nil {
		return nil, false
	}
	return o.LastModelUpdateTime, true
}

// HasLastModelUpdateTime returns a boolean if a field has been set.
func (o *ForecastInstanceAllOf) HasLastModelUpdateTime() bool {
	if o != nil && o.LastModelUpdateTime != nil {
		return true
	}

	return false
}

// SetLastModelUpdateTime gets a reference to the given time.Time and assigns it to the LastModelUpdateTime field.
func (o *ForecastInstanceAllOf) SetLastModelUpdateTime(v time.Time) {
	o.LastModelUpdateTime = &v
}

// GetMetricName returns the MetricName field value if set, zero value otherwise.
func (o *ForecastInstanceAllOf) GetMetricName() string {
	if o == nil || o.MetricName == nil {
		var ret string
		return ret
	}
	return *o.MetricName
}

// GetMetricNameOk returns a tuple with the MetricName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForecastInstanceAllOf) GetMetricNameOk() (*string, bool) {
	if o == nil || o.MetricName == nil {
		return nil, false
	}
	return o.MetricName, true
}

// HasMetricName returns a boolean if a field has been set.
func (o *ForecastInstanceAllOf) HasMetricName() bool {
	if o != nil && o.MetricName != nil {
		return true
	}

	return false
}

// SetMetricName gets a reference to the given string and assigns it to the MetricName field.
func (o *ForecastInstanceAllOf) SetMetricName(v string) {
	o.MetricName = &v
}

// GetMinDaysForForecast returns the MinDaysForForecast field value if set, zero value otherwise.
func (o *ForecastInstanceAllOf) GetMinDaysForForecast() int64 {
	if o == nil || o.MinDaysForForecast == nil {
		var ret int64
		return ret
	}
	return *o.MinDaysForForecast
}

// GetMinDaysForForecastOk returns a tuple with the MinDaysForForecast field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForecastInstanceAllOf) GetMinDaysForForecastOk() (*int64, bool) {
	if o == nil || o.MinDaysForForecast == nil {
		return nil, false
	}
	return o.MinDaysForForecast, true
}

// HasMinDaysForForecast returns a boolean if a field has been set.
func (o *ForecastInstanceAllOf) HasMinDaysForForecast() bool {
	if o != nil && o.MinDaysForForecast != nil {
		return true
	}

	return false
}

// SetMinDaysForForecast gets a reference to the given int64 and assigns it to the MinDaysForForecast field.
func (o *ForecastInstanceAllOf) SetMinDaysForForecast(v int64) {
	o.MinDaysForForecast = &v
}

// GetModel returns the Model field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ForecastInstanceAllOf) GetModel() ForecastModel {
	if o == nil || o.Model.Get() == nil {
		var ret ForecastModel
		return ret
	}
	return *o.Model.Get()
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ForecastInstanceAllOf) GetModelOk() (*ForecastModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Model.Get(), o.Model.IsSet()
}

// HasModel returns a boolean if a field has been set.
func (o *ForecastInstanceAllOf) HasModel() bool {
	if o != nil && o.Model.IsSet() {
		return true
	}

	return false
}

// SetModel gets a reference to the given NullableForecastModel and assigns it to the Model field.
func (o *ForecastInstanceAllOf) SetModel(v ForecastModel) {
	o.Model.Set(&v)
}

// SetModelNil sets the value for Model to be an explicit nil
func (o *ForecastInstanceAllOf) SetModelNil() {
	o.Model.Set(nil)
}

// UnsetModel ensures that no value is present for Model, not even an explicit nil
func (o *ForecastInstanceAllOf) UnsetModel() {
	o.Model.Unset()
}

// GetThresholdDays returns the ThresholdDays field value if set, zero value otherwise.
func (o *ForecastInstanceAllOf) GetThresholdDays() int64 {
	if o == nil || o.ThresholdDays == nil {
		var ret int64
		return ret
	}
	return *o.ThresholdDays
}

// GetThresholdDaysOk returns a tuple with the ThresholdDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForecastInstanceAllOf) GetThresholdDaysOk() (*int64, bool) {
	if o == nil || o.ThresholdDays == nil {
		return nil, false
	}
	return o.ThresholdDays, true
}

// HasThresholdDays returns a boolean if a field has been set.
func (o *ForecastInstanceAllOf) HasThresholdDays() bool {
	if o != nil && o.ThresholdDays != nil {
		return true
	}

	return false
}

// SetThresholdDays gets a reference to the given int64 and assigns it to the ThresholdDays field.
func (o *ForecastInstanceAllOf) SetThresholdDays(v int64) {
	o.ThresholdDays = &v
}

// GetForecastDef returns the ForecastDef field value if set, zero value otherwise.
func (o *ForecastInstanceAllOf) GetForecastDef() ForecastDefinitionRelationship {
	if o == nil || o.ForecastDef == nil {
		var ret ForecastDefinitionRelationship
		return ret
	}
	return *o.ForecastDef
}

// GetForecastDefOk returns a tuple with the ForecastDef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForecastInstanceAllOf) GetForecastDefOk() (*ForecastDefinitionRelationship, bool) {
	if o == nil || o.ForecastDef == nil {
		return nil, false
	}
	return o.ForecastDef, true
}

// HasForecastDef returns a boolean if a field has been set.
func (o *ForecastInstanceAllOf) HasForecastDef() bool {
	if o != nil && o.ForecastDef != nil {
		return true
	}

	return false
}

// SetForecastDef gets a reference to the given ForecastDefinitionRelationship and assigns it to the ForecastDef field.
func (o *ForecastInstanceAllOf) SetForecastDef(v ForecastDefinitionRelationship) {
	o.ForecastDef = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *ForecastInstanceAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForecastInstanceAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *ForecastInstanceAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *ForecastInstanceAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o ForecastInstanceAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Action != nil {
		toSerialize["Action"] = o.Action
	}
	if o.AltModel != nil {
		toSerialize["AltModel"] = o.AltModel
	}
	if o.DataInterval != nil {
		toSerialize["DataInterval"] = o.DataInterval
	}
	if o.DataStartDate != nil {
		toSerialize["DataStartDate"] = o.DataStartDate
	}
	if o.DeviceId != nil {
		toSerialize["DeviceId"] = o.DeviceId
	}
	if o.FullCapDays != nil {
		toSerialize["FullCapDays"] = o.FullCapDays
	}
	if o.LastModelUpdateTime != nil {
		toSerialize["LastModelUpdateTime"] = o.LastModelUpdateTime
	}
	if o.MetricName != nil {
		toSerialize["MetricName"] = o.MetricName
	}
	if o.MinDaysForForecast != nil {
		toSerialize["MinDaysForForecast"] = o.MinDaysForForecast
	}
	if o.Model.IsSet() {
		toSerialize["Model"] = o.Model.Get()
	}
	if o.ThresholdDays != nil {
		toSerialize["ThresholdDays"] = o.ThresholdDays
	}
	if o.ForecastDef != nil {
		toSerialize["ForecastDef"] = o.ForecastDef
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ForecastInstanceAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varForecastInstanceAllOf := _ForecastInstanceAllOf{}

	if err = json.Unmarshal(bytes, &varForecastInstanceAllOf); err == nil {
		*o = ForecastInstanceAllOf(varForecastInstanceAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Action")
		delete(additionalProperties, "AltModel")
		delete(additionalProperties, "DataInterval")
		delete(additionalProperties, "DataStartDate")
		delete(additionalProperties, "DeviceId")
		delete(additionalProperties, "FullCapDays")
		delete(additionalProperties, "LastModelUpdateTime")
		delete(additionalProperties, "MetricName")
		delete(additionalProperties, "MinDaysForForecast")
		delete(additionalProperties, "Model")
		delete(additionalProperties, "ThresholdDays")
		delete(additionalProperties, "ForecastDef")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableForecastInstanceAllOf struct {
	value *ForecastInstanceAllOf
	isSet bool
}

func (v NullableForecastInstanceAllOf) Get() *ForecastInstanceAllOf {
	return v.value
}

func (v *NullableForecastInstanceAllOf) Set(val *ForecastInstanceAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableForecastInstanceAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableForecastInstanceAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableForecastInstanceAllOf(val *ForecastInstanceAllOf) *NullableForecastInstanceAllOf {
	return &NullableForecastInstanceAllOf{value: val, isSet: true}
}

func (v NullableForecastInstanceAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableForecastInstanceAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
