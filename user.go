package hn

import (
	"encoding/json"
	"fmt"
	"html"
	"strings"
	"time"
)

// User is based on https://github.com/HackerNews/API#users
type User struct {
	ID        string    `json:"id"`
	Delay     uint      `json:"delay"`
	Created   time.Time `json:"created"`
	Karma     uint      `json:"karma"`
	About     string    `json:"about"`
	Submitted []uint    `json:"submitted"`
}

func (u *User) UnmarshalJSON(b []byte) error {
	input := map[string]interface{}{}
	if err := json.Unmarshal(b, &input); err != nil {
		return fmt.Errorf("json.Unmarshal: %w", err)
	}

	for k, v := range input {
		switch strings.ToLower(k) {
		case "id":
			u.ID = v.(string)
		// case "delay":
		// 	u.Delay = v.(uint)
		case "created":
			u.Created = time.Unix(int64(v.(float64)), 0)
		case "karma":
			u.Karma = uint(v.(float64))
		case "about":
			u.About = html.UnescapeString(v.(string))
		case "submitted":
			converted := []uint{}

			for _, sub := range v.([]interface{}) {
				f := sub.(float64)
				converted = append(converted, uint(f))
			}

			u.Submitted = converted
		}
	}

	return nil
}
