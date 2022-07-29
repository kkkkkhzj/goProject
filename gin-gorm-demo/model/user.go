package model

import "gorm.io/gorm"

type User struct {
	gorm.Model // embeded struct
	Username   string
	Password   string
	Email      string
}

// Create and insert user into table `users`
func (db *Database) InsertAUser(username, password, email string) error {
	return db.engine.Create(&User{Username: username, Password: password, Email: email}).Error
}

// Query all users
func (db *Database) QueryAllUsers() ([]User, error) {
	var users []User
	if err := db.engine.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
