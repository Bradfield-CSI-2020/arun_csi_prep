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
	_ "bufio"
)

func main() {
	arg := parseArgs()
	if arg == "--index" {
		scrapeXkcd()
		os.Exit(0)
	} 

	index, err := buildIndex()

	if (err != nil) {
		log.Fatalf("xkcd:main %v\n", err)
	}

	printHits(arg, index)

	// newInput := bufio.NewScanner(os.Stdin)

	// for newInput.Scan() {
	// 	fmt.Println("input:", newInput.Text())
	// }
}


func printHits(searchTerm string, searchIndex map[int]comic) {
		for _, value := range searchIndex {
		// todo: check all relevant keys for match
		safeTitleHits := strings.Contains(strings.ToLower(value.SafeTitle), strings.ToLower(searchTerm))
		transcriptHits := strings.Contains(strings.ToLower(value.Transcript), strings.ToLower(searchTerm))

		if (safeTitleHits || transcriptHits) {
			// url should look like this -> https://xkcd.com/1/
			url := "https://xkcd.com/" + strconv.Itoa(value.Num)
			fmt.Println("Url:", url)
			fmt.Println("Transcript:", value.Transcript)
			fmt.Println("-------------------------------------")
		}
	}
}

func parseArgs() string {
	if (len(os.Args) != 2) {
		log.Fatalf("xkcd:parseArgs %v\n", fmt.Errorf("Invalid number of args"))
	}

	return os.Args[1]
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

	for _, line:= range strings.Split(string(data), "\n") {
		comicLine := comic{}
		json.Unmarshal([]byte(line), &comicLine)
		index[comicLine.Num] = comicLine
	}

	return index, nil
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
