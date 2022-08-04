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
)

// ResourceSourceToPermissionResourcesHolderAllOf Definition of the list of properties defined in 'resource.SourceToPermissionResourcesHolder', excluding properties defined in parent classes.
type ResourceSourceToPermissionResourcesHolderAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType                  string                                `json:"ObjectType"`
	SourceToPermissionResources []ResourceSourceToPermissionResources `json:"SourceToPermissionResources,omitempty"`
	AdditionalProperties        map[string]interface{}
}

type _ResourceSourceToPermissionResourcesHolderAllOf ResourceSourceToPermissionResourcesHolderAllOf

// NewResourceSourceToPermissionResourcesHolderAllOf instantiates a new ResourceSourceToPermissionResourcesHolderAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceSourceToPermissionResourcesHolderAllOf(classId string, objectType string) *ResourceSourceToPermissionResourcesHolderAllOf {
	this := ResourceSourceToPermissionResourcesHolderAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewResourceSourceToPermissionResourcesHolderAllOfWithDefaults instantiates a new ResourceSourceToPermissionResourcesHolderAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceSourceToPermissionResourcesHolderAllOfWithDefaults() *ResourceSourceToPermissionResourcesHolderAllOf {
	this := ResourceSourceToPermissionResourcesHolderAllOf{}
	var classId string = "resource.SourceToPermissionResourcesHolder"
	this.ClassId = classId
	var objectType string = "resource.SourceToPermissionResourcesHolder"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ResourceSourceToPermissionResourcesHolderAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ResourceSourceToPermissionResourcesHolderAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ResourceSourceToPermissionResourcesHolderAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ResourceSourceToPermissionResourcesHolderAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ResourceSourceToPermissionResourcesHolderAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ResourceSourceToPermissionResourcesHolderAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetSourceToPermissionResources returns the SourceToPermissionResources field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourceSourceToPermissionResourcesHolderAllOf) GetSourceToPermissionResources() []ResourceSourceToPermissionResources {
	if o == nil {
		var ret []ResourceSourceToPermissionResources
		return ret
	}
	return o.SourceToPermissionResources
}

// GetSourceToPermissionResourcesOk returns a tuple with the SourceToPermissionResources field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourceSourceToPermissionResourcesHolderAllOf) GetSourceToPermissionResourcesOk() ([]ResourceSourceToPermissionResources, bool) {
	if o == nil || o.SourceToPermissionResources == nil {
		return nil, false
	}
	return o.SourceToPermissionResources, true
}

// HasSourceToPermissionResources returns a boolean if a field has been set.
func (o *ResourceSourceToPermissionResourcesHolderAllOf) HasSourceToPermissionResources() bool {
	if o != nil && o.SourceToPermissionResources != nil {
		return true
	}

	return false
}

// SetSourceToPermissionResources gets a reference to the given []ResourceSourceToPermissionResources and assigns it to the SourceToPermissionResources field.
func (o *ResourceSourceToPermissionResourcesHolderAllOf) SetSourceToPermissionResources(v []ResourceSourceToPermissionResources) {
	o.SourceToPermissionResources = v
}

func (o ResourceSourceToPermissionResourcesHolderAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.SourceToPermissionResources != nil {
		toSerialize["SourceToPermissionResources"] = o.SourceToPermissionResources
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ResourceSourceToPermissionResourcesHolderAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varResourceSourceToPermissionResourcesHolderAllOf := _ResourceSourceToPermissionResourcesHolderAllOf{}

	if err = json.Unmarshal(bytes, &varResourceSourceToPermissionResourcesHolderAllOf); err == nil {
		*o = ResourceSourceToPermissionResourcesHolderAllOf(varResourceSourceToPermissionResourcesHolderAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "SourceToPermissionResources")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResourceSourceToPermissionResourcesHolderAllOf struct {
	value *ResourceSourceToPermissionResourcesHolderAllOf
	isSet bool
}

func (v NullableResourceSourceToPermissionResourcesHolderAllOf) Get() *ResourceSourceToPermissionResourcesHolderAllOf {
	return v.value
}

func (v *NullableResourceSourceToPermissionResourcesHolderAllOf) Set(val *ResourceSourceToPermissionResourcesHolderAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceSourceToPermissionResourcesHolderAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceSourceToPermissionResourcesHolderAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceSourceToPermissionResourcesHolderAllOf(val *ResourceSourceToPermissionResourcesHolderAllOf) *NullableResourceSourceToPermissionResourcesHolderAllOf {
	return &NullableResourceSourceToPermissionResourcesHolderAllOf{value: val, isSet: true}
}

func (v NullableResourceSourceToPermissionResourcesHolderAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceSourceToPermissionResourcesHolderAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
