package service

import "context"

func (s userService) GetRootIdByUserId(ctx context.Context, userId string) (*string, error) {

	user, err := s.repository.GetUserById(ctx, userId)

	if err != nil {
		return nil, err
	}

	return &user.RootID, nil
}
