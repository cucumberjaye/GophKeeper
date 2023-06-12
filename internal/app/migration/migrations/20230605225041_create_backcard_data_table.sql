-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS backcard_data
(
    description text,
    number text not null,
    valid_thru text not null,
    cvv text not null,
    user_id text,
    last_modified int,
    CONSTRAINT fk_bankcard_data
        FOREIGN KEY(description, user_id)
            REFERENCES users_descriptions(description, user_id) ON DELETE CASCADE,
    PRIMARY KEY(description)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE backcard_data;
-- +goose StatementEnd
