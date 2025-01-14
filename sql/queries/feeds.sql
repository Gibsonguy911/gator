-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: GetFeeds :many
SELECT f.name as feed_name, f.url as feed_url, u.name as user_name
FROM feeds f
INNER JOIN users u ON f.user_id = u.id;

-- name: GetFeedByUrl :one
SELECT * FROM feeds WHERE url = $1;

-- name: MarkFeedFetched :exec
UPDATE feeds SET last_fetched_at = Now(), updated_at = Now() WHERE id = $1;

-- name: GetNextFeedToFetch :one
SELECT * FROM feeds ORDER BY last_fetched_at ASC NULLS FIRST LIMIT 1;
