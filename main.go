package main

import (
	"os/signal"
	"syscall"
	"strings"
	"time"
	"fmt"
	"os"
)

func main() {	
	fmt.Println("Instagram Autoclaimer")

	fullUsernames = openFile("./data/usernames.txt")
	fmt.Printf("Usernames: %s \n", formatNumber(int64(len(fullUsernames))))
	
	activeSessions = shuffleSlice(openFile("./data/active_sessions.txt"))
	fmt.Printf("Active Sessions: %s \n", formatNumber(int64(len(activeSessions))))
	
	bannedSessions = openFile("./data/banned_sessions.txt")
	formattedBannedSessions = formatBannedSessions()
	fmt.Printf("Banned Sessions: %s \n", formatNumber(int64(len(bannedSessions))))
	
	proxies := openFile("./data/proxies.txt")
	fmt.Printf("Proxies: %s \n", formatNumber(int64(len(proxies))))
	
	var goroutines int
	fmt.Printf("Goroutines: ")
	fmt.Scanln(&goroutines)
	fmt.Println()
	
	hostClient := createHostClient(goroutines)
	proxyClient, proxyChannel := createProxyClient(goroutines)
	
	spammerDummyRequest = []byte(getDummyRequest())
	globalConnection, globalBuffer = createTLSConnection()
	go connectionRefresher(&globalConnection, &globalBuffer)
	
	buildState()
	indexChannel := make(chan int, 1)
	bannedSessionChannel := make(chan []string)
	
	var monitorFunction func()
	switch {
		case UseGraphQLSingle:
			monitorFunction = func() { graphQLSingle(proxyClient, indexChannel) }
			
		case UseGraphQLClientRedirect:
			monitorFunction = func() { graphQLClientRedirect(proxyClient, indexChannel) }
			
		case UseGraphQLOneTap:
			monitorFunction = func() { graphQLOneTap(proxyClient, indexChannel) }
			
		case UseGraphQLSuggestions:
			monitorFunction = func() { graphQLSuggestions(proxyClient, indexChannel) }
			
		case UseBulkRouteShop:
			monitorFunction = func() { bulkRouteShopMonitor(hostClient, indexChannel, bannedSessionChannel) }
			
		case UseBulkRouteDefault:
			monitorFunction = func() { bulkRouteDefaultMonitor(hostClient, indexChannel, bannedSessionChannel) }
	}
	
	for i := 0; i < goroutines; i++ {
		go monitorFunction()
	}
	
	attemptedUsernames = make(map[string]int)
	for i := 0; i < SpammerGoroutines; i++ {
		go usernameSpammer()
	}
	go activeSessionCycler()
	go sessionStateRotater()
		
	go func() {
		for {
			for synchronize {
				time.Sleep(100 * time.Millisecond)
			}
			for i := range usernameIndexes {
				if synchronize {
					break
				}
				indexChannel <- i
			}
		}
	}()

	if UseBulkRouteShop || UseBulkRouteDefault {
		go func() {
			for {
				for _, session := range formattedBannedSessions {
					bannedSessionChannel <- session
				}
			}
		}()
	}

	go func() {
		for {
			for i := range proxies {
				proxyChannel <- proxies[i]
			}
		}
	}()

	if UseGraphQLClientRedirect {
		go func() {
			for {
				time.Sleep(time.Hour * 2)
				rotateEndpoint = !rotateEndpoint
				buildState()
			}
		}()
	}

	var rps int64
	go func() {
		for {
			var before int64 = attempts
			time.Sleep(time.Second * 1)
			rps = attempts - before
		}
	}()

	go func() {
		for {
			for _, spinner := range []string{"|", "/", "-", "\\", "|", "/", "-", "\\"} {
				fmt.Printf("[%s] Autoclaiming - Attempts: %s | R/s: %s | C/s: %s | U/s: %.3f | B/s: %.3f | RLs: %s %s \r", spinner, formatNumber(attempts), formatNumber(rps), formatNumber(rps * int64(BatchSize)), float64(rps * int64(BatchSize)) / float64(len(usernames)), float64(len(bannedSessions)) / float64(rps), formatNumber(rl), strings.Repeat(" ", 5))
				time.Sleep(time.Millisecond * 100)
			}
		}
	}()

	var kill chan os.Signal = make(chan os.Signal, 1)
	signal.Notify(kill, os.Interrupt, syscall.SIGTERM)
	go func() {
		<- kill
		fmt.Printf("A/C stopped after %s attempts %s \r\n", formatNumber(attempts), strings.Repeat(" ", 70))
		os.Exit(0)
	}()

	select {}
}
