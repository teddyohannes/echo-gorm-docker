package repository

import (
	"echo-gorm-docker/config"
	"echo-gorm-docker/entity"
	"fmt"
)

//UserRepository is contract what userRepository can do to db
type UserRepository interface {
	ProfileUser(userID string) (entity.User, error)
}

func ProfileUser(userID string) (entity.User, error) {
	fmt.Println("repository profile")
	var user entity.User
	db := config.DbManager()
	db.Preload("Books").Preload("Books.User").Find(&user, userID)
	return user, nil
}
