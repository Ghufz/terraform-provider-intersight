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

// FirmwareUpgradeValidity Upgrade Validity API that performs basic validity checks for performing a firmware upgrade on the endpoint. The API checks whether the endpoint satisfies the basic platform requirements that are needed for firmware upgrade to happen. Support is currently available only for Standalone servers.
type FirmwareUpgradeValidity struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The error string returned while checking for a target device's validity for firmware upgrade.
	ErrorMessage *string `json:"ErrorMessage,omitempty"`
	// This flag denotes whether the target device is a valid target for firmware upgrade.
	IsValid              *bool                        `json:"IsValid,omitempty"`
	Server               *ComputePhysicalRelationship `json:"Server,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FirmwareUpgradeValidity FirmwareUpgradeValidity

// NewFirmwareUpgradeValidity instantiates a new FirmwareUpgradeValidity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirmwareUpgradeValidity(classId string, objectType string) *FirmwareUpgradeValidity {
	this := FirmwareUpgradeValidity{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewFirmwareUpgradeValidityWithDefaults instantiates a new FirmwareUpgradeValidity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirmwareUpgradeValidityWithDefaults() *FirmwareUpgradeValidity {
	this := FirmwareUpgradeValidity{}
	var classId string = "firmware.UpgradeValidity"
	this.ClassId = classId
	var objectType string = "firmware.UpgradeValidity"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *FirmwareUpgradeValidity) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FirmwareUpgradeValidity) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FirmwareUpgradeValidity) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FirmwareUpgradeValidity) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FirmwareUpgradeValidity) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FirmwareUpgradeValidity) SetObjectType(v string) {
	o.ObjectType = v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *FirmwareUpgradeValidity) GetErrorMessage() string {
	if o == nil || o.ErrorMessage == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareUpgradeValidity) GetErrorMessageOk() (*string, bool) {
	if o == nil || o.ErrorMessage == nil {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *FirmwareUpgradeValidity) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage != nil {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *FirmwareUpgradeValidity) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetIsValid returns the IsValid field value if set, zero value otherwise.
func (o *FirmwareUpgradeValidity) GetIsValid() bool {
	if o == nil || o.IsValid == nil {
		var ret bool
		return ret
	}
	return *o.IsValid
}

// GetIsValidOk returns a tuple with the IsValid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareUpgradeValidity) GetIsValidOk() (*bool, bool) {
	if o == nil || o.IsValid == nil {
		return nil, false
	}
	return o.IsValid, true
}

// HasIsValid returns a boolean if a field has been set.
func (o *FirmwareUpgradeValidity) HasIsValid() bool {
	if o != nil && o.IsValid != nil {
		return true
	}

	return false
}

// SetIsValid gets a reference to the given bool and assigns it to the IsValid field.
func (o *FirmwareUpgradeValidity) SetIsValid(v bool) {
	o.IsValid = &v
}

// GetServer returns the Server field value if set, zero value otherwise.
func (o *FirmwareUpgradeValidity) GetServer() ComputePhysicalRelationship {
	if o == nil || o.Server == nil {
		var ret ComputePhysicalRelationship
		return ret
	}
	return *o.Server
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareUpgradeValidity) GetServerOk() (*ComputePhysicalRelationship, bool) {
	if o == nil || o.Server == nil {
		return nil, false
	}
	return o.Server, true
}

// HasServer returns a boolean if a field has been set.
func (o *FirmwareUpgradeValidity) HasServer() bool {
	if o != nil && o.Server != nil {
		return true
	}

	return false
}

// SetServer gets a reference to the given ComputePhysicalRelationship and assigns it to the Server field.
func (o *FirmwareUpgradeValidity) SetServer(v ComputePhysicalRelationship) {
	o.Server = &v
}

func (o FirmwareUpgradeValidity) MarshalJSON() ([]byte, error) {
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
	if o.ErrorMessage != nil {
		toSerialize["ErrorMessage"] = o.ErrorMessage
	}
	if o.IsValid != nil {
		toSerialize["IsValid"] = o.IsValid
	}
	if o.Server != nil {
		toSerialize["Server"] = o.Server
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FirmwareUpgradeValidity) UnmarshalJSON(bytes []byte) (err error) {
	type FirmwareUpgradeValidityWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The error string returned while checking for a target device's validity for firmware upgrade.
		ErrorMessage *string `json:"ErrorMessage,omitempty"`
		// This flag denotes whether the target device is a valid target for firmware upgrade.
		IsValid *bool                        `json:"IsValid,omitempty"`
		Server  *ComputePhysicalRelationship `json:"Server,omitempty"`
	}

	varFirmwareUpgradeValidityWithoutEmbeddedStruct := FirmwareUpgradeValidityWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varFirmwareUpgradeValidityWithoutEmbeddedStruct)
	if err == nil {
		varFirmwareUpgradeValidity := _FirmwareUpgradeValidity{}
		varFirmwareUpgradeValidity.ClassId = varFirmwareUpgradeValidityWithoutEmbeddedStruct.ClassId
		varFirmwareUpgradeValidity.ObjectType = varFirmwareUpgradeValidityWithoutEmbeddedStruct.ObjectType
		varFirmwareUpgradeValidity.ErrorMessage = varFirmwareUpgradeValidityWithoutEmbeddedStruct.ErrorMessage
		varFirmwareUpgradeValidity.IsValid = varFirmwareUpgradeValidityWithoutEmbeddedStruct.IsValid
		varFirmwareUpgradeValidity.Server = varFirmwareUpgradeValidityWithoutEmbeddedStruct.Server
		*o = FirmwareUpgradeValidity(varFirmwareUpgradeValidity)
	} else {
		return err
	}

	varFirmwareUpgradeValidity := _FirmwareUpgradeValidity{}

	err = json.Unmarshal(bytes, &varFirmwareUpgradeValidity)
	if err == nil {
		o.MoBaseMo = varFirmwareUpgradeValidity.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ErrorMessage")
		delete(additionalProperties, "IsValid")
		delete(additionalProperties, "Server")

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

type NullableFirmwareUpgradeValidity struct {
	value *FirmwareUpgradeValidity
	isSet bool
}

func (v NullableFirmwareUpgradeValidity) Get() *FirmwareUpgradeValidity {
	return v.value
}

func (v *NullableFirmwareUpgradeValidity) Set(val *FirmwareUpgradeValidity) {
	v.value = val
	v.isSet = true
}

func (v NullableFirmwareUpgradeValidity) IsSet() bool {
	return v.isSet
}

func (v *NullableFirmwareUpgradeValidity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirmwareUpgradeValidity(val *FirmwareUpgradeValidity) *NullableFirmwareUpgradeValidity {
	return &NullableFirmwareUpgradeValidity{value: val, isSet: true}
}

func (v NullableFirmwareUpgradeValidity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirmwareUpgradeValidity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
