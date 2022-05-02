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
	"reflect"
	"strings"
)

// NotificationSubscription Subscription is the abstract model that holds shared properties for all subscription types. Different types of subscriptions extend this model and add their own required properties.
type NotificationSubscription struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string                          `json:"ObjectType"`
	Actions    []NotificationAction            `json:"Actions,omitempty"`
	Conditions []NotificationAbstractCondition `json:"Conditions,omitempty"`
	// Subscription can be switched on/off without necessity to change the subscription settings: notification methods, conditions, etc. Ex.: Subscription MO can be configured, but switched off.
	Enabled              *bool `json:"Enabled,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NotificationSubscription NotificationSubscription

// NewNotificationSubscription instantiates a new NotificationSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationSubscription(classId string, objectType string) *NotificationSubscription {
	this := NotificationSubscription{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNotificationSubscriptionWithDefaults instantiates a new NotificationSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationSubscriptionWithDefaults() *NotificationSubscription {
	this := NotificationSubscription{}
	var classId string = "notification.AccountSubscription"
	this.ClassId = classId
	var objectType string = "notification.AccountSubscription"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NotificationSubscription) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NotificationSubscription) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NotificationSubscription) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NotificationSubscription) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NotificationSubscription) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NotificationSubscription) SetObjectType(v string) {
	o.ObjectType = v
}

// GetActions returns the Actions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NotificationSubscription) GetActions() []NotificationAction {
	if o == nil {
		var ret []NotificationAction
		return ret
	}
	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NotificationSubscription) GetActionsOk() ([]NotificationAction, bool) {
	if o == nil || o.Actions == nil {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *NotificationSubscription) HasActions() bool {
	if o != nil && o.Actions != nil {
		return true
	}

	return false
}

// SetActions gets a reference to the given []NotificationAction and assigns it to the Actions field.
func (o *NotificationSubscription) SetActions(v []NotificationAction) {
	o.Actions = v
}

// GetConditions returns the Conditions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NotificationSubscription) GetConditions() []NotificationAbstractCondition {
	if o == nil {
		var ret []NotificationAbstractCondition
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NotificationSubscription) GetConditionsOk() ([]NotificationAbstractCondition, bool) {
	if o == nil || o.Conditions == nil {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *NotificationSubscription) HasConditions() bool {
	if o != nil && o.Conditions != nil {
		return true
	}

	return false
}

// SetConditions gets a reference to the given []NotificationAbstractCondition and assigns it to the Conditions field.
func (o *NotificationSubscription) SetConditions(v []NotificationAbstractCondition) {
	o.Conditions = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *NotificationSubscription) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationSubscription) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *NotificationSubscription) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *NotificationSubscription) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o NotificationSubscription) MarshalJSON() ([]byte, error) {
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
	if o.Actions != nil {
		toSerialize["Actions"] = o.Actions
	}
	if o.Conditions != nil {
		toSerialize["Conditions"] = o.Conditions
	}
	if o.Enabled != nil {
		toSerialize["Enabled"] = o.Enabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NotificationSubscription) UnmarshalJSON(bytes []byte) (err error) {
	type NotificationSubscriptionWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string                          `json:"ObjectType"`
		Actions    []NotificationAction            `json:"Actions,omitempty"`
		Conditions []NotificationAbstractCondition `json:"Conditions,omitempty"`
		// Subscription can be switched on/off without necessity to change the subscription settings: notification methods, conditions, etc. Ex.: Subscription MO can be configured, but switched off.
		Enabled *bool `json:"Enabled,omitempty"`
	}

	varNotificationSubscriptionWithoutEmbeddedStruct := NotificationSubscriptionWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNotificationSubscriptionWithoutEmbeddedStruct)
	if err == nil {
		varNotificationSubscription := _NotificationSubscription{}
		varNotificationSubscription.ClassId = varNotificationSubscriptionWithoutEmbeddedStruct.ClassId
		varNotificationSubscription.ObjectType = varNotificationSubscriptionWithoutEmbeddedStruct.ObjectType
		varNotificationSubscription.Actions = varNotificationSubscriptionWithoutEmbeddedStruct.Actions
		varNotificationSubscription.Conditions = varNotificationSubscriptionWithoutEmbeddedStruct.Conditions
		varNotificationSubscription.Enabled = varNotificationSubscriptionWithoutEmbeddedStruct.Enabled
		*o = NotificationSubscription(varNotificationSubscription)
	} else {
		return err
	}

	varNotificationSubscription := _NotificationSubscription{}

	err = json.Unmarshal(bytes, &varNotificationSubscription)
	if err == nil {
		o.MoBaseMo = varNotificationSubscription.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Actions")
		delete(additionalProperties, "Conditions")
		delete(additionalProperties, "Enabled")

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

type NullableNotificationSubscription struct {
	value *NotificationSubscription
	isSet bool
}

func (v NullableNotificationSubscription) Get() *NotificationSubscription {
	return v.value
}

func (v *NullableNotificationSubscription) Set(val *NotificationSubscription) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationSubscription(val *NotificationSubscription) *NullableNotificationSubscription {
	return &NullableNotificationSubscription{value: val, isSet: true}
}

func (v NullableNotificationSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
