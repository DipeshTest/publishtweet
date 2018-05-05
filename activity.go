package publishtweet

import (
	"fmt"
	"strconv"
	s "strings"

	"github.com/JayDShah/TwitterAPI"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

var log = logger.GetLogger("activity-twitterpublish")

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {

	consumerKey := s.TrimSpace(context.GetInput("consumerKey").(string))
	consumerSecret := s.TrimSpace(context.GetInput("consumerSecret").(string))
	accessToken := s.TrimSpace(context.GetInput("accessToken").(string))
	accessTokenSecret := s.TrimSpace(context.GetInput("accessTokenSecret").(string))

	twitterFunction := context.GetInput("twitterFunction").(string)
	fmt.Println("test ", twitterFunction)
	if len(consumerKey) == 0 {

		context.SetOutput("statusCode", "101")

		context.SetOutput("message", "Consumer Key field is blank")
	} else if len(consumerSecret) == 0 {

		context.SetOutput("statusCode", "102")

		context.SetOutput("message", "Consumer Secret field is blank")

	} else if len(accessToken) == 0 {

		context.SetOutput("statusCode", "103")

		context.SetOutput("message", "Access Token field is blank")

	} else if len(accessTokenSecret) == 0 {

		context.SetOutput("statusCode", "104")

		context.SetOutput("message", "Access Token Secret field is blank")

	} else {
		var msg string
		var code int

		switch twitterFunction {
		case "Tweet":
			{
				tweet := s.TrimSpace(context.GetInput("text").(string))
				if len(tweet) == 0 {

					code = 105
					msg = "Tweet cannot be blank"

				} else {
					code, msg = Twitter.PostTweet(consumerKey, consumerSecret, accessToken, accessTokenSecret, tweet)
				}
			}
		case "ReTweet":
			{
				tweetId := context.GetInput("user")
				tid := s.TrimSpace(tweetId.(string))
				if len(tid) == 0 {

					code = 105
					msg = "TweetId cannot be blank"

				} else {
					value, err := strconv.Atoi(tid)

					if err == nil {
						code, msg = Twitter.ReTweet(consumerKey, consumerSecret, accessToken, accessTokenSecret, int64(value))
					} else {
						code = 1001
						msg = err.Error()
					}
				}
			}
		case "Block":
			{
				user := s.TrimSpace(context.GetInput("user").(string))
				if len(user) == 0 {

					code = 105
					msg = "Block user field cannot be blank"

				} else {
					code, msg = Twitter.BlockUser(consumerKey, consumerSecret, accessToken, accessTokenSecret, user)
				}
			}
		case "Unblock":
			{
				user := s.TrimSpace(context.GetInput("user").(string))
				if len(user) == 0 {

					code = 105
					msg = "UnBlock user field cannot be blank"

				} else {
					code, msg = Twitter.UnBlockUser(consumerKey, consumerSecret, accessToken, accessTokenSecret, user)
				}
			}
		case "Follow":
			{
				user := s.TrimSpace(context.GetInput("user").(string))
				if len(user) == 0 {

					code = 105
					msg = "Follow user field cannot be blank"

				} else {
					code, msg = Twitter.Follow(consumerKey, consumerSecret, accessToken, accessTokenSecret, user)
				}
			}
		case "Unfollow":
			{
				user := s.TrimSpace(context.GetInput("user").(string))
				if len(user) == 0 {

					code = 105
					msg = "UnFollow user field cannot be blank"

				} else {
					code, msg = Twitter.UnFollow(consumerKey, consumerSecret, accessToken, accessTokenSecret, user)
				}
			}
		case "DM":
			{
				user := s.TrimSpace(context.GetInput("user").(string))
				directmsg := s.TrimSpace(context.GetInput("text").(string))
				if len(user) == 0 || len(directmsg) == 0 {
					context.SetOutput("statusCode", "105")

					context.SetOutput("message", "User or Text field cannot be blank")

					code = 105
					msg = "User or Text field cannot be blank"

				} else {
					code, msg = Twitter.DirectMessage(consumerKey, consumerSecret, accessToken, accessTokenSecret, directmsg, user)
				}
			}
		}

		context.SetOutput("statusCode", code)

		context.SetOutput("message", msg)
	}

	return true, err
}
