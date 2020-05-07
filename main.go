package main

import (
	"log"
	"os"
	"github.com/urfave/cli"
)

var app = cli.NewApp()

func init() {
	info()
}

func main() {
	var err = app.Run(os.Args)
	if err != nil {
		log.Fatal("Error: Could not initialize Application")
	}
}

func info() {
	app.Name = "gitpodify"
	app.Description = "Takes a URL of a Gitpod repo, and Gitpodifies it."
	app.Usage = "gitpodify [options] username reponame branch"
	app.Author = "Sean Hellum"
	app.Version = "1.0.0"
	app.Action = App
}
