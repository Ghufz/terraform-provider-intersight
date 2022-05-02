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
)

// CapabilityAdapterUnitDescriptorAllOf Definition of the list of properties defined in 'capability.AdapterUnitDescriptor', excluding properties defined in parent classes.
type CapabilityAdapterUnitDescriptorAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Generation of the adapter. * `4` - Fourth generation adapters (14xx). The PIDs of these adapters end with the string 04. * `2` - Second generation VIC adapters (12xx). The PIDs of these adapters end with the string 02. * `3` - Third generation adapters (13xx). The PIDs of these adapters end with the string 03. * `5` - Fifth generation adapters (15xx). The PIDs of these adapters contain the V5 string.
	AdapterGeneration *int32 `json:"AdapterGeneration,omitempty"`
	// Order in which the ports are connected; sequential or cyclic.
	ConnectivityOrder *string `json:"ConnectivityOrder,omitempty"`
	// The port speed for ethernet ports in Mbps.
	EthernetPortSpeed *int64                    `json:"EthernetPortSpeed,omitempty"`
	Features          []CapabilityFeatureConfig `json:"Features,omitempty"`
	// The port speed for fibre channel ports in Mbps.
	FibreChannelPortSpeed *int64 `json:"FibreChannelPortSpeed,omitempty"`
	// The number of SCSI I/O Queue resources to allocate.
	FibreChannelScsiIoqLimit *int64 `json:"FibreChannelScsiIoqLimit,omitempty"`
	// Indicates that the Azure Stack Host QoS feature is supported by this adapter.
	IsAzureQosSupported *bool `json:"IsAzureQosSupported,omitempty"`
	// Indicates that the GENEVE offload feature is supported by this adapter.
	IsGeneveSupported *bool `json:"IsGeneveSupported,omitempty"`
	// Maximum Ring Size value for vNIC Receive Queue.
	MaxEthRxRingSize *int64 `json:"MaxEthRxRingSize,omitempty"`
	// Maximum Ring Size value for vNIC Transmit Queue.
	MaxEthTxRingSize *int64 `json:"MaxEthTxRingSize,omitempty"`
	// Maximum number of vNIC interfaces that can be RoCEv2 enabled.
	MaxRocev2Interfaces *int64 `json:"MaxRocev2Interfaces,omitempty"`
	// Number of Dce Ports for the adapter.
	NumDcePorts *int64 `json:"NumDcePorts,omitempty"`
	// Indicates PCI Link status of adapter.
	PciLink *int64 `json:"PciLink,omitempty"`
	// Prom card type for the adapter.
	PromCardType         *string `json:"PromCardType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CapabilityAdapterUnitDescriptorAllOf CapabilityAdapterUnitDescriptorAllOf

// NewCapabilityAdapterUnitDescriptorAllOf instantiates a new CapabilityAdapterUnitDescriptorAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilityAdapterUnitDescriptorAllOf(classId string, objectType string) *CapabilityAdapterUnitDescriptorAllOf {
	this := CapabilityAdapterUnitDescriptorAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var adapterGeneration int32 = 4
	this.AdapterGeneration = &adapterGeneration
	var isAzureQosSupported bool = true
	this.IsAzureQosSupported = &isAzureQosSupported
	var isGeneveSupported bool = true
	this.IsGeneveSupported = &isGeneveSupported
	var maxEthRxRingSize int64 = 4096
	this.MaxEthRxRingSize = &maxEthRxRingSize
	var maxEthTxRingSize int64 = 4096
	this.MaxEthTxRingSize = &maxEthTxRingSize
	var maxRocev2Interfaces int64 = 2
	this.MaxRocev2Interfaces = &maxRocev2Interfaces
	var pciLink int64 = 0
	this.PciLink = &pciLink
	return &this
}

// NewCapabilityAdapterUnitDescriptorAllOfWithDefaults instantiates a new CapabilityAdapterUnitDescriptorAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilityAdapterUnitDescriptorAllOfWithDefaults() *CapabilityAdapterUnitDescriptorAllOf {
	this := CapabilityAdapterUnitDescriptorAllOf{}
	var classId string = "capability.AdapterUnitDescriptor"
	this.ClassId = classId
	var objectType string = "capability.AdapterUnitDescriptor"
	this.ObjectType = objectType
	var adapterGeneration int32 = 4
	this.AdapterGeneration = &adapterGeneration
	var isAzureQosSupported bool = true
	this.IsAzureQosSupported = &isAzureQosSupported
	var isGeneveSupported bool = true
	this.IsGeneveSupported = &isGeneveSupported
	var maxEthRxRingSize int64 = 4096
	this.MaxEthRxRingSize = &maxEthRxRingSize
	var maxEthTxRingSize int64 = 4096
	this.MaxEthTxRingSize = &maxEthTxRingSize
	var maxRocev2Interfaces int64 = 2
	this.MaxRocev2Interfaces = &maxRocev2Interfaces
	var pciLink int64 = 0
	this.PciLink = &pciLink
	return &this
}

// GetClassId returns the ClassId field value
func (o *CapabilityAdapterUnitDescriptorAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CapabilityAdapterUnitDescriptorAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CapabilityAdapterUnitDescriptorAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CapabilityAdapterUnitDescriptorAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CapabilityAdapterUnitDescriptorAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CapabilityAdapterUnitDescriptorAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAdapterGeneration returns the AdapterGeneration field value if set, zero value otherwise.
func (o *CapabilityAdapterUnitDescriptorAllOf) GetAdapterGeneration() int32 {
	if o == nil || o.AdapterGeneration == nil {
		var ret int32
		return ret
	}
	return *o.AdapterGeneration
}

// GetAdapterGenerationOk returns a tuple with the AdapterGeneration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityAdapterUnitDescriptorAllOf) GetAdapterGenerationOk() (*int32, bool) {
	if o == nil || o.AdapterGeneration == nil {
		return nil, false
	}
	return o.AdapterGeneration, true
}

// HasAdapterGeneration returns a boolean if a field has been set.
func (o *CapabilityAdapterUnitDescriptorAllOf) HasAdapterGeneration() bool {
	if o != nil && o.AdapterGeneration != nil {
		return true
	}

	return false
}

// SetAdapterGeneration gets a reference to the given int32 and assigns it to the AdapterGeneration field.
func (o *CapabilityAdapterUnitDescriptorAllOf) SetAdapterGeneration(v int32) {
	o.AdapterGeneration = &v
}

// GetConnectivityOrder returns the ConnectivityOrder field value if set, zero value otherwise.
func (o *CapabilityAdapterUnitDescriptorAllOf) GetConnectivityOrder() string {
	if o == nil || o.ConnectivityOrder == nil {
		var ret string
		return ret
	}
	return *o.ConnectivityOrder
}

// GetConnectivityOrderOk returns a tuple with the ConnectivityOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityAdapterUnitDescriptorAllOf) GetConnectivityOrderOk() (*string, bool) {
	if o == nil || o.ConnectivityOrder == nil {
		return nil, false
	}
	return o.ConnectivityOrder, true
}

// HasConnectivityOrder returns a boolean if a field has been set.
func (o *CapabilityAdapterUnitDescriptorAllOf) HasConnectivityOrder() bool {
	if o != nil && o.ConnectivityOrder != nil {
		return true
	}

	return false
}

// SetConnectivityOrder gets a reference to the given string and assigns it to the ConnectivityOrder field.
func (o *CapabilityAdapterUnitDescriptorAllOf) SetConnectivityOrder(v string) {
	o.ConnectivityOrder = &v
}

// GetEthernetPortSpeed returns the EthernetPortSpeed field value if set, zero value otherwise.
func (o *CapabilityAdapterUnitDescriptorAllOf) GetEthernetPortSpeed() int64 {
	if o == nil || o.EthernetPortSpeed == nil {
		var ret int64
		return ret
	}
	return *o.EthernetPortSpeed
}

// GetEthernetPortSpeedOk returns a tuple with the EthernetPortSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityAdapterUnitDescriptorAllOf) GetEthernetPortSpeedOk() (*int64, bool) {
	if o == nil || o.EthernetPortSpeed == nil {
		return nil, false
	}
	return o.EthernetPortSpeed, true
}

// HasEthernetPortSpeed returns a boolean if a field has been set.
func (o *CapabilityAdapterUnitDescriptorAllOf) HasEthernetPortSpeed() bool {
	if o != nil && o.EthernetPortSpeed != nil {
		return true
	}

	return false
}

// SetEthernetPortSpeed gets a reference to the given int64 and assigns it to the EthernetPortSpeed field.
func (o *CapabilityAdapterUnitDescriptorAllOf) SetEthernetPortSpeed(v int64) {
	o.EthernetPortSpeed = &v
}

// GetFeatures returns the Features field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CapabilityAdapterUnitDescriptorAllOf) GetFeatures() []CapabilityFeatureConfig {
	if o == nil {
		var ret []CapabilityFeatureConfig
		return ret
	}
	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CapabilityAdapterUnitDescriptorAllOf) GetFeaturesOk() ([]CapabilityFeatureConfig, bool) {
	if o == nil || o.Features == nil {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *CapabilityAdapterUnitDescriptorAllOf) HasFeatures() bool {
	if o != nil && o.Features != nil {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given []CapabilityFeatureConfig and assigns it to the Features field.
func (o *CapabilityAdapterUnitDescriptorAllOf) SetFeatures(v []CapabilityFeatureConfig) {
	o.Features = v
}

// GetFibreChannelPortSpeed returns the FibreChannelPortSpeed field value if set, zero value otherwise.
func (o *CapabilityAdapterUnitDescriptorAllOf) GetFibreChannelPortSpeed() int64 {
	if o == nil || o.FibreChannelPortSpeed == nil {
		var ret int64
		return ret
	}
	return *o.FibreChannelPortSpeed
}

// GetFibreChannelPortSpeedOk returns a tuple with the FibreChannelPortSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityAdapterUnitDescriptorAllOf) GetFibreChannelPortSpeedOk() (*int64, bool) {
	if o == nil || o.FibreChannelPortSpeed == nil {
		return nil, false
	}
	return o.FibreChannelPortSpeed, true
}

// HasFibreChannelPortSpeed returns a boolean if a field has been set.
func (o *CapabilityAdapterUnitDescriptorAllOf) HasFibreChannelPortSpeed() bool {
	if o != nil && o.FibreChannelPortSpeed != nil {
		return true
	}

	return false
}

// SetFibreChannelPortSpeed gets a reference to the given int64 and assigns it to the FibreChannelPortSpeed field.
func (o *CapabilityAdapterUnitDescriptorAllOf) SetFibreChannelPortSpeed(v int64) {
	o.FibreChannelPortSpeed = &v
}

// GetFibreChannelScsiIoqLimit returns the FibreChannelScsiIoqLimit field value if set, zero value otherwise.
func (o *CapabilityAdapterUnitDescriptorAllOf) GetFibreChannelScsiIoqLimit() int64 {
	if o == nil || o.FibreChannelScsiIoqLimit == nil {
		var ret int64
		return ret
	}
	return *o.FibreChannelScsiIoqLimit
}

// GetFibreChannelScsiIoqLimitOk returns a tuple with the FibreChannelScsiIoqLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityAdapterUnitDescriptorAllOf) GetFibreChannelScsiIoqLimitOk() (*int64, bool) {
	if o == nil || o.FibreChannelScsiIoqLimit == nil {
		return nil, false
	}
	return o.FibreChannelScsiIoqLimit, true
}

// HasFibreChannelScsiIoqLimit returns a boolean if a field has been set.
func (o *CapabilityAdapterUnitDescriptorAllOf) HasFibreChannelScsiIoqLimit() bool {
	if o != nil && o.FibreChannelScsiIoqLimit != nil {
		return true
	}

	return false
}

// SetFibreChannelScsiIoqLimit gets a reference to the given int64 and assigns it to the FibreChannelScsiIoqLimit field.
func (o *CapabilityAdapterUnitDescriptorAllOf) SetFibreChannelScsiIoqLimit(v int64) {
	o.FibreChannelScsiIoqLimit = &v
}

// GetIsAzureQosSupported returns the IsAzureQosSupported field value if set, zero value otherwise.
func (o *CapabilityAdapterUnitDescriptorAllOf) GetIsAzureQosSupported() bool {
	if o == nil || o.IsAzureQosSupported == nil {
		var ret bool
		return ret
	}
	return *o.IsAzureQosSupported
}

// GetIsAzureQosSupportedOk returns a tuple with the IsAzureQosSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityAdapterUnitDescriptorAllOf) GetIsAzureQosSupportedOk() (*bool, bool) {
	if o == nil || o.IsAzureQosSupported == nil {
		return nil, false
	}
	return o.IsAzureQosSupported, true
}

// HasIsAzureQosSupported returns a boolean if a field has been set.
func (o *CapabilityAdapterUnitDescriptorAllOf) HasIsAzureQosSupported() bool {
	if o != nil && o.IsAzureQosSupported != nil {
		return true
	}

	return false
}

// SetIsAzureQosSupported gets a reference to the given bool and assigns it to the IsAzureQosSupported field.
func (o *CapabilityAdapterUnitDescriptorAllOf) SetIsAzureQosSupported(v bool) {
	o.IsAzureQosSupported = &v
}

// GetIsGeneveSupported returns the IsGeneveSupported field value if set, zero value otherwise.
func (o *CapabilityAdapterUnitDescriptorAllOf) GetIsGeneveSupported() bool {
	if o == nil || o.IsGeneveSupported == nil {
		var ret bool
		return ret
	}
	return *o.IsGeneveSupported
}

// GetIsGeneveSupportedOk returns a tuple with the IsGeneveSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityAdapterUnitDescriptorAllOf) GetIsGeneveSupportedOk() (*bool, bool) {
	if o == nil || o.IsGeneveSupported == nil {
		return nil, false
	}
	return o.IsGeneveSupported, true
}

// HasIsGeneveSupported returns a boolean if a field has been set.
func (o *CapabilityAdapterUnitDescriptorAllOf) HasIsGeneveSupported() bool {
	if o != nil && o.IsGeneveSupported != nil {
		return true
	}

	return false
}

// SetIsGeneveSupported gets a reference to the given bool and assigns it to the IsGeneveSupported field.
func (o *CapabilityAdapterUnitDescriptorAllOf) SetIsGeneveSupported(v bool) {
	o.IsGeneveSupported = &v
}

// GetMaxEthRxRingSize returns the MaxEthRxRingSize field value if set, zero value otherwise.
func (o *CapabilityAdapterUnitDescriptorAllOf) GetMaxEthRxRingSize() int64 {
	if o == nil || o.MaxEthRxRingSize == nil {
		var ret int64
		return ret
	}
	return *o.MaxEthRxRingSize
}

// GetMaxEthRxRingSizeOk returns a tuple with the MaxEthRxRingSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityAdapterUnitDescriptorAllOf) GetMaxEthRxRingSizeOk() (*int64, bool) {
	if o == nil || o.MaxEthRxRingSize == nil {
		return nil, false
	}
	return o.MaxEthRxRingSize, true
}

// HasMaxEthRxRingSize returns a boolean if a field has been set.
func (o *CapabilityAdapterUnitDescriptorAllOf) HasMaxEthRxRingSize() bool {
	if o != nil && o.MaxEthRxRingSize != nil {
		return true
	}

	return false
}

// SetMaxEthRxRingSize gets a reference to the given int64 and assigns it to the MaxEthRxRingSize field.
func (o *CapabilityAdapterUnitDescriptorAllOf) SetMaxEthRxRingSize(v int64) {
	o.MaxEthRxRingSize = &v
}

// GetMaxEthTxRingSize returns the MaxEthTxRingSize field value if set, zero value otherwise.
func (o *CapabilityAdapterUnitDescriptorAllOf) GetMaxEthTxRingSize() int64 {
	if o == nil || o.MaxEthTxRingSize == nil {
		var ret int64
		return ret
	}
	return *o.MaxEthTxRingSize
}

// GetMaxEthTxRingSizeOk returns a tuple with the MaxEthTxRingSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityAdapterUnitDescriptorAllOf) GetMaxEthTxRingSizeOk() (*int64, bool) {
	if o == nil || o.MaxEthTxRingSize == nil {
		return nil, false
	}
	return o.MaxEthTxRingSize, true
}

// HasMaxEthTxRingSize returns a boolean if a field has been set.
func (o *CapabilityAdapterUnitDescriptorAllOf) HasMaxEthTxRingSize() bool {
	if o != nil && o.MaxEthTxRingSize != nil {
		return true
	}

	return false
}

// SetMaxEthTxRingSize gets a reference to the given int64 and assigns it to the MaxEthTxRingSize field.
func (o *CapabilityAdapterUnitDescriptorAllOf) SetMaxEthTxRingSize(v int64) {
	o.MaxEthTxRingSize = &v
}

// GetMaxRocev2Interfaces returns the MaxRocev2Interfaces field value if set, zero value otherwise.
func (o *CapabilityAdapterUnitDescriptorAllOf) GetMaxRocev2Interfaces() int64 {
	if o == nil || o.MaxRocev2Interfaces == nil {
		var ret int64
		return ret
	}
	return *o.MaxRocev2Interfaces
}

// GetMaxRocev2InterfacesOk returns a tuple with the MaxRocev2Interfaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityAdapterUnitDescriptorAllOf) GetMaxRocev2InterfacesOk() (*int64, bool) {
	if o == nil || o.MaxRocev2Interfaces == nil {
		return nil, false
	}
	return o.MaxRocev2Interfaces, true
}

// HasMaxRocev2Interfaces returns a boolean if a field has been set.
func (o *CapabilityAdapterUnitDescriptorAllOf) HasMaxRocev2Interfaces() bool {
	if o != nil && o.MaxRocev2Interfaces != nil {
		return true
	}

	return false
}

// SetMaxRocev2Interfaces gets a reference to the given int64 and assigns it to the MaxRocev2Interfaces field.
func (o *CapabilityAdapterUnitDescriptorAllOf) SetMaxRocev2Interfaces(v int64) {
	o.MaxRocev2Interfaces = &v
}

// GetNumDcePorts returns the NumDcePorts field value if set, zero value otherwise.
func (o *CapabilityAdapterUnitDescriptorAllOf) GetNumDcePorts() int64 {
	if o == nil || o.NumDcePorts == nil {
		var ret int64
		return ret
	}
	return *o.NumDcePorts
}

// GetNumDcePortsOk returns a tuple with the NumDcePorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityAdapterUnitDescriptorAllOf) GetNumDcePortsOk() (*int64, bool) {
	if o == nil || o.NumDcePorts == nil {
		return nil, false
	}
	return o.NumDcePorts, true
}

// HasNumDcePorts returns a boolean if a field has been set.
func (o *CapabilityAdapterUnitDescriptorAllOf) HasNumDcePorts() bool {
	if o != nil && o.NumDcePorts != nil {
		return true
	}

	return false
}

// SetNumDcePorts gets a reference to the given int64 and assigns it to the NumDcePorts field.
func (o *CapabilityAdapterUnitDescriptorAllOf) SetNumDcePorts(v int64) {
	o.NumDcePorts = &v
}

// GetPciLink returns the PciLink field value if set, zero value otherwise.
func (o *CapabilityAdapterUnitDescriptorAllOf) GetPciLink() int64 {
	if o == nil || o.PciLink == nil {
		var ret int64
		return ret
	}
	return *o.PciLink
}

// GetPciLinkOk returns a tuple with the PciLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityAdapterUnitDescriptorAllOf) GetPciLinkOk() (*int64, bool) {
	if o == nil || o.PciLink == nil {
		return nil, false
	}
	return o.PciLink, true
}

// HasPciLink returns a boolean if a field has been set.
func (o *CapabilityAdapterUnitDescriptorAllOf) HasPciLink() bool {
	if o != nil && o.PciLink != nil {
		return true
	}

	return false
}

// SetPciLink gets a reference to the given int64 and assigns it to the PciLink field.
func (o *CapabilityAdapterUnitDescriptorAllOf) SetPciLink(v int64) {
	o.PciLink = &v
}

// GetPromCardType returns the PromCardType field value if set, zero value otherwise.
func (o *CapabilityAdapterUnitDescriptorAllOf) GetPromCardType() string {
	if o == nil || o.PromCardType == nil {
		var ret string
		return ret
	}
	return *o.PromCardType
}

// GetPromCardTypeOk returns a tuple with the PromCardType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityAdapterUnitDescriptorAllOf) GetPromCardTypeOk() (*string, bool) {
	if o == nil || o.PromCardType == nil {
		return nil, false
	}
	return o.PromCardType, true
}

// HasPromCardType returns a boolean if a field has been set.
func (o *CapabilityAdapterUnitDescriptorAllOf) HasPromCardType() bool {
	if o != nil && o.PromCardType != nil {
		return true
	}

	return false
}

// SetPromCardType gets a reference to the given string and assigns it to the PromCardType field.
func (o *CapabilityAdapterUnitDescriptorAllOf) SetPromCardType(v string) {
	o.PromCardType = &v
}

func (o CapabilityAdapterUnitDescriptorAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AdapterGeneration != nil {
		toSerialize["AdapterGeneration"] = o.AdapterGeneration
	}
	if o.ConnectivityOrder != nil {
		toSerialize["ConnectivityOrder"] = o.ConnectivityOrder
	}
	if o.EthernetPortSpeed != nil {
		toSerialize["EthernetPortSpeed"] = o.EthernetPortSpeed
	}
	if o.Features != nil {
		toSerialize["Features"] = o.Features
	}
	if o.FibreChannelPortSpeed != nil {
		toSerialize["FibreChannelPortSpeed"] = o.FibreChannelPortSpeed
	}
	if o.FibreChannelScsiIoqLimit != nil {
		toSerialize["FibreChannelScsiIoqLimit"] = o.FibreChannelScsiIoqLimit
	}
	if o.IsAzureQosSupported != nil {
		toSerialize["IsAzureQosSupported"] = o.IsAzureQosSupported
	}
	if o.IsGeneveSupported != nil {
		toSerialize["IsGeneveSupported"] = o.IsGeneveSupported
	}
	if o.MaxEthRxRingSize != nil {
		toSerialize["MaxEthRxRingSize"] = o.MaxEthRxRingSize
	}
	if o.MaxEthTxRingSize != nil {
		toSerialize["MaxEthTxRingSize"] = o.MaxEthTxRingSize
	}
	if o.MaxRocev2Interfaces != nil {
		toSerialize["MaxRocev2Interfaces"] = o.MaxRocev2Interfaces
	}
	if o.NumDcePorts != nil {
		toSerialize["NumDcePorts"] = o.NumDcePorts
	}
	if o.PciLink != nil {
		toSerialize["PciLink"] = o.PciLink
	}
	if o.PromCardType != nil {
		toSerialize["PromCardType"] = o.PromCardType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CapabilityAdapterUnitDescriptorAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varCapabilityAdapterUnitDescriptorAllOf := _CapabilityAdapterUnitDescriptorAllOf{}

	if err = json.Unmarshal(bytes, &varCapabilityAdapterUnitDescriptorAllOf); err == nil {
		*o = CapabilityAdapterUnitDescriptorAllOf(varCapabilityAdapterUnitDescriptorAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AdapterGeneration")
		delete(additionalProperties, "ConnectivityOrder")
		delete(additionalProperties, "EthernetPortSpeed")
		delete(additionalProperties, "Features")
		delete(additionalProperties, "FibreChannelPortSpeed")
		delete(additionalProperties, "FibreChannelScsiIoqLimit")
		delete(additionalProperties, "IsAzureQosSupported")
		delete(additionalProperties, "IsGeneveSupported")
		delete(additionalProperties, "MaxEthRxRingSize")
		delete(additionalProperties, "MaxEthTxRingSize")
		delete(additionalProperties, "MaxRocev2Interfaces")
		delete(additionalProperties, "NumDcePorts")
		delete(additionalProperties, "PciLink")
		delete(additionalProperties, "PromCardType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCapabilityAdapterUnitDescriptorAllOf struct {
	value *CapabilityAdapterUnitDescriptorAllOf
	isSet bool
}

func (v NullableCapabilityAdapterUnitDescriptorAllOf) Get() *CapabilityAdapterUnitDescriptorAllOf {
	return v.value
}

func (v *NullableCapabilityAdapterUnitDescriptorAllOf) Set(val *CapabilityAdapterUnitDescriptorAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilityAdapterUnitDescriptorAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilityAdapterUnitDescriptorAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilityAdapterUnitDescriptorAllOf(val *CapabilityAdapterUnitDescriptorAllOf) *NullableCapabilityAdapterUnitDescriptorAllOf {
	return &NullableCapabilityAdapterUnitDescriptorAllOf{value: val, isSet: true}
}

func (v NullableCapabilityAdapterUnitDescriptorAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilityAdapterUnitDescriptorAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
