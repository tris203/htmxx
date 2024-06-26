-- name: GetTimeline :many
SELECT sqlc.embed(tweets), likes.username IS NOT NULL AS likedByUser, bookmarks.username IS NOT NULL AS bookmarkedByUser
FROM tweets
LEFT JOIN likes ON likes.username = ? AND tweets.tweet_id = likes.tweet_id
LEFT JOIN bookmarks ON bookmarks.username = ? AND tweets.tweet_id = bookmarks.tweet_id
WHERE tweets.author LIKE ?
AND tweets.tweet_id <= ?
AND tweets.parent_tweet_id IS NULL
ORDER BY tweets.tweet_id DESC
LIMIT 10;


-- name: GetAllTimeline :many
SELECT sqlc.embed(tweets), likes.username IS NOT NULL AS likedByUser, bookmarks.username IS NOT NULL AS bookmarkedByUser
FROM tweets
LEFT JOIN likes ON likes.username = ? AND tweets.tweet_id = likes.tweet_id
LEFT JOIN bookmarks ON bookmarks.username = ? AND tweets.tweet_id = bookmarks.tweet_id
WHERE tweets.tweet_id <= ?
AND tweets.parent_tweet_id IS NULL
ORDER BY tweets.tweet_id DESC
LIMIT 10;

-- name: GetTweetLikers :many
SELECT username FROM likes
WHERE tweet_id = ?;

-- name: GetTweet :many
SELECT sqlc.embed(tweets), likes.username IS NOT NULL AS likedByUser, bookmarks.username IS NOT NULL AS bookmarkedByUser
FROM tweets
LEFT JOIN likes ON likes.username = ? AND tweets.tweet_id = likes.tweet_id
LEFT JOIN bookmarks ON bookmarks.username = ? AND tweets.tweet_id = bookmarks.tweet_id
WHERE tweets.tweet_id = ?
OR tweets.parent_tweet_id = ?;

-- name: GetLikedTweets :many
SELECT sqlc.embed(tweets), likes.username IS NOT NULL AS likedByUser, bookmarks.username IS NOT NULL AS bookmarkedByUser
FROM tweets
LEFT JOIN likes ON likes.username = ? AND tweets.tweet_id = likes.tweet_id
LEFT JOIN bookmarks ON bookmarks.username = ? AND tweets.tweet_id = bookmarks.tweet_id
WHERE tweets.tweet_id IN (SELECT tweet_id FROM likes WHERE likes.username = ?);

-- name: GetBookmarkedTweets :many
SELECT sqlc.embed(tweets), likes.username IS NOT NULL AS likedByUser
FROM tweets
LEFT JOIN likes ON likes.username = ? AND tweets.tweet_id = likes.tweet_id
WHERE tweets.tweet_id IN (SELECT tweet_id FROM bookmarks WHERE bookmarks.username = ?);

-- name: AddTweet :one
INSERT INTO tweets (author, content)
VALUES (?, ?)
RETURNING tweet_id;

-- name: DeleteTweet :one
DELETE FROM tweets
WHERE tweet_id = ?
AND author = ?
RETURNING *;

-- name: LikeTweet :exec
INSERT INTO likes (username, tweet_id)
VALUES (?, ?);

-- name: GetLikeCount :one
SELECT like_count FROM tweets
WHERE tweet_id = ?;

-- name: UnlikeTweet :exec
DELETE FROM likes
WHERE username = ?
AND tweet_id = ?;

-- name: BookmarkTweet :exec
INSERT INTO bookmarks (username, tweet_id)
VALUES (?, ?);

-- name: UnbookmarkTweet :exec
DELETE FROM bookmarks
WHERE username = ?
AND tweet_id = ?;

-- name: SearchTweets :many
SELECT * FROM tweets
WHERE content LIKE ?
ORDER BY created DESC;

-- name: ReplyTweet :one
INSERT INTO tweets (author, content, parent_tweet_id)
VALUES (?, ?, ?)
RETURNING tweet_id;
