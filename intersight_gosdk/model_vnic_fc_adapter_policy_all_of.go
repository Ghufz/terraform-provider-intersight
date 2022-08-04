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

// VnicFcAdapterPolicyAllOf Definition of the list of properties defined in 'vnic.FcAdapterPolicy', excluding properties defined in parent classes.
type VnicFcAdapterPolicyAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Error Detection Timeout, also referred to as EDTOV, is the number of milliseconds to wait before the system assumes that an error has occurred.
	ErrorDetectionTimeout *int64                              `json:"ErrorDetectionTimeout,omitempty"`
	ErrorRecoverySettings NullableVnicFcErrorRecoverySettings `json:"ErrorRecoverySettings,omitempty"`
	FlogiSettings         NullableVnicFlogiSettings           `json:"FlogiSettings,omitempty"`
	InterruptSettings     NullableVnicFcInterruptSettings     `json:"InterruptSettings,omitempty"`
	// The maximum number of data or control I/O operations that can be pending for the virtual interface at one time. If this value is exceeded, the additional I/O operations wait in the queue until the number of pending I/O operations decreases and the additional operations can be processed.
	IoThrottleCount *int64 `json:"IoThrottleCount,omitempty"`
	// The maximum number of LUNs that the Fibre Channel driver will export or show. The maximum number of LUNs is usually controlled by the operating system running on the server.
	LunCount *int64 `json:"LunCount,omitempty"`
	// The number of commands that the HBA can send and receive in a single transmission per LUN.
	LunQueueDepth *int64                    `json:"LunQueueDepth,omitempty"`
	PlogiSettings NullableVnicPlogiSettings `json:"PlogiSettings,omitempty"`
	// Resource Allocation Timeout, also referred to as RATOV, is the number of milliseconds to wait before the system assumes that a resource cannot be properly allocated.
	ResourceAllocationTimeout *int64                                `json:"ResourceAllocationTimeout,omitempty"`
	RxQueueSettings           NullableVnicFcQueueSettings           `json:"RxQueueSettings,omitempty"`
	ScsiQueueSettings         NullableVnicScsiQueueSettings         `json:"ScsiQueueSettings,omitempty"`
	TxQueueSettings           NullableVnicFcQueueSettings           `json:"TxQueueSettings,omitempty"`
	Organization              *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties      map[string]interface{}
}

type _VnicFcAdapterPolicyAllOf VnicFcAdapterPolicyAllOf

// NewVnicFcAdapterPolicyAllOf instantiates a new VnicFcAdapterPolicyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicFcAdapterPolicyAllOf(classId string, objectType string) *VnicFcAdapterPolicyAllOf {
	this := VnicFcAdapterPolicyAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var errorDetectionTimeout int64 = 2000
	this.ErrorDetectionTimeout = &errorDetectionTimeout
	var ioThrottleCount int64 = 512
	this.IoThrottleCount = &ioThrottleCount
	var lunCount int64 = 1024
	this.LunCount = &lunCount
	var lunQueueDepth int64 = 20
	this.LunQueueDepth = &lunQueueDepth
	var resourceAllocationTimeout int64 = 10000
	this.ResourceAllocationTimeout = &resourceAllocationTimeout
	return &this
}

// NewVnicFcAdapterPolicyAllOfWithDefaults instantiates a new VnicFcAdapterPolicyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicFcAdapterPolicyAllOfWithDefaults() *VnicFcAdapterPolicyAllOf {
	this := VnicFcAdapterPolicyAllOf{}
	var classId string = "vnic.FcAdapterPolicy"
	this.ClassId = classId
	var objectType string = "vnic.FcAdapterPolicy"
	this.ObjectType = objectType
	var errorDetectionTimeout int64 = 2000
	this.ErrorDetectionTimeout = &errorDetectionTimeout
	var ioThrottleCount int64 = 512
	this.IoThrottleCount = &ioThrottleCount
	var lunCount int64 = 1024
	this.LunCount = &lunCount
	var lunQueueDepth int64 = 20
	this.LunQueueDepth = &lunQueueDepth
	var resourceAllocationTimeout int64 = 10000
	this.ResourceAllocationTimeout = &resourceAllocationTimeout
	return &this
}

