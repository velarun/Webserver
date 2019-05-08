# Webserver

Pre-requistie:

Golang:
1. Golang should be installed, Go version should be > 1.10.1
2. MysqlDB should be installed. Version should be > 8.0.15
3. Use "config.json" file change username, password, db ip accordingly for connectivity
    
Python:
4. Python should be installed, python version should be > 3.5.1
5. Python scripts for webscraper available in the webscraping/htmlparser.py

How to run:

$make

Run command:
$./webserver

Logs:
Server will start at http://localhost:8000/
2019/04/25 19:07:29 Read config.json file successfully
2019/04/25 19:07:29 Database connected successfully
2019/04/25 19:07:29 Default Database Loaded successfully
Routes are Loded.

Http Get & POST request:

curl -X POST \
  http://127.0.0.1:8000/tray \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -H 'postman-token: ef29b1d5-4902-ccad-b931-42923c96d5e1' \
  -d '{
  "Id": {
        "Int64": 2,
        "Valid": true
    },
  "Timestmp" : "2019-03-15 00:00:00",
  "User_id" : "test",
  "Is_logged_in" : "1",
  "Missing_tray_title" : "hai, how are you, bye",
  "Added_tray_title" : "Good, bad, worst"
}'

Using HTTP:
http://localhost:8000/tray/{id} -- row id to get json body

Python Webscraper of hotstart webpage:

***Note: Change hotstar auth key accordingly***

$ python3 htmlparser.py 
List of Tray titles before login:
Popular in Sports
Popular Channels
Popular Shows
Popular Movies
Popular in Tamil
Login Page: <Response [200]>
List of Tray titles after login:
Popular in Sports
Popular Channels
Popular Shows
Popular Movies
Popular in Tamil
Missing Tray

Added Tray

Post to Server: <Response [200]>





