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

// IamEndPointUserInventory Endpoint User or Local User.
type IamEndPointUserInventory struct {
	PolicyAbstractInventory
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Name of the user to be created on the endpoint. It can be any string that adheres to the following constraints. It can have alphanumeric characters, dots, underscores and hyphen. It cannot be more than 16 characters.
	Name *string `json:"Name,omitempty"`
	// UserId for the end point user.
	UserId *string `json:"UserId,omitempty"`
	// An array of relationships to iamEndPointUserRoleInventory resources.
	EndPointUserRole     []IamEndPointUserRoleInventoryRelationship `json:"EndPointUserRole,omitempty"`
	TargetMo             *MoBaseMoRelationship                      `json:"TargetMo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamEndPointUserInventory IamEndPointUserInventory

// NewIamEndPointUserInventory instantiates a new IamEndPointUserInventory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamEndPointUserInventory(classId string, objectType string) *IamEndPointUserInventory {
	this := IamEndPointUserInventory{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIamEndPointUserInventoryWithDefaults instantiates a new IamEndPointUserInventory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamEndPointUserInventoryWithDefaults() *IamEndPointUserInventory {
	this := IamEndPointUserInventory{}
	var classId string = "iam.EndPointUserInventory"
	this.ClassId = classId
	var objectType string = "iam.EndPointUserInventory"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IamEndPointUserInventory) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IamEndPointUserInventory) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IamEndPointUserInventory) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IamEndPointUserInventory) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IamEndPointUserInventory) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IamEndPointUserInventory) SetObjectType(v string) {
	o.ObjectType = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IamEndPointUserInventory) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamEndPointUserInventory) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IamEndPointUserInventory) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IamEndPointUserInventory) SetName(v string) {
	o.Name = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *IamEndPointUserInventory) GetUserId() string {
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamEndPointUserInventory) GetUserIdOk() (*string, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *IamEndPointUserInventory) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *IamEndPointUserInventory) SetUserId(v string) {
	o.UserId = &v
}

// GetEndPointUserRole returns the EndPointUserRole field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamEndPointUserInventory) GetEndPointUserRole() []IamEndPointUserRoleInventoryRelationship {
	if o == nil {
		var ret []IamEndPointUserRoleInventoryRelationship
		return ret
	}
	return o.EndPointUserRole
}

// GetEndPointUserRoleOk returns a tuple with the EndPointUserRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamEndPointUserInventory) GetEndPointUserRoleOk() ([]IamEndPointUserRoleInventoryRelationship, bool) {
	if o == nil || o.EndPointUserRole == nil {
		return nil, false
	}
	return o.EndPointUserRole, true
}

// HasEndPointUserRole returns a boolean if a field has been set.
func (o *IamEndPointUserInventory) HasEndPointUserRole() bool {
	if o != nil && o.EndPointUserRole != nil {
		return true
	}

	return false
}

// SetEndPointUserRole gets a reference to the given []IamEndPointUserRoleInventoryRelationship and assigns it to the EndPointUserRole field.
func (o *IamEndPointUserInventory) SetEndPointUserRole(v []IamEndPointUserRoleInventoryRelationship) {
	o.EndPointUserRole = v
}

// GetTargetMo returns the TargetMo field value if set, zero value otherwise.
func (o *IamEndPointUserInventory) GetTargetMo() MoBaseMoRelationship {
	if o == nil || o.TargetMo == nil {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.TargetMo
}

// GetTargetMoOk returns a tuple with the TargetMo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamEndPointUserInventory) GetTargetMoOk() (*MoBaseMoRelationship, bool) {
	if o == nil || o.TargetMo == nil {
		return nil, false
	}
	return o.TargetMo, true
}

// HasTargetMo returns a boolean if a field has been set.
func (o *IamEndPointUserInventory) HasTargetMo() bool {
	if o != nil && o.TargetMo != nil {
		return true
	}

	return false
}

// SetTargetMo gets a reference to the given MoBaseMoRelationship and assigns it to the TargetMo field.
func (o *IamEndPointUserInventory) SetTargetMo(v MoBaseMoRelationship) {
	o.TargetMo = &v
}

func (o IamEndPointUserInventory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractInventory, errPolicyAbstractInventory := json.Marshal(o.PolicyAbstractInventory)
	if errPolicyAbstractInventory != nil {
		return []byte{}, errPolicyAbstractInventory
	}
	errPolicyAbstractInventory = json.Unmarshal([]byte(serializedPolicyAbstractInventory), &toSerialize)
	if errPolicyAbstractInventory != nil {
		return []byte{}, errPolicyAbstractInventory
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.UserId != nil {
		toSerialize["UserId"] = o.UserId
	}
	if o.EndPointUserRole != nil {
		toSerialize["EndPointUserRole"] = o.EndPointUserRole
	}
	if o.TargetMo != nil {
		toSerialize["TargetMo"] = o.TargetMo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IamEndPointUserInventory) UnmarshalJSON(bytes []byte) (err error) {
	type IamEndPointUserInventoryWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Name of the user to be created on the endpoint. It can be any string that adheres to the following constraints. It can have alphanumeric characters, dots, underscores and hyphen. It cannot be more than 16 characters.
		Name *string `json:"Name,omitempty"`
		// UserId for the end point user.
		UserId *string `json:"UserId,omitempty"`
		// An array of relationships to iamEndPointUserRoleInventory resources.
		EndPointUserRole []IamEndPointUserRoleInventoryRelationship `json:"EndPointUserRole,omitempty"`
		TargetMo         *MoBaseMoRelationship                      `json:"TargetMo,omitempty"`
	}

	varIamEndPointUserInventoryWithoutEmbeddedStruct := IamEndPointUserInventoryWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varIamEndPointUserInventoryWithoutEmbeddedStruct)
	if err == nil {
		varIamEndPointUserInventory := _IamEndPointUserInventory{}
		varIamEndPointUserInventory.ClassId = varIamEndPointUserInventoryWithoutEmbeddedStruct.ClassId
		varIamEndPointUserInventory.ObjectType = varIamEndPointUserInventoryWithoutEmbeddedStruct.ObjectType
		varIamEndPointUserInventory.Name = varIamEndPointUserInventoryWithoutEmbeddedStruct.Name
		varIamEndPointUserInventory.UserId = varIamEndPointUserInventoryWithoutEmbeddedStruct.UserId
		varIamEndPointUserInventory.EndPointUserRole = varIamEndPointUserInventoryWithoutEmbeddedStruct.EndPointUserRole
		varIamEndPointUserInventory.TargetMo = varIamEndPointUserInventoryWithoutEmbeddedStruct.TargetMo
		*o = IamEndPointUserInventory(varIamEndPointUserInventory)
	} else {
		return err
	}

	varIamEndPointUserInventory := _IamEndPointUserInventory{}

	err = json.Unmarshal(bytes, &varIamEndPointUserInventory)
	if err == nil {
		o.PolicyAbstractInventory = varIamEndPointUserInventory.PolicyAbstractInventory
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "UserId")
		delete(additionalProperties, "EndPointUserRole")
		delete(additionalProperties, "TargetMo")

		// remove fields from embedded structs
		reflectPolicyAbstractInventory := reflect.ValueOf(o.PolicyAbstractInventory)
		for i := 0; i < reflectPolicyAbstractInventory.Type().NumField(); i++ {
			t := reflectPolicyAbstractInventory.Type().Field(i)

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

type NullableIamEndPointUserInventory struct {
	value *IamEndPointUserInventory
	isSet bool
}

func (v NullableIamEndPointUserInventory) Get() *IamEndPointUserInventory {
	return v.value
}

func (v *NullableIamEndPointUserInventory) Set(val *IamEndPointUserInventory) {
	v.value = val
	v.isSet = true
}

func (v NullableIamEndPointUserInventory) IsSet() bool {
	return v.isSet
}

func (v *NullableIamEndPointUserInventory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamEndPointUserInventory(val *IamEndPointUserInventory) *NullableIamEndPointUserInventory {
	return &NullableIamEndPointUserInventory{value: val, isSet: true}
}

func (v NullableIamEndPointUserInventory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamEndPointUserInventory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
