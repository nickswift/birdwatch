package actions

import (
	"github.com/nickswift498/birdwatch/cli/tasks"
)

/*
 * An Action is a higher-order function representing some kind of thing one is able to do to the result of a task. It
 * takes a function pointer and a variadic number of string arguments. It invokes the given function with said
 * arguments and performs some task upon the output of the task.
 */
type Action func(task tasks.Task, args ...string)
