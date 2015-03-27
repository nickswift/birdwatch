package actions

import (
	"fmt"
	"github.com/nickswift498/birdwatch/cli/tasks"
)

// the Print action simply prints the output of a task so the user can read it.
func ActionPrint(task tasks.Task, args ...string) {
	tr := task(args...)
	for _, r := range tr {
		fmt.Printf("%s ", r)
	}
	fmt.Printf("\n")
}
