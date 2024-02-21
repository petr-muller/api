package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// UpdateStatus is the API for aggregated insights about in-progress updates (oc adm update status).
// The Update Status Controller keeps this object populated by querying UpdateInformers.
//
// Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).
// +openshift:compatibility-gen:level=1
type UpdateStatus struct {
	metav1.TypeMeta `json:",inline"`

	// metadata is the standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +kubebuilder:validation:Required
	// +required
	Spec UpdateStatusSpec `json:"spec"`
	// +optional
	Status UpdateStatusStatus `json:"status"`
}

// UpdateStatusSpec is not used now, but it can be used to configure the Update Status Controller in the future
// +k8s:deepcopy-gen=true
type UpdateStatusSpec struct {
}

// UpdateStatusStatus contains an aggregated view of the UpdateInsights collected by the Update Status Controller.`
// +k8s:deepcopy-gen=true
type UpdateStatusStatus struct {
}
