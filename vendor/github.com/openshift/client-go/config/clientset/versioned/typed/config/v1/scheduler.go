// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/openshift/api/config/v1"
	scheme "github.com/openshift/client-go/config/clientset/versioned/scheme"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SchedulersGetter has a method to return a SchedulerInterface.
// A group's client should implement this interface.
type SchedulersGetter interface {
	Schedulers() SchedulerInterface
}

// SchedulerInterface has methods to work with Scheduler resources.
type SchedulerInterface interface {
	Create(*v1.Scheduler) (*v1.Scheduler, error)
	Update(*v1.Scheduler) (*v1.Scheduler, error)
	UpdateStatus(*v1.Scheduler) (*v1.Scheduler, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.Scheduler, error)
	List(opts meta_v1.ListOptions) (*v1.SchedulerList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Scheduler, err error)
	SchedulerExpansion
}

// schedulers implements SchedulerInterface
type schedulers struct {
	client rest.Interface
}

// newSchedulers returns a Schedulers
func newSchedulers(c *ConfigV1Client) *schedulers {
	return &schedulers{
		client: c.RESTClient(),
	}
}

// Get takes name of the scheduler, and returns the corresponding scheduler object, and an error if there is any.
func (c *schedulers) Get(name string, options meta_v1.GetOptions) (result *v1.Scheduler, err error) {
	result = &v1.Scheduler{}
	err = c.client.Get().
		Resource("schedulers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Schedulers that match those selectors.
func (c *schedulers) List(opts meta_v1.ListOptions) (result *v1.SchedulerList, err error) {
	result = &v1.SchedulerList{}
	err = c.client.Get().
		Resource("schedulers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested schedulers.
func (c *schedulers) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Resource("schedulers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a scheduler and creates it.  Returns the server's representation of the scheduler, and an error, if there is any.
func (c *schedulers) Create(scheduler *v1.Scheduler) (result *v1.Scheduler, err error) {
	result = &v1.Scheduler{}
	err = c.client.Post().
		Resource("schedulers").
		Body(scheduler).
		Do().
		Into(result)
	return
}

// Update takes the representation of a scheduler and updates it. Returns the server's representation of the scheduler, and an error, if there is any.
func (c *schedulers) Update(scheduler *v1.Scheduler) (result *v1.Scheduler, err error) {
	result = &v1.Scheduler{}
	err = c.client.Put().
		Resource("schedulers").
		Name(scheduler.Name).
		Body(scheduler).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *schedulers) UpdateStatus(scheduler *v1.Scheduler) (result *v1.Scheduler, err error) {
	result = &v1.Scheduler{}
	err = c.client.Put().
		Resource("schedulers").
		Name(scheduler.Name).
		SubResource("status").
		Body(scheduler).
		Do().
		Into(result)
	return
}

// Delete takes name of the scheduler and deletes it. Returns an error if one occurs.
func (c *schedulers) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("schedulers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *schedulers) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Resource("schedulers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched scheduler.
func (c *schedulers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Scheduler, err error) {
	result = &v1.Scheduler{}
	err = c.client.Patch(pt).
		Resource("schedulers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}