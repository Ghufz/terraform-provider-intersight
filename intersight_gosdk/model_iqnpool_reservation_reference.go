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
	"reflect"
	"strings"
)

// IqnpoolReservationReference The reference to the reservation object.
type IqnpoolReservationReference struct {
	PoolReservationReference
	AdditionalProperties map[string]interface{}
}

type _IqnpoolReservationReference IqnpoolReservationReference

// NewIqnpoolReservationReference instantiates a new IqnpoolReservationReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIqnpoolReservationReference(classId string, objectType string) *IqnpoolReservationReference {
	this := IqnpoolReservationReference{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIqnpoolReservationReferenceWithDefaults instantiates a new IqnpoolReservationReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIqnpoolReservationReferenceWithDefaults() *IqnpoolReservationReference {
	this := IqnpoolReservationReference{}
	return &this
}

func (o IqnpoolReservationReference) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPoolReservationReference, errPoolReservationReference := json.Marshal(o.PoolReservationReference)
	if errPoolReservationReference != nil {
		return []byte{}, errPoolReservationReference
	}
	errPoolReservationReference = json.Unmarshal([]byte(serializedPoolReservationReference), &toSerialize)
	if errPoolReservationReference != nil {
		return []byte{}, errPoolReservationReference
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IqnpoolReservationReference) UnmarshalJSON(bytes []byte) (err error) {
	type IqnpoolReservationReferenceWithoutEmbeddedStruct struct {
	}

	varIqnpoolReservationReferenceWithoutEmbeddedStruct := IqnpoolReservationReferenceWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varIqnpoolReservationReferenceWithoutEmbeddedStruct)
	if err == nil {
		varIqnpoolReservationReference := _IqnpoolReservationReference{}
		*o = IqnpoolReservationReference(varIqnpoolReservationReference)
	} else {
		return err
	}

	varIqnpoolReservationReference := _IqnpoolReservationReference{}

	err = json.Unmarshal(bytes, &varIqnpoolReservationReference)
	if err == nil {
		o.PoolReservationReference = varIqnpoolReservationReference.PoolReservationReference
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {

		// remove fields from embedded structs
		reflectPoolReservationReference := reflect.ValueOf(o.PoolReservationReference)
		for i := 0; i < reflectPoolReservationReference.Type().NumField(); i++ {
			t := reflectPoolReservationReference.Type().Field(i)

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

type NullableIqnpoolReservationReference struct {
	value *IqnpoolReservationReference
	isSet bool
}

func (v NullableIqnpoolReservationReference) Get() *IqnpoolReservationReference {
	return v.value
}

func (v *NullableIqnpoolReservationReference) Set(val *IqnpoolReservationReference) {
	v.value = val
	v.isSet = true
}

func (v NullableIqnpoolReservationReference) IsSet() bool {
	return v.isSet
}

func (v *NullableIqnpoolReservationReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIqnpoolReservationReference(val *IqnpoolReservationReference) *NullableIqnpoolReservationReference {
	return &NullableIqnpoolReservationReference{value: val, isSet: true}
}

func (v NullableIqnpoolReservationReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIqnpoolReservationReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
