-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users_descriptions
(
    user_id text,
    description text,
    PRIMARY KEY(description, user_id),
    CONSTRAINT fk_users_descriptionsn
        FOREIGN KEY(user_id)
            REFERENCES users(user_id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users_descriptions;
-- +goose StatementEnd
