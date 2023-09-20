CREATE TABLE IF NOT EXISTS users(
    id uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    email varchar(255) NOT NULL,
    "password" text NOT NULL,
    name varchar(255),
    is_active bool NOT NULL,
    phone_number varchar(15),
    has_verified_phone bool NOT NULL DEFAULT false,
    last_authenticated_at timestamptz,
    created_at timestamptz NOT NULL DEFAULT (now()),
    updated_at timestamptz NOT NULL
);

ALTER TABLE users
    ADD CONSTRAINT users_password_key UNIQUE (id);

ALTER TABLE users
    ADD CONSTRAINT users_email_key UNIQUE (email);

