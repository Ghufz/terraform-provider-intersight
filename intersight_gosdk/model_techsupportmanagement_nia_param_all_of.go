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

// TechsupportmanagementNiaParamAllOf Definition of the list of properties defined in 'techsupportmanagement.NiaParam', excluding properties defined in parent classes.
type TechsupportmanagementNiaParamAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// CollectionLevel controls the size / depth of the tech support files collected. * `1` - Use cached data in the returned collection. * `2` - Use current data in the returned collection.
	CollectionLevel *int32 `json:"CollectionLevel,omitempty"`
	// Filename specifies an individual filename to collect from the endpoint.
	Filename *string `json:"Filename,omitempty"`
	// ForceFresh controls whether to return pre-collected files or force the collection of new files.
	ForceFresh           *bool    `json:"ForceFresh,omitempty"`
	Pids                 []string `json:"Pids,omitempty"`
	SerialNumbers        []string `json:"SerialNumbers,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TechsupportmanagementNiaParamAllOf TechsupportmanagementNiaParamAllOf

// NewTechsupportmanagementNiaParamAllOf instantiates a new TechsupportmanagementNiaParamAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTechsupportmanagementNiaParamAllOf(classId string, objectType string) *TechsupportmanagementNiaParamAllOf {
	this := TechsupportmanagementNiaParamAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var collectionLevel int32 = 1
	this.CollectionLevel = &collectionLevel
	return &this
}

// NewTechsupportmanagementNiaParamAllOfWithDefaults instantiates a new TechsupportmanagementNiaParamAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTechsupportmanagementNiaParamAllOfWithDefaults() *TechsupportmanagementNiaParamAllOf {
	this := TechsupportmanagementNiaParamAllOf{}
	var classId string = "techsupportmanagement.NiaParam"
	this.ClassId = classId
	var objectType string = "techsupportmanagement.NiaParam"
	this.ObjectType = objectType
	var collectionLevel int32 = 1
	this.CollectionLevel = &collectionLevel
	return &this
}

// GetClassId returns the ClassId field value
func (o *TechsupportmanagementNiaParamAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementNiaParamAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *TechsupportmanagementNiaParamAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *TechsupportmanagementNiaParamAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementNiaParamAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *TechsupportmanagementNiaParamAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCollectionLevel returns the CollectionLevel field value if set, zero value otherwise.
func (o *TechsupportmanagementNiaParamAllOf) GetCollectionLevel() int32 {
	if o == nil || o.CollectionLevel == nil {
		var ret int32
		return ret
	}
	return *o.CollectionLevel
}

// GetCollectionLevelOk returns a tuple with the CollectionLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementNiaParamAllOf) GetCollectionLevelOk() (*int32, bool) {
	if o == nil || o.CollectionLevel == nil {
		return nil, false
	}
	return o.CollectionLevel, true
}

// HasCollectionLevel returns a boolean if a field has been set.
func (o *TechsupportmanagementNiaParamAllOf) HasCollectionLevel() bool {
	if o != nil && o.CollectionLevel != nil {
		return true
	}

	return false
}

// SetCollectionLevel gets a reference to the given int32 and assigns it to the CollectionLevel field.
func (o *TechsupportmanagementNiaParamAllOf) SetCollectionLevel(v int32) {
	o.CollectionLevel = &v
}

// GetFilename returns the Filename field value if set, zero value otherwise.
func (o *TechsupportmanagementNiaParamAllOf) GetFilename() string {
	if o == nil || o.Filename == nil {
		var ret string
		return ret
	}
	return *o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementNiaParamAllOf) GetFilenameOk() (*string, bool) {
	if o == nil || o.Filename == nil {
		return nil, false
	}
	return o.Filename, true
}

// HasFilename returns a boolean if a field has been set.
func (o *TechsupportmanagementNiaParamAllOf) HasFilename() bool {
	if o != nil && o.Filename != nil {
		return true
	}

	return false
}

// SetFilename gets a reference to the given string and assigns it to the Filename field.
func (o *TechsupportmanagementNiaParamAllOf) SetFilename(v string) {
	o.Filename = &v
}

// GetForceFresh returns the ForceFresh field value if set, zero value otherwise.
func (o *TechsupportmanagementNiaParamAllOf) GetForceFresh() bool {
	if o == nil || o.ForceFresh == nil {
		var ret bool
		return ret
	}
	return *o.ForceFresh
}

// GetForceFreshOk returns a tuple with the ForceFresh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementNiaParamAllOf) GetForceFreshOk() (*bool, bool) {
	if o == nil || o.ForceFresh == nil {
		return nil, false
	}
	return o.ForceFresh, true
}

// HasForceFresh returns a boolean if a field has been set.
func (o *TechsupportmanagementNiaParamAllOf) HasForceFresh() bool {
	if o != nil && o.ForceFresh != nil {
		return true
	}

	return false
}

// SetForceFresh gets a reference to the given bool and assigns it to the ForceFresh field.
func (o *TechsupportmanagementNiaParamAllOf) SetForceFresh(v bool) {
	o.ForceFresh = &v
}

// GetPids returns the Pids field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TechsupportmanagementNiaParamAllOf) GetPids() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Pids
}

// GetPidsOk returns a tuple with the Pids field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TechsupportmanagementNiaParamAllOf) GetPidsOk() ([]string, bool) {
	if o == nil || o.Pids == nil {
		return nil, false
	}
	return o.Pids, true
}

// HasPids returns a boolean if a field has been set.
func (o *TechsupportmanagementNiaParamAllOf) HasPids() bool {
	if o != nil && o.Pids != nil {
		return true
	}

	return false
}

// SetPids gets a reference to the given []string and assigns it to the Pids field.
func (o *TechsupportmanagementNiaParamAllOf) SetPids(v []string) {
	o.Pids = v
}

// GetSerialNumbers returns the SerialNumbers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TechsupportmanagementNiaParamAllOf) GetSerialNumbers() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SerialNumbers
}

// GetSerialNumbersOk returns a tuple with the SerialNumbers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TechsupportmanagementNiaParamAllOf) GetSerialNumbersOk() ([]string, bool) {
	if o == nil || o.SerialNumbers == nil {
		return nil, false
	}
	return o.SerialNumbers, true
}

// HasSerialNumbers returns a boolean if a field has been set.
func (o *TechsupportmanagementNiaParamAllOf) HasSerialNumbers() bool {
	if o != nil && o.SerialNumbers != nil {
		return true
	}

	return false
}

// SetSerialNumbers gets a reference to the given []string and assigns it to the SerialNumbers field.
func (o *TechsupportmanagementNiaParamAllOf) SetSerialNumbers(v []string) {
	o.SerialNumbers = v
}

func (o TechsupportmanagementNiaParamAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.CollectionLevel != nil {
		toSerialize["CollectionLevel"] = o.CollectionLevel
	}
	if o.Filename != nil {
		toSerialize["Filename"] = o.Filename
	}
	if o.ForceFresh != nil {
		toSerialize["ForceFresh"] = o.ForceFresh
	}
	if o.Pids != nil {
		toSerialize["Pids"] = o.Pids
	}
	if o.SerialNumbers != nil {
		toSerialize["SerialNumbers"] = o.SerialNumbers
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TechsupportmanagementNiaParamAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varTechsupportmanagementNiaParamAllOf := _TechsupportmanagementNiaParamAllOf{}

	if err = json.Unmarshal(bytes, &varTechsupportmanagementNiaParamAllOf); err == nil {
		*o = TechsupportmanagementNiaParamAllOf(varTechsupportmanagementNiaParamAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CollectionLevel")
		delete(additionalProperties, "Filename")
		delete(additionalProperties, "ForceFresh")
		delete(additionalProperties, "Pids")
		delete(additionalProperties, "SerialNumbers")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTechsupportmanagementNiaParamAllOf struct {
	value *TechsupportmanagementNiaParamAllOf
	isSet bool
}

func (v NullableTechsupportmanagementNiaParamAllOf) Get() *TechsupportmanagementNiaParamAllOf {
	return v.value
}

func (v *NullableTechsupportmanagementNiaParamAllOf) Set(val *TechsupportmanagementNiaParamAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTechsupportmanagementNiaParamAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTechsupportmanagementNiaParamAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTechsupportmanagementNiaParamAllOf(val *TechsupportmanagementNiaParamAllOf) *NullableTechsupportmanagementNiaParamAllOf {
	return &NullableTechsupportmanagementNiaParamAllOf{value: val, isSet: true}
}

func (v NullableTechsupportmanagementNiaParamAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTechsupportmanagementNiaParamAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
