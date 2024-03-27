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

// FcpoolPool Pool represents a collection of WWN addresses that can be allocated to VHBAs of a server profile.
type FcpoolPool struct {
	PoolAbstractPool
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string        `json:"ObjectType"`
	IdBlocks   []FcpoolBlock `json:"IdBlocks,omitempty"`
	// Purpose of this WWN pool.
	PoolPurpose *string `json:"PoolPurpose,omitempty"`
	// An array of relationships to fcpoolFcBlock resources.
	BlockHeads   []FcpoolFcBlockRelationship           `json:"BlockHeads,omitempty"`
	Organization *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	// An array of relationships to fcpoolReservation resources.
	Reservations         []FcpoolReservationRelationship `json:"Reservations,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FcpoolPool FcpoolPool

// NewFcpoolPool instantiates a new FcpoolPool object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFcpoolPool(classId string, objectType string) *FcpoolPool {
	this := FcpoolPool{}
	this.ClassId = classId
	this.ObjectType = objectType
	var assignmentOrder string = "sequential"
	this.AssignmentOrder = &assignmentOrder
	return &this
}

// NewFcpoolPoolWithDefaults instantiates a new FcpoolPool object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFcpoolPoolWithDefaults() *FcpoolPool {
	this := FcpoolPool{}
	var classId string = "fcpool.Pool"
	this.ClassId = classId
	var objectType string = "fcpool.Pool"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *FcpoolPool) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FcpoolPool) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FcpoolPool) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FcpoolPool) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FcpoolPool) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FcpoolPool) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIdBlocks returns the IdBlocks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FcpoolPool) GetIdBlocks() []FcpoolBlock {
	if o == nil {
		var ret []FcpoolBlock
		return ret
	}
	return o.IdBlocks
}

// GetIdBlocksOk returns a tuple with the IdBlocks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FcpoolPool) GetIdBlocksOk() ([]FcpoolBlock, bool) {
	if o == nil || o.IdBlocks == nil {
		return nil, false
	}
	return o.IdBlocks, true
}

// HasIdBlocks returns a boolean if a field has been set.
func (o *FcpoolPool) HasIdBlocks() bool {
	if o != nil && o.IdBlocks != nil {
		return true
	}

	return false
}

// SetIdBlocks gets a reference to the given []FcpoolBlock and assigns it to the IdBlocks field.
func (o *FcpoolPool) SetIdBlocks(v []FcpoolBlock) {
	o.IdBlocks = v
}

// GetPoolPurpose returns the PoolPurpose field value if set, zero value otherwise.
func (o *FcpoolPool) GetPoolPurpose() string {
	if o == nil || o.PoolPurpose == nil {
		var ret string
		return ret
	}
	return *o.PoolPurpose
}

// GetPoolPurposeOk returns a tuple with the PoolPurpose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcpoolPool) GetPoolPurposeOk() (*string, bool) {
	if o == nil || o.PoolPurpose == nil {
		return nil, false
	}
	return o.PoolPurpose, true
}

// HasPoolPurpose returns a boolean if a field has been set.
func (o *FcpoolPool) HasPoolPurpose() bool {
	if o != nil && o.PoolPurpose != nil {
		return true
	}

	return false
}

// SetPoolPurpose gets a reference to the given string and assigns it to the PoolPurpose field.
func (o *FcpoolPool) SetPoolPurpose(v string) {
	o.PoolPurpose = &v
}

// GetBlockHeads returns the BlockHeads field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FcpoolPool) GetBlockHeads() []FcpoolFcBlockRelationship {
	if o == nil {
		var ret []FcpoolFcBlockRelationship
		return ret
	}
	return o.BlockHeads
}

// GetBlockHeadsOk returns a tuple with the BlockHeads field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FcpoolPool) GetBlockHeadsOk() ([]FcpoolFcBlockRelationship, bool) {
	if o == nil || o.BlockHeads == nil {
		return nil, false
	}
	return o.BlockHeads, true
}

// HasBlockHeads returns a boolean if a field has been set.
func (o *FcpoolPool) HasBlockHeads() bool {
	if o != nil && o.BlockHeads != nil {
		return true
	}

	return false
}

// SetBlockHeads gets a reference to the given []FcpoolFcBlockRelationship and assigns it to the BlockHeads field.
func (o *FcpoolPool) SetBlockHeads(v []FcpoolFcBlockRelationship) {
	o.BlockHeads = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *FcpoolPool) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcpoolPool) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *FcpoolPool) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *FcpoolPool) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetReservations returns the Reservations field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FcpoolPool) GetReservations() []FcpoolReservationRelationship {
	if o == nil {
		var ret []FcpoolReservationRelationship
		return ret
	}
	return o.Reservations
}

// GetReservationsOk returns a tuple with the Reservations field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FcpoolPool) GetReservationsOk() ([]FcpoolReservationRelationship, bool) {
	if o == nil || o.Reservations == nil {
		return nil, false
	}
	return o.Reservations, true
}

// HasReservations returns a boolean if a field has been set.
func (o *FcpoolPool) HasReservations() bool {
	if o != nil && o.Reservations != nil {
		return true
	}

	return false
}

// SetReservations gets a reference to the given []FcpoolReservationRelationship and assigns it to the Reservations field.
func (o *FcpoolPool) SetReservations(v []FcpoolReservationRelationship) {
	o.Reservations = v
}

func (o FcpoolPool) MarshalJSON() ([]byte, error) {
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
	if o.IdBlocks != nil {
		toSerialize["IdBlocks"] = o.IdBlocks
	}
	if o.PoolPurpose != nil {
		toSerialize["PoolPurpose"] = o.PoolPurpose
	}
	if o.BlockHeads != nil {
		toSerialize["BlockHeads"] = o.BlockHeads
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.Reservations != nil {
		toSerialize["Reservations"] = o.Reservations
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FcpoolPool) UnmarshalJSON(bytes []byte) (err error) {
	type FcpoolPoolWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string        `json:"ObjectType"`
		IdBlocks   []FcpoolBlock `json:"IdBlocks,omitempty"`
		// Purpose of this WWN pool.
		PoolPurpose *string `json:"PoolPurpose,omitempty"`
		// An array of relationships to fcpoolFcBlock resources.
		BlockHeads   []FcpoolFcBlockRelationship           `json:"BlockHeads,omitempty"`
		Organization *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
		// An array of relationships to fcpoolReservation resources.
		Reservations []FcpoolReservationRelationship `json:"Reservations,omitempty"`
	}

	varFcpoolPoolWithoutEmbeddedStruct := FcpoolPoolWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varFcpoolPoolWithoutEmbeddedStruct)
	if err == nil {
		varFcpoolPool := _FcpoolPool{}
		varFcpoolPool.ClassId = varFcpoolPoolWithoutEmbeddedStruct.ClassId
		varFcpoolPool.ObjectType = varFcpoolPoolWithoutEmbeddedStruct.ObjectType
		varFcpoolPool.IdBlocks = varFcpoolPoolWithoutEmbeddedStruct.IdBlocks
		varFcpoolPool.PoolPurpose = varFcpoolPoolWithoutEmbeddedStruct.PoolPurpose
		varFcpoolPool.BlockHeads = varFcpoolPoolWithoutEmbeddedStruct.BlockHeads
		varFcpoolPool.Organization = varFcpoolPoolWithoutEmbeddedStruct.Organization
		varFcpoolPool.Reservations = varFcpoolPoolWithoutEmbeddedStruct.Reservations
		*o = FcpoolPool(varFcpoolPool)
	} else {
		return err
	}

	varFcpoolPool := _FcpoolPool{}

	err = json.Unmarshal(bytes, &varFcpoolPool)
	if err == nil {
		o.PoolAbstractPool = varFcpoolPool.PoolAbstractPool
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "IdBlocks")
		delete(additionalProperties, "PoolPurpose")
		delete(additionalProperties, "BlockHeads")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "Reservations")

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

type NullableFcpoolPool struct {
	value *FcpoolPool
	isSet bool
}

func (v NullableFcpoolPool) Get() *FcpoolPool {
	return v.value
}

func (v *NullableFcpoolPool) Set(val *FcpoolPool) {
	v.value = val
	v.isSet = true
}

func (v NullableFcpoolPool) IsSet() bool {
	return v.isSet
}

func (v *NullableFcpoolPool) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFcpoolPool(val *FcpoolPool) *NullableFcpoolPool {
	return &NullableFcpoolPool{value: val, isSet: true}
}

func (v NullableFcpoolPool) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFcpoolPool) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
