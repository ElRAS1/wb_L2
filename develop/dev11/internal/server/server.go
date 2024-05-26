package server

import (
	"log"
	"net/http"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Port string `yaml:"Addr" env-default:"8080"`
}

var cfg Config

func NewServer() *http.Server {
	err := cleanenv.ReadConfig("config/config.yml", &cfg)
	if err != nil {
		log.Fatalln(err)
	}
	srv := &http.Server{Addr: cfg.Port, Handler: Router()}
	return srv
}
