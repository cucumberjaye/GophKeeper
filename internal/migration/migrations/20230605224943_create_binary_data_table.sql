-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS binary_data
(
    description text,
    tablename text DEFAULT 'binary_data',
    data text not null,
    user_id text,
    CONSTRAINT fk_binary_data
        FOREIGN KEY(user_id)
            REFERENCES users(user_id),
    PRIMARY KEY(description, user_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE binary_data;
-- +goose StatementEnd
