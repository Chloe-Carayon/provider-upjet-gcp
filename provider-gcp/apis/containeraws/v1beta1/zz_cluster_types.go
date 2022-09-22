/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AdminUsersObservation struct {
}

type AdminUsersParameters struct {

	// The name of the user, e.g. my-gcp-id@gmail.com.
	// +kubebuilder:validation:Required
	Username *string `json:"username" tf:"username,omitempty"`
}

type AuthorizationObservation struct {
}

type AuthorizationParameters struct {

	// Users to perform operations as a cluster admin. A managed ClusterRoleBinding will be created to grant the cluster-admin ClusterRole to the users. Up to ten admin users can be provided. For more info on RBAC, see https://kubernetes.io/docs/reference/access-authn-authz/rbac/#user-facing-roles
	// +kubebuilder:validation:Required
	AdminUsers []AdminUsersParameters `json:"adminUsers" tf:"admin_users,omitempty"`
}

type AwsServicesAuthenticationObservation struct {
}

type AwsServicesAuthenticationParameters struct {

	// The Amazon Resource Name (ARN) of the role that the Anthos Multi-Cloud API will assume when managing AWS resources on your account.
	// +kubebuilder:validation:Required
	RoleArn *string `json:"roleArn" tf:"role_arn,omitempty"`

	// Optional. An identifier for the assumed role session. When unspecified, it defaults to multicloud-service-agent.
	// +kubebuilder:validation:Optional
	RoleSessionName *string `json:"roleSessionName,omitempty" tf:"role_session_name,omitempty"`
}

