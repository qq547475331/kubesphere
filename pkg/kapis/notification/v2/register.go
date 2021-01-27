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

package v2

import (
	"github.com/emicklei/go-restful"
	openapi "github.com/emicklei/go-restful-openapi"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes"
	"kubesphere.io/kubesphere/pkg/api"
	"kubesphere.io/kubesphere/pkg/apiserver/query"
	"kubesphere.io/kubesphere/pkg/apiserver/runtime"
	kubesphere "kubesphere.io/kubesphere/pkg/client/clientset/versioned"
	"kubesphere.io/kubesphere/pkg/constants"
	"kubesphere.io/kubesphere/pkg/informers"
	"kubesphere.io/kubesphere/pkg/server/errors"
	"net/http"
)

const (
	GroupName      = "notification.kubesphere.io"
	KeyOpenAPITags = openapi.KeyOpenAPITags
)

var GroupVersion = schema.GroupVersion{Group: GroupName, Version: "v2"}

func AddToContainer(
	container *restful.Container,
	informers informers.InformerFactory,
	k8sClient kubernetes.Interface,
	ksClient kubesphere.Interface) error {

	ws := runtime.NewWebService(GroupVersion)
	h := newNotificationHandler(informers, k8sClient, ksClient)

	// apis for secrets
	ws.Route(ws.GET("/secrets").
		To(h.ListSecret).
		Doc("list the secrets").
		Metadata(KeyOpenAPITags, []string{constants.NotificationTag}).
		Param(ws.QueryParameter(query.ParameterName, "name used to do filtering").Required(false)).
		Param(ws.QueryParameter(query.ParameterLabelSelector, "label selector used to do filtering").Required(false)).
		Param(ws.QueryParameter(query.ParameterPage, "page").Required(false).DataFormat("page=%d").DefaultValue("page=1")).
		Param(ws.QueryParameter(query.ParameterLimit, "limit").Required(false)).
		Param(ws.QueryParameter(query.ParameterAscending, "sort parameters, e.g. ascending=false").Required(false).DefaultValue("ascending=false")).
		Param(ws.QueryParameter(query.ParameterOrderBy, "sort parameters, e.g. orderBy=createTime")).
		Returns(http.StatusOK, api.StatusOK, api.ListResult{Items: []interface{}{}}))

	ws.Route(ws.GET("/secrets/{secret}").
		To(h.GetSecret).
		Doc("get the secret").
		Metadata(KeyOpenAPITags, []string{constants.NotificationTag}).
		Param(ws.PathParameter("secret", "secret name")).
		Returns(http.StatusOK, api.StatusOK, []v1.Secret{}))

	ws.Route(ws.POST("/secrets").
		To(h.CreateOrUpdateSecret).
		Doc("create a secret").
		Metadata(KeyOpenAPITags, []string{constants.NotificationTag}).
		Returns(http.StatusOK, api.StatusOK, []v1.Secret{}))

	ws.Route(ws.PUT("/secrets/{secret}").
		To(h.CreateOrUpdateSecret).
		Doc("update the secret").
		Metadata(KeyOpenAPITags, []string{constants.NotificationTag}).
		Param(ws.PathParameter("secret", "secret name")).
		Returns(http.StatusOK, api.StatusOK, []v1.Secret{}))

	ws.Route(ws.DELETE("/secrets/{secret}").
		To(h.DeleteSecret).
		Doc("delete the secret").
		Metadata(KeyOpenAPITags, []string{constants.NotificationTag}).
		Param(ws.PathParameter("secret", "secret name")).
		Returns(http.StatusOK, api.StatusOK, errors.None))

	// apis for global notification config and receiver
	ws.Route(ws.GET("/{resources}").
		To(h.ListResource).
		Doc("list the notification config or receiver").
		Metadata(KeyOpenAPITags, []string{constants.NotificationTag}).
		Param(ws.PathParameter("resources", "notification config or receiver, known values are dingtalkconfigs, dingtalkreceivers, emailconfigs. emailreceivers, slackconfigs, slackreceivers, webhookconfigs, webhookreceivers, wechatconfigs, wechatreceivers")).
		Param(ws.QueryParameter(query.ParameterName, "name used to do filtering").Required(false)).
		Param(ws.QueryParameter(query.ParameterLabelSelector, "label selector used to do filtering").Required(false)).
		Param(ws.QueryParameter(query.ParameterPage, "page").Required(false).DataFormat("page=%d").DefaultValue("page=1")).
		Param(ws.QueryParameter(query.ParameterLimit, "limit").Required(false)).
		Param(ws.QueryParameter(query.ParameterAscending, "sort parameters, e.g. ascending=false").Required(false).DefaultValue("ascending=false")).
		Param(ws.QueryParameter(query.ParameterOrderBy, "sort parameters, e.g. orderBy=createTime")).
		Returns(http.StatusOK, api.StatusOK, api.ListResult{Items: []interface{}{}}))

	ws.Route(ws.GET("/{resources}/{name}").
		To(h.GetResource).
		Doc("get the specified notification config or receiver").
		Metadata(KeyOpenAPITags, []string{constants.NotificationTag}).
		Param(ws.PathParameter("resources", "notification config or receiver, known values are dingtalkconfigs, dingtalkreceivers, emailconfigs. emailreceivers, slackconfigs, slackreceivers, webhookconfigs, webhookreceivers, wechatconfigs, wechatreceivers")).
		Param(ws.PathParameter("name", "the name of the resource")).
		Returns(http.StatusOK, api.StatusOK, nil))

	ws.Route(ws.POST("/{resources}").
		To(h.CreateResource).
		Doc("create a notification config or receiver").
		Metadata(KeyOpenAPITags, []string{constants.NotificationTag}).
		Param(ws.PathParameter("resource", "notification config or receiver, known values are dingtalkconfigs, dingtalkreceivers, emailconfigs. emailreceivers, slackconfigs, slackreceivers, webhookconfigs, webhookreceivers, wechatconfigs, wechatreceivers")).
		Returns(http.StatusOK, api.StatusOK, nil))

	ws.Route(ws.PUT("/{resources}/{name}").
		To(h.UpdateResource).
		Doc("update the specified notification config or receiver").
		Metadata(KeyOpenAPITags, []string{constants.NotificationTag}).
		Param(ws.PathParameter("resources", "notification config or receiver, known values are dingtalkconfigs, dingtalkreceivers, emailconfigs. emailreceivers, slackconfigs, slackreceivers, webhookconfigs, webhookreceivers, wechatconfigs, wechatreceivers")).
		Param(ws.PathParameter("name", "the name of the resource")).
		Returns(http.StatusOK, api.StatusOK, nil))

	ws.Route(ws.DELETE("/{resources}/{name}").
		To(h.DeleteResource).
		Doc("delete the specified notification config or receiver").
		Metadata(KeyOpenAPITags, []string{constants.NotificationTag}).
		Param(ws.PathParameter("resources", "notification config or receiver, known values are dingtalkconfigs, dingtalkreceivers, emailconfigs. emailreceivers, slackconfigs, slackreceivers, webhookconfigs, webhookreceivers, wechatconfigs, wechatreceivers")).
		Param(ws.PathParameter("name", "the name of the resource")).
		Returns(http.StatusOK, api.StatusOK, errors.None))

	// apis for tenant notification config and receiver
	ws.Route(ws.GET("/users/{user}/{resources}").
		To(h.ListResource).
		Doc("list the notification config or receiver").
		Metadata(KeyOpenAPITags, []string{constants.NotificationTag}).
		Param(ws.PathParameter("user", "user name")).
		Param(ws.QueryParameter(query.ParameterName, "name used to do filtering").Required(false)).
		Param(ws.QueryParameter(query.ParameterLabelSelector, "label selector used to do filtering").Required(false)).
		Param(ws.QueryParameter(query.ParameterPage, "page").Required(false).DataFormat("page=%d").DefaultValue("page=1")).
		Param(ws.QueryParameter(query.ParameterLimit, "limit").Required(false)).
		Param(ws.QueryParameter(query.ParameterAscending, "sort parameters, e.g. ascending=false").Required(false).DefaultValue("ascending=false")).
		Param(ws.QueryParameter(query.ParameterOrderBy, "sort parameters, e.g. orderBy=createTime")).
		Returns(http.StatusOK, api.StatusOK, api.ListResult{Items: []interface{}{}}))

	ws.Route(ws.GET("/users/{user}/{resources}/{name}").
		To(h.GetResource).
		Doc("get the specified notification config or receiver").
		Metadata(KeyOpenAPITags, []string{constants.NotificationTag}).
		Param(ws.PathParameter("user", "user name")).
		Param(ws.PathParameter("resources", "notification config or receiver, known values are dingtalkconfigs, dingtalkreceivers, emailconfigs. emailreceivers, slackconfigs, slackreceivers, webhookconfigs, webhookreceivers, wechatconfigs, wechatreceivers")).
		Param(ws.PathParameter("name", "the name of the resource")).
		Returns(http.StatusOK, api.StatusOK, nil))

	ws.Route(ws.POST("/users/{user}/{resources}").
		To(h.CreateResource).
		Doc("create the specified notification config or receiver").
		Metadata(KeyOpenAPITags, []string{constants.NotificationTag}).
		Param(ws.PathParameter("user", "user name")).
		Returns(http.StatusOK, api.StatusOK, nil))

	ws.Route(ws.PUT("/users/{user}/{resources}/{name}").
		To(h.UpdateResource).
		Doc("update the specified notification config or receiver").
		Metadata(KeyOpenAPITags, []string{constants.NotificationTag}).
		Param(ws.PathParameter("user", "user name")).
		Param(ws.PathParameter("resources", "notification config or receiver, known values are dingtalkconfigs, dingtalkreceivers, emailconfigs. emailreceivers, slackconfigs, slackreceivers, webhookconfigs, webhookreceivers, wechatconfigs, wechatreceivers")).
		Param(ws.PathParameter("name", "the name of the resource")).
		Returns(http.StatusOK, api.StatusOK, nil))

	ws.Route(ws.DELETE("/users/{user}/{resources}/{name}").
		To(h.DeleteResource).
		Doc("delete the specified notification config or receiver").
		Metadata(KeyOpenAPITags, []string{constants.NotificationTag}).
		Param(ws.PathParameter("user", "user name")).
		Param(ws.PathParameter("resources", "notification config or receiver, known values are dingtalkconfigs, dingtalkreceivers, emailconfigs. emailreceivers, slackconfigs, slackreceivers, webhookconfigs, webhookreceivers, wechatconfigs, wechatreceivers")).
		Param(ws.PathParameter("name", "the name of the resource")).
		Returns(http.StatusOK, api.StatusOK, errors.None))

	container.Add(ws)
	return nil
}
