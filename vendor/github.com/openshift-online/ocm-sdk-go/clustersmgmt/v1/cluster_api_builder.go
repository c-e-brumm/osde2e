/*
Copyright (c) 2019 Red Hat, Inc.

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

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1

// ClusterAPIBuilder contains the data and logic needed to build 'cluster_API' objects.
//
// Information about the API of a cluster.
type ClusterAPIBuilder struct {
	url *string
}

// NewClusterAPI creates a new builder of 'cluster_API' objects.
func NewClusterAPI() *ClusterAPIBuilder {
	return new(ClusterAPIBuilder)
}

// URL sets the value of the 'URL' attribute
// to the given value.
//
//
func (b *ClusterAPIBuilder) URL(value string) *ClusterAPIBuilder {
	b.url = &value
	return b
}

// Copy copies the attributes of the given object into this builder, discarding any previous values.
func (b *ClusterAPIBuilder) Copy(object *ClusterAPI) *ClusterAPIBuilder {
	if object == nil {
		return b
	}
	b.url = object.url
	return b
}

// Build creates a 'cluster_API' object using the configuration stored in the builder.
func (b *ClusterAPIBuilder) Build() (object *ClusterAPI, err error) {
	object = new(ClusterAPI)
	if b.url != nil {
		object.url = b.url
	}
	return
}
