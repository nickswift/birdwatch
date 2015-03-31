package tasks

const(
	MaxFollowersIdsCount = 5000
)

// retrieve a list of IDs of followers of a specified screen name
func TaskFollowers(tc *client.TwitterClient, args ...string) []string {
	var users []anaconda.User
	var cursor anaconda.Cursor
	var err error
	res := []string{}

	// first, we batch the arguments together to keep ourselves from making too many queries. Then we ask twitter for
	// their IDs
	if len(args) > 100 {
		fmt.Println("Too many screen names specified. Max is 100.")
		return res
	}
	names := strings.Join(args, `,`)

	v := url.Values{}
	if users, err = tc.Api.GetUsersLookup(names, v); err != nil {
		fmt.Println("ERR: %s\n", err.Error())
		return res
	}

	// Now, we use these screen names to generate a list of users following them. We need to prevent duplicates as we do 
	// this.
	for _, user := range users {
		// set url arguments
		v = url.Values{}
		v.Set("user_id", fmt.Sprintf("%d", user.Id))
		v.Set("count", fmt.Sprintf("%d", MaxFollowersIdsCount))

		// do API call
		if cursor, err = tc.Api.GetFollowersIds(v); err != nil {
			fmt.Println("ERR: %s\n", err.Error())
			return res
		}

		// add users
		for _, i := range cursor.Ids {
			res = append(res, fmt.Sprintf("%d", i))
		}
	}

	return res
}