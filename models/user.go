package models

import (
	"errors"
	"rest-api/db"
	"rest-api/utils"
)

// User represents user data
type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

// Save saves a new user in the database
func (u *User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?, ?)"

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}

	u.ID, err = result.LastInsertId()

	return err
}

// ValidateCredentials validates a user's credentials
func (u *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"

	var retrievedPassword string

	row := db.DB.QueryRow(query, u.Email)
	err := row.Scan(&u.ID, &retrievedPassword)
	if err != nil {
		return err
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)

	if !passwordIsValid {
		return errors.New("credentials invalid")
	}

	return nil
}