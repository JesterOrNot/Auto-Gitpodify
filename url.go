package main

import (
	"regexp"
	"errors"
)

func ExtractDetailsFromURL(url string) []string {
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

func ValidateURL(url string) error {
	regex := `(?:https?:\/\/)?github\.com\/([^\s\/]+)\/([^\s\/]+)(?:\/tree\/?)?([^\s\/]+)?\/?`
	compiledRegex := regexp.MustCompile(regex)
	if !compiledRegex.MatchString(url) {
		return errors.New("Invalid URL")
	}
	return nil
}
