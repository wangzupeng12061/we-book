package repository

import (
	"context"
	"github.com/wangzupeng12061/we-book/internal/domain"
	"github.com/wangzupeng12061/we-book/internal/repository/dao"
)

var ErrUserDuplicatedEmail = dao.ErrDuplicateEmail

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
