package repository

import (
	"context"
	"github.com/wangzupeng12061/we-book/internal/domain"
	"github.com/wangzupeng12061/we-book/internal/repository/cache"
	"github.com/wangzupeng12061/we-book/internal/repository/dao"
)

var ErrUserDuplicatedEmail = dao.ErrDuplicateEmail
var ErrUserNotFound = dao.ErrUserNotFound

type UserRepository struct {
	dao *dao.UserDAO
	uc  *cache.UserCache
}

func NewUserRepository(dao *dao.UserDAO, uc *cache.UserCache) *UserRepository {
	return &UserRepository{
		dao: dao,
		uc:  uc,
	}
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
func (repo *UserRepository) FindById(ctx context.Context, id int64) (domain.User, error) {
	u, err := repo.uc.Get(ctx, id)
	if err == nil {
		return u, nil
	}
	//if errors.Is(err, cache.ErrKeyNotFound) {
	//
	//}
	ud, err := repo.dao.FindById(ctx, id)

	if err != nil {
		return domain.User{}, err
	}
	u = domain.User{
		ID:       ud.Id,
		Email:    ud.Email,
		Password: ud.Password,
	}
	go func() {
		err = repo.uc.Set(ctx, u)
		if err != nil {
			//return domain.User{}, err
		}
	}()

	return u, nil

}
