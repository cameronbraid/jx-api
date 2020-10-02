/*
Copyright 2020 The Jenkins X Authors.

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

package v1

import (
	v1 "github.com/jenkins-x/jx-api/v3/pkg/apis/jenkins.io/v1"
	"github.com/jenkins-x/jx-api/v3/pkg/client/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type JenkinsV1Interface interface {
	RESTClient() rest.Interface
	AppsGetter
	BuildPacksGetter
	CommitStatusesGetter
	EnvironmentsGetter
	EnvironmentRoleBindingsGetter
	ExtensionsGetter
	GitServicesGetter
	PipelineActivitiesGetter
	PluginsGetter
	ReleasesGetter
	SourceRepositoriesGetter
	SourceRepositoryGroupsGetter
	TeamsGetter
	UsersGetter
}

// JenkinsV1Client is used to interact with features provided by the jenkins.io group.
type JenkinsV1Client struct {
	restClient rest.Interface
}

func (c *JenkinsV1Client) Apps(namespace string) AppInterface {
	return newApps(c, namespace)
}

func (c *JenkinsV1Client) BuildPacks(namespace string) BuildPackInterface {
	return newBuildPacks(c, namespace)
}

func (c *JenkinsV1Client) CommitStatuses(namespace string) CommitStatusInterface {
	return newCommitStatuses(c, namespace)
}

func (c *JenkinsV1Client) Environments(namespace string) EnvironmentInterface {
	return newEnvironments(c, namespace)
}

func (c *JenkinsV1Client) EnvironmentRoleBindings(namespace string) EnvironmentRoleBindingInterface {
	return newEnvironmentRoleBindings(c, namespace)
}

func (c *JenkinsV1Client) Extensions(namespace string) ExtensionInterface {
	return newExtensions(c, namespace)
}

func (c *JenkinsV1Client) GitServices(namespace string) GitServiceInterface {
	return newGitServices(c, namespace)
}

func (c *JenkinsV1Client) PipelineActivities(namespace string) PipelineActivityInterface {
	return newPipelineActivities(c, namespace)
}

func (c *JenkinsV1Client) Plugins(namespace string) PluginInterface {
	return newPlugins(c, namespace)
}

func (c *JenkinsV1Client) Releases(namespace string) ReleaseInterface {
	return newReleases(c, namespace)
}

func (c *JenkinsV1Client) SourceRepositories(namespace string) SourceRepositoryInterface {
	return newSourceRepositories(c, namespace)
}

func (c *JenkinsV1Client) SourceRepositoryGroups(namespace string) SourceRepositoryGroupInterface {
	return newSourceRepositoryGroups(c, namespace)
}

func (c *JenkinsV1Client) Teams(namespace string) TeamInterface {
	return newTeams(c, namespace)
}

func (c *JenkinsV1Client) Users(namespace string) UserInterface {
	return newUsers(c, namespace)
}

// NewForConfig creates a new JenkinsV1Client for the given config.
func NewForConfig(c *rest.Config) (*JenkinsV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &JenkinsV1Client{client}, nil
}

// NewForConfigOrDie creates a new JenkinsV1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *JenkinsV1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new JenkinsV1Client for the given RESTClient.
func New(c rest.Interface) *JenkinsV1Client {
	return &JenkinsV1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *JenkinsV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
