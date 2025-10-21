# ConcertFrenzy Ticket System Kata

This is a design kata focusing on distributed systems consistency patterns from "Designing Data Intensive Applications" Chapter 9. The scenario tackles a real-world ticket booking system with consistency problems.

## The Problem

We're fixing ConcertFrenzy's ticket hold system that has:

- Double-selling tickets during high traffic
- Users seeing "Hold Failed" but getting charged later
- Inconsistent views of hold status after purchase

## Key Observations From My Solution

I initially understood linearizability conceptually but missed the architectural shift with Total Order Broadcast. My thinking was still stuck in the 2PC coordination mindset rather than seeing the log as the single source of truth.

The big insight was that total order broadcast isn't about coordinating services - it's about making them all follow the same immutable command log. The atomicity moves from "coordinating multiple resources" to "appending one command to the log."

## Lessons Learned

- **Linearizability** gives you the "single up-to-date copy" illusion but costs you availability during partitions
- **Total Order Broadcast** re-architects everything around a replicated log - services become subscribers, not coordinated resources
- **Causality violations** happen when reads go to stale replicas, fixed by tracking sequence numbers client-side
- **2PC** blocks when coordinators fail mid-transaction, creating operational headaches

## Implementation Rationale

If I were to code this, I'd structure it around:

1. A **linearizable ticket counter** using something like etcd for the initial availability check
2. A **command log** using a consensus algorithm (Raft) for all state changes
3. **Service subscribers** that process the log in the same order to maintain consistency
4. **Client-side sequence tracking** to ensure read-your-writes consistency

The key would be making the log the source of truth, not trying to keep multiple databases in sync.

## Questions for Further Exploration

One thing I'm still thinking about: in the total order broadcast approach, what's the right way to handle service-specific failures? Like if the payment service is down but the log keeps advancing - do we block the entire log or build a retry mechanism?

Also, is it practical to have the initial linearizable check AND the command log, or does that create two sources of truth that could get out of sync?
