// Package hn provides an HTTP client with methods mapped to https://github.com/HackerNews/API
package hn

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

// Client represents the "HN Client" - it's just an http client with a base
// URL. Methods are mapped to the hackernews firebase API described here:
// https://github.com/HackerNews/API
type Client struct {
	cl      *http.Client
	BaseURL url.URL
}

// NewClient creates an "HN Client". Its methods wrap requests to firebase json
// API endpoints. If an *http.Client is not provided, an *http.DefaultClient is
// used.
func NewClient(c *http.Client) *Client {
	u := url.URL{
		Scheme: "https",
		Host:   "hacker-news.firebaseio.com",
		Path:   "v0",
	}

	// Use passed *http.Client if provided.
	if c != nil {
		return &Client{cl: c, BaseURL: u}
	}

	// No *http.Client was provided - use default.
	return &Client{cl: http.DefaultClient, BaseURL: u}
}

func (c Client) Item(id int) (*Item, error) {
	route := c.BaseURL
	route.Path += fmt.Sprintf("/item/%v.json", id)

	result := &Item{}

	b, err := c.get(route.String())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(b, result); err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %w", err)
	}

	return result, nil
}

func (c Client) User(id string) (*User, error) {
	route := c.BaseURL
	route.Path += fmt.Sprintf("/user/%v.json", id)

	result := &User{}

	b, err := c.get(route.String())
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(b, result); err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %w", err)
	}

	return result, nil
}

func (c Client) MaxItem() (int, error) {
	route := c.BaseURL
	route.Path += "/maxitem.json"

	data, err := c.get(route.String())
	if err != nil {
		return 0, err
	}

	result, err := strconv.Atoi(string(data))
	if err != nil {
		return 0, fmt.Errorf("strconv.Atoi: %w", err)
	}

	return result, nil
}

func (c Client) TopStories() ([]int, error) {
	route := c.BaseURL
	route.Path += "/topstories.json"

	data, err := c.get(route.String())
	if err != nil {
		return nil, err
	}

	result := make([]int, 0)
	if err = json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %w", err)
	}

	return result, nil
}

func (c Client) NewStories() ([]int, error) {
	route := c.BaseURL
	route.Path += "/newstories.json"

	data, err := c.get(route.String())
	if err != nil {
		return nil, err
	}

	result := make([]int, 0)
	if err = json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %w", err)
	}

	return result, nil
}

func (c Client) BestStories() ([]int, error) {
	route := c.BaseURL
	route.Path += "/beststories.json"

	data, err := c.get(route.String())
	if err != nil {
		return nil, err
	}

	result := make([]int, 0)
	if err = json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %w", err)
	}

	return result, nil
}

func (c Client) AskStories() ([]int, error) {
	route := c.BaseURL
	route.Path += "/askstories.json"

	data, err := c.get(route.String())
	if err != nil {
		return nil, err
	}

	result := make([]int, 0)
	if err = json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %w", err)
	}

	return result, nil
}

func (c Client) ShowStories() ([]int, error) {
	route := c.BaseURL
	route.Path += "/showstories.json"

	data, err := c.get(route.String())
	if err != nil {
		return nil, err
	}

	result := make([]int, 0)
	if err = json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %w", err)
	}

	return result, nil
}

func (c Client) JobStories() ([]int, error) {
	route := c.BaseURL
	route.Path += "/jobstories.json"

	data, err := c.get(route.String())
	if err != nil {
		return nil, err
	}

	result := make([]int, 0)
	if err = json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %w", err)
	}

	return result, nil
}

func (c Client) Updates() (*Updates, error) {
	route := c.BaseURL
	route.Path += "/updates.json"

	data, err := c.get(route.String())
	if err != nil {
		return nil, err
	}

	result := &Updates{}
	if err := json.Unmarshal(data, result); err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %w", err)
	}

	return result, nil
}

func (c Client) get(url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodGet,
		url,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("NewRequestWithContext: %w", err)
	}

	res, err := c.cl.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http.Client.Do: %w", err)
	}
	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadAll: %w", err)
	}

	return b, nil
}
