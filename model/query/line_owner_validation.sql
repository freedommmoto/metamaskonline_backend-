-- name: SelectLineOwnerValidationByID :one
select * from line_owner_validation where id_user=$1 order by created_at desc limit 1 ;

-- name: InsertLineOwnerValidation :one
INSERT INTO line_owner_validation (code, id_user, created_at)
VALUES ($1, $2 ,$3) RETURNING *;

-- name: SelectCodeActive :one
select code from line_owner_validation where code=$1 and deleted is null ;