package main

import (
	"io"
	"log"
	"os"
	"reddit/fetcher"
)

func main() {
	var f fetcher.RedditFetcher // do not change
	var w io.Writer             // do not change

	f = &fetcher.RedditFetcherImpl{}
	err := f.Fetch()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	w, err = os.Create("reddit.txt")

	err = f.Save(w)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
