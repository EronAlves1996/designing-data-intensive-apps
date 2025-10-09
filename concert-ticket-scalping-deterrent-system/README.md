## Serializability Kata: Concert Ticket System

This is a code kata to practice the concepts of Transaction Serializability from "Designing Data-Intensive Applications". The goal is to think through how different isolation mechanisms handle concurrent seat reservations without writing actual code.

### The Problem

How do you prevent two users (or bots) from buying the same concert ticket at the same time? The business rule is simple: a seat can only be reserved once. But under high concurrency, this is hard. This kata explores three ways to guarantee serializability.

### What I Learned

- **Actual Serial Execution** is simple and safe but kills performance. It's a non-starter for a ticketing system that needs high throughput.
- **Two-Phase Locking (2PL)** feels like the "correct" way but is prone to deadlocks. You get safety but pay with complexity and potential performance hits from locking. My initial thought about lock promotion was close, but the real issue is just holding shared locks that block exclusive ones.
- **Serializable Snapshot Isolation (SSI)** is the most modern approach. It's optimistic: it lets transactions run freely and only checks for conflicts at commit time. It's faster when conflicts are rare, which fits this scenario—most users pick different seats.

My solution showed I understood the outcomes but needed more precision on the mechanics, like how SSI actually tracks the read/write sets.

### Rationale for Implementation

If I were to code this, the logic would be simple, but the database configuration would be everything. The `reserve_seats` transaction would just be those three steps. The magic is in telling the database to use `SERIALIZABLE` isolation.

For SSI, I'd need a retry mechanism. If the database aborts a transaction due to a serialization failure, the application should catch that error and transparently retry the reservation. The user might see a slight delay but never a double-booking.

### Questions for Further Exploration

In a real system, is it better to lock on a per-seat basis or use a coarser lock? What's the performance hit of true serializability vs. a practical workaround like using a Redis-based lock for the entire concert for a few seconds? How many retries are acceptable in SSI before telling the user "try again"?