// GetClassId returns the ClassId field value
func (o *VnicFcAdapterPolicyAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VnicFcAdapterPolicyAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VnicFcAdapterPolicyAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VnicFcAdapterPolicyAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VnicFcAdapterPolicyAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VnicFcAdapterPolicyAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetErrorDetectionTimeout returns the ErrorDetectionTimeout field value if set, zero value otherwise.
func (o *VnicFcAdapterPolicyAllOf) GetErrorDetectionTimeout() int64 {
	if o == nil || o.ErrorDetectionTimeout == nil {
		var ret int64
		return ret
	}
	return *o.ErrorDetectionTimeout
}

// GetErrorDetectionTimeoutOk returns a tuple with the ErrorDetectionTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicFcAdapterPolicyAllOf) GetErrorDetectionTimeoutOk() (*int64, bool) {
	if o == nil || o.ErrorDetectionTimeout == nil {
		return nil, false
	}
	return o.ErrorDetectionTimeout, true
}

// HasErrorDetectionTimeout returns a boolean if a field has been set.
func (o *VnicFcAdapterPolicyAllOf) HasErrorDetectionTimeout() bool {
	if o != nil && o.ErrorDetectionTimeout != nil {
		return true
	}

	return false
}

// SetErrorDetectionTimeout gets a reference to the given int64 and assigns it to the ErrorDetectionTimeout field.
func (o *VnicFcAdapterPolicyAllOf) SetErrorDetectionTimeout(v int64) {
	o.ErrorDetectionTimeout = &v
}

// GetErrorRecoverySettings returns the ErrorRecoverySettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicFcAdapterPolicyAllOf) GetErrorRecoverySettings() VnicFcErrorRecoverySettings {
	if o == nil || o.ErrorRecoverySettings.Get() == nil {
		var ret VnicFcErrorRecoverySettings
		return ret
	}
	return *o.ErrorRecoverySettings.Get()
}

// GetErrorRecoverySettingsOk returns a tuple with the ErrorRecoverySettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicFcAdapterPolicyAllOf) GetErrorRecoverySettingsOk() (*VnicFcErrorRecoverySettings, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorRecoverySettings.Get(), o.ErrorRecoverySettings.IsSet()
}

// HasErrorRecoverySettings returns a boolean if a field has been set.
func (o *VnicFcAdapterPolicyAllOf) HasErrorRecoverySettings() bool {
	if o != nil && o.ErrorRecoverySettings.IsSet() {
		return true
	}

	return false
}

// SetErrorRecoverySettings gets a reference to the given NullableVnicFcErrorRecoverySettings and assigns it to the ErrorRecoverySettings field.
func (o *VnicFcAdapterPolicyAllOf) SetErrorRecoverySettings(v VnicFcErrorRecoverySettings) {
	o.ErrorRecoverySettings.Set(&v)
}

// SetErrorRecoverySettingsNil sets the value for ErrorRecoverySettings to be an explicit nil
func (o *VnicFcAdapterPolicyAllOf) SetErrorRecoverySettingsNil() {
	o.ErrorRecoverySettings.Set(nil)
}

// UnsetErrorRecoverySettings ensures that no value is present for ErrorRecoverySettings, not even an explicit nil
func (o *VnicFcAdapterPolicyAllOf) UnsetErrorRecoverySettings() {
	o.ErrorRecoverySettings.Unset()
}

// GetFlogiSettings returns the FlogiSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicFcAdapterPolicyAllOf) GetFlogiSettings() VnicFlogiSettings {
	if o == nil || o.FlogiSettings.Get() == nil {
		var ret VnicFlogiSettings
		return ret
	}
	return *o.FlogiSettings.Get()
}

// GetFlogiSettingsOk returns a tuple with the FlogiSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicFcAdapterPolicyAllOf) GetFlogiSettingsOk() (*VnicFlogiSettings, bool) {
	if o == nil {
		return nil, false
	}
	return o.FlogiSettings.Get(), o.FlogiSettings.IsSet()
}

// HasFlogiSettings returns a boolean if a field has been set.
func (o *VnicFcAdapterPolicyAllOf) HasFlogiSettings() bool {
	if o != nil && o.FlogiSettings.IsSet() {
		return true
	}

	return false
}

