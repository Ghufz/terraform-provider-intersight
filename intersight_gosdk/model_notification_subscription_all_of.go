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

// NotificationSubscriptionAllOf Definition of the list of properties defined in 'notification.Subscription', excluding properties defined in parent classes.
type NotificationSubscriptionAllOf struct {
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

type _NotificationSubscriptionAllOf NotificationSubscriptionAllOf

// NewNotificationSubscriptionAllOf instantiates a new NotificationSubscriptionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationSubscriptionAllOf(classId string, objectType string) *NotificationSubscriptionAllOf {
	this := NotificationSubscriptionAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNotificationSubscriptionAllOfWithDefaults instantiates a new NotificationSubscriptionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationSubscriptionAllOfWithDefaults() *NotificationSubscriptionAllOf {
	this := NotificationSubscriptionAllOf{}
	var classId string = "notification.AccountSubscription"
	this.ClassId = classId
	var objectType string = "notification.AccountSubscription"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NotificationSubscriptionAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NotificationSubscriptionAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NotificationSubscriptionAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NotificationSubscriptionAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NotificationSubscriptionAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NotificationSubscriptionAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetActions returns the Actions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NotificationSubscriptionAllOf) GetActions() []NotificationAction {
	if o == nil {
		var ret []NotificationAction
		return ret
	}
	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NotificationSubscriptionAllOf) GetActionsOk() ([]NotificationAction, bool) {
	if o == nil || o.Actions == nil {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *NotificationSubscriptionAllOf) HasActions() bool {
	if o != nil && o.Actions != nil {
		return true
	}

	return false
}

// SetActions gets a reference to the given []NotificationAction and assigns it to the Actions field.
func (o *NotificationSubscriptionAllOf) SetActions(v []NotificationAction) {
	o.Actions = v
}

// GetConditions returns the Conditions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NotificationSubscriptionAllOf) GetConditions() []NotificationAbstractCondition {
	if o == nil {
		var ret []NotificationAbstractCondition
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NotificationSubscriptionAllOf) GetConditionsOk() ([]NotificationAbstractCondition, bool) {
	if o == nil || o.Conditions == nil {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *NotificationSubscriptionAllOf) HasConditions() bool {
	if o != nil && o.Conditions != nil {
		return true
	}

	return false
}

// SetConditions gets a reference to the given []NotificationAbstractCondition and assigns it to the Conditions field.
func (o *NotificationSubscriptionAllOf) SetConditions(v []NotificationAbstractCondition) {
	o.Conditions = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *NotificationSubscriptionAllOf) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationSubscriptionAllOf) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *NotificationSubscriptionAllOf) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *NotificationSubscriptionAllOf) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o NotificationSubscriptionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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

func (o *NotificationSubscriptionAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varNotificationSubscriptionAllOf := _NotificationSubscriptionAllOf{}

	if err = json.Unmarshal(bytes, &varNotificationSubscriptionAllOf); err == nil {
		*o = NotificationSubscriptionAllOf(varNotificationSubscriptionAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Actions")
		delete(additionalProperties, "Conditions")
		delete(additionalProperties, "Enabled")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNotificationSubscriptionAllOf struct {
	value *NotificationSubscriptionAllOf
	isSet bool
}

func (v NullableNotificationSubscriptionAllOf) Get() *NotificationSubscriptionAllOf {
	return v.value
}

func (v *NullableNotificationSubscriptionAllOf) Set(val *NotificationSubscriptionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationSubscriptionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationSubscriptionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationSubscriptionAllOf(val *NotificationSubscriptionAllOf) *NullableNotificationSubscriptionAllOf {
	return &NullableNotificationSubscriptionAllOf{value: val, isSet: true}
}

func (v NullableNotificationSubscriptionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationSubscriptionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
