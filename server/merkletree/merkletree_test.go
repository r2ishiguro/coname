// Copyright 2015 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package merkletree

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"testing"

	"github.com/syndtr/goleveldb/leveldb"

	"github.com/yahoo/coname/server/kv"
	"github.com/yahoo/coname/server/kv/leveldbkv"
)

func withDB(f func(kv.DB)) {
	dir, err := ioutil.TempDir("", "merkletree")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(dir)
	db, err := leveldb.OpenFile(dir, nil)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	f(leveldbkv.Wrap(db))
}

func TestOneEntry(t *testing.T) {
	withDB(func(db kv.DB) {
		m, err := AccessMerkleTree(db, nil)
		if err != nil {
			panic(err)
		}
		e := m.GetSnapshot(0)
		index := []byte{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
		val1, proofix, proof1, err := e.Lookup(index)
		if val1 != nil || proofix != nil || proof1 != nil || err != nil {
			panic("bad lookup in empty tree")
		}
		val2 := []byte{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2}
		ne, err := e.BeginModification()
		if err != nil {
			panic(err)
		}
		err = ne.Set(index, val2)
		if err != nil {
			panic(err)
		}
		wb := new(leveldb.Batch)
		flushed := ne.Flush(wb)
		err = db.Write(wb)
		if err != nil {
			panic(err)
		}
		e2 := m.GetSnapshot(flushed.Nr)
		v, pix, _, err := e2.Lookup(index)
		if err != nil {
			panic(err)
		}
		if pix != nil {
			v = nil
		}
		if !bytes.Equal(v, val2) {
			panic(fmt.Errorf("Value mismatch: %x / %x", v, val2))
		}
		// TODO: verify proof
	})
}

func TestTwoEntriesOneEpoch(t *testing.T) {
	withDB(func(db kv.DB) {
		m, err := AccessMerkleTree(db, []byte("abcd"))
		if err != nil {
			panic(err)
		}
		e := m.GetSnapshot(0)
		index := []byte{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
		val := []byte{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2}
		ne, err := e.BeginModification()
		if err != nil {
			panic(err)
		}
		_ = "breakpoint"
		err = ne.Set(index, val)
		if err != nil {
			panic(err)
		}
		// Make sure mutating the slices afterwards is okay
		index[15]++
		val[15]++
		err = ne.Set(index, val)
		if err != nil {
			panic(err)
		}
		wb := new(leveldb.Batch)
		flushed := ne.Flush(wb)
		err = db.Write(wb)
		if err != nil {
			panic(err)
		}
		e2 := m.GetSnapshot(flushed.Nr)
		for i := 0; i < 2; i++ {
			v, pix, _, err := e2.Lookup(index)
			if err != nil {
				panic(err)
			}
			if pix != nil {
				v = nil
			}
			if !bytes.Equal(v, val) {
				panic(fmt.Errorf("Value mismatch: %x vs %x", v, val))
			}
			// TODO: verify proof
			index[15]--
			val[15]--
		}
	})
}

func TestThreeEntriesThreeEpochs(t *testing.T) {
	withDB(func(db kv.DB) {
		m, err := AccessMerkleTree(db, []byte("xyz"))
		if err != nil {
			panic(err)
		}
		indices := [][]byte{
			[]byte{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			[]byte{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
			[]byte{0xf0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		}
		values := [][]byte{
			[]byte{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
			[]byte{3, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3},
			[]byte{4, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 4},
		}
		snapshotNrs := []uint64{0}
		for i := 0; i < len(indices); i++ {
			e := m.GetSnapshot(snapshotNrs[i])
			ne, err := e.BeginModification()
			if err != nil {
				panic(err)
			}
			err = ne.Set(indices[i], values[i])
			if err != nil {
				panic(err)
			}
			wb := new(leveldb.Batch)
			flushed := ne.Flush(wb)
			m, err = AccessMerkleTree(db, []byte("xyz"))
			if err != nil {
				panic(err)
			}
			snapshotNrs = append(snapshotNrs, flushed.Nr)
			err = db.Write(wb)
			if err != nil {
				panic(err)
			}
			for j := 0; j <= i+1; j++ {
				e2 := m.GetSnapshot(snapshotNrs[j])
				for k := 0; k < len(indices); k++ {
					v, pix, _, err := e2.Lookup(indices[k])
					if err != nil {
						panic(err)
					}
					if pix != nil {
						v = nil
					}
					if k < j {
						if !bytes.Equal(v, values[k]) {
							panic(fmt.Errorf("Value mismatch: %x / %x", v, values[k]))
						}
						// TODO: verify proof
					} else {
						if v != nil {
							if !bytes.Equal(v, values[k]) {
								panic(fmt.Errorf("bled value mismatch: %d into %d; %x / %x", i, j, v, values[k]))
							} else {
								panic(fmt.Errorf("epoch bled: %d into %d", i, j))
							}
						}
					}
				}
			}
		}
	})
}

type Map interface {
	GetSnapshot(nr uint64) MapSnapshot
}

type MapSnapshot interface {
	GetNr() uint64
	Lookup(indexBytes []byte) (value []byte)
	BeginModification() NewMapSnapshot
}

type NewMapSnapshot interface {
	Set(indexBytes, value []byte)
	Flush() MapSnapshot
}

type TestMerkleTree struct {
	MerkleTree
}

type TestSnapshot Snapshot
type TestNewSnapshot NewSnapshot

var _ Map = (*TestMerkleTree)(nil)
var _ MapSnapshot = (*TestSnapshot)(nil)
var _ NewMapSnapshot = (*TestNewSnapshot)(nil)

func (t *TestMerkleTree) GetSnapshot(nr uint64) MapSnapshot {
	return (*TestSnapshot)(t.MerkleTree.GetSnapshot(nr))
}

func (t *TestSnapshot) GetNr() uint64 {
	return t.Nr
}

func (t *TestSnapshot) Lookup(indexBytes []byte) (value []byte) {
	value, proofIndex, _, err := (*Snapshot)(t).Lookup(indexBytes)
	// TODO check proof
	if err != nil {
		panic(err)
	}
	if proofIndex != nil {
		// actually absence!
		return nil
	}
	return
}

func (t *TestSnapshot) BeginModification() NewMapSnapshot {
	newSnap, err := (*Snapshot)(t).BeginModification()
	if err != nil {
		panic(err)
	}
	return (*TestNewSnapshot)(newSnap)
}

func (t *TestNewSnapshot) Set(indexBytes, value []byte) {
	err := (*NewSnapshot)(t).Set(indexBytes, value)
	if err != nil {
		panic(err)
	}
}

func (t *TestNewSnapshot) Flush() MapSnapshot {
	newSnap := (*NewSnapshot)(t)
	db := newSnap.Snapshot.tree.db
	wb := db.NewBatch()
	flushed := (*TestSnapshot)(newSnap.Flush(wb))
	err := db.Write(wb)
	if err != nil {
		panic(err)
	}
	return flushed
}

type SimpleMap struct {
	snapshots []*SimpleSnapshot
}

type SimpleSnapshot struct {
	nr       uint64
	entries  map[[HashBytes]byte][HashBytes]byte
	wholeMap *SimpleMap
}

type SimpleNewSnapshot struct {
	SimpleSnapshot
}

var _ Map = (*SimpleMap)(nil)
var _ MapSnapshot = (*SimpleSnapshot)(nil)
var _ NewMapSnapshot = (*SimpleNewSnapshot)(nil)

func (m *SimpleMap) GetSnapshot(nr uint64) MapSnapshot {
	if len(m.snapshots) == 0 {
		m.snapshots = append(m.snapshots, &SimpleSnapshot{wholeMap: m, entries: make(map[[HashBytes]byte][HashBytes]byte)})
	}
	return m.snapshots[nr]
}

func (s *SimpleSnapshot) GetNr() uint64 {
	return s.nr
}

func (s *SimpleSnapshot) Lookup(indexBytes []byte) (value []byte) {
	var bytes [HashBytes]byte
	copy(bytes[:], indexBytes)
	if comm, ok := s.entries[bytes]; ok {
		return comm[:]
	} else {
		return nil
	}
}

func (s *SimpleSnapshot) clone() (cloned *SimpleSnapshot) {
	cloned = &SimpleSnapshot{nr: s.nr, entries: make(map[[HashBytes]byte][HashBytes]byte), wholeMap: s.wholeMap}
	for k, v := range s.entries {
		cloned.entries[k] = v
	}
	return
}

func (s *SimpleSnapshot) BeginModification() NewMapSnapshot {
	return &SimpleNewSnapshot{*s.clone()}
}

func (s *SimpleNewSnapshot) Set(indexBytes, value []byte) {
	var bytes, comm [HashBytes]byte
	copy(bytes[:], indexBytes)
	copy(comm[:], value)
	s.entries[bytes] = comm
}

func (s *SimpleNewSnapshot) Flush() MapSnapshot {
	snap := s.SimpleSnapshot.clone()
	snap.nr = uint64(len(s.wholeMap.snapshots))
	s.wholeMap.snapshots = append(s.wholeMap.snapshots, snap)
	return snap
}

type ComparingMap struct {
	Implementations []Map
}

type ComparingSnapshot struct {
	Snapshots []MapSnapshot
}

type NewComparingSnapshot struct {
	NewSnapshots []NewMapSnapshot
}

var _ Map = (*ComparingMap)(nil)
var _ MapSnapshot = (*ComparingSnapshot)(nil)
var _ NewMapSnapshot = (*NewComparingSnapshot)(nil)

func (m *ComparingMap) GetSnapshot(nr uint64) MapSnapshot {
	if nr == 0 {
		snapshots := []MapSnapshot{}
		for _, impl := range m.Implementations {
			snapshots = append(snapshots, impl.GetSnapshot(0))
		}
		return &ComparingSnapshot{snapshots}
	} else {
		panic("not implemented")
	}
}

func (s *ComparingSnapshot) GetNr() uint64 {
	panic("not implemented") // might be a sign this design is somewhat silly
}

func (s *ComparingSnapshot) Lookup(indexBytes []byte) (value []byte) {
	value = s.Snapshots[0].Lookup(indexBytes)
	for _, impl := range s.Snapshots {
		c := impl.Lookup(indexBytes)
		if !bytes.Equal(c, value) {
			log.Panicf("Lookup %x differed: %x vs %x", indexBytes, value, c)
		}
	}
	return
}

func (s *ComparingSnapshot) BeginModification() NewMapSnapshot {
	snapshots := []NewMapSnapshot{}
	for _, impl := range s.Snapshots {
		snapshots = append(snapshots, impl.BeginModification())
	}
	return &NewComparingSnapshot{snapshots}
}

func (s *NewComparingSnapshot) Set(indexBytes, value []byte) {
	for _, snap := range s.NewSnapshots {
		snap.Set(indexBytes, value)
	}
}

func (s *NewComparingSnapshot) Flush() MapSnapshot {
	snapshots := []MapSnapshot{}
	for _, newSnap := range s.NewSnapshots {
		snapshots = append(snapshots, newSnap.Flush())
	}
	return &ComparingSnapshot{snapshots}
}

var dbg = 1

func compareImplementationsRandomly(implementations []Map, itCount, byteRange int, t testing.TB) {
	bytez := func(b byte) [HashBytes]byte {
		var bytes [HashBytes]byte
		for i := range bytes {
			bytes[i] = b<<4 | b
		}
		return bytes
	}
	randBytes := func() []byte {
		var bs [HashBytes]byte
		if byteRange < 0 {
			bs = bytez(byte(rand.Intn(-byteRange)))
			bs[0] = byte(rand.Intn(-byteRange))
			bs[HashBytes-1] = byte(rand.Intn(-byteRange))
			return bs[:]
		} else {
			for i := range bs {
				bs[i] = byte(rand.Intn(byteRange))
			}
			return bs[:]
		}
	}
	comparer := ComparingMap{implementations}
	snapshots := []MapSnapshot{comparer.GetSnapshot(0)}
	changingSnapshots := []NewMapSnapshot{}
	existingKeys := [][]byte{}
	// TODO: sometimes bias towards long snapshot chains, refresh snapshots, refresh the whole thing
	// (and then all snapshots), check only new snapshots most of the time
	for i := 0; i < itCount; i++ {
		switch op := rand.Intn(3); op {
		case 0:
			// new update
			i := rand.Intn(len(snapshots))
			if dbg > 1 {
				log.Printf("begin %x -> *%x", i, len(changingSnapshots))
			}
			changingSnapshots = append(changingSnapshots, snapshots[i].BeginModification())
		case 1:
			// do changes
			if len(changingSnapshots) > 0 {
				i := rand.Intn(len(changingSnapshots))
				nrSets := rand.Intn(10)
				for j := 0; j < nrSets; j++ {
					var k []byte
					if rand.Intn(3) == 0 && len(existingKeys) > 0 {
						// update existing key
						k = existingKeys[rand.Intn(len(existingKeys))]
					} else {
						// create new key
						k = randBytes()
						existingKeys = append(existingKeys, k)
					}
					v := randBytes()
					if dbg > 1 {
						log.Printf("snap *%x: [%x] <- %x", i, k, v)
					}
					changingSnapshots[i].Set(k, v)
				}
			}
		case 2:
			// finish update
			if len(changingSnapshots) > 0 {
				i := rand.Intn(len(changingSnapshots))
				if dbg > 1 {
					log.Printf("flush *%x -> %x", i, len(snapshots))
				}
				snapshots = append(snapshots, changingSnapshots[i].Flush())
				changingSnapshots = append(changingSnapshots[:i], changingSnapshots[i+1:]...)
			}
		}
		if i+1 == itCount || (i < 200 && i%10 == 0) {
			// check all consistency
			keysToQuery := existingKeys
			for i := 0; i < 5; i++ {
				keysToQuery = append(keysToQuery, randBytes())
			}
			for _, k := range keysToQuery {
				if dbg > 2 {
					log.Printf("query %x", k)
				}
				for i := range snapshots {
					if dbg > 2 {
						log.Printf("  in %x:", i)
					}
					v := snapshots[i].Lookup(k)
					if dbg > 2 {
						log.Printf("   %x", v)
					}
				}
			}
		}
	}
}

func TestBrutally(t *testing.T) {
	rand.Seed(4)
	if testing.Verbose() {
		dbg = 2
	}
	withDB(func(db kv.DB) {
		m, err := AccessMerkleTree(db, []byte("tree1"))
		if err != nil {
			panic(err)
		}
		// Test with randomly distributed keys
		impls := []Map{&SimpleMap{}, &TestMerkleTree{*m}}
		compareImplementationsRandomly(impls, 2000, 255, t)
		m2, err := AccessMerkleTree(db, []byte("tree2"))
		if err != nil {
			panic(err)
		}
		// Test with not as random keys
		impls = []Map{&SimpleMap{}, &TestMerkleTree{*m2}}
		compareImplementationsRandomly(impls, 200, -16, t)
	})
}