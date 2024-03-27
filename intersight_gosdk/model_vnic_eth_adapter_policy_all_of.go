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

// VnicEthAdapterPolicyAllOf Definition of the list of properties defined in 'vnic.EthAdapterPolicy', excluding properties defined in parent classes.
type VnicEthAdapterPolicyAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Enables advanced filtering on the interface.
	AdvancedFilter          *bool                               `json:"AdvancedFilter,omitempty"`
	ArfsSettings            NullableVnicArfsSettings            `json:"ArfsSettings,omitempty"`
	CompletionQueueSettings NullableVnicCompletionQueueSettings `json:"CompletionQueueSettings,omitempty"`
	// GENEVE offload protocol allows you to create logical networks that span physical network boundaries by allowing any information to be encoded in a packet and passed between tunnel endpoints.
	GeneveEnabled *bool `json:"GeneveEnabled,omitempty"`
	// Enables Interrupt Scaling on the interface.
	InterruptScaling  *bool                            `json:"InterruptScaling,omitempty"`
	InterruptSettings NullableVnicEthInterruptSettings `json:"InterruptSettings,omitempty"`
	NvgreSettings     NullableVnicNvgreSettings        `json:"NvgreSettings,omitempty"`
	PtpSettings       NullableVnicPtpSettings          `json:"PtpSettings,omitempty"`
	RoceSettings      NullableVnicRoceSettings         `json:"RoceSettings,omitempty"`
	RssHashSettings   NullableVnicRssHashSettings      `json:"RssHashSettings,omitempty"`
	// Receive Side Scaling allows the incoming traffic to be spread across multiple CPU cores.
	RssSettings        *bool                          `json:"RssSettings,omitempty"`
	RxQueueSettings    NullableVnicEthRxQueueSettings `json:"RxQueueSettings,omitempty"`
	TcpOffloadSettings NullableVnicTcpOffloadSettings `json:"TcpOffloadSettings,omitempty"`
	TxQueueSettings    NullableVnicEthTxQueueSettings `json:"TxQueueSettings,omitempty"`
	// Uplink Failback Timeout in seconds when uplink failover is enabled for a vNIC. After a vNIC has started using its secondary interface, this setting controls how long the primary interface must be available before the system resumes using the primary interface for the vNIC.
	UplinkFailbackTimeout *int64                                `json:"UplinkFailbackTimeout,omitempty"`
	VxlanSettings         NullableVnicVxlanSettings             `json:"VxlanSettings,omitempty"`
	Organization          *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _VnicEthAdapterPolicyAllOf VnicEthAdapterPolicyAllOf

// NewVnicEthAdapterPolicyAllOf instantiates a new VnicEthAdapterPolicyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicEthAdapterPolicyAllOf(classId string, objectType string) *VnicEthAdapterPolicyAllOf {
	this := VnicEthAdapterPolicyAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var advancedFilter bool = false
	this.AdvancedFilter = &advancedFilter
	var geneveEnabled bool = false
	this.GeneveEnabled = &geneveEnabled
	var interruptScaling bool = false
	this.InterruptScaling = &interruptScaling
	var rssSettings bool = true
	this.RssSettings = &rssSettings
	var uplinkFailbackTimeout int64 = 5
	this.UplinkFailbackTimeout = &uplinkFailbackTimeout
	return &this
}

// NewVnicEthAdapterPolicyAllOfWithDefaults instantiates a new VnicEthAdapterPolicyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicEthAdapterPolicyAllOfWithDefaults() *VnicEthAdapterPolicyAllOf {
	this := VnicEthAdapterPolicyAllOf{}
	var classId string = "vnic.EthAdapterPolicy"
	this.ClassId = classId
	var objectType string = "vnic.EthAdapterPolicy"
	this.ObjectType = objectType
	var advancedFilter bool = false
	this.AdvancedFilter = &advancedFilter
	var geneveEnabled bool = false
	this.GeneveEnabled = &geneveEnabled
	var interruptScaling bool = false
	this.InterruptScaling = &interruptScaling
	var rssSettings bool = true
	this.RssSettings = &rssSettings
	var uplinkFailbackTimeout int64 = 5
	this.UplinkFailbackTimeout = &uplinkFailbackTimeout
	return &this
}

