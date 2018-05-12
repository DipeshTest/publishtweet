package publishtweet

import (
	"io/ioutil"
	"testing"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

var activityMetadata *activity.Metadata

func getActivityMetadata() *activity.Metadata {

	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil {
			panic("No Json Metadata found for activity.json path")
		}

		activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	}

	return activityMetadata
}

func TestCreate(t *testing.T) {

	act := NewActivity(getActivityMetadata())

	if act == nil {
		t.Error("Activity Not Created")
		t.Fail()
		return
	}
}

// func TestTwitterFunction_Tweet(t *testing.T) {
//
// 	defer func() {
// 		if r := recover(); r != nil {
// 			t.Failed()
// 			t.Errorf("panic during execution: %v", r)
// 		}
// 	}()
//
// 	act := NewActivity(getActivityMetadata())
//
// 	tc := test.NewTestActivityContext(getActivityMetadata())
//
// 	//setup attrs
//
// 	tc.SetInput("consumerKey", "")
// 	tc.SetInput("consumerSecret", "")
// 	tc.SetInput("accessToken", "")
// 	tc.SetInput("accessTokenSecret", "")
// 	tc.SetInput("twitterFunction", "Tweet")
// 	tc.SetInput("text", "Testing twitter new connector AllStars :)")
//
// 	act.Eval(tc)
//
// 	//check result attr
//
// 	code := tc.GetOutput("statusCode")
//
// 	msg := tc.GetOutput("message")
// 	fmt.Print(msg)
// 	assert.Equal(t, code, "200")
//
// }

// func TestTwitterFunction_ReTweet(t *testing.T) {
//
// 	defer func() {
// 		if r := recover(); r != nil {
// 			t.Failed()
// 			t.Errorf("panic during execution: %v", r)
// 		}
// 	}()
//
// 	act := NewActivity(getActivityMetadata())
//
// 	tc := test.NewTestActivityContext(getActivityMetadata())
//
// 	//setup attrs
//
// 	tc.SetInput("consumerKey", "")
// 	tc.SetInput("consumerSecret", "")
// 	tc.SetInput("accessToken", "")
// 	tc.SetInput("accessTokenSecret", "")
// 	tc.SetInput("twitterFunction", "ReTweet")
// 	tc.SetInput("user", "992714282587455488")
// 	tc.SetInput("text", "Retweeting my previous tweet again, I know weird but connector test karne k liye karna padta hai :p")
//
// 	act.Eval(tc)
//
// 	//check result attr
//
// 	code := tc.GetOutput("statusCode")
//
// 	msg := tc.GetOutput("message")
// 	fmt.Print(msg)
// 	assert.Equal(t, code, "200")
//
// }

// func TestTwitterFunction_Block(t *testing.T) {
//
// 	defer func() {
// 		if r := recover(); r != nil {
// 			t.Failed()
// 			t.Errorf("panic during execution: %v", r)
// 		}
// 	}()
//
// 	act := NewActivity(getActivityMetadata())
//
// 	tc := test.NewTestActivityContext(getActivityMetadata())
//
// 	//setup attrs
//
// 	tc.SetInput("consumerKey", "")
// 	tc.SetInput("consumerSecret", "")
// 	tc.SetInput("accessToken", "")
// 	tc.SetInput("accessTokenSecret", "")
// 	tc.SetInput("twitterFunction", "Block")
// 	tc.SetInput("user", "@FLOGOALLSTARS")
// 	//tc.SetInput("text", "Retweeting my previous tweet again, I know weird but connector test karne k liye karna padta hai :p")
//
// 	act.Eval(tc)
//
// 	//check result attr
//
// 	code := tc.GetOutput("statusCode")
//
// 	msg := tc.GetOutput("message")
// 	fmt.Print(msg)
// 	assert.Equal(t, code, "200")
//
// }
//
//
// func TestTwitterFunction_Unblock(t *testing.T) {
//
// 	defer func() {
// 		if r := recover(); r != nil {
// 			t.Failed()
// 			t.Errorf("panic during execution: %v", r)
// 		}
// 	}()
//
// 	act := NewActivity(getActivityMetadata())
//
// 	tc := test.NewTestActivityContext(getActivityMetadata())
//
// 	//setup attrs
//
// 	tc.SetInput("consumerKey", "")
// 	tc.SetInput("consumerSecret", "")
// 	tc.SetInput("accessToken", "")
// 	tc.SetInput("accessTokenSecret", "")
// 	tc.SetInput("twitterFunction", "Unblock")
// 	tc.SetInput("user", "@FLOGOALLSTARS")
// 	//tc.SetInput("text", "Retweeting my previous tweet again, I know weird but connector test karne k liye karna padta hai :p")
//
// 	act.Eval(tc)
//
// 	//check result attr
//
// 	code := tc.GetOutput("statusCode")
//
// 	msg := tc.GetOutput("message")
// 	fmt.Print(msg)
// 	assert.Equal(t, code, "200")
//
// }

