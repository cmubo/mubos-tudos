BEGIN;

ALTER TABLE users REMOVE UNIQUE (email);

COMMIT;