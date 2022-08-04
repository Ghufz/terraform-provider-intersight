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

// HyperflexServerModelAllOf Definition of the list of properties defined in 'hyperflex.ServerModel', excluding properties defined in parent classes.
type HyperflexServerModelAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                           `json:"ObjectType"`
	ServerModelEntries   []HyperflexServerModelEntry      `json:"ServerModelEntries,omitempty"`
	AppCatalog           *HyperflexAppCatalogRelationship `json:"AppCatalog,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexServerModelAllOf HyperflexServerModelAllOf

// NewHyperflexServerModelAllOf instantiates a new HyperflexServerModelAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexServerModelAllOf(classId string, objectType string) *HyperflexServerModelAllOf {
	this := HyperflexServerModelAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexServerModelAllOfWithDefaults instantiates a new HyperflexServerModelAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexServerModelAllOfWithDefaults() *HyperflexServerModelAllOf {
	this := HyperflexServerModelAllOf{}
	var classId string = "hyperflex.ServerModel"
	this.ClassId = classId
	var objectType string = "hyperflex.ServerModel"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexServerModelAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexServerModelAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexServerModelAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexServerModelAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexServerModelAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexServerModelAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetServerModelEntries returns the ServerModelEntries field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexServerModelAllOf) GetServerModelEntries() []HyperflexServerModelEntry {
	if o == nil {
		var ret []HyperflexServerModelEntry
		return ret
	}
	return o.ServerModelEntries
}

// GetServerModelEntriesOk returns a tuple with the ServerModelEntries field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexServerModelAllOf) GetServerModelEntriesOk() ([]HyperflexServerModelEntry, bool) {
	if o == nil || o.ServerModelEntries == nil {
		return nil, false
	}
	return o.ServerModelEntries, true
}

// HasServerModelEntries returns a boolean if a field has been set.
func (o *HyperflexServerModelAllOf) HasServerModelEntries() bool {
	if o != nil && o.ServerModelEntries != nil {
		return true
	}

	return false
}

// SetServerModelEntries gets a reference to the given []HyperflexServerModelEntry and assigns it to the ServerModelEntries field.
func (o *HyperflexServerModelAllOf) SetServerModelEntries(v []HyperflexServerModelEntry) {
	o.ServerModelEntries = v
}

// GetAppCatalog returns the AppCatalog field value if set, zero value otherwise.
func (o *HyperflexServerModelAllOf) GetAppCatalog() HyperflexAppCatalogRelationship {
	if o == nil || o.AppCatalog == nil {
		var ret HyperflexAppCatalogRelationship
		return ret
	}
	return *o.AppCatalog
}

// GetAppCatalogOk returns a tuple with the AppCatalog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexServerModelAllOf) GetAppCatalogOk() (*HyperflexAppCatalogRelationship, bool) {
	if o == nil || o.AppCatalog == nil {
		return nil, false
	}
	return o.AppCatalog, true
}

// HasAppCatalog returns a boolean if a field has been set.
func (o *HyperflexServerModelAllOf) HasAppCatalog() bool {
	if o != nil && o.AppCatalog != nil {
		return true
	}

	return false
}

// SetAppCatalog gets a reference to the given HyperflexAppCatalogRelationship and assigns it to the AppCatalog field.
func (o *HyperflexServerModelAllOf) SetAppCatalog(v HyperflexAppCatalogRelationship) {
	o.AppCatalog = &v
}

func (o HyperflexServerModelAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ServerModelEntries != nil {
		toSerialize["ServerModelEntries"] = o.ServerModelEntries
	}
	if o.AppCatalog != nil {
		toSerialize["AppCatalog"] = o.AppCatalog
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexServerModelAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexServerModelAllOf := _HyperflexServerModelAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexServerModelAllOf); err == nil {
		*o = HyperflexServerModelAllOf(varHyperflexServerModelAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ServerModelEntries")
		delete(additionalProperties, "AppCatalog")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexServerModelAllOf struct {
	value *HyperflexServerModelAllOf
	isSet bool
}

func (v NullableHyperflexServerModelAllOf) Get() *HyperflexServerModelAllOf {
	return v.value
}

func (v *NullableHyperflexServerModelAllOf) Set(val *HyperflexServerModelAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexServerModelAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexServerModelAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexServerModelAllOf(val *HyperflexServerModelAllOf) *NullableHyperflexServerModelAllOf {
	return &NullableHyperflexServerModelAllOf{value: val, isSet: true}
}

func (v NullableHyperflexServerModelAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexServerModelAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
