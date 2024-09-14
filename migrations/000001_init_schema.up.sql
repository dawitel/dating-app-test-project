-- Create 'users' table
CREATE TABLE users (
    user_id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    age INT NOT NULL,
    gender VARCHAR(50) NOT NULL,
    latitude FLOAT8 NOT NULL,
    longitude FLOAT8 NOT NULL,
    interests TEXT[] NOT NULL,
    preferences JSONB NOT NULL,
    last_active TIMESTAMP NOT NULL
);

-- Indexes for performance optimization
CREATE INDEX idx_users_gender ON users(gender);
CREATE INDEX idx_users_age ON users(age);
CREATE INDEX idx_users_location ON users USING GIST (ST_MakePoint(longitude, latitude));
CREATE INDEX idx_users_last_active ON users(last_active);

-- Ensure that 'preferences' JSONB field can be efficiently queried
CREATE INDEX idx_users_preferences ON users USING gin (preferences);
CREATE INDEX idx_users_interests ON users USING gin (interests);
CREATE INDEX idx_users_location ON users USING gist (ST_MakePoint(longitude, latitude));
