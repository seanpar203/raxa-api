CREATE TABLE IF NOT EXISTS unregistered_user_contacts(
    id uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    user_id uuid REFERENCES users (id),
    email varchar(255) NOT NULL,
    phone_number varchar(15)
);

CREATE INDEX unregistered_user_contacts_user_id_idx 
    ON unregistered_user_contacts(user_id);
