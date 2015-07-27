package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"runtime"
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




func worker(urlCh chan string, sizeCh chan string, id int) {

    fmt.Printf("wrk(%d) - ready \n", id)
	for {
        url := <- urlCh
        fmt.Printf("wrk(%d) - working on %s\n", id, url)
        length, err := getPage(url)

		if err == nil {
			sizeCh <- fmt.Sprintf("wrk(%d) - %s is length %d \n", id, url, length)
		} else {
			sizeCh <- fmt.Sprintf("wrk(%d) - Error getting %s: %s (%d)\n", id, url, err)
		}

	}
	
}

func workGenerator(url string, urlCh chan string) {
	urlCh <- url
}

func main() {

 	numCPUs := runtime.NumCPU()
	fmt.Printf("Number cpus(%d)\n", numCPUs)

	runtime.GOMAXPROCS(numCPUs)

    urls := []string {
    	"http://www.google.com/",
    	"http://www.yahoo.com/",
    	"http://www.bing.com/",
    	"http://www.cnn.com/",
    }

    urlChan  := make(chan string)
    lengthChan := make(chan string)

    // Start workers
    for i := 0; i < numCPUs; i++ {
    	go worker(urlChan, lengthChan, i)
    }

    fmt.Println("Sending work...")
	
	// Send work to workers
	for _, url := range urls {
		go workGenerator(url, urlChan)
	}
	
	fmt.Println("Waiting for results...")
    // Display results from workers
	for i := 0; i < len(urls); i++ {
		fmt.Printf(<- lengthChan)
	}
	

}