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

// NiatelemetryApicCoreFileDetailsAllOf Definition of the list of properties defined in 'niatelemetry.ApicCoreFileDetails', excluding properties defined in parent classes.
type NiatelemetryApicCoreFileDetailsAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Annotation of the Core file in APIC.
	Annotation *string `json:"Annotation,omitempty"`
	// Child action of the Core file in APIC.
	ChildAction *string `json:"ChildAction,omitempty"`
	// Collection Time of the Core file in APIC.
	CollectionTime *string `json:"CollectionTime,omitempty"`
	// Data type of the Core file in APIC.
	DataType *string `json:"DataType,omitempty"`
	// Dn for the Core file in the inventory.
	Dn *string `json:"Dn,omitempty"`
	// Export file URI of the Core file in APIC.
	ExportFileUri *string `json:"ExportFileUri,omitempty"`
	// Export status of the Core file in APIC.
	ExportStatus *string `json:"ExportStatus,omitempty"`
	// Export status str of the Core file in APIC.
	ExportStatusStr *int64 `json:"ExportStatusStr,omitempty"`
	// Export tech Sup file URI of the Core file in APIC.
	ExportTechSupFileUri *string `json:"ExportTechSupFileUri,omitempty"`
	// Return if file is exported To Controller or not in APIC.
	ExportedToController *string `json:"ExportedToController,omitempty"`
	// Ext Mngd By of the Core file in APIC.
	ExtMngdBy *string `json:"ExtMngdBy,omitempty"`
	// File size of the Core file in APIC.
	FileSize *int64 `json:"FileSize,omitempty"`
	// Host Name of the Core file in APIC.
	HostName *string `json:"HostName,omitempty"`
	// Lc owner of the Core file in APIC.
	LcOwn *string `json:"LcOwn,omitempty"`
	// Mod Ts of the Core file in APIC.
	ModTs *int64 `json:"ModTs,omitempty"`
	// Node Id of the Core file in APIC.
	NodeId *string `json:"NodeId,omitempty"`
	// Pol Name of the Core file in APIC.
	PolName *string `json:"PolName,omitempty"`
	// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
	RecordType *string `json:"RecordType,omitempty"`
	// Version of record being pushed. This determines what was the API version for data available from the device.
	RecordVersion *string `json:"RecordVersion,omitempty"`
	// Name of the APIC site from which this data is being collected.
	SiteName *string `json:"SiteName,omitempty"`
	// Status of the Core file in APIC.
	Status *string `json:"Status,omitempty"`
	// UId of the Core file in the APIC.
	Uid *string `json:"Uid,omitempty"`
	// User dom of the Core file in APIC.
	Userdom              *string                              `json:"Userdom,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryApicCoreFileDetailsAllOf NiatelemetryApicCoreFileDetailsAllOf

// NewNiatelemetryApicCoreFileDetailsAllOf instantiates a new NiatelemetryApicCoreFileDetailsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryApicCoreFileDetailsAllOf(classId string, objectType string) *NiatelemetryApicCoreFileDetailsAllOf {
	this := NiatelemetryApicCoreFileDetailsAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryApicCoreFileDetailsAllOfWithDefaults instantiates a new NiatelemetryApicCoreFileDetailsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryApicCoreFileDetailsAllOfWithDefaults() *NiatelemetryApicCoreFileDetailsAllOf {
	this := NiatelemetryApicCoreFileDetailsAllOf{}
	var classId string = "niatelemetry.ApicCoreFileDetails"
	this.ClassId = classId
	var objectType string = "niatelemetry.ApicCoreFileDetails"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryApicCoreFileDetailsAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicCoreFileDetailsAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryApicCoreFileDetailsAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryApicCoreFileDetailsAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicCoreFileDetailsAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryApicCoreFileDetailsAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAnnotation returns the Annotation field value if set, zero value otherwise.
func (o *NiatelemetryApicCoreFileDetailsAllOf) GetAnnotation() string {
	if o == nil || o.Annotation == nil {
		var ret string
		return ret
	}
	return *o.Annotation
}

// GetAnnotationOk returns a tuple with the Annotation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicCoreFileDetailsAllOf) GetAnnotationOk() (*string, bool) {
	if o == nil || o.Annotation == nil {
		return nil, false
	}
	return o.Annotation, true
}

// HasAnnotation returns a boolean if a field has been set.
func (o *NiatelemetryApicCoreFileDetailsAllOf) HasAnnotation() bool {
	if o != nil && o.Annotation != nil {
		return true
	}

	return false
}

// SetAnnotation gets a reference to the given string and assigns it to the Annotation field.
func (o *NiatelemetryApicCoreFileDetailsAllOf) SetAnnotation(v string) {
	o.Annotation = &v
}

// GetChildAction returns the ChildAction field value if set, zero value otherwise.
func (o *NiatelemetryApicCoreFileDetailsAllOf) GetChildAction() string {
	if o == nil || o.ChildAction == nil {
		var ret string
		return ret
	}
	return *o.ChildAction
}

// GetChildActionOk returns a tuple with the ChildAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicCoreFileDetailsAllOf) GetChildActionOk() (*string, bool) {
	if o == nil || o.ChildAction == nil {
		return nil, false
	}
	return o.ChildAction, true
}

// HasChildAction returns a boolean if a field has been set.
func (o *NiatelemetryApicCoreFileDetailsAllOf) HasChildAction() bool {
	if o != nil && o.ChildAction != nil {
		return true
	}

	return false
}

// SetChildAction gets a reference to the given string and assigns it to the ChildAction field.
func (o *NiatelemetryApicCoreFileDetailsAllOf) SetChildAction(v string) {
	o.ChildAction = &v
}

// GetCollectionTime returns the CollectionTime field value if set, zero value otherwise.
func (o *NiatelemetryApicCoreFileDetailsAllOf) GetCollectionTime() string {
	if o == nil || o.CollectionTime == nil {
		var ret string
		return ret
	}
	return *o.CollectionTime
}

// GetCollectionTimeOk returns a tuple with the CollectionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicCoreFileDetailsAllOf) GetCollectionTimeOk() (*string, bool) {
	if o == nil || o.CollectionTime == nil {
		return nil, false
	}
	return o.CollectionTime, true
}

// HasCollectionTime returns a boolean if a field has been set.
func (o *NiatelemetryApicCoreFileDetailsAllOf) HasCollectionTime() bool {
	if o != nil && o.CollectionTime != nil {
		return true
	}

	return false
}

// SetCollectionTime gets a reference to the given string and assigns it to the CollectionTime field.
func (o *NiatelemetryApicCoreFileDetailsAllOf) SetCollectionTime(v string) {
	o.CollectionTime = &v
}

// GetDataType returns the DataType field value if set, zero value otherwise.
func (o *NiatelemetryApicCoreFileDetailsAllOf) GetDataType() string {
	if o == nil || o.DataType == nil {
		var ret string
		return ret
	}
	return *o.DataType
}

// GetDataTypeOk returns a tuple with the DataType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicCoreFileDetailsAllOf) GetDataTypeOk() (*string, bool) {
	if o == nil || o.DataType == nil {
		return nil, false
	}
	return o.DataType, true
}

// HasDataType returns a boolean if a field has been set.
func (o *NiatelemetryApicCoreFileDetailsAllOf) HasDataType() bool {
	if o != nil && o.DataType != nil {
		return true
	}

	return false
}

// SetDataType gets a reference to the given string and assigns it to the DataType field.
func (o *NiatelemetryApicCoreFileDetailsAllOf) SetDataType(v string) {
	o.DataType = &v
}

// GetDn returns the Dn field value if set, zero value otherwise.
func (o *NiatelemetryApicCoreFileDetailsAllOf) GetDn() string {
	if o == nil || o.Dn == nil {
		var ret string
		return ret
	}
	return *o.Dn
}

// GetDnOk returns a tuple with the Dn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicCoreFileDetailsAllOf) GetDnOk() (*string, bool) {
	if o == nil || o.Dn == nil {
		return nil, false
	}
	return o.Dn, true
}

// HasDn returns a boolean if a field has been set.
func (o *NiatelemetryApicCoreFileDetailsAllOf) HasDn() bool {
	if o != nil && o.Dn != nil {
		return true
	}

	return false
}

// SetDn gets a reference to the given string and assigns it to the Dn field.
func (o *NiatelemetryApicCoreFileDetailsAllOf) SetDn(v string) {
	o.Dn = &v
}

// GetExportFileUri returns the ExportFileUri field value if set, zero value otherwise.
func (o *NiatelemetryApicCoreFileDetailsAllOf) GetExportFileUri() string {
	if o == nil || o.ExportFileUri == nil {
		var ret string
		return ret
	}
	return *o.ExportFileUri
}

// GetExportFileUriOk returns a tuple with the ExportFileUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicCoreFileDetailsAllOf) GetExportFileUriOk() (*string, bool) {
	if o == nil || o.ExportFileUri == nil {
		return nil, false
	}
	return o.ExportFileUri, true
}

// HasExportFileUri returns a boolean if a field has been set.
func (o *NiatelemetryApicCoreFileDetailsAllOf) HasExportFileUri() bool {
	if o != nil && o.ExportFileUri != nil {
		return true
	}

	return false
}

// SetExportFileUri gets a reference to the given string and assigns it to the ExportFileUri field.
func (o *NiatelemetryApicCoreFileDetailsAllOf) SetExportFileUri(v string) {
	o.ExportFileUri = &v
}

// GetExportStatus returns the ExportStatus field value if set, zero value otherwise.
func (o *NiatelemetryApicCoreFileDetailsAllOf) GetExportStatus() string {
	if o == nil || o.ExportStatus == nil {
		var ret string
		return ret
	}
	return *o.ExportStatus
}

// GetExportStatusOk returns a tuple with the ExportStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicCoreFileDetailsAllOf) GetExportStatusOk() (*string, bool) {
	if o == nil || o.ExportStatus == nil {
		return nil, false
	}
	return o.ExportStatus, true
}

// HasExportStatus returns a boolean if a field has been set.
func (o *NiatelemetryApicCoreFileDetailsAllOf) HasExportStatus() bool {
	if o != nil && o.ExportStatus != nil {
		return true
	}

	return false
}

// SetExportStatus gets a reference to the given string and assigns it to the ExportStatus field.
func (o *NiatelemetryApicCoreFileDetailsAllOf) SetExportStatus(v string) {
	o.ExportStatus = &v
}

// GetExportStatusStr returns the ExportStatusStr field value if set, zero value otherwise.
func (o *NiatelemetryApicCoreFileDetailsAllOf) GetExportStatusStr() int64 {
	if o == nil || o.ExportStatusStr == nil {
		var ret int64
		return ret
	}
	return *o.ExportStatusStr
}

// GetExportStatusStrOk returns a tuple with the ExportStatusStr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicCoreFileDetailsAllOf) GetExportStatusStrOk() (*int64, bool) {
	if o == nil || o.ExportStatusStr == nil {
		return nil, false
	}
	return o.ExportStatusStr, true
}

// HasExportStatusStr returns a boolean if a field has been set.
func (o *NiatelemetryApicCoreFileDetailsAllOf) HasExportStatusStr() bool {
	if o != nil && o.ExportStatusStr != nil {
		return true
	}

	return false
}

// SetExportStatusStr gets a reference to the given int64 and assigns it to the ExportStatusStr field.
func (o *NiatelemetryApicCoreFileDetailsAllOf) SetExportStatusStr(v int64) {
	o.ExportStatusStr = &v
}

// GetExportTechSupFileUri returns the ExportTechSupFileUri field value if set, zero value otherwise.
func (o *NiatelemetryApicCoreFileDetailsAllOf) GetExportTechSupFileUri() string {
	if o == nil || o.ExportTechSupFileUri == nil {
		var ret string
		return ret
	}
	return *o.ExportTechSupFileUri
}

// GetExportTechSupFileUriOk returns a tuple with the ExportTechSupFileUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicCoreFileDetailsAllOf) GetExportTechSupFileUriOk() (*string, bool) {
	if o == nil || o.ExportTechSupFileUri == nil {
		return nil, false
	}
	return o.ExportTechSupFileUri, true
}

// HasExportTechSupFileUri returns a boolean if a field has been set.
func (o *NiatelemetryApicCoreFileDetailsAllOf) HasExportTechSupFileUri() bool {
	if o != nil && o.ExportTechSupFileUri != nil {
		return true
	}

	return false
}

// SetExportTechSupFileUri gets a reference to the given string and assigns it to the ExportTechSupFileUri field.
func (o *NiatelemetryApicCoreFileDetailsAllOf) SetExportTechSupFileUri(v string) {
	o.ExportTechSupFileUri = &v
}

// GetExportedToController returns the ExportedToController field value if set, zero value otherwise.
func (o *NiatelemetryApicCoreFileDetailsAllOf) GetExportedToController() string {
	if o == nil || o.ExportedToController == nil {
		var ret string
		return ret
	}
	return *o.ExportedToController
}

// GetExportedToControllerOk returns a tuple with the ExportedToController field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicCoreFileDetailsAllOf) GetExportedToControllerOk() (*string, bool) {
	if o == nil || o.ExportedToController == nil {
		return nil, false
	}
	return o.ExportedToController, true
}

// HasExportedToController returns a boolean if a field has been set.
func (o *NiatelemetryApicCoreFileDetailsAllOf) HasExportedToController() bool {
	if o != nil && o.ExportedToController != nil {
		return true
	}

	return false
}

// SetExportedToController gets a reference to the given string and assigns it to the ExportedToController field.
func (o *NiatelemetryApicCoreFileDetailsAllOf) SetExportedToController(v string) {
	o.ExportedToController = &v
}

// GetExtMngdBy returns the ExtMngdBy field value if set, zero value otherwise.
func (o *NiatelemetryApicCoreFileDetailsAllOf) GetExtMngdBy() string {
	if o == nil || o.ExtMngdBy == nil {
		var ret string
		return ret
	}
	return *o.ExtMngdBy
}

// GetExtMngdByOk returns a tuple with the ExtMngdBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicCoreFileDetailsAllOf) GetExtMngdByOk() (*string, bool) {
	if o == nil || o.ExtMngdBy == nil {
		return nil, false
	}
	return o.ExtMngdBy, true
}

// HasExtMngdBy returns a boolean if a field has been set.
func (o *NiatelemetryApicCoreFileDetailsAllOf) HasExtMngdBy() bool {
	if o != nil && o.ExtMngdBy != nil {
		return true
	}

	return false
}

// SetExtMngdBy gets a reference to the given string and assigns it to the ExtMngdBy field.
func (o *NiatelemetryApicCoreFileDetailsAllOf) SetExtMngdBy(v string) {
	o.ExtMngdBy = &v
}

// GetFileSize returns the FileSize field value if set, zero value otherwise.
func (o *NiatelemetryApicCoreFileDetailsAllOf) GetFileSize() int64 {
	if o == nil || o.FileSize == nil {
		var ret int64
		return ret
	}
	return *o.FileSize
}

// GetFileSizeOk returns a tuple with the FileSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicCoreFileDetailsAllOf) GetFileSizeOk() (*int64, bool) {
	if o == nil || o.FileSize == nil {
		return nil, false
	}
	return o.FileSize, true
}

// HasFileSize returns a boolean if a field has been set.
func (o *NiatelemetryApicCoreFileDetailsAllOf) HasFileSize() bool {
	if o != nil && o.FileSize != nil {
		return true
	}

	return false
}

// SetFileSize gets a reference to the given int64 and assigns it to the FileSize field.
func (o *NiatelemetryApicCoreFileDetailsAllOf) SetFileSize(v int64) {
	o.FileSize = &v
}

// GetHostName returns the HostName field value if set, zero value otherwise.
func (o *NiatelemetryApicCoreFileDetailsAllOf) GetHostName() string {
	if o == nil || o.HostName == nil {
		var ret string
		return ret
	}
	return *o.HostName
}

// GetHostNameOk returns a tuple with the HostName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicCoreFileDetailsAllOf) GetHostNameOk() (*string, bool) {
	if o == nil || o.HostName == nil {
		return nil, false
	}
	return o.HostName, true
}

// HasHostName returns a boolean if a field has been set.
func (o *NiatelemetryApicCoreFileDetailsAllOf) HasHostName() bool {
	if o != nil && o.HostName != nil {
		return true
	}

	return false
}

// SetHostName gets a reference to the given string and assigns it to the HostName field.
func (o *NiatelemetryApicCoreFileDetailsAllOf) SetHostName(v string) {
	o.HostName = &v
}

// GetLcOwn returns the LcOwn field value if set, zero value otherwise.
func (o *NiatelemetryApicCoreFileDetailsAllOf) GetLcOwn() string {
	if o == nil || o.LcOwn == nil {
		var ret string
		return ret
	}
	return *o.LcOwn
}

// GetLcOwnOk returns a tuple with the LcOwn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicCoreFileDetailsAllOf) GetLcOwnOk() (*string, bool) {
	if o == nil || o.LcOwn == nil {
		return nil, false
	}
	return o.LcOwn, true
}

// HasLcOwn returns a boolean if a field has been set.
func (o *NiatelemetryApicCoreFileDetailsAllOf) HasLcOwn() bool {
	if o != nil && o.LcOwn != nil {
		return true
	}

	return false
}

// SetLcOwn gets a reference to the given string and assigns it to the LcOwn field.
func (o *NiatelemetryApicCoreFileDetailsAllOf) SetLcOwn(v string) {
	o.LcOwn = &v
}

// GetModTs returns the ModTs field value if set, zero value otherwise.
func (o *NiatelemetryApicCoreFileDetailsAllOf) GetModTs() int64 {
	if o == nil || o.ModTs == nil {
		var ret int64
		return ret
	}
	return *o.ModTs
}

// GetModTsOk returns a tuple with the ModTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicCoreFileDetailsAllOf) GetModTsOk() (*int64, bool) {
	if o == nil || o.ModTs == nil {
		return nil, false
	}
	return o.ModTs, true
}

// HasModTs returns a boolean if a field has been set.
func (o *NiatelemetryApicCoreFileDetailsAllOf) HasModTs() bool {
	if o != nil && o.ModTs != nil {
		return true
	}

	return false
}

// SetModTs gets a reference to the given int64 and assigns it to the ModTs field.
func (o *NiatelemetryApicCoreFileDetailsAllOf) SetModTs(v int64) {
	o.ModTs = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *NiatelemetryApicCoreFileDetailsAllOf) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicCoreFileDetailsAllOf) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *NiatelemetryApicCoreFileDetailsAllOf) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *NiatelemetryApicCoreFileDetailsAllOf) SetNodeId(v string) {
	o.NodeId = &v
}

// GetPolName returns the PolName field value if set, zero value otherwise.
func (o *NiatelemetryApicCoreFileDetailsAllOf) GetPolName() string {
	if o == nil || o.PolName == nil {
		var ret string
		return ret
	}
	return *o.PolName
}

// GetPolNameOk returns a tuple with the PolName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicCoreFileDetailsAllOf) GetPolNameOk() (*string, bool) {
	if o == nil || o.PolName == nil {
		return nil, false
	}
	return o.PolName, true
}

// HasPolName returns a boolean if a field has been set.
func (o *NiatelemetryApicCoreFileDetailsAllOf) HasPolName() bool {
	if o != nil && o.PolName != nil {
		return true
	}

	return false
}

// SetPolName gets a reference to the given string and assigns it to the PolName field.
func (o *NiatelemetryApicCoreFileDetailsAllOf) SetPolName(v string) {
	o.PolName = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *NiatelemetryApicCoreFileDetailsAllOf) GetRecordType() string {
	if o == nil || o.RecordType == nil {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicCoreFileDetailsAllOf) GetRecordTypeOk() (*string, bool) {
	if o == nil || o.RecordType == nil {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *NiatelemetryApicCoreFileDetailsAllOf) HasRecordType() bool {
	if o != nil && o.RecordType != nil {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *NiatelemetryApicCoreFileDetailsAllOf) SetRecordType(v string) {
	o.RecordType = &v
}

// GetRecordVersion returns the RecordVersion field value if set, zero value otherwise.
func (o *NiatelemetryApicCoreFileDetailsAllOf) GetRecordVersion() string {
	if o == nil || o.RecordVersion == nil {
		var ret string
		return ret
	}
	return *o.RecordVersion
}

// GetRecordVersionOk returns a tuple with the RecordVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicCoreFileDetailsAllOf) GetRecordVersionOk() (*string, bool) {
	if o == nil || o.RecordVersion == nil {
		return nil, false
	}
	return o.RecordVersion, true
}

// HasRecordVersion returns a boolean if a field has been set.
func (o *NiatelemetryApicCoreFileDetailsAllOf) HasRecordVersion() bool {
	if o != nil && o.RecordVersion != nil {
		return true
	}

	return false
}

// SetRecordVersion gets a reference to the given string and assigns it to the RecordVersion field.
func (o *NiatelemetryApicCoreFileDetailsAllOf) SetRecordVersion(v string) {
	o.RecordVersion = &v
}

// GetSiteName returns the SiteName field value if set, zero value otherwise.
func (o *NiatelemetryApicCoreFileDetailsAllOf) GetSiteName() string {
	if o == nil || o.SiteName == nil {
		var ret string
		return ret
	}
	return *o.SiteName
}

// GetSiteNameOk returns a tuple with the SiteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicCoreFileDetailsAllOf) GetSiteNameOk() (*string, bool) {
	if o == nil || o.SiteName == nil {
		return nil, false
	}
	return o.SiteName, true
}

// HasSiteName returns a boolean if a field has been set.
func (o *NiatelemetryApicCoreFileDetailsAllOf) HasSiteName() bool {
	if o != nil && o.SiteName != nil {
		return true
	}

	return false
}

// SetSiteName gets a reference to the given string and assigns it to the SiteName field.
func (o *NiatelemetryApicCoreFileDetailsAllOf) SetSiteName(v string) {
	o.SiteName = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *NiatelemetryApicCoreFileDetailsAllOf) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicCoreFileDetailsAllOf) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *NiatelemetryApicCoreFileDetailsAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *NiatelemetryApicCoreFileDetailsAllOf) SetStatus(v string) {
	o.Status = &v
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *NiatelemetryApicCoreFileDetailsAllOf) GetUid() string {
	if o == nil || o.Uid == nil {
		var ret string
		return ret
	}
	return *o.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicCoreFileDetailsAllOf) GetUidOk() (*string, bool) {
	if o == nil || o.Uid == nil {
		return nil, false
	}
	return o.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *NiatelemetryApicCoreFileDetailsAllOf) HasUid() bool {
	if o != nil && o.Uid != nil {
		return true
	}

	return false
}

// SetUid gets a reference to the given string and assigns it to the Uid field.
func (o *NiatelemetryApicCoreFileDetailsAllOf) SetUid(v string) {
	o.Uid = &v
}

// GetUserdom returns the Userdom field value if set, zero value otherwise.
func (o *NiatelemetryApicCoreFileDetailsAllOf) GetUserdom() string {
	if o == nil || o.Userdom == nil {
		var ret string
		return ret
	}
	return *o.Userdom
}

// GetUserdomOk returns a tuple with the Userdom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicCoreFileDetailsAllOf) GetUserdomOk() (*string, bool) {
	if o == nil || o.Userdom == nil {
		return nil, false
	}
	return o.Userdom, true
}

// HasUserdom returns a boolean if a field has been set.
func (o *NiatelemetryApicCoreFileDetailsAllOf) HasUserdom() bool {
	if o != nil && o.Userdom != nil {
		return true
	}

	return false
}

// SetUserdom gets a reference to the given string and assigns it to the Userdom field.
func (o *NiatelemetryApicCoreFileDetailsAllOf) SetUserdom(v string) {
	o.Userdom = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NiatelemetryApicCoreFileDetailsAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicCoreFileDetailsAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetryApicCoreFileDetailsAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetryApicCoreFileDetailsAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NiatelemetryApicCoreFileDetailsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Annotation != nil {
		toSerialize["Annotation"] = o.Annotation
	}
	if o.ChildAction != nil {
		toSerialize["ChildAction"] = o.ChildAction
	}
	if o.CollectionTime != nil {
		toSerialize["CollectionTime"] = o.CollectionTime
	}
	if o.DataType != nil {
		toSerialize["DataType"] = o.DataType
	}
	if o.Dn != nil {
		toSerialize["Dn"] = o.Dn
	}
	if o.ExportFileUri != nil {
		toSerialize["ExportFileUri"] = o.ExportFileUri
	}
	if o.ExportStatus != nil {
		toSerialize["ExportStatus"] = o.ExportStatus
	}
	if o.ExportStatusStr != nil {
		toSerialize["ExportStatusStr"] = o.ExportStatusStr
	}
	if o.ExportTechSupFileUri != nil {
		toSerialize["ExportTechSupFileUri"] = o.ExportTechSupFileUri
	}
	if o.ExportedToController != nil {
		toSerialize["ExportedToController"] = o.ExportedToController
	}
	if o.ExtMngdBy != nil {
		toSerialize["ExtMngdBy"] = o.ExtMngdBy
	}
	if o.FileSize != nil {
		toSerialize["FileSize"] = o.FileSize
	}
	if o.HostName != nil {
		toSerialize["HostName"] = o.HostName
	}
	if o.LcOwn != nil {
		toSerialize["LcOwn"] = o.LcOwn
	}
	if o.ModTs != nil {
		toSerialize["ModTs"] = o.ModTs
	}
	if o.NodeId != nil {
		toSerialize["NodeId"] = o.NodeId
	}
	if o.PolName != nil {
		toSerialize["PolName"] = o.PolName
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
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.Uid != nil {
		toSerialize["Uid"] = o.Uid
	}
	if o.Userdom != nil {
		toSerialize["Userdom"] = o.Userdom
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryApicCoreFileDetailsAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varNiatelemetryApicCoreFileDetailsAllOf := _NiatelemetryApicCoreFileDetailsAllOf{}

	if err = json.Unmarshal(bytes, &varNiatelemetryApicCoreFileDetailsAllOf); err == nil {
		*o = NiatelemetryApicCoreFileDetailsAllOf(varNiatelemetryApicCoreFileDetailsAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Annotation")
		delete(additionalProperties, "ChildAction")
		delete(additionalProperties, "CollectionTime")
		delete(additionalProperties, "DataType")
		delete(additionalProperties, "Dn")
		delete(additionalProperties, "ExportFileUri")
		delete(additionalProperties, "ExportStatus")
		delete(additionalProperties, "ExportStatusStr")
		delete(additionalProperties, "ExportTechSupFileUri")
		delete(additionalProperties, "ExportedToController")
		delete(additionalProperties, "ExtMngdBy")
		delete(additionalProperties, "FileSize")
		delete(additionalProperties, "HostName")
		delete(additionalProperties, "LcOwn")
		delete(additionalProperties, "ModTs")
		delete(additionalProperties, "NodeId")
		delete(additionalProperties, "PolName")
		delete(additionalProperties, "RecordType")
		delete(additionalProperties, "RecordVersion")
		delete(additionalProperties, "SiteName")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "Uid")
		delete(additionalProperties, "Userdom")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNiatelemetryApicCoreFileDetailsAllOf struct {
	value *NiatelemetryApicCoreFileDetailsAllOf
	isSet bool
}

func (v NullableNiatelemetryApicCoreFileDetailsAllOf) Get() *NiatelemetryApicCoreFileDetailsAllOf {
	return v.value
}

func (v *NullableNiatelemetryApicCoreFileDetailsAllOf) Set(val *NiatelemetryApicCoreFileDetailsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryApicCoreFileDetailsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryApicCoreFileDetailsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryApicCoreFileDetailsAllOf(val *NiatelemetryApicCoreFileDetailsAllOf) *NullableNiatelemetryApicCoreFileDetailsAllOf {
	return &NullableNiatelemetryApicCoreFileDetailsAllOf{value: val, isSet: true}
}

func (v NullableNiatelemetryApicCoreFileDetailsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryApicCoreFileDetailsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
