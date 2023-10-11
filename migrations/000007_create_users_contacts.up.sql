CREATE TABLE IF NOT EXISTS users_contacts(
    user_id uuid REFERENCES users (id),
    contact_id uuid REFERENCES users (id),
    CONSTRAINT user_contact_pkey PRIMARY KEY (user_id, contact_id)
);

CREATE INDEX users_contacts_user_id_idx 
    ON users_contacts(user_id);