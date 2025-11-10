package models

import (
	"errors"
	"fmt"

	db "example.com/rest-api/database"
	"example.com/rest-api/helper"
	"github.com/google/uuid"
)

type User struct {
	Id       uuid.UUID
	Email    string `binding required`
	Password string `binding required`
}

func (e User) Save() error {

	e.Id = uuid.New()

	hashedPassword, err := helper.HashPassword(e.Password)

	if err != nil {
		return err
	}

	query := `INSERT INTO Users(id, email, password) VALUES(?,?,?)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	fmt.Println(e.Id.String())

	_, err = stmt.Exec(e.Id.String(), e.Email, hashedPassword)

	return err
}

func (e User) ValidateCredentials() error {
	query := `SELECT id, password FROM Users WHERE email=?`
	row := db.DB.QueryRow(query, e.Email)

	fmt.Println(row)
	var userPassword string
	err := row.Scan(&e.Id, &userPassword)
	if err != nil {
		return errors.New("please provide a password")
	}

	authenticated := helper.CompareHashedPassword(userPassword, e.Password)

	if !authenticated {
		return errors.New("incorrect password")
	}

	return nil
}
