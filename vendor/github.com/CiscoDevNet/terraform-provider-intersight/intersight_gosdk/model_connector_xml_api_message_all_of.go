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

// ConnectorXmlApiMessageAllOf Definition of the list of properties defined in 'connector.XmlApiMessage', excluding properties defined in parent classes.
type ConnectorXmlApiMessageAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Flag to disable authentication bypassing. If set to true it is expected a valid cookie/login is provided within the XML API request body.
	WithAuth *bool `json:"WithAuth,omitempty"`
	// The XML request body to proxy to the platform.
	XmlRequest           *string `json:"XmlRequest,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConnectorXmlApiMessageAllOf ConnectorXmlApiMessageAllOf

// NewConnectorXmlApiMessageAllOf instantiates a new ConnectorXmlApiMessageAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorXmlApiMessageAllOf(classId string, objectType string) *ConnectorXmlApiMessageAllOf {
	this := ConnectorXmlApiMessageAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewConnectorXmlApiMessageAllOfWithDefaults instantiates a new ConnectorXmlApiMessageAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorXmlApiMessageAllOfWithDefaults() *ConnectorXmlApiMessageAllOf {
	this := ConnectorXmlApiMessageAllOf{}
	var classId string = "connector.XmlApiMessage"
	this.ClassId = classId
	var objectType string = "connector.XmlApiMessage"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ConnectorXmlApiMessageAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ConnectorXmlApiMessageAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ConnectorXmlApiMessageAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ConnectorXmlApiMessageAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ConnectorXmlApiMessageAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ConnectorXmlApiMessageAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetWithAuth returns the WithAuth field value if set, zero value otherwise.
func (o *ConnectorXmlApiMessageAllOf) GetWithAuth() bool {
	if o == nil || o.WithAuth == nil {
		var ret bool
		return ret
	}
	return *o.WithAuth
}

// GetWithAuthOk returns a tuple with the WithAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorXmlApiMessageAllOf) GetWithAuthOk() (*bool, bool) {
	if o == nil || o.WithAuth == nil {
		return nil, false
	}
	return o.WithAuth, true
}

// HasWithAuth returns a boolean if a field has been set.
func (o *ConnectorXmlApiMessageAllOf) HasWithAuth() bool {
	if o != nil && o.WithAuth != nil {
		return true
	}

	return false
}

// SetWithAuth gets a reference to the given bool and assigns it to the WithAuth field.
func (o *ConnectorXmlApiMessageAllOf) SetWithAuth(v bool) {
	o.WithAuth = &v
}

// GetXmlRequest returns the XmlRequest field value if set, zero value otherwise.
func (o *ConnectorXmlApiMessageAllOf) GetXmlRequest() string {
	if o == nil || o.XmlRequest == nil {
		var ret string
		return ret
	}
	return *o.XmlRequest
}

// GetXmlRequestOk returns a tuple with the XmlRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorXmlApiMessageAllOf) GetXmlRequestOk() (*string, bool) {
	if o == nil || o.XmlRequest == nil {
		return nil, false
	}
	return o.XmlRequest, true
}

// HasXmlRequest returns a boolean if a field has been set.
func (o *ConnectorXmlApiMessageAllOf) HasXmlRequest() bool {
	if o != nil && o.XmlRequest != nil {
		return true
	}

	return false
}

// SetXmlRequest gets a reference to the given string and assigns it to the XmlRequest field.
func (o *ConnectorXmlApiMessageAllOf) SetXmlRequest(v string) {
	o.XmlRequest = &v
}

func (o ConnectorXmlApiMessageAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.WithAuth != nil {
		toSerialize["WithAuth"] = o.WithAuth
	}
	if o.XmlRequest != nil {
		toSerialize["XmlRequest"] = o.XmlRequest
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ConnectorXmlApiMessageAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varConnectorXmlApiMessageAllOf := _ConnectorXmlApiMessageAllOf{}

	if err = json.Unmarshal(bytes, &varConnectorXmlApiMessageAllOf); err == nil {
		*o = ConnectorXmlApiMessageAllOf(varConnectorXmlApiMessageAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "WithAuth")
		delete(additionalProperties, "XmlRequest")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConnectorXmlApiMessageAllOf struct {
	value *ConnectorXmlApiMessageAllOf
	isSet bool
}

func (v NullableConnectorXmlApiMessageAllOf) Get() *ConnectorXmlApiMessageAllOf {
	return v.value
}

func (v *NullableConnectorXmlApiMessageAllOf) Set(val *ConnectorXmlApiMessageAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorXmlApiMessageAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorXmlApiMessageAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorXmlApiMessageAllOf(val *ConnectorXmlApiMessageAllOf) *NullableConnectorXmlApiMessageAllOf {
	return &NullableConnectorXmlApiMessageAllOf{value: val, isSet: true}
}

func (v NullableConnectorXmlApiMessageAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorXmlApiMessageAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
