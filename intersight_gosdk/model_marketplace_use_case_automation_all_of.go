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

// MarketplaceUseCaseAutomationAllOf Definition of the list of properties defined in 'marketplace.UseCaseAutomation', excluding properties defined in parent classes.
type MarketplaceUseCaseAutomationAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// A description for the automation
	Description *string `json:"Description,omitempty"`
	// A name for the automation
	Name                 *string `json:"Name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MarketplaceUseCaseAutomationAllOf MarketplaceUseCaseAutomationAllOf

// NewMarketplaceUseCaseAutomationAllOf instantiates a new MarketplaceUseCaseAutomationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarketplaceUseCaseAutomationAllOf(classId string, objectType string) *MarketplaceUseCaseAutomationAllOf {
	this := MarketplaceUseCaseAutomationAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewMarketplaceUseCaseAutomationAllOfWithDefaults instantiates a new MarketplaceUseCaseAutomationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarketplaceUseCaseAutomationAllOfWithDefaults() *MarketplaceUseCaseAutomationAllOf {
	this := MarketplaceUseCaseAutomationAllOf{}
	var classId string = "marketplace.UseCaseAutomation"
	this.ClassId = classId
	var objectType string = "marketplace.UseCaseAutomation"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *MarketplaceUseCaseAutomationAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *MarketplaceUseCaseAutomationAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *MarketplaceUseCaseAutomationAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *MarketplaceUseCaseAutomationAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *MarketplaceUseCaseAutomationAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *MarketplaceUseCaseAutomationAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *MarketplaceUseCaseAutomationAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketplaceUseCaseAutomationAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MarketplaceUseCaseAutomationAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MarketplaceUseCaseAutomationAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MarketplaceUseCaseAutomationAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketplaceUseCaseAutomationAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MarketplaceUseCaseAutomationAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MarketplaceUseCaseAutomationAllOf) SetName(v string) {
	o.Name = &v
}

func (o MarketplaceUseCaseAutomationAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MarketplaceUseCaseAutomationAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varMarketplaceUseCaseAutomationAllOf := _MarketplaceUseCaseAutomationAllOf{}

	if err = json.Unmarshal(bytes, &varMarketplaceUseCaseAutomationAllOf); err == nil {
		*o = MarketplaceUseCaseAutomationAllOf(varMarketplaceUseCaseAutomationAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "Name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMarketplaceUseCaseAutomationAllOf struct {
	value *MarketplaceUseCaseAutomationAllOf
	isSet bool
}

func (v NullableMarketplaceUseCaseAutomationAllOf) Get() *MarketplaceUseCaseAutomationAllOf {
	return v.value
}

func (v *NullableMarketplaceUseCaseAutomationAllOf) Set(val *MarketplaceUseCaseAutomationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMarketplaceUseCaseAutomationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMarketplaceUseCaseAutomationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarketplaceUseCaseAutomationAllOf(val *MarketplaceUseCaseAutomationAllOf) *NullableMarketplaceUseCaseAutomationAllOf {
	return &NullableMarketplaceUseCaseAutomationAllOf{value: val, isSet: true}
}

func (v NullableMarketplaceUseCaseAutomationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarketplaceUseCaseAutomationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
