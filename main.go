package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func trace(s string) (string, time.Time) {
	log.Println("START:", s)
	return s, time.Now()
}

func un(s string, startTime time.Time) {
	endTime := time.Now()
	log.Println("  END:", s, "ElapsedTime in seconds:", endTime.Sub(startTime))
}

func main() {
	filepath := flag.String("path", "", "Path to files.")
	flag.Parse()

	defer un(trace("DURATION:"))

	file, err := os.Open(*filepath)
	defer file.Close()

	if err != nil {
		print(err)
	}

	fileInfo, _ := file.Stat()
	urls := make([][]string, fileInfo.Size())

	for i := range urls {
		urls[i] = make([]string, 2)
	}

	var counter int32 = 0
	var url0 string
	var url1 string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var line string = scanner.Text()
		if counter%2 == 0 {
			url0 = line
		}
		if counter%2 == 1 {
			url1 = line
			urls[(counter-1)/2][0] = url0
			urls[(counter-1)/2][1] = url1
		}
		counter++
	}

	for i := range urls {
		urls[i] = make([]string, 2)
	}

	var diffResponses []bool = compareResponses(urls)
	fmt.Println(diffResponses)
}
