package server

import (
	"log/slog"
	"net/http"
	"os"
	"sync"

	"github.com/ElRAS1/wb_L2/develop/dev11/internal/event"
	"github.com/ilyakaznacheev/cleanenv"
)

type config struct {
	Port string `yaml:"Addr" env-default:"8080"`
}

type Server struct {
	mu     sync.RWMutex
	data   map[string]event.Event
	Server *http.Server
	Logger *slog.Logger
}

func NewSrv() *Server {
	return &Server{
		data:   make(map[string]event.Event),
		mu:     sync.RWMutex{},
		Server: &http.Server{},
		Logger: slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})),
	}
}

func (s *Server) NewServer() (*Server, error) {
	cfg := config{}
	s = NewSrv()
	slog.SetDefault(s.Logger)
	err := cleanenv.ReadConfig("config/config.yaml", &cfg)
	if err != nil {
		return s, err
	}
	s.Server = &http.Server{Addr: cfg.Port, Handler: s.Router()}
	return s, nil
}
