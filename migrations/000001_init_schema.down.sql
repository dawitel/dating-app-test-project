-- Drop 'users' table and associated indexes
DROP TABLE IF EXISTS users CASCADE;
DROP INDEX IF EXISTS idx_users_gender;
DROP INDEX IF EXISTS idx_users_age;
DROP INDEX IF EXISTS idx_users_location;
DROP INDEX IF EXISTS idx_users_last_active;
DROP INDEX IF EXISTS idx_users_preferences;
