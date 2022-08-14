package handler

import (
	"fmt"
	"github.com/VadimGossip/tcpBsonServerExample/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"net"
	"time"
)

type Handler struct {
	connectionHandler ConnectionHandler
}

type ConnectionHandler interface {
	HandleConnection(conn net.Conn, handlerFunc func(conn net.Conn))
	ReadChan() []byte
	WriteChan(conn net.Conn, msgBody []byte)
}

func NewHandler(connectionHandler ConnectionHandler) *Handler {
	return &Handler{connectionHandler: connectionHandler}
}

func (h *Handler) routeRequest(req domain.RouteRequest) domain.RouteResponse {
	var res domain.RouteResponse
	res.SendTime = req.SendTime
	res.RouteStartTime = time.Now()

	if req.RouteDuration > 0 {
		time.Sleep(req.RouteDuration)
	} else {
		res.Err = fmt.Sprintf("error empty route duration")
	}

	res.RouteEndTime = time.Now()

	return res
}

func (h *Handler) handleRequest(conn net.Conn) {
	for {
		msgBody := h.connectionHandler.ReadChan()
		var req domain.RouteRequest
		var res domain.RouteResponse
		var responseBytes []byte
		err := bson.Unmarshal(msgBody, &req)
		if err != nil {
			res.Err = fmt.Sprintf("error occurred while unmarshal request to router: %s", err.Error())
		}

		if err == nil {
			if req.RequestType == "route" {
				res = h.routeRequest(req)
			} else {
				res.Err = fmt.Sprintf("unknown request type %s", req.RequestType)
			}
			responseBytes, err = bson.Marshal(res)
			if err != nil {
				responseBytes, _ = bson.Marshal(domain.RouteResponse{Err: fmt.Sprintf("error occurred while marshal response from router: %s", err.Error())})
			}
		}
		h.connectionHandler.WriteChan(conn, responseBytes)
	}
}

func (h *Handler) HandleConnection(conn net.Conn) {
	h.connectionHandler.HandleConnection(conn, h.handleRequest)
}
