package repo

import (
	"context"

	"github.com/Nazhgam/paramTracker/internal/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type IPoint interface {
	GetByAlias(ctx context.Context, alias string) (*models.Point, error)
	Create(ctx context.Context, data models.Point) error
	GetLastID(ctx context.Context, pointID int) (*models.LastIDResult, error)
}

type point struct {
	db *pgxpool.Pool
}

func NewPoint(db *pgxpool.Pool) IPoint {
	return &point{
		db: db,
	}
}

func (p *point) GetByAlias(ctx context.Context, alias string) (*models.Point, error) {
	var data models.Point

	if err := p.db.QueryRow(ctx, getByAliasPoint, alias).Scan(
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

func (p *point) GetLastID(ctx context.Context, pointID int) (*models.LastIDResult, error) {
	var r models.LastIDResult
	if err := p.db.QueryRow(ctx, getLastIDByAlias, pointID).Scan(
		&r.PointGpsID,
		&r.ParamsStr.ID,
		&r.ParamsStr.StringPointID,
		&r.ParamsStr.PointID,
		&r.ParamsStr.Channel,
		&r.ParamsStr.Value,
		&r.ParamsStr.Time,
		&r.IntPointID,
	); err != nil {
		return nil, err
	}
	return &r, nil
}
