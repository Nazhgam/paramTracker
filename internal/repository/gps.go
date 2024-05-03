package repo

import (
	"context"

	"github.com/Nazhgam/paramTracker/internal/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type IGps interface {
	Create(ctx context.Context, data models.GPS) error
	GetByPointID(ctx context.Context, pointID int) (*models.GPS, error)
}

type gps struct {
	db pgxpool.Pool
}

func NewGps(db pgxpool.Pool) IGps {
	return &gps{db: db}
}

func (g *gps) Create(ctx context.Context, data models.GPS) error {

	if _, err := g.db.Exec(ctx, insertGps, data.PointID, data.PointGpsID, data.Lat, data.Lon, data.Speed, data.Time); err != nil {
		return err
	}

	return nil
}

func (g *gps) GetByPointID(ctx context.Context, pointID int) (*models.GPS, error) {

	var data models.GPS

	if err := g.db.QueryRow(ctx, getGpsByPointID, pointID).Scan(
		&data.GpsID,
		&data.PointID,
		&data.PointGpsID,
		&data.Lat,
		&data.Lon,
		&data.Speed,
		&data.Time,
	); err != nil {
		return nil, err
	}

	return &data, nil
}