type ClusterObservation struct {

	// Output only. The time at which this cluster was created.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// Output only. The endpoint of the cluster's API server.
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// Allows clients to perform consistent read-modify-writes through optimistic concurrency control. May be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	// Fleet configuration.
	// +kubebuilder:validation:Required
	Fleet []FleetObservation `json:"fleet,omitempty" tf:"fleet,omitempty"`

	// an identifier for the resource with format projects/{{project}}/locations/{{location}}/awsClusters/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Output only. If set, there are currently changes in flight to the cluster.
	Reconciling *bool `json:"reconciling,omitempty" tf:"reconciling,omitempty"`

	// Output only. The current state of the cluster. Possible values: STATE_UNSPECIFIED, PROVISIONING, RUNNING, RECONCILING, STOPPING, ERROR, DEGRADED
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// Output only. A globally unique identifier for the cluster.
	UID *string `json:"uid,omitempty" tf:"uid,omitempty"`

	// Output only. The time at which this cluster was last updated.
	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`

	// Output only. Workload Identity settings.
	WorkloadIdentityConfig []WorkloadIdentityConfigObservation `json:"workloadIdentityConfig,omitempty" tf:"workload_identity_config,omitempty"`
}

type ClusterParameters struct {

	// Optional. Annotations on the cluster. This field has the same restrictions as Kubernetes annotations. The total size of all keys and values combined is limited to 256k. Key can have 2 segments: prefix  and name , separated by a slash (/). Prefix must be a DNS subdomain. Name must be 63 characters or less, begin and end with alphanumerics, with dashes (-), underscores (_), dots (.), and alphanumerics between.
	// +kubebuilder:validation:Optional
	Annotations map[string]*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// Configuration related to the cluster RBAC settings.
	// +kubebuilder:validation:Required
	Authorization []AuthorizationParameters `json:"authorization" tf:"authorization,omitempty"`

	// The AWS region where the cluster runs. Each Google Cloud region supports a subset of nearby AWS regions. You can call to list all supported AWS regions within a given Google Cloud region.
	// +kubebuilder:validation:Required
	AwsRegion *string `json:"awsRegion" tf:"aws_region,omitempty"`

	// Configuration related to the cluster control plane.
	// +kubebuilder:validation:Required
	ControlPlane []ControlPlaneParameters `json:"controlPlane" tf:"control_plane,omitempty"`

	// Optional. A human readable description of this cluster. Cannot be longer than 255 UTF-8 encoded bytes.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Fleet configuration.
	// +kubebuilder:validation:Required
	Fleet []FleetParameters `json:"fleet" tf:"fleet,omitempty"`

	// The location for the resource
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// Cluster-wide networking configuration.
	// +kubebuilder:validation:Required
	Networking []NetworkingParameters `json:"networking" tf:"networking,omitempty"`

	// The project for the resource
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type ConfigEncryptionObservation struct {
}

type ConfigEncryptionParameters struct {

	// The ARN of the AWS KMS key used to encrypt cluster configuration.
	// +kubebuilder:validation:Required
	KMSKeyArn *string `json:"kmsKeyArn" tf:"kms_key_arn,omitempty"`
}

type ControlPlaneObservation struct {
}

type ControlPlaneParameters struct {

	// Authentication configuration for management of AWS resources.
	// +kubebuilder:validation:Required
	AwsServicesAuthentication []AwsServicesAuthenticationParameters `json:"awsServicesAuthentication" tf:"aws_services_authentication,omitempty"`

	// The ARN of the AWS KMS key used to encrypt cluster configuration.
	// +kubebuilder:validation:Required
	ConfigEncryption []ConfigEncryptionParameters `json:"configEncryption" tf:"config_encryption,omitempty"`

	// The ARN of the AWS KMS key used to encrypt cluster secrets.
	// +kubebuilder:validation:Required
	DatabaseEncryption []DatabaseEncryptionParameters `json:"databaseEncryption" tf:"database_encryption,omitempty"`

	// The name of the AWS IAM instance pofile to assign to each control plane replica.
	// +kubebuilder:validation:Required
	IAMInstanceProfile *string `json:"iamInstanceProfile" tf:"iam_instance_profile,omitempty"`

	// Optional. The AWS instance type. When unspecified, it defaults to m5.large.
	// +kubebuilder:validation:Optional
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type,omitempty"`

	// Optional. Configuration related to the main volume provisioned for each control plane replica. The main volume is in charge of storing all of the cluster's etcd state. Volumes will be provisioned in the availability zone associated with the corresponding subnet. When unspecified, it defaults to 8 GiB with the GP2 volume type.
	// +kubebuilder:validation:Optional
	MainVolume []MainVolumeParameters `json:"mainVolume,omitempty" tf:"main_volume,omitempty"`

	// Proxy configuration for outbound HTTP(S) traffic.
	// +kubebuilder:validation:Optional
	ProxyConfig []ProxyConfigParameters `json:"proxyConfig,omitempty" tf:"proxy_config,omitempty"`

	// Optional. Configuration related to the root volume provisioned for each control plane replica. Volumes will be provisioned in the availability zone associated with the corresponding subnet. When unspecified, it defaults to 32 GiB with the GP2 volume type.
	// +kubebuilder:validation:Optional
	RootVolume []RootVolumeParameters `json:"rootVolume,omitempty" tf:"root_volume,omitempty"`

	// Optional. SSH configuration for how to access the underlying control plane machines.
	// +kubebuilder:validation:Optional
	SSHConfig []SSHConfigParameters `json:"sshConfig,omitempty" tf:"ssh_config,omitempty"`

	// Optional. The IDs of additional security groups to add to control plane replicas. The Anthos Multi-Cloud API will automatically create and manage security groups with the minimum rules needed for a functioning cluster.
	// +kubebuilder:validation:Optional
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// The list of subnets where control plane replicas will run. A replica will be provisioned on each subnet and up to three values can be provided. Each subnet must be in a different AWS Availability Zone (AZ).
	// +kubebuilder:validation:Required
	SubnetIds []*string `json:"subnetIds" tf:"subnet_ids,omitempty"`

	// Optional. A set of AWS resource tags to propagate to all underlying managed AWS resources. Specify at most 50 pairs containing alphanumerics, spaces, and symbols (.+-=_:@/). Keys can be up to 127 Unicode characters. Values can be up to 255 Unicode characters.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The Kubernetes version to run on control plane replicas (e.g. 1.19.10-gke.1000). You can list all supported versions on a given Google Cloud region by calling .
	// +kubebuilder:validation:Required
	Version *string `json:"version" tf:"version,omitempty"`
}

type DatabaseEncryptionObservation struct {
}

type DatabaseEncryptionParameters struct {

	// The ARN of the AWS KMS key used to encrypt cluster secrets.
	// +kubebuilder:validation:Required
	KMSKeyArn *string `json:"kmsKeyArn" tf:"kms_key_arn,omitempty"`
}

type FleetObservation struct {

	// The name of the managed Hub Membership resource associated to this cluster. Membership names are formatted as projects//locations/global/membership/.
	Membership *string `json:"membership,omitempty" tf:"membership,omitempty"`
}

type FleetParameters struct {

	// The number of the Fleet host project where this cluster will be registered.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type MainVolumeObservation struct {
}

