package actions

import (
	"github.com/nickswift498/birdwatch/client"
	"github.com/nickswift498/birdwatch/cli/tasks"
	"os"
	"bufio"
	"fmt"
)

const(
	ActionCancelledMsg = "\naction cancelled\n"
)

/*
 * An Action is a higher-order function representing some kind of thing one is able to do to the result of a task. It
 * takes a function pointer and a variadic number of string arguments. It invokes the given function with said
 * arguments and performs some task upon the output of the task.
 */
type Action func(tc *client.TwitterClient, task tasks.Task, args ...string)

// dangerous action confirmation -- get confirmation of intent from the user before doing something that could be 
// stupid/harmful
func ConfirmActionIntent(reason string) bool {
	fmt.Println("WARNING: this action will make changes to your Twitter account."
	fmt.Println("This may do things you don't want or did not intend: ")
	fmt.Println(reason)
	fmt.Printf("\nAre you *sure* you intend to do this? (y/n) [n] > ")

	cmd, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	cmd = cmd[:len(cmd)-1]
	if cmd == "y" || cmd == "Y" {
		return true
	}
	return false
}