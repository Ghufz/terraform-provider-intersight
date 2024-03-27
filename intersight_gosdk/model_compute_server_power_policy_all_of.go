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

// ComputeServerPowerPolicyAllOf Definition of the list of properties defined in 'compute.ServerPowerPolicy', excluding properties defined in parent classes.
type ComputeServerPowerPolicyAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// User configured power state of server. * `Policy` - Power state is set to the default value in the policy. * `PowerOn` - Power state of the server is set to On. * `PowerOff` - Power state is the server set to Off. * `PowerCycle` - Power state the server is reset. * `HardReset` - Power state the server is hard reset. * `Shutdown` - Operating system on the server is shut down. * `Reboot` - Power state of IMC is rebooted.
	PowerState *string `json:"PowerState,omitempty"`
	// The name of the server it is associated with.
	ServerName   *string                               `json:"ServerName,omitempty"`
	Organization *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	// An array of relationships to policyAbstractConfigProfile resources.
	Profiles             []PolicyAbstractConfigProfileRelationship `json:"Profiles,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship      `json:"RegisteredDevice,omitempty"`
	Server               *ComputePhysicalRelationship              `json:"Server,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ComputeServerPowerPolicyAllOf ComputeServerPowerPolicyAllOf

// NewComputeServerPowerPolicyAllOf instantiates a new ComputeServerPowerPolicyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputeServerPowerPolicyAllOf(classId string, objectType string) *ComputeServerPowerPolicyAllOf {
	this := ComputeServerPowerPolicyAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var powerState string = "Policy"
	this.PowerState = &powerState
	return &this
}

// NewComputeServerPowerPolicyAllOfWithDefaults instantiates a new ComputeServerPowerPolicyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputeServerPowerPolicyAllOfWithDefaults() *ComputeServerPowerPolicyAllOf {
	this := ComputeServerPowerPolicyAllOf{}
	var classId string = "compute.ServerPowerPolicy"
	this.ClassId = classId
	var objectType string = "compute.ServerPowerPolicy"
	this.ObjectType = objectType
	var powerState string = "Policy"
	this.PowerState = &powerState
	return &this
}

// GetClassId returns the ClassId field value
func (o *ComputeServerPowerPolicyAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ComputeServerPowerPolicyAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ComputeServerPowerPolicyAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ComputeServerPowerPolicyAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ComputeServerPowerPolicyAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ComputeServerPowerPolicyAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetPowerState returns the PowerState field value if set, zero value otherwise.
func (o *ComputeServerPowerPolicyAllOf) GetPowerState() string {
	if o == nil || o.PowerState == nil {
		var ret string
		return ret
	}
	return *o.PowerState
}

// GetPowerStateOk returns a tuple with the PowerState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeServerPowerPolicyAllOf) GetPowerStateOk() (*string, bool) {
	if o == nil || o.PowerState == nil {
		return nil, false
	}
	return o.PowerState, true
}

// HasPowerState returns a boolean if a field has been set.
func (o *ComputeServerPowerPolicyAllOf) HasPowerState() bool {
	if o != nil && o.PowerState != nil {
		return true
	}

	return false
}

// SetPowerState gets a reference to the given string and assigns it to the PowerState field.
func (o *ComputeServerPowerPolicyAllOf) SetPowerState(v string) {
	o.PowerState = &v
}

// GetServerName returns the ServerName field value if set, zero value otherwise.
func (o *ComputeServerPowerPolicyAllOf) GetServerName() string {
	if o == nil || o.ServerName == nil {
		var ret string
		return ret
	}
	return *o.ServerName
}

// GetServerNameOk returns a tuple with the ServerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeServerPowerPolicyAllOf) GetServerNameOk() (*string, bool) {
	if o == nil || o.ServerName == nil {
		return nil, false
	}
	return o.ServerName, true
}

// HasServerName returns a boolean if a field has been set.
func (o *ComputeServerPowerPolicyAllOf) HasServerName() bool {
	if o != nil && o.ServerName != nil {
		return true
	}

	return false
}

