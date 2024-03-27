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

// MarketplaceUseCaseLocaleAllOf Definition of the list of properties defined in 'marketplace.UseCaseLocale', excluding properties defined in parent classes.
type MarketplaceUseCaseLocaleAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType  string                         `json:"ObjectType"`
	Automations []MarketplaceUseCaseAutomation `json:"Automations,omitempty"`
	// The string field to hold the contents value
	Contents *string `json:"Contents,omitempty"`
	// The string field to hold the description value
	Description *string `json:"Description,omitempty"`
	// A base64-encoded image for the use case
	Icon *string `json:"Icon,omitempty"`
	// The string field to hold the locale
	Locale *string `json:"Locale,omitempty"`
	// The string field to hold the name value
	Name *string `json:"Name,omitempty"`
	// The string field to hold the summary value
	Summary              *string `json:"Summary,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MarketplaceUseCaseLocaleAllOf MarketplaceUseCaseLocaleAllOf

// NewMarketplaceUseCaseLocaleAllOf instantiates a new MarketplaceUseCaseLocaleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarketplaceUseCaseLocaleAllOf(classId string, objectType string) *MarketplaceUseCaseLocaleAllOf {
	this := MarketplaceUseCaseLocaleAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewMarketplaceUseCaseLocaleAllOfWithDefaults instantiates a new MarketplaceUseCaseLocaleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarketplaceUseCaseLocaleAllOfWithDefaults() *MarketplaceUseCaseLocaleAllOf {
	this := MarketplaceUseCaseLocaleAllOf{}
	var classId string = "marketplace.UseCaseLocale"
	this.ClassId = classId
	var objectType string = "marketplace.UseCaseLocale"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *MarketplaceUseCaseLocaleAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *MarketplaceUseCaseLocaleAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *MarketplaceUseCaseLocaleAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *MarketplaceUseCaseLocaleAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *MarketplaceUseCaseLocaleAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *MarketplaceUseCaseLocaleAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAutomations returns the Automations field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MarketplaceUseCaseLocaleAllOf) GetAutomations() []MarketplaceUseCaseAutomation {
	if o == nil {
		var ret []MarketplaceUseCaseAutomation
		return ret
	}
	return o.Automations
}

// GetAutomationsOk returns a tuple with the Automations field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MarketplaceUseCaseLocaleAllOf) GetAutomationsOk() ([]MarketplaceUseCaseAutomation, bool) {
	if o == nil || o.Automations == nil {
		return nil, false
	}
	return o.Automations, true
}

// HasAutomations returns a boolean if a field has been set.
func (o *MarketplaceUseCaseLocaleAllOf) HasAutomations() bool {
	if o != nil && o.Automations != nil {
		return true
	}

	return false
}

// SetAutomations gets a reference to the given []MarketplaceUseCaseAutomation and assigns it to the Automations field.
func (o *MarketplaceUseCaseLocaleAllOf) SetAutomations(v []MarketplaceUseCaseAutomation) {
	o.Automations = v
}

// GetContents returns the Contents field value if set, zero value otherwise.
func (o *MarketplaceUseCaseLocaleAllOf) GetContents() string {
	if o == nil || o.Contents == nil {
		var ret string
		return ret
	}
	return *o.Contents
}

// GetContentsOk returns a tuple with the Contents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketplaceUseCaseLocaleAllOf) GetContentsOk() (*string, bool) {
	if o == nil || o.Contents == nil {
		return nil, false
	}
	return o.Contents, true
}

// HasContents returns a boolean if a field has been set.
func (o *MarketplaceUseCaseLocaleAllOf) HasContents() bool {
	if o != nil && o.Contents != nil {
		return true
	}

	return false
}

// SetContents gets a reference to the given string and assigns it to the Contents field.
func (o *MarketplaceUseCaseLocaleAllOf) SetContents(v string) {
	o.Contents = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *MarketplaceUseCaseLocaleAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketplaceUseCaseLocaleAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MarketplaceUseCaseLocaleAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MarketplaceUseCaseLocaleAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *MarketplaceUseCaseLocaleAllOf) GetIcon() string {
	if o == nil || o.Icon == nil {
		var ret string
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketplaceUseCaseLocaleAllOf) GetIconOk() (*string, bool) {
	if o == nil || o.Icon == nil {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *MarketplaceUseCaseLocaleAllOf) HasIcon() bool {
	if o != nil && o.Icon != nil {
		return true
	}

	return false
}

// SetIcon gets a reference to the given string and assigns it to the Icon field.
func (o *MarketplaceUseCaseLocaleAllOf) SetIcon(v string) {
	o.Icon = &v
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *MarketplaceUseCaseLocaleAllOf) GetLocale() string {
	if o == nil || o.Locale == nil {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketplaceUseCaseLocaleAllOf) GetLocaleOk() (*string, bool) {
	if o == nil || o.Locale == nil {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *MarketplaceUseCaseLocaleAllOf) HasLocale() bool {
	if o != nil && o.Locale != nil {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *MarketplaceUseCaseLocaleAllOf) SetLocale(v string) {
	o.Locale = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MarketplaceUseCaseLocaleAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketplaceUseCaseLocaleAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MarketplaceUseCaseLocaleAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MarketplaceUseCaseLocaleAllOf) SetName(v string) {
	o.Name = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *MarketplaceUseCaseLocaleAllOf) GetSummary() string {
	if o == nil || o.Summary == nil {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketplaceUseCaseLocaleAllOf) GetSummaryOk() (*string, bool) {
	if o == nil || o.Summary == nil {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *MarketplaceUseCaseLocaleAllOf) HasSummary() bool {
	if o != nil && o.Summary != nil {
		return true
	}

	return false
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *MarketplaceUseCaseLocaleAllOf) SetSummary(v string) {
	o.Summary = &v
}

func (o MarketplaceUseCaseLocaleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Automations != nil {
		toSerialize["Automations"] = o.Automations
	}
	if o.Contents != nil {
		toSerialize["Contents"] = o.Contents
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.Icon != nil {
		toSerialize["Icon"] = o.Icon
	}
	if o.Locale != nil {
		toSerialize["Locale"] = o.Locale
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Summary != nil {
		toSerialize["Summary"] = o.Summary
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MarketplaceUseCaseLocaleAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varMarketplaceUseCaseLocaleAllOf := _MarketplaceUseCaseLocaleAllOf{}

	if err = json.Unmarshal(bytes, &varMarketplaceUseCaseLocaleAllOf); err == nil {
		*o = MarketplaceUseCaseLocaleAllOf(varMarketplaceUseCaseLocaleAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Automations")
		delete(additionalProperties, "Contents")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "Icon")
		delete(additionalProperties, "Locale")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Summary")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMarketplaceUseCaseLocaleAllOf struct {
	value *MarketplaceUseCaseLocaleAllOf
	isSet bool
}

func (v NullableMarketplaceUseCaseLocaleAllOf) Get() *MarketplaceUseCaseLocaleAllOf {
	return v.value
}

func (v *NullableMarketplaceUseCaseLocaleAllOf) Set(val *MarketplaceUseCaseLocaleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMarketplaceUseCaseLocaleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMarketplaceUseCaseLocaleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarketplaceUseCaseLocaleAllOf(val *MarketplaceUseCaseLocaleAllOf) *NullableMarketplaceUseCaseLocaleAllOf {
	return &NullableMarketplaceUseCaseLocaleAllOf{value: val, isSet: true}
}

func (v NullableMarketplaceUseCaseLocaleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarketplaceUseCaseLocaleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
