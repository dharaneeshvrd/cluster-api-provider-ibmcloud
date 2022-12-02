/*
Copyright 2022 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1beta2

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.
const (
	// MachineFinalizer allows IBMVPCMachineReconciler to clean up resources associated with IBMVPCMachine before
	// removing it from the apiserver.
	MachineFinalizer = "ibmvpcmachine.infrastructure.cluster.x-k8s.io"
)

// IBMVPCMachineSpec defines the desired state of IBMVPCMachine.
type IBMVPCMachineSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Name of the instance.
	Name string `json:"name,omitempty"`

	// Image is the id of OS image which would be install on the instance.
	// Example: r134-ed3f775f-ad7e-4e37-ae62-7199b4988b00
	// TODO: allow user to specify a image name is much reasonable. Example: ibm-ubuntu-18-04-1-minimal-amd64-2
	Image string `json:"image"`

	// Zone is the place where the instance should be created. Example: us-south-3
	// TODO: Actually zone is transparent to user. The field user can access is location. Example: Dallas 2
	Zone string `json:"zone"`

	// Profile indicates the flavor of instance. Example: bx2-8x32	means 8 vCPUs	32 GB RAM	16 Gbps
	// TODO: add a reference link of profile
	// +optional
	Profile string `json:"profile,omitempty"`

	// BootVolume contains machines's boot volume configurations like size, iops etc..
	// +optional
	BootVolume *VPCVolume `json:"bootVolume,omitempty"`

	// ProviderID is the unique identifier as specified by the cloud provider.
	// +optional
	ProviderID *string `json:"providerID,omitempty"`

	// PrimaryNetworkInterface is required to specify subnet.
	PrimaryNetworkInterface NetworkInterface `json:"primaryNetworkInterface,omitempty"`

	// SSHKeys is the SSH pub keys that will be used to access VM.
	SSHKeys []*string `json:"sshKeys,omitempty"`
}

// VPCVolume defines the volume information for the instance.
type VPCVolume struct {
	// DeleteVolumeOnInstanceDelete If set to true, when deleting the instance the volume will also be deleted.
	// Default is set as true
	// +kubebuilder:default=true
	// +optional
	DeleteVolumeOnInstanceDelete bool `json:"deleteVolumeOnInstanceDelete,omitempty"`

	// Name is the unique user-defined name for this volume.
	// Default will be autogenerated
	// +optional
	Name string `json:"name,omitempty"`

	// SizeGiB is the size of the virtual server's boot disk in GiB.
	// Default to the size of the image's `minimum_provisioned_size`.
	// +optional
	SizeGiB int64 `json:"sizeGiB,omitempty"`

	// Profile is the volume profile for the bootdisk, refer https://cloud.ibm.com/docs/vpc?topic=vpc-block-storage-profiles
	// for more information.
	// Default to general-purpose
	// +kubebuilder:validation:Enum="general-purpose";"5iops-tier";"10iops-tier";"custom"
	// +kubebuilder:default=general-purpose
	// +optional
	Profile string `json:"profile,omitempty"`

	// Iops is the maximum I/O operations per second (IOPS) to use for the volume. Applicable only to volumes using a profile
	// family of `custom`.
	// +optional
	Iops int64 `json:"iops,omitempty"`

	// EncryptionKey is the root key to use to wrap the data encryption key for the volume and this points to the CRN
	// and possible values are as follows.
	// The CRN of the [Key Protect Root
	// Key](https://cloud.ibm.com/docs/key-protect?topic=key-protect-getting-started-tutorial) or [Hyper Protect Crypto
	// Service Root Key](https://cloud.ibm.com/docs/hs-crypto?topic=hs-crypto-get-started) for this resource.
	// If unspecified, the `encryption` type for the volume will be `provider_managed`.
	// +optional
	EncryptionKeyCRN string `json:"encryptionKeyCRN,omitempty"`
}

// IBMVPCMachineStatus defines the observed state of IBMVPCMachine.
type IBMVPCMachineStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	InstanceID string `json:"instanceID,omitempty"`

	// Ready is true when the provider resource is ready.
	// +optional
	Ready bool `json:"ready"`

	// Addresses contains the GCP instance associated addresses.
	Addresses []corev1.NodeAddress `json:"addresses,omitempty"`

	// InstanceStatus is the status of the GCP instance for this machine.
	// +optional
	InstanceStatus string `json:"instanceState,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=ibmvpcmachines,scope=Namespaced,categories=cluster-api
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.ready",description="Cluster infrastructure is ready for IBM VPC instances"

// IBMVPCMachine is the Schema for the ibmvpcmachines API.
type IBMVPCMachine struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IBMVPCMachineSpec   `json:"spec,omitempty"`
	Status IBMVPCMachineStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// IBMVPCMachineList contains a list of IBMVPCMachine.
type IBMVPCMachineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IBMVPCMachine `json:"items"`
}

func init() {
	SchemeBuilder.Register(&IBMVPCMachine{}, &IBMVPCMachineList{})
}
