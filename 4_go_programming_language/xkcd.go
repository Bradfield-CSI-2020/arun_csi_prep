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
)

/*
	Design consideration
	1. usage: xkcd --index
		- fetch and build an index of all comics available on xkcd?
	2. usage: xkcd 'donald trump'

*/
func main() {
	parseArgs()
}

func parseArgs() {
	// var s string = ""
	// var sep string = ""

	// for i := 0; i < len(os.Args); i++ {
	// 	s += os.Args[i] + "," 
	// }

	// fmt.Println(s)

	fmt.Println(strings.Join(os.Args[0:], " "))
}



