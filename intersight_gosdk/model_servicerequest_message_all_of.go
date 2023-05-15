/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-11765
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// ServicerequestMessageAllOf Definition of the list of properties defined in 'servicerequest.Message', excluding properties defined in parent classes.
type ServicerequestMessageAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The service item in which operation is perfomed.
	ServiceItemName      *string `json:"ServiceItemName,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServicerequestMessageAllOf ServicerequestMessageAllOf

// NewServicerequestMessageAllOf instantiates a new ServicerequestMessageAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServicerequestMessageAllOf(classId string, objectType string) *ServicerequestMessageAllOf {
	this := ServicerequestMessageAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewServicerequestMessageAllOfWithDefaults instantiates a new ServicerequestMessageAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServicerequestMessageAllOfWithDefaults() *ServicerequestMessageAllOf {
	this := ServicerequestMessageAllOf{}
	var classId string = "servicerequest.Message"
	this.ClassId = classId
	var objectType string = "servicerequest.Message"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ServicerequestMessageAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ServicerequestMessageAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ServicerequestMessageAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ServicerequestMessageAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ServicerequestMessageAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ServicerequestMessageAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetServiceItemName returns the ServiceItemName field value if set, zero value otherwise.
func (o *ServicerequestMessageAllOf) GetServiceItemName() string {
	if o == nil || o.ServiceItemName == nil {
		var ret string
		return ret
	}
	return *o.ServiceItemName
}

// GetServiceItemNameOk returns a tuple with the ServiceItemName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicerequestMessageAllOf) GetServiceItemNameOk() (*string, bool) {
	if o == nil || o.ServiceItemName == nil {
		return nil, false
	}
	return o.ServiceItemName, true
}

// HasServiceItemName returns a boolean if a field has been set.
func (o *ServicerequestMessageAllOf) HasServiceItemName() bool {
	if o != nil && o.ServiceItemName != nil {
		return true
	}

	return false
}

// SetServiceItemName gets a reference to the given string and assigns it to the ServiceItemName field.
func (o *ServicerequestMessageAllOf) SetServiceItemName(v string) {
	o.ServiceItemName = &v
}

func (o ServicerequestMessageAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ServiceItemName != nil {
		toSerialize["ServiceItemName"] = o.ServiceItemName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ServicerequestMessageAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varServicerequestMessageAllOf := _ServicerequestMessageAllOf{}

	if err = json.Unmarshal(bytes, &varServicerequestMessageAllOf); err == nil {
		*o = ServicerequestMessageAllOf(varServicerequestMessageAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ServiceItemName")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServicerequestMessageAllOf struct {
	value *ServicerequestMessageAllOf
	isSet bool
}

func (v NullableServicerequestMessageAllOf) Get() *ServicerequestMessageAllOf {
	return v.value
}

func (v *NullableServicerequestMessageAllOf) Set(val *ServicerequestMessageAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableServicerequestMessageAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableServicerequestMessageAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServicerequestMessageAllOf(val *ServicerequestMessageAllOf) *NullableServicerequestMessageAllOf {
	return &NullableServicerequestMessageAllOf{value: val, isSet: true}
}

func (v NullableServicerequestMessageAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServicerequestMessageAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
