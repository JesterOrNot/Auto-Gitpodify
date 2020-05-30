package main

import (
	"io/ioutil"
	"encoding/json"
	"net/http"
	"fmt"
	"log"
)

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
