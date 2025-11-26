package domain

import (
	"context"

	"github.com/BAITOEYSRN/test-Technical-Skill/internal/infrastructure/models"
	"github.com/google/uuid"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user models.User) (*models.User, error)
	GetListUsers(ctx context.Context) ([]models.User, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (*models.User, error)
}
