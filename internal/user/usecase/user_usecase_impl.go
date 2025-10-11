package usecase

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"gilangnyan/point-of-sales/internal/user/model"
	"gilangnyan/point-of-sales/internal/user/repository"
	userRoleRepo "gilangnyan/point-of-sales/internal/user_role/repository"
	"gilangnyan/point-of-sales/package/request"
	"gilangnyan/point-of-sales/package/response"
	"gilangnyan/point-of-sales/package/transaction"
	"gilangnyan/point-of-sales/package/utils"
)

type UserUsecaseImpl struct {
	repo         repository.UserRepository
	repo2        repository.UserProfileRepository
	userRoleRepo userRoleRepo.UserRoleRepository
	tx           transaction.TransactionManager
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

	var createdUserID string

	err = u.tx.WithTransaction(func(tx *sql.Tx) error {
		user := &model.User{
			Username: data.Username,
			Email:    data.Email,
			Password: hashedPw,
		}

		createdUser, err := u.repo.Create(ctx, tx, *user)
		if err != nil {
			return fmt.Errorf("failed to create user: %w", err)
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

		_, err = u.repo2.Create(ctx, tx, *profile)
		if err != nil {
			return fmt.Errorf("failed to create user profile: %w", err)
		}

		if len(data.RoleIDs) > 0 {
			err = u.userRoleRepo.AssignRolesWithTx(ctx, tx, createdUser, data.RoleIDs)
			if err != nil {
				return fmt.Errorf("failed to assign roles: %w", err)
			}
		}

		createdUserID = createdUser
		return nil
	})

	if err != nil {
		return "", fmt.Errorf("transaction failed: %w", err)
	}

	return createdUserID, nil
}

func (u *UserUsecaseImpl) Update(ctx context.Context, id string, data model.UpdateUserDto) (string, error) {
	var updatedUserID string

	err := u.tx.WithTransaction(func(tx *sql.Tx) error {
		checkUser, err := u.repo.FindByID(ctx, id)
		if err != nil {
			return errors.New("user not found")
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

		updatedUser, err := u.repo.Update(ctx, tx, id, *user)
		if err != nil {
			return fmt.Errorf("failed to update user: %w", err)
		}

		profile := &model.UserProfile{
			FullName:       checkUser.FullName,
			DateOfBirth:    checkUser.DateOfBirth,
			PhoneNumber:    checkUser.PhoneNumber,
			Address:        checkUser.Address,
			ProfilePicture: checkUser.ProfilePicture,
			UserID:         checkUser.ID,
		}

		if data.FullName != nil && *data.FullName != "" {
			profile.FullName = *data.FullName
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

		_, err = u.repo2.Update(ctx, tx, id, *profile)
		if err != nil {
			return fmt.Errorf("failed to update user profile: %w", err)
		}

		if data.RoleIDs != nil {
			if len(*data.RoleIDs) > 0 {
				err = u.userRoleRepo.RemoveAllUserRolesWithTx(ctx, tx, updatedUser)
				if err != nil {
					return fmt.Errorf("failed to remove existing roles: %w", err)
				}
				err = u.userRoleRepo.AssignRolesWithTx(ctx, tx, updatedUser, *data.RoleIDs)
				if err != nil {
					return fmt.Errorf("failed to assign roles: %w", err)
				}
			}
		}

		updatedUserID = updatedUser
		return nil
	})

	if err != nil {
		return "", fmt.Errorf("transaction failed: %w", err)
	}

	return updatedUserID, nil
}

func (u *UserUsecaseImpl) Delete(ctx context.Context, id string) error {
	return u.tx.WithTransaction(func(tx *sql.Tx) error {
		checkUser, err := u.repo.FindByID(ctx, id)
		if err != nil {
			return errors.New("user not found")
		}

		err = u.repo2.Delete(ctx, tx, checkUser.ID)
		if err != nil {
			return fmt.Errorf("failed to delete user profile: %w", err)
		}

		err = u.repo.Delete(ctx, tx, checkUser.ID)
		if err != nil {
			return fmt.Errorf("failed to delete user: %w", err)
		}

		return nil
	})
}

func NewUserUsecase(repo repository.UserRepository, repo2 repository.UserProfileRepository, userRoleRepo userRoleRepo.UserRoleRepository, tx transaction.TransactionManager) UserUsecase {
	return &UserUsecaseImpl{
		repo:         repo,
		repo2:        repo2,
		userRoleRepo: userRoleRepo,
		tx:           tx,
	}
}
