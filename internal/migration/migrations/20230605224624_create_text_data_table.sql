-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS text_data
(
    description text,
    tablename text DEFAULT 'text_data',
    data text not null,
    user_id text,
    CONSTRAINT fk_text_data
        FOREIGN KEY(user_id)
            REFERENCES users(user_id),
    PRIMARY KEY(description, user_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE text_data;
-- +goose StatementEnd
