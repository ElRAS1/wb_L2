package server

import (
	"log/slog"
	"net/http"
	"os"
	"sync"

	"github.com/ElRAS1/wb_L2/develop/dev11/event"
	"github.com/ilyakaznacheev/cleanenv"
)

type config struct {
	Port string `yaml:"Addr" env-default:"8080"`
}

type Server struct {
	mu     sync.Mutex
	data   *event.Data
	Server *http.Server
	Logger *slog.Logger
}

func NewSrv() *Server {
	return &Server{
		data:   event.NewData(),
		mu:     sync.Mutex{},
		Server: &http.Server{},
		Logger: slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})),
	}
}

func NewServer() (*Server, error) {
	cfg := config{}
	srv := NewSrv()
	err := cleanenv.ReadConfig("config/config.yaml", &cfg)
	if err != nil {
		return srv, err
	}
	srv.Server = &http.Server{Addr: cfg.Port, Handler: srv.Router()}
	return srv, nil
}
