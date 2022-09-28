package main

import (
	  "fmt"
      "io/ioutil"
      "log"
      "net/http"
	  "strings"
      "os"
)

var (
	tweetURL string 
	userAgent string
    err error
)

func check(err error) {
    if err != nil {
        log.Fatal(err)
    }
}

func substr(str string, start string, end string) (result string) {
    s := strings.Index(str, start)
    s += len(start)
    substring := str[s:]
    e := strings.Index(substring, end)
    return substring[:e]
}

func main() {
	client := &http.Client{}
    url := os.Args[1] + "|"
    tweetURL = "https://nitter.net/" + substr(url, "https://twitter.com/", "|")
    userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36"
    
    req, err := http.NewRequest("GET", tweetURL, nil)
    check(err)

	req.Header.Set("User-Agent", userAgent)

	resp, err := client.Do(req)
    check(err)
	
	defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    check(err)
	
	tweet := substr(string(body), "<div class=\"tweet-content media-body\" dir=\"auto\">", "</div>")
    fmt.Println(tweet)
}