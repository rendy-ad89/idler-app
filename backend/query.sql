-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByUsername :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (
  username, password
) VALUES (
  $1, $2
)
RETURNING *;

-- name: SaveProgress :exec
UPDATE users
SET cash = $2, last_save = NOW()
WHERE id = $1;

-- name: GetMachines :many
SELECT * FROM machines
ORDER BY priority ASC;

-- name: GetUsersMachines :many
SELECT user_id, machine_id, priority, name, description, level, generation, price, increment, type FROM users_machines um
JOIN machines m ON m.id = um.machine_id
WHERE user_id = $1
ORDER BY priority ASC;

-- name: GetUsersGenerators :many
SELECT priority, name, description, level, generation, price, increment FROM users_machines um
JOIN machines m ON m.id = um.machine_id
WHERE user_id = $1 AND type = 'generator'
ORDER BY priority ASC;

-- name: GetUsersAmplifiers :many
SELECT priority, name, description, level, generation, price, increment FROM users_machines um
JOIN machines m ON m.id = um.machine_id
WHERE user_id = $1 AND type = 'amplifier'
ORDER BY priority ASC;

-- name: CreateUsersMachines :batchexec
INSERT INTO users_machines (
  user_id, machine_id
) VALUES (
  $1, $2
);

-- name: UpdateUsersMachines :batchexec
UPDATE users_machines
SET level = $3
WHERE user_id = $1 AND machine_id = $2;
