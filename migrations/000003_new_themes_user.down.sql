ALTER TABLE progress.themes_user
    DROP COLUMN IF EXISTS total_lessons,
    DROP COLUMN IF EXISTS completed_lessons;