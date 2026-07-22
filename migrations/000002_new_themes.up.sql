CREATE SCHEMA IF NOT EXISTS progress;

CREATE TABLE IF NOT EXISTS progress.themes_user (
    id              SERIAL                PRIMARY KEY,
    version         BIGINT       NOT NULL DEFAULT 1,
    completed       BOOLEAN      NOT NULL DEFAULT FALSE,
    addition_at     TIMESTAMPTZ  NOT NULL, 
    completed_at    TIMESTAMPTZ,
    percentages     INTEGER      NOT NULL DEFAULT 0,
    theme_id        INTEGER      NOT NULL REFERENCES progress.themes(id),
    user_id         INTEGER      NOT NULL REFERENCES progress.users(id),

    CHECK (
        (completed = FALSE AND completed_at IS NULL AND percentages BETWEEN 0 AND 99)
        OR
        (completed = TRUE AND completed_at IS NOT NULL AND completed_at >= addition_at AND percentages = 100)
    ),

    UNIQUE (theme_id, user_id)
);

ALTER TABLE progress.themes
    ADD COLUMN IF NOT EXISTS subject VARCHAR(1000)
        CHECK (char_length(subject) BETWEEN 1 AND 1000);

UPDATE progress.themes
SET subject = 'Без названия'
WHERE subject IS NULL;

ALTER TABLE progress.themes ALTER COLUMN subject SET NOT NULL;

ALTER TABLE progress.themes
    ADD COLUMN IF NOT EXISTS rating NUMERIC(2,1)
        CHECK (rating >= 0 AND rating <= 5);

ALTER TABLE progress.themes
    ADD COLUMN IF NOT EXISTS all_ratings INT DEFAULT 0;

ALTER TABLE progress.themes
    ADD COLUMN IF NOT EXISTS number_of_ratings INT DEFAULT 0;

ALTER TABLE progress.themes
    ADD COLUMN IF NOT EXISTS number_of_users INT DEFAULT 0;

ALTER TABLE progress.themes
    ADD COLUMN IF NOT EXISTS price INT NOT NULL DEFAULT 0;

ALTER TABLE progress.themes
    ADD COLUMN IF NOT EXISTS level VARCHAR(40) NOT NULL DEFAULT 'beginner';

ALTER TABLE progress.themes
    ADD COLUMN IF NOT EXISTS duration VARCHAR(40) NOT NULL DEFAULT '3 месяца';

ALTER TABLE progress.themes
    ADD COLUMN IF NOT EXISTS format VARCHAR(40) NOT NULL DEFAULT 'text';


ALTER TABLE progress.themes DROP COLUMN IF EXISTS completed;
ALTER TABLE progress.themes DROP COLUMN IF EXISTS completed_at;
ALTER TABLE progress.themes DROP COLUMN IF EXISTS percentages;