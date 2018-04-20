// Copyright 2018 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package cluster

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMemberAddr(t *testing.T) {
	m := Member{
		Host: "127.0.0.1",
		Port: 1337,
	}

	assert.Equal(t, "127.0.0.1:1337", m.Addr())
}
