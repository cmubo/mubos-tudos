CREATE TRIGGER update_todos_updated_at
    BEFORE UPDATE
    ON
        todos
    FOR EACH ROW
EXECUTE PROCEDURE update_updated_at_todo();