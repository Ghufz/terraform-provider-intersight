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

// WorkflowPowerShellApi This models a single PowerShell script execution that can be sent to a claimed PowerShell target.
type WorkflowPowerShellApi struct {
	WorkflowApi
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The response of a PowerShell script is an object, since PowerShell is an Object based language. Each object can contain multiple objects as properties, each of which in turn can contain multiple objects and so on and so forth. The depth field specifies how many levels of contained objects are included in the JSON representation.
	Depth *int64 `json:"Depth,omitempty"`
	// The timeout in seconds for the execution of the script against the given endpoint.
	OperationTimeout *string `json:"OperationTimeout,omitempty"`
	// The grammar specification to parse the response and extract the required values.
	PowerShellResponseSpec interface{} `json:"PowerShellResponseSpec,omitempty"`
	AdditionalProperties   map[string]interface{}
}

type _WorkflowPowerShellApi WorkflowPowerShellApi

// NewWorkflowPowerShellApi instantiates a new WorkflowPowerShellApi object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowPowerShellApi(classId string, objectType string) *WorkflowPowerShellApi {
	this := WorkflowPowerShellApi{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowPowerShellApiWithDefaults instantiates a new WorkflowPowerShellApi object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowPowerShellApiWithDefaults() *WorkflowPowerShellApi {
	this := WorkflowPowerShellApi{}
	var classId string = "workflow.PowerShellApi"
	this.ClassId = classId
	var objectType string = "workflow.PowerShellApi"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowPowerShellApi) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowPowerShellApi) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowPowerShellApi) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowPowerShellApi) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowPowerShellApi) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowPowerShellApi) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDepth returns the Depth field value if set, zero value otherwise.
func (o *WorkflowPowerShellApi) GetDepth() int64 {
	if o == nil || o.Depth == nil {
		var ret int64
		return ret
	}
	return *o.Depth
}

// GetDepthOk returns a tuple with the Depth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowPowerShellApi) GetDepthOk() (*int64, bool) {
	if o == nil || o.Depth == nil {
		return nil, false
	}
	return o.Depth, true
}

// HasDepth returns a boolean if a field has been set.
func (o *WorkflowPowerShellApi) HasDepth() bool {
	if o != nil && o.Depth != nil {
		return true
	}

	return false
}

// SetDepth gets a reference to the given int64 and assigns it to the Depth field.
func (o *WorkflowPowerShellApi) SetDepth(v int64) {
	o.Depth = &v
}

// GetOperationTimeout returns the OperationTimeout field value if set, zero value otherwise.
func (o *WorkflowPowerShellApi) GetOperationTimeout() string {
	if o == nil || o.OperationTimeout == nil {
		var ret string
		return ret
	}
	return *o.OperationTimeout
}

// GetOperationTimeoutOk returns a tuple with the OperationTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowPowerShellApi) GetOperationTimeoutOk() (*string, bool) {
	if o == nil || o.OperationTimeout == nil {
		return nil, false
	}
	return o.OperationTimeout, true
}

// HasOperationTimeout returns a boolean if a field has been set.
func (o *WorkflowPowerShellApi) HasOperationTimeout() bool {
	if o != nil && o.OperationTimeout != nil {
		return true
	}

	return false
}

// SetOperationTimeout gets a reference to the given string and assigns it to the OperationTimeout field.
func (o *WorkflowPowerShellApi) SetOperationTimeout(v string) {
	o.OperationTimeout = &v
}

// GetPowerShellResponseSpec returns the PowerShellResponseSpec field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowPowerShellApi) GetPowerShellResponseSpec() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.PowerShellResponseSpec
}

// GetPowerShellResponseSpecOk returns a tuple with the PowerShellResponseSpec field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowPowerShellApi) GetPowerShellResponseSpecOk() (*interface{}, bool) {
	if o == nil || o.PowerShellResponseSpec == nil {
		return nil, false
	}
	return &o.PowerShellResponseSpec, true
}

// HasPowerShellResponseSpec returns a boolean if a field has been set.
func (o *WorkflowPowerShellApi) HasPowerShellResponseSpec() bool {
	if o != nil && o.PowerShellResponseSpec != nil {
		return true
	}

	return false
}

