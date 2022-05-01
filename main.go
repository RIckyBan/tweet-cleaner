package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/joho/godotenv"
)

type TweetData struct {
	ID        string `json:"id"`
	FullText  string `json:"full_text"`
	CreatedAt string `json:"created_at"`
}

type Tweet struct {
	Tweet TweetData `json:"tweet"`
}

func loadJson(filePath string) []Tweet {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	var tweets []Tweet
	if err := json.Unmarshal(bytes, &tweets); err != nil {
		log.Fatal(err)
	}
	return tweets
}

func loadSecrets() (string, string, string, string) {
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

func deleteTweet(client *twitter.Client, id int64) {
	// Delete a Tweet
	log.Println("Deleting Tweet with ID:", id)
	params := &twitter.StatusDestroyParams{TrimUser: twitter.Bool(true)}
	_, resp, err := client.Statuses.Destroy(id, params)
	if err != nil {
		log.Println(err)
	}
	if resp.StatusCode == 200 {
		log.Println("Successfully deleted")
	} else {
		log.Println("Failed to delete")
	}
}

func main() {
	fromDateString := flag.String("from", "", "Date to start from")
	toDateString := flag.String("to", "", "Date to end at")
	flag.Parse()

	if *fromDateString == "" || *toDateString == "" {
		log.Fatal("Please provide a date range")
	}

	layout := "2006-01-02"
	fromDate, err := time.Parse(layout, *fromDateString)
	if err != nil {
		log.Fatal(err)
	}
	toDate, err := time.Parse(layout, *toDateString)
	if err != nil {
		log.Fatal(err)
	}

	if fromDate.After(toDate) {
		log.Fatal("From-date must be before to-date")
	}

	log.Println("Fetching tweets from ", fromDate.Format(layout), " to ", toDate.Format(layout))

	ck, cs, at, as := loadSecrets()
	config := oauth1.NewConfig(ck, cs)
	token := oauth1.NewToken(at, as)
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	tweets := loadJson("tweet.json")

	// filter tweets
	var deleteIDs []int64
	layout = "Mon Jan 02 15:04:05 -0700 2006"
	for _, tw := range tweets {
		tweetDate, _ := time.Parse(layout, tw.Tweet.CreatedAt)
		if tweetDate.After(fromDate) && tweetDate.Before(toDate) {
			id, _ := strconv.ParseInt(tw.Tweet.ID, 10, 64)
			deleteIDs = append(deleteIDs, id)
		}
	}

	log.Printf("Deleting %d tweets", len(deleteIDs))

	// Delete tweets
	for _, id := range deleteIDs {
		deleteTweet(client, id)
		// sleep for 300ms
		time.Sleep(300 * time.Millisecond)
	}
}
