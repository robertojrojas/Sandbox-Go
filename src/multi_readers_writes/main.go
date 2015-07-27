package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)


func getPage(url string) (int, error) {
	res, err := http.Get(url)
	if err != nil {
		return 0, err
	}

    defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0, err
	}


    return len(body), nil
    

}

type UrlLength struct {
	url string
	length int
}

func getter(url string, size chan UrlLength) {
	length, err := getPage(url)

	if err == nil {
		currentUrlLength := UrlLength{
			url: url,
			length: length,
		}
		size <- currentUrlLength
	}
}

func main() {

    urls := []string {
    	"http://www.google.com/",
    	"http://www.yahoo.com/",
    	"http://www.bing.com/",
    	"http://www.cnn.com/",
    }

    urlLengthChan := make(chan UrlLength)

	for _, url := range urls {
		go getter(url, urlLengthChan)
	}

	for i := 0; i < len(urls); i++ {
		urlLength := <- urlLengthChan
		fmt.Printf("%s is length %d\n", urlLength.url, urlLength.length)
	}
	

}