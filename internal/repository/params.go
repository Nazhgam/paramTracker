package repo

import (
	"context"

	"github.com/Nazhgam/paramTracker/internal/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type IParam interface {
}
type param struct {
	db pgxpool.Pool
}

func NewParam(db pgxpool.Pool) IParam {
	return &param{db: db}
}

func (g *gps) CreateStr(ctx context.Context, data models.ParamStr) error {

	if _, err := g.db.Exec(ctx, insertParamStr, data.StringPointID, data.PointID, data.Channel, data.Value, data.Time); err != nil {
		return err
	}

	return nil
}

func (g *gps) GetStrByPointID(ctx context.Context, pointID int) (*models.ParamStr, error) {

	var data models.ParamStr

	if err := g.db.QueryRow(ctx, getGpsByPointID, pointID).Scan(
		&data.ID,
		&data.StringPointID,
		&data.PointID,
		&data.Channel,
		&data.Value,
		&data.Time,
	); err != nil {
		return nil, err
	}

	return &data, nil
}

func (g *gps) CreateInt(ctx context.Context, data models.ParamInt) error {

	if _, err := g.db.Exec(ctx, insertParamInt, data.IntPointID, data.PointID, data.Channel, data.Value, data.Time); err != nil {
		return err
	}

	return nil
}

func (g *gps) GetIntByPointID(ctx context.Context, pointID int) (*models.ParamInt, error) {

	var data models.ParamInt

	if err := g.db.QueryRow(ctx, getParamIntByPointID, pointID).Scan(
		&data.ID,
		&data.IntPointID,
		&data.PointID,
		&data.Channel,
		&data.Value,
		&data.Time,
	); err != nil {
		return nil, err
	}

	return &data, nil
}
