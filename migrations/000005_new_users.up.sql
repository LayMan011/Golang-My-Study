ALTER TABLE progress.users
    DROP CONSTRAINT users_login_unique,
    DROP COLUMN login,
    DROP COLUMN phone_number,
    ADD COLUMN created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    ADD COLUMN email TEXT;

UPDATE progress.users
SET email = CONCAT('user_', id, '@example.com');

ALTER TABLE progress.users
    ALTER COLUMN email SET NOT NULL,
    ADD CONSTRAINT users_email_unique UNIQUE (email);