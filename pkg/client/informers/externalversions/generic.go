/*
Copyright 2019 The KubeSphere authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package externalversions

import (
	"fmt"

	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
	v1alpha1 "kubesphere.io/kubesphere/pkg/apis/cluster/v1alpha1"
	devopsv1alpha1 "kubesphere.io/kubesphere/pkg/apis/devops/v1alpha1"
	v1alpha3 "kubesphere.io/kubesphere/pkg/apis/devops/v1alpha3"
	v1alpha2 "kubesphere.io/kubesphere/pkg/apis/iam/v1alpha2"
	networkv1alpha1 "kubesphere.io/kubesphere/pkg/apis/network/v1alpha1"
	servicemeshv1alpha2 "kubesphere.io/kubesphere/pkg/apis/servicemesh/v1alpha2"
	storagev1alpha1 "kubesphere.io/kubesphere/pkg/apis/storage/v1alpha1"
	tenantv1alpha1 "kubesphere.io/kubesphere/pkg/apis/tenant/v1alpha1"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=cluster.kubesphere.io, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithResource("clusters"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Cluster().V1alpha1().Clusters().Informer()}, nil

		// Group=devops.kubesphere.io, Version=v1alpha1
	case devopsv1alpha1.SchemeGroupVersion.WithResource("s2ibinaries"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Devops().V1alpha1().S2iBinaries().Informer()}, nil
	case devopsv1alpha1.SchemeGroupVersion.WithResource("s2ibuilders"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Devops().V1alpha1().S2iBuilders().Informer()}, nil
	case devopsv1alpha1.SchemeGroupVersion.WithResource("s2ibuildertemplates"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Devops().V1alpha1().S2iBuilderTemplates().Informer()}, nil
	case devopsv1alpha1.SchemeGroupVersion.WithResource("s2iruns"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Devops().V1alpha1().S2iRuns().Informer()}, nil

		// Group=devops.kubesphere.io, Version=v1alpha3
	case v1alpha3.SchemeGroupVersion.WithResource("devopsprojects"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Devops().V1alpha3().DevOpsProjects().Informer()}, nil
	case v1alpha3.SchemeGroupVersion.WithResource("pipelines"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Devops().V1alpha3().Pipelines().Informer()}, nil

		// Group=iam.kubesphere.io, Version=v1alpha2
	case v1alpha2.SchemeGroupVersion.WithResource("globalroles"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Iam().V1alpha2().GlobalRoles().Informer()}, nil
	case v1alpha2.SchemeGroupVersion.WithResource("globalrolebindings"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Iam().V1alpha2().GlobalRoleBindings().Informer()}, nil
	case v1alpha2.SchemeGroupVersion.WithResource("users"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Iam().V1alpha2().Users().Informer()}, nil
	case v1alpha2.SchemeGroupVersion.WithResource("workspaceroles"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Iam().V1alpha2().WorkspaceRoles().Informer()}, nil
	case v1alpha2.SchemeGroupVersion.WithResource("workspacerolebindings"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Iam().V1alpha2().WorkspaceRoleBindings().Informer()}, nil

		// Group=network.kubesphere.io, Version=v1alpha1
	case networkv1alpha1.SchemeGroupVersion.WithResource("namespacenetworkpolicies"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Network().V1alpha1().NamespaceNetworkPolicies().Informer()}, nil

		// Group=servicemesh.kubesphere.io, Version=v1alpha2
	case servicemeshv1alpha2.SchemeGroupVersion.WithResource("servicepolicies"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Servicemesh().V1alpha2().ServicePolicies().Informer()}, nil
	case servicemeshv1alpha2.SchemeGroupVersion.WithResource("strategies"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Servicemesh().V1alpha2().Strategies().Informer()}, nil

		// Group=storage.kubesphere.io, Version=v1alpha1
	case storagev1alpha1.SchemeGroupVersion.WithResource("provisionercapabilities"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Storage().V1alpha1().ProvisionerCapabilities().Informer()}, nil
	case storagev1alpha1.SchemeGroupVersion.WithResource("storageclasscapabilities"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Storage().V1alpha1().StorageClassCapabilities().Informer()}, nil

		// Group=tenant.kubesphere.io, Version=v1alpha1
	case tenantv1alpha1.SchemeGroupVersion.WithResource("workspaces"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Tenant().V1alpha1().Workspaces().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
