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

// KubernetesAbstractNode Abstract Node represents Kubernetes Node.
type KubernetesAbstractNode struct {
	KubernetesKubernetesResource
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

type _KubernetesAbstractNode KubernetesAbstractNode

// NewKubernetesAbstractNode instantiates a new KubernetesAbstractNode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesAbstractNode(classId string, objectType string) *KubernetesAbstractNode {
	this := KubernetesAbstractNode{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKubernetesAbstractNodeWithDefaults instantiates a new KubernetesAbstractNode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesAbstractNodeWithDefaults() *KubernetesAbstractNode {
	this := KubernetesAbstractNode{}
	var classId string = "kubernetes.Node"
	this.ClassId = classId
	var objectType string = "kubernetes.Node"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesAbstractNode) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesAbstractNode) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesAbstractNode) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesAbstractNode) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesAbstractNode) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesAbstractNode) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAnnotations returns the Annotations field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesAbstractNode) GetAnnotations() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesAbstractNode) GetAnnotationsOk() (*interface{}, bool) {
	if o == nil || o.Annotations == nil {
		return nil, false
	}
	return &o.Annotations, true
}

// HasAnnotations returns a boolean if a field has been set.
func (o *KubernetesAbstractNode) HasAnnotations() bool {
	if o != nil && o.Annotations != nil {
		return true
	}

	return false
}

// SetAnnotations gets a reference to the given interface{} and assigns it to the Annotations field.
func (o *KubernetesAbstractNode) SetAnnotations(v interface{}) {
	o.Annotations = v
}

// GetLabels returns the Labels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesAbstractNode) GetLabels() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesAbstractNode) GetLabelsOk() (*interface{}, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return &o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *KubernetesAbstractNode) HasLabels() bool {
	if o != nil && o.Labels != nil {
		return true
	}

	return false
}

// SetLabels gets a reference to the given interface{} and assigns it to the Labels field.
func (o *KubernetesAbstractNode) SetLabels(v interface{}) {
	o.Labels = v
}

// GetTaints returns the Taints field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesAbstractNode) GetTaints() []KubernetesTaint {
	if o == nil {
		var ret []KubernetesTaint
		return ret
	}
	return o.Taints
}

// GetTaintsOk returns a tuple with the Taints field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesAbstractNode) GetTaintsOk() ([]KubernetesTaint, bool) {
	if o == nil || o.Taints == nil {
		return nil, false
	}
	return o.Taints, true
}

// HasTaints returns a boolean if a field has been set.
func (o *KubernetesAbstractNode) HasTaints() bool {
	if o != nil && o.Taints != nil {
		return true
	}

	return false
}

// SetTaints gets a reference to the given []KubernetesTaint and assigns it to the Taints field.
func (o *KubernetesAbstractNode) SetTaints(v []KubernetesTaint) {
	o.Taints = v
}

func (o KubernetesAbstractNode) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKubernetesKubernetesResource, errKubernetesKubernetesResource := json.Marshal(o.KubernetesKubernetesResource)
	if errKubernetesKubernetesResource != nil {
		return []byte{}, errKubernetesKubernetesResource
	}
	errKubernetesKubernetesResource = json.Unmarshal([]byte(serializedKubernetesKubernetesResource), &toSerialize)
	if errKubernetesKubernetesResource != nil {
		return []byte{}, errKubernetesKubernetesResource
	}
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

func (o *KubernetesAbstractNode) UnmarshalJSON(bytes []byte) (err error) {
	type KubernetesAbstractNodeWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// Kubernetes metadata annotations for a Node.
		Annotations interface{} `json:"Annotations,omitempty"`
		// Kubernetes metadata labels for a Node.
		Labels interface{}       `json:"Labels,omitempty"`
		Taints []KubernetesTaint `json:"Taints,omitempty"`
	}

	varKubernetesAbstractNodeWithoutEmbeddedStruct := KubernetesAbstractNodeWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varKubernetesAbstractNodeWithoutEmbeddedStruct)
	if err == nil {
		varKubernetesAbstractNode := _KubernetesAbstractNode{}
		varKubernetesAbstractNode.ClassId = varKubernetesAbstractNodeWithoutEmbeddedStruct.ClassId
		varKubernetesAbstractNode.ObjectType = varKubernetesAbstractNodeWithoutEmbeddedStruct.ObjectType
		varKubernetesAbstractNode.Annotations = varKubernetesAbstractNodeWithoutEmbeddedStruct.Annotations
		varKubernetesAbstractNode.Labels = varKubernetesAbstractNodeWithoutEmbeddedStruct.Labels
		varKubernetesAbstractNode.Taints = varKubernetesAbstractNodeWithoutEmbeddedStruct.Taints
		*o = KubernetesAbstractNode(varKubernetesAbstractNode)
	} else {
		return err
	}

	varKubernetesAbstractNode := _KubernetesAbstractNode{}

	err = json.Unmarshal(bytes, &varKubernetesAbstractNode)
	if err == nil {
		o.KubernetesKubernetesResource = varKubernetesAbstractNode.KubernetesKubernetesResource
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Annotations")
		delete(additionalProperties, "Labels")
		delete(additionalProperties, "Taints")

		// remove fields from embedded structs
		reflectKubernetesKubernetesResource := reflect.ValueOf(o.KubernetesKubernetesResource)
		for i := 0; i < reflectKubernetesKubernetesResource.Type().NumField(); i++ {
			t := reflectKubernetesKubernetesResource.Type().Field(i)

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

type NullableKubernetesAbstractNode struct {
	value *KubernetesAbstractNode
	isSet bool
}

func (v NullableKubernetesAbstractNode) Get() *KubernetesAbstractNode {
	return v.value
}

func (v *NullableKubernetesAbstractNode) Set(val *KubernetesAbstractNode) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesAbstractNode) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesAbstractNode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesAbstractNode(val *KubernetesAbstractNode) *NullableKubernetesAbstractNode {
	return &NullableKubernetesAbstractNode{value: val, isSet: true}
}

func (v NullableKubernetesAbstractNode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesAbstractNode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
