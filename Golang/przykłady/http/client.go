package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"net/http"
	"time"
)

type Data struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
	Tags  []Tag  `json:"tags"`
}

type Tag struct {
	Name  string      `json:"tag,omitempty"`
	Value json.Number `json:"value,omitempty"`
}

func main() {
	c := NewExampleClient("http://example:4444/", time.Second*3)
	d, err := c.FetchData(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v", d)
}

type ExampleClient struct {
	c    *http.Client
	host string
}

func NewExampleClient(host string, t time.Duration) *ExampleClient {
	return &ExampleClient{
		host: host,
		c: &http.Client{
			Timeout: t,
		},
	}
}

func (e *ExampleClient) FetchData(ctx context.Context) (Data, error) {
	ctx = context.WithValue(ctx, "requestID", time.Now().Unix())
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, e.host, http.NoBody)
	if err != nil {
		return Data{}, fmt.Errorf("cannot create request: %w", err)
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := e.c.Do(req)
	if err != nil {
		return Data{}, fmt.Errorf("cannot get data: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return Data{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var data Data
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return Data{}, fmt.Errorf("cannot unmarshal data: %w", err)
	}

	return data, nil
}
