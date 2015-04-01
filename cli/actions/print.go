package actions

import (
	"fmt"
	"github.com/nickswift/birdwatch/client"
	"github.com/nickswift/birdwatch/cli/tasks"
)

// the Print action simply prints the output of a task so the user can read it.
func ActionPrint(tc *client.TwitterClient, task tasks.Task, args ...string) {
	tr := task(tc, args...)
	for _, r := range tr {
		fmt.Printf("\n%s", r)
	}
	fmt.Printf("\n")
}
