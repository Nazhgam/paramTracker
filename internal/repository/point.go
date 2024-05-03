package repo

import (
	"context"

	"github.com/Nazhgam/paramTracker/internal/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type IPoint interface {
	GetAll(ctx context.Context, id int) ([]models.Point, error)
	GetByID(ctx context.Context, id int) (*models.Point, error)
	Create(ctx context.Context, data models.Point) error
}

type point struct {
	db pgxpool.Pool
}

func NewPoint(db pgxpool.Pool) IPoint {
	return &point{
		db: db,
	}
}

func (p *point) GetAll(ctx context.Context, id int) ([]models.Point, error) {
	rows, err := p.db.Query(ctx, getAllPoint)
	if err != nil {
		return nil, err
	}

	var ponts []models.Point
	for rows.Next() {
		var p models.Point
		if err := rows.Scan(&p.PointID, &p.Name, &p.Alias); err != nil {
			return nil, err
		}

		ponts = append(ponts, p)
	}
	return ponts, nil
}

func (p *point) GetByID(ctx context.Context, id int) (*models.Point, error) {
	var data models.Point

	if err := p.db.QueryRow(ctx, getByIDPoint, id).Scan(
		&data.PointID,
		&data.Name,
		&data.Alias,
	); err != nil {
		return nil, err
	}

	return &data, nil
}

func (p *point) Create(ctx context.Context, data models.Point) error {
	if _, err := p.db.Exec(ctx, insertPoint, data.Name, data.Alias); err != nil {
		return err
	}

	return nil
}