// GetClassId returns the ClassId field value
func (o *VnicEthAdapterPolicyAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VnicEthAdapterPolicyAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VnicEthAdapterPolicyAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VnicEthAdapterPolicyAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VnicEthAdapterPolicyAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VnicEthAdapterPolicyAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAdvancedFilter returns the AdvancedFilter field value if set, zero value otherwise.
func (o *VnicEthAdapterPolicyAllOf) GetAdvancedFilter() bool {
	if o == nil || o.AdvancedFilter == nil {
		var ret bool
		return ret
	}
	return *o.AdvancedFilter
}

// GetAdvancedFilterOk returns a tuple with the AdvancedFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthAdapterPolicyAllOf) GetAdvancedFilterOk() (*bool, bool) {
	if o == nil || o.AdvancedFilter == nil {
		return nil, false
	}
	return o.AdvancedFilter, true
}

// HasAdvancedFilter returns a boolean if a field has been set.
func (o *VnicEthAdapterPolicyAllOf) HasAdvancedFilter() bool {
	if o != nil && o.AdvancedFilter != nil {
		return true
	}

	return false
}

// SetAdvancedFilter gets a reference to the given bool and assigns it to the AdvancedFilter field.
func (o *VnicEthAdapterPolicyAllOf) SetAdvancedFilter(v bool) {
	o.AdvancedFilter = &v
}

// GetArfsSettings returns the ArfsSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicEthAdapterPolicyAllOf) GetArfsSettings() VnicArfsSettings {
	if o == nil || o.ArfsSettings.Get() == nil {
		var ret VnicArfsSettings
		return ret
	}
	return *o.ArfsSettings.Get()
}

// GetArfsSettingsOk returns a tuple with the ArfsSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicEthAdapterPolicyAllOf) GetArfsSettingsOk() (*VnicArfsSettings, bool) {
	if o == nil {
		return nil, false
	}
	return o.ArfsSettings.Get(), o.ArfsSettings.IsSet()
}

// HasArfsSettings returns a boolean if a field has been set.
func (o *VnicEthAdapterPolicyAllOf) HasArfsSettings() bool {
	if o != nil && o.ArfsSettings.IsSet() {
		return true
	}

	return false
}

// SetArfsSettings gets a reference to the given NullableVnicArfsSettings and assigns it to the ArfsSettings field.
func (o *VnicEthAdapterPolicyAllOf) SetArfsSettings(v VnicArfsSettings) {
	o.ArfsSettings.Set(&v)
}

// SetArfsSettingsNil sets the value for ArfsSettings to be an explicit nil
func (o *VnicEthAdapterPolicyAllOf) SetArfsSettingsNil() {
	o.ArfsSettings.Set(nil)
}

// UnsetArfsSettings ensures that no value is present for ArfsSettings, not even an explicit nil
func (o *VnicEthAdapterPolicyAllOf) UnsetArfsSettings() {
	o.ArfsSettings.Unset()
}

// GetCompletionQueueSettings returns the CompletionQueueSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicEthAdapterPolicyAllOf) GetCompletionQueueSettings() VnicCompletionQueueSettings {
	if o == nil || o.CompletionQueueSettings.Get() == nil {
		var ret VnicCompletionQueueSettings
		return ret
	}
	return *o.CompletionQueueSettings.Get()
}

// GetCompletionQueueSettingsOk returns a tuple with the CompletionQueueSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicEthAdapterPolicyAllOf) GetCompletionQueueSettingsOk() (*VnicCompletionQueueSettings, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompletionQueueSettings.Get(), o.CompletionQueueSettings.IsSet()
}

// HasCompletionQueueSettings returns a boolean if a field has been set.
func (o *VnicEthAdapterPolicyAllOf) HasCompletionQueueSettings() bool {
	if o != nil && o.CompletionQueueSettings.IsSet() {
		return true
	}

	return false
}

