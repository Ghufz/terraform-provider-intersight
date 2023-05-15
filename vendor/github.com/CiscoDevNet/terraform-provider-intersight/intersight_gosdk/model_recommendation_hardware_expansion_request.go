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

// RecommendationHardwareExpansionRequest Entity representing the user request for HyperFlex cluster expansion.
type RecommendationHardwareExpansionRequest struct {
	RecommendationExpansionRequest
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Action to be triggered for computing recommendation. Default value is None. * `None` - The Enum value None represents that no action is triggered on the forecast Instance managed object. * `Evaluate` - The Enum value Evaluate represents that a re-evaluation of the forecast needs to be triggered.
	Action *string `json:"Action,omitempty"`
	// User visible error message for the Hardware Expansion request.
	Message *string `json:"Message,omitempty"`
	// Status of the Hardware Expansion request. * `None` - The Enum value None represents the default status value before any API call is made. * `Success` - The Enum value Success represents that the API call returned with success. * `Fail` - The Enum value Fail represents that the API call returned with a failure.
	Status *string `json:"Status,omitempty"`
	// An array of relationships to recommendationHardwareExpansionRequestItem resources.
	HwExpansionRequestItems []RecommendationHardwareExpansionRequestItemRelationship `json:"HwExpansionRequestItems,omitempty"`
	RegisteredDevice        *AssetDeviceRegistrationRelationship                     `json:"RegisteredDevice,omitempty"`
	AdditionalProperties    map[string]interface{}
}

type _RecommendationHardwareExpansionRequest RecommendationHardwareExpansionRequest

// NewRecommendationHardwareExpansionRequest instantiates a new RecommendationHardwareExpansionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecommendationHardwareExpansionRequest(classId string, objectType string) *RecommendationHardwareExpansionRequest {
	this := RecommendationHardwareExpansionRequest{}
	this.ClassId = classId
	this.ObjectType = objectType
	var action string = "None"
	this.Action = &action
	return &this
}

// NewRecommendationHardwareExpansionRequestWithDefaults instantiates a new RecommendationHardwareExpansionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecommendationHardwareExpansionRequestWithDefaults() *RecommendationHardwareExpansionRequest {
	this := RecommendationHardwareExpansionRequest{}
	var classId string = "recommendation.HardwareExpansionRequest"
	this.ClassId = classId
	var objectType string = "recommendation.HardwareExpansionRequest"
	this.ObjectType = objectType
	var action string = "None"
	this.Action = &action
	return &this
}

// GetClassId returns the ClassId field value
func (o *RecommendationHardwareExpansionRequest) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *RecommendationHardwareExpansionRequest) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *RecommendationHardwareExpansionRequest) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *RecommendationHardwareExpansionRequest) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *RecommendationHardwareExpansionRequest) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *RecommendationHardwareExpansionRequest) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *RecommendationHardwareExpansionRequest) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationHardwareExpansionRequest) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *RecommendationHardwareExpansionRequest) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *RecommendationHardwareExpansionRequest) SetAction(v string) {
	o.Action = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *RecommendationHardwareExpansionRequest) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationHardwareExpansionRequest) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *RecommendationHardwareExpansionRequest) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *RecommendationHardwareExpansionRequest) SetMessage(v string) {
	o.Message = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *RecommendationHardwareExpansionRequest) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationHardwareExpansionRequest) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *RecommendationHardwareExpansionRequest) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *RecommendationHardwareExpansionRequest) SetStatus(v string) {
	o.Status = &v
}

// GetHwExpansionRequestItems returns the HwExpansionRequestItems field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecommendationHardwareExpansionRequest) GetHwExpansionRequestItems() []RecommendationHardwareExpansionRequestItemRelationship {
	if o == nil {
		var ret []RecommendationHardwareExpansionRequestItemRelationship
		return ret
	}
	return o.HwExpansionRequestItems
}

// GetHwExpansionRequestItemsOk returns a tuple with the HwExpansionRequestItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecommendationHardwareExpansionRequest) GetHwExpansionRequestItemsOk() ([]RecommendationHardwareExpansionRequestItemRelationship, bool) {
	if o == nil || o.HwExpansionRequestItems == nil {
		return nil, false
	}
	return o.HwExpansionRequestItems, true
}

// HasHwExpansionRequestItems returns a boolean if a field has been set.
func (o *RecommendationHardwareExpansionRequest) HasHwExpansionRequestItems() bool {
	if o != nil && o.HwExpansionRequestItems != nil {
		return true
	}

	return false
}