// func TestTwitterFunction_Follow(t *testing.T) {
//
// 	defer func() {
// 		if r := recover(); r != nil {
// 			t.Failed()
// 			t.Errorf("panic during execution: %v", r)
// 		}
// 	}()
//
// 	act := NewActivity(getActivityMetadata())
//
// 	tc := test.NewTestActivityContext(getActivityMetadata())
//
// 	//setup attrs
//
// 	tc.SetInput("consumerKey", "")
// 	tc.SetInput("consumerSecret", "")
// 	tc.SetInput("accessToken", "")
// 	tc.SetInput("accessTokenSecret", "")
// 	tc.SetInput("twitterFunction", "Follow")
// 	tc.SetInput("user", "@FLOGOALLSTARS")
// 	//tc.SetInput("text", "Retweeting my previous tweet again, I know weird but connector test karne k liye karna padta hai :p")
//
// 	act.Eval(tc)
//
// 	//check result attr
//
// 	code := tc.GetOutput("statusCode")
//
// 	msg := tc.GetOutput("message")
// 	fmt.Print(msg)
// 	assert.Equal(t, code, "200")
//
// }

// func TestTwitterFunction_Unfollow(t *testing.T) {
//
// 	defer func() {
// 		if r := recover(); r != nil {
// 			t.Failed()
// 			t.Errorf("panic during execution: %v", r)
// 		}
// 	}()
//
// 	act := NewActivity(getActivityMetadata())
//
// 	tc := test.NewTestActivityContext(getActivityMetadata())
//
// 	//setup attrs
//
// 	tc.SetInput("consumerKey", "")
// 	tc.SetInput("consumerSecret", "")
// 	tc.SetInput("accessToken", "")
// 	tc.SetInput("accessTokenSecret", "")
// 	tc.SetInput("twitterFunction", "Unfollow")
// 	tc.SetInput("user", "@FLOGOALLSTARS")
// 	//tc.SetInput("text", "Retweeting my previous tweet again, I know weird but connector test karne k liye karna padta hai :p")
//
// 	act.Eval(tc)
//
// 	//check result attr
//
// 	code := tc.GetOutput("statusCode")
//
// 	msg := tc.GetOutput("message")
// 	fmt.Print(msg)
// 	assert.Equal(t, code, "200")
//
// }

// func TestTwitterFunction_DM(t *testing.T) {
//
// 	defer func() {
// 		if r := recover(); r != nil {
// 			t.Failed()
// 			t.Errorf("panic during execution: %v", r)
// 		}
// 	}()
//
// 	act := NewActivity(getActivityMetadata())
//
// 	tc := test.NewTestActivityContext(getActivityMetadata())
//
// 	//setup attrs
//
// 	tc.SetInput("consumerKey", "")
// 	tc.SetInput("consumerSecret", "")
// 	tc.SetInput("accessToken", "")
// 	tc.SetInput("accessTokenSecret", "")
// 	tc.SetInput("twitterFunction", "DM")
// 	tc.SetInput("user", "@FLOGOALLSTARS")
// 	tc.SetInput("text", "ping")
//
// 	act.Eval(tc)
//
// 	//check result attr
//
// 	code := tc.GetOutput("statusCode")
//
// 	msg := tc.GetOutput("message")
// 	fmt.Print(msg)
// 	assert.Equal(t, code, "200")
//
// }

