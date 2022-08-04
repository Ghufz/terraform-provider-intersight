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

// IamApiKey An API key is used to authenticate and authorize API requests sent by a client using the HTTP signature scheme. API keys can be used by unattended, daemon clients that need to send requests to Intersight programmatically. API keys are based on public key cryptography. To create an API key, the user must specify: 1. The purpose (description) of the API key, 2. The cryptographic hash algorithm, which is used to compute the digest of the body of HTTP requests, 3. The cryptographic parameters to generate a private/public key pair, e.g. RSA, ECDSA, EDDSA, key modulus, and 4. The signing algorithm, e.g. RSA PKCS v1.5, RSA PSS, ECDSA, EDDSA. The generated private key and public key are encoded in PEM format. The client owns the private key and is responsible for maintaining the confidentiality of the private key. The server holds the public key. The client must have a cryptographic provider compatible with the cryptographic parameters specified in the API key. For example, if you use the powershell SDK to write the client, make sure the appropriate cryptographic providers are installed on the local system. If you create an RSA key pair with modulus set to 2048, the client must support 2048-bit private keys. A maximum of 3 API keys per user is allowed. API keys are used to sign HTTP requests as follows: 1. A cryptographic digest of the body of the HTTP request is calculated using one of the supported cryptographic hash algorithms. 2. The value of the digest is base-64 encoded in the `Digest` HTTP header. 3. A signature is calculated as specified in the HTTP signature scheme, and the signature is added to the `Authorization` HTTP request header. All published Intersight SDKs support API keys.
type IamApiKey struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The cryptographic hash algorithm to calculate the message digest. * `SHA256` - The SHA-256 cryptographic hash, as defined by NIST in FIPS 180-4. * `SHA384` - The SHA-384 cryptographic hash, as defined by NIST in FIPS 180-4. * `SHA512` - The SHA-512 cryptographic hash, as defined by NIST in FIPS 180-4. * `SHA512_224` - The SHA-512/224 cryptographic hash, as defined by NIST in FIPS 180-4. * `SHA512_256` - The SHA-512/256 cryptographic hash, as defined by NIST in FIPS 180-4.
	HashAlgorithm *string                       `json:"HashAlgorithm,omitempty"`
	KeySpec       NullablePkixKeyGenerationSpec `json:"KeySpec,omitempty"`
	// Holds the private key for the API key.
	PrivateKey *string `json:"PrivateKey,omitempty"`
	// The purpose of the API Key.
	Purpose *string `json:"Purpose,omitempty"`
	// The signing algorithm used by the client to authenticate API requests to Intersight. The signing algorithm must be compatible with the key generation specification. * `RSASSA-PKCS1-v1_5` - RSASSA-PKCS1-v1_5 is a RSA signature scheme specified in [RFC 8017](https://tools.ietf.org/html/rfc8017).RSASSA-PKCS1-v1_5 is included only for compatibility with existing applications. * `RSASSA-PSS` - RSASSA-PSS is a RSA signature scheme specified in [RFC 8017](https://tools.ietf.org/html/rfc8017).It combines the RSASP1 and RSAVP1 primitives with the EMSA-PSS encoding method.In the interest of increased robustness, RSASSA-PSS is required in new applications. * `Ed25519` - The Ed25519 signature algorithm, as specified in [RFC 8032](https://tools.ietf.org/html/rfc8032).Ed25519 is a public-key signature system with several attractive features, includingfast single-signature verification, very fast signing, fast key generation and high security level. * `Ecdsa` - The Elliptic Curve Digital Signature Standard (ECDSA), as defined by NIST in FIPS 186-4 and ANSI X9.62.The signature is encoded as a ASN.1 DER SEQUENCE with two INTEGERs (r and s), as defined in RFC3279.When using ECDSA signatures, configure the client to use the same signature encoding as specified on the server side. * `EcdsaP1363Format` - The Elliptic Curve Digital Signature Standard (ECDSA), as defined by NIST in FIPS 186-4 and ANSI X9.62.The signature is the raw concatenation of r and s, as defined in the ISO/IEC 7816-8 IEEE P.1363 standard.In that format, r and s are represented as unsigned, big endian numbers.Extra padding bytes (of value 0x00) is applied so that both r and s encodings have the same size.When using ECDSA signatures, configure the client to use the same signature encoding as specified on the server side.
	SigningAlgorithm     *string                    `json:"SigningAlgorithm,omitempty"`
	Permission           *IamPermissionRelationship `json:"Permission,omitempty"`
	User                 *IamUserRelationship       `json:"User,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamApiKey IamApiKey

// NewIamApiKey instantiates a new IamApiKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamApiKey(classId string, objectType string) *IamApiKey {
	this := IamApiKey{}
	this.ClassId = classId
	this.ObjectType = objectType
	var hashAlgorithm string = "SHA256"
	this.HashAlgorithm = &hashAlgorithm
	var signingAlgorithm string = "RSASSA-PKCS1-v1_5"
	this.SigningAlgorithm = &signingAlgorithm
	return &this
}

// NewIamApiKeyWithDefaults instantiates a new IamApiKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamApiKeyWithDefaults() *IamApiKey {
	this := IamApiKey{}
	var classId string = "iam.ApiKey"
	this.ClassId = classId
	var objectType string = "iam.ApiKey"
	this.ObjectType = objectType
	var hashAlgorithm string = "SHA256"
	this.HashAlgorithm = &hashAlgorithm
	var signingAlgorithm string = "RSASSA-PKCS1-v1_5"
	this.SigningAlgorithm = &signingAlgorithm
	return &this
}

// GetClassId returns the ClassId field value
func (o *IamApiKey) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IamApiKey) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IamApiKey) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IamApiKey) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IamApiKey) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IamApiKey) SetObjectType(v string) {
	o.ObjectType = v
}

// GetHashAlgorithm returns the HashAlgorithm field value if set, zero value otherwise.
func (o *IamApiKey) GetHashAlgorithm() string {
	if o == nil || o.HashAlgorithm == nil {
		var ret string
		return ret
	}
	return *o.HashAlgorithm
}

// GetHashAlgorithmOk returns a tuple with the HashAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamApiKey) GetHashAlgorithmOk() (*string, bool) {
	if o == nil || o.HashAlgorithm == nil {
		return nil, false
	}
	return o.HashAlgorithm, true
}

// HasHashAlgorithm returns a boolean if a field has been set.
func (o *IamApiKey) HasHashAlgorithm() bool {
	if o != nil && o.HashAlgorithm != nil {
		return true
	}

	return false
}

// SetHashAlgorithm gets a reference to the given string and assigns it to the HashAlgorithm field.
func (o *IamApiKey) SetHashAlgorithm(v string) {
	o.HashAlgorithm = &v
}

// GetKeySpec returns the KeySpec field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamApiKey) GetKeySpec() PkixKeyGenerationSpec {
	if o == nil || o.KeySpec.Get() == nil {
		var ret PkixKeyGenerationSpec
		return ret
	}
	return *o.KeySpec.Get()
}

// GetKeySpecOk returns a tuple with the KeySpec field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamApiKey) GetKeySpecOk() (*PkixKeyGenerationSpec, bool) {
	if o == nil {
		return nil, false
	}
	return o.KeySpec.Get(), o.KeySpec.IsSet()
}

// HasKeySpec returns a boolean if a field has been set.
func (o *IamApiKey) HasKeySpec() bool {
	if o != nil && o.KeySpec.IsSet() {
		return true
	}

	return false
}

// SetKeySpec gets a reference to the given NullablePkixKeyGenerationSpec and assigns it to the KeySpec field.
func (o *IamApiKey) SetKeySpec(v PkixKeyGenerationSpec) {
	o.KeySpec.Set(&v)
}

// SetKeySpecNil sets the value for KeySpec to be an explicit nil
func (o *IamApiKey) SetKeySpecNil() {
	o.KeySpec.Set(nil)
}

// UnsetKeySpec ensures that no value is present for KeySpec, not even an explicit nil
func (o *IamApiKey) UnsetKeySpec() {
	o.KeySpec.Unset()
}

// GetPrivateKey returns the PrivateKey field value if set, zero value otherwise.
func (o *IamApiKey) GetPrivateKey() string {
	if o == nil || o.PrivateKey == nil {
		var ret string
		return ret
	}
	return *o.PrivateKey
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamApiKey) GetPrivateKeyOk() (*string, bool) {
	if o == nil || o.PrivateKey == nil {
		return nil, false
	}
	return o.PrivateKey, true
}

// HasPrivateKey returns a boolean if a field has been set.
func (o *IamApiKey) HasPrivateKey() bool {
	if o != nil && o.PrivateKey != nil {
		return true
	}

	return false
}

// SetPrivateKey gets a reference to the given string and assigns it to the PrivateKey field.
func (o *IamApiKey) SetPrivateKey(v string) {
	o.PrivateKey = &v
}

// GetPurpose returns the Purpose field value if set, zero value otherwise.
func (o *IamApiKey) GetPurpose() string {
	if o == nil || o.Purpose == nil {
		var ret string
		return ret
	}
	return *o.Purpose
}

// GetPurposeOk returns a tuple with the Purpose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamApiKey) GetPurposeOk() (*string, bool) {
	if o == nil || o.Purpose == nil {
		return nil, false
	}
	return o.Purpose, true
}

// HasPurpose returns a boolean if a field has been set.
func (o *IamApiKey) HasPurpose() bool {
	if o != nil && o.Purpose != nil {
		return true
	}

	return false
}

// SetPurpose gets a reference to the given string and assigns it to the Purpose field.
func (o *IamApiKey) SetPurpose(v string) {
	o.Purpose = &v
}

// GetSigningAlgorithm returns the SigningAlgorithm field value if set, zero value otherwise.
func (o *IamApiKey) GetSigningAlgorithm() string {
	if o == nil || o.SigningAlgorithm == nil {
		var ret string
		return ret
	}
	return *o.SigningAlgorithm
}

// GetSigningAlgorithmOk returns a tuple with the SigningAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamApiKey) GetSigningAlgorithmOk() (*string, bool) {
	if o == nil || o.SigningAlgorithm == nil {
		return nil, false
	}
	return o.SigningAlgorithm, true
}

// HasSigningAlgorithm returns a boolean if a field has been set.
func (o *IamApiKey) HasSigningAlgorithm() bool {
	if o != nil && o.SigningAlgorithm != nil {
		return true
	}

	return false
}

// SetSigningAlgorithm gets a reference to the given string and assigns it to the SigningAlgorithm field.
func (o *IamApiKey) SetSigningAlgorithm(v string) {
	o.SigningAlgorithm = &v
}

// GetPermission returns the Permission field value if set, zero value otherwise.
func (o *IamApiKey) GetPermission() IamPermissionRelationship {
	if o == nil || o.Permission == nil {
		var ret IamPermissionRelationship
		return ret
	}
	return *o.Permission
}

// GetPermissionOk returns a tuple with the Permission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamApiKey) GetPermissionOk() (*IamPermissionRelationship, bool) {
	if o == nil || o.Permission == nil {
		return nil, false
	}
	return o.Permission, true
}

// HasPermission returns a boolean if a field has been set.
func (o *IamApiKey) HasPermission() bool {
	if o != nil && o.Permission != nil {
		return true
	}

	return false
}

// SetPermission gets a reference to the given IamPermissionRelationship and assigns it to the Permission field.
func (o *IamApiKey) SetPermission(v IamPermissionRelationship) {
	o.Permission = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *IamApiKey) GetUser() IamUserRelationship {
	if o == nil || o.User == nil {
		var ret IamUserRelationship
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamApiKey) GetUserOk() (*IamUserRelationship, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *IamApiKey) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given IamUserRelationship and assigns it to the User field.
func (o *IamApiKey) SetUser(v IamUserRelationship) {
	o.User = &v
}

func (o IamApiKey) MarshalJSON() ([]byte, error) {
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
	if o.HashAlgorithm != nil {
		toSerialize["HashAlgorithm"] = o.HashAlgorithm
	}
	if o.KeySpec.IsSet() {
		toSerialize["KeySpec"] = o.KeySpec.Get()
	}
	if o.PrivateKey != nil {
		toSerialize["PrivateKey"] = o.PrivateKey
	}
	if o.Purpose != nil {
		toSerialize["Purpose"] = o.Purpose
	}
	if o.SigningAlgorithm != nil {
		toSerialize["SigningAlgorithm"] = o.SigningAlgorithm
	}
	if o.Permission != nil {
		toSerialize["Permission"] = o.Permission
	}
	if o.User != nil {
		toSerialize["User"] = o.User
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IamApiKey) UnmarshalJSON(bytes []byte) (err error) {
	type IamApiKeyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The cryptographic hash algorithm to calculate the message digest. * `SHA256` - The SHA-256 cryptographic hash, as defined by NIST in FIPS 180-4. * `SHA384` - The SHA-384 cryptographic hash, as defined by NIST in FIPS 180-4. * `SHA512` - The SHA-512 cryptographic hash, as defined by NIST in FIPS 180-4. * `SHA512_224` - The SHA-512/224 cryptographic hash, as defined by NIST in FIPS 180-4. * `SHA512_256` - The SHA-512/256 cryptographic hash, as defined by NIST in FIPS 180-4.
		HashAlgorithm *string                       `json:"HashAlgorithm,omitempty"`
		KeySpec       NullablePkixKeyGenerationSpec `json:"KeySpec,omitempty"`
		// Holds the private key for the API key.
		PrivateKey *string `json:"PrivateKey,omitempty"`
		// The purpose of the API Key.
		Purpose *string `json:"Purpose,omitempty"`
		// The signing algorithm used by the client to authenticate API requests to Intersight. The signing algorithm must be compatible with the key generation specification. * `RSASSA-PKCS1-v1_5` - RSASSA-PKCS1-v1_5 is a RSA signature scheme specified in [RFC 8017](https://tools.ietf.org/html/rfc8017).RSASSA-PKCS1-v1_5 is included only for compatibility with existing applications. * `RSASSA-PSS` - RSASSA-PSS is a RSA signature scheme specified in [RFC 8017](https://tools.ietf.org/html/rfc8017).It combines the RSASP1 and RSAVP1 primitives with the EMSA-PSS encoding method.In the interest of increased robustness, RSASSA-PSS is required in new applications. * `Ed25519` - The Ed25519 signature algorithm, as specified in [RFC 8032](https://tools.ietf.org/html/rfc8032).Ed25519 is a public-key signature system with several attractive features, includingfast single-signature verification, very fast signing, fast key generation and high security level. * `Ecdsa` - The Elliptic Curve Digital Signature Standard (ECDSA), as defined by NIST in FIPS 186-4 and ANSI X9.62.The signature is encoded as a ASN.1 DER SEQUENCE with two INTEGERs (r and s), as defined in RFC3279.When using ECDSA signatures, configure the client to use the same signature encoding as specified on the server side. * `EcdsaP1363Format` - The Elliptic Curve Digital Signature Standard (ECDSA), as defined by NIST in FIPS 186-4 and ANSI X9.62.The signature is the raw concatenation of r and s, as defined in the ISO/IEC 7816-8 IEEE P.1363 standard.In that format, r and s are represented as unsigned, big endian numbers.Extra padding bytes (of value 0x00) is applied so that both r and s encodings have the same size.When using ECDSA signatures, configure the client to use the same signature encoding as specified on the server side.
		SigningAlgorithm *string                    `json:"SigningAlgorithm,omitempty"`
		Permission       *IamPermissionRelationship `json:"Permission,omitempty"`
		User             *IamUserRelationship       `json:"User,omitempty"`
	}

	varIamApiKeyWithoutEmbeddedStruct := IamApiKeyWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varIamApiKeyWithoutEmbeddedStruct)
	if err == nil {
		varIamApiKey := _IamApiKey{}
		varIamApiKey.ClassId = varIamApiKeyWithoutEmbeddedStruct.ClassId
		varIamApiKey.ObjectType = varIamApiKeyWithoutEmbeddedStruct.ObjectType
		varIamApiKey.HashAlgorithm = varIamApiKeyWithoutEmbeddedStruct.HashAlgorithm
		varIamApiKey.KeySpec = varIamApiKeyWithoutEmbeddedStruct.KeySpec
		varIamApiKey.PrivateKey = varIamApiKeyWithoutEmbeddedStruct.PrivateKey
		varIamApiKey.Purpose = varIamApiKeyWithoutEmbeddedStruct.Purpose
		varIamApiKey.SigningAlgorithm = varIamApiKeyWithoutEmbeddedStruct.SigningAlgorithm
		varIamApiKey.Permission = varIamApiKeyWithoutEmbeddedStruct.Permission
		varIamApiKey.User = varIamApiKeyWithoutEmbeddedStruct.User
		*o = IamApiKey(varIamApiKey)
	} else {
		return err
	}

	varIamApiKey := _IamApiKey{}

	err = json.Unmarshal(bytes, &varIamApiKey)
	if err == nil {
		o.MoBaseMo = varIamApiKey.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "HashAlgorithm")
		delete(additionalProperties, "KeySpec")
		delete(additionalProperties, "PrivateKey")
		delete(additionalProperties, "Purpose")
		delete(additionalProperties, "SigningAlgorithm")
		delete(additionalProperties, "Permission")
		delete(additionalProperties, "User")

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

type NullableIamApiKey struct {
	value *IamApiKey
	isSet bool
}

func (v NullableIamApiKey) Get() *IamApiKey {
	return v.value
}

func (v *NullableIamApiKey) Set(val *IamApiKey) {
	v.value = val
	v.isSet = true
}

func (v NullableIamApiKey) IsSet() bool {
	return v.isSet
}

func (v *NullableIamApiKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamApiKey(val *IamApiKey) *NullableIamApiKey {
	return &NullableIamApiKey{value: val, isSet: true}
}

func (v NullableIamApiKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamApiKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
