#!/usr/bin/env bash
# Development script - runs both backend and frontend with proper signal handling

set -e

# Configuration
BACKEND_CMD="air"
FRONTEND_CMD="cd web && pnpm run dev"

# Store PIDs
BACKEND_PID=""
FRONTEND_PID=""

# Cleanup function - called on exit or signal
cleanup() {
    echo ""
    echo "🛑 Shutting down development environment..."

    if [ -n "$BACKEND_PID" ]; then
        echo "  Stopping backend (PID: $BACKEND_PID)..."
        # Send SIGTERM for graceful shutdown
        if kill -0 "$BACKEND_PID" 2>/dev/null; then
            kill -TERM "$BACKEND_PID" 2>/dev/null || true
            # Wait up to 5 seconds for graceful shutdown
            timeout 5s tail --pid="$BACKEND_PID" -f /dev/null 2>/dev/null || \
                kill -KILL "$BACKEND_PID" 2>/dev/null || true
        fi
    fi

    if [ -n "$FRONTEND_PID" ]; then
        echo "  Stopping frontend (PID: $FRONTEND_PID)..."
        if kill -0 "$FRONTEND_PID" 2>/dev/null; then
            kill -TERM "$FRONTEND_PID" 2>/dev/null || true
            timeout 5s tail --pid="$FRONTEND_PID" -f /dev/null 2>/dev/null || \
                kill -KILL "$FRONTEND_PID" 2>/dev/null || true
        fi
    fi

    echo "✓ Development environment stopped"
    exit 0
}

# Trap signals for cleanup
trap cleanup INT TERM EXIT

# Start backend in background
echo "Starting backend..."
eval "$BACKEND_CMD" &
BACKEND_PID=$!

# Start frontend in background
echo "Starting frontend..."
eval "$FRONTEND_CMD" &
FRONTEND_PID=$!

echo ""
echo "✓ Both services started"
echo "  Backend PID: $BACKEND_PID"
echo "  Frontend PID: $FRONTEND_PID"
echo ""
echo "Press Ctrl+C to stop all services"
echo ""

# Wait for either process to exit (or signal to be caught)
wait -n $BACKEND_PID $FRONTEND_PID 2>/dev/null || true

# If we reach here, one process exited - run cleanup
cleanup
