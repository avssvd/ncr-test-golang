-- name: ListIndications :many
SELECT * FROM indications
ORDER BY sent_at;

-- name: ListIndicationsByController :many
SELECT * FROM indications
WHERE controller_serial = $1
ORDER BY sent_at;

-- name: CreateIndication :one
INSERT INTO indications (
    indication, controller_serial, sent_at
) VALUES (
             $1, $2, $3
         )
RETURNING *;
