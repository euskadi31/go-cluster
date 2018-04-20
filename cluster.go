// Copyright 2018 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package cluster

// Cluster struct
type Cluster struct {
	config  *Configuration
	members *MemberPool
}

// New Cluster
func New(config *Configuration) *Cluster {
	return &Cluster{
		config:  config,
		members: NewMemberPool(),
	}
}

// Members return cluster members
func (c *Cluster) Members() []*Member {
	return c.members.Members()
}

// Run cluster
func (c *Cluster) Run() error {

	return nil
}
