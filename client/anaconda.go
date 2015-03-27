package client

import (
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

func NewTwitterClient(ckey, ssecret, atoken, asecret string) *TwitterClient {
	anaconda.SetConsumerKey(ckey)
	anaconda.SetConsumerSecret(ssecret)
	api := anaconda.NewTwitterApi(atoken, asecret)

	return &TwitterClient{
		Consumer: NewConsumer(ckey, ssecret),
		Access:   NewAccess(atoken, asecret),
		Api:      api,
	}
}
