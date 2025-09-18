package service

import (
	"context"
	"errors"
	"fox-storage/model"
)

func (s *Service) GetItems(ctx context.Context, folderId string, userId string) ([]model.Item, error) {

	permission, err := s.repository.GetPermissionByUserID(&ctx, folderId, userId)

	if err != nil {
		return nil, err
	}

	perm := permission.Permission
	
	if !(perm == "read" || perm == "write" || perm == "owner") {
		return nil, errors.New("Unauthorized")
	}
	
	return s.repository.GetItems(ctx, folderId)
}