type MainVolumeParameters struct {

	// Optional. The number of I/O operations per second (IOPS) to provision for GP3 volume.
	// +kubebuilder:validation:Optional
	Iops *float64 `json:"iops,omitempty" tf:"iops,omitempty"`

	// Optional. The Amazon Resource Name (ARN) of the Customer Managed Key (CMK) used to encrypt AWS EBS volumes. If not specified, the default Amazon managed key associated to the AWS region where this cluster runs will be used.
	// +kubebuilder:validation:Optional
	KMSKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`

	// Optional. The size of the volume, in GiBs. When unspecified, a default value is provided. See the specific reference in the parent resource.
	// +kubebuilder:validation:Optional
	SizeGib *float64 `json:"sizeGib,omitempty" tf:"size_gib,omitempty"`

	// Optional. Type of the EBS volume. When unspecified, it defaults to GP2 volume. Possible values: VOLUME_TYPE_UNSPECIFIED, GP2, GP3
	// +kubebuilder:validation:Optional
	VolumeType *string `json:"volumeType,omitempty" tf:"volume_type,omitempty"`
}

type NetworkingObservation struct {
}

type NetworkingParameters struct {

	// All pods in the cluster are assigned an RFC1918 IPv4 address from these ranges. Only a single range is supported. This field cannot be changed after creation.
	// +kubebuilder:validation:Required
	PodAddressCidrBlocks []*string `json:"podAddressCidrBlocks" tf:"pod_address_cidr_blocks,omitempty"`

	// All services in the cluster are assigned an RFC1918 IPv4 address from these ranges. Only a single range is supported. This field cannot be changed after creation.
	// +kubebuilder:validation:Required
	ServiceAddressCidrBlocks []*string `json:"serviceAddressCidrBlocks" tf:"service_address_cidr_blocks,omitempty"`

	// The VPC associated with the cluster. All component clusters (i.e. control plane and node pools) run on a single VPC. This field cannot be changed after creation.
	// +kubebuilder:validation:Required
	VPCID *string `json:"vpcId" tf:"vpc_id,omitempty"`
}

type ProxyConfigObservation struct {
}

type ProxyConfigParameters struct {

	// The ARN of the AWS Secret Manager secret that contains the HTTP(S) proxy configuration.
	// +kubebuilder:validation:Required
	SecretArn *string `json:"secretArn" tf:"secret_arn,omitempty"`

	// The version string of the AWS Secret Manager secret that contains the HTTP(S) proxy configuration.
	// +kubebuilder:validation:Required
	SecretVersion *string `json:"secretVersion" tf:"secret_version,omitempty"`
}

type RootVolumeObservation struct {
}

type RootVolumeParameters struct {

	// Optional. The number of I/O operations per second (IOPS) to provision for GP3 volume.
	// +kubebuilder:validation:Optional
	Iops *float64 `json:"iops,omitempty" tf:"iops,omitempty"`

	// Optional. The Amazon Resource Name (ARN) of the Customer Managed Key (CMK) used to encrypt AWS EBS volumes. If not specified, the default Amazon managed key associated to the AWS region where this cluster runs will be used.
	// +kubebuilder:validation:Optional
	KMSKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`

	// Optional. The size of the volume, in GiBs. When unspecified, a default value is provided. See the specific reference in the parent resource.
	// +kubebuilder:validation:Optional
	SizeGib *float64 `json:"sizeGib,omitempty" tf:"size_gib,omitempty"`

	// Optional. Type of the EBS volume. When unspecified, it defaults to GP2 volume. Possible values: VOLUME_TYPE_UNSPECIFIED, GP2, GP3
	// +kubebuilder:validation:Optional
	VolumeType *string `json:"volumeType,omitempty" tf:"volume_type,omitempty"`
}

type SSHConfigObservation struct {
}

type SSHConfigParameters struct {

	// The name of the EC2 key pair used to login into cluster machines.
	// +kubebuilder:validation:Required
	EC2KeyPair *string `json:"ec2KeyPair" tf:"ec2_key_pair,omitempty"`
}

type WorkloadIdentityConfigObservation struct {
	IdentityProvider *string `json:"identityProvider,omitempty" tf:"identity_provider,omitempty"`

	IssuerURI *string `json:"issuerUri,omitempty" tf:"issuer_uri,omitempty"`

	WorkloadPool *string `json:"workloadPool,omitempty" tf:"workload_pool,omitempty"`
}

type WorkloadIdentityConfigParameters struct {
}

// ClusterSpec defines the desired state of Cluster
type ClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClusterParameters `json:"forProvider"`
}

// ClusterStatus defines the observed state of Cluster.
type ClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Cluster is the Schema for the Clusters API. An Anthos cluster running on AWS.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterSpec   `json:"spec"`
	Status            ClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterList contains a list of Clusters
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cluster `json:"items"`
}

// Repository type metadata.
var (
	Cluster_Kind             = "Cluster"
	Cluster_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Cluster_Kind}.String()
	Cluster_KindAPIVersion   = Cluster_Kind + "." + CRDGroupVersion.String()
	Cluster_GroupVersionKind = CRDGroupVersion.WithKind(Cluster_Kind)
)

func init() {
	SchemeBuilder.Register(&Cluster{}, &ClusterList{})
}
