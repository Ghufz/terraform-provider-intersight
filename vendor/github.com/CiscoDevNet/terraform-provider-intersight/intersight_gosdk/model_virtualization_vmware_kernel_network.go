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

// VirtualizationVmwareKernelNetwork Details of VMware Kernel Network.
type VirtualizationVmwareKernelNetwork struct {
	VirtualizationBaseKernelNetwork
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Indicates that fault tolerance logging is enabled on this kernel network.
	FaultToleranceLogging *bool    `json:"FaultToleranceLogging,omitempty"`
	IpAddress             []string `json:"IpAddress,omitempty"`
	// Standard MAC address assigned to this kernel network.
	MacAddress *string `json:"MacAddress,omitempty"`
	// Indicates that management traffic is enabled on this kernel network.
	Management *bool `json:"Management,omitempty"`
	// Maximum transmission unit configured on a kernel network.
	Mtu *int64 `json:"Mtu,omitempty"`
	// Subnet mask of the kernel network.
	SubnetMask *string `json:"SubnetMask,omitempty"`
	// Type of stack for the kernel network. It can be custom, default, vMotion or provisioning.
	TcpIpStack *string `json:"TcpIpStack,omitempty"`
	// Indicates that vmotion is enabled on this kernel network.
	Vmotion *bool `json:"Vmotion,omitempty"`
	// Indicates that vsan traffic is enabled on this kernel network.
	Vsan *bool `json:"Vsan,omitempty"`
	// Indicates that vsphere provisioning traffic is enabled on this kernel network.
	VsphereProvisioning *bool `json:"VsphereProvisioning,omitempty"`
	// Indicates that vsphere replication is enabled on this kernel network.
	VsphereReplication *bool `json:"VsphereReplication,omitempty"`
	// Indicates that vsphere replication nfc is enabled on this kernel network.
	VsphereReplicationNfc *bool                                               `json:"VsphereReplicationNfc,omitempty"`
	DistributedNetwork    *VirtualizationVmwareDistributedNetworkRelationship `json:"DistributedNetwork,omitempty"`
	Host                  *VirtualizationVmwareHostRelationship               `json:"Host,omitempty"`
	Network               *VirtualizationVmwareNetworkRelationship            `json:"Network,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _VirtualizationVmwareKernelNetwork VirtualizationVmwareKernelNetwork

// NewVirtualizationVmwareKernelNetwork instantiates a new VirtualizationVmwareKernelNetwork object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationVmwareKernelNetwork(classId string, objectType string) *VirtualizationVmwareKernelNetwork {
	this := VirtualizationVmwareKernelNetwork{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVirtualizationVmwareKernelNetworkWithDefaults instantiates a new VirtualizationVmwareKernelNetwork object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationVmwareKernelNetworkWithDefaults() *VirtualizationVmwareKernelNetwork {
	this := VirtualizationVmwareKernelNetwork{}
	var classId string = "virtualization.VmwareKernelNetwork"
	this.ClassId = classId
	var objectType string = "virtualization.VmwareKernelNetwork"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationVmwareKernelNetwork) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareKernelNetwork) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationVmwareKernelNetwork) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationVmwareKernelNetwork) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareKernelNetwork) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationVmwareKernelNetwork) SetObjectType(v string) {
	o.ObjectType = v
}

// GetFaultToleranceLogging returns the FaultToleranceLogging field value if set, zero value otherwise.
func (o *VirtualizationVmwareKernelNetwork) GetFaultToleranceLogging() bool {
	if o == nil || o.FaultToleranceLogging == nil {
		var ret bool
		return ret
	}
	return *o.FaultToleranceLogging
}

// GetFaultToleranceLoggingOk returns a tuple with the FaultToleranceLogging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareKernelNetwork) GetFaultToleranceLoggingOk() (*bool, bool) {
	if o == nil || o.FaultToleranceLogging == nil {
		return nil, false
	}
	return o.FaultToleranceLogging, true
}

// HasFaultToleranceLogging returns a boolean if a field has been set.
func (o *VirtualizationVmwareKernelNetwork) HasFaultToleranceLogging() bool {
	if o != nil && o.FaultToleranceLogging != nil {
		return true
	}

	return false
}

// SetFaultToleranceLogging gets a reference to the given bool and assigns it to the FaultToleranceLogging field.
func (o *VirtualizationVmwareKernelNetwork) SetFaultToleranceLogging(v bool) {
	o.FaultToleranceLogging = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationVmwareKernelNetwork) GetIpAddress() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationVmwareKernelNetwork) GetIpAddressOk() ([]string, bool) {
	if o == nil || o.IpAddress == nil {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *VirtualizationVmwareKernelNetwork) HasIpAddress() bool {
	if o != nil && o.IpAddress != nil {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given []string and assigns it to the IpAddress field.
func (o *VirtualizationVmwareKernelNetwork) SetIpAddress(v []string) {
	o.IpAddress = v
}

// GetMacAddress returns the MacAddress field value if set, zero value otherwise.
func (o *VirtualizationVmwareKernelNetwork) GetMacAddress() string {
	if o == nil || o.MacAddress == nil {
		var ret string
		return ret
	}
	return *o.MacAddress
}

// GetMacAddressOk returns a tuple with the MacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareKernelNetwork) GetMacAddressOk() (*string, bool) {
	if o == nil || o.MacAddress == nil {
		return nil, false
	}
	return o.MacAddress, true
}

// HasMacAddress returns a boolean if a field has been set.
func (o *VirtualizationVmwareKernelNetwork) HasMacAddress() bool {
	if o != nil && o.MacAddress != nil {
		return true
	}

	return false
}

// SetMacAddress gets a reference to the given string and assigns it to the MacAddress field.
func (o *VirtualizationVmwareKernelNetwork) SetMacAddress(v string) {
	o.MacAddress = &v
}

// GetManagement returns the Management field value if set, zero value otherwise.
func (o *VirtualizationVmwareKernelNetwork) GetManagement() bool {
	if o == nil || o.Management == nil {
		var ret bool
		return ret
	}
	return *o.Management
}

// GetManagementOk returns a tuple with the Management field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareKernelNetwork) GetManagementOk() (*bool, bool) {
	if o == nil || o.Management == nil {
		return nil, false
	}
	return o.Management, true
}

// HasManagement returns a boolean if a field has been set.
func (o *VirtualizationVmwareKernelNetwork) HasManagement() bool {
	if o != nil && o.Management != nil {
		return true
	}

	return false
}

// SetManagement gets a reference to the given bool and assigns it to the Management field.
func (o *VirtualizationVmwareKernelNetwork) SetManagement(v bool) {
	o.Management = &v
}

// GetMtu returns the Mtu field value if set, zero value otherwise.
func (o *VirtualizationVmwareKernelNetwork) GetMtu() int64 {
	if o == nil || o.Mtu == nil {
		var ret int64
		return ret
	}
	return *o.Mtu
}

// GetMtuOk returns a tuple with the Mtu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareKernelNetwork) GetMtuOk() (*int64, bool) {
	if o == nil || o.Mtu == nil {
		return nil, false
	}
	return o.Mtu, true
}

// HasMtu returns a boolean if a field has been set.
func (o *VirtualizationVmwareKernelNetwork) HasMtu() bool {
	if o != nil && o.Mtu != nil {
		return true
	}

	return false
}

// SetMtu gets a reference to the given int64 and assigns it to the Mtu field.
func (o *VirtualizationVmwareKernelNetwork) SetMtu(v int64) {
	o.Mtu = &v
}

// GetSubnetMask returns the SubnetMask field value if set, zero value otherwise.
func (o *VirtualizationVmwareKernelNetwork) GetSubnetMask() string {
	if o == nil || o.SubnetMask == nil {
		var ret string
		return ret
	}
	return *o.SubnetMask
}

// GetSubnetMaskOk returns a tuple with the SubnetMask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareKernelNetwork) GetSubnetMaskOk() (*string, bool) {
	if o == nil || o.SubnetMask == nil {
		return nil, false
	}
	return o.SubnetMask, true
}

// HasSubnetMask returns a boolean if a field has been set.
func (o *VirtualizationVmwareKernelNetwork) HasSubnetMask() bool {
	if o != nil && o.SubnetMask != nil {
		return true
	}

	return false
}

// SetSubnetMask gets a reference to the given string and assigns it to the SubnetMask field.
func (o *VirtualizationVmwareKernelNetwork) SetSubnetMask(v string) {
	o.SubnetMask = &v
}

// GetTcpIpStack returns the TcpIpStack field value if set, zero value otherwise.
func (o *VirtualizationVmwareKernelNetwork) GetTcpIpStack() string {
	if o == nil || o.TcpIpStack == nil {
		var ret string
		return ret
	}
	return *o.TcpIpStack
}

// GetTcpIpStackOk returns a tuple with the TcpIpStack field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareKernelNetwork) GetTcpIpStackOk() (*string, bool) {
	if o == nil || o.TcpIpStack == nil {
		return nil, false
	}
	return o.TcpIpStack, true
}

// HasTcpIpStack returns a boolean if a field has been set.
func (o *VirtualizationVmwareKernelNetwork) HasTcpIpStack() bool {
	if o != nil && o.TcpIpStack != nil {
		return true
	}

	return false
}

// SetTcpIpStack gets a reference to the given string and assigns it to the TcpIpStack field.
func (o *VirtualizationVmwareKernelNetwork) SetTcpIpStack(v string) {
	o.TcpIpStack = &v
}

// GetVmotion returns the Vmotion field value if set, zero value otherwise.
func (o *VirtualizationVmwareKernelNetwork) GetVmotion() bool {
	if o == nil || o.Vmotion == nil {
		var ret bool
		return ret
	}
	return *o.Vmotion
}

// GetVmotionOk returns a tuple with the Vmotion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareKernelNetwork) GetVmotionOk() (*bool, bool) {
	if o == nil || o.Vmotion == nil {
		return nil, false
	}
	return o.Vmotion, true
}

// HasVmotion returns a boolean if a field has been set.
func (o *VirtualizationVmwareKernelNetwork) HasVmotion() bool {
	if o != nil && o.Vmotion != nil {
		return true
	}

	return false
}

// SetVmotion gets a reference to the given bool and assigns it to the Vmotion field.
func (o *VirtualizationVmwareKernelNetwork) SetVmotion(v bool) {
	o.Vmotion = &v
}

// GetVsan returns the Vsan field value if set, zero value otherwise.
func (o *VirtualizationVmwareKernelNetwork) GetVsan() bool {
	if o == nil || o.Vsan == nil {
		var ret bool
		return ret
	}
	return *o.Vsan
}

// GetVsanOk returns a tuple with the Vsan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareKernelNetwork) GetVsanOk() (*bool, bool) {
	if o == nil || o.Vsan == nil {
		return nil, false
	}
	return o.Vsan, true
}

// HasVsan returns a boolean if a field has been set.
func (o *VirtualizationVmwareKernelNetwork) HasVsan() bool {
	if o != nil && o.Vsan != nil {
		return true
	}

	return false
}

// SetVsan gets a reference to the given bool and assigns it to the Vsan field.
func (o *VirtualizationVmwareKernelNetwork) SetVsan(v bool) {
	o.Vsan = &v
}

// GetVsphereProvisioning returns the VsphereProvisioning field value if set, zero value otherwise.
func (o *VirtualizationVmwareKernelNetwork) GetVsphereProvisioning() bool {
	if o == nil || o.VsphereProvisioning == nil {
		var ret bool
		return ret
	}
	return *o.VsphereProvisioning
}

// GetVsphereProvisioningOk returns a tuple with the VsphereProvisioning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareKernelNetwork) GetVsphereProvisioningOk() (*bool, bool) {
	if o == nil || o.VsphereProvisioning == nil {
		return nil, false
	}
	return o.VsphereProvisioning, true
}

// HasVsphereProvisioning returns a boolean if a field has been set.
func (o *VirtualizationVmwareKernelNetwork) HasVsphereProvisioning() bool {
	if o != nil && o.VsphereProvisioning != nil {
		return true
	}

	return false
}

// SetVsphereProvisioning gets a reference to the given bool and assigns it to the VsphereProvisioning field.
func (o *VirtualizationVmwareKernelNetwork) SetVsphereProvisioning(v bool) {
	o.VsphereProvisioning = &v
}

// GetVsphereReplication returns the VsphereReplication field value if set, zero value otherwise.
func (o *VirtualizationVmwareKernelNetwork) GetVsphereReplication() bool {
	if o == nil || o.VsphereReplication == nil {
		var ret bool
		return ret
	}
	return *o.VsphereReplication
}

// GetVsphereReplicationOk returns a tuple with the VsphereReplication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareKernelNetwork) GetVsphereReplicationOk() (*bool, bool) {
	if o == nil || o.VsphereReplication == nil {
		return nil, false
	}
	return o.VsphereReplication, true
}

// HasVsphereReplication returns a boolean if a field has been set.
func (o *VirtualizationVmwareKernelNetwork) HasVsphereReplication() bool {
	if o != nil && o.VsphereReplication != nil {
		return true
	}

	return false
}

// SetVsphereReplication gets a reference to the given bool and assigns it to the VsphereReplication field.
func (o *VirtualizationVmwareKernelNetwork) SetVsphereReplication(v bool) {
	o.VsphereReplication = &v
}

// GetVsphereReplicationNfc returns the VsphereReplicationNfc field value if set, zero value otherwise.
func (o *VirtualizationVmwareKernelNetwork) GetVsphereReplicationNfc() bool {
	if o == nil || o.VsphereReplicationNfc == nil {
		var ret bool
		return ret
	}
	return *o.VsphereReplicationNfc
}

// GetVsphereReplicationNfcOk returns a tuple with the VsphereReplicationNfc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareKernelNetwork) GetVsphereReplicationNfcOk() (*bool, bool) {
	if o == nil || o.VsphereReplicationNfc == nil {
		return nil, false
	}
	return o.VsphereReplicationNfc, true
}

// HasVsphereReplicationNfc returns a boolean if a field has been set.
func (o *VirtualizationVmwareKernelNetwork) HasVsphereReplicationNfc() bool {
	if o != nil && o.VsphereReplicationNfc != nil {
		return true
	}

	return false
}

// SetVsphereReplicationNfc gets a reference to the given bool and assigns it to the VsphereReplicationNfc field.
func (o *VirtualizationVmwareKernelNetwork) SetVsphereReplicationNfc(v bool) {
	o.VsphereReplicationNfc = &v
}

// GetDistributedNetwork returns the DistributedNetwork field value if set, zero value otherwise.
func (o *VirtualizationVmwareKernelNetwork) GetDistributedNetwork() VirtualizationVmwareDistributedNetworkRelationship {
	if o == nil || o.DistributedNetwork == nil {
		var ret VirtualizationVmwareDistributedNetworkRelationship
		return ret
	}
	return *o.DistributedNetwork
}

// GetDistributedNetworkOk returns a tuple with the DistributedNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareKernelNetwork) GetDistributedNetworkOk() (*VirtualizationVmwareDistributedNetworkRelationship, bool) {
	if o == nil || o.DistributedNetwork == nil {
		return nil, false
	}
	return o.DistributedNetwork, true
}

// HasDistributedNetwork returns a boolean if a field has been set.
func (o *VirtualizationVmwareKernelNetwork) HasDistributedNetwork() bool {
	if o != nil && o.DistributedNetwork != nil {
		return true
	}

	return false
}

// SetDistributedNetwork gets a reference to the given VirtualizationVmwareDistributedNetworkRelationship and assigns it to the DistributedNetwork field.
func (o *VirtualizationVmwareKernelNetwork) SetDistributedNetwork(v VirtualizationVmwareDistributedNetworkRelationship) {
	o.DistributedNetwork = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *VirtualizationVmwareKernelNetwork) GetHost() VirtualizationVmwareHostRelationship {
	if o == nil || o.Host == nil {
		var ret VirtualizationVmwareHostRelationship
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareKernelNetwork) GetHostOk() (*VirtualizationVmwareHostRelationship, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *VirtualizationVmwareKernelNetwork) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given VirtualizationVmwareHostRelationship and assigns it to the Host field.
func (o *VirtualizationVmwareKernelNetwork) SetHost(v VirtualizationVmwareHostRelationship) {
	o.Host = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *VirtualizationVmwareKernelNetwork) GetNetwork() VirtualizationVmwareNetworkRelationship {
	if o == nil || o.Network == nil {
		var ret VirtualizationVmwareNetworkRelationship
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareKernelNetwork) GetNetworkOk() (*VirtualizationVmwareNetworkRelationship, bool) {
	if o == nil || o.Network == nil {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *VirtualizationVmwareKernelNetwork) HasNetwork() bool {
	if o != nil && o.Network != nil {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given VirtualizationVmwareNetworkRelationship and assigns it to the Network field.
func (o *VirtualizationVmwareKernelNetwork) SetNetwork(v VirtualizationVmwareNetworkRelationship) {
	o.Network = &v
}

func (o VirtualizationVmwareKernelNetwork) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedVirtualizationBaseKernelNetwork, errVirtualizationBaseKernelNetwork := json.Marshal(o.VirtualizationBaseKernelNetwork)
	if errVirtualizationBaseKernelNetwork != nil {
		return []byte{}, errVirtualizationBaseKernelNetwork
	}
	errVirtualizationBaseKernelNetwork = json.Unmarshal([]byte(serializedVirtualizationBaseKernelNetwork), &toSerialize)
	if errVirtualizationBaseKernelNetwork != nil {
		return []byte{}, errVirtualizationBaseKernelNetwork
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.FaultToleranceLogging != nil {
		toSerialize["FaultToleranceLogging"] = o.FaultToleranceLogging
	}
	if o.IpAddress != nil {
		toSerialize["IpAddress"] = o.IpAddress
	}
	if o.MacAddress != nil {
		toSerialize["MacAddress"] = o.MacAddress
	}
	if o.Management != nil {
		toSerialize["Management"] = o.Management
	}
	if o.Mtu != nil {
		toSerialize["Mtu"] = o.Mtu
	}
	if o.SubnetMask != nil {
		toSerialize["SubnetMask"] = o.SubnetMask
	}
	if o.TcpIpStack != nil {
		toSerialize["TcpIpStack"] = o.TcpIpStack
	}
	if o.Vmotion != nil {
		toSerialize["Vmotion"] = o.Vmotion
	}
	if o.Vsan != nil {
		toSerialize["Vsan"] = o.Vsan
	}
	if o.VsphereProvisioning != nil {
		toSerialize["VsphereProvisioning"] = o.VsphereProvisioning
	}
	if o.VsphereReplication != nil {
		toSerialize["VsphereReplication"] = o.VsphereReplication
	}
	if o.VsphereReplicationNfc != nil {
		toSerialize["VsphereReplicationNfc"] = o.VsphereReplicationNfc
	}
	if o.DistributedNetwork != nil {
		toSerialize["DistributedNetwork"] = o.DistributedNetwork
	}
	if o.Host != nil {
		toSerialize["Host"] = o.Host
	}
	if o.Network != nil {
		toSerialize["Network"] = o.Network
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationVmwareKernelNetwork) UnmarshalJSON(bytes []byte) (err error) {
	type VirtualizationVmwareKernelNetworkWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Indicates that fault tolerance logging is enabled on this kernel network.
		FaultToleranceLogging *bool    `json:"FaultToleranceLogging,omitempty"`
		IpAddress             []string `json:"IpAddress,omitempty"`
		// Standard MAC address assigned to this kernel network.
		MacAddress *string `json:"MacAddress,omitempty"`
		// Indicates that management traffic is enabled on this kernel network.
		Management *bool `json:"Management,omitempty"`
		// Maximum transmission unit configured on a kernel network.
		Mtu *int64 `json:"Mtu,omitempty"`
		// Subnet mask of the kernel network.
		SubnetMask *string `json:"SubnetMask,omitempty"`
		// Type of stack for the kernel network. It can be custom, default, vMotion or provisioning.
		TcpIpStack *string `json:"TcpIpStack,omitempty"`
		// Indicates that vmotion is enabled on this kernel network.
		Vmotion *bool `json:"Vmotion,omitempty"`
		// Indicates that vsan traffic is enabled on this kernel network.
		Vsan *bool `json:"Vsan,omitempty"`
		// Indicates that vsphere provisioning traffic is enabled on this kernel network.
		VsphereProvisioning *bool `json:"VsphereProvisioning,omitempty"`
		// Indicates that vsphere replication is enabled on this kernel network.
		VsphereReplication *bool `json:"VsphereReplication,omitempty"`
		// Indicates that vsphere replication nfc is enabled on this kernel network.
		VsphereReplicationNfc *bool                                               `json:"VsphereReplicationNfc,omitempty"`
		DistributedNetwork    *VirtualizationVmwareDistributedNetworkRelationship `json:"DistributedNetwork,omitempty"`
		Host                  *VirtualizationVmwareHostRelationship               `json:"Host,omitempty"`
		Network               *VirtualizationVmwareNetworkRelationship            `json:"Network,omitempty"`
	}

	varVirtualizationVmwareKernelNetworkWithoutEmbeddedStruct := VirtualizationVmwareKernelNetworkWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVirtualizationVmwareKernelNetworkWithoutEmbeddedStruct)
	if err == nil {
		varVirtualizationVmwareKernelNetwork := _VirtualizationVmwareKernelNetwork{}
		varVirtualizationVmwareKernelNetwork.ClassId = varVirtualizationVmwareKernelNetworkWithoutEmbeddedStruct.ClassId
		varVirtualizationVmwareKernelNetwork.ObjectType = varVirtualizationVmwareKernelNetworkWithoutEmbeddedStruct.ObjectType
		varVirtualizationVmwareKernelNetwork.FaultToleranceLogging = varVirtualizationVmwareKernelNetworkWithoutEmbeddedStruct.FaultToleranceLogging
		varVirtualizationVmwareKernelNetwork.IpAddress = varVirtualizationVmwareKernelNetworkWithoutEmbeddedStruct.IpAddress
		varVirtualizationVmwareKernelNetwork.MacAddress = varVirtualizationVmwareKernelNetworkWithoutEmbeddedStruct.MacAddress
		varVirtualizationVmwareKernelNetwork.Management = varVirtualizationVmwareKernelNetworkWithoutEmbeddedStruct.Management
		varVirtualizationVmwareKernelNetwork.Mtu = varVirtualizationVmwareKernelNetworkWithoutEmbeddedStruct.Mtu
		varVirtualizationVmwareKernelNetwork.SubnetMask = varVirtualizationVmwareKernelNetworkWithoutEmbeddedStruct.SubnetMask
		varVirtualizationVmwareKernelNetwork.TcpIpStack = varVirtualizationVmwareKernelNetworkWithoutEmbeddedStruct.TcpIpStack
		varVirtualizationVmwareKernelNetwork.Vmotion = varVirtualizationVmwareKernelNetworkWithoutEmbeddedStruct.Vmotion
		varVirtualizationVmwareKernelNetwork.Vsan = varVirtualizationVmwareKernelNetworkWithoutEmbeddedStruct.Vsan
		varVirtualizationVmwareKernelNetwork.VsphereProvisioning = varVirtualizationVmwareKernelNetworkWithoutEmbeddedStruct.VsphereProvisioning
		varVirtualizationVmwareKernelNetwork.VsphereReplication = varVirtualizationVmwareKernelNetworkWithoutEmbeddedStruct.VsphereReplication
		varVirtualizationVmwareKernelNetwork.VsphereReplicationNfc = varVirtualizationVmwareKernelNetworkWithoutEmbeddedStruct.VsphereReplicationNfc
		varVirtualizationVmwareKernelNetwork.DistributedNetwork = varVirtualizationVmwareKernelNetworkWithoutEmbeddedStruct.DistributedNetwork
		varVirtualizationVmwareKernelNetwork.Host = varVirtualizationVmwareKernelNetworkWithoutEmbeddedStruct.Host
		varVirtualizationVmwareKernelNetwork.Network = varVirtualizationVmwareKernelNetworkWithoutEmbeddedStruct.Network
		*o = VirtualizationVmwareKernelNetwork(varVirtualizationVmwareKernelNetwork)
	} else {
		return err
	}

	varVirtualizationVmwareKernelNetwork := _VirtualizationVmwareKernelNetwork{}

	err = json.Unmarshal(bytes, &varVirtualizationVmwareKernelNetwork)
	if err == nil {
		o.VirtualizationBaseKernelNetwork = varVirtualizationVmwareKernelNetwork.VirtualizationBaseKernelNetwork
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "FaultToleranceLogging")
		delete(additionalProperties, "IpAddress")
		delete(additionalProperties, "MacAddress")
		delete(additionalProperties, "Management")
		delete(additionalProperties, "Mtu")
		delete(additionalProperties, "SubnetMask")
		delete(additionalProperties, "TcpIpStack")
		delete(additionalProperties, "Vmotion")
		delete(additionalProperties, "Vsan")
		delete(additionalProperties, "VsphereProvisioning")
		delete(additionalProperties, "VsphereReplication")
		delete(additionalProperties, "VsphereReplicationNfc")
		delete(additionalProperties, "DistributedNetwork")
		delete(additionalProperties, "Host")
		delete(additionalProperties, "Network")

		// remove fields from embedded structs
		reflectVirtualizationBaseKernelNetwork := reflect.ValueOf(o.VirtualizationBaseKernelNetwork)
		for i := 0; i < reflectVirtualizationBaseKernelNetwork.Type().NumField(); i++ {
			t := reflectVirtualizationBaseKernelNetwork.Type().Field(i)

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

type NullableVirtualizationVmwareKernelNetwork struct {
	value *VirtualizationVmwareKernelNetwork
	isSet bool
}

func (v NullableVirtualizationVmwareKernelNetwork) Get() *VirtualizationVmwareKernelNetwork {
	return v.value
}

func (v *NullableVirtualizationVmwareKernelNetwork) Set(val *VirtualizationVmwareKernelNetwork) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationVmwareKernelNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationVmwareKernelNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationVmwareKernelNetwork(val *VirtualizationVmwareKernelNetwork) *NullableVirtualizationVmwareKernelNetwork {
	return &NullableVirtualizationVmwareKernelNetwork{value: val, isSet: true}
}

func (v NullableVirtualizationVmwareKernelNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationVmwareKernelNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
