package hn_test

import (
	"net/http"
	"testing"

	"github.com/pkuca/hn"
)

func TestItem(t *testing.T) {
	id := 8863

	result, err := hn.NewClient(nil).Item(id)
	if err != nil {
		t.Fatal(err)
	}

	if result.ID != id {
		t.Fatalf("expected %v, got %v", id, result.ID)
	}
}

func TestUser(t *testing.T) {
	id := "asicsp"

	result, err := hn.NewClient(nil).User(id)
	if err != nil {
		t.Fatal(err)
	}

	if result.ID != id {
		t.Fatalf("expected %v, got %v", id, result.ID)
	}
}

func TestItemComment(t *testing.T) {
	id := 2921983

	result, err := hn.NewClient(nil).Item(id)
	if err != nil {
		t.Fatal(err)
	}

	if result.ID != id {
		t.Fatalf("expected %v, got %v\n", id, result.ID)
	}
}

func TestItemAsk(t *testing.T) {
	id := 121003

	result, err := hn.NewClient(nil).Item(id)
	if err != nil {
		t.Fatal(err)
	}

	if result.ID != id {
		t.Fatalf("expected %v, got %v\n", id, result.ID)
	}
}

func TestItemJob(t *testing.T) {
	id := 192327

	result, err := hn.NewClient(nil).Item(id)
	if err != nil {
		t.Fatal(err)
	}

	if result.ID != id {
		t.Fatalf("expected %v, got %v\n", id, result.ID)
	}
}

func TestItemPoll(t *testing.T) {
	id := 126809

	result, err := hn.NewClient(nil).Item(id)
	if err != nil {
		t.Fatal(err)
	}

	if result.ID != id {
		t.Fatalf("expected %v, got %v\n", id, result.ID)
	}
}

func TestItemPollOpt(t *testing.T) {
	id := 126810

	result, err := hn.NewClient(nil).Item(id)
	if err != nil {
		t.Fatal(err)
	}

	if result.ID != id {
		t.Fatalf("expected %v, got %v\n", id, result.ID)
	}
}

func TestMaxItem(t *testing.T) {
	result, err := hn.NewClient(nil).MaxItem()
	if err != nil {
		t.Fatal(err)
	}

	if result == 0 {
		t.Fatal("no result")
	}
}

func TestTopStories(t *testing.T) {
	result, err := hn.NewClient(nil).TopStories()
	if err != nil {
		t.Fatal(err)
	}

	if len(result) == 0 {
		t.Fatal("no result")
	}
}

func TestNewStories(t *testing.T) {
	result, err := hn.NewClient(nil).NewStories()
	if err != nil {
		t.Fatal(err)
	}

	if len(result) == 0 {
		t.Fatal("no result")
	}
}

func TestBestStories(t *testing.T) {
	result, err := hn.NewClient(nil).BestStories()
	if err != nil {
		t.Fatal(err)
	}

	if len(result) == 0 {
		t.Fatal("no result")
	}
}

func TestAskStories(t *testing.T) {
	result, err := hn.NewClient(nil).AskStories()
	if err != nil {
		t.Fatal(err)
	}

	if len(result) == 0 {
		t.Fatal("no result")
	}
}

func TestShowStories(t *testing.T) {
	result, err := hn.NewClient(nil).ShowStories()
	if err != nil {
		t.Fatal(err)
	}

	if len(result) == 0 {
		t.Fatal("no result")
	}
}

func TestJobStories(t *testing.T) {
	result, err := hn.NewClient(nil).JobStories()
	if err != nil {
		t.Fatal(err)
	}

	if len(result) == 0 {
		t.Fatal("no result")
	}
}

func TestUpdates(t *testing.T) {
	result, err := hn.NewClient(nil).Updates()
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

func TestWithProvidedHTTPClient(t *testing.T) {
	cl := &http.Client{}

	result, err := hn.NewClient(cl).TopStories()
	if err != nil {
		t.Fatal(err)
	}

	if len(result) == 0 {
		t.Fatal("no result")
	}
}
