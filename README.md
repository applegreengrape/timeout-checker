# timeout-checker
A simple wait and response api to test retry& timeout. Yep! Everyone loves retry and timeout test 🤠

Install via brew 
```
$ brew tap applegreengrape/timeout-checker https://github.com/applegreengrape/timeout-checker
$ brew install timeout-checker
```
server side:
```
$ checker-api -wait 5 -port 8081
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
