package parse

import (
	"fmt"
	"github.com/google/go-github/github"
)

func Listevents(events []github.Event) {
	fmt.Printf("Number of events = %v\n", len(events))

	for i, v := range events {
		fmt.Printf("array value at [%d]=%s\n", i, *v.Type)
	}
}
