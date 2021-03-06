package main

import (
	"log"
	"os"
)

func CreateConfig() {
	file, err := createGitpodFile()
	if err != nil {
		log.Fatalf("Error!")
	}
	addTasks(file)
}

func addTasks(config *os.File) {
	if len(Languages) != 0 {
		config.WriteString("tasks:\n")
		for _, i := range Languages {
			if i == "go" {
				config.WriteString("  - name: Go\n    init: go get && go build ./... && go test ./...\n    command: go run\n")
			} else if i == "rust" {
				config.WriteString("  - name: Rust\n    init: cargo build\n")
			} else if i == "yarn" {
				config.WriteString("  - name: Yarn\n    init: yarn install\n")
			} else if i == "npm" {
				config.WriteString("  - name: npm\n    init: npm install\n")
			} else if i == "make" {
				config.WriteString("  - name: Make\n    init: make")
			}
		}
	}
}

func createGitpodFile() (*os.File, error) {
	return os.Create(".gitpod.yml")
}
