
-- +migrate Up
CREATE TABLE users (
    id                 CHAR(36) PRIMARY KEY,
    user_name          VARCHAR(255) NOT NULL,
    email              VARCHAR(255) UNIQUE,
    hashed_password    BLOB NOT NULL,
    salt               BLOB NOT NULL,
    created_at         DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at         DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);


-- +migrate Down
DROP TABLE `users`;