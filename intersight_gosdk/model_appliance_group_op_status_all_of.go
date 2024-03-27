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

// ApplianceGroupOpStatusAllOf Definition of the list of properties defined in 'appliance.GroupOpStatus', excluding properties defined in parent classes.
type ApplianceGroupOpStatusAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Description of the group.
	Description *string `json:"Description,omitempty"`
	// The name of group, which includes Identity Management, Device Connector Service, Core Service, Analytics, Infra and Appliance.
	GroupName *string `json:"GroupName,omitempty"`
	// The overall API status from this group's applications.
	OverallStatus *string `json:"OverallStatus,omitempty"`
	// An array of relationships to applianceAppOpStatus resources.
	Apps                 []ApplianceAppOpStatusRelationship   `json:"Apps,omitempty"`
	SystemOpStatus       *ApplianceSystemOpStatusRelationship `json:"SystemOpStatus,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplianceGroupOpStatusAllOf ApplianceGroupOpStatusAllOf

// NewApplianceGroupOpStatusAllOf instantiates a new ApplianceGroupOpStatusAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceGroupOpStatusAllOf(classId string, objectType string) *ApplianceGroupOpStatusAllOf {
	this := ApplianceGroupOpStatusAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewApplianceGroupOpStatusAllOfWithDefaults instantiates a new ApplianceGroupOpStatusAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceGroupOpStatusAllOfWithDefaults() *ApplianceGroupOpStatusAllOf {
	this := ApplianceGroupOpStatusAllOf{}
	var classId string = "appliance.GroupOpStatus"
	this.ClassId = classId
	var objectType string = "appliance.GroupOpStatus"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ApplianceGroupOpStatusAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ApplianceGroupOpStatusAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ApplianceGroupOpStatusAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ApplianceGroupOpStatusAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ApplianceGroupOpStatusAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ApplianceGroupOpStatusAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ApplianceGroupOpStatusAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceGroupOpStatusAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ApplianceGroupOpStatusAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ApplianceGroupOpStatusAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetGroupName returns the GroupName field value if set, zero value otherwise.
func (o *ApplianceGroupOpStatusAllOf) GetGroupName() string {
	if o == nil || o.GroupName == nil {
		var ret string
		return ret
	}
	return *o.GroupName
}

// GetGroupNameOk returns a tuple with the GroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceGroupOpStatusAllOf) GetGroupNameOk() (*string, bool) {
	if o == nil || o.GroupName == nil {
		return nil, false
	}
	return o.GroupName, true
}

// HasGroupName returns a boolean if a field has been set.
func (o *ApplianceGroupOpStatusAllOf) HasGroupName() bool {
	if o != nil && o.GroupName != nil {
		return true
	}

	return false
}

// SetGroupName gets a reference to the given string and assigns it to the GroupName field.
func (o *ApplianceGroupOpStatusAllOf) SetGroupName(v string) {
	o.GroupName = &v
}

// GetOverallStatus returns the OverallStatus field value if set, zero value otherwise.
func (o *ApplianceGroupOpStatusAllOf) GetOverallStatus() string {
	if o == nil || o.OverallStatus == nil {
		var ret string
		return ret
	}
	return *o.OverallStatus
}

// GetOverallStatusOk returns a tuple with the OverallStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceGroupOpStatusAllOf) GetOverallStatusOk() (*string, bool) {
	if o == nil || o.OverallStatus == nil {
		return nil, false
	}
	return o.OverallStatus, true
}

// HasOverallStatus returns a boolean if a field has been set.
func (o *ApplianceGroupOpStatusAllOf) HasOverallStatus() bool {
	if o != nil && o.OverallStatus != nil {
		return true
	}

	return false
}

// SetOverallStatus gets a reference to the given string and assigns it to the OverallStatus field.
func (o *ApplianceGroupOpStatusAllOf) SetOverallStatus(v string) {
	o.OverallStatus = &v
}

// GetApps returns the Apps field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplianceGroupOpStatusAllOf) GetApps() []ApplianceAppOpStatusRelationship {
	if o == nil {
		var ret []ApplianceAppOpStatusRelationship
		return ret
	}
	return o.Apps
}

// GetAppsOk returns a tuple with the Apps field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplianceGroupOpStatusAllOf) GetAppsOk() ([]ApplianceAppOpStatusRelationship, bool) {
	if o == nil || o.Apps == nil {
		return nil, false
	}
	return o.Apps, true
}

// HasApps returns a boolean if a field has been set.
func (o *ApplianceGroupOpStatusAllOf) HasApps() bool {
	if o != nil && o.Apps != nil {
		return true
	}

	return false
}

// SetApps gets a reference to the given []ApplianceAppOpStatusRelationship and assigns it to the Apps field.
func (o *ApplianceGroupOpStatusAllOf) SetApps(v []ApplianceAppOpStatusRelationship) {
	o.Apps = v
}

// GetSystemOpStatus returns the SystemOpStatus field value if set, zero value otherwise.
func (o *ApplianceGroupOpStatusAllOf) GetSystemOpStatus() ApplianceSystemOpStatusRelationship {
	if o == nil || o.SystemOpStatus == nil {
		var ret ApplianceSystemOpStatusRelationship
		return ret
	}
	return *o.SystemOpStatus
}

// GetSystemOpStatusOk returns a tuple with the SystemOpStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceGroupOpStatusAllOf) GetSystemOpStatusOk() (*ApplianceSystemOpStatusRelationship, bool) {
	if o == nil || o.SystemOpStatus == nil {
		return nil, false
	}
	return o.SystemOpStatus, true
}

// HasSystemOpStatus returns a boolean if a field has been set.
func (o *ApplianceGroupOpStatusAllOf) HasSystemOpStatus() bool {
	if o != nil && o.SystemOpStatus != nil {
		return true
	}

	return false
}

// SetSystemOpStatus gets a reference to the given ApplianceSystemOpStatusRelationship and assigns it to the SystemOpStatus field.
func (o *ApplianceGroupOpStatusAllOf) SetSystemOpStatus(v ApplianceSystemOpStatusRelationship) {
	o.SystemOpStatus = &v
}

func (o ApplianceGroupOpStatusAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.GroupName != nil {
		toSerialize["GroupName"] = o.GroupName
	}
	if o.OverallStatus != nil {
		toSerialize["OverallStatus"] = o.OverallStatus
	}
	if o.Apps != nil {
		toSerialize["Apps"] = o.Apps
	}
	if o.SystemOpStatus != nil {
		toSerialize["SystemOpStatus"] = o.SystemOpStatus
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ApplianceGroupOpStatusAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varApplianceGroupOpStatusAllOf := _ApplianceGroupOpStatusAllOf{}

	if err = json.Unmarshal(bytes, &varApplianceGroupOpStatusAllOf); err == nil {
		*o = ApplianceGroupOpStatusAllOf(varApplianceGroupOpStatusAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "GroupName")
		delete(additionalProperties, "OverallStatus")
		delete(additionalProperties, "Apps")
		delete(additionalProperties, "SystemOpStatus")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApplianceGroupOpStatusAllOf struct {
	value *ApplianceGroupOpStatusAllOf
	isSet bool
}

func (v NullableApplianceGroupOpStatusAllOf) Get() *ApplianceGroupOpStatusAllOf {
	return v.value
}

func (v *NullableApplianceGroupOpStatusAllOf) Set(val *ApplianceGroupOpStatusAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceGroupOpStatusAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceGroupOpStatusAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceGroupOpStatusAllOf(val *ApplianceGroupOpStatusAllOf) *NullableApplianceGroupOpStatusAllOf {
	return &NullableApplianceGroupOpStatusAllOf{value: val, isSet: true}
}

func (v NullableApplianceGroupOpStatusAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceGroupOpStatusAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
