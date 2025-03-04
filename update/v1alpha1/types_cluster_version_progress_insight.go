package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterVersionProgressInsight provides summary information about an ongoing cluster control plane update in Standalone clusters
//
// Compatibility level 4: No compatibility is provided, the API can change at any point for any reason. These capabilities should not be used by applications needing long term support.
// +openshift:compatibility-gen:level=4
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:path=clusterversionprogressinsights,scope=Cluster
// +openshift:api-approved.openshift.io=https://github.com/openshift/api/pull/2012
// +openshift:file-pattern=cvoRunLevel=0000_00,operatorName=cluster-version-operator,operatorOrdering=02
// +openshift:enable:FeatureGate=UpgradeStatus
// +kubebuilder:metadata:annotations="description=Provides summary information about an ongoing cluster control plane update in Standalone clusters."
// +kubebuilder:metadata:annotations="displayName=ClusterVersionProgressInsights"

// ClusterVersionStatusInsight reports the state of a ClusterVersion resource (which represents a control plane
// update in standalone clusters), during the update.
type ClusterVersionProgressInsight struct {
	metav1.TypeMeta `json:",inline"`

	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// spec is empty for now, ClusterVersionStatusInsight is purely status-reporting API. In the future spec may be used to hold
	// configuration to drive what information is surfaced and how
	// +required
	Spec ClusterVersionProgressInsightSpec `json:"spec"`
	// status exposes the health and status of the ongoing cluster update
	// +optional
	Status ClusterVersionProgressInsightStatus `json:"status"`
}

// ClusterVersionProgressInsightSpec is empty for now, ClusterVersionProgressInsightSpec is purely status-reporting API. In the future spec may be used
// to hold configuration to drive what information is surfaced and how
type ClusterVersionProgressInsightSpec struct {
}

