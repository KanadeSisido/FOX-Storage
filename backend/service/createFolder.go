package service

import (
	"context"
	"errors"
)

func (s *Service) CreateFolder(ctx *context.Context, name string, parentId *string, userId string) error {

	permission, err := s.repository.GetPermissionByUserID(ctx, *parentId, userId);

	if err != nil {
		return err
	}

	if permission.Permission != "owner" {
		return errors.New("access denied")
	}

	_, err =  s.repository.CreateItem(ctx, name, "folder", parentId, userId, nil, nil)

	return err
}