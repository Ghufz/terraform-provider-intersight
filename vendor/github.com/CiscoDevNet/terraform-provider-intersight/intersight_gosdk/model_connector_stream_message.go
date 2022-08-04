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
	"reflect"
	"strings"
)

// ConnectorStreamMessage The base message to be sent to the stream plugin. Carries information to route the message to the appropriate running stream or create a new stream.
type ConnectorStreamMessage struct {
	ConnectorBaseMessage
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// The requested stream name. Stream names are unique per device endpoint.
	StreamName           *string `json:"StreamName,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConnectorStreamMessage ConnectorStreamMessage

// NewConnectorStreamMessage instantiates a new ConnectorStreamMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorStreamMessage(classId string, objectType string) *ConnectorStreamMessage {
	this := ConnectorStreamMessage{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewConnectorStreamMessageWithDefaults instantiates a new ConnectorStreamMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorStreamMessageWithDefaults() *ConnectorStreamMessage {
	this := ConnectorStreamMessage{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *ConnectorStreamMessage) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ConnectorStreamMessage) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ConnectorStreamMessage) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ConnectorStreamMessage) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ConnectorStreamMessage) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ConnectorStreamMessage) SetObjectType(v string) {
	o.ObjectType = v
}

// GetStreamName returns the StreamName field value if set, zero value otherwise.
func (o *ConnectorStreamMessage) GetStreamName() string {
	if o == nil || o.StreamName == nil {
		var ret string
		return ret
	}
	return *o.StreamName
}

// GetStreamNameOk returns a tuple with the StreamName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorStreamMessage) GetStreamNameOk() (*string, bool) {
	if o == nil || o.StreamName == nil {
		return nil, false
	}
	return o.StreamName, true
}

// HasStreamName returns a boolean if a field has been set.
func (o *ConnectorStreamMessage) HasStreamName() bool {
	if o != nil && o.StreamName != nil {
		return true
	}

	return false
}

// SetStreamName gets a reference to the given string and assigns it to the StreamName field.
func (o *ConnectorStreamMessage) SetStreamName(v string) {
	o.StreamName = &v
}

func (o ConnectorStreamMessage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedConnectorBaseMessage, errConnectorBaseMessage := json.Marshal(o.ConnectorBaseMessage)
	if errConnectorBaseMessage != nil {
		return []byte{}, errConnectorBaseMessage
	}
	errConnectorBaseMessage = json.Unmarshal([]byte(serializedConnectorBaseMessage), &toSerialize)
	if errConnectorBaseMessage != nil {
		return []byte{}, errConnectorBaseMessage
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.StreamName != nil {
		toSerialize["StreamName"] = o.StreamName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ConnectorStreamMessage) UnmarshalJSON(bytes []byte) (err error) {
	type ConnectorStreamMessageWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// The requested stream name. Stream names are unique per device endpoint.
		StreamName *string `json:"StreamName,omitempty"`
	}

	varConnectorStreamMessageWithoutEmbeddedStruct := ConnectorStreamMessageWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varConnectorStreamMessageWithoutEmbeddedStruct)
	if err == nil {
		varConnectorStreamMessage := _ConnectorStreamMessage{}
		varConnectorStreamMessage.ClassId = varConnectorStreamMessageWithoutEmbeddedStruct.ClassId
		varConnectorStreamMessage.ObjectType = varConnectorStreamMessageWithoutEmbeddedStruct.ObjectType
		varConnectorStreamMessage.StreamName = varConnectorStreamMessageWithoutEmbeddedStruct.StreamName
		*o = ConnectorStreamMessage(varConnectorStreamMessage)
	} else {
		return err
	}

	varConnectorStreamMessage := _ConnectorStreamMessage{}

	err = json.Unmarshal(bytes, &varConnectorStreamMessage)
	if err == nil {
		o.ConnectorBaseMessage = varConnectorStreamMessage.ConnectorBaseMessage
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "StreamName")

		// remove fields from embedded structs
		reflectConnectorBaseMessage := reflect.ValueOf(o.ConnectorBaseMessage)
		for i := 0; i < reflectConnectorBaseMessage.Type().NumField(); i++ {
			t := reflectConnectorBaseMessage.Type().Field(i)

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

type NullableConnectorStreamMessage struct {
	value *ConnectorStreamMessage
	isSet bool
}

func (v NullableConnectorStreamMessage) Get() *ConnectorStreamMessage {
	return v.value
}

func (v *NullableConnectorStreamMessage) Set(val *ConnectorStreamMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorStreamMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorStreamMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorStreamMessage(val *ConnectorStreamMessage) *NullableConnectorStreamMessage {
	return &NullableConnectorStreamMessage{value: val, isSet: true}
}

func (v NullableConnectorStreamMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorStreamMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
