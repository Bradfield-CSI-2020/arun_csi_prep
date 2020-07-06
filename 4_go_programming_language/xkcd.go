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
)

/*
	Design consideration
	1. usage: xkcd --index
		- fetch and build an index of all comics available on xkcd?
	2. usage: xkcd 'donald trump'
*/
func main() {
	// fmt.Println(parseArgs())
	// fmt.Println(fetch("1"))
	indexXkcd()
	// fmt.Println(readFile())
}

func parseArgs() string {
	var search = os.Args[1]
	return search
}

func indexXkcd() {
	for i := 1; i < 10; i++ {
		fmt.Println(fetch(strconv.Itoa(i)))
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

func readFile() string {
	fileName := "db.jsonl"

	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
	}

	return string(data)
}


// todo: write an index file that has a init function that creates the in-memory index of comics?