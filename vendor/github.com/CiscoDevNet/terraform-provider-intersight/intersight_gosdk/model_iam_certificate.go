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
	"reflect"
	"strings"
)

// IamCertificate Holds a certificate, signed by a CAcert.
type IamCertificate struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType  string                  `json:"ObjectType"`
	Certificate NullableX509Certificate `json:"Certificate,omitempty"`
	// Status of the certificate. * `PendingValidation` - The certificate has not been validated. * `Valid` - The certificate is valid. * `Invalid` - Ther certificate is invalid.
	Status               *string                            `json:"Status,omitempty"`
	CertificateRequest   *IamCertificateRequestRelationship `json:"CertificateRequest,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamCertificate IamCertificate

// NewIamCertificate instantiates a new IamCertificate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamCertificate(classId string, objectType string) *IamCertificate {
	this := IamCertificate{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIamCertificateWithDefaults instantiates a new IamCertificate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamCertificateWithDefaults() *IamCertificate {
	this := IamCertificate{}
	var classId string = "iam.Certificate"
	this.ClassId = classId
	var objectType string = "iam.Certificate"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IamCertificate) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IamCertificate) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IamCertificate) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IamCertificate) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IamCertificate) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IamCertificate) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCertificate returns the Certificate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamCertificate) GetCertificate() X509Certificate {
	if o == nil || o.Certificate.Get() == nil {
		var ret X509Certificate
		return ret
	}
	return *o.Certificate.Get()
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamCertificate) GetCertificateOk() (*X509Certificate, bool) {
	if o == nil {
		return nil, false
	}
	return o.Certificate.Get(), o.Certificate.IsSet()
}

// HasCertificate returns a boolean if a field has been set.
func (o *IamCertificate) HasCertificate() bool {
	if o != nil && o.Certificate.IsSet() {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given NullableX509Certificate and assigns it to the Certificate field.
func (o *IamCertificate) SetCertificate(v X509Certificate) {
	o.Certificate.Set(&v)
}

// SetCertificateNil sets the value for Certificate to be an explicit nil
func (o *IamCertificate) SetCertificateNil() {
	o.Certificate.Set(nil)
}

// UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
func (o *IamCertificate) UnsetCertificate() {
	o.Certificate.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *IamCertificate) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamCertificate) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *IamCertificate) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *IamCertificate) SetStatus(v string) {
	o.Status = &v
}

// GetCertificateRequest returns the CertificateRequest field value if set, zero value otherwise.
func (o *IamCertificate) GetCertificateRequest() IamCertificateRequestRelationship {
	if o == nil || o.CertificateRequest == nil {
		var ret IamCertificateRequestRelationship
		return ret
	}
	return *o.CertificateRequest
}

// GetCertificateRequestOk returns a tuple with the CertificateRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamCertificate) GetCertificateRequestOk() (*IamCertificateRequestRelationship, bool) {
	if o == nil || o.CertificateRequest == nil {
		return nil, false
	}
	return o.CertificateRequest, true
}

// HasCertificateRequest returns a boolean if a field has been set.
func (o *IamCertificate) HasCertificateRequest() bool {
	if o != nil && o.CertificateRequest != nil {
		return true
	}

	return false
}

// SetCertificateRequest gets a reference to the given IamCertificateRequestRelationship and assigns it to the CertificateRequest field.
func (o *IamCertificate) SetCertificateRequest(v IamCertificateRequestRelationship) {
	o.CertificateRequest = &v
}

func (o IamCertificate) MarshalJSON() ([]byte, error) {
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
	if o.Certificate.IsSet() {
		toSerialize["Certificate"] = o.Certificate.Get()
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.CertificateRequest != nil {
		toSerialize["CertificateRequest"] = o.CertificateRequest
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IamCertificate) UnmarshalJSON(bytes []byte) (err error) {
	type IamCertificateWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType  string                  `json:"ObjectType"`
		Certificate NullableX509Certificate `json:"Certificate,omitempty"`
		// Status of the certificate. * `PendingValidation` - The certificate has not been validated. * `Valid` - The certificate is valid. * `Invalid` - Ther certificate is invalid.
		Status             *string                            `json:"Status,omitempty"`
		CertificateRequest *IamCertificateRequestRelationship `json:"CertificateRequest,omitempty"`
	}

	varIamCertificateWithoutEmbeddedStruct := IamCertificateWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varIamCertificateWithoutEmbeddedStruct)
	if err == nil {
		varIamCertificate := _IamCertificate{}
		varIamCertificate.ClassId = varIamCertificateWithoutEmbeddedStruct.ClassId
		varIamCertificate.ObjectType = varIamCertificateWithoutEmbeddedStruct.ObjectType
		varIamCertificate.Certificate = varIamCertificateWithoutEmbeddedStruct.Certificate
		varIamCertificate.Status = varIamCertificateWithoutEmbeddedStruct.Status
		varIamCertificate.CertificateRequest = varIamCertificateWithoutEmbeddedStruct.CertificateRequest
		*o = IamCertificate(varIamCertificate)
	} else {
		return err
	}

	varIamCertificate := _IamCertificate{}

	err = json.Unmarshal(bytes, &varIamCertificate)
	if err == nil {
		o.MoBaseMo = varIamCertificate.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Certificate")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "CertificateRequest")

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

type NullableIamCertificate struct {
	value *IamCertificate
	isSet bool
}

func (v NullableIamCertificate) Get() *IamCertificate {
	return v.value
}

func (v *NullableIamCertificate) Set(val *IamCertificate) {
	v.value = val
	v.isSet = true
}

func (v NullableIamCertificate) IsSet() bool {
	return v.isSet
}

func (v *NullableIamCertificate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamCertificate(val *IamCertificate) *NullableIamCertificate {
	return &NullableIamCertificate{value: val, isSet: true}
}

func (v NullableIamCertificate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamCertificate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
