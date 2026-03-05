CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE criminal_fugitives (
    id UUID PRIMARY KEY,
    full_name VARCHAR(150) NOT NULL,
    crime_description TEXT,
    danger_level INTEGER NOT NULL,
    last_known_location VARCHAR(200),
    is_captured BOOLEAN DEFAULT FALSE,
    reward_amount NUMERIC(15,2) DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP NULL
);

CREATE TABLE audit_logs (
    id UUID PRIMARY KEY,
    entity_name VARCHAR(100) NOT NULL,
    entity_id UUID NOT NULL,
    action VARCHAR(20) NOT NULL,
    old_data JSONB,
    new_data JSONB,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);