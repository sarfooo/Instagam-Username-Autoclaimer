package main

import (
	"github.com/valyala/fasthttp"
	"crypto/tls"
	"strings"
	"bufio"
	"time"
	"net"
	_"fmt"
	"log"
)

func createTLSConfig() *tls.Config {
	return &tls.Config {
		MinVersion: tls.VersionTLS13,
		MaxVersion: tls.VersionTLS13,
		InsecureSkipVerify: true,
		SessionTicketsDisabled: false,
		DynamicRecordSizingDisabled: false,
		ClientSessionCache: tls.NewLRUClientSessionCache(10000),
	}
}

func createTCPConnection() *net.TCPConn {
	for {
		address, _ := net.ResolveTCPAddr("tcp4", "31.13.80.52:443")
		if connection, err := net.DialTCP("tcp4", nil, address); err == nil {
			connection.SetNoDelay(true)
			connection.SetKeepAlive(true)    
			connection.SetLinger(0)
            connection.SetReadBuffer(8192)
			connection.SetWriteBuffer(4096) 
			return connection
		}
		log.Printf("Failed to create TCP Connection %s \n\n", strings.Repeat(" ", 45))
	}
}

func createTLSConnection() (*tls.Conn, *bufio.Reader) {
	for {
		connection := tls.Client(createTCPConnection(), createTLSConfig())
		if connection.Handshake() == nil {
			buffer := bufio.NewReaderSize(connection, 8192)
			warmUpConnection(connection, buffer)
			return connection, buffer
		}
		log.Printf("Failed to establish TLS handshake %s \n\n", strings.Repeat(" ", 45))
	}
}

func warmUpConnection(connection *tls.Conn, buffer *bufio.Reader) {
	connection.Write([]byte(getDummyRequest()))
	response := createFastHTTPResponse(true)
	response.Read(buffer)
	fasthttp.ReleaseResponse(response)
}

func connectionRefresher(connection **tls.Conn, buffer **bufio.Reader) {
    for {
        time.Sleep(time.Second * 15)
        if activeSessionRequest != nil {
            time.Sleep(time.Second * 5)
        }
        oldConnection := *connection
		*connection, *buffer = createTLSConnection()
        oldConnection.Close()
    }
}

func getDummyRequest() string {
	return "HEAD /graphql_www HTTP/1.1\r\n" +
		   "Host: i.instagram.com\r\n" +
		   "Connection: keep-alive\r\n" +
		   "User-Agent: Instagram 404.0.0.48.76 Android )\r\n\r\n"
}

func buildChangeUsernameRequest(username, bearer string) string {
	return `HEAD /graphql_www?client_doc_id=385233008511036066687298726293&variables={"bk_context":{},"params":{"bloks_versioning_id":"02aa82d2510c3c91e0f953bebbdd0b36aedd15b580d2f6557e7b4b3b116ee5a3","app_id":"com.bloks.www.bloks.caa.reg.username.async","params":"{\"params\":\"{\\\"server_params\\\":{\\\"screen_id\\\":0,\\\"text_input_id\\\":0,\\\"input_id\\\":0,\\\"suggestions_container_id\\\":0,\\\"action\\\":1,\\\"post_tos\\\":1,\\\"reg_info\\\":\\\"{}\\\",\\\"flow_info\\\":\\\"{\\\\\\\"flow_name\\\\\\\":\\\\\\\"new_to_family_fb_default\\\\\\\",\\\\\\\"flow_type\\\\\\\":\\\\\\\"ntf\\\\\\\"}\\\"},\\\"client_input_params\\\":{\\\"validation_text\\\":\\\"` + username + `\\\"}}\"}"}} HTTP/1.1` + "\r\n" +
		   "Host: i.instagram.com\r\n" +
		   "Connection: keep-alive\r\n" +
		   "Authorization: " + bearer + "\r\n" +
		   "User-Agent: Instagram 404.0.0.48.76 Android )\r\n\r\n"
}

func buildUsernameChangeRequests() (requests [][]byte) {
	for i := range usernames {
		requests = append(requests, []byte(buildChangeUsernameRequest(usernames[i], activeSession[4])))
	}
	return requests
}