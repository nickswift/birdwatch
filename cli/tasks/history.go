package tasks

import(
	"github.com/nickswift/birdwatch/storage"
	"github.com/nickswift/birdwatch/client"
	"strings"
	"sync"
	"fmt"
)

// find tweets with keywords in the user's history
func historySearch(tc *client.TwitterClient, args ...string) []string {
	var res []string
	results := make(map[string]string)

	history, err := storage.ReadHistory()
	if err != nil {
		return res
	}

	// process all the args concurrently
	var wg sync.WaitGroup
	for _, arg := range args {
		wg.Add(1)
		go func() {
			for _, tweet := range history.UserTweets {
				if strings.Contains(tweet.Text, arg) {
					// insert into dictionary -- avoid dupes
					if _, ok := results[tweet.TweetId]; !ok {
						results[tweet.TweetId] = tweet.Text
					}
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()

	// build final result
	for k, v := range results {
		res = append(res, fmt.Sprintf("%20s : %s", k, v))
	}
	return res
}