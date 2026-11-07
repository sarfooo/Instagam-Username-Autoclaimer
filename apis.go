package main

import (
	"github.com/valyala/fasthttp"
	"math/rand"
	"strings"
	"log"
)

func isSessionActive(bearer string) bool {
	request := createFastHTTPRequest("GET", `https://i.instagram.com/graphql_www?client_doc_id=385233008511036066687298726293&variables={"bk_context":{},"params":{"bloks_versioning_id":"02aa82d2510c3c91e0f953bebbdd0b36aedd15b580d2f6557e7b4b3b116ee5a3","app_id":"com.bloks.www.bloks.caa.reg.username.async","params":"{\"params\":\"{\\\"server_params\\\":{\\\"screen_id\\\":0,\\\"text_input_id\\\":0,\\\"input_id\\\":0,\\\"suggestions_container_id\\\":0,\\\"action\\\":1,\\\"post_tos\\\":1,\\\"reg_info\\\":\\\"{}\\\",\\\"flow_info\\\":\\\"{\\\\\\\"flow_name\\\\\\\":\\\\\\\"new_to_family_fb_default\\\\\\\",\\\\\\\"flow_type\\\\\\\":\\\\\\\"ntf\\\\\\\"}\\\"},\\\"client_input_params\\\":{\\\"validation_text\\\":\\\"god\\\"}}\"}"}}`, false, false, "")
	response := createFastHTTPResponse(false)

	defer fasthttp.ReleaseRequest(request)
	defer fasthttp.ReleaseResponse(response)

	request.Header.Set("Authorization", bearer)

	fasthttp.Do(request, response)

	return strings.Contains(string(response.Body()), "A user with that username")
}

func fetchUsername(usernameId string) string {
	request := createFastHTTPRequest("GET", `https://i.instagram.com/graphql_www?doc_id=31080414001602564&__a=1&lsd=AdTcZz5w8BF6urmnI-JHdkbVA7w&variables={"location_id":null,"shid":null,"skip_dialog":true,"shared_entity_id":"","id":"` + usernameId + `"}`, false, false, "")
	response := createFastHTTPResponse(false)

	defer fasthttp.ReleaseRequest(request)
	defer fasthttp.ReleaseResponse(response)

	request.Header.Set("Cookie", InstagramWebSessionId + "; ds_user_id=" + randomIntString(11))

	fasthttp.Do(request, response)
	body := string(response.Body())

	if !strings.Contains(body, `{"data":{"xig_user_by_igid_v2"`) {
		log.Printf("Failed to fetch username %s \n\n", strings.Repeat(" ", 45))
		return fetchUsername(usernameId)
	}
	if !strings.Contains(body, `"username":"`) || !MonitorBans && strings.Contains(body, "Instagram User") {
		return "Unknown"
	}
	return strings.Split(strings.Split(body, `username":"`)[1], `"`)[0]
}


func fetchUsernameInfo(username string) (string, string) {
	request := createFastHTTPRequest("GET", `https://www.instagram.com/graphql/query/?variables={"username":"` + username + `","__relay_internal__pv__BarcelonaIsInternalUserrelayprovider":true,"__relay_internal__pv__BarcelonaIsLoggedInrelayprovider":true,"__relay_internal__pv__BarcelonaHasSpoilerStylingInforelayprovider":true,"__relay_internal__pv__BarcelonaShouldShowFediverseM1Featuresrelayprovider":true,"__relay_internal__pv__BarcelonaHasEventBadgerelayprovider":true}&doc_id=25366393769625328`, false, false, "")
	response := createFastHTTPResponse(false)
	
	defer fasthttp.ReleaseRequest(request)
	defer fasthttp.ReleaseResponse(response)

	request.Header.Set("Cookie", InstagramWebSessionId + "; ds_user_id=" + randomIntString(11))

	createHostClient(1).Do(request, response)
	body := string(response.Body())

	if !strings.Contains(body, `{"data":{"user":`) {
		log.Printf("Failed to fetch username ID %s \n\n", strings.Repeat(" ", 45))
		return fetchUsernameInfo(username)
	}
	if !strings.Contains(body, `"username":"`) || !MonitorBans && strings.Contains(body, "Instagram User") {
		return "Unknown", "https://upload.wikimedia.org/wikipedia/commons/a/ac/Default_pfp.jpg"
	}
	return strings.Split(strings.Split(body, `"id":"`)[1], `"`)[0], strings.Split(strings.Split(body, `url":"`)[1], `"`)[0]
}

