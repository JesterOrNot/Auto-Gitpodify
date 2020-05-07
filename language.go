package main

import (
	"sync"
)
var waitGroup sync.WaitGroup = sync.WaitGroup {}
var Languages []string = make([]string, 0, 2)

func checkAll(files []string) {
	waitGroup.Add(2)
	go checkRust(files)
	go checkYarn(files)
	waitGroup.Wait()
}

func checkRust(files []string) {
	defer waitGroup.Done()
	for _, file := range files {
		if file == "Cargo.toml" {
			Languages = append(Languages, "rust")
		}
	}
}

func checkYarn(files []string) {
	defer waitGroup.Done()
	for _, file := range files {
		if file == "yarn.lock" {
			Languages = append(Languages, "Yarn")
		}
	}
}
