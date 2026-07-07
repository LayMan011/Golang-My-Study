-- УДАЛЯЕМ НОВУЮ ТАБЛИЦУ
DROP TABLE IF EXISTS progress.themes_for_user CASCADE;

-- УДАЛЯЕМ НОВЫЕ КОЛОНКИ ИЗ themes
ALTER TABLE progress.themes 
    DROP COLUMN IF EXISTS subject,
    DROP COLUMN IF EXISTS rating,
    DROP COLUMN IF EXISTS number_of_ratings,
    DROP COLUMN IF EXISTS number_of_users,
    DROP COLUMN IF EXISTS price;

-- ВОССТАНАВЛИВАЕМ СТАРЫЕ КОЛОНКИ
ALTER TABLE progress.themes 
    ADD COLUMN IF NOT EXISTS completed BOOLEAN NOT NULL DEFAULT FALSE,
    ADD COLUMN IF NOT EXISTS completed_at TIMESTAMPTZ,
    ADD COLUMN IF NOT EXISTS percentages INTEGER NOT NULL DEFAULT 0;

-- ВОССТАНАВЛИВАЕМ CHECK CONSTRAINT
ALTER TABLE progress.themes 
    ADD CHECK (
        (completed=FALSE AND completed_at IS NULL AND percentages BETWEEN 0 AND 99)
        OR
        (completed=TRUE AND completed_at IS NOT NULL AND percentages >= 0 AND percentages=100)
    );