// SetCompletionQueueSettings gets a reference to the given NullableVnicCompletionQueueSettings and assigns it to the CompletionQueueSettings field.
func (o *VnicEthAdapterPolicyAllOf) SetCompletionQueueSettings(v VnicCompletionQueueSettings) {
	o.CompletionQueueSettings.Set(&v)
}

// SetCompletionQueueSettingsNil sets the value for CompletionQueueSettings to be an explicit nil
func (o *VnicEthAdapterPolicyAllOf) SetCompletionQueueSettingsNil() {
	o.CompletionQueueSettings.Set(nil)
}

// UnsetCompletionQueueSettings ensures that no value is present for CompletionQueueSettings, not even an explicit nil
func (o *VnicEthAdapterPolicyAllOf) UnsetCompletionQueueSettings() {
	o.CompletionQueueSettings.Unset()
}

// GetGeneveEnabled returns the GeneveEnabled field value if set, zero value otherwise.
func (o *VnicEthAdapterPolicyAllOf) GetGeneveEnabled() bool {
	if o == nil || o.GeneveEnabled == nil {
		var ret bool
		return ret
	}
	return *o.GeneveEnabled
}

// GetGeneveEnabledOk returns a tuple with the GeneveEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthAdapterPolicyAllOf) GetGeneveEnabledOk() (*bool, bool) {
	if o == nil || o.GeneveEnabled == nil {
		return nil, false
	}
	return o.GeneveEnabled, true
}

// HasGeneveEnabled returns a boolean if a field has been set.
func (o *VnicEthAdapterPolicyAllOf) HasGeneveEnabled() bool {
	if o != nil && o.GeneveEnabled != nil {
		return true
	}

	return false
}

// SetGeneveEnabled gets a reference to the given bool and assigns it to the GeneveEnabled field.
func (o *VnicEthAdapterPolicyAllOf) SetGeneveEnabled(v bool) {
	o.GeneveEnabled = &v
}

// GetInterruptScaling returns the InterruptScaling field value if set, zero value otherwise.
func (o *VnicEthAdapterPolicyAllOf) GetInterruptScaling() bool {
	if o == nil || o.InterruptScaling == nil {
		var ret bool
		return ret
	}
	return *o.InterruptScaling
}

// GetInterruptScalingOk returns a tuple with the InterruptScaling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthAdapterPolicyAllOf) GetInterruptScalingOk() (*bool, bool) {
	if o == nil || o.InterruptScaling == nil {
		return nil, false
	}
	return o.InterruptScaling, true
}

// HasInterruptScaling returns a boolean if a field has been set.
func (o *VnicEthAdapterPolicyAllOf) HasInterruptScaling() bool {
	if o != nil && o.InterruptScaling != nil {
		return true
	}

	return false
}

// SetInterruptScaling gets a reference to the given bool and assigns it to the InterruptScaling field.
func (o *VnicEthAdapterPolicyAllOf) SetInterruptScaling(v bool) {
	o.InterruptScaling = &v
}

// GetInterruptSettings returns the InterruptSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicEthAdapterPolicyAllOf) GetInterruptSettings() VnicEthInterruptSettings {
	if o == nil || o.InterruptSettings.Get() == nil {
		var ret VnicEthInterruptSettings
		return ret
	}
	return *o.InterruptSettings.Get()
}

// GetInterruptSettingsOk returns a tuple with the InterruptSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicEthAdapterPolicyAllOf) GetInterruptSettingsOk() (*VnicEthInterruptSettings, bool) {
	if o == nil {
		return nil, false
	}
	return o.InterruptSettings.Get(), o.InterruptSettings.IsSet()
}

// HasInterruptSettings returns a boolean if a field has been set.
func (o *VnicEthAdapterPolicyAllOf) HasInterruptSettings() bool {
	if o != nil && o.InterruptSettings.IsSet() {
		return true
	}

	return false
}

// SetInterruptSettings gets a reference to the given NullableVnicEthInterruptSettings and assigns it to the InterruptSettings field.
func (o *VnicEthAdapterPolicyAllOf) SetInterruptSettings(v VnicEthInterruptSettings) {
	o.InterruptSettings.Set(&v)
}

