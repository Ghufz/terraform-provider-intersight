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
	"reflect"
	"strings"
)

// OpenapiApiMethodMeta Contains metadata about APIs in the previously uploaded OpenAPI specification file.
type OpenapiApiMethodMeta struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The description of the given API.
	Description *string `json:"Description,omitempty"`
	// The display label of the given API.
	DisplayLabel *string `json:"DisplayLabel,omitempty"`
	// The method type for the given API. * `GET` - Method type which indicates it is a GET API call. * `POST` - Method type which indicates it is a POST API call. * `PUT` - Method type which indicates it is a PUT API call. * `PATCH` - Method type which indicates it is a PATCH API call. * `DELETE` - Method type which indicates it is a DELETE API call.
	Method *string `json:"Method,omitempty"`
	// The description of the given API.
	Name *string `json:"Name,omitempty"`
	// Path of the selected API endpoint.
	Path                 *string                         `json:"Path,omitempty"`
	Source               *OpenapiProcessFileRelationship `json:"Source,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OpenapiApiMethodMeta OpenapiApiMethodMeta

// NewOpenapiApiMethodMeta instantiates a new OpenapiApiMethodMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenapiApiMethodMeta(classId string, objectType string) *OpenapiApiMethodMeta {
	this := OpenapiApiMethodMeta{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewOpenapiApiMethodMetaWithDefaults instantiates a new OpenapiApiMethodMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenapiApiMethodMetaWithDefaults() *OpenapiApiMethodMeta {
	this := OpenapiApiMethodMeta{}
	var classId string = "openapi.ApiMethodMeta"
	this.ClassId = classId
	var objectType string = "openapi.ApiMethodMeta"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *OpenapiApiMethodMeta) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *OpenapiApiMethodMeta) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *OpenapiApiMethodMeta) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *OpenapiApiMethodMeta) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *OpenapiApiMethodMeta) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *OpenapiApiMethodMeta) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OpenapiApiMethodMeta) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenapiApiMethodMeta) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OpenapiApiMethodMeta) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OpenapiApiMethodMeta) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayLabel returns the DisplayLabel field value if set, zero value otherwise.
func (o *OpenapiApiMethodMeta) GetDisplayLabel() string {
	if o == nil || o.DisplayLabel == nil {
		var ret string
		return ret
	}
	return *o.DisplayLabel
}

// GetDisplayLabelOk returns a tuple with the DisplayLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenapiApiMethodMeta) GetDisplayLabelOk() (*string, bool) {
	if o == nil || o.DisplayLabel == nil {
		return nil, false
	}
	return o.DisplayLabel, true
}

// HasDisplayLabel returns a boolean if a field has been set.
func (o *OpenapiApiMethodMeta) HasDisplayLabel() bool {
	if o != nil && o.DisplayLabel != nil {
		return true
	}

	return false
}

// SetDisplayLabel gets a reference to the given string and assigns it to the DisplayLabel field.
func (o *OpenapiApiMethodMeta) SetDisplayLabel(v string) {
	o.DisplayLabel = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *OpenapiApiMethodMeta) GetMethod() string {
	if o == nil || o.Method == nil {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenapiApiMethodMeta) GetMethodOk() (*string, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *OpenapiApiMethodMeta) HasMethod() bool {
	if o != nil && o.Method != nil {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *OpenapiApiMethodMeta) SetMethod(v string) {
	o.Method = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OpenapiApiMethodMeta) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenapiApiMethodMeta) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OpenapiApiMethodMeta) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OpenapiApiMethodMeta) SetName(v string) {
	o.Name = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *OpenapiApiMethodMeta) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenapiApiMethodMeta) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *OpenapiApiMethodMeta) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *OpenapiApiMethodMeta) SetPath(v string) {
	o.Path = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *OpenapiApiMethodMeta) GetSource() OpenapiProcessFileRelationship {
	if o == nil || o.Source == nil {
		var ret OpenapiProcessFileRelationship
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenapiApiMethodMeta) GetSourceOk() (*OpenapiProcessFileRelationship, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *OpenapiApiMethodMeta) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given OpenapiProcessFileRelationship and assigns it to the Source field.
func (o *OpenapiApiMethodMeta) SetSource(v OpenapiProcessFileRelationship) {
	o.Source = &v
}

func (o OpenapiApiMethodMeta) MarshalJSON() ([]byte, error) {
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
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.DisplayLabel != nil {
		toSerialize["DisplayLabel"] = o.DisplayLabel
	}
	if o.Method != nil {
		toSerialize["Method"] = o.Method
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Path != nil {
		toSerialize["Path"] = o.Path
	}
	if o.Source != nil {
		toSerialize["Source"] = o.Source
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OpenapiApiMethodMeta) UnmarshalJSON(bytes []byte) (err error) {
	type OpenapiApiMethodMetaWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The description of the given API.
		Description *string `json:"Description,omitempty"`
		// The display label of the given API.
		DisplayLabel *string `json:"DisplayLabel,omitempty"`
		// The method type for the given API. * `GET` - Method type which indicates it is a GET API call. * `POST` - Method type which indicates it is a POST API call. * `PUT` - Method type which indicates it is a PUT API call. * `PATCH` - Method type which indicates it is a PATCH API call. * `DELETE` - Method type which indicates it is a DELETE API call.
		Method *string `json:"Method,omitempty"`
		// The description of the given API.
		Name *string `json:"Name,omitempty"`
		// Path of the selected API endpoint.
		Path   *string                         `json:"Path,omitempty"`
		Source *OpenapiProcessFileRelationship `json:"Source,omitempty"`
	}

	varOpenapiApiMethodMetaWithoutEmbeddedStruct := OpenapiApiMethodMetaWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varOpenapiApiMethodMetaWithoutEmbeddedStruct)
	if err == nil {
		varOpenapiApiMethodMeta := _OpenapiApiMethodMeta{}
		varOpenapiApiMethodMeta.ClassId = varOpenapiApiMethodMetaWithoutEmbeddedStruct.ClassId
		varOpenapiApiMethodMeta.ObjectType = varOpenapiApiMethodMetaWithoutEmbeddedStruct.ObjectType
		varOpenapiApiMethodMeta.Description = varOpenapiApiMethodMetaWithoutEmbeddedStruct.Description
		varOpenapiApiMethodMeta.DisplayLabel = varOpenapiApiMethodMetaWithoutEmbeddedStruct.DisplayLabel
		varOpenapiApiMethodMeta.Method = varOpenapiApiMethodMetaWithoutEmbeddedStruct.Method
		varOpenapiApiMethodMeta.Name = varOpenapiApiMethodMetaWithoutEmbeddedStruct.Name
		varOpenapiApiMethodMeta.Path = varOpenapiApiMethodMetaWithoutEmbeddedStruct.Path
		varOpenapiApiMethodMeta.Source = varOpenapiApiMethodMetaWithoutEmbeddedStruct.Source
		*o = OpenapiApiMethodMeta(varOpenapiApiMethodMeta)
	} else {
		return err
	}

	varOpenapiApiMethodMeta := _OpenapiApiMethodMeta{}

	err = json.Unmarshal(bytes, &varOpenapiApiMethodMeta)
	if err == nil {
		o.MoBaseMo = varOpenapiApiMethodMeta.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "DisplayLabel")
		delete(additionalProperties, "Method")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Path")
		delete(additionalProperties, "Source")

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

type NullableOpenapiApiMethodMeta struct {
	value *OpenapiApiMethodMeta
	isSet bool
}

func (v NullableOpenapiApiMethodMeta) Get() *OpenapiApiMethodMeta {
	return v.value
}

func (v *NullableOpenapiApiMethodMeta) Set(val *OpenapiApiMethodMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenapiApiMethodMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenapiApiMethodMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenapiApiMethodMeta(val *OpenapiApiMethodMeta) *NullableOpenapiApiMethodMeta {
	return &NullableOpenapiApiMethodMeta{value: val, isSet: true}
}

func (v NullableOpenapiApiMethodMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenapiApiMethodMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
