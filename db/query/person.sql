-- name: CreatePerson :one
INSERT INTO person (
    name, age
) VALUES (
             $1, $2
         )
    RETURNING *;