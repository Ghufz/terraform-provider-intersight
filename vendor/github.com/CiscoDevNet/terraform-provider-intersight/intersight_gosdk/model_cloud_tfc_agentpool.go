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

// CloudTfcAgentpool An agent pool represents a group of agents that can be used to allow Terraform Cloud to communicate with isolated, private, or on-premises infrastructure.
type CloudTfcAgentpool struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The ID of the Agent Pool.
	Identity *string `json:"Identity,omitempty"`
	// The name of the agent pool.
	Name *string `json:"Name,omitempty"`
	// The number of active agents used by this pool. The total active agent are sum of idle, busy and unknown agent counts.
	NumActiveAgents *int64 `json:"NumActiveAgents,omitempty"`
	// The number of Tokens in this agent Pool.
	NumTokens            *int64                            `json:"NumTokens,omitempty"`
	Organization         *CloudTfcOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CloudTfcAgentpool CloudTfcAgentpool

// NewCloudTfcAgentpool instantiates a new CloudTfcAgentpool object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudTfcAgentpool(classId string, objectType string) *CloudTfcAgentpool {
	this := CloudTfcAgentpool{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCloudTfcAgentpoolWithDefaults instantiates a new CloudTfcAgentpool object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudTfcAgentpoolWithDefaults() *CloudTfcAgentpool {
	this := CloudTfcAgentpool{}
	var classId string = "cloud.TfcAgentpool"
	this.ClassId = classId
	var objectType string = "cloud.TfcAgentpool"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CloudTfcAgentpool) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CloudTfcAgentpool) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CloudTfcAgentpool) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CloudTfcAgentpool) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CloudTfcAgentpool) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CloudTfcAgentpool) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *CloudTfcAgentpool) GetIdentity() string {
	if o == nil || o.Identity == nil {
		var ret string
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudTfcAgentpool) GetIdentityOk() (*string, bool) {
	if o == nil || o.Identity == nil {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *CloudTfcAgentpool) HasIdentity() bool {
	if o != nil && o.Identity != nil {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given string and assigns it to the Identity field.
func (o *CloudTfcAgentpool) SetIdentity(v string) {
	o.Identity = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CloudTfcAgentpool) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudTfcAgentpool) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CloudTfcAgentpool) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CloudTfcAgentpool) SetName(v string) {
	o.Name = &v
}

// GetNumActiveAgents returns the NumActiveAgents field value if set, zero value otherwise.
func (o *CloudTfcAgentpool) GetNumActiveAgents() int64 {
	if o == nil || o.NumActiveAgents == nil {
		var ret int64
		return ret
	}
	return *o.NumActiveAgents
}

// GetNumActiveAgentsOk returns a tuple with the NumActiveAgents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudTfcAgentpool) GetNumActiveAgentsOk() (*int64, bool) {
	if o == nil || o.NumActiveAgents == nil {
		return nil, false
	}
	return o.NumActiveAgents, true
}

// HasNumActiveAgents returns a boolean if a field has been set.
func (o *CloudTfcAgentpool) HasNumActiveAgents() bool {
	if o != nil && o.NumActiveAgents != nil {
		return true
	}

	return false
}

// SetNumActiveAgents gets a reference to the given int64 and assigns it to the NumActiveAgents field.
func (o *CloudTfcAgentpool) SetNumActiveAgents(v int64) {
	o.NumActiveAgents = &v
}

// GetNumTokens returns the NumTokens field value if set, zero value otherwise.
func (o *CloudTfcAgentpool) GetNumTokens() int64 {
	if o == nil || o.NumTokens == nil {
		var ret int64
		return ret
	}
	return *o.NumTokens
}

// GetNumTokensOk returns a tuple with the NumTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudTfcAgentpool) GetNumTokensOk() (*int64, bool) {
	if o == nil || o.NumTokens == nil {
		return nil, false
	}
	return o.NumTokens, true
}

// HasNumTokens returns a boolean if a field has been set.
func (o *CloudTfcAgentpool) HasNumTokens() bool {
	if o != nil && o.NumTokens != nil {
		return true
	}

	return false
}

