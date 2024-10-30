\c blog_db;

-- Enable UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create blogs table if not exists
CREATE TABLE IF NOT EXISTS blogs (
                                     id SERIAL PRIMARY KEY,
                                     title VARCHAR(255) NOT NULL,
                                     content TEXT NOT NULL,
                                     author_id INTEGER NOT NULL,
                                     created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                                     updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes if not exist
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

-- Drop trigger if exists and recreate
DROP TRIGGER IF EXISTS update_blogs_modtime ON blogs;
CREATE TRIGGER update_blogs_modtime
    BEFORE UPDATE ON blogs
    FOR EACH ROW
EXECUTE FUNCTION update_modified_column();