// func TestTwitterFunction_EmptyConsumerKey(t *testing.T) {
//
// 	defer func() {
// 		if r := recover(); r != nil {
// 			t.Failed()
// 			t.Errorf("panic during execution: %v", r)
// 		}
// 	}()
//
// 	act := NewActivity(getActivityMetadata())
//
// 	tc := test.NewTestActivityContext(getActivityMetadata())
//
// 	//setup attrs
//
// 	tc.SetInput("consumerKey", "")
// 	tc.SetInput("consumerSecret", "1234")
// 	tc.SetInput("accessToken", "xtz")
// 	tc.SetInput("accessTokenSecret", "qwer")
// 	tc.SetInput("twitterFunction", "Follow")
// 	tc.SetInput("user", "@FLOGOALLSTARS")
// 	//tc.SetInput("text", "Retweeting my previous tweet again, I know weird but connector test karne k liye karna padta hai :p")
//
// 	act.Eval(tc)
//
// 	//check result attr
//
// 	code := tc.GetOutput("statusCode")
//
// 	msg := tc.GetOutput("message")
// 	fmt.Print(msg)
// 	assert.Equal(t, code, "101")
//
// }
//
// func TestTwitterFunction_EmptyConsumerSecret(t *testing.T) {
//
// 	defer func() {
// 		if r := recover(); r != nil {
// 			t.Failed()
// 			t.Errorf("panic during execution: %v", r)
// 		}
// 	}()
//
// 	act := NewActivity(getActivityMetadata())
//
// 	tc := test.NewTestActivityContext(getActivityMetadata())
//
// 	//setup attrs
//
// 	tc.SetInput("consumerKey", "xyz")
// 	tc.SetInput("consumerSecret", "")
// 	tc.SetInput("accessToken", "1234")
// 	tc.SetInput("accessTokenSecret", "1234")
// 	tc.SetInput("twitterFunction", "Follow")
// 	tc.SetInput("user", "@FLOGOALLSTARS")
// 	//tc.SetInput("text", "Retweeting my previous tweet again, I know weird but connector test karne k liye karna padta hai :p")
//
// 	act.Eval(tc)
//
// 	//check result attr
//
// 	code := tc.GetOutput("statusCode")
//
// 	msg := tc.GetOutput("message")
// 	fmt.Print(msg)
// 	assert.Equal(t, code, "102")
//
// }

// func TestTwitterFunction_EmptyAccessToken(t *testing.T) {
//
// 	defer func() {
// 		if r := recover(); r != nil {
// 			t.Failed()
// 			t.Errorf("panic during execution: %v", r)
// 		}
// 	}()
//
// 	act := NewActivity(getActivityMetadata())
//
// 	tc := test.NewTestActivityContext(getActivityMetadata())
//
// 	//setup attrs
//
// 	tc.SetInput("consumerKey", "123")
// 	tc.SetInput("consumerSecret", "1234")
// 	tc.SetInput("accessToken", "")
// 	tc.SetInput("accessTokenSecret", "qwer")
// 	tc.SetInput("twitterFunction", "Follow")
// 	tc.SetInput("user", "@FLOGOALLSTARS")
// 	//tc.SetInput("text", "Retweeting my previous tweet again, I know weird but connector test karne k liye karna padta hai :p")
//
// 	act.Eval(tc)
//
// 	//check result attr
//
// 	code := tc.GetOutput("statusCode")
//
// 	msg := tc.GetOutput("message")
// 	fmt.Print(msg)
// 	assert.Equal(t, code, "103")
//
// }
//
// func TestTwitterFunction_EmptyAccessTokenSecret(t *testing.T) {
//
// 	defer func() {
// 		if r := recover(); r != nil {
// 			t.Failed()
// 			t.Errorf("panic during execution: %v", r)
// 		}
// 	}()
//
// 	act := NewActivity(getActivityMetadata())
//
// 	tc := test.NewTestActivityContext(getActivityMetadata())
//
// 	//setup attrs
//
// 	tc.SetInput("consumerKey", "xyz")
// 	tc.SetInput("consumerSecret", "asd")
// 	tc.SetInput("accessToken", "qwe")
// 	tc.SetInput("accessTokenSecret", "")
// 	tc.SetInput("twitterFunction", "Follow")
// 	tc.SetInput("user", "@FLOGOALLSTARS")
// 	//tc.SetInput("text", "Retweeting my previous tweet again, I know weird but connector test karne k liye karna padta hai :p")
//
// 	act.Eval(tc)
//
// 	//check result attr
//
// 	code := tc.GetOutput("statusCode")
//
// 	msg := tc.GetOutput("message")
// 	fmt.Print(msg)
// 	assert.Equal(t, code, "104")
//
// }
//
// func TestTwitterFunction_EmptyTwitterFunction(t *testing.T) {
//
// 	defer func() {
// 		if r := recover(); r != nil {
// 			t.Failed()
// 			t.Errorf("panic during execution: %v", r)
// 		}
// 	}()
//
// 	act := NewActivity(getActivityMetadata())
//
// 	tc := test.NewTestActivityContext(getActivityMetadata())
//
// 	//setup attrs
//
// 	tc.SetInput("consumerKey", "xyz")
// 	tc.SetInput("consumerSecret", "asd")
// 	tc.SetInput("accessToken", "qwe")
// 	tc.SetInput("accessTokenSecret", "qq")
// 	tc.SetInput("twitterFunction", "")
// 	tc.SetInput("user", "@FLOGOALLSTARS")
// 	//tc.SetInput("text", "Retweeting my previous tweet again, I know weird but connector test karne k liye karna padta hai :p")
//
// 	act.Eval(tc)
//
// 	//check result attr
//
// 	code := tc.GetOutput("statusCode")
//
// 	msg := tc.GetOutput("message")
// 	fmt.Print(msg)
// 	assert.Equal(t, code, "105")
//
// }
