CREATE INDEX IF NOT EXISTS idx_themes_author_user_id
    ON progress.themes (author_user_id);

CREATE INDEX IF NOT EXISTS idx_themes_user_user_id
    ON progress.themes_user (user_id);