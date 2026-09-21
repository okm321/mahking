-- name: CreateMembers :copyfrom
INSERT INTO members (group_id, name) VALUES (@group_id, @name);

-- name: ListMembersByGroupID :many
SELECT
  id,
  group_id,
  name,
  created_at,
  updated_at
FROM
  members
WHERE
  group_id = @group_id
ORDER BY
  id;
