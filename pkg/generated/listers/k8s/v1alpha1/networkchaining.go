/*
Copyright The Kubernetes Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "ovn4nfv-k8s-plugin/pkg/apis/k8s/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// NetworkChainingLister helps list NetworkChainings.
type NetworkChainingLister interface {
	// List lists all NetworkChainings in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.NetworkChaining, err error)
	// NetworkChainings returns an object that can list and get NetworkChainings.
	NetworkChainings(namespace string) NetworkChainingNamespaceLister
	NetworkChainingListerExpansion
}

// networkChainingLister implements the NetworkChainingLister interface.
type networkChainingLister struct {
	indexer cache.Indexer
}

// NewNetworkChainingLister returns a new NetworkChainingLister.
func NewNetworkChainingLister(indexer cache.Indexer) NetworkChainingLister {
	return &networkChainingLister{indexer: indexer}
}

// List lists all NetworkChainings in the indexer.
func (s *networkChainingLister) List(selector labels.Selector) (ret []*v1alpha1.NetworkChaining, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.NetworkChaining))
	})
	return ret, err
}

// NetworkChainings returns an object that can list and get NetworkChainings.
func (s *networkChainingLister) NetworkChainings(namespace string) NetworkChainingNamespaceLister {
	return networkChainingNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// NetworkChainingNamespaceLister helps list and get NetworkChainings.
type NetworkChainingNamespaceLister interface {
	// List lists all NetworkChainings in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.NetworkChaining, err error)
	// Get retrieves the NetworkChaining from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.NetworkChaining, error)
	NetworkChainingNamespaceListerExpansion
}

// networkChainingNamespaceLister implements the NetworkChainingNamespaceLister
// interface.
type networkChainingNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all NetworkChainings in the indexer for a given namespace.
func (s networkChainingNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.NetworkChaining, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.NetworkChaining))
	})
	return ret, err
}

// Get retrieves the NetworkChaining from the indexer for a given namespace and name.
func (s networkChainingNamespaceLister) Get(name string) (*v1alpha1.NetworkChaining, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("networkchaining"), name)
	}
	return obj.(*v1alpha1.NetworkChaining), nil
}