package main

import (
	"github.com/valyala/fasthttp"
	"strings"
	"bytes"
	"time"
	"log"
)

func graphQLSingle(client *fasthttp.Client, indexChannel chan int) {
	request := createFastHTTPRequest("POST", "", true, false, "br")
	POSTResponse := createFastHTTPResponse(false)
    HEADResponse := createFastHTTPResponse(true)
	for {
		index := <- indexChannel
		request.SetRequestURI(usernameBodies[index])
		request.Header.Set("Cookie", "ds_user_id=" + randomIntString(11))
		client.Do(request, POSTResponse)
		size := len(POSTResponse.Body())

		if size == 53 && activeSessionRequest == nil {
			activeSessionRequest = activeSessionRequests[index]
			globalConnection.Write(activeSessionRequest)										
			HEADResponse.Read(globalBuffer)
			
			log.Printf("Detected @%s %s \n\n", usernames[index], strings.Repeat(" ", 60))
			time.Sleep(SpammerDuration)
			isUsernameClaimed(index)
			activeSessionRequest = nil

		} else if size > 300 {
			attempts += 1
		} else if size > 0 {
			rl += 1
		}
		request.Header.DelAllCookies()
	}
}

func graphQLClientRedirect(client *fasthttp.Client, indexChannel chan int) {
	request := createFastHTTPRequest("POST", "", true, false, "br")
	POSTResponse := createFastHTTPResponse(false)
    HEADResponse := createFastHTTPResponse(true)
	for {
		index := <- indexChannel
		request.SetRequestURI(usernameBodies[index])
		request.Header.Set("Cookie", InstagramAPICookie + randomIntString(11))
		client.Do(request, POSTResponse)
		body, _ := POSTResponse.BodyUnbrotli()

		if bytes.Contains(body, clientRedirectVerificationBody) {
			for _, index := range usernameIndexes[index] {
				if !bytes.Contains(body, usernameMatchesA[index]) && bytes.Count(body, usernameMatchesB[index]) > 2 && activeSessionRequest == nil {
					activeSessionRequest = activeSessionRequests[index]
					globalConnection.Write(activeSessionRequest)										
					HEADResponse.Read(globalBuffer)
					
					log.Printf("Detected @%s %s \n\n", usernames[index], strings.Repeat(" ", 60))
					time.Sleep(SpammerDuration)
					isUsernameClaimed(index)
					activeSessionRequest = nil
				}
			}
			attempts++
		} else if len(body) > 0 {
			rl++
		}
		request.Header.DelAllCookies()
	}
}

func graphQLOneTap(client *fasthttp.Client, indexChannel chan int) {
	request := createFastHTTPRequest("GET", "", true, false, "gzip")
	POSTResponse := createFastHTTPResponse(false)
    HEADResponse := createFastHTTPResponse(true)
	for {
		index := <- indexChannel
		request.SetRequestURI(usernameBodies[index])
		request.Header.Set("Cookie", "ds_user_id=" + randomIntString(11))
		client.Do(request, POSTResponse)
		body, _ := POSTResponse.BodyGunzip()

		if bytes.Contains(body, oneTapBodyVerification) {
			for _, index := range usernameIndexes[index] {
				if !bytes.Contains(body, usernameMatchesA[index]) && activeSessionRequest == nil {
					activeSessionRequest = activeSessionRequests[index]
					globalConnection.Write(activeSessionRequest)										
					HEADResponse.Read(globalBuffer)
					
					log.Printf("Detected @%s %s \n\n", usernames[index], strings.Repeat(" ", 60))
					time.Sleep(SpammerDuration)
					isUsernameClaimed(index)
					activeSessionRequest = nil
				}
			}
			attempts++	
		} else if len(body) > 0 {
			rl++
		}
		request.Header.DelAllCookies()
	}
}

func graphQLSuggestions(client *fasthttp.Client, indexChannel chan int) {
	request := createFastHTTPRequest("POST", "https://www.instagram.com/api/graphql", true, true, "br")
	POSTResponse := createFastHTTPResponse(false)
    HEADResponse := createFastHTTPResponse(true)
    for {
		index := <- indexChannel
		request.Header.Set("Cookie", "ig_did=89AEF1E6-2E15-45DE-98FE-616F1F6E19CE; ds_user_id=" + randomIntString(11))
		request.SetBody([]byte(usernameBodies[index]))
		client.Do(request, POSTResponse)
		body, _ := POSTResponse.BodyUnbrotli()

		if bytes.Contains(body, suggestionsBodyVerification) {
			for _, index := range usernameIndexes[index] {
				if bytes.Contains(body, usernameMatchesA[index]) && activeSessionRequest == nil {
					activeSessionRequest = activeSessionRequests[index]
					globalConnection.Write(activeSessionRequest)										
					HEADResponse.Read(globalBuffer)
					
					log.Printf("Detected @%s %s \n\n", usernames[index], strings.Repeat(" ", 60))
					time.Sleep(SpammerDuration)
					isUsernameClaimed(index)
					activeSessionRequest = nil
				}
			}
			attempts++
		} else if len(body) > 0 {
			rl++
		}
		request.Header.DelAllCookies()
    }
}