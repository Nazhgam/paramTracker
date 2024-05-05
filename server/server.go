package server

import (
	"context"
	"log"

	"github.com/Nazhgam/paramTracker/internal/controller"
	repo "github.com/Nazhgam/paramTracker/internal/repository"
	"github.com/Nazhgam/paramTracker/internal/service"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/echo/v4"
)

type Server struct {
	log     *log.Logger
	cfg     Config
	handler controller.Handler
}

func ServerEnv(log *log.Logger) (Server, error) {
	cfg, err := NewConfig("config/")
	if err != nil {
		log.Printf("error while get Config datas. error is: %s", err)
		return Server{}, err
	}

	pgx, err := pgxpool.Connect(context.Background(), cfg.Database.ConnStr())
	if err != nil {
		log.Printf("error while parseConfig for pgxpool. error: %s", err)
		return Server{}, err
	}
	db := repo.New(pgx)
	srv := service.New(log, db)
	handler := controller.New(log, srv)

	return Server{
		log:     log,
		handler: handler,
		cfg:     cfg,
	}, nil
}

func (s Server) Start() {
	// Создаем экземпляр Echo
	e := echo.New()
	s.Route(e)
	// Запускаем веб-сервер
	s.log.Fatal(e.Start(s.cfg.ToString()))
}
