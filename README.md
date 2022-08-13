# tcpBsonServerExample
An example of a basic TCP server which can receive BSON packets, process them in several threads 
(I use a special thread handler https://github.com/VadimGossip/tcpConHandler) and return the responses. 
Reading and writing are done independently.

In addition to this project I made a request generator which can also be used for mastering client-server interaction.
https://github.com/VadimGossip/tcpBsonServerReqGenerator

Request to server example:
```
type RouteRequest struct {
  RequestType   string         `bson:"requestType"`
  MsgId	        int            `bson:"msgId"`
  SendTime      time.Time      `bson:"sendTime"`
  RouteDuration time.Duration  `bson:"routeDur"`
}
```
The RouteDuration parameter is needed to emulate some work on the server side. 

Server response example:
```
type RouteResponse struct {
  Err            string     `bson:"err,omitempty"`
  SendTime       time.Time  `bson:"sendTime,omitempty"`
  RouteStartTime time.Time  `bson:"routeBegin,omitempty"`
  RouteEndTime   time.Time  `bson:"routeEnd,omitempty"`
}
```
