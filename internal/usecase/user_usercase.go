package usecase

import (
	"context"

	"github.com/BAITOEYSRN/test-Technical-Skill/internal/domain/entity"
	domain "github.com/BAITOEYSRN/test-Technical-Skill/internal/domain/reppository"
	"github.com/BAITOEYSRN/test-Technical-Skill/internal/infrastructure/models"
	"github.com/google/uuid"
)

type userRepo struct {
	userRepository domain.UserRepository
}

func NewUserRepository(userRepository domain.UserRepository) *userRepo {
	return &userRepo{
		userRepository: userRepository,
	}
}

func (r *userRepo) GetListUsers(ctx context.Context) ([]entity.User, error) {
	users, err := r.userRepository.GetListUsers(ctx)
	if err != nil {
		return nil, err
	}
	var usersEntity []entity.User
	for _, user := range users {
		usersEntity = append(usersEntity, entity.User{
			ID:          user.ID,
			FirstName:   user.FirstName,
			LastName:    user.LastName,
			DateOfBirth: user.DateOfBirth,
			Age:         user.Age,
			Address:     user.Address,
			CreatedAt:   user.CreatedAt,
			UpdatedAt:   user.UpdatedAt,
		})
	}
	return usersEntity, nil
}

func (r *userRepo) CreateUser(ctx context.Context, user entity.User) (*entity.User, error) {

	userCreated, err := r.userRepository.CreateUser(ctx, models.User{
		ID:          user.ID,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		DateOfBirth: user.DateOfBirth,
		Age:         user.Age,
		Address:     user.Address,
	})
	if err != nil {
		return nil, err
	}
	return &entity.User{
		ID:          userCreated.ID,
		FirstName:   userCreated.FirstName,
		LastName:    userCreated.LastName,
		DateOfBirth: userCreated.DateOfBirth,
		Age:         userCreated.Age,
		Address:     userCreated.Address,
	}, nil
}

func (u *userRepo) GetUserByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	user, err := u.userRepository.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &entity.User{
		ID:          user.ID,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		DateOfBirth: user.DateOfBirth,
		Age:         user.Age,
		Address:     user.Address,
	}, nil
}
