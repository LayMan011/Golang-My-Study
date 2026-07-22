ALTER TABLE progress.users
    ADD COLUMN IF NOT EXISTS login VARCHAR(50),
    ADD COLUMN IF NOT EXISTS password TEXT;

UPDATE progress.users
SET login = CONCAT('user_', id),
    password = 'some_password_hash'
WHERE login IS NULL OR password IS NULL;

ALTER TABLE progress.users ALTER COLUMN login SET NOT NULL;
ALTER TABLE progress.users ALTER COLUMN password SET NOT NULL;
ALTER TABLE progress.users ADD CONSTRAINT users_login_unique UNIQUE (login);