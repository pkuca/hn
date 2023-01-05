package hn

import (
	"encoding/json"
	"errors"
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

var (
	errUserUnmarshalID        = errors.New("user unmarshal failed on 'id' field")
	errUserUnmarshalDelay     = errors.New("user unmarshal failed on 'delay' field")
	errUserUnmarshalCreated   = errors.New("user unmarshal failed on 'created' field")
	errUserUnmarshalKarma     = errors.New("user unmarshal failed on 'karma' field")
	errUserUnmarshalAbout     = errors.New("user unmarshal failed on 'about' field")
	errUserUnmarshalSubmitted = errors.New("user unmarshal failed on 'submitted' field")

	errUserUnmarshalUnknownKey = errors.New("user unmarshal failed on unknown key")
)

func (u *User) UnmarshalJSON(b []byte) error {
	input := map[string]interface{}{}
	if err := json.Unmarshal(b, &input); err != nil {
		return fmt.Errorf("json.Unmarshal: %w", err)
	}

	for k, v := range input {
		switch strings.ToLower(k) {
		case "id":
			strID, ok := v.(string)
			if !ok {
				return fmt.Errorf("%w: %v", errUserUnmarshalID, v)
			}

			u.ID = strID
		case "delay":
			uintDelay, ok := v.(uint)
			if !ok {
				return fmt.Errorf("%w: %v", errUserUnmarshalDelay, v)
			}

			u.Delay = uintDelay
		case "created":
			floatCreated, ok := v.(float64)
			if !ok {
				return fmt.Errorf("%w: %v", errUserUnmarshalCreated, v)
			}

			u.Created = time.Unix(int64(floatCreated), 0)
		case "karma":
			floatKarma, ok := v.(float64)
			if !ok {
				return fmt.Errorf("%w: %v", errUserUnmarshalKarma, v)
			}

			u.Karma = uint(floatKarma)
		case "about":
			strAbout, ok := v.(string)
			if !ok {
				return fmt.Errorf("%w: %v", errUserUnmarshalAbout, v)
			}

			u.About = html.UnescapeString(strAbout)
		case "submitted":
			ret := []uint{}

			submitted, ok := v.([]interface{})
			if !ok {
				return fmt.Errorf("%w: %v", errUserUnmarshalSubmitted, v)
			}

			for _, s := range submitted {
				f, ok := s.(float64)
				if !ok {
					return fmt.Errorf("%w: %v", errUserUnmarshalSubmitted, v)
				}

				ret = append(ret, uint(f))
			}

			u.Submitted = ret
		default:
			return fmt.Errorf("%w: %v", errUserUnmarshalUnknownKey, strings.ToLower(k))
		}
	}

	return nil
}
