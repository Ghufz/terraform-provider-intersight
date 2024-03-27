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

// WorkflowServiceItemActionProperties Definition to capture the action properties which will be used in the action of service item.
type WorkflowServiceItemActionProperties struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Type of action operation to be executed on the service item. * `PostDeployment` - This represents the post-deployment actions for the resources created or defined through the deployment action. There can be more than one post-deployment operations associated with a service item. * `Deployment` - This represents the deploy action, for the service item action definition. This operation type is used to create or define resources that is managed by the service item. There can only be one Service Item Action Definition that can be marked with the operation type as Deployment and this is a mandatory operation type. All valid Service Items must have one and only one operation type marked as type Deployment. * `Decommission` - This represents the decommission action, used to decommission the created resources. All valid Service Items must have one and only one operation type marked as type Decommission. Once a decommission action is run on a Service Item, no further operations are allowed on that Service Item.
	OperationType *string `json:"OperationType,omitempty"`
	// When true, the action on the service item will be stopped when it reaches a failure by either calling the configured stop workflow or by calling the rollback workflow. By default value is set to true.
	StopOnFailure        *bool `json:"StopOnFailure,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowServiceItemActionProperties WorkflowServiceItemActionProperties

// NewWorkflowServiceItemActionProperties instantiates a new WorkflowServiceItemActionProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowServiceItemActionProperties(classId string, objectType string) *WorkflowServiceItemActionProperties {
	this := WorkflowServiceItemActionProperties{}
	this.ClassId = classId
	this.ObjectType = objectType
	var operationType string = "PostDeployment"
	this.OperationType = &operationType
	var stopOnFailure bool = true
	this.StopOnFailure = &stopOnFailure
	return &this
}

// NewWorkflowServiceItemActionPropertiesWithDefaults instantiates a new WorkflowServiceItemActionProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowServiceItemActionPropertiesWithDefaults() *WorkflowServiceItemActionProperties {
	this := WorkflowServiceItemActionProperties{}
	var classId string = "workflow.ServiceItemActionProperties"
	this.ClassId = classId
	var objectType string = "workflow.ServiceItemActionProperties"
	this.ObjectType = objectType
	var operationType string = "PostDeployment"
	this.OperationType = &operationType
	var stopOnFailure bool = true
	this.StopOnFailure = &stopOnFailure
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowServiceItemActionProperties) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemActionProperties) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowServiceItemActionProperties) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowServiceItemActionProperties) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemActionProperties) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowServiceItemActionProperties) SetObjectType(v string) {
	o.ObjectType = v
}

// GetOperationType returns the OperationType field value if set, zero value otherwise.
func (o *WorkflowServiceItemActionProperties) GetOperationType() string {
	if o == nil || o.OperationType == nil {
		var ret string
		return ret
	}
	return *o.OperationType
}

// GetOperationTypeOk returns a tuple with the OperationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemActionProperties) GetOperationTypeOk() (*string, bool) {
	if o == nil || o.OperationType == nil {
		return nil, false
	}
	return o.OperationType, true
}

// HasOperationType returns a boolean if a field has been set.
func (o *WorkflowServiceItemActionProperties) HasOperationType() bool {
	if o != nil && o.OperationType != nil {
		return true
	}

	return false
}

// SetOperationType gets a reference to the given string and assigns it to the OperationType field.
func (o *WorkflowServiceItemActionProperties) SetOperationType(v string) {
	o.OperationType = &v
}

// GetStopOnFailure returns the StopOnFailure field value if set, zero value otherwise.
func (o *WorkflowServiceItemActionProperties) GetStopOnFailure() bool {
	if o == nil || o.StopOnFailure == nil {
		var ret bool
		return ret
	}
	return *o.StopOnFailure
}

// GetStopOnFailureOk returns a tuple with the StopOnFailure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemActionProperties) GetStopOnFailureOk() (*bool, bool) {
	if o == nil || o.StopOnFailure == nil {
		return nil, false
	}
	return o.StopOnFailure, true
}

// HasStopOnFailure returns a boolean if a field has been set.
func (o *WorkflowServiceItemActionProperties) HasStopOnFailure() bool {
	if o != nil && o.StopOnFailure != nil {
		return true
	}

	return false
}

// SetStopOnFailure gets a reference to the given bool and assigns it to the StopOnFailure field.
func (o *WorkflowServiceItemActionProperties) SetStopOnFailure(v bool) {
	o.StopOnFailure = &v
}

func (o WorkflowServiceItemActionProperties) MarshalJSON() ([]byte, error) {
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
	if o.OperationType != nil {
		toSerialize["OperationType"] = o.OperationType
	}
	if o.StopOnFailure != nil {
		toSerialize["StopOnFailure"] = o.StopOnFailure
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowServiceItemActionProperties) UnmarshalJSON(bytes []byte) (err error) {
	type WorkflowServiceItemActionPropertiesWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Type of action operation to be executed on the service item. * `PostDeployment` - This represents the post-deployment actions for the resources created or defined through the deployment action. There can be more than one post-deployment operations associated with a service item. * `Deployment` - This represents the deploy action, for the service item action definition. This operation type is used to create or define resources that is managed by the service item. There can only be one Service Item Action Definition that can be marked with the operation type as Deployment and this is a mandatory operation type. All valid Service Items must have one and only one operation type marked as type Deployment. * `Decommission` - This represents the decommission action, used to decommission the created resources. All valid Service Items must have one and only one operation type marked as type Decommission. Once a decommission action is run on a Service Item, no further operations are allowed on that Service Item.
		OperationType *string `json:"OperationType,omitempty"`
		// When true, the action on the service item will be stopped when it reaches a failure by either calling the configured stop workflow or by calling the rollback workflow. By default value is set to true.
		StopOnFailure *bool `json:"StopOnFailure,omitempty"`
	}

	varWorkflowServiceItemActionPropertiesWithoutEmbeddedStruct := WorkflowServiceItemActionPropertiesWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varWorkflowServiceItemActionPropertiesWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowServiceItemActionProperties := _WorkflowServiceItemActionProperties{}
		varWorkflowServiceItemActionProperties.ClassId = varWorkflowServiceItemActionPropertiesWithoutEmbeddedStruct.ClassId
		varWorkflowServiceItemActionProperties.ObjectType = varWorkflowServiceItemActionPropertiesWithoutEmbeddedStruct.ObjectType
		varWorkflowServiceItemActionProperties.OperationType = varWorkflowServiceItemActionPropertiesWithoutEmbeddedStruct.OperationType
		varWorkflowServiceItemActionProperties.StopOnFailure = varWorkflowServiceItemActionPropertiesWithoutEmbeddedStruct.StopOnFailure
		*o = WorkflowServiceItemActionProperties(varWorkflowServiceItemActionProperties)
	} else {
		return err
	}

	varWorkflowServiceItemActionProperties := _WorkflowServiceItemActionProperties{}

	err = json.Unmarshal(bytes, &varWorkflowServiceItemActionProperties)
	if err == nil {
		o.MoBaseComplexType = varWorkflowServiceItemActionProperties.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "OperationType")
		delete(additionalProperties, "StopOnFailure")

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

type NullableWorkflowServiceItemActionProperties struct {
	value *WorkflowServiceItemActionProperties
	isSet bool
}

func (v NullableWorkflowServiceItemActionProperties) Get() *WorkflowServiceItemActionProperties {
	return v.value
}

func (v *NullableWorkflowServiceItemActionProperties) Set(val *WorkflowServiceItemActionProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowServiceItemActionProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowServiceItemActionProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowServiceItemActionProperties(val *WorkflowServiceItemActionProperties) *NullableWorkflowServiceItemActionProperties {
	return &NullableWorkflowServiceItemActionProperties{value: val, isSet: true}
}

func (v NullableWorkflowServiceItemActionProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowServiceItemActionProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
