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

// IamEndPointUserRoleInventoryAllOf Definition of the list of properties defined in 'iam.EndPointUserRoleInventory', excluding properties defined in parent classes.
type IamEndPointUserRoleInventoryAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Denotes whether password has changed.
	ChangePassword *bool `json:"ChangePassword,omitempty"`
	// Enables the user account on the endpoint.
	Enabled *bool `json:"Enabled,omitempty"`
	// Indicates whether the value of the 'password' property has been set.
	IsPasswordSet *bool `json:"IsPasswordSet,omitempty"`
	// Valid login password of the user.
	Password *string `json:"Password,omitempty"`
	// An array of relationships to iamEndPointRole resources.
	EndPointRole         []IamEndPointRoleRelationship               `json:"EndPointRole,omitempty"`
	EndPointUser         *IamEndPointUserInventoryRelationship       `json:"EndPointUser,omitempty"`
	EndPointUserPolicy   *IamEndPointUserPolicyInventoryRelationship `json:"EndPointUserPolicy,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamEndPointUserRoleInventoryAllOf IamEndPointUserRoleInventoryAllOf

// NewIamEndPointUserRoleInventoryAllOf instantiates a new IamEndPointUserRoleInventoryAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamEndPointUserRoleInventoryAllOf(classId string, objectType string) *IamEndPointUserRoleInventoryAllOf {
	this := IamEndPointUserRoleInventoryAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIamEndPointUserRoleInventoryAllOfWithDefaults instantiates a new IamEndPointUserRoleInventoryAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamEndPointUserRoleInventoryAllOfWithDefaults() *IamEndPointUserRoleInventoryAllOf {
	this := IamEndPointUserRoleInventoryAllOf{}
	var classId string = "iam.EndPointUserRoleInventory"
	this.ClassId = classId
	var objectType string = "iam.EndPointUserRoleInventory"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IamEndPointUserRoleInventoryAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IamEndPointUserRoleInventoryAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IamEndPointUserRoleInventoryAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IamEndPointUserRoleInventoryAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IamEndPointUserRoleInventoryAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IamEndPointUserRoleInventoryAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetChangePassword returns the ChangePassword field value if set, zero value otherwise.
func (o *IamEndPointUserRoleInventoryAllOf) GetChangePassword() bool {
	if o == nil || o.ChangePassword == nil {
		var ret bool
		return ret
	}
	return *o.ChangePassword
}

// GetChangePasswordOk returns a tuple with the ChangePassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamEndPointUserRoleInventoryAllOf) GetChangePasswordOk() (*bool, bool) {
	if o == nil || o.ChangePassword == nil {
		return nil, false
	}
	return o.ChangePassword, true
}

// HasChangePassword returns a boolean if a field has been set.
func (o *IamEndPointUserRoleInventoryAllOf) HasChangePassword() bool {
	if o != nil && o.ChangePassword != nil {
		return true
	}

	return false
}

// SetChangePassword gets a reference to the given bool and assigns it to the ChangePassword field.
func (o *IamEndPointUserRoleInventoryAllOf) SetChangePassword(v bool) {
	o.ChangePassword = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *IamEndPointUserRoleInventoryAllOf) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamEndPointUserRoleInventoryAllOf) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *IamEndPointUserRoleInventoryAllOf) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *IamEndPointUserRoleInventoryAllOf) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetIsPasswordSet returns the IsPasswordSet field value if set, zero value otherwise.
func (o *IamEndPointUserRoleInventoryAllOf) GetIsPasswordSet() bool {
	if o == nil || o.IsPasswordSet == nil {
		var ret bool
		return ret
	}
	return *o.IsPasswordSet
}

// GetIsPasswordSetOk returns a tuple with the IsPasswordSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamEndPointUserRoleInventoryAllOf) GetIsPasswordSetOk() (*bool, bool) {
	if o == nil || o.IsPasswordSet == nil {
		return nil, false
	}
	return o.IsPasswordSet, true
}

// HasIsPasswordSet returns a boolean if a field has been set.
func (o *IamEndPointUserRoleInventoryAllOf) HasIsPasswordSet() bool {
	if o != nil && o.IsPasswordSet != nil {
		return true
	}

	return false
}

// SetIsPasswordSet gets a reference to the given bool and assigns it to the IsPasswordSet field.
func (o *IamEndPointUserRoleInventoryAllOf) SetIsPasswordSet(v bool) {
	o.IsPasswordSet = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *IamEndPointUserRoleInventoryAllOf) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamEndPointUserRoleInventoryAllOf) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *IamEndPointUserRoleInventoryAllOf) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *IamEndPointUserRoleInventoryAllOf) SetPassword(v string) {
	o.Password = &v
}

// GetEndPointRole returns the EndPointRole field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamEndPointUserRoleInventoryAllOf) GetEndPointRole() []IamEndPointRoleRelationship {
	if o == nil {
		var ret []IamEndPointRoleRelationship
		return ret
	}
	return o.EndPointRole
}

// GetEndPointRoleOk returns a tuple with the EndPointRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamEndPointUserRoleInventoryAllOf) GetEndPointRoleOk() ([]IamEndPointRoleRelationship, bool) {
	if o == nil || o.EndPointRole == nil {
		return nil, false
	}
	return o.EndPointRole, true
}

// HasEndPointRole returns a boolean if a field has been set.
func (o *IamEndPointUserRoleInventoryAllOf) HasEndPointRole() bool {
	if o != nil && o.EndPointRole != nil {
		return true
	}

	return false
}

// SetEndPointRole gets a reference to the given []IamEndPointRoleRelationship and assigns it to the EndPointRole field.
func (o *IamEndPointUserRoleInventoryAllOf) SetEndPointRole(v []IamEndPointRoleRelationship) {
	o.EndPointRole = v
}

// GetEndPointUser returns the EndPointUser field value if set, zero value otherwise.
func (o *IamEndPointUserRoleInventoryAllOf) GetEndPointUser() IamEndPointUserInventoryRelationship {
	if o == nil || o.EndPointUser == nil {
		var ret IamEndPointUserInventoryRelationship
		return ret
	}
	return *o.EndPointUser
}

// GetEndPointUserOk returns a tuple with the EndPointUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamEndPointUserRoleInventoryAllOf) GetEndPointUserOk() (*IamEndPointUserInventoryRelationship, bool) {
	if o == nil || o.EndPointUser == nil {
		return nil, false
	}
	return o.EndPointUser, true
}

// HasEndPointUser returns a boolean if a field has been set.
func (o *IamEndPointUserRoleInventoryAllOf) HasEndPointUser() bool {
	if o != nil && o.EndPointUser != nil {
		return true
	}

	return false
}

// SetEndPointUser gets a reference to the given IamEndPointUserInventoryRelationship and assigns it to the EndPointUser field.
func (o *IamEndPointUserRoleInventoryAllOf) SetEndPointUser(v IamEndPointUserInventoryRelationship) {
	o.EndPointUser = &v
}

// GetEndPointUserPolicy returns the EndPointUserPolicy field value if set, zero value otherwise.
func (o *IamEndPointUserRoleInventoryAllOf) GetEndPointUserPolicy() IamEndPointUserPolicyInventoryRelationship {
	if o == nil || o.EndPointUserPolicy == nil {
		var ret IamEndPointUserPolicyInventoryRelationship
		return ret
	}
	return *o.EndPointUserPolicy
}

// GetEndPointUserPolicyOk returns a tuple with the EndPointUserPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamEndPointUserRoleInventoryAllOf) GetEndPointUserPolicyOk() (*IamEndPointUserPolicyInventoryRelationship, bool) {
	if o == nil || o.EndPointUserPolicy == nil {
		return nil, false
	}
	return o.EndPointUserPolicy, true
}

// HasEndPointUserPolicy returns a boolean if a field has been set.
func (o *IamEndPointUserRoleInventoryAllOf) HasEndPointUserPolicy() bool {
	if o != nil && o.EndPointUserPolicy != nil {
		return true
	}

	return false
}

// SetEndPointUserPolicy gets a reference to the given IamEndPointUserPolicyInventoryRelationship and assigns it to the EndPointUserPolicy field.
func (o *IamEndPointUserRoleInventoryAllOf) SetEndPointUserPolicy(v IamEndPointUserPolicyInventoryRelationship) {
	o.EndPointUserPolicy = &v
}

func (o IamEndPointUserRoleInventoryAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ChangePassword != nil {
		toSerialize["ChangePassword"] = o.ChangePassword
	}
	if o.Enabled != nil {
		toSerialize["Enabled"] = o.Enabled
	}
	if o.IsPasswordSet != nil {
		toSerialize["IsPasswordSet"] = o.IsPasswordSet
	}
	if o.Password != nil {
		toSerialize["Password"] = o.Password
	}
	if o.EndPointRole != nil {
		toSerialize["EndPointRole"] = o.EndPointRole
	}
	if o.EndPointUser != nil {
		toSerialize["EndPointUser"] = o.EndPointUser
	}
	if o.EndPointUserPolicy != nil {
		toSerialize["EndPointUserPolicy"] = o.EndPointUserPolicy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IamEndPointUserRoleInventoryAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varIamEndPointUserRoleInventoryAllOf := _IamEndPointUserRoleInventoryAllOf{}

	if err = json.Unmarshal(bytes, &varIamEndPointUserRoleInventoryAllOf); err == nil {
		*o = IamEndPointUserRoleInventoryAllOf(varIamEndPointUserRoleInventoryAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ChangePassword")
		delete(additionalProperties, "Enabled")
		delete(additionalProperties, "IsPasswordSet")
		delete(additionalProperties, "Password")
		delete(additionalProperties, "EndPointRole")
		delete(additionalProperties, "EndPointUser")
		delete(additionalProperties, "EndPointUserPolicy")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIamEndPointUserRoleInventoryAllOf struct {
	value *IamEndPointUserRoleInventoryAllOf
	isSet bool
}

func (v NullableIamEndPointUserRoleInventoryAllOf) Get() *IamEndPointUserRoleInventoryAllOf {
	return v.value
}

func (v *NullableIamEndPointUserRoleInventoryAllOf) Set(val *IamEndPointUserRoleInventoryAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIamEndPointUserRoleInventoryAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIamEndPointUserRoleInventoryAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamEndPointUserRoleInventoryAllOf(val *IamEndPointUserRoleInventoryAllOf) *NullableIamEndPointUserRoleInventoryAllOf {
	return &NullableIamEndPointUserRoleInventoryAllOf{value: val, isSet: true}
}

func (v NullableIamEndPointUserRoleInventoryAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamEndPointUserRoleInventoryAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
