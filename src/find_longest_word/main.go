package main

import (
	"strings"
	"fmt"
	"log"
	"errors"
	"net/http"
	"io/ioutil"
	"sort"
)

type words []string

func (theWords words) Len() int {
	return len(theWords)
}
func (theWords words) Swap(i, j int) {
    theWords[i], theWords[j] = theWords[j], theWords[i]
}

func (theWords words) Less(left, right int) bool {

    leftWord  := theWords[left]
    rightWord := theWords[right]

    // Descending 
	return len(leftWord) > len(rightWord)
}

type Request struct {
	Url string
	body []byte
	Err error
}

func (r *Request) Get() {


	client := &http.Client{}
    req, err := http.NewRequest("GET", r.Url, nil)
    req.Header.Add("Accept", `text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8`)
    req.Header.Add("User-Agent", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_5) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11`)
    resp, err := client.Do(req)

	if err != nil {
		r.Err = err
		return
	}
   
    //fmt.Printf("URL: %s Status %s\n", r.Url, resp.Status)
	if resp.StatusCode != http.StatusOK {
		r.Err = errors.New(resp.Status)
		return
	}

	defer resp.Body.Close()

	r.body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		r.Err = err
		return
	}

}

func (r *Request) IsOK() (bool) {
	return r.Err == nil
}

func (r *Request) ToString() (string) {
	return string(r.body)
}

func (r *Request) Reset() {
	r.Url  = ""
	r.body = nil
	r.Err  = nil

}

func parseData(rawData string) words{

	wordsToReturn := words{}
	for _, word  := range strings.Split(rawData, "\n") {
		if word == "" {
			continue
		}
		wordsToReturn = append(wordsToReturn, word)

	}

	return wordsToReturn

}

func displayLongestWord(wordsToProcess words ) {

    // Sort Descending based on length so it's easier to find longest word
	sort.Sort(wordsToProcess)

	lastWordIdx := len(wordsToProcess)

    longestWord     := ""
    copyLongestWord := ""
    wordComposition := make([]string, 10)

	for outerIdx := 0; outerIdx < lastWordIdx-1; outerIdx++ {
		longestWord     = wordsToProcess[outerIdx]
		copyLongestWord = longestWord

        wordsSizeFound   := 0
        foundLongestWord := false 
        longestWordLen   := len(longestWord)

        // Compare the current Longest Word against the words the follow in the list
		for innerIdx := outerIdx + 1; innerIdx < lastWordIdx; innerIdx++ {

			candidateWord    := wordsToProcess[innerIdx]
			firstInstanceIdx := strings.Index(copyLongestWord, candidateWord)

			if firstInstanceIdx > -1 {

				wordParts       := strings.Split(copyLongestWord, candidateWord)
				copyLongestWord = strings.Join(wordParts, "")

				wordComposition = append(wordComposition, candidateWord)
				wordsSizeFound += len(candidateWord)
				
				if wordsSizeFound == longestWordLen {
					 foundLongestWord = true
					 break
				}

			}
		}

		if foundLongestWord {
			break
		} else {
			wordsSizeFound = 0
			wordComposition = make([]string, 10)

		}
   
	}
    
    fmt.Printf("Longest word [%s]\n", longestWord)
    fmt.Printf("Composed of words: %v\n", wordComposition)
}

func main() {

	request := &Request{}
	request.Url = "http://norvig.com/ngrams/word.list"

	request.Get()

	if !request.IsOK() {
		log.Fatalf("Unable to get list of words: %s\n", request.Err)
	}

	wordsToProcess := parseData(request.ToString())

	displayLongestWord(wordsToProcess)

}