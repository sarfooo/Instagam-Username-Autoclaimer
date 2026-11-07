package main

import (
	"crypto/tls"
	"bufio"
	"time"
	"os"
)

var InstagramWebSessionId = os.Getenv("INSTAGRAM_WEB_SESSION_ID")

const (
	// Instagram Hashes
	InstagramLSD = 		 	 "AdRS8Yr0bk4d6oawJQu9exLPofw"
	InstagramCSRF = 	 	 "ciFdupQEwp4490e1rVhwPA"
	InstagramDATR = 	 	 "y0SmaXI7JfSaS1yL-ltrzoxG"
	InstagramAPICookie = 	 "datr=" + InstagramDATR + "; csrftoken=" + InstagramCSRF + "; " + "ds_user_id="
    // Checking options
    UseBulkRouteDefault 	 = false
    UseBulkRouteShop    	 = true
    UseGraphQLSingle         = false
    UseGraphQLClientRedirect = false
    UseGraphQLOneTap         = false
    UseGraphQLSuggestions    = false

	// A/C Configuration
	MonitorBans 	   		 = true
    BatchSize          		 = 15
	ActiveSessionPoolSize	 = 1
    SpammerGoroutines  		 = 300
    SpammerDuration    		 = time.Second * 1
	
	// Discord Webhooking
	DiscordAutoClaimWebhook string = ""
	DiscordMonitorWebhook string = ""
	DiscordAutoClaimedTemplate string = `{"content":null,"embeds":[{"title":"Instagram Autoclaimer","description":"@[%s](https://www.instagram.com/%s) has been autoclaimed","color":%v,"thumbnail":{"url":"%s"}}],"attachments":[]}`
	DiscordSwappedTemplate string = `{"content":null,"embeds":[{"title":"Instagram Monitor","description":"@[%s](https://www.instagram.com/%s) has been swapped to @[%s](https://www.instagram.com/%s)","color":%v,"fields":[{"name":"Previous Profile ID","value":"%s"},{"name":"New Profile ID","value":"%s"}],"thumbnail":{"url":"%s"}}],"attachments":[]}`
	DiscordDeactivatedTemplate string = `{"content":null,"embeds":[{"title":"Instagram Monitor","description":"@[%s](https://www.instagram.com/%s) has been deactivated","color":5918792,"fields":[{"name":"ID","value":"%s"}],"thumbnail":{"url":"%s"}}],"attachments":[]}`
)

var (
	fullUsernames []string
	usernames []string
	usernameIds []string
	usernameIndexes [][]int

	usernameBodies []string
	usernameMatchesA [][]byte
	usernameMatchesB [][]byte
	usernameMatchesC [][]byte
	
	activeSessions []string
	activeSession []string
	activeSessionPool [][]string
	activeSessionRequest []byte
	activeSessionRequests [][]byte
	
	bannedSessions []string
	formattedBannedSessions [][]string
	removeBannedSession bool
	
	clientRedirectVerificationBody []byte = []byte("credential_type")
	oneTapBodyVerification = []byte(`rname":`)
	suggestionsBodyVerification []byte = []byte("SUCCESS")

	spammerDummyRequest []byte
	spammerUsernameIndex int
	attemptedUsernames map[string]int
	
	globalConnection *tls.Conn
	globalBuffer *bufio.Reader
	
	synchronize bool
	rotateEndpoint bool

	attempts int64
	rl int64
	rotated int64
)