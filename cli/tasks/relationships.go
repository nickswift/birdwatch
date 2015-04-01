package tasks

import(
	"github.com/nickswift/birdwatch/client"
	"github.com/chimeracoder/anaconda"
	"net/url"
	"strings"
	"sync"
	"fmt"
)

const(
	MaxFollowersIdsCount = 5000
)

// TODO: handle errors more elegantly
// retrieve a list of IDs of followers of a specified screen name
func followingUsers(tc *client.TwitterClient, args ...string) []anaconda.User {
	var users []anaconda.User
	var pages chan anaconda.FollowersPage
	var err error

	// first, we batch the arguments together to keep ourselves from making too many queries. Then we ask twitter for
	// their IDs
	if len(args) > 100 {
		fmt.Println("Too many screen names specified. Max is 100.")
		return nil
	}
	names := strings.Join(args, `,`)

	v := url.Values{}
	if users, err = tc.Api.GetUsersLookup(names, v); err != nil {
		fmt.Println("ERR: %s\n", err.Error())
		return nil
	}

	// set up wait group
	var wg sync.WaitGroup

	// Now, we use these screen names to generate a list of users following them. We need to prevent duplicates as we do 
	// this.
	for _, user := range users {
		// set url arguments
		v = url.Values{}
		v.Set("user_id", fmt.Sprintf("%d", user.Id))
		v.Set("count", fmt.Sprintf("%d", MaxFollowersIdsCount))

		// do API call
		if pages = tc.Api.GetFollowersListAll(v); err != nil {
			fmt.Println("ERR: %s\n", err.Error())
			return nil
		}

		// concurrently add users to list while we query for the next user
		for page := range pages{
			wg.Add(1)
			go func(){
				for _, user := range page.Followers {
					users = append(users, user)
				}
				wg.Done()
			}()
		}

	}
	wg.Wait()

	return users
}

// extract IDs from following data
func followingIds(tc *client.TwitterClient, args ...string) []string {
	// invoke ur-task
	var res []string
	users := followingUsers(tc, args...)

	if users == nil {
		return []string{}
	}

	for _, user := range users {
		res = append(res, fmt.Sprintf("%d", user.Id))
	}

	return res
}

// extract names from following data
func followingNames(tc *client.TwitterClient, args ...string) []string {
	// invoke ur-task
	var res []string
	users := followingUsers(tc, args...)

	if users == nil {
		return []string{}
	}

	for _, user := range users {
		res = append(res, fmt.Sprintf("%s", user.ScreenName))
	}

	return res
}