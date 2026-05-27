-- Creates foods table

CREATE TABLE IF NOT EXISTS foods (
    id UUID PRIMARY KEY,
    image_url TEXT NOT NULL,
    comment TEXT,
    calories INT,
    proteins NUMERIC,
    fats NUMERIC,
    carbs NUMERIC,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);