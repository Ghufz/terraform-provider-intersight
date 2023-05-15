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
)

// OpenapiTaskGenerationRequestAllOf Definition of the list of properties defined in 'openapi.TaskGenerationRequest', excluding properties defined in parent classes.
type OpenapiTaskGenerationRequestAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Indicates if target endpoint is external or internal. An endpoint is internal if the target is an Intersight resource. For instance, configuring an intersight object using a Task. * `External` - Denotes that the target endpoint is an external API endpoint * `Internal` - Denotes that the target endpoint is a Intersight API endpoint
	EndpointType     *string               `json:"EndpointType,omitempty"`
	FailedTasks      []OpenapiFailedTask   `json:"FailedTasks,omitempty"`
	HeaderParameters []OpenapiKeyValuePair `json:"HeaderParameters,omitempty"`
	// When this value is set to true, the task name validations happen and provides the task validation status against each of the selected API on the property selectedApis When this value is set to false, an internal workflow to generate the tasks is submitted,  conflicting task names are created with an incremented version.
	IsValidateRequest *bool `json:"IsValidateRequest,omitempty"`
	// Specifies the REST API protocol being used, which can be one of HTTP or HTTPS. * `HTTPS` - Denotes that the API call uses the HTTPS protocol type * `HTTP` - Denotes that the API call uses the HTTP protocol type
	Protocol        *string               `json:"Protocol,omitempty"`
	QueryParameters []OpenapiKeyValuePair `json:"QueryParameters,omitempty"`
	SelectedApis    []OpenapiApiInfo      `json:"SelectedApis,omitempty"`
	// Depicts the status of the task creation request. * `none` - Indicates the default status. * `InProgress` - Request has been picked up for generating tasks from the OpenAPI Specification file. * `Completed` - All the tasks from the request have been created. * `Failed` - There were failures in generating one or more tasks in the request.
	Status *string `json:"Status,omitempty"`
	// Optional string that can be prefixed to the name of created tasks.
	TaskPrefix *string               `json:"TaskPrefix,omitempty"`
	TaskTags   []OpenapiKeyValuePair `json:"TaskTags,omitempty"`
	// Specifies the URL of the endpoint that the created task communicates with. It is defaulted to intersight.com for internal endpoints.
	Url                  *string                               `json:"Url,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	Source               *OpenapiProcessFileRelationship       `json:"Source,omitempty"`
	Workflow             *WorkflowWorkflowInfoRelationship     `json:"Workflow,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OpenapiTaskGenerationRequestAllOf OpenapiTaskGenerationRequestAllOf

// NewOpenapiTaskGenerationRequestAllOf instantiates a new OpenapiTaskGenerationRequestAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenapiTaskGenerationRequestAllOf(classId string, objectType string) *OpenapiTaskGenerationRequestAllOf {
	this := OpenapiTaskGenerationRequestAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var endpointType string = "External"
	this.EndpointType = &endpointType
	var protocol string = "HTTPS"
	this.Protocol = &protocol
	return &this
}

// NewOpenapiTaskGenerationRequestAllOfWithDefaults instantiates a new OpenapiTaskGenerationRequestAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenapiTaskGenerationRequestAllOfWithDefaults() *OpenapiTaskGenerationRequestAllOf {
	this := OpenapiTaskGenerationRequestAllOf{}
	var classId string = "openapi.TaskGenerationRequest"
	this.ClassId = classId
	var objectType string = "openapi.TaskGenerationRequest"
	this.ObjectType = objectType
	var endpointType string = "External"
	this.EndpointType = &endpointType
	var protocol string = "HTTPS"
	this.Protocol = &protocol
	return &this
}

// GetClassId returns the ClassId field value
func (o *OpenapiTaskGenerationRequestAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *OpenapiTaskGenerationRequestAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *OpenapiTaskGenerationRequestAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *OpenapiTaskGenerationRequestAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *OpenapiTaskGenerationRequestAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *OpenapiTaskGenerationRequestAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEndpointType returns the EndpointType field value if set, zero value otherwise.
func (o *OpenapiTaskGenerationRequestAllOf) GetEndpointType() string {
	if o == nil || o.EndpointType == nil {
		var ret string
		return ret
	}
	return *o.EndpointType
}

// GetEndpointTypeOk returns a tuple with the EndpointType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenapiTaskGenerationRequestAllOf) GetEndpointTypeOk() (*string, bool) {
	if o == nil || o.EndpointType == nil {
		return nil, false
	}
	return o.EndpointType, true
}

