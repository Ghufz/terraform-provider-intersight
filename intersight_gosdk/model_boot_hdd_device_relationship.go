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
	"fmt"
)

// BootHddDeviceRelationship - A relationship to the 'boot.HddDevice' resource, or the expanded 'boot.HddDevice' resource, or the 'null' value.
type BootHddDeviceRelationship struct {
	BootHddDevice *BootHddDevice
	MoMoRef       *MoMoRef
}

// BootHddDeviceAsBootHddDeviceRelationship is a convenience function that returns BootHddDevice wrapped in BootHddDeviceRelationship
func BootHddDeviceAsBootHddDeviceRelationship(v *BootHddDevice) BootHddDeviceRelationship {
	return BootHddDeviceRelationship{
		BootHddDevice: v,
	}
}

// MoMoRefAsBootHddDeviceRelationship is a convenience function that returns MoMoRef wrapped in BootHddDeviceRelationship
func MoMoRefAsBootHddDeviceRelationship(v *MoMoRef) BootHddDeviceRelationship {
	return BootHddDeviceRelationship{
		MoMoRef: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BootHddDeviceRelationship) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'boot.HddDevice'
	if jsonDict["ClassId"] == "boot.HddDevice" {
		// try to unmarshal JSON data into BootHddDevice
		err = json.Unmarshal(data, &dst.BootHddDevice)
		if err == nil {
			return nil // data stored in dst.BootHddDevice, return on the first match
		} else {
			dst.BootHddDevice = nil
			return fmt.Errorf("Failed to unmarshal BootHddDeviceRelationship as BootHddDevice: %s", err.Error())
		}
	}

	// check if the discriminator value is 'mo.MoRef'
	if jsonDict["ClassId"] == "mo.MoRef" {
		// try to unmarshal JSON data into MoMoRef
		err = json.Unmarshal(data, &dst.MoMoRef)
		if err == nil {
			return nil // data stored in dst.MoMoRef, return on the first match
		} else {
			dst.MoMoRef = nil
			return fmt.Errorf("Failed to unmarshal BootHddDeviceRelationship as MoMoRef: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BootHddDeviceRelationship) MarshalJSON() ([]byte, error) {
	if src.BootHddDevice != nil {
		return json.Marshal(&src.BootHddDevice)
	}

	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *BootHddDeviceRelationship) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.BootHddDevice != nil {
		return obj.BootHddDevice
	}

	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	// all schemas are nil
	return nil
}

type NullableBootHddDeviceRelationship struct {
	value *BootHddDeviceRelationship
	isSet bool
}

func (v NullableBootHddDeviceRelationship) Get() *BootHddDeviceRelationship {
	return v.value
}

func (v *NullableBootHddDeviceRelationship) Set(val *BootHddDeviceRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableBootHddDeviceRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableBootHddDeviceRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBootHddDeviceRelationship(val *BootHddDeviceRelationship) *NullableBootHddDeviceRelationship {
	return &NullableBootHddDeviceRelationship{value: val, isSet: true}
}

func (v NullableBootHddDeviceRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBootHddDeviceRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
