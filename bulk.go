package main

import (
	"github.com/valyala/fasthttp"
	"strings"
	"bytes"
	"time"
	"log"
)

func bulkRouteDefaultMonitor(client *fasthttp.HostClient, usernameIndexChannel chan int, bannedSessionChannel chan []string) {
	request := createFastHTTPRequest("POST", "https://www.instagram.com/ajax/bulk-route-definitions/", true, true, "br")
	POSTResponse := createFastHTTPResponse(false)
    HEADResponse := createFastHTTPResponse(true)
	for {
		index, session := <- usernameIndexChannel, <- bannedSessionChannel
		request.Header.Set("cookie", session[0] + randomIntString(11))
		request.SetBody([]byte(usernameBodies[index] + session[1]))
		client.Do(request, POSTResponse)
		body, _ := POSTResponse.BodyUnbrotli()
		statusCode := POSTResponse.Header.StatusCode()
		
		if statusCode == 200 {
			for _, index := range usernameIndexes[index] {
				if !bytes.Contains(body, usernameMatchesA[index]) && !synchronize {
					synchronize = true
					activeSessionRequest = activeSessionRequests[index]
					globalConnection.Write(activeSessionRequest)										
					HEADResponse.Read(globalBuffer)
					
					log.Printf("Detected @%s %s \n\n", usernames[index], strings.Repeat(" ", 60))
					time.Sleep(SpammerDuration)
					activeSessionRequest = nil
					
					isUsernameClaimed(index)
					synchronize = false
				}
			}
			attempts++
		} else if statusCode == 429 && !removeBannedSession {
			removeBannedSession = true
			sessionReactivated(session)
			rl++
			removeBannedSession = false
		}
		request.Header.DelAllCookies()
    }
}

func bulkRouteShopMonitor(client *fasthttp.HostClient, usernameIndexChannel chan int, bannedSessionChannel chan []string) {
	request := createFastHTTPRequest("POST", "https://www.instagram.com/ajax/bulk-route-definitions/", true, true, "br")
	POSTResponse := createFastHTTPResponse(false)
    HEADResponse := createFastHTTPResponse(true)
    for {
		index, session := <- usernameIndexChannel, <- bannedSessionChannel
		request.Header.Set("cookie", session[0] + randomIntString(11))
		request.SetBody([]byte(usernameBodies[index] + session[1]))
		client.Do(request, POSTResponse)
		body, _ := POSTResponse.BodyUnbrotli()
		statusCode := POSTResponse.Header.StatusCode()

		if statusCode == 200 {
			for _, index := range usernameIndexes[index] {
				if bytes.Contains(body, usernameMatchesA[index]) && !synchronize {
					synchronize = true
					activeSessionRequest = activeSessionRequests[index]
					globalConnection.Write(activeSessionRequest)										
					HEADResponse.Read(globalBuffer)
					
					log.Printf("Detected @%s %s \n\n", usernames[index], strings.Repeat(" ", 75))
					time.Sleep(SpammerDuration)
					activeSessionRequest = nil
					
					isUsernameClaimed(index)
					synchronize = false
				}
			}
			attempts++
		} else if statusCode == 429 && !removeBannedSession {
			removeBannedSession = true
			sessionReactivated(session)
			rl++
			removeBannedSession = false
		}
		request.Header.DelAllCookies()
    }
}