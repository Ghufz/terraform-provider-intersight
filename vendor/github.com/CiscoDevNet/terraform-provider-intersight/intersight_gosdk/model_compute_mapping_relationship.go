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
	"fmt"
)

// ComputeMappingRelationship - A relationship to the 'compute.Mapping' resource, or the expanded 'compute.Mapping' resource, or the 'null' value.
type ComputeMappingRelationship struct {
	ComputeMapping *ComputeMapping
	MoMoRef        *MoMoRef
}

// ComputeMappingAsComputeMappingRelationship is a convenience function that returns ComputeMapping wrapped in ComputeMappingRelationship
func ComputeMappingAsComputeMappingRelationship(v *ComputeMapping) ComputeMappingRelationship {
	return ComputeMappingRelationship{
		ComputeMapping: v,
	}
}

// MoMoRefAsComputeMappingRelationship is a convenience function that returns MoMoRef wrapped in ComputeMappingRelationship
func MoMoRefAsComputeMappingRelationship(v *MoMoRef) ComputeMappingRelationship {
	return ComputeMappingRelationship{
		MoMoRef: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ComputeMappingRelationship) UnmarshalJSON(data []byte) error {
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

	// check if the discriminator value is 'compute.Mapping'
	if jsonDict["ClassId"] == "compute.Mapping" {
		// try to unmarshal JSON data into ComputeMapping
		err = json.Unmarshal(data, &dst.ComputeMapping)
		if err == nil {
			return nil // data stored in dst.ComputeMapping, return on the first match
		} else {
			dst.ComputeMapping = nil
			return fmt.Errorf("Failed to unmarshal ComputeMappingRelationship as ComputeMapping: %s", err.Error())
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
			return fmt.Errorf("Failed to unmarshal ComputeMappingRelationship as MoMoRef: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ComputeMappingRelationship) MarshalJSON() ([]byte, error) {
	if src.ComputeMapping != nil {
		return json.Marshal(&src.ComputeMapping)
	}

	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ComputeMappingRelationship) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ComputeMapping != nil {
		return obj.ComputeMapping
	}

	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	// all schemas are nil
	return nil
}

type NullableComputeMappingRelationship struct {
	value *ComputeMappingRelationship
	isSet bool
}

func (v NullableComputeMappingRelationship) Get() *ComputeMappingRelationship {
	return v.value
}

func (v *NullableComputeMappingRelationship) Set(val *ComputeMappingRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableComputeMappingRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableComputeMappingRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputeMappingRelationship(val *ComputeMappingRelationship) *NullableComputeMappingRelationship {
	return &NullableComputeMappingRelationship{value: val, isSet: true}
}

func (v NullableComputeMappingRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputeMappingRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
