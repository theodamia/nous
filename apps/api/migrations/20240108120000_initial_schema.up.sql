-- Enable TimescaleDB extension
CREATE EXTENSION IF NOT EXISTS timescaledb;

-- Tool calls table (hypertable for time-series data)
-- Note: For hypertables, primary key must include the partitioning column
CREATE TABLE IF NOT EXISTS tool_calls (
    id UUID DEFAULT gen_random_uuid(),
    request_id UUID NOT NULL,
    tool_name VARCHAR(255) NOT NULL,
    duration_ms INTEGER NOT NULL,
    status VARCHAR(20) NOT NULL CHECK (status IN ('success', 'failed')),
    input_tokens INTEGER DEFAULT 0,
    output_tokens INTEGER DEFAULT 0,
    error_message TEXT,
    metadata JSONB,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (id, created_at)
);

-- Convert to hypertable for time-series optimization
SELECT create_hypertable('tool_calls', 'created_at', if_not_exists => TRUE);

-- Indexes for common queries
CREATE INDEX IF NOT EXISTS idx_tool_calls_request_id ON tool_calls(request_id);
CREATE INDEX IF NOT EXISTS idx_tool_calls_tool_name ON tool_calls(tool_name);
CREATE INDEX IF NOT EXISTS idx_tool_calls_status ON tool_calls(status);
CREATE INDEX IF NOT EXISTS idx_tool_calls_created_at ON tool_calls(created_at DESC);

-- Composite index for tool + time queries
CREATE INDEX IF NOT EXISTS idx_tool_calls_tool_time ON tool_calls(tool_name, created_at DESC);
