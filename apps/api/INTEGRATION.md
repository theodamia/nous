# Agent Integration Guide

This guide explains how external agents can send tool call events to the Nous API.

## Overview

Agents send tool call events via HTTP POST requests to the Nous API. Each event represents a single tool invocation with its results, timing, and metadata.

## API Endpoint

```
POST http://localhost:8080/api/v1/events
Content-Type: application/json
```

**Production:** Replace `localhost:8080` with your Nous API server URL.

## Event Payload

### Required Fields

```json
{
  "request_id": "550e8400-e29b-41d4-a716-446655440000",  // UUID string
  "tool_name": "SearchWeb",                               // Tool identifier
  "duration_ms": 245,                                     // Duration in milliseconds
  "status": "success"                                     // "success" or "failed"
}
```

### Optional Fields

```json
{
  "input_tokens": 1250,                    // Number of input tokens
  "output_tokens": 890,                    // Number of output tokens
  "error_message": "Error details",        // Error message if failed
  "metadata": {                            // Custom metadata object
    "query": "example search",
    "results_count": 10
  },
  "timestamp": "2024-01-08T12:00:00Z"     // ISO 8601 timestamp (optional)
}
```

### Complete Example

```json
{
  "request_id": "550e8400-e29b-41d4-a716-446655440000",
  "tool_name": "SearchWeb",
  "duration_ms": 245,
  "status": "success",
  "input_tokens": 1250,
  "output_tokens": 890,
  "metadata": {
    "query": "example search query",
    "results_count": 10
  }
}
```

## Integration Examples

### Python

```python
import requests
import uuid
import time
from typing import Optional, Dict, Any

NOUS_API_URL = "http://localhost:8080/api/v1/events"

def send_tool_call_event(
    request_id: str,
    tool_name: str,
    duration_ms: int,
    status: str,
    input_tokens: Optional[int] = None,
    output_tokens: Optional[int] = None,
    error_message: Optional[str] = None,
    metadata: Optional[Dict[str, Any]] = None
):
    """Send a tool call event to Nous API."""
    payload = {
        "request_id": request_id,
        "tool_name": tool_name,
        "duration_ms": duration_ms,
        "status": status,
    }
    
    if input_tokens is not None:
        payload["input_tokens"] = input_tokens
    if output_tokens is not None:
        payload["output_tokens"] = output_tokens
    if error_message:
        payload["error_message"] = error_message
    if metadata:
        payload["metadata"] = metadata
    
    try:
        response = requests.post(
            NOUS_API_URL,
            json=payload,
            headers={"Content-Type": "application/json"},
            timeout=5
        )
        response.raise_for_status()
        return response.json()
    except requests.exceptions.RequestException as e:
        print(f"Failed to send event: {e}")
        return None

# Example usage
request_id = str(uuid.uuid4())
start_time = time.time()

# ... execute tool ...

duration_ms = int((time.time() - start_time) * 1000)
send_tool_call_event(
    request_id=request_id,
    tool_name="SearchWeb",
    duration_ms=duration_ms,
    status="success",
    input_tokens=1250,
    output_tokens=890,
    metadata={"query": "example"}
)
```

### JavaScript/TypeScript

```typescript
const NOUS_API_URL = 'http://localhost:8080/api/v1/events';

interface ToolCallEvent {
  request_id: string;
  tool_name: string;
  duration_ms: number;
  status: 'success' | 'failed';
  input_tokens?: number;
  output_tokens?: number;
  error_message?: string;
  metadata?: Record<string, any>;
}

async function sendToolCallEvent(event: ToolCallEvent): Promise<void> {
  try {
    const response = await fetch(NOUS_API_URL, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(event),
    });

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }

    const result = await response.json();
    console.log('Event sent:', result);
  } catch (error) {
    console.error('Failed to send event:', error);
  }
}

// Example usage
const requestId = crypto.randomUUID();
const startTime = Date.now();

// ... execute tool ...

const durationMs = Date.now() - startTime;
sendToolCallEvent({
  request_id: requestId,
  tool_name: 'SearchWeb',
  duration_ms: durationMs,
  status: 'success',
  input_tokens: 1250,
  output_tokens: 890,
  metadata: { query: 'example' },
});
```

### Go