// SetInterruptSettingsNil sets the value for InterruptSettings to be an explicit nil
func (o *VnicEthAdapterPolicyAllOf) SetInterruptSettingsNil() {
	o.InterruptSettings.Set(nil)
}

// UnsetInterruptSettings ensures that no value is present for InterruptSettings, not even an explicit nil
func (o *VnicEthAdapterPolicyAllOf) UnsetInterruptSettings() {
	o.InterruptSettings.Unset()
}

// GetNvgreSettings returns the NvgreSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicEthAdapterPolicyAllOf) GetNvgreSettings() VnicNvgreSettings {
	if o == nil || o.NvgreSettings.Get() == nil {
		var ret VnicNvgreSettings
		return ret
	}
	return *o.NvgreSettings.Get()
}

// GetNvgreSettingsOk returns a tuple with the NvgreSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicEthAdapterPolicyAllOf) GetNvgreSettingsOk() (*VnicNvgreSettings, bool) {
	if o == nil {
		return nil, false
	}
	return o.NvgreSettings.Get(), o.NvgreSettings.IsSet()
}

// HasNvgreSettings returns a boolean if a field has been set.
func (o *VnicEthAdapterPolicyAllOf) HasNvgreSettings() bool {
	if o != nil && o.NvgreSettings.IsSet() {
		return true
	}

	return false
}

// SetNvgreSettings gets a reference to the given NullableVnicNvgreSettings and assigns it to the NvgreSettings field.
func (o *VnicEthAdapterPolicyAllOf) SetNvgreSettings(v VnicNvgreSettings) {
	o.NvgreSettings.Set(&v)
}

// SetNvgreSettingsNil sets the value for NvgreSettings to be an explicit nil
func (o *VnicEthAdapterPolicyAllOf) SetNvgreSettingsNil() {
	o.NvgreSettings.Set(nil)
}

// UnsetNvgreSettings ensures that no value is present for NvgreSettings, not even an explicit nil
func (o *VnicEthAdapterPolicyAllOf) UnsetNvgreSettings() {
	o.NvgreSettings.Unset()
}

// GetPtpSettings returns the PtpSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicEthAdapterPolicyAllOf) GetPtpSettings() VnicPtpSettings {
	if o == nil || o.PtpSettings.Get() == nil {
		var ret VnicPtpSettings
		return ret
	}
	return *o.PtpSettings.Get()
}

// GetPtpSettingsOk returns a tuple with the PtpSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicEthAdapterPolicyAllOf) GetPtpSettingsOk() (*VnicPtpSettings, bool) {
	if o == nil {
		return nil, false
	}
	return o.PtpSettings.Get(), o.PtpSettings.IsSet()
}

// HasPtpSettings returns a boolean if a field has been set.
func (o *VnicEthAdapterPolicyAllOf) HasPtpSettings() bool {
	if o != nil && o.PtpSettings.IsSet() {
		return true
	}

	return false
}

// SetPtpSettings gets a reference to the given NullableVnicPtpSettings and assigns it to the PtpSettings field.
func (o *VnicEthAdapterPolicyAllOf) SetPtpSettings(v VnicPtpSettings) {
	o.PtpSettings.Set(&v)
}

// SetPtpSettingsNil sets the value for PtpSettings to be an explicit nil
func (o *VnicEthAdapterPolicyAllOf) SetPtpSettingsNil() {
	o.PtpSettings.Set(nil)
}

// UnsetPtpSettings ensures that no value is present for PtpSettings, not even an explicit nil
func (o *VnicEthAdapterPolicyAllOf) UnsetPtpSettings() {
	o.PtpSettings.Unset()
}

// GetRoceSettings returns the RoceSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicEthAdapterPolicyAllOf) GetRoceSettings() VnicRoceSettings {
	if o == nil || o.RoceSettings.Get() == nil {
		var ret VnicRoceSettings
		return ret
	}
	return *o.RoceSettings.Get()
}