type ClusterVersionProgressInsightStatus struct {
	// conditions provides detailed observed conditions about ClusterVersion
	// +listType=map
	// +listMapKey=type
	// +patchStrategy=merge
	// +patchMergeKey=type
	// +optional
	// +kubebuilder:validation:MaxItems=10
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`

	// resource is the ClusterVersion resource that represents the control plane
	//
	// Note: By OpenShift API conventions, in isolation this should be a specialized reference that refers just to
	// resource name (because the rest is implied by status insight type). However, because we use resource references in
	// many places and this API is intended to be consumed by clients, not produced, consistency seems to be more valuable
	// than type safety for producers.
	// +required
	Resource ResourceRef `json:"resource"`

	// assessment is the assessment of the control plane update process
	// +required
	Assessment ControlPlaneAssessment `json:"assessment"`

	// versions contains the original and target versions of the upgrade
	// +required
	Versions ControlPlaneUpdateVersions `json:"versions"`

	// completion is a percentage of the update completion (0-100)
	// +required
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=100
	Completion int32 `json:"completion"`

	// startedAt is the time when the update started
	// +required
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Format=date-time
	StartedAt metav1.Time `json:"startedAt"`

	// completedAt is the time when the update completed
	// +optional
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Format=date-time
	CompletedAt *metav1.Time `json:"completedAt,omitempty"`

	// estimatedCompletedAt is the estimated time when the update will complete
	// +optional
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Format=date-time
	EstimatedCompletedAt *metav1.Time `json:"estimatedCompletedAt,omitempty"`
}

// ControlPlaneAssessment is the assessment of the control plane update process
// +kubebuilder:validation:Enum=Unknown;Progressing;Completed;Degraded
type ControlPlaneAssessment string

const (
	// Unknown means the update status and health cannot be determined
	ControlPlaneAssessmentUnknown ControlPlaneAssessment = "Unknown"
	// Progressing means the control plane is updating and no problems or slowness are detected
	ControlPlaneAssessmentProgressing ControlPlaneAssessment = "Progressing"
	// Completed means the control plane successfully completed updating and no problems are detected
	ControlPlaneAssessmentCompleted ControlPlaneAssessment = "Completed"
	// Degraded means the process of updating the control plane suffers from an observed problem
	ControlPlaneAssessmentDegraded ControlPlaneAssessment = "Degraded"
)

// ControlPlaneUpdateVersions contains the original and target versions of the upgrade
type ControlPlaneUpdateVersions struct {
	// previous is the version of the control plane before the update. When the cluster is being installed
	// for the first time, the version will have a placeholder value '<none>' and carry 'Installation' metadata
	// +required
	// +kubebuilder:validation:XValidation:rule="self.version == '<none>' ? (has(self.metadata) && self.metadata.exists(m, m.key == 'Installation')) : !(has(self.metadata) && self.metadata.exists(m, m.key == 'Installation'))",message="previous version must be '<none>' iff marked with Installation metadata"
	Previous Version `json:"previous"`

	// target is the version of the control plane after the update. It may never be '<none>' or have `Installation` metadata
	// +required
	// +kubebuilder:validation:XValidation:rule="self.version != '<none>' && !(has(self.metadata) && self.metadata.exists(m, m.key == 'Installation'))",message="target version must not be '<none>' or have Installation metadata"
	Target Version `json:"target"`
}

// Version describes a version involved in an update, typically on one side of an update edge
type Version struct {
	// version is a semantic version string, or a placeholder '<none>' for the special case where this
	// is a "previous" version in a new installation, in which case the metadata must contain an item
	// with key 'Installation'
	// +required
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:MinLength=5
	// +kubebuilder:validation:MaxLength=64
	Version string `json:"version"`

	// metadata is a list of metadata associated with the version. It is a list of key-value pairs. The value is optional
	// and when not provided, the metadata item has boolean semantics (presence indicates true)
	// +listType=map
	// +listMapKey=key
	// +patchStrategy=merge
	// +patchMergeKey=key
	// +optional
	// +kubebuilder:validation:MaxItems=5
	Metadata []VersionMetadata `json:"metadata,omitempty" patchStrategy:"merge" patchMergeKey:"key"`
}

// VersionMetadata is a key:value item assigned to version involved in the update. Value can be empty, then the metadata
// have boolean semantics (true when present, false when absent)
type VersionMetadata struct {
	// key is the name of this metadata value
	// +required
	Key VersionMetadataKey `json:"key"`

	// value is the value for the metadata
	// +optional
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:MaxLength=32
	Value string `json:"value,omitempty"`
}

// VersionMetadataKey is a key for a metadata value associated with a version
// +kubebuilder:validation:Enum=Installation;Partial;Architecture
type VersionMetadataKey string

const (
	// Installation denotes a boolean that indicates the update was initiated as an installation
	InstallationMetadata VersionMetadataKey = "Installation"
	// Partial denotes a boolean that indicates the update was initiated in a state where the previous upgrade
	// (to the original version) was not fully completed
	PartialMetadata VersionMetadataKey = "Partial"
	// Architecture denotes a string that indicates the architecture of the payload image of the version,
	// when relevant
	ArchitectureMetadata VersionMetadataKey = "Architecture"
)

// ClusterVersionStatusInsightConditionType are types of conditions that can be reported on ClusterVersion status insight
type ClusterVersionStatusInsightConditionType string

const (
	// Updating condition communicates whether the ClusterVersion is updating
	ClusterVersionStatusInsightUpdating ClusterVersionStatusInsightConditionType = "Updating"
)

// ClusterVersionStatusInsightUpdatingReason are well-known reasons for the Updating condition on ClusterVersion status insights
type ClusterVersionStatusInsightUpdatingReason string

const (
	// CannotDetermineUpdating is used with Updating=Unknown
	ClusterVersionCannotDetermineUpdating ClusterVersionStatusInsightUpdatingReason = "CannotDetermineUpdating"
	// ClusterVersionProgressing means that ClusterVersion is considered to be Updating=True because it has a Progressing=True condition
	ClusterVersionProgressing ClusterVersionStatusInsightUpdatingReason = "ClusterVersionProgressing"
	// ClusterVersionNotProgressing means that ClusterVersion is considered to be Updating=False because it has a Progressing=False condition
	ClusterVersionNotProgressing ClusterVersionStatusInsightUpdatingReason = "ClusterVersionNotProgressing"
)

// ResourceRef is a reference to a kubernetes resource, typically involved in an insight
type ResourceRef struct {
	// group of the object being referenced, if any
	// +optional
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:MaxLength=253
	Group string `json:"group,omitempty"`

	// resource of object being referenced
	// +required
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:MaxLength=253
	// +kubebuilder:validation:MinLength=1
	Resource string `json:"resource"`

	// name of the object being referenced
	// +required
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:MaxLength=253
	// +kubebuilder:validation:MinLength=1
	Name string `json:"name"`

	// namespace of the object being referenced, if any
	// +optional
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:MaxLength=253
	// +kubebuilder:validation:Pattern=`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`
	Namespace string `json:"namespace,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterVersionProgressInsightList is a list of ClusterVersionProgressInsightList resources
//
// Compatibility level 4: No compatibility is provided, the API can change at any point for any reason. These capabilities should not be used by applications needing long term support.
// +openshift:compatibility-gen:level=4
type ClusterVersionProgressInsightList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata"`

	// items is a list of ClusterVersionStatusInsight resources
	// +optional
	// +kubebuilder:validation:MaxItems=32
	Items []ClusterVersionProgressInsight `json:"items"`
}
