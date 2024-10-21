-- +goose Up
-- +goose StatementBegin
CREATE TABLE feeds (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    url TEXT UNIQUE NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    user_id UUID NOT NULL,
    CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE feeds;
-- +goose StatementEnd
