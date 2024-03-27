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

// NiatelemetryDigitalOpticalMonitoring Digital optical monitoring details for aci nodes.
type NiatelemetryDigitalOpticalMonitoring struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Alerts count for the interface in the node.
	Alerts *string `json:"Alerts,omitempty"`
	// Dn with interface name for the aci nodes.
	Dn *string `json:"Dn,omitempty"`
	// RxLos count for the interface in the node.
	RxLos *string `json:"RxLos,omitempty"`
	// TxfaultCount for the interface in the node.
	TxFaultCount         *string `json:"TxFaultCount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryDigitalOpticalMonitoring NiatelemetryDigitalOpticalMonitoring

// NewNiatelemetryDigitalOpticalMonitoring instantiates a new NiatelemetryDigitalOpticalMonitoring object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryDigitalOpticalMonitoring(classId string, objectType string) *NiatelemetryDigitalOpticalMonitoring {
	this := NiatelemetryDigitalOpticalMonitoring{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryDigitalOpticalMonitoringWithDefaults instantiates a new NiatelemetryDigitalOpticalMonitoring object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryDigitalOpticalMonitoringWithDefaults() *NiatelemetryDigitalOpticalMonitoring {
	this := NiatelemetryDigitalOpticalMonitoring{}
	var classId string = "niatelemetry.DigitalOpticalMonitoring"
	this.ClassId = classId
	var objectType string = "niatelemetry.DigitalOpticalMonitoring"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryDigitalOpticalMonitoring) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryDigitalOpticalMonitoring) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryDigitalOpticalMonitoring) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryDigitalOpticalMonitoring) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryDigitalOpticalMonitoring) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryDigitalOpticalMonitoring) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAlerts returns the Alerts field value if set, zero value otherwise.
func (o *NiatelemetryDigitalOpticalMonitoring) GetAlerts() string {
	if o == nil || o.Alerts == nil {
		var ret string
		return ret
	}
	return *o.Alerts
}

// GetAlertsOk returns a tuple with the Alerts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryDigitalOpticalMonitoring) GetAlertsOk() (*string, bool) {
	if o == nil || o.Alerts == nil {
		return nil, false
	}
	return o.Alerts, true
}

// HasAlerts returns a boolean if a field has been set.
func (o *NiatelemetryDigitalOpticalMonitoring) HasAlerts() bool {
	if o != nil && o.Alerts != nil {
		return true
	}

	return false
}

// SetAlerts gets a reference to the given string and assigns it to the Alerts field.
func (o *NiatelemetryDigitalOpticalMonitoring) SetAlerts(v string) {
	o.Alerts = &v
}

// GetDn returns the Dn field value if set, zero value otherwise.
func (o *NiatelemetryDigitalOpticalMonitoring) GetDn() string {
	if o == nil || o.Dn == nil {
		var ret string
		return ret
	}
	return *o.Dn
}

// GetDnOk returns a tuple with the Dn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryDigitalOpticalMonitoring) GetDnOk() (*string, bool) {
	if o == nil || o.Dn == nil {
		return nil, false
	}
	return o.Dn, true
}

// HasDn returns a boolean if a field has been set.
func (o *NiatelemetryDigitalOpticalMonitoring) HasDn() bool {
	if o != nil && o.Dn != nil {
		return true
	}

	return false
}

// SetDn gets a reference to the given string and assigns it to the Dn field.
func (o *NiatelemetryDigitalOpticalMonitoring) SetDn(v string) {
	o.Dn = &v
}

// GetRxLos returns the RxLos field value if set, zero value otherwise.
func (o *NiatelemetryDigitalOpticalMonitoring) GetRxLos() string {
	if o == nil || o.RxLos == nil {
		var ret string
		return ret
	}
	return *o.RxLos
}

// GetRxLosOk returns a tuple with the RxLos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryDigitalOpticalMonitoring) GetRxLosOk() (*string, bool) {
	if o == nil || o.RxLos == nil {
		return nil, false
	}
	return o.RxLos, true
}

// HasRxLos returns a boolean if a field has been set.
func (o *NiatelemetryDigitalOpticalMonitoring) HasRxLos() bool {
	if o != nil && o.RxLos != nil {
		return true
	}

	return false
}

// SetRxLos gets a reference to the given string and assigns it to the RxLos field.
func (o *NiatelemetryDigitalOpticalMonitoring) SetRxLos(v string) {
	o.RxLos = &v
}

// GetTxFaultCount returns the TxFaultCount field value if set, zero value otherwise.
func (o *NiatelemetryDigitalOpticalMonitoring) GetTxFaultCount() string {
	if o == nil || o.TxFaultCount == nil {
		var ret string
		return ret
	}
	return *o.TxFaultCount
}

// GetTxFaultCountOk returns a tuple with the TxFaultCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryDigitalOpticalMonitoring) GetTxFaultCountOk() (*string, bool) {
	if o == nil || o.TxFaultCount == nil {
		return nil, false
	}
	return o.TxFaultCount, true
}

// HasTxFaultCount returns a boolean if a field has been set.
func (o *NiatelemetryDigitalOpticalMonitoring) HasTxFaultCount() bool {
	if o != nil && o.TxFaultCount != nil {
		return true
	}

	return false
}

// SetTxFaultCount gets a reference to the given string and assigns it to the TxFaultCount field.
func (o *NiatelemetryDigitalOpticalMonitoring) SetTxFaultCount(v string) {
	o.TxFaultCount = &v
}

func (o NiatelemetryDigitalOpticalMonitoring) MarshalJSON() ([]byte, error) {
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
	if o.Alerts != nil {
		toSerialize["Alerts"] = o.Alerts
	}
	if o.Dn != nil {
		toSerialize["Dn"] = o.Dn
	}
	if o.RxLos != nil {
		toSerialize["RxLos"] = o.RxLos
	}
	if o.TxFaultCount != nil {
		toSerialize["TxFaultCount"] = o.TxFaultCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryDigitalOpticalMonitoring) UnmarshalJSON(bytes []byte) (err error) {
	type NiatelemetryDigitalOpticalMonitoringWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Alerts count for the interface in the node.
		Alerts *string `json:"Alerts,omitempty"`
		// Dn with interface name for the aci nodes.
		Dn *string `json:"Dn,omitempty"`
		// RxLos count for the interface in the node.
		RxLos *string `json:"RxLos,omitempty"`
		// TxfaultCount for the interface in the node.
		TxFaultCount *string `json:"TxFaultCount,omitempty"`
	}

	varNiatelemetryDigitalOpticalMonitoringWithoutEmbeddedStruct := NiatelemetryDigitalOpticalMonitoringWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNiatelemetryDigitalOpticalMonitoringWithoutEmbeddedStruct)
	if err == nil {
		varNiatelemetryDigitalOpticalMonitoring := _NiatelemetryDigitalOpticalMonitoring{}
		varNiatelemetryDigitalOpticalMonitoring.ClassId = varNiatelemetryDigitalOpticalMonitoringWithoutEmbeddedStruct.ClassId
		varNiatelemetryDigitalOpticalMonitoring.ObjectType = varNiatelemetryDigitalOpticalMonitoringWithoutEmbeddedStruct.ObjectType
		varNiatelemetryDigitalOpticalMonitoring.Alerts = varNiatelemetryDigitalOpticalMonitoringWithoutEmbeddedStruct.Alerts
		varNiatelemetryDigitalOpticalMonitoring.Dn = varNiatelemetryDigitalOpticalMonitoringWithoutEmbeddedStruct.Dn
		varNiatelemetryDigitalOpticalMonitoring.RxLos = varNiatelemetryDigitalOpticalMonitoringWithoutEmbeddedStruct.RxLos
		varNiatelemetryDigitalOpticalMonitoring.TxFaultCount = varNiatelemetryDigitalOpticalMonitoringWithoutEmbeddedStruct.TxFaultCount
		*o = NiatelemetryDigitalOpticalMonitoring(varNiatelemetryDigitalOpticalMonitoring)
	} else {
		return err
	}

	varNiatelemetryDigitalOpticalMonitoring := _NiatelemetryDigitalOpticalMonitoring{}

	err = json.Unmarshal(bytes, &varNiatelemetryDigitalOpticalMonitoring)
	if err == nil {
		o.MoBaseComplexType = varNiatelemetryDigitalOpticalMonitoring.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Alerts")
		delete(additionalProperties, "Dn")
		delete(additionalProperties, "RxLos")
		delete(additionalProperties, "TxFaultCount")

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

type NullableNiatelemetryDigitalOpticalMonitoring struct {
	value *NiatelemetryDigitalOpticalMonitoring
	isSet bool
}

func (v NullableNiatelemetryDigitalOpticalMonitoring) Get() *NiatelemetryDigitalOpticalMonitoring {
	return v.value
}

func (v *NullableNiatelemetryDigitalOpticalMonitoring) Set(val *NiatelemetryDigitalOpticalMonitoring) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryDigitalOpticalMonitoring) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryDigitalOpticalMonitoring) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryDigitalOpticalMonitoring(val *NiatelemetryDigitalOpticalMonitoring) *NullableNiatelemetryDigitalOpticalMonitoring {
	return &NullableNiatelemetryDigitalOpticalMonitoring{value: val, isSet: true}
}

func (v NullableNiatelemetryDigitalOpticalMonitoring) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryDigitalOpticalMonitoring) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
