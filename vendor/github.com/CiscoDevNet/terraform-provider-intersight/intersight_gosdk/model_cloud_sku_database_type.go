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
	"reflect"
	"strings"
)

// CloudSkuDatabaseType Stores details of instance type which handle databases.
type CloudSkuDatabaseType struct {
	CloudBaseSku
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The database edition. For e.g. standard or enterprise.
	DatabaseEdition *string `json:"DatabaseEdition,omitempty"`
	// The database engine. For e.g. SQL Server, Oracle, PostgreSQL.
	DatabaseEngine *string `json:"DatabaseEngine,omitempty"`
	// The licensing option for the database. For e.g. license required or not.
	LicenseModel *string `json:"LicenseModel,omitempty"`
	// Network performance of this instance type.
	NetworkPerformance   *string `json:"NetworkPerformance,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CloudSkuDatabaseType CloudSkuDatabaseType

// NewCloudSkuDatabaseType instantiates a new CloudSkuDatabaseType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudSkuDatabaseType(classId string, objectType string) *CloudSkuDatabaseType {
	this := CloudSkuDatabaseType{}
	this.ClassId = classId
	this.ObjectType = objectType
	var currency string = "USD"
	this.Currency = &currency
	var isActive bool = true
	this.IsActive = &isActive
	var platformType string = ""
	this.PlatformType = &platformType
	var serviceCategory string = "Compute"
	this.ServiceCategory = &serviceCategory
	return &this
}

// NewCloudSkuDatabaseTypeWithDefaults instantiates a new CloudSkuDatabaseType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudSkuDatabaseTypeWithDefaults() *CloudSkuDatabaseType {
	this := CloudSkuDatabaseType{}
	var classId string = "cloud.SkuDatabaseType"
	this.ClassId = classId
	var objectType string = "cloud.SkuDatabaseType"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CloudSkuDatabaseType) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CloudSkuDatabaseType) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CloudSkuDatabaseType) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CloudSkuDatabaseType) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CloudSkuDatabaseType) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CloudSkuDatabaseType) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDatabaseEdition returns the DatabaseEdition field value if set, zero value otherwise.
func (o *CloudSkuDatabaseType) GetDatabaseEdition() string {
	if o == nil || o.DatabaseEdition == nil {
		var ret string
		return ret
	}
	return *o.DatabaseEdition
}

// GetDatabaseEditionOk returns a tuple with the DatabaseEdition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudSkuDatabaseType) GetDatabaseEditionOk() (*string, bool) {
	if o == nil || o.DatabaseEdition == nil {
		return nil, false
	}
	return o.DatabaseEdition, true
}

// HasDatabaseEdition returns a boolean if a field has been set.
func (o *CloudSkuDatabaseType) HasDatabaseEdition() bool {
	if o != nil && o.DatabaseEdition != nil {
		return true
	}

	return false
}

// SetDatabaseEdition gets a reference to the given string and assigns it to the DatabaseEdition field.
func (o *CloudSkuDatabaseType) SetDatabaseEdition(v string) {
	o.DatabaseEdition = &v
}

// GetDatabaseEngine returns the DatabaseEngine field value if set, zero value otherwise.
func (o *CloudSkuDatabaseType) GetDatabaseEngine() string {
	if o == nil || o.DatabaseEngine == nil {
		var ret string
		return ret
	}
	return *o.DatabaseEngine
}

// GetDatabaseEngineOk returns a tuple with the DatabaseEngine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudSkuDatabaseType) GetDatabaseEngineOk() (*string, bool) {
	if o == nil || o.DatabaseEngine == nil {
		return nil, false
	}
	return o.DatabaseEngine, true
}

// HasDatabaseEngine returns a boolean if a field has been set.
func (o *CloudSkuDatabaseType) HasDatabaseEngine() bool {
	if o != nil && o.DatabaseEngine != nil {
		return true
	}

	return false
}

// SetDatabaseEngine gets a reference to the given string and assigns it to the DatabaseEngine field.
func (o *CloudSkuDatabaseType) SetDatabaseEngine(v string) {
	o.DatabaseEngine = &v
}

// GetLicenseModel returns the LicenseModel field value if set, zero value otherwise.
func (o *CloudSkuDatabaseType) GetLicenseModel() string {
	if o == nil || o.LicenseModel == nil {
		var ret string
		return ret
	}
	return *o.LicenseModel
}

// GetLicenseModelOk returns a tuple with the LicenseModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudSkuDatabaseType) GetLicenseModelOk() (*string, bool) {
	if o == nil || o.LicenseModel == nil {
		return nil, false
	}
	return o.LicenseModel, true
}

// HasLicenseModel returns a boolean if a field has been set.
func (o *CloudSkuDatabaseType) HasLicenseModel() bool {
	if o != nil && o.LicenseModel != nil {
		return true
	}

	return false
}

// SetLicenseModel gets a reference to the given string and assigns it to the LicenseModel field.
func (o *CloudSkuDatabaseType) SetLicenseModel(v string) {
	o.LicenseModel = &v
}

// GetNetworkPerformance returns the NetworkPerformance field value if set, zero value otherwise.
func (o *CloudSkuDatabaseType) GetNetworkPerformance() string {
	if o == nil || o.NetworkPerformance == nil {
		var ret string
		return ret
	}
	return *o.NetworkPerformance
}

// GetNetworkPerformanceOk returns a tuple with the NetworkPerformance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudSkuDatabaseType) GetNetworkPerformanceOk() (*string, bool) {
	if o == nil || o.NetworkPerformance == nil {
		return nil, false
	}
	return o.NetworkPerformance, true
}

// HasNetworkPerformance returns a boolean if a field has been set.
func (o *CloudSkuDatabaseType) HasNetworkPerformance() bool {
	if o != nil && o.NetworkPerformance != nil {
		return true
	}

	return false
}

// SetNetworkPerformance gets a reference to the given string and assigns it to the NetworkPerformance field.
func (o *CloudSkuDatabaseType) SetNetworkPerformance(v string) {
	o.NetworkPerformance = &v
}

func (o CloudSkuDatabaseType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedCloudBaseSku, errCloudBaseSku := json.Marshal(o.CloudBaseSku)
	if errCloudBaseSku != nil {
		return []byte{}, errCloudBaseSku
	}
	errCloudBaseSku = json.Unmarshal([]byte(serializedCloudBaseSku), &toSerialize)
	if errCloudBaseSku != nil {
		return []byte{}, errCloudBaseSku
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DatabaseEdition != nil {
		toSerialize["DatabaseEdition"] = o.DatabaseEdition
	}
	if o.DatabaseEngine != nil {
		toSerialize["DatabaseEngine"] = o.DatabaseEngine
	}
	if o.LicenseModel != nil {
		toSerialize["LicenseModel"] = o.LicenseModel
	}
	if o.NetworkPerformance != nil {
		toSerialize["NetworkPerformance"] = o.NetworkPerformance
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CloudSkuDatabaseType) UnmarshalJSON(bytes []byte) (err error) {
	type CloudSkuDatabaseTypeWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The database edition. For e.g. standard or enterprise.
		DatabaseEdition *string `json:"DatabaseEdition,omitempty"`
		// The database engine. For e.g. SQL Server, Oracle, PostgreSQL.
		DatabaseEngine *string `json:"DatabaseEngine,omitempty"`
		// The licensing option for the database. For e.g. license required or not.
		LicenseModel *string `json:"LicenseModel,omitempty"`
		// Network performance of this instance type.
		NetworkPerformance *string `json:"NetworkPerformance,omitempty"`
	}

	varCloudSkuDatabaseTypeWithoutEmbeddedStruct := CloudSkuDatabaseTypeWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varCloudSkuDatabaseTypeWithoutEmbeddedStruct)
	if err == nil {
		varCloudSkuDatabaseType := _CloudSkuDatabaseType{}
		varCloudSkuDatabaseType.ClassId = varCloudSkuDatabaseTypeWithoutEmbeddedStruct.ClassId
		varCloudSkuDatabaseType.ObjectType = varCloudSkuDatabaseTypeWithoutEmbeddedStruct.ObjectType
		varCloudSkuDatabaseType.DatabaseEdition = varCloudSkuDatabaseTypeWithoutEmbeddedStruct.DatabaseEdition
		varCloudSkuDatabaseType.DatabaseEngine = varCloudSkuDatabaseTypeWithoutEmbeddedStruct.DatabaseEngine
		varCloudSkuDatabaseType.LicenseModel = varCloudSkuDatabaseTypeWithoutEmbeddedStruct.LicenseModel
		varCloudSkuDatabaseType.NetworkPerformance = varCloudSkuDatabaseTypeWithoutEmbeddedStruct.NetworkPerformance
		*o = CloudSkuDatabaseType(varCloudSkuDatabaseType)
	} else {
		return err
	}

	varCloudSkuDatabaseType := _CloudSkuDatabaseType{}

	err = json.Unmarshal(bytes, &varCloudSkuDatabaseType)
	if err == nil {
		o.CloudBaseSku = varCloudSkuDatabaseType.CloudBaseSku
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DatabaseEdition")
		delete(additionalProperties, "DatabaseEngine")
		delete(additionalProperties, "LicenseModel")
		delete(additionalProperties, "NetworkPerformance")

		// remove fields from embedded structs
		reflectCloudBaseSku := reflect.ValueOf(o.CloudBaseSku)
		for i := 0; i < reflectCloudBaseSku.Type().NumField(); i++ {
			t := reflectCloudBaseSku.Type().Field(i)

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

type NullableCloudSkuDatabaseType struct {
	value *CloudSkuDatabaseType
	isSet bool
}

func (v NullableCloudSkuDatabaseType) Get() *CloudSkuDatabaseType {
	return v.value
}

func (v *NullableCloudSkuDatabaseType) Set(val *CloudSkuDatabaseType) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudSkuDatabaseType) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudSkuDatabaseType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudSkuDatabaseType(val *CloudSkuDatabaseType) *NullableCloudSkuDatabaseType {
	return &NullableCloudSkuDatabaseType{value: val, isSet: true}
}

func (v NullableCloudSkuDatabaseType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudSkuDatabaseType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
