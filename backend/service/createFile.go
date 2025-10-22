package service

import (
	"context"
	"errors"
	"io"
	"path/filepath"
)

func (s itemService) CreateFile(ctx context.Context, name string, file io.Reader, bytesSize int64, mime string, parentId *string, userId string) error {

	permission, err := s.repository.GetPermissionByUserID(ctx, *parentId, userId)

	if err != nil {
		return err
	}

	if permission.Permission != "owner" {
		return errors.New("access denied")
	}

	_name := filepath.Base(name) //example.txt
	f, err := s.repository.CreateItem(ctx, _name, "file", parentId, userId, &mime, &bytesSize)

	if err != nil {
		return err
	}

	locatedName := f.ID + filepath.Ext(name)
	err = s.repository.LocateFile(ctx, f.ID, locatedName, file)

	return err
}
