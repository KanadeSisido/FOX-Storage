package service

import (
	"context"
	"errors"
)

func (s itemService) GetFileName(ctx context.Context, fileId string, userId string) (*string, *string, error) {

	perm, err := s.repository.GetPermissionByUserID(ctx, fileId, userId)

	if err != nil {
		return nil, nil, errors.New("access denied")
	}

	if !(perm.Permission == "owner" || perm.Permission == "write" || perm.Permission == "read") {
		return nil, nil, errors.New("access denied")
	}

	item, err := s.repository.GetItemLocationById(ctx, fileId)

	filepath := "./storage/" + item.StorageKey

	return &filepath, &item.Name, err
}
