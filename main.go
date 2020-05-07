package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"sync"

	"github.com/urfave/cli"
)

var app = cli.NewApp()
var waitGroup sync.WaitGroup = sync.WaitGroup{}
var Languages []string

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
	if validateURL(url) != nil {
		log.Fatal("Error Invalid URL")
	}
	details := extractDetailsFromURL(url)
	username := details[0]
	repoName := details[1]
	branch := details[2]
	repo := newRepo(username, repoName, branch)
	files := GetFiles(repo)
	checkAll(files)
	for i := 0; i < len(files); i++ {
		fmt.Println(files[i])
	}
	fmt.Println(Languages)
}

func validateArgs(c *cli.Context) error {
	if len(c.Args().Tail()) > 0 || !c.Args().Present() {
		return errors.New("Error: Invalid Arguments")
	}
	return nil
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

func checkAll(files []string) {
	waitGroup.Add(3)
	go checkGo(files)
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
			Languages = append(Languages, "yarn")
		}
	}
}

func checkGo(files []string) {
	defer waitGroup.Done()
	for _, file := range files {
		if file == "go.mod" {
			Languages = append(Languages, "go")
		}
	}
}

type Response struct {
	Tree []GitFile `json:"tree"`
}

type GitFile struct {
	Name string `json:"path"`
	Type string `json:"type"`
}

type Repository struct {
	User     string
	RepoName string
	Branch   string
}

func newRepo(user, repoName, branch string) Repository {
	return Repository{
		User:     user,
		RepoName: repoName,
		Branch:   branch,
	}
}
