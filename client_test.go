package hn_test

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/pkuca/hn"
)

// http://hassansin.github.io/Unit-Testing-http-client-in-Go#2-by-replacing-httptransport
type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func newTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}

func TestItem(t *testing.T) {
	id := 8863

	testContentPath := fmt.Sprintf("testdata/%s", t.Name())
	testContent, err := os.Open(testContentPath)
	if err != nil {
		t.Fatal(err)
	}

	testClient := newTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(testContent)}
	})

	result, err := hn.NewClient(testClient).Item(id)
	if err != nil {
		t.Fatal(err)
	}

	if result.ID != id {
		t.Fatalf("expected %v, got %v", id, result.ID)
	}
}

func TestUser(t *testing.T) {
	id := "asicsp"

	testContentPath := fmt.Sprintf("testdata/%s", t.Name())
	testContent, err := os.Open(testContentPath)
	if err != nil {
		t.Fatal(err)
	}

	testClient := newTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(testContent)}
	})

	result, err := hn.NewClient(testClient).User(id)
	if err != nil {
		t.Fatal(err)
	}

	if result.ID != id {
		t.Fatalf("expected %v, got %v", id, result.ID)
	}
}

func TestItemComment(t *testing.T) {
	id := 2921983

	testContentPath := fmt.Sprintf("testdata/%s", t.Name())
	testContent, err := os.Open(testContentPath)
	if err != nil {
		t.Fatal(err)
	}

	testClient := newTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(testContent)}
	})

	result, err := hn.NewClient(testClient).Item(id)
	if err != nil {
		t.Fatal(err)
	}

	if result.ID != id {
		t.Fatalf("expected %v, got %v\n", id, result.ID)
	}
}

func TestItemAsk(t *testing.T) {
	id := 121003

	testContentPath := fmt.Sprintf("testdata/%s", t.Name())
	testContent, err := os.Open(testContentPath)
	if err != nil {
		t.Fatal(err)
	}

	testClient := newTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(testContent)}
	})

	result, err := hn.NewClient(testClient).Item(id)
	if err != nil {
		t.Fatal(err)
	}

	if result.ID != id {
		t.Fatalf("expected %v, got %v\n", id, result.ID)
	}
}

func TestItemJob(t *testing.T) {
	id := 192327

	testContentPath := fmt.Sprintf("testdata/%s", t.Name())
	testContent, err := os.Open(testContentPath)
	if err != nil {
		t.Fatal(err)
	}

	testClient := newTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(testContent)}
	})

	result, err := hn.NewClient(testClient).Item(id)
	if err != nil {
		t.Fatal(err)
	}

	if result.ID != id {
		t.Fatalf("expected %v, got %v\n", id, result.ID)
	}
}

func TestItemPoll(t *testing.T) {
	id := 126809

	testContentPath := fmt.Sprintf("testdata/%s", t.Name())
	testContent, err := os.Open(testContentPath)
	if err != nil {
		t.Fatal(err)
	}

	testClient := newTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(testContent)}
	})

	result, err := hn.NewClient(testClient).Item(id)
	if err != nil {
		t.Fatal(err)
	}

	if result.ID != id {
		t.Fatalf("expected %v, got %v\n", id, result.ID)
	}
}

func TestItemPollOpt(t *testing.T) {
	id := 126810

	testContentPath := fmt.Sprintf("testdata/%s", t.Name())
	testContent, err := os.Open(testContentPath)
	if err != nil {
		t.Fatal(err)
	}

	testClient := newTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(testContent)}
	})

	result, err := hn.NewClient(testClient).Item(id)
	if err != nil {
		t.Fatal(err)
	}

	if result.ID != id {
		t.Fatalf("expected %v, got %v\n", id, result.ID)
	}
}

func TestMaxItem(t *testing.T) {
	testContentPath := fmt.Sprintf("testdata/%s", t.Name())
	testContent, err := os.Open(testContentPath)
	if err != nil {
		t.Fatal(err)
	}

	testClient := newTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(testContent)}
	})

	result, err := hn.NewClient(testClient).MaxItem()
	if err != nil {
		t.Fatal(err)
	}

	if result == 0 {
		t.Fatal("no result")
	}
}

func TestTopStories(t *testing.T) {
	testContentPath := fmt.Sprintf("testdata/%s", t.Name())
	testContent, err := os.Open(testContentPath)
	if err != nil {
		t.Fatal(err)
	}

	testClient := newTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(testContent)}
	})

	result, err := hn.NewClient(testClient).TopStories()
	if err != nil {
		t.Fatal(err)
	}

	if len(result) == 0 {
		t.Fatal("no result")
	}
}

func TestNewStories(t *testing.T) {
	testContentPath := fmt.Sprintf("testdata/%s", t.Name())
	testContent, err := os.Open(testContentPath)
	if err != nil {
		t.Fatal(err)
	}

	testClient := newTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(testContent)}
	})

	result, err := hn.NewClient(testClient).NewStories()
	if err != nil {
		t.Fatal(err)
	}

	if len(result) == 0 {
		t.Fatal("no result")
	}
}

func TestBestStories(t *testing.T) {
	testContentPath := fmt.Sprintf("testdata/%s", t.Name())
	testContent, err := os.Open(testContentPath)
	if err != nil {
		t.Fatal(err)
	}

	testClient := newTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(testContent)}
	})

	result, err := hn.NewClient(testClient).BestStories()
	if err != nil {
		t.Fatal(err)
	}

	if len(result) == 0 {
		t.Fatal("no result")
	}
}

func TestAskStories(t *testing.T) {
	testContentPath := fmt.Sprintf("testdata/%s", t.Name())
	testContent, err := os.Open(testContentPath)
	if err != nil {
		t.Fatal(err)
	}

	testClient := newTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(testContent)}
	})

	result, err := hn.NewClient(testClient).AskStories()
	if err != nil {
		t.Fatal(err)
	}

	if len(result) == 0 {
		t.Fatal("no result")
	}
}

func TestShowStories(t *testing.T) {
	testContentPath := fmt.Sprintf("testdata/%s", t.Name())
	testContent, err := os.Open(testContentPath)
	if err != nil {
		t.Fatal(err)
	}

	testClient := newTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(testContent)}
	})

	result, err := hn.NewClient(testClient).ShowStories()
	if err != nil {
		t.Fatal(err)
	}

	if len(result) == 0 {
		t.Fatal("no result")
	}
}

func TestJobStories(t *testing.T) {
	testContentPath := fmt.Sprintf("testdata/%s", t.Name())
	testContent, err := os.Open(testContentPath)
	if err != nil {
		t.Fatal(err)
	}

	testClient := newTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(testContent)}
	})

	result, err := hn.NewClient(testClient).JobStories()
	if err != nil {
		t.Fatal(err)
	}

	if len(result) == 0 {
		t.Fatal("no result")
	}
}

func TestUpdates(t *testing.T) {
	testContentPath := fmt.Sprintf("testdata/%s", t.Name())
	testContent, err := os.Open(testContentPath)
	if err != nil {
		t.Fatal(err)
	}

	testClient := newTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(testContent)}
	})

	result, err := hn.NewClient(testClient).Updates()
	if err != nil {
		t.Fatal(err)
	}

	if len(result.Items) == 0 {
		t.Fatal("no result")
	}

	if len(result.Profiles) == 0 {
		t.Fatal("no result")
	}
}
