package repository

import (
	"context"
	"github.com/wangzupeng12061/we-book/internal/domain"
	"github.com/wangzupeng12061/we-book/internal/repository/dao"
)

var ErrUserDuplicatedEmail = dao.ErrDuplicateEmail
var ErrUserNotFound = dao.ErrUserNotFound

type UserRepository struct {
	dao *dao.UserDAO
}

func NewUserRepository(dao *dao.UserDAO) *UserRepository {
	return &UserRepository{dao: dao}
}

func (repo *UserRepository) Create(ctx context.Context, u domain.User) error {
	return repo.dao.Insert(ctx, dao.User{
		Email:    u.Email,
		Password: u.Password,
	})

}

func (repo *UserRepository) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	u, err := repo.dao.FindByEmail(ctx, email)
	if err != nil {
		return domain.User{}, err
	}
	return domain.User{
		ID:       u.Id,
		Email:    u.Email,
		Password: u.Password,
	}, nil
}
