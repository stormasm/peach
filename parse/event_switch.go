package parse

import (
	"fmt"
	"github.com/stormasm/go-github/github"
	"strings"
)

func EventSwitch(event github.Event) {
	eventype := *event.Type

	switch true {
	case strings.EqualFold(eventype, "WatchEvent"):
		fmt.Println("Hit on watch event")
	case strings.EqualFold(eventype, "PushEvent"):
		fmt.Println("Hit on push event")
	default:
		fmt.Printf("\n")
	}
}
