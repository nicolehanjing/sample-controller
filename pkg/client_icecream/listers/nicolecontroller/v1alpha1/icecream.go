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
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha1 "k8s.io/sample-controller/pkg/apis/nicolecontroller/v1alpha1"
)

// IcecreamLister helps list Icecreams.
// All objects returned here must be treated as read-only.
type IcecreamLister interface {
	// List lists all Icecreams in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Icecream, err error)
	// Icecreams returns an object that can list and get Icecreams.
	Icecreams(namespace string) IcecreamNamespaceLister
	IcecreamListerExpansion
}

// icecreamLister implements the IcecreamLister interface.
type icecreamLister struct {
	indexer cache.Indexer
}

// NewIcecreamLister returns a new IcecreamLister.
func NewIcecreamLister(indexer cache.Indexer) IcecreamLister {
	return &icecreamLister{indexer: indexer}
}

// List lists all Icecreams in the indexer.
func (s *icecreamLister) List(selector labels.Selector) (ret []*v1alpha1.Icecream, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Icecream))
	})
	return ret, err
}

// Icecreams returns an object that can list and get Icecreams.
func (s *icecreamLister) Icecreams(namespace string) IcecreamNamespaceLister {
	return icecreamNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// IcecreamNamespaceLister helps list and get Icecreams.
// All objects returned here must be treated as read-only.
type IcecreamNamespaceLister interface {
	// List lists all Icecreams in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Icecream, err error)
	// Get retrieves the Icecream from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Icecream, error)
	IcecreamNamespaceListerExpansion
}

// icecreamNamespaceLister implements the IcecreamNamespaceLister
// interface.
type icecreamNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Icecreams in the indexer for a given namespace.
func (s icecreamNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Icecream, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Icecream))
	})
	return ret, err
}

// Get retrieves the Icecream from the indexer for a given namespace and name.
func (s icecreamNamespaceLister) Get(name string) (*v1alpha1.Icecream, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("icecream"), name)
	}
	return obj.(*v1alpha1.Icecream), nil
}