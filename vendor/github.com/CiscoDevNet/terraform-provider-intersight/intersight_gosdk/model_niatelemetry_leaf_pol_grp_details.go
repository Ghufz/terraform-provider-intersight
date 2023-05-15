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
)

// NiatelemetryLeafPolGrpDetails Object to capture leaf pol group details.
type NiatelemetryLeafPolGrpDetails struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Dn of the Pol group and leaf profile relational object.
	Dn *string `json:"Dn,omitempty"`
	// Fabric node control dn associated with the pol group.
	FabricNodeControlDn *string `json:"FabricNodeControlDn,omitempty"`
	// Fabric node control policy name associated with the pol group.
	FabricNodeControlPolName *string `json:"FabricNodeControlPolName,omitempty"`
	// Leaf policy group name in APIC.
	LeafPolGroupName *string `json:"LeafPolGroupName,omitempty"`
	// Leaf profile group name in APIC.
	LeafProfileName *string `json:"LeafProfileName,omitempty"`
	// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
	RecordType *string `json:"RecordType,omitempty"`
	// Version of record being pushed. This determines what was the API version for data available from the device.
	RecordVersion *string `json:"RecordVersion,omitempty"`
	// Name of the APIC site from which this data is being collected.
	SiteName *string `json:"SiteName,omitempty"`
	// State of fabric node control pol.
	State                *string                              `json:"State,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryLeafPolGrpDetails NiatelemetryLeafPolGrpDetails

// NewNiatelemetryLeafPolGrpDetails instantiates a new NiatelemetryLeafPolGrpDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryLeafPolGrpDetails(classId string, objectType string) *NiatelemetryLeafPolGrpDetails {
	this := NiatelemetryLeafPolGrpDetails{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryLeafPolGrpDetailsWithDefaults instantiates a new NiatelemetryLeafPolGrpDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryLeafPolGrpDetailsWithDefaults() *NiatelemetryLeafPolGrpDetails {
	this := NiatelemetryLeafPolGrpDetails{}
	var classId string = "niatelemetry.LeafPolGrpDetails"
	this.ClassId = classId
	var objectType string = "niatelemetry.LeafPolGrpDetails"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryLeafPolGrpDetails) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryLeafPolGrpDetails) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryLeafPolGrpDetails) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryLeafPolGrpDetails) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryLeafPolGrpDetails) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryLeafPolGrpDetails) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDn returns the Dn field value if set, zero value otherwise.
func (o *NiatelemetryLeafPolGrpDetails) GetDn() string {
	if o == nil || o.Dn == nil {
		var ret string
		return ret
	}
	return *o.Dn
}

// GetDnOk returns a tuple with the Dn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryLeafPolGrpDetails) GetDnOk() (*string, bool) {
	if o == nil || o.Dn == nil {
		return nil, false
	}
	return o.Dn, true
}

// HasDn returns a boolean if a field has been set.
func (o *NiatelemetryLeafPolGrpDetails) HasDn() bool {
	if o != nil && o.Dn != nil {
		return true
	}

	return false
}

// SetDn gets a reference to the given string and assigns it to the Dn field.
func (o *NiatelemetryLeafPolGrpDetails) SetDn(v string) {
	o.Dn = &v
}

// GetFabricNodeControlDn returns the FabricNodeControlDn field value if set, zero value otherwise.
func (o *NiatelemetryLeafPolGrpDetails) GetFabricNodeControlDn() string {
	if o == nil || o.FabricNodeControlDn == nil {
		var ret string
		return ret
	}
	return *o.FabricNodeControlDn
}

// GetFabricNodeControlDnOk returns a tuple with the FabricNodeControlDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryLeafPolGrpDetails) GetFabricNodeControlDnOk() (*string, bool) {
	if o == nil || o.FabricNodeControlDn == nil {
		return nil, false
	}
	return o.FabricNodeControlDn, true
}

// HasFabricNodeControlDn returns a boolean if a field has been set.
func (o *NiatelemetryLeafPolGrpDetails) HasFabricNodeControlDn() bool {
	if o != nil && o.FabricNodeControlDn != nil {
		return true
	}

	return false
}

// SetFabricNodeControlDn gets a reference to the given string and assigns it to the FabricNodeControlDn field.
func (o *NiatelemetryLeafPolGrpDetails) SetFabricNodeControlDn(v string) {
	o.FabricNodeControlDn = &v
}

// GetFabricNodeControlPolName returns the FabricNodeControlPolName field value if set, zero value otherwise.
func (o *NiatelemetryLeafPolGrpDetails) GetFabricNodeControlPolName() string {
	if o == nil || o.FabricNodeControlPolName == nil {
		var ret string
		return ret
	}
	return *o.FabricNodeControlPolName
}

// GetFabricNodeControlPolNameOk returns a tuple with the FabricNodeControlPolName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryLeafPolGrpDetails) GetFabricNodeControlPolNameOk() (*string, bool) {
	if o == nil || o.FabricNodeControlPolName == nil {
		return nil, false
	}
	return o.FabricNodeControlPolName, true
}

// HasFabricNodeControlPolName returns a boolean if a field has been set.
func (o *NiatelemetryLeafPolGrpDetails) HasFabricNodeControlPolName() bool {
	if o != nil && o.FabricNodeControlPolName != nil {
		return true
	}

	return false
}

// SetFabricNodeControlPolName gets a reference to the given string and assigns it to the FabricNodeControlPolName field.
func (o *NiatelemetryLeafPolGrpDetails) SetFabricNodeControlPolName(v string) {
	o.FabricNodeControlPolName = &v
}

// GetLeafPolGroupName returns the LeafPolGroupName field value if set, zero value otherwise.
func (o *NiatelemetryLeafPolGrpDetails) GetLeafPolGroupName() string {
	if o == nil || o.LeafPolGroupName == nil {
		var ret string
		return ret
	}
	return *o.LeafPolGroupName
}

// GetLeafPolGroupNameOk returns a tuple with the LeafPolGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryLeafPolGrpDetails) GetLeafPolGroupNameOk() (*string, bool) {
	if o == nil || o.LeafPolGroupName == nil {
		return nil, false
	}
	return o.LeafPolGroupName, true
}

// HasLeafPolGroupName returns a boolean if a field has been set.
func (o *NiatelemetryLeafPolGrpDetails) HasLeafPolGroupName() bool {
	if o != nil && o.LeafPolGroupName != nil {
		return true
	}

	return false
}

// SetLeafPolGroupName gets a reference to the given string and assigns it to the LeafPolGroupName field.
func (o *NiatelemetryLeafPolGrpDetails) SetLeafPolGroupName(v string) {
	o.LeafPolGroupName = &v
}

// GetLeafProfileName returns the LeafProfileName field value if set, zero value otherwise.
func (o *NiatelemetryLeafPolGrpDetails) GetLeafProfileName() string {
	if o == nil || o.LeafProfileName == nil {
		var ret string
		return ret
	}
	return *o.LeafProfileName
}

// GetLeafProfileNameOk returns a tuple with the LeafProfileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryLeafPolGrpDetails) GetLeafProfileNameOk() (*string, bool) {
	if o == nil || o.LeafProfileName == nil {
		return nil, false
	}
	return o.LeafProfileName, true
}

// HasLeafProfileName returns a boolean if a field has been set.
func (o *NiatelemetryLeafPolGrpDetails) HasLeafProfileName() bool {
	if o != nil && o.LeafProfileName != nil {
		return true
	}

	return false
}

// SetLeafProfileName gets a reference to the given string and assigns it to the LeafProfileName field.
func (o *NiatelemetryLeafPolGrpDetails) SetLeafProfileName(v string) {
	o.LeafProfileName = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *NiatelemetryLeafPolGrpDetails) GetRecordType() string {
	if o == nil || o.RecordType == nil {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryLeafPolGrpDetails) GetRecordTypeOk() (*string, bool) {
	if o == nil || o.RecordType == nil {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *NiatelemetryLeafPolGrpDetails) HasRecordType() bool {
	if o != nil && o.RecordType != nil {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *NiatelemetryLeafPolGrpDetails) SetRecordType(v string) {
	o.RecordType = &v
}

// GetRecordVersion returns the RecordVersion field value if set, zero value otherwise.
func (o *NiatelemetryLeafPolGrpDetails) GetRecordVersion() string {
	if o == nil || o.RecordVersion == nil {
		var ret string
		return ret
	}
	return *o.RecordVersion
}

// GetRecordVersionOk returns a tuple with the RecordVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryLeafPolGrpDetails) GetRecordVersionOk() (*string, bool) {
	if o == nil || o.RecordVersion == nil {
		return nil, false
	}
	return o.RecordVersion, true
}

// HasRecordVersion returns a boolean if a field has been set.
func (o *NiatelemetryLeafPolGrpDetails) HasRecordVersion() bool {
	if o != nil && o.RecordVersion != nil {
		return true
	}

	return false
}

// SetRecordVersion gets a reference to the given string and assigns it to the RecordVersion field.
func (o *NiatelemetryLeafPolGrpDetails) SetRecordVersion(v string) {
	o.RecordVersion = &v
}

// GetSiteName returns the SiteName field value if set, zero value otherwise.
func (o *NiatelemetryLeafPolGrpDetails) GetSiteName() string {
	if o == nil || o.SiteName == nil {
		var ret string
		return ret
	}
	return *o.SiteName
}

// GetSiteNameOk returns a tuple with the SiteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryLeafPolGrpDetails) GetSiteNameOk() (*string, bool) {
	if o == nil || o.SiteName == nil {
		return nil, false
	}
	return o.SiteName, true
}

// HasSiteName returns a boolean if a field has been set.
func (o *NiatelemetryLeafPolGrpDetails) HasSiteName() bool {
	if o != nil && o.SiteName != nil {
		return true
	}

	return false
}

// SetSiteName gets a reference to the given string and assigns it to the SiteName field.
func (o *NiatelemetryLeafPolGrpDetails) SetSiteName(v string) {
	o.SiteName = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *NiatelemetryLeafPolGrpDetails) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryLeafPolGrpDetails) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *NiatelemetryLeafPolGrpDetails) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *NiatelemetryLeafPolGrpDetails) SetState(v string) {
	o.State = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NiatelemetryLeafPolGrpDetails) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryLeafPolGrpDetails) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetryLeafPolGrpDetails) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetryLeafPolGrpDetails) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NiatelemetryLeafPolGrpDetails) MarshalJSON() ([]byte, error) {
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
	if o.Dn != nil {
		toSerialize["Dn"] = o.Dn
	}
	if o.FabricNodeControlDn != nil {
		toSerialize["FabricNodeControlDn"] = o.FabricNodeControlDn
	}
	if o.FabricNodeControlPolName != nil {
		toSerialize["FabricNodeControlPolName"] = o.FabricNodeControlPolName
	}
	if o.LeafPolGroupName != nil {
		toSerialize["LeafPolGroupName"] = o.LeafPolGroupName
	}
	if o.LeafProfileName != nil {
		toSerialize["LeafProfileName"] = o.LeafProfileName
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
	if o.State != nil {
		toSerialize["State"] = o.State
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryLeafPolGrpDetails) UnmarshalJSON(bytes []byte) (err error) {
	type NiatelemetryLeafPolGrpDetailsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Dn of the Pol group and leaf profile relational object.
		Dn *string `json:"Dn,omitempty"`
		// Fabric node control dn associated with the pol group.
		FabricNodeControlDn *string `json:"FabricNodeControlDn,omitempty"`
		// Fabric node control policy name associated with the pol group.
		FabricNodeControlPolName *string `json:"FabricNodeControlPolName,omitempty"`
		// Leaf policy group name in APIC.
		LeafPolGroupName *string `json:"LeafPolGroupName,omitempty"`
		// Leaf profile group name in APIC.
		LeafProfileName *string `json:"LeafProfileName,omitempty"`
		// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
		RecordType *string `json:"RecordType,omitempty"`
		// Version of record being pushed. This determines what was the API version for data available from the device.
		RecordVersion *string `json:"RecordVersion,omitempty"`
		// Name of the APIC site from which this data is being collected.
		SiteName *string `json:"SiteName,omitempty"`
		// State of fabric node control pol.
		State            *string                              `json:"State,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varNiatelemetryLeafPolGrpDetailsWithoutEmbeddedStruct := NiatelemetryLeafPolGrpDetailsWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNiatelemetryLeafPolGrpDetailsWithoutEmbeddedStruct)
	if err == nil {
		varNiatelemetryLeafPolGrpDetails := _NiatelemetryLeafPolGrpDetails{}
		varNiatelemetryLeafPolGrpDetails.ClassId = varNiatelemetryLeafPolGrpDetailsWithoutEmbeddedStruct.ClassId
		varNiatelemetryLeafPolGrpDetails.ObjectType = varNiatelemetryLeafPolGrpDetailsWithoutEmbeddedStruct.ObjectType
		varNiatelemetryLeafPolGrpDetails.Dn = varNiatelemetryLeafPolGrpDetailsWithoutEmbeddedStruct.Dn
		varNiatelemetryLeafPolGrpDetails.FabricNodeControlDn = varNiatelemetryLeafPolGrpDetailsWithoutEmbeddedStruct.FabricNodeControlDn
		varNiatelemetryLeafPolGrpDetails.FabricNodeControlPolName = varNiatelemetryLeafPolGrpDetailsWithoutEmbeddedStruct.FabricNodeControlPolName
		varNiatelemetryLeafPolGrpDetails.LeafPolGroupName = varNiatelemetryLeafPolGrpDetailsWithoutEmbeddedStruct.LeafPolGroupName
		varNiatelemetryLeafPolGrpDetails.LeafProfileName = varNiatelemetryLeafPolGrpDetailsWithoutEmbeddedStruct.LeafProfileName
		varNiatelemetryLeafPolGrpDetails.RecordType = varNiatelemetryLeafPolGrpDetailsWithoutEmbeddedStruct.RecordType
		varNiatelemetryLeafPolGrpDetails.RecordVersion = varNiatelemetryLeafPolGrpDetailsWithoutEmbeddedStruct.RecordVersion
		varNiatelemetryLeafPolGrpDetails.SiteName = varNiatelemetryLeafPolGrpDetailsWithoutEmbeddedStruct.SiteName
		varNiatelemetryLeafPolGrpDetails.State = varNiatelemetryLeafPolGrpDetailsWithoutEmbeddedStruct.State
		varNiatelemetryLeafPolGrpDetails.RegisteredDevice = varNiatelemetryLeafPolGrpDetailsWithoutEmbeddedStruct.RegisteredDevice
		*o = NiatelemetryLeafPolGrpDetails(varNiatelemetryLeafPolGrpDetails)
	} else {
		return err
	}

	varNiatelemetryLeafPolGrpDetails := _NiatelemetryLeafPolGrpDetails{}

	err = json.Unmarshal(bytes, &varNiatelemetryLeafPolGrpDetails)
	if err == nil {
		o.MoBaseMo = varNiatelemetryLeafPolGrpDetails.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Dn")
		delete(additionalProperties, "FabricNodeControlDn")
		delete(additionalProperties, "FabricNodeControlPolName")
		delete(additionalProperties, "LeafPolGroupName")
		delete(additionalProperties, "LeafProfileName")
		delete(additionalProperties, "RecordType")
		delete(additionalProperties, "RecordVersion")
		delete(additionalProperties, "SiteName")
		delete(additionalProperties, "State")
		delete(additionalProperties, "RegisteredDevice")

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

type NullableNiatelemetryLeafPolGrpDetails struct {
	value *NiatelemetryLeafPolGrpDetails
	isSet bool
}

func (v NullableNiatelemetryLeafPolGrpDetails) Get() *NiatelemetryLeafPolGrpDetails {
	return v.value
}

func (v *NullableNiatelemetryLeafPolGrpDetails) Set(val *NiatelemetryLeafPolGrpDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryLeafPolGrpDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryLeafPolGrpDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryLeafPolGrpDetails(val *NiatelemetryLeafPolGrpDetails) *NullableNiatelemetryLeafPolGrpDetails {
	return &NullableNiatelemetryLeafPolGrpDetails{value: val, isSet: true}
}

func (v NullableNiatelemetryLeafPolGrpDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryLeafPolGrpDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
