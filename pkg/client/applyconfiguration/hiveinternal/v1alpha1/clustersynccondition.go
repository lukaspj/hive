// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/openshift/hive/apis/hiveinternal/v1alpha1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ClusterSyncConditionApplyConfiguration represents an declarative configuration of the ClusterSyncCondition type for use
// with apply.
type ClusterSyncConditionApplyConfiguration struct {
	Type               *v1alpha1.ClusterSyncConditionType `json:"type,omitempty"`
	Status             *v1.ConditionStatus                `json:"status,omitempty"`
	LastProbeTime      *metav1.Time                       `json:"lastProbeTime,omitempty"`
	LastTransitionTime *metav1.Time                       `json:"lastTransitionTime,omitempty"`
	Reason             *string                            `json:"reason,omitempty"`
	Message            *string                            `json:"message,omitempty"`
}

// ClusterSyncConditionApplyConfiguration constructs an declarative configuration of the ClusterSyncCondition type for use with
// apply.
func ClusterSyncCondition() *ClusterSyncConditionApplyConfiguration {
	return &ClusterSyncConditionApplyConfiguration{}
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *ClusterSyncConditionApplyConfiguration) WithType(value v1alpha1.ClusterSyncConditionType) *ClusterSyncConditionApplyConfiguration {
	b.Type = &value
	return b
}

// WithStatus sets the Status field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Status field is set to the value of the last call.
func (b *ClusterSyncConditionApplyConfiguration) WithStatus(value v1.ConditionStatus) *ClusterSyncConditionApplyConfiguration {
	b.Status = &value
	return b
}

// WithLastProbeTime sets the LastProbeTime field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LastProbeTime field is set to the value of the last call.
func (b *ClusterSyncConditionApplyConfiguration) WithLastProbeTime(value metav1.Time) *ClusterSyncConditionApplyConfiguration {
	b.LastProbeTime = &value
	return b
}

// WithLastTransitionTime sets the LastTransitionTime field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LastTransitionTime field is set to the value of the last call.
func (b *ClusterSyncConditionApplyConfiguration) WithLastTransitionTime(value metav1.Time) *ClusterSyncConditionApplyConfiguration {
	b.LastTransitionTime = &value
	return b
}

// WithReason sets the Reason field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Reason field is set to the value of the last call.
func (b *ClusterSyncConditionApplyConfiguration) WithReason(value string) *ClusterSyncConditionApplyConfiguration {
	b.Reason = &value
	return b
}

// WithMessage sets the Message field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Message field is set to the value of the last call.
func (b *ClusterSyncConditionApplyConfiguration) WithMessage(value string) *ClusterSyncConditionApplyConfiguration {
	b.Message = &value
	return b
}