```go
package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
    "time"
    "github.com/google/uuid"
)

const NousAPIURL = "http://localhost:8080/api/v1/events"

type ToolCallEvent struct {
    RequestID    string                 `json:"request_id"`
    ToolName     string                 `json:"tool_name"`
    DurationMs   int                    `json:"duration_ms"`
    Status       string                 `json:"status"`
    InputTokens  *int                   `json:"input_tokens,omitempty"`
    OutputTokens *int                   `json:"output_tokens,omitempty"`
    ErrorMessage *string                `json:"error_message,omitempty"`
    Metadata     map[string]interface{} `json:"metadata,omitempty"`
}

func SendToolCallEvent(event ToolCallEvent) error {
    jsonData, err := json.Marshal(event)
    if err != nil {
        return fmt.Errorf("failed to marshal event: %w", err)
    }

    resp, err := http.Post(NousAPIURL, "application/json", bytes.NewBuffer(jsonData))
    if err != nil {
        return fmt.Errorf("failed to send event: %w", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusCreated {
        return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
    }

    return nil
}

// Example usage
func main() {
    requestID := uuid.New().String()
    startTime := time.Now()

    // ... execute tool ...

    durationMs := int(time.Since(startTime).Milliseconds())
    inputTokens := 1250
    outputTokens := 890

    event := ToolCallEvent{
        RequestID:    requestID,
        ToolName:     "SearchWeb",
        DurationMs:   durationMs,
        Status:       "success",
        InputTokens:  &inputTokens,
        OutputTokens: &outputTokens,
        Metadata: map[string]interface{}{
            "query": "example",
        },
    }

    if err := SendToolCallEvent(event); err != nil {
        fmt.Printf("Error: %v\n", err)
    }
}
```

### cURL

```bash
curl -X POST http://localhost:8080/api/v1/events \
  -H "Content-Type: application/json" \
  -d '{
    "request_id": "550e8400-e29b-41d4-a716-446655440000",
    "tool_name": "SearchWeb",
    "duration_ms": 245,
    "status": "success",
    "input_tokens": 1250,
    "output_tokens": 890,
    "metadata": {
      "query": "example search query"
    }
  }'
```

## Integration Patterns

### 1. Wrapper Function Pattern

Wrap your tool calls to automatically track events:

```python
def tracked_tool_call(tool_name, func, *args, **kwargs):
    """Wrapper that tracks tool calls automatically."""
    request_id = str(uuid.uuid4())
    start_time = time.time()
    
    try:
        result = func(*args, **kwargs)
        duration_ms = int((time.time() - start_time) * 1000)
        
        send_tool_call_event(
            request_id=request_id,
            tool_name=tool_name,
            duration_ms=duration_ms,
            status="success",
            metadata={"args": str(args), "kwargs": str(kwargs)}
        )
        return result
    except Exception as e:
        duration_ms = int((time.time() - start_time) * 1000)
        
        send_tool_call_event(
            request_id=request_id,
            tool_name=tool_name,
            duration_ms=duration_ms,
            status="failed",
            error_message=str(e),
            metadata={"args": str(args), "kwargs": str(kwargs)}
        )
        raise

# Usage
result = tracked_tool_call("SearchWeb", search_web, query="example")
```

### 2. Decorator Pattern (Python)

```python
from functools import wraps
import uuid
import time

def track_tool_call(tool_name: str):
    """Decorator to track tool calls."""
    def decorator(func):
        @wraps(func)
        def wrapper(*args, **kwargs):
            request_id = str(uuid.uuid4())
            start_time = time.time()
            
            try:
                result = func(*args, **kwargs)
                duration_ms = int((time.time() - start_time) * 1000)
                
                send_tool_call_event(
                    request_id=request_id,
                    tool_name=tool_name,
                    duration_ms=duration_ms,
                    status="success"
                )
                return result
            except Exception as e:
                duration_ms = int((time.time() - start_time) * 1000)
                
                send_tool_call_event(
                    request_id=request_id,
                    tool_name=tool_name,
                    duration_ms=duration_ms,
                    status="failed",
                    error_message=str(e)
                )
                raise
        return wrapper
    return decorator

# Usage
@track_tool_call("SearchWeb")
def search_web(query: str):
    # ... tool implementation ...
    pass
```

### 3. Context Manager Pattern (Python)

```python
from contextlib import contextmanager
import uuid
import time

@contextmanager
def track_tool_call(tool_name: str, request_id: str = None):
    """Context manager to track tool calls."""
    if request_id is None:
        request_id = str(uuid.uuid4())
    
    start_time = time.time()
    error = None
    
    try:
        yield request_id
    except Exception as e:
        error = e
        raise
    finally:
        duration_ms = int((time.time() - start_time) * 1000)
        
        send_tool_call_event(
            request_id=request_id,
            tool_name=tool_name,
            duration_ms=duration_ms,
            status="failed" if error else "success",
            error_message=str(error) if error else None
        )

# Usage
with track_tool_call("SearchWeb") as request_id:
    result = search_web("example")
```

