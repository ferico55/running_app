package model

type User struct {
	Id        int64  `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	Username  string `json:"username" db:"username"`
	AvatarUrl string `json:"AvatarUrl" db:"avatar_url"`
	Password  string `json:"-" db:"password"`
}

type RegisterRequest struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Name      string `json:"name"`
	AvatarUrl string `json:"avatarUrl"`
}
