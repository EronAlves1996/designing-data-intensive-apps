# Replication Lag Katas

This repository contains two code katas focused on the practical implications of replication lag in distributed database systems. The katas are based on Chapter 5 (Replication) of "Designing Data-Intensive Applications" by Martin Kleppmann.

## What I Practiced

These katas helped me understand the real-world consequences of eventual consistency in single-leader replication setups. The main focus was on two specific consistency anomalies that affect user experience:

- **Read-After-Write Consistency** (Reading Your Own Writes): When a user writes data but can't see it immediately in subsequent reads because those reads go to lagging replicas.

- **Monotonic Reads Consistency**: When a user sees data "go backwards in time" because sequential reads are served by replicas with different replication lag levels.

## Key Lessons Learned

### First Kata (QuickPost Social Media)

I initially struggled with precise terminology. I understood the concepts but didn't use the exact terms from the literature. The main insight was that routing strategies need to be user-centric rather than data-centric.

### Second Kata (TaskFlow Project Management)

After revisiting the concepts, I improved significantly. I designed concrete solutions:

- For read-your-writes: Using logical timestamps passed via headers to ensure reads go to sufficiently fresh replicas
- For monotonic reads: Implementing session affinity through user ID-based routing to guarantee sequential consistency per user

The big trade-off I recognized: stronger consistency guarantees often mean reduced fault tolerance and less flexible load distribution. If you route specific users to specific replicas, you lose the ability to balance load dynamically and face challenges when replicas fail.

## Implementation Rationale

If I were to code these solutions, I'd focus on:

1. **Metadata tracking**: Each write operation would generate a consistency token (timestamp/LSN) that gets propagated to clients
2. **Routing layer intelligence**: A smart proxy or middleware that can interpret consistency tokens and make routing decisions
3. **Sticky session management**: Consistent hashing of user IDs to specific replicas while handling failover scenarios gracefully

The complexity isn't in the database queries themselves, but in the coordination layer between the application and multiple database replicas.

## Questions for Further Exploration

How do these strategies work in multi-datacenter scenarios where cross-region latency is significant? What happens during network partitions when some replicas become completely unreachable? How do we balance the cost of stronger consistency against actual user needs - maybe not all data requires the same level of consistency?
