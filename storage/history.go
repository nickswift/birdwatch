package storage

import(
	"github.com/nickswift/birdwatch/config"
	"encoding/csv"
	"path/filepath"
	"errors"
	"os"
)

var (
	HistoryFilePath = filepath.Join(config.DataDir, "tweets.csv")
)

type Tweet struct {
	TweetId                  string
	InReplyToStatusId        string
	InReplyToUserId          string
	Timestamp                string
	Source                   string
	Text                     string
	RetweetedStatusId        string
	RetweetedStatusUserId    string
	RetweetedStatusTimestamp string
	ExpandedUrls             string
}

// TODO: also cache things like user friends, etc...
type History struct {
	UserTweets []*Tweet
}

func historyFileExists() bool {
	_, err := os.Stat(HistoryFilePath)
	return !os.IsNotExist(err)
}

func ReadHistory() (*History, error) {
	if !historyFileExists() {
		return nil, errors.New("history file does not exist")
	}

	var tweets []*Tweet

	// open CSV history file
	file, err := os.Open(HistoryFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read from file
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 10

	raw, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	// build history
	for _, t := range raw {
		tweets = append(tweets, &Tweet{t[0],t[1],t[2],t[3],t[4],t[5],t[6],t[7],t[8],t[9]})
	}

	return &History{
		UserTweets: tweets,
	}, nil
}