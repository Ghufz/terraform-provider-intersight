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

// CapabilityActionsMetaData Metadata or constraints of various server actions supported in Intersight for 3rd Party servers. It is validated against the target provided  and the actions are allowed only upon successful validation.
type CapabilityActionsMetaData struct {
	CapabilityCapability
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Maximum firmware version of the server.
	MaxFwVersion *string `json:"MaxFwVersion,omitempty"`
	// Minimum firmware version of the server.
	MinFwVersion *string `json:"MinFwVersion,omitempty"`
	// Model of the server that supports power or storage operations.
	Model            *string  `json:"Model,omitempty"`
	SupportedActions []string `json:"SupportedActions,omitempty"`
	// The target type to which the metadata applies. * `` - An unrecognized platform type. * `APIC` - A Cisco Application Policy Infrastructure Controller (APIC) cluster. * `CAPIC` - A Cisco Cloud Application Policy Infrastructure Controller (Cloud APIC) instance. * `DCNM` - A Cisco Data Center Network Manager (DCNM) instance. * `UCSFI` - A Cisco UCS Fabric Interconnect that is managed by Cisco UCS Manager (UCSM). * `UCSFIISM` - A Cisco UCS Fabric Interconnect that is managed by Cisco Intersight. * `IMC` - A standalone Cisco UCS rack server (Deprecated). * `IMCM4` - A standalone Cisco UCS C-Series or S-Series M4 server. * `IMCM5` - A standalone Cisco UCS C-Series or S-Series M5 server. * `IMCRack` - A standalone Cisco UCS C-Series or S-Series M6 or newer server. * `UCSIOM` - A Cisco UCS Blade Chassis I/O Module (IOM). * `HX` - A Cisco HyperFlex (HX) cluster. * `HyperFlexAP` - A Cisco HyperFlex Application Platform instance (Deprecated). * `IWE` - A Cisco Intersight Workload Engine instance (Deprecated). * `UCSD` - A Cisco UCS Director (UCSD) instance. * `IntersightAppliance` - A Cisco Intersight Connected Virtual Appliance instance. * `IntersightAssist` - A Cisco Intersight Assist instance. * `PureStorageFlashArray` - A Pure Storage FlashArray that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer and storage management features are supported on this device. * `NexusDevice` - A Cisco Nexus Network Switch that is managed using Cisco Intersight Assist. * `ACISwitch` - A Cisco Nexus Network Switch with the embedded Device Connector and is a part of the Cisco ACI fabric. * `NexusSwitch` - A standalone Cisco Nexus Network Switch with the embedded Device Connector. * `MDSSwitch` - A Cisco MDS Switch that is managed using the embedded Device Connector. * `MDSDevice` - A Cisco MDS Switch that is managed using Cisco Intersight Assist. * `UCSC890` - A standalone Cisco UCS C890 server managed using Cisco Intersight Assist. * `RedfishServer` - A generic target type for servers that support Redfish APIs and is managed using Cisco Intersight Assist. Support is limited to HPE and Dell Servers. * `NetAppOntap` - A Netapp ONTAP Storage system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device. * `NetAppActiveIqUnifiedManager` - A NetApp Active IQ Unified Manager (AIQUM) that is managed using Cisco Intersight Assist. * `EmcScaleIo` - An EMC ScaleIO Software Defined Storage system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device. * `EmcVmax` - An EMC VMAX 2 or 3 series enterprise storage array that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device. * `EmcVplex` - An EMC VPLEX virtual storage system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device. * `EmcXtremIo` - An EMC XtremIO SSD storage system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device. * `VmwareVcenter` - A VMware vCenter instance that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer and Virtualization features are supported on this hypervisor. * `MicrosoftHyperV` - A Microsoft Hyper-V host that is managed using Cisco Intersight Assist. Optionally, other hosts in the cluster can be discovered through this host. Cisco Intersight Workload Optimizer features are supported on this hypervisor. * `AppDynamics` - An AppDynamics controller running in a SaaS or on-prem datacenter. On-prem AppDynamics instance is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this controller. * `Dynatrace` - A Dynatrace Server instance running in a SaaS or on-prem datacenter. On-prem Dynatrace instance is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this server. * `NewRelic` - A NewRelic user account. The NewRelic instance monitors the application infrastructure. Cisco Intersight Workload Optimizer features are supported on this server. * `ServiceNow` - A cloud-based workflow automation platform that enables enterprise organizations to improve operational efficiencies by streamlining and automating routine work tasks. * `ReadHatOpenStack` - An OpenStack target manages Virtual Machines, Physical Machines, Datacenters and Virtual Datacenters using different OpenStack services as administrative endpoints. * `CloudFoundry` - An open source cloud platform on which developers can build, deploy, run and scale applications. * `MicrosoftAzureApplicationInsights` - A feature of Azure Monitor, is an extensible Application Performance Management service for developers and DevOps professionals to monitor their live applications. * `OpenStack` - An OpenStack target manages Virtual Machines, Physical Machines, Datacenters and Virtual Datacenters using different OpenStack services as administrative endpoints. * `MicrosoftSqlServer` - A Microsoft SQL database server that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this database. * `MySqlServer` - A MySQL database server that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this database. * `OracleDatabaseServer` - An Oracle database server that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this database. * `IBMWebSphereApplicationServer` - An IBM WebSphere Application server that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this application server. * `OracleWebLogicServer` - Oracle WebLogic Server is a unified and extensible platform for developing, deploying and running enterprise applications, such as Java, for on-premises and in the cloud. WebLogic Server offers a robust, mature, and scalable implementation of Java Enterprise Edition (EE) and Jakarta EE. * `ApacheTomcatServer` - An Apache Tomcat server that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this server. * `JavaVirtualMachine` - A JVM Application with JMX configured that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this application. * `RedHatJBossApplicationServer` - JBoss Application Server is an open-source, cross-platform Java application server developed by JBoss, a division of Red Hat Inc. It is an open-source implementation of Java 2 Enterprise Edition (J2EE) that is used for implementing Java applications and other Web-based applications and software. * `Kubernetes` - A Kubernetes cluster that runs containerized applications, with Kubernetes Collector installed. Cisco Intersight Workload Optimizer features are supported on Kubernetes cluster. * `AmazonWebService` - An Amazon Web Service cloud account.  Cisco Intersight Workload Optimizer and Virtualization features are supported on this cloud. * `AmazonWebServiceBilling` - An Amazon Web Service cloud billing account used to retrieve billing information stored in S3 bucket.  Cisco Intersight Workload Optimizer features are supported on this cloud. * `GoogleCloudPlatform` - A Google Cloud Platform service account with access to one or more projects.  Cisco Intersight Workload Optimizer features are supported on this cloud. * `GoogleCloudPlatformBilling` - A Google Cloud Platform service account used to retrieve billing information from BigQuery.  Cisco Intersight Workload Optimizer features are supported on this cloud. * `MicrosoftAzureServicePrincipal` - A Microsoft Azure Service Principal account with access to Azure subscriptions.  Cisco Intersight Workload Optimizer features are supported on this cloud. * `MicrosoftAzureEnterpriseAgreement` - A Microsoft Azure Enterprise Agreement enrolment used to retrieve pricing and billing information. Cisco Intersight Workload Optimizer features are supported on this cloud. * `MicrosoftAzureBilling` - A Microsoft Azure Service Principal account with access to billing information. Cisco Intersight Workload Optimizer features are supported on this cloud. * `DellCompellent` - A Dell EMC SC Series (Compellent) storage system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device. * `HPE3Par` - A HPE 3PAR StoreServ system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device. * `RedHatEnterpriseVirtualization` - A Red Hat Enterprise Virtualization Hypervisor system that manages Virtual Machines. * `NutanixAcropolis` - A Nutanix Acropolis cluster that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this cluster. * `HPEOneView` - A HPE OneView system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this system. * `ServiceEngine` - Cisco Application Services Engine. Cisco Application Services Engine is a platform to deploy and manage applications. * `HitachiVirtualStoragePlatform` - A Hitachi Virtual Storage Platform (Hitachi VSP) that is managed using Cisco Intersight Assist. * `GenericTarget` - A generic third-party target supported only in Partner Integration Appliance. This target type is used for development purposes and will not be supported in production environment. * `IMCBlade` - A Cisco UCS blade server managed by Cisco Intersight. * `TerraformCloud` - A Terraform Cloud Business Tier account. * `TerraformAgent` - A Terraform Cloud Agent that will be deployed on Cisco Intersight Assist. The agent can be used to plan and apply Terraform runs from a Terraform Cloud workspace. * `CustomTarget` - CustomTarget is deprecated.  Use HTTPEndpoint type to claim HTTP endpoints. * `AnsibleEndpoint` - An external endpoint that is added as a target  which can be accessed through Ansible in Intersight Cloud Orchestrator automation workflows. * `HTTPEndpoint` - An HTTP endpoint that can be accessed in Intersight Orchestrator workflows  directly or using Cisco Intersight Assist.  Authentication Schemes supported are Basic and Bearer Token. * `SSHEndpoint` - An SSH endpoint that can be accessed in Intersight Orchestrator workflows using Cisco Intersight Assist. * `CiscoCatalyst` - A Cisco Catalyst networking switch device. * `PowerShellEndpoint` - A Windows operating system server on which PowerShell scripts can be executed using Cisco Intersight Assist.
	TargetType *string `json:"TargetType,omitempty"`
	// Manufacturer of the server.
	Vendor               *string `json:"Vendor,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CapabilityActionsMetaData CapabilityActionsMetaData

// NewCapabilityActionsMetaData instantiates a new CapabilityActionsMetaData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilityActionsMetaData(classId string, objectType string) *CapabilityActionsMetaData {
	this := CapabilityActionsMetaData{}
	this.ClassId = classId
	this.ObjectType = objectType
	var targetType string = ""
	this.TargetType = &targetType
	return &this
}

// NewCapabilityActionsMetaDataWithDefaults instantiates a new CapabilityActionsMetaData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilityActionsMetaDataWithDefaults() *CapabilityActionsMetaData {
	this := CapabilityActionsMetaData{}
	var classId string = "capability.ActionsMetaData"
	this.ClassId = classId
	var objectType string = "capability.ActionsMetaData"
	this.ObjectType = objectType
	var targetType string = ""
	this.TargetType = &targetType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CapabilityActionsMetaData) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CapabilityActionsMetaData) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CapabilityActionsMetaData) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CapabilityActionsMetaData) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CapabilityActionsMetaData) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CapabilityActionsMetaData) SetObjectType(v string) {
	o.ObjectType = v
}

// GetMaxFwVersion returns the MaxFwVersion field value if set, zero value otherwise.
func (o *CapabilityActionsMetaData) GetMaxFwVersion() string {
	if o == nil || o.MaxFwVersion == nil {
		var ret string
		return ret
	}
	return *o.MaxFwVersion
}

// GetMaxFwVersionOk returns a tuple with the MaxFwVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityActionsMetaData) GetMaxFwVersionOk() (*string, bool) {
	if o == nil || o.MaxFwVersion == nil {
		return nil, false
	}
	return o.MaxFwVersion, true
}

// HasMaxFwVersion returns a boolean if a field has been set.
func (o *CapabilityActionsMetaData) HasMaxFwVersion() bool {
	if o != nil && o.MaxFwVersion != nil {
		return true
	}

	return false
}

// SetMaxFwVersion gets a reference to the given string and assigns it to the MaxFwVersion field.
func (o *CapabilityActionsMetaData) SetMaxFwVersion(v string) {
	o.MaxFwVersion = &v
}

// GetMinFwVersion returns the MinFwVersion field value if set, zero value otherwise.
func (o *CapabilityActionsMetaData) GetMinFwVersion() string {
	if o == nil || o.MinFwVersion == nil {
		var ret string
		return ret
	}
	return *o.MinFwVersion
}

// GetMinFwVersionOk returns a tuple with the MinFwVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityActionsMetaData) GetMinFwVersionOk() (*string, bool) {
	if o == nil || o.MinFwVersion == nil {
		return nil, false
	}
	return o.MinFwVersion, true
}

// HasMinFwVersion returns a boolean if a field has been set.
func (o *CapabilityActionsMetaData) HasMinFwVersion() bool {
	if o != nil && o.MinFwVersion != nil {
		return true
	}

	return false
}

// SetMinFwVersion gets a reference to the given string and assigns it to the MinFwVersion field.
func (o *CapabilityActionsMetaData) SetMinFwVersion(v string) {
	o.MinFwVersion = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *CapabilityActionsMetaData) GetModel() string {
	if o == nil || o.Model == nil {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityActionsMetaData) GetModelOk() (*string, bool) {
	if o == nil || o.Model == nil {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *CapabilityActionsMetaData) HasModel() bool {
	if o != nil && o.Model != nil {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *CapabilityActionsMetaData) SetModel(v string) {
	o.Model = &v
}

// GetSupportedActions returns the SupportedActions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CapabilityActionsMetaData) GetSupportedActions() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SupportedActions
}

// GetSupportedActionsOk returns a tuple with the SupportedActions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CapabilityActionsMetaData) GetSupportedActionsOk() ([]string, bool) {
	if o == nil || o.SupportedActions == nil {
		return nil, false
	}
	return o.SupportedActions, true
}

// HasSupportedActions returns a boolean if a field has been set.
func (o *CapabilityActionsMetaData) HasSupportedActions() bool {
	if o != nil && o.SupportedActions != nil {
		return true
	}

	return false
}

// SetSupportedActions gets a reference to the given []string and assigns it to the SupportedActions field.
func (o *CapabilityActionsMetaData) SetSupportedActions(v []string) {
	o.SupportedActions = v
}

// GetTargetType returns the TargetType field value if set, zero value otherwise.
func (o *CapabilityActionsMetaData) GetTargetType() string {
	if o == nil || o.TargetType == nil {
		var ret string
		return ret
	}
	return *o.TargetType
}

// GetTargetTypeOk returns a tuple with the TargetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityActionsMetaData) GetTargetTypeOk() (*string, bool) {
	if o == nil || o.TargetType == nil {
		return nil, false
	}
	return o.TargetType, true
}

// HasTargetType returns a boolean if a field has been set.
func (o *CapabilityActionsMetaData) HasTargetType() bool {
	if o != nil && o.TargetType != nil {
		return true
	}

	return false
}

// SetTargetType gets a reference to the given string and assigns it to the TargetType field.
func (o *CapabilityActionsMetaData) SetTargetType(v string) {
	o.TargetType = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *CapabilityActionsMetaData) GetVendor() string {
	if o == nil || o.Vendor == nil {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityActionsMetaData) GetVendorOk() (*string, bool) {
	if o == nil || o.Vendor == nil {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *CapabilityActionsMetaData) HasVendor() bool {
	if o != nil && o.Vendor != nil {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *CapabilityActionsMetaData) SetVendor(v string) {
	o.Vendor = &v
}

func (o CapabilityActionsMetaData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedCapabilityCapability, errCapabilityCapability := json.Marshal(o.CapabilityCapability)
	if errCapabilityCapability != nil {
		return []byte{}, errCapabilityCapability
	}
	errCapabilityCapability = json.Unmarshal([]byte(serializedCapabilityCapability), &toSerialize)
	if errCapabilityCapability != nil {
		return []byte{}, errCapabilityCapability
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.MaxFwVersion != nil {
		toSerialize["MaxFwVersion"] = o.MaxFwVersion
	}
	if o.MinFwVersion != nil {
		toSerialize["MinFwVersion"] = o.MinFwVersion
	}
	if o.Model != nil {
		toSerialize["Model"] = o.Model
	}
	if o.SupportedActions != nil {
		toSerialize["SupportedActions"] = o.SupportedActions
	}
	if o.TargetType != nil {
		toSerialize["TargetType"] = o.TargetType
	}
	if o.Vendor != nil {
		toSerialize["Vendor"] = o.Vendor
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CapabilityActionsMetaData) UnmarshalJSON(bytes []byte) (err error) {
	type CapabilityActionsMetaDataWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Maximum firmware version of the server.
		MaxFwVersion *string `json:"MaxFwVersion,omitempty"`
		// Minimum firmware version of the server.
		MinFwVersion *string `json:"MinFwVersion,omitempty"`
		// Model of the server that supports power or storage operations.
		Model            *string  `json:"Model,omitempty"`
		SupportedActions []string `json:"SupportedActions,omitempty"`
		// The target type to which the metadata applies. * `` - An unrecognized platform type. * `APIC` - A Cisco Application Policy Infrastructure Controller (APIC) cluster. * `CAPIC` - A Cisco Cloud Application Policy Infrastructure Controller (Cloud APIC) instance. * `DCNM` - A Cisco Data Center Network Manager (DCNM) instance. * `UCSFI` - A Cisco UCS Fabric Interconnect that is managed by Cisco UCS Manager (UCSM). * `UCSFIISM` - A Cisco UCS Fabric Interconnect that is managed by Cisco Intersight. * `IMC` - A standalone Cisco UCS rack server (Deprecated). * `IMCM4` - A standalone Cisco UCS C-Series or S-Series M4 server. * `IMCM5` - A standalone Cisco UCS C-Series or S-Series M5 server. * `IMCRack` - A standalone Cisco UCS C-Series or S-Series M6 or newer server. * `UCSIOM` - A Cisco UCS Blade Chassis I/O Module (IOM). * `HX` - A Cisco HyperFlex (HX) cluster. * `HyperFlexAP` - A Cisco HyperFlex Application Platform instance (Deprecated). * `IWE` - A Cisco Intersight Workload Engine instance (Deprecated). * `UCSD` - A Cisco UCS Director (UCSD) instance. * `IntersightAppliance` - A Cisco Intersight Connected Virtual Appliance instance. * `IntersightAssist` - A Cisco Intersight Assist instance. * `PureStorageFlashArray` - A Pure Storage FlashArray that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer and storage management features are supported on this device. * `NexusDevice` - A Cisco Nexus Network Switch that is managed using Cisco Intersight Assist. * `ACISwitch` - A Cisco Nexus Network Switch with the embedded Device Connector and is a part of the Cisco ACI fabric. * `NexusSwitch` - A standalone Cisco Nexus Network Switch with the embedded Device Connector. * `MDSSwitch` - A Cisco MDS Switch that is managed using the embedded Device Connector. * `MDSDevice` - A Cisco MDS Switch that is managed using Cisco Intersight Assist. * `UCSC890` - A standalone Cisco UCS C890 server managed using Cisco Intersight Assist. * `RedfishServer` - A generic target type for servers that support Redfish APIs and is managed using Cisco Intersight Assist. Support is limited to HPE and Dell Servers. * `NetAppOntap` - A Netapp ONTAP Storage system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device. * `NetAppActiveIqUnifiedManager` - A NetApp Active IQ Unified Manager (AIQUM) that is managed using Cisco Intersight Assist. * `EmcScaleIo` - An EMC ScaleIO Software Defined Storage system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device. * `EmcVmax` - An EMC VMAX 2 or 3 series enterprise storage array that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device. * `EmcVplex` - An EMC VPLEX virtual storage system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device. * `EmcXtremIo` - An EMC XtremIO SSD storage system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device. * `VmwareVcenter` - A VMware vCenter instance that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer and Virtualization features are supported on this hypervisor. * `MicrosoftHyperV` - A Microsoft Hyper-V host that is managed using Cisco Intersight Assist. Optionally, other hosts in the cluster can be discovered through this host. Cisco Intersight Workload Optimizer features are supported on this hypervisor. * `AppDynamics` - An AppDynamics controller running in a SaaS or on-prem datacenter. On-prem AppDynamics instance is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this controller. * `Dynatrace` - A Dynatrace Server instance running in a SaaS or on-prem datacenter. On-prem Dynatrace instance is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this server. * `NewRelic` - A NewRelic user account. The NewRelic instance monitors the application infrastructure. Cisco Intersight Workload Optimizer features are supported on this server. * `ServiceNow` - A cloud-based workflow automation platform that enables enterprise organizations to improve operational efficiencies by streamlining and automating routine work tasks. * `ReadHatOpenStack` - An OpenStack target manages Virtual Machines, Physical Machines, Datacenters and Virtual Datacenters using different OpenStack services as administrative endpoints. * `CloudFoundry` - An open source cloud platform on which developers can build, deploy, run and scale applications. * `MicrosoftAzureApplicationInsights` - A feature of Azure Monitor, is an extensible Application Performance Management service for developers and DevOps professionals to monitor their live applications. * `OpenStack` - An OpenStack target manages Virtual Machines, Physical Machines, Datacenters and Virtual Datacenters using different OpenStack services as administrative endpoints. * `MicrosoftSqlServer` - A Microsoft SQL database server that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this database. * `MySqlServer` - A MySQL database server that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this database. * `OracleDatabaseServer` - An Oracle database server that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this database. * `IBMWebSphereApplicationServer` - An IBM WebSphere Application server that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this application server. * `OracleWebLogicServer` - Oracle WebLogic Server is a unified and extensible platform for developing, deploying and running enterprise applications, such as Java, for on-premises and in the cloud. WebLogic Server offers a robust, mature, and scalable implementation of Java Enterprise Edition (EE) and Jakarta EE. * `ApacheTomcatServer` - An Apache Tomcat server that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this server. * `JavaVirtualMachine` - A JVM Application with JMX configured that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this application. * `RedHatJBossApplicationServer` - JBoss Application Server is an open-source, cross-platform Java application server developed by JBoss, a division of Red Hat Inc. It is an open-source implementation of Java 2 Enterprise Edition (J2EE) that is used for implementing Java applications and other Web-based applications and software. * `Kubernetes` - A Kubernetes cluster that runs containerized applications, with Kubernetes Collector installed. Cisco Intersight Workload Optimizer features are supported on Kubernetes cluster. * `AmazonWebService` - An Amazon Web Service cloud account.  Cisco Intersight Workload Optimizer and Virtualization features are supported on this cloud. * `AmazonWebServiceBilling` - An Amazon Web Service cloud billing account used to retrieve billing information stored in S3 bucket.  Cisco Intersight Workload Optimizer features are supported on this cloud. * `GoogleCloudPlatform` - A Google Cloud Platform service account with access to one or more projects.  Cisco Intersight Workload Optimizer features are supported on this cloud. * `GoogleCloudPlatformBilling` - A Google Cloud Platform service account used to retrieve billing information from BigQuery.  Cisco Intersight Workload Optimizer features are supported on this cloud. * `MicrosoftAzureServicePrincipal` - A Microsoft Azure Service Principal account with access to Azure subscriptions.  Cisco Intersight Workload Optimizer features are supported on this cloud. * `MicrosoftAzureEnterpriseAgreement` - A Microsoft Azure Enterprise Agreement enrolment used to retrieve pricing and billing information. Cisco Intersight Workload Optimizer features are supported on this cloud. * `MicrosoftAzureBilling` - A Microsoft Azure Service Principal account with access to billing information. Cisco Intersight Workload Optimizer features are supported on this cloud. * `DellCompellent` - A Dell EMC SC Series (Compellent) storage system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device. * `HPE3Par` - A HPE 3PAR StoreServ system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this device. * `RedHatEnterpriseVirtualization` - A Red Hat Enterprise Virtualization Hypervisor system that manages Virtual Machines. * `NutanixAcropolis` - A Nutanix Acropolis cluster that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this cluster. * `HPEOneView` - A HPE OneView system that is managed using Cisco Intersight Assist. Cisco Intersight Workload Optimizer features are supported on this system. * `ServiceEngine` - Cisco Application Services Engine. Cisco Application Services Engine is a platform to deploy and manage applications. * `HitachiVirtualStoragePlatform` - A Hitachi Virtual Storage Platform (Hitachi VSP) that is managed using Cisco Intersight Assist. * `GenericTarget` - A generic third-party target supported only in Partner Integration Appliance. This target type is used for development purposes and will not be supported in production environment. * `IMCBlade` - A Cisco UCS blade server managed by Cisco Intersight. * `TerraformCloud` - A Terraform Cloud Business Tier account. * `TerraformAgent` - A Terraform Cloud Agent that will be deployed on Cisco Intersight Assist. The agent can be used to plan and apply Terraform runs from a Terraform Cloud workspace. * `CustomTarget` - CustomTarget is deprecated.  Use HTTPEndpoint type to claim HTTP endpoints. * `AnsibleEndpoint` - An external endpoint that is added as a target  which can be accessed through Ansible in Intersight Cloud Orchestrator automation workflows. * `HTTPEndpoint` - An HTTP endpoint that can be accessed in Intersight Orchestrator workflows  directly or using Cisco Intersight Assist.  Authentication Schemes supported are Basic and Bearer Token. * `SSHEndpoint` - An SSH endpoint that can be accessed in Intersight Orchestrator workflows using Cisco Intersight Assist. * `CiscoCatalyst` - A Cisco Catalyst networking switch device. * `PowerShellEndpoint` - A Windows operating system server on which PowerShell scripts can be executed using Cisco Intersight Assist.
		TargetType *string `json:"TargetType,omitempty"`
		// Manufacturer of the server.
		Vendor *string `json:"Vendor,omitempty"`
	}

	varCapabilityActionsMetaDataWithoutEmbeddedStruct := CapabilityActionsMetaDataWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varCapabilityActionsMetaDataWithoutEmbeddedStruct)
	if err == nil {
		varCapabilityActionsMetaData := _CapabilityActionsMetaData{}
		varCapabilityActionsMetaData.ClassId = varCapabilityActionsMetaDataWithoutEmbeddedStruct.ClassId
		varCapabilityActionsMetaData.ObjectType = varCapabilityActionsMetaDataWithoutEmbeddedStruct.ObjectType
		varCapabilityActionsMetaData.MaxFwVersion = varCapabilityActionsMetaDataWithoutEmbeddedStruct.MaxFwVersion
		varCapabilityActionsMetaData.MinFwVersion = varCapabilityActionsMetaDataWithoutEmbeddedStruct.MinFwVersion
		varCapabilityActionsMetaData.Model = varCapabilityActionsMetaDataWithoutEmbeddedStruct.Model
		varCapabilityActionsMetaData.SupportedActions = varCapabilityActionsMetaDataWithoutEmbeddedStruct.SupportedActions
		varCapabilityActionsMetaData.TargetType = varCapabilityActionsMetaDataWithoutEmbeddedStruct.TargetType
		varCapabilityActionsMetaData.Vendor = varCapabilityActionsMetaDataWithoutEmbeddedStruct.Vendor
		*o = CapabilityActionsMetaData(varCapabilityActionsMetaData)
	} else {
		return err
	}

	varCapabilityActionsMetaData := _CapabilityActionsMetaData{}

	err = json.Unmarshal(bytes, &varCapabilityActionsMetaData)
	if err == nil {
		o.CapabilityCapability = varCapabilityActionsMetaData.CapabilityCapability
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "MaxFwVersion")
		delete(additionalProperties, "MinFwVersion")
		delete(additionalProperties, "Model")
		delete(additionalProperties, "SupportedActions")
		delete(additionalProperties, "TargetType")
		delete(additionalProperties, "Vendor")

		// remove fields from embedded structs
		reflectCapabilityCapability := reflect.ValueOf(o.CapabilityCapability)
		for i := 0; i < reflectCapabilityCapability.Type().NumField(); i++ {
			t := reflectCapabilityCapability.Type().Field(i)

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

type NullableCapabilityActionsMetaData struct {
	value *CapabilityActionsMetaData
	isSet bool
}

func (v NullableCapabilityActionsMetaData) Get() *CapabilityActionsMetaData {
	return v.value
}

func (v *NullableCapabilityActionsMetaData) Set(val *CapabilityActionsMetaData) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilityActionsMetaData) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilityActionsMetaData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilityActionsMetaData(val *CapabilityActionsMetaData) *NullableCapabilityActionsMetaData {
	return &NullableCapabilityActionsMetaData{value: val, isSet: true}
}

func (v NullableCapabilityActionsMetaData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilityActionsMetaData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
