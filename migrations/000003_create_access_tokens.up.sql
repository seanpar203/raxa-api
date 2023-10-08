CREATE TABLE IF NOT EXISTS access_tokens(
    token uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    user_id uuid NOT NULL,
    valid_until timestamptz NOT NULL,
    created_at timestamptz NOT NULL DEFAULT (now()),
    updated_at timestamptz NOT NULL
);

ALTER TABLE access_tokens
    ADD CONSTRAINT user_id_fk FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE;
    

CREATE INDEX access_tokens_token_idx ON access_tokens(token);
CREATE INDEX access_tokens_user_id_idx ON access_tokens(user_id);