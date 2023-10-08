CREATE TABLE IF NOT EXISTS users(
    id uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    email varchar(255) NOT NULL,
    "password" text NOT NULL,
    name varchar(255),
    birthday date,
    is_active bool NOT NULL,
    phone_number varchar(15),
    has_verified_phone bool NOT NULL DEFAULT false,
    last_authenticated_at timestamptz,
    created_at timestamptz NOT NULL DEFAULT (now()),
    updated_at timestamptz NOT NULL,
    photo varchar(255)
);

ALTER TABLE users
    ADD CONSTRAINT users_password_key UNIQUE (id);

ALTER TABLE users
    ADD CONSTRAINT users_email_key UNIQUE (email);

ALTER TABLE users
    ADD CONSTRAINT users_phone_number_key UNIQUE NULLS NOT DISTINCT (phone_number);


CREATE INDEX users_email_idx 
    ON users (email);

CREATE INDEX users_id_idx 
    ON users (id);
    
CREATE INDEX users_phone_number_idx 
    ON users (phone_number);