DROP TABLE todos;
DROP TABLE users;

DROP FUNCTION IF EXISTS update_updated_at_todo();
DROP TRIGGER IF EXISTS update_todos_updated_at ON todos;