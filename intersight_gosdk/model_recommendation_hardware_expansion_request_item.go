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

// RecommendationHardwareExpansionRequestItem Entity representing the user request for expansion of each hardware item.
type RecommendationHardwareExpansionRequestItem struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Type of hardware item for which the capacity increase is requested by the user. For example, CPU. * `None` - The Enum value None represents that no value was set for the hardware type. * `CPU` - The Enum value CPU represents that the hardware type requested for expansion is a physical CPU. * `Memory` - The Enum value Memory represents that the hardware type requested for expansion is a memory unit. * `Storage` - The Enum value Storage represents that the hardware type requested for expansion is a storage unit, ie, storage drives.
	ItemType *string `json:"ItemType,omitempty"`
	// The maximum value allowed for expansion for the hardware type on the referred registered device.
	MaxValue *float32 `json:"MaxValue,omitempty"`
	// Unit type for the maximum value of the resource. For example, TB, GB, MB. * `TB` - The Enum value TB represents that the measurement unit is in terabytes. * `MB` - The Enum value MB represents that the measurement unit is in megabytes. * `GB` - The Enum value GB represents that the measurement unit is in gigabytes. * `MHz` - The Enum value MHz represents that the measurement unit is in megahertz. * `GHz` - The Enum value GHz represents that the measurement unit is in gigahertz. * `Percentage` - The Enum value Percentage represents that the expansion request is in the percentage of resource increase. For example, a 20% increase in CPU capacity.
	MaxValueUnit *string `json:"MaxValueUnit,omitempty"`
	// Unit type for the expansion request, i.e., if the increase is requested as a raw value in TB, GB, etc., or in percentage increase. * `TB` - The Enum value TB represents that the measurement unit is in terabytes. * `MB` - The Enum value MB represents that the measurement unit is in megabytes. * `GB` - The Enum value GB represents that the measurement unit is in gigabytes. * `MHz` - The Enum value MHz represents that the measurement unit is in megahertz. * `GHz` - The Enum value GHz represents that the measurement unit is in gigahertz. * `Percentage` - The Enum value Percentage represents that the expansion request is in the percentage of resource increase. For example, a 20% increase in CPU capacity.
	UnitType *string `json:"UnitType,omitempty"`
	// Value of the expansion request which can be absolute value or percentage increase.
	Value                *float32                                            `json:"Value,omitempty"`
	ExpansionRequest     *RecommendationHardwareExpansionRequestRelationship `json:"ExpansionRequest,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RecommendationHardwareExpansionRequestItem RecommendationHardwareExpansionRequestItem

// NewRecommendationHardwareExpansionRequestItem instantiates a new RecommendationHardwareExpansionRequestItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecommendationHardwareExpansionRequestItem(classId string, objectType string) *RecommendationHardwareExpansionRequestItem {
	this := RecommendationHardwareExpansionRequestItem{}
	this.ClassId = classId
	this.ObjectType = objectType
	var itemType string = "None"
	this.ItemType = &itemType
	var maxValueUnit string = "TB"
	this.MaxValueUnit = &maxValueUnit
	var unitType string = "TB"
	this.UnitType = &unitType
	return &this
}

// NewRecommendationHardwareExpansionRequestItemWithDefaults instantiates a new RecommendationHardwareExpansionRequestItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecommendationHardwareExpansionRequestItemWithDefaults() *RecommendationHardwareExpansionRequestItem {
	this := RecommendationHardwareExpansionRequestItem{}
	var classId string = "recommendation.HardwareExpansionRequestItem"
	this.ClassId = classId
	var objectType string = "recommendation.HardwareExpansionRequestItem"
	this.ObjectType = objectType
	var itemType string = "None"
	this.ItemType = &itemType
	var maxValueUnit string = "TB"
	this.MaxValueUnit = &maxValueUnit
	var unitType string = "TB"
	this.UnitType = &unitType
	return &this
}

// GetClassId returns the ClassId field value
func (o *RecommendationHardwareExpansionRequestItem) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *RecommendationHardwareExpansionRequestItem) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *RecommendationHardwareExpansionRequestItem) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *RecommendationHardwareExpansionRequestItem) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *RecommendationHardwareExpansionRequestItem) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *RecommendationHardwareExpansionRequestItem) SetObjectType(v string) {
	o.ObjectType = v
}

// GetItemType returns the ItemType field value if set, zero value otherwise.
func (o *RecommendationHardwareExpansionRequestItem) GetItemType() string {
	if o == nil || o.ItemType == nil {
		var ret string
		return ret
	}
	return *o.ItemType
}

// GetItemTypeOk returns a tuple with the ItemType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationHardwareExpansionRequestItem) GetItemTypeOk() (*string, bool) {
	if o == nil || o.ItemType == nil {
		return nil, false
	}
	return o.ItemType, true
}

// HasItemType returns a boolean if a field has been set.
func (o *RecommendationHardwareExpansionRequestItem) HasItemType() bool {
	if o != nil && o.ItemType != nil {
		return true
	}

	return false
}

// SetItemType gets a reference to the given string and assigns it to the ItemType field.
func (o *RecommendationHardwareExpansionRequestItem) SetItemType(v string) {
	o.ItemType = &v
}

// GetMaxValue returns the MaxValue field value if set, zero value otherwise.
func (o *RecommendationHardwareExpansionRequestItem) GetMaxValue() float32 {
	if o == nil || o.MaxValue == nil {
		var ret float32
		return ret
	}
	return *o.MaxValue
}

// GetMaxValueOk returns a tuple with the MaxValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationHardwareExpansionRequestItem) GetMaxValueOk() (*float32, bool) {
	if o == nil || o.MaxValue == nil {
		return nil, false
	}
	return o.MaxValue, true
}

// HasMaxValue returns a boolean if a field has been set.
func (o *RecommendationHardwareExpansionRequestItem) HasMaxValue() bool {
	if o != nil && o.MaxValue != nil {
		return true
	}

	return false
}

// SetMaxValue gets a reference to the given float32 and assigns it to the MaxValue field.
func (o *RecommendationHardwareExpansionRequestItem) SetMaxValue(v float32) {
	o.MaxValue = &v
}

// GetMaxValueUnit returns the MaxValueUnit field value if set, zero value otherwise.
func (o *RecommendationHardwareExpansionRequestItem) GetMaxValueUnit() string {
	if o == nil || o.MaxValueUnit == nil {
		var ret string
		return ret
	}
	return *o.MaxValueUnit
}

// GetMaxValueUnitOk returns a tuple with the MaxValueUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationHardwareExpansionRequestItem) GetMaxValueUnitOk() (*string, bool) {
	if o == nil || o.MaxValueUnit == nil {
		return nil, false
	}
	return o.MaxValueUnit, true
}

// HasMaxValueUnit returns a boolean if a field has been set.
func (o *RecommendationHardwareExpansionRequestItem) HasMaxValueUnit() bool {
	if o != nil && o.MaxValueUnit != nil {
		return true
	}

	return false
}

// SetMaxValueUnit gets a reference to the given string and assigns it to the MaxValueUnit field.
func (o *RecommendationHardwareExpansionRequestItem) SetMaxValueUnit(v string) {
	o.MaxValueUnit = &v
}

// GetUnitType returns the UnitType field value if set, zero value otherwise.
func (o *RecommendationHardwareExpansionRequestItem) GetUnitType() string {
	if o == nil || o.UnitType == nil {
		var ret string
		return ret
	}
	return *o.UnitType
}

// GetUnitTypeOk returns a tuple with the UnitType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationHardwareExpansionRequestItem) GetUnitTypeOk() (*string, bool) {
	if o == nil || o.UnitType == nil {
		return nil, false
	}
	return o.UnitType, true
}

// HasUnitType returns a boolean if a field has been set.
func (o *RecommendationHardwareExpansionRequestItem) HasUnitType() bool {
	if o != nil && o.UnitType != nil {
		return true
	}

	return false
}

// SetUnitType gets a reference to the given string and assigns it to the UnitType field.
func (o *RecommendationHardwareExpansionRequestItem) SetUnitType(v string) {
	o.UnitType = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *RecommendationHardwareExpansionRequestItem) GetValue() float32 {
	if o == nil || o.Value == nil {
		var ret float32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationHardwareExpansionRequestItem) GetValueOk() (*float32, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *RecommendationHardwareExpansionRequestItem) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given float32 and assigns it to the Value field.
func (o *RecommendationHardwareExpansionRequestItem) SetValue(v float32) {
	o.Value = &v
}

// GetExpansionRequest returns the ExpansionRequest field value if set, zero value otherwise.
func (o *RecommendationHardwareExpansionRequestItem) GetExpansionRequest() RecommendationHardwareExpansionRequestRelationship {
	if o == nil || o.ExpansionRequest == nil {
		var ret RecommendationHardwareExpansionRequestRelationship
		return ret
	}
	return *o.ExpansionRequest
}

// GetExpansionRequestOk returns a tuple with the ExpansionRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationHardwareExpansionRequestItem) GetExpansionRequestOk() (*RecommendationHardwareExpansionRequestRelationship, bool) {
	if o == nil || o.ExpansionRequest == nil {
		return nil, false
	}
	return o.ExpansionRequest, true
}

// HasExpansionRequest returns a boolean if a field has been set.
func (o *RecommendationHardwareExpansionRequestItem) HasExpansionRequest() bool {
	if o != nil && o.ExpansionRequest != nil {
		return true
	}

	return false
}

// SetExpansionRequest gets a reference to the given RecommendationHardwareExpansionRequestRelationship and assigns it to the ExpansionRequest field.
func (o *RecommendationHardwareExpansionRequestItem) SetExpansionRequest(v RecommendationHardwareExpansionRequestRelationship) {
	o.ExpansionRequest = &v
}

func (o RecommendationHardwareExpansionRequestItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ItemType != nil {
		toSerialize["ItemType"] = o.ItemType
	}
	if o.MaxValue != nil {
		toSerialize["MaxValue"] = o.MaxValue
	}
	if o.MaxValueUnit != nil {
		toSerialize["MaxValueUnit"] = o.MaxValueUnit
	}
	if o.UnitType != nil {
		toSerialize["UnitType"] = o.UnitType
	}
	if o.Value != nil {
		toSerialize["Value"] = o.Value
	}
	if o.ExpansionRequest != nil {
		toSerialize["ExpansionRequest"] = o.ExpansionRequest
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RecommendationHardwareExpansionRequestItem) UnmarshalJSON(bytes []byte) (err error) {
	type RecommendationHardwareExpansionRequestItemWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Type of hardware item for which the capacity increase is requested by the user. For example, CPU. * `None` - The Enum value None represents that no value was set for the hardware type. * `CPU` - The Enum value CPU represents that the hardware type requested for expansion is a physical CPU. * `Memory` - The Enum value Memory represents that the hardware type requested for expansion is a memory unit. * `Storage` - The Enum value Storage represents that the hardware type requested for expansion is a storage unit, ie, storage drives.
		ItemType *string `json:"ItemType,omitempty"`
		// The maximum value allowed for expansion for the hardware type on the referred registered device.
		MaxValue *float32 `json:"MaxValue,omitempty"`
		// Unit type for the maximum value of the resource. For example, TB, GB, MB. * `TB` - The Enum value TB represents that the measurement unit is in terabytes. * `MB` - The Enum value MB represents that the measurement unit is in megabytes. * `GB` - The Enum value GB represents that the measurement unit is in gigabytes. * `MHz` - The Enum value MHz represents that the measurement unit is in megahertz. * `GHz` - The Enum value GHz represents that the measurement unit is in gigahertz. * `Percentage` - The Enum value Percentage represents that the expansion request is in the percentage of resource increase. For example, a 20% increase in CPU capacity.
		MaxValueUnit *string `json:"MaxValueUnit,omitempty"`
		// Unit type for the expansion request, i.e., if the increase is requested as a raw value in TB, GB, etc., or in percentage increase. * `TB` - The Enum value TB represents that the measurement unit is in terabytes. * `MB` - The Enum value MB represents that the measurement unit is in megabytes. * `GB` - The Enum value GB represents that the measurement unit is in gigabytes. * `MHz` - The Enum value MHz represents that the measurement unit is in megahertz. * `GHz` - The Enum value GHz represents that the measurement unit is in gigahertz. * `Percentage` - The Enum value Percentage represents that the expansion request is in the percentage of resource increase. For example, a 20% increase in CPU capacity.
		UnitType *string `json:"UnitType,omitempty"`
		// Value of the expansion request which can be absolute value or percentage increase.
		Value            *float32                                            `json:"Value,omitempty"`
		ExpansionRequest *RecommendationHardwareExpansionRequestRelationship `json:"ExpansionRequest,omitempty"`
	}

	varRecommendationHardwareExpansionRequestItemWithoutEmbeddedStruct := RecommendationHardwareExpansionRequestItemWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varRecommendationHardwareExpansionRequestItemWithoutEmbeddedStruct)
	if err == nil {
		varRecommendationHardwareExpansionRequestItem := _RecommendationHardwareExpansionRequestItem{}
		varRecommendationHardwareExpansionRequestItem.ClassId = varRecommendationHardwareExpansionRequestItemWithoutEmbeddedStruct.ClassId
		varRecommendationHardwareExpansionRequestItem.ObjectType = varRecommendationHardwareExpansionRequestItemWithoutEmbeddedStruct.ObjectType
		varRecommendationHardwareExpansionRequestItem.ItemType = varRecommendationHardwareExpansionRequestItemWithoutEmbeddedStruct.ItemType
		varRecommendationHardwareExpansionRequestItem.MaxValue = varRecommendationHardwareExpansionRequestItemWithoutEmbeddedStruct.MaxValue
		varRecommendationHardwareExpansionRequestItem.MaxValueUnit = varRecommendationHardwareExpansionRequestItemWithoutEmbeddedStruct.MaxValueUnit
		varRecommendationHardwareExpansionRequestItem.UnitType = varRecommendationHardwareExpansionRequestItemWithoutEmbeddedStruct.UnitType
		varRecommendationHardwareExpansionRequestItem.Value = varRecommendationHardwareExpansionRequestItemWithoutEmbeddedStruct.Value
		varRecommendationHardwareExpansionRequestItem.ExpansionRequest = varRecommendationHardwareExpansionRequestItemWithoutEmbeddedStruct.ExpansionRequest
		*o = RecommendationHardwareExpansionRequestItem(varRecommendationHardwareExpansionRequestItem)
	} else {
		return err
	}

	varRecommendationHardwareExpansionRequestItem := _RecommendationHardwareExpansionRequestItem{}

	err = json.Unmarshal(bytes, &varRecommendationHardwareExpansionRequestItem)
	if err == nil {
		o.MoBaseMo = varRecommendationHardwareExpansionRequestItem.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ItemType")
		delete(additionalProperties, "MaxValue")
		delete(additionalProperties, "MaxValueUnit")
		delete(additionalProperties, "UnitType")
		delete(additionalProperties, "Value")
		delete(additionalProperties, "ExpansionRequest")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

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

type NullableRecommendationHardwareExpansionRequestItem struct {
	value *RecommendationHardwareExpansionRequestItem
	isSet bool
}

func (v NullableRecommendationHardwareExpansionRequestItem) Get() *RecommendationHardwareExpansionRequestItem {
	return v.value
}

func (v *NullableRecommendationHardwareExpansionRequestItem) Set(val *RecommendationHardwareExpansionRequestItem) {
	v.value = val
	v.isSet = true
}

func (v NullableRecommendationHardwareExpansionRequestItem) IsSet() bool {
	return v.isSet
}

func (v *NullableRecommendationHardwareExpansionRequestItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecommendationHardwareExpansionRequestItem(val *RecommendationHardwareExpansionRequestItem) *NullableRecommendationHardwareExpansionRequestItem {
	return &NullableRecommendationHardwareExpansionRequestItem{value: val, isSet: true}
}

func (v NullableRecommendationHardwareExpansionRequestItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecommendationHardwareExpansionRequestItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
