ALTER TABLE progress.themes_user 
    ADD COLUMN IF NOT EXISTS total_lessons INT NOT NULL DEFAULT 0,
    ADD COLUMN IF NOT EXISTS completed_lessons INT NOT NULL DEFAULT 0;