// GetRoceSettingsOk returns a tuple with the RoceSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicEthAdapterPolicyAllOf) GetRoceSettingsOk() (*VnicRoceSettings, bool) {
	if o == nil {
		return nil, false
	}
	return o.RoceSettings.Get(), o.RoceSettings.IsSet()
}

// HasRoceSettings returns a boolean if a field has been set.
func (o *VnicEthAdapterPolicyAllOf) HasRoceSettings() bool {
	if o != nil && o.RoceSettings.IsSet() {
		return true
	}

	return false
}

// SetRoceSettings gets a reference to the given NullableVnicRoceSettings and assigns it to the RoceSettings field.
func (o *VnicEthAdapterPolicyAllOf) SetRoceSettings(v VnicRoceSettings) {
	o.RoceSettings.Set(&v)
}

// SetRoceSettingsNil sets the value for RoceSettings to be an explicit nil
func (o *VnicEthAdapterPolicyAllOf) SetRoceSettingsNil() {
	o.RoceSettings.Set(nil)
}

// UnsetRoceSettings ensures that no value is present for RoceSettings, not even an explicit nil
func (o *VnicEthAdapterPolicyAllOf) UnsetRoceSettings() {
	o.RoceSettings.Unset()
}

// GetRssHashSettings returns the RssHashSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicEthAdapterPolicyAllOf) GetRssHashSettings() VnicRssHashSettings {
	if o == nil || o.RssHashSettings.Get() == nil {
		var ret VnicRssHashSettings
		return ret
	}
	return *o.RssHashSettings.Get()
}

// GetRssHashSettingsOk returns a tuple with the RssHashSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicEthAdapterPolicyAllOf) GetRssHashSettingsOk() (*VnicRssHashSettings, bool) {
	if o == nil {
		return nil, false
	}
	return o.RssHashSettings.Get(), o.RssHashSettings.IsSet()
}

// HasRssHashSettings returns a boolean if a field has been set.
func (o *VnicEthAdapterPolicyAllOf) HasRssHashSettings() bool {
	if o != nil && o.RssHashSettings.IsSet() {
		return true
	}

	return false
}

// SetRssHashSettings gets a reference to the given NullableVnicRssHashSettings and assigns it to the RssHashSettings field.
func (o *VnicEthAdapterPolicyAllOf) SetRssHashSettings(v VnicRssHashSettings) {
	o.RssHashSettings.Set(&v)
}

// SetRssHashSettingsNil sets the value for RssHashSettings to be an explicit nil
func (o *VnicEthAdapterPolicyAllOf) SetRssHashSettingsNil() {
	o.RssHashSettings.Set(nil)
}

// UnsetRssHashSettings ensures that no value is present for RssHashSettings, not even an explicit nil
func (o *VnicEthAdapterPolicyAllOf) UnsetRssHashSettings() {
	o.RssHashSettings.Unset()
}

// GetRssSettings returns the RssSettings field value if set, zero value otherwise.
func (o *VnicEthAdapterPolicyAllOf) GetRssSettings() bool {
	if o == nil || o.RssSettings == nil {
		var ret bool
		return ret
	}
	return *o.RssSettings
}

// GetRssSettingsOk returns a tuple with the RssSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthAdapterPolicyAllOf) GetRssSettingsOk() (*bool, bool) {
	if o == nil || o.RssSettings == nil {
		return nil, false
	}
	return o.RssSettings, true
}

// HasRssSettings returns a boolean if a field has been set.
func (o *VnicEthAdapterPolicyAllOf) HasRssSettings() bool {
	if o != nil && o.RssSettings != nil {
		return true
	}

	return false
}

// SetRssSettings gets a reference to the given bool and assigns it to the RssSettings field.
func (o *VnicEthAdapterPolicyAllOf) SetRssSettings(v bool) {
	o.RssSettings = &v
}

// GetRxQueueSettings returns the RxQueueSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicEthAdapterPolicyAllOf) GetRxQueueSettings() VnicEthRxQueueSettings {
	if o == nil || o.RxQueueSettings.Get() == nil {
		var ret VnicEthRxQueueSettings
		return ret
	}
	return *o.RxQueueSettings.Get()
}

