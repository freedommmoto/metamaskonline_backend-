-- name: SelectUserID :one
select *
from users
where id_user = $1
  and deleted is null;

-- name: InsertNewUser :one
INSERT INTO users
( username, password, owner_validation )
VALUES ( $1, $2, false)
RETURNING *;