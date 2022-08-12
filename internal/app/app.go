package app

import (
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"tpcClientServerStand/config"
	server "tpcClientServerStand/internal/server/tcp"
	"tpcClientServerStand/internal/transport/tcp"
)

func Run(configDir string) {
	cfg, err := config.Init(configDir)
	if err != nil {
		logrus.Errorf("Config initialization error %s", err)
	}
	handler := tcp.NewHandler()
	srv := server.NewServer(cfg.ServerListenerTcp)
	go func() {
		if err := srv.Run(handler); err != nil {
			logrus.Fatalf("error occured while running tcp server: %s", err.Error())
		}
	}()
	logrus.Infof("Tcp server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Infof("Tcp server stoped")
}
