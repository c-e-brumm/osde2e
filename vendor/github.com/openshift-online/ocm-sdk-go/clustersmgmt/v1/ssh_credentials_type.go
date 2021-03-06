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

// SSHCredentials represents the values of the 'SSH_credentials' type.
//
// SSH key pair of a cluster.
type SSHCredentials struct {
	privateKey *string
	publicKey  *string
}

// Empty returns true if the object is empty, i.e. no attribute has a value.
func (o *SSHCredentials) Empty() bool {
	return o == nil || (o.privateKey == nil &&
		o.publicKey == nil &&
		true)
}

// PrivateKey returns the value of the 'private_key' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// SSH private key of the cluster.
func (o *SSHCredentials) PrivateKey() string {
	if o != nil && o.privateKey != nil {
		return *o.privateKey
	}
	return ""
}

// GetPrivateKey returns the value of the 'private_key' attribute and
// a flag indicating if the attribute has a value.
//
// SSH private key of the cluster.
func (o *SSHCredentials) GetPrivateKey() (value string, ok bool) {
	ok = o != nil && o.privateKey != nil
	if ok {
		value = *o.privateKey
	}
	return
}

// PublicKey returns the value of the 'public_key' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// SSH public key of the cluster.
func (o *SSHCredentials) PublicKey() string {
	if o != nil && o.publicKey != nil {
		return *o.publicKey
	}
	return ""
}

// GetPublicKey returns the value of the 'public_key' attribute and
// a flag indicating if the attribute has a value.
//
// SSH public key of the cluster.
func (o *SSHCredentials) GetPublicKey() (value string, ok bool) {
	ok = o != nil && o.publicKey != nil
	if ok {
		value = *o.publicKey
	}
	return
}

// SSHCredentialsList is a list of values of the 'SSH_credentials' type.
type SSHCredentialsList struct {
	items []*SSHCredentials
}

// Len returns the length of the list.
func (l *SSHCredentialsList) Len() int {
	if l == nil {
		return 0
	}
	return len(l.items)
}

// Empty returns true if the list is empty.
func (l *SSHCredentialsList) Empty() bool {
	return l == nil || len(l.items) == 0
}

// Get returns the item of the list with the given index. If there is no item with
// that index it returns nil.
func (l *SSHCredentialsList) Get(i int) *SSHCredentials {
	if l == nil || i < 0 || i >= len(l.items) {
		return nil
	}
	return l.items[i]
}

// Slice returns an slice containing the items of the list. The returned slice is a
// copy of the one used internally, so it can be modified without affecting the
// internal representation.
//
// If you don't need to modify the returned slice consider using the Each or Range
// functions, as they don't need to allocate a new slice.
func (l *SSHCredentialsList) Slice() []*SSHCredentials {
	var slice []*SSHCredentials
	if l == nil {
		slice = make([]*SSHCredentials, 0)
	} else {
		slice = make([]*SSHCredentials, len(l.items))
		copy(slice, l.items)
	}
	return slice
}

// Each runs the given function for each item of the list, in order. If the function
// returns false the iteration stops, otherwise it continues till all the elements
// of the list have been processed.
func (l *SSHCredentialsList) Each(f func(item *SSHCredentials) bool) {
	if l == nil {
		return
	}
	for _, item := range l.items {
		if !f(item) {
			break
		}
	}
}

// Range runs the given function for each index and item of the list, in order. If
// the function returns false the iteration stops, otherwise it continues till all
// the elements of the list have been processed.
func (l *SSHCredentialsList) Range(f func(index int, item *SSHCredentials) bool) {
	if l == nil {
		return
	}
	for index, item := range l.items {
		if !f(index, item) {
			break
		}
	}
}
