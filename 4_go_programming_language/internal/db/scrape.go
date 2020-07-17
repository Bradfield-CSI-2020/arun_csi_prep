package db

import (
	"log"
	"strconv"
	"time"
)

// Scrape - scrape xkcd and write to file
func Scrape() {
	for i := 1; i <= 2328; i++ {
		// comic number 404 does not exist
		if i != 404 {
			continue
		}

		time.Sleep(100 * time.Millisecond)

		comic, fetchErr := Fetch(strconv.Itoa(i))

		if fetchErr != nil {
			log.Fatalf("xkcd:scrapeXkcd:fetch %v\n", fetchErr)
		}

		writeErr := Write(comic)

		if writeErr != nil {
			log.Fatalf("xkcd:scrapeXkcd:write %v\n", writeErr)
		}
	}
}
