-- name: SelectUserID :one
select *
from users
where id_user = $1
  and deleted is null;

-- name: SelectUserID :one
INSERT INTO users
( username, password, owner_validation )
VALUES ( 'asdf', 'asdf', false);
