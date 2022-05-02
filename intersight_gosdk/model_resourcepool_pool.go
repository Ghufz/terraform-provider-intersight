/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-6484
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// ResourcepoolPool Pool represents a collection of resource. The resource can be any MO which has PoolResource meta enabled. The resource in the pool can be reserved or unreserved by using Lease API, reserved/unreserved resources can be used in the entities like server profiles.
type ResourcepoolPool struct {
	PoolAbstractPool
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The resource management type in the pool, it can be either static or dynamic. * `Static` - The resources in the pool will not be changed until user manually update it. * `Dynamic` - The resources in the pool will be updated dynamically based on the condition.
	PoolType               *string                                    `json:"PoolType,omitempty"`
	ResourcePoolParameters NullableResourcepoolResourcePoolParameters `json:"ResourcePoolParameters,omitempty"`
	// The type of the resource present in the pool, example 'server' its combination of RackUnit and Blade. * `None` - The resource cannot consider for Resource Pool. * `Server` - Resource Pool holds the server kind of resources, example - RackServer, Blade.
	ResourceType         *string                               `json:"ResourceType,omitempty"`
	Selectors            []ResourceSelector                    `json:"Selectors,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResourcepoolPool ResourcepoolPool

// NewResourcepoolPool instantiates a new ResourcepoolPool object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourcepoolPool(classId string, objectType string) *ResourcepoolPool {
	this := ResourcepoolPool{}
	this.ClassId = classId
	this.ObjectType = objectType
	var assignmentOrder string = "sequential"
	this.AssignmentOrder = &assignmentOrder
	var poolType string = "Static"
	this.PoolType = &poolType
	var resourceType string = "None"
	this.ResourceType = &resourceType
	return &this
}

// NewResourcepoolPoolWithDefaults instantiates a new ResourcepoolPool object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourcepoolPoolWithDefaults() *ResourcepoolPool {
	this := ResourcepoolPool{}
	var classId string = "resourcepool.Pool"
	this.ClassId = classId
	var objectType string = "resourcepool.Pool"
	this.ObjectType = objectType
	var poolType string = "Static"
	this.PoolType = &poolType
	var resourceType string = "None"
	this.ResourceType = &resourceType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ResourcepoolPool) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ResourcepoolPool) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ResourcepoolPool) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ResourcepoolPool) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ResourcepoolPool) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ResourcepoolPool) SetObjectType(v string) {
	o.ObjectType = v
}

// GetPoolType returns the PoolType field value if set, zero value otherwise.
func (o *ResourcepoolPool) GetPoolType() string {
	if o == nil || o.PoolType == nil {
		var ret string
		return ret
	}
	return *o.PoolType
}

// GetPoolTypeOk returns a tuple with the PoolType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourcepoolPool) GetPoolTypeOk() (*string, bool) {
	if o == nil || o.PoolType == nil {
		return nil, false
	}
	return o.PoolType, true
}

// HasPoolType returns a boolean if a field has been set.
func (o *ResourcepoolPool) HasPoolType() bool {
	if o != nil && o.PoolType != nil {
		return true
	}

	return false
}

// SetPoolType gets a reference to the given string and assigns it to the PoolType field.
func (o *ResourcepoolPool) SetPoolType(v string) {
	o.PoolType = &v
}

// GetResourcePoolParameters returns the ResourcePoolParameters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourcepoolPool) GetResourcePoolParameters() ResourcepoolResourcePoolParameters {
	if o == nil || o.ResourcePoolParameters.Get() == nil {
		var ret ResourcepoolResourcePoolParameters
		return ret
	}
	return *o.ResourcePoolParameters.Get()
}

// GetResourcePoolParametersOk returns a tuple with the ResourcePoolParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourcepoolPool) GetResourcePoolParametersOk() (*ResourcepoolResourcePoolParameters, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResourcePoolParameters.Get(), o.ResourcePoolParameters.IsSet()
}

// HasResourcePoolParameters returns a boolean if a field has been set.
func (o *ResourcepoolPool) HasResourcePoolParameters() bool {
	if o != nil && o.ResourcePoolParameters.IsSet() {
		return true
	}

	return false
}

// SetResourcePoolParameters gets a reference to the given NullableResourcepoolResourcePoolParameters and assigns it to the ResourcePoolParameters field.
func (o *ResourcepoolPool) SetResourcePoolParameters(v ResourcepoolResourcePoolParameters) {
	o.ResourcePoolParameters.Set(&v)
}

// SetResourcePoolParametersNil sets the value for ResourcePoolParameters to be an explicit nil
func (o *ResourcepoolPool) SetResourcePoolParametersNil() {
	o.ResourcePoolParameters.Set(nil)
}

// UnsetResourcePoolParameters ensures that no value is present for ResourcePoolParameters, not even an explicit nil
func (o *ResourcepoolPool) UnsetResourcePoolParameters() {
	o.ResourcePoolParameters.Unset()
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *ResourcepoolPool) GetResourceType() string {
	if o == nil || o.ResourceType == nil {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourcepoolPool) GetResourceTypeOk() (*string, bool) {
	if o == nil || o.ResourceType == nil {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *ResourcepoolPool) HasResourceType() bool {
	if o != nil && o.ResourceType != nil {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *ResourcepoolPool) SetResourceType(v string) {
	o.ResourceType = &v
}

// GetSelectors returns the Selectors field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourcepoolPool) GetSelectors() []ResourceSelector {
	if o == nil {
		var ret []ResourceSelector
		return ret
	}
	return o.Selectors
}

// GetSelectorsOk returns a tuple with the Selectors field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourcepoolPool) GetSelectorsOk() ([]ResourceSelector, bool) {
	if o == nil || o.Selectors == nil {
		return nil, false
	}
	return o.Selectors, true
}

// HasSelectors returns a boolean if a field has been set.
func (o *ResourcepoolPool) HasSelectors() bool {
	if o != nil && o.Selectors != nil {
		return true
	}

	return false
}

// SetSelectors gets a reference to the given []ResourceSelector and assigns it to the Selectors field.
func (o *ResourcepoolPool) SetSelectors(v []ResourceSelector) {
	o.Selectors = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *ResourcepoolPool) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourcepoolPool) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *ResourcepoolPool) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *ResourcepoolPool) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o ResourcepoolPool) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPoolAbstractPool, errPoolAbstractPool := json.Marshal(o.PoolAbstractPool)
	if errPoolAbstractPool != nil {
		return []byte{}, errPoolAbstractPool
	}
	errPoolAbstractPool = json.Unmarshal([]byte(serializedPoolAbstractPool), &toSerialize)
	if errPoolAbstractPool != nil {
		return []byte{}, errPoolAbstractPool
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.PoolType != nil {
		toSerialize["PoolType"] = o.PoolType
	}
	if o.ResourcePoolParameters.IsSet() {
		toSerialize["ResourcePoolParameters"] = o.ResourcePoolParameters.Get()
	}
	if o.ResourceType != nil {
		toSerialize["ResourceType"] = o.ResourceType
	}
	if o.Selectors != nil {
		toSerialize["Selectors"] = o.Selectors
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ResourcepoolPool) UnmarshalJSON(bytes []byte) (err error) {
	type ResourcepoolPoolWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The resource management type in the pool, it can be either static or dynamic. * `Static` - The resources in the pool will not be changed until user manually update it. * `Dynamic` - The resources in the pool will be updated dynamically based on the condition.
		PoolType               *string                                    `json:"PoolType,omitempty"`
		ResourcePoolParameters NullableResourcepoolResourcePoolParameters `json:"ResourcePoolParameters,omitempty"`
		// The type of the resource present in the pool, example 'server' its combination of RackUnit and Blade. * `None` - The resource cannot consider for Resource Pool. * `Server` - Resource Pool holds the server kind of resources, example - RackServer, Blade.
		ResourceType *string                               `json:"ResourceType,omitempty"`
		Selectors    []ResourceSelector                    `json:"Selectors,omitempty"`
		Organization *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	}

	varResourcepoolPoolWithoutEmbeddedStruct := ResourcepoolPoolWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varResourcepoolPoolWithoutEmbeddedStruct)
	if err == nil {
		varResourcepoolPool := _ResourcepoolPool{}
		varResourcepoolPool.ClassId = varResourcepoolPoolWithoutEmbeddedStruct.ClassId
		varResourcepoolPool.ObjectType = varResourcepoolPoolWithoutEmbeddedStruct.ObjectType
		varResourcepoolPool.PoolType = varResourcepoolPoolWithoutEmbeddedStruct.PoolType
		varResourcepoolPool.ResourcePoolParameters = varResourcepoolPoolWithoutEmbeddedStruct.ResourcePoolParameters
		varResourcepoolPool.ResourceType = varResourcepoolPoolWithoutEmbeddedStruct.ResourceType
		varResourcepoolPool.Selectors = varResourcepoolPoolWithoutEmbeddedStruct.Selectors
		varResourcepoolPool.Organization = varResourcepoolPoolWithoutEmbeddedStruct.Organization
		*o = ResourcepoolPool(varResourcepoolPool)
	} else {
		return err
	}

	varResourcepoolPool := _ResourcepoolPool{}

	err = json.Unmarshal(bytes, &varResourcepoolPool)
	if err == nil {
		o.PoolAbstractPool = varResourcepoolPool.PoolAbstractPool
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "PoolType")
		delete(additionalProperties, "ResourcePoolParameters")
		delete(additionalProperties, "ResourceType")
		delete(additionalProperties, "Selectors")
		delete(additionalProperties, "Organization")

		// remove fields from embedded structs
		reflectPoolAbstractPool := reflect.ValueOf(o.PoolAbstractPool)
		for i := 0; i < reflectPoolAbstractPool.Type().NumField(); i++ {
			t := reflectPoolAbstractPool.Type().Field(i)

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

type NullableResourcepoolPool struct {
	value *ResourcepoolPool
	isSet bool
}

func (v NullableResourcepoolPool) Get() *ResourcepoolPool {
	return v.value
}

func (v *NullableResourcepoolPool) Set(val *ResourcepoolPool) {
	v.value = val
	v.isSet = true
}

func (v NullableResourcepoolPool) IsSet() bool {
	return v.isSet
}

func (v *NullableResourcepoolPool) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourcepoolPool(val *ResourcepoolPool) *NullableResourcepoolPool {
	return &NullableResourcepoolPool{value: val, isSet: true}
}

func (v NullableResourcepoolPool) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourcepoolPool) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
