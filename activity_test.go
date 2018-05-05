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

/*func TestGDriveCreateFile_Success(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs

	tc.SetInput("accessToken", "ya29.GlurBW7n5A2Fk_grpX9KMVeXLEOT4k0OhmSnF_w7626K9kgKmempF_xTDJ6uQVMkdWWWIMiNcb-ht6Rv0gcnuUb2VhtF9h7nltFw0iniwp10dmDQsFT49giOqSA8")
	tc.SetInput("fileFullPath", "C:\\Users\\asutar\\Documents\\CodeOfConduct.pdf")
	tc.SetInput("emailAddr", "jashah@tibco.com")
	tc.SetInput("sendNotification", true)
	tc.SetInput("role", "writer")

	act.Eval(tc)

	//check result attr

	code := tc.GetOutput("statusCode")
	msg := tc.GetOutput("message")
	fmt.Print(msg)
	assert.Equal(t, code, "200")

}*/

/* func TestGDriveCreateFile_TokenError(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs

	tc.SetInput("accessToken", "ya29.GluqBdPuqRNrY9cJ3elwFXSu_4Mx5leuMYXmrdvIB_MPMdU0Y-6M9iU83ELUEJo5UAqeQTN8EaDdUkMxXqhk11E9Ed8H3Cgh353dC05fr3efeKiDebR0Gf11fM7t")
	tc.SetInput("fileFullPath", "C:\\Users\\asutar\\Documents\\CodeOfConduct.pdf")
	tc.SetInput("emailAddr", "jashah@tibco.com")
	tc.SetInput("sendNotification", true)
	tc.SetInput("role", "writer")

	act.Eval(tc)

	//check result attr

	code := tc.GetOutput("statusCode")
	msg := tc.GetOutput("message")
	fmt.Print(msg)
	assert.Equal(t, code, "401")
	tc.SetInput("accessToken", "")

	act.Eval(tc)

	code = tc.GetOutput("statusCode")
	msg = tc.GetOutput("message")
	fmt.Print(msg)
	assert.Equal(t, code, "105")

} */

/* func TestGDriveCreateFile_InvalidPath(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs

	tc.SetInput("accessToken", "ya29.GlurBW7n5A2Fk_grpX9KMVeXLEOT4k0OhmSnF_w7626K9kgKmempF_xTDJ6uQVMkdWWWIMiNcb-ht6Rv0gcnuUb2VhtF9h7nltFw0iniwp10dmDQsFT49giOqSA8")
	tc.SetInput("fileFullPath", "")
	tc.SetInput("emailAddr", "jashah@tibco.com")
	tc.SetInput("sendNotification", true)
	tc.SetInput("role", "writer")

	act.Eval(tc)

	//check result attr

	code := tc.GetOutput("statusCode")
	msg := tc.GetOutput("message")
	fmt.Print(msg)
	assert.Equal(t, code, "106")




	tc.SetInput("fileFullPath", "D://MYGsfdkdsf/jdkd.pdf")
	act.Eval(tc)
	code = tc.GetOutput("statusCode")
	msg = tc.GetOutput("message")
	fmt.Print(msg)
	assert.Equal(t, code, "101")


	tc.SetInput("fileFullPath", "D:\\Programs")


	act.Eval(tc)
	code = tc.GetOutput("statusCode")
	msg = tc.GetOutput("message")
	fmt.Print(msg)
	assert.Equal(t, code, "102")

} */

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
// 	tc.SetInput("consumerKey", "BWSpdB3hoSTv2YTui0hx1rY9Q")
// 	tc.SetInput("consumerSecret", "wFCsXWw8K1eUOG7zwAu0RmGTm74zEGxsp5T3Ss4htRxYbI8VMP")
// 	tc.SetInput("accessToken", "990618908427108355-fUnlckZLtsRVDuy9XGDRkYYlokaD8DB")
// 	tc.SetInput("accessTokenSecret", "Jz186PFXQv6sOg1sYhHYqqX6WDXLLFHIFwHFnu5xM4bXK")
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
// 	tc.SetInput("consumerKey", "BWSpdB3hoSTv2YTui0hx1rY9Q")
// 	tc.SetInput("consumerSecret", "wFCsXWw8K1eUOG7zwAu0RmGTm74zEGxsp5T3Ss4htRxYbI8VMP")
// 	tc.SetInput("accessToken", "990618908427108355-fUnlckZLtsRVDuy9XGDRkYYlokaD8DB")
// 	tc.SetInput("accessTokenSecret", "Jz186PFXQv6sOg1sYhHYqqX6WDXLLFHIFwHFnu5xM4bXK")
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
// 	tc.SetInput("consumerKey", "BWSpdB3hoSTv2YTui0hx1rY9Q")
// 	tc.SetInput("consumerSecret", "wFCsXWw8K1eUOG7zwAu0RmGTm74zEGxsp5T3Ss4htRxYbI8VMP")
// 	tc.SetInput("accessToken", "990618908427108355-fUnlckZLtsRVDuy9XGDRkYYlokaD8DB")
// 	tc.SetInput("accessTokenSecret", "Jz186PFXQv6sOg1sYhHYqqX6WDXLLFHIFwHFnu5xM4bXK")
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
// 	tc.SetInput("consumerKey", "BWSpdB3hoSTv2YTui0hx1rY9Q")
// 	tc.SetInput("consumerSecret", "wFCsXWw8K1eUOG7zwAu0RmGTm74zEGxsp5T3Ss4htRxYbI8VMP")
// 	tc.SetInput("accessToken", "990618908427108355-fUnlckZLtsRVDuy9XGDRkYYlokaD8DB")
// 	tc.SetInput("accessTokenSecret", "Jz186PFXQv6sOg1sYhHYqqX6WDXLLFHIFwHFnu5xM4bXK")
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
// 	tc.SetInput("consumerKey", "BWSpdB3hoSTv2YTui0hx1rY9Q")
// 	tc.SetInput("consumerSecret", "wFCsXWw8K1eUOG7zwAu0RmGTm74zEGxsp5T3Ss4htRxYbI8VMP")
// 	tc.SetInput("accessToken", "990618908427108355-fUnlckZLtsRVDuy9XGDRkYYlokaD8DB")
// 	tc.SetInput("accessTokenSecret", "Jz186PFXQv6sOg1sYhHYqqX6WDXLLFHIFwHFnu5xM4bXK")
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
// 	tc.SetInput("consumerKey", "BWSpdB3hoSTv2YTui0hx1rY9Q")
// 	tc.SetInput("consumerSecret", "wFCsXWw8K1eUOG7zwAu0RmGTm74zEGxsp5T3Ss4htRxYbI8VMP")
// 	tc.SetInput("accessToken", "990618908427108355-fUnlckZLtsRVDuy9XGDRkYYlokaD8DB")
// 	tc.SetInput("accessTokenSecret", "Jz186PFXQv6sOg1sYhHYqqX6WDXLLFHIFwHFnu5xM4bXK")
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
