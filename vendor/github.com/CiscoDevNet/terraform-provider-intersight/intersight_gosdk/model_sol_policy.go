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
	"reflect"
	"strings"
)

// SolPolicy Policy for configuring Serial Over LAN settings on endpoint.
type SolPolicy struct {
	PolicyAbstractPolicy
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Baud Rate used for Serial Over LAN communication. * `9600` - Use baud rate 9600 for communication. * `19200` - Use baud rate 19200 for communication. * `38400` - Use baud rate 38400 for communication. * `57600` - Use baud rate 57600 for communication. * `115200` - Use baud rate 115200 for communication.
	BaudRate *int32 `json:"BaudRate,omitempty"`
	// Serial port through which the system routes Serial Over LAN communication. This field is available only on some Cisco UCS C-Series servers. If it is unavailable, the server uses COM port 0 by default. * `com0` - Use serial port com0 for communication. * `com1` - Use serial port com1 for communication.
	ComPort *string `json:"ComPort,omitempty"`
	// State of Serial Over LAN service on the endpoint.
	Enabled *bool `json:"Enabled,omitempty"`
	// SSH port used to access Serial Over LAN directly. Enables bypassing Cisco IMC shell to provide direct access to Serial Over LAN.
	SshPort      *int64                                `json:"SshPort,omitempty"`
	Organization *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	// An array of relationships to policyAbstractConfigProfile resources.
	Profiles             []PolicyAbstractConfigProfileRelationship `json:"Profiles,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SolPolicy SolPolicy

// NewSolPolicy instantiates a new SolPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSolPolicy(classId string, objectType string) *SolPolicy {
	this := SolPolicy{}
	this.ClassId = classId
	this.ObjectType = objectType
	var baudRate int32 = 9600
	this.BaudRate = &baudRate
	var comPort string = "com0"
	this.ComPort = &comPort
	var enabled bool = true
	this.Enabled = &enabled
	var sshPort int64 = 2400
	this.SshPort = &sshPort
	return &this
}

// NewSolPolicyWithDefaults instantiates a new SolPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSolPolicyWithDefaults() *SolPolicy {
	this := SolPolicy{}
	var classId string = "sol.Policy"
	this.ClassId = classId
	var objectType string = "sol.Policy"
	this.ObjectType = objectType
	var baudRate int32 = 9600
	this.BaudRate = &baudRate
	var comPort string = "com0"
	this.ComPort = &comPort
	var enabled bool = true
	this.Enabled = &enabled
	var sshPort int64 = 2400
	this.SshPort = &sshPort
	return &this
}

// GetClassId returns the ClassId field value
func (o *SolPolicy) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SolPolicy) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SolPolicy) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *SolPolicy) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SolPolicy) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SolPolicy) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBaudRate returns the BaudRate field value if set, zero value otherwise.
func (o *SolPolicy) GetBaudRate() int32 {
	if o == nil || o.BaudRate == nil {
		var ret int32
		return ret
	}
	return *o.BaudRate
}

// GetBaudRateOk returns a tuple with the BaudRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SolPolicy) GetBaudRateOk() (*int32, bool) {
	if o == nil || o.BaudRate == nil {
		return nil, false
	}
	return o.BaudRate, true
}

// HasBaudRate returns a boolean if a field has been set.
func (o *SolPolicy) HasBaudRate() bool {
	if o != nil && o.BaudRate != nil {
		return true
	}

	return false
}

// SetBaudRate gets a reference to the given int32 and assigns it to the BaudRate field.
func (o *SolPolicy) SetBaudRate(v int32) {
	o.BaudRate = &v
}

// GetComPort returns the ComPort field value if set, zero value otherwise.
func (o *SolPolicy) GetComPort() string {
	if o == nil || o.ComPort == nil {
		var ret string
		return ret
	}
	return *o.ComPort
}

// GetComPortOk returns a tuple with the ComPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SolPolicy) GetComPortOk() (*string, bool) {
	if o == nil || o.ComPort == nil {
		return nil, false
	}
	return o.ComPort, true
}

// HasComPort returns a boolean if a field has been set.
func (o *SolPolicy) HasComPort() bool {
	if o != nil && o.ComPort != nil {
		return true
	}

	return false
}

// SetComPort gets a reference to the given string and assigns it to the ComPort field.
func (o *SolPolicy) SetComPort(v string) {
	o.ComPort = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SolPolicy) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SolPolicy) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SolPolicy) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SolPolicy) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetSshPort returns the SshPort field value if set, zero value otherwise.
func (o *SolPolicy) GetSshPort() int64 {
	if o == nil || o.SshPort == nil {
		var ret int64
		return ret
	}
	return *o.SshPort
}

// GetSshPortOk returns a tuple with the SshPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SolPolicy) GetSshPortOk() (*int64, bool) {
	if o == nil || o.SshPort == nil {
		return nil, false
	}
	return o.SshPort, true
}

// HasSshPort returns a boolean if a field has been set.
func (o *SolPolicy) HasSshPort() bool {
	if o != nil && o.SshPort != nil {
		return true
	}

	return false
}

// SetSshPort gets a reference to the given int64 and assigns it to the SshPort field.
func (o *SolPolicy) SetSshPort(v int64) {
	o.SshPort = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *SolPolicy) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SolPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *SolPolicy) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *SolPolicy) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetProfiles returns the Profiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SolPolicy) GetProfiles() []PolicyAbstractConfigProfileRelationship {
	if o == nil {
		var ret []PolicyAbstractConfigProfileRelationship
		return ret
	}
	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SolPolicy) GetProfilesOk() ([]PolicyAbstractConfigProfileRelationship, bool) {
	if o == nil || o.Profiles == nil {
		return nil, false
	}
	return o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *SolPolicy) HasProfiles() bool {
	if o != nil && o.Profiles != nil {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []PolicyAbstractConfigProfileRelationship and assigns it to the Profiles field.
func (o *SolPolicy) SetProfiles(v []PolicyAbstractConfigProfileRelationship) {
	o.Profiles = v
}

func (o SolPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractPolicy, errPolicyAbstractPolicy := json.Marshal(o.PolicyAbstractPolicy)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	errPolicyAbstractPolicy = json.Unmarshal([]byte(serializedPolicyAbstractPolicy), &toSerialize)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.BaudRate != nil {
		toSerialize["BaudRate"] = o.BaudRate
	}
	if o.ComPort != nil {
		toSerialize["ComPort"] = o.ComPort
	}
	if o.Enabled != nil {
		toSerialize["Enabled"] = o.Enabled
	}
	if o.SshPort != nil {
		toSerialize["SshPort"] = o.SshPort
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.Profiles != nil {
		toSerialize["Profiles"] = o.Profiles
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SolPolicy) UnmarshalJSON(bytes []byte) (err error) {
	type SolPolicyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Baud Rate used for Serial Over LAN communication. * `9600` - Use baud rate 9600 for communication. * `19200` - Use baud rate 19200 for communication. * `38400` - Use baud rate 38400 for communication. * `57600` - Use baud rate 57600 for communication. * `115200` - Use baud rate 115200 for communication.
		BaudRate *int32 `json:"BaudRate,omitempty"`
		// Serial port through which the system routes Serial Over LAN communication. This field is available only on some Cisco UCS C-Series servers. If it is unavailable, the server uses COM port 0 by default. * `com0` - Use serial port com0 for communication. * `com1` - Use serial port com1 for communication.
		ComPort *string `json:"ComPort,omitempty"`
		// State of Serial Over LAN service on the endpoint.
		Enabled *bool `json:"Enabled,omitempty"`
		// SSH port used to access Serial Over LAN directly. Enables bypassing Cisco IMC shell to provide direct access to Serial Over LAN.
		SshPort      *int64                                `json:"SshPort,omitempty"`
		Organization *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
		// An array of relationships to policyAbstractConfigProfile resources.
		Profiles []PolicyAbstractConfigProfileRelationship `json:"Profiles,omitempty"`
	}

	varSolPolicyWithoutEmbeddedStruct := SolPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varSolPolicyWithoutEmbeddedStruct)
	if err == nil {
		varSolPolicy := _SolPolicy{}
		varSolPolicy.ClassId = varSolPolicyWithoutEmbeddedStruct.ClassId
		varSolPolicy.ObjectType = varSolPolicyWithoutEmbeddedStruct.ObjectType
		varSolPolicy.BaudRate = varSolPolicyWithoutEmbeddedStruct.BaudRate
		varSolPolicy.ComPort = varSolPolicyWithoutEmbeddedStruct.ComPort
		varSolPolicy.Enabled = varSolPolicyWithoutEmbeddedStruct.Enabled
		varSolPolicy.SshPort = varSolPolicyWithoutEmbeddedStruct.SshPort
		varSolPolicy.Organization = varSolPolicyWithoutEmbeddedStruct.Organization
		varSolPolicy.Profiles = varSolPolicyWithoutEmbeddedStruct.Profiles
		*o = SolPolicy(varSolPolicy)
	} else {
		return err
	}

	varSolPolicy := _SolPolicy{}

	err = json.Unmarshal(bytes, &varSolPolicy)
	if err == nil {
		o.PolicyAbstractPolicy = varSolPolicy.PolicyAbstractPolicy
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BaudRate")
		delete(additionalProperties, "ComPort")
		delete(additionalProperties, "Enabled")
		delete(additionalProperties, "SshPort")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "Profiles")

		// remove fields from embedded structs
		reflectPolicyAbstractPolicy := reflect.ValueOf(o.PolicyAbstractPolicy)
		for i := 0; i < reflectPolicyAbstractPolicy.Type().NumField(); i++ {
			t := reflectPolicyAbstractPolicy.Type().Field(i)

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

type NullableSolPolicy struct {
	value *SolPolicy
	isSet bool
}

func (v NullableSolPolicy) Get() *SolPolicy {
	return v.value
}

func (v *NullableSolPolicy) Set(val *SolPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableSolPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableSolPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSolPolicy(val *SolPolicy) *NullableSolPolicy {
	return &NullableSolPolicy{value: val, isSet: true}
}

func (v NullableSolPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSolPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
