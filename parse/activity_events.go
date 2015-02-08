package parse

import (
	"fmt"
	"github.com/stormasm/go-github/github"
)

func Listevents(events []github.Event) {
	fmt.Printf("Number of events = %v\n", len(events))

	for i, event := range events {
		//fmt.Printf("\n")
		fmt.Printf("[%d] %s  ", i, *event.Type)
		EventSwitch(event)
	}
}

/*
func Listevents(events []github.Event) {
	fmt.Printf("Number of events = %v\n", len(events))

	for i, event := range events {
		fmt.Printf("\n")
		fmt.Printf("[%d] %s  ", i, *event.Type)
		fmt.Printf("%s", *event.Repo.Name)
		fmt.Printf("\n")
		//payload(event)
	}
}
*/

func payload(event github.Event) {
	fmt.Println("\n")
	//mystr := event.String()
	//fmt.Println(mystr)
	myrep := event.Payload()
	fmt.Println(myrep)
}
