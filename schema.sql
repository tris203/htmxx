CREATE TABLE bookmarks (
tweet_id INTEGER NOT NULL,
username TEXT NOT NULL,
PRIMARY KEY (tweet_id, username)
);

CREATE TABLE likes (
tweet_id INTEGER NOT NULL,
username TEXT NOT NULL,
PRIMARY KEY (tweet_id, username)
);

CREATE TABLE tweets (
tweet_id INTEGER PRIMARY KEY AUTOINCREMENT,
author TEXT NOT NULL,
content TEXT NOT NULL,
created DATETIME DEFAULT CURRENT_TIMESTAMP,
like_count INTEGER DEFAULT 0 NOT NULL
);
