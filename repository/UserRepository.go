package repository

import (
	"context"

	"github.com/ferico55/running_app/model"
)

func GetUserByUsername(username string, context context.Context) (*model.User, error) {
	db := openDBConnection()
	defer db.Close()

	rows, err := db.QueryxContext(context, `SELECT id, name, username, avatar_url, password FROM users WHERE username=?`, username)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var user model.User
	if rows.Next() {
		err = rows.StructScan(&user)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, nil
	}
	return &user, nil
}

func IsUsernameFound(username string, context context.Context) error {
	db := openDBConnection()
	defer db.Close()

	rows := db.QueryRowxContext(context, `SELECT username FROM users WHERE username=?`, username)
	var emailFound string
	return rows.Scan(&emailFound)
}

func RegisterUser(user model.RegisterRequest, context context.Context) (int64, error) {
	db := openDBConnection()
	defer db.Close()

	result, err := db.ExecContext(context, `
	INSERT into users (username, name, password, avatar_url)
	VALUES (?, ?, ?, ?)`, user.Username, user.Name, user.Password, user.AvatarUrl)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func GetUserById(id int64, context context.Context) (*model.User, error) {
	db := openDBConnection() //db->dari si connection.go
	defer db.Close()

	rows, err := db.QueryxContext(context, `SELECT id, name, username, avatar_url, password FROM users where id=?`, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var user model.User
	if rows.Next() {
		err = rows.StructScan(&user)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, nil
	}
	return &user, nil
}
