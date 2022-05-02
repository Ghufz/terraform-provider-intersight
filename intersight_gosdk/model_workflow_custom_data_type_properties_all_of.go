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

// WorkflowCustomDataTypePropertiesAllOf Definition of the list of properties defined in 'workflow.CustomDataTypeProperties', excluding properties defined in parent classes.
type WorkflowCustomDataTypePropertiesAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// When set to false custom data type is not cloneable. It is set to true only if data type is not internal and it is not using any internal custom data type.
	Cloneable *bool `json:"Cloneable,omitempty"`
	// When set to false the custom data type is owned by the system and used for internal services. Such custom data type cannot be directly used by external entities.
	ExternalMeta         *bool `json:"ExternalMeta,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowCustomDataTypePropertiesAllOf WorkflowCustomDataTypePropertiesAllOf

// NewWorkflowCustomDataTypePropertiesAllOf instantiates a new WorkflowCustomDataTypePropertiesAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowCustomDataTypePropertiesAllOf(classId string, objectType string) *WorkflowCustomDataTypePropertiesAllOf {
	this := WorkflowCustomDataTypePropertiesAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowCustomDataTypePropertiesAllOfWithDefaults instantiates a new WorkflowCustomDataTypePropertiesAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowCustomDataTypePropertiesAllOfWithDefaults() *WorkflowCustomDataTypePropertiesAllOf {
	this := WorkflowCustomDataTypePropertiesAllOf{}
	var classId string = "workflow.CustomDataTypeProperties"
	this.ClassId = classId
	var objectType string = "workflow.CustomDataTypeProperties"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowCustomDataTypePropertiesAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowCustomDataTypePropertiesAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowCustomDataTypePropertiesAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowCustomDataTypePropertiesAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowCustomDataTypePropertiesAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowCustomDataTypePropertiesAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCloneable returns the Cloneable field value if set, zero value otherwise.
func (o *WorkflowCustomDataTypePropertiesAllOf) GetCloneable() bool {
	if o == nil || o.Cloneable == nil {
		var ret bool
		return ret
	}
	return *o.Cloneable
}

// GetCloneableOk returns a tuple with the Cloneable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowCustomDataTypePropertiesAllOf) GetCloneableOk() (*bool, bool) {
	if o == nil || o.Cloneable == nil {
		return nil, false
	}
	return o.Cloneable, true
}

// HasCloneable returns a boolean if a field has been set.
func (o *WorkflowCustomDataTypePropertiesAllOf) HasCloneable() bool {
	if o != nil && o.Cloneable != nil {
		return true
	}

	return false
}

// SetCloneable gets a reference to the given bool and assigns it to the Cloneable field.
func (o *WorkflowCustomDataTypePropertiesAllOf) SetCloneable(v bool) {
	o.Cloneable = &v
}

// GetExternalMeta returns the ExternalMeta field value if set, zero value otherwise.
func (o *WorkflowCustomDataTypePropertiesAllOf) GetExternalMeta() bool {
	if o == nil || o.ExternalMeta == nil {
		var ret bool
		return ret
	}
	return *o.ExternalMeta
}

// GetExternalMetaOk returns a tuple with the ExternalMeta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowCustomDataTypePropertiesAllOf) GetExternalMetaOk() (*bool, bool) {
	if o == nil || o.ExternalMeta == nil {
		return nil, false
	}
	return o.ExternalMeta, true
}

// HasExternalMeta returns a boolean if a field has been set.
func (o *WorkflowCustomDataTypePropertiesAllOf) HasExternalMeta() bool {
	if o != nil && o.ExternalMeta != nil {
		return true
	}

	return false
}

// SetExternalMeta gets a reference to the given bool and assigns it to the ExternalMeta field.
func (o *WorkflowCustomDataTypePropertiesAllOf) SetExternalMeta(v bool) {
	o.ExternalMeta = &v
}

func (o WorkflowCustomDataTypePropertiesAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Cloneable != nil {
		toSerialize["Cloneable"] = o.Cloneable
	}
	if o.ExternalMeta != nil {
		toSerialize["ExternalMeta"] = o.ExternalMeta
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowCustomDataTypePropertiesAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowCustomDataTypePropertiesAllOf := _WorkflowCustomDataTypePropertiesAllOf{}

	if err = json.Unmarshal(bytes, &varWorkflowCustomDataTypePropertiesAllOf); err == nil {
		*o = WorkflowCustomDataTypePropertiesAllOf(varWorkflowCustomDataTypePropertiesAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Cloneable")
		delete(additionalProperties, "ExternalMeta")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowCustomDataTypePropertiesAllOf struct {
	value *WorkflowCustomDataTypePropertiesAllOf
	isSet bool
}

func (v NullableWorkflowCustomDataTypePropertiesAllOf) Get() *WorkflowCustomDataTypePropertiesAllOf {
	return v.value
}

func (v *NullableWorkflowCustomDataTypePropertiesAllOf) Set(val *WorkflowCustomDataTypePropertiesAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowCustomDataTypePropertiesAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowCustomDataTypePropertiesAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowCustomDataTypePropertiesAllOf(val *WorkflowCustomDataTypePropertiesAllOf) *NullableWorkflowCustomDataTypePropertiesAllOf {
	return &NullableWorkflowCustomDataTypePropertiesAllOf{value: val, isSet: true}
}

func (v NullableWorkflowCustomDataTypePropertiesAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowCustomDataTypePropertiesAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
