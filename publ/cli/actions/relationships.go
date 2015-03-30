package actions

import(
	"github.com/nickswift498/birdwatch/client"
	"github.com/nickswift498/birdwatch/cli/actions"
	"github.com/nickswift498/birdwatch/cli/tasks"
	"fmt"
)

const (
	UnfollowDangerReason = `
		You're about to unfollow a large number of users. You're encouraged to consider why you're doing this, as this may
		greatly alter the composition and tone of the posts that appear in your feed.
`
)

/*
 * Following actions
 * these are actions which manage the user's follow status with regard to other users.
 */
// func ActionFollowUsers(tc *client.TwitterClient, tasks.Task, args ...string) {
// }

func ActionUnfollow(tc *client.TwitterClient, task tasks.Task, args ...string) {
	// TODO: if there's no cached list of users we're following, generate and save it.

	// assume the given task returns an actionable list of user Ids
	tr := task(tc, args...)

	if !actions.ConfirmActionIntent(UnfollowDangerReason) {
		fmt.Printf(actions.ActionCancelledMsg)
		return
	}

	for _, id := range tr {
		fmt.Printf("\nUNFOLLOWING: %s", id)
	}
	fmt.Println("")
}