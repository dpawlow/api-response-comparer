package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"sync"
)

func getResponses(urls []string) []Search {
	var responses = make([]Search, 2)

	var wg sync.WaitGroup
	for i, url := range urls {
		wg.Add(1)
		go func(urlPath string, target interface{}) {
			getJSONResponse(urlPath, &target)
			fmt.Println(target)
			wg.Done()
		}(url, &responses[i])
	}
	wg.Wait()

	return responses
}

func getJSONResponse(url string, target interface{}) error {
	r, err1 := http.Get(url)
	defer r.Body.Close()
	body, err2 := ioutil.ReadAll(r.Body)

	if err1 != nil && err2 != nil {
		fmt.Sprintln("Error processing json for %s", url)
		println(err1, err2)
	}
	return json.Unmarshal(body, &target)
}

func compareResponses(urls [][]string) []bool {
	var diffUrls = make([]bool, len(urls))
	var diffChan chan []Search = make(chan []Search, len(urls))

	fmt.Println(diffUrls)

	var wg sync.WaitGroup
	for i := range urls {
		wg.Add(1)
		go func(urls *[][]string) {
			diffChan <- getResponses((*urls)[i])
			wg.Done()
		}(&urls)
	}
	wg.Wait()
	close(diffChan)

	var i int = 0
	for elem := range diffChan {
		diffUrls[i] = reflect.DeepEqual(elem[0], elem[1])
		i++
	}

	return diffUrls
}
