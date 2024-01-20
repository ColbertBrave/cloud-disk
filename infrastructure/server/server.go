package server

import (
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/cloud-disk/infrastructure/auth"
	"github.com/cloud-disk/infrastructure/log"
)

var server *Server

type Server struct {
	addr       string
	httpServer *http.Server
	ginEngine  *gin.Engine
}

type Option func(engine *gin.Engine)

func InitServer(addr string) {
	if addr == "" {
		log.Error("the server addr is empty")
		return
	}

	gin.ForceConsoleColor()
	gin.SetMode(gin.ReleaseMode)

	ginEngine := gin.New()
	ginEngine.Use(gin.Recovery())
	ginEngine.Use(GetCostTimeOfRequest())
	ginEngine.Use(Authenticate())

	server = &Server{
		addr: addr,
		httpServer: &http.Server{
			Addr:    addr,
			Handler: ginEngine,
		},
		ginEngine: ginEngine,
	}
}

func Run() {
	if err := server.Start(); err != nil {
		log.Error("run the http server error|%s", err)
		return
	}
	log.Info("successfully run the http server")
}

func GetCostTimeOfRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		costTime := time.Since(startTime)
		log.Info("%s|%s|cost time %d ms", c.Request.Method, c.Request.URL, costTime)
	}
}

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := auth.VerifyRequest(auth.Auth, c.Request)
		if err != nil {
			log.Error("verify request error:%s", err)
			return
		}
		c.Next()
	}
}

func (s *Server) Start() error {
	listener, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}

	err = s.httpServer.Serve(listener)
	if err != nil {
		return err
	}
	return nil
}

func Close() {
	if err := server.httpServer.Close(); err != nil {
		log.Error("close http server err|%s", err)
		return
	}
	log.Info("the http server is closed")
}
