-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS login_password
(
    description text,
    tablename text DEFAULT 'login_password',
    login text not null,
    password text not null,
    user_id text,
    CONSTRAINT fk_login_password
        FOREIGN KEY(user_id)
            REFERENCES users(user_id),
    PRIMARY KEY(description, user_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE login_password;
-- +goose StatementEnd
