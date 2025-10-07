# Conference Room Booking Kata

This is a code kata to demonstrate the **Write Skew** concurrency anomaly and how to prevent it in a database-driven application. The scenario involves a booking system with a business rule that a maximum of 3 rooms can be booked for any given time slot.

## The Problem

The core problem is that under **Snapshot Isolation** (or Repeatable Read), two concurrent transactions can both read the same state of the database (e.g., seeing 2 existing bookings), decide it's safe to proceed, and then insert new bookings. Since they insert different rows, there's no direct write conflict, and both commits succeed, violating the business invariant (e.g., ending up with 4 bookings).

This is a classic **Write Skew** problem, caused by a **Phantom Read** where a transaction's write (an INSERT) affects the result of another transaction's previous read (a COUNT).

## My Solution & Rationale

I explored two main solutions to prevent this:

1.  **Pessimistic Locking with `SELECT FOR UPDATE`**: The idea is to lock the existing rows for the time slot during the initial read. This forces the second transaction to wait, ensuring it sees the updated count after the first one commits. The rationale is to turn the non-lockable phantom read into a lock on concrete existing rows, serializing the transactions.

2.  **Materializing Conflicts**: The idea is to create a dedicated `time_slots` table that holds a counter for each slot. All bookings must update this central counter, creating a direct write-write conflict on a single row that the database can automatically serialize. The rationale is to materialize the conflict into a tangible resource that can be locked, eliminating the phantom.

## Key Lessons Learned

- **Snapshot Isolation is not a silver bullet.** It prevents dirty reads and non-repeatable reads, but it doesn't solve all concurrency problems. Write Skew is a dangerous blind spot.
- **Phantoms are not just about reads.** The real danger is when a phantom (a new row matching a previous search) allows a business rule to be violated.
- The gap between a `SELECT` (checking a condition) and an `INSERT`/`UPDATE` (modifying data) is a critical section that needs protection, either through explicit locking or by designing the data model to force conflicts.

## Questions for Further Exploration

For the `SELECT FOR UPDATE` solution, a question arises: what's the performance impact under very high load? If 20 teams try to book the last available slot at once, 19 will be blocked and might time out. Is there a smarter, more granular locking strategy, or is an optimistic approach with retries better in that case?

For the materialized conflict solution, how do we keep the `bookings` table and the `time_slots` counter perfectly in sync? Should this be done with application-level transactions, or is a database trigger a more robust approach to avoid bugs?
