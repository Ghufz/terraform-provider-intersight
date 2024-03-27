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

// WorkflowAnsiblePlaySessionAllOf Definition of the list of properties defined in 'workflow.AnsiblePlaySession', excluding properties defined in parent classes.
type WorkflowAnsiblePlaySessionAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The command line arguments for running the Ansible playbook against the given endpoint. Escape character backslash needs to be used when the command line arguments contain double quotes in them.
	CommandLineArguments *string `json:"CommandLineArguments,omitempty"`
	// The path of the host inventory file that resides on the Ansible Endpoint target or the comma separated list of hosts on which the Ansible playbook is to be run. Make sure to suffix a comma when the list of hosts is provided as input, even if the list has only one value.
	HostInventory *string `json:"HostInventory,omitempty"`
	// The path of the Ansible playbook that resides on the Ansible Endpoint target.
	PlaybookPath *string `json:"PlaybookPath,omitempty"`
	// SSH operation timeout value in seconds. Value provided should be string representation of an interger.
	SshOpTimeout         *string `json:"SshOpTimeout,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowAnsiblePlaySessionAllOf WorkflowAnsiblePlaySessionAllOf

// NewWorkflowAnsiblePlaySessionAllOf instantiates a new WorkflowAnsiblePlaySessionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowAnsiblePlaySessionAllOf(classId string, objectType string) *WorkflowAnsiblePlaySessionAllOf {
	this := WorkflowAnsiblePlaySessionAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowAnsiblePlaySessionAllOfWithDefaults instantiates a new WorkflowAnsiblePlaySessionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowAnsiblePlaySessionAllOfWithDefaults() *WorkflowAnsiblePlaySessionAllOf {
	this := WorkflowAnsiblePlaySessionAllOf{}
	var classId string = "workflow.AnsiblePlaySession"
	this.ClassId = classId
	var objectType string = "workflow.AnsiblePlaySession"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowAnsiblePlaySessionAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowAnsiblePlaySessionAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowAnsiblePlaySessionAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowAnsiblePlaySessionAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowAnsiblePlaySessionAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowAnsiblePlaySessionAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCommandLineArguments returns the CommandLineArguments field value if set, zero value otherwise.
func (o *WorkflowAnsiblePlaySessionAllOf) GetCommandLineArguments() string {
	if o == nil || o.CommandLineArguments == nil {
		var ret string
		return ret
	}
	return *o.CommandLineArguments
}

// GetCommandLineArgumentsOk returns a tuple with the CommandLineArguments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowAnsiblePlaySessionAllOf) GetCommandLineArgumentsOk() (*string, bool) {
	if o == nil || o.CommandLineArguments == nil {
		return nil, false
	}
	return o.CommandLineArguments, true
}

// HasCommandLineArguments returns a boolean if a field has been set.
func (o *WorkflowAnsiblePlaySessionAllOf) HasCommandLineArguments() bool {
	if o != nil && o.CommandLineArguments != nil {
		return true
	}

	return false
}

// SetCommandLineArguments gets a reference to the given string and assigns it to the CommandLineArguments field.
func (o *WorkflowAnsiblePlaySessionAllOf) SetCommandLineArguments(v string) {
	o.CommandLineArguments = &v
}

// GetHostInventory returns the HostInventory field value if set, zero value otherwise.
func (o *WorkflowAnsiblePlaySessionAllOf) GetHostInventory() string {
	if o == nil || o.HostInventory == nil {
		var ret string
		return ret
	}
	return *o.HostInventory
}

// GetHostInventoryOk returns a tuple with the HostInventory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowAnsiblePlaySessionAllOf) GetHostInventoryOk() (*string, bool) {
	if o == nil || o.HostInventory == nil {
		return nil, false
	}
	return o.HostInventory, true
}

// HasHostInventory returns a boolean if a field has been set.
func (o *WorkflowAnsiblePlaySessionAllOf) HasHostInventory() bool {
	if o != nil && o.HostInventory != nil {
		return true
	}

	return false
}

// SetHostInventory gets a reference to the given string and assigns it to the HostInventory field.
func (o *WorkflowAnsiblePlaySessionAllOf) SetHostInventory(v string) {
	o.HostInventory = &v
}

// GetPlaybookPath returns the PlaybookPath field value if set, zero value otherwise.
func (o *WorkflowAnsiblePlaySessionAllOf) GetPlaybookPath() string {
	if o == nil || o.PlaybookPath == nil {
		var ret string
		return ret
	}
	return *o.PlaybookPath
}

// GetPlaybookPathOk returns a tuple with the PlaybookPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowAnsiblePlaySessionAllOf) GetPlaybookPathOk() (*string, bool) {
	if o == nil || o.PlaybookPath == nil {
		return nil, false
	}
	return o.PlaybookPath, true
}

// HasPlaybookPath returns a boolean if a field has been set.
func (o *WorkflowAnsiblePlaySessionAllOf) HasPlaybookPath() bool {
	if o != nil && o.PlaybookPath != nil {
		return true
	}

	return false
}

// SetPlaybookPath gets a reference to the given string and assigns it to the PlaybookPath field.
func (o *WorkflowAnsiblePlaySessionAllOf) SetPlaybookPath(v string) {
	o.PlaybookPath = &v
}

// GetSshOpTimeout returns the SshOpTimeout field value if set, zero value otherwise.
func (o *WorkflowAnsiblePlaySessionAllOf) GetSshOpTimeout() string {
	if o == nil || o.SshOpTimeout == nil {
		var ret string
		return ret
	}
	return *o.SshOpTimeout
}

// GetSshOpTimeoutOk returns a tuple with the SshOpTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowAnsiblePlaySessionAllOf) GetSshOpTimeoutOk() (*string, bool) {
	if o == nil || o.SshOpTimeout == nil {
		return nil, false
	}
	return o.SshOpTimeout, true
}

// HasSshOpTimeout returns a boolean if a field has been set.
func (o *WorkflowAnsiblePlaySessionAllOf) HasSshOpTimeout() bool {
	if o != nil && o.SshOpTimeout != nil {
		return true
	}

	return false
}

// SetSshOpTimeout gets a reference to the given string and assigns it to the SshOpTimeout field.
func (o *WorkflowAnsiblePlaySessionAllOf) SetSshOpTimeout(v string) {
	o.SshOpTimeout = &v
}

func (o WorkflowAnsiblePlaySessionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.CommandLineArguments != nil {
		toSerialize["CommandLineArguments"] = o.CommandLineArguments
	}
	if o.HostInventory != nil {
		toSerialize["HostInventory"] = o.HostInventory
	}
	if o.PlaybookPath != nil {
		toSerialize["PlaybookPath"] = o.PlaybookPath
	}
	if o.SshOpTimeout != nil {
		toSerialize["SshOpTimeout"] = o.SshOpTimeout
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowAnsiblePlaySessionAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowAnsiblePlaySessionAllOf := _WorkflowAnsiblePlaySessionAllOf{}

	if err = json.Unmarshal(bytes, &varWorkflowAnsiblePlaySessionAllOf); err == nil {
		*o = WorkflowAnsiblePlaySessionAllOf(varWorkflowAnsiblePlaySessionAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CommandLineArguments")
		delete(additionalProperties, "HostInventory")
		delete(additionalProperties, "PlaybookPath")
		delete(additionalProperties, "SshOpTimeout")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowAnsiblePlaySessionAllOf struct {
	value *WorkflowAnsiblePlaySessionAllOf
	isSet bool
}

func (v NullableWorkflowAnsiblePlaySessionAllOf) Get() *WorkflowAnsiblePlaySessionAllOf {
	return v.value
}

func (v *NullableWorkflowAnsiblePlaySessionAllOf) Set(val *WorkflowAnsiblePlaySessionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowAnsiblePlaySessionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowAnsiblePlaySessionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowAnsiblePlaySessionAllOf(val *WorkflowAnsiblePlaySessionAllOf) *NullableWorkflowAnsiblePlaySessionAllOf {
	return &NullableWorkflowAnsiblePlaySessionAllOf{value: val, isSet: true}
}

func (v NullableWorkflowAnsiblePlaySessionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowAnsiblePlaySessionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
