package domain

import "time"

type RouteRequest struct {
	RequestType   string        `bson:"requestType"`
	MsgId         int           `bson:"msgId"`
	SendTime      time.Time     `bson:"sendTime"`
	RouteDuration time.Duration `bson:"routeDur"`
}

type RouteResponse struct {
	Err            string    `bson:"err,omitempty"`
	SendTime       time.Time `bson:"sendTime,omitempty"`
	RouteStartTime time.Time `bson:"routeBegin,omitempty"`
	RouteEndTime   time.Time `bson:"routeEnd,omitempty"`
}
