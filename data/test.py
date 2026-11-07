import requests

cookies = {
    'ds_user_id': '57716718298',
    'ig_did': '3C59EDA3-52C9-4684-A1C6-F4AA8526884F',
    'csrftoken': 'RjQkTp5MVMhDGqXQphhZ13L5WyzfVVal',
    'mid': 'aaYNugAAAAHuKHDQQIYsnIBh9gsD',
    'ps_l': '1',
    'ps_n': '1',
    'datr': 'DUmmaXfl854XU5ENoq9s4wwZ',
    'dpr': '1.25',
    'wd': '1536x178',
    'sessionid': '57716718298%3AuQR0No72cp7nFM%3A25%3AAYgtFBzVig1Y6gO-8a9tX3MAt56-jnLjeA_qJITepA',
    'rur': '"NHA\\05457716718298\\0541804128880:01fea6297f2f6bcaf9665a1b9bdb22553fb2e5a1898d6c25928c8dda5ee9402f45d1df19"',
}

headers = {
    'accept': '*/*',
    'accept-language': 'en-US,en;q=0.9',
    'content-type': 'application/x-www-form-urlencoded',
    'origin': 'https://www.instagram.com',
    'priority': 'u=1, i',
    'referer': 'https://www.instagram.com/poshy/shop/?__coig_login=1',
    'sec-ch-prefers-color-scheme': 'dark',
    'sec-ch-ua': '"Not:A-Brand";v="99", "Google Chrome";v="145", "Chromium";v="145"',
    'sec-ch-ua-full-version-list': '"Not:A-Brand";v="99.0.0.0", "Google Chrome";v="145.0.7632.117", "Chromium";v="145.0.7632.117"',
    'sec-ch-ua-mobile': '?0',
    'sec-ch-ua-model': '""',
    'sec-ch-ua-platform': '"Windows"',
    'sec-ch-ua-platform-version': '"19.0.0"',
    'sec-fetch-dest': 'empty',
    'sec-fetch-mode': 'cors',
    'sec-fetch-site': 'same-origin',
    'user-agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/145.0.0.0 Safari/537.36',
    'x-asbd-id': '359341',
    'x-fb-lsd': 'rLkrUBPU19DKdk6cwqzsOy',
    'x-ig-d': 'www',
    # 'cookie': 'ds_user_id=57716718298; ig_did=3C59EDA3-52C9-4684-A1C6-F4AA8526884F; csrftoken=RjQkTp5MVMhDGqXQphhZ13L5WyzfVVal; mid=aaYNugAAAAHuKHDQQIYsnIBh9gsD; ps_l=1; ps_n=1; datr=DUmmaXfl854XU5ENoq9s4wwZ; dpr=1.25; wd=1536x178; sessionid=57716718298%3AuQR0No72cp7nFM%3A25%3AAYgtFBzVig1Y6gO-8a9tX3MAt56-jnLjeA_qJITepA; rur="NHA\\05457716718298\\0541804128880:01fea6297f2f6bcaf9665a1b9bdb22553fb2e5a1898d6c25928c8dda5ee9402f45d1df19"',
}

