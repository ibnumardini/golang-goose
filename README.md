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

## Goose Tracking Table

Goose automatically creates a `goose_db_version` table to track applied migrations:

``` bash
goose_db_version
├── id         INTEGER
├── version_id INTEGER NOT NULL
├── is_applied INTEGER NOT NULL
└── tstamp     TIMESTAMP
```

Example state after all migrations are applied:

| id | version_id | is_applied | tstamp |
| -- | ---------- | ---------- | ------ |
| 1 | 0 | 1 | 2026-04-01 14:34:34 |
| 2 | 1 | 1 | 2026-04-01 14:34:41 |
| 3 | 2 | 1 | 2026-04-01 14:34:41 |
| 4 | 3 | 1 | 2026-04-01 14:34:41 |
| 5 | 20260401213605 | 1 | 2026-04-01 14:39:18 |
| 6 | 20260401213905 | 1 | 2026-04-01 14:39:18 |
| 7 | 20260401144235 | 1 | 2026-04-01 14:51:58 |
| 8 | 20260401224527 | 1 | 2026-04-01 14:51:58 |

`version_id` corresponds to the numeric prefix of each migration file. `is_applied = 1` means the migration has been applied; on rollback, goose inserts a new row with `is_applied = 0`.

## Schema

``` bash
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
