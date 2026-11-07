import requests

response = requests.get("https://api.seeip.org/jsonip?", proxies = {
    "http": "http://6t2a1kxgl4-package-ipv6:wu5btqrgqf@ipv6.enigmaproxy.net:30",
    "https": "http://6t2a1kxgl4-package-ipv6:wu5btqrgqf@ipv6.enigmaproxy.net:30"
})
print(response)