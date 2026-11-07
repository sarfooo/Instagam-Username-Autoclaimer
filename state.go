package main

import (
	"strings"
	"time"
	"fmt"
)

func buildState() {
	usernames = splitSliceAtDelimiter(fullUsernames, ":", 0)
	usernameIds = splitSliceAtDelimiter(fullUsernames, ":", 1)
	
	usernameIndexes = chunkSlice(createIndexSlice(usernames), BatchSize)
	usernameBodies = createBodies()
	usernameMatchesA, usernameMatchesB, usernameMatchesC = createMatches()
	
	activeSessionPool = createActiveSessionPool()
	activeSession = activeSessionPool[0]
	activeSessionRequests = buildUsernameChangeRequests()
}

func createActiveSessionPool() (pool [][]string) {
	for _, session := range shuffleSlice(activeSessions) {
		info := strings.Split(session, ":")
		bearer := strings.Join(info[4:], ":")
		if !isSessionActive(bearer) {
			continue
		}
		pool = append(pool, append(info[:4], bearer, session))
		if len(pool) == ActiveSessionPoolSize {
			return pool
		}
	}
	if len(pool) == 0 {
		panic("Please restart A/C with at least 1 active Instagram session")
	}
	return pool
}

func sessionStateRotater() {
	for {
		time.Sleep(time.Minute * 30)
		if activeSessionRequest != nil {
			time.Sleep(time.Second * 30)
		}
		activeSessionPool = createActiveSessionPool()
		activeSession = activeSessionPool[0]
		activeSessionRequests = buildUsernameChangeRequests()
	}
}

func formatBannedSessions() (sessions [][]string) {
	for _, session := range shuffleSlice(bannedSessions) {
		info := strings.Split(session, ":")
		sessions = append(sessions, []string{"sessionid=" + info[4] + "; ds_user_id=", "fb_dtsg=" + strings.Join(info[5:], ":")})
	}
	return sessions
}