// SetFlogiSettings gets a reference to the given NullableVnicFlogiSettings and assigns it to the FlogiSettings field.
func (o *VnicFcAdapterPolicyAllOf) SetFlogiSettings(v VnicFlogiSettings) {
	o.FlogiSettings.Set(&v)
}

// SetFlogiSettingsNil sets the value for FlogiSettings to be an explicit nil
func (o *VnicFcAdapterPolicyAllOf) SetFlogiSettingsNil() {
	o.FlogiSettings.Set(nil)
}

// UnsetFlogiSettings ensures that no value is present for FlogiSettings, not even an explicit nil
func (o *VnicFcAdapterPolicyAllOf) UnsetFlogiSettings() {
	o.FlogiSettings.Unset()
}

// GetInterruptSettings returns the InterruptSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicFcAdapterPolicyAllOf) GetInterruptSettings() VnicFcInterruptSettings {
	if o == nil || o.InterruptSettings.Get() == nil {
		var ret VnicFcInterruptSettings
		return ret
	}
	return *o.InterruptSettings.Get()
}

// GetInterruptSettingsOk returns a tuple with the InterruptSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicFcAdapterPolicyAllOf) GetInterruptSettingsOk() (*VnicFcInterruptSettings, bool) {
	if o == nil {
		return nil, false
	}
	return o.InterruptSettings.Get(), o.InterruptSettings.IsSet()
}

// HasInterruptSettings returns a boolean if a field has been set.
func (o *VnicFcAdapterPolicyAllOf) HasInterruptSettings() bool {
	if o != nil && o.InterruptSettings.IsSet() {
		return true
	}

	return false
}

// SetInterruptSettings gets a reference to the given NullableVnicFcInterruptSettings and assigns it to the InterruptSettings field.
func (o *VnicFcAdapterPolicyAllOf) SetInterruptSettings(v VnicFcInterruptSettings) {
	o.InterruptSettings.Set(&v)
}

// SetInterruptSettingsNil sets the value for InterruptSettings to be an explicit nil
func (o *VnicFcAdapterPolicyAllOf) SetInterruptSettingsNil() {
	o.InterruptSettings.Set(nil)
}

// UnsetInterruptSettings ensures that no value is present for InterruptSettings, not even an explicit nil
func (o *VnicFcAdapterPolicyAllOf) UnsetInterruptSettings() {
	o.InterruptSettings.Unset()
}

// GetIoThrottleCount returns the IoThrottleCount field value if set, zero value otherwise.
func (o *VnicFcAdapterPolicyAllOf) GetIoThrottleCount() int64 {
	if o == nil || o.IoThrottleCount == nil {
		var ret int64
		return ret
	}
	return *o.IoThrottleCount
}

// GetIoThrottleCountOk returns a tuple with the IoThrottleCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicFcAdapterPolicyAllOf) GetIoThrottleCountOk() (*int64, bool) {
	if o == nil || o.IoThrottleCount == nil {
		return nil, false
	}
	return o.IoThrottleCount, true
}

// HasIoThrottleCount returns a boolean if a field has been set.
func (o *VnicFcAdapterPolicyAllOf) HasIoThrottleCount() bool {
	if o != nil && o.IoThrottleCount != nil {
		return true
	}

	return false
}

// SetIoThrottleCount gets a reference to the given int64 and assigns it to the IoThrottleCount field.
func (o *VnicFcAdapterPolicyAllOf) SetIoThrottleCount(v int64) {
	o.IoThrottleCount = &v
}

// GetLunCount returns the LunCount field value if set, zero value otherwise.
func (o *VnicFcAdapterPolicyAllOf) GetLunCount() int64 {
	if o == nil || o.LunCount == nil {
		var ret int64
		return ret
	}
	return *o.LunCount
}

// GetLunCountOk returns a tuple with the LunCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicFcAdapterPolicyAllOf) GetLunCountOk() (*int64, bool) {
	if o == nil || o.LunCount == nil {
		return nil, false
	}
	return o.LunCount, true
}

// HasLunCount returns a boolean if a field has been set.
func (o *VnicFcAdapterPolicyAllOf) HasLunCount() bool {
	if o != nil && o.LunCount != nil {
		return true
	}

	return false
}

// SetLunCount gets a reference to the given int64 and assigns it to the LunCount field.
func (o *VnicFcAdapterPolicyAllOf) SetLunCount(v int64) {
	o.LunCount = &v
}

