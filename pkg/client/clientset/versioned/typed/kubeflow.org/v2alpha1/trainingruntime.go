// Copyright 2024 The Kubeflow Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package v2alpha1

import (
	"context"
	json "encoding/json"
	"fmt"
	"time"

	v2alpha1 "github.com/kubeflow/training-operator/pkg/apis/kubeflow.org/v2alpha1"
	kubefloworgv2alpha1 "github.com/kubeflow/training-operator/pkg/client/applyconfiguration/kubeflow.org/v2alpha1"
	scheme "github.com/kubeflow/training-operator/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// TrainingRuntimesGetter has a method to return a TrainingRuntimeInterface.
// A group's client should implement this interface.
type TrainingRuntimesGetter interface {
	TrainingRuntimes(namespace string) TrainingRuntimeInterface
}

// TrainingRuntimeInterface has methods to work with TrainingRuntime resources.
type TrainingRuntimeInterface interface {
	Create(ctx context.Context, trainingRuntime *v2alpha1.TrainingRuntime, opts v1.CreateOptions) (*v2alpha1.TrainingRuntime, error)
	Update(ctx context.Context, trainingRuntime *v2alpha1.TrainingRuntime, opts v1.UpdateOptions) (*v2alpha1.TrainingRuntime, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v2alpha1.TrainingRuntime, error)
	List(ctx context.Context, opts v1.ListOptions) (*v2alpha1.TrainingRuntimeList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2alpha1.TrainingRuntime, err error)
	Apply(ctx context.Context, trainingRuntime *kubefloworgv2alpha1.TrainingRuntimeApplyConfiguration, opts v1.ApplyOptions) (result *v2alpha1.TrainingRuntime, err error)
	TrainingRuntimeExpansion
}

// trainingRuntimes implements TrainingRuntimeInterface
type trainingRuntimes struct {
	client rest.Interface
	ns     string
}

// newTrainingRuntimes returns a TrainingRuntimes
func newTrainingRuntimes(c *KubeflowV2alpha1Client, namespace string) *trainingRuntimes {
	return &trainingRuntimes{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the trainingRuntime, and returns the corresponding trainingRuntime object, and an error if there is any.
func (c *trainingRuntimes) Get(ctx context.Context, name string, options v1.GetOptions) (result *v2alpha1.TrainingRuntime, err error) {
	result = &v2alpha1.TrainingRuntime{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("trainingruntimes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of TrainingRuntimes that match those selectors.
func (c *trainingRuntimes) List(ctx context.Context, opts v1.ListOptions) (result *v2alpha1.TrainingRuntimeList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v2alpha1.TrainingRuntimeList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("trainingruntimes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested trainingRuntimes.
func (c *trainingRuntimes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("trainingruntimes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a trainingRuntime and creates it.  Returns the server's representation of the trainingRuntime, and an error, if there is any.
func (c *trainingRuntimes) Create(ctx context.Context, trainingRuntime *v2alpha1.TrainingRuntime, opts v1.CreateOptions) (result *v2alpha1.TrainingRuntime, err error) {
	result = &v2alpha1.TrainingRuntime{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("trainingruntimes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(trainingRuntime).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a trainingRuntime and updates it. Returns the server's representation of the trainingRuntime, and an error, if there is any.
func (c *trainingRuntimes) Update(ctx context.Context, trainingRuntime *v2alpha1.TrainingRuntime, opts v1.UpdateOptions) (result *v2alpha1.TrainingRuntime, err error) {
	result = &v2alpha1.TrainingRuntime{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("trainingruntimes").
		Name(trainingRuntime.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(trainingRuntime).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the trainingRuntime and deletes it. Returns an error if one occurs.
func (c *trainingRuntimes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("trainingruntimes").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *trainingRuntimes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("trainingruntimes").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched trainingRuntime.
func (c *trainingRuntimes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2alpha1.TrainingRuntime, err error) {
	result = &v2alpha1.TrainingRuntime{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("trainingruntimes").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied trainingRuntime.
func (c *trainingRuntimes) Apply(ctx context.Context, trainingRuntime *kubefloworgv2alpha1.TrainingRuntimeApplyConfiguration, opts v1.ApplyOptions) (result *v2alpha1.TrainingRuntime, err error) {
	if trainingRuntime == nil {
		return nil, fmt.Errorf("trainingRuntime provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(trainingRuntime)
	if err != nil {
		return nil, err
	}
	name := trainingRuntime.Name
	if name == nil {
		return nil, fmt.Errorf("trainingRuntime.Name must be provided to Apply")
	}
	result = &v2alpha1.TrainingRuntime{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("trainingruntimes").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}