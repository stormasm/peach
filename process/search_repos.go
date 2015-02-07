package main

import (
	"code.google.com/p/goauth2/oauth"
	"fmt"
	"github.com/google/go-github/github"
	"github.com/stormasm/peach/parse"
	"os"
)

var (
	client *github.Client

	// auth indicates whether tests are being run with an OAuth token.
	// Tests can use this flag to skip certain tests when run without auth.
	auth bool
)

func init() {
	token := os.Getenv("GITHUB_AUTH_TOKEN")
	if token == "" {
		print("!!! No OAuth token.  Some tests won't run. !!!\n\n")
		client = github.NewClient(nil)
	} else {
		t := &oauth.Transport{
			Token: &oauth.Token{AccessToken: token},
		}
		client = github.NewClient(t.Client())
		auth = true
	}
}

func main() {
/*
	opt := &github.ListOptions{Page: 2, PerPage: 100}
	events, _, err := client.Activity.ListEvents(opt)
	if err != nil {
		fmt.Println("Activities.ListEvents returned error: %v", err)
	}
*/
	opts := &github.SearchOptions{Sort: "forks", Order: "desc", ListOptions: github.ListOptions{Page: 1, PerPage: 10}}
	result, _, err := client.Search.Repositories("jekyll-poole-notes", opts)
	if err != nil {
		fmt.Println("Search.Repositories returned error: %v", err)
	}

	parse.Search_repos_results(result)
	//fmt.Println(result)

	rate, _, err := client.RateLimits()
	if err != nil {
		fmt.Printf("Error fetching rate limit: %#v\n\n", err)
	} else {
		fmt.Println("Core   Remaining = ",rate.Core.Remaining)
		fmt.Println("Search Remaining = ", rate.Search.Remaining)
		//fmt.Printf("Core API Rate Limit: %d\n", rate.core.Remaining)
		//fmt.Printf("Search API Rate Limit: %d\n", rate.search.Remaining)
	}

}