// HasEndpointType returns a boolean if a field has been set.
func (o *OpenapiTaskGenerationRequestAllOf) HasEndpointType() bool {
	if o != nil && o.EndpointType != nil {
		return true
	}

	return false
}

// SetEndpointType gets a reference to the given string and assigns it to the EndpointType field.
func (o *OpenapiTaskGenerationRequestAllOf) SetEndpointType(v string) {
	o.EndpointType = &v
}

// GetFailedTasks returns the FailedTasks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenapiTaskGenerationRequestAllOf) GetFailedTasks() []OpenapiFailedTask {
	if o == nil {
		var ret []OpenapiFailedTask
		return ret
	}
	return o.FailedTasks
}

// GetFailedTasksOk returns a tuple with the FailedTasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenapiTaskGenerationRequestAllOf) GetFailedTasksOk() ([]OpenapiFailedTask, bool) {
	if o == nil || o.FailedTasks == nil {
		return nil, false
	}
	return o.FailedTasks, true
}

// HasFailedTasks returns a boolean if a field has been set.
func (o *OpenapiTaskGenerationRequestAllOf) HasFailedTasks() bool {
	if o != nil && o.FailedTasks != nil {
		return true
	}

	return false
}

// SetFailedTasks gets a reference to the given []OpenapiFailedTask and assigns it to the FailedTasks field.
func (o *OpenapiTaskGenerationRequestAllOf) SetFailedTasks(v []OpenapiFailedTask) {
	o.FailedTasks = v
}

// GetHeaderParameters returns the HeaderParameters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenapiTaskGenerationRequestAllOf) GetHeaderParameters() []OpenapiKeyValuePair {
	if o == nil {
		var ret []OpenapiKeyValuePair
		return ret
	}
	return o.HeaderParameters
}

// GetHeaderParametersOk returns a tuple with the HeaderParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenapiTaskGenerationRequestAllOf) GetHeaderParametersOk() ([]OpenapiKeyValuePair, bool) {
	if o == nil || o.HeaderParameters == nil {
		return nil, false
	}
	return o.HeaderParameters, true
}

// HasHeaderParameters returns a boolean if a field has been set.
func (o *OpenapiTaskGenerationRequestAllOf) HasHeaderParameters() bool {
	if o != nil && o.HeaderParameters != nil {
		return true
	}

	return false
}

// SetHeaderParameters gets a reference to the given []OpenapiKeyValuePair and assigns it to the HeaderParameters field.
func (o *OpenapiTaskGenerationRequestAllOf) SetHeaderParameters(v []OpenapiKeyValuePair) {
	o.HeaderParameters = v
}

// GetIsValidateRequest returns the IsValidateRequest field value if set, zero value otherwise.
func (o *OpenapiTaskGenerationRequestAllOf) GetIsValidateRequest() bool {
	if o == nil || o.IsValidateRequest == nil {
		var ret bool
		return ret
	}
	return *o.IsValidateRequest
}

// GetIsValidateRequestOk returns a tuple with the IsValidateRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenapiTaskGenerationRequestAllOf) GetIsValidateRequestOk() (*bool, bool) {
	if o == nil || o.IsValidateRequest == nil {
		return nil, false
	}
	return o.IsValidateRequest, true
}

// HasIsValidateRequest returns a boolean if a field has been set.
func (o *OpenapiTaskGenerationRequestAllOf) HasIsValidateRequest() bool {
	if o != nil && o.IsValidateRequest != nil {
		return true
	}

	return false
}

// SetIsValidateRequest gets a reference to the given bool and assigns it to the IsValidateRequest field.
func (o *OpenapiTaskGenerationRequestAllOf) SetIsValidateRequest(v bool) {
	o.IsValidateRequest = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *OpenapiTaskGenerationRequestAllOf) GetProtocol() string {
	if o == nil || o.Protocol == nil {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenapiTaskGenerationRequestAllOf) GetProtocolOk() (*string, bool) {
	if o == nil || o.Protocol == nil {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *OpenapiTaskGenerationRequestAllOf) HasProtocol() bool {
	if o != nil && o.Protocol != nil {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *OpenapiTaskGenerationRequestAllOf) SetProtocol(v string) {
	o.Protocol = &v
}

// GetQueryParameters returns the QueryParameters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenapiTaskGenerationRequestAllOf) GetQueryParameters() []OpenapiKeyValuePair {
	if o == nil {
		var ret []OpenapiKeyValuePair
		return ret
	}
	return o.QueryParameters
}

// GetQueryParametersOk returns a tuple with the QueryParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenapiTaskGenerationRequestAllOf) GetQueryParametersOk() ([]OpenapiKeyValuePair, bool) {
	if o == nil || o.QueryParameters == nil {
		return nil, false
	}
	return o.QueryParameters, true
}

// HasQueryParameters returns a boolean if a field has been set.
func (o *OpenapiTaskGenerationRequestAllOf) HasQueryParameters() bool {
	if o != nil && o.QueryParameters != nil {
		return true
	}

	return false
}

// SetQueryParameters gets a reference to the given []OpenapiKeyValuePair and assigns it to the QueryParameters field.
func (o *OpenapiTaskGenerationRequestAllOf) SetQueryParameters(v []OpenapiKeyValuePair) {
	o.QueryParameters = v
}

// GetSelectedApis returns the SelectedApis field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenapiTaskGenerationRequestAllOf) GetSelectedApis() []OpenapiApiInfo {
	if o == nil {
		var ret []OpenapiApiInfo
		return ret
	}
	return o.SelectedApis
}

