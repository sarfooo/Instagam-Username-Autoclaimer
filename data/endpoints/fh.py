import requests
import random
import uuid

headers = {
    'content-type': 'application/x-www-form-urlencoded; charset=UTF-8',
    'ig-intended-user-id': str(random.randint(1, 10000000000)),
    'ig-u-ds-user-id': str(random.randint(1, 10000000000)),
    'user-agent': 'Instagram 381.1.2.26.83 (iPhone14,5; iOS 16_3_1; en_US; en; scale=3.00; 1170x2532; 737297623) AppleWebKit/420+',
    'x-mid': 'ZDHqpQAAAAEP2s27NnTYI9cmO1UM',
    'connection': 'close',
}


data = {
    # 'signed_body': 'SIGNATURE.{"waterfall_id":"681bd048453c4c218b8608fc3fdcfbf0","dryrun":"true","big_blue_token":"EAACgZC2tvhrgBQsnb090DIFI73kzOmFhaQ4VnFv8q16uymBVQYv16dtqrzmpnt5nrGpn2Iw7ZC5YWIsOh8tVwheGtxjxeYJ260xHP7WkG4FEX3aVRzz2qO0HrLn2hncRY2En3agHIfS5NNPKsXnqRCZBK4sOgFh55LVXDckbeo6lojW7lIk7J6VAHFHrArYqwZDZD","device_id":"EB1507B6-8959-46EA-958D-ABFEE4C9AA6D","att_permission_status":"3","surface":"login_continue_button","has_seen_aart_on":"1","fb_reg_flag":"false","force_signup_with_fb_after_cp_claiming":"false","allow_contacts_sync":"false","phone_id":"18146D2B-7691-4B12-A5C6-2932CBA2EAC5","username":"god"}',
    'signed_body': 'SIGNATURE.{"waterfall_id":"' + str(uuid.uuid4()) + '","dryrun":"false","fb_access_token":"EAABwzLixnjYBQuM8qumZCEuVNMO3ZCbZBQCr61Wnlc7feQBHsgcerD6gyZCyUvxZBK15LB7yrJB6IIUApkdF3Ev5CroYVwKlIOSZBdlvOYN4kZClEaUtPYJm83ZA6lmKM9ZB1MkiPYjZBSYU5nTFMhzs5kxStRr8RaB7f3zCG5j9JiXPtQTvThZAbwXPH4UgrXY3awemqqy6ZA6ReQgXZAdknJilD0ZAqHo2L3GujGzviWnXu72NtfK8CeMwZDZD","device_id":"' + str(uuid.uuid4()) + '","att_permission_status":"3","surface":"login_continue_button","has_seen_aart_on":"1","fb_reg_flag":"false","force_signup_with_fb_after_cp_claiming":"false","allow_contacts_sync":"false","phone_id":"' + str(uuid.uuid4()) + '","username":"god"}',
}

# print('SIGNATURE.{"waterfall_id":"' + str(uuid.uuid4()) + '","dryrun":"false","fb_access_token":"EAABwzLixnjYBQjK2nzDd1nTKJisRM6w8TtrQ9beMxPInwPdSgml3T3ZA6JKzXJCuUSpKMDEIcEXulMIBVogzSE2b3cZBXtH3tXe3P0s2FSGLDq5yJVE1LftGl8MBAOy9Lkki9JXZB38IXt2N2VGa4ExDbYmJ3XaFt2lJiPTiF3FVmZBZCjkjgPyZBE4BW2jHMx1FmIQI8wZCRR4XBZAItZB07KEemW6l8IPFPNjh6BgIdNmPTYJ1DQWIKNNUZD","device_id":"' + str(uuid.uuid4()) + '","att_permission_status":"3","surface":"login_continue_button","has_seen_aart_on":"1","fb_reg_flag":"false","force_signup_with_fb_after_cp_claiming":"false","allow_contacts_sync":"false","phone_id":"' + str(uuid.uuid4()) + '","username":"god"}')
response = requests.post('https://www.threads.net/api/v1/fb/facebook_signup/', headers=headers, data=data)
# response = requests.post('https://i.instagram.com/api/v1/web/accounts/username_suggestions/', headers=headers, data=data, proxies = {"http": "http://h91p68ig35-package-ipv6:8hna38jpvm@ipv6.enigmaproxy.net:30", "https": "http://h91p68ig35-package-ipv6:8hna38jpvm@ipv6.enigmaproxy.net:30"})
print(response.text)