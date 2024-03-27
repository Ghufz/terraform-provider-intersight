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
)

// PartnerintegrationDocIssuesAllOf Definition of the list of properties defined in 'partnerintegration.DocIssues', excluding properties defined in parent classes.
type PartnerintegrationDocIssuesAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// List of documentation issues.
	DocumentationIssues  interface{}                              `json:"DocumentationIssues,omitempty"`
	Inventory            *PartnerintegrationInventoryRelationship `json:"Inventory,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PartnerintegrationDocIssuesAllOf PartnerintegrationDocIssuesAllOf

// NewPartnerintegrationDocIssuesAllOf instantiates a new PartnerintegrationDocIssuesAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartnerintegrationDocIssuesAllOf(classId string, objectType string) *PartnerintegrationDocIssuesAllOf {
	this := PartnerintegrationDocIssuesAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewPartnerintegrationDocIssuesAllOfWithDefaults instantiates a new PartnerintegrationDocIssuesAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartnerintegrationDocIssuesAllOfWithDefaults() *PartnerintegrationDocIssuesAllOf {
	this := PartnerintegrationDocIssuesAllOf{}
	var classId string = "partnerintegration.DocIssues"
	this.ClassId = classId
	var objectType string = "partnerintegration.DocIssues"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *PartnerintegrationDocIssuesAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *PartnerintegrationDocIssuesAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *PartnerintegrationDocIssuesAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *PartnerintegrationDocIssuesAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *PartnerintegrationDocIssuesAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *PartnerintegrationDocIssuesAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDocumentationIssues returns the DocumentationIssues field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PartnerintegrationDocIssuesAllOf) GetDocumentationIssues() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.DocumentationIssues
}

// GetDocumentationIssuesOk returns a tuple with the DocumentationIssues field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartnerintegrationDocIssuesAllOf) GetDocumentationIssuesOk() (*interface{}, bool) {
	if o == nil || o.DocumentationIssues == nil {
		return nil, false
	}
	return &o.DocumentationIssues, true
}

// HasDocumentationIssues returns a boolean if a field has been set.
func (o *PartnerintegrationDocIssuesAllOf) HasDocumentationIssues() bool {
	if o != nil && o.DocumentationIssues != nil {
		return true
	}

	return false
}

// SetDocumentationIssues gets a reference to the given interface{} and assigns it to the DocumentationIssues field.
func (o *PartnerintegrationDocIssuesAllOf) SetDocumentationIssues(v interface{}) {
	o.DocumentationIssues = v
}

// GetInventory returns the Inventory field value if set, zero value otherwise.
func (o *PartnerintegrationDocIssuesAllOf) GetInventory() PartnerintegrationInventoryRelationship {
	if o == nil || o.Inventory == nil {
		var ret PartnerintegrationInventoryRelationship
		return ret
	}
	return *o.Inventory
}

// GetInventoryOk returns a tuple with the Inventory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerintegrationDocIssuesAllOf) GetInventoryOk() (*PartnerintegrationInventoryRelationship, bool) {
	if o == nil || o.Inventory == nil {
		return nil, false
	}
	return o.Inventory, true
}

// HasInventory returns a boolean if a field has been set.
func (o *PartnerintegrationDocIssuesAllOf) HasInventory() bool {
	if o != nil && o.Inventory != nil {
		return true
	}

	return false
}

// SetInventory gets a reference to the given PartnerintegrationInventoryRelationship and assigns it to the Inventory field.
func (o *PartnerintegrationDocIssuesAllOf) SetInventory(v PartnerintegrationInventoryRelationship) {
	o.Inventory = &v
}

func (o PartnerintegrationDocIssuesAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DocumentationIssues != nil {
		toSerialize["DocumentationIssues"] = o.DocumentationIssues
	}
	if o.Inventory != nil {
		toSerialize["Inventory"] = o.Inventory
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PartnerintegrationDocIssuesAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varPartnerintegrationDocIssuesAllOf := _PartnerintegrationDocIssuesAllOf{}

	if err = json.Unmarshal(bytes, &varPartnerintegrationDocIssuesAllOf); err == nil {
		*o = PartnerintegrationDocIssuesAllOf(varPartnerintegrationDocIssuesAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DocumentationIssues")
		delete(additionalProperties, "Inventory")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePartnerintegrationDocIssuesAllOf struct {
	value *PartnerintegrationDocIssuesAllOf
	isSet bool
}

func (v NullablePartnerintegrationDocIssuesAllOf) Get() *PartnerintegrationDocIssuesAllOf {
	return v.value
}

func (v *NullablePartnerintegrationDocIssuesAllOf) Set(val *PartnerintegrationDocIssuesAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePartnerintegrationDocIssuesAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePartnerintegrationDocIssuesAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartnerintegrationDocIssuesAllOf(val *PartnerintegrationDocIssuesAllOf) *NullablePartnerintegrationDocIssuesAllOf {
	return &NullablePartnerintegrationDocIssuesAllOf{value: val, isSet: true}
}

func (v NullablePartnerintegrationDocIssuesAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartnerintegrationDocIssuesAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
