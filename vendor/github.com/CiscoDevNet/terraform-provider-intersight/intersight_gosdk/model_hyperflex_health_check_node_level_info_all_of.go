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

// HyperflexHealthCheckNodeLevelInfoAllOf Definition of the list of properties defined in 'hyperflex.HealthCheckNodeLevelInfo', excluding properties defined in parent classes.
type HyperflexHealthCheckNodeLevelInfoAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Node-specific check failure cause.
	NodeCause *string `json:"NodeCause,omitempty"`
	// Node-specific check result.
	NodeCheckResult *string `json:"NodeCheckResult,omitempty"`
	// The IP Address of the ESXi server.
	NodeEsxIpAddress *string `json:"NodeEsxIpAddress,omitempty"`
	// The IP Address of cluster node.
	NodeIpAddress *string `json:"NodeIpAddress,omitempty"`
	// Cluster node name on which the check was run.
	NodeName *string `json:"NodeName,omitempty"`
	// Node-specific check failure suggested resolution.
	NodeResolution       *string `json:"NodeResolution,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexHealthCheckNodeLevelInfoAllOf HyperflexHealthCheckNodeLevelInfoAllOf

// NewHyperflexHealthCheckNodeLevelInfoAllOf instantiates a new HyperflexHealthCheckNodeLevelInfoAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexHealthCheckNodeLevelInfoAllOf(classId string, objectType string) *HyperflexHealthCheckNodeLevelInfoAllOf {
	this := HyperflexHealthCheckNodeLevelInfoAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexHealthCheckNodeLevelInfoAllOfWithDefaults instantiates a new HyperflexHealthCheckNodeLevelInfoAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexHealthCheckNodeLevelInfoAllOfWithDefaults() *HyperflexHealthCheckNodeLevelInfoAllOf {
	this := HyperflexHealthCheckNodeLevelInfoAllOf{}
	var classId string = "hyperflex.HealthCheckNodeLevelInfo"
	this.ClassId = classId
	var objectType string = "hyperflex.HealthCheckNodeLevelInfo"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexHealthCheckNodeLevelInfoAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexHealthCheckNodeLevelInfoAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexHealthCheckNodeLevelInfoAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexHealthCheckNodeLevelInfoAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexHealthCheckNodeLevelInfoAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexHealthCheckNodeLevelInfoAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetNodeCause returns the NodeCause field value if set, zero value otherwise.
func (o *HyperflexHealthCheckNodeLevelInfoAllOf) GetNodeCause() string {
	if o == nil || o.NodeCause == nil {
		var ret string
		return ret
	}
	return *o.NodeCause
}

// GetNodeCauseOk returns a tuple with the NodeCause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHealthCheckNodeLevelInfoAllOf) GetNodeCauseOk() (*string, bool) {
	if o == nil || o.NodeCause == nil {
		return nil, false
	}
	return o.NodeCause, true
}

// HasNodeCause returns a boolean if a field has been set.
func (o *HyperflexHealthCheckNodeLevelInfoAllOf) HasNodeCause() bool {
	if o != nil && o.NodeCause != nil {
		return true
	}

	return false
}

// SetNodeCause gets a reference to the given string and assigns it to the NodeCause field.
func (o *HyperflexHealthCheckNodeLevelInfoAllOf) SetNodeCause(v string) {
	o.NodeCause = &v
}

// GetNodeCheckResult returns the NodeCheckResult field value if set, zero value otherwise.
func (o *HyperflexHealthCheckNodeLevelInfoAllOf) GetNodeCheckResult() string {
	if o == nil || o.NodeCheckResult == nil {
		var ret string
		return ret
	}
	return *o.NodeCheckResult
}

// GetNodeCheckResultOk returns a tuple with the NodeCheckResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHealthCheckNodeLevelInfoAllOf) GetNodeCheckResultOk() (*string, bool) {
	if o == nil || o.NodeCheckResult == nil {
		return nil, false
	}
	return o.NodeCheckResult, true
}

// HasNodeCheckResult returns a boolean if a field has been set.
func (o *HyperflexHealthCheckNodeLevelInfoAllOf) HasNodeCheckResult() bool {
	if o != nil && o.NodeCheckResult != nil {
		return true
	}

	return false
}

// SetNodeCheckResult gets a reference to the given string and assigns it to the NodeCheckResult field.
func (o *HyperflexHealthCheckNodeLevelInfoAllOf) SetNodeCheckResult(v string) {
	o.NodeCheckResult = &v
}

// GetNodeEsxIpAddress returns the NodeEsxIpAddress field value if set, zero value otherwise.
func (o *HyperflexHealthCheckNodeLevelInfoAllOf) GetNodeEsxIpAddress() string {
	if o == nil || o.NodeEsxIpAddress == nil {
		var ret string
		return ret
	}
	return *o.NodeEsxIpAddress
}

// GetNodeEsxIpAddressOk returns a tuple with the NodeEsxIpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHealthCheckNodeLevelInfoAllOf) GetNodeEsxIpAddressOk() (*string, bool) {
	if o == nil || o.NodeEsxIpAddress == nil {
		return nil, false
	}
	return o.NodeEsxIpAddress, true
}

// HasNodeEsxIpAddress returns a boolean if a field has been set.
func (o *HyperflexHealthCheckNodeLevelInfoAllOf) HasNodeEsxIpAddress() bool {
	if o != nil && o.NodeEsxIpAddress != nil {
		return true
	}

	return false
}

// SetNodeEsxIpAddress gets a reference to the given string and assigns it to the NodeEsxIpAddress field.
func (o *HyperflexHealthCheckNodeLevelInfoAllOf) SetNodeEsxIpAddress(v string) {
	o.NodeEsxIpAddress = &v
}

// GetNodeIpAddress returns the NodeIpAddress field value if set, zero value otherwise.
func (o *HyperflexHealthCheckNodeLevelInfoAllOf) GetNodeIpAddress() string {
	if o == nil || o.NodeIpAddress == nil {
		var ret string
		return ret
	}
	return *o.NodeIpAddress
}

// GetNodeIpAddressOk returns a tuple with the NodeIpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHealthCheckNodeLevelInfoAllOf) GetNodeIpAddressOk() (*string, bool) {
	if o == nil || o.NodeIpAddress == nil {
		return nil, false
	}
	return o.NodeIpAddress, true
}

// HasNodeIpAddress returns a boolean if a field has been set.
func (o *HyperflexHealthCheckNodeLevelInfoAllOf) HasNodeIpAddress() bool {
	if o != nil && o.NodeIpAddress != nil {
		return true
	}

	return false
}

// SetNodeIpAddress gets a reference to the given string and assigns it to the NodeIpAddress field.
func (o *HyperflexHealthCheckNodeLevelInfoAllOf) SetNodeIpAddress(v string) {
	o.NodeIpAddress = &v
}

// GetNodeName returns the NodeName field value if set, zero value otherwise.
func (o *HyperflexHealthCheckNodeLevelInfoAllOf) GetNodeName() string {
	if o == nil || o.NodeName == nil {
		var ret string
		return ret
	}
	return *o.NodeName
}

// GetNodeNameOk returns a tuple with the NodeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHealthCheckNodeLevelInfoAllOf) GetNodeNameOk() (*string, bool) {
	if o == nil || o.NodeName == nil {
		return nil, false
	}
	return o.NodeName, true
}

// HasNodeName returns a boolean if a field has been set.
func (o *HyperflexHealthCheckNodeLevelInfoAllOf) HasNodeName() bool {
	if o != nil && o.NodeName != nil {
		return true
	}

	return false
}

// SetNodeName gets a reference to the given string and assigns it to the NodeName field.
func (o *HyperflexHealthCheckNodeLevelInfoAllOf) SetNodeName(v string) {
	o.NodeName = &v
}

// GetNodeResolution returns the NodeResolution field value if set, zero value otherwise.
func (o *HyperflexHealthCheckNodeLevelInfoAllOf) GetNodeResolution() string {
	if o == nil || o.NodeResolution == nil {
		var ret string
		return ret
	}
	return *o.NodeResolution
}

// GetNodeResolutionOk returns a tuple with the NodeResolution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHealthCheckNodeLevelInfoAllOf) GetNodeResolutionOk() (*string, bool) {
	if o == nil || o.NodeResolution == nil {
		return nil, false
	}
	return o.NodeResolution, true
}

// HasNodeResolution returns a boolean if a field has been set.
func (o *HyperflexHealthCheckNodeLevelInfoAllOf) HasNodeResolution() bool {
	if o != nil && o.NodeResolution != nil {
		return true
	}

	return false
}

// SetNodeResolution gets a reference to the given string and assigns it to the NodeResolution field.
func (o *HyperflexHealthCheckNodeLevelInfoAllOf) SetNodeResolution(v string) {
	o.NodeResolution = &v
}

func (o HyperflexHealthCheckNodeLevelInfoAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.NodeCause != nil {
		toSerialize["NodeCause"] = o.NodeCause
	}
	if o.NodeCheckResult != nil {
		toSerialize["NodeCheckResult"] = o.NodeCheckResult
	}
	if o.NodeEsxIpAddress != nil {
		toSerialize["NodeEsxIpAddress"] = o.NodeEsxIpAddress
	}
	if o.NodeIpAddress != nil {
		toSerialize["NodeIpAddress"] = o.NodeIpAddress
	}
	if o.NodeName != nil {
		toSerialize["NodeName"] = o.NodeName
	}
	if o.NodeResolution != nil {
		toSerialize["NodeResolution"] = o.NodeResolution
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexHealthCheckNodeLevelInfoAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexHealthCheckNodeLevelInfoAllOf := _HyperflexHealthCheckNodeLevelInfoAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexHealthCheckNodeLevelInfoAllOf); err == nil {
		*o = HyperflexHealthCheckNodeLevelInfoAllOf(varHyperflexHealthCheckNodeLevelInfoAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "NodeCause")
		delete(additionalProperties, "NodeCheckResult")
		delete(additionalProperties, "NodeEsxIpAddress")
		delete(additionalProperties, "NodeIpAddress")
		delete(additionalProperties, "NodeName")
		delete(additionalProperties, "NodeResolution")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexHealthCheckNodeLevelInfoAllOf struct {
	value *HyperflexHealthCheckNodeLevelInfoAllOf
	isSet bool
}

func (v NullableHyperflexHealthCheckNodeLevelInfoAllOf) Get() *HyperflexHealthCheckNodeLevelInfoAllOf {
	return v.value
}

func (v *NullableHyperflexHealthCheckNodeLevelInfoAllOf) Set(val *HyperflexHealthCheckNodeLevelInfoAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexHealthCheckNodeLevelInfoAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexHealthCheckNodeLevelInfoAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexHealthCheckNodeLevelInfoAllOf(val *HyperflexHealthCheckNodeLevelInfoAllOf) *NullableHyperflexHealthCheckNodeLevelInfoAllOf {
	return &NullableHyperflexHealthCheckNodeLevelInfoAllOf{value: val, isSet: true}
}

func (v NullableHyperflexHealthCheckNodeLevelInfoAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexHealthCheckNodeLevelInfoAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