// SetServerName gets a reference to the given string and assigns it to the ServerName field.
func (o *ComputeServerPowerPolicyAllOf) SetServerName(v string) {
	o.ServerName = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *ComputeServerPowerPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeServerPowerPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *ComputeServerPowerPolicyAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *ComputeServerPowerPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetProfiles returns the Profiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComputeServerPowerPolicyAllOf) GetProfiles() []PolicyAbstractConfigProfileRelationship {
	if o == nil {
		var ret []PolicyAbstractConfigProfileRelationship
		return ret
	}
	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComputeServerPowerPolicyAllOf) GetProfilesOk() ([]PolicyAbstractConfigProfileRelationship, bool) {
	if o == nil || o.Profiles == nil {
		return nil, false
	}
	return o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *ComputeServerPowerPolicyAllOf) HasProfiles() bool {
	if o != nil && o.Profiles != nil {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []PolicyAbstractConfigProfileRelationship and assigns it to the Profiles field.
func (o *ComputeServerPowerPolicyAllOf) SetProfiles(v []PolicyAbstractConfigProfileRelationship) {
	o.Profiles = v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *ComputeServerPowerPolicyAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeServerPowerPolicyAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *ComputeServerPowerPolicyAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *ComputeServerPowerPolicyAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

// GetServer returns the Server field value if set, zero value otherwise.
func (o *ComputeServerPowerPolicyAllOf) GetServer() ComputePhysicalRelationship {
	if o == nil || o.Server == nil {
		var ret ComputePhysicalRelationship
		return ret
	}
	return *o.Server
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeServerPowerPolicyAllOf) GetServerOk() (*ComputePhysicalRelationship, bool) {
	if o == nil || o.Server == nil {
		return nil, false
	}
	return o.Server, true
}

// HasServer returns a boolean if a field has been set.
func (o *ComputeServerPowerPolicyAllOf) HasServer() bool {
	if o != nil && o.Server != nil {
		return true
	}

	return false
}

// SetServer gets a reference to the given ComputePhysicalRelationship and assigns it to the Server field.
func (o *ComputeServerPowerPolicyAllOf) SetServer(v ComputePhysicalRelationship) {
	o.Server = &v
}

func (o ComputeServerPowerPolicyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.PowerState != nil {
		toSerialize["PowerState"] = o.PowerState
	}
	if o.ServerName != nil {
		toSerialize["ServerName"] = o.ServerName
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.Profiles != nil {
		toSerialize["Profiles"] = o.Profiles
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	if o.Server != nil {
		toSerialize["Server"] = o.Server
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ComputeServerPowerPolicyAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varComputeServerPowerPolicyAllOf := _ComputeServerPowerPolicyAllOf{}

	if err = json.Unmarshal(bytes, &varComputeServerPowerPolicyAllOf); err == nil {
		*o = ComputeServerPowerPolicyAllOf(varComputeServerPowerPolicyAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "PowerState")
		delete(additionalProperties, "ServerName")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "Profiles")
		delete(additionalProperties, "RegisteredDevice")
		delete(additionalProperties, "Server")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableComputeServerPowerPolicyAllOf struct {
	value *ComputeServerPowerPolicyAllOf
	isSet bool
}

func (v NullableComputeServerPowerPolicyAllOf) Get() *ComputeServerPowerPolicyAllOf {
	return v.value
}

func (v *NullableComputeServerPowerPolicyAllOf) Set(val *ComputeServerPowerPolicyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableComputeServerPowerPolicyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableComputeServerPowerPolicyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputeServerPowerPolicyAllOf(val *ComputeServerPowerPolicyAllOf) *NullableComputeServerPowerPolicyAllOf {
	return &NullableComputeServerPowerPolicyAllOf{value: val, isSet: true}
}

func (v NullableComputeServerPowerPolicyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputeServerPowerPolicyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
