-- name: SelectLineOwnerValidationByID :one
select * from line_owner_validation where id_user=$1 order by created_at desc limit 1 ;