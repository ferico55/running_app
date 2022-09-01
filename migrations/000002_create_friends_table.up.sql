CREATE TABLE IF NOT EXISTS user_friend(
    user_id bigint unsigned NOT NULL,
    friend_id bigint unsigned NOT NULL,
    status varchar(20),
    PRIMARY KEY (user_id, friend_id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (friend_id) REFERENCES users(id)
)