package automata

import (
	"crypto/sha256"
	stdmath "math"

	"github.com/goropikari/golex/collection"
)

const epsilon = 'ε'
const nonFinStateTokenID TokenID = stdmath.MaxInt

const SupportedChars = "!\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~ \t\n\r"

// const SupportedChars = "abcdefghijklmnopqrstuvwxyz+-*/.0123456789"

// const SupportedChars = "abc"

type TokenID int
type StateID int
type Sha = [sha256.Size]byte
type Nothing struct{}

var nothing = Nothing{}

type State struct {
	id      StateID
	tokenID TokenID
}

func NewState(id StateID) State {
	return State{id: id, tokenID: TokenID(nonFinStateTokenID)}
}

func NewStateWithTokenID(id StateID, tokenID TokenID) State {
	return State{
		id:      id,
		tokenID: tokenID,
	}
}

func (st State) GetID() StateID {
	return st.id
}

func (st State) GetTokenID() TokenID {
	if st.tokenID == nonFinStateTokenID {
		return 0
	}
	return st.tokenID
}

func (st State) GetRawTokenID() TokenID {
	return st.tokenID
}

func (st *State) SetTokenID(id TokenID) {
	st.tokenID = id
}

type StateSet struct {
	bs collection.Bitset
}

func NewStateSet(n int) *StateSet {
	return &StateSet{
		bs: collection.NewBitset(n),
	}
}

func (ss *StateSet) Insert(x StateID) *StateSet {
	ss.bs = ss.bs.Up(int(x))
	return ss
}

func (ss *StateSet) Copy() *StateSet {
	return &StateSet{
		bs: ss.bs.Copy(),
	}
}

func (ss *StateSet) Union(other *StateSet) *StateSet {
	return &StateSet{
		bs: ss.bs.Union(other.bs),
	}
}

func (ss *StateSet) Intersection(other *StateSet) *StateSet {
	return &StateSet{
		bs: ss.bs.Intersection(other.bs),
	}
}

func (ss *StateSet) IsAny() bool {
	return !ss.IsEmpty()
}

func (ss *StateSet) IsEmpty() bool {
	return ss.bs.IsZero()
}

func (ss *StateSet) Contains(x StateID) bool {
	return ss.bs.Contains(int(x))
}

type stateSetIterator struct {
	maxID  StateID
	currID StateID
	ss     *StateSet
}

func newStateSetIterator(ss *StateSet) *stateSetIterator {
	sid := StateID(0)
	maxID := StateID(ss.bs.GetLength() - 1)
	for sid <= maxID {
		if ss.Contains(sid) {
			break
		}
		sid++
	}

	return &stateSetIterator{
		maxID:  maxID,
		currID: sid,
		ss:     ss,
	}
}

func (iter *stateSetIterator) HasNext() bool {
	return iter.currID <= iter.maxID
}

func (iter *stateSetIterator) Next() StateID {
	ret := iter.currID
	iter.currID++
	for iter.currID <= StateID(iter.maxID) {
		if iter.ss.Contains(iter.currID) {
			break
		}
		iter.currID++
	}

	return ret
}

func (ss *StateSet) iterator() *stateSetIterator {
	return newStateSetIterator(ss)
}

type StateSetDict[T any] struct {
	dict    *collection.BitsetDict[T]
	shaToSs map[Sha]*StateSet
}

func NewStateSetDict[T any]() *StateSetDict[T] {
	return &StateSetDict[T]{
		dict:    collection.NewBitsetDict[T](),
		shaToSs: make(map[Sha]*StateSet),
	}
}

func (d *StateSetDict[T]) Set(ss *StateSet, v T) *StateSetDict[T] {
	d.dict.Set(ss.bs, v)
	sha := sha256.Sum256(ss.bs.Bytes())
	d.shaToSs[sha] = ss
	return d
}

func (d *StateSetDict[T]) Get(ss *StateSet) (T, bool) {
	v, ok := d.dict.Get(ss.bs)
	return v, ok
}

func (d *StateSetDict[T]) Contains(ss *StateSet) bool {
	return d.dict.Contains(ss.bs)
}

func (d *StateSetDict[T]) iterator() *stateSetDictIterator[T] {
	stSets := make([]*StateSet, 0)
	for _, ss := range d.shaToSs {
		stSets = append(stSets, ss)
	}

	return &stateSetDictIterator[T]{
		d:       d,
		stSets:  stSets,
		length:  len(stSets),
		currIdx: 0,
	}
}

type stateSetDictIterator[T any] struct {
	d       *StateSetDict[T]
	stSets  []*StateSet
	length  int
	currIdx int
}

func (iter *stateSetDictIterator[T]) HasNext() bool {
	return iter.currIdx < iter.length
}

func (iter *stateSetDictIterator[T]) Next() (*StateSet, T) {
	ss := iter.stSets[iter.currIdx]
	v, _ := iter.d.Get(ss)
	iter.currIdx++
	return ss, v
}

func charLabel(s string) string {
	switch s {
	case "\t":
		return "Tab"
	case "\n":
		return "Newline"
	case "\r":
		return "CR"
	case " ":
		return "Space"
	case "\\":
		return "Backslash"
	default:
		return s
	}
}
