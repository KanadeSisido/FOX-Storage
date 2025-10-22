package repository

import (
	"context"
	"fox-storage/model"
	"io"
	"os"
	"path/filepath"

	"gorm.io/gorm"
)

func (r itemRepository) LocateFile(ctx context.Context, id string, fileName string, file io.Reader) error {

	data, err := io.ReadAll(file)

	if err != nil {
		return err
	}

	fs := model.FileStorage{
		ItemID:     id,
		StorageKey: fileName,
		Checksum:   nil,
	}
	path := filepath.Join("./storage", fileName)

	err = gorm.G[model.FileStorage](r.db).Create(ctx, &fs)
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}
