package fetcher

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

type Response struct {
	Data struct {
		Children []struct {
			Data struct {
				Title string `json:"title"`
				URL   string `json:"url"`
			} `json:"data"`
		} `json:"children"`
	} `json:"data"`
}

type RedditFetcher interface {
	Fetch() error
	Save(io.Writer) error
}

type RedditFetcherImpl struct {
	redditResponse Response
}

func (r *RedditFetcherImpl) Fetch() error {
	requestURL := "https://www.reddit.com/r/golang.json"

	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	log.Printf("client: got response!")
	log.Printf("client: status code: %d\n", res.StatusCode)

	if res.StatusCode != http.StatusOK {
		return errors.New("client: bad status code")
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(resBody, &r.redditResponse)
	if err != nil {
		return err
	}

	return nil
}

func (r *RedditFetcherImpl) Save(w io.Writer) error {
	for _, child := range r.redditResponse.Data.Children {
		title := []byte(child.Data.Title + "\n")
		url := []byte(child.Data.URL + "\n")
		_, err := w.Write(title)
		if err != nil {
			return err
		}

		_, err = w.Write(url)
		if err != nil {
			return err
		}
	}
	return nil
}
