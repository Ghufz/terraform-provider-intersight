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

// StorageNetAppNode NetApp node is a controller in a NetApp cluster. Services and components are controlled and managed by the NetApp node.
type StorageNetAppNode struct {
	StorageBaseArrayController
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType            string                                  `json:"ObjectType"`
	AvgPerformanceMetrics *StorageNetAppPerformanceMetricsAverage `json:"AvgPerformanceMetrics,omitempty"`
	// Storage node option for cdpd state. * `unknown` - The cdpd option is unknown on the node. * `on` - The cdpd option is enabled on the node. * `off` - The cdpd option is disabled on the node.
	CdpdEnabled *string `json:"CdpdEnabled,omitempty"`
	// The health of the NetApp Node.
	Health           *bool                                 `json:"Health,omitempty"`
	HighAvailability NullableStorageNetAppHighAvailability `json:"HighAvailability,omitempty"`
	// Unique identifier of NetApp Node across data center.
	Key *string `json:"Key,omitempty"`
	// The system id of the NetApp Node.
	Systemid *string `json:"Systemid,omitempty"`
	// Universally unique identifier of NetApp Node.
	Uuid  *string                           `json:"Uuid,omitempty"`
	Array *StorageNetAppClusterRelationship `json:"Array,omitempty"`
	// An array of relationships to storageNetAppNodeEvent resources.
	Events               []StorageNetAppNodeEventRelationship `json:"Events,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppNode StorageNetAppNode

// NewStorageNetAppNode instantiates a new StorageNetAppNode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppNode(classId string, objectType string) *StorageNetAppNode {
	this := StorageNetAppNode{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppNodeWithDefaults instantiates a new StorageNetAppNode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppNodeWithDefaults() *StorageNetAppNode {
	this := StorageNetAppNode{}
	var classId string = "storage.NetAppNode"
	this.ClassId = classId
	var objectType string = "storage.NetAppNode"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppNode) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppNode) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppNode) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppNode) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppNode) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppNode) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAvgPerformanceMetrics returns the AvgPerformanceMetrics field value if set, zero value otherwise.
func (o *StorageNetAppNode) GetAvgPerformanceMetrics() StorageNetAppPerformanceMetricsAverage {
	if o == nil || o.AvgPerformanceMetrics == nil {
		var ret StorageNetAppPerformanceMetricsAverage
		return ret
	}
	return *o.AvgPerformanceMetrics
}

// GetAvgPerformanceMetricsOk returns a tuple with the AvgPerformanceMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppNode) GetAvgPerformanceMetricsOk() (*StorageNetAppPerformanceMetricsAverage, bool) {
	if o == nil || o.AvgPerformanceMetrics == nil {
		return nil, false
	}
	return o.AvgPerformanceMetrics, true
}

// HasAvgPerformanceMetrics returns a boolean if a field has been set.
func (o *StorageNetAppNode) HasAvgPerformanceMetrics() bool {
	if o != nil && o.AvgPerformanceMetrics != nil {
		return true
	}

	return false
}

// SetAvgPerformanceMetrics gets a reference to the given StorageNetAppPerformanceMetricsAverage and assigns it to the AvgPerformanceMetrics field.
func (o *StorageNetAppNode) SetAvgPerformanceMetrics(v StorageNetAppPerformanceMetricsAverage) {
	o.AvgPerformanceMetrics = &v
}

// GetCdpdEnabled returns the CdpdEnabled field value if set, zero value otherwise.
func (o *StorageNetAppNode) GetCdpdEnabled() string {
	if o == nil || o.CdpdEnabled == nil {
		var ret string
		return ret
	}
	return *o.CdpdEnabled
}

// GetCdpdEnabledOk returns a tuple with the CdpdEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppNode) GetCdpdEnabledOk() (*string, bool) {
	if o == nil || o.CdpdEnabled == nil {
		return nil, false
	}
	return o.CdpdEnabled, true
}

// HasCdpdEnabled returns a boolean if a field has been set.
func (o *StorageNetAppNode) HasCdpdEnabled() bool {
	if o != nil && o.CdpdEnabled != nil {
		return true
	}

	return false
}

// SetCdpdEnabled gets a reference to the given string and assigns it to the CdpdEnabled field.
func (o *StorageNetAppNode) SetCdpdEnabled(v string) {
	o.CdpdEnabled = &v
}

// GetHealth returns the Health field value if set, zero value otherwise.
func (o *StorageNetAppNode) GetHealth() bool {
	if o == nil || o.Health == nil {
		var ret bool
		return ret
	}
	return *o.Health
}

// GetHealthOk returns a tuple with the Health field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppNode) GetHealthOk() (*bool, bool) {
	if o == nil || o.Health == nil {
		return nil, false
	}
	return o.Health, true
}

// HasHealth returns a boolean if a field has been set.
func (o *StorageNetAppNode) HasHealth() bool {
	if o != nil && o.Health != nil {
		return true
	}

	return false
}

// SetHealth gets a reference to the given bool and assigns it to the Health field.
func (o *StorageNetAppNode) SetHealth(v bool) {
	o.Health = &v
}

// GetHighAvailability returns the HighAvailability field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppNode) GetHighAvailability() StorageNetAppHighAvailability {
	if o == nil || o.HighAvailability.Get() == nil {
		var ret StorageNetAppHighAvailability
		return ret
	}
	return *o.HighAvailability.Get()
}

// GetHighAvailabilityOk returns a tuple with the HighAvailability field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppNode) GetHighAvailabilityOk() (*StorageNetAppHighAvailability, bool) {
	if o == nil {
		return nil, false
	}
	return o.HighAvailability.Get(), o.HighAvailability.IsSet()
}

// HasHighAvailability returns a boolean if a field has been set.
func (o *StorageNetAppNode) HasHighAvailability() bool {
	if o != nil && o.HighAvailability.IsSet() {
		return true
	}

	return false
}

// SetHighAvailability gets a reference to the given NullableStorageNetAppHighAvailability and assigns it to the HighAvailability field.
func (o *StorageNetAppNode) SetHighAvailability(v StorageNetAppHighAvailability) {
	o.HighAvailability.Set(&v)
}

// SetHighAvailabilityNil sets the value for HighAvailability to be an explicit nil
func (o *StorageNetAppNode) SetHighAvailabilityNil() {
	o.HighAvailability.Set(nil)
}

// UnsetHighAvailability ensures that no value is present for HighAvailability, not even an explicit nil
func (o *StorageNetAppNode) UnsetHighAvailability() {
	o.HighAvailability.Unset()
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *StorageNetAppNode) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppNode) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *StorageNetAppNode) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *StorageNetAppNode) SetKey(v string) {
	o.Key = &v
}

// GetSystemid returns the Systemid field value if set, zero value otherwise.
func (o *StorageNetAppNode) GetSystemid() string {
	if o == nil || o.Systemid == nil {
		var ret string
		return ret
	}
	return *o.Systemid
}

// GetSystemidOk returns a tuple with the Systemid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppNode) GetSystemidOk() (*string, bool) {
	if o == nil || o.Systemid == nil {
		return nil, false
	}
	return o.Systemid, true
}

// HasSystemid returns a boolean if a field has been set.
func (o *StorageNetAppNode) HasSystemid() bool {
	if o != nil && o.Systemid != nil {
		return true
	}

	return false
}

// SetSystemid gets a reference to the given string and assigns it to the Systemid field.
func (o *StorageNetAppNode) SetSystemid(v string) {
	o.Systemid = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *StorageNetAppNode) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppNode) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *StorageNetAppNode) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *StorageNetAppNode) SetUuid(v string) {
	o.Uuid = &v
}

// GetArray returns the Array field value if set, zero value otherwise.
func (o *StorageNetAppNode) GetArray() StorageNetAppClusterRelationship {
	if o == nil || o.Array == nil {
		var ret StorageNetAppClusterRelationship
		return ret
	}
	return *o.Array
}

// GetArrayOk returns a tuple with the Array field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppNode) GetArrayOk() (*StorageNetAppClusterRelationship, bool) {
	if o == nil || o.Array == nil {
		return nil, false
	}
	return o.Array, true
}

// HasArray returns a boolean if a field has been set.
func (o *StorageNetAppNode) HasArray() bool {
	if o != nil && o.Array != nil {
		return true
	}

	return false
}

// SetArray gets a reference to the given StorageNetAppClusterRelationship and assigns it to the Array field.
func (o *StorageNetAppNode) SetArray(v StorageNetAppClusterRelationship) {
	o.Array = &v
}

// GetEvents returns the Events field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppNode) GetEvents() []StorageNetAppNodeEventRelationship {
	if o == nil {
		var ret []StorageNetAppNodeEventRelationship
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppNode) GetEventsOk() ([]StorageNetAppNodeEventRelationship, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *StorageNetAppNode) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []StorageNetAppNodeEventRelationship and assigns it to the Events field.
func (o *StorageNetAppNode) SetEvents(v []StorageNetAppNodeEventRelationship) {
	o.Events = v
}

func (o StorageNetAppNode) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedStorageBaseArrayController, errStorageBaseArrayController := json.Marshal(o.StorageBaseArrayController)
	if errStorageBaseArrayController != nil {
		return []byte{}, errStorageBaseArrayController
	}
	errStorageBaseArrayController = json.Unmarshal([]byte(serializedStorageBaseArrayController), &toSerialize)
	if errStorageBaseArrayController != nil {
		return []byte{}, errStorageBaseArrayController
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AvgPerformanceMetrics != nil {
		toSerialize["AvgPerformanceMetrics"] = o.AvgPerformanceMetrics
	}
	if o.CdpdEnabled != nil {
		toSerialize["CdpdEnabled"] = o.CdpdEnabled
	}
	if o.Health != nil {
		toSerialize["Health"] = o.Health
	}
	if o.HighAvailability.IsSet() {
		toSerialize["HighAvailability"] = o.HighAvailability.Get()
	}
	if o.Key != nil {
		toSerialize["Key"] = o.Key
	}
	if o.Systemid != nil {
		toSerialize["Systemid"] = o.Systemid
	}
	if o.Uuid != nil {
		toSerialize["Uuid"] = o.Uuid
	}
	if o.Array != nil {
		toSerialize["Array"] = o.Array
	}
	if o.Events != nil {
		toSerialize["Events"] = o.Events
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageNetAppNode) UnmarshalJSON(bytes []byte) (err error) {
	type StorageNetAppNodeWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType            string                                  `json:"ObjectType"`
		AvgPerformanceMetrics *StorageNetAppPerformanceMetricsAverage `json:"AvgPerformanceMetrics,omitempty"`
		// Storage node option for cdpd state. * `unknown` - The cdpd option is unknown on the node. * `on` - The cdpd option is enabled on the node. * `off` - The cdpd option is disabled on the node.
		CdpdEnabled *string `json:"CdpdEnabled,omitempty"`
		// The health of the NetApp Node.
		Health           *bool                                 `json:"Health,omitempty"`
		HighAvailability NullableStorageNetAppHighAvailability `json:"HighAvailability,omitempty"`
		// Unique identifier of NetApp Node across data center.
		Key *string `json:"Key,omitempty"`
		// The system id of the NetApp Node.
		Systemid *string `json:"Systemid,omitempty"`
		// Universally unique identifier of NetApp Node.
		Uuid  *string                           `json:"Uuid,omitempty"`
		Array *StorageNetAppClusterRelationship `json:"Array,omitempty"`
		// An array of relationships to storageNetAppNodeEvent resources.
		Events []StorageNetAppNodeEventRelationship `json:"Events,omitempty"`
	}

	varStorageNetAppNodeWithoutEmbeddedStruct := StorageNetAppNodeWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageNetAppNodeWithoutEmbeddedStruct)
	if err == nil {
		varStorageNetAppNode := _StorageNetAppNode{}
		varStorageNetAppNode.ClassId = varStorageNetAppNodeWithoutEmbeddedStruct.ClassId
		varStorageNetAppNode.ObjectType = varStorageNetAppNodeWithoutEmbeddedStruct.ObjectType
		varStorageNetAppNode.AvgPerformanceMetrics = varStorageNetAppNodeWithoutEmbeddedStruct.AvgPerformanceMetrics
		varStorageNetAppNode.CdpdEnabled = varStorageNetAppNodeWithoutEmbeddedStruct.CdpdEnabled
		varStorageNetAppNode.Health = varStorageNetAppNodeWithoutEmbeddedStruct.Health
		varStorageNetAppNode.HighAvailability = varStorageNetAppNodeWithoutEmbeddedStruct.HighAvailability
		varStorageNetAppNode.Key = varStorageNetAppNodeWithoutEmbeddedStruct.Key
		varStorageNetAppNode.Systemid = varStorageNetAppNodeWithoutEmbeddedStruct.Systemid
		varStorageNetAppNode.Uuid = varStorageNetAppNodeWithoutEmbeddedStruct.Uuid
		varStorageNetAppNode.Array = varStorageNetAppNodeWithoutEmbeddedStruct.Array
		varStorageNetAppNode.Events = varStorageNetAppNodeWithoutEmbeddedStruct.Events
		*o = StorageNetAppNode(varStorageNetAppNode)
	} else {
		return err
	}

	varStorageNetAppNode := _StorageNetAppNode{}

	err = json.Unmarshal(bytes, &varStorageNetAppNode)
	if err == nil {
		o.StorageBaseArrayController = varStorageNetAppNode.StorageBaseArrayController
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AvgPerformanceMetrics")
		delete(additionalProperties, "CdpdEnabled")
		delete(additionalProperties, "Health")
		delete(additionalProperties, "HighAvailability")
		delete(additionalProperties, "Key")
		delete(additionalProperties, "Systemid")
		delete(additionalProperties, "Uuid")
		delete(additionalProperties, "Array")
		delete(additionalProperties, "Events")

		// remove fields from embedded structs
		reflectStorageBaseArrayController := reflect.ValueOf(o.StorageBaseArrayController)
		for i := 0; i < reflectStorageBaseArrayController.Type().NumField(); i++ {
			t := reflectStorageBaseArrayController.Type().Field(i)

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

type NullableStorageNetAppNode struct {
	value *StorageNetAppNode
	isSet bool
}

func (v NullableStorageNetAppNode) Get() *StorageNetAppNode {
	return v.value
}

func (v *NullableStorageNetAppNode) Set(val *StorageNetAppNode) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppNode) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppNode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppNode(val *StorageNetAppNode) *NullableStorageNetAppNode {
	return &NullableStorageNetAppNode{value: val, isSet: true}
}

func (v NullableStorageNetAppNode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppNode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
