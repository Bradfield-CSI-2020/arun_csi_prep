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
	"net/http"
	"strconv"
	"time"
	"encoding/json"
	"log"
	"bufio"
)

func main() {
	arg := parseArgs()
	if arg == "--index" {
		scrapeXkcd()
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


func printComic(value comic) {
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
func printHits(searchTerm string, searchIndex map[int]comic) {
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
func testInvertedIndex(searchTerm string, invertedIndex map[string]*invertedIndexEntry) {
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

func scrapeXkcd() {
	for i := 1; i <= 2328; i++ {
		// comic number 404 does not exist
		if (i != 404) {
			continue
		}
		
		time.Sleep(100 * time.Millisecond)

		comic, fetchErr := fetch(strconv.Itoa(i))

		if (fetchErr != nil) {
			log.Fatalf("xkcd:scrapeXkcd:fetch %v\n", fetchErr)
		}

		writeErr := write(comic)

		if (writeErr != nil) {
			log.Fatalf("xkcd:scrapeXkcd:write %v\n", writeErr)
		}
	}
}

func fetch(num string) (string, error) {
	slice := [3] string {"https://xkcd.com/", num, "/info.0.json"}
	url := strings.Join(slice[0:], "")

	resp, reqError := http.Get(url)

	if (reqError != nil) {
		return "", reqError
	}

	body, readErr := ioutil.ReadAll(resp.Body)

	resp.Body.Close()

	if (readErr != nil) {
		return "", readErr
	}

	return string(body), nil
}

func write(line string) error {
	fileName := "db.jsonl"
	file, openErr := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, os.ModeAppend)
	
	if openErr != nil {
		fmt.Fprintf(os.Stderr, "xkcd:write:open %v\n", openErr)
		newFile, err := os.Create("./db.jsonl")

		if (err != nil) {
			return err
		}

		file = newFile
	}

	_, writeErr := file.WriteString(line + "\n")
	
	if writeErr != nil {
		fmt.Fprintf(os.Stderr, "xkcd:write:write %v\n", writeErr)
		return writeErr
	}

	syncErr := file.Sync()

	return syncErr
}

func buildIndex() (map[int]comic, error) {
	fileName := "db.jsonl"

	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Fprintf(os.Stderr, "buildIndex:read %v\n", err)
		return nil, err
	}

	var index = make(map[int]comic)

	lineCount := 0
	for _, line:= range strings.Split(string(data), "\n") {
		lineCount++
		comicLine := comic{}
		json.Unmarshal([]byte(line), &comicLine)
		index[comicLine.Num] = comicLine
	}
	fmt.Println("original line count:", lineCount)
	return index, nil
}

func buildInvertedIndex() (map[string]*invertedIndexEntry, error) {
	fileName := "db.jsonl"

	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Fprintf(os.Stderr, "buildIndex:read %v\n", err)
		return nil, err
	}

	var invertedIndex = make(map[string]*invertedIndexEntry)

	for _, line:= range strings.Split(string(data), "\n") {
		// line is a document
		// go over `safe_title` and `transcript`,
		// lower case, split at whitespace,
		// for each string entry, add to invertedIndex
		// profit
		comicLine := comic{}
		json.Unmarshal([]byte(line), &comicLine)
		
		// todo: account for searchterms being substrings of fields
		titleFields := strings.Fields(comicLine.SafeTitle)
		transcriptFields := strings.Fields(comicLine.Transcript)

		combinedList := append(titleFields, transcriptFields...)

		for _, value := range combinedList {
			// entry := invertedIndexEntry{""}
			searchValue := strings.ToLower(value)
			existingEntry := invertedIndex[searchValue]

			if (existingEntry == nil) {
				entry := invertedIndexEntry{}
				entry.SearchTerm = searchValue
				entry.Comics = make(map[int]comic)
				entry.Comics[comicLine.Num] = comicLine
				entry.Frequency = 1
				invertedIndex[searchValue] = &entry
			} else {
				existingEntry.Frequency++
				existingEntry.Comics[comicLine.Num] = comicLine
				// no need to update invertedIndex, as we update the entry directly
			}
		}
	}

	return invertedIndex, nil
}

type invertedIndexEntry struct {
	SearchTerm string
	Frequency int
	Comics map[int]comic
}
type invertedIndex struct {
	HashMap map[string]*invertedIndexEntry
}

type comic struct {
	Month string `json:"month"`
	Num int `json:"num"`
	Day string `json:"day"`
	Link string `json:"link"`
	Year string `json:"year"`
	News string `json:"news"`
	SafeTitle string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt string `json:"alt"`
	Img string `json:"img"`
	Title string `json:"title"`
}
