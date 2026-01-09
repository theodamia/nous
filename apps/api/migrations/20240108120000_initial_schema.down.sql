-- Drop indexes
DROP INDEX IF EXISTS idx_tool_calls_tool_time;
DROP INDEX IF EXISTS idx_tool_calls_created_at;
DROP INDEX IF EXISTS idx_tool_calls_status;
DROP INDEX IF EXISTS idx_tool_calls_tool_name;
DROP INDEX IF EXISTS idx_tool_calls_request_id;

-- Drop hypertable
DROP TABLE IF EXISTS tool_calls;

-- Note: We don't drop the TimescaleDB extension as it might be used by other tables
