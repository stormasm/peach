package parse

import (
	"fmt"
	"github.com/google/go-github/github"
)

func Listevents(events []github.Event) {
	fmt.Printf("Number of events = %v\n", len(events))

	for i, v := range events {
		fmt.Printf("[%d] %s  ", i, *v.Type)
		fmt.Printf("%s", *v.Repo.Name)
		fmt.Printf("\n")
	}
}
