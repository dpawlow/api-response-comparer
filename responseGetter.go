package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

type Configs struct {
	Configurations []Config
}

type Config struct {
	Version         string
	Exposure        int
	Gal_version     string
	Layout          string
	Related_version string
	Layout_holder   string
	Qcat            string
}

func (c1 *Configs) equals(c2 Configs) (answer bool) {
	answer = true
	for i := range c1.Configurations {
		c1.Configurations[i].equals(c2.Configurations[i])
	}
	return
}

func (c1 *Config) equals(c2 Config) (answer bool) {
	answer = c1.Version == c2.Version && c1.Exposure == c2.Exposure && c1.Gal_version == c2.Gal_version
	answer = answer && c1.Layout == c2.Layout && c1.Related_version == c2.Related_version
	answer = answer && c1.Layout_holder == c2.Layout_holder && c1.Qcat == c2.Qcat
	return
}

func getResponses(urls []string) []Configs {
	var responses = make([]Configs, 2)

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
	var diffChan chan []Configs = make(chan []Configs, len(urls))

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
		diffUrls[i] = elem[0].equals(elem[1])
		i++
	}

	return diffUrls
}
