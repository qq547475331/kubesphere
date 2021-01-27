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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v2 "kubesphere.io/kubesphere/pkg/apis/notification/v2"
)

// FakeEmailConfigs implements EmailConfigInterface
type FakeEmailConfigs struct {
	Fake *FakeNotificationV2
}

var emailconfigsResource = schema.GroupVersionResource{Group: "notification.kubesphere.io", Version: "v2", Resource: "emailconfigs"}

var emailconfigsKind = schema.GroupVersionKind{Group: "notification.kubesphere.io", Version: "v2", Kind: "EmailConfig"}

// Get takes name of the emailConfig, and returns the corresponding emailConfig object, and an error if there is any.
func (c *FakeEmailConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v2.EmailConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(emailconfigsResource, name), &v2.EmailConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v2.EmailConfig), err
}

// List takes label and field selectors, and returns the list of EmailConfigs that match those selectors.
func (c *FakeEmailConfigs) List(ctx context.Context, opts v1.ListOptions) (result *v2.EmailConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(emailconfigsResource, emailconfigsKind, opts), &v2.EmailConfigList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v2.EmailConfigList{ListMeta: obj.(*v2.EmailConfigList).ListMeta}
	for _, item := range obj.(*v2.EmailConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested emailConfigs.
func (c *FakeEmailConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(emailconfigsResource, opts))
}

// Create takes the representation of a emailConfig and creates it.  Returns the server's representation of the emailConfig, and an error, if there is any.
func (c *FakeEmailConfigs) Create(ctx context.Context, emailConfig *v2.EmailConfig, opts v1.CreateOptions) (result *v2.EmailConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(emailconfigsResource, emailConfig), &v2.EmailConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v2.EmailConfig), err
}

// Update takes the representation of a emailConfig and updates it. Returns the server's representation of the emailConfig, and an error, if there is any.
func (c *FakeEmailConfigs) Update(ctx context.Context, emailConfig *v2.EmailConfig, opts v1.UpdateOptions) (result *v2.EmailConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(emailconfigsResource, emailConfig), &v2.EmailConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v2.EmailConfig), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeEmailConfigs) UpdateStatus(ctx context.Context, emailConfig *v2.EmailConfig, opts v1.UpdateOptions) (*v2.EmailConfig, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(emailconfigsResource, "status", emailConfig), &v2.EmailConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v2.EmailConfig), err
}

// Delete takes name of the emailConfig and deletes it. Returns an error if one occurs.
func (c *FakeEmailConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(emailconfigsResource, name), &v2.EmailConfig{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEmailConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(emailconfigsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v2.EmailConfigList{})
	return err
}

// Patch applies the patch and returns the patched emailConfig.
func (c *FakeEmailConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2.EmailConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(emailconfigsResource, name, pt, data, subresources...), &v2.EmailConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v2.EmailConfig), err
}
