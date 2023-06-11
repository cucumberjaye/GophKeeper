-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS binary_data
(
    description text,
    data bytea not null,
    user_id text,
    last_modified timestamp default now(),
    CONSTRAINT fk_binary_data
        FOREIGN KEY(description, user_id)
            REFERENCES users_descriptions(description, user_id) ON DELETE CASCADE,
    PRIMARY KEY(description)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE binary_data;
-- +goose StatementEnd