// SetHwExpansionRequestItems gets a reference to the given []RecommendationHardwareExpansionRequestItemRelationship and assigns it to the HwExpansionRequestItems field.
func (o *RecommendationHardwareExpansionRequest) SetHwExpansionRequestItems(v []RecommendationHardwareExpansionRequestItemRelationship) {
	o.HwExpansionRequestItems = v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *RecommendationHardwareExpansionRequest) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationHardwareExpansionRequest) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *RecommendationHardwareExpansionRequest) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *RecommendationHardwareExpansionRequest) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o RecommendationHardwareExpansionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedRecommendationExpansionRequest, errRecommendationExpansionRequest := json.Marshal(o.RecommendationExpansionRequest)
	if errRecommendationExpansionRequest != nil {
		return []byte{}, errRecommendationExpansionRequest
	}
	errRecommendationExpansionRequest = json.Unmarshal([]byte(serializedRecommendationExpansionRequest), &toSerialize)
	if errRecommendationExpansionRequest != nil {
		return []byte{}, errRecommendationExpansionRequest
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Action != nil {
		toSerialize["Action"] = o.Action
	}
	if o.Message != nil {
		toSerialize["Message"] = o.Message
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.HwExpansionRequestItems != nil {
		toSerialize["HwExpansionRequestItems"] = o.HwExpansionRequestItems
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RecommendationHardwareExpansionRequest) UnmarshalJSON(bytes []byte) (err error) {
	type RecommendationHardwareExpansionRequestWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Action to be triggered for computing recommendation. Default value is None. * `None` - The Enum value None represents that no action is triggered on the forecast Instance managed object. * `Evaluate` - The Enum value Evaluate represents that a re-evaluation of the forecast needs to be triggered.
		Action *string `json:"Action,omitempty"`
		// User visible error message for the Hardware Expansion request.
		Message *string `json:"Message,omitempty"`
		// Status of the Hardware Expansion request. * `None` - The Enum value None represents the default status value before any API call is made. * `Success` - The Enum value Success represents that the API call returned with success. * `Fail` - The Enum value Fail represents that the API call returned with a failure.
		Status *string `json:"Status,omitempty"`
		// An array of relationships to recommendationHardwareExpansionRequestItem resources.
		HwExpansionRequestItems []RecommendationHardwareExpansionRequestItemRelationship `json:"HwExpansionRequestItems,omitempty"`
		RegisteredDevice        *AssetDeviceRegistrationRelationship                     `json:"RegisteredDevice,omitempty"`
	}

	varRecommendationHardwareExpansionRequestWithoutEmbeddedStruct := RecommendationHardwareExpansionRequestWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varRecommendationHardwareExpansionRequestWithoutEmbeddedStruct)
	if err == nil {
		varRecommendationHardwareExpansionRequest := _RecommendationHardwareExpansionRequest{}
		varRecommendationHardwareExpansionRequest.ClassId = varRecommendationHardwareExpansionRequestWithoutEmbeddedStruct.ClassId
		varRecommendationHardwareExpansionRequest.ObjectType = varRecommendationHardwareExpansionRequestWithoutEmbeddedStruct.ObjectType
		varRecommendationHardwareExpansionRequest.Action = varRecommendationHardwareExpansionRequestWithoutEmbeddedStruct.Action
		varRecommendationHardwareExpansionRequest.Message = varRecommendationHardwareExpansionRequestWithoutEmbeddedStruct.Message
		varRecommendationHardwareExpansionRequest.Status = varRecommendationHardwareExpansionRequestWithoutEmbeddedStruct.Status
		varRecommendationHardwareExpansionRequest.HwExpansionRequestItems = varRecommendationHardwareExpansionRequestWithoutEmbeddedStruct.HwExpansionRequestItems
		varRecommendationHardwareExpansionRequest.RegisteredDevice = varRecommendationHardwareExpansionRequestWithoutEmbeddedStruct.RegisteredDevice
		*o = RecommendationHardwareExpansionRequest(varRecommendationHardwareExpansionRequest)
	} else {
		return err
	}

	varRecommendationHardwareExpansionRequest := _RecommendationHardwareExpansionRequest{}

	err = json.Unmarshal(bytes, &varRecommendationHardwareExpansionRequest)
	if err == nil {
		o.RecommendationExpansionRequest = varRecommendationHardwareExpansionRequest.RecommendationExpansionRequest
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Action")
		delete(additionalProperties, "Message")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "HwExpansionRequestItems")
		delete(additionalProperties, "RegisteredDevice")

		// remove fields from embedded structs
		reflectRecommendationExpansionRequest := reflect.ValueOf(o.RecommendationExpansionRequest)
		for i := 0; i < reflectRecommendationExpansionRequest.Type().NumField(); i++ {
			t := reflectRecommendationExpansionRequest.Type().Field(i)

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

type NullableRecommendationHardwareExpansionRequest struct {
	value *RecommendationHardwareExpansionRequest
	isSet bool
}

func (v NullableRecommendationHardwareExpansionRequest) Get() *RecommendationHardwareExpansionRequest {
	return v.value
}

func (v *NullableRecommendationHardwareExpansionRequest) Set(val *RecommendationHardwareExpansionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRecommendationHardwareExpansionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRecommendationHardwareExpansionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecommendationHardwareExpansionRequest(val *RecommendationHardwareExpansionRequest) *NullableRecommendationHardwareExpansionRequest {
	return &NullableRecommendationHardwareExpansionRequest{value: val, isSet: true}
}

func (v NullableRecommendationHardwareExpansionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecommendationHardwareExpansionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
