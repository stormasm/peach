package main

import (
	"code.google.com/p/goauth2/oauth"
	"fmt"
	"github.com/stormasm/go-github/github"
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
	opts := &github.SearchOptions{Sort: "forks", Order: "desc",
		ListOptions: github.ListOptions{Page: 1, PerPage: 100}}
	result, _, err := client.Search.Users("location:portland,or language:go", opts)
	if err != nil {
		fmt.Printf("Search.Issues returned error: %v", err)
	}

	//fmt.Println(result)

	parse.Search_users_results(result)

	rate, _, err := client.RateLimits()
	if err != nil {
		fmt.Printf("Error fetching rate limit: %#v\n\n", err)
	} else {
		fmt.Println("Core   Remaining = ", rate.Core.Remaining)
		fmt.Println("Search Remaining = ", rate.Search.Remaining)
	}
}
