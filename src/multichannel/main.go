package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	//"os"
)

func generator(url string, urlChannel chan string) {
	urlChannel <- url
}

func getPage(url string) (int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	return len(body), nil
}

func worker(urlChannel chan string, sizeChannel chan string, id int) {
	for {
		url := <-urlChannel
		length, err := getPage(url)
		if err == nil {
			sizeChannel <- fmt.Sprintf("[%d] %s has length: %d\n", id, url, length)
		} else {
			sizeChannel <- fmt.Sprintf("[%d]Error getting %s: %s\n", id, url, err)
		}
	}
}

func main() {
	urls := []string{"http://www.google.com/", "http://www.yahoo.com/",
		"http://www.bing.com/", "http://www.bbc.com/"}
	sizeChannel := make(chan string)
	urlChannel := make(chan string)
	defer close(sizeChannel)
	defer close(urlChannel)

	for i := 0; i < 10; i++ {
		go worker(urlChannel, sizeChannel, i)
	}

	for i := 0; i < len(urls); i++ {
		go generator(urls[i], urlChannel)
	}

	for i := 0; i < len(urls); i++ {
		fmt.Printf("%s", <-sizeChannel)
	}

}
