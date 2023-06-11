-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS text_data
(
    description text,
    data text not null,
    user_id text,
    last_modified timestamp default now(),
    CONSTRAINT fk_text_data
        FOREIGN KEY(description, user_id)
            REFERENCES users_descriptions(description, user_id) ON DELETE CASCADE,
    PRIMARY KEY(description)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE text_data;
-- +goose StatementEnd
