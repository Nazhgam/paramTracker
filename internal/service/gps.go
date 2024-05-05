package service

import (
	"context"

	"github.com/Nazhgam/paramTracker/internal/models"
)

func (s *service) CreateGps(ctx context.Context, alias string, data models.GPS) error {
	point, err := s.storage.IPoint.GetByAlias(ctx, alias)
	if err != nil {
		s.log.Printf("error get point by alias: %s", err)
		return err
	}

	data.PointID = point.PointID

	if err := s.storage.IGps.Create(ctx, data); err != nil {
		s.log.Printf("error create gps in table: %s", err)
		return err
	}

	return nil
}

func (s *service) GetAllGps(ctx context.Context) ([]models.GPS, error) {
	res, err := s.storage.IGps.GetAll(ctx)
	if err != nil {
		s.log.Printf("error get all gps: %s", err)
		return nil, err
	}

	return res, nil
}
