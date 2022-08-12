package tcp

import (
	"github.com/VadimGossip/tcpBsonServerExample/internal/config"
	"github.com/VadimGossip/tcpBsonServerExample/internal/transport/tcp"
	"github.com/sirupsen/logrus"
	"net"
	"strconv"
)

type Server struct {
	config config.NetServerConfig
}

func NewServer(config config.NetServerConfig) *Server {
	return &Server{config: config}
}

func (s *Server) Run(handler *tcp.Handler) error {
	ln, err := net.Listen("tcp", s.config.Host+":"+strconv.Itoa(s.config.Port))
	if err != nil {
		logrus.Errorf("error occurred while running tcp server: %s", err.Error())
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			logrus.Errorf("error occurred while running tcp server: %s", err.Error())
		}
		go handler.HandleConnection(conn)
	}
}
