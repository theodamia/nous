# Nous

Tool-use observability dashboard for agent systems in production.

## Structure
```
nous/
├── apps/
│   ├── api/        # Go backend
│   └── web/        # React frontend (Vite)
├── packages/
│   └── sdk/        # Instrumentation SDK
├── deploy/
│   └── docker/     # Docker configs
└── Makefile
```

## Quick Start
```bash
# Start infrastructure
make docker-up

# Run both frontend and backend
make dev
```

## Services

| Service | Port | Description |
|---------|------|-------------|
| API | 8080 | Go backend + WebSocket |
| Web | 5173 | Vite dev server |
| TimescaleDB | 5432 | Time-series database |
| Redis | 6379 | Pub/sub (optional) |
