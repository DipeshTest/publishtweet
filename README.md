---
title: Publish Tweet
weight: 1
---

# Publish Tweet
This activity allows you to Tweet, Tweet with Media, Retweet a particular tweet you like, Follow/UnFollow user on Twitter, Block/Unblock user on Twitter, DirectMessage a user on Twitter, GetUserTimeline, get your timeline, GetTrendsByPlace for a particular woeid (where on earth id). This activity is developed by FLOGO AllStars team.

## Installation
### Flogo CLI
```bash
flogo install github.com/DipeshTest/publishtweet
```
### TIBCO Cloud Integration
upload via WebUI

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
		"TweetMedia",
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
	},
	{
		"name": "mediaURL",
		"type": "string"
	}],
	"outputs": [{
		"name": "statusCode",
		"type": "integer"
	},
	{
		"name": "message",
		"type": "any"
	}]
```
## Settings
| Setting           | Required  | Description                         |
|:------------------|:----------|:------------------------------------|
| consumerKey       | true      | consumerKey of your Twitter account |         
| consumerSecret    | true      | consumerSecret of your Twitter account |
| accessToken       | true      | accessToken of your Twitter account |
| accessTokenSecret | true      | accessTokenSecret of your Twitter account |
| Function          | true      | Actions you want to perform, the possible values are "Tweet", "TweetMedia","ReTweet","Block","Unblock","Follow","Unfollow","DirectMessage" |
| User              | false     | Use this field to provide tweetId for ReTweet, User for Block, UnBlock, Follow, UnFollow, DirectMessage |
| Text              | false     | Use this field to provide an text for you Tweet or Direct Message |
| MediaURL          | false     | Use this field to provide an external Image Media URL |

Note: You may use below URL to generate your Twitter tokens: https://www.slickremix.com/docs/how-to-get-api-keys-and-tokens-for-twitter/

## Examples & Tester
Please refer activity_test.go 

Tester read secure Test Properties, to store them outside any GitHub
feel free to adjust Path and enter your Keys. Tokens and Secrets
File Content Should look like follows:
- consumerKey=[your consumerKey]
- consumerSecret=[your consumerSecret]
- accessToken=[your accessToken]
- accessTokenSecret=[your accessTokenSecret]

Note: path need to be adjusted to your Properties File:
- c:\\GODev\\twitterApp.properties

## Response Codes
### Publish Tweet
| ResponseCode     | Type | Description |
|:------------|:---------|:------------|
|200 |OK| The request operation was successful.|
|101 |INVALID INPUT| Consumer Key field is blank.|
|102 |INVALID INPUT| Consumer Secret field is blank.|
|103 |INVALID INPUT| Access Token field is blank.|
|104 |INVALID INPUT| Access Token Secret field is blank.|
|105 |AUTHENTICATION ERROR| User Field is blank for Tweet or User field is blank for TweetId/Block user/UnBlock user/Follow user/Unfollow user.|

#Refer : https://developer.twitter.com/en/docs/basics/response-codes for additional Response codes