// GetRxQueueSettingsOk returns a tuple with the RxQueueSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicEthAdapterPolicyAllOf) GetRxQueueSettingsOk() (*VnicEthRxQueueSettings, bool) {
	if o == nil {
		return nil, false
	}
	return o.RxQueueSettings.Get(), o.RxQueueSettings.IsSet()
}

// HasRxQueueSettings returns a boolean if a field has been set.
func (o *VnicEthAdapterPolicyAllOf) HasRxQueueSettings() bool {
	if o != nil && o.RxQueueSettings.IsSet() {
		return true
	}

	return false
}

// SetRxQueueSettings gets a reference to the given NullableVnicEthRxQueueSettings and assigns it to the RxQueueSettings field.
func (o *VnicEthAdapterPolicyAllOf) SetRxQueueSettings(v VnicEthRxQueueSettings) {
	o.RxQueueSettings.Set(&v)
}

// SetRxQueueSettingsNil sets the value for RxQueueSettings to be an explicit nil
func (o *VnicEthAdapterPolicyAllOf) SetRxQueueSettingsNil() {
	o.RxQueueSettings.Set(nil)
}

// UnsetRxQueueSettings ensures that no value is present for RxQueueSettings, not even an explicit nil
func (o *VnicEthAdapterPolicyAllOf) UnsetRxQueueSettings() {
	o.RxQueueSettings.Unset()
}

// GetTcpOffloadSettings returns the TcpOffloadSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicEthAdapterPolicyAllOf) GetTcpOffloadSettings() VnicTcpOffloadSettings {
	if o == nil || o.TcpOffloadSettings.Get() == nil {
		var ret VnicTcpOffloadSettings
		return ret
	}
	return *o.TcpOffloadSettings.Get()
}

// GetTcpOffloadSettingsOk returns a tuple with the TcpOffloadSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicEthAdapterPolicyAllOf) GetTcpOffloadSettingsOk() (*VnicTcpOffloadSettings, bool) {
	if o == nil {
		return nil, false
	}
	return o.TcpOffloadSettings.Get(), o.TcpOffloadSettings.IsSet()
}

// HasTcpOffloadSettings returns a boolean if a field has been set.
func (o *VnicEthAdapterPolicyAllOf) HasTcpOffloadSettings() bool {
	if o != nil && o.TcpOffloadSettings.IsSet() {
		return true
	}

	return false
}

// SetTcpOffloadSettings gets a reference to the given NullableVnicTcpOffloadSettings and assigns it to the TcpOffloadSettings field.
func (o *VnicEthAdapterPolicyAllOf) SetTcpOffloadSettings(v VnicTcpOffloadSettings) {
	o.TcpOffloadSettings.Set(&v)
}

// SetTcpOffloadSettingsNil sets the value for TcpOffloadSettings to be an explicit nil
func (o *VnicEthAdapterPolicyAllOf) SetTcpOffloadSettingsNil() {
	o.TcpOffloadSettings.Set(nil)
}

// UnsetTcpOffloadSettings ensures that no value is present for TcpOffloadSettings, not even an explicit nil
func (o *VnicEthAdapterPolicyAllOf) UnsetTcpOffloadSettings() {
	o.TcpOffloadSettings.Unset()
}

// GetTxQueueSettings returns the TxQueueSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicEthAdapterPolicyAllOf) GetTxQueueSettings() VnicEthTxQueueSettings {
	if o == nil || o.TxQueueSettings.Get() == nil {
		var ret VnicEthTxQueueSettings
		return ret
	}
	return *o.TxQueueSettings.Get()
}

// GetTxQueueSettingsOk returns a tuple with the TxQueueSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicEthAdapterPolicyAllOf) GetTxQueueSettingsOk() (*VnicEthTxQueueSettings, bool) {
	if o == nil {
		return nil, false
	}
	return o.TxQueueSettings.Get(), o.TxQueueSettings.IsSet()
}

