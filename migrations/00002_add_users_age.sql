-- +goose Up
ALTER TABLE users ADD COLUMN age INTEGER;

-- +goose Down
-- SQLite does not support DROP COLUMN in older versions, so we recreate the table
CREATE TABLE users_new (
    id         INTEGER PRIMARY KEY AUTOINCREMENT,
    name       TEXT    NOT NULL,
    email      TEXT    NOT NULL UNIQUE,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);
INSERT INTO users_new SELECT id, name, email, created_at FROM users;
DROP TABLE users;
ALTER TABLE users_new RENAME TO users;
