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

// CrdCustomResourceAllOf Definition of the list of properties defined in 'crd.CustomResource', excluding properties defined in parent classes.
type CrdCustomResourceAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Type of custom resource or 'kind' in K8s.
	DcLauncher *string `json:"DcLauncher,omitempty"`
	// Docker image URL for public cloud DC.
	Image *string `json:"Image,omitempty"`
	// A string property called name which is used as part of a uniqueness constraint. See 'identity' specification in this MO definition.
	Name *string `json:"Name,omitempty"`
	// Namespace to launch the deployment associated with the custom resource.
	Namespace *string `json:"Namespace,omitempty"`
	// Port used for public cloud DC.
	Port       *int64                            `json:"Port,omitempty"`
	Properties []CrdCustomResourceConfigProperty `json:"Properties,omitempty"`
	// Target ID for public cloud DC.
	TargetId *string `json:"TargetId,omitempty"`
	// Target Moid for public cloud DC.
	TargetMoid *string `json:"TargetMoid,omitempty"`
	// Target type for public cloud DC.
	TargetType           *string                 `json:"TargetType,omitempty"`
	Account              *IamAccountRelationship `json:"Account,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CrdCustomResourceAllOf CrdCustomResourceAllOf

// NewCrdCustomResourceAllOf instantiates a new CrdCustomResourceAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCrdCustomResourceAllOf(classId string, objectType string) *CrdCustomResourceAllOf {
	this := CrdCustomResourceAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCrdCustomResourceAllOfWithDefaults instantiates a new CrdCustomResourceAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCrdCustomResourceAllOfWithDefaults() *CrdCustomResourceAllOf {
	this := CrdCustomResourceAllOf{}
	var classId string = "crd.CustomResource"
	this.ClassId = classId
	var objectType string = "crd.CustomResource"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CrdCustomResourceAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CrdCustomResourceAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CrdCustomResourceAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CrdCustomResourceAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CrdCustomResourceAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CrdCustomResourceAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDcLauncher returns the DcLauncher field value if set, zero value otherwise.
func (o *CrdCustomResourceAllOf) GetDcLauncher() string {
	if o == nil || o.DcLauncher == nil {
		var ret string
		return ret
	}
	return *o.DcLauncher
}

// GetDcLauncherOk returns a tuple with the DcLauncher field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CrdCustomResourceAllOf) GetDcLauncherOk() (*string, bool) {
	if o == nil || o.DcLauncher == nil {
		return nil, false
	}
	return o.DcLauncher, true
}

// HasDcLauncher returns a boolean if a field has been set.
func (o *CrdCustomResourceAllOf) HasDcLauncher() bool {
	if o != nil && o.DcLauncher != nil {
		return true
	}

	return false
}

// SetDcLauncher gets a reference to the given string and assigns it to the DcLauncher field.
func (o *CrdCustomResourceAllOf) SetDcLauncher(v string) {
	o.DcLauncher = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *CrdCustomResourceAllOf) GetImage() string {
	if o == nil || o.Image == nil {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CrdCustomResourceAllOf) GetImageOk() (*string, bool) {
	if o == nil || o.Image == nil {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *CrdCustomResourceAllOf) HasImage() bool {
	if o != nil && o.Image != nil {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *CrdCustomResourceAllOf) SetImage(v string) {
	o.Image = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CrdCustomResourceAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CrdCustomResourceAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CrdCustomResourceAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CrdCustomResourceAllOf) SetName(v string) {
	o.Name = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *CrdCustomResourceAllOf) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CrdCustomResourceAllOf) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *CrdCustomResourceAllOf) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *CrdCustomResourceAllOf) SetNamespace(v string) {
	o.Namespace = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *CrdCustomResourceAllOf) GetPort() int64 {
	if o == nil || o.Port == nil {
		var ret int64
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CrdCustomResourceAllOf) GetPortOk() (*int64, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *CrdCustomResourceAllOf) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int64 and assigns it to the Port field.
func (o *CrdCustomResourceAllOf) SetPort(v int64) {
	o.Port = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CrdCustomResourceAllOf) GetProperties() []CrdCustomResourceConfigProperty {
	if o == nil {
		var ret []CrdCustomResourceConfigProperty
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CrdCustomResourceAllOf) GetPropertiesOk() ([]CrdCustomResourceConfigProperty, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *CrdCustomResourceAllOf) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given []CrdCustomResourceConfigProperty and assigns it to the Properties field.
func (o *CrdCustomResourceAllOf) SetProperties(v []CrdCustomResourceConfigProperty) {
	o.Properties = v
}

// GetTargetId returns the TargetId field value if set, zero value otherwise.
func (o *CrdCustomResourceAllOf) GetTargetId() string {
	if o == nil || o.TargetId == nil {
		var ret string
		return ret
	}
	return *o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CrdCustomResourceAllOf) GetTargetIdOk() (*string, bool) {
	if o == nil || o.TargetId == nil {
		return nil, false
	}
	return o.TargetId, true
}

// HasTargetId returns a boolean if a field has been set.
func (o *CrdCustomResourceAllOf) HasTargetId() bool {
	if o != nil && o.TargetId != nil {
		return true
	}

	return false
}

// SetTargetId gets a reference to the given string and assigns it to the TargetId field.
func (o *CrdCustomResourceAllOf) SetTargetId(v string) {
	o.TargetId = &v
}

// GetTargetMoid returns the TargetMoid field value if set, zero value otherwise.
func (o *CrdCustomResourceAllOf) GetTargetMoid() string {
	if o == nil || o.TargetMoid == nil {
		var ret string
		return ret
	}
	return *o.TargetMoid
}

// GetTargetMoidOk returns a tuple with the TargetMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CrdCustomResourceAllOf) GetTargetMoidOk() (*string, bool) {
	if o == nil || o.TargetMoid == nil {
		return nil, false
	}
	return o.TargetMoid, true
}

// HasTargetMoid returns a boolean if a field has been set.
func (o *CrdCustomResourceAllOf) HasTargetMoid() bool {
	if o != nil && o.TargetMoid != nil {
		return true
	}

	return false
}

// SetTargetMoid gets a reference to the given string and assigns it to the TargetMoid field.
func (o *CrdCustomResourceAllOf) SetTargetMoid(v string) {
	o.TargetMoid = &v
}

// GetTargetType returns the TargetType field value if set, zero value otherwise.
func (o *CrdCustomResourceAllOf) GetTargetType() string {
	if o == nil || o.TargetType == nil {
		var ret string
		return ret
	}
	return *o.TargetType
}

// GetTargetTypeOk returns a tuple with the TargetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CrdCustomResourceAllOf) GetTargetTypeOk() (*string, bool) {
	if o == nil || o.TargetType == nil {
		return nil, false
	}
	return o.TargetType, true
}

// HasTargetType returns a boolean if a field has been set.
func (o *CrdCustomResourceAllOf) HasTargetType() bool {
	if o != nil && o.TargetType != nil {
		return true
	}

	return false
}

// SetTargetType gets a reference to the given string and assigns it to the TargetType field.
func (o *CrdCustomResourceAllOf) SetTargetType(v string) {
	o.TargetType = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *CrdCustomResourceAllOf) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CrdCustomResourceAllOf) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *CrdCustomResourceAllOf) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *CrdCustomResourceAllOf) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

func (o CrdCustomResourceAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DcLauncher != nil {
		toSerialize["DcLauncher"] = o.DcLauncher
	}
	if o.Image != nil {
		toSerialize["Image"] = o.Image
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Namespace != nil {
		toSerialize["Namespace"] = o.Namespace
	}
	if o.Port != nil {
		toSerialize["Port"] = o.Port
	}
	if o.Properties != nil {
		toSerialize["Properties"] = o.Properties
	}
	if o.TargetId != nil {
		toSerialize["TargetId"] = o.TargetId
	}
	if o.TargetMoid != nil {
		toSerialize["TargetMoid"] = o.TargetMoid
	}
	if o.TargetType != nil {
		toSerialize["TargetType"] = o.TargetType
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CrdCustomResourceAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varCrdCustomResourceAllOf := _CrdCustomResourceAllOf{}

	if err = json.Unmarshal(bytes, &varCrdCustomResourceAllOf); err == nil {
		*o = CrdCustomResourceAllOf(varCrdCustomResourceAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DcLauncher")
		delete(additionalProperties, "Image")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Namespace")
		delete(additionalProperties, "Port")
		delete(additionalProperties, "Properties")
		delete(additionalProperties, "TargetId")
		delete(additionalProperties, "TargetMoid")
		delete(additionalProperties, "TargetType")
		delete(additionalProperties, "Account")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCrdCustomResourceAllOf struct {
	value *CrdCustomResourceAllOf
	isSet bool
}

func (v NullableCrdCustomResourceAllOf) Get() *CrdCustomResourceAllOf {
	return v.value
}

func (v *NullableCrdCustomResourceAllOf) Set(val *CrdCustomResourceAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCrdCustomResourceAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCrdCustomResourceAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCrdCustomResourceAllOf(val *CrdCustomResourceAllOf) *NullableCrdCustomResourceAllOf {
	return &NullableCrdCustomResourceAllOf{value: val, isSet: true}
}

func (v NullableCrdCustomResourceAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCrdCustomResourceAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
