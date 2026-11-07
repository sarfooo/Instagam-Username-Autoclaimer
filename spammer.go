package main

import "time"

func cycleActiveSession(start int) {
	for i := start; i < len(activeSessionPool); i++ {
		time.Sleep(1 * time.Second)
		if activeSessionRequest != nil {
			activeSessionRequest = []byte(buildChangeUsernameRequest(usernames[spammerUsernameIndex], activeSessionPool[i][4]))
		}
	}
}

func activeSessionCycler() {
	for {
		for activeSessionRequest == nil {
			time.Sleep(time.Millisecond * 100)
		}
		cycleActiveSession(1)
		for activeSessionRequest != nil {
			cycleActiveSession(0)
		}
	}
}

func usernameSpammer() {
	response := createFastHTTPResponse(true)
	connection, buffer := createTLSConnection()
	go connectionRefresher(&connection, &buffer)
	for {
		if activeSessionRequest == nil {
			connection.Write(spammerDummyRequest)
		} else {
			connection.Write(activeSessionRequest)
		}
		response.Read(buffer)
	}
}