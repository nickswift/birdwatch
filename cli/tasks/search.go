package tasks

import(
	"github.com/nickswift498/birdwatch/client"
	"net/url"
	"fmt"
)

func search(tc *client.TwitterClient, args ...string) []string {
	res := []string{}
	var v = url.Values{}
	var lim string

	if len(args) == 0 {
		fmt.Printf("needs arguments")
		return res
	}

	// default to 10-length search result limit
	if len(args) < 2 {
		lim = "10"
	} else {
		// set limit from last element of arg string
		lim = args[len(args)-1]
		// pop that argument off the end
		args = args[:len(args)-1]
	}
	v.Set("count", lim)

	for _, arg := range args {
		res = append(res, fmt.Sprintf("RESULTS FOR %s", arg))
		result, _ := tc.Api.GetSearch(arg, v)
		for _, tweet := range result.Statuses {
			res = append(res, fmt.Sprintf("%15s said: %s", fmt.Sprintf("@%s",tweet.User.ScreenName), tweet.Text))
		}
	}

	return res
}