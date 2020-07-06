// Exercise 4.12: The popular web comic xkcd has a JSON interface. 
// For example, a request to https://xkcd.com/571/info.0.json produces a detailed description of comic 571,
// one of many favorites. 
// Download each URL (once!) and build an offline index. 
// Write a tool xkcd that, using this index, prints the URL and transcript of each comic that matches a search term 
// provided on the command line
package main


import (
	"fmt"
	"os"
	"strings"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func main() {
	var arg = parseArgs()
	if arg == "--index" {
		// todo: use flags so index can be used as a search term
		scrapeXkcd()
		os.Exit(0)
	} 

	fmt.Println(arg)
}

func parseArgs() string {
	if (len(os.Args) != 2) {
		err := fmt.Errorf("Invalid number of args")
		fmt.Fprintf(os.Stderr, "xkcd:parseArgs %v\n", err)
		os.Exit(1)
	}

	var searchTerm = os.Args[1]
	return searchTerm
}

func scrapeXkcd() {
	for i := 1; i <= 2328; i++ {
		// comic number 404 does not exist
		if (i != 404) {
			time.Sleep(100 * time.Millisecond)
			comic := fetch(strconv.Itoa(i))
			write(comic)
		}
	}
}

func fetch(num string) string {
	var slice = [3] string {"https://xkcd.com/", num, "/info.0.json"}
	var url = strings.Join(slice[0:], "")

	resp, _ := http.Get(url)
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	// fmt.Printf("%s", body)
	return string(body)
}

func write(line string) {
	// todo: calculate path via __dirname eq?	
	fileName := "db.jsonl"
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, os.ModeAppend)
	
	if err != nil {
		fmt.Fprintf(os.Stderr, "write:open: %v\n", err)
	}

	done, err := file.WriteString(line + "\n")
	
	fmt.Println("write result: " + strconv.Itoa(done))

	if err != nil {
		fmt.Fprintf(os.Stderr, "write:write %v\n", err)
	}

	file.Sync()
}

func readFile() string {
	fileName := "db.jsonl"

	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
	}

	return string(data)
}

// todo: write an index file that has a init function that creates the in-memory index of comics?