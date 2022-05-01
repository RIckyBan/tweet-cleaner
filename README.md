# tweet-cleaner
Script to delete tweets with Golang

## Prerequisites

1. Setup twitter api secrets

	First, publish your credentials from [twitter developer potal](https://developer.twitter.com/en). Add your twitter api keys to the `.env` file.

	```
	CONSUMER_KEY=
	CONSUMER_SECRET=
	ACCESS_TOKEN=
	ACCESS_SECRET=
	```

2. Prepare your archived twitter data

	Download your archived twitter data from [here](https://twitter.com/settings/your_twitter_data).
	Then place `tweeet.js` file inside to the project root.

3. Format your twitter data into json

	Run the command below.

	```shell
	$ sed '1d' tweet.js | sed '1i [' > tweet.json
	```


## How to use

```shell
$ go run main.go --from 2022-01-01 --to 2022-05-01
```

:warning: Please use this script at your own risk.