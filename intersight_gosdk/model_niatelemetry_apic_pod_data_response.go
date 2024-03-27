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
	"fmt"
)

// NiatelemetryApicPodDataResponse - The response body of a HTTP GET request for the 'niatelemetry.ApicPodData' resource. The value may be one of the following types. 1. When 'tag' is specified in the URL query, the response schema     is a summary of the tag usage. 1. When '$apply' is specified in the URL query, the response schema     is dynamically-generated schema based on the $apply value. 1. When '$count' is specified in the URL query, the response is     a simple object providing the count of the resources. 1. In all other cases, the response is a list of 'niatelemetry.ApicPodData' resources.
type NiatelemetryApicPodDataResponse struct {
	MoAggregateTransform        *MoAggregateTransform
	MoDocumentCount             *MoDocumentCount
	MoTagSummary                *MoTagSummary
	NiatelemetryApicPodDataList *NiatelemetryApicPodDataList
}

// MoAggregateTransformAsNiatelemetryApicPodDataResponse is a convenience function that returns MoAggregateTransform wrapped in NiatelemetryApicPodDataResponse
func MoAggregateTransformAsNiatelemetryApicPodDataResponse(v *MoAggregateTransform) NiatelemetryApicPodDataResponse {
	return NiatelemetryApicPodDataResponse{
		MoAggregateTransform: v,
	}
}

// MoDocumentCountAsNiatelemetryApicPodDataResponse is a convenience function that returns MoDocumentCount wrapped in NiatelemetryApicPodDataResponse
func MoDocumentCountAsNiatelemetryApicPodDataResponse(v *MoDocumentCount) NiatelemetryApicPodDataResponse {
	return NiatelemetryApicPodDataResponse{
		MoDocumentCount: v,
	}
}

// MoTagSummaryAsNiatelemetryApicPodDataResponse is a convenience function that returns MoTagSummary wrapped in NiatelemetryApicPodDataResponse
func MoTagSummaryAsNiatelemetryApicPodDataResponse(v *MoTagSummary) NiatelemetryApicPodDataResponse {
	return NiatelemetryApicPodDataResponse{
		MoTagSummary: v,
	}
}

// NiatelemetryApicPodDataListAsNiatelemetryApicPodDataResponse is a convenience function that returns NiatelemetryApicPodDataList wrapped in NiatelemetryApicPodDataResponse
func NiatelemetryApicPodDataListAsNiatelemetryApicPodDataResponse(v *NiatelemetryApicPodDataList) NiatelemetryApicPodDataResponse {
	return NiatelemetryApicPodDataResponse{
		NiatelemetryApicPodDataList: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *NiatelemetryApicPodDataResponse) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'mo.AggregateTransform'
	if jsonDict["ObjectType"] == "mo.AggregateTransform" {
		// try to unmarshal JSON data into MoAggregateTransform
		err = json.Unmarshal(data, &dst.MoAggregateTransform)
		if err == nil {
			return nil // data stored in dst.MoAggregateTransform, return on the first match
		} else {
			dst.MoAggregateTransform = nil
			return fmt.Errorf("Failed to unmarshal NiatelemetryApicPodDataResponse as MoAggregateTransform: %s", err.Error())
		}
	}

	// check if the discriminator value is 'mo.DocumentCount'
	if jsonDict["ObjectType"] == "mo.DocumentCount" {
		// try to unmarshal JSON data into MoDocumentCount
		err = json.Unmarshal(data, &dst.MoDocumentCount)
		if err == nil {
			return nil // data stored in dst.MoDocumentCount, return on the first match
		} else {
			dst.MoDocumentCount = nil
			return fmt.Errorf("Failed to unmarshal NiatelemetryApicPodDataResponse as MoDocumentCount: %s", err.Error())
		}
	}

	// check if the discriminator value is 'mo.TagSummary'
	if jsonDict["ObjectType"] == "mo.TagSummary" {
		// try to unmarshal JSON data into MoTagSummary
		err = json.Unmarshal(data, &dst.MoTagSummary)
		if err == nil {
			return nil // data stored in dst.MoTagSummary, return on the first match
		} else {
			dst.MoTagSummary = nil
			return fmt.Errorf("Failed to unmarshal NiatelemetryApicPodDataResponse as MoTagSummary: %s", err.Error())
		}
	}

	// check if the discriminator value is 'niatelemetry.ApicPodData.List'
	if jsonDict["ObjectType"] == "niatelemetry.ApicPodData.List" {
		// try to unmarshal JSON data into NiatelemetryApicPodDataList
		err = json.Unmarshal(data, &dst.NiatelemetryApicPodDataList)
		if err == nil {
			return nil // data stored in dst.NiatelemetryApicPodDataList, return on the first match
		} else {
			dst.NiatelemetryApicPodDataList = nil
			return fmt.Errorf("Failed to unmarshal NiatelemetryApicPodDataResponse as NiatelemetryApicPodDataList: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src NiatelemetryApicPodDataResponse) MarshalJSON() ([]byte, error) {
	if src.MoAggregateTransform != nil {
		return json.Marshal(&src.MoAggregateTransform)
	}

	if src.MoDocumentCount != nil {
		return json.Marshal(&src.MoDocumentCount)
	}

	if src.MoTagSummary != nil {
		return json.Marshal(&src.MoTagSummary)
	}

	if src.NiatelemetryApicPodDataList != nil {
		return json.Marshal(&src.NiatelemetryApicPodDataList)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *NiatelemetryApicPodDataResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.MoAggregateTransform != nil {
		return obj.MoAggregateTransform
	}

	if obj.MoDocumentCount != nil {
		return obj.MoDocumentCount
	}

	if obj.MoTagSummary != nil {
		return obj.MoTagSummary
	}

	if obj.NiatelemetryApicPodDataList != nil {
		return obj.NiatelemetryApicPodDataList
	}

	// all schemas are nil
	return nil
}

type NullableNiatelemetryApicPodDataResponse struct {
	value *NiatelemetryApicPodDataResponse
	isSet bool
}

func (v NullableNiatelemetryApicPodDataResponse) Get() *NiatelemetryApicPodDataResponse {
	return v.value
}

func (v *NullableNiatelemetryApicPodDataResponse) Set(val *NiatelemetryApicPodDataResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryApicPodDataResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryApicPodDataResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryApicPodDataResponse(val *NiatelemetryApicPodDataResponse) *NullableNiatelemetryApicPodDataResponse {
	return &NullableNiatelemetryApicPodDataResponse{value: val, isSet: true}
}

func (v NullableNiatelemetryApicPodDataResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryApicPodDataResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
