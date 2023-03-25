package main

import (
	"log"
	"os"
	"reddit/fetcher"
)

func main() {
	r := &fetcher.RedditFetcherImpl{}
	err := r.Fetch()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	file, err := os.Create("reddit.txt")
	err = r.Save(file)
	//err = r.Save(os.Stdout)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

}
