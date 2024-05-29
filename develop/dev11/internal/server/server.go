package server

import (
	"net/http"
	"sync"

	"github.com/ElRAS1/wb_L2/develop/dev11/event"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Port string `yaml:"Addr" env-default:"8080"`
}

type Server struct {
	data   *event.Data
	server *http.Server
	mu     sync.Mutex
}

func NewSrv() *Server {
	return &Server{
		data:   event.NewData(),
		mu:     sync.Mutex{},
		server: &http.Server{}}
}

func NewServer() (*http.Server, error) {
	cfg := Config{}
	srv := NewSrv()
	err := cleanenv.ReadConfig("config/config.yaml", &cfg)
	if err != nil {
		return srv.server, err
	}
	srv.server = &http.Server{Addr: cfg.Port, Handler: srv.Router()}
	return srv.server, nil
}
