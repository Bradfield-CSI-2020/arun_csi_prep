package ds

import (
	"github.com/Bradfield-CSI-2020/arun_csi_prep/4_go_programming_language/xkcdlib"
)

// InvertedIndexEntry is awesom
type InvertedIndexEntry struct {
	SearchTerm string
	Frequency int
	Comics map[int]xkcdlib.Comic
}

// InvertedIndex is awesom
type InvertedIndex struct {
	HashMap map[string]*InvertedIndexEntry
}