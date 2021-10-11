-- name: GetController :one
SELECT * FROM controllers
WHERE serial = $1;

-- name: ListControllers :many
SELECT * FROM controllers
ORDER BY serial;

-- name: CreateController :exec
INSERT INTO controllers (
    serial
) VALUES (
             $1
         );

-- name: DeleteController :exec
DELETE FROM controllers
WHERE serial = $1;
