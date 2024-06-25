BEGIN;

-- Create table for users
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    email VARCHAR NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Create table for submissions
CREATE TABLE IF NOT EXISTS  submissions (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    answers JSONB NOT NULL DEFAULT '{}'::JSONB,
    risk_score INTEGER NOT NULL DEFAULT 0,
    risk_category VARCHAR NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_user_id_on_submissions ON submissions (user_id);

COMMIT;