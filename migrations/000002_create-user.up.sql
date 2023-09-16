CREATE TABLE IF NOT EXISTS users(
    id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4 (),
    email varchar(255),
    "password" text,
    is_active bool NOT NULL,
    last_authenticated_at timestamptz,
    created_at timestamptz NOT NULL,
    updated_at timestamptz NOT NULL
);

ALTER TABLE users
    ADD CONSTRAINT users_password_key UNIQUE (id);

ALTER TABLE users
    ADD CONSTRAINT users_email_key UNIQUE (email);

