-- Write your migrate up statements here

ALTER TABLE users
ADD COLUMN created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
ADD COLUMN updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW();

---- create above / drop below ----

ALTER TABLE users
DROP COLUMN created_at,
DROP COLUMN updated_at;
