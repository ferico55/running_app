CREATE TABLE IF NOT EXISTS runs(
    id bigint unsigned NOT NULL AUTO_INCREMENT,
    user_id bigint unsigned NOT NULL,
    location_name varchar(300),
    date TIMESTAMP,
    duration INT,
    distance INT,
    total_steps INT,
    route TEXT,

    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users(id)
)