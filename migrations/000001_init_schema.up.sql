-- Enable PostGIS extension (if using geographical functions, you can skip this if not needed)
-- CREATE EXTENSION IF NOT EXISTS postgis;

-- Create 'users' table
CREATE TABLE users (
    user_id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,  -- Added password field
    age INT NOT NULL,
    gender VARCHAR(50) NOT NULL,
    location JSONB NOT NULL,  -- Storing as JSONB
    interests TEXT[] NOT NULL,
    preferences JSONB NOT NULL,
    last_active TIMESTAMP NOT NULL,
    score INT NOT NULL DEFAULT 0
);

-- Indexes for performance optimization
CREATE INDEX idx_users_gender ON users(gender);
CREATE INDEX idx_users_age ON users(age);
CREATE INDEX idx_users_last_active ON users(last_active);

-- Ensure that 'preferences' JSONB field can be efficiently queried
CREATE INDEX idx_users_preferences ON users USING gin (preferences);
CREATE INDEX idx_users_interests ON users USING gin (interests);
