-- +goose Up
CREATE TABLE feed_follows(
    id UUID PRIMARY KEY NOT NULL,
    feed_id UUID NOT NULL,
    user_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,  
    UNIQUE(user_id, feed_id)
);

-- +goose Down
DROP TABLE feed_follows;