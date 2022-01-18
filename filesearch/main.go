package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

var (
	matches   []string
	waitGroup = sync.WaitGroup{}
	lock      = sync.Mutex{}
)

func fileSearch(root string, filename string) {
	fmt.Println("Searching in ", root)
	files, _ := ioutil.ReadDir(root)
	for _, file := range files {
		if strings.Contains(file.Name(), filename) {
			lock.Lock()
			matches = append(matches, filepath.Join(root, file.Name()))
			lock.Unlock()
		}
		if file.IsDir() {
			waitGroup.Add(1)
			go fileSearch(filepath.Join(root, file.Name()), filename)
		}
	}
	waitGroup.Done()
}

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Missing expected args: root, filename")
	}
	root := os.Args[1]
	filename := os.Args[2]

	start := time.Now()
	waitGroup.Add(1)
	fileSearch(root, filename)
	waitGroup.Wait()
	elapsed := time.Since(start).Milliseconds()
	for _, file := range matches {
		fmt.Println("Matched", file)
	}
	fmt.Println("Took ", elapsed)
}
