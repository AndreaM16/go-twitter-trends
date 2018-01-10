package twitter

import (
	"github.com/andream16/go-twitter-trends/configuration"
	"github.com/ChimeraCoder/anaconda"
)

func InitAnaconda(conf configuration.Configuration) *anaconda.TwitterApi {
	anaconda.SetConsumerKey(conf.Credentials.Consumer.Key)
	anaconda.SetConsumerSecret(conf.Credentials.Consumer.Secret)
	api := anaconda.NewTwitterApi(conf.Credentials.Access.Token, conf.Credentials.Access.Secret)
	return api
}