// HasTxQueueSettings returns a boolean if a field has been set.
func (o *VnicEthAdapterPolicyAllOf) HasTxQueueSettings() bool {
	if o != nil && o.TxQueueSettings.IsSet() {
		return true
	}

	return false
}

// SetTxQueueSettings gets a reference to the given NullableVnicEthTxQueueSettings and assigns it to the TxQueueSettings field.
func (o *VnicEthAdapterPolicyAllOf) SetTxQueueSettings(v VnicEthTxQueueSettings) {
	o.TxQueueSettings.Set(&v)
}

// SetTxQueueSettingsNil sets the value for TxQueueSettings to be an explicit nil
func (o *VnicEthAdapterPolicyAllOf) SetTxQueueSettingsNil() {
	o.TxQueueSettings.Set(nil)
}

// UnsetTxQueueSettings ensures that no value is present for TxQueueSettings, not even an explicit nil
func (o *VnicEthAdapterPolicyAllOf) UnsetTxQueueSettings() {
	o.TxQueueSettings.Unset()
}

// GetUplinkFailbackTimeout returns the UplinkFailbackTimeout field value if set, zero value otherwise.
func (o *VnicEthAdapterPolicyAllOf) GetUplinkFailbackTimeout() int64 {
	if o == nil || o.UplinkFailbackTimeout == nil {
		var ret int64
		return ret
	}
	return *o.UplinkFailbackTimeout
}

// GetUplinkFailbackTimeoutOk returns a tuple with the UplinkFailbackTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthAdapterPolicyAllOf) GetUplinkFailbackTimeoutOk() (*int64, bool) {
	if o == nil || o.UplinkFailbackTimeout == nil {
		return nil, false
	}
	return o.UplinkFailbackTimeout, true
}

// HasUplinkFailbackTimeout returns a boolean if a field has been set.
func (o *VnicEthAdapterPolicyAllOf) HasUplinkFailbackTimeout() bool {
	if o != nil && o.UplinkFailbackTimeout != nil {
		return true
	}

	return false
}

// SetUplinkFailbackTimeout gets a reference to the given int64 and assigns it to the UplinkFailbackTimeout field.
func (o *VnicEthAdapterPolicyAllOf) SetUplinkFailbackTimeout(v int64) {
	o.UplinkFailbackTimeout = &v
}

// GetVxlanSettings returns the VxlanSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicEthAdapterPolicyAllOf) GetVxlanSettings() VnicVxlanSettings {
	if o == nil || o.VxlanSettings.Get() == nil {
		var ret VnicVxlanSettings
		return ret
	}
	return *o.VxlanSettings.Get()
}

// GetVxlanSettingsOk returns a tuple with the VxlanSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicEthAdapterPolicyAllOf) GetVxlanSettingsOk() (*VnicVxlanSettings, bool) {
	if o == nil {
		return nil, false
	}
	return o.VxlanSettings.Get(), o.VxlanSettings.IsSet()
}

// HasVxlanSettings returns a boolean if a field has been set.
func (o *VnicEthAdapterPolicyAllOf) HasVxlanSettings() bool {
	if o != nil && o.VxlanSettings.IsSet() {
		return true
	}

	return false
}

// SetVxlanSettings gets a reference to the given NullableVnicVxlanSettings and assigns it to the VxlanSettings field.
func (o *VnicEthAdapterPolicyAllOf) SetVxlanSettings(v VnicVxlanSettings) {
	o.VxlanSettings.Set(&v)
}

// SetVxlanSettingsNil sets the value for VxlanSettings to be an explicit nil
func (o *VnicEthAdapterPolicyAllOf) SetVxlanSettingsNil() {
	o.VxlanSettings.Set(nil)
}

