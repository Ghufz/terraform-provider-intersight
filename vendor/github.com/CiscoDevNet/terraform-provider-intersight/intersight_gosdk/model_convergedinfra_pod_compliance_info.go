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
	"reflect"
	"strings"
)

// ConvergedinfraPodComplianceInfo Captures the compliance information for a converged infrastructure pod.
type ConvergedinfraPodComplianceInfo struct {
	ConvergedinfraBaseComplianceDetails
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// An array of relationships to convergedinfraBaseComplianceDetails resources.
	Details              []ConvergedinfraBaseComplianceDetailsRelationship `json:"Details,omitempty"`
	Pod                  *ConvergedinfraPodRelationship                    `json:"Pod,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConvergedinfraPodComplianceInfo ConvergedinfraPodComplianceInfo

// NewConvergedinfraPodComplianceInfo instantiates a new ConvergedinfraPodComplianceInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConvergedinfraPodComplianceInfo(classId string, objectType string) *ConvergedinfraPodComplianceInfo {
	this := ConvergedinfraPodComplianceInfo{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewConvergedinfraPodComplianceInfoWithDefaults instantiates a new ConvergedinfraPodComplianceInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConvergedinfraPodComplianceInfoWithDefaults() *ConvergedinfraPodComplianceInfo {
	this := ConvergedinfraPodComplianceInfo{}
	var classId string = "convergedinfra.PodComplianceInfo"
	this.ClassId = classId
	var objectType string = "convergedinfra.PodComplianceInfo"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ConvergedinfraPodComplianceInfo) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ConvergedinfraPodComplianceInfo) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ConvergedinfraPodComplianceInfo) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ConvergedinfraPodComplianceInfo) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ConvergedinfraPodComplianceInfo) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ConvergedinfraPodComplianceInfo) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDetails returns the Details field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConvergedinfraPodComplianceInfo) GetDetails() []ConvergedinfraBaseComplianceDetailsRelationship {
	if o == nil {
		var ret []ConvergedinfraBaseComplianceDetailsRelationship
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConvergedinfraPodComplianceInfo) GetDetailsOk() ([]ConvergedinfraBaseComplianceDetailsRelationship, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *ConvergedinfraPodComplianceInfo) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given []ConvergedinfraBaseComplianceDetailsRelationship and assigns it to the Details field.
func (o *ConvergedinfraPodComplianceInfo) SetDetails(v []ConvergedinfraBaseComplianceDetailsRelationship) {
	o.Details = v
}

// GetPod returns the Pod field value if set, zero value otherwise.
func (o *ConvergedinfraPodComplianceInfo) GetPod() ConvergedinfraPodRelationship {
	if o == nil || o.Pod == nil {
		var ret ConvergedinfraPodRelationship
		return ret
	}
	return *o.Pod
}

// GetPodOk returns a tuple with the Pod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraPodComplianceInfo) GetPodOk() (*ConvergedinfraPodRelationship, bool) {
	if o == nil || o.Pod == nil {
		return nil, false
	}
	return o.Pod, true
}

// HasPod returns a boolean if a field has been set.
func (o *ConvergedinfraPodComplianceInfo) HasPod() bool {
	if o != nil && o.Pod != nil {
		return true
	}

	return false
}

// SetPod gets a reference to the given ConvergedinfraPodRelationship and assigns it to the Pod field.
func (o *ConvergedinfraPodComplianceInfo) SetPod(v ConvergedinfraPodRelationship) {
	o.Pod = &v
}

func (o ConvergedinfraPodComplianceInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedConvergedinfraBaseComplianceDetails, errConvergedinfraBaseComplianceDetails := json.Marshal(o.ConvergedinfraBaseComplianceDetails)
	if errConvergedinfraBaseComplianceDetails != nil {
		return []byte{}, errConvergedinfraBaseComplianceDetails
	}
	errConvergedinfraBaseComplianceDetails = json.Unmarshal([]byte(serializedConvergedinfraBaseComplianceDetails), &toSerialize)
	if errConvergedinfraBaseComplianceDetails != nil {
		return []byte{}, errConvergedinfraBaseComplianceDetails
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Details != nil {
		toSerialize["Details"] = o.Details
	}
	if o.Pod != nil {
		toSerialize["Pod"] = o.Pod
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ConvergedinfraPodComplianceInfo) UnmarshalJSON(bytes []byte) (err error) {
	type ConvergedinfraPodComplianceInfoWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// An array of relationships to convergedinfraBaseComplianceDetails resources.
		Details []ConvergedinfraBaseComplianceDetailsRelationship `json:"Details,omitempty"`
		Pod     *ConvergedinfraPodRelationship                    `json:"Pod,omitempty"`
	}

	varConvergedinfraPodComplianceInfoWithoutEmbeddedStruct := ConvergedinfraPodComplianceInfoWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varConvergedinfraPodComplianceInfoWithoutEmbeddedStruct)
	if err == nil {
		varConvergedinfraPodComplianceInfo := _ConvergedinfraPodComplianceInfo{}
		varConvergedinfraPodComplianceInfo.ClassId = varConvergedinfraPodComplianceInfoWithoutEmbeddedStruct.ClassId
		varConvergedinfraPodComplianceInfo.ObjectType = varConvergedinfraPodComplianceInfoWithoutEmbeddedStruct.ObjectType
		varConvergedinfraPodComplianceInfo.Details = varConvergedinfraPodComplianceInfoWithoutEmbeddedStruct.Details
		varConvergedinfraPodComplianceInfo.Pod = varConvergedinfraPodComplianceInfoWithoutEmbeddedStruct.Pod
		*o = ConvergedinfraPodComplianceInfo(varConvergedinfraPodComplianceInfo)
	} else {
		return err
	}

	varConvergedinfraPodComplianceInfo := _ConvergedinfraPodComplianceInfo{}

	err = json.Unmarshal(bytes, &varConvergedinfraPodComplianceInfo)
	if err == nil {
		o.ConvergedinfraBaseComplianceDetails = varConvergedinfraPodComplianceInfo.ConvergedinfraBaseComplianceDetails
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Details")
		delete(additionalProperties, "Pod")

		// remove fields from embedded structs
		reflectConvergedinfraBaseComplianceDetails := reflect.ValueOf(o.ConvergedinfraBaseComplianceDetails)
		for i := 0; i < reflectConvergedinfraBaseComplianceDetails.Type().NumField(); i++ {
			t := reflectConvergedinfraBaseComplianceDetails.Type().Field(i)

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

type NullableConvergedinfraPodComplianceInfo struct {
	value *ConvergedinfraPodComplianceInfo
	isSet bool
}

func (v NullableConvergedinfraPodComplianceInfo) Get() *ConvergedinfraPodComplianceInfo {
	return v.value
}

func (v *NullableConvergedinfraPodComplianceInfo) Set(val *ConvergedinfraPodComplianceInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableConvergedinfraPodComplianceInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableConvergedinfraPodComplianceInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConvergedinfraPodComplianceInfo(val *ConvergedinfraPodComplianceInfo) *NullableConvergedinfraPodComplianceInfo {
	return &NullableConvergedinfraPodComplianceInfo{value: val, isSet: true}
}

func (v NullableConvergedinfraPodComplianceInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConvergedinfraPodComplianceInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
