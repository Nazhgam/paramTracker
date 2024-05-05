package service

import (
	"context"
	"errors"
	"strings"

	"github.com/Nazhgam/paramTracker/internal/models"
)

func (s *service) CreatePoint(ctx context.Context, point models.Point) error {
	if strings.TrimSpace(point.Name) == "" {
		s.log.Println("empty name")
		return errors.New("empty name")
	}

	if strings.TrimSpace(point.Alias) == "" {
		s.log.Println("empty alias")
		return errors.New("empty alias")
	}

	err := s.storage.IPoint.Create(ctx, point)
	if err != nil {
		s.log.Printf("error create point: %s", err)
		return err
	}

	return nil
}

func (s *service) GetLastIDByAlias(ctx context.Context, alias string) (*models.LastIDResult, error) {
	point, err := s.storage.IPoint.GetByAlias(ctx, alias)
	if err != nil {
		s.log.Printf("error get point by alias: %s", err)
		return nil, err
	}

	res, err := s.storage.IPoint.GetLastID(ctx, point.PointID)
	if err != nil {
		s.log.Printf("error get gps by pointID: %s", err)
		return nil, err
	}

	return res, nil
}
