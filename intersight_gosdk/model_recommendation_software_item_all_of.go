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

// RecommendationSoftwareItemAllOf Definition of the list of properties defined in 'recommendation.SoftwareItem', excluding properties defined in parent classes.
type RecommendationSoftwareItemAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The user visible message which informs user of the type of software recommendation.
	Message *string `json:"Message,omitempty"`
	// The software-recommendation type, for example, HXDP version, HyperV or ESXi version, etc. * `None` - The Enum value None represents the default software recommendation value. * `HXDPVersion` - The Enum value HXDPVersion represents that the software recommendation is to upgrade the HyperFlex Data Platform build version. * `NodeRatioLicense` - The Enum value NodeRatioLicense represents that the software recommendation is to upgrade the HyperFlex Data Platform license for using 1:2 converged to compute node ratio limits. * `DCNoFILicense` - The Enum value DCNoFILicense represents that the software recommendation is to upgrade the HyperFlex Data Platform license for using DC-No-FI limits. * `LAZExistingStatus` - The Enum value LAZExistingStatus represents that the software recommendation indicates that the HyperFlex cluster might have LAZ enabled. * `LAZNewStatus` - The Enum value LAZNewStatus represents that the software recommendation is to enable LAZ with expansion on the HyperFlex Cluster. * `EVCStatus` - The Enum value EVCStatus represents that the software recommendation is to enable Enhanced VMotion on the HypeFlex Cluster.
	RecommendationType   *string                                     `json:"RecommendationType,omitempty"`
	ClusterExpansion     *RecommendationClusterExpansionRelationship `json:"ClusterExpansion,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RecommendationSoftwareItemAllOf RecommendationSoftwareItemAllOf

// NewRecommendationSoftwareItemAllOf instantiates a new RecommendationSoftwareItemAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecommendationSoftwareItemAllOf(classId string, objectType string) *RecommendationSoftwareItemAllOf {
	this := RecommendationSoftwareItemAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewRecommendationSoftwareItemAllOfWithDefaults instantiates a new RecommendationSoftwareItemAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecommendationSoftwareItemAllOfWithDefaults() *RecommendationSoftwareItemAllOf {
	this := RecommendationSoftwareItemAllOf{}
	var classId string = "recommendation.SoftwareItem"
	this.ClassId = classId
	var objectType string = "recommendation.SoftwareItem"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *RecommendationSoftwareItemAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *RecommendationSoftwareItemAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *RecommendationSoftwareItemAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *RecommendationSoftwareItemAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *RecommendationSoftwareItemAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *RecommendationSoftwareItemAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *RecommendationSoftwareItemAllOf) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationSoftwareItemAllOf) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *RecommendationSoftwareItemAllOf) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *RecommendationSoftwareItemAllOf) SetMessage(v string) {
	o.Message = &v
}

// GetRecommendationType returns the RecommendationType field value if set, zero value otherwise.
func (o *RecommendationSoftwareItemAllOf) GetRecommendationType() string {
	if o == nil || o.RecommendationType == nil {
		var ret string
		return ret
	}
	return *o.RecommendationType
}

// GetRecommendationTypeOk returns a tuple with the RecommendationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationSoftwareItemAllOf) GetRecommendationTypeOk() (*string, bool) {
	if o == nil || o.RecommendationType == nil {
		return nil, false
	}
	return o.RecommendationType, true
}

// HasRecommendationType returns a boolean if a field has been set.
func (o *RecommendationSoftwareItemAllOf) HasRecommendationType() bool {
	if o != nil && o.RecommendationType != nil {
		return true
	}

	return false
}

// SetRecommendationType gets a reference to the given string and assigns it to the RecommendationType field.
func (o *RecommendationSoftwareItemAllOf) SetRecommendationType(v string) {
	o.RecommendationType = &v
}

// GetClusterExpansion returns the ClusterExpansion field value if set, zero value otherwise.
func (o *RecommendationSoftwareItemAllOf) GetClusterExpansion() RecommendationClusterExpansionRelationship {
	if o == nil || o.ClusterExpansion == nil {
		var ret RecommendationClusterExpansionRelationship
		return ret
	}
	return *o.ClusterExpansion
}

// GetClusterExpansionOk returns a tuple with the ClusterExpansion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationSoftwareItemAllOf) GetClusterExpansionOk() (*RecommendationClusterExpansionRelationship, bool) {
	if o == nil || o.ClusterExpansion == nil {
		return nil, false
	}
	return o.ClusterExpansion, true
}

// HasClusterExpansion returns a boolean if a field has been set.
func (o *RecommendationSoftwareItemAllOf) HasClusterExpansion() bool {
	if o != nil && o.ClusterExpansion != nil {
		return true
	}

	return false
}

// SetClusterExpansion gets a reference to the given RecommendationClusterExpansionRelationship and assigns it to the ClusterExpansion field.
func (o *RecommendationSoftwareItemAllOf) SetClusterExpansion(v RecommendationClusterExpansionRelationship) {
	o.ClusterExpansion = &v
}

func (o RecommendationSoftwareItemAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Message != nil {
		toSerialize["Message"] = o.Message
	}
	if o.RecommendationType != nil {
		toSerialize["RecommendationType"] = o.RecommendationType
	}
	if o.ClusterExpansion != nil {
		toSerialize["ClusterExpansion"] = o.ClusterExpansion
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RecommendationSoftwareItemAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varRecommendationSoftwareItemAllOf := _RecommendationSoftwareItemAllOf{}

	if err = json.Unmarshal(bytes, &varRecommendationSoftwareItemAllOf); err == nil {
		*o = RecommendationSoftwareItemAllOf(varRecommendationSoftwareItemAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Message")
		delete(additionalProperties, "RecommendationType")
		delete(additionalProperties, "ClusterExpansion")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRecommendationSoftwareItemAllOf struct {
	value *RecommendationSoftwareItemAllOf
	isSet bool
}

func (v NullableRecommendationSoftwareItemAllOf) Get() *RecommendationSoftwareItemAllOf {
	return v.value
}

func (v *NullableRecommendationSoftwareItemAllOf) Set(val *RecommendationSoftwareItemAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableRecommendationSoftwareItemAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableRecommendationSoftwareItemAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecommendationSoftwareItemAllOf(val *RecommendationSoftwareItemAllOf) *NullableRecommendationSoftwareItemAllOf {
	return &NullableRecommendationSoftwareItemAllOf{value: val, isSet: true}
}

func (v NullableRecommendationSoftwareItemAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecommendationSoftwareItemAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
