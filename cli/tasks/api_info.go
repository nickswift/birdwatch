package tasks

import(
	"github.com/nickswift/birdwatch/client"
	"fmt"
)

// find the currently selected user's information
func twitterAPIInfo(tc *client.TwitterClient, args ...string) []string {
	var ck, cs, at, as string
	strNSet := "not set"

	if tc.Consumer.Key == "" {
		ck = strNSet
	} else {
		ck = tc.Consumer.Key
	}
	if tc.Consumer.Secret == "" {
		cs = strNSet
	} else {
		cs = tc.Consumer.Secret
	}
	if tc.Access.Token == "" {
		at = strNSet
	} else {
		at = tc.Access.Token
	}
	if tc.Access.Secret == "" {
		as = strNSet
	} else {
		as = tc.Access.Secret
	}

	return []string{
		fmt.Sprintf("\nTwitter API Information"),
		fmt.Sprintf("CONSUMER KEY    : %s", ck),
		fmt.Sprintf("CONSUMER SECRET : %s", cs),
		fmt.Sprintf("ACCESS TOKEN    : %s", at),
		fmt.Sprintf("ACCESS SECRET   : %s", as),
		fmt.Sprintf("\nIf this information doesn't look right, take a look at the relevant environment variables."),
		fmt.Sprintf("If the previous sentence didn't make sense, please consult the README.\n"),
	}
}