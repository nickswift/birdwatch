package tasks

// establish a mapping of keys to function pointers (basically a jump-table)
var (
	Tasks = map[string]tasks.Task{
		"":       nil,
		"twitapi"   : TaskTwitterAPIInfo,
		"search"    : TaskSearch,
		"following" : TaskFollowers,
	}
)