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
)

// NiatelemetryEpgAllOf Definition of the list of properties defined in 'niatelemetry.Epg', excluding properties defined in parent classes.
type NiatelemetryEpgAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Azure Pack NAT with ASA feature usage.
	AzurePackCount *int64 `json:"AzurePackCount,omitempty"`
	// Dn value for the End Point Groups present.
	Dn *string `json:"Dn,omitempty"`
	// Number of  objects with delimiter value present in EPG Delimiter attribute.
	EpgDelimiterCount *int64 `json:"EpgDelimiterCount,omitempty"`
	// Number of ports with FC path attribute of type FC.
	FcNpvCount *int64 `json:"FcNpvCount,omitempty"`
	// Number of FCoE per End Point Group.
	FcoeCount *int64 `json:"FcoeCount,omitempty"`
	// Number of FvRsDomAtt objects per End Point Group with VMware configuration.
	FvRsDomAttCount *int64 `json:"FvRsDomAttCount,omitempty"`
	// Intra End Point Group Contract for Distributed Virtual Switch and BM feature usage.
	IntraEpgDvsBmCount *int64 `json:"IntraEpgDvsBmCount,omitempty"`
	// Intra EPG Isolation for Hyper-V, enabled if pcEnfPref attribute is set to enforced.
	IntraEpgHyperv *string `json:"IntraEpgHyperv,omitempty"`
	// Gets the state of End Point Groups with isAttrBasedEPg value as configured.
	IsAttrBased *string `json:"IsAttrBased,omitempty"`
	// Gets the state of End Point Groups where microsegmentation is present.
	Microsegmentation *string `json:"Microsegmentation,omitempty"`
	// Number of FvRsDomAtt objects per End Point Group with Microsoft configuration.
	MicrosoftUsegCount *int64 `json:"MicrosoftUsegCount,omitempty"`
	// Name value for the End Point Groups present.
	Name *string `json:"Name,omitempty"`
	// Number of objects with Simplified Service Graph Integration with Windows Azure Pack.
	OrchslDevVipCfgCount *int64 `json:"OrchslDevVipCfgCount,omitempty"`
	// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
	RecordType *string `json:"RecordType,omitempty"`
	// Version of record being pushed. This determines what was the API version for data available from the device.
	RecordVersion *string `json:"RecordVersion,omitempty"`
	// The Site name represents an APIC cluster. Service Engine can onboard multiple APIC clusters / sites.
	SiteName *string `json:"SiteName,omitempty"`
	// Logical Operators for attribute based microsegmentation in a hypervisor.
	UsegHypervCount      *int64                               `json:"UsegHypervCount,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryEpgAllOf NiatelemetryEpgAllOf

// NewNiatelemetryEpgAllOf instantiates a new NiatelemetryEpgAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryEpgAllOf(classId string, objectType string) *NiatelemetryEpgAllOf {
	this := NiatelemetryEpgAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryEpgAllOfWithDefaults instantiates a new NiatelemetryEpgAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryEpgAllOfWithDefaults() *NiatelemetryEpgAllOf {
	this := NiatelemetryEpgAllOf{}
	var classId string = "niatelemetry.Epg"
	this.ClassId = classId
	var objectType string = "niatelemetry.Epg"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryEpgAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryEpgAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryEpgAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryEpgAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryEpgAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryEpgAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAzurePackCount returns the AzurePackCount field value if set, zero value otherwise.
func (o *NiatelemetryEpgAllOf) GetAzurePackCount() int64 {
	if o == nil || o.AzurePackCount == nil {
		var ret int64
		return ret
	}
	return *o.AzurePackCount
}

// GetAzurePackCountOk returns a tuple with the AzurePackCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryEpgAllOf) GetAzurePackCountOk() (*int64, bool) {
	if o == nil || o.AzurePackCount == nil {
		return nil, false
	}
	return o.AzurePackCount, true
}

// HasAzurePackCount returns a boolean if a field has been set.
func (o *NiatelemetryEpgAllOf) HasAzurePackCount() bool {
	if o != nil && o.AzurePackCount != nil {
		return true
	}

	return false
}

// SetAzurePackCount gets a reference to the given int64 and assigns it to the AzurePackCount field.
func (o *NiatelemetryEpgAllOf) SetAzurePackCount(v int64) {
	o.AzurePackCount = &v
}

// GetDn returns the Dn field value if set, zero value otherwise.
func (o *NiatelemetryEpgAllOf) GetDn() string {
	if o == nil || o.Dn == nil {
		var ret string
		return ret
	}
	return *o.Dn
}

// GetDnOk returns a tuple with the Dn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryEpgAllOf) GetDnOk() (*string, bool) {
	if o == nil || o.Dn == nil {
		return nil, false
	}
	return o.Dn, true
}

// HasDn returns a boolean if a field has been set.
func (o *NiatelemetryEpgAllOf) HasDn() bool {
	if o != nil && o.Dn != nil {
		return true
	}

	return false
}

// SetDn gets a reference to the given string and assigns it to the Dn field.
func (o *NiatelemetryEpgAllOf) SetDn(v string) {
	o.Dn = &v
}

// GetEpgDelimiterCount returns the EpgDelimiterCount field value if set, zero value otherwise.
func (o *NiatelemetryEpgAllOf) GetEpgDelimiterCount() int64 {
	if o == nil || o.EpgDelimiterCount == nil {
		var ret int64
		return ret
	}
	return *o.EpgDelimiterCount
}

// GetEpgDelimiterCountOk returns a tuple with the EpgDelimiterCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryEpgAllOf) GetEpgDelimiterCountOk() (*int64, bool) {
	if o == nil || o.EpgDelimiterCount == nil {
		return nil, false
	}
	return o.EpgDelimiterCount, true
}

// HasEpgDelimiterCount returns a boolean if a field has been set.
func (o *NiatelemetryEpgAllOf) HasEpgDelimiterCount() bool {
	if o != nil && o.EpgDelimiterCount != nil {
		return true
	}

	return false
}

// SetEpgDelimiterCount gets a reference to the given int64 and assigns it to the EpgDelimiterCount field.
func (o *NiatelemetryEpgAllOf) SetEpgDelimiterCount(v int64) {
	o.EpgDelimiterCount = &v
}

// GetFcNpvCount returns the FcNpvCount field value if set, zero value otherwise.
func (o *NiatelemetryEpgAllOf) GetFcNpvCount() int64 {
	if o == nil || o.FcNpvCount == nil {
		var ret int64
		return ret
	}
	return *o.FcNpvCount
}

// GetFcNpvCountOk returns a tuple with the FcNpvCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryEpgAllOf) GetFcNpvCountOk() (*int64, bool) {
	if o == nil || o.FcNpvCount == nil {
		return nil, false
	}
	return o.FcNpvCount, true
}

// HasFcNpvCount returns a boolean if a field has been set.
func (o *NiatelemetryEpgAllOf) HasFcNpvCount() bool {
	if o != nil && o.FcNpvCount != nil {
		return true
	}

	return false
}

// SetFcNpvCount gets a reference to the given int64 and assigns it to the FcNpvCount field.
func (o *NiatelemetryEpgAllOf) SetFcNpvCount(v int64) {
	o.FcNpvCount = &v
}

// GetFcoeCount returns the FcoeCount field value if set, zero value otherwise.
func (o *NiatelemetryEpgAllOf) GetFcoeCount() int64 {
	if o == nil || o.FcoeCount == nil {
		var ret int64
		return ret
	}
	return *o.FcoeCount
}

// GetFcoeCountOk returns a tuple with the FcoeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryEpgAllOf) GetFcoeCountOk() (*int64, bool) {
	if o == nil || o.FcoeCount == nil {
		return nil, false
	}
	return o.FcoeCount, true
}

// HasFcoeCount returns a boolean if a field has been set.
func (o *NiatelemetryEpgAllOf) HasFcoeCount() bool {
	if o != nil && o.FcoeCount != nil {
		return true
	}

	return false
}

// SetFcoeCount gets a reference to the given int64 and assigns it to the FcoeCount field.
func (o *NiatelemetryEpgAllOf) SetFcoeCount(v int64) {
	o.FcoeCount = &v
}

// GetFvRsDomAttCount returns the FvRsDomAttCount field value if set, zero value otherwise.
func (o *NiatelemetryEpgAllOf) GetFvRsDomAttCount() int64 {
	if o == nil || o.FvRsDomAttCount == nil {
		var ret int64
		return ret
	}
	return *o.FvRsDomAttCount
}

// GetFvRsDomAttCountOk returns a tuple with the FvRsDomAttCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryEpgAllOf) GetFvRsDomAttCountOk() (*int64, bool) {
	if o == nil || o.FvRsDomAttCount == nil {
		return nil, false
	}
	return o.FvRsDomAttCount, true
}

// HasFvRsDomAttCount returns a boolean if a field has been set.
func (o *NiatelemetryEpgAllOf) HasFvRsDomAttCount() bool {
	if o != nil && o.FvRsDomAttCount != nil {
		return true
	}

	return false
}

// SetFvRsDomAttCount gets a reference to the given int64 and assigns it to the FvRsDomAttCount field.
func (o *NiatelemetryEpgAllOf) SetFvRsDomAttCount(v int64) {
	o.FvRsDomAttCount = &v
}

// GetIntraEpgDvsBmCount returns the IntraEpgDvsBmCount field value if set, zero value otherwise.
func (o *NiatelemetryEpgAllOf) GetIntraEpgDvsBmCount() int64 {
	if o == nil || o.IntraEpgDvsBmCount == nil {
		var ret int64
		return ret
	}
	return *o.IntraEpgDvsBmCount
}

// GetIntraEpgDvsBmCountOk returns a tuple with the IntraEpgDvsBmCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryEpgAllOf) GetIntraEpgDvsBmCountOk() (*int64, bool) {
	if o == nil || o.IntraEpgDvsBmCount == nil {
		return nil, false
	}
	return o.IntraEpgDvsBmCount, true
}

// HasIntraEpgDvsBmCount returns a boolean if a field has been set.
func (o *NiatelemetryEpgAllOf) HasIntraEpgDvsBmCount() bool {
	if o != nil && o.IntraEpgDvsBmCount != nil {
		return true
	}

	return false
}

// SetIntraEpgDvsBmCount gets a reference to the given int64 and assigns it to the IntraEpgDvsBmCount field.
func (o *NiatelemetryEpgAllOf) SetIntraEpgDvsBmCount(v int64) {
	o.IntraEpgDvsBmCount = &v
}

// GetIntraEpgHyperv returns the IntraEpgHyperv field value if set, zero value otherwise.
func (o *NiatelemetryEpgAllOf) GetIntraEpgHyperv() string {
	if o == nil || o.IntraEpgHyperv == nil {
		var ret string
		return ret
	}
	return *o.IntraEpgHyperv
}

// GetIntraEpgHypervOk returns a tuple with the IntraEpgHyperv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryEpgAllOf) GetIntraEpgHypervOk() (*string, bool) {
	if o == nil || o.IntraEpgHyperv == nil {
		return nil, false
	}
	return o.IntraEpgHyperv, true
}

// HasIntraEpgHyperv returns a boolean if a field has been set.
func (o *NiatelemetryEpgAllOf) HasIntraEpgHyperv() bool {
	if o != nil && o.IntraEpgHyperv != nil {
		return true
	}

	return false
}

// SetIntraEpgHyperv gets a reference to the given string and assigns it to the IntraEpgHyperv field.
func (o *NiatelemetryEpgAllOf) SetIntraEpgHyperv(v string) {
	o.IntraEpgHyperv = &v
}

// GetIsAttrBased returns the IsAttrBased field value if set, zero value otherwise.
func (o *NiatelemetryEpgAllOf) GetIsAttrBased() string {
	if o == nil || o.IsAttrBased == nil {
		var ret string
		return ret
	}
	return *o.IsAttrBased
}

// GetIsAttrBasedOk returns a tuple with the IsAttrBased field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryEpgAllOf) GetIsAttrBasedOk() (*string, bool) {
	if o == nil || o.IsAttrBased == nil {
		return nil, false
	}
	return o.IsAttrBased, true
}

// HasIsAttrBased returns a boolean if a field has been set.
func (o *NiatelemetryEpgAllOf) HasIsAttrBased() bool {
	if o != nil && o.IsAttrBased != nil {
		return true
	}

	return false
}

// SetIsAttrBased gets a reference to the given string and assigns it to the IsAttrBased field.
func (o *NiatelemetryEpgAllOf) SetIsAttrBased(v string) {
	o.IsAttrBased = &v
}

// GetMicrosegmentation returns the Microsegmentation field value if set, zero value otherwise.
func (o *NiatelemetryEpgAllOf) GetMicrosegmentation() string {
	if o == nil || o.Microsegmentation == nil {
		var ret string
		return ret
	}
	return *o.Microsegmentation
}

// GetMicrosegmentationOk returns a tuple with the Microsegmentation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryEpgAllOf) GetMicrosegmentationOk() (*string, bool) {
	if o == nil || o.Microsegmentation == nil {
		return nil, false
	}
	return o.Microsegmentation, true
}

// HasMicrosegmentation returns a boolean if a field has been set.
func (o *NiatelemetryEpgAllOf) HasMicrosegmentation() bool {
	if o != nil && o.Microsegmentation != nil {
		return true
	}

	return false
}

// SetMicrosegmentation gets a reference to the given string and assigns it to the Microsegmentation field.
func (o *NiatelemetryEpgAllOf) SetMicrosegmentation(v string) {
	o.Microsegmentation = &v
}

// GetMicrosoftUsegCount returns the MicrosoftUsegCount field value if set, zero value otherwise.
func (o *NiatelemetryEpgAllOf) GetMicrosoftUsegCount() int64 {
	if o == nil || o.MicrosoftUsegCount == nil {
		var ret int64
		return ret
	}
	return *o.MicrosoftUsegCount
}

// GetMicrosoftUsegCountOk returns a tuple with the MicrosoftUsegCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryEpgAllOf) GetMicrosoftUsegCountOk() (*int64, bool) {
	if o == nil || o.MicrosoftUsegCount == nil {
		return nil, false
	}
	return o.MicrosoftUsegCount, true
}

// HasMicrosoftUsegCount returns a boolean if a field has been set.
func (o *NiatelemetryEpgAllOf) HasMicrosoftUsegCount() bool {
	if o != nil && o.MicrosoftUsegCount != nil {
		return true
	}

	return false
}

// SetMicrosoftUsegCount gets a reference to the given int64 and assigns it to the MicrosoftUsegCount field.
func (o *NiatelemetryEpgAllOf) SetMicrosoftUsegCount(v int64) {
	o.MicrosoftUsegCount = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NiatelemetryEpgAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryEpgAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NiatelemetryEpgAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NiatelemetryEpgAllOf) SetName(v string) {
	o.Name = &v
}

// GetOrchslDevVipCfgCount returns the OrchslDevVipCfgCount field value if set, zero value otherwise.
func (o *NiatelemetryEpgAllOf) GetOrchslDevVipCfgCount() int64 {
	if o == nil || o.OrchslDevVipCfgCount == nil {
		var ret int64
		return ret
	}
	return *o.OrchslDevVipCfgCount
}

// GetOrchslDevVipCfgCountOk returns a tuple with the OrchslDevVipCfgCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryEpgAllOf) GetOrchslDevVipCfgCountOk() (*int64, bool) {
	if o == nil || o.OrchslDevVipCfgCount == nil {
		return nil, false
	}
	return o.OrchslDevVipCfgCount, true
}

// HasOrchslDevVipCfgCount returns a boolean if a field has been set.
func (o *NiatelemetryEpgAllOf) HasOrchslDevVipCfgCount() bool {
	if o != nil && o.OrchslDevVipCfgCount != nil {
		return true
	}

	return false
}

// SetOrchslDevVipCfgCount gets a reference to the given int64 and assigns it to the OrchslDevVipCfgCount field.
func (o *NiatelemetryEpgAllOf) SetOrchslDevVipCfgCount(v int64) {
	o.OrchslDevVipCfgCount = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *NiatelemetryEpgAllOf) GetRecordType() string {
	if o == nil || o.RecordType == nil {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryEpgAllOf) GetRecordTypeOk() (*string, bool) {
	if o == nil || o.RecordType == nil {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *NiatelemetryEpgAllOf) HasRecordType() bool {
	if o != nil && o.RecordType != nil {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *NiatelemetryEpgAllOf) SetRecordType(v string) {
	o.RecordType = &v
}

// GetRecordVersion returns the RecordVersion field value if set, zero value otherwise.
func (o *NiatelemetryEpgAllOf) GetRecordVersion() string {
	if o == nil || o.RecordVersion == nil {
		var ret string
		return ret
	}
	return *o.RecordVersion
}

// GetRecordVersionOk returns a tuple with the RecordVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryEpgAllOf) GetRecordVersionOk() (*string, bool) {
	if o == nil || o.RecordVersion == nil {
		return nil, false
	}
	return o.RecordVersion, true
}

// HasRecordVersion returns a boolean if a field has been set.
func (o *NiatelemetryEpgAllOf) HasRecordVersion() bool {
	if o != nil && o.RecordVersion != nil {
		return true
	}

	return false
}

// SetRecordVersion gets a reference to the given string and assigns it to the RecordVersion field.
func (o *NiatelemetryEpgAllOf) SetRecordVersion(v string) {
	o.RecordVersion = &v
}

// GetSiteName returns the SiteName field value if set, zero value otherwise.
func (o *NiatelemetryEpgAllOf) GetSiteName() string {
	if o == nil || o.SiteName == nil {
		var ret string
		return ret
	}
	return *o.SiteName
}

// GetSiteNameOk returns a tuple with the SiteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryEpgAllOf) GetSiteNameOk() (*string, bool) {
	if o == nil || o.SiteName == nil {
		return nil, false
	}
	return o.SiteName, true
}

// HasSiteName returns a boolean if a field has been set.
func (o *NiatelemetryEpgAllOf) HasSiteName() bool {
	if o != nil && o.SiteName != nil {
		return true
	}

	return false
}

// SetSiteName gets a reference to the given string and assigns it to the SiteName field.
func (o *NiatelemetryEpgAllOf) SetSiteName(v string) {
	o.SiteName = &v
}

// GetUsegHypervCount returns the UsegHypervCount field value if set, zero value otherwise.
func (o *NiatelemetryEpgAllOf) GetUsegHypervCount() int64 {
	if o == nil || o.UsegHypervCount == nil {
		var ret int64
		return ret
	}
	return *o.UsegHypervCount
}

// GetUsegHypervCountOk returns a tuple with the UsegHypervCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryEpgAllOf) GetUsegHypervCountOk() (*int64, bool) {
	if o == nil || o.UsegHypervCount == nil {
		return nil, false
	}
	return o.UsegHypervCount, true
}

// HasUsegHypervCount returns a boolean if a field has been set.
func (o *NiatelemetryEpgAllOf) HasUsegHypervCount() bool {
	if o != nil && o.UsegHypervCount != nil {
		return true
	}

	return false
}

// SetUsegHypervCount gets a reference to the given int64 and assigns it to the UsegHypervCount field.
func (o *NiatelemetryEpgAllOf) SetUsegHypervCount(v int64) {
	o.UsegHypervCount = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NiatelemetryEpgAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryEpgAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetryEpgAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetryEpgAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NiatelemetryEpgAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AzurePackCount != nil {
		toSerialize["AzurePackCount"] = o.AzurePackCount
	}
	if o.Dn != nil {
		toSerialize["Dn"] = o.Dn
	}
	if o.EpgDelimiterCount != nil {
		toSerialize["EpgDelimiterCount"] = o.EpgDelimiterCount
	}
	if o.FcNpvCount != nil {
		toSerialize["FcNpvCount"] = o.FcNpvCount
	}
	if o.FcoeCount != nil {
		toSerialize["FcoeCount"] = o.FcoeCount
	}
	if o.FvRsDomAttCount != nil {
		toSerialize["FvRsDomAttCount"] = o.FvRsDomAttCount
	}
	if o.IntraEpgDvsBmCount != nil {
		toSerialize["IntraEpgDvsBmCount"] = o.IntraEpgDvsBmCount
	}
	if o.IntraEpgHyperv != nil {
		toSerialize["IntraEpgHyperv"] = o.IntraEpgHyperv
	}
	if o.IsAttrBased != nil {
		toSerialize["IsAttrBased"] = o.IsAttrBased
	}
	if o.Microsegmentation != nil {
		toSerialize["Microsegmentation"] = o.Microsegmentation
	}
	if o.MicrosoftUsegCount != nil {
		toSerialize["MicrosoftUsegCount"] = o.MicrosoftUsegCount
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.OrchslDevVipCfgCount != nil {
		toSerialize["OrchslDevVipCfgCount"] = o.OrchslDevVipCfgCount
	}
	if o.RecordType != nil {
		toSerialize["RecordType"] = o.RecordType
	}
	if o.RecordVersion != nil {
		toSerialize["RecordVersion"] = o.RecordVersion
	}
	if o.SiteName != nil {
		toSerialize["SiteName"] = o.SiteName
	}
	if o.UsegHypervCount != nil {
		toSerialize["UsegHypervCount"] = o.UsegHypervCount
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryEpgAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varNiatelemetryEpgAllOf := _NiatelemetryEpgAllOf{}

	if err = json.Unmarshal(bytes, &varNiatelemetryEpgAllOf); err == nil {
		*o = NiatelemetryEpgAllOf(varNiatelemetryEpgAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AzurePackCount")
		delete(additionalProperties, "Dn")
		delete(additionalProperties, "EpgDelimiterCount")
		delete(additionalProperties, "FcNpvCount")
		delete(additionalProperties, "FcoeCount")
		delete(additionalProperties, "FvRsDomAttCount")
		delete(additionalProperties, "IntraEpgDvsBmCount")
		delete(additionalProperties, "IntraEpgHyperv")
		delete(additionalProperties, "IsAttrBased")
		delete(additionalProperties, "Microsegmentation")
		delete(additionalProperties, "MicrosoftUsegCount")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "OrchslDevVipCfgCount")
		delete(additionalProperties, "RecordType")
		delete(additionalProperties, "RecordVersion")
		delete(additionalProperties, "SiteName")
		delete(additionalProperties, "UsegHypervCount")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNiatelemetryEpgAllOf struct {
	value *NiatelemetryEpgAllOf
	isSet bool
}

func (v NullableNiatelemetryEpgAllOf) Get() *NiatelemetryEpgAllOf {
	return v.value
}

func (v *NullableNiatelemetryEpgAllOf) Set(val *NiatelemetryEpgAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryEpgAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryEpgAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryEpgAllOf(val *NiatelemetryEpgAllOf) *NullableNiatelemetryEpgAllOf {
	return &NullableNiatelemetryEpgAllOf{value: val, isSet: true}
}

func (v NullableNiatelemetryEpgAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryEpgAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
