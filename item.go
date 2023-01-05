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

var (
	errItemUnmarshalID          = errors.New("item unmarshal failed on 'id' field")
	errItemUnmarshalType        = errors.New("item unmarshal failed on 'type' field")
	errItemUnmarshalBy          = errors.New("item unmarshal failed on 'by' field")
	errItemUnmarshalTime        = errors.New("item unmarshal failed on 'time' field")
	errItemUnmarshalText        = errors.New("item unmarshal failed on 'text' field")
	errItemUnmarshalParent      = errors.New("item unmarshal failed on 'parent' field")
	errItemUnmarshalPoll        = errors.New("item unmarshal failed on 'poll' field")
	errItemUnmarshalKids        = errors.New("item unmarshal failed on 'kids' field")
	errItemUnmarshalScore       = errors.New("item unmarshal failed on 'score' field")
	errItemUnmarshalTitle       = errors.New("item unmarshal failed on 'title' field")
	errItemUnmarshalParts       = errors.New("item unmarshal failed on 'parts' field")
	errItemUnmarshalDescendants = errors.New("item unmarshal failed on 'descendants' field")

	errItemUnmarshalUnknownKey = errors.New("item unmarshal failed on unknown key")
)

func (i *Item) UnmarshalJSON(b []byte) error {
	input := map[string]interface{}{}
	if err := json.Unmarshal(b, &input); err != nil {
		return fmt.Errorf("json.Unmarshal: %w", err)
	}

	for k, v := range input {
		switch strings.ToLower(k) {
		case "id":
			floatID, ok := v.(float64)
			if !ok {
				return fmt.Errorf("%w: %v", errItemUnmarshalID, v)
			}

			i.ID = int(floatID)
		case "type":
			strType, ok := v.(string)
			if !ok {
				return fmt.Errorf("%w: %v", errItemUnmarshalType, v)
			}

			i.Type = strType
		case "by":
			strBy, ok := v.(string)
			if !ok {
				return fmt.Errorf("%w: %v", errItemUnmarshalBy, v)
			}

			i.By = strBy
		case "time":
			floatTime, ok := v.(float64)
			if !ok {
				return fmt.Errorf("%w: %v", errItemUnmarshalTime, v)
			}

			i.Time = time.Unix(int64(floatTime), 0)
		case "text":
			strText, ok := v.(string)
			if !ok {
				return fmt.Errorf("%w: %v", errItemUnmarshalText, v)
			}

			i.Text = strText
		case "parent":
			floatParent, ok := v.(float64)
			if !ok {
				return fmt.Errorf("%w: %v", errItemUnmarshalParent, v)
			}

			i.Parent = int(floatParent)
		case "poll":
			floatPoll, ok := v.(float64)
			if !ok {
				return fmt.Errorf("%w: %v", errItemUnmarshalPoll, v)
			}

			i.Poll = int(floatPoll)
		case "kids":
			ret := []int{}

			kids, ok := v.([]interface{})
			if !ok {
				return fmt.Errorf("%w: %v", errItemUnmarshalKids, v)
			}

			for _, k := range kids {
				f, ok := k.(float64)
				if !ok {
					return fmt.Errorf("%w: %v", errItemUnmarshalKids, v)
				}

				ret = append(ret, int(f))
			}

			i.Kids = ret
		case "url":
			u, err := url.Parse(v.(string))
			if err != nil {
				return fmt.Errorf("url.Parse: %w", err)
			}

			i.URL = u
		case "score":
			floatScore, ok := v.(float64)
			if !ok {
				return fmt.Errorf("%w: %v", errItemUnmarshalScore, v)
			}

			i.Score = int(floatScore)
		case "title":
			strTitle, ok := v.(string)
			if !ok {
				return fmt.Errorf("%w: %v", errItemUnmarshalTitle, v)
			}

			i.Title = strTitle
		case "parts":
			ret := []int{}

			parts, ok := v.([]interface{})
			if !ok {
				return fmt.Errorf("%w: %v", errItemUnmarshalParts, v)
			}

			for _, p := range parts {
				f, ok := p.(float64)
				if !ok {
					return fmt.Errorf("%w: %v", errItemUnmarshalDescendants, v)
				}

				ret = append(ret, int(f))
			}

			i.Parts = ret
		case "descendants":
			floatDescendants, ok := v.(float64)
			if !ok {
				return fmt.Errorf("%w: %v", errItemUnmarshalDescendants, v)
			}

			i.Descendants = int(floatDescendants)
		default:
			return fmt.Errorf("%w: %v", errItemUnmarshalUnknownKey, strings.ToLower(k))
		}
	}

	return nil
}
