-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS backcard_data
(
    description text,
    tablename text DEFAULT 'backcard_data',
    number text not null,
    valid_thru text not null,
    cvv text not null,
    user_id text,
    CONSTRAINT fk_bankcard_data
        FOREIGN KEY(user_id)
            REFERENCES users(user_id),
    PRIMARY KEY(description, user_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE backcard_data;
-- +goose StatementEnd
