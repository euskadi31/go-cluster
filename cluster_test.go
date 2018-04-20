// Copyright 2018 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package cluster

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCluster(t *testing.T) {
	c := New(&Configuration{})

	assert.Equal(t, 0, len(c.Members()))
}
