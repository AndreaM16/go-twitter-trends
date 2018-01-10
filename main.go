package main

import (
	"github.com/andream16/go-twitter-trends/configuration"
	"github.com/andream16/go-twitter-trends/twitter"
)

func main() {
	conf := configuration.InitConfiguration()
	api := twitter.InitAnaconda(conf)
	twitter.SearchByText(api, "golang")
}