// GetSelectedApisOk returns a tuple with the SelectedApis field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenapiTaskGenerationRequestAllOf) GetSelectedApisOk() ([]OpenapiApiInfo, bool) {
	if o == nil || o.SelectedApis == nil {
		return nil, false
	}
	return o.SelectedApis, true
}

// HasSelectedApis returns a boolean if a field has been set.
func (o *OpenapiTaskGenerationRequestAllOf) HasSelectedApis() bool {
	if o != nil && o.SelectedApis != nil {
		return true
	}

	return false
}

// SetSelectedApis gets a reference to the given []OpenapiApiInfo and assigns it to the SelectedApis field.
func (o *OpenapiTaskGenerationRequestAllOf) SetSelectedApis(v []OpenapiApiInfo) {
	o.SelectedApis = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OpenapiTaskGenerationRequestAllOf) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenapiTaskGenerationRequestAllOf) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OpenapiTaskGenerationRequestAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *OpenapiTaskGenerationRequestAllOf) SetStatus(v string) {
	o.Status = &v
}

// GetTaskPrefix returns the TaskPrefix field value if set, zero value otherwise.
func (o *OpenapiTaskGenerationRequestAllOf) GetTaskPrefix() string {
	if o == nil || o.TaskPrefix == nil {
		var ret string
		return ret
	}
	return *o.TaskPrefix
}

// GetTaskPrefixOk returns a tuple with the TaskPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenapiTaskGenerationRequestAllOf) GetTaskPrefixOk() (*string, bool) {
	if o == nil || o.TaskPrefix == nil {
		return nil, false
	}
	return o.TaskPrefix, true
}

// HasTaskPrefix returns a boolean if a field has been set.
func (o *OpenapiTaskGenerationRequestAllOf) HasTaskPrefix() bool {
	if o != nil && o.TaskPrefix != nil {
		return true
	}

	return false
}

// SetTaskPrefix gets a reference to the given string and assigns it to the TaskPrefix field.
func (o *OpenapiTaskGenerationRequestAllOf) SetTaskPrefix(v string) {
	o.TaskPrefix = &v
}

// GetTaskTags returns the TaskTags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenapiTaskGenerationRequestAllOf) GetTaskTags() []OpenapiKeyValuePair {
	if o == nil {
		var ret []OpenapiKeyValuePair
		return ret
	}
	return o.TaskTags
}

// GetTaskTagsOk returns a tuple with the TaskTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenapiTaskGenerationRequestAllOf) GetTaskTagsOk() ([]OpenapiKeyValuePair, bool) {
	if o == nil || o.TaskTags == nil {
		return nil, false
	}
	return o.TaskTags, true
}

// HasTaskTags returns a boolean if a field has been set.
func (o *OpenapiTaskGenerationRequestAllOf) HasTaskTags() bool {
	if o != nil && o.TaskTags != nil {
		return true
	}

	return false
}

