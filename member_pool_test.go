// Copyright 2018 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package cluster

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMemberPool(t *testing.T) {
	mp := NewMemberPool()

	assert.Equal(t, 0, mp.Size())

	mp.Add(NewMember())

	assert.Equal(t, 1, mp.Size())

	m1 := NewMember()
	mp.Add(m1)

	mp.Add(m1)

	assert.Equal(t, 2, mp.Size())

	ok := mp.Has(m1.ID)

	assert.True(t, ok)

	m2, err := mp.Get(m1.ID)
	assert.NoError(t, err)

	assert.Equal(t, m1, m2)

	mp.Del(m1.ID)

	assert.Equal(t, 1, mp.Size())

	_, err = mp.Get(m1.ID)

	assert.Error(t, err)

	members := mp.Members()

	assert.Equal(t, mp.Size(), len(members))
}
