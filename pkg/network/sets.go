package network

import (
	"fmt"
	"sync"

	networkv1alpha1 "github.com/carbonql/networkcontroller/pkg/apis/networkcontroller/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func makeKey(obj interface{}) string {
	om := obj.(metav1.Object)
	ns := om.GetNamespace()
	n := om.GetName()

	return fmt.Sprintf("%s/%s", ns, n)
}

// serviceSet maps service key -> a set of hosts. Service key is of the form
// `namespace/name`.
type serviceSet struct {
	sync.RWMutex
	hosts map[string]*networkv1alpha1.DNSEntry
}

func (c *Controller) addServiceHosts(obj interface{}) {
	key := makeKey(obj)

	c.serviceHosts.Lock()
	defer c.serviceHosts.Unlock()

	if _, exists := c.serviceHosts.hosts[key]; !exists {
		c.serviceHosts.hosts[key] = &networkv1alpha1.DNSEntry{}
		go c.poll(key)
	}
}

func (c *Controller) deleteServiceHosts(obj interface{}) {
	key := makeKey(obj)

	c.serviceHosts.Lock()
	defer c.serviceHosts.Unlock()

	delete(c.serviceHosts.hosts, key)
}
