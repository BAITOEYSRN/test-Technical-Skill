package repository

import (
	"context"
	"errors"
	"net/http"

	"github.com/BAITOEYSRN/test-Technical-Skill/internal/infrastructure/models"
	"github.com/BAITOEYSRN/test-Technical-Skill/pkg/response"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(ctx context.Context, user models.User) (*models.User, error) {
	if err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&user).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) GetListUsers(ctx context.Context) ([]models.User, error) {
	var users []models.User
	err := r.db.WithContext(ctx).Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepository) GetUserByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.Wrap(
				err,
				http.StatusNotFound,
				uuid.New(),
				"user not found",
			)
		}

		return nil, response.Wrap(
			err,
			http.StatusInternalServerError,
			uuid.New(),
			"failed to get user by id",
		)
	}
	return &user, nil
}
