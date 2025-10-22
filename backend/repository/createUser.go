package repository

import (
	"context"
	"fox-storage/model"

	"gorm.io/gorm"
)

func (r userRepository) CreateUser(ctx context.Context, username string, email string, hashedPassword string) error {

	err := r.db.Transaction(func(tx *gorm.DB) error {

		user := model.User{Name: username, Email: email, HashedPassword: hashedPassword}

		err := gorm.G[model.User](tx).Create(ctx, &user)
		if err != nil {
			return err
		}

		root := model.Item{
			Name:    "root",
			Type:    "folder",
			OwnerID: user.ID,
		}

		err = gorm.G[model.Item](tx).Create(ctx, &root)

		if err != nil {
			return err
		}

		_, err = gorm.G[model.User](tx).Where("id=?", user.ID).Update(ctx, "root_id", root.ID)

		if err != nil {
			return err
		}

		perm := model.Permission{
			ItemID:     root.ID,
			UserID:     user.ID,
			Permission: "Owner",
		}

		err = gorm.G[model.Permission](tx).Create(ctx, &perm)
		if err != nil {
			return err
		}

		return nil
	})

	return err
}
