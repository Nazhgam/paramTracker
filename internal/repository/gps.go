package repo

import (
	"context"

	"github.com/Nazhgam/paramTracker/internal/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type IGps interface {
	Create(ctx context.Context, data models.GPS) error
	GetAll(ctx context.Context) ([]models.GPS, error)
}

type gps struct {
	db *pgxpool.Pool
}

func NewGps(db *pgxpool.Pool) IGps {
	return &gps{db: db}
}

func (g *gps) Create(ctx context.Context, data models.GPS) error {

	if _, err := g.db.Exec(ctx, insertGps, data.PointID, data.PointGpsID, data.Lat, data.Lon, data.Speed, data.Time); err != nil {
		return err
	}

	return nil
}

func (g *gps) GetAll(ctx context.Context) ([]models.GPS, error) {
	var (
		result []models.GPS
		data   models.GPS
	)

	rows, err := g.db.Query(ctx, getAllGps)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		data = models.GPS{}

		if err := rows.Scan(
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

		result = append(result, data)
	}

	return result, nil
}
