package mysql

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"

	"github.com/Yamako76/go-react-sample/domain"
)

type UserModel struct {
	ID        int64     `gorm:"column:id"`
	Name      string    `gorm:"column:name"`
	Email     string    `gorm:"column:email"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

var _ domain.UserRepository = (*UserRepository)(nil)

func (repo *UserRepository) Get(ctx context.Context, id int64) (user *domain.User, err error) {
	var userModel UserModel
	result := repo.db.Raw("SELECT id, name, email, created_at FROM users WHERE id = ?", id).Scan(&userModel)
	if result.RowsAffected == 0 {
		return nil, domain.ErrUserNotFound
	}
	if result.Error != nil {
		return nil, fmt.Errorf("ユーザーの取得に失敗しました: %w", result.Error)
	}

	user = &domain.User{
		ID:        userModel.ID,
		Name:      userModel.Name,
		Email:     userModel.Email,
		CreatedAt: userModel.CreatedAt,
	}
	return user, nil
}
