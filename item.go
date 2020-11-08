package hn

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strings"
	"time"
)

// Item is based on https://github.com/HackerNews/API#items
type Item struct {
	ID          int
	Type        string
	By          string
	Time        time.Time
	Text        string
	Parent      int
	Poll        int
	Kids        []int
	URL         *url.URL
	Score       int
	Title       string
	Parts       []int
	Descendants int
	Deleted     bool
	Dead        bool
}

var errUnknownKey = errors.New("unknown key")

func (i *Item) UnmarshalJSON(b []byte) error {
	input := map[string]interface{}{}
	if err := json.Unmarshal(b, &input); err != nil {
		return fmt.Errorf("json.Unmarshal: %w", err)
	}

	for k, v := range input {
		switch strings.ToLower(k) {
		case "id":
			i.ID = int(v.(float64))
		case "type":
			i.Type = v.(string)
		case "by":
			i.By = v.(string)
		case "time":
			i.Time = time.Unix(int64(v.(float64)), 0)
		case "text":
			i.Text = v.(string)
		case "parent":
			i.Parent = int(v.(float64))
		case "poll":
			i.Poll = int(v.(float64))
		case "kids":
			converted := []int{}

			for _, sub := range v.([]interface{}) {
				f := sub.(float64)
				converted = append(converted, int(f))
			}

			i.Kids = converted
		case "url":
			u, err := url.Parse(v.(string))
			if err != nil {
				return fmt.Errorf("url.Parse: %w", err)
			}

			i.URL = u
		case "score":
			i.Score = int(v.(float64))
		case "title":
			i.Title = v.(string)
		case "parts":
			converted := []int{}

			for _, sub := range v.([]interface{}) {
				f := sub.(float64)
				converted = append(converted, int(f))
			}

			i.Parts = converted
		case "descendants":
			i.Descendants = int(v.(float64))
		default:
			return fmt.Errorf("%w: %v", errUnknownKey, strings.ToLower(k))
		}
	}

	return nil
}
