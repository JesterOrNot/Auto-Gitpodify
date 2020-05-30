package main

import (
	"log"
	"os"
	"sync"

	"github.com/urfave/cli"
)

var app = cli.NewApp()
var waitGroup sync.WaitGroup = sync.WaitGroup{}
var Languages []string

func init() {
	Info()
}

func main() {
	var err = app.Run(os.Args)
	if err != nil {
		log.Fatal("Error: Could not initialize Application")
	}
}
