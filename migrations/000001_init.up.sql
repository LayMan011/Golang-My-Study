CREATE SCHEMA progress;

CREATE TABLE progress.users (
    id              SERIAL                   PRIMARY KEY,
    version         BIGINT          NOT NULL DEFAULT 1,
    full_name       VARCHAR(100)    NOT NULL CHECK(char_length(full_name) BETWEEN 3 AND 100),
    phone_number    VARCHAR(15)              CHECK(
        phone_number ~ '^\+[0-9]+$'
        AND
        char_length(phone_number) BETWEEN 10 AND 15
    )
);

CREATE TABLE progress.themes (
    id              SERIAL                  PRIMARY KEY,
    version         BIGINT       NOT NULL   DEFAULT 1,
    title           VARCHAR(100) NOT NULL   CHECK(char_length(title) BETWEEN 1 AND 100),
    description     VARCHAR(1000)           CHECK(char_length(description) BETWEEN 1 AND 1000),
    completed       BOOLEAN      NOT NULL,
    created_at      TIMESTAMPTZ  NOT NULL,
    completed_at    TIMESTAMPTZ,
    percentages     INTEGER      NOT NULL,

    CHECK (
        (completed=FALSE AND completed_at IS NULL AND percentages BETWEEN 0 AND 99)
        OR
        (completed=TRUE AND completed_at IS NOT NULL AND completed_at >= created_at AND percentages=100)
    ),

    author_user_id INTEGER NOT NULL REFERENCES progress.users(id)
);