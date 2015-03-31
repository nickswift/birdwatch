package tasks

// establish a mapping of keys to function pointers (basically a jump-table)
var (
	Tasks = map[string]Task{
		"":       nil,
		"twitapi"       : twitterAPIInfo,
		"search"        : search,
		"following"     : followingNames,
		"following_ids" : followingIds,
		"history"       : historySearch,
	}
)