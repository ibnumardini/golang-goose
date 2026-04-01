-- +goose Up
INSERT INTO todos (user_id, title, completed) VALUES
    (1, 'Buy groceries',       0),
    (1, 'Read a book',         1),
    (2, 'Go for a morning run',0),
    (2, 'Write unit tests',    0);

-- +goose Down
DELETE FROM todos WHERE title IN ('Buy groceries', 'Read a book', 'Go for a morning run', 'Write unit tests');
