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
