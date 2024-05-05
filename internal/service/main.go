package service

import (
	"context"
	"log"

	"github.com/Nazhgam/paramTracker/internal/models"
	repo "github.com/Nazhgam/paramTracker/internal/repository"
)

type Service interface {
	CreateGps(ctx context.Context, alias string, data models.GPS) error
	GetAllGps(ctx context.Context) ([]models.GPS, error)

	CreatePoint(ctx context.Context, point models.Point) error
	GetLastIDByAlias(ctx context.Context, alias string) (*models.LastIDResult, error)

	CreateParamStr(ctx context.Context, alias string, data models.ParamStr) error
	CreateParamInt(ctx context.Context, alias string, data models.ParamInt) error
}

type service struct {
	log     *log.Logger
	storage repo.Repo
}

func New(log *log.Logger, storage repo.Repo) Service {
	return &service{
		log:     log,
		storage: storage,
	}
}
