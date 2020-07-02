package main

// Exercise 4.12: The popular web comic xkcd has a JSON interface. 
// For example, a request to https://xkcd.com/571/info.0.json produces a detailed description of comic 571,
// one of many favorites. 
// Download each URL (once!) and build an offline index. 
// Write a tool xkcd that, using this index, prints the URL and transcript of each comic that matches a search term 
// provided on the command line

import (
	"fmt"
	"os"
	"strings"
	// "strconv"
	"io/ioutil"
	"net/http"
)

/*
	Design consideration
	1. usage: xkcd --index
		- fetch and build an index of all comics available on xkcd?
	2. usage: xkcd 'donald trump'
*/
func main() {
	parseArgs()
	fetch("1")
}

func parseArgs() string {
	var search = os.Args[1]
	return search
}

func fetch(num string) string {
	var slice = [3] string {"https://xkcd.com/", num, "/info.0.json"}
	var url = strings.Join(slice[0:], "")

	resp, _ := http.Get(url)
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Printf("%s", body)
	return string(body)
}



