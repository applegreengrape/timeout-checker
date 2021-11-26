# timeout-checker
A simple wait and response api to test retry& timeout. Yep! Everyone loves retry and timeout test 🤠

## Flags
-wait: `int` set how many seconds to wait before response, default is 3 second

-port: `int` set api server port, default is 8080

-path: `string` set api path, default is "/api/echo"

-method: `string` http method, default is "POST"

-payload: `string` path to the payload json file

## Install via brew 
```
$ brew tap applegreengrape/timeout-checker https://github.com/applegreengrape/timeout-checker
$ brew install timeout-checker
```
server side:
```
$ checker-api -wait 3 -port 8081
2021/11/22 15:21:30 wait 3 second
2021/11/22 15:21:30 received
2021/11/22 15:21:33 returned
2021/11/22 15:21:42 wait 3 second
2021/11/22 15:21:42 received
2021/11/22 15:21:45 returned
```

client side:
```
$ curl http://localhost:8081/healthcheck
{"status":"up"}

$ curl http://localhost:8081/healthcheck
{"status":"up"}
```

## Customize api response
server side:
```
$ go run main.go -path '/v1/api/echo' -payload echo.json
2021/11/25 16:25:43 wait 3 second
2021/11/25 16:25:43 received
2021/11/25 16:25:46 returned
```

client side:
```
$ curl -X POST  http://localhost:8080/v1/api/echo
{
    "_id": "619fb7e74eb6e971b55f410a",
    "index": 0,
    "guid": "32316e07-db6a-437d-887c-6d0c34ee805a",
    "isActive": true,
    "balance": "$2,259.93",
    "picture": "http://placehold.it/32x32",
    "age": 36,
    "eyeColor": "blue",
    "name": "Jan Crane",
    "gender": "female",
    "company": "COMVEY",
    "email": "jancrane@comvey.com",
    "phone": "+1 (953) 460-3063",
    "address": "962 Lorraine Street, Hayden, Texas, 2598",
    "about": "Est aute exercitation laborum ipsum ad non sit laborum officia. Exercitation cupidatat consequat ea id aute cillum minim veniam Lorem cupidatat consectetur sit esse excepteur. Ex fugiat tempor dolore Lorem et nulla veniam quis in esse. Non est deserunt consequat aute do pariatur laborum amet labore eu ex culpa sit ipsum. Laboris duis nulla voluptate proident. Fugiat sint anim minim anim cupidatat veniam nulla dolore qui labore aute fugiat anim ipsum.\r\n",
    "registered": "2018-10-29T11:51:37 -00:00",
    "latitude": 65.197781,
    "longitude": 46.290158,
    "tags": [
      "minim",
      "ipsum",
      "tempor",
      "in",
      "sit",
      "culpa",
      "dolore"
    ],
    "friends": [
      {
        "id": 0,
        "name": "Camille Hopkins"
      },
      {
        "id": 1,
        "name": "Alyssa Contreras"
      },
      {
        "id": 2,
        "name": "Burton Haney"
      }
    ],
    "greeting": "Hello, Jan Crane! You have 6 unread messages.",
    "favoriteFruit": "strawberry"
  }
```
So you should be able to customiz the api path and response
