CREATE TABLE IF NOT EXISTS users(
    id bigint unsigned NOT NULL AUTO_INCREMENT,
    username varchar(50),
    name varchar(200),
    avatar_url varchar(300),
    password varchar(200),
    PRIMARY KEY (id)
)