// GetLunQueueDepth returns the LunQueueDepth field value if set, zero value otherwise.
func (o *VnicFcAdapterPolicyAllOf) GetLunQueueDepth() int64 {
	if o == nil || o.LunQueueDepth == nil {
		var ret int64
		return ret
	}
	return *o.LunQueueDepth
}

// GetLunQueueDepthOk returns a tuple with the LunQueueDepth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicFcAdapterPolicyAllOf) GetLunQueueDepthOk() (*int64, bool) {
	if o == nil || o.LunQueueDepth == nil {
		return nil, false
	}
	return o.LunQueueDepth, true
}

// HasLunQueueDepth returns a boolean if a field has been set.
func (o *VnicFcAdapterPolicyAllOf) HasLunQueueDepth() bool {
	if o != nil && o.LunQueueDepth != nil {
		return true
	}

	return false
}

// SetLunQueueDepth gets a reference to the given int64 and assigns it to the LunQueueDepth field.
func (o *VnicFcAdapterPolicyAllOf) SetLunQueueDepth(v int64) {
	o.LunQueueDepth = &v
}

// GetPlogiSettings returns the PlogiSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicFcAdapterPolicyAllOf) GetPlogiSettings() VnicPlogiSettings {
	if o == nil || o.PlogiSettings.Get() == nil {
		var ret VnicPlogiSettings
		return ret
	}
	return *o.PlogiSettings.Get()
}

// GetPlogiSettingsOk returns a tuple with the PlogiSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicFcAdapterPolicyAllOf) GetPlogiSettingsOk() (*VnicPlogiSettings, bool) {
	if o == nil {
		return nil, false
	}
	return o.PlogiSettings.Get(), o.PlogiSettings.IsSet()
}

// HasPlogiSettings returns a boolean if a field has been set.
func (o *VnicFcAdapterPolicyAllOf) HasPlogiSettings() bool {
	if o != nil && o.PlogiSettings.IsSet() {
		return true
	}

	return false
}

// SetPlogiSettings gets a reference to the given NullableVnicPlogiSettings and assigns it to the PlogiSettings field.
func (o *VnicFcAdapterPolicyAllOf) SetPlogiSettings(v VnicPlogiSettings) {
	o.PlogiSettings.Set(&v)
}

// SetPlogiSettingsNil sets the value for PlogiSettings to be an explicit nil
func (o *VnicFcAdapterPolicyAllOf) SetPlogiSettingsNil() {
	o.PlogiSettings.Set(nil)
}

// UnsetPlogiSettings ensures that no value is present for PlogiSettings, not even an explicit nil
func (o *VnicFcAdapterPolicyAllOf) UnsetPlogiSettings() {
	o.PlogiSettings.Unset()
}

// GetResourceAllocationTimeout returns the ResourceAllocationTimeout field value if set, zero value otherwise.
func (o *VnicFcAdapterPolicyAllOf) GetResourceAllocationTimeout() int64 {
	if o == nil || o.ResourceAllocationTimeout == nil {
		var ret int64
		return ret
	}
	return *o.ResourceAllocationTimeout
}

// GetResourceAllocationTimeoutOk returns a tuple with the ResourceAllocationTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicFcAdapterPolicyAllOf) GetResourceAllocationTimeoutOk() (*int64, bool) {
	if o == nil || o.ResourceAllocationTimeout == nil {
		return nil, false
	}
	return o.ResourceAllocationTimeout, true
}

// HasResourceAllocationTimeout returns a boolean if a field has been set.
func (o *VnicFcAdapterPolicyAllOf) HasResourceAllocationTimeout() bool {
	if o != nil && o.ResourceAllocationTimeout != nil {
		return true
	}

	return false
}

// SetResourceAllocationTimeout gets a reference to the given int64 and assigns it to the ResourceAllocationTimeout field.
func (o *VnicFcAdapterPolicyAllOf) SetResourceAllocationTimeout(v int64) {
	o.ResourceAllocationTimeout = &v
}

// GetRxQueueSettings returns the RxQueueSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicFcAdapterPolicyAllOf) GetRxQueueSettings() VnicFcQueueSettings {
	if o == nil || o.RxQueueSettings.Get() == nil {
		var ret VnicFcQueueSettings
		return ret
	}
	return *o.RxQueueSettings.Get()
}

