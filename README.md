---
title: Publish Tweet
weight: 1
---

# Publish Tweet
This activity allows you to Tweet, Retweet a particular tweet you like,Follow/UnFollow user on Twitter,,Block/Unblock user on Twitter,DirectMessage a user on Twitter,GetUserTimeline, get your timeline,GetTrendsByPlace for a particular woeid(where on earth id).This activity is developed by FLOGO AllStars team.
## Installation
### Flogo CLI
```bash
flogo install github.com/DipeshTest/publishtweet
```

### Third-party libraries used
- #### package anaconda - "github.com/ChimeraCoder/anaconda":
	Anaconda is a simple, transparent Go package for accessing version 1.1 of the Twitter API.Successful API queries return native Go structs that can be used immediately, with no need for type assertions
## Schema
Inputs and Outputs:

```json
 "inputs": [{
		"name": "consumerKey",
		"type": "string",
		"required": true
	},
	{
		"name": "consumerSecret",
		"type": "string",
		"required": true
	},
	{
		"name": "accessToken",
		"type": "string",
		"required": true
	},
	{
		"name": "accessTokenSecret",
		"type": "string",
		"required": true
	},
	{
		"name": "twitterFunction",
		"type": "string",
		"allowed": ["Tweet",
		"ReTweet",
		"Block",
		"Unblock",
		"Follow",
		"Unfollow",
		"DM"],
		"value": "Tweet",
		"required": true
	},
	{
		"name": "user",
		"type": "string"
	},
	{
		"name": "text",
		"type": "string"
	}],
	"outputs": [{
		"name": "statusCode",
		"type": "string"
	},
	{
		"name": "message",
		"type": "any"
	}]
```
## Settings
| Setting     | Required | Description |
|:------------|:---------|:------------|
| consumerKey | True     | The consumerKey of your Twitter account |         
| consumerSecret   | True    | The consumerSecret of your Twitter account|
| accessToken       | True    | The accessToken of your Twitter account |
| accessTokenSecret   | True    | The accessTokenSecret of your Twitter account|
| twitterFunction   | True    | Select the action you want to perform using the PublishTweet connector, the possible values are "Tweet","ReTweet","Block","Unblock","Follow","Unfollow","DM"|
| user   | False    | Use this field to provide tweetId for ReTweet, userHandle for Block,UnBlock,Follow,UnFollow,DirectMessage(DM)|
|text|False| Use this field to provide text for you Tweet,message for you DM|

Note: You may use below URL to generate your Twitter tokens: https://www.slickremix.com/docs/how-to-get-api-keys-and-tokens-for-twitter/
## Examples
Please refer activity_test.go 


## Response Codes
### Google Drive Create
| ResponseCode     | Type | Description |
|:------------|:---------|:------------|
|200 |OK| The request operation was successful.|
|101 |INVALID INPUT| Consumer Key field is blank.|
|102 |INVALID INPUT| Consumer Secret field is blank.|
|103 |INVALID INPUT| Access Token field is blank.|
|104 |INVALID INPUT| Access Token Secret field is blank.|
|105 |AUTHENTICATION ERROR| User Field is blank for Tweet or User field is blank for TweetId/Block user/UnBlock user/Follow user/Unfollow user.|
Refer : https://developer.twitter.com/en/docs/basics/response-codes for additional Response codes
