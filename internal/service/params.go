package service

import (
	"context"

	"github.com/Nazhgam/paramTracker/internal/models"
)

func (s *service) CreateParamStr(ctx context.Context, alias string, data models.ParamStr) error {
	point, err := s.storage.IPoint.GetByAlias(ctx, alias)
	if err != nil {
		s.log.Printf("error get point by alias in param str: %s", err)
		return err
	}

	data.PointID = point.PointID

	err = s.storage.IParam.CreateStr(ctx, data)
	if err != nil {
		s.log.Printf("error create param str: %s", err)
		return err
	}

	return nil
}

func (s *service) CreateParamInt(ctx context.Context, alias string, data models.ParamInt) error {
	point, err := s.storage.IPoint.GetByAlias(ctx, alias)
	if err != nil {
		s.log.Printf("error get point by alias in param str: %s", err)
		return err
	}

	data.PointID = point.PointID

	err = s.storage.IParam.CreateInt(ctx, data)
	if err != nil {
		s.log.Printf("error create param str: %s", err)
		return err
	}

	return nil
}
