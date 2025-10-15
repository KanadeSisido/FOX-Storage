package service

import (
	"context"
	"errors"
	"fmt"
)

func (s *Service) GetFileName(ctx context.Context, fileId string, userId string) (*string, *string, error) {

	perm, err := s.repository.GetPermissionByUserID(&ctx, fileId, userId)

	if err != nil {
		return nil, nil, errors.New("access denied")
	}

	if !(perm.Permission == "owner" || perm.Permission == "write" || perm.Permission == "read") {
		return nil, nil, errors.New("access denied")
	}

	item, err := s.repository.GetInsItemById(ctx, fileId)

	fmt.Println(item)

	filepath := "./storage/" + item.StorageKey

	return &filepath, &item.Name, err
}