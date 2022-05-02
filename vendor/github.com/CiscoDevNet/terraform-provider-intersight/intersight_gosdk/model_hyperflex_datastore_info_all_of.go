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

// HyperflexDatastoreInfoAllOf Definition of the list of properties defined in 'hyperflex.DatastoreInfo', excluding properties defined in parent classes.
type HyperflexDatastoreInfoAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// This datastore's backend unique id.
	DsBackendId *string `json:"DsBackendId,omitempty"`
	// This datastore's frontend id.
	DsFrontendId         *string `json:"DsFrontendId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexDatastoreInfoAllOf HyperflexDatastoreInfoAllOf

// NewHyperflexDatastoreInfoAllOf instantiates a new HyperflexDatastoreInfoAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexDatastoreInfoAllOf(classId string, objectType string) *HyperflexDatastoreInfoAllOf {
	this := HyperflexDatastoreInfoAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexDatastoreInfoAllOfWithDefaults instantiates a new HyperflexDatastoreInfoAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexDatastoreInfoAllOfWithDefaults() *HyperflexDatastoreInfoAllOf {
	this := HyperflexDatastoreInfoAllOf{}
	var classId string = "hyperflex.DatastoreInfo"
	this.ClassId = classId
	var objectType string = "hyperflex.DatastoreInfo"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexDatastoreInfoAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexDatastoreInfoAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexDatastoreInfoAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexDatastoreInfoAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexDatastoreInfoAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexDatastoreInfoAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDsBackendId returns the DsBackendId field value if set, zero value otherwise.
func (o *HyperflexDatastoreInfoAllOf) GetDsBackendId() string {
	if o == nil || o.DsBackendId == nil {
		var ret string
		return ret
	}
	return *o.DsBackendId
}

// GetDsBackendIdOk returns a tuple with the DsBackendId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexDatastoreInfoAllOf) GetDsBackendIdOk() (*string, bool) {
	if o == nil || o.DsBackendId == nil {
		return nil, false
	}
	return o.DsBackendId, true
}

// HasDsBackendId returns a boolean if a field has been set.
func (o *HyperflexDatastoreInfoAllOf) HasDsBackendId() bool {
	if o != nil && o.DsBackendId != nil {
		return true
	}

	return false
}

// SetDsBackendId gets a reference to the given string and assigns it to the DsBackendId field.
func (o *HyperflexDatastoreInfoAllOf) SetDsBackendId(v string) {
	o.DsBackendId = &v
}

// GetDsFrontendId returns the DsFrontendId field value if set, zero value otherwise.
func (o *HyperflexDatastoreInfoAllOf) GetDsFrontendId() string {
	if o == nil || o.DsFrontendId == nil {
		var ret string
		return ret
	}
	return *o.DsFrontendId
}

// GetDsFrontendIdOk returns a tuple with the DsFrontendId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexDatastoreInfoAllOf) GetDsFrontendIdOk() (*string, bool) {
	if o == nil || o.DsFrontendId == nil {
		return nil, false
	}
	return o.DsFrontendId, true
}

// HasDsFrontendId returns a boolean if a field has been set.
func (o *HyperflexDatastoreInfoAllOf) HasDsFrontendId() bool {
	if o != nil && o.DsFrontendId != nil {
		return true
	}

	return false
}

// SetDsFrontendId gets a reference to the given string and assigns it to the DsFrontendId field.
func (o *HyperflexDatastoreInfoAllOf) SetDsFrontendId(v string) {
	o.DsFrontendId = &v
}

func (o HyperflexDatastoreInfoAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DsBackendId != nil {
		toSerialize["DsBackendId"] = o.DsBackendId
	}
	if o.DsFrontendId != nil {
		toSerialize["DsFrontendId"] = o.DsFrontendId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexDatastoreInfoAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexDatastoreInfoAllOf := _HyperflexDatastoreInfoAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexDatastoreInfoAllOf); err == nil {
		*o = HyperflexDatastoreInfoAllOf(varHyperflexDatastoreInfoAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DsBackendId")
		delete(additionalProperties, "DsFrontendId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexDatastoreInfoAllOf struct {
	value *HyperflexDatastoreInfoAllOf
	isSet bool
}

func (v NullableHyperflexDatastoreInfoAllOf) Get() *HyperflexDatastoreInfoAllOf {
	return v.value
}

func (v *NullableHyperflexDatastoreInfoAllOf) Set(val *HyperflexDatastoreInfoAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexDatastoreInfoAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexDatastoreInfoAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexDatastoreInfoAllOf(val *HyperflexDatastoreInfoAllOf) *NullableHyperflexDatastoreInfoAllOf {
	return &NullableHyperflexDatastoreInfoAllOf{value: val, isSet: true}
}

func (v NullableHyperflexDatastoreInfoAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexDatastoreInfoAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
