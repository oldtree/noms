package types

import (
	"crypto/sha1"
	"sort"

	"github.com/attic-labs/noms/chunks"
	"github.com/attic-labs/noms/ref"
)

func isSequenceOrderedByIndexedType(t Type) bool {
	return t.Desc.(CompoundDesc).ElemTypes[0].IsOrdered()
}

// Given a leaf in an ordered sequence, returns the values in that leaf which define the ordering of the sequence.
type getLeafOrderedValuesFn func(Value) []Value

// Returns a cursor to |key| in |ms|, plus the leaf + index that |key| is in. |t| is the type of the ordered values.
func findLeafInOrderedSequence(ms metaSequence, t Type, key Value, getValues getLeafOrderedValuesFn, cs chunks.ChunkSource) (cursor *sequenceCursor, leaf Value, idx int) {
	cursor, leaf = newMetaSequenceCursor(ms, cs)

	if isSequenceOrderedByIndexedType(t) {
		orderedKey := key.(OrderedValue)

		cursor.seekBinary(func(mt sequenceItem) bool {
			return !mt.(metaTuple).value.(OrderedValue).Less(orderedKey)
		})
	} else {
		cursor.seekBinary(func(mt sequenceItem) bool {
			return !mt.(metaTuple).value.(Ref).TargetRef().Less(key.Ref())
		})
	}

	if current := cursor.current().(metaTuple); current.ChildRef() != valueFromType(leaf, leaf.Type()).Ref() {
		leaf = readMetaTupleValue(current, cs)
	}

	if leafData := getValues(leaf); isSequenceOrderedByIndexedType(t) {
		orderedKey := key.(OrderedValue)

		idx = sort.Search(len(leafData), func(i int) bool {
			return !leafData[i].(OrderedValue).Less(orderedKey)
		})
	} else {
		idx = sort.Search(len(leafData), func(i int) bool {
			return !leafData[i].Ref().Less(key.Ref())
		})
	}

	return
}

func newOrderedMetaSequenceBoundaryChecker() boundaryChecker {
	return newBuzHashBoundaryChecker(orderedSequenceWindowSize, sha1.Size, objectPattern, func(item sequenceItem) []byte {
		digest := item.(metaTuple).ChildRef().Digest()
		return digest[:]
	})
}

func newOrderedMetaSequenceChunkFn(t Type) makeChunkFn {
	return func(items []sequenceItem) (sequenceItem, Value) {
		tuples := make(metaSequenceData, len(items))

		for i, v := range items {
			tuples[i] = v.(metaTuple) // chunk is written when the root sequence is written
		}

		meta := newMetaSequenceFromData(tuples, t, nil)
		return metaTuple{meta, ref.Ref{}, tuples.last().value}, meta
	}
}
