ALTER TABLE progress.users 
    DROP CONSTRAINT IF EXISTS users_login_unique;

DO $$
BEGIN
  IF EXISTS (
    SELECT 1
    FROM information_schema.columns
    WHERE table_schema = 'progress'
      AND table_name   = 'users'
      AND column_name  = 'login'
  ) THEN
    ALTER TABLE progress.users
      ALTER COLUMN login DROP NOT NULL;

    ALTER TABLE progress.users
      DROP COLUMN login;
  END IF;
END $$;

ALTER TABLE progress.users
    ADD COLUMN IF NOT EXISTS created_at TIMESTAMPTZ;

UPDATE progress.users
SET created_at = NOW()
WHERE created_at IS NULL;

ALTER TABLE progress.users
    ALTER COLUMN created_at SET NOT NULL;

ALTER TABLE progress.users
    ADD COLUMN IF NOT EXISTS email TEXT;

UPDATE progress.users
SET email = CONCAT('user_', id, '@example.com');

ALTER TABLE progress.users
    ALTER COLUMN email SET NOT NULL;

ALTER TABLE progress.users
    DROP COLUMN IF EXISTS phone_number;

DO $$
BEGIN
  IF NOT EXISTS (
    SELECT 1
    FROM information_schema.table_constraints
    WHERE table_schema = 'progress'
      AND table_name   = 'users'
      AND constraint_name = 'users_email_unique'
  ) THEN
    ALTER TABLE progress.users
      ADD CONSTRAINT users_email_unique UNIQUE (email);
  END IF;
END $$;