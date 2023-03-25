package main

import (
	"io"
	"reddit/fetcher"
)

func main() {
	var f fetcher.RedditFetcher // do not change
	var w io.Writer             // do not change

	f.Fetch()
	f.Save(w)
}
