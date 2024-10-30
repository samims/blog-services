-- Create user with password if not exists
DO
$do$
BEGIN
   IF NOT EXISTS (
      SELECT FROM pg_catalog.pg_roles
      WHERE  rolname = 'blog_user') THEN

      CREATE ROLE blog_user LOGIN PASSWORD 'blog_password';
   END IF;
END
$do$;

-- Grant privileges
ALTER ROLE blog_user WITH LOGIN;
ALTER USER blog_user WITH SUPERUSER;

-- Create database if not exists
CREATE DATABASE blog_db
    WITH 
    OWNER = blog_user
    ENCODING = 'UTF8'
    LC_COLLATE = 'en_US.utf8'
    LC_CTYPE = 'en_US.utf8'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1;

-- Connect to the blog_db database
\c blog_db;

-- Grant privileges on database
GRANT ALL PRIVILEGES ON DATABASE blog_db TO blog_user;

-- Enable UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create blogs table
CREATE TABLE IF NOT EXISTS blogs (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    author_id INTEGER NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Grant privileges on all tables
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO blog_user;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO blog_user;

-- Create indexes
CREATE INDEX IF NOT EXISTS idx_blogs_author_id ON blogs(author_id);
CREATE INDEX IF NOT EXISTS idx_blogs_created_at ON blogs(created_at DESC);

-- Function to update updated_at timestamp
CREATE OR REPLACE FUNCTION update_modified_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ language 'plpgsql';

-- Create trigger
CREATE TRIGGER update_blogs_modtime
    BEFORE UPDATE ON blogs
    FOR EACH ROW
    EXECUTE FUNCTION update_modified_column();