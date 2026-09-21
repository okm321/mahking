-- name: CreateGameScores :copyfrom
INSERT INTO game_scores (
  game_id,
  group_id,
  member_id,
  seat,
  ranking,
  raw_score,
  point,
  chip_count,
  is_busted
)
VALUES (
  @game_id,
  @group_id,
  @member_id,
  @seat,
  @ranking,
  @raw_score,
  @point,
  @chip_count,
  @is_busted
);

-- name: ListGameScoresByGameID :many
SELECT
  id,
  game_id,
  group_id,
  member_id,
  seat,
  ranking,
  raw_score,
  point,
  chip_count,
  is_busted,
  created_at,
  updated_at
FROM
  game_scores
WHERE
  game_id = @game_id
ORDER BY
  ranking;
