CREATE SCHEMA IF NOT EXISTS progress;

CREATE TABLE progress.themes_for_user (
    id              SERIAL                  PRIMARY KEY,
    version         BIGINT       NOT NULL   DEFAULT 1,
    completed       BOOLEAN      NOT NULL   DEFAULT FALSE,
    completed_at    TIMESTAMPTZ,
    percentages     INTEGER      NOT NULL   DEFAULT 0,
    theme_id        INTEGER      NOT NULL   REFERENCES progress.themes(id),
    user_id         INTEGER      NOT NULL   REFERENCES progress.users(id),

    CHECK (
        (completed=FALSE AND completed_at IS NULL AND percentages BETWEEN 0 AND 99)
        OR
        (completed=TRUE AND completed_at IS NOT NULL AND percentages=100)
    ),
    
    UNIQUE(theme_id, user_id)
);

-- ИЗМЕНЯЕМ ТАБЛИЦУ themes (добавляем новые колонки)

ALTER TABLE progress.themes ADD COLUMN subject VARCHAR(1000) NOT NULL CHECK(char_length(subject) BETWEEN 1 AND 1000);

ALTER TABLE progress.themes ADD COLUMN rating NUMERIC(2, 1) CHECK(rating >= 0 AND rating <= 5);

ALTER TABLE progress.themes ADD COLUMN number_of_ratings INT DEFAULT 0;

ALTER TABLE progress.themes ADD COLUMN number_of_users INT DEFAULT 0;

ALTER TABLE progress.themes ADD COLUMN price INT NOT NULL DEFAULT 0;

-- УДАЛЯЕМ СТАРЫЕ КОЛОНКИ

ALTER TABLE progress.themes DROP COLUMN IF EXISTS completed;

ALTER TABLE progress.themes DROP COLUMN IF EXISTS completed_at;

ALTER TABLE progress.themes DROP COLUMN IF EXISTS percentages;
