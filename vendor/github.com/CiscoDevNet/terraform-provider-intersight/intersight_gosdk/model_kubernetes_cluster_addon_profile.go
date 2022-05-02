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

// KubernetesClusterAddonProfile A profile that describes a collection of Addons for a particular cluster.
type KubernetesClusterAddonProfile struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string            `json:"ObjectType"`
	Addons     []KubernetesAddon `json:"Addons,omitempty"`
	// Name of the cluster addon profile.
	Name                 *string                               `json:"Name,omitempty"`
	AssociatedCluster    *KubernetesClusterRelationship        `json:"AssociatedCluster,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesClusterAddonProfile KubernetesClusterAddonProfile

// NewKubernetesClusterAddonProfile instantiates a new KubernetesClusterAddonProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesClusterAddonProfile(classId string, objectType string) *KubernetesClusterAddonProfile {
	this := KubernetesClusterAddonProfile{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKubernetesClusterAddonProfileWithDefaults instantiates a new KubernetesClusterAddonProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesClusterAddonProfileWithDefaults() *KubernetesClusterAddonProfile {
	this := KubernetesClusterAddonProfile{}
	var classId string = "kubernetes.ClusterAddonProfile"
	this.ClassId = classId
	var objectType string = "kubernetes.ClusterAddonProfile"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesClusterAddonProfile) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesClusterAddonProfile) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesClusterAddonProfile) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesClusterAddonProfile) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesClusterAddonProfile) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesClusterAddonProfile) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAddons returns the Addons field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesClusterAddonProfile) GetAddons() []KubernetesAddon {
	if o == nil {
		var ret []KubernetesAddon
		return ret
	}
	return o.Addons
}

// GetAddonsOk returns a tuple with the Addons field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesClusterAddonProfile) GetAddonsOk() ([]KubernetesAddon, bool) {
	if o == nil || o.Addons == nil {
		return nil, false
	}
	return o.Addons, true
}

// HasAddons returns a boolean if a field has been set.
func (o *KubernetesClusterAddonProfile) HasAddons() bool {
	if o != nil && o.Addons != nil {
		return true
	}

	return false
}

// SetAddons gets a reference to the given []KubernetesAddon and assigns it to the Addons field.
func (o *KubernetesClusterAddonProfile) SetAddons(v []KubernetesAddon) {
	o.Addons = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *KubernetesClusterAddonProfile) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterAddonProfile) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *KubernetesClusterAddonProfile) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *KubernetesClusterAddonProfile) SetName(v string) {
	o.Name = &v
}

// GetAssociatedCluster returns the AssociatedCluster field value if set, zero value otherwise.
func (o *KubernetesClusterAddonProfile) GetAssociatedCluster() KubernetesClusterRelationship {
	if o == nil || o.AssociatedCluster == nil {
		var ret KubernetesClusterRelationship
		return ret
	}
	return *o.AssociatedCluster
}

// GetAssociatedClusterOk returns a tuple with the AssociatedCluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterAddonProfile) GetAssociatedClusterOk() (*KubernetesClusterRelationship, bool) {
	if o == nil || o.AssociatedCluster == nil {
		return nil, false
	}
	return o.AssociatedCluster, true
}

// HasAssociatedCluster returns a boolean if a field has been set.
func (o *KubernetesClusterAddonProfile) HasAssociatedCluster() bool {
	if o != nil && o.AssociatedCluster != nil {
		return true
	}

	return false
}

// SetAssociatedCluster gets a reference to the given KubernetesClusterRelationship and assigns it to the AssociatedCluster field.
func (o *KubernetesClusterAddonProfile) SetAssociatedCluster(v KubernetesClusterRelationship) {
	o.AssociatedCluster = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *KubernetesClusterAddonProfile) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterAddonProfile) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *KubernetesClusterAddonProfile) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *KubernetesClusterAddonProfile) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o KubernetesClusterAddonProfile) MarshalJSON() ([]byte, error) {
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
	if o.Addons != nil {
		toSerialize["Addons"] = o.Addons
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.AssociatedCluster != nil {
		toSerialize["AssociatedCluster"] = o.AssociatedCluster
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesClusterAddonProfile) UnmarshalJSON(bytes []byte) (err error) {
	type KubernetesClusterAddonProfileWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string            `json:"ObjectType"`
		Addons     []KubernetesAddon `json:"Addons,omitempty"`
		// Name of the cluster addon profile.
		Name              *string                               `json:"Name,omitempty"`
		AssociatedCluster *KubernetesClusterRelationship        `json:"AssociatedCluster,omitempty"`
		Organization      *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	}

	varKubernetesClusterAddonProfileWithoutEmbeddedStruct := KubernetesClusterAddonProfileWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varKubernetesClusterAddonProfileWithoutEmbeddedStruct)
	if err == nil {
		varKubernetesClusterAddonProfile := _KubernetesClusterAddonProfile{}
		varKubernetesClusterAddonProfile.ClassId = varKubernetesClusterAddonProfileWithoutEmbeddedStruct.ClassId
		varKubernetesClusterAddonProfile.ObjectType = varKubernetesClusterAddonProfileWithoutEmbeddedStruct.ObjectType
		varKubernetesClusterAddonProfile.Addons = varKubernetesClusterAddonProfileWithoutEmbeddedStruct.Addons
		varKubernetesClusterAddonProfile.Name = varKubernetesClusterAddonProfileWithoutEmbeddedStruct.Name
		varKubernetesClusterAddonProfile.AssociatedCluster = varKubernetesClusterAddonProfileWithoutEmbeddedStruct.AssociatedCluster
		varKubernetesClusterAddonProfile.Organization = varKubernetesClusterAddonProfileWithoutEmbeddedStruct.Organization
		*o = KubernetesClusterAddonProfile(varKubernetesClusterAddonProfile)
	} else {
		return err
	}

	varKubernetesClusterAddonProfile := _KubernetesClusterAddonProfile{}

	err = json.Unmarshal(bytes, &varKubernetesClusterAddonProfile)
	if err == nil {
		o.MoBaseMo = varKubernetesClusterAddonProfile.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Addons")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "AssociatedCluster")
		delete(additionalProperties, "Organization")

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

type NullableKubernetesClusterAddonProfile struct {
	value *KubernetesClusterAddonProfile
	isSet bool
}

func (v NullableKubernetesClusterAddonProfile) Get() *KubernetesClusterAddonProfile {
	return v.value
}

func (v *NullableKubernetesClusterAddonProfile) Set(val *KubernetesClusterAddonProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesClusterAddonProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesClusterAddonProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesClusterAddonProfile(val *KubernetesClusterAddonProfile) *NullableKubernetesClusterAddonProfile {
	return &NullableKubernetesClusterAddonProfile{value: val, isSet: true}
}

func (v NullableKubernetesClusterAddonProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesClusterAddonProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
