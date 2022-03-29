package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

const allLetters = "abcdefghijklmnopqrstuvwxyz"

var wg sync.WaitGroup

func main() {
	var frequency [26]int32
	start := time.Now()
	for i := 1000; i <= 1200; i++ {
		wg.Add(1)
		go countLetters(fmt.Sprintf("https://www.rfc-editor.org/rfc/rfc%d.txt", i), &frequency)
	}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("Done, took %s\n", elapsed)
	for i, f := range frequency {
		wg.Add(1)
		fmt.Printf("%s -> %d\n", string(allLetters[i]), f)
	}

}

func countLetters(url string, frequency *[26]int32) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	for i := 0; i <= 20; i++ {
		for _, b := range body {
			c := strings.ToLower(string(b))
			index := strings.Index(allLetters, c)
			if index >= 0 {
				atomic.AddInt32(&frequency[index], 1)
			}
		}
	}
	wg.Done()
}
