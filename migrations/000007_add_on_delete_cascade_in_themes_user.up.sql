ALTER TABLE progress.themes_user
    DROP CONSTRAINT IF EXISTS themes_user_user_id_fkey,
    DROP CONSTRAINT IF EXISTS themes_user_theme_id_fkey;

ALTER TABLE progress.themes_user
    ADD CONSTRAINT themes_user_user_id_fkey
        FOREIGN KEY (user_id)
        REFERENCES progress.users(id)
        ON DELETE CASCADE;

ALTER TABLE progress.themes_user
    ADD CONSTRAINT themes_user_theme_id_fkey
        FOREIGN KEY (theme_id)
        REFERENCES progress.themes(id)
        ON DELETE CASCADE;