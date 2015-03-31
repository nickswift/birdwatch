package actions

// establish a mapping of keys to function pointers (basically a jump-table)
var (
	Actions = map[string]actions.Action{
		""         : nil,
		"print"    : ActionPrint,
		"unfollow" : ActionUnfollow,
	}
)
