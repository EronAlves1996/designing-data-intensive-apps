### Kata: The "ConcertFrenzy" Ticket System Overhaul

**Business Context:**

You are a backend engineer at "ConcertFrenzy," a platform notorious for high-demand ticket sales. The current system for the final "hold-and-checkout" step is a mess. When a user clicks "Buy Now," the system places a 5-minute hold on the tickets and starts the payment process. However, due to eventual consistency across data centers, users sometimes:

1.  See a "Hold Failed. No Tickets Available" message, only to receive a confirmation email and credit card charge minutes later.
2.  See a "Purchase Successful" screen, but then receive a cancellation email because the tickets were double-sold.

The CEO has declared this a "company-ending bug." Your team is tasked with designing the core logic for a new, consistent ticket hold service.

**Your Mission:**

Design and reason about the core decision flow for the `placeHold(userId, eventId, quantity)` function. You will not write code, but you will draft a technical narrative and make critical architectural choices based on the concepts from Chapter 9.

**Part 1: The Linearizability Dilemma (10 mins)**

- **Scenario A:** A user, Alice, clicks "Buy Now" for the last 2 tickets from a node in Europe. At the exact same time, user Bob clicks for the last 2 tickets from a node in North America.
- **Question:** Describe what can happen in the current, non-linearizable system. How could both Alice and Bob be led to believe they have secured the tickets?
  > **A:** As the system actually is not linearizable, actually the data storage acts as splitted, and we have splitted brains, and the writes occur concurrently. We have a bunch of holds occuring at the same time on various nodes, which we loose information and causality guarantees. Alice and Bob actually think they have secured the tickets, but the final result is actually undefined: may they have secured, may they not have
- **Decision:** Your team proposes using a linearizable datastore (like etcd or ZooKeeper) to manage the ticket inventory counter. Explain _why_ linearizability is the correct guarantee here. What is the specific "illusion" it creates that solves the double-selling problem?
  > **A:** Linearizability is make the system act like it have a single data storage. That way, the ticket inventory is supposed to be only one, and double selling is now not a problem, because, if we have only a single ticket inventory, and not two, now, every time someone hold a ticket, actually he holds the ticket, and he's not fooled about splitted data storage and brains problems
- **Trade-off:** The product manager is worried about performance. What is the primary **Cost of Linearizability** you must explain to them? (Hint: Think about what happens during a network partition).
  > **A:** The primary cost here is, if the network is failed apart, then we should choose between availability or consistency. If we choose availability, then we need to cope with the possible consequences of database failovers and leader election, and possible data loss. If we choose consistency, then we have to cope with possible angry users that cannot use our system because it is unavailable.

**Part 2: Causality and User Confusion (10 mins)**

- **Scenario B:** After placing a hold, the user goes to their "My Holds" page. The system shows "Hold Active." They then click "Refresh" and for a moment, the page shows "No Active Holds," before switching back to "Hold Active." This is confusing but not financially damaging.
- **Analysis:** This is a violation of causality. The user's action of _seeing the hold_ causally precedes their action of _refreshing the page_. Why is it possible for them to see a state that appears to be from _before_ the hold was placed?
  > **A:** Probably because, first, he hits some replica where the hold is already hold by him. Then, when he refreshes the page, he hits another replica where the hold is not hold by him _yet_. It's like the user is travelling backwards in time, but it's only a monotonic read and/or read-your-writes guarantee violation
- **Solution:** Your system uses a leader-based replication. How can you use **Sequence Number Ordering** to prevent this read-your-writes inconsistency? Describe what the client (web browser) and server must do to guarantee the user always sees their own updates.
  > **A:** When the user make a write, the system will return in the response, on a header, the actual logic timestamp to indicate what's the ordering write timestamp where this user should always see the updated data. On the browser, for every request, it sends this logic timestamp in a request header, where the server gonna see this logic timestamp and direct it to the adequately updated database, where this logic timestamp is equal or higher in any of them. We can reinforce the monotonic read guarantee by routing the user to a single database always, by using the user id as the criteria for this.

**Part 3: Achieving Consensus on a Hold (10 mins)**

- **The Final Hurdle:** The `placeHold` operation is not just about decrementing a counter. It must also create a hold record in a database and emit a `HoldPlaced` event to a message queue for the payment service to consume. All of this must be atomic: either all steps happen, or none do (e.g., if the database is unavailable, the ticket counter should not be decremented).
- **Proposal 1:** Use a **Two-Phase Commit (2PC)** protocol between the ticket inventory service (the linearizable store), the holds database, and the message queue.
  - What is the role of the "coordinator" in this setup?
  - What is a major operational drawback of 2PC that might make your team hesitant to use it?
- **Proposal 2:** Use a **Total Order Broadcast** mechanism, implemented via a consensus algorithm like Raft.
  - How does this change the architecture? Instead of three separate resources, what becomes the single "source of truth" that all nodes agree on?
  - Explain how broadcasting a message like `[tx_id: 789, operation: placeHold, user: Alice, tickets: 2]` through a total order broadcast log ensures consistency across all services (inventory, database, queue).

This kata forces you to apply the abstract concepts of linearizability, causality, and consensus to a high-stakes, realistic business problem, moving from identifying the problem to evaluating different solution architectures.
