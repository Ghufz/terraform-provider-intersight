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
)

// NiatelemetryNodeAllOf Definition of the list of properties defined in 'niatelemetry.Node', excluding properties defined in parent classes.
type NiatelemetryNodeAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Returns hostname of the node.
	Hostname *string `json:"Hostname,omitempty"`
	// Returns management IP of the node.
	ManagementtIp *string `json:"ManagementtIp,omitempty"`
	// Returns out of band IP of the node.
	OutofbandIp          *string `json:"OutofbandIp,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryNodeAllOf NiatelemetryNodeAllOf

// NewNiatelemetryNodeAllOf instantiates a new NiatelemetryNodeAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryNodeAllOf(classId string, objectType string) *NiatelemetryNodeAllOf {
	this := NiatelemetryNodeAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryNodeAllOfWithDefaults instantiates a new NiatelemetryNodeAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryNodeAllOfWithDefaults() *NiatelemetryNodeAllOf {
	this := NiatelemetryNodeAllOf{}
	var classId string = "niatelemetry.Node"
	this.ClassId = classId
	var objectType string = "niatelemetry.Node"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryNodeAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryNodeAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryNodeAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryNodeAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryNodeAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryNodeAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *NiatelemetryNodeAllOf) GetHostname() string {
	if o == nil || o.Hostname == nil {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNodeAllOf) GetHostnameOk() (*string, bool) {
	if o == nil || o.Hostname == nil {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *NiatelemetryNodeAllOf) HasHostname() bool {
	if o != nil && o.Hostname != nil {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *NiatelemetryNodeAllOf) SetHostname(v string) {
	o.Hostname = &v
}

// GetManagementtIp returns the ManagementtIp field value if set, zero value otherwise.
func (o *NiatelemetryNodeAllOf) GetManagementtIp() string {
	if o == nil || o.ManagementtIp == nil {
		var ret string
		return ret
	}
	return *o.ManagementtIp
}

// GetManagementtIpOk returns a tuple with the ManagementtIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNodeAllOf) GetManagementtIpOk() (*string, bool) {
	if o == nil || o.ManagementtIp == nil {
		return nil, false
	}
	return o.ManagementtIp, true
}

// HasManagementtIp returns a boolean if a field has been set.
func (o *NiatelemetryNodeAllOf) HasManagementtIp() bool {
	if o != nil && o.ManagementtIp != nil {
		return true
	}

	return false
}

// SetManagementtIp gets a reference to the given string and assigns it to the ManagementtIp field.
func (o *NiatelemetryNodeAllOf) SetManagementtIp(v string) {
	o.ManagementtIp = &v
}

// GetOutofbandIp returns the OutofbandIp field value if set, zero value otherwise.
func (o *NiatelemetryNodeAllOf) GetOutofbandIp() string {
	if o == nil || o.OutofbandIp == nil {
		var ret string
		return ret
	}
	return *o.OutofbandIp
}

// GetOutofbandIpOk returns a tuple with the OutofbandIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNodeAllOf) GetOutofbandIpOk() (*string, bool) {
	if o == nil || o.OutofbandIp == nil {
		return nil, false
	}
	return o.OutofbandIp, true
}

// HasOutofbandIp returns a boolean if a field has been set.
func (o *NiatelemetryNodeAllOf) HasOutofbandIp() bool {
	if o != nil && o.OutofbandIp != nil {
		return true
	}

	return false
}

// SetOutofbandIp gets a reference to the given string and assigns it to the OutofbandIp field.
func (o *NiatelemetryNodeAllOf) SetOutofbandIp(v string) {
	o.OutofbandIp = &v
}

func (o NiatelemetryNodeAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Hostname != nil {
		toSerialize["Hostname"] = o.Hostname
	}
	if o.ManagementtIp != nil {
		toSerialize["ManagementtIp"] = o.ManagementtIp
	}
	if o.OutofbandIp != nil {
		toSerialize["OutofbandIp"] = o.OutofbandIp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryNodeAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varNiatelemetryNodeAllOf := _NiatelemetryNodeAllOf{}

	if err = json.Unmarshal(bytes, &varNiatelemetryNodeAllOf); err == nil {
		*o = NiatelemetryNodeAllOf(varNiatelemetryNodeAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Hostname")
		delete(additionalProperties, "ManagementtIp")
		delete(additionalProperties, "OutofbandIp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNiatelemetryNodeAllOf struct {
	value *NiatelemetryNodeAllOf
	isSet bool
}

func (v NullableNiatelemetryNodeAllOf) Get() *NiatelemetryNodeAllOf {
	return v.value
}

func (v *NullableNiatelemetryNodeAllOf) Set(val *NiatelemetryNodeAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryNodeAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryNodeAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryNodeAllOf(val *NiatelemetryNodeAllOf) *NullableNiatelemetryNodeAllOf {
	return &NullableNiatelemetryNodeAllOf{value: val, isSet: true}
}

func (v NullableNiatelemetryNodeAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryNodeAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
