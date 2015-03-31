package config

import(
	"os"
)

var (
	APIConsumerKey = os.Getenv("BW_CONSUMER_KEY")
	APIConsumerSecret = os.Getenv("BW_CONSUMER_SECRET")
	APIAccessToken = os.Getenv("BW_ACCESS_TOKEN")
	APIAccessSecret = os.Getenv("BW_ACCESS_SECRET")

	DataDir = os.Getenv("BW_DATA_DIR")
)