// GetRxQueueSettingsOk returns a tuple with the RxQueueSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicFcAdapterPolicyAllOf) GetRxQueueSettingsOk() (*VnicFcQueueSettings, bool) {
	if o == nil {
		return nil, false
	}
	return o.RxQueueSettings.Get(), o.RxQueueSettings.IsSet()
}

// HasRxQueueSettings returns a boolean if a field has been set.
func (o *VnicFcAdapterPolicyAllOf) HasRxQueueSettings() bool {
	if o != nil && o.RxQueueSettings.IsSet() {
		return true
	}

	return false
}

// SetRxQueueSettings gets a reference to the given NullableVnicFcQueueSettings and assigns it to the RxQueueSettings field.
func (o *VnicFcAdapterPolicyAllOf) SetRxQueueSettings(v VnicFcQueueSettings) {
	o.RxQueueSettings.Set(&v)
}

// SetRxQueueSettingsNil sets the value for RxQueueSettings to be an explicit nil
func (o *VnicFcAdapterPolicyAllOf) SetRxQueueSettingsNil() {
	o.RxQueueSettings.Set(nil)
}

// UnsetRxQueueSettings ensures that no value is present for RxQueueSettings, not even an explicit nil
func (o *VnicFcAdapterPolicyAllOf) UnsetRxQueueSettings() {
	o.RxQueueSettings.Unset()
}

// GetScsiQueueSettings returns the ScsiQueueSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicFcAdapterPolicyAllOf) GetScsiQueueSettings() VnicScsiQueueSettings {
	if o == nil || o.ScsiQueueSettings.Get() == nil {
		var ret VnicScsiQueueSettings
		return ret
	}
	return *o.ScsiQueueSettings.Get()
}

// GetScsiQueueSettingsOk returns a tuple with the ScsiQueueSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicFcAdapterPolicyAllOf) GetScsiQueueSettingsOk() (*VnicScsiQueueSettings, bool) {
	if o == nil {
		return nil, false
	}
	return o.ScsiQueueSettings.Get(), o.ScsiQueueSettings.IsSet()
}

// HasScsiQueueSettings returns a boolean if a field has been set.
func (o *VnicFcAdapterPolicyAllOf) HasScsiQueueSettings() bool {
	if o != nil && o.ScsiQueueSettings.IsSet() {
		return true
	}

	return false
}

// SetScsiQueueSettings gets a reference to the given NullableVnicScsiQueueSettings and assigns it to the ScsiQueueSettings field.
func (o *VnicFcAdapterPolicyAllOf) SetScsiQueueSettings(v VnicScsiQueueSettings) {
	o.ScsiQueueSettings.Set(&v)
}

// SetScsiQueueSettingsNil sets the value for ScsiQueueSettings to be an explicit nil
func (o *VnicFcAdapterPolicyAllOf) SetScsiQueueSettingsNil() {
	o.ScsiQueueSettings.Set(nil)
}

// UnsetScsiQueueSettings ensures that no value is present for ScsiQueueSettings, not even an explicit nil
func (o *VnicFcAdapterPolicyAllOf) UnsetScsiQueueSettings() {
	o.ScsiQueueSettings.Unset()
}

// GetTxQueueSettings returns the TxQueueSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicFcAdapterPolicyAllOf) GetTxQueueSettings() VnicFcQueueSettings {
	if o == nil || o.TxQueueSettings.Get() == nil {
		var ret VnicFcQueueSettings
		return ret
	}
	return *o.TxQueueSettings.Get()
}

// GetTxQueueSettingsOk returns a tuple with the TxQueueSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicFcAdapterPolicyAllOf) GetTxQueueSettingsOk() (*VnicFcQueueSettings, bool) {
	if o == nil {
		return nil, false
	}
	return o.TxQueueSettings.Get(), o.TxQueueSettings.IsSet()
}

// HasTxQueueSettings returns a boolean if a field has been set.
func (o *VnicFcAdapterPolicyAllOf) HasTxQueueSettings() bool {
	if o != nil && o.TxQueueSettings.IsSet() {
		return true
	}

	return false
}