// SetNumTokens gets a reference to the given int64 and assigns it to the NumTokens field.
func (o *CloudTfcAgentpool) SetNumTokens(v int64) {
	o.NumTokens = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *CloudTfcAgentpool) GetOrganization() CloudTfcOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret CloudTfcOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudTfcAgentpool) GetOrganizationOk() (*CloudTfcOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *CloudTfcAgentpool) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given CloudTfcOrganizationRelationship and assigns it to the Organization field.
func (o *CloudTfcAgentpool) SetOrganization(v CloudTfcOrganizationRelationship) {
	o.Organization = &v
}

func (o CloudTfcAgentpool) MarshalJSON() ([]byte, error) {
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
	if o.Identity != nil {
		toSerialize["Identity"] = o.Identity
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.NumActiveAgents != nil {
		toSerialize["NumActiveAgents"] = o.NumActiveAgents
	}
	if o.NumTokens != nil {
		toSerialize["NumTokens"] = o.NumTokens
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CloudTfcAgentpool) UnmarshalJSON(bytes []byte) (err error) {
	type CloudTfcAgentpoolWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The ID of the Agent Pool.
		Identity *string `json:"Identity,omitempty"`
		// The name of the agent pool.
		Name *string `json:"Name,omitempty"`
		// The number of active agents used by this pool. The total active agent are sum of idle, busy and unknown agent counts.
		NumActiveAgents *int64 `json:"NumActiveAgents,omitempty"`
		// The number of Tokens in this agent Pool.
		NumTokens    *int64                            `json:"NumTokens,omitempty"`
		Organization *CloudTfcOrganizationRelationship `json:"Organization,omitempty"`
	}

	varCloudTfcAgentpoolWithoutEmbeddedStruct := CloudTfcAgentpoolWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varCloudTfcAgentpoolWithoutEmbeddedStruct)
	if err == nil {
		varCloudTfcAgentpool := _CloudTfcAgentpool{}
		varCloudTfcAgentpool.ClassId = varCloudTfcAgentpoolWithoutEmbeddedStruct.ClassId
		varCloudTfcAgentpool.ObjectType = varCloudTfcAgentpoolWithoutEmbeddedStruct.ObjectType
		varCloudTfcAgentpool.Identity = varCloudTfcAgentpoolWithoutEmbeddedStruct.Identity
		varCloudTfcAgentpool.Name = varCloudTfcAgentpoolWithoutEmbeddedStruct.Name
		varCloudTfcAgentpool.NumActiveAgents = varCloudTfcAgentpoolWithoutEmbeddedStruct.NumActiveAgents
		varCloudTfcAgentpool.NumTokens = varCloudTfcAgentpoolWithoutEmbeddedStruct.NumTokens
		varCloudTfcAgentpool.Organization = varCloudTfcAgentpoolWithoutEmbeddedStruct.Organization
		*o = CloudTfcAgentpool(varCloudTfcAgentpool)
	} else {
		return err
	}

	varCloudTfcAgentpool := _CloudTfcAgentpool{}

	err = json.Unmarshal(bytes, &varCloudTfcAgentpool)
	if err == nil {
		o.MoBaseMo = varCloudTfcAgentpool.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Identity")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "NumActiveAgents")
		delete(additionalProperties, "NumTokens")
		delete(additionalProperties, "Organization")

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

type NullableCloudTfcAgentpool struct {
	value *CloudTfcAgentpool
	isSet bool
}

func (v NullableCloudTfcAgentpool) Get() *CloudTfcAgentpool {
	return v.value
}

func (v *NullableCloudTfcAgentpool) Set(val *CloudTfcAgentpool) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudTfcAgentpool) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudTfcAgentpool) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudTfcAgentpool(val *CloudTfcAgentpool) *NullableCloudTfcAgentpool {
	return &NullableCloudTfcAgentpool{value: val, isSet: true}
}

func (v NullableCloudTfcAgentpool) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudTfcAgentpool) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
