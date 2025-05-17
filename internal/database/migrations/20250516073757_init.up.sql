CREATE TABLE api_keys (
    id SERIAL PRIMARY KEY,
    user_id UUID NOT NULL,
    api_key TEXT NOT NULL UNIQUE,
    duration INTERVAL, 
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);