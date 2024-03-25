CREATE TABLE IF NOT EXISTS todos(
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    completed BOOLEAN NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS users(
    id SERIAL PRIMARY KEY,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP 
);

-- UPDATED AT functionality on todo
DROP FUNCTION IF EXISTS update_updated_at_todo();

CREATE FUNCTION update_updated_at_todo()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_todos_updated_at
    BEFORE UPDATE
    ON todos
    FOR EACH ROW
EXECUTE PROCEDURE update_updated_at_todo();
-- UPDATED AT functionality on todo ends