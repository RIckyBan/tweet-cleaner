package main

import (
	"log"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/joho/godotenv"
)

func loadDotenv() (string, string, string, string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ConsumerKey := os.Getenv("CONSUMER_KEY")
	ConsumerSecret := os.Getenv("CONSUMER_SECRET")
	AccessToken := os.Getenv("ACCESS_TOKEN")
	AccessSecret := os.Getenv("ACCESS_SECRET")

	return ConsumerKey, ConsumerSecret, AccessToken, AccessSecret
}

func main() {
	ck, cs, at, as := loadDotenv()
	config := oauth1.NewConfig(ck, cs)
	token := oauth1.NewToken(at, as)
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	// Home Timeline
	tweets, _, err := client.Timelines.HomeTimeline(&twitter.HomeTimelineParams{
		Count: 20,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(tweets)
}
