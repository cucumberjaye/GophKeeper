-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS login_password
(
    description text,
    login text not null,
    password text not null,
    user_id text,
    last_modified timestamp default now(),
    CONSTRAINT fk_login_password
       FOREIGN KEY(description, user_id)
            REFERENCES users_descriptions(description, user_id) ON DELETE CASCADE,
    PRIMARY KEY(description)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE login_password;
-- +goose StatementEnd
