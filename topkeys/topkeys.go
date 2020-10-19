package topkeys

import (
	"container/heap"
	// "fmt"
)

type BigKey struct {
	KeyName string
	Bytes   uint64
	index   int
}

type TopKeyList []*BigKey

func (tkl TopKeyList) Len() int { return len(tkl) }

func (tkl TopKeyList) Less(i, j int) bool {
	return tkl[i].Bytes > tkl[j].Bytes
}

func (tkl TopKeyList) Swap(i, j int) {
	tkl[i], tkl[j] = tkl[j], tkl[i]
	tkl[i].index = i
	tkl[j].index = j
}

func (tkl *TopKeyList) Push(x interface{}) {
	n := len(*tkl)
	bigKey := x.(*BigKey)
	bigKey.index = n
	*tkl = append(*tkl, bigKey)
}

func (tkl *TopKeyList) Pop() interface{} {
	old := *tkl
	n := len(old)
	bigKey := old[n-1]
	bigKey.index = -1
	*tkl = old[0 : n-1]
	return bigKey
}

func (tkl *TopKeyList) update(bigKey *BigKey, keyName string, keyBytes uint64) {
	bigKey.KeyName = keyName
	bigKey.Bytes = keyBytes
	heap.Fix(tkl, bigKey.index)
}