// SetTxQueueSettings gets a reference to the given NullableVnicFcQueueSettings and assigns it to the TxQueueSettings field.
func (o *VnicFcAdapterPolicyAllOf) SetTxQueueSettings(v VnicFcQueueSettings) {
	o.TxQueueSettings.Set(&v)
}

// SetTxQueueSettingsNil sets the value for TxQueueSettings to be an explicit nil
func (o *VnicFcAdapterPolicyAllOf) SetTxQueueSettingsNil() {
	o.TxQueueSettings.Set(nil)
}

// UnsetTxQueueSettings ensures that no value is present for TxQueueSettings, not even an explicit nil
func (o *VnicFcAdapterPolicyAllOf) UnsetTxQueueSettings() {
	o.TxQueueSettings.Unset()
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *VnicFcAdapterPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicFcAdapterPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *VnicFcAdapterPolicyAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *VnicFcAdapterPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o VnicFcAdapterPolicyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ErrorDetectionTimeout != nil {
		toSerialize["ErrorDetectionTimeout"] = o.ErrorDetectionTimeout
	}
	if o.ErrorRecoverySettings.IsSet() {
		toSerialize["ErrorRecoverySettings"] = o.ErrorRecoverySettings.Get()
	}
	if o.FlogiSettings.IsSet() {
		toSerialize["FlogiSettings"] = o.FlogiSettings.Get()
	}
	if o.InterruptSettings.IsSet() {
		toSerialize["InterruptSettings"] = o.InterruptSettings.Get()
	}
	if o.IoThrottleCount != nil {
		toSerialize["IoThrottleCount"] = o.IoThrottleCount
	}
	if o.LunCount != nil {
		toSerialize["LunCount"] = o.LunCount
	}
	if o.LunQueueDepth != nil {
		toSerialize["LunQueueDepth"] = o.LunQueueDepth
	}
	if o.PlogiSettings.IsSet() {
		toSerialize["PlogiSettings"] = o.PlogiSettings.Get()
	}
	if o.ResourceAllocationTimeout != nil {
		toSerialize["ResourceAllocationTimeout"] = o.ResourceAllocationTimeout
	}
	if o.RxQueueSettings.IsSet() {
		toSerialize["RxQueueSettings"] = o.RxQueueSettings.Get()
	}
	if o.ScsiQueueSettings.IsSet() {
		toSerialize["ScsiQueueSettings"] = o.ScsiQueueSettings.Get()
	}
	if o.TxQueueSettings.IsSet() {
		toSerialize["TxQueueSettings"] = o.TxQueueSettings.Get()
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VnicFcAdapterPolicyAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varVnicFcAdapterPolicyAllOf := _VnicFcAdapterPolicyAllOf{}

	if err = json.Unmarshal(bytes, &varVnicFcAdapterPolicyAllOf); err == nil {
		*o = VnicFcAdapterPolicyAllOf(varVnicFcAdapterPolicyAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ErrorDetectionTimeout")
		delete(additionalProperties, "ErrorRecoverySettings")
		delete(additionalProperties, "FlogiSettings")
		delete(additionalProperties, "InterruptSettings")
		delete(additionalProperties, "IoThrottleCount")
		delete(additionalProperties, "LunCount")
		delete(additionalProperties, "LunQueueDepth")
		delete(additionalProperties, "PlogiSettings")
		delete(additionalProperties, "ResourceAllocationTimeout")
		delete(additionalProperties, "RxQueueSettings")
		delete(additionalProperties, "ScsiQueueSettings")
		delete(additionalProperties, "TxQueueSettings")
		delete(additionalProperties, "Organization")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVnicFcAdapterPolicyAllOf struct {
	value *VnicFcAdapterPolicyAllOf
	isSet bool
}

func (v NullableVnicFcAdapterPolicyAllOf) Get() *VnicFcAdapterPolicyAllOf {
	return v.value
}

func (v *NullableVnicFcAdapterPolicyAllOf) Set(val *VnicFcAdapterPolicyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicFcAdapterPolicyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicFcAdapterPolicyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicFcAdapterPolicyAllOf(val *VnicFcAdapterPolicyAllOf) *NullableVnicFcAdapterPolicyAllOf {
	return &NullableVnicFcAdapterPolicyAllOf{value: val, isSet: true}
}

func (v NullableVnicFcAdapterPolicyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicFcAdapterPolicyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
