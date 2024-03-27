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
)

// KubernetesAbstractNodeAllOf Definition of the list of properties defined in 'kubernetes.AbstractNode', excluding properties defined in parent classes.
type KubernetesAbstractNodeAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Kubernetes metadata annotations for a Node.
	Annotations interface{} `json:"Annotations,omitempty"`
	// Kubernetes metadata labels for a Node.
	Labels               interface{}       `json:"Labels,omitempty"`
	Taints               []KubernetesTaint `json:"Taints,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesAbstractNodeAllOf KubernetesAbstractNodeAllOf

// NewKubernetesAbstractNodeAllOf instantiates a new KubernetesAbstractNodeAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesAbstractNodeAllOf(classId string, objectType string) *KubernetesAbstractNodeAllOf {
	this := KubernetesAbstractNodeAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKubernetesAbstractNodeAllOfWithDefaults instantiates a new KubernetesAbstractNodeAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesAbstractNodeAllOfWithDefaults() *KubernetesAbstractNodeAllOf {
	this := KubernetesAbstractNodeAllOf{}
	var classId string = "kubernetes.Node"
	this.ClassId = classId
	var objectType string = "kubernetes.Node"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesAbstractNodeAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesAbstractNodeAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesAbstractNodeAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesAbstractNodeAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesAbstractNodeAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesAbstractNodeAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAnnotations returns the Annotations field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesAbstractNodeAllOf) GetAnnotations() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesAbstractNodeAllOf) GetAnnotationsOk() (*interface{}, bool) {
	if o == nil || o.Annotations == nil {
		return nil, false
	}
	return &o.Annotations, true
}

// HasAnnotations returns a boolean if a field has been set.
func (o *KubernetesAbstractNodeAllOf) HasAnnotations() bool {
	if o != nil && o.Annotations != nil {
		return true
	}

	return false
}

// SetAnnotations gets a reference to the given interface{} and assigns it to the Annotations field.
func (o *KubernetesAbstractNodeAllOf) SetAnnotations(v interface{}) {
	o.Annotations = v
}

// GetLabels returns the Labels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesAbstractNodeAllOf) GetLabels() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesAbstractNodeAllOf) GetLabelsOk() (*interface{}, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return &o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *KubernetesAbstractNodeAllOf) HasLabels() bool {
	if o != nil && o.Labels != nil {
		return true
	}

	return false
}

// SetLabels gets a reference to the given interface{} and assigns it to the Labels field.
func (o *KubernetesAbstractNodeAllOf) SetLabels(v interface{}) {
	o.Labels = v
}

// GetTaints returns the Taints field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesAbstractNodeAllOf) GetTaints() []KubernetesTaint {
	if o == nil {
		var ret []KubernetesTaint
		return ret
	}
	return o.Taints
}

// GetTaintsOk returns a tuple with the Taints field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesAbstractNodeAllOf) GetTaintsOk() ([]KubernetesTaint, bool) {
	if o == nil || o.Taints == nil {
		return nil, false
	}
	return o.Taints, true
}

// HasTaints returns a boolean if a field has been set.
func (o *KubernetesAbstractNodeAllOf) HasTaints() bool {
	if o != nil && o.Taints != nil {
		return true
	}

	return false
}

// SetTaints gets a reference to the given []KubernetesTaint and assigns it to the Taints field.
func (o *KubernetesAbstractNodeAllOf) SetTaints(v []KubernetesTaint) {
	o.Taints = v
}

func (o KubernetesAbstractNodeAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Annotations != nil {
		toSerialize["Annotations"] = o.Annotations
	}
	if o.Labels != nil {
		toSerialize["Labels"] = o.Labels
	}
	if o.Taints != nil {
		toSerialize["Taints"] = o.Taints
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesAbstractNodeAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varKubernetesAbstractNodeAllOf := _KubernetesAbstractNodeAllOf{}

	if err = json.Unmarshal(bytes, &varKubernetesAbstractNodeAllOf); err == nil {
		*o = KubernetesAbstractNodeAllOf(varKubernetesAbstractNodeAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Annotations")
		delete(additionalProperties, "Labels")
		delete(additionalProperties, "Taints")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKubernetesAbstractNodeAllOf struct {
	value *KubernetesAbstractNodeAllOf
	isSet bool
}

func (v NullableKubernetesAbstractNodeAllOf) Get() *KubernetesAbstractNodeAllOf {
	return v.value
}

func (v *NullableKubernetesAbstractNodeAllOf) Set(val *KubernetesAbstractNodeAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesAbstractNodeAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesAbstractNodeAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesAbstractNodeAllOf(val *KubernetesAbstractNodeAllOf) *NullableKubernetesAbstractNodeAllOf {
	return &NullableKubernetesAbstractNodeAllOf{value: val, isSet: true}
}

func (v NullableKubernetesAbstractNodeAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesAbstractNodeAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
