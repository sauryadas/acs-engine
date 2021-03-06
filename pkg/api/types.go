package api

import (
	neturl "net/url"
)

// TypeMeta describes an individual API model object
type TypeMeta struct {
	// APIVersion is on every object
	APIVersion string `json:"apiVersion"`
}

// SubscriptionState represents the state of the subscription
type SubscriptionState int

// Subscription represents the customer subscription
type Subscription struct {
	ID    string
	State SubscriptionState
}

// ResourcePurchasePlan defines resource plan as required by ARM
// for billing purposes.
type ResourcePurchasePlan struct {
	Name          string `json:"name"`
	Product       string `json:"product"`
	PromotionCode string `json:"promotionCode"`
	Publisher     string `json:"publisher"`
}

// ContainerService complies with the ARM model of
// resource definition in a JSON template.
type ContainerService struct {
	APIVersion string               `json:"apiVersion"`
	ID         string               `json:"id"`
	Location   string               `json:"location"`
	Name       string               `json:"name"`
	Plan       ResourcePurchasePlan `json:"plan"`
	Tags       map[string]string    `json:"tags"`
	Type       string               `json:"type"`

	Properties Properties `json:"properties"`
}

// Properties represents the ACS cluster definition
type Properties struct {
	ProvisioningState       ProvisioningState       `json:"provisioningState"`
	OrchestratorProfile     OrchestratorProfile     `json:"orchestratorProfile"`
	MasterProfile           MasterProfile           `json:"masterProfile"`
	AgentPoolProfiles       []AgentPoolProfile      `json:"agentPoolProfiles"`
	LinuxProfile            LinuxProfile            `json:"linuxProfile"`
	WindowsProfile          WindowsProfile          `json:"windowsProfile"`
	DiagnosticsProfile      DiagnosticsProfile      `json:"diagnosticsProfile"`
	JumpboxProfile          JumpboxProfile          `json:"jumpboxProfile"`
	ServicePrincipalProfile ServicePrincipalProfile `json:"servicePrincipalProfile"`
	CertificateProfile      CertificateProfile      `json:"certificateProfile"`
}

// ServicePrincipalProfile contains the client and secret used by the cluster for Azure Resource CRUD
type ServicePrincipalProfile struct {
	ClientID string `json:"servicePrincipalClientID,omitempty"`
	Secret   string `json:"servicePrincipalClientSecret,omitempty"`
}

// CertificateProfile represents the definition of the master cluster
type CertificateProfile struct {
	// CaCertificate is the certificate authority certificate.
	CaCertificate string `json:"caCertificate,omitempty"`
	// ApiServerCertificate is the rest api server certificate, and signed by the CA
	APIServerCertificate string `json:"apiServerCertificate,omitempty"`
	// ApiServerPrivateKey is the rest api server private key, and signed by the CA
	APIServerPrivateKey string `json:"apiServerPrivateKey,omitempty"`
	// ClientCertificate is the certificate used by the client kubelet services and signed by the CA
	ClientCertificate string `json:"clientCertificate,omitempty"`
	// ClientPrivateKey is the private key used by the client kubelet services and signed by the CA
	ClientPrivateKey string `json:"clientPrivateKey,omitempty"`
	// KubeConfigCertificate is the client certificate used for kubectl cli and signed by the CA
	KubeConfigCertificate string `json:"kubeConfigCertificate,omitempty"`
	// KubeConfigPrivateKey is the client private key used for kubectl cli and signed by the CA
	KubeConfigPrivateKey string `json:"kubeConfigPrivateKey,omitempty"`
	// caPrivateKey is an internal field only set if generation required
	caPrivateKey string
}

// LinuxProfile represents the linux parameters passed to the cluster
type LinuxProfile struct {
	AdminUsername string `json:"adminUsername"`
	SSH           struct {
		PublicKeys []struct {
			KeyData string `json:"keyData"`
		} `json:"publicKeys"`
	} `json:"ssh"`
}

// WindowsProfile represents the windows parameters passed to the cluster
type WindowsProfile struct {
	AdminUsername string `json:"adminUsername"`
	AdminPassword string `json:"adminPassword"`
}

// ProvisioningState represents the current state of container service resource.
type ProvisioningState string

const (
	// Creating means ContainerService resource is being created.
	Creating ProvisioningState = "Creating"
	// Updating means an existing ContainerService resource is being updated
	Updating ProvisioningState = "Updating"
	// Failed means resource is in failed state
	Failed ProvisioningState = "Failed"
	// Succeeded means resource created succeeded during last create/update
	Succeeded ProvisioningState = "Succeeded"
	// Deleting means resource is in the process of being deleted
	Deleting ProvisioningState = "Deleting"
	// Migrating means resource is being migrated from one subscription or
	// resource group to another
	Migrating ProvisioningState = "Migrating"
)

// OrchestratorProfile contains Orchestrator properties
type OrchestratorProfile struct {
	OrchestratorType OrchestratorType `json:"orchestratorType"`
	KubernetesConfig KubernetesConfig `json:"kubernetesConfig,omitempty"`
}

