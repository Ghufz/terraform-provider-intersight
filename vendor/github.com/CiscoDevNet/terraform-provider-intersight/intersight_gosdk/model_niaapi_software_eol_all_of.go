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

// NiaapiSoftwareEolAllOf Definition of the list of properties defined in 'niaapi.SoftwareEol', excluding properties defined in parent classes.
type NiaapiSoftwareEolAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// String contains the Release versions affected by this notice, seperated by comma.
	AffectedVersions *string `json:"AffectedVersions,omitempty"`
	// Date time of this notice Announced.
	AnnouncementDate *time.Time `json:"AnnouncementDate,omitempty"`
	// Epoch time of this notice Announced.
	AnnouncementDateEpoch *int64 `json:"AnnouncementDateEpoch,omitempty"`
	// The bulletinno of this software release end of life notice.
	BulletinNo *string `json:"BulletinNo,omitempty"`
	// The description of this software release end of life notice.
	Description *string `json:"Description,omitempty"`
	// Date time of End of New service attachment .
	EndofNewServiceAttachmentDate *time.Time `json:"EndofNewServiceAttachmentDate,omitempty"`
	// Epoch time of End of New service attachment .
	EndofNewServiceAttachmentDateEpoch *int64 `json:"EndofNewServiceAttachmentDateEpoch,omitempty"`
	// Date time of End of Renewal of service contract .
	EndofServiceContractRenewalDate *time.Time `json:"EndofServiceContractRenewalDate,omitempty"`
	// Epoch time of End of Renewal of service contract .
	EndofServiceContractRenewalDateEpoch *int64 `json:"EndofServiceContractRenewalDateEpoch,omitempty"`
	// Date time of End of Maintenance.
	EndofSwMaintenanceDate *time.Time `json:"EndofSwMaintenanceDate,omitempty"`
	// Epoch time of End of Maintenance.
	EndofSwMaintenanceDateEpoch *int64 `json:"EndofSwMaintenanceDateEpoch,omitempty"`
	// The title of this software release end of life notice.
	Headline *string `json:"Headline,omitempty"`
	// Date time of Last day of Support .
	LastDateofSupport *time.Time `json:"LastDateofSupport,omitempty"`
	// Epoch time of Last day of Support .
	LastDateofSupportEpoch *int64 `json:"LastDateofSupportEpoch,omitempty"`
	// Date time of Lastship Date.
	LastShipDate *time.Time `json:"LastShipDate,omitempty"`
	// Epoch time of Lastship Date.
	LastShipDateEpoch *int64 `json:"LastShipDateEpoch,omitempty"`
	// The URL of this migration notice.
	MigrationUrl *string `json:"MigrationUrl,omitempty"`
	// Software end of life notice URL link to the notice webpage.
	SoftwareEolUrl       *string `json:"SoftwareEolUrl,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiaapiSoftwareEolAllOf NiaapiSoftwareEolAllOf

// NewNiaapiSoftwareEolAllOf instantiates a new NiaapiSoftwareEolAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiaapiSoftwareEolAllOf(classId string, objectType string) *NiaapiSoftwareEolAllOf {
	this := NiaapiSoftwareEolAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiaapiSoftwareEolAllOfWithDefaults instantiates a new NiaapiSoftwareEolAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiaapiSoftwareEolAllOfWithDefaults() *NiaapiSoftwareEolAllOf {
	this := NiaapiSoftwareEolAllOf{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiaapiSoftwareEolAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiaapiSoftwareEolAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiaapiSoftwareEolAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiaapiSoftwareEolAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiaapiSoftwareEolAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiaapiSoftwareEolAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAffectedVersions returns the AffectedVersions field value if set, zero value otherwise.
func (o *NiaapiSoftwareEolAllOf) GetAffectedVersions() string {
	if o == nil || o.AffectedVersions == nil {
		var ret string
		return ret
	}
	return *o.AffectedVersions
}

// GetAffectedVersionsOk returns a tuple with the AffectedVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiSoftwareEolAllOf) GetAffectedVersionsOk() (*string, bool) {
	if o == nil || o.AffectedVersions == nil {
		return nil, false
	}
	return o.AffectedVersions, true
}

// HasAffectedVersions returns a boolean if a field has been set.
func (o *NiaapiSoftwareEolAllOf) HasAffectedVersions() bool {
	if o != nil && o.AffectedVersions != nil {
		return true
	}

	return false
}

// SetAffectedVersions gets a reference to the given string and assigns it to the AffectedVersions field.
func (o *NiaapiSoftwareEolAllOf) SetAffectedVersions(v string) {
	o.AffectedVersions = &v
}

// GetAnnouncementDate returns the AnnouncementDate field value if set, zero value otherwise.
func (o *NiaapiSoftwareEolAllOf) GetAnnouncementDate() time.Time {
	if o == nil || o.AnnouncementDate == nil {
		var ret time.Time
		return ret
	}
	return *o.AnnouncementDate
}

// GetAnnouncementDateOk returns a tuple with the AnnouncementDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiSoftwareEolAllOf) GetAnnouncementDateOk() (*time.Time, bool) {
	if o == nil || o.AnnouncementDate == nil {
		return nil, false
	}
	return o.AnnouncementDate, true
}

// HasAnnouncementDate returns a boolean if a field has been set.
func (o *NiaapiSoftwareEolAllOf) HasAnnouncementDate() bool {
	if o != nil && o.AnnouncementDate != nil {
		return true
	}

	return false
}

// SetAnnouncementDate gets a reference to the given time.Time and assigns it to the AnnouncementDate field.
func (o *NiaapiSoftwareEolAllOf) SetAnnouncementDate(v time.Time) {
	o.AnnouncementDate = &v
}

// GetAnnouncementDateEpoch returns the AnnouncementDateEpoch field value if set, zero value otherwise.
func (o *NiaapiSoftwareEolAllOf) GetAnnouncementDateEpoch() int64 {
	if o == nil || o.AnnouncementDateEpoch == nil {
		var ret int64
		return ret
	}
	return *o.AnnouncementDateEpoch
}

// GetAnnouncementDateEpochOk returns a tuple with the AnnouncementDateEpoch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiSoftwareEolAllOf) GetAnnouncementDateEpochOk() (*int64, bool) {
	if o == nil || o.AnnouncementDateEpoch == nil {
		return nil, false
	}
	return o.AnnouncementDateEpoch, true
}

// HasAnnouncementDateEpoch returns a boolean if a field has been set.
func (o *NiaapiSoftwareEolAllOf) HasAnnouncementDateEpoch() bool {
	if o != nil && o.AnnouncementDateEpoch != nil {
		return true
	}

	return false
}

// SetAnnouncementDateEpoch gets a reference to the given int64 and assigns it to the AnnouncementDateEpoch field.
func (o *NiaapiSoftwareEolAllOf) SetAnnouncementDateEpoch(v int64) {
	o.AnnouncementDateEpoch = &v
}

// GetBulletinNo returns the BulletinNo field value if set, zero value otherwise.
func (o *NiaapiSoftwareEolAllOf) GetBulletinNo() string {
	if o == nil || o.BulletinNo == nil {
		var ret string
		return ret
	}
	return *o.BulletinNo
}

// GetBulletinNoOk returns a tuple with the BulletinNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiSoftwareEolAllOf) GetBulletinNoOk() (*string, bool) {
	if o == nil || o.BulletinNo == nil {
		return nil, false
	}
	return o.BulletinNo, true
}

// HasBulletinNo returns a boolean if a field has been set.
func (o *NiaapiSoftwareEolAllOf) HasBulletinNo() bool {
	if o != nil && o.BulletinNo != nil {
		return true
	}

	return false
}

// SetBulletinNo gets a reference to the given string and assigns it to the BulletinNo field.
func (o *NiaapiSoftwareEolAllOf) SetBulletinNo(v string) {
	o.BulletinNo = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *NiaapiSoftwareEolAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiSoftwareEolAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *NiaapiSoftwareEolAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *NiaapiSoftwareEolAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetEndofNewServiceAttachmentDate returns the EndofNewServiceAttachmentDate field value if set, zero value otherwise.
func (o *NiaapiSoftwareEolAllOf) GetEndofNewServiceAttachmentDate() time.Time {
	if o == nil || o.EndofNewServiceAttachmentDate == nil {
		var ret time.Time
		return ret
	}
	return *o.EndofNewServiceAttachmentDate
}

// GetEndofNewServiceAttachmentDateOk returns a tuple with the EndofNewServiceAttachmentDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiSoftwareEolAllOf) GetEndofNewServiceAttachmentDateOk() (*time.Time, bool) {
	if o == nil || o.EndofNewServiceAttachmentDate == nil {
		return nil, false
	}
	return o.EndofNewServiceAttachmentDate, true
}

// HasEndofNewServiceAttachmentDate returns a boolean if a field has been set.
func (o *NiaapiSoftwareEolAllOf) HasEndofNewServiceAttachmentDate() bool {
	if o != nil && o.EndofNewServiceAttachmentDate != nil {
		return true
	}

	return false
}

// SetEndofNewServiceAttachmentDate gets a reference to the given time.Time and assigns it to the EndofNewServiceAttachmentDate field.
func (o *NiaapiSoftwareEolAllOf) SetEndofNewServiceAttachmentDate(v time.Time) {
	o.EndofNewServiceAttachmentDate = &v
}

// GetEndofNewServiceAttachmentDateEpoch returns the EndofNewServiceAttachmentDateEpoch field value if set, zero value otherwise.
func (o *NiaapiSoftwareEolAllOf) GetEndofNewServiceAttachmentDateEpoch() int64 {
	if o == nil || o.EndofNewServiceAttachmentDateEpoch == nil {
		var ret int64
		return ret
	}
	return *o.EndofNewServiceAttachmentDateEpoch
}

// GetEndofNewServiceAttachmentDateEpochOk returns a tuple with the EndofNewServiceAttachmentDateEpoch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiSoftwareEolAllOf) GetEndofNewServiceAttachmentDateEpochOk() (*int64, bool) {
	if o == nil || o.EndofNewServiceAttachmentDateEpoch == nil {
		return nil, false
	}
	return o.EndofNewServiceAttachmentDateEpoch, true
}

// HasEndofNewServiceAttachmentDateEpoch returns a boolean if a field has been set.
func (o *NiaapiSoftwareEolAllOf) HasEndofNewServiceAttachmentDateEpoch() bool {
	if o != nil && o.EndofNewServiceAttachmentDateEpoch != nil {
		return true
	}

	return false
}

// SetEndofNewServiceAttachmentDateEpoch gets a reference to the given int64 and assigns it to the EndofNewServiceAttachmentDateEpoch field.
func (o *NiaapiSoftwareEolAllOf) SetEndofNewServiceAttachmentDateEpoch(v int64) {
	o.EndofNewServiceAttachmentDateEpoch = &v
}

// GetEndofServiceContractRenewalDate returns the EndofServiceContractRenewalDate field value if set, zero value otherwise.
func (o *NiaapiSoftwareEolAllOf) GetEndofServiceContractRenewalDate() time.Time {
	if o == nil || o.EndofServiceContractRenewalDate == nil {
		var ret time.Time
		return ret
	}
	return *o.EndofServiceContractRenewalDate
}

// GetEndofServiceContractRenewalDateOk returns a tuple with the EndofServiceContractRenewalDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiSoftwareEolAllOf) GetEndofServiceContractRenewalDateOk() (*time.Time, bool) {
	if o == nil || o.EndofServiceContractRenewalDate == nil {
		return nil, false
	}
	return o.EndofServiceContractRenewalDate, true
}

// HasEndofServiceContractRenewalDate returns a boolean if a field has been set.
func (o *NiaapiSoftwareEolAllOf) HasEndofServiceContractRenewalDate() bool {
	if o != nil && o.EndofServiceContractRenewalDate != nil {
		return true
	}

	return false
}

// SetEndofServiceContractRenewalDate gets a reference to the given time.Time and assigns it to the EndofServiceContractRenewalDate field.
func (o *NiaapiSoftwareEolAllOf) SetEndofServiceContractRenewalDate(v time.Time) {
	o.EndofServiceContractRenewalDate = &v
}

// GetEndofServiceContractRenewalDateEpoch returns the EndofServiceContractRenewalDateEpoch field value if set, zero value otherwise.
func (o *NiaapiSoftwareEolAllOf) GetEndofServiceContractRenewalDateEpoch() int64 {
	if o == nil || o.EndofServiceContractRenewalDateEpoch == nil {
		var ret int64
		return ret
	}
	return *o.EndofServiceContractRenewalDateEpoch
}

// GetEndofServiceContractRenewalDateEpochOk returns a tuple with the EndofServiceContractRenewalDateEpoch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiSoftwareEolAllOf) GetEndofServiceContractRenewalDateEpochOk() (*int64, bool) {
	if o == nil || o.EndofServiceContractRenewalDateEpoch == nil {
		return nil, false
	}
	return o.EndofServiceContractRenewalDateEpoch, true
}

// HasEndofServiceContractRenewalDateEpoch returns a boolean if a field has been set.
func (o *NiaapiSoftwareEolAllOf) HasEndofServiceContractRenewalDateEpoch() bool {
	if o != nil && o.EndofServiceContractRenewalDateEpoch != nil {
		return true
	}

	return false
}

// SetEndofServiceContractRenewalDateEpoch gets a reference to the given int64 and assigns it to the EndofServiceContractRenewalDateEpoch field.
func (o *NiaapiSoftwareEolAllOf) SetEndofServiceContractRenewalDateEpoch(v int64) {
	o.EndofServiceContractRenewalDateEpoch = &v
}

// GetEndofSwMaintenanceDate returns the EndofSwMaintenanceDate field value if set, zero value otherwise.
func (o *NiaapiSoftwareEolAllOf) GetEndofSwMaintenanceDate() time.Time {
	if o == nil || o.EndofSwMaintenanceDate == nil {
		var ret time.Time
		return ret
	}
	return *o.EndofSwMaintenanceDate
}

// GetEndofSwMaintenanceDateOk returns a tuple with the EndofSwMaintenanceDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiSoftwareEolAllOf) GetEndofSwMaintenanceDateOk() (*time.Time, bool) {
	if o == nil || o.EndofSwMaintenanceDate == nil {
		return nil, false
	}
	return o.EndofSwMaintenanceDate, true
}

// HasEndofSwMaintenanceDate returns a boolean if a field has been set.
func (o *NiaapiSoftwareEolAllOf) HasEndofSwMaintenanceDate() bool {
	if o != nil && o.EndofSwMaintenanceDate != nil {
		return true
	}

	return false
}

// SetEndofSwMaintenanceDate gets a reference to the given time.Time and assigns it to the EndofSwMaintenanceDate field.
func (o *NiaapiSoftwareEolAllOf) SetEndofSwMaintenanceDate(v time.Time) {
	o.EndofSwMaintenanceDate = &v
}

// GetEndofSwMaintenanceDateEpoch returns the EndofSwMaintenanceDateEpoch field value if set, zero value otherwise.
func (o *NiaapiSoftwareEolAllOf) GetEndofSwMaintenanceDateEpoch() int64 {
	if o == nil || o.EndofSwMaintenanceDateEpoch == nil {
		var ret int64
		return ret
	}
	return *o.EndofSwMaintenanceDateEpoch
}

// GetEndofSwMaintenanceDateEpochOk returns a tuple with the EndofSwMaintenanceDateEpoch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiSoftwareEolAllOf) GetEndofSwMaintenanceDateEpochOk() (*int64, bool) {
	if o == nil || o.EndofSwMaintenanceDateEpoch == nil {
		return nil, false
	}
	return o.EndofSwMaintenanceDateEpoch, true
}

// HasEndofSwMaintenanceDateEpoch returns a boolean if a field has been set.
func (o *NiaapiSoftwareEolAllOf) HasEndofSwMaintenanceDateEpoch() bool {
	if o != nil && o.EndofSwMaintenanceDateEpoch != nil {
		return true
	}

	return false
}

// SetEndofSwMaintenanceDateEpoch gets a reference to the given int64 and assigns it to the EndofSwMaintenanceDateEpoch field.
func (o *NiaapiSoftwareEolAllOf) SetEndofSwMaintenanceDateEpoch(v int64) {
	o.EndofSwMaintenanceDateEpoch = &v
}

// GetHeadline returns the Headline field value if set, zero value otherwise.
func (o *NiaapiSoftwareEolAllOf) GetHeadline() string {
	if o == nil || o.Headline == nil {
		var ret string
		return ret
	}
	return *o.Headline
}

// GetHeadlineOk returns a tuple with the Headline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiSoftwareEolAllOf) GetHeadlineOk() (*string, bool) {
	if o == nil || o.Headline == nil {
		return nil, false
	}
	return o.Headline, true
}

// HasHeadline returns a boolean if a field has been set.
func (o *NiaapiSoftwareEolAllOf) HasHeadline() bool {
	if o != nil && o.Headline != nil {
		return true
	}

	return false
}

// SetHeadline gets a reference to the given string and assigns it to the Headline field.
func (o *NiaapiSoftwareEolAllOf) SetHeadline(v string) {
	o.Headline = &v
}

// GetLastDateofSupport returns the LastDateofSupport field value if set, zero value otherwise.
func (o *NiaapiSoftwareEolAllOf) GetLastDateofSupport() time.Time {
	if o == nil || o.LastDateofSupport == nil {
		var ret time.Time
		return ret
	}
	return *o.LastDateofSupport
}

// GetLastDateofSupportOk returns a tuple with the LastDateofSupport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiSoftwareEolAllOf) GetLastDateofSupportOk() (*time.Time, bool) {
	if o == nil || o.LastDateofSupport == nil {
		return nil, false
	}
	return o.LastDateofSupport, true
}

// HasLastDateofSupport returns a boolean if a field has been set.
func (o *NiaapiSoftwareEolAllOf) HasLastDateofSupport() bool {
	if o != nil && o.LastDateofSupport != nil {
		return true
	}

	return false
}

// SetLastDateofSupport gets a reference to the given time.Time and assigns it to the LastDateofSupport field.
func (o *NiaapiSoftwareEolAllOf) SetLastDateofSupport(v time.Time) {
	o.LastDateofSupport = &v
}

// GetLastDateofSupportEpoch returns the LastDateofSupportEpoch field value if set, zero value otherwise.
func (o *NiaapiSoftwareEolAllOf) GetLastDateofSupportEpoch() int64 {
	if o == nil || o.LastDateofSupportEpoch == nil {
		var ret int64
		return ret
	}
	return *o.LastDateofSupportEpoch
}

// GetLastDateofSupportEpochOk returns a tuple with the LastDateofSupportEpoch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiSoftwareEolAllOf) GetLastDateofSupportEpochOk() (*int64, bool) {
	if o == nil || o.LastDateofSupportEpoch == nil {
		return nil, false
	}
	return o.LastDateofSupportEpoch, true
}

// HasLastDateofSupportEpoch returns a boolean if a field has been set.
func (o *NiaapiSoftwareEolAllOf) HasLastDateofSupportEpoch() bool {
	if o != nil && o.LastDateofSupportEpoch != nil {
		return true
	}

	return false
}

// SetLastDateofSupportEpoch gets a reference to the given int64 and assigns it to the LastDateofSupportEpoch field.
func (o *NiaapiSoftwareEolAllOf) SetLastDateofSupportEpoch(v int64) {
	o.LastDateofSupportEpoch = &v
}

// GetLastShipDate returns the LastShipDate field value if set, zero value otherwise.
func (o *NiaapiSoftwareEolAllOf) GetLastShipDate() time.Time {
	if o == nil || o.LastShipDate == nil {
		var ret time.Time
		return ret
	}
	return *o.LastShipDate
}

// GetLastShipDateOk returns a tuple with the LastShipDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiSoftwareEolAllOf) GetLastShipDateOk() (*time.Time, bool) {
	if o == nil || o.LastShipDate == nil {
		return nil, false
	}
	return o.LastShipDate, true
}

// HasLastShipDate returns a boolean if a field has been set.
func (o *NiaapiSoftwareEolAllOf) HasLastShipDate() bool {
	if o != nil && o.LastShipDate != nil {
		return true
	}

	return false
}

// SetLastShipDate gets a reference to the given time.Time and assigns it to the LastShipDate field.
func (o *NiaapiSoftwareEolAllOf) SetLastShipDate(v time.Time) {
	o.LastShipDate = &v
}

// GetLastShipDateEpoch returns the LastShipDateEpoch field value if set, zero value otherwise.
func (o *NiaapiSoftwareEolAllOf) GetLastShipDateEpoch() int64 {
	if o == nil || o.LastShipDateEpoch == nil {
		var ret int64
		return ret
	}
	return *o.LastShipDateEpoch
}

// GetLastShipDateEpochOk returns a tuple with the LastShipDateEpoch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiSoftwareEolAllOf) GetLastShipDateEpochOk() (*int64, bool) {
	if o == nil || o.LastShipDateEpoch == nil {
		return nil, false
	}
	return o.LastShipDateEpoch, true
}

// HasLastShipDateEpoch returns a boolean if a field has been set.
func (o *NiaapiSoftwareEolAllOf) HasLastShipDateEpoch() bool {
	if o != nil && o.LastShipDateEpoch != nil {
		return true
	}

	return false
}

// SetLastShipDateEpoch gets a reference to the given int64 and assigns it to the LastShipDateEpoch field.
func (o *NiaapiSoftwareEolAllOf) SetLastShipDateEpoch(v int64) {
	o.LastShipDateEpoch = &v
}

// GetMigrationUrl returns the MigrationUrl field value if set, zero value otherwise.
func (o *NiaapiSoftwareEolAllOf) GetMigrationUrl() string {
	if o == nil || o.MigrationUrl == nil {
		var ret string
		return ret
	}
	return *o.MigrationUrl
}

// GetMigrationUrlOk returns a tuple with the MigrationUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiSoftwareEolAllOf) GetMigrationUrlOk() (*string, bool) {
	if o == nil || o.MigrationUrl == nil {
		return nil, false
	}
	return o.MigrationUrl, true
}

// HasMigrationUrl returns a boolean if a field has been set.
func (o *NiaapiSoftwareEolAllOf) HasMigrationUrl() bool {
	if o != nil && o.MigrationUrl != nil {
		return true
	}

	return false
}

// SetMigrationUrl gets a reference to the given string and assigns it to the MigrationUrl field.
func (o *NiaapiSoftwareEolAllOf) SetMigrationUrl(v string) {
	o.MigrationUrl = &v
}

// GetSoftwareEolUrl returns the SoftwareEolUrl field value if set, zero value otherwise.
func (o *NiaapiSoftwareEolAllOf) GetSoftwareEolUrl() string {
	if o == nil || o.SoftwareEolUrl == nil {
		var ret string
		return ret
	}
	return *o.SoftwareEolUrl
}

// GetSoftwareEolUrlOk returns a tuple with the SoftwareEolUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiSoftwareEolAllOf) GetSoftwareEolUrlOk() (*string, bool) {
	if o == nil || o.SoftwareEolUrl == nil {
		return nil, false
	}
	return o.SoftwareEolUrl, true
}

// HasSoftwareEolUrl returns a boolean if a field has been set.
func (o *NiaapiSoftwareEolAllOf) HasSoftwareEolUrl() bool {
	if o != nil && o.SoftwareEolUrl != nil {
		return true
	}

	return false
}

// SetSoftwareEolUrl gets a reference to the given string and assigns it to the SoftwareEolUrl field.
func (o *NiaapiSoftwareEolAllOf) SetSoftwareEolUrl(v string) {
	o.SoftwareEolUrl = &v
}

func (o NiaapiSoftwareEolAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AffectedVersions != nil {
		toSerialize["AffectedVersions"] = o.AffectedVersions
	}
	if o.AnnouncementDate != nil {
		toSerialize["AnnouncementDate"] = o.AnnouncementDate
	}
	if o.AnnouncementDateEpoch != nil {
		toSerialize["AnnouncementDateEpoch"] = o.AnnouncementDateEpoch
	}
	if o.BulletinNo != nil {
		toSerialize["BulletinNo"] = o.BulletinNo
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.EndofNewServiceAttachmentDate != nil {
		toSerialize["EndofNewServiceAttachmentDate"] = o.EndofNewServiceAttachmentDate
	}
	if o.EndofNewServiceAttachmentDateEpoch != nil {
		toSerialize["EndofNewServiceAttachmentDateEpoch"] = o.EndofNewServiceAttachmentDateEpoch
	}
	if o.EndofServiceContractRenewalDate != nil {
		toSerialize["EndofServiceContractRenewalDate"] = o.EndofServiceContractRenewalDate
	}
	if o.EndofServiceContractRenewalDateEpoch != nil {
		toSerialize["EndofServiceContractRenewalDateEpoch"] = o.EndofServiceContractRenewalDateEpoch
	}
	if o.EndofSwMaintenanceDate != nil {
		toSerialize["EndofSwMaintenanceDate"] = o.EndofSwMaintenanceDate
	}
	if o.EndofSwMaintenanceDateEpoch != nil {
		toSerialize["EndofSwMaintenanceDateEpoch"] = o.EndofSwMaintenanceDateEpoch
	}
	if o.Headline != nil {
		toSerialize["Headline"] = o.Headline
	}
	if o.LastDateofSupport != nil {
		toSerialize["LastDateofSupport"] = o.LastDateofSupport
	}
	if o.LastDateofSupportEpoch != nil {
		toSerialize["LastDateofSupportEpoch"] = o.LastDateofSupportEpoch
	}
	if o.LastShipDate != nil {
		toSerialize["LastShipDate"] = o.LastShipDate
	}
	if o.LastShipDateEpoch != nil {
		toSerialize["LastShipDateEpoch"] = o.LastShipDateEpoch
	}
	if o.MigrationUrl != nil {
		toSerialize["MigrationUrl"] = o.MigrationUrl
	}
	if o.SoftwareEolUrl != nil {
		toSerialize["SoftwareEolUrl"] = o.SoftwareEolUrl
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiaapiSoftwareEolAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varNiaapiSoftwareEolAllOf := _NiaapiSoftwareEolAllOf{}

	if err = json.Unmarshal(bytes, &varNiaapiSoftwareEolAllOf); err == nil {
		*o = NiaapiSoftwareEolAllOf(varNiaapiSoftwareEolAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AffectedVersions")
		delete(additionalProperties, "AnnouncementDate")
		delete(additionalProperties, "AnnouncementDateEpoch")
		delete(additionalProperties, "BulletinNo")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "EndofNewServiceAttachmentDate")
		delete(additionalProperties, "EndofNewServiceAttachmentDateEpoch")
		delete(additionalProperties, "EndofServiceContractRenewalDate")
		delete(additionalProperties, "EndofServiceContractRenewalDateEpoch")
		delete(additionalProperties, "EndofSwMaintenanceDate")
		delete(additionalProperties, "EndofSwMaintenanceDateEpoch")
		delete(additionalProperties, "Headline")
		delete(additionalProperties, "LastDateofSupport")
		delete(additionalProperties, "LastDateofSupportEpoch")
		delete(additionalProperties, "LastShipDate")
		delete(additionalProperties, "LastShipDateEpoch")
		delete(additionalProperties, "MigrationUrl")
		delete(additionalProperties, "SoftwareEolUrl")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNiaapiSoftwareEolAllOf struct {
	value *NiaapiSoftwareEolAllOf
	isSet bool
}

func (v NullableNiaapiSoftwareEolAllOf) Get() *NiaapiSoftwareEolAllOf {
	return v.value
}

func (v *NullableNiaapiSoftwareEolAllOf) Set(val *NiaapiSoftwareEolAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNiaapiSoftwareEolAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNiaapiSoftwareEolAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiaapiSoftwareEolAllOf(val *NiaapiSoftwareEolAllOf) *NullableNiaapiSoftwareEolAllOf {
	return &NullableNiaapiSoftwareEolAllOf{value: val, isSet: true}
}

func (v NullableNiaapiSoftwareEolAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiaapiSoftwareEolAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
