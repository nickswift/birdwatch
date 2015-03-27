package tasks

/*
 * A task is a function which takes a variadic number of string arguments, does some task and returns a list of strings
 * representing the result of that task. Tasks are supposed to be called and acted upon by an Action.
 */
type Task func(args ...string) []string