func createMatches() (matchesA, matchesB, matchesC [][]byte) {
	for i := 0; i < len(usernames); i++ {
		switch {
			case UseGraphQLOneTap:
				matchesA = append(matchesA, []byte(`rname":"` + usernames[i] + `"`))

			case UseGraphQLClientRedirect:
				matchesA = append(matchesA, []byte(`"` + usernames[i] + `\`))
				matchesB = append(matchesB, []byte(`"` + usernameIds[i] + `\`))

			case UseGraphQLSuggestions:
				matchesA = append(matchesA, []byte(`"` + usernames[i] + `","`))
			
			case UseBulkRouteShop:
				matchesA = append(matchesA, []byte("/" + usernames[i] + `/shop/":{"error":t`))
				
			case UseBulkRouteDefault:
				matchesA = append(matchesA, []byte(`(\u0040` + usernames[i] + `)`))
				matchesB = append(matchesB, []byte(`"` + usernames[i] + ` \u`))
				matchesC = append(matchesC, []byte("/" + usernames[i] + `/":{"error":false,"result":{`))
		}
	}
	return matchesA, matchesB, matchesC
}

func createBodies() (bodies []string) {
	if UseGraphQLSingle {
		return buildGraphQLSingle()
	}
	var chunks [][]string
	var builder func([]string) string
	switch {		
		case UseGraphQLClientRedirect:
			chunks, builder = chunkSlice(usernameIds, BatchSize), buildGraphQLClientRedirect

		case UseGraphQLOneTap:
			chunks, builder = chunkSlice(usernameIds, BatchSize), buildGraphQLOneTap

		case UseGraphQLSuggestions:
			chunks, builder = chunkSlice(usernames, BatchSize), buildGraphQLSuggestions

		case UseBulkRouteShop || UseBulkRouteDefault:
			chunks, builder = chunkSlice(usernames, BatchSize), buildBulkBody
	}
	for _, chunk := range chunks {
		bodies = append(bodies, builder(chunk))
	}
	return bodies
}

func buildGraphQLSingle() (bodies []string)  {
	for i := range usernames {
		bodies = append(bodies, `https://i.instagram.com/graphql_www?doc_id=6881983411865519&variables={"username":"` + usernames[i] + `"}`)
	}
	return bodies
}

func buildGraphQLOneTap(batch []string) (body string) {
	body = `https://i.instagram.com/graphql/query?__a=1&doc_id=24047751574823150&variables={"fb":false,"fb_access_token":"","fb_oidc_flow":false,"logged_in":false,"one_tap":true,"one_tap_user_ids":["`
	return body + strings.Join(batch, `","`) + 	`"]}`
}	
	
func buildGraphQLClientRedirect(batch []string) (url string) {
	bodyIds := strings.Join(batch, `\",\"credential_type\":\"none\",\"token\":\"\"},{\"uid\":\"`)
	if rotateEndpoint {
		url = `https://i.instagram.com/graphql_www?doc_id=7765850536785467&variables={"input":{"appid":"com.bloks.www.bloks.caa.login.process_client_data_and_redirect","bloks_versioning_id":"6ffa76cd52f4fb43c58effe983d6ad7dd8ead63a7d2e4f5b41b54fa30b5ce39d","params":"{\"family_device_id\":\"5a3b9397-3c33-43bb-a783-dbcec7fe0d22\",\"account_list\":[{\"uid\":\"`
		return url + bodyIds + `\",\"credential_type\":\"none\",\"token\":\"\"}]}"}}`	
	}	
	url = `https://www.instagram.com/async/wbloks/fetch/?__a=1&lsd=` + InstagramLSD + `&appid=com.bloks.www.bloks.caa.login.process_client_data_and_redirect&__bkv=6ffa76cd52f4fb43c58effe983d6ad7dd8ead63a7d2e4f5b41b54fa30b5ce39d&type=app&params={"params":"{\"server_params\":{\"family_device_id\":\"5a3b9397-3c33-43bb-a783-dbcec7fe0d22\",\"account_list\":[{\"uid\":\"`
	return url + bodyIds + `\",\"credential_type\":\"none\",\"token\":\"\"}]}}"}`
}	

func buildGraphQLSuggestions(batch []string) string {
	var usernames string
	if len(batch) > 1 {
		usernames = strings.Join(batch[1:], " ")
	}
	return `doc_id=25391252800555418&__a=1&lsd=AdRVc4hh2gu3jkkCSY95yriX4_8&variables={"scale":1,"input":{"contactpoint_type":"EMAIL","field_name":"CONTACTPOINT","contactpoint":{"sensitive_string_value":"` + batch[0] + `@gmail.com"},"fullname":{"sensitive_string_value":"` + usernames + `"}}}`
}

func buildBulkBody(batch []string) (body string) {
	body = "__a=1&__comet_req=117&routing_namespace=igx_www_logged_out%24a%2487a091182d5bd65bcb043a2888004e09&__csr=gZJRb45c9Wnp44VqQ9a48tBzVe8y85-U1W44kOoa9U8p8iyo01Uso0hH4zobk30c31g3Oy54xG1tzrBmg80jnQ1ewgp2w8ho0gKo1Yo0kWg3TU&__hblp=08C7458cqwllwSw5_w2lE2no1tUaU1wocEqG0hm0QE21w9i&__sjsp=n1T1h9vi977gMk4Hw&__spin_r=1034302344&__spin_b=trunk&__spin_t=1772478466&"
	suffix := "/%s/&"
	if UseBulkRouteShop {
		body = "routing_namespace=igx_www%24a%2487a091182d5bd65bcb043a2888004e09&__d=www&__user=0&__a=1&__req=5&__hs=20516.HYP%3Ainstagram_web_pkg.2.1...0&dpr=1&__ccg=EXCELLENT&__rev=1034413161&__s=155cwv%3Anvaxst%3At7az6s&__hsi=7613228449832872992&__dyn=7xeUjG1mxu1syUbFp41twpUnwgU29zEdEc8co2qwJwbS0E8hw2nVE4W0qa0FE2awt81s8hwGwQwoEcE7O2l0Fwqo31w9O0H8jwae4UaEW2G0AEco5G1Wxfxm16wUwtE1wEbUGdG1QwTwFwIwbS1LwTwKG1pg2Xwr86C1mg6LhAq4rwh8nAwHxW1owLwfybyohw5yweu&__csr=hIj91p34lvbtRkll8gYCAATiQzylGql4gOl5xqbzrCGVoyczpVAWDDBDhoa9UhAgd8gDx-VoSU-GxCl7zUZ5wwAx3KqjwqQEWcUkUlBK6oG14hKifBCyokWxlbogzoy3C6Ufo-0tu0crwlU6R120mWvg01vgmEkw4QwkuU0C10cZ4w1eG06r207TggxZ0aUxox2Vi0ywyw9vwgA08R81Rydw9xw6HwoE4Mi1nx21z-0NU5dw2f80X6m01Dgw0joE0nPw&__hsdp=l2QbB0A84gapaekoCqRlauqkIY9V18IwybAkEZwyg2sw9le2phw1m809rU0qkwcW3q0o_woE3YDwbm0fXwOwqU0wC&__hblp=0e20PE9E1Uopw9W14x28xG1qwqVE3kwxAwbem0P80Pe09Qw4SwSweynGE5-0aWw4ky82DxSu5Umwba1fwf60_9U2Rwj80Fa0U8eEcEcokCxS6820wby0JEa8&__sjsp=l2QbB0A84gapaekoCqRlauqkI_4PV18IwybAkwgo30w9le2phw055y&__comet_req=7&__spin_r=1034413161&__spin_b=trunk&__spin_t=1772592880&__crn=comet.igweb.PolarisCommerceStorefrontRoute&"
		suffix = "/%s/shop/&"
	}
	for i := range batch {
		body += fmt.Sprintf("route_urls[%d]=" + suffix, i, batch[i])
	}
	return body
}