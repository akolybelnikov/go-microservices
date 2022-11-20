-- name: GetAll :many
select id,
       email,
       first_name,
       last_name,
       password,
       user_active,
       created_at,
       updated_at
from users
order by last_name;

-- name: GetByEmail :one
select id,
       email,
       first_name,
       last_name,
       password,
       user_active,
       created_at,
       updated_at
from users
where email = $1;

-- name: GetOne :one
select id,
       email,
       first_name,
       last_name,
       password,
       user_active,
       created_at,
       updated_at
from users
where id = $1;

-- name: Update :exec
update users
set email       = $1,
    first_name  = $2,
    last_name   = $3,
    user_active = $4,
    updated_at  = $5
where id = $6;

-- name: Delete :exec
delete
from users
where id = $1;

-- name: Insert :one
insert into users (email, first_name, last_name, password, user_active, created_at, updated_at)
values ($1, $2, $3, $4, $5, $6, $7)
returning id;

-- name: ResetPassword :exec
update users
set password = $1
where id = $2;
