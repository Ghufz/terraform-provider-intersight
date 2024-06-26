/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-16342
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// MarketplaceUseCaseVersionResource A MO describing the resources that belong to a UseCaseVersion.
type MarketplaceUseCaseVersionResource struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// A string ID for each use case.
	ResourceId *string `json:"ResourceId,omitempty"`
	// A string resource type for each use case.
	ResourceType         *string `json:"ResourceType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MarketplaceUseCaseVersionResource MarketplaceUseCaseVersionResource

// NewMarketplaceUseCaseVersionResource instantiates a new MarketplaceUseCaseVersionResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarketplaceUseCaseVersionResource(classId string, objectType string) *MarketplaceUseCaseVersionResource {
	this := MarketplaceUseCaseVersionResource{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewMarketplaceUseCaseVersionResourceWithDefaults instantiates a new MarketplaceUseCaseVersionResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarketplaceUseCaseVersionResourceWithDefaults() *MarketplaceUseCaseVersionResource {
	this := MarketplaceUseCaseVersionResource{}
	var classId string = "marketplace.UseCaseVersionResource"
	this.ClassId = classId
	var objectType string = "marketplace.UseCaseVersionResource"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *MarketplaceUseCaseVersionResource) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *MarketplaceUseCaseVersionResource) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *MarketplaceUseCaseVersionResource) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *MarketplaceUseCaseVersionResource) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *MarketplaceUseCaseVersionResource) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *MarketplaceUseCaseVersionResource) SetObjectType(v string) {
	o.ObjectType = v
}

// GetResourceId returns the ResourceId field value if set, zero value otherwise.
func (o *MarketplaceUseCaseVersionResource) GetResourceId() string {
	if o == nil || o.ResourceId == nil {
		var ret string
		return ret
	}
	return *o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketplaceUseCaseVersionResource) GetResourceIdOk() (*string, bool) {
	if o == nil || o.ResourceId == nil {
		return nil, false
	}
	return o.ResourceId, true
}

// HasResourceId returns a boolean if a field has been set.
func (o *MarketplaceUseCaseVersionResource) HasResourceId() bool {
	if o != nil && o.ResourceId != nil {
		return true
	}

	return false
}

// SetResourceId gets a reference to the given string and assigns it to the ResourceId field.
func (o *MarketplaceUseCaseVersionResource) SetResourceId(v string) {
	o.ResourceId = &v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *MarketplaceUseCaseVersionResource) GetResourceType() string {
	if o == nil || o.ResourceType == nil {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketplaceUseCaseVersionResource) GetResourceTypeOk() (*string, bool) {
	if o == nil || o.ResourceType == nil {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *MarketplaceUseCaseVersionResource) HasResourceType() bool {
	if o != nil && o.ResourceType != nil {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *MarketplaceUseCaseVersionResource) SetResourceType(v string) {
	o.ResourceType = &v
}

func (o MarketplaceUseCaseVersionResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ResourceId != nil {
		toSerialize["ResourceId"] = o.ResourceId
	}
	if o.ResourceType != nil {
		toSerialize["ResourceType"] = o.ResourceType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MarketplaceUseCaseVersionResource) UnmarshalJSON(bytes []byte) (err error) {
	type MarketplaceUseCaseVersionResourceWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// A string ID for each use case.
		ResourceId *string `json:"ResourceId,omitempty"`
		// A string resource type for each use case.
		ResourceType *string `json:"ResourceType,omitempty"`
	}

	varMarketplaceUseCaseVersionResourceWithoutEmbeddedStruct := MarketplaceUseCaseVersionResourceWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varMarketplaceUseCaseVersionResourceWithoutEmbeddedStruct)
	if err == nil {
		varMarketplaceUseCaseVersionResource := _MarketplaceUseCaseVersionResource{}
		varMarketplaceUseCaseVersionResource.ClassId = varMarketplaceUseCaseVersionResourceWithoutEmbeddedStruct.ClassId
		varMarketplaceUseCaseVersionResource.ObjectType = varMarketplaceUseCaseVersionResourceWithoutEmbeddedStruct.ObjectType
		varMarketplaceUseCaseVersionResource.ResourceId = varMarketplaceUseCaseVersionResourceWithoutEmbeddedStruct.ResourceId
		varMarketplaceUseCaseVersionResource.ResourceType = varMarketplaceUseCaseVersionResourceWithoutEmbeddedStruct.ResourceType
		*o = MarketplaceUseCaseVersionResource(varMarketplaceUseCaseVersionResource)
	} else {
		return err
	}

	varMarketplaceUseCaseVersionResource := _MarketplaceUseCaseVersionResource{}

	err = json.Unmarshal(bytes, &varMarketplaceUseCaseVersionResource)
	if err == nil {
		o.MoBaseComplexType = varMarketplaceUseCaseVersionResource.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ResourceId")
		delete(additionalProperties, "ResourceType")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

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

type NullableMarketplaceUseCaseVersionResource struct {
	value *MarketplaceUseCaseVersionResource
	isSet bool
}

func (v NullableMarketplaceUseCaseVersionResource) Get() *MarketplaceUseCaseVersionResource {
	return v.value
}

func (v *NullableMarketplaceUseCaseVersionResource) Set(val *MarketplaceUseCaseVersionResource) {
	v.value = val
	v.isSet = true
}

func (v NullableMarketplaceUseCaseVersionResource) IsSet() bool {
	return v.isSet
}

func (v *NullableMarketplaceUseCaseVersionResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarketplaceUseCaseVersionResource(val *MarketplaceUseCaseVersionResource) *NullableMarketplaceUseCaseVersionResource {
	return &NullableMarketplaceUseCaseVersionResource{value: val, isSet: true}
}

func (v NullableMarketplaceUseCaseVersionResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarketplaceUseCaseVersionResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
