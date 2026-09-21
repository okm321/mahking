-- name: ListGamesByGroupID :many
SELECT
  id,
  group_id,
  note,
  played_at,
  created_at,
  updated_at
FROM
  games
WHERE
  group_id = @group_id
ORDER BY
  played_at DESC, id DESC;

-- name: CreateGame :one
INSERT INTO games (group_id, note, played_at)
VALUES (@group_id, @note, @played_at)
RETURNING id, group_id, note, played_at, created_at, updated_at;
