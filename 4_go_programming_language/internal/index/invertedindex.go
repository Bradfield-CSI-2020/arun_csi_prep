package index

// InvertedIndexEntry is awesom
type InvertedIndexEntry struct {
	SearchTerm string
	Frequency  int
	Comics     map[int]Comic
}

// InvertedIndex is awesom
type InvertedIndex struct {
	HashMap map[string]*InvertedIndexEntry
}
