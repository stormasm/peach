package parse

import (
	"fmt"
	"github.com/google/go-github/github"
)

func Search_repos_results(results *github.RepositoriesSearchResult) {
	fmt.Printf("Number of repositories = %d\n", results.Total)

	for i, repository := range results.Repositories {
		//fmt.Printf("\n")
		fmt.Printf("[%d] %s  ", i, *repository.Name)
	}
	fmt.Println()
}
