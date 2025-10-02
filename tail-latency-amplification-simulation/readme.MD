# Tail Latency Amplification Kata

This is a simple demonstration of tail latency amplification effect from *Designing Data Intensive Applications* book, implemented in Go.

## What I Built

I created a simulation of a distributed storage system with 10 nodes where client requests must query multiple nodes. The kata shows how a few slow nodes can dramatically increase overall response times, and implements a mitigation strategy using early completion.

## Key Observations

The core problem is simple but profound: when you need to wait for all nodes in a system, your performance is dictated by the slowest node. What seems like rare slow responses (5% tail latency) becomes the common case when you're waiting for 10 nodes - there's a 40% chance at least one will be slow.

## Lessons Learned

- **Concurrency is tricky**: My first implementation had race conditions with the response counter and context cancellation. Getting concurrent coordination right requires careful thinking.

- **Trade-offs are fundamental**: The mitigation strategy trades data consistency for latency. By not waiting for all nodes, we risk incomplete or stale data, but gain much better response times.

- **Real-world relevance**: This pattern explains why systems like DynamoDB use quorum reads rather than reading from all replicas.

## Code Rationale

I used Go's concurrency primitives because they're perfect for modeling distributed systems:
- Goroutines to simulate parallel node queries
- Context for cancellation when we reach our response threshold
- Atomic operations for thread-safe counting
- Channels for coordinating responses

The tricky part was handling the early completion correctly - making sure we cancel pending requests once we have enough responses without introducing race conditions.

## Questions for Further Exploration

- How would this change if nodes could return different versions of data?
- What's the optimal threshold for balancing latency and consistency?
- How do real databases like Cassandra or DynamoDB handle this in practice?
- What other mitigation strategies exist besides early completion?