# data = 'route_urls[0]=/poshy/shop/&routing_namespace=igx_www%24a%2487a091182d5bd65bcb043a2888004e09&__d=www&__user=0&__a=1&__req=5&__hs=20515.HYP%3Ainstagram_web_pkg.2.1...0&dpr=1&__ccg=EXCELLENT&__rev=1034367274&__s=dcht5y%3A23eh27%3Ayrrwor&__hsi=7612890238644497813&__dyn=7xeUjG1mxu1syUbFp41twpUnwgU29zEdEc8co2qwJwbS0E8hw2nVE4W0qa0FE2awt81s8hwGwQwoEcE7O2l0Fwqo31w9O0H8jwae4UaEW2G0AEco5G1Wxfxm16wUwtE1wEbUGdG1QwTwFwIwbS1LwTwKG1pg2Xwr86C1mg6LhAq4rwh8nAwHxW1owLwfybyohw5yweu&__csr=gB5gz7hIp1h5EysDEZ-xquGEDb9F9bqGHVp9Vk8WyWGfzbx2e8lzFEDzAmAayp8aF8O4pGwgE9bLGmu7Ukx2dx6VoCaK-5pWUGidVrxiivgcawGx6dzoGmleFo56ubybGieU8k4-FpecyoLwBXxyquHK7o26w0Oeo06Di0pm0tei4okwmF45U1g8mw0t-Uc8O04Am05ME1i5w6rw5xw11la0ji5Khe480h1o5-064k0EA0mi0RK0ixG08oGU8GZaJ1a0Hm048U2Nw3lo09uE0gKw098h0eWi1hwl80iQw3FXU&__hsdp=l3k9A5M-geE4gapBkBPOpsOGaDb4jNex7gMyzKHqx1xLg11o6lwz70e602yy06TE2_wWw6CU1yUc87y04vomw5Uwey&__hblp=0e62-0BU1X8eoG1ywsUdE7e1ExK3y483kADwam085zaw7Rw2AU1g8fU16UqxS3B0sE0Qe5E2_y8co520Ho0zm321yxm1aw2P83OwsEmwiUC2q0A8f8bo3ozU24w&__sjsp=l3k9A5M-geE4gapBkBPOpsOGaDbezEhEhQc8EXGha466U1286lwz70&__comet_req=7&jazoest=26519&lsd=XIUBoGmmr4H-A45IOq_AgZ&__spin_r=1034367274&__spin_b=trunk&__spin_t=1772514134&__crn=comet.igweb.PolarisCommerceStorefrontRoute&fb_dtsg=NAftfjrui6PUl_l5tEZlm5EnzQCynZ7EnWlsxMooGtnBHKUWVOg8-GA%3A17843676607167008%3A1772592872'
data = 'route_urls[0]=%2Fposhy%2Fshop%2F%3F__coig_login%3D1&routing_namespace=igx_www%24a%2487a091182d5bd65bcb043a2888004e09&__d=www&__user=0&__a=1&__req=5&__hs=20516.HYP%3Ainstagram_web_pkg.2.1...0&dpr=1&__ccg=EXCELLENT&__rev=1034413161&__s=155cwv%3Anvaxst%3At7az6s&__hsi=7613228449832872992&__dyn=7xeUjG1mxu1syUbFp41twpUnwgU29zEdEc8co2qwJwbS0E8hw2nVE4W0qa0FE2awt81s8hwGwQwoEcE7O2l0Fwqo31w9O0H8jwae4UaEW2G0AEco5G1Wxfxm16wUwtE1wEbUGdG1QwTwFwIwbS1LwTwKG1pg2Xwr86C1mg6LhAq4rwh8nAwHxW1owLwfybyohw5yweu&__csr=hIj91p34lvbtRkll8gYCAATiQzylGql4gOl5xqbzrCGVoyczpVAWDDBDhoa9UhAgd8gDx-VoSU-GxCl7zUZ5wwAx3KqjwqQEWcUkUlBK6oG14hKifBCyokWxlbogzoy3C6Ufo-0tu0crwlU6R120mWvg01vgmEkw4QwkuU0C10cZ4w1eG06r207TggxZ0aUxox2Vi0ywyw9vwgA08R81Rydw9xw6HwoE4Mi1nx21z-0NU5dw2f80X6m01Dgw0joE0nPw&__hsdp=l2QbB0A84gapaekoCqRlauqkIY9V18IwybAkEZwyg2sw9le2phw1m809rU0qkwcW3q0o_woE3YDwbm0fXwOwqU0wC&__hblp=0e20PE9E1Uopw9W14x28xG1qwqVE3kwxAwbem0P80Pe09Qw4SwSweynGE5-0aWw4ky82DxSu5Umwba1fwf60_9U2Rwj80Fa0U8eEcEcokCxS6820wby0JEa8&__sjsp=l2QbB0A84gapaekoCqRlauqkI_4PV18IwybAkwgo30w9le2phw055y&__comet_req=7&__spin_r=1034413161&__spin_b=trunk&__spin_t=1772592880&__crn=comet.igweb.PolarisCommerceStorefrontRoute'


response = requests.post('https://www.instagram.com/ajax/bulk-route-definitions/', cookies=cookies, headers=headers, data=data)
print(len(response.text))