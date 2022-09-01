package repository

import (
	"context"

	"github.com/ferico55/running_app/model"
)

func GetFriendForUser(userId int64, context context.Context) ([]model.User, error) {
	db := openDBConnection()
	defer db.Close()

	rows, err := db.QueryxContext(context, `
	select u.id, u.username, u.name, u.avatar_url FROM user_friend f
JOIN users u ON f.friend_id = u.id
WHERE f.user_id = ? AND f.status = "accepted"
UNION
select u.id, u.username, u.name, u.avatar_url FROM user_friend f
JOIN users u ON f.user_id = u.id
WHERE f.friend_id = ? AND f.status = "accepted"
	`, userId, userId)
	if err != nil {
		return []model.User{}, err
	}

	defer rows.Close()
	var friends = []model.User{}
	for rows.Next() {
		var user model.User
		err = rows.StructScan(&user)
		if err != nil {
			return []model.User{}, err
		}
		friends = append(friends, user)
	}
	return friends, nil
}

func IsFriended(userId int64, friendId int64, context context.Context) error {
	db := openDBConnection()
	defer db.Close()

	row := db.QueryRowxContext(context, `SELECT user_id FROM user_friend
	WHERE (user_id = ? AND friend_id = ?) OR (friend_id = ? AND user_id = ?)`, userId, friendId, userId, friendId)
	var id int64
	return row.Scan(&id)
}

func AddFriend(userId int64, friendId int64, context context.Context) error {
	db := openDBConnection()
	defer db.Close()

	_, err := db.ExecContext(context, `INSERT INTO user_friend VALUES(?, ?, 'pending')`, userId, friendId)
	return err
}

func RemoveFriend(userId int64, friendId int64, context context.Context) error {
	db := openDBConnection()
	defer db.Close()

	_, err := db.ExecContext(context, `DELETE FROM user_friend WHERE (user_id = ? AND friend_id = ?) OR (friend_id = ? AND user_id = ?)`, userId, friendId, userId, friendId)
	return err
}

func GetPendingFriendRequest(userId int64, context context.Context) ([]model.User, error) {
	db := openDBConnection()
	defer db.Close()

	rows, err := db.QueryxContext(context, `
select u.id, u.username, u.name, u.avatar_url FROM user_friend f
JOIN users u ON f.user_id = u.id
WHERE f.friend_id = ? AND f.status = "pending"
	`, userId)
	if err != nil {
		return []model.User{}, err
	}

	defer rows.Close()
	var friends = []model.User{}
	for rows.Next() {
		var user model.User
		err = rows.StructScan(&user)
		if err != nil {
			return []model.User{}, err
		}
		friends = append(friends, user)
	}
	return friends, nil
}
