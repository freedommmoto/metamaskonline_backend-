-- name: AddNewWallet :one
INSERT INTO wallet ( metamask_wallet_id, follow_wallet, id_user, id_chain, last_block_number,  wallet_name)
VALUES ( $1, true, $2, 1, 0, '') RETURNING *;



