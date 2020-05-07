package main

import (
	"errors"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"

	"github.com/urfave/cli"
)

func App(c *cli.Context) {
	url := c.Args().First()
	err := validateURL(url)
	if err != nil {
		log.Fatal("Error Invalid URL")
	}
	details := extractDetailsFromURL(url)
	username := details[0]
	repoName := details[1]
	branch := details[2]
	repo := newRepo(username, repoName, branch)
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

func extractDetailsFromURL(url string) []string {
	regex := `(?:https?:\/\/)?github\.com\/([^\s\/]+)\/([^\s\/]+)(?:\/tree\/?)?([^\s\/]+)?\/?`
	compiledRegex := regexp.MustCompile(regex)
	var result []string
	for _, i := range compiledRegex.FindStringSubmatch(url)[1:] {
		if i != "" && i != "/" {
			result = append(result, i)
		} else {
			result = append(result, "master")
		}
	}
	return result
}

func validateURL(url string) error {
	regex := `(?:https?:\/\/)?github\.com\/([^\s\/]+)\/([^\s\/]+)(?:\/tree\/?)?([^\s\/]+)?\/?`
	compiledRegex := regexp.MustCompile(regex)
	if !compiledRegex.MatchString(url) {
		return errors.New("Invalid URL")
	}
	return nil
}
