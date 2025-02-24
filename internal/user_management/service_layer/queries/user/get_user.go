package queries

import (
	"clean-hex/internal/user_management/domain/entities"
	"clean-hex/pkg/errors"
	"gorm.io/gorm"
)

func ViewGetUser(db *gorm.DB, id uint) (*entities.User, error) {
	user := new(entities.User)
	if err := db.First(user, id).Error; err != nil {
		return nil, errors.NotFound("User.NotFound")
	}
	return user, nil
}
