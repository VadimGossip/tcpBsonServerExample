package app

import (
	"github.com/VadimGossip/tcpBsonServerExample/internal/config"
	"github.com/VadimGossip/tcpBsonServerExample/internal/server/tcp"
	"github.com/VadimGossip/tcpBsonServerExample/internal/transport/tcp"
	"github.com/VadimGossip/tcpConHandler"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(configDir string) {
	cfg, err := config.Init(configDir)
	if err != nil {
		logrus.Errorf("Config initialization error %s", err)
	}
	ch := tcpConHandler.NewConnectionHandler(100, 100, time.Second*2)
	hlr := handler.NewHandler(ch)
	srv := tcp.NewServer(cfg.ServerListenerTcp)
	go func() {
		if err := srv.Run(hlr); err != nil {
			logrus.Fatalf("error occured while running tcp server: %s", err.Error())
		}
	}()
	logrus.Infof("Tcp server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Infof("Tcp server stoped")
}