// KubernetesConfig contains the Kubernetes config structure, containing
// Kubernetes specific configuration
type KubernetesConfig struct {
	KubernetesHyperkubeSpec string `json:"kubernetesHyperkubeSpec,omitempty"`
	KubectlVersion          string `json:"kubectlVersion,omitempty"`
}

// MasterProfile represents the definition of the master cluster
type MasterProfile struct {
	Count                    int    `json:"count"`
	DNSPrefix                string `json:"dnsPrefix"`
	VMSize                   string `json:"vmSize"`
	VnetSubnetID             string `json:"vnetSubnetID,omitempty"`
	FirstConsecutiveStaticIP string `json:"firstConsecutiveStaticIP,omitempty"`
	Subnet                   string `json:"subnet"`

	// Master LB public endpoint/FQDN with port
	// The format will be FQDN:2376
	// Not used during PUT, returned as part of GET
	FQDN string `json:"fqdn,omitempty"`
}

// AgentPoolProfile represents an agent pool definition
type AgentPoolProfile struct {
	Name                string `json:"name"`
	Count               int    `json:"count"`
	VMSize              string `json:"vmSize"`
	DNSPrefix           string `json:"dnsPrefix,omitempty"`
	OSType              OSType `json:"osType,omitempty"`
	Ports               []int  `json:"ports,omitempty"`
	AvailabilityProfile string `json:"availabilityProfile"`
	StorageProfile      string `json:"storageProfile,omitempty"`
	DiskSizesGB         []int  `json:"diskSizesGB,omitempty"`
	VnetSubnetID        string `json:"vnetSubnetID,omitempty"`
	Subnet              string `json:"subnet"`

	FQDN string `json:"fqdn,omitempty"`
}

// DiagnosticsProfile setting to enable/disable capturing
// diagnostics for VMs hosting container cluster.
type DiagnosticsProfile struct {
	VMDiagnostics VMDiagnostics `json:"vmDiagnostics"`
}

// VMDiagnostics contains settings to on/off boot diagnostics collection
// in RD Host
type VMDiagnostics struct {
	Enabled bool `json:"enabled"`

	// Specifies storage account Uri where Boot Diagnostics (CRP &
	// VMSS BootDiagostics) and VM Diagnostics logs (using Linux
	// Diagnostics Extension) will be stored. Uri will be of standard
	// blob domain. i.e. https://storageaccount.blob.core.windows.net/
	// This field is readonly as ACS RP will create a storage account
	// for the customer.
	StorageURL neturl.URL `json:"storageUrl"`
}

// OrchestratorType defines orchestrators supported by ACS
type OrchestratorType string

// JumpboxProfile dscribes properties of the jumpbox setup
// in the ACS container cluster.
type JumpboxProfile struct {
	OSType    OSType `json:"osType"`
	DNSPrefix string `json:"dnsPrefix"`

	// Jumpbox public endpoint/FQDN with port
	// The format will be FQDN:2376
	// Not used during PUT, returned as part of GET
	FQDN string `json:"fqdn,omitempty"`
}

// OSType represents OS types of agents
type OSType string

// HasWindows returns true if the cluster contains windows
func (a *Properties) HasWindows() bool {
	for _, agentPoolProfile := range a.AgentPoolProfiles {
		if agentPoolProfile.OSType == Windows {
			return true
		}
	}
	return false
}

// HasManagedDisks returns true if the cluster contains Managed Disks
func (a *Properties) HasManagedDisks() bool {
	for _, agentPoolProfile := range a.AgentPoolProfiles {
		if agentPoolProfile.StorageProfile == ManagedDisks {
			return true
		}
	}
	return false
}

// GetCAPrivateKey returns the ca private key
func (c *CertificateProfile) GetCAPrivateKey() string {
	return c.caPrivateKey
}

// SetCAPrivateKey sets the ca private key
func (c *CertificateProfile) SetCAPrivateKey(caPrivateKey string) {
	c.caPrivateKey = caPrivateKey
}

// IsCustomVNET returns true if the customer brought their own VNET
func (m *MasterProfile) IsCustomVNET() bool {
	return len(m.VnetSubnetID) > 0
}

// IsCustomVNET returns true if the customer brought their own VNET
func (a *AgentPoolProfile) IsCustomVNET() bool {
	return len(a.VnetSubnetID) > 0
}

// IsWindows returns true if the agent pool is windows
func (a *AgentPoolProfile) IsWindows() bool {
	return a.OSType == Windows
}

// IsAvailabilitySets returns true if the customer specified disks
func (a *AgentPoolProfile) IsAvailabilitySets() bool {
	return a.AvailabilityProfile == AvailabilitySet
}

// IsManagedDisks returns true if the customer specified disks
func (a *AgentPoolProfile) IsManagedDisks() bool {
	return a.StorageProfile == ManagedDisks
}

// IsStorageAccount returns true if the customer specified storage account
func (a *AgentPoolProfile) IsStorageAccount() bool {
	return a.StorageProfile == StorageAccount
}

// HasDisks returns true if the customer specified disks
func (a *AgentPoolProfile) HasDisks() bool {
	return len(a.DiskSizesGB) > 0
}
