ALTER TABLE progress.users 
    DROP CONSTRAINT IF EXISTS users_login_unique;

ALTER TABLE progress.users 
    ALTER COLUMN login DROP NOT NULL,
    ALTER COLUMN password DROP NOT NULL;

ALTER TABLE progress.users 
    DROP COLUMN IF EXISTS login,
    DROP COLUMN IF EXISTS password;
