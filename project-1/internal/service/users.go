package service

import (
	"context"
	"errors"
	"github.com/Ilja-R/TeachMeSkillsHW/project-1/internal/errs"
	"github.com/Ilja-R/TeachMeSkillsHW/project-1/internal/models"
)

func (s *UsersService) GetAllUsers(ctx context.Context) (users []models.User, err error) {
	users, err = s.repository.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *UsersService) GetUserByID(ctx context.Context, id int) (user models.User, err error) {
	user, err = s.repository.GetUserByID(ctx, id)
	if err != nil {
		if errors.Is(err, errs.ErrNotfound) {
			return models.User{}, errs.ErrUserNotfound
		}
		return models.User{}, err
	}

	return user, nil
}

func (s *UsersService) CreateUser(ctx context.Context, user models.User) (err error) {
	err = s.repository.CreateUser(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (s *UsersService) UpdateUserByID(ctx context.Context, user models.User) (err error) {
	_, err = s.repository.GetUserByID(ctx, user.ID)
	if err != nil {
		if errors.Is(err, errs.ErrNotfound) {
			return errs.ErrUserNotfound
		}
		return err
	}

	err = s.repository.UpdateUserByID(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (s *UsersService) DeleteUserByID(ctx context.Context, id int) (err error) {
	_, err = s.repository.GetUserByID(ctx, id)
	if err != nil {
		if errors.Is(err, errs.ErrNotfound) {
			return errs.ErrUserNotfound
		}
		return err
	}

	err = s.repository.DeleteUserByID(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
