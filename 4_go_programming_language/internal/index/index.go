package index

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// BuildIndex by reading the db file line by line
func BuildIndex() (map[int]Comic, error) {
	fileName := "db.jsonl"

	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Fprintf(os.Stderr, "buildIndex:read %v\n", err)
		return nil, err
	}

	var index = make(map[int]Comic)

	lineCount := 0
	for _, line := range strings.Split(string(data), "\n") {
		lineCount++
		comicLine := Comic{}
		err := json.Unmarshal([]byte(line), &comicLine)

		if err != nil {
			fmt.Fprintf(os.Stderr, "buildIndex:marshal %v\n", err)
			return nil, err
		}
		index[comicLine.Num] = comicLine
	}
	fmt.Println("original line count:", lineCount)
	return index, nil
}

// BuildInvertedIndex by reading the db file line by line
func BuildInvertedIndex() (map[string]*InvertedIndexEntry, error) {
	fileName := "db.jsonl"

	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Fprintf(os.Stderr, "buildIndex:read %v\n", err)
		return nil, err
	}

	var invertedIndex = make(map[string]*InvertedIndexEntry)

	for _, line := range strings.Split(string(data), "\n") {
		// line is a document
		// go over `safe_title` and `transcript`,
		// lower case, split at whitespace,
		// for each string entry, add to invertedIndex
		// profit
		comicLine := Comic{}
		err := json.Unmarshal([]byte(line), &comicLine)

		if err != nil {
			fmt.Fprintf(os.Stderr, "buildIndex:marshal %v\n", err)
			return nil, err
		}

		// todo: account for searchterms being substrings of fields
		titleFields := strings.Fields(comicLine.SafeTitle)
		transcriptFields := strings.Fields(comicLine.Transcript)

		combinedList := append(titleFields, transcriptFields...)

		for _, value := range combinedList {
			searchValue := strings.ToLower(value)
			existingEntry := invertedIndex[searchValue]

			if existingEntry == nil {
				entry := InvertedIndexEntry{}
				entry.SearchTerm = searchValue
				entry.Comics = make(map[int]Comic)
				entry.Comics[comicLine.Num] = comicLine
				entry.Frequency = 1
				invertedIndex[searchValue] = &entry
			} else {
				existingEntry.Frequency++
				existingEntry.Comics[comicLine.Num] = comicLine
			}
		}
	}

	return invertedIndex, nil
}
