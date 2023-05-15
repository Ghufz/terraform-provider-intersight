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
)

// MetaDisplayNameDefinitionAllOf Definition of the list of properties defined in 'meta.DisplayNameDefinition', excluding properties defined in parent classes.
type MetaDisplayNameDefinitionAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// A specification for constructing the displayname from the MO's properties.
	Format *string `json:"Format,omitempty"`
	// An indication of whether the displayname should be contructed 'recursively' including the displayname of the first ancestor with a similarly named displayname.
	IncludeAncestor *bool `json:"IncludeAncestor,omitempty"`
	// The name of the displayname used as a key in the DisplayName map which is returned as part of an MO for a Rest request.
	Name                 *string `json:"Name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MetaDisplayNameDefinitionAllOf MetaDisplayNameDefinitionAllOf

// NewMetaDisplayNameDefinitionAllOf instantiates a new MetaDisplayNameDefinitionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetaDisplayNameDefinitionAllOf(classId string, objectType string) *MetaDisplayNameDefinitionAllOf {
	this := MetaDisplayNameDefinitionAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewMetaDisplayNameDefinitionAllOfWithDefaults instantiates a new MetaDisplayNameDefinitionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetaDisplayNameDefinitionAllOfWithDefaults() *MetaDisplayNameDefinitionAllOf {
	this := MetaDisplayNameDefinitionAllOf{}
	var classId string = "meta.DisplayNameDefinition"
	this.ClassId = classId
	var objectType string = "meta.DisplayNameDefinition"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *MetaDisplayNameDefinitionAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *MetaDisplayNameDefinitionAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *MetaDisplayNameDefinitionAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *MetaDisplayNameDefinitionAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *MetaDisplayNameDefinitionAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *MetaDisplayNameDefinitionAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *MetaDisplayNameDefinitionAllOf) GetFormat() string {
	if o == nil || o.Format == nil {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaDisplayNameDefinitionAllOf) GetFormatOk() (*string, bool) {
	if o == nil || o.Format == nil {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *MetaDisplayNameDefinitionAllOf) HasFormat() bool {
	if o != nil && o.Format != nil {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *MetaDisplayNameDefinitionAllOf) SetFormat(v string) {
	o.Format = &v
}

// GetIncludeAncestor returns the IncludeAncestor field value if set, zero value otherwise.
func (o *MetaDisplayNameDefinitionAllOf) GetIncludeAncestor() bool {
	if o == nil || o.IncludeAncestor == nil {
		var ret bool
		return ret
	}
	return *o.IncludeAncestor
}

// GetIncludeAncestorOk returns a tuple with the IncludeAncestor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaDisplayNameDefinitionAllOf) GetIncludeAncestorOk() (*bool, bool) {
	if o == nil || o.IncludeAncestor == nil {
		return nil, false
	}
	return o.IncludeAncestor, true
}

// HasIncludeAncestor returns a boolean if a field has been set.
func (o *MetaDisplayNameDefinitionAllOf) HasIncludeAncestor() bool {
	if o != nil && o.IncludeAncestor != nil {
		return true
	}

	return false
}

// SetIncludeAncestor gets a reference to the given bool and assigns it to the IncludeAncestor field.
func (o *MetaDisplayNameDefinitionAllOf) SetIncludeAncestor(v bool) {
	o.IncludeAncestor = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MetaDisplayNameDefinitionAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaDisplayNameDefinitionAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MetaDisplayNameDefinitionAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MetaDisplayNameDefinitionAllOf) SetName(v string) {
	o.Name = &v
}

func (o MetaDisplayNameDefinitionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Format != nil {
		toSerialize["Format"] = o.Format
	}
	if o.IncludeAncestor != nil {
		toSerialize["IncludeAncestor"] = o.IncludeAncestor
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MetaDisplayNameDefinitionAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varMetaDisplayNameDefinitionAllOf := _MetaDisplayNameDefinitionAllOf{}

	if err = json.Unmarshal(bytes, &varMetaDisplayNameDefinitionAllOf); err == nil {
		*o = MetaDisplayNameDefinitionAllOf(varMetaDisplayNameDefinitionAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Format")
		delete(additionalProperties, "IncludeAncestor")
		delete(additionalProperties, "Name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMetaDisplayNameDefinitionAllOf struct {
	value *MetaDisplayNameDefinitionAllOf
	isSet bool
}

func (v NullableMetaDisplayNameDefinitionAllOf) Get() *MetaDisplayNameDefinitionAllOf {
	return v.value
}

func (v *NullableMetaDisplayNameDefinitionAllOf) Set(val *MetaDisplayNameDefinitionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMetaDisplayNameDefinitionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMetaDisplayNameDefinitionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetaDisplayNameDefinitionAllOf(val *MetaDisplayNameDefinitionAllOf) *NullableMetaDisplayNameDefinitionAllOf {
	return &NullableMetaDisplayNameDefinitionAllOf{value: val, isSet: true}
}

func (v NullableMetaDisplayNameDefinitionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetaDisplayNameDefinitionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
