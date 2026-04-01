# golang-goose

A demo project for database migrations using [pressly/goose](https://github.com/pressly/goose) with SQLite.

## Requirements

- Go 1.25+

## Setup

```bash
git clone https://github.com/ibnumardini/golang-goose
cd golang-goose
go mod download
```

## Usage

```bash
go run main.go up       # apply all pending migrations
go run main.go down     # rollback the last migration
go run main.go reset    # rollback all migrations
go run main.go status   # show migration status
go run main.go version  # show current version
```

## Migrations

| File | Description |
| ---- | ----------- |
| `00001_create_users.sql` | Create `users` table |
| `00002_add_users_age.sql` | Add `age` column to `users` |
| `00003_seed_users.sql` | Seed initial users data |
| `20260401144235_todo_groups.sql` | Create `todo_groups` table |
| `20260401213605_create_todos.sql` | Create `todos` table (relation to `users`) |
| `20260401213905_seed_todos.sql` | Seed initial todos data |
| `20260401224527_add_todo_group_id_to_todos_table.sql` | Add `todo_group_id` column to `todos` |

## Schema

```
users
├── id
├── name
├── email
├── age
└── created_at

todo_groups
├── id
├── name
├── created_at
└── updated_at

todos
├── id
├── user_id       → users.id
├── todo_group_id → todo_groups.id
├── title
├── completed
└── created_at
```
