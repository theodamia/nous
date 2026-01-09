package models

import (
	"time"

	"github.com/google/uuid"
)

// ToolCall represents a single tool call event
type ToolCall struct {
	ID           uuid.UUID              `json:"id"`
	RequestID    uuid.UUID              `json:"request_id"`
	ToolName     string                 `json:"tool_name"`
	DurationMs   int                    `json:"duration_ms"`
	Status       string                 `json:"status"` // "success" or "failed"
	InputTokens  int                    `json:"input_tokens"`
	OutputTokens int                    `json:"output_tokens"`
	ErrorMessage *string                `json:"error_message,omitempty"`
	Metadata     map[string]interface{} `json:"metadata,omitempty"`
	CreatedAt    time.Time              `json:"created_at"`
}

// ToolCallEvent is the incoming event from agents
type ToolCallEvent struct {
	RequestID    string                 `json:"request_id"`
	ToolName     string                 `json:"tool_name"`
	DurationMs   int                    `json:"duration_ms"`
	Status       string                 `json:"status"`
	InputTokens  *int                   `json:"input_tokens,omitempty"`
	OutputTokens *int                   `json:"output_tokens,omitempty"`
	ErrorMessage *string                `json:"error_message,omitempty"`
	Metadata     map[string]interface{} `json:"metadata,omitempty"`
	Timestamp    *time.Time             `json:"timestamp,omitempty"`
}

// ToolCallDataPoint represents aggregated tool call data for a time period
type ToolCallDataPoint struct {
	Hour     string `json:"hour"`
	Success  int    `json:"success"`
	Failures int    `json:"failures"`
}

// LatencyDataPoint represents latency percentiles for a tool
type LatencyDataPoint struct {
	Tool string  `json:"tool"`
	P50  float64 `json:"p50"`
	P95  float64 `json:"p95"`
	P99  float64 `json:"p99"`
}

// TokenUsageDataPoint represents token usage for a time period
type TokenUsageDataPoint struct {
	Hour   string `json:"hour"`
	Input  int    `json:"input"`
	Output int    `json:"output"`
}

// FailureRateDataPoint represents failure rate for a time period
type FailureRateDataPoint struct {
	Hour           string  `json:"hour"`
	FailurePercent float64 `json:"failurePercent"`
}

// MetricsOverview represents aggregated metrics
type MetricsOverview struct {
	TotalCalls    int64   `json:"total_calls"`
	AvgLatencyMs  float64 `json:"avg_latency_ms"`
	TotalTokens   int64   `json:"total_tokens"`
	FailureRate   float64 `json:"failure_rate"`
	ChangePercent float64 `json:"change_percent"`
}
