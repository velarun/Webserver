import requests 
from bs4 import BeautifulSoup
import re

x = 5
url = 'https://uk.hotstar.com/subscribe/get-started'
signin = "https://api.hotstar.com/gb/aadhar/v2/web/gb/user/login"
post = "http://127.0.0.1:8000/tray"
userid = "testprod-999@hotstar.com"
pwd = "test"

response = requests.get(url)
html_soup = BeautifulSoup(response.text, 'html.parser')
movie_containers = html_soup.find_all('script')
val = str(movie_containers[0])
m = re.findall('\"title\"\:\"([A-Za-z ]+)\"', val)
tray1 = []
for i in m:
    if re.search(".*This Content is Best Experienced.*", i):
        pass
    else:
        tray1.append(i)

print("List of Tray titles before login:")
for i in range(0, x):
    print(tray1[i])

session = requests.session()
session.headers.update({'user-agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.108 Safari/537.36'})
session.headers.update({'accept': '*/*'})
session.headers.update({'content-type': 'application/json'})
session.headers.update({'hotstarauth': 'st=1557301246~exp=1557307246~acl=/*~hmac=b639173ff70d9a4da933d391a7439112830c8db9ffc676ad8ceab197c3d5c6e9'})
session.headers.update({'origin': 'https://uk.hotstar.com'})
root_page = session.post(signin, json={"isProfileRequired":False,"userData":{"deviceId":"gdprDeviceId","password": pwd,"username": userid,"usertype":"email"},"verification":{}})
print("Login Page:", root_page)

response = requests.get(url)
html_soup = BeautifulSoup(response.text, 'html.parser')
movie_containers = html_soup.find_all('script')
val = str(movie_containers[0])
m = re.findall('\"title\"\:\"([A-Za-z ]+)\"', val)
tray2 = []
for i in m:
    if re.search(".*This Content is Best Experienced.*", i):
        pass
    else:
        tray2.append(i)

print("List of Tray titles after login:")
for i in range(0, x):
    print(tray2[i])

missing_tray = ""
added_tray = ""

for i in tray1:
    if i not in tray2:
        missing_tray = missing_tray + i + ","

for i in tray2:
    if i not in tray1:
        added_tray = added_tray + i +  ","

print("Missing Tray")
print(missing_tray)

print("Added Tray")
print(added_tray)


session = requests.session()
json_body = { "Id": { "Int64": 2, "Valid": True }, "Timestmp" : "2019-03-15 00:00:00", "User_id" : userid, "Is_logged_in" : "1", "Missing_tray_title" : missing_tray, "Added_tray_title" : added_tray}
root_page = session.post(post, json=json_body)
print("Post to Server:", root_page)