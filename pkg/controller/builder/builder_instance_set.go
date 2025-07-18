/*
Copyright (C) 2022-2025 ApeCloud Co., Ltd

This file is part of KubeBlocks project

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package builder

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"

	kbappsv1 "github.com/apecloud/kubeblocks/apis/apps/v1"
	workloads "github.com/apecloud/kubeblocks/apis/workloads/v1"
)

type InstanceSetBuilder struct {
	BaseBuilder[workloads.InstanceSet, *workloads.InstanceSet, InstanceSetBuilder]
}

func NewInstanceSetBuilder(namespace, name string) *InstanceSetBuilder {
	builder := &InstanceSetBuilder{}
	replicas := int32(1)
	builder.init(namespace, name,
		&workloads.InstanceSet{
			Spec: workloads.InstanceSetSpec{
				Replicas: &replicas,
			},
		}, builder)
	return builder
}

func (builder *InstanceSetBuilder) SetReplicas(replicas int32) *InstanceSetBuilder {
	builder.get().Spec.Replicas = &replicas
	return builder
}

func (builder *InstanceSetBuilder) SetMinReadySeconds(minReadySeconds int32) *InstanceSetBuilder {
	builder.get().Spec.MinReadySeconds = minReadySeconds
	return builder
}

func (builder *InstanceSetBuilder) SetSelectorMatchLabel(labels map[string]string) *InstanceSetBuilder {
	selector := builder.get().Spec.Selector
	if selector == nil {
		selector = &metav1.LabelSelector{}
		builder.get().Spec.Selector = selector
	}
	matchLabels := make(map[string]string, len(labels))
	for k, v := range labels {
		matchLabels[k] = v
	}
	builder.get().Spec.Selector.MatchLabels = matchLabels
	return builder
}

func (builder *InstanceSetBuilder) SetRoles(roles []workloads.ReplicaRole) *InstanceSetBuilder {
	builder.get().Spec.Roles = roles
	return builder
}

func (builder *InstanceSetBuilder) SetTemplate(template corev1.PodTemplateSpec) *InstanceSetBuilder {
	builder.get().Spec.Template = template
	return builder
}

func (builder *InstanceSetBuilder) AddVolumeClaimTemplates(templates ...corev1.PersistentVolumeClaim) *InstanceSetBuilder {
	templateList := builder.get().Spec.VolumeClaimTemplates
	templateList = append(templateList, templates...)
	builder.get().Spec.VolumeClaimTemplates = templateList
	return builder
}

func (builder *InstanceSetBuilder) SetVolumeClaimTemplates(templates ...corev1.PersistentVolumeClaim) *InstanceSetBuilder {
	builder.get().Spec.VolumeClaimTemplates = templates
	return builder
}

func (builder *InstanceSetBuilder) SetPVCRetentionPolicy(retentionPolicy *workloads.PersistentVolumeClaimRetentionPolicy) *InstanceSetBuilder {
	builder.get().Spec.PersistentVolumeClaimRetentionPolicy = retentionPolicy
	return builder
}

func (builder *InstanceSetBuilder) SetPodManagementPolicy(policy appsv1.PodManagementPolicyType) *InstanceSetBuilder {
	builder.get().Spec.PodManagementPolicy = policy
	return builder
}

func (builder *InstanceSetBuilder) SetParallelPodManagementConcurrency(parallelPodManagementConcurrency *intstr.IntOrString) *InstanceSetBuilder {
	builder.get().Spec.ParallelPodManagementConcurrency = parallelPodManagementConcurrency
	return builder
}
func (builder *InstanceSetBuilder) SetPodUpdatePolicy(policy workloads.PodUpdatePolicyType) *InstanceSetBuilder {
	builder.get().Spec.PodUpdatePolicy = policy
	return builder
}

func (builder *InstanceSetBuilder) SetInstanceUpdateStrategy(strategy *workloads.InstanceUpdateStrategy) *InstanceSetBuilder {
	builder.get().Spec.InstanceUpdateStrategy = strategy
	return builder
}

func (builder *InstanceSetBuilder) SetMemberUpdateStrategy(strategy *workloads.MemberUpdateStrategy) *InstanceSetBuilder {
	builder.get().Spec.MemberUpdateStrategy = strategy
	return builder
}

func (builder *InstanceSetBuilder) SetLifecycleActions(lifecycleActions *kbappsv1.ComponentLifecycleActions) *InstanceSetBuilder {
	if lifecycleActions != nil && lifecycleActions.Switchover != nil {
		if builder.get().Spec.MembershipReconfiguration == nil {
			builder.get().Spec.MembershipReconfiguration = &workloads.MembershipReconfiguration{}
		}
		builder.get().Spec.MembershipReconfiguration.Switchover = lifecycleActions.Switchover
	}
	return builder
}

func (builder *InstanceSetBuilder) SetTemplateVars(templateVars map[string]any) *InstanceSetBuilder {
	if templateVars != nil {
		builder.get().Spec.TemplateVars = make(map[string]string)
		for k, v := range templateVars {
			builder.get().Spec.TemplateVars[k] = v.(string)
		}
	}
	return builder
}

func (builder *InstanceSetBuilder) SetPaused(paused bool) *InstanceSetBuilder {
	builder.get().Spec.Paused = paused
	return builder
}

func (builder *InstanceSetBuilder) SetInstances(instances []workloads.InstanceTemplate) *InstanceSetBuilder {
	builder.get().Spec.Instances = instances
	return builder
}

func (builder *InstanceSetBuilder) SetFlatInstanceOrdinal(flatInstanceOrdinal bool) *InstanceSetBuilder {
	builder.get().Spec.FlatInstanceOrdinal = flatInstanceOrdinal
	return builder
}

func (builder *InstanceSetBuilder) SetOfflineInstances(offlineInstances []string) *InstanceSetBuilder {
	builder.get().Spec.OfflineInstances = offlineInstances
	return builder
}

func (builder *InstanceSetBuilder) SetDisableDefaultHeadlessService(disable bool) *InstanceSetBuilder {
	builder.get().Spec.DisableDefaultHeadlessService = disable
	return builder
}
