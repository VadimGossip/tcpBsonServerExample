package handler

import (
	"fmt"
	"github.com/VadimGossip/tcpBsonServerExample/internal/domain"
	"github.com/VadimGossip/tcpConHandler"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"net"
	"time"
)

type Handler struct {
	connectionHandler *tcpConHandler.ConnectionHandler
}

func NewHandler() *Handler {
	return &Handler{}
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

func (h *Handler) handleRequest() {
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
		h.connectionHandler.WriteChan(responseBytes)
	}
}

func (h *Handler) HandleConnection(conn net.Conn) {
	h.connectionHandler = tcpConHandler.NewConnectionHandler(conn, 2*time.Second, h.handleRequest, 100, 100)
	err := h.connectionHandler.HandleConnection()
	if err != nil {
		logrus.Errorf("Error while handle connection %s", err)
	}
}