// UnsetVxlanSettings ensures that no value is present for VxlanSettings, not even an explicit nil
func (o *VnicEthAdapterPolicyAllOf) UnsetVxlanSettings() {
	o.VxlanSettings.Unset()
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *VnicEthAdapterPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthAdapterPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *VnicEthAdapterPolicyAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *VnicEthAdapterPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o VnicEthAdapterPolicyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AdvancedFilter != nil {
		toSerialize["AdvancedFilter"] = o.AdvancedFilter
	}
	if o.ArfsSettings.IsSet() {
		toSerialize["ArfsSettings"] = o.ArfsSettings.Get()
	}
	if o.CompletionQueueSettings.IsSet() {
		toSerialize["CompletionQueueSettings"] = o.CompletionQueueSettings.Get()
	}
	if o.GeneveEnabled != nil {
		toSerialize["GeneveEnabled"] = o.GeneveEnabled
	}
	if o.InterruptScaling != nil {
		toSerialize["InterruptScaling"] = o.InterruptScaling
	}
	if o.InterruptSettings.IsSet() {
		toSerialize["InterruptSettings"] = o.InterruptSettings.Get()
	}
	if o.NvgreSettings.IsSet() {
		toSerialize["NvgreSettings"] = o.NvgreSettings.Get()
	}
	if o.PtpSettings.IsSet() {
		toSerialize["PtpSettings"] = o.PtpSettings.Get()
	}
	if o.RoceSettings.IsSet() {
		toSerialize["RoceSettings"] = o.RoceSettings.Get()
	}
	if o.RssHashSettings.IsSet() {
		toSerialize["RssHashSettings"] = o.RssHashSettings.Get()
	}
	if o.RssSettings != nil {
		toSerialize["RssSettings"] = o.RssSettings
	}
	if o.RxQueueSettings.IsSet() {
		toSerialize["RxQueueSettings"] = o.RxQueueSettings.Get()
	}
	if o.TcpOffloadSettings.IsSet() {
		toSerialize["TcpOffloadSettings"] = o.TcpOffloadSettings.Get()
	}
	if o.TxQueueSettings.IsSet() {
		toSerialize["TxQueueSettings"] = o.TxQueueSettings.Get()
	}
	if o.UplinkFailbackTimeout != nil {
		toSerialize["UplinkFailbackTimeout"] = o.UplinkFailbackTimeout
	}
	if o.VxlanSettings.IsSet() {
		toSerialize["VxlanSettings"] = o.VxlanSettings.Get()
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VnicEthAdapterPolicyAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varVnicEthAdapterPolicyAllOf := _VnicEthAdapterPolicyAllOf{}

	if err = json.Unmarshal(bytes, &varVnicEthAdapterPolicyAllOf); err == nil {
		*o = VnicEthAdapterPolicyAllOf(varVnicEthAdapterPolicyAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AdvancedFilter")
		delete(additionalProperties, "ArfsSettings")
		delete(additionalProperties, "CompletionQueueSettings")
		delete(additionalProperties, "GeneveEnabled")
		delete(additionalProperties, "InterruptScaling")
		delete(additionalProperties, "InterruptSettings")
		delete(additionalProperties, "NvgreSettings")
		delete(additionalProperties, "PtpSettings")
		delete(additionalProperties, "RoceSettings")
		delete(additionalProperties, "RssHashSettings")
		delete(additionalProperties, "RssSettings")
		delete(additionalProperties, "RxQueueSettings")
		delete(additionalProperties, "TcpOffloadSettings")
		delete(additionalProperties, "TxQueueSettings")
		delete(additionalProperties, "UplinkFailbackTimeout")
		delete(additionalProperties, "VxlanSettings")
		delete(additionalProperties, "Organization")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVnicEthAdapterPolicyAllOf struct {
	value *VnicEthAdapterPolicyAllOf
	isSet bool
}

func (v NullableVnicEthAdapterPolicyAllOf) Get() *VnicEthAdapterPolicyAllOf {
	return v.value
}

func (v *NullableVnicEthAdapterPolicyAllOf) Set(val *VnicEthAdapterPolicyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicEthAdapterPolicyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicEthAdapterPolicyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicEthAdapterPolicyAllOf(val *VnicEthAdapterPolicyAllOf) *NullableVnicEthAdapterPolicyAllOf {
	return &NullableVnicEthAdapterPolicyAllOf{value: val, isSet: true}
}

func (v NullableVnicEthAdapterPolicyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicEthAdapterPolicyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
