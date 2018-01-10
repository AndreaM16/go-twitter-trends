package twitter

import (
	"github.com/ChimeraCoder/anaconda"
	"fmt"
)

func SearchByText(api *anaconda.TwitterApi, text string) {
	searchResult, _ := api.GetSearch(text, nil)
	for _ , tweet := range searchResult.Statuses {
		fmt.Print(tweet.Text)
	}
}
