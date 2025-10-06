package usecase

import (
	"context"
	"errors"
	"gilangnyan/point-of-sales/internal/user/model"
	"gilangnyan/point-of-sales/internal/user/repository"
	"gilangnyan/point-of-sales/package/utils"
)

type UserUsecaseImpl struct {
	repo repository.UserRepository
}

func (u *UserUsecaseImpl) FindAll(ctx context.Context) ([]*model.User, error) {
	data, err := u.repo.FindAll(ctx)
	if err != nil {
		return nil, errors.New("failed to retrieve users")
	}
	return data, nil
}

func (u *UserUsecaseImpl) Create(ctx context.Context, data model.CreateUserDto) (*model.User, error) {
	checkEmail, _ := u.repo.FindByEmail(ctx, data.Email)
	if checkEmail != nil {
		return nil, errors.New("email already exists")
	}
	checkUsername, _ := u.repo.FindByUsername(ctx, data.Username)
	if checkUsername != nil {
		return nil, errors.New("username already exists")
	}

	hashedPw, err := utils.HashPassword(data.Password)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	user := &model.User{
		Username: data.Username,
		Email:    data.Email,
		Password: hashedPw,
	}

	createdUser, err := u.repo.Create(ctx, *user)
	if err != nil {
		return nil, errors.New("failed to create user")
	}

	return createdUser, nil
}

func (u *UserUsecaseImpl) Update(ctx context.Context, id string, data model.UpdateUserDto) (*model.User, error) {
	checkUser, err := u.repo.FindByID(ctx, id)
	if err != nil {
		return nil, errors.New("user not found")
	}

	if data.Username != nil && *data.Username != "" {
		checkUser.Username = *data.Username
	}
	if data.Email != nil && *data.Email != "" {
		checkUser.Email = *data.Email
	}

	updatedUser, err := u.repo.Update(ctx, id, *checkUser)
	if err != nil {
		return nil, errors.New("failed to update user")
	}

	return updatedUser, nil
}

func (u *UserUsecaseImpl) Delete(ctx context.Context, id string) error {
	checkUser, err := u.repo.FindByID(ctx, id)
	if err != nil {
		return errors.New("user not found")
	}

	err = u.repo.Delete(ctx, checkUser.ID)
	if err != nil {
		return errors.New("failed to delete user")
	}

	return nil
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &UserUsecaseImpl{
		repo: repo,
	}
}
