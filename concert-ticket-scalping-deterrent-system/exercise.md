### Kata: The Concert Ticket Scalping Deterrent System

**Business Context:**

You are a backend engineer at "BeatDrop," a popular ticketing platform. Your system is plagued by scalper bots that exploit race conditions during high-demand sales. Multiple bots try to grab the same block of seats simultaneously, leading to overselling or lost sales.

Your team is implementing a new, robust "Seat Reservation Engine" to guarantee that each seat can only be sold once, no matter the traffic volume. The product manager has mandated that the transaction isolation for the reservation process must be **Serializable**.

**Your Mission:**

You are tasked with designing and reasoning about the core transaction for reserving a set of seats. You will not write code, but you will design the transaction logic and analyze how it would behave under the three serializability schemes.

**The Schema & Pre-condition:**

- A `seats` table with columns: `concert_id`, `seat_id`, `status` ('available', 'reserved', 'sold').
- A `reservations` table with columns: `reservation_id`, `user_id`, `concert_id`, `seat_ids` (an array), `status` ('pending', 'confirmed').
- Initial State: For a popular concert `C123`, seats `S1`, `S2`, and `S3` are all `'available'`.

**The Core Transaction Logic (`reserve_seats`):**

This transaction is called when a user (or a bot) requests to reserve specific seats.

1.  **Check Availability:** Verify that all requested seats (e.g., `[S1, S2, S3]`) have a `status` of `'available'` for concert `C123`.
2.  **Create Reservation:** If all are available, insert a new record into the `reservations` table with `status = 'pending'`.
3.  **Update Seats:** Update the `status` of all the requested seats from `'available'` to `'reserved'`.

---

**Kata Tasks (30 minutes):**

**1. Analyze Under Actual Serial Execution (5 mins):**
Two transactions, `T-A` (reserving `[S1, S2]`) and `T-B` (reserving `[S2, S3]`), arrive simultaneously. Describe the final state of the database if the system uses Actual Serial Execution. What is the primary drawback for the business in this scenario?

**2. Analyze Under Two-Phase Locking (2PL) (10 mins):**
The same two transactions, `T-A` and `T-B`, arrive. Describe the sequence of events. What locks would each transaction request and when? What is a likely outcome (e.g., one succeeds, one fails, or a deadlock)? If a deadlock occurs, how would the database resolve it, and what should happen from the user's perspective?

**3. Analyze Under Serializable Snapshot Isolation (SSI) (10 mins):**
Again, `T-A` and `T-B` arrive. Under SSI, both transactions might initially seem to succeed in their own snapshots. Walk through the steps. At what point would SSI detect a conflict? Which transaction would be committed and which would be aborted? What is the key piece of information SSI uses to make this decision that isn't a lock?

**4. Design Recommendation & Justification (5 mins):**
Based on your analysis, which serializability implementation (2PL or SSI) would you recommend for the "Seat Reservation Engine"? Justify your choice considering the business need for both correctness (no double-booking) and high performance under extreme load.

This kata forces you to think through the mechanics and consequences of each serializability method in a high-stakes, real-world scenario, moving beyond abstract definitions into practical system design.
