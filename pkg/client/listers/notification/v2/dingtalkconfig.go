/*
Copyright 2020 The KubeSphere Authors.

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

package v2

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v2 "kubesphere.io/kubesphere/pkg/apis/notification/v2"
)

// DingTalkConfigLister helps list DingTalkConfigs.
type DingTalkConfigLister interface {
	// List lists all DingTalkConfigs in the indexer.
	List(selector labels.Selector) (ret []*v2.DingTalkConfig, err error)
	// Get retrieves the DingTalkConfig from the index for a given name.
	Get(name string) (*v2.DingTalkConfig, error)
	DingTalkConfigListerExpansion
}

// dingTalkConfigLister implements the DingTalkConfigLister interface.
type dingTalkConfigLister struct {
	indexer cache.Indexer
}

// NewDingTalkConfigLister returns a new DingTalkConfigLister.
func NewDingTalkConfigLister(indexer cache.Indexer) DingTalkConfigLister {
	return &dingTalkConfigLister{indexer: indexer}
}

// List lists all DingTalkConfigs in the indexer.
func (s *dingTalkConfigLister) List(selector labels.Selector) (ret []*v2.DingTalkConfig, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v2.DingTalkConfig))
	})
	return ret, err
}

// Get retrieves the DingTalkConfig from the index for a given name.
func (s *dingTalkConfigLister) Get(name string) (*v2.DingTalkConfig, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v2.Resource("dingtalkconfig"), name)
	}
	return obj.(*v2.DingTalkConfig), nil
}
