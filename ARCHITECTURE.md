# DILEX Architecture
DIL Exchange

**Status:** In development. V1 core sync engine in progress.

---

## Core Problem

In DIL (Disconnected, Intermittent, Limited-bandwidth) environments, connectivity loss is the normal operating mode. When a connection window opens, a naive queue delivers in arrival order — bulk transfers block high-priority payloads.

**DILEX's guarantee: highest-priority data moves first, always.**

---

## V1 Components

### Write-Ahead Log (WAL)
Every payload is written to disk before any other operation. On crash recovery, the WAL is replayed to restore queue state.

**Implementation:** SQLite in WAL journal mode, `PRAGMA synchronous = FULL`.

### Priority Queue
Payloads ordered by: (1) priority level (1–5, 1 = highest), then (2) age (FIFO within level to prevent starvation).

**Implementation:** SQLite table with composite index on `(priority ASC, created_at ASC)`. Items remain in queue until the remote node sends an explicit acknowledgment.

### Sync Manager
Manages the connection lifecycle: monitors connectivity, drains the queue in priority order, handles disconnection, and resumes partial transfers on reconnect.

**Transport:** gRPC. Each payload has a unique ID assigned at WAL write time; the Sync Manager resumes from the last acknowledged ID on reconnect.

### Client Interface
Three public methods. Nothing else is exported.

```go
Send(data []byte, priority int) error
Receive() ([]byte, error)
Status() (Status, error)
```

---

## Design Constraints

**V1 is one-directional.** Node A sends; Node B receives. There is no bidirectional sync in V1. Callers that need both directions run two DILEX instances pointing at each other. This eliminates any need for conflict resolution — the same payload cannot originate from both nodes simultaneously.

Bidirectional sync as a first-class primitive is deferred to a future version.

---

## Deployment Model

DILEX ships as an embeddable Go library (`pkg/client`). Go callers import it directly. A binary API server for non-Go callers is a post-V1 deliverable.

### Deployment Constraints

These constraints apply to every build and every environment:

- **Air-gapped ready.** No runtime dependencies, no external fetches, no phone-home.
- **Low-resource.** Functional with as little as 2 GB RAM shared with other processes.
- **Environment parity.** Identical behavior on a developer laptop, a ruggedized field server, and AWS GovCloud.

---

## Repository Structure

```
dilex/
├── internal/
│   ├── queue/
│   ├── wal/
│   └── sync/
├── pkg/client/       # Only public API surface
├── proto/dilex.proto
└── examples/basic/
```

`internal/` is not importable by external callers. `pkg/client` is the only supported public interface.

---

## Key Technology Decisions

| Concern | Choice | Status |
|---|---|---|
| Language | Go | Decided |
| Storage | SQLite | Decided |
| Transport | gRPC | Decided |

---

## Open Questions

1. **Connectivity probe:** TCP SYN? gRPC health check? Something else?
2. **Queue depth limits:** `queue_eviction: reject` is the default — should lowest-priority eviction be supported in V1 or deferred?
3. **Clock skew:** What behavior is guaranteed when nodes have unsynchronized clocks?
4. **Caller-supplied item keys:** How are logical keys namespaced? What if the caller doesn't supply one?
