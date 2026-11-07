package main

import (
	"strings"
	"time"
	"fmt"
	"log"
)

func isUsernameClaimed(index int) {
	newUsernameId, profilePicture := fetchUsernameInfo(usernames[index])
	if newUsernameId == usernameIds[index] {
		logFalseDetection(index)
	} else {
		getUsernameStatus(index, newUsernameId, profilePicture)
		if session := findClaimedSession(newUsernameId); session != nil {
			claimedUsername(index, profilePicture, session)
		}
	}
	time.Sleep(time.Second)
	buildState()
}

func claimedUsername(index int, profilePicture string, session []string) {
	log.Printf("Successfully autoclaimed @%s %s \n\n", usernames[index], strings.Repeat(" ", 70))
	writeFile("./data/logs/" + usernames[index] + ".log", fmt.Sprintf("Username: %s\nPassword: %s\nEmail: %s\nEmail Password: %s\nAuthorization: %s\n", usernames[index], session[1], session[2], session[3], session[4]))
	sendDiscordWebhook(DiscordAutoClaimWebhook, fmt.Sprintf(DiscordAutoClaimedTemplate, usernames[index], usernames[index], randomRGBColor(), profilePicture))

	activeSessions = removeString(activeSessions, session[5])
	createFile("./data/active_sessions.txt", activeSessions)
}

func getUsernameStatus(index int, newUsernameId, profilePicture string) {
	if newUsernameId == "Unknown" {
		log.Printf("@%s has been deactivated %s \n\n", usernames[index], strings.Repeat(" " , 70))
		sendDiscordWebhook(DiscordMonitorWebhook, fmt.Sprintf(DiscordDeactivatedTemplate, usernames[index], usernames[index], usernameIds[index], profilePicture))
	} else {
		newUsername := fetchUsername(usernameIds[index])
		log.Printf("@%s has been swapped to @%s %s \n\n", usernames[index], newUsername, strings.Repeat(" " , 70))
		sendDiscordWebhook(DiscordMonitorWebhook, fmt.Sprintf(DiscordSwappedTemplate, usernames[index], usernames[index], newUsername, newUsername, randomRGBColor(), usernameIds[index], newUsernameId, profilePicture))
		fullUsernames = append(fullUsernames, usernames[index] + ":" + newUsernameId)
	}
	fullUsernames = removeIndex(fullUsernames, index)
	createFile("./data/usernames.txt", fullUsernames)
}

func logFalseDetection(index int) {
	log.Printf("False detection on @%s %s \n\n", usernames[index], strings.Repeat(" ", 70))
	if attemptedUsernames[usernames[index]] == 2 {
		log.Printf("Removed @%s due to multiple claim failures %s \n\n", usernames[index], strings.Repeat(" ", 70))
		fullUsernames = removeIndex(fullUsernames, index)
		createFile("./data/usernames.txt", fullUsernames)
	}
	attemptedUsernames[usernames[index]]++
}

func sessionReactivated(session []string) {
	sessionId := strings.Split(strings.Split(session[0], "sessionid=")[1], ";")[0]
	if containsStringSplit(bannedSessions, sessionId, ":", 4) {
		fullSession := strings.Join(findStringSplit(bannedSessions, sessionId, ":", 4), ":")
		writeFile("./data/reactivated_sessions.txt", fullSession)	
	
		bannedSessions = removeString(bannedSessions, fullSession)
		createFile("./data/banned_sessions.txt", bannedSessions)
		formattedBannedSessions = formatBannedSessions()
	}
}