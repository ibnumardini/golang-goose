-- +goose Up
INSERT INTO users (name, email, age) VALUES
    ('Alice', 'alice@example.com', 30),
    ('Bob',   'bob@example.com',   25);

-- +goose Down
DELETE FROM users WHERE email IN ('alice@example.com', 'bob@example.com');