## Request ID Management

**Important:** Use the same `request_id` for all tool calls within a single agent request/chain.

```python
# At the start of an agent request
request_id = str(uuid.uuid4())

# All tool calls in this request use the same ID
send_tool_call_event(request_id, "SearchWeb", ...)
send_tool_call_event(request_id, "ReadFile", ...)
send_tool_call_event(request_id, "WriteFile", ...)

# This allows tracking the full chain in the dashboard
```

## Error Handling

### Best Practices

1. **Non-blocking:** Don't let event sending failures break your agent
2. **Retry logic:** Implement retries for transient failures
3. **Timeout:** Set reasonable timeouts (e.g., 5 seconds)
4. **Logging:** Log failures for debugging

```python
def send_tool_call_event_safe(event_data):
    """Send event with retry and error handling."""
    max_retries = 3
    for attempt in range(max_retries):
        try:
            response = requests.post(
                NOUS_API_URL,
                json=event_data,
                timeout=5
            )
            response.raise_for_status()
            return True
        except requests.exceptions.RequestException as e:
            if attempt == max_retries - 1:
                # Log but don't fail
                logger.warning(f"Failed to send event after {max_retries} attempts: {e}")
                return False
            time.sleep(0.5 * (attempt + 1))  # Exponential backoff
    return False
```

## Response Format

### Success Response

```json
{
  "status": "ok"
}
```

**HTTP Status:** `201 Created`

### Error Responses

**400 Bad Request** - Invalid payload:
```json
{
  "error": "Missing required fields"
}
```

**500 Internal Server Error** - Server error:
```json
{
  "error": "Failed to ingest event"
}
```

## Field Validation

- **`request_id`**: Must be a valid UUID (e.g., `550e8400-e29b-41d4-a716-446655440000`)
- **`tool_name`**: Non-empty string
- **`status`**: Must be exactly `"success"` or `"failed"`
- **`duration_ms`**: Non-negative integer
- **`input_tokens`**: Non-negative integer (optional)
- **`output_tokens`**: Non-negative integer (optional)

## Performance Considerations

1. **Async sending:** Send events asynchronously to avoid blocking agent execution
2. **Batching:** Consider batching multiple events (future feature)
3. **Connection pooling:** Reuse HTTP connections when possible
4. **Fire and forget:** Don't wait for response in critical paths

```python
import asyncio
import aiohttp

async def send_tool_call_event_async(event_data):
    """Send event asynchronously."""
    async with aiohttp.ClientSession() as session:
        try:
            async with session.post(
                NOUS_API_URL,
                json=event_data,
                timeout=aiohttp.ClientTimeout(total=5)
            ) as response:
                return await response.json()
        except Exception as e:
            print(f"Failed to send event: {e}")
            return None

# Fire and forget
asyncio.create_task(send_tool_call_event_async(event_data))
```

## Authentication (Future)

Currently, the API has no authentication. In production, you'll likely need:

- **API Key** authentication
- **JWT** tokens
- **OAuth2** integration

When implemented, include the token in headers:

```python
headers = {
    "Content-Type": "application/json",
    "Authorization": f"Bearer {api_key}"
}
```

## SDK (Future)

A Nous SDK will be available in `packages/sdk/` for easier integration:

```python
from nous_sdk import NousClient

client = NousClient(api_url="http://localhost:8080", api_key="...")

with client.track_tool_call("SearchWeb") as event:
    result = search_web("example")
    event.metadata = {"query": "example"}
```

## Testing

Use the example script to test your integration:

```bash
cd apps/api/examples
./agent-example.sh
```

Or manually test with curl:

```bash
curl -X POST http://localhost:8080/api/v1/events \
  -H "Content-Type: application/json" \
  -d '{
    "request_id": "550e8400-e29b-41d4-a716-446655440000",
    "tool_name": "TestTool",
    "duration_ms": 100,
    "status": "success"
  }'
```

## Troubleshooting

### Connection Refused

- Check if Nous API is running: `curl http://localhost:8080/health`
- Verify the API URL is correct
- Check firewall/network settings

### Invalid Request Body

- Ensure `Content-Type: application/json` header is set
- Validate JSON syntax
- Check all required fields are present

### UUID Validation Error

- Ensure `request_id` is a valid UUID format
- Use `uuid.uuid4()` or equivalent in your language

## Next Steps

1. **Implement tracking** in your agent code
2. **Test integration** with example script
3. **Monitor dashboard** at `http://localhost:5173/observability`
4. **Watch for SDK** release for easier integration
