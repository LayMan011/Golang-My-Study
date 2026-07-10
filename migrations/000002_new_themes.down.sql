DROP TABLE IF EXISTS progress.themes_for_user;

ALTER TABLE progress.themes
    DROP COLUMN IF EXISTS subject,
    DROP COLUMN IF EXISTS rating,
    DROP COLUMN IF EXISTS all_ratings,
    DROP COLUMN IF EXISTS number_of_ratings,
    DROP COLUMN IF EXISTS number_of_users,
    DROP COLUMN IF EXISTS price;

ALTER TABLE progress.themes
    ADD COLUMN IF NOT EXISTS completed BOOLEAN NOT NULL DEFAULT FALSE,
    ADD COLUMN IF NOT EXISTS completed_at TIMESTAMPTZ,
    ADD COLUMN IF NOT EXISTS percentages INTEGER NOT NULL DEFAULT 0;

ALTER TABLE progress.themes
    ADD
    CHECK (
        (completed = FALSE AND completed_at IS NULL AND percentages BETWEEN 0 AND 99)
        OR
        (completed = TRUE AND completed_at IS NOT NULL AND percentages = 100)
    );