func fetchUsernameTitle(username string) string {
	request := createFastHTTPRequest("POST", "https://www.instagram.com/ajax/bulk-route-definitions/", false, true, "")
	response := createFastHTTPResponse(false)
	session := formattedBannedSessions[rand.Intn(len(formattedBannedSessions))]

	defer fasthttp.ReleaseRequest(request)
	defer fasthttp.ReleaseResponse(response)

	request.Header.Set("Cookie", session[0] + randomIntString(11))
	request.SetBody([]byte(session[1] + "&route_urls[0]=/" + username + "/routing_namespace=igx_www$a$87a091182d5bd65bcb043a2888004e09&__d=www&__user=0&__a=1&__req=4&__hs=20508.HYP:instagram_web_pkg.2.1...0&dpr=1&__ccg=EXCELLENT&__rev=1033916097&__s=ljw4ew:pajs9i:dbfc2o&__hsi=7610267266820470448&__dyn=7xeUjG1mxu1syUbFp41twpUnwgU7SbzEdF8aUco2qwJw5ux609vCwjE1EE2Cw8G11wBz81s8hwGxu786a3a1YwBgao6C0Mo2swlo8od8-U2zxe2GewGw9a361qw8Xxm16wUwtE1wEbUGdwtUd-2u2J0bS1LwTwKG1pg2fwxyo6O1FwlA3a3zhA6bwIxe6V89F8uwm8jxK0-8KmUhw4rwhEcE4y1Hwj8&__csr=il5MDsY5Yp4EavTORiQLq4b8AObTZ-KGkPLBFdqlKGHKiFXTiy4jyeKRpKEOO9aqWAHhKidhoTSl2RxV4LWyFeAja5FpVbLBhmHF6AAGvUHLUCWiy8jXBWh-QFk8Q5VGKi8BACZe22uq4FpoyaDKmqp124UFaQml5h9Uy4poWAFenF4ByHyEly4uHyUGa-dBABzbgjGibyokx6EWibzaBwi9Ala9zRASQuhpap84m01E4w04k1wed0l9m3S0KV98mwkU0m7XAEE7m0A8K0cng1-O06cwuo7ME0Ldw6rwfG2K0Fq80hS1KwnUB0rE6twF4p3z9dIkazk2Gi0Dd0l81aUm8m9gK8wSw2m48yaBa3KnlU0Xam3W2m8Naxh1speE-bw3Qiw0yOw2BE0aTo0bU80jOCw49w&__hsdp=kwmA3Q684VdXkQn58IQO6Pt43XdR8ijKlcwK63DQ4pzzePk5E9ZB3Xxkg6E7-14DgiwZghCwmF21i2a4polyNFomjakwrAx2exy1Sxl2FjzFUdEbGwywwGi2e2y0E8qG3e6oS0zo5jy9Ef84i4oO17BxF04pwcW0d8w2kVE1IE2Ywam69E4m0G8118K0Oo5u0G81fE1K82bg3-xm2K3G8m4Eeo36who2gx6&__hblp=0dS1CwpoW1hwWwmE4-17w-wKx-jwp87-aGuhaHh44EyibAzGzVoC5o4qfwxzEKujCwExS4awywwGi2e2y3m7VVEiwExGEb9ElKdw8LxGdxm5u8CwYwh8hz8-4U8Voqg3pw_xW3e1Xwj80Ia3ycJqweK1oyU6y1ewPFbw6MwbO0FooCDwZyo-0yWxp7xm1ewqU3FwlU6y3-0jW1xU2gwce5Q1Og24wjEaEybgaUeAtohgeu3u16whpUfo8Hwqohw&__sjsp=kwmA3Q684VdXkQn58IQO6PqA4gIPti4CykO2UoExWwu8AfK3u0hV161xwkpolyNE8u&__comet_req=7&lsd=AdQvnH_jcja9j9MkozqZ9unZ7IA&jazoest=22456&__spin_r=1033916097&__spin_b=trunk&__spin_t=1771903426&__crn=comet.igweb.PolarisProfileRoute"))

	createHostClient(1).Do(request, response)
	body := string(response.Body())

	if !strings.Contains(body, `"error":`) || strings.Contains(body, `"error":true`) {
		log.Printf("Failed to fetch username title %s \n\n", strings.Repeat(" ", 45))
		return fetchUsernameTitle(username)
	}
	return strings.Split(strings.Split(body, `"title":"`)[1], `","`)[0]
}

func sendDiscordWebhook(webhook, template string) {
	request := createFastHTTPRequest("POST", webhook, false, false, "")
	response := createFastHTTPResponse(false)
	
	defer fasthttp.ReleaseRequest(request)
	defer fasthttp.ReleaseResponse(response)

	request.Header.Set("Content-Type", "application/json")
	request.SetBody([]byte(template))

	fasthttp.Do(request, response)
}

func setBiography() {
	request := createFastHTTPRequest("POST", "https://i.instagram.com/api/v1/accounts/set_biography/", true, true, "")
	response := createFastHTTPResponse(false)
	
    request.Header.Set("x-csrftoken", "xSQQ9vbIVUDQG6LKkSgcT6ZDcR6Q6hl1")
	request.Header.Set("cookie", "sessionid=" + activeSession[4])
    request.SetBody([]byte("raw_text=t.me/pClaimed"))

    fasthttp.Do(request, response)
}
