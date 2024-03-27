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

// PartnerintegrationDcLogs Logs from the build operation.
type PartnerintegrationDcLogs struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Stage in the build process these logs belong to. * `None` - Default value for the log stage. * `Backend` - Logs corresponding to backend build. * `Ui` - Logs corresponding to ui build stage. * `Apidocs` - Logs corresponding to the apidocs build stage.
	Stage                *string                                        `json:"Stage,omitempty"`
	Stderr               []string                                       `json:"Stderr,omitempty"`
	Stdout               []string                                       `json:"Stdout,omitempty"`
	DeviceConnector      *PartnerintegrationDeviceConnectorRelationship `json:"DeviceConnector,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PartnerintegrationDcLogs PartnerintegrationDcLogs

// NewPartnerintegrationDcLogs instantiates a new PartnerintegrationDcLogs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartnerintegrationDcLogs(classId string, objectType string) *PartnerintegrationDcLogs {
	this := PartnerintegrationDcLogs{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewPartnerintegrationDcLogsWithDefaults instantiates a new PartnerintegrationDcLogs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartnerintegrationDcLogsWithDefaults() *PartnerintegrationDcLogs {
	this := PartnerintegrationDcLogs{}
	var classId string = "partnerintegration.DcLogs"
	this.ClassId = classId
	var objectType string = "partnerintegration.DcLogs"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *PartnerintegrationDcLogs) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *PartnerintegrationDcLogs) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *PartnerintegrationDcLogs) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *PartnerintegrationDcLogs) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *PartnerintegrationDcLogs) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *PartnerintegrationDcLogs) SetObjectType(v string) {
	o.ObjectType = v
}

// GetStage returns the Stage field value if set, zero value otherwise.
func (o *PartnerintegrationDcLogs) GetStage() string {
	if o == nil || o.Stage == nil {
		var ret string
		return ret
	}
	return *o.Stage
}

// GetStageOk returns a tuple with the Stage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerintegrationDcLogs) GetStageOk() (*string, bool) {
	if o == nil || o.Stage == nil {
		return nil, false
	}
	return o.Stage, true
}

// HasStage returns a boolean if a field has been set.
func (o *PartnerintegrationDcLogs) HasStage() bool {
	if o != nil && o.Stage != nil {
		return true
	}

	return false
}

// SetStage gets a reference to the given string and assigns it to the Stage field.
func (o *PartnerintegrationDcLogs) SetStage(v string) {
	o.Stage = &v
}

// GetStderr returns the Stderr field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PartnerintegrationDcLogs) GetStderr() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Stderr
}

// GetStderrOk returns a tuple with the Stderr field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartnerintegrationDcLogs) GetStderrOk() ([]string, bool) {
	if o == nil || o.Stderr == nil {
		return nil, false
	}
	return o.Stderr, true
}

// HasStderr returns a boolean if a field has been set.
func (o *PartnerintegrationDcLogs) HasStderr() bool {
	if o != nil && o.Stderr != nil {
		return true
	}

	return false
}

// SetStderr gets a reference to the given []string and assigns it to the Stderr field.
func (o *PartnerintegrationDcLogs) SetStderr(v []string) {
	o.Stderr = v
}

// GetStdout returns the Stdout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PartnerintegrationDcLogs) GetStdout() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Stdout
}

// GetStdoutOk returns a tuple with the Stdout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartnerintegrationDcLogs) GetStdoutOk() ([]string, bool) {
	if o == nil || o.Stdout == nil {
		return nil, false
	}
	return o.Stdout, true
}

// HasStdout returns a boolean if a field has been set.
func (o *PartnerintegrationDcLogs) HasStdout() bool {
	if o != nil && o.Stdout != nil {
		return true
	}

	return false
}

// SetStdout gets a reference to the given []string and assigns it to the Stdout field.
func (o *PartnerintegrationDcLogs) SetStdout(v []string) {
	o.Stdout = v
}

// GetDeviceConnector returns the DeviceConnector field value if set, zero value otherwise.
func (o *PartnerintegrationDcLogs) GetDeviceConnector() PartnerintegrationDeviceConnectorRelationship {
	if o == nil || o.DeviceConnector == nil {
		var ret PartnerintegrationDeviceConnectorRelationship
		return ret
	}
	return *o.DeviceConnector
}

// GetDeviceConnectorOk returns a tuple with the DeviceConnector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerintegrationDcLogs) GetDeviceConnectorOk() (*PartnerintegrationDeviceConnectorRelationship, bool) {
	if o == nil || o.DeviceConnector == nil {
		return nil, false
	}
	return o.DeviceConnector, true
}

// HasDeviceConnector returns a boolean if a field has been set.
func (o *PartnerintegrationDcLogs) HasDeviceConnector() bool {
	if o != nil && o.DeviceConnector != nil {
		return true
	}

	return false
}

// SetDeviceConnector gets a reference to the given PartnerintegrationDeviceConnectorRelationship and assigns it to the DeviceConnector field.
func (o *PartnerintegrationDcLogs) SetDeviceConnector(v PartnerintegrationDeviceConnectorRelationship) {
	o.DeviceConnector = &v
}

func (o PartnerintegrationDcLogs) MarshalJSON() ([]byte, error) {
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
	if o.Stage != nil {
		toSerialize["Stage"] = o.Stage
	}
	if o.Stderr != nil {
		toSerialize["Stderr"] = o.Stderr
	}
	if o.Stdout != nil {
		toSerialize["Stdout"] = o.Stdout
	}
	if o.DeviceConnector != nil {
		toSerialize["DeviceConnector"] = o.DeviceConnector
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PartnerintegrationDcLogs) UnmarshalJSON(bytes []byte) (err error) {
	type PartnerintegrationDcLogsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Stage in the build process these logs belong to. * `None` - Default value for the log stage. * `Backend` - Logs corresponding to backend build. * `Ui` - Logs corresponding to ui build stage. * `Apidocs` - Logs corresponding to the apidocs build stage.
		Stage           *string                                        `json:"Stage,omitempty"`
		Stderr          []string                                       `json:"Stderr,omitempty"`
		Stdout          []string                                       `json:"Stdout,omitempty"`
		DeviceConnector *PartnerintegrationDeviceConnectorRelationship `json:"DeviceConnector,omitempty"`
	}

	varPartnerintegrationDcLogsWithoutEmbeddedStruct := PartnerintegrationDcLogsWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varPartnerintegrationDcLogsWithoutEmbeddedStruct)
	if err == nil {
		varPartnerintegrationDcLogs := _PartnerintegrationDcLogs{}
		varPartnerintegrationDcLogs.ClassId = varPartnerintegrationDcLogsWithoutEmbeddedStruct.ClassId
		varPartnerintegrationDcLogs.ObjectType = varPartnerintegrationDcLogsWithoutEmbeddedStruct.ObjectType
		varPartnerintegrationDcLogs.Stage = varPartnerintegrationDcLogsWithoutEmbeddedStruct.Stage
		varPartnerintegrationDcLogs.Stderr = varPartnerintegrationDcLogsWithoutEmbeddedStruct.Stderr
		varPartnerintegrationDcLogs.Stdout = varPartnerintegrationDcLogsWithoutEmbeddedStruct.Stdout
		varPartnerintegrationDcLogs.DeviceConnector = varPartnerintegrationDcLogsWithoutEmbeddedStruct.DeviceConnector
		*o = PartnerintegrationDcLogs(varPartnerintegrationDcLogs)
	} else {
		return err
	}

	varPartnerintegrationDcLogs := _PartnerintegrationDcLogs{}

	err = json.Unmarshal(bytes, &varPartnerintegrationDcLogs)
	if err == nil {
		o.MoBaseMo = varPartnerintegrationDcLogs.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Stage")
		delete(additionalProperties, "Stderr")
		delete(additionalProperties, "Stdout")
		delete(additionalProperties, "DeviceConnector")

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

type NullablePartnerintegrationDcLogs struct {
	value *PartnerintegrationDcLogs
	isSet bool
}

func (v NullablePartnerintegrationDcLogs) Get() *PartnerintegrationDcLogs {
	return v.value
}

func (v *NullablePartnerintegrationDcLogs) Set(val *PartnerintegrationDcLogs) {
	v.value = val
	v.isSet = true
}

func (v NullablePartnerintegrationDcLogs) IsSet() bool {
	return v.isSet
}

func (v *NullablePartnerintegrationDcLogs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartnerintegrationDcLogs(val *PartnerintegrationDcLogs) *NullablePartnerintegrationDcLogs {
	return &NullablePartnerintegrationDcLogs{value: val, isSet: true}
}

func (v NullablePartnerintegrationDcLogs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartnerintegrationDcLogs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
