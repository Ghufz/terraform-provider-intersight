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

// IamUserGroupAllOf Definition of the list of properties defined in 'iam.UserGroup', excluding properties defined in parent classes.
type IamUserGroupAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The name of the user group which the dynamic user belongs to.
	Name         *string                      `json:"Name,omitempty"`
	Idp          *IamIdpRelationship          `json:"Idp,omitempty"`
	Idpreference *IamIdpReferenceRelationship `json:"Idpreference,omitempty"`
	// An array of relationships to iamPermission resources.
	Permissions []IamPermissionRelationship `json:"Permissions,omitempty"`
	Qualifier   *IamQualifierRelationship   `json:"Qualifier,omitempty"`
	// An array of relationships to iamUser resources.
	Users                []IamUserRelationship `json:"Users,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamUserGroupAllOf IamUserGroupAllOf

// NewIamUserGroupAllOf instantiates a new IamUserGroupAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamUserGroupAllOf(classId string, objectType string) *IamUserGroupAllOf {
	this := IamUserGroupAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIamUserGroupAllOfWithDefaults instantiates a new IamUserGroupAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamUserGroupAllOfWithDefaults() *IamUserGroupAllOf {
	this := IamUserGroupAllOf{}
	var classId string = "iam.UserGroup"
	this.ClassId = classId
	var objectType string = "iam.UserGroup"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IamUserGroupAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IamUserGroupAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IamUserGroupAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IamUserGroupAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IamUserGroupAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IamUserGroupAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IamUserGroupAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamUserGroupAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IamUserGroupAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IamUserGroupAllOf) SetName(v string) {
	o.Name = &v
}

// GetIdp returns the Idp field value if set, zero value otherwise.
func (o *IamUserGroupAllOf) GetIdp() IamIdpRelationship {
	if o == nil || o.Idp == nil {
		var ret IamIdpRelationship
		return ret
	}
	return *o.Idp
}

// GetIdpOk returns a tuple with the Idp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamUserGroupAllOf) GetIdpOk() (*IamIdpRelationship, bool) {
	if o == nil || o.Idp == nil {
		return nil, false
	}
	return o.Idp, true
}

// HasIdp returns a boolean if a field has been set.
func (o *IamUserGroupAllOf) HasIdp() bool {
	if o != nil && o.Idp != nil {
		return true
	}

	return false
}

// SetIdp gets a reference to the given IamIdpRelationship and assigns it to the Idp field.
func (o *IamUserGroupAllOf) SetIdp(v IamIdpRelationship) {
	o.Idp = &v
}

// GetIdpreference returns the Idpreference field value if set, zero value otherwise.
func (o *IamUserGroupAllOf) GetIdpreference() IamIdpReferenceRelationship {
	if o == nil || o.Idpreference == nil {
		var ret IamIdpReferenceRelationship
		return ret
	}
	return *o.Idpreference
}

// GetIdpreferenceOk returns a tuple with the Idpreference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamUserGroupAllOf) GetIdpreferenceOk() (*IamIdpReferenceRelationship, bool) {
	if o == nil || o.Idpreference == nil {
		return nil, false
	}
	return o.Idpreference, true
}

// HasIdpreference returns a boolean if a field has been set.
func (o *IamUserGroupAllOf) HasIdpreference() bool {
	if o != nil && o.Idpreference != nil {
		return true
	}

	return false
}

// SetIdpreference gets a reference to the given IamIdpReferenceRelationship and assigns it to the Idpreference field.
func (o *IamUserGroupAllOf) SetIdpreference(v IamIdpReferenceRelationship) {
	o.Idpreference = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamUserGroupAllOf) GetPermissions() []IamPermissionRelationship {
	if o == nil {
		var ret []IamPermissionRelationship
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamUserGroupAllOf) GetPermissionsOk() ([]IamPermissionRelationship, bool) {
	if o == nil || o.Permissions == nil {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *IamUserGroupAllOf) HasPermissions() bool {
	if o != nil && o.Permissions != nil {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []IamPermissionRelationship and assigns it to the Permissions field.
func (o *IamUserGroupAllOf) SetPermissions(v []IamPermissionRelationship) {
	o.Permissions = v
}

// GetQualifier returns the Qualifier field value if set, zero value otherwise.
func (o *IamUserGroupAllOf) GetQualifier() IamQualifierRelationship {
	if o == nil || o.Qualifier == nil {
		var ret IamQualifierRelationship
		return ret
	}
	return *o.Qualifier
}

// GetQualifierOk returns a tuple with the Qualifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamUserGroupAllOf) GetQualifierOk() (*IamQualifierRelationship, bool) {
	if o == nil || o.Qualifier == nil {
		return nil, false
	}
	return o.Qualifier, true
}

// HasQualifier returns a boolean if a field has been set.
func (o *IamUserGroupAllOf) HasQualifier() bool {
	if o != nil && o.Qualifier != nil {
		return true
	}

	return false
}

// SetQualifier gets a reference to the given IamQualifierRelationship and assigns it to the Qualifier field.
func (o *IamUserGroupAllOf) SetQualifier(v IamQualifierRelationship) {
	o.Qualifier = &v
}

// GetUsers returns the Users field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamUserGroupAllOf) GetUsers() []IamUserRelationship {
	if o == nil {
		var ret []IamUserRelationship
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamUserGroupAllOf) GetUsersOk() ([]IamUserRelationship, bool) {
	if o == nil || o.Users == nil {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *IamUserGroupAllOf) HasUsers() bool {
	if o != nil && o.Users != nil {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []IamUserRelationship and assigns it to the Users field.
func (o *IamUserGroupAllOf) SetUsers(v []IamUserRelationship) {
	o.Users = v
}

func (o IamUserGroupAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Idp != nil {
		toSerialize["Idp"] = o.Idp
	}
	if o.Idpreference != nil {
		toSerialize["Idpreference"] = o.Idpreference
	}
	if o.Permissions != nil {
		toSerialize["Permissions"] = o.Permissions
	}
	if o.Qualifier != nil {
		toSerialize["Qualifier"] = o.Qualifier
	}
	if o.Users != nil {
		toSerialize["Users"] = o.Users
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IamUserGroupAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varIamUserGroupAllOf := _IamUserGroupAllOf{}

	if err = json.Unmarshal(bytes, &varIamUserGroupAllOf); err == nil {
		*o = IamUserGroupAllOf(varIamUserGroupAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Idp")
		delete(additionalProperties, "Idpreference")
		delete(additionalProperties, "Permissions")
		delete(additionalProperties, "Qualifier")
		delete(additionalProperties, "Users")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIamUserGroupAllOf struct {
	value *IamUserGroupAllOf
	isSet bool
}

func (v NullableIamUserGroupAllOf) Get() *IamUserGroupAllOf {
	return v.value
}

func (v *NullableIamUserGroupAllOf) Set(val *IamUserGroupAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIamUserGroupAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIamUserGroupAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamUserGroupAllOf(val *IamUserGroupAllOf) *NullableIamUserGroupAllOf {
	return &NullableIamUserGroupAllOf{value: val, isSet: true}
}

func (v NullableIamUserGroupAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamUserGroupAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
