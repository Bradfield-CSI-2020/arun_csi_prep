// Exercise 4.12: The popular web comic xkcd has a JSON interface. 
// For example, a request to https://xkcd.com/571/info.0.json produces a detailed description of comic 571,
// one of many favorites. 
// Download each URL (once!) and build an offline index. 
// Write a tool xkcd that, using this index, prints the URL and transcript of each comic that matches a search term 
// provided on the command line
package main

// todo: write an index file that has a init function that creates the in-memory index of comics?

import (
	"fmt"
	"os"
	"strings"
	"io/ioutil"
	"strconv"
	"encoding/json"
	"log"
	"bufio"
	"github.com/Bradfield-CSI-2020/arun_csi_prep/4_go_programming_language/xkcdlib"
	"github.com/Bradfield-CSI-2020/arun_csi_prep/4_go_programming_language/index"
)

func main() {
	arg := parseArgs()
	if arg == "--index" {
		xkcdlib.Scrape()
		os.Exit(0)
	} 

	invertedIndex, err := buildInvertedIndex()

	if (err != nil) {
		log.Fatalf("xkcd:main %v\n", err)
	}
	
	if (arg != "") {
		testInvertedIndex(arg, invertedIndex)
		// exit if invoked with arg
		os.Exit(0)
	}

	newInput := bufio.NewScanner(os.Stdin)

	for newInput.Scan() {
		cleanInput := strings.TrimSpace(newInput.Text())
		if (cleanInput != "") {
			testInvertedIndex(cleanInput, invertedIndex)
		}
	}
}


func printComic(value xkcdlib.Comic) {
	// url should look like this -> https://xkcd.com/1/
	url := "https://xkcd.com/" + strconv.Itoa(value.Num)
	fmt.Println("-------------------------------------")
	fmt.Println("Url:", url)
	fmt.Println("Title:", value.SafeTitle)
	fmt.Println("Transcript:")
	fmt.Println(value.Transcript)
	fmt.Println("-------------------------------------")
}

// not used, keeping for future benchmarking
func printHits(searchTerm string, searchIndex map[int]xkcdlib.Comic) {
	count := 0
	for _, value := range searchIndex {
		safeTitleHits := strings.Contains(strings.ToLower(value.SafeTitle), strings.ToLower(searchTerm))
		transcriptHits := strings.Contains(strings.ToLower(value.Transcript), strings.ToLower(searchTerm))
		
		if (safeTitleHits || transcriptHits) {
			count++
			printComic(value)
		}
	}
	fmt.Println("count original:", count)
}

// only checks for exact (but caseinsensitive) matches:
// 'space' will match 'Space', 'spAce', 'SPACE' --
// -- but will not match 'Myspace'
// todo: match substrings too
func testInvertedIndex(searchTerm string, invertedIndex map[string]*index.InvertedIndexEntry) {
	hits := invertedIndex[strings.ToLower(searchTerm)]
	if (hits == nil) {
		fmt.Println("no hits !!")
	} else {
		foundComics := hits.Comics

		for _, value := range foundComics {
			printComic(value)
		}
	}
}

func parseArgs() string {
	if (len(os.Args) > 2) {
		log.Fatalf("xkcd:parseArgs %v\n", fmt.Errorf("Invalid number of args"))
	}

	if (len(os.Args) == 1) {
		return ""
	}

	return strings.TrimSpace(os.Args[1])
}

func buildIndex() (map[int]xkcdlib.Comic, error) {
	fileName := "db.jsonl"

	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Fprintf(os.Stderr, "buildIndex:read %v\n", err)
		return nil, err
	}

	var index = make(map[int]xkcdlib.Comic)

	lineCount := 0
	for _, line:= range strings.Split(string(data), "\n") {
		lineCount++
		comicLine := xkcdlib.Comic{}
		json.Unmarshal([]byte(line), &comicLine)
		index[comicLine.Num] = comicLine
	}
	fmt.Println("original line count:", lineCount)
	return index, nil
}

func buildInvertedIndex() (map[string]*index.InvertedIndexEntry, error) {
	fileName := "db.jsonl"

	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Fprintf(os.Stderr, "buildIndex:read %v\n", err)
		return nil, err
	}

	var invertedIndex = make(map[string]*index.InvertedIndexEntry)

	for _, line:= range strings.Split(string(data), "\n") {
		// line is a document
		// go over `safe_title` and `transcript`,
		// lower case, split at whitespace,
		// for each string entry, add to invertedIndex
		// profit
		comicLine := xkcdlib.Comic{}
		json.Unmarshal([]byte(line), &comicLine)
		
		// todo: account for searchterms being substrings of fields
		titleFields := strings.Fields(comicLine.SafeTitle)
		transcriptFields := strings.Fields(comicLine.Transcript)

		combinedList := append(titleFields, transcriptFields...)

		for _, value := range combinedList {
			searchValue := strings.ToLower(value)
			existingEntry := invertedIndex[searchValue]

			if (existingEntry == nil) {
				entry := index.InvertedIndexEntry{}
				entry.SearchTerm = searchValue
				entry.Comics = make(map[int]xkcdlib.Comic)
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
