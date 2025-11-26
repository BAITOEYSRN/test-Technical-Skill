package domain

import (
	"context"

	"github.com/BAITOEYSRN/test-Technical-Skill/internal/domain/entity"
	"github.com/google/uuid"
)

type UserUsecase interface {
	GetListUsers(ctx context.Context) ([]entity.User, error)
	CreateUser(ctx context.Context, user entity.User) (*uuid.UUID, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (*entity.User, error)
}
