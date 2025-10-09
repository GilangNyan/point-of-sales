package usecase

import (
	"context"
	"errors"
	"fmt"
	"gilangnyan/point-of-sales/internal/user/model"
	"gilangnyan/point-of-sales/internal/user/repository"
	"gilangnyan/point-of-sales/package/request"
	"gilangnyan/point-of-sales/package/response"
	"gilangnyan/point-of-sales/package/utils"
)

type UserUsecaseImpl struct {
	repo  repository.UserRepository
	repo2 repository.UserProfileRepository
}

func (u *UserUsecaseImpl) FindAll(ctx context.Context, params *request.PaginationParams) (*response.PaginationResponse[*model.UserWithProfile], error) {
	data, total, err := u.repo.FindAll(ctx, params)
	if err != nil {
		return nil, errors.New("failed to retrieve users")
	}
	return response.NewPaginationResponse(data, total, params.Page, params.PageSize), nil
}

func (u *UserUsecaseImpl) Create(ctx context.Context, data model.CreateUserDto) (string, error) {
	checkEmail, _ := u.repo.FindByEmail(ctx, data.Email)
	if checkEmail != nil {
		return "", errors.New("email already exists")
	}
	checkUsername, _ := u.repo.FindByUsername(ctx, data.Username)
	if checkUsername != nil {
		return "", errors.New("username already exists")
	}

	hashedPw, err := utils.HashPassword(data.Password)
	if err != nil {
		return "", errors.New("failed to hash password")
	}

	user := &model.User{
		Username: data.Username,
		Email:    data.Email,
		Password: hashedPw,
	}

	createdUser, err := u.repo.Create(ctx, *user)
	if err != nil {
		return "", errors.New("failed to create user")
	}

	profile := &model.UserProfile{
		FullName:    data.FullName,
		UserID:      createdUser,
		PhoneNumber: nil,
		DateOfBirth: nil,
		Address:     nil,
	}
	if data.PhoneNumber != nil {
		profile.PhoneNumber = data.PhoneNumber
	}
	if data.DateOfBirth != nil {
		profile.DateOfBirth = data.DateOfBirth
	}
	if data.Address != nil {
		profile.Address = data.Address
	}

	_, err = u.repo2.Create(ctx, *profile)
	if err != nil {
		fmt.Printf("Error creating user profile: %v\n", err)
		return "", errors.New("failed to create user profile")
	}

	return createdUser, nil
}

func (u *UserUsecaseImpl) Update(ctx context.Context, id string, data model.UpdateUserDto) (string, error) {
	checkUser, err := u.repo.FindByID(ctx, id)
	if err != nil {
		return "", errors.New("user not found")
	}

	if data.Username != nil && *data.Username != "" {
		checkUser.Username = *data.Username
	}
	if data.Email != nil && *data.Email != "" {
		checkUser.Email = *data.Email
	}
	if data.IsActive != nil && *data.IsActive != checkUser.IsActive {
		checkUser.IsActive = *data.IsActive
	}
	if data.IsBlocked != nil && *data.IsBlocked != checkUser.IsBlocked {
		checkUser.IsBlocked = *data.IsBlocked
	}

	user := &model.User{
		Username:  checkUser.Username,
		Email:     checkUser.Email,
		IsActive:  checkUser.IsActive,
		IsBlocked: checkUser.IsBlocked,
	}

	updatedUser, err := u.repo.Update(ctx, id, *user)
	if err != nil {
		return "", errors.New("failed to update user")
	}

	profile := &model.UserProfile{
		FullName:       checkUser.FullName,
		DateOfBirth:    checkUser.DateOfBirth,
		PhoneNumber:    checkUser.PhoneNumber,
		Address:        checkUser.Address,
		ProfilePicture: checkUser.ProfilePicture,
		UserID:         checkUser.ID,
	}

	_, err = u.repo2.Update(ctx, id, *profile)
	if err != nil {
		return "", errors.New("failed to update user profile")
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

func NewUserUsecase(repo repository.UserRepository, repo2 repository.UserProfileRepository) UserUsecase {
	return &UserUsecaseImpl{
		repo:  repo,
		repo2: repo2,
	}
}
