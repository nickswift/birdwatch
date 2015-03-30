package client

import (
	// "time"
	"github.com/chimeracoder/anaconda"
)

type Consumer struct {
	Key    string
	Secret string
}

func NewConsumer(key, secret string) *Consumer {
	return &Consumer{key, secret}
}

type Access struct {
	Token  string
	Secret string
}

func NewAccess(token, secret string) *Access {
	return &Access{token, secret}
}

type TwitterClient struct {
	Consumer *Consumer
	Access   *Access
	Api      *anaconda.TwitterApi
}

func NewTwitterClient(ckey, csecret, atoken, asecret string) *TwitterClient {
	anaconda.SetConsumerKey(ckey)
	anaconda.SetConsumerSecret(csecret)
	api := anaconda.NewTwitterApi(atoken, asecret)

	// set rate-limiting stuff -- right now, it's at 180 queries per 15 minutes
	// api.SetDelay(5 * time.Second)

	return &TwitterClient{
		Consumer: NewConsumer(ckey, csecret),
		Access:   NewAccess(atoken, asecret),
		Api:      api,
	}
}

// Our main use-case is API info stored as environment variables
// TODO: allow this to work with command line arguments at runtime
func TwitterClientFromEnv() *TwitterClient {
	consumerKey := os.Getenv("BW_CONSUMER_KEY")
	consumerSecret := os.Getenv("BW_CONSUMER_SECRET")
	accessToken := os.Getenv("BW_ACCESS_TOKEN")
	accessSecret := os.Getenv("BW_ACCESS_SECRET")

	return NewTwitterClient(consumerKey, consumerSecret, accessToken, accessSecret)
}

