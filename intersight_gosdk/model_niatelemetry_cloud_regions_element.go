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
	"reflect"
	"strings"
)

// NiatelemetryCloudRegionsElement Stores information related to cloud region.
type NiatelemetryCloudRegionsElement struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Return value of adminState attribute.
	AdminState *string `json:"AdminState,omitempty"`
	// Return whether CAPIC is deployed in the cloud region or not.
	CapicDeployed *string `json:"CapicDeployed,omitempty"`
	// Return whether any user deployment is configured in the cloud region or not.
	InUse *string `json:"InUse,omitempty"`
	// Return value of name of the cloud region.
	Name                 *string `json:"Name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryCloudRegionsElement NiatelemetryCloudRegionsElement

// NewNiatelemetryCloudRegionsElement instantiates a new NiatelemetryCloudRegionsElement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryCloudRegionsElement(classId string, objectType string) *NiatelemetryCloudRegionsElement {
	this := NiatelemetryCloudRegionsElement{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryCloudRegionsElementWithDefaults instantiates a new NiatelemetryCloudRegionsElement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryCloudRegionsElementWithDefaults() *NiatelemetryCloudRegionsElement {
	this := NiatelemetryCloudRegionsElement{}
	var classId string = "niatelemetry.CloudRegionsElement"
	this.ClassId = classId
	var objectType string = "niatelemetry.CloudRegionsElement"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryCloudRegionsElement) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryCloudRegionsElement) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryCloudRegionsElement) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryCloudRegionsElement) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryCloudRegionsElement) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryCloudRegionsElement) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAdminState returns the AdminState field value if set, zero value otherwise.
func (o *NiatelemetryCloudRegionsElement) GetAdminState() string {
	if o == nil || o.AdminState == nil {
		var ret string
		return ret
	}
	return *o.AdminState
}

// GetAdminStateOk returns a tuple with the AdminState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryCloudRegionsElement) GetAdminStateOk() (*string, bool) {
	if o == nil || o.AdminState == nil {
		return nil, false
	}
	return o.AdminState, true
}

// HasAdminState returns a boolean if a field has been set.
func (o *NiatelemetryCloudRegionsElement) HasAdminState() bool {
	if o != nil && o.AdminState != nil {
		return true
	}

	return false
}

// SetAdminState gets a reference to the given string and assigns it to the AdminState field.
func (o *NiatelemetryCloudRegionsElement) SetAdminState(v string) {
	o.AdminState = &v
}

// GetCapicDeployed returns the CapicDeployed field value if set, zero value otherwise.
func (o *NiatelemetryCloudRegionsElement) GetCapicDeployed() string {
	if o == nil || o.CapicDeployed == nil {
		var ret string
		return ret
	}
	return *o.CapicDeployed
}

// GetCapicDeployedOk returns a tuple with the CapicDeployed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryCloudRegionsElement) GetCapicDeployedOk() (*string, bool) {
	if o == nil || o.CapicDeployed == nil {
		return nil, false
	}
	return o.CapicDeployed, true
}

// HasCapicDeployed returns a boolean if a field has been set.
func (o *NiatelemetryCloudRegionsElement) HasCapicDeployed() bool {
	if o != nil && o.CapicDeployed != nil {
		return true
	}

	return false
}

// SetCapicDeployed gets a reference to the given string and assigns it to the CapicDeployed field.
func (o *NiatelemetryCloudRegionsElement) SetCapicDeployed(v string) {
	o.CapicDeployed = &v
}

// GetInUse returns the InUse field value if set, zero value otherwise.
func (o *NiatelemetryCloudRegionsElement) GetInUse() string {
	if o == nil || o.InUse == nil {
		var ret string
		return ret
	}
	return *o.InUse
}

// GetInUseOk returns a tuple with the InUse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryCloudRegionsElement) GetInUseOk() (*string, bool) {
	if o == nil || o.InUse == nil {
		return nil, false
	}
	return o.InUse, true
}

// HasInUse returns a boolean if a field has been set.
func (o *NiatelemetryCloudRegionsElement) HasInUse() bool {
	if o != nil && o.InUse != nil {
		return true
	}

	return false
}

// SetInUse gets a reference to the given string and assigns it to the InUse field.
func (o *NiatelemetryCloudRegionsElement) SetInUse(v string) {
	o.InUse = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NiatelemetryCloudRegionsElement) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryCloudRegionsElement) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NiatelemetryCloudRegionsElement) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NiatelemetryCloudRegionsElement) SetName(v string) {
	o.Name = &v
}

func (o NiatelemetryCloudRegionsElement) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AdminState != nil {
		toSerialize["AdminState"] = o.AdminState
	}
	if o.CapicDeployed != nil {
		toSerialize["CapicDeployed"] = o.CapicDeployed
	}
	if o.InUse != nil {
		toSerialize["InUse"] = o.InUse
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryCloudRegionsElement) UnmarshalJSON(bytes []byte) (err error) {
	type NiatelemetryCloudRegionsElementWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Return value of adminState attribute.
		AdminState *string `json:"AdminState,omitempty"`
		// Return whether CAPIC is deployed in the cloud region or not.
		CapicDeployed *string `json:"CapicDeployed,omitempty"`
		// Return whether any user deployment is configured in the cloud region or not.
		InUse *string `json:"InUse,omitempty"`
		// Return value of name of the cloud region.
		Name *string `json:"Name,omitempty"`
	}

	varNiatelemetryCloudRegionsElementWithoutEmbeddedStruct := NiatelemetryCloudRegionsElementWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNiatelemetryCloudRegionsElementWithoutEmbeddedStruct)
	if err == nil {
		varNiatelemetryCloudRegionsElement := _NiatelemetryCloudRegionsElement{}
		varNiatelemetryCloudRegionsElement.ClassId = varNiatelemetryCloudRegionsElementWithoutEmbeddedStruct.ClassId
		varNiatelemetryCloudRegionsElement.ObjectType = varNiatelemetryCloudRegionsElementWithoutEmbeddedStruct.ObjectType
		varNiatelemetryCloudRegionsElement.AdminState = varNiatelemetryCloudRegionsElementWithoutEmbeddedStruct.AdminState
		varNiatelemetryCloudRegionsElement.CapicDeployed = varNiatelemetryCloudRegionsElementWithoutEmbeddedStruct.CapicDeployed
		varNiatelemetryCloudRegionsElement.InUse = varNiatelemetryCloudRegionsElementWithoutEmbeddedStruct.InUse
		varNiatelemetryCloudRegionsElement.Name = varNiatelemetryCloudRegionsElementWithoutEmbeddedStruct.Name
		*o = NiatelemetryCloudRegionsElement(varNiatelemetryCloudRegionsElement)
	} else {
		return err
	}

	varNiatelemetryCloudRegionsElement := _NiatelemetryCloudRegionsElement{}

	err = json.Unmarshal(bytes, &varNiatelemetryCloudRegionsElement)
	if err == nil {
		o.MoBaseComplexType = varNiatelemetryCloudRegionsElement.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AdminState")
		delete(additionalProperties, "CapicDeployed")
		delete(additionalProperties, "InUse")
		delete(additionalProperties, "Name")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

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

type NullableNiatelemetryCloudRegionsElement struct {
	value *NiatelemetryCloudRegionsElement
	isSet bool
}

func (v NullableNiatelemetryCloudRegionsElement) Get() *NiatelemetryCloudRegionsElement {
	return v.value
}

func (v *NullableNiatelemetryCloudRegionsElement) Set(val *NiatelemetryCloudRegionsElement) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryCloudRegionsElement) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryCloudRegionsElement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryCloudRegionsElement(val *NiatelemetryCloudRegionsElement) *NullableNiatelemetryCloudRegionsElement {
	return &NullableNiatelemetryCloudRegionsElement{value: val, isSet: true}
}

func (v NullableNiatelemetryCloudRegionsElement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryCloudRegionsElement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
