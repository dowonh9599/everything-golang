package goroutineAndChannel

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	errGETRequestFailed = errors.New("error: GET request failed")
)

type hitResult struct {
	url        string
	statusCode int
	err        error
}

func GetURLList() []string {
	return []string{
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.naver.com/",
		"https://www.youtube.com/",
		"https://go.dev/",
		"https://www.cloudflare.com/",
		"https://www.hasura.io/",
	}
}

func PrintResult(m map[string]string) {
	for k, v := range m {
		fmt.Printf("%s ...%s\n", k, v)
	}
}

func HitURL(url string) error {
	fmt.Printf("hitting url: %s\n", url)
	resp, err := http.Get(url)
	if err != nil || (resp != nil && resp.StatusCode >= 400) {
		return errGETRequestFailed
	}
	return nil
}

// c chan<- hitResult: SEND ONLY channel
func HitURLWithChannel(url string, c chan<- hitResult) {
	fmt.Printf("hitting url: %s\n", url)
	resp, err := http.Get(url)
	statusCode := 200
	if err != nil && resp == nil {
		statusCode = 400
	} else {
		statusCode = resp.StatusCode
	}
	c <- hitResult{url, statusCode, err}
}

func TestURLChecker() {
	var results = make(map[string]string)
	urls := GetURLList()
	for _, url := range urls {
		result := "OK"
		err := HitURL(url)
		if err != nil {
			result = "FAILED"
		}
		results[url] = result
	}
	PrintResult(results)
}

func TestURLCheckerWithGoroutine() {
	var results = make(map[string]string)
	c := make(chan hitResult)
	urls := GetURLList()
	for _, url := range urls {
		go HitURLWithChannel(url, c)
	}

	for i := 0; i < len(urls); i++ {
		hitResultRes := <-c
		result := "OK"
		if hitResultRes.statusCode >= 400 {
			result = "FAILED"
		}
		results[hitResultRes.url] = result
	}

	PrintResult(results)
}
