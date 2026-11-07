package main

import (
	"github.com/valyala/fasthttp/fasthttpproxy"
	"github.com/valyala/fasthttp"
	"time"
	"net"
)

func createHostClient(connections int) *fasthttp.HostClient {
	return &fasthttp.HostClient {
		Addr: "31.13.80.52:443",
		TLSConfig: createTLSConfig(),
		ReadBufferSize: 8196,
		WriteBufferSize: 8196,
		MaxConns: connections,
		IsTLS: true,
		DisablePathNormalizing: true,
		NoDefaultUserAgentHeader: true,
		DisableHeaderNamesNormalizing: true,
	}
}

func createProxyClient(connections int) (client *fasthttp.Client, channel chan string) {
	channel = make(chan string)
	return &fasthttp.Client {
		TLSConfig: createTLSConfig(),
		Dial: func(addr string) (net.Conn, error) {
			return fasthttpproxy.FasthttpHTTPDialer(<- channel)(addr)
		},
		ReadBufferSize: 8196,
		WriteBufferSize: 8196,
		MaxConnsPerHost: connections,
		MaxConnDuration: time.Millisecond * 100,
		DisablePathNormalizing: true,
		NoDefaultUserAgentHeader: true,
		DisableHeaderNamesNormalizing: true,
		
	}, channel
}

func createFastHTTPRequest(method, url string, useIOSUserAgent, useDefaultContentType bool, compressionMode string) *fasthttp.Request {
    request := fasthttp.AcquireRequest()
    request.Header.SetMethod(method)
    request.SetRequestURI(url)
	
	request.Header.Set("Connection", "keep-alive")
	request.Header.Set("Cache-Control", "no-store")
	request.Header.Set("Sec-Fetch-Site", "same-origin")
    request.Header.Set("User-Agent", "Instagram 404.0.0.48.76 Android )")
	if useIOSUserAgent {
        request.Header.Set("User-Agent", "Instagram 361.0.0.46.88 (iPhone14,5; iOS 16_0_2; en_GB; en; scale=3.00; 1170x2532; 521442846) AppleWebKit/420+")
    }
    if useDefaultContentType {
        request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    }
    if compressionMode != "" {
        request.Header.Set("Accept-Encoding", compressionMode)
    }
    if UseGraphQLSuggestions {
        request.Header.Set("X-FB-LSD", "M5-hy7Vh-5Fd8KTsfam1jW")
    }
    return request
}

func createFastHTTPResponse(skipBody bool) *fasthttp.Response {
	response := fasthttp.AcquireResponse()
	response.SkipBody = skipBody
	return response
}