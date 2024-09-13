-- 001_create_users_table.up.sql
CREATE TABLE users (
    user_id UUID PRIMARY KEY,
    name VARCHAR(100),
    age INT,
    gender VARCHAR(10),
    location POINT,
    interests TEXT[],
    preferences JSONB,
    last_active TIMESTAMPTZ
);
