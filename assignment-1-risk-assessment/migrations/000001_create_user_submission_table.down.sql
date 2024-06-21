BEGIN;

DROP INDEX IF EXISTS idx_user_id_on_submissions;
DROP TABLE IF EXISTS submissions;
DROP TABLE IF EXISTS users;

COMMIT;