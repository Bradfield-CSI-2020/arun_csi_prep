// Exercise 4.12: The popular web comic xkcd has a JSON interface.
// For example, a request to https://xkcd.com/571/info.0.json produces a detailed description of comic 571,
// one of many favorites.
// Download each URL (once!) and build an offline index.
// Write a tool xkcd that, using this index, prints the URL and transcript of each comic that matches a search term
// provided on the command line
package main

// todo: write an index file that has a init function that creates the in-memory index of comics?
import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/Bradfield-CSI-2020/arun_csi_prep/4_go_programming_language/internal/db"
	"github.com/Bradfield-CSI-2020/arun_csi_prep/4_go_programming_language/internal/index"
)

func main() {
	arg := parseArgs()
	if arg == "--index" {
		db.Scrape()
		os.Exit(0)
	}

	invertedIndex, err := index.BuildInvertedIndex()

	if err != nil {
		log.Fatalf("xkcd:main %v\n", err)
	}

	if arg != "" {
		testInvertedIndex(arg, invertedIndex)
		// exit if invoked with arg
		os.Exit(0)
	}

	newInput := bufio.NewScanner(os.Stdin)

	for newInput.Scan() {
		cleanInput := strings.TrimSpace(newInput.Text())
		if cleanInput != "" {
			testInvertedIndex(cleanInput, invertedIndex)
		}
	}
}

func printComic(value index.Comic) {
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
func printHits(searchTerm string, searchIndex map[int]index.Comic) {
	count := 0
	for _, value := range searchIndex {
		safeTitleHits := strings.Contains(strings.ToLower(value.SafeTitle), strings.ToLower(searchTerm))
		transcriptHits := strings.Contains(strings.ToLower(value.Transcript), strings.ToLower(searchTerm))

		if safeTitleHits || transcriptHits {
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
	if hits == nil {
		fmt.Println("no hits !!")
	} else {
		foundComics := hits.Comics

		for _, value := range foundComics {
			printComic(value)
		}
	}
}

func parseArgs() string {
	if len(os.Args) > 2 {
		log.Fatalf("xkcd:parseArgs %v\n", fmt.Errorf("Invalid number of args"))
	}

	if len(os.Args) == 1 {
		return ""
	}

	return strings.TrimSpace(os.Args[1])
}
