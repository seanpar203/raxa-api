CREATE TYPE otp_channel AS ENUM ('sms', 'email');

CREATE TABLE IF NOT EXISTS otp(
    id uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    user_id uuid NOT NULL,
    channel otp_channel NOT NULL,
    code varchar(6) NOT NULL,
    created_at timestamptz NOT NULL DEFAULT (now()),
    expires_at timestamptz NOT NULL
);

ALTER TABLE otp
    ADD CONSTRAINT user_id_fk FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE;

CREATE INDEX idx_otp_user_id ON refresh_tokens(user_id);