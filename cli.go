package main

import (
	"github.com/urfave/cli"
	"fmt"
	"errors"
	"log"
)

func Info() {
	app.Name = "gitpodify"
	app.Description = "Takes a URL of a Gitpod repo, and Gitpodifies it."
	app.Usage = "gitpodify [options] 8rl"
	app.Author = "Sean Hellum"
	app.Version = "1.0.0"
	app.Action = runApp
}

func runApp(c *cli.Context) {
	if validateArgs(c) != nil {
		log.Fatal("Error Invalid Arguments")
	}
	url := c.Args().First()
	if ValidateURL(url) != nil {
		log.Fatal("Error Invalid URL")
	}
	details := ExtractDetailsFromURL(url)
	username := details[0]
	repoName := details[1]
	branch := details[2]
	repo := newRepo(username, repoName, branch)
	files := GetFiles(repo)
	CheckAll(files)
	for i := 0; i < len(files); i++ {
		fmt.Println(files[i])
	}
	fmt.Println(Languages)
	CreateConfig()
}

func validateArgs(c *cli.Context) error {
	if len(c.Args().Tail()) > 0 || !c.Args().Present() {
		return errors.New("Error: Invalid Arguments")
	}
	return nil
}
