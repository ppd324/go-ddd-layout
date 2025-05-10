package http

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/ppd324/go-ddd-layout/app/interfaces/http/router"
	"github.com/ppd324/go-ddd-layout/conf"
	"log"
	"strconv"
)

type HttpServer struct {
	cfg     *conf.AppConfig
	routers []router.Router
}

var ProviderSet = wire.NewSet(NewHttpServer)

func NewHttpServer(cfg *conf.AppConfig, routers []router.Router) *HttpServer {
	return &HttpServer{
		cfg:     cfg,
		routers: routers,
	}

}

func (s *HttpServer) Start() {
	gin.SetMode(gin.ReleaseMode)
	g := gin.New()
	g.Use(gin.Recovery())
	for _, r := range s.routers {
		r.Mount(g)
	}
	log.Println("http server start,run env", s.cfg.Server.Env, "will serving HTTP on", s.cfg.Server.Addr, ":", s.cfg.Server.Port)
	err := g.Run(s.cfg.Server.Addr + ":" + strconv.Itoa(s.cfg.Server.Port))
	if err != nil {
		panic(err)
	}
}
