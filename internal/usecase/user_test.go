package usecase_test

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/BAITOEYSRN/test-Technical-Skill/internal/domain/entity"
	"github.com/BAITOEYSRN/test-Technical-Skill/internal/infrastructure/models"
	"github.com/BAITOEYSRN/test-Technical-Skill/internal/infrastructure/repository"
	"github.com/BAITOEYSRN/test-Technical-Skill/internal/usecase"
	"github.com/BAITOEYSRN/test-Technical-Skill/pkg/response"
	"github.com/BAITOEYSRN/test-Technical-Skill/pkg/utils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var (
	timeMock = utils.TimeMock()
	uuidMock = uuid.New()
	userID   = uuid.MustParse("123e4567-e89b-12d3-a456-426614174000")
)

func TestGetListUsers(t *testing.T) {
	userID := uuid.New()

	type given struct {
		ctx context.Context
	}

	type when struct {
		repo []models.User
		err  error
	}

	type expect struct {
		users []entity.User
		err   error
	}

	type testCase struct {
		name   string
		given  given
		when   when
		expect expect
	}

	tests := []testCase{
		{
			name: "success",
			given: given{
				ctx: context.Background(),
			},
			when: when{
				repo: []models.User{
					{
						ID:          userID,
						FirstName:   "John",
						LastName:    "Doe",
						DateOfBirth: "2000-01-01",
						Age:         25,
						Address:     "Bangkok",
						CreatedAt:   timeMock,
						UpdatedAt:   &timeMock,
					},
				},
			},
			expect: expect{
				users: []entity.User{
					{
						ID:          userID,
						FirstName:   "John",
						LastName:    "Doe",
						DateOfBirth: "2000-01-01",
						Age:         25,
						Address:     "Bangkok",
						CreatedAt:   timeMock,
						UpdatedAt:   &timeMock,
					},
				},
				err: nil,
			},
		},
		{
			name: "error",
			given: given{
				ctx: context.Background(),
			},
			when: when{
				repo: nil,
				err:  response.Wrap(errors.New("error"), http.StatusNotFound, uuidMock, "user not found"),
			},
			expect: expect{
				users: nil,
				err:   response.Wrap(errors.New("error"), http.StatusNotFound, uuidMock, "user not found"),
			},
		},
		{
			name: "get list user []models.User{}",
			given: given{
				ctx: context.Background(),
			},
			when: when{
				repo: []models.User{},
				err:  nil,
			},
			expect: expect{
				users: []entity.User{},
				err:   nil,
			},
		},
	}

	for _, test := range tests {
		mockRepo := repository.NewUserRepositoryMock()
		mockRepo.On("GetListUsers", test.given.ctx).Return(test.when.repo, test.when.err)
		usecase := usecase.NewUserRepository(mockRepo)
		users, err := usecase.GetListUsers(test.given.ctx)
		if test.expect.err != nil {
			assert.Error(t, err)
			assert.Equal(t, test.expect.err, err)
			continue
		}
		assert.NoError(t, err)
		assert.Equal(t, test.expect.users, users)

	}
}

func TestCreateUser(t *testing.T) {
	type given struct {
		ctx  context.Context
		user entity.User
	}

	type when struct {
		repo *models.User
		err  error
	}

	type expect struct {
		user *uuid.UUID
		err  error
	}

	type testCase struct {
		name   string
		given  given
		when   when
		expect expect
	}

	tests := []testCase{
		{
			name: "success",
			given: given{
				ctx: context.Background(),
				user: entity.User{
					FirstName:   "John",
					LastName:    "Doe",
					DateOfBirth: "2000-01-01",
					Age:         25,
					Address:     "Bangkok",
				},
			},
			when: when{
				repo: &models.User{
					ID:          userID,
					FirstName:   "John",
					LastName:    "Doe",
					DateOfBirth: "2000-01-01",
					Age:         25,
					Address:     "Bangkok",
				},
				err: nil,
			},
			expect: expect{
				user: &userID,
				err:  nil,
			},
		},

		{
			name: "error while creating user",
			given: given{
				ctx: context.Background(),
				user: entity.User{
					FirstName:   "John",
					LastName:    "Doe",
					DateOfBirth: "2000-01-01",
					Age:         25,
					Address:     "Bangkok",
				},
			},
			when: when{
				repo: &models.User{},
				err: response.Wrap(
					errors.New("error"),
					http.StatusInternalServerError,
					uuidMock,
					"failed to create user",
				),
			},
			expect: expect{
				user: nil,
				err: response.Wrap(
					errors.New("error"),
					http.StatusInternalServerError,
					uuidMock,
					"failed to create user",
				),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockRepo := repository.NewUserRepositoryMock()

			mockRepo.On("CreateUser", test.given.ctx, models.User{
				FirstName:   test.given.user.FirstName,
				LastName:    test.given.user.LastName,
				DateOfBirth: test.given.user.DateOfBirth,
				Age:         test.given.user.Age,
				Address:     test.given.user.Address,
			}).
				Return(test.when.repo, test.when.err)

			usecase := usecase.NewUserRepository(mockRepo)

			user, err := usecase.CreateUser(test.given.ctx, test.given.user)

			if test.expect.err != nil {
				assert.Error(t, err)
				assert.EqualError(t, err, test.expect.err.Error())
				assert.Nil(t, user)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, test.expect.user, user)
		})
	}
}

func TestGetUserByID(t *testing.T) {
	type given struct {
		ctx context.Context
		id  uuid.UUID
	}

	type when struct {
		repo *models.User
		err  error
	}

	type expect struct {
		user *entity.User
		err  error
	}

	type testCase struct {
		name   string
		given  given
		when   when
		expect expect
	}

	tests := []testCase{
		{
			name: "success",
			given: given{
				ctx: context.Background(),
				id:  userID,
			},
			when: when{
				repo: &models.User{
					ID:          userID,
					FirstName:   "John",
					LastName:    "Doe",
					DateOfBirth: "2000-01-01",
					Age:         25,
					Address:     "Bangkok",
					CreatedAt:   timeMock,
					UpdatedAt:   &timeMock,
				},
				err: nil,
			},
			expect: expect{
				user: &entity.User{
					ID:          userID,
					FirstName:   "John",
					LastName:    "Doe",
					DateOfBirth: "2000-01-01",
					Age:         25,
					Address:     "Bangkok",
					CreatedAt:   timeMock,
					UpdatedAt:   &timeMock,
				},
				err: nil,
			},
		},
		{
			name: "user not found",
			given: given{
				ctx: context.Background(),
				id:  uuid.New(),
			},
			when: when{
				repo: nil,
				err: response.Wrap(
					errors.New("not found user by id"),
					http.StatusNotFound,
					uuidMock,
					"not found user by id",
				),
			},
			expect: expect{
				user: nil,
				err: response.Wrap(
					errors.New("not found user by id"),
					http.StatusNotFound,
					uuidMock,
					"not found user by id",
				),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockRepo := repository.NewUserRepositoryMock()

			mockRepo.On("GetUserByID", test.given.ctx, test.given.id).
				Return(test.when.repo, test.when.err)

			usecase := usecase.NewUserRepository(mockRepo)
			user, err := usecase.GetUserByID(test.given.ctx, test.given.id)
			if test.expect.err != nil {
				assert.Error(t, err)
				assert.Equal(t, test.expect.err, err)
				assert.Nil(t, user)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, test.expect.user, user)
		})
	}
}
