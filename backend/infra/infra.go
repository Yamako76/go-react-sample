package infra

import (
	"gorm.io/gorm"

	"github.com/Yamako76/go-react/infra/mysql"
)

func NewDB() (*gorm.DB, error) {
	return mysql.NewDB()
}
