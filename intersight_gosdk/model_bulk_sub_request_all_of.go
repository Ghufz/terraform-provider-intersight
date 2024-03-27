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

// BulkSubRequestAllOf Definition of the list of properties defined in 'bulk.SubRequest', excluding properties defined in parent classes.
type BulkSubRequestAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// The URI on which this action is to be performed.
	Uri *string `json:"Uri,omitempty"`
	// The type of operation to be performed. One of - Post (Create), Patch (Update) or Delete (Remove). The value is used to override the top level verb. * `POST` - Used to create a REST resource. * `PATCH` - Used to update a REST resource. * `DELETE` - Used to delete a REST resource.
	Verb                 *string `json:"Verb,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BulkSubRequestAllOf BulkSubRequestAllOf

// NewBulkSubRequestAllOf instantiates a new BulkSubRequestAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkSubRequestAllOf(classId string, objectType string) *BulkSubRequestAllOf {
	this := BulkSubRequestAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var verb string = "POST"
	this.Verb = &verb
	return &this
}

// NewBulkSubRequestAllOfWithDefaults instantiates a new BulkSubRequestAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkSubRequestAllOfWithDefaults() *BulkSubRequestAllOf {
	this := BulkSubRequestAllOf{}
	var classId string = "bulk.RestSubRequest"
	this.ClassId = classId
	var objectType string = "bulk.RestSubRequest"
	this.ObjectType = objectType
	var verb string = "POST"
	this.Verb = &verb
	return &this
}

// GetClassId returns the ClassId field value
func (o *BulkSubRequestAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *BulkSubRequestAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *BulkSubRequestAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *BulkSubRequestAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *BulkSubRequestAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *BulkSubRequestAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *BulkSubRequestAllOf) GetUri() string {
	if o == nil || o.Uri == nil {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkSubRequestAllOf) GetUriOk() (*string, bool) {
	if o == nil || o.Uri == nil {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *BulkSubRequestAllOf) HasUri() bool {
	if o != nil && o.Uri != nil {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *BulkSubRequestAllOf) SetUri(v string) {
	o.Uri = &v
}

// GetVerb returns the Verb field value if set, zero value otherwise.
func (o *BulkSubRequestAllOf) GetVerb() string {
	if o == nil || o.Verb == nil {
		var ret string
		return ret
	}
	return *o.Verb
}

// GetVerbOk returns a tuple with the Verb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkSubRequestAllOf) GetVerbOk() (*string, bool) {
	if o == nil || o.Verb == nil {
		return nil, false
	}
	return o.Verb, true
}

// HasVerb returns a boolean if a field has been set.
func (o *BulkSubRequestAllOf) HasVerb() bool {
	if o != nil && o.Verb != nil {
		return true
	}

	return false
}

// SetVerb gets a reference to the given string and assigns it to the Verb field.
func (o *BulkSubRequestAllOf) SetVerb(v string) {
	o.Verb = &v
}

func (o BulkSubRequestAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Uri != nil {
		toSerialize["Uri"] = o.Uri
	}
	if o.Verb != nil {
		toSerialize["Verb"] = o.Verb
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BulkSubRequestAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varBulkSubRequestAllOf := _BulkSubRequestAllOf{}

	if err = json.Unmarshal(bytes, &varBulkSubRequestAllOf); err == nil {
		*o = BulkSubRequestAllOf(varBulkSubRequestAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Uri")
		delete(additionalProperties, "Verb")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBulkSubRequestAllOf struct {
	value *BulkSubRequestAllOf
	isSet bool
}

func (v NullableBulkSubRequestAllOf) Get() *BulkSubRequestAllOf {
	return v.value
}

func (v *NullableBulkSubRequestAllOf) Set(val *BulkSubRequestAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkSubRequestAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkSubRequestAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkSubRequestAllOf(val *BulkSubRequestAllOf) *NullableBulkSubRequestAllOf {
	return &NullableBulkSubRequestAllOf{value: val, isSet: true}
}

func (v NullableBulkSubRequestAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkSubRequestAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
