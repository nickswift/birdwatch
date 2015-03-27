package tasks

import (
	"github.com/nickswift498/birdwatch/cli/tasks"
)

// establish a mapping of keys to function pointers (basically a jump-table)
var (
	Tasks = map[string]tasks.Task{
		"":       nil,
		"simple": TaskSimple,
	}
)