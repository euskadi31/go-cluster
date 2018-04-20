// Copyright 2018 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package cluster

import (
	"net"
	"strconv"

	"github.com/satori/go.uuid"
)

// Member struct
type Member struct {
	Topology
	ID   string
	Host string
	Port int
}

// NewMember constructor
func NewMember() *Member {
	return &Member{
		ID: uuid.NewV4().String(),
	}
}

// Addr string
func (m *Member) Addr() string {
	return net.JoinHostPort(m.Host, strconv.Itoa(m.Port))
}
