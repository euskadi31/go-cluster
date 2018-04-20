// Copyright 2018 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package cluster

import (
	"errors"
	"sync"
)

var ErrMemberNotFound = errors.New("Member not found")

// MemberPool struct
type MemberPool struct {
	members    map[string]*Member
	membersMtx *sync.RWMutex
}

// NewMemberPool constructor
func NewMemberPool() *MemberPool {
	return &MemberPool{
		members:    make(map[string]*Member),
		membersMtx: &sync.RWMutex{},
	}
}

// Add Member to MemberPool
func (p *MemberPool) Add(member *Member) {
	p.membersMtx.Lock()
	defer p.membersMtx.Unlock()

	p.members[member.ID] = member
}

// Has Member exists by ID
func (p *MemberPool) Has(id string) bool {
	p.membersMtx.RLock()
	defer p.membersMtx.RUnlock()

	_, ok := p.members[id]

	return ok
}

// Get Member by ID
func (p *MemberPool) Get(id string) (*Member, error) {
	p.membersMtx.RLock()
	defer p.membersMtx.RUnlock()

	if member, ok := p.members[id]; ok {
		return member, nil
	}

	return nil, ErrMemberNotFound
}

// Del Member by ID
func (p *MemberPool) Del(id string) {
	p.membersMtx.Lock()
	defer p.membersMtx.Unlock()

	delete(p.members, id)
}

// Size of members
func (p *MemberPool) Size() int {
	return len(p.members)
}

// Members return
func (p *MemberPool) Members() []*Member {
	p.membersMtx.RLock()
	defer p.membersMtx.RUnlock()

	size := len(p.members)

	members := make([]*Member, size)

	i := 0
	for _, member := range p.members {
		members[i] = member
		i++
	}

	return members
}
