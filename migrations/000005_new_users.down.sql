ALTER TABLE progress.users
    DROP CONSTRAINT IF EXISTS users_email_unique;

ALTER TABLE progress.users
    DROP COLUMN IF EXISTS email,
    DROP COLUMN IF EXISTS created_at;

ALTER TABLE progress.users
    ADD COLUMN login VARCHAR(50);

UPDATE progress.users
SET login = CONCAT('user_', id);

ALTER TABLE progress.users
    ALTER COLUMN login SET NOT NULL;

ALTER TABLE progress.users
    ADD CONSTRAINT users_login_unique UNIQUE (login);

ALTER TABLE progress.users
    ADD COLUMN phone_number VARCHAR(15)
        CHECK (
            phone_number ~ '^\+[0-9]+$'
            AND char_length(phone_number) BETWEEN 10 AND 15
        );