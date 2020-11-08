// Package hn is based on https://github.com/HackerNews/API
package hn

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type Client struct {
	cl      *http.Client
	BaseURL url.URL
}

// NewClient is the init func for package use. The returned Client's methods
// wrap requests to firebase json endpoints. If an *http.Client is not provided,
// there is a fallback to *http.DefaultClient.
func NewClient(c *http.Client) *Client {
	u := url.URL{
		Scheme: "https",
		Host:   "hacker-news.firebaseio.com",
		Path:   "v0",
	}

	// Use the optional *http.Client from invocation.
	if c != nil {
		return &Client{cl: c, BaseURL: u}
	}

	// No *http.Client was provied - use default.
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

func (c Client) MaxItem() (uint64, error) {
	route := c.BaseURL
	route.Path += "/maxitem.json"

	data, err := c.get(route.String())
	if err != nil {
		return 0, err
	}

	result, err := strconv.ParseUint(string(data), 10, 64)
	if err != nil {
		return 0, fmt.Errorf("strconv.ParseUint: %w", err)
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

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadAll: %w", err)
	}

	return b, nil
}