// SetPowerShellResponseSpec gets a reference to the given interface{} and assigns it to the PowerShellResponseSpec field.
func (o *WorkflowPowerShellApi) SetPowerShellResponseSpec(v interface{}) {
	o.PowerShellResponseSpec = v
}

func (o WorkflowPowerShellApi) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedWorkflowApi, errWorkflowApi := json.Marshal(o.WorkflowApi)
	if errWorkflowApi != nil {
		return []byte{}, errWorkflowApi
	}
	errWorkflowApi = json.Unmarshal([]byte(serializedWorkflowApi), &toSerialize)
	if errWorkflowApi != nil {
		return []byte{}, errWorkflowApi
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Depth != nil {
		toSerialize["Depth"] = o.Depth
	}
	if o.OperationTimeout != nil {
		toSerialize["OperationTimeout"] = o.OperationTimeout
	}
	if o.PowerShellResponseSpec != nil {
		toSerialize["PowerShellResponseSpec"] = o.PowerShellResponseSpec
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowPowerShellApi) UnmarshalJSON(bytes []byte) (err error) {
	type WorkflowPowerShellApiWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The response of a PowerShell script is an object, since PowerShell is an Object based language. Each object can contain multiple objects as properties, each of which in turn can contain multiple objects and so on and so forth. The depth field specifies how many levels of contained objects are included in the JSON representation.
		Depth *int64 `json:"Depth,omitempty"`
		// The timeout in seconds for the execution of the script against the given endpoint.
		OperationTimeout *string `json:"OperationTimeout,omitempty"`
		// The grammar specification to parse the response and extract the required values.
		PowerShellResponseSpec interface{} `json:"PowerShellResponseSpec,omitempty"`
	}

	varWorkflowPowerShellApiWithoutEmbeddedStruct := WorkflowPowerShellApiWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varWorkflowPowerShellApiWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowPowerShellApi := _WorkflowPowerShellApi{}
		varWorkflowPowerShellApi.ClassId = varWorkflowPowerShellApiWithoutEmbeddedStruct.ClassId
		varWorkflowPowerShellApi.ObjectType = varWorkflowPowerShellApiWithoutEmbeddedStruct.ObjectType
		varWorkflowPowerShellApi.Depth = varWorkflowPowerShellApiWithoutEmbeddedStruct.Depth
		varWorkflowPowerShellApi.OperationTimeout = varWorkflowPowerShellApiWithoutEmbeddedStruct.OperationTimeout
		varWorkflowPowerShellApi.PowerShellResponseSpec = varWorkflowPowerShellApiWithoutEmbeddedStruct.PowerShellResponseSpec
		*o = WorkflowPowerShellApi(varWorkflowPowerShellApi)
	} else {
		return err
	}

	varWorkflowPowerShellApi := _WorkflowPowerShellApi{}

	err = json.Unmarshal(bytes, &varWorkflowPowerShellApi)
	if err == nil {
		o.WorkflowApi = varWorkflowPowerShellApi.WorkflowApi
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Depth")
		delete(additionalProperties, "OperationTimeout")
		delete(additionalProperties, "PowerShellResponseSpec")

		// remove fields from embedded structs
		reflectWorkflowApi := reflect.ValueOf(o.WorkflowApi)
		for i := 0; i < reflectWorkflowApi.Type().NumField(); i++ {
			t := reflectWorkflowApi.Type().Field(i)

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

type NullableWorkflowPowerShellApi struct {
	value *WorkflowPowerShellApi
	isSet bool
}

func (v NullableWorkflowPowerShellApi) Get() *WorkflowPowerShellApi {
	return v.value
}

func (v *NullableWorkflowPowerShellApi) Set(val *WorkflowPowerShellApi) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowPowerShellApi) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowPowerShellApi) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowPowerShellApi(val *WorkflowPowerShellApi) *NullableWorkflowPowerShellApi {
	return &NullableWorkflowPowerShellApi{value: val, isSet: true}
}

func (v NullableWorkflowPowerShellApi) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowPowerShellApi) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
