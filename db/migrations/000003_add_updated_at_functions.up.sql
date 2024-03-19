DROP FUNCTION IF EXISTS update_updated_at_todo();

CREATE FUNCTION update_updated_at_todo()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ language 'plpgsql';