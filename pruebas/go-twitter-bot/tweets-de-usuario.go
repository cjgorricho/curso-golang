package main

import (
	// other imports
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

// Credentials stores all of our access/consumer tokens
// and secret keys needed for authentication against
// the twitter REST API.
type Credentials struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

// getClient is a helper function that will return a twitter client
// that we can subsequently use to send tweets, or to stream new tweets
// this will take in a pointer to a Credential struct which will contain
// everything needed to authenticate and return a pointer to a twitter Client
// or an error
func getClient(creds *Credentials) (*twitter.Client, error) {
	// Pass in your consumer key (API Key) and your Consumer Secret (API Secret)
	config := oauth1.NewConfig(creds.ConsumerKey, creds.ConsumerSecret)
	// Pass in your Access Token and your Access Token Secret
	token := oauth1.NewToken(creds.AccessToken, creds.AccessTokenSecret)

	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	// Verify Credentials
	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}

	// we can retrieve the user and verify if the credentials
	// we have used successfully allow us to log in!
	user, _, err := client.Accounts.VerifyCredentials(verifyParams)
	if err != nil {
		return nil, err
	}

	log.Printf("User's ACCOUNT:\n%+v\n", user)
	return client, nil
}

func main() {
	fmt.Println("Go-Twitter Bot v0.01")
	creds := Credentials{
		AccessToken:       os.Getenv("ACCESS_TOKEN"),
		AccessTokenSecret: os.Getenv("ACCESS_TOKEN_SECRET"),
		ConsumerKey:       os.Getenv("CONSUMER_KEY"),
		ConsumerSecret:    os.Getenv("CONSUMER_SECRET"),
	}

	log.Printf("%+v\n", creds)

	client, err := getClient(&creds)
	if err != nil {
		log.Println("Error getting Twitter Client")
		log.Println(err)
	}

	tweets, resp, err := client.Timelines.UserTimeline(&twitter.UserTimelineParams{
		ScreenName: "alvarouribevel",
		Count:      200,
	})

	if err != nil {
		log.Println(err)
	}
	log.Printf("Largo del Header: %+v\n", len(resp.Header))
	log.Printf("Numero de tweets: %+v\n", len(tweets))

	fmt.Println("Rate limit remaining:", resp.Header["X-Rate-Limit-Remaining"][0])
	ratereset, err := strconv.ParseInt(resp.Header["X-Rate-Limit-Reset"][0], 10, 64)
	if err != nil {
		log.Println(err)
	}
	ratetime := time.Unix(ratereset, 0)
	fmt.Printf("Rate limit reset: %v\n", ratetime)

	fmt.Println("App limit remaining:", resp.Header["X-App-Rate-Limit-Remaining"][0])
	appreset, err := strconv.ParseInt(resp.Header["X-App-Rate-Limit-Reset"][0], 10, 64)
	if err != nil {
		log.Println(err)
	}
	apptime := time.Unix(appreset, 0)
	fmt.Printf("Rate limit reset: %v\n", apptime)

	for i, m := range tweets {
		fmt.Printf("Tuit %v: @%+v %+v %+v %v\n", i, m.User.ScreenName, m.CreatedAt, m.ID, m.Text)
	}

}
