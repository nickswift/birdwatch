package actions

import (
	"github.com/nickswift498/birdwatch/cli/actions"
)

// establish a mapping of keys to function pointers (basically a jump-table)
var (
	Actions = map[string]actions.Action{
		""         : nil,
		"print"    : actions.ActionPrint,
		"unfollow" : ActionUnfollow,
	}
)
