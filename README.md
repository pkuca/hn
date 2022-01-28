# hn
[![Go Reference](https://pkg.go.dev/badge/github.com/pkuca/hn.svg)](https://pkg.go.dev/github.com/pkuca/hn)

Package hn provides a client wrapping requests to the hackernews [firebase API](https://github.com/HackerNews/API).

## Installation

```bash
go get github.com/pkuca/hn
```

## Basic Usage
```golang
client := hn.NewClient(nil)
topStories, _ := client.TopStories()
```

## Custom Client

```golang
httpClient := &http.Client{
    Timeout: time.Second * 20,
}

hnClient := hn.NewClient(httpClient)
topStories, _ := hnClient.TopStories()
```
