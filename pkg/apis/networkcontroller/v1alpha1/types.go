/*
Copyright 2017 The Kubernetes Authors.

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

package v1alpha1

import (
	"sort"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Assert is a specification for an Assert resource
type Assert struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AssertSpec   `json:"spec"`
	Status AssertStatus `json:"status"`
}

// AssertSpec is the spec for an Assert resource
type AssertSpec struct {
	DeploymentName string `json:"deploymentName"`
	Replicas       *int32 `json:"replicas"`
}

// AssertStatus is the status for an Assert resource
type AssertStatus struct {
	DNSEntries []*DNSEntry `json:"dnsEntries"`
}

// DNSEntry is the status for an DNS entry resource
type DNSEntry struct {
	APIVersion string   `json:"apiVersion"`
	Kind       string   `json:"kind"`
	Namespace  string   `json:"namespace"`
	Name       string   `json:"name"`
	Hosts      []string `json:"hosts"`
	Error      *string  `json:"error"`
}

// ServiceDNSEntry creates a `DNSEntry` for a `Service`.
func ServiceDNSEntry(namespace, name string, hosts ...string) *DNSEntry {
	sort.Strings(hosts)
	return &DNSEntry{
		APIVersion: "v1",
		Kind:       "Service",
		Namespace:  namespace,
		Name:       name,
		Hosts:      hosts,
		Error:      nil,
	}
}

// ErrorServiceDNSEntry creates a `DNSEntry` for a `Service` which can't be
// reached via DNS.
func ErrorServiceDNSEntry(namespace, name, err string) *DNSEntry {
	return &DNSEntry{
		APIVersion: "v1",
		Kind:       "Service",
		Namespace:  namespace,
		Name:       name,
		Hosts:      nil,
		Error:      &err,
	}
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AssertList is a list of Assert resources
type AssertList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Assert `json:"items"`
}
