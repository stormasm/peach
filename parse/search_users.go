package parse

import (
	"fmt"
	"github.com/google/go-github/github"
)

func Search_users_results(results *github.UsersSearchResult) {
	fmt.Printf("Number of users = %d\n", *results.Total)

	for i, user := range results.Users {
		//fmt.Printf("\n")
		fmt.Printf("[%d] %s  ", i, *user.Login)
	}
	fmt.Println()
}