// SetTaskTags gets a reference to the given []OpenapiKeyValuePair and assigns it to the TaskTags field.
func (o *OpenapiTaskGenerationRequestAllOf) SetTaskTags(v []OpenapiKeyValuePair) {
	o.TaskTags = v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *OpenapiTaskGenerationRequestAllOf) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenapiTaskGenerationRequestAllOf) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *OpenapiTaskGenerationRequestAllOf) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *OpenapiTaskGenerationRequestAllOf) SetUrl(v string) {
	o.Url = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *OpenapiTaskGenerationRequestAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenapiTaskGenerationRequestAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *OpenapiTaskGenerationRequestAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *OpenapiTaskGenerationRequestAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *OpenapiTaskGenerationRequestAllOf) GetSource() OpenapiProcessFileRelationship {
	if o == nil || o.Source == nil {
		var ret OpenapiProcessFileRelationship
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenapiTaskGenerationRequestAllOf) GetSourceOk() (*OpenapiProcessFileRelationship, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *OpenapiTaskGenerationRequestAllOf) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given OpenapiProcessFileRelationship and assigns it to the Source field.
func (o *OpenapiTaskGenerationRequestAllOf) SetSource(v OpenapiProcessFileRelationship) {
	o.Source = &v
}

// GetWorkflow returns the Workflow field value if set, zero value otherwise.
func (o *OpenapiTaskGenerationRequestAllOf) GetWorkflow() WorkflowWorkflowInfoRelationship {
	if o == nil || o.Workflow == nil {
		var ret WorkflowWorkflowInfoRelationship
		return ret
	}
	return *o.Workflow
}

// GetWorkflowOk returns a tuple with the Workflow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenapiTaskGenerationRequestAllOf) GetWorkflowOk() (*WorkflowWorkflowInfoRelationship, bool) {
	if o == nil || o.Workflow == nil {
		return nil, false
	}
	return o.Workflow, true
}

// HasWorkflow returns a boolean if a field has been set.
func (o *OpenapiTaskGenerationRequestAllOf) HasWorkflow() bool {
	if o != nil && o.Workflow != nil {
		return true
	}

	return false
}

// SetWorkflow gets a reference to the given WorkflowWorkflowInfoRelationship and assigns it to the Workflow field.
func (o *OpenapiTaskGenerationRequestAllOf) SetWorkflow(v WorkflowWorkflowInfoRelationship) {
	o.Workflow = &v
}

func (o OpenapiTaskGenerationRequestAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.EndpointType != nil {
		toSerialize["EndpointType"] = o.EndpointType
	}
	if o.FailedTasks != nil {
		toSerialize["FailedTasks"] = o.FailedTasks
	}
	if o.HeaderParameters != nil {
		toSerialize["HeaderParameters"] = o.HeaderParameters
	}
	if o.IsValidateRequest != nil {
		toSerialize["IsValidateRequest"] = o.IsValidateRequest
	}
	if o.Protocol != nil {
		toSerialize["Protocol"] = o.Protocol
	}
	if o.QueryParameters != nil {
		toSerialize["QueryParameters"] = o.QueryParameters
	}
	if o.SelectedApis != nil {
		toSerialize["SelectedApis"] = o.SelectedApis
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.TaskPrefix != nil {
		toSerialize["TaskPrefix"] = o.TaskPrefix
	}
	if o.TaskTags != nil {
		toSerialize["TaskTags"] = o.TaskTags
	}
	if o.Url != nil {
		toSerialize["Url"] = o.Url
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.Source != nil {
		toSerialize["Source"] = o.Source
	}
	if o.Workflow != nil {
		toSerialize["Workflow"] = o.Workflow
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OpenapiTaskGenerationRequestAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varOpenapiTaskGenerationRequestAllOf := _OpenapiTaskGenerationRequestAllOf{}

	if err = json.Unmarshal(bytes, &varOpenapiTaskGenerationRequestAllOf); err == nil {
		*o = OpenapiTaskGenerationRequestAllOf(varOpenapiTaskGenerationRequestAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "EndpointType")
		delete(additionalProperties, "FailedTasks")
		delete(additionalProperties, "HeaderParameters")
		delete(additionalProperties, "IsValidateRequest")
		delete(additionalProperties, "Protocol")
		delete(additionalProperties, "QueryParameters")
		delete(additionalProperties, "SelectedApis")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "TaskPrefix")
		delete(additionalProperties, "TaskTags")
		delete(additionalProperties, "Url")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "Source")
		delete(additionalProperties, "Workflow")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOpenapiTaskGenerationRequestAllOf struct {
	value *OpenapiTaskGenerationRequestAllOf
	isSet bool
}

func (v NullableOpenapiTaskGenerationRequestAllOf) Get() *OpenapiTaskGenerationRequestAllOf {
	return v.value
}

func (v *NullableOpenapiTaskGenerationRequestAllOf) Set(val *OpenapiTaskGenerationRequestAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenapiTaskGenerationRequestAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenapiTaskGenerationRequestAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenapiTaskGenerationRequestAllOf(val *OpenapiTaskGenerationRequestAllOf) *NullableOpenapiTaskGenerationRequestAllOf {
	return &NullableOpenapiTaskGenerationRequestAllOf{value: val, isSet: true}
}

func (v NullableOpenapiTaskGenerationRequestAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenapiTaskGenerationRequestAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
