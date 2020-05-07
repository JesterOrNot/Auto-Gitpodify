package main

import (
	"os"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/urfave/cli"
)

func App(c *cli.Context) {
	username := c.Args().Get(0)
	if username == "" {
		log.Fatal("Error!")
		os.Exit(1)
	}
	repoName := c.Args().Get(1)
	if repoName == "" {
		log.Fatal("Error!")
		os.Exit(1)
	}
	branch := c.Args().Get(2)
	if branch == "" {
		branch = "master"
	}
	repo := newRepo(username, c.Args().Get(1), branch)
	// fmt.Println(repo)
	// fmt.Println("repoName:", repoName, "branch:", branch, "org:", username)
	files := GetFiles(repo)
	for i := 0; i < len(files); i++ {
		fmt.Println(files[i])
	}
}

func GetFiles(repo Repository) []string {
	response, err := http.Get(fmt.Sprintf("https://api.github.com/repos/%s/%s/git/trees/%s?recursive=1", repo.User, repo.RepoName, repo.Branch))
	if err != nil {
		log.Fatal(err)
	}
	responseData, _ := ioutil.ReadAll(response.Body)
	var responseObject Response
	json.Unmarshal(responseData, &responseObject)
	var files []string
	for i := 0; i < len(responseObject.Tree); i++ {
		item := responseObject.Tree[i]
		if isFile(item) {
			files = append(files, responseObject.Tree[i].Name)
		}
	}
	return files
}

func isFile(file GitFile) bool {
	if file.Type == "tree" {
		return false
	}
	return true
}
