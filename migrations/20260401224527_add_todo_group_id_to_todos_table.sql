-- +goose Up
-- +goose StatementBegin
ALTER TABLE todos
ADD COLUMN todo_group_id INTEGER REFERENCES todo_groups(id) ON DELETE SET NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE todos
DROP COLUMN todo_group_id;
-- +goose StatementEnd
